package retriever_test

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"io"
	"testing"

	"github.com/data-preservation-programs/singularity/retriever"
	lassietypes "github.com/filecoin-project/lassie/pkg/types"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-unixfsnode"
	"github.com/ipfs/go-unixfsnode/testutil"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/linking"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/storage/memstore"
	"github.com/ipld/go-ipld-prime/traversal"
	"github.com/ipld/go-ipld-prime/traversal/selector"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/stretchr/testify/require"
)

func TestRetrieve(t *testing.T) {
	lsys := cidlink.DefaultLinkSystem()
	memSys := memstore.Store{
		Bag: make(map[string][]byte),
	}
	lsys.SetReadStorage(&memSys)
	lsys.SetWriteStorage(&memSys)
	expectedBytes := make([]byte, 4<<20)
	expectedBytesWriter := bytes.NewBuffer(expectedBytes)
	fileReader := io.TeeReader(rand.Reader, expectedBytesWriter)
	file := testutil.GenerateFile(t, &lsys, fileReader, 4<<20)
	fl := &fakeLassie{lsys: &lsys}
	ef := &fakeEndpointFinder{
		endpoints: map[string]peer.AddrInfo{
			"apples": {
				ID: peer.ID("apple tree"),
			},
			"oranges": {
				ID: peer.ID("orange tree"),
			},
			"cheese": {
				ID: peer.ID("cheese cave"),
			},
		},
	}
	retriever := retriever.NewRetriever(fl, ef)
	verifyRetrieval(t, expectedBytes, retriever, file.Root, 0, 4<<20, []string{"apples", "oranges"})
	var providerAddrs []peer.AddrInfo
	for _, provider := range fl.lastRequest.Providers {
		providerAddrs = append(providerAddrs, provider.Peer)
	}
	require.Equal(t, []peer.AddrInfo{{ID: peer.ID("apple tree")}, {ID: peer.ID("orange tree")}}, providerAddrs)
	require.Equal(t, int64(0), fl.lastRequest.Bytes.From)
	require.Equal(t, int64(4<<20)-1, *fl.lastRequest.Bytes.To)
	verifyRetrieval(t, expectedBytes, retriever, file.Root, 1, (1<<18)+1, []string{"apples", "cheese"})
	providerAddrs = providerAddrs[:0]
	for _, provider := range fl.lastRequest.Providers {
		providerAddrs = append(providerAddrs, provider.Peer)
	}
	require.Equal(t, []peer.AddrInfo{{ID: peer.ID("apple tree")}, {ID: peer.ID("cheese cave")}}, providerAddrs)
	require.Equal(t, int64(1), fl.lastRequest.Bytes.From)
	require.Equal(t, int64(1<<18), *fl.lastRequest.Bytes.To)
}

func verifyRetrieval(t *testing.T, expectedBytes []byte, retriever *retriever.Retriever, root cid.Cid, start int64, end int64, sps []string) {
	outBytes := make([]byte, end-start)
	outBuff := bytes.NewBuffer(outBytes)
	err := retriever.Retrieve(context.Background(), root, start, end, sps, outBuff)
	require.NoError(t, err)
	require.Equal(t, expectedBytes[start:end], outBytes)
}

type fakeEndpointFinder struct {
	endpoints map[string]peer.AddrInfo
}

func (ef *fakeEndpointFinder) FindHTTPEndpoints(ctx context.Context, sps []string) ([]peer.AddrInfo, error) {
	addrs := make([]peer.AddrInfo, 0, len(sps))
	for _, sp := range sps {
		addr, ok := ef.endpoints[sp]
		if ok {
			addrs = append(addrs, addr)
		}
	}
	return addrs, nil
}

type fakeLassie struct {
	lastRequest lassietypes.RetrievalRequest
	lsys        *linking.LinkSystem
}

func (fl *fakeLassie) Fetch(ctx context.Context, request lassietypes.RetrievalRequest, opts ...lassietypes.FetchOption) (*lassietypes.RetrievalStats, error) {
	fl.lastRequest = request
	if request.Path != "" {
		return nil, errors.New("Path must be empty")
	}
	carLsys := request.LinkSystem
	unixfsnode.AddUnixFSReificationToLinkSystem(&carLsys)
	carLsys.StorageReadOpener = func(lc linking.LinkContext, l datamodel.Link) (io.Reader, error) {
		r, err := fl.lsys.StorageReadOpener(lc, l)
		if err != nil {
			return nil, err
		}
		data, err := io.ReadAll(r)
		if err != nil {
			return nil, err
		}
		w, wc, err := carLsys.StorageWriteOpener(lc)
		if err != nil {
			return nil, err
		}
		rdr := bytes.NewReader(data)
		if _, err := io.Copy(w, rdr); err != nil {
			return nil, err
		}
		if err := wc(l); err != nil {
			return nil, err
		}
		if _, err := rdr.Seek(0, io.SeekStart); err != nil {
			return nil, err
		}
		return io.NopCloser(rdr), nil
	}
	nd, err := carLsys.Load(linking.LinkContext{Ctx: context.Background()}, cidlink.Link{Cid: request.Root}, dagpb.Type.PBNode)
	if err != nil {
		return nil, err
	}
	progress := traversal.Progress{
		Cfg: &traversal.Config{
			Ctx:                            context.Background(),
			LinkSystem:                     carLsys,
			LinkTargetNodePrototypeChooser: dagpb.AddSupportToChooser(basicnode.Chooser),
		},
	}
	sel, err := selector.CompileSelector(request.GetSelector())
	if err != nil {
		return nil, err
	}
	err = progress.WalkMatching(nd, sel, unixfsnode.BytesConsumingMatcher)
	if err != nil {
		return nil, err
	}
	return &lassietypes.RetrievalStats{}, nil
}
