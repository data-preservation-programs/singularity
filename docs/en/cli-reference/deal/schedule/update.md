# Update an existing schedule

{% code fullWidth="true" %}
```
NAME:
   singularity deal schedule update - Update an existing schedule

USAGE:
   singularity deal schedule update [command options] <schedule_id>

DESCRIPTION:
   CRON pattern '--schedule-cron': The CRON pattern can either be a descriptor or a standard CRON pattern with optional second field
     Standard CRON:
       ┌───────────── minute (0 - 59)
       │ ┌───────────── hour (0 - 23)
       │ │ ┌───────────── day of the month (1 - 31)
       │ │ │ ┌───────────── month (1 - 12)
       │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday)
       │ │ │ │ │                                   
       │ │ │ │ │
       │ │ │ │ │
       * * * * *

     Optional Second field:
       ┌─────────────  second (0 - 59)
       │ ┌─────────────  minute (0 - 59)
       │ │ ┌─────────────  hour (0 - 23)
       │ │ │ ┌─────────────  day of the month (1 - 31)
       │ │ │ │ ┌─────────────  month (1 - 12)
       │ │ │ │ │ ┌─────────────  day of the week (0 - 6) (Sunday to Saturday)
       │ │ │ │ │ │
       │ │ │ │ │ │
       * * * * * *

     Descriptor:
       @yearly, @annually - Equivalent to 0 0 1 1 *
       @monthly           - Equivalent to 0 0 1 * *
       @weekly            - Equivalent to 0 0 * * 0
       @daily,  @midnight - Equivalent to 0 0 * * *
       @hourly            - Equivalent to 0 * * * *

OPTIONS:
   --help, -h  show help

   Boost Only

   --http-header value, -H value [ --http-header value, -H value ]  HTTP headers to be passed with the request (i.e. key=value). This will replace the existing header values. To remove a header, use --http-header "key="". To remove all headers, use --http-header ""
   --ipni                                                           Whether to announce the deal to IPNI (default: true)
   --url-template value, -u value                                   URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car

   Deal Proposal

   --duration value, -d value     Duration in epoch or in duration format, i.e. 1500000, 2400h
   --keep-unsealed                Whether to keep unsealed copy (default: true)
   --price-per-deal value         Price in FIL per deal (default: 0)
   --price-per-gb value           Price in FIL per GiB (default: 0)
   --price-per-gb-epoch value     Price in FIL per GiB per epoch (default: 0)
   --start-delay value, -s value  Deal start delay in epoch or in duration format, i.e. 1000, 72h
   --verified                     Whether to propose deals as verified (default: true)

   Restrictions

   --allowed-piece-cid value, --piece-cid value [ --allowed-piece-cid value, --piece-cid value ]                      List of allowed piece CIDs in this schedule. Append only.
   --allowed-piece-cid-file value, --piece-cid-file value [ --allowed-piece-cid-file value, --piece-cid-file value ]  List of files that contains a list of piece CIDs to allow. Append only.
   --force                                                                                                            Force to send out deals regardless of replication restriction (default: false)
   --max-pending-deal-number value, --pending-number value                                                            Max pending deal number overall for this request, i.e. 100TiB (default: 0)
   --max-pending-deal-size value, --pending-size value                                                                Max pending deal sizes overall for this request, i.e. 1000
   --total-deal-number value, --total-number value                                                                    Max total deal number for this request, i.e. 1000 (default: 0)
   --total-deal-size value, --total-size value                                                                        Max total deal sizes for this request, i.e. 100TiB

   Scheduling

   --schedule-cron value, --cron value           Cron schedule to send out batch deals
   --schedule-deal-number value, --number value  Max deal number per triggered schedule, i.e. 30 (default: 0)
   --schedule-deal-size value, --size value      Max deal sizes per triggered schedule, i.e. 500GiB

   Tracking

   --notes value, -n value  Any notes or tag to store along with the request, for tracking purpose

```
{% endcode %}
