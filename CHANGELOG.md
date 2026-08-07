# Changelog

All notable changes to Singularity are recorded here. This file tracks **final
releases only** -- release candidates ship as GitHub pre-releases and are not
listed. The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and the project aims to follow [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Each release's full narrative notes -- migration steps, infrastructure readiness,
and the complete PR list -- live in its
[GitHub release](https://github.com/data-preservation-programs/singularity/releases).

## [Unreleased]

## [1.0.0] "Cygnus" -- 2026-08-07

First major release since v0.5.17, superseding the v0.6/v0.7 release-candidate
line. See the [full release notes](https://github.com/data-preservation-programs/singularity/releases/tag/v1.0.0)
for migration steps and the PDP/DDO infrastructure-readiness matrix.

### Highlights

- Experimental **PDP** warm-storage and **DDO** cold-storage deal types
  (`--deal-type pdp|ddo`) on the FWSS registry; legacy `market` (f05) remains
  the default.
- **Trustless IPFS gateway** (`/ipfs/`) on the content-provider, with parallel
  span-read block retrieval.
- **Wallet/Actor model split**: private keys move out of the database into a
  filesystem keystore (`singularity wallet export-keys`).
- **Bounded-time preparation deletes** via nullable foreign keys plus a reaper
  task, instead of cascades.
- **Devcontainer-based CI on podman**, mirroring the local dev environment.

### Breaking

- **MongoDB and MySQL/MariaDB backends removed** -- export to PostgreSQL before
  upgrading.
- **Wallet keys must be exported to a keystore**; the old `wallets` table is
  split into `wallets` and `actors`, and private keys are dropped from the DB.
- **Schedules carry an explicit `--deal-type`**; `--replication N` is removed
  from `create-batch` (use repeatable `--deal-type` / `--provider`).
- **Foreign keys switched to `SET NULL`** -- `admin init` drops and re-creates
  them; the first run on a large dataset may take several minutes for index
  rebuilds, so run it standalone.
- **API defaults (operator-only)**: default bind is now `127.0.0.1:9090` (was
  `0.0.0.0:9090`; widen with `--bind`), the wildcard CORS middleware is removed,
  and secret-named config values are masked in JSON API responses.

### Fixed

- Large preparation and job deletes no longer stall on FK cascade scans.
- `car_blocks.file_id IS NULL` is treated as valid intermediate data, not an
  orphan.
- Numerous smaller correctness fixes in the deal tracker, dataset worker, and
  API error paths (see the full release notes).

[Unreleased]: https://github.com/data-preservation-programs/singularity/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/data-preservation-programs/singularity/releases/tag/v1.0.0
