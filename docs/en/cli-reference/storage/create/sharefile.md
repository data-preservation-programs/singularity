# Citrix Sharefile

{% code fullWidth="true" %}
```
NAME:
   singularity storage create sharefile - Citrix Sharefile

USAGE:
   singularity storage create sharefile [command options] <name> <path>

DESCRIPTION:
   --upload_cutoff
      Cutoff for switching to multipart upload.

   --root_folder_id
      ID of the root folder.
      
      Leave blank to access "Personal Folders".  You can use one of the
      standard values here or any folder ID (long hex number ID).

      Examples:
         | <unset>    | Access the Personal Folders (default).
         | favorites  | Access the Favorites folder.
         | allshared  | Access all the shared folders.
         | connectors | Access all the individual connectors.
         | top        | Access the home, favorites, and shared folders as well as the connectors.

   --chunk_size
      Upload chunk size.
      
      Must a power of 2 >= 256k.
      
      Making this larger will improve performance, but note that each chunk
      is buffered in memory one per transfer.
      
      Reducing this will reduce memory usage but decrease performance.

   --endpoint
      Endpoint for API calls.
      
      This is usually auto discovered as part of the oauth process, but can
      be set manually to something like: https://XXX.sharefile.com
      

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h              show help
   --root_folder_id value  ID of the root folder. [$ROOT_FOLDER_ID]

   Advanced

   --chunk_size value     Upload chunk size. (default: "64Mi") [$CHUNK_SIZE]
   --encoding value       The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value       Endpoint for API calls. [$ENDPOINT]
   --upload_cutoff value  Cutoff for switching to multipart upload. (default: "128Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}
