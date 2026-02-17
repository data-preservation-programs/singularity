package pdptracker

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/go-synapse"
	"github.com/data-preservation-programs/go-synapse/constants"
	"github.com/data-preservation-programs/go-synapse/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-varint"
)

const pdpDefaultPageSize uint64 = 100

type activePiecesResult struct {
	Pieces  []contracts.CidsCid
	HasMore bool
}

type pdpVerifierAPI interface {
	GetNextDataSetId(opts *bind.CallOpts) (uint64, error)
	GetDataSetListener(opts *bind.CallOpts, setId *big.Int) (common.Address, error)
	GetDataSetStorageProvider(opts *bind.CallOpts, setId *big.Int) (common.Address, common.Address, error)
	DataSetLive(opts *bind.CallOpts, setId *big.Int) (bool, error)
	GetNextChallengeEpoch(opts *bind.CallOpts, setId *big.Int) (*big.Int, error)
	GetActivePieces(opts *bind.CallOpts, setId *big.Int, offset *big.Int, limit *big.Int) (activePiecesResult, error)
}

type pdpVerifierContract struct {
	contract *contracts.PDPVerifier
}

func (c pdpVerifierContract) GetNextDataSetId(opts *bind.CallOpts) (uint64, error) {
	return c.contract.GetNextDataSetId(opts)
}

func (c pdpVerifierContract) GetDataSetListener(opts *bind.CallOpts, setId *big.Int) (common.Address, error) {
	return c.contract.GetDataSetListener(opts, setId)
}

func (c pdpVerifierContract) GetDataSetStorageProvider(opts *bind.CallOpts, setId *big.Int) (common.Address, common.Address, error) {
	return c.contract.GetDataSetStorageProvider(opts, setId)
}

func (c pdpVerifierContract) DataSetLive(opts *bind.CallOpts, setId *big.Int) (bool, error) {
	return c.contract.DataSetLive(opts, setId)
}

func (c pdpVerifierContract) GetNextChallengeEpoch(opts *bind.CallOpts, setId *big.Int) (*big.Int, error) {
	return c.contract.GetNextChallengeEpoch(opts, setId)
}

func (c pdpVerifierContract) GetActivePieces(opts *bind.CallOpts, setId *big.Int, offset *big.Int, limit *big.Int) (activePiecesResult, error) {
	result, err := c.contract.GetActivePieces(opts, setId, offset, limit)
	if err != nil {
		return activePiecesResult{}, err
	}
	return activePiecesResult{
		Pieces:  result.Pieces,
		HasMore: result.HasMore,
	}, nil
}

// ChainPDPClient implements PDPClient using the PDPVerifier contract on-chain.
type ChainPDPClient struct {
	ethClient *ethclient.Client
	contract  pdpVerifierAPI
	pageSize  uint64
}

// NewPDPClient creates a new PDP client backed by the PDPVerifier contract.
func NewPDPClient(ctx context.Context, rpcURL string) (*ChainPDPClient, error) {
	if rpcURL == "" {
		return nil, errors.New("rpc URL is required")
	}

	ethClient, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to PDP RPC")
	}

	network, _, err := synapse.DetectNetwork(ctx, ethClient)
	if err != nil {
		ethClient.Close()
		return nil, errors.Wrap(err, "failed to detect PDP network")
	}

	contractAddr := constants.GetPDPVerifierAddress(network)
	if contractAddr == (common.Address{}) {
		ethClient.Close()
		return nil, errors.New("unsupported PDP network: missing contract address")
	}

	verifier, err := contracts.NewPDPVerifier(contractAddr, ethClient)
	if err != nil {
		ethClient.Close()
		return nil, errors.Wrap(err, "failed to initialize PDP verifier contract")
	}

	return &ChainPDPClient{
		ethClient: ethClient,
		contract:  pdpVerifierContract{contract: verifier},
		pageSize:  pdpDefaultPageSize,
	}, nil
}

// Close releases the underlying RPC client.
func (c *ChainPDPClient) Close() error {
	if c.ethClient == nil {
		return nil
	}
	c.ethClient.Close()
	return nil
}

// GetProofSetsForClient returns all proof sets associated with a client address.
func (c *ChainPDPClient) GetProofSetsForClient(ctx context.Context, clientAddress address.Address) ([]ProofSetInfo, error) {
	listenerAddr, err := delegatedAddressToCommon(clientAddress)
	if err != nil {
		return nil, err
	}

	nextID, err := c.contract.GetNextDataSetId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get next data set ID")
	}

	var proofSets []ProofSetInfo
	for setID := uint64(0); setID < nextID; setID++ {
		setIDBig := new(big.Int).SetUint64(setID)

		listener, err := c.contract.GetDataSetListener(&bind.CallOpts{Context: ctx}, setIDBig)
		if err != nil {
			Logger.Debugw("failed to get PDP data set listener", "setID", setID, "error", err)
			continue
		}
		if listener != listenerAddr {
			continue
		}

		info, err := c.buildProofSetInfo(ctx, setID, listener)
		if err != nil {
			Logger.Warnw("failed to build PDP proof set info", "setID", setID, "error", err)
			continue
		}
		proofSets = append(proofSets, *info)
	}

	return proofSets, nil
}

// GetProofSetInfo returns detailed information about a specific proof set.
func (c *ChainPDPClient) GetProofSetInfo(ctx context.Context, proofSetID uint64) (*ProofSetInfo, error) {
	listener, err := c.contract.GetDataSetListener(&bind.CallOpts{Context: ctx}, new(big.Int).SetUint64(proofSetID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get PDP data set listener")
	}
	return c.buildProofSetInfo(ctx, proofSetID, listener)
}

// IsProofSetLive checks if a proof set is actively being challenged.
func (c *ChainPDPClient) IsProofSetLive(ctx context.Context, proofSetID uint64) (bool, error) {
	live, err := c.contract.DataSetLive(&bind.CallOpts{Context: ctx}, new(big.Int).SetUint64(proofSetID))
	if err != nil {
		return false, errors.Wrap(err, "failed to check PDP data set live status")
	}
	return live, nil
}

// GetNextChallengeEpoch returns the next challenge epoch for a proof set.
func (c *ChainPDPClient) GetNextChallengeEpoch(ctx context.Context, proofSetID uint64) (int32, error) {
	epoch, err := c.contract.GetNextChallengeEpoch(&bind.CallOpts{Context: ctx}, new(big.Int).SetUint64(proofSetID))
	if err != nil {
		return 0, errors.Wrap(err, "failed to get PDP next challenge epoch")
	}
	if !epoch.IsInt64() || epoch.Int64() > math.MaxInt32 {
		return 0, fmt.Errorf("PDP next challenge epoch out of range: %s", epoch.String())
	}
	return int32(epoch.Int64()), nil
}

func (c *ChainPDPClient) buildProofSetInfo(ctx context.Context, setID uint64, listener common.Address) (*ProofSetInfo, error) {
	setIDBig := new(big.Int).SetUint64(setID)

	storageProvider, _, err := c.contract.GetDataSetStorageProvider(&bind.CallOpts{Context: ctx}, setIDBig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get PDP data set storage provider")
	}

	isLive, err := c.contract.DataSetLive(&bind.CallOpts{Context: ctx}, setIDBig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check PDP data set live status")
	}

	nextChallenge, err := c.contract.GetNextChallengeEpoch(&bind.CallOpts{Context: ctx}, setIDBig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get PDP next challenge epoch")
	}
	if !nextChallenge.IsInt64() || nextChallenge.Int64() > math.MaxInt32 {
		return nil, fmt.Errorf("PDP next challenge epoch out of range: %s", nextChallenge.String())
	}

	pieces, err := c.getPieceCIDs(ctx, setIDBig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get PDP active pieces")
	}

	clientAddr, err := commonToDelegatedAddress(listener)
	if err != nil {
		return nil, err
	}
	providerAddr, err := commonToDelegatedAddress(storageProvider)
	if err != nil {
		return nil, err
	}

	return &ProofSetInfo{
		ProofSetID:         setID,
		ClientAddress:      clientAddr,
		ProviderAddress:    providerAddr,
		IsLive:             isLive,
		NextChallengeEpoch: int32(nextChallenge.Int64()),
		PieceCIDs:          pieces,
	}, nil
}

func (c *ChainPDPClient) getPieceCIDs(ctx context.Context, setID *big.Int) ([]cid.Cid, error) {
	var (
		offset uint64
		result []cid.Cid
	)

	for {
		pieces, err := c.contract.GetActivePieces(
			&bind.CallOpts{Context: ctx},
			setID,
			new(big.Int).SetUint64(offset),
			new(big.Int).SetUint64(c.pageSize),
		)
		if err != nil {
			return nil, err
		}

		for i, piece := range pieces.Pieces {
			parsed, err := cid.Cast(piece.Data)
			if err != nil {
				return nil, errors.Wrapf(err, "invalid piece CID at index %d", i)
			}
			result = append(result, parsed)
		}

		if !pieces.HasMore || len(pieces.Pieces) == 0 {
			break
		}
		offset += uint64(len(pieces.Pieces))
	}

	return result, nil
}

func delegatedAddressToCommon(addr address.Address) (common.Address, error) {
	if addr == address.Undef {
		return common.Address{}, errors.New("client address is required")
	}
	if addr.Protocol() != address.Delegated {
		return common.Address{}, fmt.Errorf("client address must be delegated (f4), got protocol %d", addr.Protocol())
	}

	namespace, n, err := varint.FromUvarint(addr.Payload())
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to decode delegated namespace")
	}
	subaddr := addr.Payload()[n:]
	if namespace != 10 {
		return common.Address{}, fmt.Errorf("unsupported delegated namespace %d", namespace)
	}
	if len(subaddr) != common.AddressLength {
		return common.Address{}, fmt.Errorf("invalid delegated address length: %d", len(subaddr))
	}

	return common.BytesToAddress(subaddr), nil
}

func commonToDelegatedAddress(subaddr common.Address) (address.Address, error) {
	addr, err := address.NewDelegatedAddress(10, subaddr.Bytes())
	if err != nil {
		return address.Undef, errors.Wrap(err, "failed to encode delegated address")
	}
	return addr, nil
}
