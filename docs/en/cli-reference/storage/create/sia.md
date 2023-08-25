# Sia Decentralized Cloud

{% code fullWidth="true" %}
```
NAME:
   singularity storage create sia - Sia Decentralized Cloud

USAGE:
   singularity storage create sia [command options] <name> <path>

DESCRIPTION:
   --api_url
      Sia daemon API URL, like http://sia.daemon.host:9980.
      
      Note that siad must run with --disable-api-security to open API port for other hosts (not recommended).
      Keep default if Sia daemon runs on localhost.

   --api_password
      Sia Daemon API Password.
      
      Can be found in the apipassword file located in HOME/.sia/ or in the daemon directory.

   --user_agent
      Siad User Agent
      
      Sia daemon requires the 'Sia-Agent' user agent by default for security

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --api_password value  Sia Daemon API Password. [$API_PASSWORD]
   --api_url value       Sia daemon API URL, like http://sia.daemon.host:9980. (default: "http://127.0.0.1:9980") [$API_URL]
   --help, -h            show help

   Advanced

   --encoding value    The encoding for the backend. (default: "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --user_agent value  Siad User Agent (default: "Sia-Agent") [$USER_AGENT]

```
{% endcode %}
