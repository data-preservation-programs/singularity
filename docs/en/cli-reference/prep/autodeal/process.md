# Process all ready preparations for auto-deal creation

> **💡 Note:** This is a manual batch processing command. For automated workflows, use the [`onboard` command](../../../README.md#basic-usage) or [`run unified`](../../run/README.md) service which handle batch processing automatically.

This command scans all preparations and creates deal schedules for any that are ready and have auto-deal enabled.

{% code fullWidth="true" %}
```
NAME:
   singularity prep autodeal process - Process all ready preparations for auto-deal creation

USAGE:
   singularity prep autodeal process [command options]

OPTIONS:
   --help, -h  show help
```
{% endcode %}

## Example usage

```bash
# Process all ready preparations in batch
singularity prep autodeal process

# Example output:
# {
#   "message": "Auto-deal processing completed successfully"
# }
```

## What this command does

1. **Scans all preparations** in the database
2. **Identifies ready preparations** that have:
   - `autoCreateDeals: true`
   - All jobs completed
   - No existing deal schedule
3. **Creates deal schedules** for each ready preparation
4. **Reports results** of the batch operation

## When to use this command

- **Manual batch operations**: When you want to process multiple preparations at once
- **Scheduled processing**: In cron jobs or other automated scripts
- **Recovery scenarios**: After system downtime to catch up on pending deal creations
- **Testing**: To verify that multiple preparations are processed correctly

## Alternative approaches

For continuous processing, consider:

```bash
# Automated unified service (recommended)
singularity run unified --max-workers 5

# Single preparation onboarding
singularity onboard --name "dataset" --source "/data" --enable-deals
```
