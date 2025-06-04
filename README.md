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

1. **Create a preparation with auto-deal enabled:**
```bash
singularity prep create \
  --name "my-dataset" \
  --source "/path/to/data" \
  --auto-create-deals \
  --deal-provider "f01234" \
  --deal-verified \
  --deal-price-per-gb 0.0000001
```

2. **Attach a wallet:**
```bash
singularity prep attach-wallet "my-dataset" "f3your-wallet-address"
```

3. **Start the services:**
```bash
# Start dataset worker (processes data)
singularity run dataset-worker --enable-pack &

# Start auto-deal daemon (creates deals automatically)
singularity run autodeal &
```

4. **Watch the magic happen!** ✨
   - Data gets processed automatically
   - When all jobs complete, deal schedules are created
   - No manual intervention required!

## 🤖 Auto-Deal System

The Auto-Deal System automatically creates deal schedules when data preparation jobs complete, eliminating manual intervention.

### How It Works

```
Data Processing → Job Completion → Auto-Deal Trigger → Deal Schedule Created ✅
```

### Configuration Options

| Flag | Description | Default |
|------|-------------|---------|
| `--auto-create-deals` | Enable automatic deal creation | `false` |
| `--deal-provider` | Storage provider ID (e.g., f01234) | Required |
| `--deal-verified` | Create verified deals | `false` |
| `--deal-price-per-gb` | Price per GB per epoch | `0` |
| `--deal-duration` | Deal duration (e.g., "8760h") | `12840h` |
| `--wallet-validation` | Validate wallets before creating deals | `true` |
| `--sp-validation` | Validate storage provider | `true` |

### Auto-Deal Daemon

Monitor and batch process ready preparations:

```bash
singularity run autodeal \
  --check-interval 30s \
  --enable-batch-mode \
  --enable-job-hooks \
  --max-retries 3
```

### Manual Commands

```bash
# Create deal schedule for specific preparation
singularity prep autodeal create --preparation "my-dataset"

# Process all ready preparations
singularity prep autodeal process

# Check if preparation is ready
singularity prep autodeal check --preparation "my-dataset"
```

## 📖 Documentation
[Read the Full Documentation](https://data-programs.gitbook.io/singularity/overview/readme)

## 🛠️ Advanced Usage

### Multiple Storage Providers

Create different preparations for different providers:

```bash
# Hot storage with fast provider
singularity prep create --name "hot-data" --deal-provider "f01234" --deal-price-per-gb 0.000001

# Cold storage with economical provider  
singularity prep create --name "cold-data" --deal-provider "f05678" --deal-price-per-gb 0.0000001
```

### Conditional Auto-Deals

Use validation to control when deals are created:

```bash
# Only create deals if wallet has sufficient balance
singularity prep create --name "conditional" --auto-create-deals --wallet-validation

# Only create deals if provider is verified
singularity prep create --name "verified-only" --auto-create-deals --sp-validation
```

### Monitoring

```bash
# View auto-deal notifications
singularity admin notification list --source "auto-deal-service"

# Monitor daemon logs
singularity run autodeal 2>&1 | grep "autodeal-trigger"
```

## 🏗️ Architecture

### Component Overview

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Dataset Worker │    │   Pack Handler  │    │ Auto-Deal Daemon│
│                 │────▶│                 │────▶│                 │
│ Job Completion  │    │ Job Completion  │    │ Trigger Service │
│     Hooks       │    │     Hooks       │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                        │                        │
         └────────────────────────┼────────────────────────┘
                                  ▼
                    ┌─────────────────────────────┐
                    │     Auto-Deal Service       │
                    │                             │
                    │ • Check Readiness          │
                    │ • Validate Wallets/SPs    │
                    │ • Create Deal Schedules    │
                    │ • Send Notifications       │
                    └─────────────────────────────┘
```

### Key Services

- **Dataset Worker**: Processes data and triggers auto-deals on completion
- **Auto-Deal Daemon**: Monitors preparations and batch processes ready ones
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
singularity prep autodeal check --help
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
# Disable automatic triggering temporarily
singularity run autodeal --enable-job-hooks=false

# Increase monitoring frequency
singularity run autodeal --check-interval=10s
```

## 🚨 Troubleshooting

### Common Issues

**Auto-deal not triggering:**
- Ensure `--auto-create-deals` is enabled
- Verify wallet is attached: `singularity prep list-wallets <prep>`
- Check all jobs are complete
- Verify daemon is running with `--enable-job-hooks`

**Deal creation failing:**
- Check provider ID is correct
- Ensure wallet has sufficient balance
- Verify network connectivity to Lotus
- Review validation settings

**Performance issues:**
- Increase `--check-interval` for less frequent checks
- Use `--max-retries` to limit retry attempts
- Monitor database performance

### Debug Commands

```bash
# Test auto-deal creation manually
singularity prep autodeal create --preparation "my-dataset"

# View detailed logs
singularity run autodeal --log-level debug

# Check preparation status
singularity prep autodeal check --preparation "my-dataset"
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
