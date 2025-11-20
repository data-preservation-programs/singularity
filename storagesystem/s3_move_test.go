package storagesystem

import (
	"context"
	"strings"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/localstack"
	"github.com/stretchr/testify/require"
)

func TestS3MoveWithCopyDelete(t *testing.T) {
	// Skip if Docker not available
	p := localstack.Preset(
		localstack.WithServices(localstack.S3),
	)
	localS3, err := gnomock.Start(p)
	if err != nil && strings.HasPrefix(err.Error(), "can't start container") {
		t.Skip("Docker required for S3 tests")
	}
	require.NoError(t, err)
	defer func() { _ = gnomock.Stop(localS3) }()

	ctx := context.Background()
	bucketName := "test-bucket"

	// Create S3 storage config
	storage := model.Storage{
		Type: "s3",
		Path: bucketName,
		Config: map[string]string{
			"provider":          "Other",
			"region":            "us-east-1",
			"access_key_id":     "test",
			"secret_access_key": "test",
			"endpoint":          "http://" + localS3.Address(localstack.APIPort),
			"force_path_style":  "true",
			"chunk_size":        "5Mi",
			"upload_cutoff":     "5Mi",
			"copy_cutoff":       "5Mi",
		},
	}

	handler, err := NewRCloneHandler(ctx, storage)
	require.NoError(t, err)

	// Write a test file with UUID name
	testContent := []byte("test content for move")
	uuidName := "12345678-1234-1234-1234-123456789abc.car"

	obj, err := handler.Write(ctx, uuidName, strings.NewReader(string(testContent)))
	require.NoError(t, err)
	require.Equal(t, uuidName, obj.Remote())

	// Now try to move it to piece CID name
	pieceCidName := "baga6ea4seaqmk65y3wzeg277zsfvad5ovnpmlcnch6avl2nltop4m4suxft2moa.car"

	newObj, err := handler.Move(ctx, obj, pieceCidName)
	require.NoError(t, err, "Move should succeed using Copy+Delete fallback")
	require.NotNil(t, newObj)
	require.Equal(t, pieceCidName, newObj.Remote(), "New object should have piece CID name")

	// Verify the original UUID file is gone
	entries := handler.Scan(ctx, "")
	var foundFiles []string
	for entry := range entries {
		if entry.Info != nil {
			foundFiles = append(foundFiles, entry.Info.Remote())
		}
	}

	require.Contains(t, foundFiles, pieceCidName, "Should find piece CID file")
	require.NotContains(t, foundFiles, uuidName, "UUID file should be deleted")

	t.Logf("Successfully moved %s to %s using Copy+Delete", uuidName, pieceCidName)
}