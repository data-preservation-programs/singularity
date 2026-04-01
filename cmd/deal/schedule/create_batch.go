package schedule

import (
	"slices"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	handlerschedule "github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var CreateBatchCmd = &cli.Command{
	Name:  "create-batch",
	Usage: "Create the cross-product of schedules for multiple preparations and providers",
	Description: `Create all schedules for preparations x providers x deal types.

Examples:
  # Cold DDO copies to 2 providers
  singularity deal schedule create-batch \
    --group dataset-a-cold \
    --preparation prep-a --preparation prep-b \
    --provider f01000 --provider f02000 \
    --deal-type ddo

  # Warm + cold to 1 provider
  singularity deal schedule create-batch \
    --group dataset-a-warm \
    --preparation prep-a \
    --provider f03000 \
    --deal-type ddo --deal-type pdp
`,
	Flags: append([]cli.Flag{
		&cli.StringSliceFlag{
			Name:     "preparation",
			Usage:    "Preparation IDs or names to include. Repeat this flag.",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:     "provider",
			Usage:    "Storage provider IDs to include. Repeat this flag.",
			Required: true,
		},
		&cli.StringSliceFlag{
			Name:  "deal-type",
			Usage: "Deal types to create schedules for: market, pdp, or ddo. Repeat for multiple types.",
			Value: cli.NewStringSlice(string(model.DealTypeMarket)),
		},
	}, scheduleCreateFlags(false)...),
	Action: func(c *cli.Context) error {
		if c.String("group") == "" {
			return errors.New("group label is required for create-batch")
		}

		preparations := c.StringSlice("preparation")
		if len(preparations) == 0 {
			return errors.New("at least one preparation is required")
		}
		providers := c.StringSlice("provider")
		if len(providers) == 0 {
			return errors.New("at least one provider is required")
		}

		dealTypes, err := parseDealTypes(c.StringSlice("deal-type"))
		if err != nil {
			return err
		}

		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		allowedPieceCIDs, err := allowedPieceCIDsFromContext(c)
		if err != nil {
			return errors.WithStack(err)
		}

		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))
		created := make([]model.Schedule, 0, len(preparations)*len(dealTypes)*len(providers))

		for _, preparation := range preparations {
			for _, dealType := range dealTypes {
				for _, provider := range providers {
					request := createRequest(c, preparation, provider, string(dealType), allowedPieceCIDs)
					schedule, err := handlerschedule.Default.CreateHandler(c.Context, db, lotusClient, request)
					if err != nil {
						return errors.Wrapf(err, "failed to create schedule for preparation %q, provider %q, deal type %q", preparation, provider, dealType)
					}

					created = append(created, *schedule)
				}
			}
		}

		cliutil.Print(c, created)
		return nil
	},
}

func parseDealTypes(values []string) ([]model.DealType, error) {
	dealTypes := make([]model.DealType, 0, len(values))
	for _, v := range values {
		dt := model.DealType(v)
		if !slices.Contains(model.DealTypes, dt) {
			return nil, errors.Newf("invalid deal type %q: must be one of market, pdp, ddo", v)
		}
		dealTypes = append(dealTypes, dt)
	}
	if len(dealTypes) == 0 {
		return nil, errors.New("at least one deal type is required")
	}
	return dealTypes, nil
}
