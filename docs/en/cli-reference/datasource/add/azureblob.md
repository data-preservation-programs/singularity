# Microsoft Azure Blob Storage

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add azureblob - Microsoft Azure Blob Storage

USAGE:
   singularity datasource add azureblob [command options] <dataset_name> <source_path>

DESCRIPTION:
   --azureblob-access-tier
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

   --azureblob-account
      Azure Storage Account Name.
      
      Set this to the Azure Storage Account Name in use.
      
      Leave blank to use SAS URL or Emulator, otherwise it needs to be set.
      
      If this is blank and if env_auth is set it will be read from the
      environment variable `AZURE_STORAGE_ACCOUNT_NAME` if possible.
      

   --azureblob-archive-tier-delete
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
      

   --azureblob-chunk-size
      Upload chunk size.
      
      Note that this is stored in memory and there may be up to
      "--transfers" * "--azureblob-upload-concurrency" chunks stored at once
      in memory.

   --azureblob-client-certificate-password
      Password for the certificate file (optional).
      
      Optionally set this if using
      - Service principal with certificate
      
      And the certificate has a password.
      

   --azureblob-client-certificate-path
      Path to a PEM or PKCS12 certificate file including the private key.
      
      Set this if using
      - Service principal with certificate
      

   --azureblob-client-id
      The ID of the client in use.
      
      Set this if using
      - Service principal with client secret
      - Service principal with certificate
      - User with username and password
      

   --azureblob-client-secret
      One of the service principal's client secrets
      
      Set this if using
      - Service principal with client secret
      

   --azureblob-client-send-certificate-chain
      Send the certificate chain when using certificate auth.
      
      Specifies whether an authentication request will include an x5c header
      to support subject name / issuer based authentication. When set to
      true, authentication requests include the x5c header.
      
      Optionally set this if using
      - Service principal with certificate
      

   --azureblob-disable-checksum
      Don't store MD5 checksum with object metadata.
      
      Normally rclone will calculate the MD5 checksum of the input before
      uploading it so it can add it to metadata on the object. This is great
      for data integrity checking but can cause long delays for large files
      to start uploading.

   --azureblob-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --azureblob-endpoint
      Endpoint for the service.
      
      Leave blank normally.

   --azureblob-env-auth
      Read credentials from runtime (environment variables, CLI or MSI).
      
      See the [authentication docs](/azureblob#authentication) for full info.

   --azureblob-key
      Storage Account Shared Key.
      
      Leave blank to use SAS URL or Emulator.

   --azureblob-list-chunk
      Size of blob list.
      
      This sets the number of blobs requested in each listing chunk. Default
      is the maximum, 5000. "List blobs" requests are permitted 2 minutes
      per megabyte to complete. If an operation is taking longer than 2
      minutes per megabyte on average, it will time out (
      [source](https://docs.microsoft.com/en-us/rest/api/storageservices/setting-timeouts-for-blob-service-operations#exceptions-to-default-timeout-interval)
      ). This can be used to limit the number of blobs items to return, to
      avoid the time out.

   --azureblob-memory-pool-flush-time
      How often internal memory buffer pools will be flushed.
      
      Uploads which requires additional buffers (f.e multipart) will use memory pool for allocations.
      This option controls how often unused buffers will be removed from the pool.

   --azureblob-memory-pool-use-mmap
      Whether to use mmap buffers in internal memory pool.

   --azureblob-msi-client-id
      Object ID of the user-assigned MSI to use, if any.
      
      Leave blank if msi_object_id or msi_mi_res_id specified.

   --azureblob-msi-mi-res-id
      Azure resource ID of the user-assigned MSI to use, if any.
      
      Leave blank if msi_client_id or msi_object_id specified.

   --azureblob-msi-object-id
      Object ID of the user-assigned MSI to use, if any.
      
      Leave blank if msi_client_id or msi_mi_res_id specified.

   --azureblob-no-check-container
      If set, don't attempt to check the container exists or create it.
      
      This can be useful when trying to minimise the number of transactions
      rclone does if you know the container exists already.
      

   --azureblob-no-head-object
      If set, do not do HEAD before GET when getting objects.

   --azureblob-password
      The user's password
      
      Set this if using
      - User with username and password
      

   --azureblob-public-access
      Public access level of a container: blob or container.

      Examples:
         | <unset>   | The container and its blobs can be accessed only with an authorized request.
                     | It's a default value.
         | blob      | Blob data within this container can be read via anonymous request.
         | container | Allow full public read access for container and blob data.

   --azureblob-sas-url
      SAS URL for container level access only.
      
      Leave blank if using account/key or Emulator.

   --azureblob-service-principal-file
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
      

   --azureblob-tenant
      ID of the service principal's tenant. Also called its directory ID.
      
      Set this if using
      - Service principal with client secret
      - Service principal with certificate
      - User with username and password
      

   --azureblob-upload-concurrency
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

   --azureblob-upload-cutoff
      Cutoff for switching to chunked upload (<= 256 MiB) (deprecated).

   --azureblob-use-emulator
      Uses local storage emulator if provided as 'true'.
      
      Leave blank if using real azure storage endpoint.

   --azureblob-use-msi
      Use a managed service identity to authenticate (only works in Azure).
      
      When true, use a [managed service identity](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/)
      to authenticate to Azure Storage instead of a SAS token or account key.
      
      If the VM(SS) on which this program is running has a system-assigned identity, it will
      be used by default. If the resource has no system-assigned but exactly one user-assigned identity,
      the user-assigned identity will be used by default. If the resource has multiple user-assigned
      identities, the identity to use must be explicitly specified using exactly one of the msi_object_id,
      msi_client_id, or msi_mi_res_id parameters.

   --azureblob-username
      User name (usually an email address)
      
      Set this if using
      - User with username and password
      


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)
   --scanning-state value   set the initial scanning state (default: ready)

   Options for azureblob

   --azureblob-access-tier value                    Access tier of blob: hot, cool or archive. [$AZUREBLOB_ACCESS_TIER]
   --azureblob-account value                        Azure Storage Account Name. [$AZUREBLOB_ACCOUNT]
   --azureblob-archive-tier-delete value            Delete archive tier blobs before overwriting. (default: "false") [$AZUREBLOB_ARCHIVE_TIER_DELETE]
   --azureblob-chunk-size value                     Upload chunk size. (default: "4Mi") [$AZUREBLOB_CHUNK_SIZE]
   --azureblob-client-certificate-password value    Password for the certificate file (optional). [$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-certificate-path value        Path to a PEM or PKCS12 certificate file including the private key. [$AZUREBLOB_CLIENT_CERTIFICATE_PATH]
   --azureblob-client-id value                      The ID of the client in use. [$AZUREBLOB_CLIENT_ID]
   --azureblob-client-secret value                  One of the service principal's client secrets [$AZUREBLOB_CLIENT_SECRET]
   --azureblob-client-send-certificate-chain value  Send the certificate chain when using certificate auth. (default: "false") [$AZUREBLOB_CLIENT_SEND_CERTIFICATE_CHAIN]
   --azureblob-disable-checksum value               Don't store MD5 checksum with object metadata. (default: "false") [$AZUREBLOB_DISABLE_CHECKSUM]
   --azureblob-encoding value                       The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8") [$AZUREBLOB_ENCODING]
   --azureblob-endpoint value                       Endpoint for the service. [$AZUREBLOB_ENDPOINT]
   --azureblob-env-auth value                       Read credentials from runtime (environment variables, CLI or MSI). (default: "false") [$AZUREBLOB_ENV_AUTH]
   --azureblob-key value                            Storage Account Shared Key. [$AZUREBLOB_KEY]
   --azureblob-list-chunk value                     Size of blob list. (default: "5000") [$AZUREBLOB_LIST_CHUNK]
   --azureblob-memory-pool-flush-time value         How often internal memory buffer pools will be flushed. (default: "1m0s") [$AZUREBLOB_MEMORY_POOL_FLUSH_TIME]
   --azureblob-memory-pool-use-mmap value           Whether to use mmap buffers in internal memory pool. (default: "false") [$AZUREBLOB_MEMORY_POOL_USE_MMAP]
   --azureblob-msi-client-id value                  Object ID of the user-assigned MSI to use, if any. [$AZUREBLOB_MSI_CLIENT_ID]
   --azureblob-msi-mi-res-id value                  Azure resource ID of the user-assigned MSI to use, if any. [$AZUREBLOB_MSI_MI_RES_ID]
   --azureblob-msi-object-id value                  Object ID of the user-assigned MSI to use, if any. [$AZUREBLOB_MSI_OBJECT_ID]
   --azureblob-no-check-container value             If set, don't attempt to check the container exists or create it. (default: "false") [$AZUREBLOB_NO_CHECK_CONTAINER]
   --azureblob-no-head-object value                 If set, do not do HEAD before GET when getting objects. (default: "false") [$AZUREBLOB_NO_HEAD_OBJECT]
   --azureblob-password value                       The user's password [$AZUREBLOB_PASSWORD]
   --azureblob-public-access value                  Public access level of a container: blob or container. [$AZUREBLOB_PUBLIC_ACCESS]
   --azureblob-sas-url value                        SAS URL for container level access only. [$AZUREBLOB_SAS_URL]
   --azureblob-service-principal-file value         Path to file containing credentials for use with a service principal. [$AZUREBLOB_SERVICE_PRINCIPAL_FILE]
   --azureblob-tenant value                         ID of the service principal's tenant. Also called its directory ID. [$AZUREBLOB_TENANT]
   --azureblob-upload-concurrency value             Concurrency for multipart uploads. (default: "16") [$AZUREBLOB_UPLOAD_CONCURRENCY]
   --azureblob-upload-cutoff value                  Cutoff for switching to chunked upload (<= 256 MiB) (deprecated). [$AZUREBLOB_UPLOAD_CUTOFF]
   --azureblob-use-emulator value                   Uses local storage emulator if provided as 'true'. (default: "false") [$AZUREBLOB_USE_EMULATOR]
   --azureblob-use-msi value                        Use a managed service identity to authenticate (only works in Azure). (default: "false") [$AZUREBLOB_USE_MSI]
   --azureblob-username value                       User name (usually an email address) [$AZUREBLOB_USERNAME]

```
{% endcode %}
