# Pcloud

{% code fullWidth="true" %}
```
NAME:
   singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity storage create pcloud - Pcloud

USAGE:
   singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity storage create pcloud [command options] <name> <path>

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

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --root-folder-id
      Fill in for rclone to use a non root folder as its starting point.

   --hostname
      Hostname to connect to.
      
      This is normally set when rclone initially does the oauth connection,
      however you will need to set it by hand if you are using remote config
      with rclone authorize.
      

      Examples:
         | api.pcloud.com  | Original/US region
         | eapi.pcloud.com | EU region

   --username
      Your pcloud username.
            
      This is only required when you want to use the cleanup command. Due to a bug
      in the pcloud API the required API does not support OAuth authentication so
      we have to rely on user password authentication for it.

   --password
      Your pcloud password.


OPTIONS:
   --client-id value      OAuth Client Id. [$CLIENT_ID]
   --client-secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help

   Advanced

   --auth-url value        Auth server URL. [$AUTH_URL]
   --encoding value        The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hostname value        Hostname to connect to. (default: "api.pcloud.com") [$HOSTNAME]
   --password value        Your pcloud password. [$PASSWORD]
   --root-folder-id value  Fill in for rclone to use a non root folder as its starting point. (default: "d0") [$ROOT_FOLDER_ID]
   --token value           OAuth Access Token as a JSON blob. [$TOKEN]
   --token-url value       Token server url. [$TOKEN_URL]
   --username value        Your pcloud username. [$USERNAME]

```
{% endcode %}
