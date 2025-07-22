# Proof CLI Commands

The proof commands provide functionality to list and synchronize Filecoin proofs from the blockchain.

## Commands

### `singularity proof list`

Lists proofs with optional filtering and pagination.

#### Usage

```bash
singularity proof list [OPTIONS]
```

#### Options

- `--deal-id <ID>` - Filter proofs by deal ID
- `--proof-type <TYPE>` - Filter proofs by type: `replication`, `spacetime`
- `--provider <PROVIDER>` - Filter proofs by storage provider (e.g., `f01000`)
- `--verified` - Show only verified proofs
- `--unverified` - Show only unverified proofs
- `--limit <NUMBER>` - Limit number of results (default: 100)
- `--offset <NUMBER>` - Offset for pagination (default: 0)

#### Examples

```bash
# List all proofs
singularity proof list

# List proofs for a specific deal
singularity proof list --deal-id 12345

# List only replication proofs
singularity proof list --proof-type replication

# List only spacetime proofs
singularity proof list --proof-type spacetime

# List proofs from a specific provider
singularity proof list --provider f01000

# List only verified proofs
singularity proof list --verified

# List only unverified proofs
singularity proof list --unverified

# List with pagination
singularity proof list --limit 50 --offset 100

# Combined filters
singularity proof list --provider f01000 --proof-type replication --verified
```

#### Output

The command outputs a table with the following columns:
- `ID` - Proof record ID
- `DealID` - Associated deal ID
- `ProofType` - Type of proof (replication/spacetime)
- `MessageID` - Blockchain message CID
- `Height` - Block height
- `Method` - Proof method name
- `Verified` - Whether proof was verified
- `Provider` - Storage provider ID
- `CreatedAt` - When proof was recorded

With `--verbose` flag, additional columns are shown:
- `BlockCID` - Block CID where proof was included
- `SectorID` - Sector ID (if available)
- `ErrorMsg` - Error message (if any)
- `UpdatedAt` - Last update time

### `singularity proof sync`

Synchronizes proofs from the Filecoin blockchain into the local database.

#### Usage

```bash
singularity proof sync [OPTIONS]
```

#### Options

- `--deal-id <ID>` - Sync proofs for specific deal ID
- `--provider <PROVIDER>` - Sync proofs for specific storage provider

#### Examples

```bash
# Sync proofs for all active deals
singularity proof sync

# Sync proofs for a specific deal
singularity proof sync --deal-id 12345

# Sync proofs for a specific provider
singularity proof sync --provider f01000

# Sync proofs for specific provider and deal (both filters applied)
singularity proof sync --provider f01000 --deal-id 12345
```

#### Behavior

- If no options are provided, syncs proofs for all active deals
- If `--deal-id` is provided, syncs proofs for that specific deal
- If `--provider` is provided, syncs proofs for that specific provider
- If both are provided, both filters are applied
- The command looks back 2000 epochs (about 16 hours) for proof messages
- Duplicate proofs are automatically skipped
- Errors for individual messages are logged but don't stop the sync process

#### Output

The command outputs a success message upon completion:
```json
{
  "status": "success"
}
```

## Global Options

These options can be used with any proof command:

- `--json` - Output results in JSON format
- `--verbose` - Show verbose output with additional details
- `--database-connection-string <CONNECTION>` - Database connection string
- `--lotus-api <URL>` - Lotus API endpoint (default: https://api.node.glif.io/rpc/v1)
- `--lotus-token <TOKEN>` - Lotus API token

## Examples

### Basic Usage

```bash
# List first 10 proofs
singularity proof list --limit 10

# Sync proofs for all deals
singularity proof sync

# Check specific deal's proofs
singularity proof list --deal-id 12345 --verbose
```

### JSON Output

```bash
# Get proof data in JSON format
singularity --json proof list --deal-id 12345

# Sync with JSON status
singularity --json proof sync --provider f01000
```

### Filtering Examples

```bash
# Find all failed proofs
singularity proof list --unverified --verbose

# Check replication proofs for a provider
singularity proof list --provider f01000 --proof-type replication

# Paginate through large result sets
singularity proof list --limit 100 --offset 0    # First 100
singularity proof list --limit 100 --offset 100  # Next 100
```

### Monitoring Examples

```bash
# Monitor recent proofs
singularity proof list --limit 20 --verbose

# Sync and then check results
singularity proof sync --provider f01000
singularity proof list --provider f01000 --verbose
```

## Integration with Other Commands

The proof commands work alongside other Singularity commands:

```bash
# List deals first, then check their proofs
singularity deal list --provider f01000
singularity proof list --provider f01000

# Sync proofs after creating schedules
singularity schedule create --provider f01000 ...
singularity proof sync --provider f01000
```

## Error Handling

- Database connection errors are displayed immediately
- Lotus API errors during sync are logged but don't stop the process
- Invalid command line arguments show usage help
- Use `--verbose` flag to see detailed error information

## Performance Notes

- Listing proofs is fast due to database indexes
- Syncing proofs may take time depending on the number of messages
- Use specific filters (deal-id, provider) for faster sync operations
- The sync process looks back 2000 epochs by default