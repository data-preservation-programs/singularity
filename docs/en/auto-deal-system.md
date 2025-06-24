# Singularity Auto-Deal System

> **🚀 Quick Start: Use the [`onboard` command](../../README.md#-auto-deal-system) for complete automated data onboarding**

This document provides technical details for the Singularity Auto-Deal System, which automates storage deal creation when data preparation completes.

## Overview

The Auto-Deal System provides **automated deal creation** as part of the unified data onboarding workflow. Instead of manually managing multiple steps, users can now onboard data from source to storage deals with a single command.

## Primary Interface: `onboard` Command

The main entry point for auto-deal functionality is the unified `onboard` command:

```bash
./singularity onboard \
  --name "my-dataset" \
  --source "/path/to/data" \
  --enable-deals \
  --deal-provider "f01234" \
  --deal-verified \
  --start-workers \
  --wait-for-completion
```

This single command:
1. ✅ Creates storage connections automatically
2. ✅ Sets up data preparation with deal parameters
3. ✅ Starts managed workers to process jobs
4. ✅ Automatically progresses through scan → pack → daggen
5. ✅ Creates storage deals when preparation completes

## System Architecture

The simplified Auto-Deal System consists of two main components:

### 1. **Workflow Orchestrator** (`service/workflow/orchestrator.go`)
- **Event-driven job progression**: scan → pack → daggen → deals
- **Automatic triggering**: No polling, responds to job completion events
- **Integration point**: Called by dataset workers when jobs complete

### 2. **Auto-Deal Trigger Service** (`service/autodeal/trigger.go`)
- **Core auto-deal logic**: Creates deal schedules when preparations are ready
- **Manual overrides**: Supports manual triggering via CLI commands
- **Validation**: Handles wallet and storage provider validation

## Technical Implementation

### Event-Driven Triggering

When a job completes, the workflow orchestrator automatically:

```go
// Job completion triggers workflow progression
func (o *WorkflowOrchestrator) OnJobComplete(ctx context.Context, jobID model.JobID) error {
    // Check job type and trigger next stage
    switch job.Type {
    case model.Scan:
        return o.handleScanCompletion(ctx, db, lotusClient, preparation)
    case model.Pack:
        return o.handlePackCompletion(ctx, db, lotusClient, preparation)
    case model.DagGen:
        return o.handleDagGenCompletion(ctx, db, lotusClient, preparation)
    }
}
```

### Database Schema

The `Preparation` model includes auto-deal configuration:

```go
type Preparation struct {
    // ... existing fields
    
    // Deal configuration (encapsulated in DealConfig struct)
    DealConfig       DealConfig      `gorm:"embedded;embeddedPrefix:deal_config_"`
    DealTemplateID   *DealTemplateID // Optional deal template to use
    WalletValidation bool            // Enable wallet balance validation  
    SPValidation     bool            // Enable storage provider validation
    // ... additional fields
}
```

## Manual Control

For advanced users who need granular control, you can:

```bash
# Monitor preparation status
./singularity prep status <preparation-name>

# Check all deal schedules
./singularity deal schedule list

# Use the unified service for background processing
./singularity run unified --max-workers 10
```

## Configuration Options

### Deal Parameters (via `onboard` command)
- `--deal-provider`: Storage Provider ID (e.g., f01234)
- `--deal-verified`: Whether deals should be verified (default: false)
- `--deal-price-per-gb`: Price in FIL per GiB (default: 0.0)
- `--deal-duration`: Deal duration (default: ~535 days)
- `--deal-start-delay`: Start delay (default: 72h)

### Validation Options
- `--validate-wallet`: Enable wallet balance validation
- `--validate-provider`: Enable storage provider validation

### Worker Management
- `--start-workers`: Start managed workers (default: true)
- `--max-workers`: Maximum number of workers (default: 3)
- `--wait-for-completion`: Monitor until completion

## Advanced Workflow Control

The unified service provides fine-grained control over workflow progression:

```bash
# Run with custom workflow settings
./singularity run unified \
  --disable-auto-deals \
  --disable-pack-to-daggen \
  --max-workers 10
```

## Migration from Complex Multi-Step Approach

**Old approach** (complex, manual):
```bash
# Multiple manual steps
./singularity prep create --auto-create-deals ...
./singularity run dataset-worker --enable-pack &
./singularity run unified
# ... monitor manually
```

**New approach** (simple, automated):
```bash
# Single command
./singularity onboard --name "dataset" --source "/data" --enable-deals --deal-provider "f01234"
```

## Best Practices

1. **Use `onboard` for new workflows** - It provides the simplest and most reliable experience
2. **Enable auto-deal by default** - `--enable-deals` is recommended for most use cases
3. **Set appropriate deal parameters** - Configure provider, pricing, and duration upfront
4. **Use `--wait-for-completion`** - For automated scripts and monitoring
5. **Validate providers and wallets** - Use validation flags for production use

## Troubleshooting

```bash
# Check preparation status
./singularity prep status <preparation-name>

# List all deal schedules
./singularity deal schedule list

# View schedules for specific preparation
curl http://localhost:7005/api/preparation/<name>/schedules
```

For issues with the unified service:
```bash
# Check unified service status
./singularity run unified --dry-run
```

## API Integration

For programmatic access, use the preparation creation API with auto-deal parameters:

```bash
curl -X POST http://localhost:7005/api/preparation \
  -H "Content-Type: application/json" \
  -d '{
    "name": "api-dataset",
    "sourceStorages": ["source-storage"],
    "outputStorages": ["output-storage"],
    "autoCreateDeals": true,
    "dealProvider": "f01234",
    "dealVerified": true,
    "dealPricePerGb": 0.0000001
  }'
```

The auto-deal system will automatically create deal schedules when all jobs complete, providing a seamless integration experience for both CLI and API users.