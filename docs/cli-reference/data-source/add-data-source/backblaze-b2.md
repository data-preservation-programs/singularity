# Backblaze B2

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add b2 - Backblaze B2

USAGE:
   singularity datasource add b2 [command options] <dataset_name> <source_path>

DESCRIPTION:
   --b2-chunk-size
      Upload chunk size.

      When uploading large files, chunk the file into this size.

      Must fit in memory. These chunks are buffered in memory and there
      might a maximum of "--transfers" chunks in progress at once.

      5,000,000 Bytes is the minimum size.

   --b2-encoding
      The encoding for the backend.

      See the [encoding section in the overview](/overview/#encoding) for more info.

   --b2-endpoint
      Endpoint for the service.

      Leave blank normally.

   --b2-hard-delete
      Permanently delete files on remote removal, otherwise hide files.

   --b2-version-at
      Show file versions as they were at the specified time.

      Note that when using this no file write operations are permitted,
      so you can't upload files or delete them.

   --b2-disable-checksum
      Disable checksums for large (> upload cutoff) files.

      Normally rclone will calculate the SHA1 checksum of the input before
      uploading it so it can add it to metadata on the object. This is great
      for data integrity checking but can cause long delays for large files
      to start uploading.

   --b2-download-auth-duration
      Time before the authorization token will expire in s or suffix ms|s|m|h|d.

      The duration before the download authorization token will expire.
      The minimum value is 1 second. The maximum value is one week.

   --b2-account
      Account ID or Application Key ID.

   --b2-key
      Application Key.

   --b2-memory-pool-use-mmap
      Whether to use mmap buffers in internal memory pool.

   --b2-test-mode
      A flag string for X-Bz-Test-Mode header for debugging.

      This is for debugging purposes only. Setting it to one of the strings
      below will cause b2 to return specific errors:

        * "fail_some_uploads"
        * "expire_some_account_authorization_tokens"
        * "force_cap_exceeded"

      These will be set in the "X-Bz-Test-Mode" header which is documented
      in the [b2 integrations checklist](https://www.backblaze.com/b2/docs/integration_checklist.html).

   --b2-memory-pool-flush-time
      How often internal memory buffer pools will be flushed.
      Uploads which requires additional buffers (f.e multipart) will use memory pool for allocations.
      This option controls how often unused buffers will be removed from the pool.

   --b2-copy-cutoff
      Cutoff for switching to multipart copy.

      Any files larger than this that need to be server-side copied will be
      copied in chunks of this size.

      The minimum is 0 and the maximum is 4.6 GiB.

   --b2-download-url
      Custom endpoint for downloads.

      This is usually set to a Cloudflare CDN URL as Backblaze offers
      free egress for data downloaded through the Cloudflare network.
      Rclone works with private buckets by sending an "Authorization" header.
      If the custom endpoint rewrites the requests for authentication,
      e.g., in Cloudflare Workers, this header needs to be handled properly.
      Leave blank if you want to use the endpoint provided by Backblaze.

      The URL provided here SHOULD have the protocol and SHOULD NOT have
      a trailing slash or specify the /file/bucket subpath as rclone will
      request files with "{download_url}/file/{bucket_name}/{path}".

      Example:
      > https://mysubdomain.mydomain.tld
      (No trailing "/", "file" or "bucket")

   --b2-versions
      Include old versions in directory listings.

      Note that when using this no file write operations are permitted,
      so you can't upload files or delete them.

   --b2-upload-cutoff
      Cutoff for switching to chunked upload.

      Files above this size will be uploaded in chunks of "--b2-chunk-size".

      This value should be set no larger than 4.657 GiB (== 5 GB).


OPTIONS:
   --b2-account value      Account ID or Application Key ID. [$B2_ACCOUNT]
   --b2-hard-delete value  Permanently delete files on remote removal, otherwise hide files. (default: "false") [$B2_HARD_DELETE]
   --b2-key value          Application Key. [$B2_KEY]
   --help, -h              show help

   Advanced Options

   --b2-chunk-size value              Upload chunk size. (default: "96Mi") [$B2_CHUNK_SIZE]
   --b2-copy-cutoff value             Cutoff for switching to multipart copy. (default: "4Gi") [$B2_COPY_CUTOFF]
   --b2-disable-checksum value        Disable checksums for large (> upload cutoff) files. (default: "false") [$B2_DISABLE_CHECKSUM]
   --b2-download-auth-duration value  Time before the authorization token will expire in s or suffix ms|s|m|h|d. (default: "1w") [$B2_DOWNLOAD_AUTH_DURATION]
   --b2-download-url value            Custom endpoint for downloads. [$B2_DOWNLOAD_URL]
   --b2-encoding value                The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$B2_ENCODING]
   --b2-endpoint value                Endpoint for the service. [$B2_ENDPOINT]
   --b2-memory-pool-flush-time value  How often internal memory buffer pools will be flushed. (default: "1m0s") [$B2_MEMORY_POOL_FLUSH_TIME]
   --b2-memory-pool-use-mmap value    Whether to use mmap buffers in internal memory pool. (default: "false") [$B2_MEMORY_POOL_USE_MMAP]
   --b2-test-mode value               A flag string for X-Bz-Test-Mode header for debugging. [$B2_TEST_MODE]
   --b2-upload-cutoff value           Cutoff for switching to chunked upload. (default: "200Mi") [$B2_UPLOAD_CUTOFF]
   --b2-version-at value              Show file versions as they were at the specified time. (default: "off") [$B2_VERSION_AT]
   --b2-versions value                Include old versions in directory listings. (default: "false") [$B2_VERSIONS]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)

```
{% endcode %}
