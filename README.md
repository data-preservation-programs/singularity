# Singularity
[![codecov](https://codecov.io/github/data-preservation-programs/singularity/branch/main/graph/badge.svg?token=1A3BMQU3LM)](https://codecov.io/github/data-preservation-programs/singularity)
[![Go Report Card](https://goreportcard.com/badge/github.com/data-preservation-programs/singularity)](https://goreportcard.com/report/github.com/data-preservation-programs/singularity)
[![Go Reference](https://pkg.go.dev/badge/github.com/data-preservation-programs/singularity.svg)](https://pkg.go.dev/github.com/data-preservation-programs/singularity)
[![Build](https://github.com/data-preservation-programs/singularity/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/data-preservation-programs/singularity/actions/workflows/go.yml)

The new pure-go implementation of Singularity provides everything you need to onboard your, or your client's data to Filecoin network, with **automatic deal creation** and intelligent workflow management.

## ✨ Key Features

- **🚀 Automatic Deal Creation** - Deal schedules created automatically when data preparation completes
- **📦 Data Preparation** - Efficient scanning, packing, and CAR file generation
- **🔗 Deal Management** - Comprehensive deal scheduling and tracking
- **💼 Wallet Management** - Full wallet lifecycle with metadata support for organization
- **🏪 Storage Integration** - Support for multiple storage backends (local, S3, etc.)
- **📊 Monitoring & Notifications** - Real-time status updates and error handling
- **🔧 Flexible Configuration** - Extensive customization options for different workflows

## 🚀 Quick Start

### Installation

```bash
# Download the latest release
wget https://github.com/data-preservation-programs/singularity/releases/latest/download/singularity-linux-amd64
chmod +x singularity-linux-amd64
sudo mv singularity-linux-amd64 /usr/local/bin/singularity

# Or build from source
git clone https://github.com/data-preservation-programs/singularity.git
cd singularity
go build -o singularity .
```

### Basic Usage

**Single command data onboarding with automatic deal creation:**

```bash
singularity onboard \
  --name "my-dataset" \
  --source "/path/to/data" \
  --auto-create-deals \
  --deal-provider "f01234" \
  --deal-verified \
  --deal-price-per-gb 0.0000001 \
  --start-workers \
  --wait-for-completion
```

**That's it!** ✨ This single command will:
1. Create storage connections automatically
2. Set up data preparation with deal parameters  
3. Start managed workers to process jobs
4. Automatically progress through scan → pack → daggen
5. Create storage deals when preparation completes
6. Monitor progress until completion

## 🤖 Auto-Deal System

The Auto-Deal System automatically creates deal schedules when data preparation jobs complete, eliminating manual intervention. The `onboard` command provides the simplest interface for complete automated workflows.

### How It Works

```
Source Data → Scan → Pack → DAG Gen → Deal Schedule Created ✅
```

All stages progress automatically with event-driven triggering - no polling or manual monitoring required.

### Configuration Options (`onboard` command)

| Flag | Description | Default |
|------|-------------|---------|
| `--auto-create-deals` | Enable automatic deal creation | `true` |
| `--deal-provider` | Storage provider ID (e.g., f01234) | Required |
| `--deal-verified` | Create verified deals | `false` |
| `--deal-price-per-gb` | Price per GB per epoch | `0` |
| `--deal-duration` | Deal duration (e.g., "8760h") | `535 days` |
| `--deal-start-delay` | Deal start delay | `72h` |
| `--validate-wallet` | Validate wallets before creating deals | `false` |
| `--validate-provider` | Validate storage provider | `false` |
| `--start-workers` | Start managed workers automatically | `true` |
| `--wait-for-completion` | Monitor until completion | `false` |

### Manual Monitoring

```bash
# Check preparation status
singularity prep status "my-dataset"

# List all deal schedules
singularity deal schedule list

# Run background processing service
singularity run unified --max-workers 5
```

## 📖 Documentation
[Read the Full Documentation](https://data-programs.gitbook.io/singularity/overview/readme)

## 🛠️ Advanced Usage

### Multiple Storage Providers

Onboard data to different providers with different strategies:

```bash
# Hot storage with fast provider
singularity onboard --name "hot-data" --source "/critical/data" \
  --deal-provider "f01234" --deal-price-per-gb 0.000001 --auto-create-deals

# Cold storage with economical provider  
singularity onboard --name "cold-data" --source "/archive/data" \
  --deal-provider "f05678" --deal-price-per-gb 0.0000001 --auto-create-deals
```

### Conditional Auto-Deals

Use validation to control when deals are created:

```bash
# Only create deals if wallet has sufficient balance
singularity onboard --name "conditional" --source "/data" --auto-create-deals \
  --deal-provider "f01234" --wallet-validation

# Only create deals if provider is verified  
singularity onboard --name "verified-only" --source "/data" --auto-create-deals \
  --deal-provider "f01234" --sp-validation
```

### Monitoring

```bash
# Check preparation status
singularity prep status "my-dataset"

# List all deal schedules
singularity deal schedule list

# Run unified service with monitoring
singularity run unified --max-workers 5
```

### Error Logging

Singularity automatically logs errors and events during operations. View and query error logs to troubleshoot issues:

```bash
# List recent error logs
singularity error log list

# Filter by component and level
singularity error log list --component onboard --level error

# Find errors for specific preparation
singularity error log list --entity-type preparation --entity-id 123

# View only critical errors
singularity error log list --level critical

# Check deal scheduling issues
singularity error log list --component deal_schedule --level warning

# Query logs from the last hour
singularity error log list --start-time $(date -d '1 hour ago' -Iseconds)

# Paginate through results
singularity error log list --limit 20 --offset 40

# Export logs as JSON for analysis
singularity error log list --json --component onboard > error_logs.json
```

**Error Levels:**
- **info**: General information and successful operations
- **warning**: Non-critical issues that should be monitored
- **error**: Failures that need attention but don't stop operations
- **critical**: Severe failures requiring immediate action

**Common Components:**
- **onboard**: Data scanning, preparation, and CAR file generation
- **deal_schedule**: Deal creation and scheduling operations
- **pack**: Data packing and piece generation
- **daggen**: DAG generation from packed data

## 🏗️ Architecture

### Simplified Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Onboard        │    │ Worker Manager  │    │ Workflow        │
│  Command        │────▶│                 │────▶│ Orchestrator    │
│                 │    │ • Auto-scaling  │    │                 │
│ • One command   │    │ • Job processing│    │ • Event-driven  │
│ • Full workflow │    │ • Monitoring    │    │ • Auto-progress │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                  │                        │
                                  ▼                        ▼
                    ┌─────────────────────────────┐ ┌──────────────┐
                    │     Auto-Deal Service       │ │ Deal Schedule│
                    │                             │ │    Created   │
                    │ • Check Readiness          │ │      ✅      │
                    │ • Validate Wallets/SPs    │ │              │
                    │ • Create Deal Schedules    │ │              │
                    └─────────────────────────────┘ └──────────────┘
```

### Key Components

- **Onboard Command**: Single entry point for complete automated workflows
- **Worker Manager**: Auto-scaling workers that process jobs intelligently  
- **Workflow Orchestrator**: Event-driven progression through data preparation stages
- **Auto-Deal Service**: Creates deal schedules when preparations complete
- **Trigger Service**: Handles automatic deal creation logic
- **Validation System**: Ensures wallets and providers are ready for deals
- **Notification System**: Provides observability and error reporting

## 🧪 Testing

```bash
# Run auto-deal tests
go test ./service/autodeal/ -v

# Run integration tests
go test ./service/autodeal/ -v -run "TestTrigger"

# Test CLI functionality
singularity onboard --help
```

## 🔧 Configuration

### Environment Variables

```bash
# Lotus connection
export LOTUS_API="https://api.node.glif.io/rpc/v1"
export LOTUS_TOKEN="your-token"

# Database
export DATABASE_CONNECTION_STRING="sqlite:singularity.db"
```

### Runtime Configuration

```bash
# Run unified service with custom settings
singularity run unified --max-workers 5

# Run with specific worker configuration
singularity run unified --max-workers 10
```

## 🚨 Troubleshooting

### Common Issues

**Auto-deal not triggering:**
- Ensure `--auto-create-deals` is enabled when using `onboard`
- Verify wallet is attached: `singularity prep list-wallets <prep>`
- Check all jobs are complete
- Verify unified service is running: `singularity run unified`

**Deal creation failing:**
- Check provider ID is correct
- Ensure wallet has sufficient balance
- Verify network connectivity to Lotus
- Review validation settings

**Performance issues:**
- Adjust `--max-workers` in unified service for better throughput
- Monitor database performance and connections
- Use appropriate hardware resources for large datasets

### Debug Commands

```bash
# Test onboard workflow
singularity onboard --name "test-dataset" --source "/test/data" --auto-create-deals

# View detailed logs
singularity run unified --max-workers 3

# Check preparation status
singularity prep status "my-dataset"
```

## 🤝 Migration from Manual Workflows

Existing preparations work unchanged! Auto-deal is completely opt-in:

```bash
# Existing workflow (still works)
singularity prep create --name "manual"
singularity deal schedule create --preparation "manual" --provider "f01234"

# New automated workflow
singularity prep create --name "automatic" --auto-create-deals --deal-provider "f01234"
```

## 📊 Monitoring & Observability

### Key Metrics
- Preparations processed per minute
- Deal schedules created automatically
- Validation success/failure rates
- Error frequencies and types

### Log Analysis
```bash
# Monitor auto-deal activity
tail -f singularity.log | grep "autodeal-trigger\|auto-deal"

# View successful deal creations
grep "Auto-Deal Schedule Created Successfully" singularity.log
```

## 🌟 Benefits

### Before Auto-Deal System
- ❌ Manual deal schedule creation required
- ❌ Risk of forgetting to create deals
- ❌ No automation for completed preparations
- ❌ Time-consuming manual monitoring

### After Auto-Deal System
- ✅ Zero-touch deal creation for completed preparations
- ✅ Configurable validation and error handling
- ✅ Background monitoring and batch processing
- ✅ Comprehensive logging and notifications
- ✅ Full backward compatibility

## 🔮 Future Enhancements

- **Dynamic provider selection** based on reputation/pricing
- **Deal success monitoring** and automatic retries
- **Cost optimization** algorithms
- **Advanced scheduling** (time-based, capacity-based)
- **Multi-wallet load balancing**
- **Integration with deal marketplaces**

## 📞 Support

For issues or questions:

1. **Check logs**: `tail -f singularity.log | grep auto-deal`
2. **Review notifications**: `singularity admin notification list`
3. **Run tests**: `go test ./service/autodeal/ -v`
4. **Consult documentation**: [Full Documentation](https://data-programs.gitbook.io/singularity/overview/readme)

## Related Projects
- [js-singularity](https://github.com/tech-greedy/singularity) -
The predecessor that was implemented in Node.js
- [js-singularity-import-boost](https://github.com/tech-greedy/singularity-import) -
Automatically import deals to boost for Filecoin storage providers
- [js-singularity-browser](https://github.com/tech-greedy/singularity-browser) -
A next.js app for browsing singularity made deals
- [go-generate-car](https://github.com/tech-greedy/generate-car) -
The internal tool used by `js-singularity` to generate car files as well as commp
- [go-generate-ipld-car](https://github.com/tech-greedy/generate-car#generate-ipld-car) -
The internal tool used by `js-singularity` to regenerate the CAR that captures the unixfs dag of the dataset.

## License
Dual-licensed under [MIT](https://github.com/filecoin-project/lotus/blob/master/LICENSE-MIT) + [Apache 2.0](https://github.com/filecoin-project/lotus/blob/master/LICENSE-APACHE)
