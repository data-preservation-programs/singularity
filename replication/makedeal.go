package replication

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/filecoin-project/go-address"
	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket/network"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/filecoin-shipyard/boostly"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/jellydator/ttlcache/v3"
	"github.com/jsign/go-filsigner/wallet"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
	"github.com/ybbus/jsonrpc/v3"
	"golang.org/x/exp/slices"
)

const (
	StorageProposalV120 = "/fil/storage/mk/1.2.0"
	StorageProposalV111 = "/fil/storage/mk/1.1.1"
)

var ErrNoSupportedProtocols = errors.New("no supported protocols")

// nolint: tagliatelle
type MinerInfo struct {
	PeerIDEncoded           string `json:"PeerID"`
	PeerID                  peer.ID
	MultiaddrsBase64Encoded []string `json:"Multiaddrs"`
	Multiaddrs              []multiaddr.Multiaddr
}

type DealProviderCollateralBound struct {
	Min string
	Max string
}

type DealMaker interface {
	MakeDeal(ctx context.Context, walletObj model.Wallet, car model.Car, dealConfig DealConfig) (*model.Deal, error)
}

// DealMakerImpl is an implementation of a deal-making component for a Filecoin-like network.
// It is responsible for negotiating storage deals between a client and a provider (miner).
//
// Fields:
//   - lotusClient: A JSON-RPC client to interact with a Filecoin Lotus node.
//     This client is used to query the blockchain and interact with the Filecoin network.
//   - host: A libp2p host used for networking operations, such as connecting to miners
//     and sending/receiving deal proposals.
//   - requestTimeout: The maximum duration that the DealMakerImpl should wait for a response
//     from the miner when making a deal. After this timeout, the deal request
//     is considered as failed.
//   - minerInfoCache: A cache that stores information about miners. The key is the miner's
//     address, and the value is a struct containing detailed information about the miner.
//   - protocolsCache: A cache that stores the supported protocols of different peers in the network.
//     The key is the peer ID of the node, and the value is a slice of supported protocol IDs.
//   - collateralCache: A cache that stores the minimum required collateral for making a deal.
//     The key is a string representing some parameters of the deal, and the value is
//     the minimum required collateral in attoFIL (or similar units).
type DealMakerImpl struct {
	lotusClient     jsonrpc.RPCClient
	host            host.Host
	requestTimeout  time.Duration
	minerInfoCache  *ttlcache.Cache[string, *MinerInfo]
	protocolsCache  *ttlcache.Cache[peer.ID, []protocol.ID]
	collateralCache *ttlcache.Cache[string, big.Int]
}

func (d DealMakerImpl) Close() error {
	if d.host != nil {
		return d.host.Close()
	}

	return nil
}

func NewDealMaker(
	lotusClient jsonrpc.RPCClient,
	libp2p host.Host,
	cacheTTL time.Duration,
	requestTimeout time.Duration) DealMakerImpl {
	minerInfoCache := ttlcache.New[string, *MinerInfo](
		ttlcache.WithTTL[string, *MinerInfo](cacheTTL),
		ttlcache.WithDisableTouchOnHit[string, *MinerInfo]())
	protocolsCache := ttlcache.New[peer.ID, []protocol.ID](
		ttlcache.WithTTL[peer.ID, []protocol.ID](cacheTTL),
		ttlcache.WithDisableTouchOnHit[peer.ID, []protocol.ID]())
	collateralCache := ttlcache.New[string, big.Int](
		ttlcache.WithTTL[string, big.Int](cacheTTL),
		ttlcache.WithDisableTouchOnHit[string, big.Int]())

	return DealMakerImpl{
		lotusClient:     lotusClient,
		requestTimeout:  requestTimeout,
		host:            libp2p,
		minerInfoCache:  minerInfoCache,
		protocolsCache:  protocolsCache,
		collateralCache: collateralCache,
	}
}

// GetProviderInfo retrieves information about a given Filecoin provider (miner).
//
// This function checks a cache for the requested miner's information. If the
// miner information is found in the cache and is not expired, it returns the
// cached information. Otherwise, it queries the information using the Lotus
// client, decodes multiaddresses and peer ID, and stores this newly fetched
// information in the cache with a default TTL (Time To Live).
//
// Parameters:
//   - ctx context.Context: The context to use for the Lotus client call, allowing
//     for cancellation and timeouts.
//   - provider string: The address of the provider (miner) whose information is to
//     be retrieved.
//
// Returns:
//   - *MinerInfo: A pointer to the MinerInfo structure that contains various
//     details about the provider (miner), including its multiaddresses
//     and peer ID.
//   - error: An error that will be returned if any issues were encountered while
//     trying to retrieve the provider's information. This could be due to
//     errors in communicating with the Lotus client, decoding base64 encoded
//     multiaddresses, creating new multiaddresses, or decoding the peer ID.
func (d DealMakerImpl) GetProviderInfo(ctx context.Context, provider string) (*MinerInfo, error) {
	file := d.minerInfoCache.Get(provider)
	if file != nil && !file.IsExpired() {
		return file.Value(), nil
	}

	logger.Debugw("getting miner info", "miner", provider)
	minerInfo := new(MinerInfo)
	err := d.lotusClient.CallFor(ctx, minerInfo, "Filecoin.StateMinerInfo", provider, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get miner info, miner: %s", provider)
	}

	logger.Debug("got miner info", "miner", provider, "minerInfo", minerInfo)
	minerInfo.Multiaddrs = make([]multiaddr.Multiaddr, len(minerInfo.MultiaddrsBase64Encoded))
	for i, addr := range minerInfo.MultiaddrsBase64Encoded {
		decoded, err := base64.StdEncoding.DecodeString(addr)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to decode multiaddr %s", addr)
		}
		minerInfo.Multiaddrs[i], err = multiaddr.NewMultiaddrBytes(decoded)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse multiaddr %s", addr)
		}
	}
	minerInfo.PeerID, err = peer.Decode(minerInfo.PeerIDEncoded)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decode peer id %s", minerInfo.PeerIDEncoded)
	}

	d.minerInfoCache.Set(provider, minerInfo, ttlcache.DefaultTTL)
	return minerInfo, nil
}

// GetProtocols retrieves the supported protocols of a given Filecoin miner.
//
// This function checks a cache for the miner's supported protocols. If the
// protocols for the given miner are found in the cache and are not expired,
// it returns the cached protocols. Otherwise, it performs the following steps:
// 1. Adds the miner's multiaddresses to the peerstore with a temporary TTL (Time To Live).
// 2. Attempts to establish a network connection with the miner using the host's Connect method.
// 3. Queries the host's peerstore for the supported protocols of the miner.
// 4. Stores the newly fetched protocols in the cache with a default TTL.
//
// Parameters:
//   - ctx context.Context: The context to use when connecting to the miner, allowing
//     for cancellation and timeouts.
//   - minerInfo peer.AddrInfo: A structure containing the peer ID and multiaddresses of
//     the miner whose supported protocols are to be retrieved.
//
// Returns:
//   - []protocol.ID: A slice of protocol IDs representing the protocols supported
//     by the miner.
//   - error: An error that will be returned if any issues were encountered while
//     trying to retrieve the miner's supported protocols. This could be due
//     to errors in connecting to the miner or fetching protocols from the peerstore.
func (d DealMakerImpl) GetProtocols(ctx context.Context, minerInfo peer.AddrInfo) ([]protocol.ID, error) {
	file := d.protocolsCache.Get(minerInfo.ID)
	if file != nil && !file.IsExpired() {
		return file.Value(), nil
	}

	logger.Debugw("getting protocols", "miner", minerInfo.ID)
	d.host.Peerstore().AddAddrs(minerInfo.ID, minerInfo.Addrs, peerstore.TempAddrTTL)
	if err := d.host.Connect(ctx, minerInfo); err != nil {
		return nil, errors.Wrapf(err, "failed to connect to miner %s - %v", minerInfo.ID, minerInfo.Addrs)
	}

	protocols, err := d.host.Peerstore().GetProtocols(minerInfo.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get protocols from peer %s", minerInfo.ID)
	}

	logger.Debugw("got protocols", "miner", minerInfo.ID, "protocols", protocols)
	d.protocolsCache.Set(minerInfo.ID, protocols, ttlcache.DefaultTTL)
	return protocols, nil
}

// GetMinCollateral retrieves the minimum collateral that a Filecoin miner needs
// to put up in order to make a deal, considering the given piece size and
// whether the deal is verified.
//
// This function checks a cache for the minimum collateral based on the piece size
// and verified flag. If a valid cached value is found, it returns this value.
// Otherwise, it performs the following steps:
//  1. Makes a call to the Lotus client (Filecoin node) to fetch the deal provider
//     collateral bounds for the specified piece size and verified flag.
//  2. Parses the minimum bound from the response and converts it into a big.Int.
//  3. Stores the newly fetched minimum collateral in the cache with a default TTL.
//
// Parameters:
//   - ctx context.Context: The context for the function, allowing for cancellation
//     and timeouts.
//   - pieceSize int64: The size of the piece (in bytes) for which to retrieve the
//     minimum collateral.
//   - verified bool: A flag indicating whether the deal is verified.
func (d DealMakerImpl) GetMinCollateral(ctx context.Context, pieceSize int64, verified bool) (big.Int, error) {
	file := d.collateralCache.Get(fmt.Sprintf("%d-%t", pieceSize, verified))
	if file != nil && !file.IsExpired() {
		return file.Value(), nil
	}

	logger.Debugw("getting deal provider collateral bounds", "pieceSize", pieceSize, "verified", verified)
	bound := new(DealProviderCollateralBound)
	err := d.lotusClient.CallFor(ctx, bound, "Filecoin.StateDealProviderCollateralBounds", pieceSize, verified, nil)
	if err != nil {
		return big.Int{}, errors.Wrapf(err, "failed to get deal provider collateral bounds with pieceSize %d and verified %s", pieceSize, verified)
	}

	logger.Debugw("got deal provider collateral bounds", "pieceSize", pieceSize, "verified", verified, "bound", bound)
	value, err := big.FromString(bound.Min)
	if err != nil {
		return big.Int{}, errors.Wrapf(err, "failed to parse min collateral %s", bound.Min)
	}
	d.collateralCache.Set(fmt.Sprintf("%d-%t", pieceSize, verified), value, ttlcache.DefaultTTL)
	return value, nil
}

// MakeDeal120 attempts to make a storage deal with a Filecoin miner, following
// the version 1.2.0 deal making protocol. It constructs and sends the deal proposal
// and associated data transfer instructions to the miner, and waits for a response.
//
// The function accepts various parameters related to the deal, including deal proposal,
// deal configuration, file size, root CID of the data, and miner information. It also
// generates a transfer object based on the input parameters, which includes instructions
// for how the miner should retrieve the data.
//
// Parameters:
//   - ctx context.Context: The context for the function, allowing for cancellation
//     and timeouts.
//   - deal market.ClientDealProposal: The storage deal proposal to be sent to the miner.
//   - dealID uuid.UUID: The unique identifier for this deal.
//   - dealConfig DealConfig: Configuration options for this deal.
//   - fileSize int64: The size of the file (in bytes) being stored in this deal.
//   - rootCID cid.Cid: The root CID of the data to be stored in this deal.
//   - minerInfo peer.AddrInfo: Network information for the miner with whom the deal is made.
//
// Returns:
//   - *boostly.DealResponse: A pointer to the response from the miner regarding the
//     proposed deal, including whether it was accepted and
//     any associated messages.
//   - error: An error that will be returned if any issues were encountered while making
//     the deal. This could be due to communication errors with the miner,
//     serialization/deserialization issues, or stream handling errors.
func (d DealMakerImpl) MakeDeal120(
	ctx context.Context,
	deal market.ClientDealProposal,
	dealID uuid.UUID,
	dealConfig DealConfig,
	fileSize int64,
	rootCID cid.Cid,
	minerInfo peer.AddrInfo) (*boostly.DealResponse, error) {
	logger.Debugw("making deal 120", "dealID", dealID, "deal", deal,
		"dealConfig", dealConfig, "fileSize", fileSize, "rootCID", rootCID.String(), "minerInfo", minerInfo)
	transfer := boostly.Transfer{
		Size: uint64(fileSize),
	}
	url := strings.Replace(dealConfig.URLTemplate, "{PIECE_CID}", deal.Proposal.PieceCID.String(), 1)
	isOnline := url != ""
	if isOnline {
		transferParams := &boostly.HttpRequest{URL: url}
		if len(dealConfig.HTTPHeaders) > 0 {
			transferParams.Headers = make(map[string]string)
			for _, header := range dealConfig.HTTPHeaders {
				sp := strings.Split(header, "=")
				if len(sp) != 2 {
					return nil, errors.Errorf("invalid http header %s", header)
				}
				transferParams.Headers[sp[0]] = sp[1]
			}
		}
		paramsBytes, err := json.Marshal(transferParams)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to serialize transfer params %v", transferParams)
		}
		transfer.Type = "http"
		transfer.Params = paramsBytes
	}

	dealParams := boostly.DealParams{
		DealUUID:           dealID,
		ClientDealProposal: deal,
		DealDataRoot:       rootCID,
		IsOffline:          !isOnline,
		Transfer:           transfer,
		RemoveUnsealedCopy: !dealConfig.KeepUnsealed,
		SkipIPNIAnnounce:   !dealConfig.AnnounceToIPNI,
	}

	d.host.Peerstore().AddAddrs(minerInfo.ID, minerInfo.Addrs, peerstore.TempAddrTTL)
	if err := d.host.Connect(ctx, minerInfo); err != nil {
		return nil, errors.Wrapf(err, "failed to connect to miner %s - %s %v", dealConfig.Provider, minerInfo.ID, minerInfo.Addrs)
	}

	stream, err := d.host.NewStream(ctx, minerInfo.ID, StorageProposalV120)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open stream with %s using %s", dealConfig.Provider, StorageProposalV120)
	}
	defer stream.Close()
	if deadline, ok := ctx.Deadline(); ok {
		err := stream.SetDeadline(deadline)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		//nolint:errcheck
		defer stream.SetDeadline(time.Time{})
	}

	logger.Debugw("sending deal params", "dealParams", dealParams)
	err = cborutil.WriteCborRPC(stream, &dealParams)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to write deal params %v", dealParams)
	}

	var resp boostly.DealResponse
	err = cborutil.ReadCborRPC(stream, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read deal response")
	}

	logger.Debugw("got deal response", "resp", resp)
	return &resp, nil
}

// MakeDeal111 attempts to make a storage deal with a Filecoin miner, following
// the version 1.1.1 deal making protocol. It constructs and sends the deal proposal
// to the miner, and waits for a response.
//
// The function accepts various parameters related to the deal, including the deal proposal,
// deal configuration, the root CID of the data, and miner network information. It also
// generates a proposal object based on the input parameters and sends this to the miner.
//
// Parameters:
//   - ctx context.Context: The context for the function, allowing for cancellation
//     and timeouts.
//   - deal market.ClientDealProposal: The storage deal proposal to be sent to the miner.
//   - dealConfig DealConfig: Configuration options for this deal.
//   - rootCID cid.Cid: The root CID of the data to be stored in this deal.
//   - minerInfo peer.AddrInfo: Network information for the miner with whom the deal is made.
//
// Returns:
//   - *network.SignedResponse: A pointer to the signed response from the miner regarding the
//     proposed deal, including whether it was accepted and
//     any associated messages.
//   - error: An error that will be returned if any issues were encountered while making
//     the deal. This could be due to communication errors with the miner,
//     serialization/deserialization issues, or stream handling errors.
func (d DealMakerImpl) MakeDeal111(
	ctx context.Context,
	deal market.ClientDealProposal,
	dealConfig DealConfig,
	rootCID cid.Cid,
	minerInfo peer.AddrInfo) (*network.SignedResponse, error) {
	logger.Debugw("making deal 111", "deal", deal, "dealConfig", dealConfig, "rootCID", rootCID.String(), "minerInfo", minerInfo)
	proposal := network.Proposal{
		FastRetrieval: dealConfig.KeepUnsealed,
		DealProposal:  &deal,
		Piece: &storagemarket.DataRef{
			TransferType: storagemarket.TTManual,
			Root:         rootCID,
			PieceCid:     &deal.Proposal.PieceCID,
			PieceSize:    deal.Proposal.PieceSize.Unpadded(),
		},
	}

	d.host.Peerstore().AddAddrs(minerInfo.ID, minerInfo.Addrs, peerstore.TempAddrTTL)
	if err := d.host.Connect(ctx, minerInfo); err != nil {
		return nil, errors.Wrapf(err, "failed to connect to miner %s - %s %v", dealConfig.Provider, minerInfo.ID, minerInfo.Addrs)
	}

	stream, err := d.host.NewStream(ctx, minerInfo.ID, StorageProposalV111)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open stream with %s using %s", dealConfig.Provider, StorageProposalV111)
	}
	defer stream.Close()
	if deadline, ok := ctx.Deadline(); ok {
		err = stream.SetDeadline(deadline)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		//nolint:errcheck
		defer stream.SetDeadline(time.Time{})
	}

	logger.Debugw("sending deal params", "proposal", proposal)
	var resp network.SignedResponse
	err = cborutil.WriteCborRPC(stream, &proposal)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to write deal params %v", proposal)
	}

	err = cborutil.ReadCborRPC(stream, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read deal response")
	}

	logger.Debugw("got deal response", "resp", resp)
	return &resp, nil
}

// DealConfig represents the configuration parameters for a storage deal in a Filecoin-like network.
//
// Fields:
// - Provider: The Filecoin address of the storage provider (miner) with whom the deal is being made.
// - StartDelay: The duration before the deal is expected to start. It gives the miner time to prepare for the deal.
// - Duration: The duration for which the deal will be in effect, i.e., the time period the data is stored.
// - Verified: A flag indicating whether the deal is for verified storage.
// - HTTPHeaders: Custom HTTP headers to be used when fetching data from a URL.
// - URLTemplate: A URL template string which can be used to fetch data for storage.
// - KeepUnsealed: A flag indicating whether the miner should keep the data unsealed (quick retrieval) or not.
// - AnnounceToIPNI: A flag indicating whether the deal should be announced to the InterPlanetary Name Identifier (IPNI).
// - PricePerDeal: The upfront cost of the deal, independent of the amount of data being stored, in FIL (or a smaller unit like attoFIL).
// - PricePerGB: The upfront cost per GB of data being stored, in FIL (or a smaller unit like attoFIL).
// - PricePerGBEpoch: The cost per GB per epoch (e.g., per block or per minute), in FIL (or a smaller unit like attoFIL).
type DealConfig struct {
	Provider        string
	StartDelay      time.Duration
	Duration        time.Duration
	Verified        bool
	HTTPHeaders     []string
	URLTemplate     string
	KeepUnsealed    bool
	AnnounceToIPNI  bool
	PricePerDeal    float64
	PricePerGB      float64
	PricePerGBEpoch float64
}

// GetPrice calculates the price of a deal based on the size of the piece being stored,
// the duration of the storage, and various price parameters that are part of the DealConfig.
//
// The method considers three potential price calculations:
//  1. Price based on the size of the piece (in GB), the duration of the deal (in epochs),
//     and a price rate per GB per epoch.
//  2. Price based on the size of the piece (in GB) and a flat price rate per GB.
//  3. A flat price per deal, independent of the size and duration.
//
// The method returns the maximum of these three calculated prices.
//
// Parameters:
// - pieceSize int64: The size of the piece to be stored, in bytes.
// - duration time.Duration: The duration for which the piece will be stored.
//
// Returns:
// - big.Int: The calculated price for the deal in attoFIL (1e-18 FIL).
func (d DealConfig) GetPrice(pieceSize int64, duration time.Duration) big.Int {
	gb := float64(pieceSize) / 1e9
	epoch := duration.Minutes() * 2
	p1 := big.NewIntUnsigned(uint64(d.PricePerGBEpoch * 1e18 * gb * epoch))
	p2 := big.NewIntUnsigned(uint64(d.PricePerGB * 1e18 * gb))
	p3 := big.NewIntUnsigned(uint64(d.PricePerDeal * 1e18))
	return big.Max(big.Max(p1, p2), p3)
}

// MakeDeal initiates a storage deal between a client and a provider in a Filecoin-like network.
//
// It constructs a deal proposal based on input parameters including a car file,
// a wallet, a deal configuration, and more. It connects to the provider (miner), negotiates the deal
// terms based on the protocols supported by the provider, signs the deal proposal using the client's private key,
// and then sends the proposal to the provider.
//
// Parameters:
// - ctx context.Context: The context to use for timeouts and cancellation.
// - walletObj model.Wallet: The client's wallet, containing the client's addresses and private key.
// - car model.Car: The car file that contains the data to be stored.
// - dealConfig DealConfig: The configuration for the deal, including price and duration.
//
// Returns:
//   - *model.Deal: The resulting deal object, if the deal was successful. Contains various details
//     about the deal, including its state and price.
//   - error: An error if the deal could not be completed, or nil if the deal was successful.
//
// Possible Errors:
// - Failed to parse wallet or provider addresses.
// - Failed to get provider info or supported protocol.
// - Failed to connect to the provider.
// - Failed to serialize the deal proposal.
// - Failed to sign the deal proposal.
// - Deal proposal rejected by the provider.
// - No supported protocol found between client and provider.
func (d DealMakerImpl) MakeDeal(ctx context.Context, walletObj model.Wallet,
	car model.Car, dealConfig DealConfig) (*model.Deal, error) {
	logger.Infow("making deal", "client", walletObj.ID, "pieceCID", car.PieceCID.String(), "provider", dealConfig.Provider)
	now := time.Now().UTC()
	addr, err := address.NewFromString(walletObj.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse wallet address %s", walletObj.Address)
	}

	pvd, err := address.NewFromString(dealConfig.Provider)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse provider address %s", dealConfig.Provider)
	}

	label, err := market.NewLabelFromString(cid.Cid(car.RootCID).String())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse label %s", cid.Cid(car.RootCID).String())
	}

	ctx, cancel := context.WithTimeout(ctx, d.requestTimeout)
	defer cancel()

	minerInfo, err := d.GetProviderInfo(ctx, dealConfig.Provider)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	d.host.Peerstore().AddAddrs(minerInfo.PeerID, minerInfo.Multiaddrs, peerstore.TempAddrTTL)
	addrInfo := peer.AddrInfo{ID: minerInfo.PeerID, Addrs: minerInfo.Multiaddrs}

	protocols, err := d.GetProtocols(ctx, addrInfo)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	startEpoch := epochutil.TimeToEpoch(now.Add(dealConfig.StartDelay))
	endEpoch := epochutil.TimeToEpoch(now.Add(dealConfig.StartDelay + dealConfig.Duration))
	price := dealConfig.GetPrice(car.PieceSize, dealConfig.Duration)
	verified := dealConfig.Verified
	pieceCID := cid.Cid(car.PieceCID)
	pieceSize := abi.PaddedPieceSize(car.PieceSize)
	collateral, err := d.GetMinCollateral(ctx, car.PieceSize, verified)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	proposal := market.DealProposal{
		PieceCID:             pieceCID,
		PieceSize:            pieceSize,
		VerifiedDeal:         verified,
		Client:               addr,
		Provider:             pvd,
		Label:                label,
		StartEpoch:           startEpoch,
		EndEpoch:             endEpoch,
		StoragePricePerEpoch: price,
		ProviderCollateral:   collateral,
	}
	proposalBytes, err := cborutil.Dump(&proposal)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to serialize deal proposal %s", proposal)
	}

	signature, err := wallet.WalletSign(walletObj.PrivateKey, proposalBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign deal proposal")
	}

	deal := market.ClientDealProposal{
		Proposal:        proposal,
		ClientSignature: *signature,
	}

	dealModel := &model.Deal{
		State:      model.DealProposed,
		ClientID:   walletObj.ID,
		Provider:   dealConfig.Provider,
		Label:      cid.Cid(car.RootCID).String(),
		PieceCID:   car.PieceCID,
		PieceSize:  car.PieceSize,
		StartEpoch: int32(startEpoch),
		EndEpoch:   int32(endEpoch),
		Price:      dealConfig.GetPrice(car.PieceSize, dealConfig.Duration).String(),
		Verified:   dealConfig.Verified,
	}
	if slices.Contains(protocols, StorageProposalV120) {
		dealID := uuid.New()
		resp, err := d.MakeDeal120(ctx, deal, dealID, dealConfig, car.FileSize, cid.Cid(car.RootCID), addrInfo)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if resp.Accepted {
			dealModel.ProposalID = dealID.String()
			return dealModel, nil
		}

		return nil, errors.Errorf("deal rejected: %s", resp.Message)
	} else if slices.Contains(protocols, StorageProposalV111) {
		resp, err := d.MakeDeal111(ctx, deal, dealConfig, cid.Cid(car.RootCID), addrInfo)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		dealModel.ProposalID = resp.Response.Proposal.String()
		return dealModel, nil
	}

	return nil, errors.Wrapf(ErrNoSupportedProtocols, "protocols: %v", protocols)
}
