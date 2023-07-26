package epochutil

import (
	"context"
	"strings"
	"time"

	"github.com/data-preservation-programs/singularity/util"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/pkg/errors"
)

var GENESIS_TIMESTAMP = int32(1598306400)

type result struct {
	Blocks []block
}

type block struct {
	Timestamp int32
}

func Initialize(ctx context.Context, lotusAPI string, lotusToken string) error {
	if strings.HasPrefix(lotusAPI, "https://api.node.glif.io/rpc") {
		return nil
	}

	client := util.NewLotusClient(lotusAPI, lotusToken)
	var r result
	err := client.CallFor(ctx, &r, "Filecoin.ChainGetGenesis")
	if err != nil {
		return errors.Wrap(err, "failed to decide genesis timestamp")
	}

	if len(r.Blocks) == 0 {
		return errors.New("length of blocks from genesis is 0")
	}

	GENESIS_TIMESTAMP = r.Blocks[0].Timestamp
	return nil
}

func EpochToTime(epoch int32) time.Time {
	return time.Unix(int64(epoch)*30+int64(GENESIS_TIMESTAMP), 0)
}

func UnixToEpoch(unix int64) int32 {
	return (int32(unix) - GENESIS_TIMESTAMP) / 30
}

func TimeToEpoch(t time.Time) abi.ChainEpoch {
	return abi.ChainEpoch((t.Unix() - int64(GENESIS_TIMESTAMP)) / 30)
}
