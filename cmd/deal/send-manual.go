package deal

import (
	"context"
	"fmt"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	"github.com/data-preservation-programs/singularity/service/statetracker"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/data-preservation-programs/singularity/util/errorcategorization"
	"github.com/dustin/go-humanize"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
)

var SendManualCmd = &cli.Command{
	Name:  "send-manual",
	Usage: "Send a manual deal proposal to boost or legacy market",
	Description: `Send a manual deal proposal to boost or legacy market
  Example: singularity deal send-manual --client f01234 --provider f05678 --piece-cid bagaxxxx --piece-size 32GiB
Notes:
  * The client address must have been imported to the wallet using 'singularity wallet import'
  * The deal proposal will not be saved in the database however will eventually be tracked if the deal tracker is running
  * There is a quick address verification using GLIF API which can be made faster by setting LOTUS_API and LOTUS_TOKEN to your own lotus node`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "save",
			Usage: "Whether to save the deal proposal to the database for tracking purpose",
		},
		&cli.StringFlag{
			Name:     "client",
			Category: "Deal Proposal",
			Usage:    "Client address to send deal from",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "provider",
			Category: "Deal Proposal",
			Usage:    "Storage Provider ID to send deal to",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "piece-cid",
			Category: "Deal Proposal",
			Usage:    "Piece CID of the deal",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "piece-size",
			Category: "Deal Proposal",
			Usage:    "Piece Size of the deal",
			Value:    "32GiB",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:     "http-header",
			Category: "Boost Only",
			Usage:    "http headers to be passed with the request (i.e. key=value)",
		},
		&cli.StringFlag{
			Name:     "http-url",
			Category: "Boost Only",
			Usage:    "URL or URL template with PIECE_CID placeholder for boost to fetch the CAR file, e.g. http://127.0.0.1/piece/{PIECE_CID}.car",
			Aliases:  []string{"url-template"},
			Value:    "",
		},
		&cli.Uint64Flag{
			Name:     "file-size",
			Category: "Boost Only",
			Usage:    "File size in bytes for boost to fetch the CAR file",
			Value:    0,
		},
		&cli.Float64Flag{
			Name:     "price-per-gb-epoch",
			Category: "Deal Proposal",
			Usage:    "Price in FIL per GiB per epoch",
			Value:    0,
		},
		&cli.Float64Flag{
			Name:     "price-per-gb",
			Category: "Deal Proposal",
			Usage:    "Price in FIL  per GiB",
			Value:    0,
		},
		&cli.Float64Flag{
			Name:     "price-per-deal",
			Category: "Deal Proposal",
			Usage:    "Price in FIL per deal",
			Value:    0,
		},
		&cli.StringFlag{
			Name:        "root-cid",
			Category:    "Deal Proposal",
			Usage:       "Root CID that is required as part of the deal proposal, if empty, will be set to empty CID",
			Value:       "bafkqaaa",
			DefaultText: "Empty CID",
		},
		&cli.BoolFlag{
			Name:     "verified",
			Category: "Deal Proposal",
			Usage:    "Whether to propose deals as verified",
			Value:    true,
		},
		&cli.BoolFlag{
			Name:     "ipni",
			Category: "Boost Only",
			Usage:    "Whether to announce the deal to IPNI",
			Value:    true,
		},
		&cli.BoolFlag{
			Name:     "keep-unsealed",
			Category: "Deal Proposal",
			Usage:    "Whether to keep unsealed copy",
			Value:    true,
		},
		&cli.StringFlag{
			Name:        "start-delay",
			Category:    "Deal Proposal",
			Aliases:     []string{"s"},
			Usage:       "Deal start delay in epoch or in duration format, i.e. 1000, 72h",
			Value:       "72h",
			DefaultText: "72h[3 days]",
		},
		&cli.StringFlag{
			Name:        "duration",
			Category:    "Deal Proposal",
			Aliases:     []string{"d"},
			Usage:       "Duration in epoch or in duration format, i.e. 1500000, 2400h",
			Value:       "12840h",
			DefaultText: "12840h[535 days]",
		},
		&cli.DurationFlag{
			Name:        "timeout",
			Usage:       "Timeout for the deal proposal",
			Value:       time.Minute,
			DefaultText: "1m",
		},
	},
	Action: func(c *cli.Context) error {
		lotusAPI := c.String("lotus-api")
		lotusToken := c.String("lotus-token")
		err := epochutil.Initialize(c.Context, lotusAPI, lotusToken)
		if err != nil {
			return errors.WithStack(err)
		}
		proposal := deal.Proposal{
			HTTPHeaders:     c.StringSlice("http-header"),
			URLTemplate:     c.String("url-template"),
			PricePerGBEpoch: c.Float64("price-per-gb-epoch"),
			PricePerGB:      c.Float64("price-per-gb"),
			PricePerDeal:    c.Float64("price-per-deal"),
			RootCID:         c.String("root-cid"),
			Verified:        c.Bool("verified"),
			IPNI:            c.Bool("ipni"),
			KeepUnsealed:    c.Bool("keep-unsealed"),
			StartDelay:      c.String("start-delay"),
			Duration:        c.String("duration"),
			ClientAddress:   c.String("client"),
			ProviderID:      c.String("provider"),
			PieceCID:        c.String("piece-cid"),
			PieceSize:       c.String("piece-size"),
			FileSize:        c.Uint64("file-size"),
		}
		timeout := c.Duration("timeout")
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		ctx, cancel := context.WithTimeout(c.Context, timeout)
		defer cancel()
		h, err := util.InitHost(nil)
		if err != nil {
			return errors.Wrap(err, "failed to init host")
		}
		defer func() { _ = h.Close() }()
		dealMaker := replication.NewDealMaker(
			util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token")),
			h,
			10*timeout,
			timeout,
		)
		// Initialize state tracker for deal lifecycle tracking
		stateTracker := statetracker.NewStateChangeTracker(db)

		dealModel, err := deal.Default.SendManualHandler(ctx, db, dealMaker, proposal)

		// Handle deal failures with state tracking (similar to dealpusher pattern)
		if err != nil {
			if c.Bool("save") {
				// Create context metadata for enhanced error categorization
				contextMetadata := &errorcategorization.ErrorMetadata{
					ProviderID:      proposal.ProviderID,
					PieceCID:        proposal.PieceCID,
					ClientAddress:   proposal.ClientAddress,
					StartEpoch:      func() *int32 { epoch := int32(epochutil.TimeToEpoch(time.Now())); return &epoch }(),
					EndEpoch:        func() *int32 { epoch := int32(epochutil.TimeToEpoch(time.Now().Add(24 * time.Hour))); return &epoch }(),
					AttemptNumber:   func() *int { attempt := 1; return &attempt }(),
					LastAttemptTime: func() *time.Time { t := time.Now(); return &t }(),
				}

				// Try to parse piece size for metadata
				if pieceSize, sizeErr := humanize.ParseBytes(proposal.PieceSize); sizeErr == nil {
					size := int64(pieceSize)
					contextMetadata.PieceSize = &size
				}

				// Categorize the error with context
				categorization := errorcategorization.CategorizeErrorWithContext(err.Error(), contextMetadata)

				// Get wallet for deal record
				var wallet model.Wallet
				if walletErr := wallet.FindByIDOrAddr(db, proposal.ClientAddress); walletErr == nil {
					// Create failed deal record with proper structure
					failedDeal := &model.Deal{
						State:         categorization.DealState,
						ClientActorID: wallet.ActorID,
						Provider:      proposal.ProviderID,
						PieceCID:      model.CID{}, // Will be set from proposal if valid
						PieceSize:     0,           // Will be set from proposal if valid
						StartEpoch:    int32(epochutil.TimeToEpoch(time.Now())),
						EndEpoch:      int32(epochutil.TimeToEpoch(time.Now().Add(24 * time.Hour))),
						Price:         "0", // No price for failed deals
						Verified:      proposal.Verified,
						ErrorMessage:  err.Error(),
						ClientID:      &wallet.ID,
					}

					// Try to parse piece info if possible
					if pieceCID, parseErr := cid.Parse(proposal.PieceCID); parseErr == nil {
						failedDeal.PieceCID = model.CID(pieceCID)
						if pieceSize, sizeErr := humanize.ParseBytes(proposal.PieceSize); sizeErr == nil {
							failedDeal.PieceSize = int64(pieceSize)
						}
					}

					// Save the failed deal record
					if dbErr := database.DoRetry(ctx, func() error { return db.Create(failedDeal).Error }); dbErr == nil {
						// Track the failure state change using enhanced error tracking
						if trackErr := stateTracker.TrackErrorStateChange(ctx, failedDeal, nil, err.Error(), contextMetadata); trackErr == nil {
							fmt.Printf("Tracked deal failure with enhanced categorization - ID: %d, state: %s, category: %s, severity: %s, retryable: %t\n",
								failedDeal.ID, categorization.DealState, categorization.Category, categorization.Severity, categorization.Retryable)
						}
					}
				}
			}
			return errors.WithStack(err)
		}

		// Handle successful deal creation
		if c.Bool("save") {
			db = db.WithContext(ctx)
			err = database.DoRetry(ctx, func() error {
				return db.Create(dealModel).Error
			})
			if err != nil {
				return errors.WithStack(err)
			}

			// Track successful deal proposal creation
			metadata := &statetracker.StateChangeMetadata{
				Reason:       "Manual deal proposal created and sent to storage provider",
				StoragePrice: dealModel.Price,
			}
			if trackErr := stateTracker.TrackStateChange(ctx, dealModel, nil, dealModel.State, metadata); trackErr != nil {
				fmt.Printf("Warning: Failed to track deal state change: %v\n", trackErr)
			}
		}
		cliutil.Print(c, dealModel)
		return nil
	},
}
