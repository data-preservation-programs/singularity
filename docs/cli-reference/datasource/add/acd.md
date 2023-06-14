# Amazon Drive

```
NAME:
   singularity datasource add acd - Amazon Drive

USAGE:
   singularity datasource add acd [command options] <dataset_name> <source_path>

DESCRIPTION:
   --acd-checkpoint
      Checkpoint for internal polling (debug).

   --acd-token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --acd-client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --acd-token
      OAuth Access Token as a JSON blob.

   --acd-auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --acd-upload-wait-per-gb
      Additional time per GiB to wait after a failed complete upload to see if it appears.
      
      Sometimes Amazon Drive gives an error when a file has been fully
      uploaded but the file appears anyway after a little while.  This
      happens sometimes for files over 1 GiB in size and nearly every time for
      files bigger than 10 GiB. This parameter controls the time rclone waits
      for the file to appear.
      
      The default value for this parameter is 3 minutes per GiB, so by
      default it will wait 3 minutes for every GiB uploaded to see if the
      file appears.
      
      You can disable this feature by setting it to 0. This may cause
      conflict errors as rclone retries the failed upload but the file will
      most likely appear correctly eventually.
      
      These values were determined empirically by observing lots of uploads
      of big files for a range of file sizes.
      
      Upload with the "-v" flag to see more info about what rclone is doing
      in this situation.

   --acd-templink-threshold
      Files >= this size will be downloaded via their tempLink.
      
      Files this size or more will be downloaded via their "tempLink". This
      is to work around a problem with Amazon Drive which blocks downloads
      of files bigger than about 10 GiB. The default for this is 9 GiB which
      shouldn't need to be changed.
      
      To download files above this threshold, rclone requests a "tempLink"
      which downloads the file through a temporary URL directly from the
      underlying S3 storage.

   --acd-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --acd-client-id
      OAuth Client Id.
      
      Leave blank normally.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for acd

   --acd-auth-url value            Auth server URL. [$ACD_AUTH_URL]
   --acd-client-id value           OAuth Client Id. [$ACD_CLIENT_ID]
   --acd-client-secret value       OAuth Client Secret. [$ACD_CLIENT_SECRET]
   --acd-encoding value            The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$ACD_ENCODING]
   --acd-templink-threshold value  Files >= this size will be downloaded via their tempLink. (default: "9Gi") [$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               OAuth Access Token as a JSON blob. [$ACD_TOKEN]
   --acd-token-url value           Token server url. [$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  Additional time per GiB to wait after a failed complete upload to see if it appears. (default: "3m0s") [$ACD_UPLOAD_WAIT_PER_GB]

```
