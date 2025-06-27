package file

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
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

				wallet := &model.Wallet{ActorID: "f01", Address: "f11"}
				err = db.Create(wallet).Error
				require.NoError(t, err)

				deals := make([]model.Deal, 0, 4)
				for i, testCid := range testCids {
					deal := model.Deal{
						State:    model.DealActive,
						PieceCID: model.CID(testCid),
						Provider: "apples" + strconv.Itoa(i),
						Wallet:   wallet,
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
						Wallet:   wallet,
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

				// tinyWriter forces copying through the small buffer created with io.CopyN.
				tinyW := &tinyWriter{
					dst: outBuf,
				}
				_, err = io.CopyN(tinyW, seeker, 1<<20)

				require.NoError(t, err)
				expected := bytes.Join([][]byte{ranges[0].expectedBytes[1<<19 : 1<<20], ranges[1].expectedBytes[0 : 1<<19]}, nil)
				require.Len(t, expected, len(outBuf))
				require.Equal(t, expected, outBuf)
				if !testCase.keepLocalFile {
					require.Len(t, fr.requests, 2)
					require.Equal(t, retrieveRequest{ranges[0].file.Root, 1 << 19, 1 << 20, []string{deals[0].Provider, deals[1].Provider}}, fr.requests[0])
					require.Equal(t, retrieveRequest{ranges[1].file.Root, 0, 1 << 20, []string{deals[0].Provider, deals[1].Provider}}, fr.requests[1])
					fr.requests = nil
				}
				_, err = seeker.Read(outBuf)
				require.NoError(t, err)
				expected = bytes.Join([][]byte{ranges[1].expectedBytes[1<<19 : 1<<20], ranges[2].expectedBytes[0 : 1<<19]}, nil)
				require.Len(t, expected, len(outBuf))
				require.Equal(t, expected, outBuf)
				if !testCase.keepLocalFile {
					require.Len(t, fr.requests, 1) // only one because 1st read got leftover data already in rangeReader
					// The following retrieveRequest is skipped it was included in the previous read.
					//     retrieveRequest{ranges[1].file.Root, 1 << 19, 1 << 20, []string{deals[0].Provider, deals[1].Provider}
					require.Equal(t, retrieveRequest{ranges[2].file.Root, 0, 1 << 20, []string{deals[2].Provider}}, fr.requests[0])
					fr.requests = nil
				}
				_, err = seeker.Read(outBuf)
				require.NoError(t, err)
				expected = bytes.Join([][]byte{ranges[2].expectedBytes[1<<19 : 1<<20], ranges[3].expectedBytes[0 : 1<<19]}, nil)
				require.Equal(t, expected, outBuf)
				if !testCase.keepLocalFile {
					require.Len(t, fr.requests, 1) // only one because 1st read got leftover data already in rangeReader
					// The following retrieveRequest is skipped it was included in the previous read.
					//     retrieveRequest{ranges[2].file.Root, 1 << 19, 1 << 20, []string{deals[2].Provider}}
					require.Equal(t, retrieveRequest{ranges[3].file.Root, 0, 1 << 20, []string{deals[2].Provider}}, fr.requests[0])
					fr.requests = nil
				}
				n, err := seeker.Read(outBuf)
				require.NoError(t, err)
				require.Equal(t, 1<<19, n)
				require.Equal(t, ranges[3].expectedBytes[1<<19:1<<20], outBuf[:n])
				if !testCase.keepLocalFile {
					require.Len(t, fr.requests, 0) // zero because 1st read was got leftover data already in rangeReader
				}
				require.NoError(t, fr.checkDone())

				// Test seeking to 16Kib from end of file, and reading all
				// remaining data. This also tests the seeker's WriteTo
				// function.
				const seekBack = int64(16384)
				_, _ = seeker.Seek(-seekBack, io.SeekEnd)
				buf := bytes.NewBuffer(nil)
				copied, err := io.Copy(buf, seeker)
				require.NoError(t, err)
				require.Equal(t, seekBack, copied)
				require.Equal(t, int(seekBack), buf.Len())

				// Reading again should result in EOF.
				buf.Reset()
				copied, err = io.Copy(buf, seeker)
				require.Zero(t, copied)
				require.ErrorIs(t, err, io.EOF)

				// Check that close does not return error.
				require.NoError(t, seeker.Close())

				// Check that trying to read bad ID gets correct error.
				_, _, _, err = Default.RetrieveFileHandler(ctx, db, fr, uint64(file.ID+37))
				require.ErrorIs(t, err, handlererror.ErrNotFound)
			})
		})
	}
}

func TestUnableToServeRangeError(t *testing.T) {
	origErr := UnableToServeRangeError{
		Start: 0,
		End:   1 << 10,
		Err:   ErrNoJobRecord,
	}

	err := fmt.Errorf("cannot serve data: %w", origErr)
	var uerr UnableToServeRangeError
	require.ErrorAs(t, err, &uerr)

	err = uerr.Unwrap()
	require.ErrorIs(t, err, ErrNoJobRecord)
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
	wg       sync.WaitGroup
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

	// Create pipe and goroutines to simulate Retriever.
	reader, writer := io.Pipe()
	errChan := make(chan error, 2)
	go func() {
		// Simulate deserialize goroutine.
		_, err := io.Copy(out, reader)
		errChan <- err
		_ = reader.Close()
	}()
	go func() {
		// Simulate getContent goroutine.
		_, err := io.Copy(writer, io.LimitReader(nlr, rangeEnd-rangeStart))
		errChan <- err
		_ = writer.Close()
	}()

	// collect errors
	for i := 0; i < 2; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case nextErr := <-errChan:
			if nextErr != nil {
				err = nextErr
			}
		}
	}
	return err
}

func (fr *fakeRetriever) RetrieveReader(ctx context.Context, c cid.Cid, rangeStart int64, rangeEnd int64, sps []string) (io.ReadCloser, error) {
	fr.requests = append(fr.requests, retrieveRequest{c, rangeStart, rangeEnd, sps})

	node, err := fr.lsys.Load(linking.LinkContext{Ctx: ctx}, cidlink.Link{Cid: c}, dagpb.Type.PBNode)
	if err != nil {
		return nil, err
	}
	fnode, err := file.NewUnixFSFile(ctx, node, fr.lsys)
	if err != nil {
		return nil, err
	}
	nlr, err := fnode.AsLargeBytes()
	if err != nil {
		return nil, err
	}
	_, err = nlr.Seek(rangeStart, io.SeekStart)
	if err != nil {
		return nil, err
	}

	fr.wg.Add(2)

	// Create pipe and goroutines to simulate Retriever.
	reader, writer := io.Pipe()
	go func() {
		// Simulate getContent goroutine.
		_, err := io.Copy(writer, io.LimitReader(nlr, rangeEnd-rangeStart))
		writer.CloseWithError(err)
		fr.wg.Done()
	}()
	outReader, outWriter := io.Pipe()
	go func() {
		// Simulate deserialize goroutine.
		_, err := io.Copy(outWriter, reader)
		_ = reader.Close()
		outWriter.CloseWithError(err)
		fr.wg.Done()
	}()

	return outReader, nil
}

// checkDone checks that all pipe goroutines have finished.
func (fr *fakeRetriever) checkDone() error {
	done := make(chan struct{})
	go func() {
		fr.wg.Wait()
		close(done)
	}()
	timer := time.NewTimer(time.Second)
	select {
	case <-done:
		timer.Stop()
	case <-timer.C:
		return errors.New("retrieve goroutines have not exited")
	}
	return nil
}

type tinyWriter struct {
	dst    []byte
	offset int
}

func (w *tinyWriter) Write(p []byte) (int, error) {
	n := copy(w.dst[w.offset:], p)
	w.offset += n
	return n, nil
}

func (w *tinyWriter) reset() {
	w.offset = 0
}

func BenchmarkFilecoinRetrieve(b *testing.B) {
	connStr := "sqlite:" + b.TempDir() + "/singularity.db"
	db, closer, err := database.OpenWithLogger(connStr)
	require.NoError(b, err)
	defer func() { _ = closer.Close() }()
	b.Setenv("DATABASE_CONNECTION_STRING", connStr)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	db = db.WithContext(ctx)
	require.NoError(b, model.GetMigrator(db).Migrate())

	path := b.TempDir()
	lsys := cidlink.DefaultLinkSystem()
	memSys := memstore.Store{
		Bag: make(map[string][]byte),
	}
	lsys.SetReadStorage(&memSys)
	lsys.SetWriteStorage(&memSys)
	lsys.TrustedStorage = true
	ranges := make([]testRange, 0, 4)
	for i := 0; i < 4; i++ {
		expectedByteRange := make([]byte, 1<<24)
		expectedBytesWriter := bytes.NewBuffer(expectedByteRange)
		expectedBytesWriter.Reset()
		fileReader := io.TeeReader(rand.Reader, expectedBytesWriter)
		file := ufstestutil.GenerateFile(b, &lsys, fileReader, 1<<24)
		ranges = append(ranges, testRange{expectedByteRange, file})
	}

	name := "deletedFile.txt"
	file := model.File{
		Path: name,
		Size: 4 << 24,
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
	err = db.Create(&file).Error
	require.NoError(b, err)

	jobs := make([]model.Job, 2)
	for i := 0; i < 2; i++ {
		job := model.Job{
			AttachmentID: file.Attachment.ID,
		}
		err = db.Create(&job).Error
		require.NoError(b, err)
		jobs[i] = job
	}

	for i, testRange := range ranges {
		fileRange := model.FileRange{
			FileID: file.ID,
			CID:    model.CID(testRange.file.Root),
			Offset: int64(i) * (1 << 24),
			Length: 1 << 24,
			JobID:  ptr.Of(jobs[i/2].ID),
		}
		err = db.Create(&fileRange).Error
		require.NoError(b, err)
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
		require.NoError(b, err)
	}

	wallet := &model.Wallet{ActorID: "f01", Address: "f11"}
	err = db.Create(wallet).Error
	require.NoError(b, err)

	for i, testCid := range testCids {
		deal := model.Deal{
			State:    model.DealActive,
			PieceCID: model.CID(testCid),
			Provider: "apples" + strconv.Itoa(i),
			Wallet:   wallet,
		}
		err = db.Create(&deal).Error
		require.NoError(b, err)

		state := model.DealPublished
		if i > 0 {
			state = model.DealProposed
		}
		deal = model.Deal{
			State:    state,
			PieceCID: model.CID(testCid),
			Provider: "oranges" + strconv.Itoa(i),
			Wallet:   wallet,
		}
		err = db.Create(&deal).Error
		require.NoError(b, err)
	}
	fr := &fakeRetriever{
		lsys: &lsys,
	}
	outBuf := make([]byte, 1<<20)

	// tinyWriter forces copying through the small buffer created with io.CopyN.
	tinyW := &tinyWriter{
		dst: outBuf,
	}
	readLen := int64(len(outBuf))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		seeker, _, _, _ := Default.RetrieveFileHandler(ctx, db, fr, uint64(file.ID))
		fr.requests = nil

		// Read the entire file in 1Mib chunks.
		for {
			n, _ := io.CopyN(tinyW, seeker, readLen)
			tinyW.reset()
			if n == 0 {
				break
			}
		}

		_ = seeker.Close()
	}

	b.StopTimer()
	b.Log("Number of retrieve requests:", len(fr.requests))
}
