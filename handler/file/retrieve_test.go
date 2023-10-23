package file

import (
	"bytes"
	"context"
	"crypto/rand"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	util "github.com/ipfs/boxo/util"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-unixfsnode/file"
	ufstestutil "github.com/ipfs/go-unixfsnode/testutil"
	dagpb "github.com/ipld/go-codec-dagpb"
	"github.com/ipld/go-ipld-prime/linking"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/storage/memstore"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type testRange struct {
	expectedBytes []byte
	file          ufstestutil.DirEntry
}

func TestRetrieveFileHandler(t *testing.T) {
	testCases := []struct {
		name          string
		keepLocalFile bool
	}{
		{
			name:          "from available local file",
			keepLocalFile: true,
		},
		{
			name:          "via filecoin retriever",
			keepLocalFile: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {

				path := t.TempDir()
				lsys := cidlink.DefaultLinkSystem()
				memSys := memstore.Store{
					Bag: make(map[string][]byte),
				}
				lsys.SetReadStorage(&memSys)
				lsys.SetWriteStorage(&memSys)
				lsys.TrustedStorage = true
				ranges := make([]testRange, 0, 4)
				for i := 0; i < 4; i++ {
					expectedByteRange := make([]byte, 1<<20)
					expectedBytesWriter := bytes.NewBuffer(expectedByteRange)
					expectedBytesWriter.Reset()
					fileReader := io.TeeReader(rand.Reader, expectedBytesWriter)
					file := ufstestutil.GenerateFile(t, &lsys, fileReader, 1<<20)
					ranges = append(ranges, testRange{expectedByteRange, file})
				}

				name := "deletedFile.txt"
				if testCase.keepLocalFile {
					osFile, err := os.CreateTemp(path, "push-*")
					require.NoError(t, err)
					for _, testRange := range ranges {
						_, err := osFile.Write(testRange.expectedBytes)
						require.NoError(t, err)
					}
					name = filepath.Base(osFile.Name())
					err = osFile.Close()
					require.NoError(t, err)
				}
				file := model.File{
					Path: name,
					Size: 4 << 20,
					Attachment: &model.SourceAttachment{
						Preparation: &model.Preparation{
							Name: "prep",
						},
						Storage: &model.Storage{
							Name: "source",
							Type: "local",
							Path: path,
						},
					},
				}
				err := db.Create(&file).Error
				require.NoError(t, err)
				jobs := make([]model.Job, 2)
				for i := 0; i < 2; i++ {
					job := model.Job{
						AttachmentID: file.Attachment.ID,
					}
					err = db.Create(&job).Error
					require.NoError(t, err)
					jobs[i] = job
				}

				for i, testRange := range ranges {
					fileRange := model.FileRange{
						FileID: file.ID,
						CID:    model.CID(testRange.file.Root),
						Offset: int64(i) * (1 << 20),
						Length: 1 << 20,
						JobID:  ptr.Of(jobs[i/2].ID),
					}
					err = db.Create(&fileRange).Error
					require.NoError(t, err)
				}

				testCids := make([]cid.Cid, 0, 2)
				for i := 0; i < 2; i++ {
					testCids = append(testCids, cid.NewCidV1(cid.Raw, util.Hash([]byte("test"+strconv.Itoa(i)))))
				}

				for i, job := range jobs {
					car := model.Car{
						JobID:         ptr.Of(job.ID),
						PieceCID:      model.CID(testCids[i]),
						PreparationID: file.Attachment.PreparationID,
					}
					err = db.Create(&car).Error
					require.NoError(t, err)
				}

				deals := make([]model.Deal, 0, 4)
				for i, testCid := range testCids {

					deal := model.Deal{
						State:    model.DealActive,
						PieceCID: model.CID(testCid),
						Provider: "apples" + strconv.Itoa(i),
						Wallet:   &model.Wallet{},
					}
					err = db.Create(&deal).Error
					require.NoError(t, err)

					deals = append(deals, deal)
					state := model.DealPublished
					if i > 0 {
						state = model.DealProposed
					}
					deal = model.Deal{
						State:    state,
						PieceCID: model.CID(testCid),
						Provider: "oranges" + strconv.Itoa(i),
						Wallet:   &model.Wallet{},
					}
					err = db.Create(&deal).Error
					require.NoError(t, err)
					deals = append(deals, deal)
				}
				fr := &fakeRetriever{
					lsys: &lsys,
				}
				seeker, _, _, err := Default.RetrieveFileHandler(ctx, db, fr, uint64(file.ID))
				require.NoError(t, err)
				_, err = seeker.Seek(1<<19, io.SeekStart)
				require.NoError(t, err)
				outBuf := make([]byte, 1<<20)
				_, err = seeker.Read(outBuf)
				require.NoError(t, err)
				expected := bytes.Join([][]byte{ranges[0].expectedBytes[1<<19 : 1<<20], ranges[1].expectedBytes[0 : 1<<19]}, nil)
				require.Equal(t, len(expected), len(outBuf))
				require.Equal(t, expected, outBuf)
				if !testCase.keepLocalFile {
					require.Len(t, fr.requests, 2)
					require.Equal(t, retrieveRequest{ranges[0].file.Root, 1 << 19, 1 << 20, []string{deals[0].Provider, deals[1].Provider}}, fr.requests[0])
					require.Equal(t, retrieveRequest{ranges[1].file.Root, 0, 1 << 19, []string{deals[0].Provider, deals[1].Provider}}, fr.requests[1])
					fr.requests = nil
				}
				_, err = seeker.Read(outBuf)
				require.NoError(t, err)
				expected = bytes.Join([][]byte{ranges[1].expectedBytes[1<<19 : 1<<20], ranges[2].expectedBytes[0 : 1<<19]}, nil)
				require.Equal(t, expected, outBuf)
				if !testCase.keepLocalFile {
					require.Len(t, fr.requests, 2)
					require.Equal(t, retrieveRequest{ranges[1].file.Root, 1 << 19, 1 << 20, []string{deals[0].Provider, deals[1].Provider}}, fr.requests[0])
					require.Equal(t, retrieveRequest{ranges[2].file.Root, 0, 1 << 19, []string{deals[2].Provider}}, fr.requests[1])
					fr.requests = nil
				}
				_, err = seeker.Read(outBuf)
				require.NoError(t, err)
				expected = bytes.Join([][]byte{ranges[2].expectedBytes[1<<19 : 1<<20], ranges[3].expectedBytes[0 : 1<<19]}, nil)
				require.Equal(t, expected, outBuf)
				if !testCase.keepLocalFile {
					require.Len(t, fr.requests, 2)
					require.Equal(t, retrieveRequest{ranges[2].file.Root, 1 << 19, 1 << 20, []string{deals[2].Provider}}, fr.requests[0])
					require.Equal(t, retrieveRequest{ranges[3].file.Root, 0, 1 << 19, []string{deals[2].Provider}}, fr.requests[1])
					fr.requests = nil
				}
				n, err := seeker.Read(outBuf)
				require.NoError(t, err)
				require.Equal(t, 1<<19, n)
				require.Equal(t, ranges[3].expectedBytes[1<<19:1<<20], outBuf[:n])
				if !testCase.keepLocalFile {
					require.Len(t, fr.requests, 1)
					require.Equal(t, retrieveRequest{ranges[3].file.Root, 1 << 19, 1 << 20, []string{deals[2].Provider}}, fr.requests[0])
					fr.requests = nil
				}
			})
		})
	}
}

type retrieveRequest struct {
	c          cid.Cid
	rangeStart int64
	rangeEnd   int64
	sps        []string
}

type fakeRetriever struct {
	requests []retrieveRequest
	lsys     *linking.LinkSystem
}

func (fr *fakeRetriever) Retrieve(ctx context.Context, c cid.Cid, rangeStart int64, rangeEnd int64, sps []string, out io.Writer) error {
	fr.requests = append(fr.requests, retrieveRequest{c, rangeStart, rangeEnd, sps})
	node, err := fr.lsys.Load(linking.LinkContext{Ctx: ctx}, cidlink.Link{Cid: c}, dagpb.Type.PBNode)
	if err != nil {
		return err
	}
	fnode, err := file.NewUnixFSFile(ctx, node, fr.lsys)
	if err != nil {
		return err
	}
	nlr, err := fnode.AsLargeBytes()
	if err != nil {
		return err
	}
	_, err = nlr.Seek(rangeStart, io.SeekStart)
	if err != nil {
		return err
	}
	rangeLeftReader := io.LimitReader(nlr, rangeEnd-rangeStart)
	_, err = io.Copy(out, rangeLeftReader)
	return err
}
