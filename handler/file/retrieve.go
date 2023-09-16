package file

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var logger = log.Logger("singularity/handler/file")

var ErrNoFileRangeRecord = errors.New("missing file range record")
var ErrNoJobRecord = errors.New("missing job record")
var ErrNoFilecoinDeals = errors.New("no filecoin deals available")
var ErrByteOffsetBeyondFile = errors.New("byte offset would exceed file size")

type UnableToServeRangeError struct {
	Start int64
	End   int64
	Err   error
}

func (e UnableToServeRangeError) Unwrap() error {
	return e.Err
}

func (e UnableToServeRangeError) Error() string {
	return fmt.Sprintf("unable to serve byte range %d to %d: %s", e.Start, e.End, e.Err.Error())
}

// RetrieveFileHandler retrieves the actual bytes for a file on disk using a given file ID.
//
// # For now, this function only works if the file remains available in its original source storage
//
// Parameters:
// - ctx: The context for managing timeouts and cancellation.
// - db: The gorm.DB instance for database operations.
// - id: The ID of the file to be retrieved.
//
// Returns:
// - A ReadSeekCloser for the given file
// - the name of the file
// - An error if any issues occur during the database operation, including when the file is not found.
func (DefaultHandler) RetrieveFileHandler(
	ctx context.Context,
	db *gorm.DB,
	filecoinRetriever FilecoinRetriever,
	id uint64,
) (data io.ReadSeekCloser, name string, modTime time.Time, err error) {
	db = db.WithContext(ctx)
	var file model.File
	err = db.Preload("Attachment.Storage").First(&file, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, "", time.Time{}, errors.Wrapf(handlererror.ErrNotFound, "file '%d' does not exist", id)
	}
	if err != nil {
		return nil, "", time.Time{}, errors.WithStack(err)
	}

	rclone, err := storagesystem.NewRCloneHandler(ctx, *file.Attachment.Storage)
	if err != nil {
		return nil, file.FileName(), time.Unix(0, file.LastModifiedNano), errors.WithStack(err)
	}

	seeker, obj, err := storagesystem.Open(rclone, ctx, file.Path)
	if err != nil {
		// we have no local copy, so we have to return a filecoin based reader
		return &filecoinReader{
			ctx:       ctx,
			db:        db,
			retriever: filecoinRetriever,
			size:      file.Size,
			id:        id,
		}, file.FileName(), time.Unix(0, file.LastModifiedNano), nil
	}

	return seeker, file.FileName(), obj.ModTime(ctx), nil
}

// io.ReadSeekCloser implementation that reads from remote singularity
type filecoinReader struct {
	ctx       context.Context
	db        *gorm.DB
	retriever FilecoinRetriever
	offset    int64
	size      int64
	id        uint64
}

func (r *filecoinReader) Read(p []byte) (int, error) {
	logger.Infof("buffer size: %v", len(p))

	buf := bytes.NewBuffer(p)
	buf.Reset()

	if r.offset >= r.size {
		return 0, io.EOF
	}

	// Figure out how many bytes to read
	readLen := int64(len(p))
	remainingBytes := r.size - r.offset
	if readLen > remainingBytes {
		readLen = remainingBytes
	}

	fileRanges, err := findFileRanges(r.db, r.id, r.offset, r.offset+readLen)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve file range deals: %w", err)
	}

	read := 0
	for _, fileRange := range fileRanges {
		if readLen == 0 {
			// this shouldn't happen
			logger.Warnw("retrieval reader retrieved file ranges beyond end of range", "fileRangeStart", fileRange.Offset, "fileRangeEnd", fileRange.Offset+fileRange.Length)
			break
		}
		if fileRange.Offset > r.offset {
			return read, UnableToServeRangeError{Start: r.offset, End: fileRange.Offset, Err: ErrNoFileRangeRecord}
		}
		rangeReadLen := readLen
		remainingRange := (fileRange.Offset + fileRange.Length) - r.offset
		if rangeReadLen > remainingRange {
			rangeReadLen = remainingRange
		}
		if fileRange.JobID == nil {
			return read, UnableToServeRangeError{Start: r.offset, End: r.offset + rangeReadLen, Err: ErrNoJobRecord}
		}
		providers, err := findProviders(r.db, *fileRange.JobID)
		if err != nil || len(providers) == 0 {
			return read, UnableToServeRangeError{Start: r.offset, End: r.offset + rangeReadLen, Err: ErrNoFilecoinDeals}
		}
		err = r.retriever.Retrieve(r.ctx, cid.Cid(fileRange.CID), r.offset, r.offset+rangeReadLen, providers, buf)
		if err != nil {
			return read, UnableToServeRangeError{
				Start: r.offset,
				End:   r.offset + rangeReadLen,
				Err:   fmt.Errorf("unable to retrieve data from filecoin: %w", err),
			}
		}
		r.offset += rangeReadLen
		readLen -= rangeReadLen
		read += int(rangeReadLen)
	}

	// check for missing file ranges at the end
	if readLen > 0 {
		return read, UnableToServeRangeError{Start: r.offset, End: r.offset + readLen, Err: ErrNoFileRangeRecord}
	}
	return read, nil
}

func (r *filecoinReader) Seek(offset int64, whence int) (int64, error) {
	var newOffset int64

	switch whence {
	case io.SeekStart:
		newOffset = offset
	case io.SeekCurrent:
		newOffset = r.offset + offset
	case io.SeekEnd:
		newOffset = r.size + offset
	}

	if newOffset > r.size {
		return 0, ErrByteOffsetBeyondFile
	}

	r.offset = newOffset

	return r.offset, nil
}

func (r *filecoinReader) Close() error {
	return nil
}

func findFileRanges(db *gorm.DB, id uint64, startRange int64, endRange int64) ([]model.FileRange, error) {
	var fileRanges []model.FileRange
	err := db.Model(&model.FileRange{}).Where("file_ranges.file_id = ? AND file_ranges.offset < ? AND (file_ranges.offset+file_ranges.length) > ?", id, endRange, startRange).
		Order("file_ranges.offset ASC").Find(&fileRanges).Error
	if err != nil {
		return nil, err
	}
	return fileRanges, nil
}

type deal struct {
	Provider string
}

func findProviders(db *gorm.DB, jobID model.JobID) ([]string, error) {
	var deals []deal
	err := db.Table("deals").Select("provider").
		Joins("JOIN cars ON deals.piece_cid = cars.piece_cid").
		Where("cars.job_id = ? and deals.state IN (?)", jobID, []model.DealState{
			model.DealPublished,
			model.DealActive,
		}).Find(&deals).Error
	if err != nil {
		return nil, err
	}
	providers := make([]string, 0, len(deals))
	for _, deal := range deals {
		providers = append(providers, deal.Provider)
	}
	return providers, nil
}
