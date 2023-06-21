# Google Photos

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add gphotos - Google Photos

USAGE:
   singularity datasource add gphotos [command options] <dataset_name> <source_path>

DESCRIPTION:
   --gphotos-client-id
      OAuth Client Id.
      
      Leave blank normally.

   --gphotos-token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --gphotos-read-size
      Set to read the size of media items.
      
      Normally rclone does not read the size of media items since this takes
      another transaction.  This isn't necessary for syncing.  However
      rclone mount needs to know the size of files in advance of reading
      them, so setting this flag when using rclone mount is recommended if
      you want to read the media.

   --gphotos-start-year
      Year limits the photos to be downloaded to those which are uploaded after the given year.

   --gphotos-include-archived
      Also view and download archived media.
      
      By default, rclone does not request archived media. Thus, when syncing,
      archived media is not visible in directory listings or transferred.
      
      Note that media in albums is always visible and synced, no matter
      their archive status.
      
      With this flag, archived media are always visible in directory
      listings and transferred.
      
      Without this flag, archived media will not be visible in directory
      listings and won't be transferred.

   --gphotos-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --gphotos-client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --gphotos-token
      OAuth Access Token as a JSON blob.

   --gphotos-auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --gphotos-read-only
      Set to make the Google Photos backend read only.
      
      If you choose read only then rclone will only request read only access
      to your photos, otherwise rclone will request full access.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for gphotos

   --gphotos-auth-url value          Auth server URL. [$GPHOTOS_AUTH_URL]
   --gphotos-client-id value         OAuth Client Id. [$GPHOTOS_CLIENT_ID]
   --gphotos-client-secret value     OAuth Client Secret. [$GPHOTOS_CLIENT_SECRET]
   --gphotos-encoding value          The encoding for the backend. (default: "Slash,CrLf,InvalidUtf8,Dot") [$GPHOTOS_ENCODING]
   --gphotos-include-archived value  Also view and download archived media. (default: "false") [$GPHOTOS_INCLUDE_ARCHIVED]
   --gphotos-read-only value         Set to make the Google Photos backend read only. (default: "false") [$GPHOTOS_READ_ONLY]
   --gphotos-read-size value         Set to read the size of media items. (default: "false") [$GPHOTOS_READ_SIZE]
   --gphotos-start-year value        Year limits the photos to be downloaded to those which are uploaded after the given year. (default: "2000") [$GPHOTOS_START_YEAR]
   --gphotos-token value             OAuth Access Token as a JSON blob. [$GPHOTOS_TOKEN]
   --gphotos-token-url value         Token server url. [$GPHOTOS_TOKEN_URL]

```
{% endcode %}
