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
   2. Creates data preparation with deal parameters
   3. Starts scanning immediately
   4. Enables automatic job progression (scan → pack → daggen → deals)
   5. Optionally starts managed workers to process jobs

   This is the simplest way to onboard data from source to storage deals.

OPTIONS:
   --auto-create-deals                Enable automatic deal creation after preparation completion (default: true)
   --max-size value                   Maximum size of a single CAR file (default: "31.5GiB")
   --max-workers value                Maximum number of workers to run (default: 3)
   --name value                       Name for the preparation
   --no-dag                           Disable maintaining folder DAG structure (default: false)
   --output value [ --output value ]  Local output path(s) for CAR files (optional)
   --source value [ --source value ]  Local source path(s) to onboard
   --start-workers                    Start managed workers to process jobs automatically (default: true)
   --timeout value                    Timeout for waiting for completion (0 = no timeout) (default: 0s)
   --sp-validation                    Enable storage provider validation (default: false)
   --wallet-validation                Enable wallet balance validation (default: false)
   --wait-for-completion              Wait and monitor until all jobs complete (default: false)

   Deal Settings

   --deal-duration value      Duration for storage deals (e.g., 535 days) (default: 12840h0m0s)
   --deal-price-per-gb value  Price in FIL per GiB for storage deals (default: 0)
   --deal-provider value      Storage Provider ID for deals (e.g., f01000)
   --deal-start-delay value   Start delay for storage deals (e.g., 72h) (default: 72h0m0s)
   --deal-verified            Whether deals should be verified (default: false)

```
{% endcode %}
