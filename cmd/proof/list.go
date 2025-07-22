package proof

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/proof"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:  "list",
	Usage: "List proofs with optional filtering",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:  "deal-id",
			Usage: "Filter proofs by deal ID",
		},
		&cli.StringFlag{
			Name:  "proof-type",
			Usage: "Filter proofs by type: replication, spacetime",
		},
		&cli.StringFlag{
			Name:  "provider",
			Usage: "Filter proofs by storage provider",
		},
		&cli.BoolFlag{
			Name:  "verified",
			Usage: "Filter proofs by verification status",
		},
		&cli.BoolFlag{
			Name:  "unverified",
			Usage: "Filter proofs by unverified status",
		},
		&cli.IntFlag{
			Name:  "limit",
			Usage: "Limit number of results",
			Value: 100,
		},
		&cli.IntFlag{
			Name:  "offset",
			Usage: "Offset for pagination",
			Value: 0,
		},
	},
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer func() { _ = closer.Close() }()

		request := proof.ListProofRequest{
			Limit:  c.Int("limit"),
			Offset: c.Int("offset"),
		}

		if c.IsSet("deal-id") {
			dealID := c.Uint64("deal-id")
			request.DealID = &dealID
		}

		if c.IsSet("proof-type") {
			proofType := model.ProofType(c.String("proof-type"))
			request.ProofType = &proofType
		}

		if c.IsSet("provider") {
			provider := c.String("provider")
			request.Provider = &provider
		}

		if c.IsSet("verified") {
			verified := true
			request.Verified = &verified
		}

		if c.IsSet("unverified") {
			verified := false
			request.Verified = &verified
		}

		proofs, err := proof.Default.ListHandler(c.Context, db, request)
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, proofs)
		return nil
	},
}
