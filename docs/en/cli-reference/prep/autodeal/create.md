# Create automatic deal schedule for a specific preparation

> **💡 Note:** This is a manual override command. For automated workflows, use the [`onboard` command](../../../README.md#basic-usage) which creates deal schedules automatically when jobs complete.

This command manually triggers the creation of a deal schedule for a specific preparation that has auto-deal enabled and all jobs completed.

{% code fullWidth="true" %}
```
NAME:
   singularity prep autodeal create - Create automatic deal schedule for a specific preparation

USAGE:
   singularity prep autodeal create [command options]

OPTIONS:
   --preparation value  Preparation ID or name
   --help, -h           show help
```
{% endcode %}

## Example usage

```bash
# Manually create deal schedule for a preparation
singularity prep autodeal create --preparation "my-dataset"

# Example output:
# {
#   "id": 123,
#   "preparation_id": 45,
#   "provider": "f01234",
#   "verified": true,
#   "price_per_gb": 0.0000001,
#   "duration": "8760h0m0s"
# }
```

## When to use this command

- **Manual timing control**: When you want to control exactly when deals are created
- **Debugging**: To test auto-deal functionality manually
- **Custom workflows**: When integrating with external systems that need manual triggering
- **Recovery**: If automatic deal creation failed and you need to retry manually

## Prerequisites

Before using this command, ensure:
- The preparation has `autoCreateDeals: true` 
- All jobs for the preparation are complete
- No deal schedule already exists for the preparation
- Required deal parameters are configured (provider, pricing, etc.)
