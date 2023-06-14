# Citrix Sharefile

```
NAME:
   singularity datasource add sharefile - Citrix Sharefile

USAGE:
   singularity datasource add sharefile [command options] <dataset_name> <source_path>

DESCRIPTION:
   --sharefile-upload-cutoff
      Cutoff for switching to multipart upload.

   --sharefile-root-folder-id
      ID of the root folder.
      
      Leave blank to access "Personal Folders".  You can use one of the
      standard values here or any folder ID (long hex number ID).

      Examples:
         | <unset>    | Access the Personal Folders (default).
         | favorites  | Access the Favorites folder.
         | allshared  | Access all the shared folders.
         | connectors | Access all the individual connectors.
         | top        | Access the home, favorites, and shared folders as well as the connectors.

   --sharefile-chunk-size
      Upload chunk size.
      
      Must a power of 2 >= 256k.
      
      Making this larger will improve performance, but note that each chunk
      is buffered in memory one per transfer.
      
      Reducing this will reduce memory usage but decrease performance.

   --sharefile-endpoint
      Endpoint for API calls.
      
      This is usually auto discovered as part of the oauth process, but can
      be set manually to something like: https://XXX.sharefile.com
      

   --sharefile-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for sharefile

   --sharefile-chunk-size value      Upload chunk size. (default: "64Mi") [$SHAREFILE_CHUNK_SIZE]
   --sharefile-encoding value        The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot") [$SHAREFILE_ENCODING]
   --sharefile-endpoint value        Endpoint for API calls. [$SHAREFILE_ENDPOINT]
   --sharefile-root-folder-id value  ID of the root folder. [$SHAREFILE_ROOT_FOLDER_ID]
   --sharefile-upload-cutoff value   Cutoff for switching to multipart upload. (default: "128Mi") [$SHAREFILE_UPLOAD_CUTOFF]

```
