# wallet balance

Get FIL balance and FIL+ datacap balance information for any Filecoin wallet address.

## Usage

```bash
singularity wallet balance <wallet_address>
```

## Description

The `wallet balance` command queries the Filecoin network via Lotus to retrieve real-time balance information for any wallet address. This includes both the standard FIL balance and FIL+ verified client datacap allocations.

## Arguments

- `<wallet_address>` - The Filecoin wallet address to query (required)

## Output

The command returns detailed balance information:

| Field | Description | Example |
|-------|-------------|---------|
| **address** | The wallet address queried | `f12syf7zd3lfsv43aj2kb454ymaqw7debhumjnbqa` |
| **balance** | FIL balance in human-readable format | `1.000000 FIL` |
| **balanceAttoFIL** | Raw balance in attoFIL (10^-18 FIL) | `1000000000000000000` |
| **dataCap** | FIL+ datacap balance in GiB | `1024.50 GiB` |
| **dataCapBytes** | Raw datacap in bytes | `1100048498688` |
| **error** | Error message if any operation failed | `null` or error details |

## Examples

### Basic Usage

```bash
# Check balance for a wallet
singularity wallet balance f12syf7zd3lfsv43aj2kb454ymaqw7debhumjnbqa
```

**Output:**
```
Address                                    Balance       BalanceAttoFIL       DataCap      DataCapBytes   Error  
f12syf7zd3lfsv43aj2kb454ymaqw7debhumjnbqa  1.000000 FIL  1000000000000000000  1024.50 GiB  1100048498688  <nil>  
```

### JSON Output

```bash
# Get balance in JSON format for programmatic use
singularity --json wallet balance f12syf7zd3lfsv43aj2kb454ymaqw7debhumjnbqa
```

**Output:**
```json
{
  "address": "f12syf7zd3lfsv43aj2kb454ymaqw7debhumjnbqa",
  "balance": "1.000000 FIL",
  "balanceAttoFIL": "1000000000000000000",
  "dataCap": "1024.50 GiB",
  "dataCapBytes": 1100048498688
}
```

## Error Handling

The command is designed to be resilient and provide partial results even if some operations fail:

- **Balance API Error**: If FIL balance retrieval fails, datacap information is still attempted
- **Datacap API Error**: If datacap retrieval fails, FIL balance information is still provided
- **Invalid Address**: Returns an error for malformed wallet addresses
- **Network Issues**: Returns descriptive error messages for connectivity problems

Example with errors:
```json
{
  "address": "f1invalidaddress",
  "balance": "0",
  "balanceAttoFIL": "",
  "dataCap": "0.00 GiB",
  "dataCapBytes": 0,
  "error": "failed to get wallet balance: invalid wallet address format"
}
```

## Use Cases

### Storage Providers
- Monitor wallet balances before making deals
- Verify client datacap availability
- Track payment wallets across multiple addresses

### Verified Clients
- Check remaining FIL+ datacap allocation
- Monitor wallet balances for deal payments
- Verify datacap transfers and usage

### Developers & Integrators
- Build balance monitoring dashboards
- Automate balance checks in scripts
- Integrate balance data into applications

## Related Commands

- [`singularity wallet list`](../list/) - List all imported wallets
- [`singularity wallet create`](../create/) - Create new wallets
- [`singularity wallet import`](../import/) - Import existing wallets

## Network Configuration

The command uses the Lotus API endpoint configured via:
- `--lotus-api` flag or `LOTUS_API` environment variable
- `--lotus-token` flag or `LOTUS_TOKEN` environment variable (if needed)

For different networks:
- **Mainnet**: `https://api.node.glif.io/rpc/v1` (default)
- **Calibration**: `https://api.calibration.node.glif.io/rpc/v1`
