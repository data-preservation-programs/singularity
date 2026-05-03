# Send a manual PDP deal via the FWSS-pull flow

{% code fullWidth="true" %}
```
NAME:
   singularity deal send-manual-pdp - Send a manual PDP deal via the FWSS-pull flow

USAGE:
   singularity deal send-manual-pdp [command options]

DESCRIPTION:
   Push a single piece to an SP via Curio's /pdp/piece/pull, then trigger the
   SP's on-chain commit (createDataSet+addPieces if no assembling set yet, or addPieces
   into the existing one). Useful for e2e/diagnostic testing of the FWSS pull path.
     Example: singularity deal send-manual-pdp --client f1xxx --provider t410fxxx --piece-cid bagaxxxx --piece-size 1048576 --eth-rpc http://localhost:5700/rpc/v1 --source-url-base https://static.example.org

OPTIONS:
   --client value           Client wallet address (must be imported)
   --provider value         Storage provider f4/t4 address
   --piece-cid value        Piece CID (commp v1)
   --piece-size value       Padded piece size in bytes (default: 0)
   --eth-rpc value          FEVM JSON-RPC endpoint [$ETH_RPC_URL]
   --source-url-base value  HTTPS base where Curio fetches the piece (sourceUrl = <base>/piece/<pieceCidV2>) [$PDP_SOURCE_URL_BASE]
   --record-keeper value    FWSS contract address. Defaults to network FWSS from go-synapse. [$PDP_RECORD_KEEPER]
   --pull-timeout value     How long to wait for Curio to finish each phase (default: 5m0s)
   --help, -h               show help
```
{% endcode %}
