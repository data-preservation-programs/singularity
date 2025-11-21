package storagesystem

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/config/configmap"
	_ "github.com/rclone/rclone/backend/s3" // Import S3 backend
)

func TestS3Features(t *testing.T) {
	// Create a test S3 configuration
	testConfig := map[string]string{
		"provider": "AWS",
		"region":   "us-east-1",
		"access_key_id": "test",
		"secret_access_key": "test",
		"endpoint": "http://localhost:9000", // MinIO or localstack endpoint
		"chunk_size": "5Mi", // Required for S3
		"upload_cutoff": "5Mi",
		"copy_cutoff": "5Mi",
		"disable_checksum": "true",
		"force_path_style": "true",
	}

	registry, err := fs.Find("s3")
	if err != nil {
		t.Fatalf("Failed to find S3 backend: %v", err)
	}

	ctx := context.Background()
	ctx, _ = fs.AddConfig(ctx)

	// Create S3 filesystem
	s3fs, err := registry.NewFs(ctx, "s3", "test-bucket", configmap.Simple(testConfig))
	if err != nil {
		// This might fail if no S3 endpoint is available, which is OK for feature inspection
		t.Logf("Warning: Failed to create S3 filesystem (this is OK if no S3 endpoint is available): %v", err)
		// Continue anyway to inspect features
	}

	if s3fs != nil {
		features := s3fs.Features()
		t.Logf("S3 Backend Features:")
		t.Logf("  Name: %s", s3fs.Name())
		t.Logf("  Root: %s", s3fs.Root())
		t.Logf("  String: %s", s3fs.String())
		t.Logf("  Precision: %s", s3fs.Precision())
		t.Logf("  Hashes: %v", s3fs.Hashes())

		t.Logf("\nFeature Support:")
		t.Logf("  Move: %v", features.Move != nil)
		t.Logf("  Copy: %v", features.Copy != nil)
		t.Logf("  DirMove: %v", features.DirMove != nil)
		t.Logf("  Purge: %v", features.Purge != nil)
		t.Logf("  PutStream: %v", features.PutStream != nil)
		t.Logf("  About: %v", features.About != nil)
		t.Logf("  ServerSideAcrossConfigs: %v", features.ServerSideAcrossConfigs)
		t.Logf("  CleanUp: %v", features.CleanUp != nil)
		t.Logf("  ListR: %v", features.ListR != nil)
		t.Logf("  SetTier: %v", features.SetTier)
		t.Logf("  GetTier: %v", features.GetTier)

		if features.Move == nil {
			t.Logf("\nNote: S3 backend does not expose Move feature")
		}
		if features.Copy != nil {
			t.Logf("Note: S3 backend does expose Copy feature")
		}
	}
}

func TestS3MoveImplementation(t *testing.T) {
	// Let's also check what rclone operations would do
	ctx := context.Background()

	// Try to create a minimal S3 storage config
	storage := model.Storage{
		Type: "s3",
		Path: "test-bucket",
		Config: map[string]string{
			"provider": "AWS",
			"region": "us-east-1",
			"access_key_id": "test",
			"secret_access_key": "test",
		},
	}

	handler, err := NewRCloneHandler(ctx, storage)
	if err != nil {
		t.Logf("Could not create RCloneHandler for S3: %v", err)
		t.Logf("This is expected if S3 is not configured")
		return
	}

	// Check the features
	if handler.fs != nil {
		features := handler.fs.Features()
		t.Logf("\nRCloneHandler S3 Features via NewRCloneHandler:")
		t.Logf("  Move available: %v", features.Move != nil)
		t.Logf("  Copy available: %v", features.Copy != nil)

		if features.Move == nil {
			t.Logf("\nConfirmed: S3 Move feature is NOT available through RCloneHandler")
			t.Logf("This will cause ErrMoveNotSupported and files will keep UUID names")
		}
	}
}