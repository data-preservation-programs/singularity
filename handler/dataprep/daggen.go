package dataprep

import (
	"context"

	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// StartDagGenHandler initiates the start of a Directed Acyclic Graph (DAG) generation job for a given source storage.
//
// This function is a wrapper around the more general `StartJobHandler` function and sets the job type to 'Scan'.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - id: The unique identifier for the desired Preparation record.
// - name: The name of the source storage.
//
// Returns:
// - A pointer to the model.Job record that was initiated.
// - An error, if any occurred during the operation.
func StartDagGenHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	name string) (*model.Job, error) {
	return StartJobHandler(ctx, db, id, name, model.Scan)
}

// @Summary Start a new DAG generation job
// @Tags Job
// @Accept json
// @Produce json
// @Param id path int true "Preparation ID"
// @Param name path string true "Storage name"
// @Success 200 {object} Job
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /preparation/{id}/source/{name}/start-daggen [post]
func _() {}

// PauseDagGenHandler pauses an ongoing Directed Acyclic Graph (DAG) generation job for a given source storage.
//
// This function is a wrapper around the more general `PauseJobHandler` function, specifically for pausing 'Scan' type jobs.
//
// Parameters:
// - ctx: The context for database transactions and other operations.
// - db: A pointer to the gorm.DB instance representing the database connection.
// - id: The unique identifier for the desired Preparation record.
// - name: The name of the source storage.
//
// Returns:
// - A pointer to the model.Job record that was paused.
// - An error, if any occurred during the operation.
func PauseDagGenHandler(
	ctx context.Context,
	db *gorm.DB,
	id uint32,
	name string) (*model.Job, error) {
	return PauseJobHandler(ctx, db, id, name, model.Scan)
}

// @Summary Pause an ongoing DAG generation job
// @Tags Job
// @Accept json
// @Produce json
// @Param id path int true "Preparation ID"
// @Param name path string true "Storage name"
// @Success 200 {object} Job
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /preparation/{id}/source/{name}/pause-daggen [post]
func _() {}
