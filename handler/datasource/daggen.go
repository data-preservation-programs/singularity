package datasource

import (
	"strconv"

	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func DagGenHandler(
	db *gorm.DB,
	id string,
) (*model.Source, error) {
	return dagGenHandler(db, id)
}

// @Summary Mark a source as ready for DAG generation
// @Tags Data Source
// @Produce json
// @Param id path string true "Source ID"
// @Success 200 {object} model.Source
// @Failure 400 {object} handler.HTTPError
// @Failure 500 {object} handler.HTTPError
// @Router /source/{id}/daggen [post]
func dagGenHandler(
	db *gorm.DB,
	id string,
) (*model.Source, error) {
	sourceID, err := strconv.Atoi(id)
	if err != nil {
		return nil, handler.NewBadRequestString("invalid source id")
	}
	var source model.Source
	err = db.Where("id = ?", sourceID).First(&source).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, handler.NewBadRequestString("source not found")
	}
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}

	err = db.Model(&source).Update("dag_gen_state", model.Ready).Error
	if err != nil {
		return nil, handler.NewHandlerError(err)
	}
	source.DagGenState = model.Ready
	return &source, nil
}
