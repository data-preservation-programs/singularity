# Create a new deal template with unified flags and defaults

{% code fullWidth="true" %}
```
NAME:
   singularity deal-schedule-template create - Create a new deal template with unified flags and defaults

USAGE:
   singularity deal-schedule-template create [command options]

DESCRIPTION:
   Create a new deal template using the same flags and default values as deal schedule create.

   Key flags:
     --provider           Storage Provider ID (e.g., f01234)
     --duration           Deal duration (default: 12840h)
     --start-delay        Deal start delay (default: 72h)
     --verified           Propose deals as verified (default: true)
     --keep-unsealed      Keep unsealed copy (default: true)
     --ipni               Announce deals to IPNI (default: true)
     --http-header        HTTP headers (key=value)
     --allowed-piece-cid  List of allowed piece CIDs
     --allowed-piece-cid-file File with allowed piece CIDs

   See --help for all options.

OPTIONS:
   --allowed-piece-cid value [ --allowed-piece-cid value ]  List of allowed piece CIDs for this template
   --allowed-piece-cid-file value                           File containing list of allowed piece CIDs
   --duration value                                         Duration for storage deals (e.g., 12840h for 535 days) (default: 12840h0m0s)
   --force                                                  Force deals regardless of replication restrictions (overrides max pending/total deal limits and piece CID restrictions) (default: false)
   --help, -h                                               show help
   --http-header value [ --http-header value ]              HTTP headers to be passed with the request (key=value format)
   --ipni                                                   Whether to announce deals to IPNI (default: true)
   --keep-unsealed                                          Whether to keep unsealed copy of deals (default: true)
   --name value                                             Name of the deal template
   --notes value                                            Notes or tags for tracking purposes
   --price-per-deal value                                   Price in FIL per deal for storage deals (default: 0)
   --price-per-gb value                                     Price in FIL per GiB for storage deals (default: 0)
   --price-per-gb-epoch value                               Price in FIL per GiB per epoch for storage deals (default: 0)
   --provider value                                         Storage Provider ID (e.g., f01000)
   --start-delay value                                      Start delay for storage deals (default: 72h0m0s)
   --url-template value                                     URL template for deals
   --verified                                               Whether deals should be verified (default: true)

   Restrictions

   --max-pending-deal-number value  Max pending deal number overall (0 = unlimited) (default: 0)
   --max-pending-deal-size value    Max pending deal sizes overall (e.g., 1000GiB, 0 = unlimited) (default: "0")
   --total-deal-number value        Max total deal number for this template (0 = unlimited) (default: 0)
   --total-deal-size value          Max total deal sizes for this template (e.g., 100TiB, 0 = unlimited) (default: "0")

   Scheduling

   --schedule-cron value         Cron schedule to send out batch deals (e.g., @daily, @hourly, '0 0 * * *')
   --schedule-deal-number value  Max deal number per triggered schedule (0 = unlimited) (default: 0)
   --schedule-deal-size value    Max deal sizes per triggered schedule (e.g., 500GiB, 0 = unlimited) (default: "0")

```
{% endcode %}
