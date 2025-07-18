# Update an existing deal template

{% code fullWidth="true" %}
```
NAME:
   singularity deal-template update - Update an existing deal template

USAGE:
   singularity deal-template update [command options] <template_id_or_name>

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
   --name value                     New name for the deal template
   --description value              Description of the deal template
   --provider value                 Storage Provider ID (e.g., f01000)
   --price-per-gb value             Price in FIL per GiB for storage deals
   --price-per-gb-epoch value       Price in FIL per GiB per epoch for storage deals
   --price-per-deal value           Price in FIL per deal for storage deals
   --duration value                 Duration for storage deals (e.g., 12840h for 535 days)
   --start-delay value              Start delay for storage deals
   --verified                       Whether deals should be verified
   --keep-unsealed                  Whether to keep unsealed copy of deals
   --ipni                           Whether to announce deals to IPNI
   --url-template value             URL template for deals
   --http-header value              HTTP headers to be passed with the request (key=value format)
   --notes value                    Notes or tags for tracking purposes
   --force                          Force deals regardless of replication restrictions
   --allowed-piece-cid value        List of allowed piece CIDs for this template
   --allowed-piece-cid-file value   File containing list of allowed piece CIDs
   --replace-piece-cids             Replace existing piece CIDs instead of merging (use with --allowed-piece-cid or --allowed-piece-cid-file)

   Scheduling:
   --schedule-cron value            Cron schedule to send out batch deals (e.g., @daily, @hourly, '0 0 * * *')
   --schedule-deal-number value     Max deal number per triggered schedule (0 = unlimited) (default: 0)
   --schedule-deal-size value       Max deal sizes per triggered schedule (e.g., 500GiB, 0 = unlimited)

   Restrictions:
   --total-deal-number value        Max total deal number for this template (0 = unlimited) (default: 0)
   --total-deal-size value          Max total deal sizes for this template (e.g., 100TiB, 0 = unlimited)
   --max-pending-deal-number value  Max pending deal number overall (0 = unlimited) (default: 0)
   --max-pending-deal-size value    Max pending deal sizes overall (e.g., 1000GiB, 0 = unlimited)

   --help, -h                       show help
```
{% endcode %}

## Examples

### Update template description
```bash
singularity deal-template update my-template --description "Updated description"
```

### Update multiple fields
```bash
singularity deal-template update my-template \
  --name "new-template-name" \
  --provider "f01234" \
  --price-per-gb 0.002 \
  --verified
```

### Update scheduling configuration
```bash
singularity deal-template update my-template \
  --schedule-cron "@daily" \
  --schedule-deal-number 10 \
  --schedule-deal-size "100GiB"
```

### Add piece CIDs (merge with existing)
```bash
singularity deal-template update my-template \
  --allowed-piece-cid "baga6ea4seaqao7s73y24kcutaosvacpdjgfe5pw76ooefnyqw4ynr3d2y6x2mpq" \
  --allowed-piece-cid "baga6ea4seaqjtcl7dtqxe4vx4kqpb5f2zwz2ydbgshkjlx6jfhpbwjjh3gqjl6a"
```

### Replace all piece CIDs
```bash
singularity deal-template update my-template \
  --allowed-piece-cid-file /path/to/piece-cids.txt \
  --replace-piece-cids
```

### Update by template ID
```bash
singularity deal-template update 1 --description "Updated via ID"
```

## Notes

- **Partial Updates**: Only fields specified with flags will be updated. Unspecified fields remain unchanged.
- **Template Identification**: You can specify templates by either their ID (numeric) or name.
- **Piece CID Handling**: By default, new piece CIDs are merged with existing ones. Use `--replace-piece-cids` to completely replace the list.
- **Validation**: All input parameters are validated before updating. Invalid values will result in an error.
- **Atomic Updates**: Updates are performed atomically - either all changes succeed or none are applied.
- **JSON Output**: The updated template is returned in JSON format for verification.

## Common Use Cases

1. **Adjust Pricing**: Update deal pricing based on market conditions
2. **Change Provider**: Switch to a different storage provider
3. **Update Scheduling**: Modify deal scheduling frequency or limits
4. **Manage Piece CIDs**: Add or replace allowed piece CIDs for content filtering
5. **Template Maintenance**: Update descriptions, notes, or other metadata