# Pcloud

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add pcloud - Pcloud

USAGE:
   singularity datasource add pcloud [command options] <dataset_name> <source_path>

DESCRIPTION:
   --pcloud-auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --pcloud-client-id
      OAuth Client Id.
      
      Leave blank normally.

   --pcloud-client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --pcloud-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --pcloud-hostname
      Hostname to connect to.
      
      This is normally set when rclone initially does the oauth connection,
      however you will need to set it by hand if you are using remote config
      with rclone authorize.
      

      Examples:
         | api.pcloud.com  | Original/US region
         | eapi.pcloud.com | EU region

   --pcloud-password
      Your pcloud password.

   --pcloud-root-folder-id
      Fill in for rclone to use a non root folder as its starting point.

   --pcloud-token
      OAuth Access Token as a JSON blob.

   --pcloud-token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --pcloud-username
      Your pcloud username.
            
      This is only required when you want to use the cleanup command. Due to a bug
      in the pcloud API the required API does not support OAuth authentication so
      we have to rely on user password authentication for it.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)
   --scanning-state value   set the initial scanning state (default: ready)

   Options for pcloud

   --pcloud-auth-url value        Auth server URL. [$PCLOUD_AUTH_URL]
   --pcloud-client-id value       OAuth Client Id. [$PCLOUD_CLIENT_ID]
   --pcloud-client-secret value   OAuth Client Secret. [$PCLOUD_CLIENT_SECRET]
   --pcloud-encoding value        The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PCLOUD_ENCODING]
   --pcloud-hostname value        Hostname to connect to. (default: "api.pcloud.com") [$PCLOUD_HOSTNAME]
   --pcloud-password value        Your pcloud password. [$PCLOUD_PASSWORD]
   --pcloud-root-folder-id value  Fill in for rclone to use a non root folder as its starting point. (default: "d0") [$PCLOUD_ROOT_FOLDER_ID]
   --pcloud-token value           OAuth Access Token as a JSON blob. [$PCLOUD_TOKEN]
   --pcloud-token-url value       Token server url. [$PCLOUD_TOKEN_URL]
   --pcloud-username value        Your pcloud username. [$PCLOUD_USERNAME]

```
{% endcode %}
