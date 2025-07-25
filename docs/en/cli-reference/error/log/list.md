# List error logs with filtering and pagination

{% code fullWidth="true" %}
```
NAME:
   singularity error log list - List error logs with filtering and pagination

USAGE:
   singularity error log list [command options]

CATEGORY:
   Error Log Management

OPTIONS:
   --entity-type value  Filter by entity type (e.g., deal, preparation, schedule)
   --entity-id value    Filter by entity ID
   --component value    Filter by component (e.g., onboard, deal_schedule)
   --level value        Filter by error level (info, warning, error, critical)
   --event-type value   Filter by event type
   --start-time value   Filter logs after this time (RFC3339 format, e.g., 2023-01-01T00:00:00Z)
   --end-time value     Filter logs before this time (RFC3339 format, e.g., 2023-12-31T23:59:59Z)
   --limit value        Maximum number of logs to return (default: 50)
   --offset value       Number of logs to skip (default: 0)
   --help, -h           show help
```
{% endcode %}
