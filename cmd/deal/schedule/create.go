package schedule

import (
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	handlerschedule "github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/urfave/cli/v2"
)

var CreateCmd = &cli.Command{
	Name:  "create",
	Usage: "Create a schedule to send out deals to a storage provider",
	Description: `CRON pattern '--schedule-cron': The CRON pattern can either be a descriptor or a standard CRON pattern with optional second field
  Standard CRON:
    ┌───────────── minute (0 - 59)
    │ ┌───────────── hour (0 - 23)
    │ │ ┌───────────── day of the month (1 - 31)
    │ │ │ ┌───────────── month (1 - 12)
    │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday)
    │ │ │ │ │                                   
    │ │ │ │ │
    │ │ │ │ │
    * * * * *

  Optional Second field:
    ┌─────────────  second (0 - 59)
    │ ┌─────────────  minute (0 - 59)
    │ │ ┌─────────────  hour (0 - 23)
    │ │ │ ┌─────────────  day of the month (1 - 31)
    │ │ │ │ ┌─────────────  month (1 - 12)
    │ │ │ │ │ ┌─────────────  day of the week (0 - 6) (Sunday to Saturday)
    │ │ │ │ │ │
    │ │ │ │ │ │
    * * * * * *

  Descriptor:
    @yearly, @annually - Equivalent to 0 0 1 1 *
    @monthly           - Equivalent to 0 0 1 * *
    @weekly            - Equivalent to 0 0 * * 0
    @daily,  @midnight - Equivalent to 0 0 * * *
    @hourly            - Equivalent to 0 * * * *`,
	Flags: append(scheduleTargetFlags(), scheduleCreateFlags(true)...),
	Action: func(c *cli.Context) error {
		db, closer, err := database.OpenFromCLI(c)
		if err != nil {
			return errors.WithStack(err)
		}
		defer closer.Close()

		allowedPieceCIDs, err := allowedPieceCIDsFromContext(c)
		if err != nil {
			return errors.WithStack(err)
		}

		request := createRequest(c, c.String("preparation"), c.String("provider"), c.String("deal-type"), allowedPieceCIDs)
		lotusClient := util.NewLotusClient(c.String("lotus-api"), c.String("lotus-token"))
		schedule, err := handlerschedule.Default.CreateHandler(c.Context, db, lotusClient, request)
		if err != nil {
			return errors.WithStack(err)
		}

		cliutil.Print(c, schedule)
		return nil
	},
}
