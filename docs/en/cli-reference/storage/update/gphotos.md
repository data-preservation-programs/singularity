# Google Photos

{% code fullWidth="true" %}
```
NAME:
   singularity storage update gphotos - Google Photos

USAGE:
   singularity storage update gphotos [command options] <name|id>

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

   --read-only
      Set to make the Google Photos backend read only.
      
      If you choose read only then rclone will only request read only access
      to your photos, otherwise rclone will request full access.

   --read-size
      Set to read the size of media items.
      
      Normally rclone does not read the size of media items since this takes
      another transaction.  This isn't necessary for syncing.  However
      rclone mount needs to know the size of files in advance of reading
      them, so setting this flag when using rclone mount is recommended if
      you want to read the media.

   --start-year
      Year limits the photos to be downloaded to those which are uploaded after the given year.

   --include-archived
      Also view and download archived media.
      
      By default, rclone does not request archived media. Thus, when syncing,
      archived media is not visible in directory listings or transferred.
      
      Note that media in albums is always visible and synced, no matter
      their archive status.
      
      With this flag, archived media are always visible in directory
      listings and transferred.
      
      Without this flag, archived media will not be visible in directory
      listings and won't be transferred.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --client-id value      OAuth Client Id. [$CLIENT_ID]
   --client-secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help
   --read-only            Set to make the Google Photos backend read only. (default: false) [$READ_ONLY]

   Advanced

   --auth-url value    Auth server URL. [$AUTH_URL]
   --encoding value    The encoding for the backend. (default: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --include-archived  Also view and download archived media. (default: false) [$INCLUDE_ARCHIVED]
   --read-size         Set to read the size of media items. (default: false) [$READ_SIZE]
   --start-year value  Year limits the photos to be downloaded to those which are uploaded after the given year. (default: 2000) [$START_YEAR]
   --token value       OAuth Access Token as a JSON blob. [$TOKEN]
   --token-url value   Token server url. [$TOKEN_URL]

```
{% endcode %}
