# Sia Decentralized Cloud

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add sia - Sia Decentralized Cloud

USAGE:
   singularity datasource add sia [command options] <dataset_name> <source_path>

DESCRIPTION:
   --sia-api-url
      Sia daemon API URL, like http://sia.daemon.host:9980.
      
      Note that siad must run with --disable-api-security to open API port for other hosts (not recommended).
      Keep default if Sia daemon runs on localhost.

   --sia-api-password
      Sia Daemon API Password.
      
      Can be found in the apipassword file located in HOME/.sia/ or in the daemon directory.

   --sia-user-agent
      Siad User Agent
      
      Sia daemon requires the 'Sia-Agent' user agent by default for security

   --sia-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for sia

   --sia-api-password value  Sia Daemon API Password. [$SIA_API_PASSWORD]
   --sia-api-url value       Sia daemon API URL, like http://sia.daemon.host:9980. (default: "http://127.0.0.1:9980") [$SIA_API_URL]
   --sia-encoding value      The encoding for the backend. (default: "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot") [$SIA_ENCODING]
   --sia-user-agent value    Siad User Agent (default: "Sia-Agent") [$SIA_USER_AGENT]

```
{% endcode %}
