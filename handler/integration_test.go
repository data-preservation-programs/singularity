package handler

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoreFunctionality(t *testing.T) {
	assert := assert.New(t)
	db := database.OpenInMemory()
	defer model.DropAll(db)
	err := admin.InitHandler(db)
	assert.Nil(err)
}
