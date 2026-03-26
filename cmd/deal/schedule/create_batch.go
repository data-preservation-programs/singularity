package schedule

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

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
	Description: `Create all schedules for preparations x providers x replication policy.

Examples:
  singularity deal schedule create-batch \
    --group dataset-a \
    --preparation prep-a --preparation prep-b \
    --provider f01000 --provider f02000 \
    --replication market=1 --replication pdp=1
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
			Name:     "replication",
			Category: "Batching",
			Usage:    "Replication policy entry in dealType=count form, e.g. market=1 or pdp=2. Repeat this flag.",
			Value:    cli.NewStringSlice(fmt.Sprintf("%s=1", model.DealTypeMarket)),
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

		replicationPolicy, err := parseReplicationPolicy(c.StringSlice("replication"))
		if err != nil {
			return errors.WithStack(err)
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
		created := make([]model.Schedule, 0, len(preparations)*len(providers)*len(replicationPolicy))

		for _, preparation := range preparations {
			for _, provider := range providers {
				for _, dealType := range replicationPolicy {
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

func parseReplicationPolicy(values []string) ([]model.DealType, error) {
	policy := make([]model.DealType, 0)
	for _, value := range values {
		parts := strings.SplitN(strings.TrimSpace(value), "=", 2)
		if len(parts) != 2 {
			return nil, errors.Newf("invalid replication policy entry %q: expected dealType=count", value)
		}

		dealType := model.DealType(strings.TrimSpace(parts[0]))
		if !slices.Contains(model.DealTypes, dealType) {
			return nil, errors.Newf("invalid replication deal type %q", dealType)
		}

		count, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil || count < 1 {
			return nil, errors.Newf("invalid replication count %q", parts[1])
		}

		for i := 0; i < count; i++ {
			policy = append(policy, dealType)
		}
	}

	return policy, nil
}
