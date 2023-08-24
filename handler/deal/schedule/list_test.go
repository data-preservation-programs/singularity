package schedule

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/require"
)

func TestListHandler(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	err = db.Create(&model.Preparation{
		Wallets: []model.Wallet{{
			ID: "f01",
		}},
	}).Error
	require.NoError(t, err)
	err = db.Create(&model.Schedule{
		PreparationID: 1,
	}).Error
	require.NoError(t, err)
	schedules, err := ListHandler(context.Background(), db)
	require.NoError(t, err)
	require.Len(t, schedules, 1)
}
