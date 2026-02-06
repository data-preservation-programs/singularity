# Start a PDP deal tracker that tracks f41 PDP deals for all relevant wallets

{% code fullWidth="true" %}
```
NAME:
   singularity run pdp-tracker - Start a PDP deal tracker that tracks f41 PDP deals for all relevant wallets

USAGE:
   singularity run pdp-tracker [command options]

DESCRIPTION:
   The PDP tracker monitors Proof of Data Possession (PDP) deals on the Filecoin network.
   Unlike legacy f05 market deals, PDP deals use proof sets managed through the PDPVerifier contract
   where data is verified through cryptographic challenges.

   This tracker:
   - Monitors proof sets for tracked wallets
   - Updates deal status based on on-chain proof set state
   - Tracks challenge epochs and live status

   Note: Full functionality requires the go-synapse library integration.
   See: https://github.com/data-preservation-programs/go-synapse

OPTIONS:
   --interval value   How often to check for PDP deal updates (default: 10m0s)
   --lotus-api value  Lotus RPC API endpoint [$LOTUS_API]
   --help, -h         show help
```
{% endcode %}
