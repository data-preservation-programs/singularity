package must

import (
	"github.com/ipfs/go-log/v2"
)

var logger = log.Logger("must")

func String(s string, err error) string {
	if err != nil {
		logger.Panic(err)
	}
	return s
}
