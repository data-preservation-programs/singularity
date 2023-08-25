# Sugarsync

{% code fullWidth="true" %}
```
NAME:
   singularity storage update sugarsync - Sugarsync

USAGE:
   singularity storage update sugarsync [command options] <name>

DESCRIPTION:
   --app_id
      Sugarsync App ID.
      
      Leave blank to use rclone's.

   --access_key_id
      Sugarsync Access Key ID.
      
      Leave blank to use rclone's.

   --private_access_key
      Sugarsync Private Access Key.
      
      Leave blank to use rclone's.

   --hard_delete
      Permanently delete files if true
      otherwise put them in the deleted files.

   --refresh_token
      Sugarsync refresh token.
      
      Leave blank normally, will be auto configured by rclone.

   --authorization
      Sugarsync authorization.
      
      Leave blank normally, will be auto configured by rclone.

   --authorization_expiry
      Sugarsync authorization expiry.
      
      Leave blank normally, will be auto configured by rclone.

   --user
      Sugarsync user.
      
      Leave blank normally, will be auto configured by rclone.

   --root_id
      Sugarsync root id.
      
      Leave blank normally, will be auto configured by rclone.

   --deleted_id
      Sugarsync deleted folder id.
      
      Leave blank normally, will be auto configured by rclone.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --access_key_id value       Sugarsync Access Key ID. [$ACCESS_KEY_ID]
   --app_id value              Sugarsync App ID. [$APP_ID]
   --hard_delete               Permanently delete files if true (default: false) [$HARD_DELETE]
   --help, -h                  show help
   --private_access_key value  Sugarsync Private Access Key. [$PRIVATE_ACCESS_KEY]

   Advanced

   --authorization value         Sugarsync authorization. [$AUTHORIZATION]
   --authorization_expiry value  Sugarsync authorization expiry. [$AUTHORIZATION_EXPIRY]
   --deleted_id value            Sugarsync deleted folder id. [$DELETED_ID]
   --encoding value              The encoding for the backend. (default: "Slash,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --refresh_token value         Sugarsync refresh token. [$REFRESH_TOKEN]
   --root_id value               Sugarsync root id. [$ROOT_ID]
   --user value                  Sugarsync user. [$USER]

```
{% endcode %}
