# PDP Calibnet End-to-End Runbook (Real Run)

This documents an actual run completed on **2026-02-27** against calibnet using Singularity in this repo.

## Scope

- Create/prepare data
- Create PDP schedule
- Execute on-chain PDP flow (`createDataSet` + `addPieces`)
- Verify resulting schedule/deal state in Singularity DB
- Record concrete on-chain transactions

## Environment

- Repo: `data-preservation-programs/singularity`
- Binary: `./singularity-current`
- Network: calibnet
- DB: `postgres://anjor@127.0.0.1:5432/singularity_v2?sslmode=disable`
- Lotus RPC: `https://api.calibration.node.glif.io/rpc/v1`
- Eth RPC: `https://api.calibration.node.glif.io/rpc/v1`

Common env used:

```bash
export DATABASE_CONNECTION_STRING='postgres://anjor@127.0.0.1:5432/singularity_v2?sslmode=disable'
export LOTUS_API='https://api.calibration.node.glif.io/rpc/v1'
export ETH_RPC_URL='https://api.calibration.node.glif.io/rpc/v1'
export MARKET_DEAL_URL='https://marketdeals-calibration.s3.amazonaws.com/StateMarketDeals.json.zst'
export LOTUS_TEST=1
```

## Data + Preparation

Local source path used:

- `/tmp/pdp-demo-20260227`

Preparation in this run:

- `preparation_id=1`
- source storage name: `pdp-src-v2`
- source attachment id: `1`

Piece CIDs involved:

- CommPv1-style piece (rejected by PDP contract):
  - `baga6ea4seaqiuv5czxqikj7bna2ygzxyfhjm46al2s3lldfbiopu5ay6wnwscmq`
- CommPv2 piece (accepted):
  - `bafkzcibd6adqmxxmxtat74d6vqbczd5v3kvg5mdbhtkuawcyfeqvegqd4skzqqy3`

## Wallet

Imported key file:

- `/Users/anjor/Downloads/calibnet-wallet.key`

Addresses:

- secp/f1: `t1qptkn7el6ui2gibs5kximks52ff7onscwn65hly`
- delegated/f410 (used for PDP tx signing path): `t410fmcffelu4okurnsxqcs352bgqmvxww566spsdeii`
- EVM: `0x608a522E9C72a916CAF014b7dD04D0656f6B77De`

Wallet attached to preparation:

```bash
./singularity-current prep attach-wallet 1 1
```

## Schedules Used

### Schedule 1 (initial)

- `schedule_id=1`
- provider: `t01000`
- deal type: `pdp`
- Final state: `completed`
- Contains historical error from earlier attempts in `error_message`.

### Schedule 2 (successful clean path)

Created specifically with allowed CommPv2 piece:

```bash
./singularity-current deal schedule create \
  --preparation 1 \
  --provider t01000 \
  --deal-type pdp \
  --piece-cid bafkzcibd6adqmxxmxtat74d6vqbczd5v3kvg5mdbhtkuawcyfeqvegqd4skzqqy3 \
  --max-pending-deal-number 1 \
  --total-deal-number 1 \
  --notes auto-pdp-commpv2
```

Then later:

```bash
./singularity-current deal schedule update --max-pending-deal-number 0 2
```

Final state: `completed`.

## On-Chain Transactions (Actual)

Contract:

- `0x85e366Cf9DD2c0aE37E963d9556F5f4718d6417C` (PDPVerifier on calibnet)

Proof set created in this run:

- `proof_set_id=11849`

### Successful `createDataSet`

- Eth tx: `0xb0c98f4d18f401c8d4aeb0042887ae1dea0f74f71a5fd003b51501cdd74f7576`
- Filecoin msg CID: `bafy2bzacechskzxpw7krv3tdoi433hfipqgjczr3ncvztrqfrdpg6qdmbso4k`
- Receipt: `status=0x1`, `blockNumber=0x355a3c` (3496508)
- Event: `DataSetCreated(setId=11849, storageProvider=0x608a522E9C72a916CAF014b7dD04D0656f6B77De)`

### Successful `addPieces`

- Eth tx: `0xb1aacc2469a081743d3a1a8d48af671ff63d5c729422115592e89645e543c7c0`
- Filecoin msg CID: `bafy2bzaceaiqtzx2rjcrzskpshqjpowgrkx7jm2r4zkd7ztrflvhtmcmenqzg`
- Receipt: `status=0x1`, `blockNumber=0x355a90` (3496592)
- Event: `PiecesAdded(setId=11849, pieceCids=[bafkzcibd6adqmxxmxtat74d6vqbczd5v3kvg5mdbhtkuawcyfeqvegqd4skzqqy3])`

### Prior failed `addPieces` (for reference)

- Eth tx: `0x86fbed6d5ac86716992761880da64b52702caf5c7425c5e1bc7d3781af0cb60d`
- Filecoin msg CID: `bafy2bzacec33cynncg55yfyxaj5tv4kdswu26drdxabxsqug7crtmtfop6lze`
- Receipt: `status=0x0`
- `StateSearchMsg` receipt: `ExitCode=7`, `GasUsed=5000000` (fixed gas cap issue)

## Final Singularity DB State

Schedules:

- `schedule_id=1`: `completed` (historical error text retained)
- `schedule_id=2`: `completed` (clean successful path)

Deal rows:

- one PDP deal row exists from `schedule_id=2`
  - `state=proposed`
  - `proof_set_id=11849`
  - `provider=t01000`
  - `wallet_id=1`
  - piece CID bytes correspond to `bafkzcibd6adqmxxmxtat74d6vqbczd5v3kvg5mdbhtkuawcyfeqvegqd4skzqqy3`

## Code Changes Applied During This Run

- `service/dealpusher/pdp_onchain.go`
  - `createDataSet` path updated to send value for sybil fee.
  - listener handling changed (see issue context).
  - `addPieces` path switched to manager-based flow (avoids hard fixed gas cap path).
- `util/piececid.go`
  - Added helper to accept both legacy CommP and CommPv2 CIDs.
- `handler/dataprep/piece.go`
- `handler/deal/schedule/create.go`
- `handler/deal/schedule/update.go`
  - Validation updated to accept CommPv2 piece CIDs.

## How to Re-Run Worker

```bash
./singularity-current run deal-pusher --eth-rpc "$ETH_RPC_URL"
```

Optional tracker:

```bash
./singularity-current run pdp-tracker --eth-rpc "$ETH_RPC_URL"
```
