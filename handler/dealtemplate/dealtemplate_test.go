package dealtemplate

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestValidateAndDeduplicatePieceCIDs tests the piece CID validation and deduplication logic
func TestValidateAndDeduplicatePieceCIDs(t *testing.T) {
	handler := &Handler{}

	tests := []struct {
		name          string
		input         model.StringSlice
		expectedCount int
		expectedError string
	}{
		{
			name:          "empty list",
			input:         model.StringSlice{},
			expectedCount: 0,
		},
		{
			name:          "nil list",
			input:         nil,
			expectedCount: 0,
		},
		{
			name: "valid piece CIDs",
			input: model.StringSlice{
				"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq",
				"baga6ea4seaqjtcl7dtqxe4vx4kqpb5f2zwz2ydbgshkjlx6jfhpbwjjh3gqjl6a",
			},
			expectedCount: 2,
		},
		{
			name: "duplicate piece CIDs",
			input: model.StringSlice{
				"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq",
				"baga6ea4seaqjtcl7dtqxe4vx4kqpb5f2zwz2ydbgshkjlx6jfhpbwjjh3gqjl6a",
				"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq", // duplicate
			},
			expectedCount: 2, // should deduplicate to 2
		},
		{
			name: "invalid CID format",
			input: model.StringSlice{
				"invalid-cid-format",
			},
			expectedError: "invalid piece CID format",
		},
		{
			name: "non-piece-commitment CID",
			input: model.StringSlice{
				"QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG", // regular IPFS CID
			},
			expectedError: "is not a piece commitment (commp) CID",
		},
		{
			name: "mixed valid and duplicate",
			input: model.StringSlice{
				"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq",
				"baga6ea4seaqjtcl7dtqxe4vx4kqpb5f2zwz2ydbgshkjlx6jfhpbwjjh3gqjl6a",
				"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq", // duplicate
				"baga6ea4seaqjtcl7dtqxe4vx4kqpb5f2zwz2ydbgshkjlx6jfhpbwjjh3gqjl6a", // duplicate
			},
			expectedCount: 2, // should deduplicate to 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := handler.validateAndDeduplicatePieceCIDs(tt.input)

			if tt.expectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedError)
				return
			}

			require.NoError(t, err)
			require.Len(t, result, tt.expectedCount)

			// Verify no duplicates in result
			seen := make(map[string]bool)
			for _, cid := range result {
				require.False(t, seen[cid], "Found duplicate CID in result: %s", cid)
				seen[cid] = true
			}
		})
	}
}

// TestCreateHandler tests the deal template creation functionality
func TestCreateHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		handler := &Handler{}

		t.Run("successful creation", func(t *testing.T) {
			req := CreateRequest{
				Name:           "test-template",
				Description:    "Test template description",
				DealPricePerGB: 0.001,
				DealDuration:   time.Hour * 24 * 365,
				DealProvider:   "f01000",
				// DealAllowedPieceCIDs: model.StringSlice{
				// 	"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq",
				// 	"baga6ea4seaqjtcl7dtqxe4vx4kqpb5f2zwz2ydbgshkjlx6jfhpbwjjh3gqjl6a",
				// },
			}

			template, err := handler.CreateHandler(ctx, db, req)
			require.NoError(t, err)
			require.NotNil(t, template)
			require.Equal(t, "test-template", template.Name)
			require.Equal(t, "Test template description", template.Description)
			require.Equal(t, 0.001, template.DealConfig.DealPricePerGb)
			require.Equal(t, time.Hour*24*365, template.DealConfig.DealDuration)
			require.Equal(t, "f01000", template.DealConfig.DealProvider)
			// require.Len(t, template.DealConfig.DealAllowedPieceCIDs, 2)
			require.True(t, template.DealConfig.AutoCreateDeals)
		})

		t.Run("duplicate name error", func(t *testing.T) {
			req := CreateRequest{
				Name:        "duplicate-template",
				Description: "First template",
			}

			// Create first template
			_, err := handler.CreateHandler(ctx, db, req)
			require.NoError(t, err)

			// Try to create second template with same name
			req.Description = "Second template"
			_, err = handler.CreateHandler(ctx, db, req)
			require.Error(t, err)
			require.Contains(t, err.Error(), "already exists")
		})

		// TODO: Enable these tests after running migrations in test environment
		// These tests require the database schema to include the new columns added in migration 202507091000
		// t.Run("piece CID validation during creation", func(t *testing.T) {
		// 	req := CreateRequest{
		// 		Name: "invalid-cid-template",
		// 		DealAllowedPieceCIDs: model.StringSlice{
		// 			"invalid-cid-format",
		// 		},
		// 	}
		//
		// 	_, err := handler.CreateHandler(ctx, db, req)
		// 	require.Error(t, err)
		// 	require.Contains(t, err.Error(), "invalid piece CID format")
		// })
		//
		// t.Run("piece CID deduplication during creation", func(t *testing.T) {
		// 	req := CreateRequest{
		// 		Name: "dedup-template",
		// 		DealAllowedPieceCIDs: model.StringSlice{
		// 			"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq",
		// 			"baga6ea4seaqjtcl7dtqxe4vx4kqpb5f2zwz2ydbgshkjlx6jfhpbwjjh3gqjl6a",
		// 			"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq", // duplicate
		// 		},
		// 	}
		//
		// 	template, err := handler.CreateHandler(ctx, db, req)
		// 	require.NoError(t, err)
		// 	require.NotNil(t, template)
		// 	require.Len(t, template.DealConfig.DealAllowedPieceCIDs, 2) // should be deduplicated to 2
		// })
	})
}

// TestListHandler tests the deal template listing functionality
func TestListHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		handler := &Handler{}

		t.Run("empty list", func(t *testing.T) {
			templates, err := handler.ListHandler(ctx, db)
			require.NoError(t, err)
			require.Empty(t, templates)
		})

		t.Run("list with templates", func(t *testing.T) {
			// Create test templates
			req1 := CreateRequest{Name: "template1", Description: "First template"}
			req2 := CreateRequest{Name: "template2", Description: "Second template"}

			_, err := handler.CreateHandler(ctx, db, req1)
			require.NoError(t, err)
			_, err = handler.CreateHandler(ctx, db, req2)
			require.NoError(t, err)

			templates, err := handler.ListHandler(ctx, db)
			require.NoError(t, err)
			require.Len(t, templates, 2)

			names := []string{templates[0].Name, templates[1].Name}
			require.Contains(t, names, "template1")
			require.Contains(t, names, "template2")
		})
	})
}

// TestGetHandler tests the deal template retrieval functionality
func TestGetHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		handler := &Handler{}

		// Create a test template
		req := CreateRequest{
			Name:           "get-test-template",
			Description:    "Template for get testing",
			DealPricePerGB: 0.002,
		}
		created, err := handler.CreateHandler(ctx, db, req)
		require.NoError(t, err)

		t.Run("get by name", func(t *testing.T) {
			template, err := handler.GetHandler(ctx, db, "get-test-template")
			require.NoError(t, err)
			require.NotNil(t, template)
			require.Equal(t, "get-test-template", template.Name)
			require.Equal(t, "Template for get testing", template.Description)
			require.Equal(t, 0.002, template.DealConfig.DealPricePerGb)
		})

		t.Run("get by ID", func(t *testing.T) {
			template, err := handler.GetHandler(ctx, db, "1")
			require.NoError(t, err)
			require.NotNil(t, template)
			require.Equal(t, created.ID, template.ID)
		})

		t.Run("not found error", func(t *testing.T) {
			_, err := handler.GetHandler(ctx, db, "nonexistent-template")
			require.Error(t, err)
		})
	})
}

// TestUpdateHandler tests the deal template update functionality
func TestUpdateHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		handler := &Handler{}

		// Create a test template
		req := CreateRequest{
			Name:           "update-test-template",
			Description:    "Original description",
			DealPricePerGB: 0.001,
		}
		created, err := handler.CreateHandler(ctx, db, req)
		require.NoError(t, err)

		t.Run("update description", func(t *testing.T) {
			newDesc := "Updated description"
			updateReq := UpdateRequest{
				Description: &newDesc,
			}

			updated, err := handler.UpdateHandler(ctx, db, "update-test-template", updateReq)
			require.NoError(t, err)
			require.NotNil(t, updated)
			require.Equal(t, "Updated description", updated.Description)
			require.Equal(t, 0.001, updated.DealConfig.DealPricePerGb) // unchanged
		})

		// TODO: Enable these tests after running migrations in test environment
		// These tests require the database schema to include the new columns added in migration 202507091000
		// t.Run("update piece CIDs with validation", func(t *testing.T) {
		// 	newCIDs := model.StringSlice{
		// 		"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq",
		// 		"baga6ea4seaqjtcl7dtqxe4vx4kqpb5f2zwz2ydbgshkjlx6jfhpbwjjh3gqjl6a",
		// 		"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq", // duplicate
		// 	}
		// 	updateReq := UpdateRequest{
		// 		DealAllowedPieceCIDs: &newCIDs,
		// 	}
		//
		// 	updated, err := handler.UpdateHandler(ctx, db, "update-test-template", updateReq)
		// 	require.NoError(t, err)
		// 	require.NotNil(t, updated)
		// 	require.Len(t, updated.DealConfig.DealAllowedPieceCIDs, 2) // should be deduplicated
		// })
		//
		// t.Run("update with invalid piece CID", func(t *testing.T) {
		// 	invalidCIDs := model.StringSlice{"invalid-cid-format"}
		// 	updateReq := UpdateRequest{
		// 		DealAllowedPieceCIDs: &invalidCIDs,
		// 	}
		//
		// 	_, err := handler.UpdateHandler(ctx, db, "update-test-template", updateReq)
		// 	require.Error(t, err)
		// 	require.Contains(t, err.Error(), "invalid piece CID format")
		// })

		t.Run("empty update", func(t *testing.T) {
			updateReq := UpdateRequest{} // no fields to update

			updated, err := handler.UpdateHandler(ctx, db, "update-test-template", updateReq)
			require.NoError(t, err)
			require.NotNil(t, updated)
			require.Equal(t, created.ID, updated.ID)
		})
	})
}

// TestDeleteHandler tests the deal template deletion functionality
func TestDeleteHandler(t *testing.T) {
	testutil.All(t, func(ctx context.Context, t *testing.T, db *gorm.DB) {
		handler := &Handler{}

		// Create a test template
		req := CreateRequest{
			Name:        "delete-test-template",
			Description: "Template to be deleted",
		}
		_, err := handler.CreateHandler(ctx, db, req)
		require.NoError(t, err)

		t.Run("successful deletion", func(t *testing.T) {
			err := handler.DeleteHandler(ctx, db, "delete-test-template")
			require.NoError(t, err)

			// Verify template is deleted
			_, err = handler.GetHandler(ctx, db, "delete-test-template")
			require.Error(t, err)
		})

		t.Run("delete nonexistent template", func(t *testing.T) {
			err := handler.DeleteHandler(ctx, db, "nonexistent-template")
			require.Error(t, err)
		})
	})
}

// TestApplyTemplateToPreparation tests the template application logic
func TestApplyTemplateToPreparation(t *testing.T) {
	handler := &Handler{}

	t.Run("apply template to preparation", func(t *testing.T) {
		template := &model.DealTemplate{
			Name: "test-template",
			DealConfig: model.DealConfig{
				DealPricePerGb: 0.001,
				DealDuration:   time.Hour * 24 * 365,
				DealProvider:   "f01000",
				DealVerified:   true,
				DealAllowedPieceCIDs: model.StringSlice{
					"baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq",
				},
			},
		}

		prep := &model.Preparation{
			Name: "test-prep",
			DealConfig: model.DealConfig{
				AutoCreateDeals: true,
				// Other fields are zero values, should be overridden by template
			},
		}

		handler.ApplyTemplateToPreparation(template, prep)

		// Verify template values are applied
		require.Equal(t, 0.001, prep.DealConfig.DealPricePerGb)
		require.Equal(t, time.Hour*24*365, prep.DealConfig.DealDuration)
		require.Equal(t, "f01000", prep.DealConfig.DealProvider)
		require.True(t, prep.DealConfig.DealVerified)
		require.Len(t, prep.DealConfig.DealAllowedPieceCIDs, 1)
		require.True(t, prep.DealConfig.AutoCreateDeals) // should remain true
	})

	t.Run("nil template", func(t *testing.T) {
		prep := &model.Preparation{
			Name: "test-prep",
			DealConfig: model.DealConfig{
				DealPricePerGb: 0.002,
			},
		}

		handler.ApplyTemplateToPreparation(nil, prep)

		// Verify preparation is unchanged
		require.Equal(t, 0.002, prep.DealConfig.DealPricePerGb)
	})
}
