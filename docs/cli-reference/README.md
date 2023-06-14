# CLI Reference

```
NAME:
   singularity - A tool for large-scale clients with PB-scale data onboarding to Filecoin network

USAGE:
   singularity [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command
   Daemons:
     run  Run different singularity components
   Easy Commands:
     ez-prep  Prepare a dataset from a local path
   Operations:
     admin       Admin commands
     deal        [Alpha] Replication / Deal making management
     dataset     Dataset management
     datasource  Data source management
     wallet      [Alpha] Wallet management
   Utility:
     download  Download a CAR file from the metadata API

GLOBAL OPTIONS:
   --database-connection-string CREATE DATABASE <dbname> DEFAULT CHARACTER SET ascii  Connection string to the database.
      Supported database: sqlite3, postgres, mysql
      Example for postgres  - postgres://user:pass@example.com:5432/dbname
      Example for mysql     - mysql://user:pass@tcp(localhost:3306)/dbname?charset=ascii&parseTime=true
                                Note: the database needs to be created using ascii Character Set:                                CREATE DATABASE <dbname> DEFAULT CHARACTER SET ascii
      Example for sqlite3   - sqlite:/absolute/path/to/database.db
                  or        - sqlite:relative/path/to/database.db
       (default: sqlite:/home/shane/.singularity/singularity.db) [$DATABASE_CONNECTION_STRING]
   --verbose   Enable verbose logging (default: false)
   --json      Enable JSON output (default: false)
   --help, -h  show help
```
