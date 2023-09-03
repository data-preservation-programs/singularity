package deal

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestListHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		err := db.Create(&model.Preparation{
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
		deals, err := Default.ListHandler(ctx, db, ListDealRequest{
			Preparations: []string{"1"},
			Sources:      []string{"storage"},
			Schedules:    []uint32{1},
			Providers:    []string{"provider"},
			States:       []model.DealState{model.DealActive},
		})
		require.NoError(t, err)
		require.Len(t, deals, 1)
	})

}
