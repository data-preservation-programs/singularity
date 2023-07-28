# Oracle Cloud Infrastructure Object Storage

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add oos - Oracle Cloud Infrastructure Object Storage

USAGE:
   singularity datasource add oos [command options] <dataset_name> <source_path>

DESCRIPTION:
   --oos-chunk-size
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
      

   --oos-compartment
      [Provider] - user_principal_auth
         Object storage compartment OCID

   --oos-config-file
      [Provider] - user_principal_auth
         Path to OCI config file

         Examples:
            | ~/.oci/config | oci configuration file location

   --oos-config-profile
      [Provider] - user_principal_auth
         Profile name inside the oci config file

         Examples:
            | Default | Use the default profile

   --oos-copy-cutoff
      Cutoff for switching to multipart copy.
      
      Any files larger than this that need to be server-side copied will be
      copied in chunks of this size.
      
      The minimum is 0 and the maximum is 5 GiB.

   --oos-copy-timeout
      Timeout for copy.
      
      Copy is an asynchronous operation, specify timeout to wait for copy to succeed
      

   --oos-disable-checksum
      Don't store MD5 checksum with object metadata.
      
      Normally rclone will calculate the MD5 checksum of the input before
      uploading it so it can add it to metadata on the object. This is great
      for data integrity checking but can cause long delays for large files
      to start uploading.

   --oos-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --oos-endpoint
      Endpoint for Object storage API.
      
      Leave blank to use the default endpoint for the region.

   --oos-leave-parts-on-error
      If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
      
      It should be set to true for resuming uploads across different sessions.
      
      WARNING: Storing parts of an incomplete multipart upload counts towards space usage on object storage and will add
      additional costs if not cleaned up.
      

   --oos-namespace
      Object storage namespace

   --oos-no-check-bucket
      If set, don't attempt to check the bucket exists or create it.
      
      This can be useful when trying to minimise the number of transactions
      rclone does if you know the bucket exists already.
      
      It can also be needed if the user you are using does not have bucket
      creation permissions.
      

   --oos-provider
      Choose your Auth Provider

      Examples:
         | env_auth                | automatically pickup the credentials from runtime(env), first one to provide auth wins
         | user_principal_auth     | use an OCI user and an API key for authentication.
                                   | you’ll need to put in a config file your tenancy OCID, user OCID, region, the path, fingerprint to an API key.
                                   | https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm
         | instance_principal_auth | use instance principals to authorize an instance to make API calls. 
                                   | each instance has its own identity, and authenticates using the certificates that are read from instance metadata. 
                                   | https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm
         | resource_principal_auth | use resource principals to make API calls
         | no_auth                 | no credentials needed, this is typically for reading public buckets

   --oos-region
      Object storage Region

   --oos-sse-customer-algorithm
      If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
      Object Storage supports "AES256" as the encryption algorithm. For more information, see
      Using Your Own Keys for Server-Side Encryption (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm).

      Examples:
         | <unset> | None
         | AES256  | AES256

   --oos-sse-customer-key
      To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
      encrypt or  decrypt the data. Please note only one of sse_customer_key_file|sse_customer_key|sse_kms_key_id is
      needed. For more information, see Using Your Own Keys for Server-Side Encryption 
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)

      Examples:
         | <unset> | None

   --oos-sse-customer-key-file
      To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
      with the object. Please note only one of sse_customer_key_file|sse_customer_key|sse_kms_key_id is needed.'

      Examples:
         | <unset> | None

   --oos-sse-customer-key-sha256
      If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
      key. This value is used to check the integrity of the encryption key. see Using Your Own Keys for 
      Server-Side Encryption (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm).

      Examples:
         | <unset> | None

   --oos-sse-kms-key-id
      if using using your own master key in vault, this header specifies the 
      OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of a master encryption key used to call
      the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.
      Please note only one of sse_customer_key_file|sse_customer_key|sse_kms_key_id is needed.

      Examples:
         | <unset> | None

   --oos-storage-tier
      The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      Examples:
         | Standard         | Standard storage tier, this is the default tier
         | InfrequentAccess | InfrequentAccess storage tier
         | Archive          | Archive storage tier

   --oos-upload-concurrency
      Concurrency for multipart uploads.
      
      This is the number of chunks of the same file that are uploaded
      concurrently.
      
      If you are uploading small numbers of large files over high-speed links
      and these uploads do not fully utilize your bandwidth, then increasing
      this may help to speed up the transfers.

   --oos-upload-cutoff
      Cutoff for switching to chunked upload.
      
      Any files larger than this will be uploaded in chunks of chunk_size.
      The minimum is 0 and the maximum is 5 GiB.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)
   --scanning-state value   set the initial scanning state (default: ready)

   Options for oos

   --oos-chunk-size value               Chunk size to use for uploading. (default: "5Mi") [$OOS_CHUNK_SIZE]
   --oos-compartment value              Object storage compartment OCID [$OOS_COMPARTMENT]
   --oos-config-file value              Path to OCI config file (default: "~/.oci/config") [$OOS_CONFIG_FILE]
   --oos-config-profile value           Profile name inside the oci config file (default: "Default") [$OOS_CONFIG_PROFILE]
   --oos-copy-cutoff value              Cutoff for switching to multipart copy. (default: "4.656Gi") [$OOS_COPY_CUTOFF]
   --oos-copy-timeout value             Timeout for copy. (default: "1m0s") [$OOS_COPY_TIMEOUT]
   --oos-disable-checksum value         Don't store MD5 checksum with object metadata. (default: "false") [$OOS_DISABLE_CHECKSUM]
   --oos-encoding value                 The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$OOS_ENCODING]
   --oos-endpoint value                 Endpoint for Object storage API. [$OOS_ENDPOINT]
   --oos-leave-parts-on-error value     If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery. (default: "false") [$OOS_LEAVE_PARTS_ON_ERROR]
   --oos-namespace value                Object storage namespace [$OOS_NAMESPACE]
   --oos-no-check-bucket value          If set, don't attempt to check the bucket exists or create it. (default: "false") [$OOS_NO_CHECK_BUCKET]
   --oos-provider value                 Choose your Auth Provider (default: "env_auth") [$OOS_PROVIDER]
   --oos-region value                   Object storage Region [$OOS_REGION]
   --oos-sse-customer-algorithm value   If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm. [$OOS_SSE_CUSTOMER_ALGORITHM]
   --oos-sse-customer-key value         To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to [$OOS_SSE_CUSTOMER_KEY]
   --oos-sse-customer-key-file value    To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated [$OOS_SSE_CUSTOMER_KEY_FILE]
   --oos-sse-customer-key-sha256 value  If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption [$OOS_SSE_CUSTOMER_KEY_SHA256]
   --oos-sse-kms-key-id value           if using using your own master key in vault, this header specifies the [$OOS_SSE_KMS_KEY_ID]
   --oos-storage-tier value             The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (default: "Standard") [$OOS_STORAGE_TIER]
   --oos-upload-concurrency value       Concurrency for multipart uploads. (default: "10") [$OOS_UPLOAD_CONCURRENCY]
   --oos-upload-cutoff value            Cutoff for switching to chunked upload. (default: "200Mi") [$OOS_UPLOAD_CUTOFF]

```
{% endcode %}
