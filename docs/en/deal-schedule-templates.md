# Deal Schedule Templates

Deal schedule templates are reusable configurations that store deal parameters for data preparation workflows. They simplify the process of creating preparations with consistent deal settings and reduce the need to specify deal parameters manually each time.

## Overview

Deal schedule templates allow you to:
- Define and store a complete set of deal parameters once
- Reuse the same deal configuration across multiple preparations
- Ensure consistency in deal pricing and settings
- Simplify the onboarding process for new users
- Maintain organization-wide deal standards

## Creating Deal Schedule Templates

Use the `singularity deal-schedule-template create` command to create a new deal schedule template:

```bash
singularity deal-template create \
  --name "standard-archive" \
  --description "Standard archival storage deals" \
  --deal-price-per-gb 0.0000000001 \
  --deal-duration 535days \
  --deal-start-delay 72h \
  --deal-verified \
  --deal-keep-unsealed \
  --deal-announce-to-ipni \
  --deal-provider f01000
```

### Available Parameters

| Parameter | Description | Example |
|-----------|-------------|---------|
| `--name` | Unique name for the template (required) | `"enterprise-tier"` |
| `--description` | Human-readable description | `"High-performance storage deals"` |
| `--deal-price-per-gb` | Price in FIL per GiB | `0.0000000001` |
| `--deal-price-per-gb-epoch` | Price in FIL per GiB per epoch | `0.0000000001` |
| `--deal-price-per-deal` | Fixed price in FIL per deal | `0.01` |
| `--deal-duration` | Deal duration | `535days`, `1y`, `8760h` |
| `--deal-start-delay` | Delay before deal starts | `72h`, `3days` |
| `--deal-verified` | Enable verified deals (datacap) | Flag |
| `--deal-keep-unsealed` | Keep unsealed copy | Flag |
| `--deal-announce-to-ipni` | Announce to IPNI network | Flag |
| `--deal-provider` | Storage Provider ID | `f01000` |
| `--deal-url-template` | URL template for content | `"https://example.com/{PIECE_CID}"` |
| `--deal-http-headers` | HTTP headers as JSON | `'{"Authorization":"Bearer token"}'` |

## Managing Deal Templates

### List Templates
```bash
# List all deal schedule templates
singularity deal-schedule-template list

# Output as JSON
singularity deal-schedule-template list --json
```

### View Template Details
```bash
# View specific template
singularity deal-schedule-template get standard-archive

# View by ID
singularity deal-schedule-template get 1
```

### Delete Templates
```bash
# Delete by name
singularity deal-schedule-template delete standard-archive

# Delete by ID  
singularity deal-schedule-template delete 1
```

## Using Deal Schedule Templates

### In Preparation Creation

Apply a deal schedule template when creating a preparation:

```bash
singularity prep create \
  --name "my-dataset" \
  --source /path/to/data \
  --auto-create-deals \
  --deal-schedule-template standard-archive
```

### Override Template Values

You can override specific template values by providing parameters directly:

```bash
singularity prep create \
  --name "my-dataset" \
  --source /path/to/data \
  --auto-create-deals \
  --deal-schedule-template standard-archive \
  --deal-price-per-gb 0.0000000002  # Override template price
```

### Manual Parameters (No Template)

You can still specify all deal parameters manually without using a template:

```bash
singularity prep create \
  --name "my-dataset" \
  --source /path/to/data \
  --auto-create-deals \
  --deal-price-per-gb 0.0000000001 \
  --deal-duration 535days \
  --deal-verified \
  --deal-provider f01000
```

## Template Priority

When both a template and direct parameters are provided:
1. **Direct parameters always override template values**
2. **Template values are used for unspecified parameters**
3. **Default values are used if neither template nor direct parameters specify a value**

Example:
```bash
# Template has: price=0.0000000001, duration=535days, verified=true
# Command specifies: price=0.0000000002, provider=f02000
# Result: price=0.0000000002 (overridden), duration=535days (from template), 
#         verified=true (from template), provider=f02000 (from command)
```

## Best Practices

### Template Naming
- Use descriptive names: `enterprise-tier`, `budget-storage`, `research-archive`
- Include version numbers for evolving templates: `standard-v1`, `standard-v2`
- Use organization prefixes: `acme-standard`, `research-lab-default`

### Template Organization
```bash
# Create templates for different use cases
singularity deal-template create --name "hot-storage" --deal-duration 180days --deal-price-per-gb 0.0000000005
singularity deal-template create --name "cold-archive" --deal-duration 1460days --deal-price-per-gb 0.0000000001
singularity deal-template create --name "research-tier" --deal-verified --deal-duration 1095days
```

### Parameter Guidelines
- **Duration**: Match your data retention requirements
  - Short-term: 180-365 days
  - Medium-term: 1-3 years  
  - Long-term: 3+ years
- **Pricing**: Consider storage provider economics
  - Research current market rates
  - Factor in deal duration and data size
- **Verification**: Use `--deal-verified` for datacap deals
- **Provider Selection**: Research provider reliability and pricing

## Examples

### Enterprise Template
```bash
singularity deal-template create \
  --name "enterprise-standard" \
  --description "Enterprise-grade storage with 3-year retention" \
  --deal-duration 1095days \
  --deal-price-per-gb 0.0000000002 \
  --deal-verified \
  --deal-keep-unsealed \
  --deal-announce-to-ipni \
  --deal-start-delay 72h
```

### Research Archive Template
```bash
singularity deal-template create \
  --name "research-archive" \
  --description "Long-term research data archive with datacap" \
  --deal-duration 1460days \
  --deal-price-per-gb 0.0000000001 \
  --deal-verified \
  --deal-keep-unsealed \
  --deal-announce-to-ipni
```

### Budget Storage Template
```bash
singularity deal-template create \
  --name "budget-tier" \
  --description "Cost-effective storage for non-critical data" \
  --deal-duration 365days \
  --deal-price-per-gb 0.00000000005 \
  --deal-start-delay 168h
```

## Integration with Workflows

Deal templates integrate seamlessly with Singularity's automated workflows:

```bash
# Create template
singularity deal-template create --name "workflow-standard" --deal-verified --deal-duration 1095days

# Use in automated preparation
singularity prep create \
  --source /data/dataset1 \
  --deal-schedule-template workflow-standard \
  --auto-create-deals \
  --auto-start \
  --auto-progress
```

This approach ensures consistent deal parameters across all your data preparation workflows while maintaining the flexibility to override specific values when needed.