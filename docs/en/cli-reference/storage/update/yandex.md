# Yandex Disk

{% code fullWidth="true" %}
```
NAME:
   singularity storage update yandex - Yandex Disk

USAGE:
   singularity storage update yandex [command options] <name|id>

DESCRIPTION:
   --client-id
      OAuth Client Id.
      
      Leave blank normally.

   --client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --token
      OAuth Access Token as a JSON blob.

   --auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --hard-delete
      Delete files permanently rather than putting them into the trash.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --client-id value      OAuth Client Id. [$CLIENT_ID]
   --client-secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help

   Advanced

   --auth-url value   Auth server URL. [$AUTH_URL]
   --encoding value   The encoding for the backend. (default: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete      Delete files permanently rather than putting them into the trash. (default: false) [$HARD_DELETE]
   --token value      OAuth Access Token as a JSON blob. [$TOKEN]
   --token-url value  Token server url. [$TOKEN_URL]

```
{% endcode %}
