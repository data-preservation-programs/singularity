# Create a deal schedule

Time to make some deals with storage providers. Start with running the deal pusher service

```
singularity run deal-pusher
```

For PDP (`--deal-type pdp`) schedules, run deal pusher with FEVM RPC configured:

```sh
singularity run deal-pusher --eth-rpc "$ETH_RPC_URL"
```

For DDO (`--deal-type ddo`) schedules, run deal pusher with DDO contract configuration and run deal tracker with DDO tracking enabled:

```sh
singularity run deal-pusher \
  --eth-rpc "$ETH_RPC_URL" \
  --ddo-contract "$DDO_CONTRACT_ADDRESS" \
  --ddo-payments-contract "$DDO_PAYMENTS_CONTRACT_ADDRESS" \
  --ddo-payment-token "$DDO_PAYMENT_TOKEN"

singularity run deal-tracker \
  --eth-rpc "$ETH_RPC_URL" \
  --ddo-contract "$DDO_CONTRACT_ADDRESS"
```

## Send all deals at once

With smaller dataset, you could send all deals to your storage providers all at once. To achieve this, you can use below command

```sh
singularity deal schedule create <preparation> <provider_id>
```

However, if the dataset is large, it may be too much for storage providers to ingest that many deals before the deal proposal expiration, so you can create a schedule

## Send deals with schedule

With the same command, you can create your own schedule to control how fast and how often should the deals be made to storage providers
```sh
singularity deal schedule create -h
```

## PDP caveats

- Accepted PDP piece CID formats: legacy CommP and CommPv2.
- Current PDP proofset piece-size cap: **1 GiB minus FR32 overhead** (`1,065,353,216` bytes).
- On calibnet/devnet, `createDataSet` path requires sybil-fee value and uses compatibility listener handling.

For a full real-network walkthrough with concrete transactions, see:

- [PDP Calibnet End-to-End Runbook](../topics/pdp-calibnet-e2e.md)

## DDO caveats

- DDO schedules use deal type `ddo`.
- DDO schedules require a non-empty `urlTemplate`.
- The provider must be active in the DDO contract and support the configured payment token.
- DDO deals stay in `proposed` state until `deal-tracker` observes allocation activation on-chain.

For a full DDO setup guide, see:

- [DDO Contract Deal Guide](../topics/ddo-contract-deal-guide.md)
