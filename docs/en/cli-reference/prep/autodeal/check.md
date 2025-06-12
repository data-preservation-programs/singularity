# Check if a preparation is ready for auto-deal creation

> **💡 Note:** This is a manual override command. For automated workflows, use the [`onboard` command](../../../README.md#basic-usage) which handles readiness checking automatically.

This command checks whether all jobs for a preparation have completed and the preparation is ready for automatic deal schedule creation.

{% code fullWidth="true" %}
```
NAME:
   singularity prep autodeal check - Check if a preparation is ready for auto-deal creation

USAGE:
   singularity prep autodeal check [command options]

OPTIONS:
   --preparation value  Preparation ID or name
   --help, -h           show help
```
{% endcode %}

## Example usage

```bash
# Check if preparation is ready for auto-deal
singularity prep autodeal check --preparation "my-dataset"

# Example output when ready:
# {
#   "preparation": "my-dataset", 
#   "ready": true
# }
```

## When a preparation is "ready"

A preparation is considered ready when:
- All scan jobs are complete
- All pack jobs are complete  
- All daggen jobs are complete (if enabled)
- Auto-deal creation is enabled for the preparation
- No deal schedule already exists for the preparation
