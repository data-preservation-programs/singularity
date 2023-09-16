package deserializer_test

import (
	"bytes"
	"context"
	"crypto/rand"
	"io"
	"testing"

	"github.com/data-preservation-programs/singularity/retriever/deserializer"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-unixfsnode"
	"github.com/ipfs/go-unixfsnode/testutil"
	"github.com/ipld/go-car/v2"
	"github.com/ipld/go-car/v2/storage"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/linking"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/storage/memstore"
	"github.com/ipld/go-ipld-prime/traversal"
	"github.com/ipld/go-ipld-prime/traversal/selector/builder"
	"github.com/stretchr/testify/require"
)

func TestDeserialize(t *testing.T) {
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
	verifyCarDeserialization(t, &lsys, file, expectedBytes, 0, 4<<20)
	verifyCarDeserialization(t, &lsys, file, expectedBytes, 1, (1<<18)+1)

}

func verifyCarDeserialization(t *testing.T, lsys *linking.LinkSystem, file testutil.DirEntry, expectedBytes []byte, start int64, end int64) {

	fullCarBuff := new(bytes.Buffer)
	fullCarStorage, err := storage.NewWritable(fullCarBuff, []cid.Cid{file.Root}, car.WriteAsCarV1(true))
	require.NoError(t, err)
	carLsys := cidlink.DefaultLinkSystem()
	carLsys.TrustedStorage = true
	carLsys.SetWriteStorage(fullCarStorage)
	unixfsnode.AddUnixFSReificationToLinkSystem(&carLsys)
	carLsys.StorageReadOpener = func(lc linking.LinkContext, l datamodel.Link) (io.Reader, error) {
		r, err := lsys.StorageReadOpener(lc, l)
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
	nd, err := carLsys.Load(linking.LinkContext{Ctx: context.Background()}, cidlink.Link{Cid: file.Root}, dagpb.Type.PBNode)
	require.NoError(t, err)
	ssb := builder.NewSelectorSpecBuilder(basicnode.Prototype.Any)
	sel, err := ssb.ExploreInterpretAs("unixfs", ssb.MatcherSubset(start, end)).Selector()
	require.NoError(t, err)
	progress := traversal.Progress{
		Cfg: &traversal.Config{
			Ctx:                            context.Background(),
			LinkSystem:                     carLsys,
			LinkTargetNodePrototypeChooser: dagpb.AddSupportToChooser(basicnode.Chooser),
		},
	}
	err = progress.WalkMatching(nd, sel, unixfsnode.BytesConsumingMatcher)
	require.NoError(t, err)
	err = fullCarStorage.Finalize()
	require.NoError(t, err)
	carReader, err := car.NewBlockReader(fullCarBuff)
	require.NoError(t, err)
	outBytes := make([]byte, end-start)
	outBuff := bytes.NewBuffer(outBytes)
	written, err := deserializer.Deserialize(context.Background(), carReader, file.Root, start, end, outBuff)
	require.NoError(t, err)
	require.Equal(t, written, end-start)
	require.Equal(t, expectedBytes[start:end], outBytes)
}
