package admin

import (
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestInitHandler(t *testing.T) {
	db := database.OpenInMemory()
	defer model.DropAll(db)
	err := InitHandler(db)
	require.NoError(t, err)
}
