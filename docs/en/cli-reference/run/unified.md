# Start the unified auto-preparation service

> **💡 Alternative:** For most users, the [`onboard` command](../onboard.md) provides a simpler single-command workflow. Use this service for continuous background processing.

The unified service combines worker management and workflow orchestration into a single background service that automatically processes data preparation jobs and creates deals.

{% code fullWidth="true" %}
```
NAME:
   singularity run unified - Start the unified auto-preparation service with worker management and workflow orchestration

USAGE:
   singularity run unified [command options]

DESCRIPTION:
   The unified service provides automated data preparation processing:
   
   1. Worker Management: Auto-scaling workers that process scan, pack, and daggen jobs
   2. Workflow Orchestration: Event-driven progression through job stages
   3. Auto-Deal Creation: Automatic deal schedule creation when preparations complete
   
   This is the recommended way to run fully automated data preparation.

OPTIONS:
   Worker Management:
   --min-workers value           Minimum number of workers to keep running (default: 1)
   --max-workers value           Maximum number of workers to run (default: 5)
   --scale-up-threshold value    Number of ready jobs to trigger worker scale-up (default: 5)
   --scale-down-threshold value  Number of ready jobs below which to scale down (default: 2)
   --worker-idle-timeout value   How long a worker can be idle before shutdown (default: 5m0s)
   --disable-auto-scaling        Disable automatic worker scaling (default: false)

   Service Configuration:
   --check-interval value        How often to check for work and status (default: 30s)

   Workflow Control:
   --disable-workflow-orchestration  Disable automatic job progression (default: false)
   --disable-auto-deals             Disable automatic deal creation (default: false)
   --disable-scan-to-pack           Disable automatic scan → pack transitions (default: false)
   --disable-pack-to-daggen         Disable automatic pack → daggen transitions (default: false)
   --disable-daggen-to-deals        Disable automatic daggen → deals transitions (default: false)

   --help, -h                    show help
```
{% endcode %}

## Basic Usage

### Start unified service with defaults

```bash
# Start with default settings (1-5 workers, full automation)
singularity run unified
```

### Production configuration

```bash
# High-capacity production setup
singularity run unified \
  --min-workers 3 \
  --max-workers 20 \
  --scale-up-threshold 10 \
  --scale-down-threshold 5 \
  --check-interval 15s
```

### Custom workflow control

```bash
# Disable auto-deals, only do data preparation
singularity run unified \
  --disable-auto-deals \
  --max-workers 10
```

## What the unified service does

### 🔄 **Worker Management**
- **Auto-scaling**: Automatically scales workers based on job availability
- **Intelligent job distribution**: Workers are assigned to job types based on current needs
- **Resource optimization**: Scales down when work is light, scales up when busy
- **Fault tolerance**: Restarts failed workers automatically

### ⚡ **Workflow Orchestration**
- **Event-driven progression**: Jobs automatically progress through stages
- **Scan → Pack**: When scan jobs complete, pack jobs start automatically
- **Pack → DagGen**: When pack jobs complete, daggen jobs start automatically  
- **DagGen → Deals**: When all jobs complete, deal schedules are created automatically

### 📊 **Monitoring and Status**
- **Regular status updates**: Logs comprehensive service status every 30 seconds
- **Job queue monitoring**: Tracks ready/processing/complete jobs by type
- **Worker status**: Shows active workers and their job types

## Service Output

### Startup
```
INFO Starting unified auto-preparation service
INFO Worker manager started (min: 1, max: 5, auto-scaling: enabled)
INFO Workflow orchestrator enabled (scan→pack→daggen→deals)
INFO Starting 1 minimum workers
INFO Started managed worker managed-worker-1234 (total workers: 1)
```

### Regular Status Updates
```
=== UNIFIED SERVICE STATUS ===
Workers: 3 active (enabled: true)
Orchestrator enabled: true
Jobs - Scan: 5 ready/15 total, Pack: 0 ready/8 total, DagGen: 2 ready/2 total
Worker 12345678: types=[scan pack daggen], uptime=5m23s
Worker 87654321: types=[scan pack], uptime=3m45s  
Worker abcdefgh: types=[daggen], uptime=1m12s
===============================
```

### Auto-scaling Events
```
INFO Scaling up: adding 2 workers (ready jobs: 12)
INFO Started managed worker managed-worker-5678 (total workers: 4)
INFO Scaling down: removing 1 workers (ready jobs: 1)
INFO Stopping managed worker managed-worker-1234
```

## Configuration Options

### Worker Scaling Strategy

| Setting | Description | Default | Production Recommendation |
|---------|-------------|---------|---------------------------|
| `min-workers` | Always keep this many workers running | 1 | 3-5 |
| `max-workers` | Never exceed this many workers | 5 | 10-20 |
| `scale-up-threshold` | Start new workers when ready jobs ≥ this | 5 | 10-15 |
| `scale-down-threshold` | Stop workers when ready jobs ≤ this | 2 | 3-5 |
| `worker-idle-timeout` | Stop idle workers after this time | 5m | 10m |

### Workflow Control

Disable specific transitions for custom workflows:

```bash
# Only do scanning and packing, no DAG generation or deals
singularity run unified \
  --disable-pack-to-daggen \
  --disable-daggen-to-deals

# Prepare data but create deals manually
singularity run unified --disable-auto-deals
```

## Comparison with Other Approaches

### vs. `onboard` command
- **Onboard**: Best for single dataset workflows, includes monitoring
- **Unified service**: Best for continuous background processing of multiple datasets

### vs. Manual worker management
- **Manual**: `singularity run dataset-worker` (requires separate auto-deal setup)
- **Unified**: Automatic worker management + workflow orchestration + auto-deals

## When to use this service

✅ **Use unified service when:**
- You have continuous data onboarding needs
- You want background processing without manual intervention
- You need automatic scaling based on workload
- You're processing multiple datasets over time

✅ **Use onboard command when:**
- You have a single dataset to process
- You want built-in progress monitoring
- You prefer a simple one-command approach

## Monitoring and Troubleshooting

### Check service status
```bash
# View current configuration (dry run)
singularity run unified --dry-run

# Check preparations being processed
singularity prep status <preparation-name>

# View created deal schedules
singularity deal schedule list
```

### Service logs
The unified service provides comprehensive logging:
- Worker lifecycle events (start/stop/scaling)
- Job progression through workflow stages
- Auto-deal creation events
- Error handling and recovery

## See Also

- [Onboard command](../onboard.md) - Single-command data onboarding
- [Dataset worker](./dataset-worker.md) - Manual worker management
- [Auto-deal commands](../prep/autodeal/README.md) - Manual deal creation controls