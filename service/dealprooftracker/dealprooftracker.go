package dealprooftracker

import (
    "context"
    "log"
    "time"
    "gorm.io/gorm"
    "github.com/data-preservation-programs/singularity/lotusclient"
)

type ProofTracker struct {
    db         *gorm.DB
    lotusURL   string
    lotusToken string
    interval   time.Duration
    stopCh     chan struct{}
}

func NewProofTracker(db *gorm.DB, lotusURL, lotusToken string, interval time.Duration) *ProofTracker {
    return &ProofTracker{
        db:         db,
        lotusURL:   lotusURL,
        lotusToken: lotusToken,
        interval:   interval,
        stopCh:     make(chan struct{}),
    }
}

// Start runs the tracker loop
func (pt *ProofTracker) Start(ctx context.Context) {
    ticker := time.NewTicker(pt.interval)
    defer ticker.Stop()
    for {
        select {
        case <-ticker.C:
            pt.pollAndUpdate(ctx)
        case <-pt.stopCh:
            return
        case <-ctx.Done():
            return
        }
    }
}

// pollAndUpdate queries Lotus and updates DB for tracked deals (MVP: single deal example)
func (pt *ProofTracker) pollAndUpdate(ctx context.Context) {
    log.Println("[ProofTracker] Polling Lotus and updating DB...")
    lotus := lotusclient.NewLotusClient(pt.lotusURL, pt.lotusToken)

    // Example: fetch a list of deal IDs to track (MVP: hardcoded or from DB)
    var trackedDeals []uint64
    // TODO: Replace with DB query for all deals to track
    trackedDeals = []uint64{113718925} // Example deal

    for _, dealID := range trackedDeals {
        var dealInfo map[string]interface{}
        err := lotus.CallFor(ctx, &dealInfo, "Filecoin.StateMarketStorageDeal", dealID, nil)
        if err != nil {
            log.Printf("[ProofTracker] Failed to get deal info for %d: %v", dealID, err)
            continue
        }
        state := dealInfo["State"].(map[string]interface{})
        sectorNumber := int64(state["SectorNumber"].(float64))
        sectorStartEpoch := int32(state["SectorStartEpoch"].(float64))

        proposal := dealInfo["Proposal"].(map[string]interface{})
        provider := proposal["Provider"].(string)

        // Get proving deadline info
        var deadlineInfo map[string]interface{}
        err = lotus.CallFor(ctx, &deadlineInfo, "Filecoin.StateMinerProvingDeadline", provider, nil)
        if err != nil {
            log.Printf("[ProofTracker] Failed to get proving deadline for %s: %v", provider, err)
            continue
        }
        currentDeadline := int32(deadlineInfo["Index"].(float64))
        periodStartEpoch := int32(deadlineInfo["PeriodStart"].(float64))

        // Estimate next proof time (MVP: now + 30min)
        estimatedNextProofTime := time.Now().Add(30 * time.Minute)

        // Get partition info (faults/recoveries)
        var partitions []map[string]interface{}
        err = lotus.CallFor(ctx, &partitions, "Filecoin.StateMinerPartitions", provider, currentDeadline, nil)
        faults := 0
        recoveries := 0
        if err == nil && len(partitions) > 0 {
            for _, p := range partitions {
                if f, ok := p["FaultySectors"].(float64); ok {
                    faults += int(f)
                }
                if r, ok := p["RecoveringSectors"].(float64); ok {
                    recoveries += int(r)
                }
            }
        }

        // Upsert into DB
        tracking := DealProofTracking{
            DealID:                dealID,
            Provider:              provider,
            SectorID:              sectorNumber,
            SectorStartEpoch:      sectorStartEpoch,
            CurrentDeadlineIndex:  currentDeadline,
            PeriodStartEpoch:      periodStartEpoch,
            EstimatedNextProofTime: estimatedNextProofTime,
            Faults:                faults,
            Recoveries:            recoveries,
            LastUpdatedAt:         time.Now(),
        }
        err = pt.db.WithContext(ctx).Save(&tracking).Error
        if err != nil {
            log.Printf("[ProofTracker] Failed to upsert deal proof tracking for %d: %v", dealID, err)
        } else {
            log.Printf("[ProofTracker] Updated deal proof tracking for %d", dealID)
        }
    }
}

// (Removed duplicate Stop() method)

// GetLiveProofInfo queries Lotus for real-time proof info for a deal
func (pt *ProofTracker) GetLiveProofInfo(ctx context.Context, dealID uint64) (*DealProofTracking, error) {
    lotus := NewLotusClient(pt.lotusURL, pt.lotusToken)
    var dealInfo map[string]interface{}
    err := lotus.CallFor(ctx, &dealInfo, "Filecoin.StateMarketStorageDeal", dealID, nil)
    if err != nil {
        return nil, err
    }
    state := dealInfo["State"].(map[string]interface{})
    sectorNumber := int64(state["SectorNumber"].(float64))
    sectorStartEpoch := int32(state["SectorStartEpoch"].(float64))
    proposal := dealInfo["Proposal"].(map[string]interface{})
    provider := proposal["Provider"].(string)
    var deadlineInfo map[string]interface{}
    err = lotus.CallFor(ctx, &deadlineInfo, "Filecoin.StateMinerProvingDeadline", provider, nil)
    if err != nil {
        return nil, err
    }
    currentDeadline := int32(deadlineInfo["Index"].(float64))
    periodStartEpoch := int32(deadlineInfo["PeriodStart"].(float64))
    estimatedNextProofTime := time.Now().Add(30 * time.Minute)
    var partitions []map[string]interface{}
    err = lotus.CallFor(ctx, &partitions, "Filecoin.StateMinerPartitions", provider, currentDeadline, nil)
    faults := 0
    recoveries := 0
    if err == nil && len(partitions) > 0 {
        for _, p := range partitions {
            if f, ok := p["FaultySectors"].(float64); ok {
                faults += int(f)
            }
            if r, ok := p["RecoveringSectors"].(float64); ok {
                recoveries += int(r)
            }
        }
    }
    tracking := &DealProofTracking{
        DealID:                dealID,
        Provider:              provider,
        SectorID:              sectorNumber,
        SectorStartEpoch:      sectorStartEpoch,
        CurrentDeadlineIndex:  currentDeadline,
        PeriodStartEpoch:      periodStartEpoch,
        EstimatedNextProofTime: estimatedNextProofTime,
        Faults:                faults,
        Recoveries:            recoveries,
        LastUpdatedAt:         time.Now(),
    }
    return tracking, nil
}

// GetDBProofInfo fetches proof info from the DB for a deal
func (pt *ProofTracker) GetDBProofInfo(ctx context.Context, dealID uint64) (*DealProofTracking, error) {
    var tracking DealProofTracking
    err := pt.db.WithContext(ctx).First(&tracking, "deal_id = ?", dealID).Error
    if err != nil {
        return nil, err
    }
    return &tracking, nil
}

// HealthCheck returns nil if the tracker can reach Lotus and DB
func (pt *ProofTracker) HealthCheck(ctx context.Context) error {
    // Check DB
    if err := pt.db.WithContext(ctx).Exec("SELECT 1").Error; err != nil {
        return err
    }
    // Check Lotus
    lotus := lotusclient.NewLotusClient(pt.lotusURL, pt.lotusToken)
    var chainHead map[string]interface{}
    err := lotus.CallFor(ctx, &chainHead, "Filecoin.ChainHead")
    if err != nil {
        return err
    }
    return nil
}
// (Removed stray code block that was outside any function)

// Stop signals the tracker to stop
func (pt *ProofTracker) Stop() {
    close(pt.stopCh)
}
