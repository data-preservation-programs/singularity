# Enterprise File Fabric

{% code fullWidth="true" %}
```
NAME:
   singularity storage create filefabric - Enterprise File Fabric

USAGE:
   singularity storage create filefabric [command options] <name> <path>

DESCRIPTION:
   --url
      URL of the Enterprise File Fabric to connect to.

      Examples:
         | https://storagemadeeasy.com       | Storage Made Easy US
         | https://eu.storagemadeeasy.com    | Storage Made Easy EU
         | https://yourfabric.smestorage.com | Connect to your Enterprise File Fabric

   --root_folder_id
      ID of the root folder.
      
      Leave blank normally.
      
      Fill in to make rclone start with directory of a given ID.
      

   --permanent_token
      Permanent Authentication Token.
      
      A Permanent Authentication Token can be created in the Enterprise File
      Fabric, on the users Dashboard under Security, there is an entry
      you'll see called "My Authentication Tokens". Click the Manage button
      to create one.
      
      These tokens are normally valid for several years.
      
      For more info see: https://docs.storagemadeeasy.com/organisationcloud/api-tokens
      

   --token
      Session Token.
      
      This is a session token which rclone caches in the config file. It is
      usually valid for 1 hour.
      
      Don't set this value - rclone will set it automatically.
      

   --token_expiry
      Token expiry time.
      
      Don't set this value - rclone will set it automatically.
      

   --version
      Version read from the file fabric.
      
      Don't set this value - rclone will set it automatically.
      

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h               show help
   --permanent_token value  Permanent Authentication Token. [$PERMANENT_TOKEN]
   --root_folder_id value   ID of the root folder. [$ROOT_FOLDER_ID]
   --url value              URL of the Enterprise File Fabric to connect to. [$URL]

   Advanced

   --encoding value      The encoding for the backend. (default: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --token value         Session Token. [$TOKEN]
   --token_expiry value  Token expiry time. [$TOKEN_EXPIRY]
   --version value       Version read from the file fabric. [$VERSION]

```
{% endcode %}
