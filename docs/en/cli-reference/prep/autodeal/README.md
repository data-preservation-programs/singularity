# Auto-deal management commands

> **💡 Recommendation:** For most users, the [`onboard` command](../../../README.md#basic-usage) provides a simpler single-command workflow that handles auto-deal creation automatically.

These commands provide **manual override capabilities** for advanced users who need granular control over auto-deal creation.

{% code fullWidth="true" %}
```
NAME:
   singularity prep autodeal - Auto-deal management commands

USAGE:
   singularity prep autodeal command [command options]

COMMANDS:
   create   Create automatic deal schedule for a specific preparation
   process  Process all ready preparations for auto-deal creation
   check    Check if a preparation is ready for auto-deal creation
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```
{% endcode %}

## When to use these commands

- **Advanced workflows**: When you need to control auto-deal timing manually
- **Debugging**: To check preparation readiness or manually trigger deal creation
- **Integration**: For custom scripts that need granular control over deal creation
- **Batch processing**: To process multiple preparations programmatically

## Simple alternative

For most use cases, prefer the unified workflow:

```bash
singularity onboard --name "dataset" --source "/data" --enable-deals --deal-provider "f01234"
```
