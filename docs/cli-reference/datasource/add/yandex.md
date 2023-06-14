# Yandex Disk

```
NAME:
   singularity datasource add yandex - Yandex Disk

USAGE:
   singularity datasource add yandex [command options] <dataset_name> <source_path>

DESCRIPTION:
   --yandex-auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --yandex-token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --yandex-hard-delete
      Delete files permanently rather than putting them into the trash.

   --yandex-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --yandex-client-id
      OAuth Client Id.
      
      Leave blank normally.

   --yandex-client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --yandex-token
      OAuth Access Token as a JSON blob.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for yandex

   --yandex-auth-url value       Auth server URL. [$YANDEX_AUTH_URL]
   --yandex-client-id value      OAuth Client Id. [$YANDEX_CLIENT_ID]
   --yandex-client-secret value  OAuth Client Secret. [$YANDEX_CLIENT_SECRET]
   --yandex-encoding value       The encoding for the backend. (default: "Slash,Del,Ctl,InvalidUtf8,Dot") [$YANDEX_ENCODING]
   --yandex-hard-delete value    Delete files permanently rather than putting them into the trash. (default: "false") [$YANDEX_HARD_DELETE]
   --yandex-token value          OAuth Access Token as a JSON blob. [$YANDEX_TOKEN]
   --yandex-token-url value      Token server url. [$YANDEX_TOKEN_URL]

```
