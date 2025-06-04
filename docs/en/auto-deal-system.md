# Singularity Auto-Deal System - Technical Reference

> **📖 For quick start and basic usage, see the [main README](../../README.md#-auto-deal-system)**

This document provides technical details and advanced configuration options for the Singularity Auto-Deal System.

## System Architecture

The Auto-Deal System consists of three main components:

1. **Auto-Deal Trigger Service** (`service/autodeal/trigger.go`) - Core logic for automatic triggering
2. **Job Completion Hooks** - Integrated into dataset worker and pack handler
3. **Auto-Deal Monitor Daemon** (`cmd/run/autodeal.go`) - Background service for batch processing

### Component Integration

```go
// Job completion triggers auto-deal creation
func (w *Thread) handleWorkComplete(ctx context.Context, jobID model.JobID) error {
    // ... update job state to complete
    
    // Trigger auto-deal creation if enabled
    w.triggerAutoDealIfReady(ctx, jobID)
    return nil
}
```

## Technical Configuration

### Database Schema Extensions

The auto-deal system extends the `Preparation` model with additional fields:

```go
type Preparation struct {
    // ... existing fields
    
    // Auto-deal configuration
    AutoCreateDeals     bool          `gorm:"default:false"`
    DealProvider        string        
    DealVerified        bool          `gorm:"default:false"`
    DealPricePerGB      float64       
    DealDuration        time.Duration 
    DealStartDelay      time.Duration `gorm:"default:72h"`
    WalletValidation    bool          `gorm:"default:true"`
    SPValidation        bool          `gorm:"default:true"`
    DealHTTPHeaders     model.ConfigMap
    // ... additional deal parameters
}
```

### Service Interface

The trigger service implements a testable interface:

```go
type AutoDealServiceInterface interface {
    CheckPreparationReadiness(ctx context.Context, db *gorm.DB, preparationID string) (bool, error)
    CreateAutomaticDealSchedule(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, preparationID string) (*model.Schedule, error)
    ProcessReadyPreparations(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient) error
}
```

### Configuration Options Reference

#### Preparation Auto-Deal Settings

| Field | CLI Flag | Type | Description |
|-------|----------|------|-------------|
| `AutoCreateDeals` | `--auto-create-deals` | bool | Enable automatic deal creation |
| `DealProvider` | `--deal-provider` | string | Storage provider ID (required) |
| `DealVerified` | `--deal-verified` | bool | Create verified deals |
| `DealPricePerGB` | `--deal-price-per-gb` | float64 | Price per GB per epoch |
| `DealDuration` | `--deal-duration` | duration | Deal duration |
| `WalletValidation` | `--wallet-validation` | bool | Validate wallets before creation |
| `SPValidation` | `--sp-validation` | bool | Validate storage provider |

#### Monitor Daemon Settings

| Field | CLI Flag | Type | Description |
|-------|----------|------|-------------|
| `CheckInterval` | `--check-interval` | duration | Polling frequency |
| `EnableBatchMode` | `--enable-batch-mode` | bool | Enable batch processing |
| `EnableJobHooks` | `--enable-job-hooks` | bool | Enable job completion triggers |
| `MaxRetries` | `--max-retries` | int | Retry limit before backoff |

## Implementation Details

### Trigger Logic Flow

```go
func (s *TriggerService) TriggerForJobCompletion(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, jobID model.JobID) error {
    // 1. Get job and check if preparation has auto-deal enabled
    // 2. Check if all jobs for preparation are complete
    // 3. Check if deal schedule already exists (avoid duplicates)
    // 4. Create automatic deal schedule
    // 5. Log success/failure notifications
}
```

### Readiness Detection

The system determines preparation readiness by:

```sql
SELECT COUNT(*) FROM jobs 
JOIN source_attachments ON jobs.attachment_id = source_attachments.id 
WHERE source_attachments.preparation_id = ? AND jobs.state != 'Complete'
```

If count is 0, the preparation is ready for deal creation.

### Error Handling Strategy

1. **Network Failures**: Exponential backoff with configurable max retries
2. **Validation Failures**: Log warnings but continue processing
3. **Missing Dependencies**: Skip with detailed notifications
4. **Context Cancellation**: Clean shutdown without data loss

### Notification System Integration

All auto-deal events are stored in the notifications table:

```go
type Notification struct {
    ID       uint            `gorm:"primaryKey"`
    Source   string          // "auto-deal-service"
    Title    string          // Event title
    Message  string          // Detailed message
    Level    string          // "info", "warning", "error"
    Metadata model.ConfigMap // Structured data
}
```

## Advanced Usage Patterns

### Custom Validation Logic

Disable built-in validation for custom workflows:

```bash
singularity prep create \
  --name "custom-workflow" \
  --auto-create-deals \
  --wallet-validation=false \
  --sp-validation=false \
  --deal-provider "f01234"
```

### High-Frequency Monitoring

For time-sensitive workflows:

```bash
singularity run autodeal \
  --check-interval=5s \
  --max-retries=1 \
  --exit-on-error
```

### Testing and Development

```bash
# Test auto-deal creation without side effects
singularity prep autodeal create --preparation "test" --dry-run

# Enable debug logging
singularity run autodeal --log-level=debug
```

## Validation System Details

### Wallet Validation Implementation

```go
type BalanceValidator struct {
    lotusClient jsonrpc.RPCClient
}

func (v *BalanceValidator) ValidateWalletExists(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, address string, context string) (*ValidationResult, error) {
    // 1. Convert address format (f1 ↔ t1) if needed
    // 2. Check wallet exists on-chain
    // 3. Validate balance if thresholds configured
    // 4. Log validation results to notifications
}
```

### Storage Provider Validation

```go
type SPValidator struct {
    lotusClient jsonrpc.RPCClient
}

func (v *SPValidator) ValidateStorageProvider(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, providerID string, context string) (*ValidationResult, error) {
    // 1. Look up provider on-chain using StateLookupID
    // 2. Check provider power and status
    // 3. Verify provider is accepting deals
    // 4. Fall back to default provider if validation fails
}
```

### Error Recovery Mechanisms

| Scenario | Recovery Strategy | Implementation |
|----------|------------------|----------------|
| Lotus API timeout | Retry with exponential backoff | `database.DoRetry()` wrapper |
| Invalid provider ID | Use default provider from config | `GetDefaultStorageProvider()` |
| Insufficient balance | Log warning, skip deal creation | Notification with metadata |
| Database constraint violation | Rollback transaction, log error | GORM transaction handling |

## Monitoring and Observability

### Structured Logging

The auto-deal system uses structured logging with context:

```go
logger.Infof("Successfully created automatic deal schedule %d for preparation %s", 
    schedule.ID, preparation.Name)

// With metadata
logger.With("preparation_id", prepID, "schedule_id", scheduleID).
    Info("Auto-deal schedule created")
```

### Notification Categories

| Source | Title Pattern | Level | Purpose |
|--------|---------------|-------|---------|
| `auto-deal-service` | `Auto-Deal Schedule Created` | info | Success notifications |
| `auto-deal-service` | `Wallet Validation Failed` | warning | Validation issues |
| `auto-deal-service` | `Deal Schedule Creation Failed` | error | Critical failures |

### Metrics Collection

Key metrics to monitor:

```go
// Track processing rates
autodeal_preparations_processed_total
autodeal_schedules_created_total
autodeal_validation_failures_total

// Monitor performance
autodeal_processing_duration_seconds
autodeal_readiness_check_duration_seconds
```

### Health Checks

The daemon provides health check endpoints:

```bash
# Check daemon status
curl http://localhost:8080/health/autodeal

# View processing statistics
curl http://localhost:8080/metrics/autodeal
```

## Performance Tuning

### Database Optimization

For high-throughput scenarios:

```sql
-- Index for readiness checks
CREATE INDEX idx_jobs_attachment_state ON jobs(attachment_id, state);

-- Index for auto-deal preparations
CREATE INDEX idx_preparations_autodeal ON preparations(auto_create_deals) WHERE auto_create_deals = true;
```

### Memory Management

Configure worker memory limits:

```bash
# Limit concurrent processing
singularity run autodeal --max-concurrent-preparations=5

# Adjust check intervals based on load
singularity run autodeal --check-interval=60s  # Lower frequency for high load
```

### Network Optimization

For remote Lotus nodes:

```bash
# Increase timeouts for slow networks
export LOTUS_API_TIMEOUT=30s

# Use connection pooling
export LOTUS_MAX_CONNECTIONS=10
```

### Batch Processing Optimization

```go
// Process preparations in batches to reduce database load
const BatchSize = 50

func (s *TriggerService) ProcessReadyPreparations(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient) error {
    offset := 0
    for {
        preparations := s.getReadyPreparations(db, offset, BatchSize)
        if len(preparations) == 0 {
            break
        }
        s.processBatch(preparations)
        offset += BatchSize
    }
}
```

## API Reference

### Auto-Deal Service Interface

```go
type TriggerService struct {
    autoDealService AutoDealServiceInterface
    mutex           sync.RWMutex
    enabled         bool
}

// Core methods
func (s *TriggerService) TriggerForJobCompletion(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, jobID model.JobID) error
func (s *TriggerService) TriggerForPreparation(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient, preparationID string) error
func (s *TriggerService) BatchProcessReadyPreparations(ctx context.Context, db *gorm.DB, lotusClient jsonrpc.RPCClient) error

// Configuration
func (s *TriggerService) SetEnabled(enabled bool)
func (s *TriggerService) IsEnabled() bool
```

### Monitor Service Configuration

```go
type MonitorConfig struct {
    CheckInterval     time.Duration // How often to check for ready preparations
    EnableBatchMode   bool          // Enable batch processing mode
    ExitOnComplete    bool          // Exit when no work found
    ExitOnError       bool          // Exit on any error
    MaxRetries        int           // Max retries before backing off
    RetryInterval     time.Duration // Base interval for retry backoff
}

func DefaultMonitorConfig() MonitorConfig {
    return MonitorConfig{
        CheckInterval:   30 * time.Second,
        EnableBatchMode: true,
        ExitOnComplete:  false,
        ExitOnError:     false,
        MaxRetries:      3,
        RetryInterval:   5 * time.Minute,
    }
}
```

## Troubleshooting Guide

### Diagnostic Commands

```bash
# Test trigger service functionality
go test ./service/autodeal/ -v -run TestTriggerService

# Check database connectivity
singularity admin notification list --limit 1

# Verify Lotus connectivity
singularity wallet list

# Test preparation readiness detection
singularity prep autodeal check --preparation "test-prep"
```

### Common Error Patterns

**"preparation not found" errors:**
```bash
# Check preparation exists and ID is correct
singularity prep list --name "my-prep"
```

**"no wallet attached" errors:**
```bash
# Verify wallet attachment
singularity prep list-wallets "my-prep"

# Attach wallet if missing
singularity prep attach-wallet "my-prep" "f3wallet"
```

**"provider cannot be resolved" errors:**
```bash
# Check provider ID format
echo "f01234" | grep -E "^f0[0-9]+$"

# Test provider on-chain
lotus state lookup-id f01234
```

### Performance Debugging

```bash
# Monitor database query performance
export DATABASE_LOG_LEVEL=debug

# Profile memory usage
go tool pprof http://localhost:6060/debug/pprof/heap

# Check goroutine leaks
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

### Log Analysis

```bash
# Extract auto-deal timeline
grep "autodeal-trigger" singularity.log | cut -d' ' -f1,2,5- | sort

# Count success/failure rates
grep "Auto-Deal Schedule Created Successfully" singularity.log | wc -l
grep "Auto-Deal Creation Failed" singularity.log | wc -l

# Identify slow operations
grep "processing_duration" singularity.log | awk '{print $NF}' | sort -n
```

## Testing and Development

### Unit Testing

```bash
# Run all auto-deal tests
go test ./service/autodeal/ -v

# Run specific test patterns
go test ./service/autodeal/ -v -run "TestTrigger.*Success"

# Test with race detection
go test ./service/autodeal/ -race -v
```

### Integration Testing

```bash
# Test full workflow with mocks
go test ./service/autodeal/ -v -run "TestTriggerIntegration"

# Test error scenarios
go test ./service/autodeal/ -v -run "TestErrorScenarios"
```

### Development Setup

```bash
# Build with debug symbols
go build -gcflags="all=-N -l" -o singularity-debug .

# Run with debug logging
./singularity-debug run autodeal --log-level=debug

# Enable pprof for profiling
go run . run autodeal --enable-pprof &
```

### Mock Development

```go
// Create mock auto-deal service for testing
mockService := &MockAutoDealer{}
triggerService.SetAutoDealService(mockService)

// Set up expectations
mockService.On("CheckPreparationReadiness", mock.Anything, mock.Anything, "1").Return(true, nil)
mockService.On("CreateAutomaticDealSchedule", mock.Anything, mock.Anything, mock.Anything, "1").Return(&model.Schedule{ID: 1}, nil)
```

## Production Deployment Guide

### Best Practices

1. **Resource Planning**
   - Allocate sufficient database connections for concurrent processing
   - Monitor disk space for CAR file storage
   - Plan network bandwidth for Lotus API calls

2. **Security Considerations**
   - Store wallet private keys securely (never in auto-deal config)
   - Use read-only Lotus tokens where possible
   - Implement proper access controls for auto-deal daemon

3. **Reliability Patterns**
   - Run auto-deal daemon with process supervisor (systemd, supervisord)
   - Implement health checks and alerting
   - Use database backups and transaction logging

### Production Configuration

```bash
# Production daemon configuration
singularity run autodeal \
  --check-interval=60s \
  --max-retries=5 \
  --retry-interval=10m \
  --enable-batch-mode \
  --enable-job-hooks

# With systemd service
[Unit]
Description=Singularity Auto-Deal Daemon
After=network.target

[Service]
Type=simple
User=singularity
WorkingDirectory=/opt/singularity
ExecStart=/opt/singularity/bin/singularity run autodeal
Restart=always
RestartSec=30

[Install]
WantedBy=multi-user.target
```

### Monitoring Integration

```yaml
# Prometheus scrape config
- job_name: 'singularity-autodeal'
  static_configs:
    - targets: ['localhost:8080']
  metrics_path: '/metrics/autodeal'
  scrape_interval: 30s
```

### Capacity Planning

```bash
# Estimate resource requirements
echo "Preparations with auto-deal: $(singularity prep list --auto-deal-only | wc -l)"
echo "Average jobs per preparation: $(sqlite3 singularity.db 'SELECT AVG(job_count) FROM (SELECT COUNT(*) as job_count FROM jobs GROUP BY attachment_id)')"
echo "Database size: $(du -h singularity.db)"
```