package deal

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
		SourceStorages: []model.Storage{{
			Name: "storage",
		}},
	}).Error
	require.NoError(t, err)

	err = db.Create([]model.Deal{
		{
			Schedule: &model.Schedule{
				PreparationID: 1,
			},
			State:    model.DealActive,
			ClientID: "f01",
			Provider: "provider",
		},
	}).Error
	require.NoError(t, err)
	deals, err := ListHandler(context.Background(), db, ListDealRequest{
		Preparations: []uint32{1},
		Sources:      []string{"storage"},
		Schedules:    []uint32{1},
		Providers:    []string{"provider"},
		States:       []model.DealState{model.DealActive},
	})
	require.NoError(t, err)
	require.Len(t, deals, 1)
}
