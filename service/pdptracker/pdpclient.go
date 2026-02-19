package pdptracker

import (
	"context"
	"fmt"
	"math/big"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/go-synapse/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-varint"
)

const pdpDefaultPageSize uint64 = 100

// activePiecesResult wraps the contract return for testability.
type activePiecesResult struct {
	Pieces  []contracts.CidsCid
	HasMore bool
}

// pdpContractCaller is the subset of PDPVerifier calls needed by the
// event processor. Extracted as interface for unit testing.
type pdpContractCaller interface {
	GetDataSetListener(opts *bind.CallOpts, setId *big.Int) (common.Address, error)
	GetActivePieces(opts *bind.CallOpts, setId *big.Int, offset *big.Int, limit *big.Int) (activePiecesResult, error)
}

// pdpVerifierCaller wraps the generated contract binding.
type pdpVerifierCaller struct {
	contract *contracts.PDPVerifier
}

func (c pdpVerifierCaller) GetDataSetListener(opts *bind.CallOpts, setId *big.Int) (common.Address, error) {
	return c.contract.GetDataSetListener(opts, setId)
}

func (c pdpVerifierCaller) GetActivePieces(opts *bind.CallOpts, setId *big.Int, offset *big.Int, limit *big.Int) (activePiecesResult, error) {
	result, err := c.contract.GetActivePieces(opts, setId, offset, limit)
	if err != nil {
		return activePiecesResult{}, err
	}
	return activePiecesResult{Pieces: result.Pieces, HasMore: result.HasMore}, nil
}

// ChainPDPClient provides the minimal RPC calls needed by the event processor:
// getDataSetListener (client address hydration on DataSetCreated) and
// getActivePieces (piece reconciliation on PiecesAdded/Removed).
type ChainPDPClient struct {
	ethClient *ethclient.Client
	contract  pdpContractCaller
	pageSize  uint64
}

// NewPDPClient creates a PDP RPC client for the given contract address.
func NewPDPClient(ctx context.Context, rpcURL string, contractAddr common.Address) (*ChainPDPClient, error) {
	if rpcURL == "" {
		return nil, errors.New("rpc URL is required")
	}

	ethClient, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to PDP RPC")
	}

	verifier, err := contracts.NewPDPVerifier(contractAddr, ethClient)
	if err != nil {
		ethClient.Close()
		return nil, errors.Wrap(err, "failed to init PDP verifier contract")
	}

	return &ChainPDPClient{
		ethClient: ethClient,
		contract:  pdpVerifierCaller{contract: verifier},
		pageSize:  pdpDefaultPageSize,
	}, nil
}

// Close releases the underlying RPC client.
func (c *ChainPDPClient) Close() error {
	if c.ethClient != nil {
		c.ethClient.Close()
	}
	return nil
}

// GetDataSetListener returns the listener (client) address for a proof set.
func (c *ChainPDPClient) GetDataSetListener(ctx context.Context, setID uint64) (common.Address, error) {
	addr, err := c.contract.GetDataSetListener(&bind.CallOpts{Context: ctx}, new(big.Int).SetUint64(setID))
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to get dataset listener")
	}
	return addr, nil
}

// GetActivePieces returns all currently active piece CIDs in a proof set,
// handling pagination internally.
func (c *ChainPDPClient) GetActivePieces(ctx context.Context, setID uint64) ([]cid.Cid, error) {
	setIDBig := new(big.Int).SetUint64(setID)
	var (
		offset uint64
		result []cid.Cid
	)

	for {
		pieces, err := c.contract.GetActivePieces(
			&bind.CallOpts{Context: ctx},
			setIDBig,
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
