# Deal State Management CLI Commands

This document describes the new CLI commands implemented for deal state management as part of issue #573.

## Overview

The state management CLI provides comprehensive tools for monitoring, exporting, and repairing deal state changes in Singularity. These commands enable operators to:

- View and filter deal state changes with comprehensive query options
- Export state change history to standard formats (CSV, JSON)
- Perform manual recovery and repair operations for deals
- Get detailed statistics about deal state changes

## Commands Structure

All state management commands are organized under the `state` subcommand:

```bash
singularity state <subcommand> [options]
```

## Available Commands

### 1. List State Changes (`list` / `ls`)

Lists deal state changes with optional filtering and pagination.

#### Usage
```bash
singularity state list [options]
```

#### Options
- `--deal-id <id>`: Filter by specific deal ID
- `--state <state>`: Filter by deal state (proposed, published, active, expired, proposal_expired, rejected, slashed, error)
- `--provider <provider>`: Filter by storage provider ID (e.g., f01234)
- `--client <address>`: Filter by client wallet address
- `--start-time <time>`: Filter changes after this time (RFC3339 format)
- `--end-time <time>`: Filter changes before this time (RFC3339 format)
- `--offset <number>`: Number of records to skip for pagination (default: 0)
- `--limit <number>`: Maximum number of records to return (default: 100)
- `--order-by <field>`: Field to sort by (timestamp, dealId, newState, providerId, clientAddress, default: timestamp)
- `--order <direction>`: Sort order (asc, desc, default: desc)
- `--export <format>`: Export format (csv, json)
- `--output <path>`: Output file path for export

#### Examples
```bash
# List all state changes
singularity state list

# List state changes for a specific deal
singularity state list --deal-id 123

# List active deals from a specific provider
singularity state list --state active --provider f01234

# List state changes in a time range
singularity state list --start-time 2023-01-01T00:00:00Z --end-time 2023-12-31T23:59:59Z

# Export to CSV
singularity state list --export csv --output my-state-changes.csv

# Paginated results
singularity state list --limit 50 --offset 100
```

### 2. Get Deal State Changes (`get`)

Retrieves all state changes for a specific deal, ordered by timestamp.

#### Usage
```bash
singularity state get <deal-id> [options]
```

#### Options
- `--export <format>`: Export format (csv, json)
- `--output <path>`: Output file path for export

#### Examples
```bash
# Get state changes for deal 123
singularity state get 123

# Export deal state changes to JSON
singularity state get 123 --export json --output deal-123-history.json
```

### 3. State Statistics (`stats`)

Retrieves comprehensive statistics about deal state changes.

#### Usage
```bash
singularity state stats
```

#### Example Output
```json
{
  "totalStateChanges": 1250,
  "stateDistribution": {
    "proposed": 300,
    "published": 250,
    "active": 500,
    "expired": 150,
    "proposal_expired": 30,
    "rejected": 15,
    "slashed": 3,
    "error": 2
  },
  "recentActivity": {
    "last24Hours": 45,
    "last7Days": 320,
    "last30Days": 890
  },
  "providerStats": {
    "totalProviders": 25,
    "topProviders": [
      {"providerId": "f01234", "stateChanges": 125},
      {"providerId": "f05678", "stateChanges": 98}
    ]
  },
  "clientStats": {
    "totalClients": 15,
    "topClients": [
      {"clientAddress": "f1abcdef", "stateChanges": 200}
    ]
  }
}
```

### 4. Repair Operations (`repair`)

Provides manual recovery and repair capabilities for deal state management.

#### Subcommands

##### Force State Transition (`force-transition`)

Forces a deal to transition to a new state. **Use with caution!**

```bash
singularity state repair force-transition <deal-id> <new-state> [options]
```

**Options:**
- `--reason <text>`: Reason for the forced transition (default: "Manual repair operation")
- `--epoch <number>`: Filecoin epoch height for the state change
- `--sector-id <id>`: Storage provider sector ID
- `--dry-run`: Show what would be done without making changes

**Examples:**
```bash
# Force deal to active state
singularity state repair force-transition 123 active --reason "Manual activation after verification"

# Dry run to see what would happen
singularity state repair force-transition 123 active --dry-run
```

##### Reset Error Deals (`reset-error-deals`)

Resets deals in error state to allow retry operations.

```bash
singularity state repair reset-error-deals [options]
```

**Options:**
- `--deal-id <id>`: Specific deal IDs to reset (can be specified multiple times)
- `--provider <provider>`: Reset error deals for a specific provider
- `--reset-to-state <state>`: State to reset deals to (default: proposed)
- `--limit <number>`: Maximum number of deals to reset (default: 100)
- `--dry-run`: Show what would be done without making changes

**Examples:**
```bash
# Reset all error deals to proposed state
singularity state repair reset-error-deals

# Reset specific deals
singularity state repair reset-error-deals --deal-id 123 --deal-id 456

# Reset error deals for specific provider
singularity state repair reset-error-deals --provider f01234

# Dry run to see affected deals
singularity state repair reset-error-deals --dry-run
```

##### Cleanup Orphaned Changes (`cleanup-orphaned-changes`)

Removes state change records that reference deals that no longer exist.

```bash
singularity state repair cleanup-orphaned-changes [options]
```

**Options:**
- `--dry-run`: Show what would be deleted without making changes

**Examples:**
```bash
# Clean up orphaned state changes
singularity state repair cleanup-orphaned-changes

# See what would be cleaned up
singularity state repair cleanup-orphaned-changes --dry-run
```

## Export Formats

### CSV Format

CSV exports include the following columns:
- ID: State change record ID
- DealID: Deal ID
- PreviousState: Previous deal state
- NewState: New deal state
- Timestamp: When the change occurred
- EpochHeight: Filecoin epoch (if available)
- SectorID: Storage provider sector ID (if available)
- ProviderID: Storage provider ID
- ClientAddress: Client wallet address
- Metadata: Additional metadata as JSON

### JSON Format

JSON exports include metadata about the export and an array of state change objects:

```json
{
  "metadata": {
    "exportTime": "2023-07-24T10:30:00Z",
    "totalCount": 150
  },
  "stateChanges": [
    {
      "id": 1,
      "dealId": 123,
      "previousState": "proposed",
      "newState": "published",
      "timestamp": "2023-07-24T10:00:00Z",
      "epochHeight": 123456,
      "sectorId": "sector-123",
      "providerId": "f01234",
      "clientAddress": "f1abcdef",
      "metadata": "{\"reason\":\"Deal accepted\"}"
    }
  ]
}
```

## Common Use Cases

### 1. Monitoring Deal Progress

```bash
# Check recent state changes
singularity state list --limit 20

# Monitor specific provider's deals
singularity state list --provider f01234 --limit 10

# Check deals in error state
singularity state list --state error
```

### 2. Troubleshooting Failed Deals

```bash
# Get complete history for a failed deal
singularity state get 123

# Find all error deals for investigation
singularity state list --state error --export csv --output error-deals.csv

# Reset error deals after fixing underlying issues
singularity state repair reset-error-deals --provider f01234
```

### 3. Reporting and Analytics

```bash
# Get overall statistics
singularity state stats

# Export monthly report
singularity state list --start-time 2023-07-01T00:00:00Z --end-time 2023-07-31T23:59:59Z --export json --output july-report.json

# Track provider performance
singularity state list --provider f01234 --export csv --output provider-performance.csv
```

### 4. Emergency Recovery

```bash
# Force stuck deal to correct state (use carefully!)
singularity state repair force-transition 123 active --reason "Recovery after manual verification"

# Reset multiple error deals
singularity state repair reset-error-deals --limit 50 --provider f01234

# Clean up database inconsistencies
singularity state repair cleanup-orphaned-changes
```

## Database Requirements

These commands require the Singularity database to be properly initialized with the deal state change tracking tables. Ensure that:

1. The database connection string is properly configured
2. Database migrations have been run (`singularity admin init`)
3. The `deal_state_changes` table exists and is populated by the deal tracking services

## Security Considerations

- **Repair operations are powerful**: The `repair` subcommands can modify deal states and should be used with caution
- **Use dry-run**: Always test repair operations with `--dry-run` first
- **Backup before bulk operations**: Consider backing up your database before performing bulk repair operations
- **Audit trail**: All repair operations create state change records for audit purposes

## Testing

The implementation includes comprehensive test coverage:

- **Unit tests**: Individual command functionality and edge cases
- **Integration tests**: End-to-end workflows and database interactions
- **Export tests**: Verification of CSV and JSON export formats
- **Error handling tests**: Validation of error conditions and user input

To run the tests:

```bash
go test ./cmd/statechange/...
```

## Implementation Details

The state management CLI is implemented with:

- **Modular design**: Each command is in its own file with dedicated test files
- **Consistent interface**: All commands follow the same pattern for options and output
- **Error handling**: Comprehensive error messages and validation
- **Export capabilities**: Pluggable export system supporting multiple formats
- **Database integration**: Direct integration with GORM models and state tracking services
- **Audit logging**: All repair operations are logged with metadata

The commands integrate with the existing Singularity architecture and use the same database models and service layers as the rest of the application.