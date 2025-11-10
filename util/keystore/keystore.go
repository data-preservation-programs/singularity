package keystore

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/filecoin-project/go-address"
	"github.com/jsign/go-filsigner/wallet"
)

type KeyStore interface {
	Put(privateKey string) (path string, addr address.Address, err error) // saves key, returns path and address
	Get(path string) (privateKey string, err error)                       // loads key from path
	List() ([]KeyInfo, error)                                             // lists all keys
	Delete(path string) error                                             // removes key
	Has(path string) bool                                                 // checks if key exists
}

type KeyInfo struct {
	Address address.Address
	Path    string
}

// filesystem keystore implementation
type LocalKeyStore struct {
	dir string
}

func NewLocalKeyStore(dir string) (*LocalKeyStore, error) {
	if err := os.MkdirAll(dir, 0700); err != nil {
		return nil, fmt.Errorf("failed to create keystore directory: %w", err)
	}
	return &LocalKeyStore{dir: dir}, nil
}

// lotus/go-filsigner export format expected
func (ks *LocalKeyStore) Put(privateKey string) (string, address.Address, error) {
	addr, err := wallet.PublicKey(privateKey)
	if err != nil {
		return "", address.Undef, fmt.Errorf("failed to derive address from private key: %w", err)
	}

	// file named by address (f1.../f3...)
	filename := addr.String()
	path := filepath.Join(ks.dir, filename)

	if err := os.WriteFile(path, []byte(privateKey), 0600); err != nil {
		return "", address.Undef, fmt.Errorf("failed to write key file: %w", err)
	}

	return path, addr, nil
}

func (ks *LocalKeyStore) Get(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read key file: %w", err)
	}
	return string(data), nil
}

func (ks *LocalKeyStore) List() ([]KeyInfo, error) {
	entries, err := os.ReadDir(ks.dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read keystore directory: %w", err)
	}

	var keys []KeyInfo
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(ks.dir, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			continue // skip unreadable
		}

		// verify valid key by deriving address
		addr, err := wallet.PublicKey(string(data))
		if err != nil {
			continue // skip invalid
		}

		keys = append(keys, KeyInfo{
			Address: addr,
			Path:    path,
		})
	}

	return keys, nil
}

func (ks *LocalKeyStore) Delete(path string) error {
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("failed to delete key file: %w", err)
	}
	return nil
}

func (ks *LocalKeyStore) Has(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
