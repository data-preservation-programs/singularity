# Union Storage Support for Data Preparation

This document describes the implementation of union storage support in Singularity, which allows users to treat multiple storage locations as a single source while maintaining control over how pieces are created.

## Overview

Union storage support allows users to:
1. Mount multiple storage locations as a single unified storage
2. Optionally create one piece per upstream folder when using union storage
3. Control piece creation granularity through CLI flags and API parameters

## API Changes

### Data Preparation API

The Data Preparation API has been extended with a new field:

```go
type CreateRequest struct {
    // ... other fields ...
    OnePiecePerUpstream bool `default:"false" json:"onePiecePerUpstream"` // When using union storage, create one piece per upstream folder
}
```

### REST API Endpoints

- `POST /api/v1/preparation`: Create a new preparation with union storage support
  - New parameter: `onePiecePerUpstream` (boolean)

### Command Line Interface

A new flag has been added to relevant commands:

```bash
singularity prep create [command options] [arguments...]
   --one-piece-per-upstream  When using union storage, create one piece per upstream folder (default: false)
```

## Usage Examples

### Basic Union Storage Preparation

```bash
# Create a preparation with default behavior (combine all files)
singularity prep create --name "my_prep" --source union://path/to/storage

# Create a preparation with one piece per upstream folder
singularity prep create --name "my_prep" --source union://path/to/storage --one-piece-per-upstream
```

### API Example

```json
{
    "name": "my_preparation",
    "sourceStorages": ["union_storage"],
    "onePiecePerUpstream": true,
    "maxSize": "32GiB"
}
```

## Implementation Details

1. The feature uses RClone's union storage backend capabilities
2. When `onePiecePerUpstream` is true:
   - Each upstream folder in the union storage is treated as a separate unit
   - One piece is created per upstream folder, regardless of size
   - Each piece gets its own CID
3. When `onePiecePerUpstream` is false (default):
   - Files are combined based on the regular piece size rules
   - Multiple upstream folders may be combined into a single piece

## Error Handling

- Invalid union storage configurations will return appropriate error messages
- If a union storage upstream is inaccessible, the system will handle it gracefully
- Size limits are still enforced even when creating one piece per upstream

## Testing

The implementation includes:
1. Unit tests for union storage scanning logic
2. Integration tests for the complete preparation workflow
3. CLI flag handling tests

## Limitations

- Maximum piece size limits still apply
- Union storage performance characteristics depend on the underlying storage systems
- Not all storage backends support union storage operations
