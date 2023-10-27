//nolint:forcetypeassert
package file

import (
	"context"
	"io"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type FilecoinRetriever interface {
	Retrieve(ctx context.Context, c cid.Cid, rangeStart int64, rangeEnd int64, sps []string, out io.Writer) error
	RetrieveReader(ctx context.Context, c cid.Cid, rangeStart int64, rangeEnd int64, sps []string) (io.ReadCloser, error)
}

type Handler interface {
	PrepareToPackFileHandler(
		ctx context.Context,
		db *gorm.DB,
		fileID uint64) (int64, error)

	GetFileDealsHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint64,
	) ([]model.Deal, error)

	GetFileHandler(
		ctx context.Context,
		db *gorm.DB,
		id uint64,
	) (*model.File, error)

	PushFileHandler(
		ctx context.Context,
		db *gorm.DB,
		preparation string,
		source string,
		fileInfo Info,
	) (*model.File, error)

	RetrieveFileHandler(
		ctx context.Context,
		db *gorm.DB,
		retriever FilecoinRetriever,
		id uint64,
	) (data io.ReadSeekCloser, name string, modTime time.Time, err error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

var _ Handler = &MockFile{}

type MockFile struct {
	mock.Mock
}

func (m *MockFile) PrepareToPackFileHandler(ctx context.Context, db *gorm.DB, fileID uint64) (int64, error) {
	args := m.Called(ctx, db, fileID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockFile) PushFileHandler(ctx context.Context, db *gorm.DB, preparation string, source string, fileInfo Info) (*model.File, error) {
	args := m.Called(ctx, db, preparation, source, fileInfo)
	return args.Get(0).(*model.File), args.Error(1)
}

func (m *MockFile) GetFileHandler(ctx context.Context, db *gorm.DB, id uint64) (*model.File, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).(*model.File), args.Error(1)
}

func (m *MockFile) GetFileDealsHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint64,
) ([]model.Deal, error) {
	args := m.Called(ctx, db, id)
	return args.Get(0).([]model.Deal), args.Error(1)
}

func (m *MockFile) RetrieveFileHandler(
	ctx context.Context,
	db *gorm.DB,
	retriever FilecoinRetriever,
	id uint64,
) (data io.ReadSeekCloser, name string, modTime time.Time, err error) {
	args := m.Called(ctx, db, retriever, id)
	return args.Get(0).(io.ReadSeekCloser), args.Get(1).(string), args.Get(2).(time.Time), args.Error(3)
}
