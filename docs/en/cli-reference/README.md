# CLI Reference

{% code fullWidth="true" %}
```
NAME:
   singularity - A tool for large-scale clients with PB-scale data onboarding to Filecoin network.

USAGE:
   singularity [global options] command [command options] [arguments...]

DESCRIPTION:
   Database Backend Support:
     Singularity supports multiple database backend: sqlite3, postgres, mysql5.7+
     Use '--database-connection-string' or $DATABASE_CONNECTION_STRING to specify the database connection string.
       Example for postgres  - postgres://user:pass@example.com:5432/dbname
       Example for mysql     - mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true
       Example for sqlite3   - sqlite:/absolute/path/to/database.db
                   or        - sqlite:relative/path/to/database.db

   Network Support:
     Default settings in Singularity are for Mainnet. You may set below environment values for other network:
       For Calibration network:
         * Set LOTUS_API to https://api.calibration.node.glif.io/rpc/v1
         * Set MARKET_DEAL_URL to https://marketdeals-calibration.s3.amazonaws.com/StateMarketDeals.json.zst
         * Set LOTUS_TEST to 1
       For all other networks:
         * Set LOTUS_API to your network's Lotus API endpoint
         * Set MARKET_DEAL_URL to empty string
         * Set LOTUS_TEST to 0 or 1 based on whether the network address starts with 'f' or 't'
       Switching between different networks with the same database instance is not recommended.

   Logging:
     Singularity uses go-log for logging and can be controlled by below environment variables:
       * GOLOG_LOG_LEVEL  - example values: debug, info, warn, error, dpanic, panic, fatal
       * GOLOG_LOG_FMT    - example values: color, nocolor, json
       * More details can be found at https://github.com/ipfs/go-log


COMMANDS:
   version, v  Print version information
   help, h     Shows a list of commands or help for one command
   Daemons:
     run  run different singularity components
   Operations:
     admin    Admin commands
     deal     Replication / Deal making management
     wallet   Wallet management
     storage  Create and manage storage system connections
     prep     Create and manage dataset preparations
   Utility:
     ez-prep      Prepare a dataset from a local path
     download     Download a CAR file from the metadata API
     extract-car  Extract folders or files from a folder of CAR files to a local directory

GLOBAL OPTIONS:
   --database-connection-string value  Connection string to the database (default: sqlite:./singularity.db) [$DATABASE_CONNECTION_STRING]
   --help, -h                          show help
   --json                              Enable JSON output (default: false)
   --verbose                           Enable verbose output. This will print more columns for the result as well as full error trace (default: false)

   Lotus

   --lotus-api value    Lotus RPC API endpoint (default: "https://api.node.glif.io/rpc/v1") [$LOTUS_API]
   --lotus-test         Whether the runtime environment is using Testnet. (default: false) [$LOTUS_TEST]
   --lotus-token value  Lotus RPC API token [$LOTUS_TOKEN]

```
{% endcode %}
