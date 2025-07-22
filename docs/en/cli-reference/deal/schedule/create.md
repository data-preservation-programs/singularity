# Create a schedule to send out deals to a storage provider with unified flags and defaults

{% code fullWidth="true" %}
```
NAME:
   singularity deal schedule create - Create a schedule to send out deals to a storage provider with unified flags and defaults

USAGE:
   singularity deal schedule create [command options]

DESCRIPTION:
   Create a new deal schedule with unified flags and default values.

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
   --help, -h           show help
   --preparation value  Preparation ID or name
   --provider value     Storage Provider ID to send deals to

   Boost Only

   --http-header value, -H value [ --http-header value, -H value ]  HTTP headers to be passed with the request (i.e. key=value)
   --ipni                                                           Whether to announce the deal to IPNI (default: true)
   --url-template value, -u value                                   URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car

   Deal Proposal

   --duration value, -d value     Duration in epoch or in duration format, i.e. 1500000, 2400h (default: 12840h[535 days])
   --keep-unsealed                Whether to keep unsealed copy (default: true)
   --price-per-deal value         Price in FIL per deal (default: 0)
   --price-per-gb value           Price in FIL per GiB (default: 0)
   --price-per-gb-epoch value     Price in FIL per GiB per epoch (default: 0)
   --start-delay value, -s value  Deal start delay in epoch or in duration format, i.e. 1000, 72h (default: 72h[3 days])
   --verified                     Whether to propose deals as verified (default: true)

   Restrictions

   --allowed-piece-cid value, --piece-cid value [ --allowed-piece-cid value, --piece-cid value ]                      List of allowed piece CIDs in this schedule (default: Any)
   --allowed-piece-cid-file value, --piece-cid-file value [ --allowed-piece-cid-file value, --piece-cid-file value ]  List of files that contains a list of piece CIDs to allow
   --force                                                                                                            Force to send out deals regardless of replication restriction (default: false)
   --max-pending-deal-number value, --pending-number value                                                            Max pending deal number overall for this request, i.e. 100TiB (default: Unlimited)
   --max-pending-deal-size value, --pending-size value                                                                Max pending deal sizes overall for this request, i.e. 1000 (default: Unlimited)
   --total-deal-number value, --total-number value                                                                    Max total deal number for this request, i.e. 1000 (default: Unlimited)
   --total-deal-size value, --total-size value                                                                        Max total deal sizes for this request, i.e. 100TiB (default: Unlimited)

   Scheduling

   --schedule-cron value, --cron value           Cron schedule to send out batch deals (default: disabled)
   --schedule-deal-number value, --number value  Max deal number per triggered schedule, i.e. 30 (default: Unlimited)
   --schedule-deal-size value, --size value      Max deal sizes per triggered schedule, i.e. 500GiB (default: Unlimited)

   Tracking

   --notes value, -n value  Any notes or tag to store along with the request, for tracking purpose

```
{% endcode %}
