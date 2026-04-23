# Migrating from legacy deals

This guide is for operators moving from legacy market deals (`market`) to the contract-backed deal types now supported by Singularity:

- PDP deals (`pdp`)
- DDO allocation deals (`ddo`)

If your program refers to the second path as PoRep onboarding or allocation-backed deals, Singularity exposes it as `ddo`.

## What `admin init` does

Run:

```bash
singularity admin init
```

This upgrades the schema and backfills missing `deal_type` values on older rows:

- legacy deals are marked as `market`
- legacy schedules are marked as `market`

This is a metadata migration only. It does **not** convert existing market deals into PDP or DDO deals.

## What stays the same

- Existing market deals remain market deals.
- Existing market schedules continue to run as market schedules unless you change them.
- Existing prep data, wallets, providers, and CAR generation do not need to be recreated just because deal types were added.

## Recommended migration path

The safest migration is to create a new schedule for the new deal type rather than mutating an active legacy schedule in place.

1. Run `singularity admin init`.
2. List your existing schedules and identify the legacy market schedule you want to replace.
3. Confirm the preparation has a wallet attached.
4. Pause the legacy schedule before enabling the new one.
5. Create a fresh PDP or DDO schedule against the same preparation.
6. Run the workers needed for that deal type.
7. Verify new deals are being created and tracked.
8. Remove the old schedule only after the replacement path is behaving as expected.

## Migrating to PDP

Use PDP when you want proof-set based contract deals.

Before creating a PDP schedule:

- run `deal-pusher` with `--eth-rpc`
- attach a wallet that can sign FEVM transactions
- make sure the preparation's piece sizes fit the current PDP limit

Current PDP caveats:

- accepted piece CID formats are legacy CommP and CommPv2
- the current proof-set piece-size cap is `1,065,353,216` bytes

Example:

```bash
singularity run deal-pusher --eth-rpc "$ETH_RPC_URL"

singularity run pdp-tracker --eth-rpc "$ETH_RPC_URL"

singularity deal schedule create \
  --preparation 1 \
  --provider t01000 \
  --deal-type pdp \
  --piece-cid bafkzcibd... \
  --max-pending-deal-number 1 \
  --total-deal-number 1
```

For a real-network walkthrough, see [PDP Calibnet E2E Runbook](../topics/pdp-calibnet-e2e.md).

## Migrating to DDO

Use DDO when you want allocation-based contract deals.

Before creating a DDO schedule:

- run `deal-pusher` with `--eth-rpc`, `--ddo-contract`, `--ddo-payments-contract`, and `--ddo-payment-token`
- run `deal-tracker` with `--eth-rpc` and `--ddo-contract`
- attach a wallet that can sign FEVM transactions
- fund the wallet with FIL for gas and with the configured payment token for deals — `deal-pusher` logs a warning at schedule startup if the FIL balance is below the gas threshold or if there are no payment tokens and no deposited funds; in the latter case the on-chain deposit/approval step will fail at the first deal
- make sure the provider is registered and active in the DDO contract
- make sure the provider supports the payment token you configured
- provide a non-empty `--url-template` so the storage provider can fetch each piece

Example:

```bash
singularity run deal-pusher \
  --eth-rpc "$ETH_RPC_URL" \
  --ddo-contract "$DDO_CONTRACT_ADDRESS" \
  --ddo-payments-contract "$DDO_PAYMENTS_CONTRACT_ADDRESS" \
  --ddo-payment-token "$DDO_PAYMENT_TOKEN"

singularity run deal-tracker \
  --eth-rpc "$ETH_RPC_URL" \
  --ddo-contract "$DDO_CONTRACT_ADDRESS"

singularity deal schedule create \
  --preparation 1 \
  --provider t01000 \
  --deal-type ddo \
  --url-template "https://downloads.example.com/piece/{PIECE_CID}.car" \
  --max-pending-deal-number 10 \
  --total-deal-number 100
```

For DDO, the most important operational difference is that `deal-tracker` is responsible for moving deals from `proposed` to `active` once the allocation is activated on-chain.

## In-place schedule updates

Singularity does allow `deal_type` updates on schedules, but use that path carefully.

- Existing deal rows keep their original `deal_type`.
- DDO schedules require a non-empty `url_template`.
- Contract-backed deal types have different runtime dependencies than legacy market deals.

If you need a low-risk migration, create a new schedule instead of editing the old one.

## Rollback

If the new path is not behaving correctly:

1. Pause the PDP or DDO schedule.
2. Re-enable the legacy market schedule.
3. Fix worker configuration or contract settings.
4. Retry with a new replacement schedule once the environment is correct.
