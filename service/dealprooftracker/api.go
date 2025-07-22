package dealprooftracker

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

// RegisterRoutes registers dealprooftracker API endpoints
func RegisterRoutes(r *gin.Engine, tracker *ProofTracker) {
    r.GET("/api/v1/dealproofs/:deal_id", func(c *gin.Context) {
        dealID, err := strconv.ParseUint(c.Param("deal_id"), 10, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deal_id"})
            return
        }
        ctx := c.Request.Context()
        info, err := tracker.GetDBProofInfo(ctx, dealID)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, info)
    })
    r.GET("/api/v1/dealproofs/:deal_id/live", func(c *gin.Context) {
        dealID, err := strconv.ParseUint(c.Param("deal_id"), 10, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid deal_id"})
            return
        }
        ctx := c.Request.Context()
        info, err := tracker.GetLiveProofInfo(ctx, dealID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, info)
    })
}
