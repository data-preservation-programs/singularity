# Yandex Disk

{% code fullWidth="true" %}
```
NAME:
   singularity storage create yandex - Yandex Disk

USAGE:
   singularity storage create yandex [command options] <name> <path>

DESCRIPTION:
   --client_id
      OAuth Client Id.
      
      Leave blank normally.

   --client_secret
      OAuth Client Secret.
      
      Leave blank normally.

   --token
      OAuth Access Token as a JSON blob.

   --auth_url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --token_url
      Token server url.
      
      Leave blank to use the provider defaults.

   --hard_delete
      Delete files permanently rather than putting them into the trash.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --client_id value      OAuth Client Id. [$CLIENT_ID]
   --client_secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help

   Advanced

   --auth_url value   Auth server URL. [$AUTH_URL]
   --encoding value   The encoding for the backend. (default: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hard_delete      Delete files permanently rather than putting them into the trash. (default: false) [$HARD_DELETE]
   --token value      OAuth Access Token as a JSON blob. [$TOKEN]
   --token_url value  Token server url. [$TOKEN_URL]

```
{% endcode %}
