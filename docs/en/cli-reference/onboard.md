# Complete data onboarding workflow (storage → preparation → scanning → deal creation)

{% code fullWidth="true" %}
```
NAME:
   singularity onboard - Complete data onboarding workflow (storage → preparation → scanning → deal creation)

USAGE:
   singularity onboard [command options]

DESCRIPTION:
   The onboard command provides a unified workflow for complete data onboarding.

   It performs the following steps automatically:
   1. Creates storage connections (if paths provided)
   2. Creates data preparation with deal parameters from a deal template
   3. Starts scanning immediately
   4. Enables automatic job progression (scan → pack → daggen → deals)
   5. Optionally starts managed workers to process jobs

   This is the simplest way to onboard data from source to storage deals.
   
   IMPORTANT: Deal configuration is managed through deal templates (--deal-template-id).
   Create a deal template first using 'singularity deal-template create' with your
   desired deal parameters, then reference it during onboarding for consistent,
   auditable deal configuration.

OPTIONS:
   --auto-create-deals                Enable automatic deal creation after preparation completion (default: true)
   --json                             Output result in JSON format for automation (default: false)
   --max-size value                   Maximum size of a single CAR file (default: "31.5GiB")
   --max-workers value                Maximum number of workers to run (default: 3)
   --name value                       Name for the preparation
   --no-dag                           Disable maintaining folder DAG structure (default: false)
   --output value [ --output value ]  Local output path(s) for CAR files (optional)
   --source value [ --source value ]  Local source path(s) to onboard
   --sp-validation                    Enable storage provider validation (default: false)
   --start-workers                    Start managed workers to process jobs automatically (default: true)
   --timeout value                    Timeout for waiting for completion (0 = no timeout) (default: 0s)
   --wait-for-completion              Wait and monitor until all jobs complete (default: false)
   --wallet-validation                Enable wallet balance validation (default: false)

   Deal Settings

   --deal-template-id value   Deal template ID to use for deal configuration (required when auto-create-deals is enabled)

```
{% endcode %}
