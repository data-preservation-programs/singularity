# Microsoft Azure Blob Storage

{% code fullWidth="true" %}
```
NAME:
   singularity storage update azureblob - Microsoft Azure Blob Storage

USAGE:
   singularity storage update azureblob [command options] <name|id>

DESCRIPTION:
   --account
      Azure Storage Account Name.
      
      Set this to the Azure Storage Account Name in use.
      
      Leave blank to use SAS URL or Emulator, otherwise it needs to be set.
      
      If this is blank and if env_auth is set it will be read from the
      environment variable `AZURE_STORAGE_ACCOUNT_NAME` if possible.
      

   --env-auth
      Read credentials from runtime (environment variables, CLI or MSI).
      
      See the [authentication docs](/azureblob#authentication) for full info.

   --key
      Storage Account Shared Key.
      
      Leave blank to use SAS URL or Emulator.

   --sas-url
      SAS URL for container level access only.
      
      Leave blank if using account/key or Emulator.

   --tenant
      ID of the service principal's tenant. Also called its directory ID.
      
      Set this if using
      - Service principal with client secret
      - Service principal with certificate
      - User with username and password
      

   --client-id
      The ID of the client in use.
      
      Set this if using
      - Service principal with client secret
      - Service principal with certificate
      - User with username and password
      

   --client-secret
      One of the service principal's client secrets
      
      Set this if using
      - Service principal with client secret
      

   --client-certificate-path
      Path to a PEM or PKCS12 certificate file including the private key.
      
      Set this if using
      - Service principal with certificate
      

   --client-certificate-password
      Password for the certificate file (optional).
      
      Optionally set this if using
      - Service principal with certificate
      
      And the certificate has a password.
      

   --client-send-certificate-chain
      Send the certificate chain when using certificate auth.
      
      Specifies whether an authentication request will include an x5c header
      to support subject name / issuer based authentication. When set to
      true, authentication requests include the x5c header.
      
      Optionally set this if using
      - Service principal with certificate
      

   --username
      User name (usually an email address)
      
      Set this if using
      - User with username and password
      

   --password
      The user's password
      
      Set this if using
      - User with username and password
      

   --service-principal-file
      Path to file containing credentials for use with a service principal.
      
      Leave blank normally. Needed only if you want to use a service principal instead of interactive login.
      
          $ az ad sp create-for-rbac --name "<name>" \
            --role "Storage Blob Data Owner" \
            --scopes "/subscriptions/<subscription>/resourceGroups/<resource-group>/providers/Microsoft.Storage/storageAccounts/<storage-account>/blobServices/default/containers/<container>" \
            > azure-principal.json
      
      See ["Create an Azure service principal"](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli) and ["Assign an Azure role for access to blob data"](https://docs.microsoft.com/en-us/azure/storage/common/storage-auth-aad-rbac-cli) pages for more details.
      
      It may be more convenient to put the credentials directly into the
      rclone config file under the `client_id`, `tenant` and `client_secret`
      keys instead of setting `service_principal_file`.
      

   --use-msi
      Use a managed service identity to authenticate (only works in Azure).
      
      When true, use a [managed service identity](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/)
      to authenticate to Azure Storage instead of a SAS token or account key.
      
      If the VM(SS) on which this program is running has a system-assigned identity, it will
      be used by default. If the resource has no system-assigned but exactly one user-assigned identity,
      the user-assigned identity will be used by default. If the resource has multiple user-assigned
      identities, the identity to use must be explicitly specified using exactly one of the msi_object_id,
      msi_client_id, or msi_mi_res_id parameters.

   --msi-object-id
      Object ID of the user-assigned MSI to use, if any.
      
      Leave blank if msi_client_id or msi_mi_res_id specified.

   --msi-client-id
      Object ID of the user-assigned MSI to use, if any.
      
      Leave blank if msi_object_id or msi_mi_res_id specified.

   --msi-mi-res-id
      Azure resource ID of the user-assigned MSI to use, if any.
      
      Leave blank if msi_client_id or msi_object_id specified.

   --use-emulator
      Uses local storage emulator if provided as 'true'.
      
      Leave blank if using real azure storage endpoint.

   --endpoint
      Endpoint for the service.
      
      Leave blank normally.

   --upload-cutoff
      Cutoff for switching to chunked upload (<= 256 MiB) (deprecated).

   --chunk-size
      Upload chunk size.
      
      Note that this is stored in memory and there may be up to
      "--transfers" * "--azureblob-upload-concurrency" chunks stored at once
      in memory.

   --upload-concurrency
      Concurrency for multipart uploads.
      
      This is the number of chunks of the same file that are uploaded
      concurrently.
      
      If you are uploading small numbers of large files over high-speed
      links and these uploads do not fully utilize your bandwidth, then
      increasing this may help to speed up the transfers.
      
      In tests, upload speed increases almost linearly with upload
      concurrency. For example to fill a gigabit pipe it may be necessary to
      raise this to 64. Note that this will use more memory.
      
      Note that chunks are stored in memory and there may be up to
      "--transfers" * "--azureblob-upload-concurrency" chunks stored at once
      in memory.

   --list-chunk
      Size of blob list.
      
      This sets the number of blobs requested in each listing chunk. Default
      is the maximum, 5000. "List blobs" requests are permitted 2 minutes
      per megabyte to complete. If an operation is taking longer than 2
      minutes per megabyte on average, it will time out (
      [source](https://docs.microsoft.com/en-us/rest/api/storageservices/setting-timeouts-for-blob-service-operations#exceptions-to-default-timeout-interval)
      ). This can be used to limit the number of blobs items to return, to
      avoid the time out.

   --access-tier
      Access tier of blob: hot, cool or archive.
      
      Archived blobs can be restored by setting access tier to hot or
      cool. Leave blank if you intend to use default access tier, which is
      set at account level
      
      If there is no "access tier" specified, rclone doesn't apply any tier.
      rclone performs "Set Tier" operation on blobs while uploading, if objects
      are not modified, specifying "access tier" to new one will have no effect.
      If blobs are in "archive tier" at remote, trying to perform data transfer
      operations from remote will not be allowed. User should first restore by
      tiering blob to "Hot" or "Cool".

   --archive-tier-delete
      Delete archive tier blobs before overwriting.
      
      Archive tier blobs cannot be updated. So without this flag, if you
      attempt to update an archive tier blob, then rclone will produce the
      error:
      
          can't update archive tier blob without --azureblob-archive-tier-delete
      
      With this flag set then before rclone attempts to overwrite an archive
      tier blob, it will delete the existing blob before uploading its
      replacement.  This has the potential for data loss if the upload fails
      (unlike updating a normal blob) and also may cost more since deleting
      archive tier blobs early may be chargable.
      

   --disable-checksum
      Don't store MD5 checksum with object metadata.
      
      Normally rclone will calculate the MD5 checksum of the input before
      uploading it so it can add it to metadata on the object. This is great
      for data integrity checking but can cause long delays for large files
      to start uploading.

   --memory-pool-flush-time
      How often internal memory buffer pools will be flushed.
      
      Uploads which requires additional buffers (f.e multipart) will use memory pool for allocations.
      This option controls how often unused buffers will be removed from the pool.

   --memory-pool-use-mmap
      Whether to use mmap buffers in internal memory pool.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --public-access
      Public access level of a container: blob or container.

      Examples:
         | <unset>   | The container and its blobs can be accessed only with an authorized request.
         |           | It's a default value.
         | blob      | Blob data within this container can be read via anonymous request.
         | container | Allow full public read access for container and blob data.

   --no-check-container
      If set, don't attempt to check the container exists or create it.
      
      This can be useful when trying to minimise the number of transactions
      rclone does if you know the container exists already.
      

   --no-head-object
      If set, do not do HEAD before GET when getting objects.


OPTIONS:
   --account value                      Azure Storage Account Name. [$ACCOUNT]
   --client-certificate-password value  Password for the certificate file (optional). [$CLIENT_CERTIFICATE_PASSWORD]
   --client-certificate-path value      Path to a PEM or PKCS12 certificate file including the private key. [$CLIENT_CERTIFICATE_PATH]
   --client-id value                    The ID of the client in use. [$CLIENT_ID]
   --client-secret value                One of the service principal's client secrets [$CLIENT_SECRET]
   --env-auth                           Read credentials from runtime (environment variables, CLI or MSI). (default: false) [$ENV_AUTH]
   --help, -h                           show help
   --key value                          Storage Account Shared Key. [$KEY]
   --sas-url value                      SAS URL for container level access only. [$SAS_URL]
   --tenant value                       ID of the service principal's tenant. Also called its directory ID. [$TENANT]

   Advanced

   --access-tier value              Access tier of blob: hot, cool or archive. [$ACCESS_TIER]
   --archive-tier-delete            Delete archive tier blobs before overwriting. (default: false) [$ARCHIVE_TIER_DELETE]
   --chunk-size value               Upload chunk size. (default: "4Mi") [$CHUNK_SIZE]
   --client-send-certificate-chain  Send the certificate chain when using certificate auth. (default: false) [$CLIENT_SEND_CERTIFICATE_CHAIN]
   --disable-checksum               Don't store MD5 checksum with object metadata. (default: false) [$DISABLE_CHECKSUM]
   --encoding value                 The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8") [$ENCODING]
   --endpoint value                 Endpoint for the service. [$ENDPOINT]
   --list-chunk value               Size of blob list. (default: 5000) [$LIST_CHUNK]
   --memory-pool-flush-time value   How often internal memory buffer pools will be flushed. (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           Whether to use mmap buffers in internal memory pool. (default: false) [$MEMORY_POOL_USE_MMAP]
   --msi-client-id value            Object ID of the user-assigned MSI to use, if any. [$MSI_CLIENT_ID]
   --msi-mi-res-id value            Azure resource ID of the user-assigned MSI to use, if any. [$MSI_MI_RES_ID]
   --msi-object-id value            Object ID of the user-assigned MSI to use, if any. [$MSI_OBJECT_ID]
   --no-check-container             If set, don't attempt to check the container exists or create it. (default: false) [$NO_CHECK_CONTAINER]
   --no-head-object                 If set, do not do HEAD before GET when getting objects. (default: false) [$NO_HEAD_OBJECT]
   --password value                 The user's password [$PASSWORD]
   --public-access value            Public access level of a container: blob or container. [$PUBLIC_ACCESS]
   --service-principal-file value   Path to file containing credentials for use with a service principal. [$SERVICE_PRINCIPAL_FILE]
   --upload-concurrency value       Concurrency for multipart uploads. (default: 16) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            Cutoff for switching to chunked upload (<= 256 MiB) (deprecated). [$UPLOAD_CUTOFF]
   --use-emulator                   Uses local storage emulator if provided as 'true'. (default: false) [$USE_EMULATOR]
   --use-msi                        Use a managed service identity to authenticate (only works in Azure). (default: false) [$USE_MSI]
   --username value                 User name (usually an email address) [$USERNAME]

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
   --client-user-agent value                        Set the user-agent to a specified string. To remove, use empty string. (default: rclone/v1.62.2-DEV)

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
