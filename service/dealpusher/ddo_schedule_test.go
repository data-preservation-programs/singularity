package dealpusher

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/data-preservation-programs/go-synapse/signer"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type ddoDealManagerMock struct {
	spConfig      *DDOSPConfig
	allocIDs      []uint64
	pieces        []DDOPieceSubmission
	ensureCalled  bool
	lastCreateCfg DDOSchedulingConfig
}

func (m *ddoDealManagerMock) ValidateSP(_ context.Context, _ uint64) (*DDOSPConfig, error) {
	return m.spConfig, nil
}

func (m *ddoDealManagerMock) CheckBalance(_ context.Context, _ common.Address) (*DDOBalanceStatus, error) {
	return &DDOBalanceStatus{
		NativeFIL:      new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18)),
		TokenBalance:   new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e18)),
		DepositedFunds: big.NewInt(0),
		LockupCurrent:  big.NewInt(0),
		Available:      big.NewInt(0),
	}, nil
}

func (m *ddoDealManagerMock) EnsurePayments(_ context.Context, _ signer.EVMSigner, _ []DDOPieceSubmission, _ DDOSchedulingConfig) error {
	m.ensureCalled = true
	return nil
}

func (m *ddoDealManagerMock) CreateAllocations(_ context.Context, _ signer.EVMSigner, pieces []DDOPieceSubmission, cfg DDOSchedulingConfig) (*DDOQueuedTx, error) {
	m.pieces = append([]DDOPieceSubmission(nil), pieces...)
	m.lastCreateCfg = cfg
	return &DDOQueuedTx{Hash: "0xddo123"}, nil
}

func (m *ddoDealManagerMock) WaitForConfirmations(_ context.Context, txHash string, _ uint64, _ time.Duration) (*DDOTransactionReceipt, error) {
	return &DDOTransactionReceipt{Hash: txHash, Status: 1}, nil
}

func (m *ddoDealManagerMock) ParseAllocationIDs(_ context.Context, _ string) ([]uint64, error) {
	return m.allocIDs, nil
}

func TestDealPusher_RunSchedule_DDOWithoutDependenciesErrors(t *testing.T) {
	d := &DealPusher{
		scheduleDealTypeResolver: func(_ *model.Schedule) model.DealType { return model.DealTypeDDO },
	}
	state, err := d.runSchedule(context.Background(), &model.Schedule{})
	require.Error(t, err)
	require.Equal(t, model.ScheduleError, state)
	require.Contains(t, err.Error(), "ddo scheduling dependencies are not configured")
}

func TestDealPusher_RunSchedule_DDOCreatesDealsWithAllocations(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		clientSubaddr := make([]byte, 20)
		clientSubaddr[19] = 10
		clientAddr, err := address.NewDelegatedAddress(10, clientSubaddr)
		require.NoError(t, err)

		prep := model.Preparation{Name: "prep"}
		require.NoError(t, db.Create(&prep).Error)

		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)
		keyPath, _, err := ks.Put(testutil.TestPrivateKeyHex)
		require.NoError(t, err)

		actorID := "f01001"
		require.NoError(t, db.Create(&model.Actor{ID: actorID, Address: clientAddr.String()}).Error)

		wallet := model.Wallet{
			Address:  clientAddr.String(),
			KeyPath:  keyPath,
			KeyStore: "local",
			ActorID:  &actorID,
		}
		require.NoError(t, db.Create(&wallet).Error)
		require.NoError(t, db.Model(&prep).Update("wallet_id", wallet.ID).Error)

		storage := model.Storage{Name: "src-storage"}
		require.NoError(t, db.Create(&storage).Error)
		attachment := model.SourceAttachment{PreparationID: prep.ID, StorageID: storage.ID}
		require.NoError(t, db.Create(&attachment).Error)

		pieceCID := model.CID(calculateCommp(t, generateRandomBytes(1000), 1024))
		car := model.Car{
			AttachmentID:  &attachment.ID,
			PreparationID: &prep.ID,
			PieceCID:      pieceCID,
			PieceSize:     1024,
			StoragePath:   "car-1",
		}
		require.NoError(t, db.Create(&car).Error)

		schedule := model.Schedule{
			PreparationID:   prep.ID,
			State:           model.ScheduleActive,
			Provider:        "f01234",
			TotalDealNumber: 1,
			URLTemplate:     "https://example.com/{PIECE_CID}",
		}
		require.NoError(t, db.Create(&schedule).Error)
		schedule.Preparation = &model.Preparation{Wallet: &wallet}

		mock := &ddoDealManagerMock{
			spConfig: &DDOSPConfig{
				IsActive:     true,
				MinPieceSize: 128,
				MaxPieceSize: 1 << 30,
				MinTermLen:   100,
				MaxTermLen:   1000000,
			},
			allocIDs: []uint64{42},
		}

		d := &DealPusher{
			dbNoContext:              db,
			keyStore:                 ks,
			ddoDealManager:           mock,
			ddoSchedulingConfig:      defaultDDOSchedulingConfig(),
			scheduleDealTypeResolver: func(_ *model.Schedule) model.DealType { return model.DealTypeDDO },
		}

		state, err := d.runSchedule(ctx, &schedule)
		require.NoError(t, err)
		require.Equal(t, model.ScheduleCompleted, state)

		require.True(t, mock.ensureCalled)
		require.Len(t, mock.pieces, 1)
		require.Equal(t, cid.Cid(pieceCID), mock.pieces[0].PieceCID)
		require.Equal(t, uint64(1024), mock.pieces[0].PieceSize)
		require.Equal(t, uint64(1234), mock.pieces[0].ProviderID)
		require.Contains(t, mock.pieces[0].DownloadURL, cid.Cid(pieceCID).String())

		var deals []model.Deal
		require.NoError(t, db.Where("schedule_id = ?", schedule.ID).Find(&deals).Error)
		require.Len(t, deals, 1)
		require.Equal(t, model.DealTypeDDO, deals[0].DealType)
		require.Equal(t, model.DealProposed, deals[0].State)
		require.NotNil(t, deals[0].DDOAllocationID)
		require.Equal(t, uint64(42), *deals[0].DDOAllocationID)
		require.NotNil(t, deals[0].WalletID)
		require.Equal(t, wallet.ID, *deals[0].WalletID)
	})
}

func TestDealPusher_RunSchedule_DDOInactiveSPErrors(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		clientSubaddr := make([]byte, 20)
		clientSubaddr[19] = 10
		clientAddr, err := address.NewDelegatedAddress(10, clientSubaddr)
		require.NoError(t, err)

		prep := model.Preparation{Name: "prep"}
		require.NoError(t, db.Create(&prep).Error)

		ks, err := keystore.NewLocalKeyStore(t.TempDir())
		require.NoError(t, err)
		keyPath, _, err := ks.Put(testutil.TestPrivateKeyHex)
		require.NoError(t, err)

		actorID := "f01001"
		require.NoError(t, db.Create(&model.Actor{ID: actorID, Address: clientAddr.String()}).Error)
		wallet := model.Wallet{
			Address:  clientAddr.String(),
			KeyPath:  keyPath,
			KeyStore: "local",
			ActorID:  &actorID,
		}
		require.NoError(t, db.Create(&wallet).Error)
		require.NoError(t, db.Model(&prep).Update("wallet_id", wallet.ID).Error)

		storage := model.Storage{Name: "src-storage"}
		require.NoError(t, db.Create(&storage).Error)
		attachment := model.SourceAttachment{PreparationID: prep.ID, StorageID: storage.ID}
		require.NoError(t, db.Create(&attachment).Error)

		schedule := model.Schedule{
			PreparationID:   prep.ID,
			State:           model.ScheduleActive,
			Provider:        "f09999",
			TotalDealNumber: 1,
		}
		require.NoError(t, db.Create(&schedule).Error)
		schedule.Preparation = &model.Preparation{Wallet: &wallet}

		mock := &ddoDealManagerMock{
			spConfig: &DDOSPConfig{IsActive: false},
		}

		d := &DealPusher{
			dbNoContext:              db,
			keyStore:                 ks,
			ddoDealManager:           mock,
			ddoSchedulingConfig:      defaultDDOSchedulingConfig(),
			scheduleDealTypeResolver: func(_ *model.Schedule) model.DealType { return model.DealTypeDDO },
		}

		state, err := d.runSchedule(ctx, &schedule)
		require.Error(t, err)
		require.Equal(t, model.ScheduleError, state)
		require.Contains(t, err.Error(), "not active")
	})
}

func TestDealPusher_RunSchedule_DDODurationClampsTermMax(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		cases := []struct {
			name        string
			duration    time.Duration
			daemonMax   int64
			wantTermMax int64
		}{
			{"duration below daemon ceiling", 100 * 30 * time.Second, 5256000, 100},
			{"duration at daemon ceiling", 5256000 * 30 * time.Second, 5256000, 5256000},
			{"duration above daemon ceiling -- clamped", 10000000 * 30 * time.Second, 5256000, 5256000},
			{"zero duration falls back to daemon", 0, 5256000, 5256000},
		}

		// Per-case privkeys so each wallet row has a unique (address, key_path)
		// without fighting the keystore's content-addressed filenames.
		privs := []string{
			"7b2254797065223a22736563703235366b31222c22507269766174654b6579223a226b35507976337148327349586343595a58594f5775453149326e32554539436861556b6c4e36695a5763453d227d",
			"7b2254797065223a22736563703235366b31222c22507269766174654b6579223a224f4c6d6f575a6b6a33346b422f612b716351716f48456779374f6558646d784f2b5444562b5352735433673d227d",
			"7b2254797065223a22736563703235366b31222c22507269766174654b6579223a22375768664c3943672b444b516f46562f5643434f4b346c536f4d505437312b43786379797434735031356b3d227d",
			"7b2254797065223a22736563703235366b31222c22507269766174654b6579223a2279616a746d62743459307a7a4578306964646837562f6f6c624b42762f4a786838757179766241356342383d227d",
		}
		require.GreaterOrEqual(t, len(privs), len(cases))

		for i, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				prep := model.Preparation{Name: "prep-" + tc.name}
				require.NoError(t, db.Create(&prep).Error)

				ks, err := keystore.NewLocalKeyStore(t.TempDir())
				require.NoError(t, err)
				keyPath, _, err := ks.Put(privs[i])
				require.NoError(t, err)
				clientAddr, err := keystore.AddressFromExport(privs[i])
				require.NoError(t, err)

				actorID := fmt.Sprintf("f0100%d", i+1)
				require.NoError(t, db.Create(&model.Actor{ID: actorID, Address: clientAddr.String()}).Error)

				wallet := model.Wallet{
					Address:  clientAddr.String(),
					KeyPath:  keyPath,
					KeyStore: "local",
					ActorID:  &actorID,
				}
				require.NoError(t, db.Create(&wallet).Error)
				require.NoError(t, db.Model(&prep).Update("wallet_id", wallet.ID).Error)

				storage := model.Storage{Name: "src-" + tc.name}
				require.NoError(t, db.Create(&storage).Error)
				attachment := model.SourceAttachment{PreparationID: prep.ID, StorageID: storage.ID}
				require.NoError(t, db.Create(&attachment).Error)

				pieceCID := model.CID(calculateCommp(t, generateRandomBytes(1000), 1024))
				car := model.Car{
					AttachmentID:  &attachment.ID,
					PreparationID: &prep.ID,
					PieceCID:      pieceCID,
					PieceSize:     1024,
					StoragePath:   "car-" + tc.name,
				}
				require.NoError(t, db.Create(&car).Error)

				schedule := model.Schedule{
					PreparationID:   prep.ID,
					State:           model.ScheduleActive,
					Provider:        "f01234",
					TotalDealNumber: 1,
					URLTemplate:     "https://example.com/{PIECE_CID}",
					Duration:        tc.duration,
				}
				require.NoError(t, db.Create(&schedule).Error)
				schedule.Preparation = &model.Preparation{Wallet: &wallet}

				mock := &ddoDealManagerMock{
					spConfig: &DDOSPConfig{
						IsActive:     true,
						MinPieceSize: 128,
						MaxPieceSize: 1 << 30,
						MinTermLen:   100,
						MaxTermLen:   10000000,
					},
					allocIDs: []uint64{42},
				}

				cfg := defaultDDOSchedulingConfig()
				cfg.TermMax = tc.daemonMax
				cfg.TermMin = 1 // keep under any test duration above

				d := &DealPusher{
					dbNoContext:              db,
					keyStore:                 ks,
					ddoDealManager:           mock,
					ddoSchedulingConfig:      cfg,
					scheduleDealTypeResolver: func(_ *model.Schedule) model.DealType { return model.DealTypeDDO },
				}

				state, err := d.runSchedule(ctx, &schedule)
				require.NoError(t, err)
				require.Equal(t, model.ScheduleCompleted, state)
				require.Equal(t, tc.wantTermMax, mock.lastCreateCfg.TermMax,
					"expected TermMax=%d with duration=%v and daemonMax=%d, got %d",
					tc.wantTermMax, tc.duration, tc.daemonMax, mock.lastCreateCfg.TermMax)
			})
		}
	})
}

func TestParseProviderActorID(t *testing.T) {
	tests := []struct {
		input string
		want  uint64
		err   bool
	}{
		{"f01234", 1234, false},
		{"t01234", 1234, false},
		{"1234", 1234, false},
		{"f0", 0, true},
		{"bogus", 0, true},
	}
	for _, tc := range tests {
		got, err := parseProviderActorID(tc.input)
		if tc.err {
			require.Error(t, err, "input: %s", tc.input)
		} else {
			require.NoError(t, err, "input: %s", tc.input)
			require.Equal(t, tc.want, got, "input: %s", tc.input)
		}
	}
}
