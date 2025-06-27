// Package retriever wraps functionality for fetching data from filecoin.
package retriever

import (
	"context"
	"io"

	"github.com/data-preservation-programs/singularity/retriever/deserializer"
	lassietypes "github.com/filecoin-project/lassie/pkg/types"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"github.com/ipld/go-car/v2"
	"github.com/ipld/go-car/v2/storage"
	trustlessutils "github.com/ipld/go-trustless-utils"
	"github.com/ipni/go-libipni/metadata"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multicodec"
	"go.uber.org/multierr"
)

var logger = log.Logger("singularity/retriever")

// EndpointFinder finds http endpoints for set of SPs
type EndpointFinder interface {
	FindHTTPEndpoints(ctx context.Context, sps []string) ([]peer.AddrInfo, error)
}

// Retriever wraps Lassie and HTTP endpoint lookup to returns a byte stream for
// a cid at the root of a unixfs tree, from a list of Filecoin providers
type Retriever struct {
	lassie         lassietypes.Fetcher
	endpointFinder EndpointFinder
}

// NewRetriever returns a new retriever instance
func NewRetriever(lassie lassietypes.Fetcher, endpointFinder EndpointFinder) *Retriever {
	return &Retriever{
		lassie:         lassie,
		endpointFinder: endpointFinder,
	}
}

// Retrieve retrieves a byte range from a cid representing a unixfstree from a given list of SPs, writing the output to a car file
func (r *Retriever) Retrieve(ctx context.Context, c cid.Cid, rangeStart int64, rangeEnd int64, sps []string, out io.Writer) error {
	logger.Infow("retrieving from filecoin", "cid", c, "rangeStart", rangeStart, "rangeEnd", rangeEnd, "sps", sps)
	reader, writer := io.Pipe()
	errChan := make(chan error, 2)
	go func() {
		errChan <- r.deserialize(ctx, c, rangeStart, rangeEnd, reader, out)
		_ = reader.Close()
	}()
	go func() {
		errChan <- r.getContent(ctx, c, rangeStart, rangeEnd, sps, writer)
		_ = writer.Close()
	}()

	// collect errors
	var err error
	for range 2 {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case nextErr := <-errChan:
			err = multierr.Append(err, nextErr)
		}
	}
	return err
}

func (r *Retriever) RetrieveReader(ctx context.Context, c cid.Cid, rangeStart int64, rangeEnd int64, sps []string) (io.ReadCloser, error) {
	reader, writer := io.Pipe()
	go func() {
		err := r.getContent(ctx, c, rangeStart, rangeEnd, sps, writer)
		writer.CloseWithError(err)
	}()

	outReader, outWriter := io.Pipe()
	go func() {
		err := r.deserialize(ctx, c, rangeStart, rangeEnd, reader, outWriter)
		_ = reader.Close()
		outWriter.CloseWithError(err)
	}()

	return outReader, nil
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
	if err != nil {
		return err
	}
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
	if err != nil {
		return err
	}
	request.Duplicates = true
	request.Protocols = []multicodec.Code{multicodec.TransportIpfsGatewayHttp}

	providers := make([]lassietypes.Provider, len(providerAddrs))
	for i := range providerAddrs {
		providers[i].Peer = providerAddrs[i]
		providers[i].Protocols = []metadata.Protocol{&metadata.GraphsyncFilecoinV1{}, &metadata.Bitswap{}, &metadata.IpfsGatewayHttp{}}
	}
	request.Providers = providers
	_, err = r.lassie.Fetch(ctx, request)
	if err != nil {
		return err
	}
	return writable.Finalize()
}
