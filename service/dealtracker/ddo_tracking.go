package dealtracker

import (
	"context"

	ddocontract "github.com/Eastore-project/ddo-client/pkg/contract/ddo"
	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
)

type DDOTrackingClient struct {
	client *ddocontract.Client
}

func NewDDOTrackingClient(rpcURL, ddoAddr string) (*DDOTrackingClient, error) {
	client, err := ddocontract.NewReadOnlyClientWithParams(rpcURL, ddoAddr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize DDO tracking client")
	}
	return &DDOTrackingClient{client: client}, nil
}

func (c *DDOTrackingClient) GetAllocationInfo(ctx context.Context, allocationID uint64) (*model.DDOAllocationStatus, error) {
	info, err := c.client.GetAllocationInfo(allocationID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch DDO allocation %d", allocationID)
	}
	return &model.DDOAllocationStatus{
		Activated:    info.Activated,
		SectorNumber: info.SectorNumber,
	}, nil
}

func (c *DDOTrackingClient) Close() {
	if c.client != nil {
		c.client.Close()
	}
}
