package admin

import (
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestInitHandler(t *testing.T) {
	db, err := database.OpenInMemory()
	require.NoError(t, err)
	defer model.DropAll(db)
	require.NoError(t, InitHandler(db))
}
