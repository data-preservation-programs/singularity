package cmd

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/handler/admin"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func swapAdminHandler(mockHandler admin.Handler) func() {
	actual := admin.Default
	admin.Default = mockHandler
	return func() {
		admin.Default = actual
	}
}

func TestAdminInit(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(MockAdmin)
		defer swapAdminHandler(mockHandler)()
		mockHandler.On("InitHandler", mock.Anything, mock.Anything).Return(nil)
		_, _, err := runner.Run(ctx, "singularity admin init")
		require.NoError(t, err)
	})
}

func TestAdminReset(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(MockAdmin)
		defer swapAdminHandler(mockHandler)()
		mockHandler.On("ResetHandler", mock.Anything, mock.Anything).Return(nil)
		_, _, err := runner.Run(ctx, "singularity admin reset --really-do-it")
		require.NoError(t, err)
	})
}

func TestAdminReset_NoReallyDoIt(t *testing.T) {
	testutil.OneWithoutReset(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		runner := NewRunner()
		defer runner.Save(t)
		mockHandler := new(MockAdmin)
		defer swapAdminHandler(mockHandler)()
		mockHandler.On("ResetHandler", mock.Anything, mock.Anything).Return(nil)
		_, _, err := runner.Run(ctx, "singularity admin reset")
		require.ErrorIs(t, err, cliutil.ErrReallyDoIt)
	})
}
