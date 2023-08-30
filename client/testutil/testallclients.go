package testutil

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/client"
)

func TestWithAllClients(ctx context.Context, t *testing.T, test func(*testing.T, client.Client)) {
}
