# use resource principals to make API calls

{% code fullWidth="true" %}
```
NAME:
   singularity storage update oos resource_principal_auth - use resource principals to make API calls

USAGE:
   singularity storage update oos resource_principal_auth [command options] <name|id>

DESCRIPTION:
   --namespace
      Object storage namespace

   --compartment
      Object storage compartment OCID

   --region
      Object storage Region

   --endpoint
      Endpoint for Object storage API.
      
      Leave blank to use the default endpoint for the region.

   --storage-tier
      The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      Examples:
         | Standard         | Standard storage tier, this is the default tier
         | InfrequentAccess | InfrequentAccess storage tier
         | Archive          | Archive storage tier

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
      

   --upload-concurrency
      Concurrency for multipart uploads.
      
      This is the number of chunks of the same file that are uploaded
      concurrently.
      
      If you are uploading small numbers of large files over high-speed links
      and these uploads do not fully utilize your bandwidth, then increasing
      this may help to speed up the transfers.

   --copy-cutoff
      Cutoff for switching to multipart copy.
      
      Any files larger than this that need to be server-side copied will be
      copied in chunks of this size.
      
      The minimum is 0 and the maximum is 5 GiB.

   --copy-timeout
      Timeout for copy.
      
      Copy is an asynchronous operation, specify timeout to wait for copy to succeed
      

   --disable-checksum
      Don't store MD5 checksum with object metadata.
      
      Normally rclone will calculate the MD5 checksum of the input before
      uploading it so it can add it to metadata on the object. This is great
      for data integrity checking but can cause long delays for large files
      to start uploading.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --leave-parts-on-error
      If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
      
      It should be set to true for resuming uploads across different sessions.
      
      WARNING: Storing parts of an incomplete multipart upload counts towards space usage on object storage and will add
      additional costs if not cleaned up.
      

   --no-check-bucket
      If set, don't attempt to check the bucket exists or create it.
      
      This can be useful when trying to minimise the number of transactions
      rclone does if you know the bucket exists already.
      
      It can also be needed if the user you are using does not have bucket
      creation permissions.
      

   --sse-customer-key-file
      To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
      with the object. Please note only one of sse_customer_key_file|sse_customer_key|sse_kms_key_id is needed.'

      Examples:
         | <unset> | None

   --sse-customer-key
      To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
      encrypt or  decrypt the data. Please note only one of sse_customer_key_file|sse_customer_key|sse_kms_key_id is
      needed. For more information, see Using Your Own Keys for Server-Side Encryption 
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)

      Examples:
         | <unset> | None

   --sse-customer-key-sha256
      If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
      key. This value is used to check the integrity of the encryption key. see Using Your Own Keys for 
      Server-Side Encryption (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm).

      Examples:
         | <unset> | None

   --sse-kms-key-id
      if using using your own master key in vault, this header specifies the 
      OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of a master encryption key used to call
      the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.
      Please note only one of sse_customer_key_file|sse_customer_key|sse_kms_key_id is needed.

      Examples:
         | <unset> | None

   --sse-customer-algorithm
      If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
      Object Storage supports "AES256" as the encryption algorithm. For more information, see
      Using Your Own Keys for Server-Side Encryption (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm).

      Examples:
         | <unset> | None
         | AES256  | AES256


OPTIONS:
   --compartment value  Object storage compartment OCID [$COMPARTMENT]
   --endpoint value     Endpoint for Object storage API. [$ENDPOINT]
   --help, -h           show help
   --namespace value    Object storage namespace [$NAMESPACE]
   --region value       Object storage Region [$REGION]

   Advanced

   --chunk-size value               Chunk size to use for uploading. (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              Cutoff for switching to multipart copy. (default: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             Timeout for copy. (default: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               Don't store MD5 checksum with object metadata. (default: false) [$DISABLE_CHECKSUM]
   --encoding value                 The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery. (default: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                If set, don't attempt to check the bucket exists or create it. (default: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           if using using your own master key in vault, this header specifies the [$SSE_KMS_KEY_ID]
   --storage-tier value             The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (default: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       Concurrency for multipart uploads. (default: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            Cutoff for switching to chunked upload. (default: "200Mi") [$UPLOAD_CUTOFF]

   HTTP Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers. To remove, use empty string.
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth. To remove, use empty string.
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value). This will replace the existing header values. To remove a header, use --http-header "key="". To remove all headers, use --http-header ""
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth. To remove, use empty string.
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-user-agent value                        Set the user-agent to a specified string. To remove, use empty string. (default: rclone/v1.62.2-DEV)

   Retry Strategy

   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
