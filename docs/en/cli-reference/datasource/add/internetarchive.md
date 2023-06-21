# Internet Archive

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add internetarchive - Internet Archive

USAGE:
   singularity datasource add internetarchive [command options] <dataset_name> <source_path>

DESCRIPTION:
   --internetarchive-access-key-id
      IAS3 Access Key.
      
      Leave blank for anonymous access.
      You can find one here: https://archive.org/account/s3.php

   --internetarchive-disable-checksum
      Don't ask the server to test against MD5 checksum calculated by rclone.
      Normally rclone will calculate the MD5 checksum of the input before
      uploading it so it can ask the server to check the object against checksum.
      This is great for data integrity checking but can cause long delays for
      large files to start uploading.

   --internetarchive-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --internetarchive-endpoint
      IAS3 Endpoint.
      
      Leave blank for default value.

   --internetarchive-front-endpoint
      Host of InternetArchive Frontend.
      
      Leave blank for default value.

   --internetarchive-secret-access-key
      IAS3 Secret Key (password).
      
      Leave blank for anonymous access.

   --internetarchive-wait-archive
      Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
      Only enable if you need to be guaranteed to be reflected after write operations.
      0 to disable waiting. No errors to be thrown in case of timeout.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for internetarchive

   --internetarchive-access-key-id value      IAS3 Access Key. [$INTERNETARCHIVE_ACCESS_KEY_ID]
   --internetarchive-disable-checksum value   Don't ask the server to test against MD5 checksum calculated by rclone. (default: "true") [$INTERNETARCHIVE_DISABLE_CHECKSUM]
   --internetarchive-encoding value           The encoding for the backend. (default: "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot") [$INTERNETARCHIVE_ENCODING]
   --internetarchive-endpoint value           IAS3 Endpoint. (default: "https://s3.us.archive.org") [$INTERNETARCHIVE_ENDPOINT]
   --internetarchive-front-endpoint value     Host of InternetArchive Frontend. (default: "https://archive.org") [$INTERNETARCHIVE_FRONT_ENDPOINT]
   --internetarchive-secret-access-key value  IAS3 Secret Key (password). [$INTERNETARCHIVE_SECRET_ACCESS_KEY]
   --internetarchive-wait-archive value       Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish. (default: "0s") [$INTERNETARCHIVE_WAIT_ARCHIVE]

```
{% endcode %}
