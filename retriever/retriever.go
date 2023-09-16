package retriever

import (
	"context"
	"io"
	"sync"

	"github.com/data-preservation-programs/singularity/retriever/deserializer"
	lassietypes "github.com/filecoin-project/lassie/pkg/types"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car/v2"
	"github.com/ipld/go-car/v2/storage"
	trustlessutils "github.com/ipld/go-trustless-utils"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multicodec"
	"go.uber.org/multierr"
)

type EndpointFinder interface {
	FindHTTPEndpoints(ctx context.Context, sps []string) ([]peer.AddrInfo, error)
}

// a retriever returns a byte stream for a cid at the root of a unixfs tree,
// from a list of Filecoin providers
type Retriever struct {
	lassie         lassietypes.Fetcher
	endpointFinder EndpointFinder
}

func NewRetriever(lassie lassietypes.Fetcher, endpointFinder EndpointFinder) *Retriever {
	return &Retriever{
		lassie:         lassie,
		endpointFinder: endpointFinder,
	}
}

// deserialize takes an reader of a carFile and writes the deserialized output
func (r *Retriever) deserialize(ctx context.Context, c cid.Cid, rangeStart int64, rangeEnd int64, carInput io.Reader, carOutput io.Writer) error {
	cr, err := car.NewBlockReader(carInput)
	if err != nil {
		return err
	}
	_, err = deserializer.Deserialize(ctx, cr, c, rangeStart, rangeEnd, carOutput)
	return err
}

// getContent fetches content through Lassie and writes the CAR file to an output writer
func (r *Retriever) getContent(ctx context.Context, c cid.Cid, rangeStart int64, rangeEnd int64, sps []string, carOutput io.Writer) error {
	writable, err := storage.NewWritable(carOutput, []cid.Cid{c}, car.WriteAsCarV1(true))
	providerAddrs, err := r.endpointFinder.FindHTTPEndpoints(ctx, sps)
	if err != nil {
		return err
	}
	inclusiveRangeEnd := rangeEnd - 1
	request, err := lassietypes.NewRequestForPath(writable, c, "", trustlessutils.DagScopeEntity, &trustlessutils.ByteRange{
		From: rangeStart,
		// byte range is INCLUSIVE in the lassie/trustless HTTP context, so decrement
		To: &inclusiveRangeEnd,
	})
	request.Duplicates = true
	request.Protocols = []multicodec.Code{multicodec.TransportIpfsGatewayHttp}
	request.FixedPeers = providerAddrs
	_, err = r.lassie.Fetch(ctx, request, func(lassietypes.RetrievalEvent) {})
	if err != nil {
		return err
	}
	return writable.Finalize()
}

// Retrieve retrieves a range from a cid representing a unixfstree from a given list of SPs, writing the output to a car file
func (r *Retriever) Retrieve(ctx context.Context, c cid.Cid, rangeStart int64, rangeEnd int64, sps []string, out io.Writer) error {
	reader, writer := io.Pipe()
	errChan := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := r.deserialize(ctx, c, rangeStart, rangeEnd, reader, out)
		reader.Close()
		select {
		case <-ctx.Done():
		case errChan <- err:
		}
	}()
	go func() {
		defer wg.Done()
		err := r.getContent(ctx, c, rangeStart, rangeEnd, sps, writer)
		writer.Close()
		select {
		case <-ctx.Done():
		case errChan <- err:
		}
	}()
	wg.Wait()

	// collect errors
	var err error
	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case nextErr := <-errChan:
			multierr.Append(err, nextErr)
		}
	}
	return err
}
