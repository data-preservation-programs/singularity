# Amazon Drive

{% code fullWidth="true" %}
```
NAME:
   singularity storage create acd - Amazon Drive

USAGE:
   singularity storage create acd [command options] <name> <path>

DESCRIPTION:
   --client_id
      OAuth Client Id.
      
      Leave blank normally.

   --client_secret
      OAuth Client Secret.
      
      Leave blank normally.

   --token
      OAuth Access Token as a JSON blob.

   --auth_url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --token_url
      Token server url.
      
      Leave blank to use the provider defaults.

   --checkpoint
      Checkpoint for internal polling (debug).

   --upload_wait_per_gb
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

   --templink_threshold
      Files >= this size will be downloaded via their tempLink.
      
      Files this size or more will be downloaded via their "tempLink". This
      is to work around a problem with Amazon Drive which blocks downloads
      of files bigger than about 10 GiB. The default for this is 9 GiB which
      shouldn't need to be changed.
      
      To download files above this threshold, rclone requests a "tempLink"
      which downloads the file through a temporary URL directly from the
      underlying S3 storage.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --client_id value      OAuth Client Id. [$CLIENT_ID]
   --client_secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help

   Advanced

   --auth_url value            Auth server URL. [$AUTH_URL]
   --checkpoint value          Checkpoint for internal polling (debug). [$CHECKPOINT]
   --encoding value            The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --templink_threshold value  Files >= this size will be downloaded via their tempLink. (default: "9Gi") [$TEMPLINK_THRESHOLD]
   --token value               OAuth Access Token as a JSON blob. [$TOKEN]
   --token_url value           Token server url. [$TOKEN_URL]
   --upload_wait_per_gb value  Additional time per GiB to wait after a failed complete upload to see if it appears. (default: "3m0s") [$UPLOAD_WAIT_PER_GB]

```
{% endcode %}
