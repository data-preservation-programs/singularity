# no credentials needed, this is typically for reading public buckets

{% code fullWidth="true" %}
```
NAME:
   singularity storage update oos no_auth - no credentials needed, this is typically for reading public buckets

USAGE:
   singularity storage update oos no_auth [command options] <name>

DESCRIPTION:
   --namespace
      Object storage namespace

   --region
      Object storage Region

   --endpoint
      Endpoint for Object storage API.
      
      Leave blank to use the default endpoint for the region.

   --storage_tier
      The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      Examples:
         | Standard         | Standard storage tier, this is the default tier
         | InfrequentAccess | InfrequentAccess storage tier
         | Archive          | Archive storage tier

   --upload_cutoff
      Cutoff for switching to chunked upload.
      
      Any files larger than this will be uploaded in chunks of chunk_size.
      The minimum is 0 and the maximum is 5 GiB.

   --chunk_size
      Chunk size to use for uploading.
      
      When uploading files larger than upload_cutoff or files with unknown
      size (e.g. from "rclone rcat" or uploaded with "rclone mount" or google
      photos or google docs) they will be uploaded as multipart uploads
      using this chunk size.
      
      Note that "upload_concurrency" chunks of this size are buffered
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
      statistics displayed with "-P" flag.
      

   --upload_concurrency
      Concurrency for multipart uploads.
      
      This is the number of chunks of the same file that are uploaded
      concurrently.
      
      If you are uploading small numbers of large files over high-speed links
      and these uploads do not fully utilize your bandwidth, then increasing
      this may help to speed up the transfers.

   --copy_cutoff
      Cutoff for switching to multipart copy.
      
      Any files larger than this that need to be server-side copied will be
      copied in chunks of this size.
      
      The minimum is 0 and the maximum is 5 GiB.

   --copy_timeout
      Timeout for copy.
      
      Copy is an asynchronous operation, specify timeout to wait for copy to succeed
      

   --disable_checksum
      Don't store MD5 checksum with object metadata.
      
      Normally rclone will calculate the MD5 checksum of the input before
      uploading it so it can add it to metadata on the object. This is great
      for data integrity checking but can cause long delays for large files
      to start uploading.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --leave_parts_on_error
      If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
      
      It should be set to true for resuming uploads across different sessions.
      
      WARNING: Storing parts of an incomplete multipart upload counts towards space usage on object storage and will add
      additional costs if not cleaned up.
      

   --no_check_bucket
      If set, don't attempt to check the bucket exists or create it.
      
      This can be useful when trying to minimise the number of transactions
      rclone does if you know the bucket exists already.
      
      It can also be needed if the user you are using does not have bucket
      creation permissions.
      

   --sse_customer_key_file
      To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
      with the object. Please note only one of sse_customer_key_file|sse_customer_key|sse_kms_key_id is needed.'

      Examples:
         | <unset> | None

   --sse_customer_key
      To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
      encrypt or  decrypt the data. Please note only one of sse_customer_key_file|sse_customer_key|sse_kms_key_id is
      needed. For more information, see Using Your Own Keys for Server-Side Encryption 
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)

      Examples:
         | <unset> | None

   --sse_customer_key_sha256
      If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
      key. This value is used to check the integrity of the encryption key. see Using Your Own Keys for 
      Server-Side Encryption (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm).

      Examples:
         | <unset> | None

   --sse_kms_key_id
      if using using your own master key in vault, this header specifies the 
      OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of a master encryption key used to call
      the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.
      Please note only one of sse_customer_key_file|sse_customer_key|sse_kms_key_id is needed.

      Examples:
         | <unset> | None

   --sse_customer_algorithm
      If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
      Object Storage supports "AES256" as the encryption algorithm. For more information, see
      Using Your Own Keys for Server-Side Encryption (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm).

      Examples:
         | <unset> | None
         | AES256  | AES256


OPTIONS:
   --endpoint value   Endpoint for Object storage API. [$ENDPOINT]
   --help, -h         show help
   --namespace value  Object storage namespace [$NAMESPACE]
   --region value     Object storage Region [$REGION]

   Advanced

   --chunk_size value               Chunk size to use for uploading. (default: "5Mi") [$CHUNK_SIZE]
   --copy_cutoff value              Cutoff for switching to multipart copy. (default: "4.656Gi") [$COPY_CUTOFF]
   --copy_timeout value             Timeout for copy. (default: "1m0s") [$COPY_TIMEOUT]
   --disable_checksum               Don't store MD5 checksum with object metadata. (default: false) [$DISABLE_CHECKSUM]
   --encoding value                 The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave_parts_on_error           If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery. (default: false) [$LEAVE_PARTS_ON_ERROR]
   --no_check_bucket                If set, don't attempt to check the bucket exists or create it. (default: false) [$NO_CHECK_BUCKET]
   --sse_customer_algorithm value   If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm. [$SSE_CUSTOMER_ALGORITHM]
   --sse_customer_key value         To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to [$SSE_CUSTOMER_KEY]
   --sse_customer_key_file value    To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated [$SSE_CUSTOMER_KEY_FILE]
   --sse_customer_key_sha256 value  If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption [$SSE_CUSTOMER_KEY_SHA256]
   --sse_kms_key_id value           if using using your own master key in vault, this header specifies the [$SSE_KMS_KEY_ID]
   --storage_tier value             The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (default: "Standard") [$STORAGE_TIER]
   --upload_concurrency value       Concurrency for multipart uploads. (default: 10) [$UPLOAD_CONCURRENCY]
   --upload_cutoff value            Cutoff for switching to chunked upload. (default: "200Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}
