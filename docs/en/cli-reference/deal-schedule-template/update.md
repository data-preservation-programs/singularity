# Update an existing deal template

{% code fullWidth="true" %}
```
NAME:
   singularity deal-schedule-template update - Update an existing deal template

USAGE:
   singularity deal-schedule-template update [command options] <template_id_or_name>

CATEGORY:
   Deal Template Management

DESCRIPTION:
   Update an existing deal template with new values. Only specified flags will be updated.

   Key flags:
     --name               New name for the template
     --provider           Storage Provider ID (e.g., f01234)
     --duration           Deal duration (e.g., 12840h)
     --start-delay        Deal start delay (e.g., 72h)
     --verified           Propose deals as verified
     --keep-unsealed      Keep unsealed copy
     --ipni               Announce deals to IPNI
     --http-header        HTTP headers (key=value)
     --allowed-piece-cid  List of allowed piece CIDs
     --allowed-piece-cid-file File with allowed piece CIDs

   Piece CID Handling:
     By default, piece CIDs are merged with existing ones. 
     Use --replace-piece-cids to completely replace the existing list.

   See --help for all options.

OPTIONS:
   --allowed-piece-cid value [ --allowed-piece-cid value ]  List of allowed piece CIDs for this template
   --allowed-piece-cid-file value                           File containing list of allowed piece CIDs
   --description value                                      Description of the deal template
   --duration value                                         Duration for storage deals (e.g., 12840h for 535 days) (default: 0s)
   --force                                                  Force deals regardless of replication restrictions (default: false)
   --help, -h                                               show help
   --http-header value [ --http-header value ]              HTTP headers to be passed with the request (key=value format)
   --ipni                                                   Whether to announce deals to IPNI (default: false)
   --keep-unsealed                                          Whether to keep unsealed copy of deals (default: false)
   --name value                                             New name for the deal template
   --notes value                                            Notes or tags for tracking purposes
   --price-per-deal value                                   Price in FIL per deal for storage deals (default: 0)
   --price-per-gb value                                     Price in FIL per GiB for storage deals (default: 0)
   --price-per-gb-epoch value                               Price in FIL per GiB per epoch for storage deals (default: 0)
   --provider value                                         Storage Provider ID (e.g., f01000)
   --replace-piece-cids                                     Replace existing piece CIDs instead of merging (use with --allowed-piece-cid or --allowed-piece-cid-file) (default: false)
   --start-delay value                                      Start delay for storage deals (default: 0s)
   --url-template value                                     URL template for deals
   --verified                                               Whether deals should be verified (default: false)

   Restrictions

   --max-pending-deal-number value  Max pending deal number overall (0 = unlimited) (default: 0)
   --max-pending-deal-size value    Max pending deal sizes overall (e.g., 1000GiB, 0 = unlimited)
   --total-deal-number value        Max total deal number for this template (0 = unlimited) (default: 0)
   --total-deal-size value          Max total deal sizes for this template (e.g., 100TiB, 0 = unlimited)

   Scheduling

   --schedule-cron value         Cron schedule to send out batch deals (e.g., @daily, @hourly, '0 0 * * *')
   --schedule-deal-number value  Max deal number per triggered schedule (0 = unlimited) (default: 0)
   --schedule-deal-size value    Max deal sizes per triggered schedule (e.g., 500GiB, 0 = unlimited)

```
{% endcode %}
