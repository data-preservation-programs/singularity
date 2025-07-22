package handler

import (
	"net/http"
	"strconv"

	"github.com/data-preservation-programs/singularity/service"
	"github.com/filecoin-project/go-address"
	"github.com/labstack/echo/v4"
)

// ProofHandler handles HTTP requests for proof operations
type ProofHandler struct {
	proofService service.ProofService
}

// NewProofHandler creates a new ProofHandler
func NewProofHandler(proofService service.ProofService) *ProofHandler {
	return &ProofHandler{
		proofService: proofService,
	}
}

// RegisterRoutes registers the proof API routes
func (h *ProofHandler) RegisterRoutes(e *echo.Echo) {
	proofs := e.Group("/api/v1/proofs")
	proofs.POST("", h.StoreProof)
	proofs.POST("/:dealId/verify", h.VerifyProof)
	proofs.GET("/:dealId", h.GetProof)
	proofs.GET("/client/:address", h.ListClientProofs)
	proofs.GET("/provider/:address", h.ListProviderProofs)
	proofs.GET("/:dealId/history", h.GetVerificationHistory)
}

// StoreProof handles the request to store a new proof
func (h *ProofHandler) StoreProof(c echo.Context) error {
	var req StoreProofRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	proof, err := req.ToModel()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.proofService.StoreProof(c.Request().Context(), proof); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, ProofResponseFromModel(proof))
}

// VerifyProof handles the request to verify a proof
func (h *ProofHandler) VerifyProof(c echo.Context) error {
	dealID, err := strconv.ParseUint(c.Param("dealId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid deal ID"})
	}

	var req VerifyProofRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	verifierAddr, err := address.NewFromString(req.VerifierAddress)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid verifier address"})
	}

	verification, err := h.proofService.VerifyProof(c.Request().Context(), dealID, verifierAddr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"dealId":            verification.DealID,
		"verifiedBy":        verification.VerifiedBy,
		"verificationResult": verification.VerificationResult,
		"verificationTime":   verification.VerificationTime.UTC().Format("2006-01-02T15:04:05Z"),
	})
}

// GetProof handles the request to get a proof by deal ID
func (h *ProofHandler) GetProof(c echo.Context) error {
	dealID, err := strconv.ParseUint(c.Param("dealId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid deal ID"})
	}

	proof, err := h.proofService.GetProofByDealID(c.Request().Context(), dealID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, ProofResponseFromModel(proof))
}

// ListClientProofs handles the request to list proofs by client address
func (h *ProofHandler) ListClientProofs(c echo.Context) error {
	clientAddr, err := address.NewFromString(c.Param("address"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid client address"})
	}

	proofs, err := h.proofService.ListClientProofs(c.Request().Context(), clientAddr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := make([]*ProofResponse, len(proofs))
	for i, proof := range proofs {
		response[i] = ProofResponseFromModel(proof)
	}

	return c.JSON(http.StatusOK, response)
}

// ListProviderProofs handles the request to list proofs by provider address
func (h *ProofHandler) ListProviderProofs(c echo.Context) error {
	providerAddr, err := address.NewFromString(c.Param("address"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid provider address"})
	}

	proofs, err := h.proofService.ListProviderProofs(c.Request().Context(), providerAddr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := make([]*ProofResponse, len(proofs))
	for i, proof := range proofs {
		response[i] = ProofResponseFromModel(proof)
	}

	return c.JSON(http.StatusOK, response)
}

// GetVerificationHistory handles the request to get verification history for a proof
func (h *ProofHandler) GetVerificationHistory(c echo.Context) error {
	dealID, err := strconv.ParseUint(c.Param("dealId"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid deal ID"})
	}

	history, err := h.proofService.GetProofVerificationHistory(c.Request().Context(), dealID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := make([]*VerificationResponse, len(history))
	for i, v := range history {
		response[i] = &VerificationResponse{
			ID:                v.ID,
			VerifiedBy:        v.VerifiedBy,
			VerificationResult: v.VerificationResult,
			VerificationTime:   v.VerificationTime.UTC().Format("2006-01-02T15:04:05Z"),
		}
	}

	return c.JSON(http.StatusOK, response)
}
