# DDO Contract Deal Guide

This guide covers Singularity's allocation-based contract deal path, exposed as deal type `ddo`.

If your program refers to this path as PoRep onboarding or allocation-backed contract usage, use `ddo` in Singularity commands and API payloads.

## What DDO changes

Unlike legacy market deals:

- Singularity submits allocations on-chain instead of sending a boost proposal
- the storage provider fetches data from the download URL attached to each allocation
- payment setup happens before allocation submission
- deals become `active` only after `deal-tracker` sees the allocation activate on-chain

## Prerequisites

You need:

- a preparation with CAR files ready to schedule
- a wallet attached to that preparation
- an FEVM RPC endpoint
- the DDO Diamond proxy contract address
- the DDO Payments contract address
- an ERC20 payment token address supported by the target provider
- a reachable download URL template for each piece

The target provider must also be:

- resolvable by Lotus
- active in the DDO contract
- configured to accept the payment token you pass to `deal-pusher`

## Required worker configuration

Start `deal-pusher` with the DDO contract settings:

```bash
singularity run deal-pusher \
  --eth-rpc "$ETH_RPC_URL" \
  --ddo-contract "$DDO_CONTRACT_ADDRESS" \
  --ddo-payments-contract "$DDO_PAYMENTS_CONTRACT_ADDRESS" \
  --ddo-payment-token "$DDO_PAYMENT_TOKEN"
```

Start `deal-tracker` with DDO tracking enabled:

```bash
singularity run deal-tracker \
  --eth-rpc "$ETH_RPC_URL" \
  --ddo-contract "$DDO_CONTRACT_ADDRESS"
```

Without the DDO-enabled `deal-tracker`, DDO deals will be created in `proposed` state but will not transition to `active` when the allocation is sealed.

## Creating a DDO schedule

Create a schedule with:

- `--deal-type ddo`
- a provider address that Lotus can resolve
- a non-empty `--url-template`

Example:

```bash
singularity deal schedule create \
  --preparation 1 \
  --provider t01000 \
  --deal-type ddo \
  --url-template "https://downloads.example.com/piece/{PIECE_CID}.car" \
  --max-pending-deal-number 10 \
  --total-deal-number 100
```

`--url-template` should include `{PIECE_CID}` so each allocation points at the correct CAR file.

## How scheduling works

For each batch, Singularity:

1. validates the provider against the DDO contract
2. checks that each piece size fits the provider's configured min and max limits
3. builds per-piece download URLs from `url_template`
4. ensures deposit and operator approval through the payments contract
5. submits the allocation transaction
6. waits for confirmation depth
7. parses allocation IDs from the receipt
8. writes DDO deal rows into the database

## Runtime flags

The main DDO tuning flags on `deal-pusher` are:

- `--ddo-batch-size`
- `--ddo-confirmation-depth`
- `--ddo-poll-interval`
- `--ddo-term-min`
- `--ddo-term-max`
- `--ddo-expiration-offset`

Defaults today are:

- batch size: `10`
- confirmation depth: `5`
- poll interval: `30s`
- min term: `518400` epochs
- max term: `5256000` epochs
- expiration offset: `172800` epochs

## Observing state

You can inspect DDO deals with the normal deal list APIs and CLI filters:

```bash
singularity deal list --deal-type ddo
```

Expected state flow:

- `proposed`: allocation submitted, waiting for activation
- `active`: allocation activated on-chain and picked up by `deal-tracker`

## Batch DDO schedules

To schedule DDO deals across multiple preparations and providers in one command, use `create-batch` with `--deal-type ddo`:

```bash
singularity deal schedule create-batch \
  --group my-ddo-dataset \
  --preparation prep-a --preparation prep-b \
  --provider t01000 --provider t02000 \
  --deal-type ddo \
  --url-template "https://downloads.example.com/piece/{PIECE_CID}.car"
```

This creates the cross-product of DDO schedules (2 preparations x 1 deal type x 2 providers = 4 schedules), all tagged with the `--group` label for easy management. See [Create a deal schedule](../deal-making/create-a-deal-schedule.md#batch-schedule-creation) for more on batch creation.

## Updating an existing DDO schedule

DDO schedules can be updated, but keep these invariants in mind:

- `deal_type` must remain a valid deal type
- `url_template` cannot be empty while the schedule is `ddo`

If you are migrating from a legacy market schedule, creating a fresh DDO schedule is safer than editing the old one in place.

## Current limitations

- DDO activation tracking is implemented, but expiry or terminal-state tracking is not yet defined.
- The current implementation waits for tx confirmation before writing deal rows; crash recovery for the gap between confirmation and DB insert is not yet hardened.
