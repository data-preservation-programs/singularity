# Filecoin Proof Management System

## Overview

The proof management system in Singularity provides comprehensive functionality for storing, verifying, and managing Filecoin storage proofs. It supports both client and storage provider perspectives and includes features for proof verification history tracking.

## Features

- Store proofs associated with storage deals
- Verify proofs with role-based access control
- Track proof verification history
- List proofs by client or provider address
- Manage proof lifecycle and status

## API Endpoints

### Store a New Proof

```http
POST /api/v1/proofs
```

Request body:
```json
{
    "dealId": 123,
    "pieceCid": "bafk2bzaced...",
    "proofBytes": "base64-encoded-proof-data",
    "sectorId": 456,
    "clientAddress": "f01234",
    "providerAddress": "f05678"
}
```

### Verify a Proof

```http
POST /api/v1/proofs/{dealId}/verify
```

Request body:
```json
{
    "verifierAddress": "f09876"
}
```

### List Client Proofs

```http
GET /api/v1/proofs/client/{address}
```

### List Provider Proofs

```http
GET /api/v1/proofs/provider/{address}
```

### Get Verification History

```http
GET /api/v1/proofs/{dealId}/history
```

## CLI Commands

### Store a Proof

```bash
singularity proof store --deal-id 123 \
                       --piece-cid bafk2bzaced... \
                       --proof-file /path/to/proof \
                       --sector-id 456 \
                       --client f01234 \
                       --provider f05678
```

### Verify a Proof

```bash
singularity proof verify --deal-id 123 --verifier f09876
```

### List Proofs

```bash
# List by client
singularity proof list client --address f01234

# List by provider
singularity proof list provider --address f05678
```

### Get Proof History

```bash
singularity proof history --deal-id 123
```

## Database Schema

### Deal Proofs Table

```sql
CREATE TABLE deal_proofs (
    deal_id BIGINT PRIMARY KEY,
    piece_cid TEXT NOT NULL,
    sector_id BIGINT NOT NULL,
    proof_bytes BYTEA NOT NULL,
    client_address TEXT NOT NULL,
    provider_address TEXT NOT NULL,
    proof_status TEXT NOT NULL,
    verification_timestamp TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Proof Verifications Table

```sql
CREATE TABLE proof_verifications (
    id SERIAL PRIMARY KEY,
    deal_id BIGINT NOT NULL,
    verified_by TEXT NOT NULL,
    verification_result BOOLEAN NOT NULL,
    verification_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (deal_id) REFERENCES deal_proofs(deal_id)
);
```

## Development

### Running Tests

```bash
# Run unit tests
make test

# Run integration tests
make integration-test
```

### Adding New Features

1. Define the feature in the service interface
2. Implement the feature in the service implementation
3. Add API endpoint in the handler
4. Add CLI command
5. Add tests
6. Update documentation

### Error Handling

The system uses standard error types for common scenarios:

- Invalid address format
- Invalid piece CID
- Proof not found
- Unauthorized access
- Verification failures

## Security Considerations

1. Address Validation
   - All Filecoin addresses are validated before use
   - Both client and provider addresses must be valid

2. Proof Verification
   - Only authorized verifiers can verify proofs
   - All verification attempts are recorded
   - Verification results are immutable

3. Access Control
   - Role-based access for different operations
   - Separate client and provider views
   - Audit trail for all operations

## Performance Considerations

1. Database Indexes
   - Indexes on client_address and provider_address
   - Index on piece_cid for quick lookups

2. Proof Storage
   - Efficient storage of proof bytes
   - Proper handling of large proofs

3. Concurrent Operations
   - Support for concurrent verifications
   - Thread-safe operations

## Integration Guidelines

1. Service Integration
```go
proofService := service.NewProofService(db)
```

2. API Integration
```go
proofHandler := handler.NewProofHandler(proofService)
proofHandler.RegisterRoutes(router)
```

3. CLI Integration
```go
app.Commands = append(app.Commands, cmd.ProofCmd)
```
