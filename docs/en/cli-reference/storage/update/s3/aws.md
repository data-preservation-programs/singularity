# Amazon Web Services (AWS) S3

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 aws - Amazon Web Services (AWS) S3

USAGE:
   singularity storage update s3 aws [command options] <name|id>

DESCRIPTION:
   --env-auth
      Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
      
      Only applies if access_key_id and secret_access_key is blank.

      Examples:
         | false | Enter AWS credentials in the next step.
         | true  | Get AWS credentials from the environment (env vars or IAM).

   --access-key-id
      AWS Access Key ID.
      
      Leave blank for anonymous access or runtime credentials.

   --secret-access-key
      AWS Secret Access Key (password).
      
      Leave blank for anonymous access or runtime credentials.

   --region
      Region to connect to.

      Examples:
         | us-east-1      | The default endpoint - a good choice if you are unsure.
         |                | US Region, Northern Virginia, or Pacific Northwest.
         |                | Leave location constraint empty.
         | us-east-2      | US East (Ohio) Region.
         |                | Needs location constraint us-east-2.
         | us-west-1      | US West (Northern California) Region.
         |                | Needs location constraint us-west-1.
         | us-west-2      | US West (Oregon) Region.
         |                | Needs location constraint us-west-2.
         | ca-central-1   | Canada (Central) Region.
         |                | Needs location constraint ca-central-1.
         | eu-west-1      | EU (Ireland) Region.
         |                | Needs location constraint EU or eu-west-1.
         | eu-west-2      | EU (London) Region.
         |                | Needs location constraint eu-west-2.
         | eu-west-3      | EU (Paris) Region.
         |                | Needs location constraint eu-west-3.
         | eu-north-1     | EU (Stockholm) Region.
         |                | Needs location constraint eu-north-1.
         | eu-south-1     | EU (Milan) Region.
         |                | Needs location constraint eu-south-1.
         | eu-central-1   | EU (Frankfurt) Region.
         |                | Needs location constraint eu-central-1.
         | ap-southeast-1 | Asia Pacific (Singapore) Region.
         |                | Needs location constraint ap-southeast-1.
         | ap-southeast-2 | Asia Pacific (Sydney) Region.
         |                | Needs location constraint ap-southeast-2.
         | ap-northeast-1 | Asia Pacific (Tokyo) Region.
         |                | Needs location constraint ap-northeast-1.
         | ap-northeast-2 | Asia Pacific (Seoul).
         |                | Needs location constraint ap-northeast-2.
         | ap-northeast-3 | Asia Pacific (Osaka-Local).
         |                | Needs location constraint ap-northeast-3.
         | ap-south-1     | Asia Pacific (Mumbai).
         |                | Needs location constraint ap-south-1.
         | ap-east-1      | Asia Pacific (Hong Kong) Region.
         |                | Needs location constraint ap-east-1.
         | sa-east-1      | South America (Sao Paulo) Region.
         |                | Needs location constraint sa-east-1.
         | il-central-1   | Israel (Tel Aviv) Region.
         |                | Needs location constraint il-central-1.
         | me-south-1     | Middle East (Bahrain) Region.
         |                | Needs location constraint me-south-1.
         | af-south-1     | Africa (Cape Town) Region.
         |                | Needs location constraint af-south-1.
         | cn-north-1     | China (Beijing) Region.
         |                | Needs location constraint cn-north-1.
         | cn-northwest-1 | China (Ningxia) Region.
         |                | Needs location constraint cn-northwest-1.
         | us-gov-east-1  | AWS GovCloud (US-East) Region.
         |                | Needs location constraint us-gov-east-1.
         | us-gov-west-1  | AWS GovCloud (US) Region.
         |                | Needs location constraint us-gov-west-1.

   --endpoint
      Endpoint for S3 API.
      
      Leave blank if using AWS to use the default endpoint for the region.

   --location-constraint
      Location constraint - must be set to match the Region.
      
      Used when creating buckets only.

      Examples:
         | <unset>        | Empty for US Region, Northern Virginia, or Pacific Northwest
         | us-east-2      | US East (Ohio) Region
         | us-west-1      | US West (Northern California) Region
         | us-west-2      | US West (Oregon) Region
         | ca-central-1   | Canada (Central) Region
         | eu-west-1      | EU (Ireland) Region
         | eu-west-2      | EU (London) Region
         | eu-west-3      | EU (Paris) Region
         | eu-north-1     | EU (Stockholm) Region
         | eu-south-1     | EU (Milan) Region
         | EU             | EU Region
         | ap-southeast-1 | Asia Pacific (Singapore) Region
         | ap-southeast-2 | Asia Pacific (Sydney) Region
         | ap-northeast-1 | Asia Pacific (Tokyo) Region
         | ap-northeast-2 | Asia Pacific (Seoul) Region
         | ap-northeast-3 | Asia Pacific (Osaka-Local) Region
         | ap-south-1     | Asia Pacific (Mumbai) Region
         | ap-east-1      | Asia Pacific (Hong Kong) Region
         | sa-east-1      | South America (Sao Paulo) Region
         | il-central-1   | Israel (Tel Aviv) Region
         | me-south-1     | Middle East (Bahrain) Region
         | af-south-1     | Africa (Cape Town) Region
         | cn-north-1     | China (Beijing) Region
         | cn-northwest-1 | China (Ningxia) Region
         | us-gov-east-1  | AWS GovCloud (US-East) Region
         | us-gov-west-1  | AWS GovCloud (US) Region

   --acl
      Canned ACL used when creating buckets and storing or copying objects.
      
      This ACL is used for creating objects and if bucket_acl isn't set, for creating buckets too.
      
      For more info visit https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      Note that this ACL is applied when server-side copying objects as S3
      doesn't copy the ACL from the source but rather writes a fresh one.
      
      If the acl is an empty string then no X-Amz-Acl: header is added and
      the default (private) will be used.
      

   --bucket-acl
      Canned ACL used when creating buckets.
      
      For more info visit https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      Note that this ACL is applied when only when creating buckets.  If it
      isn't set then "acl" is used instead.
      
      If the "acl" and "bucket_acl" are empty strings then no X-Amz-Acl:
      header is added and the default (private) will be used.
      

      Examples:
         | private            | Owner gets FULL_CONTROL.
         |                    | No one else has access rights (default).
         | public-read        | Owner gets FULL_CONTROL.
         |                    | The AllUsers group gets READ access.
         | public-read-write  | Owner gets FULL_CONTROL.
         |                    | The AllUsers group gets READ and WRITE access.
         |                    | Granting this on a bucket is generally not recommended.
         | authenticated-read | Owner gets FULL_CONTROL.
         |                    | The AuthenticatedUsers group gets READ access.

   --requester-pays
      Enables requester pays option when interacting with S3 bucket.

   --server-side-encryption
      The server-side encryption algorithm used when storing this object in S3.

      Examples:
         | <unset> | None
         | AES256  | AES256

   --sse-customer-algorithm
      If using SSE-C, the server-side encryption algorithm used when storing this object in S3.

      Examples:
         | <unset> | None
         | AES256  | AES256

   --sse-kms-key-id
      If using KMS ID you must provide the ARN of Key.

      Examples:
         | <unset>                 | None
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
      
      Alternatively you can provide --sse-customer-key-base64.

      Examples:
         | <unset> | None

   --sse-customer-key-base64
      If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
      
      Alternatively you can provide --sse-customer-key.

      Examples:
         | <unset> | None

   --sse-customer-key-md5
      If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
      
      If you leave it blank, this is calculated automatically from the sse_customer_key provided.
      

      Examples:
         | <unset> | None

   --storage-class
      The storage class to use when storing new objects in S3.

      Examples:
         | <unset>             | Default
         | STANDARD            | Standard storage class
         | REDUCED_REDUNDANCY  | Reduced redundancy storage class
         | STANDARD_IA         | Standard Infrequent Access storage class
         | ONEZONE_IA          | One Zone Infrequent Access storage class
         | GLACIER             | Glacier storage class
         | DEEP_ARCHIVE        | Glacier Deep Archive storage class
         | INTELLIGENT_TIERING | Intelligent-Tiering storage class
         | GLACIER_IR          | Glacier Instant Retrieval storage class

   --upload-cutoff
      Cutoff for switching to chunked upload.
      
      Any files larger than this will be uploaded in chunks of chunk_size.
      The minimum is 0 and the maximum is 5 GiB.

   --chunk-size
      Chunk size to use for uploading.
      
      When uploading files larger than upload_cutoff or files with unknown
      size (e.g. from "rclone rcat" or uploaded with "rclone mount" or google
      photos or google docs) they will be uploaded as multipart uploads
      using this chunk size.
      
      Note that "--s3-upload-concurrency" chunks of this size are buffered
      in memory per transfer.
      
      If you are transferring large files over high-speed links and you have
      enough memory, then increasing this will speed up the transfers.
      
      Rclone will automatically increase the chunk size when uploading a
      large file of known size to stay below the 10,000 chunks limit.
      
      Files of unknown size are uploaded with the configured
      chunk_size. Since the default chunk size is 5 MiB and there can be at
      most 10,000 chunks, this means that by default the maximum size of
      a file you can stream upload is 48 GiB.  If you wish to stream upload
      larger files then you will need to increase chunk_size.
      
      Increasing the chunk size decreases the accuracy of the progress
      statistics displayed with "-P" flag. Rclone treats chunk as sent when
      it's buffered by the AWS SDK, when in fact it may still be uploading.
      A bigger chunk size means a bigger AWS SDK buffer and progress
      reporting more deviating from the truth.
      

   --max-upload-parts
      Maximum number of parts in a multipart upload.
      
      This option defines the maximum number of multipart chunks to use
      when doing a multipart upload.
      
      This can be useful if a service does not support the AWS S3
      specification of 10,000 chunks.
      
      Rclone will automatically increase the chunk size when uploading a
      large file of a known size to stay below this number of chunks limit.
      

   --copy-cutoff
      Cutoff for switching to multipart copy.
      
      Any files larger than this that need to be server-side copied will be
      copied in chunks of this size.
      
      The minimum is 0 and the maximum is 5 GiB.

   --disable-checksum
      Don't store MD5 checksum with object metadata.
      
      Normally rclone will calculate the MD5 checksum of the input before
      uploading it so it can add it to metadata on the object. This is great
      for data integrity checking but can cause long delays for large files
      to start uploading.

   --shared-credentials-file
      Path to the shared credentials file.
      
      If env_auth = true then rclone can use a shared credentials file.
      
      If this variable is empty rclone will look for the
      "AWS_SHARED_CREDENTIALS_FILE" env variable. If the env value is empty
      it will default to the current user's home directory.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      Profile to use in the shared credentials file.
      
      If env_auth = true then rclone can use a shared credentials file. This
      variable controls which profile is used in that file.
      
      If empty it will default to the environment variable "AWS_PROFILE" or
      "default" if that environment variable is also not set.
      

   --session-token
      An AWS session token.

   --upload-concurrency
      Concurrency for multipart uploads and copies.
      
      This is the number of chunks of the same file that are uploaded
      concurrently for multipart uploads and copies.
      
      If you are uploading small numbers of large files over high-speed links
      and these uploads do not fully utilize your bandwidth, then increasing
      this may help to speed up the transfers.

   --force-path-style
      If true use path style access if false use virtual hosted style.
      
      If this is true (the default) then rclone will use path style access,
      if false then rclone will use virtual path style. See [the AWS S3
      docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      for more info.
      
      Some providers (e.g. AWS, Aliyun OSS, Netease COS, or Tencent COS) require this set to
      false - rclone will do this automatically based on the provider
      setting.
      
      Note that if your bucket isn't a valid DNS name, i.e. has '.' or '_' in,
      you'll need to set this to true.
      

   --v2-auth
      If true use v2 authentication.
      
      If this is false (the default) then rclone will use v4 authentication.
      If it is set then rclone will use v2 authentication.
      
      Use this only if v4 signatures don't work, e.g. pre Jewel/v10 CEPH.

   --use-dual-stack
      If true use AWS S3 dual-stack endpoint (IPv6 support).
      
      See [AWS Docs on Dualstack Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/dual-stack-endpoints.html)

   --use-accelerate-endpoint
      If true use the AWS S3 accelerated endpoint.
      
      See: [AWS S3 Transfer acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration-examples.html)

   --leave-parts-on-error
      If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
      
      It should be set to true for resuming uploads across different sessions.
      
      WARNING: Storing parts of an incomplete multipart upload counts towards space usage on S3 and will add additional costs if not cleaned up.
      

   --list-chunk
      Size of listing chunk (response list for each ListObject S3 request).
      
      This option is also known as "MaxKeys", "max-items", or "page-size" from the AWS S3 specification.
      Most services truncate the response list to 1000 objects even if requested more than that.
      In AWS S3 this is a global maximum and cannot be changed, see [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html).
      In Ceph, this can be increased with the "rgw list buckets max chunk" option.
      

   --list-version
      Version of ListObjects to use: 1,2 or 0 for auto.
      
      When S3 originally launched it only provided the ListObjects call to
      enumerate objects in a bucket.
      
      However in May 2016 the ListObjectsV2 call was introduced. This is
      much higher performance and should be used if at all possible.
      
      If set to the default, 0, rclone will guess according to the provider
      set which list objects method to call. If it guesses wrong, then it
      may be set manually here.
      

   --list-url-encode
      Whether to url encode listings: true/false/unset
      
      Some providers support URL encoding listings and where this is
      available this is more reliable when using control characters in file
      names. If this is set to unset (the default) then rclone will choose
      according to the provider setting what to apply, but you can override
      rclone's choice here.
      

   --no-check-bucket
      If set, don't attempt to check the bucket exists or create it.
      
      This can be useful when trying to minimise the number of transactions
      rclone does if you know the bucket exists already.
      
      It can also be needed if the user you are using does not have bucket
      creation permissions. Before v1.52.0 this would have passed silently
      due to a bug.
      

   --no-head
      If set, don't HEAD uploaded objects to check integrity.
      
      This can be useful when trying to minimise the number of transactions
      rclone does.
      
      Setting it means that if rclone receives a 200 OK message after
      uploading an object with PUT then it will assume that it got uploaded
      properly.
      
      In particular it will assume:
      
      - the metadata, including modtime, storage class and content type was as uploaded
      - the size was as uploaded
      
      It reads the following items from the response for a single part PUT:
      
      - the MD5SUM
      - The uploaded date
      
      For multipart uploads these items aren't read.
      
      If an source object of unknown length is uploaded then rclone **will** do a
      HEAD request.
      
      Setting this flag increases the chance for undetected upload failures,
      in particular an incorrect size, so it isn't recommended for normal
      operation. In practice the chance of an undetected upload failure is
      very small even with this flag.
      

   --no-head-object
      If set, do not do HEAD before GET when getting objects.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --memory-pool-flush-time
      How often internal memory buffer pools will be flushed. (no longer used)

   --memory-pool-use-mmap
      Whether to use mmap buffers in internal memory pool. (no longer used)

   --disable-http2
      Disable usage of http2 for S3 backends.
      
      There is currently an unsolved issue with the s3 (specifically minio) backend
      and HTTP/2.  HTTP/2 is enabled by default for the s3 backend but can be
      disabled here.  When the issue is solved this flag will be removed.
      
      See: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      Custom endpoint for downloads.
      This is usually set to a CloudFront CDN URL as AWS S3 offers
      cheaper egress for data downloaded through the CloudFront network.

   --directory-markers
      Upload an empty object with a trailing slash when a new directory is created
      
      Empty folders are unsupported for bucket based remotes, this option creates an empty
      object ending with "/", to persist the folder.
      

   --use-multipart-etag
      Whether to use ETag in multipart uploads for verification
      
      This should be true, false or left unset to use the default for the provider.
      

   --use-unsigned-payload
      Whether to use an unsigned payload in PutObject
      
      Rclone has to avoid the AWS SDK seeking the body when calling
      PutObject. The AWS provider can add checksums in the trailer to avoid
      seeking but other providers can't.
      
      This should be true, false or left unset to use the default for the provider.
      

   --use-presigned-request
      Whether to use a presigned request or PutObject for single part uploads
      
      If this is false rclone will use PutObject from the AWS SDK to upload
      an object.
      
      Versions of rclone < 1.59 use presigned requests to upload a single
      part object and setting this flag to true will re-enable that
      functionality. This shouldn't be necessary except in exceptional
      circumstances or for testing.
      

   --versions
      Include old versions in directory listings.

   --version-at
      Show file versions as they were at the specified time.
      
      The parameter should be a date, "2006-01-02", datetime "2006-01-02
      15:04:05" or a duration for that long ago, eg "100d" or "1h".
      
      Note that when using this no file write operations are permitted,
      so you can't upload files or delete them.
      
      See [the time option docs](/docs/#time-option) for valid formats.
      

   --version-deleted
      Show deleted file markers when using versions.
      
      This shows deleted file markers in the listing when using versions. These will appear
      as 0 size files. The only operation which can be performed on them is deletion.
      
      Deleting a delete marker will reveal the previous version.
      
      Deleted files will always show with a timestamp.
      

   --decompress
      If set this will decompress gzip encoded objects.
      
      It is possible to upload objects to S3 with "Content-Encoding: gzip"
      set. Normally rclone will download these files as compressed objects.
      
      If this flag is set then rclone will decompress these files with
      "Content-Encoding: gzip" as they are received. This means that rclone
      can't check the size and hash but the file contents will be decompressed.
      

   --might-gzip
      Set this if the backend might gzip objects.
      
      Normally providers will not alter objects when they are downloaded. If
      an object was not uploaded with `Content-Encoding: gzip` then it won't
      be set on download.
      
      However some providers may gzip objects even if they weren't uploaded
      with `Content-Encoding: gzip` (eg Cloudflare).
      
      A symptom of this would be receiving errors like
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      If you set this flag and rclone downloads an object with
      Content-Encoding: gzip set and chunked transfer encoding, then rclone
      will decompress the object on the fly.
      
      If this is set to unset (the default) then rclone will choose
      according to the provider setting what to apply, but you can override
      rclone's choice here.
      

   --use-accept-encoding-gzip
      Whether to send `Accept-Encoding: gzip` header.
      
      By default, rclone will append `Accept-Encoding: gzip` to the request to download
      compressed objects whenever possible.
      
      However some providers such as Google Cloud Storage may alter the HTTP headers, breaking
      the signature of the request.
      
      A symptom of this would be receiving errors like
      
        SignatureDoesNotMatch: The request signature we calculated does not match the signature you provided.
      
      In this case, you might want to try disabling this option.
      

   --no-system-metadata
      Suppress setting and reading of system metadata

   --sts-endpoint
      Endpoint for STS (deprecated).
      
      Leave blank if using AWS to use the default endpoint for the region.

   --use-already-exists
      Set if rclone should report BucketAlreadyExists errors on bucket creation.
      
      At some point during the evolution of the s3 protocol, AWS started
      returning an `AlreadyOwnedByYou` error when attempting to create a
      bucket that the user already owned, rather than a
      `BucketAlreadyExists` error.
      
      Unfortunately exactly what has been implemented by s3 clones is a
      little inconsistent, some return `AlreadyOwnedByYou`, some return
      `BucketAlreadyExists` and some return no error at all.
      
      This is important to rclone because it ensures the bucket exists by
      creating it on quite a lot of operations (unless
      `--s3-no-check-bucket` is used).
      
      If rclone knows the provider can return `AlreadyOwnedByYou` or returns
      no error then it can report `BucketAlreadyExists` errors when the user
      attempts to create a bucket not owned by them. Otherwise rclone
      ignores the `BucketAlreadyExists` error which can lead to confusion.
      
      This should be automatically set correctly for all providers rclone
      knows about - please make a bug report if not.
      

   --use-multipart-uploads
      Set if rclone should use multipart uploads.
      
      You can change this if you want to disable the use of multipart uploads.
      This shouldn't be necessary in normal operation.
      
      This should be automatically set correctly for all providers rclone
      knows about - please make a bug report if not.
      

   --sdk-log-mode
      Set to debug the SDK
      
      This can be set to a comma separated list of the following functions:
      
      - `Signing`
      - `Retries`
      - `Request`
      - `RequestWithBody`
      - `Response`
      - `ResponseWithBody`
      - `DeprecatedUsage`
      - `RequestEventMessage`
      - `ResponseEventMessage`
      
      Use `Off` to disable and `All` to set all log levels. You will need to
      use `-vv` to see the debug level logs.
      

   --description
      Description of the remote.


OPTIONS:
   --access-key-id value           AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                     Canned ACL used when creating buckets and storing or copying objects. [$ACL]
   --endpoint value                Endpoint for S3 API. [$ENDPOINT]
   --env-auth                      Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars). (default: false) [$ENV_AUTH]
   --help, -h                      show help
   --location-constraint value     Location constraint - must be set to match the Region. [$LOCATION_CONSTRAINT]
   --region value                  Region to connect to. [$REGION]
   --secret-access-key value       AWS Secret Access Key (password). [$SECRET_ACCESS_KEY]
   --server-side-encryption value  The server-side encryption algorithm used when storing this object in S3. [$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          If using KMS ID you must provide the ARN of Key. [$SSE_KMS_KEY_ID]
   --storage-class value           The storage class to use when storing new objects in S3. [$STORAGE_CLASS]

   Advanced

   --bucket-acl value                                Canned ACL used when creating buckets. [$BUCKET_ACL]
   --chunk-size value                                Chunk size to use for uploading. (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value                               Cutoff for switching to multipart copy. (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                                      If set this will decompress gzip encoded objects. (default: false) [$DECOMPRESS]
   --description value                               Description of the remote. [$DESCRIPTION]
   --directory-markers                               Upload an empty object with a trailing slash when a new directory is created (default: false) [$DIRECTORY_MARKERS]
   --disable-checksum                                Don't store MD5 checksum with object metadata. (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                                   Disable usage of http2 for S3 backends. (default: false) [$DISABLE_HTTP2]
   --download-url value                              Custom endpoint for downloads. [$DOWNLOAD_URL]
   --encoding value                                  The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style                                If true use path style access if false use virtual hosted style. (default: true) [$FORCE_PATH_STYLE]
   --leave-parts-on-error                            If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery. (default: false) [$LEAVE_PARTS_ON_ERROR]
   --list-chunk value                                Size of listing chunk (response list for each ListObject S3 request). (default: 1000) [$LIST_CHUNK]
   --list-url-encode value                           Whether to url encode listings: true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value                              Version of ListObjects to use: 1,2 or 0 for auto. (default: 0) [$LIST_VERSION]
   --max-upload-parts value                          Maximum number of parts in a multipart upload. (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value                    How often internal memory buffer pools will be flushed. (no longer used) (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap                            Whether to use mmap buffers in internal memory pool. (no longer used) (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value                                Set this if the backend might gzip objects. (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                                 If set, don't attempt to check the bucket exists or create it. (default: false) [$NO_CHECK_BUCKET]
   --no-head                                         If set, don't HEAD uploaded objects to check integrity. (default: false) [$NO_HEAD]
   --no-head-object                                  If set, do not do HEAD before GET when getting objects. (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata                              Suppress setting and reading of system metadata (default: false) [$NO_SYSTEM_METADATA]
   --profile value                                   Profile to use in the shared credentials file. [$PROFILE]
   --requester-pays                                  Enables requester pays option when interacting with S3 bucket. (default: false) [$REQUESTER_PAYS]
   --sdk-log-mode value                              Set to debug the SDK (default: "Off") [$SDK_LOG_MODE]
   --session-token value                             An AWS session token. [$SESSION_TOKEN]
   --shared-credentials-file value                   Path to the shared credentials file. [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value                    If using SSE-C, the server-side encryption algorithm used when storing this object in S3. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value                          To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value                   If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data. [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value                      If using SSE-C you may provide the secret encryption key MD5 checksum (optional). [$SSE_CUSTOMER_KEY_MD5]
   --sts-endpoint value                              Endpoint for STS (deprecated). [$STS_ENDPOINT]
   --upload-concurrency value                        Concurrency for multipart uploads and copies. (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value                             Cutoff for switching to chunked upload. (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-accelerate-endpoint                         If true use the AWS S3 accelerated endpoint. (default: false) [$USE_ACCELERATE_ENDPOINT]
   --use-accept-encoding-gzip Accept-Encoding: gzip  Whether to send Accept-Encoding: gzip header. (default: "unset") [$USE_ACCEPT_ENCODING_GZIP]
   --use-already-exists value                        Set if rclone should report BucketAlreadyExists errors on bucket creation. (default: "unset") [$USE_ALREADY_EXISTS]
   --use-dual-stack                                  If true use AWS S3 dual-stack endpoint (IPv6 support). (default: false) [$USE_DUAL_STACK]
   --use-multipart-etag value                        Whether to use ETag in multipart uploads for verification (default: "unset") [$USE_MULTIPART_ETAG]
   --use-multipart-uploads value                     Set if rclone should use multipart uploads. (default: "unset") [$USE_MULTIPART_UPLOADS]
   --use-presigned-request                           Whether to use a presigned request or PutObject for single part uploads (default: false) [$USE_PRESIGNED_REQUEST]
   --use-unsigned-payload value                      Whether to use an unsigned payload in PutObject (default: "unset") [$USE_UNSIGNED_PAYLOAD]
   --v2-auth                                         If true use v2 authentication. (default: false) [$V2_AUTH]
   --version-at value                                Show file versions as they were at the specified time. (default: "off") [$VERSION_AT]
   --version-deleted                                 Show deleted file markers when using versions. (default: false) [$VERSION_DELETED]
   --versions                                        Include old versions in directory listings. (default: false) [$VERSIONS]

   Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers. To remove, use empty string.
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth. To remove, use empty string.
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value). This will replace the existing header values. To remove a header, use --http-header "key="". To remove all headers, use --http-header ""
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth. To remove, use empty string.
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-scan-concurrency value                  Max number of concurrent listing requests when scanning data source (default: 1)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-use-server-mod-time                     Use server modified time if possible (default: false)
   --client-user-agent value                        Set the user-agent to a specified string. To remove, use empty string. (default: rclone default)

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
