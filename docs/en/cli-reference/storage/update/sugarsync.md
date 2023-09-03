# Sugarsync

{% code fullWidth="true" %}
```
NAME:
   singularity storage update sugarsync - Sugarsync

USAGE:
   singularity storage update sugarsync [command options] <name|id>

DESCRIPTION:
   --app-id
      Sugarsync App ID.
      
      Leave blank to use rclone's.

   --access-key-id
      Sugarsync Access Key ID.
      
      Leave blank to use rclone's.

   --private-access-key
      Sugarsync Private Access Key.
      
      Leave blank to use rclone's.

   --hard-delete
      Permanently delete files if true
      otherwise put them in the deleted files.

   --refresh-token
      Sugarsync refresh token.
      
      Leave blank normally, will be auto configured by rclone.

   --authorization
      Sugarsync authorization.
      
      Leave blank normally, will be auto configured by rclone.

   --authorization-expiry
      Sugarsync authorization expiry.
      
      Leave blank normally, will be auto configured by rclone.

   --user
      Sugarsync user.
      
      Leave blank normally, will be auto configured by rclone.

   --root-id
      Sugarsync root id.
      
      Leave blank normally, will be auto configured by rclone.

   --deleted-id
      Sugarsync deleted folder id.
      
      Leave blank normally, will be auto configured by rclone.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --access-key-id value       Sugarsync Access Key ID. [$ACCESS_KEY_ID]
   --app-id value              Sugarsync App ID. [$APP_ID]
   --hard-delete               Permanently delete files if true (default: false) [$HARD_DELETE]
   --help, -h                  show help
   --private-access-key value  Sugarsync Private Access Key. [$PRIVATE_ACCESS_KEY]

   Advanced

   --authorization value         Sugarsync authorization. [$AUTHORIZATION]
   --authorization-expiry value  Sugarsync authorization expiry. [$AUTHORIZATION_EXPIRY]
   --deleted-id value            Sugarsync deleted folder id. [$DELETED_ID]
   --encoding value              The encoding for the backend. (default: "Slash,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --refresh-token value         Sugarsync refresh token. [$REFRESH_TOKEN]
   --root-id value               Sugarsync root id. [$ROOT_ID]
   --user value                  Sugarsync user. [$USER]

```
{% endcode %}
