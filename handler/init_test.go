package handler

import (
	"github.com/data-preservation-programs/go-singularity/database"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitHandler(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	err := InitHandler(db)
	assert.NoError(err)
}
