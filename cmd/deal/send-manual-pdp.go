package deal

import (
	"fmt"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/wallet"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/dealpusher"
	"github.com/data-preservation-programs/singularity/util/keystore"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var SendManualPDPCmd = &cli.Command{
	Name:  "send-manual-pdp",
	Usage: "Send a manual PDP deal via the FWSS-pull flow",
	Description: `Push a single piece to an SP via Curio's /pdp/piece/pull, then trigger the
SP's on-chain commit (createDataSet+addPieces if no assembling set yet, or addPieces
into the existing one). Useful for e2e/diagnostic testing of the FWSS pull path.
  Example: singularity deal send-manual-pdp --client f1xxx --provider t410fxxx --piece-cid bagaxxxx --piece-size 1048576 --eth-rpc http://localhost:5700/rpc/v1 --source-url-base https://static.example.org`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "client",
			Usage:    "Client wallet address (must be imported)",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "provider",
			Usage:    "Storage provider f4/t4 address",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "piece-cid",
			Usage:    "Piece CID (commp v1)",
			Required: true,
		},
		&cli.Int64Flag{
			Name:     "piece-size",
			Usage:    "Padded piece size in bytes",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "eth-rpc",
			Usage:    "FEVM JSON-RPC endpoint",
			EnvVars:  []string{"ETH_RPC_URL"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "source-url-base",
			Usage:    "HTTPS base where Curio fetches the piece (sourceUrl = <base>/piece/<pieceCidV2>)",
			EnvVars:  []string{"PDP_SOURCE_URL_BASE"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    "record-keeper",
			Usage:   "FWSS contract address. Defaults to network FWSS from go-synapse.",
			EnvVars: []string{"PDP_RECORD_KEEPER"},
		},
		&cli.DurationFlag{
			Name:  "pull-timeout",
			Usage: "How long to wait for Curio to finish each phase",
			Value: 5 * time.Minute,
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		pdp, err := dealpusher.NewOnChainPDP(c.Context, dealpusher.OnChainPDPConfig{
			DB:            db,
			RPCURL:        c.String("eth-rpc"),
			SourceURLBase: c.String("source-url-base"),
			RecordKeeper:  c.String("record-keeper"),
		})
		if err != nil {
			return errors.WithStack(err)
		}
		defer pdp.Close()

		var walletObj model.Wallet
		err = db.WithContext(c.Context).Where("address = ?", c.String("client")).First(&walletObj).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("wallet %s not found -- import it first with 'singularity wallet import'", c.String("client"))
		}
		if err != nil {
			return errors.WithStack(err)
		}

		ks, err := keystore.NewLocalKeyStore(wallet.GetKeystoreDir())
		if err != nil {
			return errors.Wrap(err, "failed to init keystore")
		}
		evmSigner, err := keystore.EVMSigner(ks, walletObj)
		if err != nil {
			return errors.WithStack(err)
		}

		pieceCID, err := cid.Parse(c.String("piece-cid"))
		if err != nil {
			return errors.Wrap(err, "invalid piece CID")
		}
		pieceSize := c.Int64("piece-size")

		cfg := dealpusher.PDPSchedulingConfig{
			BatchSize:            1,
			MaxPiecesPerProofSet: 1024,
			PullTimeout:          c.Duration("pull-timeout"),
		}

		fmt.Println("pushing piece to SP via /pdp/piece/pull + on-chain commit...")
		result, err := pdp.PullPiecesToFWSS(
			c.Context,
			evmSigner,
			c.String("provider"),
			[]dealpusher.PDPPieceInput{{PieceCID: pieceCID, PieceSize: pieceSize}},
			cfg,
		)
		if err != nil {
			return errors.Wrap(err, "FWSS pull push failed")
		}
		fmt.Printf("data set ID: %d\n", result.DataSetID)

		dataSetIDCopy := result.DataSetID
		dealModel := &model.Deal{
			State:      model.DealProposed,
			DealType:   model.DealTypePDP,
			Provider:   c.String("provider"),
			PieceCID:   model.CID(pieceCID),
			PieceSize:  pieceSize,
			WalletID:   &walletObj.ID,
			ProofSetID: &dataSetIDCopy,
		}
		if err := db.WithContext(c.Context).Create(dealModel).Error; err != nil {
			return errors.Wrap(err, "failed to save deal")
		}
		cliutil.Print(c, dealModel)
		return nil
	},
}
