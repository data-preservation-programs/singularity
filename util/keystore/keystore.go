package keystore

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/data-preservation-programs/go-synapse/signer"
	"github.com/filecoin-project/go-address"
)

// KeyStore stores keys by a short name (typically the filecoin address).
// The name is a relative identifier -- implementations resolve it against
// their own storage root. Callers must pass back the name returned by Put.
type KeyStore interface {
	Put(privateKey string) (name string, addr address.Address, err error)
	Get(name string) (privateKey string, err error)
	List() ([]KeyInfo, error)
	Delete(name string) error
	Has(name string) bool
}

type KeyInfo struct {
	Address address.Address
	Name    string // relative identifier within the keystore
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

// reject names that would escape the keystore dir or address another directory
func validateName(name string) error {
	if name == "" {
		return fmt.Errorf("empty key name")
	}
	if strings.ContainsRune(name, os.PathSeparator) || name == "." || name == ".." {
		return fmt.Errorf("invalid key name %q: must be a basename", name)
	}
	return nil
}

// lotus wallet export format expected (hex-encoded JSON with Type and PrivateKey)
func (ks *LocalKeyStore) Put(privateKey string) (string, address.Address, error) {
	addr, err := AddressFromExport(privateKey)
	if err != nil {
		return "", address.Undef, fmt.Errorf("failed to derive address from private key: %w", err)
	}

	name := addr.String()
	path := filepath.Join(ks.dir, name)

	if err := os.WriteFile(path, []byte(privateKey), 0600); err != nil {
		return "", address.Undef, fmt.Errorf("failed to write key file: %w", err)
	}

	return name, addr, nil
}

func (ks *LocalKeyStore) Get(name string) (string, error) {
	if err := validateName(name); err != nil {
		return "", err
	}
	data, err := os.ReadFile(filepath.Join(ks.dir, name))
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

		name := entry.Name()
		data, err := os.ReadFile(filepath.Join(ks.dir, name))
		if err != nil {
			continue // skip unreadable
		}

		addr, err := AddressFromExport(string(data))
		if err != nil {
			continue // skip invalid
		}

		keys = append(keys, KeyInfo{
			Address: addr,
			Name:    name,
		})
	}

	return keys, nil
}

func (ks *LocalKeyStore) Delete(name string) error {
	if err := validateName(name); err != nil {
		return err
	}
	if err := os.Remove(filepath.Join(ks.dir, name)); err != nil {
		return fmt.Errorf("failed to delete key file: %w", err)
	}
	return nil
}

func (ks *LocalKeyStore) Has(name string) bool {
	if err := validateName(name); err != nil {
		return false
	}
	_, err := os.Stat(filepath.Join(ks.dir, name))
	return err == nil
}

// AddressFromExport derives filecoin address from a lotus wallet export string
func AddressFromExport(exported string) (address.Address, error) {
	s, err := signer.FromLotusExport(exported)
	if err != nil {
		return address.Undef, err
	}
	return s.FilecoinAddress(), nil
}
