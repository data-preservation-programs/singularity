package wallet

import (
	"testing"
)

// TODO: ImportHandler was removed as part of wallet/actor separation (#590)
// This test needs to be replaced with tests for ImportKeystoreHandler which uses
// the new keystore-based approach instead of storing private keys in the database.
// See handler/wallet/import_keystore.go for the new implementation.
func TestImportHandler(t *testing.T) {
	t.Skip("ImportHandler removed - needs replacement with ImportKeystoreHandler tests")
}
