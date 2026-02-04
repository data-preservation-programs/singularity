# Changelog - v0.7.0

## Migration Guide

### From v0.5.x / v0.6.x

#### Database Migration

The database schema has changed significantly. Run the following after upgrading:

```bash
singularity admin init
```

This will auto-migrate the schema. Key changes:
- Foreign keys changed from CASCADE to SET NULL for performance
- New indexes added for job and preparation cleanup

`admin init` is generally cheap to run if schema is healthy so it's advisable to automate as a pre-start script in orchestration, however the initial migration for 0.7.0 may need some time due to index rebuilds so it is advisable to halt other services and run `admin init` standalone in this one case.

#### Breaking: MongoDB Removed

MongoDB backend import is no longer supported. If you were using MongoDB (singularity v1):

1. Export your data from v1
2. Upgrade to v0.6.0-RC2 and import to PostgreSQL/YugabyteDB (MySQL strongly discouraged, sqlite only for very small deployments)
3. Then upgrade to v0.7.0

#### Sequence Reset (PostgreSQL/YugabyteDB)

**Automatic**: `singularity admin init` now detects and fixes stale sequences automatically. This handles the common case of importing data with explicit IDs (e.g., from MySQL backup).

**Manual**: If needed, a standalone script is available at `scripts/fix-sequences-after-import.sql`.

#### Piece Type Classification

**Automatic**: `singularity admin init` now infers `piece_type` for CAR files that predate this column, classifying them as `data` (contains file content) or `dag` (directory metadata only). This is a no-op if all pieces are already labeled.

**Manual**: A standalone script is available at `scripts/infer_piece_type.sql` for inspection or manual runs.

#### Curio Compatibility (DAG Pieces)

If DAG pieces were rejected by Curio, you will need to regenerate them:

1. Run `singularity admin init` to classify existing pieces
2. Delete the rejected DAG pieces using `singularity prep delete-piece`
3. Re-run `singularity prep start-daggen` on affected preparations to create Curio-compatible pieces

---

## Features

- **prep delete-piece command** - Delete specific pieces from a preparation (#602)
- **Small piece padding** - Pad very small pieces with literal zeroes for curio compatibility (#595)

## Bug Fixes

- **S3 car file renames** - Fix car file renames on S3-type remotes (#598)
- **Remote storage padding** - Use local tmpfile for padding small pieces on remote/non-appendable storage (#597)
- **Deadlocks resolved** - Fix remaining deadlocks in prep deletion and job claiming (#596)
- **Download flags** - Harmonize download flags across commands (#599, closes #460)
- **FK re-scanning** - Avoid re-scanning known valid foreign keys during cleanup (#601)
- **docker-compose env var** - Fix singularity_init container environment variable (#526)
- **prep add-piece CLI** - Fix parser bug treating positional args as subcommands; auto-lookup existing pieces by CID for piece consolidation workflows (#607)

## Database & Performance

- **Nullable FKs for massive deletes** - Use nullable foreign keys and pointers for fast bulk deletion; orphaned objects cleaned up via reaper task in batches (#600)
- **Indexes for FK cleanup** - Add indexes on jobs/preps foreign keys for faster operations
- **SKIP LOCKED for job claiming** - Prevent deadlocks during concurrent job acquisition
- **Deferred association loading** - Load associations after claiming to reduce lock contention

## Infrastructure

- **MongoDB removed** - Remove legacy MongoDB support (#588)
- **Devcontainer CI** - Add devcontainer-based CI workflow with PostgreSQL and MySQL (#586)
- **Distroless runner** - Use distroless container image for smaller footprint (#592)
- **Dependency upgrades** - Significantly update gorm, boxo, rclone, libp2p, and security packages (#591, #589)
- **Swagger fixes** - Remove hardcoded host, dedupe storage type fields, bump go-swagger to v0.33.1 (#605, #606, #608)

## Breaking Changes

- MongoDB backend no longer supported

---

## PRs Included

| PR | Title |
|----|-------|
| #608 | bump go-swagger to v0.33.1 |
| #607 | fix add-piece to lookup existing pieces by CID |
| #606 | dedupe fields in storage types codegen |
| #605 | remove hardcoded @host in swagger |
| #602 | feat: add prep delete-piece command |
| #601 | hotfix: avoid re-scanning known valid FKs |
| #600 | give in and use nullable FKs/pointers for massive deletes |
| #599 | harmonize download flags, closes #460 |
| #598 | fix: correctly rename cars on S3 type remotes |
| #597 | hotfix: use local tmpfile for padding small pieces |
| #596 | fix remaining deadlocks + prep deletion |
| #595 | Pad very small pieces with literal zeroes |
| #592 | update runner image |
| #591 | improve lotus API test survivability |
| #589 | chore/update |
| #588 | remove legacy mongodb |
| #586 | Devcontainer-based CI flow |
| #526 | fix docker-compose.yml singularity_init container env var |
| #519 | chore: bump version to v0.6.0-RC3 |
| #516 | Fix: restore version.json to v0.6.0-RC2 |
