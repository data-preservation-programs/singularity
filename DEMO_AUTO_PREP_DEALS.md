# Auto-Prep Deal Scheduling Demo

This demo showcases the new **Auto-Prep Deal Scheduling** feature that provides complete data onboarding in a single command - from data source to storage deals.

## Overview

The auto-prep deal scheduling feature eliminates manual intervention by providing a unified `onboard` command that:
- Creates storage connections automatically
- Sets up data preparation with deal parameters
- Starts scanning, packing, and DAG generation automatically
- Creates storage deals when preparation completes
- Manages workers to process jobs automatically

## Prerequisites

```bash
# Ensure Singularity is built with the latest changes
go build -o singularity

# Start the Singularity API (in separate terminal if not using --start-workers)
./singularity run api
```

## Simple Demo - Single Command Onboarding

The simplest way to onboard data with automatic deal creation:

```bash
# Complete onboarding in one command
./singularity onboard \
  --name "my-dataset" \
  --source "/path/to/your/data" \
  --output "/path/to/output" \
  --enable-deals \
  --deal-provider "f01234" \
  --deal-verified \
  --deal-price-per-gb 0.0000001 \
  --deal-duration "8760h" \
  --deal-start-delay "72h" \
  --start-workers \
  --wait-for-completion
```

That's it! This single command will:
1. ✅ Create source and output storage automatically
2. ✅ Create preparation with auto-deal configuration
3. ✅ Start managed workers to process jobs
4. ✅ Begin scanning immediately
5. ✅ Automatically progress through scan → pack → daggen → deals
6. ✅ Monitor progress until completion

## Demo Script

Here's a complete demo script:

```bash
#!/bin/bash

echo "=== Single Command Auto-Prep Deal Scheduling Demo ==="
echo

echo "🚀 Starting complete data onboarding with automatic deal creation..."
echo "This will take your data from source files to Filecoin storage deals automatically."
echo

# Create some demo data if needed
mkdir -p ./demo-data ./demo-output
echo "Sample file for demo" > ./demo-data/sample.txt

echo "Running onboard command..."
./singularity onboard \
  --name "demo-auto-dataset" \
  --source "./demo-data" \
  --output "./demo-output" \
  --enable-deals \
  --deal-provider "f01234" \
  --deal-verified \
  --deal-price-per-gb 0.0000001 \
  --deal-duration "8760h" \
  --deal-start-delay "72h" \
  --start-workers \
  --max-workers 2 \
  --wait-for-completion \
  --timeout "30m"

echo
echo "🎉 Demo Complete!"
echo "Your data has been automatically processed and storage deals have been created."
```

## Manual Monitoring (Alternative to --wait-for-completion)

If you prefer to monitor manually instead of using `--wait-for-completion`:

```bash
# Start onboarding without waiting
./singularity onboard \
  --name "my-dataset" \
  --source "/path/to/data" \
  --enable-deals \
  --deal-provider "f01234" \
  --start-workers

# Monitor progress manually
./singularity prep status my-dataset

# Check if deals were created
./singularity deal schedule list

# View schedules for this preparation
curl http://localhost:7005/api/preparation/my-dataset/schedules
```

## Key Features Demonstrated

1. **Single Command Workflow**: Complete data onboarding in one command
2. **Automatic Storage Creation**: No need to pre-create storage connections
3. **Integrated Worker Management**: Built-in workers process jobs automatically  
4. **Automatic Job Progression**: Seamless flow from scanning to deal creation
5. **Progress Monitoring**: Built-in monitoring with timeout support
6. **Deal Configuration**: All deal parameters configured upfront

## Expected Output

When the demo completes successfully, you should see:
- ✅ Storage connections created automatically
- ✅ Preparation created with auto-deal configuration
- ✅ Workers started and processing jobs automatically
- ✅ Progress updates showing scan → pack → daggen → deals
- ✅ Storage deals created and visible in schedule list

## Advanced Usage

```bash
# Onboard multiple sources with validation
./singularity onboard \
  --name "multi-source-dataset" \
  --source "/path/to/source1" \
  --source "/path/to/source2" \
  --output "/path/to/output1" \
  --output "/path/to/output2" \
  --enable-deals \
  --deal-provider "f01234" \
  --validate-wallet \
  --validate-provider \
  --start-workers \
  --max-workers 5

# Onboard without automatic deal creation
./singularity onboard \
  --name "prep-only-dataset" \
  --source "/path/to/data" \
  --enable-deals=false \
  --start-workers

# Run with different deal parameters
./singularity onboard \
  --name "custom-deals-dataset" \
  --source "/path/to/data" \
  --enable-deals \
  --deal-provider "f01000" \
  --deal-verified=false \
  --deal-price-per-gb 0.1 \
  --deal-duration "17520h" \
  --deal-start-delay "168h"
```

## Troubleshooting

```bash
# Check preparation status
./singularity prep status <preparation-name>

# List all deal schedules
./singularity deal schedule list

# Check worker status (if using separate terminals)
./singularity run unified --dry-run
```

This streamlined approach reduces what used to be a complex multi-step process into a single command, making large-scale data onboarding to Filecoin much simpler and more accessible.