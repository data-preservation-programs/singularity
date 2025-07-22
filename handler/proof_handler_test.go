package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestHandler(t *testing.T) (*ProofHandler, *echo.Echo) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&model.DealProof{}, &model.ProofVerification{})
	require.NoError(t, err)

	proofService := service.NewProofService(db)
	handler := NewProofHandler(proofService)

	e := echo.New()
	handler.RegisterRoutes(e)

	return handler, e
}

func TestProofHandler_StoreProof(t *testing.T) {
	_, e := setupTestHandler(t)

	// Create test request
	req := StoreProofRequest{
		DealID:          1,
		PieceCID:        "bafk2bzacedbootlz5dgj5migqrjmwjrqygrmyqwqrpxs7x5n5uncgqkqrw75y",
		ProofBytes:      base64.StdEncoding.EncodeToString([]byte("test proof")),
		SectorID:        1,
		ClientAddress:   "f01234",
		ProviderAddress: "f05678",
	}

	body, err := json.Marshal(req)
	require.NoError(t, err)

	// Create HTTP request
	rec := httptest.NewRecorder()
	httpReq := httptest.NewRequest("POST", "/api/v1/proofs", bytes.NewReader(body))
	httpReq.Header.Set("Content-Type", "application/json")

	// Serve the request
	e.ServeHTTP(rec, httpReq)

	require.Equal(t, http.StatusCreated, rec.Code)

	var response ProofResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Equal(t, req.DealID, response.DealID)
	require.Equal(t, model.ProofStatusPending, response.ProofStatus)
}

func TestProofHandler_GetProof(t *testing.T) {
	handler, e := setupTestHandler(t)

	// Store a proof first
	proof := &model.DealProof{
		DealID:          1,
		PieceCID:        "bafk2bzacedbootlz5dgj5migqrjmwjrqygrmyqwqrpxs7x5n5uncgqkqrw75y",
		ProofBytes:      []byte("test proof"),
		SectorID:        1,
		ClientAddress:   "f01234",
		ProviderAddress: "f05678",
	}

	err := handler.proofService.StoreProof(nil, proof)
	require.NoError(t, err)

	// Create HTTP request
	rec := httptest.NewRecorder()
	httpReq := httptest.NewRequest("GET", "/api/v1/proofs/1", nil)

	// Serve the request
	e.ServeHTTP(rec, httpReq)

	require.Equal(t, http.StatusOK, rec.Code)

	var response ProofResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Equal(t, proof.DealID, response.DealID)
	require.Equal(t, proof.PieceCID, response.PieceCID)
}

func TestProofHandler_VerifyProof(t *testing.T) {
	handler, e := setupTestHandler(t)

	// Store a proof first
	proof := &model.DealProof{
		DealID:          1,
		PieceCID:        "bafk2bzacedbootlz5dgj5migqrjmwjrqygrmyqwqrpxs7x5n5uncgqkqrw75y",
		ProofBytes:      []byte("test proof"),
		SectorID:        1,
		ClientAddress:   "f01234",
		ProviderAddress: "f05678",
	}

	err := handler.proofService.StoreProof(nil, proof)
	require.NoError(t, err)

	// Create verification request
	req := VerifyProofRequest{
		VerifierAddress: "f09876",
	}

	body, err := json.Marshal(req)
	require.NoError(t, err)

	// Create HTTP request
	rec := httptest.NewRecorder()
	httpReq := httptest.NewRequest("POST", "/api/v1/proofs/1/verify", bytes.NewReader(body))
	httpReq.Header.Set("Content-Type", "application/json")

	// Serve the request
	e.ServeHTTP(rec, httpReq)

	require.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Equal(t, float64(1), response["dealId"])
	require.Equal(t, true, response["verificationResult"])
}

func TestProofHandler_ListClientProofs(t *testing.T) {
	handler, e := setupTestHandler(t)

	// Store multiple proofs
	clientAddr := "f01234"
	for i := uint64(1); i <= 3; i++ {
		proof := &model.DealProof{
			DealID:          i,
			PieceCID:        "bafk2bzacedbootlz5dgj5migqrjmwjrqygrmyqwqrpxs7x5n5uncgqkqrw75y",
			ProofBytes:      []byte("test proof"),
			SectorID:        i,
			ClientAddress:   clientAddr,
			ProviderAddress: "f05678",
		}
		err := handler.proofService.StoreProof(nil, proof)
		require.NoError(t, err)
	}

	// Create HTTP request
	rec := httptest.NewRecorder()
	httpReq := httptest.NewRequest("GET", "/api/v1/proofs/client/"+clientAddr, nil)

	// Serve the request
	e.ServeHTTP(rec, httpReq)

	require.Equal(t, http.StatusOK, rec.Code)

	var response []ProofResponse
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Len(t, response, 3)
}
