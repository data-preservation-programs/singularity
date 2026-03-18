# Run the singularity API

{% code fullWidth="true" %}
```
NAME:
   singularity run api - Run the singularity API

USAGE:
   singularity run api [command options]

OPTIONS:
   --no-automigrate  skip automatic database migration and correctness checks on startup; only use if you run 'admin init' on every upgrade or manually before starting daemons (default: false)
   --bind value      Bind address for the API server (default: ":9090")
   --help, -h        show help
```
{% endcode %}
