# Send a manual PDP deal on-chain

{% code fullWidth="true" %}
```
NAME:
   singularity deal send-manual-pdp - Send a manual PDP deal on-chain

USAGE:
   singularity deal send-manual-pdp [command options]

DESCRIPTION:
   Create/reuse a proof set and add a piece to it on-chain via PDPVerifier.
     Example: singularity deal send-manual-pdp --client f1xxx --provider t410fxxx --piece-cid bagaxxxx --piece-size 1048576 --eth-rpc http://localhost:5700/rpc/v1

OPTIONS:
   --client value                Client wallet address (must be imported)
   --provider value              Storage provider f4/t4 address
   --piece-cid value             Piece CID (commp)
   --piece-size value            Piece size in bytes (default: 0)
   --eth-rpc value               FEVM JSON-RPC endpoint [$ETH_RPC_URL]
   --pdp-contract-address value  Override PDPVerifier contract address (for devnet/testing) [$PDP_CONTRACT_ADDRESS]
   --confirmation-depth value    Blocks to wait for tx confirmation (default: 5)
   --help, -h                    show help
```
{% endcode %}
