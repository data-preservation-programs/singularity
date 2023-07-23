package replication

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication/internal/proposal110"
	"github.com/data-preservation-programs/singularity/replication/internal/proposal120"
	"github.com/filecoin-project/go-address"
	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/jellydator/ttlcache/v3"
	"github.com/jsign/go-filsigner/wallet"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"github.com/ybbus/jsonrpc/v3"
	"golang.org/x/exp/slices"
)

const (
	StorageProposalV120 = "/fil/storage/mk/1.2.0"
	StorageProposalV111 = "/fil/storage/mk/1.1.1"
)

func TimeToEpoch(t time.Time) abi.ChainEpoch {
	return abi.ChainEpoch((t.Unix() - 1598306400) / 30)
}

func EpochToTime(epoch abi.ChainEpoch) time.Time {
	return time.Unix(int64(epoch*30+1598306400), 0)
}

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

func (d DealMakerImpl) getProviderInfo(ctx context.Context, provider string) (*MinerInfo, error) {
	item := d.minerInfoCache.Get(provider)
	if item != nil && !item.IsExpired() {
		return item.Value(), nil
	}

	minerInfo := new(MinerInfo)
	err := d.lotusClient.CallFor(ctx, minerInfo, "Filecoin.StateMinerInfo", provider, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get miner info")
	}

	minerInfo.Multiaddrs = make([]multiaddr.Multiaddr, len(minerInfo.MultiaddrsBase64Encoded))
	for i, addr := range minerInfo.MultiaddrsBase64Encoded {
		decoded, err := base64.StdEncoding.DecodeString(addr)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode multiaddr")
		}
		minerInfo.Multiaddrs[i], err = multiaddr.NewMultiaddrBytes(decoded)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create multiaddr")
		}
	}
	minerInfo.PeerID, err = peer.Decode(minerInfo.PeerIDEncoded)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode peer id")
	}

	d.minerInfoCache.Set(provider, minerInfo, ttlcache.DefaultTTL)
	return minerInfo, nil
}

func (d DealMakerImpl) getProtocols(ctx context.Context, minerInfo peer.AddrInfo) ([]protocol.ID, error) {
	item := d.protocolsCache.Get(minerInfo.ID)
	if item != nil && !item.IsExpired() {
		return item.Value(), nil
	}

	d.host.Peerstore().AddAddrs(minerInfo.ID, minerInfo.Addrs, peerstore.TempAddrTTL)
	if err := d.host.Connect(ctx, minerInfo); err != nil {
		return nil, errors.Wrap(err, "failed to connect to miner")
	}

	protocols, err := d.host.Peerstore().GetProtocols(minerInfo.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get protocols")
	}

	d.protocolsCache.Set(minerInfo.ID, protocols, ttlcache.DefaultTTL)
	return protocols, nil
}

func (d DealMakerImpl) getMinCollateral(ctx context.Context, pieceSize int64, verified bool) (big.Int, error) {
	item := d.collateralCache.Get(fmt.Sprintf("%d-%t", pieceSize, verified))
	if item != nil && !item.IsExpired() {
		return item.Value(), nil
	}

	bound := new(DealProviderCollateralBound)
	err := d.lotusClient.CallFor(ctx, bound, "Filecoin.StateDealProviderCollateralBounds", pieceSize, verified, nil)
	if err != nil {
		return big.Int{}, errors.Wrap(err, "failed to get deal provider collateral bounds")
	}

	value, err := big.FromString(bound.Min)
	if err != nil {
		return big.Int{}, errors.Wrap(err, "failed to parse min collateral")
	}
	d.collateralCache.Set(fmt.Sprintf("%d-%t", pieceSize, verified), value, ttlcache.DefaultTTL)
	return value, nil
}

func (d DealMakerImpl) makeDeal120(
	ctx context.Context,
	deal proposal110.ClientDealProposal,
	dealID uuid.UUID,
	dealConfig DealConfig,
	fileSize int64,
	rootCID cid.Cid,
	minerInfo peer.AddrInfo) (*proposal120.DealResponse, error) {
	transfer := proposal120.Transfer{
		Size: uint64(fileSize),
	}
	url := strings.Replace(dealConfig.URLTemplate, "{PIECE_CID}", deal.Proposal.PieceCID.String(), 1)
	isOnline := url != ""
	if isOnline {
		transferParams := &proposal120.HttpRequest{URL: url}
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
			return nil, errors.Wrap(err, "failed to serialize transfer params")
		}
		transfer.Type = "http"
		transfer.Params = paramsBytes
	}

	dealParams := proposal120.DealParams{
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
		return nil, errors.Wrap(err, "failed to connect to miner")
	}

	stream, err := d.host.NewStream(ctx, minerInfo.ID, StorageProposalV120)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open stream")
	}
	defer stream.Close()
	if deadline, ok := ctx.Deadline(); ok {
		err := stream.SetDeadline(deadline)
		if err != nil {
			return nil, errors.Wrap(err, "failed to set stream deadline")
		}
		//nolint:errcheck
		defer stream.SetDeadline(time.Time{})
	}

	err = cborutil.WriteCborRPC(stream, &dealParams)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write deal params")
	}

	var resp proposal120.DealResponse
	err = cborutil.ReadCborRPC(stream, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read deal response")
	}

	return &resp, nil
}

func (d DealMakerImpl) makeDeal111(
	ctx context.Context,
	deal proposal110.ClientDealProposal,
	dealConfig DealConfig,
	rootCID cid.Cid,
	minerInfo peer.AddrInfo) (*proposal110.SignedResponse, error) {
	proposal := proposal110.Proposal{
		FastRetrieval: dealConfig.KeepUnsealed,
		DealProposal:  &deal,
		Piece: &proposal110.DataRef{
			TransferType: proposal110.TTManual,
			Root:         rootCID,
			PieceCid:     &deal.Proposal.PieceCID,
			PieceSize:    deal.Proposal.PieceSize.Unpadded(),
		},
	}

	d.host.Peerstore().AddAddrs(minerInfo.ID, minerInfo.Addrs, peerstore.TempAddrTTL)
	if err := d.host.Connect(ctx, minerInfo); err != nil {
		return nil, errors.Wrap(err, "failed to connect to miner")
	}

	stream, err := d.host.NewStream(ctx, minerInfo.ID, StorageProposalV111)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open stream")
	}
	defer stream.Close()
	if deadline, ok := ctx.Deadline(); ok {
		err = stream.SetDeadline(deadline)
		if err != nil {
			return nil, errors.Wrap(err, "failed to set stream deadline")
		}
		//nolint:errcheck
		defer stream.SetDeadline(time.Time{})
	}

	var resp proposal110.SignedResponse
	err = cborutil.WriteCborRPC(stream, &proposal)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write deal params")
	}

	err = cborutil.ReadCborRPC(stream, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read deal response")
	}

	return &resp, nil
}

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

func (d DealConfig) GetPrice(pieceSize int64, duration time.Duration) big.Int {
	gb := float64(pieceSize) / 1e9
	epoch := duration.Minutes() * 2
	p1 := big.NewIntUnsigned(uint64(d.PricePerGBEpoch * 1e18 * gb * epoch))
	p2 := big.NewIntUnsigned(uint64(d.PricePerGB * 1e18 * gb))
	p3 := big.NewIntUnsigned(uint64(d.PricePerDeal * 1e18))
	return big.Max(big.Max(p1, p2), p3)
}

func (d DealMakerImpl) MakeDeal(ctx context.Context, walletObj model.Wallet,
	car model.Car, dealConfig DealConfig) (*model.Deal, error) {
	now := time.Now().UTC()
	addr, err := address.NewFromString(walletObj.Address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse wallet address")
	}

	pvd, err := address.NewFromString(dealConfig.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse provider address")
	}

	label, err := proposal110.NewLabelFromString(cid.Cid(car.RootCID).String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse label")
	}

	ctx, cancel := context.WithTimeout(ctx, d.requestTimeout)
	defer cancel()

	minerInfo, err := d.getProviderInfo(ctx, dealConfig.Provider)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get provider info")
	}

	d.host.Peerstore().AddAddrs(minerInfo.PeerID, minerInfo.Multiaddrs, peerstore.TempAddrTTL)
	addrInfo := peer.AddrInfo{ID: minerInfo.PeerID, Addrs: minerInfo.Multiaddrs}

	protocols, err := d.getProtocols(ctx, addrInfo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get supported protocol")
	}

	startEpoch := TimeToEpoch(now.Add(dealConfig.StartDelay))
	endEpoch := TimeToEpoch(now.Add(dealConfig.StartDelay + dealConfig.Duration))
	price := dealConfig.GetPrice(car.PieceSize, dealConfig.Duration)
	verified := dealConfig.Verified
	pieceCID := cid.Cid(car.PieceCID)
	pieceSize := abi.PaddedPieceSize(car.PieceSize)
	collateral, err := d.getMinCollateral(ctx, car.PieceSize, verified)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get min collateral")
	}
	proposal := proposal110.DealProposal{
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
		return nil, errors.Wrap(err, "failed to serialize deal proposal")
	}

	signature, err := wallet.WalletSign(walletObj.PrivateKey, proposalBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign deal proposal")
	}

	deal := proposal110.ClientDealProposal{
		Proposal:        proposal,
		ClientSignature: *signature,
	}

	dealModel := &model.Deal{
		DatasetID:  &car.DatasetID,
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
		resp, err := d.makeDeal120(ctx, deal, dealID, dealConfig, car.FileSize, cid.Cid(car.RootCID), addrInfo)
		if err != nil {
			return nil, errors.Wrap(err, "failed to make deal")
		}
		if resp.Accepted {
			dealModel.ProposalID = dealID.String()
			return dealModel, nil
		}

		return nil, errors.Errorf("deal rejected: %s", resp.Message)
	} else if slices.Contains(protocols, StorageProposalV111) {
		resp, err := d.makeDeal111(ctx, deal, dealConfig, cid.Cid(car.RootCID), addrInfo)
		if err != nil {
			return nil, errors.Wrap(err, "failed to make deal")
		}

		dealModel.ProposalID = resp.Response.Proposal.String()
		return dealModel, nil
	}

	return nil, errors.New("no supported protocol found")
}
