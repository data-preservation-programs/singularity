# Koofr / Digi Storage

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add koofr - Koofr, Digi Storage and other Koofr-compatible storage providers

USAGE:
   singularity datasource add koofr [command options] <dataset_name> <source_path>

DESCRIPTION:
   --koofr-user
      Your user name.

   --koofr-password
      [Provider] - koofr
         Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password).

      [Provider] - digistorage
         Your password for rclone (generate one at https://storage.rcs-rds.ro/app/admin/preferences/password).

      [Provider] - other
         Your password for rclone (generate one at your service's settings page).

   --koofr-encoding
      The encoding for the backend.

      See the [encoding section in the overview](/overview/#encoding) for more info.

   --koofr-provider
      Choose your storage provider.

      Examples:
         | koofr       | Koofr, https://app.koofr.net/
         | digistorage | Digi Storage, https://storage.rcs-rds.ro/
         | other       | Any other Koofr API compatible storage service

   --koofr-endpoint
      [Provider] - other
         The Koofr API endpoint to use.

   --koofr-mountid
      Mount ID of the mount to use.

      If omitted, the primary mount is used.

   --koofr-setmtime
      Does the backend support setting modification time.

      Set this to false if you use a mount ID that points to a Dropbox or Amazon Drive backend.


OPTIONS:
   --help, -h              show help
   --koofr-endpoint value  The Koofr API endpoint to use. [$KOOFR_ENDPOINT]
   --koofr-password value  Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password). [$KOOFR_PASSWORD]
   --koofr-provider value  Choose your storage provider. [$KOOFR_PROVIDER]
   --koofr-user value      Your user name. [$KOOFR_USER]

   Advanced Options

   --koofr-encoding value  The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$KOOFR_ENCODING]
   --koofr-mountid value   Mount ID of the mount to use. [$KOOFR_MOUNTID]
   --koofr-setmtime value  Does the backend support setting modification time. (default: "true") [$KOOFR_SETMTIME]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)

```
{% endcode %}
