# Sugarsync

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add sugarsync - Sugarsync

USAGE:
   singularity datasource add sugarsync [command options] <dataset_name> <source_path>

DESCRIPTION:
   --sugarsync-app-id
      Sugarsync App ID.

      Leave blank to use rclone's.

   --sugarsync-access-key-id
      Sugarsync Access Key ID.

      Leave blank to use rclone's.

   --sugarsync-private-access-key
      Sugarsync Private Access Key.

      Leave blank to use rclone's.

   --sugarsync-hard-delete
      Permanently delete files if true
      otherwise put them in the deleted files.

   --sugarsync-authorization
      Sugarsync authorization.

      Leave blank normally, will be auto configured by rclone.

   --sugarsync-user
      Sugarsync user.

      Leave blank normally, will be auto configured by rclone.

   --sugarsync-encoding
      The encoding for the backend.

      See the [encoding section in the overview](/overview/#encoding) for more info.

   --sugarsync-refresh-token
      Sugarsync refresh token.

      Leave blank normally, will be auto configured by rclone.

   --sugarsync-authorization-expiry
      Sugarsync authorization expiry.

      Leave blank normally, will be auto configured by rclone.

   --sugarsync-root-id
      Sugarsync root id.

      Leave blank normally, will be auto configured by rclone.

   --sugarsync-deleted-id
      Sugarsync deleted folder id.

      Leave blank normally, will be auto configured by rclone.


OPTIONS:
   --help, -h                            show help
   --sugarsync-access-key-id value       Sugarsync Access Key ID. [$SUGARSYNC_ACCESS_KEY_ID]
   --sugarsync-app-id value              Sugarsync App ID. [$SUGARSYNC_APP_ID]
   --sugarsync-hard-delete value         Permanently delete files if true (default: "false") [$SUGARSYNC_HARD_DELETE]
   --sugarsync-private-access-key value  Sugarsync Private Access Key. [$SUGARSYNC_PRIVATE_ACCESS_KEY]

   Advanced Options

   --sugarsync-authorization value         Sugarsync authorization. [$SUGARSYNC_AUTHORIZATION]
   --sugarsync-authorization-expiry value  Sugarsync authorization expiry. [$SUGARSYNC_AUTHORIZATION_EXPIRY]
   --sugarsync-deleted-id value            Sugarsync deleted folder id. [$SUGARSYNC_DELETED_ID]
   --sugarsync-encoding value              The encoding for the backend. (default: "Slash,Ctl,InvalidUtf8,Dot") [$SUGARSYNC_ENCODING]
   --sugarsync-refresh-token value         Sugarsync refresh token. [$SUGARSYNC_REFRESH_TOKEN]
   --sugarsync-root-id value               Sugarsync root id. [$SUGARSYNC_ROOT_ID]
   --sugarsync-user value                  Sugarsync user. [$SUGARSYNC_USER]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)

```
{% endcode %}
