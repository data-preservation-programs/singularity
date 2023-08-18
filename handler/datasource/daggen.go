package datasource

import (
	"context"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func DagGenHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) (*model.Source, error) {
	return dagGenHandler(ctx, db.WithContext(ctx), id)
}

// @Summary Mark a source as ready for DAG generation
// @Tags Data Source
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {object} model.Source
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /source/{id}/daggen [post]
func dagGenHandler(
	ctx context.Context,
	db *gorm.DB,
	id string,
) (*model.Source, error) {
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewInvalidParameterErr("invalid source id")
	}
	var source model.Source
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewInvalidParameterErr("source not found")
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = database.DoRetry(ctx, func() error {
		return db.Model(&source).Update("dag_gen_state", model.Ready).Error
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	source.DagGenState = model.Ready
	return &source, nil
}
