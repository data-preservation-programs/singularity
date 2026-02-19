package pdptracker

import (
	"context"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexsupply/shovel/dig"
	"github.com/indexsupply/shovel/shovel"
	"github.com/indexsupply/shovel/shovel/config"
	"github.com/indexsupply/shovel/wpg"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PDPIndexer runs an embedded Shovel instance that indexes PDPVerifier
// contract events into Postgres tables for consumption by the event processor.
type PDPIndexer struct {
	pgp  *pgxpool.Pool
	conf config.Root
}

// NewPDPIndexer builds the Shovel configuration and runs schema migrations.
// Call Start to begin indexing.
func NewPDPIndexer(ctx context.Context, pgURL string, rpcURL string, chainID uint64, contractAddr common.Address) (*PDPIndexer, error) {
	conf := buildShovelConfig(pgURL, rpcURL, chainID, contractAddr)
	if err := config.ValidateFix(&conf); err != nil {
		return nil, errors.Wrap(err, "invalid shovel config")
	}

	pgp, err := wpg.NewPool(ctx, pgURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create indexer pg pool")
	}

	tx, err := pgp.Begin(ctx)
	if err != nil {
		pgp.Close()
		return nil, errors.Wrap(err, "failed to begin migration tx")
	}
	if _, err := tx.Exec(ctx, shovel.Schema); err != nil {
		//nolint:errcheck
		tx.Rollback(ctx)
		pgp.Close()
		return nil, errors.Wrap(err, "failed to apply shovel schema")
	}
	if err := config.Migrate(ctx, tx, conf); err != nil {
		//nolint:errcheck
		tx.Rollback(ctx)
		pgp.Close()
		return nil, errors.Wrap(err, "failed to migrate integration tables")
	}
	if err := tx.Commit(ctx); err != nil {
		pgp.Close()
		return nil, errors.Wrap(err, "failed to commit migration")
	}

	return &PDPIndexer{pgp: pgp, conf: conf}, nil
}

// Start begins Shovel indexing in the background. Implements service.Server.
func (idx *PDPIndexer) Start(ctx context.Context, exitErr chan<- error) error {
	mgr := shovel.NewManager(ctx, idx.pgp, idx.conf)
	ec := make(chan error, 1)
	go mgr.Run(ec)
	if err := <-ec; err != nil {
		return errors.Wrap(err, "shovel indexer startup failed")
	}
	Logger.Info("shovel indexer started")

	go func() {
		<-ctx.Done()
		idx.pgp.Close()
		Logger.Info("shovel indexer stopped")
		if exitErr != nil {
			exitErr <- nil
		}
	}()

	return nil
}

// Name returns the service name. Implements service.Server.
func (*PDPIndexer) Name() string { return "PDPIndexer" }

const srcName = "fevm"

func buildShovelConfig(pgURL, rpcURL string, chainID uint64, contract common.Address) config.Root {
	addrHex := strings.ToLower(contract.Hex())
	src := config.Source{
		Name:    srcName,
		ChainID: chainID,
		URLs:    []string{rpcURL},
	}

	af := func() dig.BlockData {
		return dig.BlockData{
			Name:   "log_addr",
			Filter: dig.Filter{Op: "contains", Arg: []string{addrHex}},
		}
	}

	return config.Root{
		PGURL:   pgURL,
		Sources: []config.Source{src},
		Integrations: []config.Integration{
			dataSetCreatedIG(src, af()),
			piecesAddedIG(src, af()),
			piecesRemovedIG(src, af()),
			nextProvingPeriodIG(src, af()),
			possessionProvenIG(src, af()),
			dataSetDeletedIG(src, af()),
			spChangedIG(src, af()),
		},
	}
}

func dataSetCreatedIG(src config.Source, af dig.BlockData) config.Integration {
	return config.Integration{
		Name:    "pdp_dataset_created",
		Enabled: true,
		Sources: []config.Source{{Name: src.Name}},
		Table: wpg.Table{
			Name: "pdp_dataset_created",
			Columns: []wpg.Column{
				{Name: "set_id", Type: "numeric"},
				{Name: "storage_provider", Type: "bytea"},
			},
		},
		Block: []dig.BlockData{af},
		Event: dig.Event{
			Name: "DataSetCreated",
			Type: "event",
			Inputs: []dig.Input{
				{Indexed: true, Name: "setId", Type: "uint256", Column: "set_id"},
				{Indexed: true, Name: "storageProvider", Type: "address", Column: "storage_provider"},
			},
		},
	}
}

// piecesAddedIG captures only set_id from the indexed topic. The array fields
// (pieceIds, pieceCids) are not decoded by Shovel; the event processor
// reconciles via getActivePieces RPC instead.
func piecesAddedIG(src config.Source, af dig.BlockData) config.Integration {
	return config.Integration{
		Name:    "pdp_pieces_added",
		Enabled: true,
		Sources: []config.Source{{Name: src.Name}},
		Table: wpg.Table{
			Name: "pdp_pieces_added",
			Columns: []wpg.Column{
				{Name: "set_id", Type: "numeric"},
			},
		},
		Block: []dig.BlockData{af},
		Event: dig.Event{
			Name: "PiecesAdded",
			Type: "event",
			Inputs: []dig.Input{
				{Indexed: true, Name: "setId", Type: "uint256", Column: "set_id"},
				// non-indexed array fields listed for correct signature, not selected
				{Name: "pieceIds", Type: "uint256[]"},
				{Name: "pieceCids", Type: "tuple[]", Components: []dig.Input{
					{Name: "data", Type: "bytes"},
				}},
			},
		},
	}
}

func piecesRemovedIG(src config.Source, af dig.BlockData) config.Integration {
	return config.Integration{
		Name:    "pdp_pieces_removed",
		Enabled: true,
		Sources: []config.Source{{Name: src.Name}},
		Table: wpg.Table{
			Name: "pdp_pieces_removed",
			Columns: []wpg.Column{
				{Name: "set_id", Type: "numeric"},
			},
		},
		Block: []dig.BlockData{af},
		Event: dig.Event{
			Name: "PiecesRemoved",
			Type: "event",
			Inputs: []dig.Input{
				{Indexed: true, Name: "setId", Type: "uint256", Column: "set_id"},
				{Name: "pieceIds", Type: "uint256[]"},
			},
		},
	}
}

func nextProvingPeriodIG(src config.Source, af dig.BlockData) config.Integration {
	return config.Integration{
		Name:    "pdp_next_proving_period",
		Enabled: true,
		Sources: []config.Source{{Name: src.Name}},
		Table: wpg.Table{
			Name: "pdp_next_proving_period",
			Columns: []wpg.Column{
				{Name: "set_id", Type: "numeric"},
				{Name: "challenge_epoch", Type: "numeric"},
				{Name: "leaf_count", Type: "numeric"},
			},
		},
		Block: []dig.BlockData{af},
		Event: dig.Event{
			Name: "NextProvingPeriod",
			Type: "event",
			Inputs: []dig.Input{
				{Indexed: true, Name: "setId", Type: "uint256", Column: "set_id"},
				{Name: "challengeEpoch", Type: "uint256", Column: "challenge_epoch"},
				{Name: "leafCount", Type: "uint256", Column: "leaf_count"},
			},
		},
	}
}

// possessionProvenIG captures only set_id; the challenges tuple array is not
// needed for deal tracking.
func possessionProvenIG(src config.Source, af dig.BlockData) config.Integration {
	return config.Integration{
		Name:    "pdp_possession_proven",
		Enabled: true,
		Sources: []config.Source{{Name: src.Name}},
		Table: wpg.Table{
			Name: "pdp_possession_proven",
			Columns: []wpg.Column{
				{Name: "set_id", Type: "numeric"},
			},
		},
		Block: []dig.BlockData{af},
		Event: dig.Event{
			Name: "PossessionProven",
			Type: "event",
			Inputs: []dig.Input{
				{Indexed: true, Name: "setId", Type: "uint256", Column: "set_id"},
				{Name: "challenges", Type: "tuple[]", Components: []dig.Input{
					{Name: "pieceId", Type: "uint256"},
					{Name: "offset", Type: "uint256"},
				}},
			},
		},
	}
}

func dataSetDeletedIG(src config.Source, af dig.BlockData) config.Integration {
	return config.Integration{
		Name:    "pdp_dataset_deleted",
		Enabled: true,
		Sources: []config.Source{{Name: src.Name}},
		Table: wpg.Table{
			Name: "pdp_dataset_deleted",
			Columns: []wpg.Column{
				{Name: "set_id", Type: "numeric"},
				{Name: "deleted_leaf_count", Type: "numeric"},
			},
		},
		Block: []dig.BlockData{af},
		Event: dig.Event{
			Name: "DataSetDeleted",
			Type: "event",
			Inputs: []dig.Input{
				{Indexed: true, Name: "setId", Type: "uint256", Column: "set_id"},
				{Name: "deletedLeafCount", Type: "uint256", Column: "deleted_leaf_count"},
			},
		},
	}
}

func spChangedIG(src config.Source, af dig.BlockData) config.Integration {
	return config.Integration{
		Name:    "pdp_sp_changed",
		Enabled: true,
		Sources: []config.Source{{Name: src.Name}},
		Table: wpg.Table{
			Name: "pdp_sp_changed",
			Columns: []wpg.Column{
				{Name: "set_id", Type: "numeric"},
				{Name: "old_sp", Type: "bytea"},
				{Name: "new_sp", Type: "bytea"},
			},
		},
		Block: []dig.BlockData{af},
		Event: dig.Event{
			Name: "StorageProviderChanged",
			Type: "event",
			Inputs: []dig.Input{
				{Indexed: true, Name: "setId", Type: "uint256", Column: "set_id"},
				{Indexed: true, Name: "oldStorageProvider", Type: "address", Column: "old_sp"},
				{Indexed: true, Name: "newStorageProvider", Type: "address", Column: "new_sp"},
			},
		},
	}
}
