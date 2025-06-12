# Complete data onboarding workflow

> **🚀 This is the recommended command** for most users. It provides a unified workflow from source data to storage deals.

The `onboard` command automates the entire data onboarding process in a single command, eliminating the need for manual job management and monitoring.

{% code fullWidth="true" %}
```
NAME:
   singularity onboard - Complete data onboarding workflow (storage → preparation → scanning → deal creation)

USAGE:
   singularity onboard [command options]

DESCRIPTION:
   The onboard command provides a unified workflow for complete data onboarding.

   It performs the following steps automatically:
   1. Creates storage connections (if paths provided)
   2. Creates data preparation with deal parameters
   3. Starts scanning immediately
   4. Enables automatic job progression (scan → pack → daggen → deals)
   5. Optionally starts managed workers to process jobs

   This is the simplest way to onboard data from source to storage deals.

OPTIONS:
   Data source flags:
   --name value             Name for the preparation
   --source value           Local source path(s) to onboard
   --output value           Local output path(s) for CAR files (optional)

   Preparation settings:
   --max-size value         Maximum size of a single CAR file (default: "31.5GiB")
   --no-dag                 Disable maintaining folder DAG structure (default: false)

   Deal Settings:
   --enable-deals           Enable automatic deal creation after preparation completion (default: true)
   --deal-provider value    Storage Provider ID for deals (e.g., f01000)
   --deal-price-per-gb value Price in FIL per GiB for storage deals (default: 0)
   --deal-duration value    Duration for storage deals (e.g., 535 days) (default: 12840h0m0s)
   --deal-start-delay value Start delay for storage deals (e.g., 72h) (default: 72h0m0s)
   --deal-verified          Whether deals should be verified (default: false)

   Worker management:
   --start-workers          Start managed workers to process jobs automatically (default: true)
   --max-workers value      Maximum number of workers to run (default: 3)

   Progress monitoring:
   --wait-for-completion    Wait and monitor until all jobs complete (default: false)
   --timeout value          Timeout for waiting for completion (0 = no timeout) (default: 0s)

   Validation:
   --validate-wallet        Enable wallet balance validation (default: false)
   --validate-provider      Enable storage provider validation (default: false)

   --help, -h               show help
```
{% endcode %}

## Basic Examples

### Simple data onboarding with deals

```bash
singularity onboard \
  --name "my-dataset" \
  --source "/path/to/data" \
  --enable-deals \
  --deal-provider "f01234"
```

### Complete workflow with monitoring

```bash
singularity onboard \
  --name "important-data" \
  --source "/critical/data" \
  --output "/backup/cars" \
  --enable-deals \
  --deal-provider "f01234" \
  --deal-verified \
  --deal-price-per-gb 0.0000001 \
  --start-workers \
  --max-workers 5 \
  --wait-for-completion \
  --timeout "2h"
```

### Multiple sources to single preparation

```bash
singularity onboard \
  --name "multi-source-dataset" \
  --source "/data/folder1" \
  --source "/data/folder2" \
  --source "/data/folder3" \
  --enable-deals \
  --deal-provider "f01234"
```

### Data preparation without deals

```bash
singularity onboard \
  --name "prep-only" \
  --source "/data" \
  --enable-deals=false \
  --start-workers
```

## Advanced Configuration

### With validation and custom parameters

```bash
singularity onboard \
  --name "production-dataset" \
  --source "/production/data" \
  --enable-deals \
  --deal-provider "f01000" \
  --deal-verified \
  --deal-duration "17520h" \
  --deal-start-delay "168h" \
  --validate-wallet \
  --validate-provider \
  --max-workers 10
```

### Custom CAR file settings

```bash
singularity onboard \
  --name "large-files" \
  --source "/huge/data" \
  --max-size "63GiB" \
  --no-dag \
  --enable-deals \
  --deal-provider "f01234"
```

## What happens when you run this command

1. **🏗️ Storage Setup**: Automatically creates local storage connections for source and output paths
2. **📋 Preparation Creation**: Creates a new preparation with specified name and deal parameters
3. **⚙️ Workflow Activation**: Enables automatic job progression (scan → pack → daggen → deals)
4. **👷 Worker Management**: Starts managed workers to process jobs (if `--start-workers` is enabled)
5. **🔍 Job Initiation**: Begins scanning source data immediately
6. **📊 Progress Monitoring**: Shows real-time progress updates (if `--wait-for-completion` is enabled)
7. **🎯 Deal Creation**: Automatically creates deal schedules when all jobs complete

## Output and Monitoring

### Success Output
```
🚀 Starting unified data onboarding...

📋 Creating data preparation...
✓ Created preparation: my-dataset (ID: 123)

⚙️ Enabling workflow orchestration...
✓ Automatic job progression enabled (scan → pack → daggen → deals)

👷 Starting managed workers...
✓ Started 3 managed workers

🔍 Starting initial scanning...
✓ Scanning started for all source attachments

✅ Onboarding initiated successfully!
```

### Progress Monitoring (with `--wait-for-completion`)
```
📊 Monitoring progress...
📊 Progress: 0/15 jobs complete | Scan: 15 ready, 0 processing, 0 complete
📊 Progress: 3/15 jobs complete | Scan: 12 ready, 3 processing, 0 complete
📊 Progress: 15/15 jobs complete | Scan: 0 ready, 0 processing, 15 complete | Pack: 8 ready, 0 processing, 0 complete
...
📊 Progress: 23/23 jobs complete | Deals: 1 schedule(s) created
🎉 Onboarding completed successfully!
```

## Alternative Commands

If you need more granular control, you can use the traditional multi-step approach:

```bash
# Traditional approach (more complex)
singularity prep create --name "dataset" --auto-create-deals --deal-provider "f01234"
singularity run unified --max-workers 5
```

But for most use cases, the `onboard` command is simpler and more reliable.

## Troubleshooting

If the onboard command fails:

```bash
# Check preparation status
singularity prep status "my-dataset"

# Check deal schedules created
singularity deal schedule list

# Check worker status
singularity run unified --dry-run
```

## See Also

- [Auto-deal management commands](./prep/autodeal/README.md) - Manual override commands
- [Unified service](./run/README.md) - Background service for continuous processing
- [Auto-Deal System documentation](../auto-deal-system.md) - Technical details