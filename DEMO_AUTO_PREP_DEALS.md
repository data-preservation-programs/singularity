# Auto-Prep Deal Scheduling Demo

This demo showcases the new **Auto-Prep Deal Scheduling** feature that automatically creates storage deals when data preparation completes.

## Overview

The auto-prep deal scheduling feature eliminates manual intervention by:
- Automatically creating deal schedules when all preparation jobs complete
- Providing configurable deal parameters during preparation setup
- Running a background monitoring service
- Supporting batch processing of multiple preparations

## Demo Scenario

We'll demonstrate:
1. Setting up auto-deal configuration
2. Creating a preparation with auto-deal enabled
3. Monitoring the preparation progress
4. Automatic deal schedule creation
5. Manual commands for checking and triggering

## Prerequisites

```bash
# Ensure Singularity is built with the latest changes
go build -o singularity

# Start the Singularity daemon (in separate terminal)
./singularity daemon
```

## Demo Steps

### Step 1: Setup Storage and Wallet

```bash
# Add a local storage source
./singularity storage create local source --path /path/to/your/data

# Add a local storage output
./singularity storage create local output --path /path/to/output

# Create and fund a wallet
./singularity wallet create
# Fund the wallet with test FIL (testnet)
```

### Step 2: Create Preparation with Auto-Deal Configuration

```bash
# Create preparation with auto-deal enabled
curl -X POST http://localhost:7005/preparation \
  -H "Content-Type: application/json" \
  -d '{
    "name": "demo-auto-dataset",
    "source": {
      "storageId": 1,
      "path": ""
    },
    "output": {
      "storageId": 2
    },
    "autoCreateDeals": true,
    "dealProvider": "f01234",
    "dealVerified": true,
    "dealPricePerGb": 0.0000001,
    "dealDuration": "8760h",
    "dealStartDelay": "72h",
    "dealKeepUnsealed": false,
    "dealAnnounceToIpni": true,
    "walletValidation": true,
    "spValidation": true
  }'
```

### Step 3: Start Auto-Deal Monitoring Service

```bash
# Start the auto-deal daemon (in separate terminal)
./singularity run autodeal \
  --check-interval 30s \
  --enable-batch-mode \
  --enable-job-hooks \
  --max-retries 3 \
  --retry-interval 5m
```

### Step 4: Monitor Preparation Progress

```bash
# Check preparation status
./singularity preparation status demo-auto-dataset

# Check if preparation is ready for auto-deal
./singularity dataprep autodeal check --preparation demo-auto-dataset

# List all jobs for the preparation
./singularity job list --preparation demo-auto-dataset
```

### Step 5: Manual Auto-Deal Commands (Optional)

```bash
# Manually trigger auto-deal for specific preparation
./singularity dataprep autodeal create --preparation demo-auto-dataset

# Process all ready preparations in batch
./singularity dataprep autodeal process

# Check the created deal schedules
./singularity schedule list --preparation demo-auto-dataset
```

### Step 6: Verify Deal Schedule Creation

```bash
# List deal schedules created by auto-deal system
./singularity schedule list

# Check specific schedule details
./singularity schedule status <schedule-id>

# View schedule with detailed info
curl http://localhost:7005/preparation/demo-auto-dataset/schedules
```

## Demo Script

Here's a complete demo script you can run:

```bash
#!/bin/bash

echo "=== Auto-Prep Deal Scheduling Demo ==="
echo

echo "Step 1: Setting up storages..."
./singularity storage create local source --path ./demo-data
./singularity storage create local output --path ./demo-output
echo

echo "Step 2: Creating preparation with auto-deal enabled..."
PREP_RESPONSE=$(curl -s -X POST http://localhost:7005/preparation \
  -H "Content-Type: application/json" \
  -d '{
    "name": "demo-auto-dataset",
    "source": {"storageId": 1, "path": ""},
    "output": {"storageId": 2},
    "autoCreateDeals": true,
    "dealProvider": "f01234",
    "dealVerified": true,
    "dealPricePerGb": 0.0000001,
    "dealDuration": "8760h",
    "dealStartDelay": "72h"
  }')

echo "Preparation created: $PREP_RESPONSE"
echo

echo "Step 3: Starting auto-deal daemon..."
echo "Run in separate terminal: ./singularity run autodeal --check-interval 10s"
echo

echo "Step 4: Monitoring preparation..."
while true; do
  STATUS=$(./singularity dataprep autodeal check --preparation demo-auto-dataset 2>/dev/null || echo "not ready")
  echo "Preparation status: $STATUS"
  
  if [[ "$STATUS" == *"ready"* ]]; then
    echo "✅ Preparation is ready for auto-deal!"
    break
  fi
  
  sleep 5
done
echo

echo "Step 5: Auto-deal will trigger automatically, or run manually:"
echo "./singularity dataprep autodeal create --preparation demo-auto-dataset"
echo

echo "Step 6: Verify deal schedules created:"
./singularity schedule list --preparation demo-auto-dataset
echo

echo "=== Demo Complete! ==="
echo "The auto-deal system has automatically created deal schedules for your completed preparation."
```

## Key Features Demonstrated

1. **Automatic Triggering**: Deal schedules are created automatically when all jobs complete
2. **Configuration Flexibility**: Deal parameters set during preparation creation
3. **Monitoring Service**: Background daemon continuously checks for ready preparations
4. **Manual Override**: Commands to manually check and trigger auto-deal creation
5. **Batch Processing**: Ability to process multiple preparations simultaneously
6. **Validation**: Optional wallet and storage provider validation

## Expected Output

When the demo completes successfully, you should see:
- ✅ Preparation created with auto-deal configuration
- ✅ All preparation jobs completed
- ✅ Deal schedule automatically created with correct parameters
- ✅ Schedule visible in the schedules list

## Troubleshooting

```bash
# Check daemon logs
./singularity run autodeal --enable-batch-mode --exit-on-complete

# Verify preparation completion
./singularity job list --preparation demo-auto-dataset

# Check for errors in deal schedule creation
./singularity schedule list --all
```

## Advanced Usage

```bash
# Run daemon with custom settings
./singularity run autodeal \
  --check-interval 60s \
  --max-retries 5 \
  --retry-interval 10m \
  --exit-on-complete

# Batch process all ready preparations
./singularity dataprep autodeal process

# Check multiple preparations
for prep in prep1 prep2 prep3; do
  ./singularity dataprep autodeal check --preparation $prep
done
```

This demo showcases how the auto-prep deal scheduling feature streamlines the storage deal creation process, reducing manual overhead and enabling automated large-scale data onboarding workflows.