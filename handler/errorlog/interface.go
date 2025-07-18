//nolint:forcetypeassert
package errorlog

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service/errorlog"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type Handler interface {
	ListErrorLogsHandler(
		ctx context.Context,
		db *gorm.DB,
		filters errorlog.QueryFilters,
	) ([]model.ErrorLog, int64, error)
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}

type MockErrorLog struct {
	mock.Mock
}

func (m *MockErrorLog) ListErrorLogsHandler(
	ctx context.Context,
	db *gorm.DB,
	filters errorlog.QueryFilters,
) ([]model.ErrorLog, int64, error) {
	args := m.Called(ctx, db, filters)
	return args.Get(0).([]model.ErrorLog), args.Get(1).(int64), args.Error(2)
}

var _ Handler = &MockErrorLog{}