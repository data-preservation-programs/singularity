# Create a new SP Pool with default deal parameters

{% code fullWidth="true" %}
```
NAME:
   singularity deal sp-pool create - Create a new SP Pool with default deal parameters

USAGE:
   singularity deal sp-pool create [command options]

OPTIONS:
   --help, -h     show help
   --name value   Unique name for the pool
   --notes value  Notes

   Boost Only

   --http-header value, -H value [ --http-header value, -H value ]  HTTP headers to be passed with the request (i.e. key=value)
   --ipni                                                           Whether to announce the deal to IPNI (default: true)
   --url-template value, -u value                                   URL template with PIECE_CID placeholder for boost to fetch the CAR file

   Deal Proposal

   --duration value, -d value     Duration in epoch or duration format, i.e. 1500000, 2400h (default: 12840h[535 days])
   --keep-unsealed                Whether to keep unsealed copy (default: true)
   --price-per-deal value         Price in FIL per deal (default: 0)
   --price-per-gb value           Price in FIL per GiB (default: 0)
   --price-per-gb-epoch value     Price in FIL per GiB per epoch (default: 0)
   --start-delay value, -s value  Deal start delay in epoch or duration format, i.e. 1000, 72h (default: 72h[3 days])
   --verified                     Whether to propose deals as verified (default: true)

   Restrictions

   --force                                                  Force to send out deals regardless of replication restriction (default: false)
   --max-pending-deal-number value, --pending-number value  Max pending deal number overall (default: Unlimited)
   --max-pending-deal-size value, --pending-size value      Max pending deal sizes overall, i.e. 100TiB (default: Unlimited)

   Scheduling

   --schedule-cron value, --cron value           Cron schedule to send out batch deals (default: disabled)
   --schedule-cron-perpetual                     Whether the cron schedule runs indefinitely (default: false)
   --schedule-deal-number value, --number value  Max deal number per triggered schedule (default: Unlimited)
   --schedule-deal-size value, --size value      Max deal sizes per triggered schedule, i.e. 500GiB (default: Unlimited)

```
{% endcode %}
