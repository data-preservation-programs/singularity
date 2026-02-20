package keystore

import (
	"fmt"

	"github.com/data-preservation-programs/go-synapse/signer"
	"github.com/data-preservation-programs/singularity/model"
)

// Signer loads a wallet's private key from the keystore and returns
// a go-synapse Signer. For secp256k1 keys the result also satisfies
// signer.EVMSigner (check with signer.AsEVM).
func Signer(ks KeyStore, w model.Wallet) (signer.Signer, error) {
	exported, err := ks.Get(w.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("loading key for wallet %d: %w", w.ID, err)
	}
	s, err := signer.FromLotusExport(exported)
	if err != nil {
		return nil, fmt.Errorf("parsing key for wallet %d: %w", w.ID, err)
	}
	return s, nil
}

// EVMSigner loads a wallet's key and returns an EVMSigner for Ethereum/FEVM
// transaction signing. Returns an error if the key type is BLS.
func EVMSigner(ks KeyStore, w model.Wallet) (signer.EVMSigner, error) {
	s, err := Signer(ks, w)
	if err != nil {
		return nil, err
	}
	evm, ok := signer.AsEVM(s)
	if !ok {
		return nil, fmt.Errorf("wallet %d (%s) is not an EVM-capable key", w.ID, w.Address)
	}
	return evm, nil
}
