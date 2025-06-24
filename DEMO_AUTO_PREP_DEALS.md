# Auto-Prep Deal Scheduling Demo

This demo showcases the new **Auto-Prep Deal Scheduling** feature that provides complete data onboarding in a single command - from data source to storage deals.

## Overview

The auto-prep deal scheduling feature eliminates manual intervention by providing:
- **Deal Templates**: Reusable deal configurations for consistent parameters
- **Unified Onboarding**: Complete data preparation with automated deal creation
- **Automatic Storage**: Creates storage connections automatically  
- **Seamless Workflow**: Automatic progression from scanning to deal creation
- **Worker Management**: Built-in workers process jobs automatically

## Prerequisites

```bash
# Ensure Singularity is built with the latest changes
go build -o singularity

# No additional setup required - the onboard command manages everything automatically
```

## Demo 1: Using Deal Templates (Recommended)

The most efficient way to onboard data with reusable deal configurations:

```bash
# First, create a deal template (one-time setup)
./singularity deal-template create \
  --name "standard-archive" \
  --description "Standard archival storage with 18-month retention" \
  --deal-price-per-gb 0.0000000001 \
  --deal-duration 535days \
  --deal-start-delay 72h \
  --deal-verified \
  --deal-keep-unsealed \
  --deal-announce-to-ipni \
  --deal-provider "f01234"

# Now onboard data using the template
./singularity prep create \
  --name "my-dataset" \
  --source "/path/to/your/data" \
  --output "/path/to/output" \
  --auto-create-deals \
  --deal-template "standard-archive" \
  --auto-start \
  --auto-progress
```

## Demo 2: Direct Parameters (No Template)

You can still specify deal parameters directly without using templates:

```bash
# Complete onboarding with direct parameters
./singularity prep create \
  --name "my-dataset" \
  --source "/path/to/your/data" \
  --output "/path/to/output" \
  --auto-create-deals \
  --deal-provider "f01234" \
  --deal-verified \
  --deal-price-per-gb 0.0000001 \
  --deal-duration 535days \
  --deal-start-delay 72h \
  --auto-start \
  --auto-progress
```

That's it! This single command will:
1. ✅ Create source and output storage automatically
2. ✅ Create preparation with auto-deal configuration
3. ✅ Start managed workers to process jobs
4. ✅ Begin scanning immediately
5. ✅ Automatically progress through scan → pack → daggen → deals
6. ✅ Monitor progress until completion

## Demo Script

Here's a complete demo script showcasing both deal templates and direct parameters:

```bash
#!/bin/bash

echo "=== Auto-Prep Deal Scheduling Demo with Templates ==="
echo

echo "📋 Step 1: Creating deal templates for reuse..."

# Create enterprise template
./singularity deal-template create \
  --name "enterprise-tier" \
  --description "Enterprise-grade storage with 3-year retention" \
  --deal-duration 1095days \
  --deal-price-per-gb 0.0000000002 \
  --deal-verified \
  --deal-keep-unsealed \
  --deal-announce-to-ipni \
  --deal-start-delay 72h

# Create research template  
./singularity deal-template create \
  --name "research-archive" \
  --description "Long-term research data archive" \
  --deal-duration 1460days \
  --deal-price-per-gb 0.0000000001 \
  --deal-verified \
  --deal-keep-unsealed

echo "✅ Deal templates created!"
echo

# List templates
echo "📋 Available deal templates:"
./singularity deal-template list
echo

echo "🚀 Step 2: Onboarding data using templates..."

# Create some demo data if needed (check if directories already exist)
if [ -d "./demo-data" ] && [ "$(ls -A ./demo-data)" ]; then
    echo "Warning: ./demo-data directory already exists and contains files. Please remove or backup existing content before proceeding."
    echo "Use: rm -rf ./demo-data ./demo-output"
    exit 1
fi
mkdir -p ./demo-data ./demo-output
echo "Sample file for enterprise demo" > ./demo-data/enterprise-data.txt
echo "Sample file for research demo" > ./demo-data/research-data.txt

echo "Creating enterprise dataset with template..."
./singularity prep create \
  --name "enterprise-dataset" \
  --source "./demo-data" \
  --output "./demo-output" \
  --auto-create-deals \
  --deal-template "enterprise-tier" \
  --auto-start \
  --auto-progress

echo
echo "Creating research dataset with template override..."
./singularity prep create \
  --name "research-dataset" \
  --source "./demo-data" \
  --auto-create-deals \
  --deal-template "research-archive" \
  --deal-provider "f01000" \
  --auto-start \
  --auto-progress

echo
echo "🎉 Demo Complete!"
echo "✅ Deal templates created for reuse"
echo "✅ Multiple datasets prepared with consistent deal parameters"
echo "✅ Template values overridden when needed"
```

## Deal Template Management

Manage your deal templates for reuse across projects:

```bash
# List all templates
./singularity deal-template list

# View template details
./singularity deal-template get enterprise-tier

# Create additional templates for different use cases
./singularity deal-template create \
  --name "budget-tier" \
  --description "Cost-effective storage for non-critical data" \
  --deal-duration 365days \
  --deal-price-per-gb 0.00000000005 \
  --deal-start-delay 168h

# Delete templates when no longer needed
./singularity deal-template delete old-template
```

## Manual Monitoring

Monitor your preparations and deal creation:

```bash
# Monitor preparation progress
./singularity prep status my-dataset

# Check if deals were created
./singularity deal schedule list

# View specific template details
./singularity deal-template get enterprise-tier

# View schedules for this preparation via API
curl http://localhost:7005/api/preparation/my-dataset/schedules
```

## Key Features Demonstrated

1. **Deal Templates**: Reusable deal configurations for consistency across projects
2. **Template Override**: Ability to override specific template values per preparation
3. **Automatic Storage Creation**: Local storage connections created automatically
4. **Integrated Auto-Progress**: Seamless flow from scanning to deal creation
5. **Parameter Flexibility**: Choose between templates or direct parameter specification
6. **Template Management**: Full CRUD operations for deal template lifecycle

## Expected Output

When the demo completes successfully, you should see:
- ✅ Deal templates created and available for reuse
- ✅ Storage connections created automatically for each preparation
- ✅ Preparations created with auto-deal configuration from templates
- ✅ Template values applied with option to override specific parameters
- ✅ Progress updates showing scan → pack → daggen → deals
- ✅ Storage deals created using template configurations

## Advanced Usage

```bash
# Create multiple sources with template
./singularity prep create \
  --name "multi-source-dataset" \
  --source "/path/to/source1" \
  --source "/path/to/source2" \
  --output "/path/to/output" \
  --auto-create-deals \
  --deal-template "enterprise-tier" \
  --wallet-validation \
  --sp-validation \
  --auto-start \
  --auto-progress

# Preparation without automatic deal creation
./singularity prep create \
  --name "prep-only-dataset" \
  --source "/path/to/data" \
  --auto-start \
  --auto-progress

# Override template with custom parameters
./singularity prep create \
  --name "custom-deals-dataset" \
  --source "/path/to/data" \
  --auto-create-deals \
  --deal-template "research-archive" \
  --deal-provider "f01000" \
  --deal-verified=false \
  --deal-price-per-gb 0.0000000005

# Multiple templates for different tiers
./singularity deal-template create --name "hot-storage" --deal-duration 180days --deal-price-per-gb 0.0000000005
./singularity deal-template create --name "cold-archive" --deal-duration 1460days --deal-price-per-gb 0.0000000001
```

## Troubleshooting

```bash
# Check preparation status
./singularity prep status <preparation-name>

# List all deal schedules
./singularity deal schedule list

# View available deal templates
./singularity deal-template list

# Check specific template configuration
./singularity deal-template get <template-name>

# Check worker status (if using separate terminals)
./singularity run unified --dry-run
```

## Benefits of Deal Templates

This approach offers several advantages over manual parameter specification:

1. **Consistency**: Ensure all datasets use the same deal parameters
2. **Reusability**: Create templates once, use across multiple projects
3. **Organization**: Maintain different templates for different data tiers
4. **Simplification**: Reduce complex command-line arguments to simple template names
5. **Flexibility**: Override specific parameters when needed while keeping template defaults
6. **Maintenance**: Update deal parameters organization-wide by modifying templates

This streamlined approach with deal templates reduces what used to be a complex multi-step process into a standardized, reusable workflow, making large-scale data onboarding to Filecoin much simpler and more accessible.