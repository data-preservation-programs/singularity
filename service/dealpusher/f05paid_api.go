package dealpusher

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
)

// F05PaidDealManager owns the paid f05 schedule execution path.
// The first scaffold PR wires the type into Singularity; a later PR will
// provide the concrete implementation backed by the Singularity payments contract.
type F05PaidDealManager interface {
	RunSchedule(ctx context.Context, schedule *model.Schedule) (model.ScheduleState, error)
}
