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
	Usage: "Send a manual PDP deal on-chain",
	Description: `Create/reuse a proof set and add a piece to it on-chain via PDPVerifier.
  Example: singularity deal send-manual-pdp --client f1xxx --provider t410fxxx --piece-cid bagaxxxx --piece-size 1048576 --eth-rpc http://localhost:5700/rpc/v1`,
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
			Usage:    "Piece CID (commp)",
			Required: true,
		},
		&cli.Int64Flag{
			Name:     "piece-size",
			Usage:    "Piece size in bytes",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "eth-rpc",
			Usage:    "FEVM JSON-RPC endpoint",
			EnvVars:  []string{"ETH_RPC_URL"},
			Required: true,
		},
		&cli.Uint64Flag{
			Name:  "confirmation-depth",
			Usage: "Blocks to wait for tx confirmation",
			Value: 5,
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		pdp, err := dealpusher.NewOnChainPDP(c.Context, db, c.String("eth-rpc"))
		if err != nil {
			return errors.WithStack(err)
		}
		defer pdp.Close()

		// load wallet
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

		provider := c.String("provider")

		cfg := dealpusher.PDPSchedulingConfig{
			BatchSize:            1,
			MaxPiecesPerProofSet: 1024,
			ConfirmationDepth:    c.Uint64("confirmation-depth"),
			PollingInterval:      5 * time.Second,
		}

		// ensure proof set exists (or create one)
		fmt.Println("ensuring proof set...")
		proofSetID, err := pdp.EnsureProofSet(c.Context, evmSigner, provider, cfg)
		if err != nil {
			return errors.Wrap(err, "failed to ensure proof set")
		}
		fmt.Printf("proof set ID: %d\n", proofSetID)

		// parse piece cid
		pieceCID, err := cid.Parse(c.String("piece-cid"))
		if err != nil {
			return errors.Wrap(err, "invalid piece CID")
		}

		// add piece to proof set
		fmt.Println("submitting add-roots tx...")
		pieceSize := c.Int64("piece-size")
		queuedTx, err := pdp.QueueAddRoots(c.Context, evmSigner, proofSetID, []cid.Cid{pieceCID}, []int64{pieceSize}, cfg)
		if err != nil {
			return errors.Wrap(err, "failed to add roots")
		}
		fmt.Printf("tx: %s\n", queuedTx.Hash)

		// wait for confirmation
		fmt.Println("waiting for confirmation...")
		receipt, err := pdp.WaitForConfirmations(c.Context, queuedTx.Hash, cfg.ConfirmationDepth, cfg.PollingInterval)
		if err != nil {
			return errors.Wrap(err, "tx failed")
		}
		fmt.Printf("confirmed at block %d (gas: %d)\n", receipt.BlockNumber, receipt.GasUsed)

		// save deal record
		proofSetIDCopy := proofSetID
		dealModel := &model.Deal{
			State:      model.DealProposed,
			DealType:   model.DealTypePDP,
			Provider:   provider,
			PieceCID:   model.CID(pieceCID),
			PieceSize:  pieceSize,
			WalletID:   &walletObj.ID,
			ProofSetID: &proofSetIDCopy,
		}
		if err := db.WithContext(c.Context).Create(dealModel).Error; err != nil {
			return errors.Wrap(err, "failed to save deal")
		}
		cliutil.Print(c, dealModel)
		return nil
	},
}
