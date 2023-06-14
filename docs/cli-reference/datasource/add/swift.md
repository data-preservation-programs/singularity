# OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

```
NAME:
   singularity datasource add swift - OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

USAGE:
   singularity datasource add swift [command options] <dataset_name> <source_path>

DESCRIPTION:
   --swift-tenant
      Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME).

   --swift-tenant-id
      Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID).

   --swift-leave-parts-on-error
      If true avoid calling abort upload on a failure.
      
      It should be set to true for resuming uploads across different sessions.

   --swift-no-chunk
      Don't chunk files during streaming upload.
      
      When doing streaming uploads (e.g. using rcat or mount) setting this
      flag will cause the swift backend to not upload chunked files.
      
      This will limit the maximum upload size to 5 GiB. However non chunked
      files are easier to deal with and have an MD5SUM.
      
      Rclone will still chunk files bigger than chunk_size when doing normal
      copy operations.

   --swift-env-auth
      Get swift credentials from environment variables in standard OpenStack form.

      Examples:
         | false | Enter swift credentials in the next step.
         | true  | Get swift credentials from environment vars.
                 | Leave other fields blank if using this.

   --swift-key
      API key or password (OS_PASSWORD).

   --swift-region
      Region name - optional (OS_REGION_NAME).

   --swift-storage-url
      Storage URL - optional (OS_STORAGE_URL).

   --swift-application-credential-id
      Application Credential ID (OS_APPLICATION_CREDENTIAL_ID).

   --swift-user
      User name to log in (OS_USERNAME).

   --swift-auth
      Authentication URL for server (OS_AUTH_URL).

      Examples:
         | https://auth.api.rackspacecloud.com/v1.0     | Rackspace US
         | https://lon.auth.api.rackspacecloud.com/v1.0 | Rackspace UK
         | https://identity.api.rackspacecloud.com/v2.0 | Rackspace v2
         | https://auth.storage.memset.com/v1.0         | Memset Memstore UK
         | https://auth.storage.memset.com/v2.0         | Memset Memstore UK v2
         | https://auth.cloud.ovh.net/v3                | OVH

   --swift-auth-version
      AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).

   --swift-endpoint-type
      Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).

      Examples:
         | public   | Public (default, choose this if not sure)
         | internal | Internal (use internal service net)
         | admin    | Admin

   --swift-storage-policy
      The storage policy to use when creating a new container.
      
      This applies the specified storage policy when creating a new
      container. The policy cannot be changed afterwards. The allowed
      configuration values and their meaning depend on your Swift storage
      provider.

      Examples:
         | <unset> | Default
         | pcs     | OVH Public Cloud Storage
         | pca     | OVH Public Cloud Archive

   --swift-no-large-objects
      Disable support for static and dynamic large objects
      
      Swift cannot transparently store files bigger than 5 GiB. There are
      two schemes for doing that, static or dynamic large objects, and the
      API does not allow rclone to determine whether a file is a static or
      dynamic large object without doing a HEAD on the object. Since these
      need to be treated differently, this means rclone has to issue HEAD
      requests for objects for example when reading checksums.
      
      When `no_large_objects` is set, rclone will assume that there are no
      static or dynamic large objects stored. This means it can stop doing
      the extra HEAD calls which in turn increases performance greatly
      especially when doing a swift to swift transfer with `--checksum` set.
      
      Setting this option implies `no_chunk` and also that no files will be
      uploaded in chunks, so files bigger than 5 GiB will just fail on
      upload.
      
      If you set this option and there *are* static or dynamic large objects,
      then this will give incorrect hashes for them. Downloads will succeed,
      but other operations such as Remove and Copy will fail.
      

   --swift-chunk-size
      Above this size files will be chunked into a _segments container.
      
      Above this size files will be chunked into a _segments container.  The
      default for this is 5 GiB which is its maximum value.

   --swift-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --swift-user-id
      User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID).

   --swift-domain
      User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME)

   --swift-tenant-domain
      Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME).

   --swift-auth-token
      Auth Token from alternate authentication - optional (OS_AUTH_TOKEN).

   --swift-application-credential-name
      Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME).

   --swift-application-credential-secret
      Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET).


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for swift

   --swift-application-credential-id value      Application Credential ID (OS_APPLICATION_CREDENTIAL_ID). [$SWIFT_APPLICATION_CREDENTIAL_ID]
   --swift-application-credential-name value    Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME). [$SWIFT_APPLICATION_CREDENTIAL_NAME]
   --swift-application-credential-secret value  Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET). [$SWIFT_APPLICATION_CREDENTIAL_SECRET]
   --swift-auth value                           Authentication URL for server (OS_AUTH_URL). [$SWIFT_AUTH]
   --swift-auth-token value                     Auth Token from alternate authentication - optional (OS_AUTH_TOKEN). [$SWIFT_AUTH_TOKEN]
   --swift-auth-version value                   AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION). (default: "0") [$SWIFT_AUTH_VERSION]
   --swift-chunk-size value                     Above this size files will be chunked into a _segments container. (default: "5Gi") [$SWIFT_CHUNK_SIZE]
   --swift-domain value                         User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME) [$SWIFT_DOMAIN]
   --swift-encoding value                       The encoding for the backend. (default: "Slash,InvalidUtf8") [$SWIFT_ENCODING]
   --swift-endpoint-type value                  Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE). (default: "public") [$SWIFT_ENDPOINT_TYPE]
   --swift-env-auth value                       Get swift credentials from environment variables in standard OpenStack form. (default: "false") [$SWIFT_ENV_AUTH]
   --swift-key value                            API key or password (OS_PASSWORD). [$SWIFT_KEY]
   --swift-leave-parts-on-error value           If true avoid calling abort upload on a failure. (default: "false") [$SWIFT_LEAVE_PARTS_ON_ERROR]
   --swift-no-chunk value                       Don't chunk files during streaming upload. (default: "false") [$SWIFT_NO_CHUNK]
   --swift-no-large-objects value               Disable support for static and dynamic large objects (default: "false") [$SWIFT_NO_LARGE_OBJECTS]
   --swift-region value                         Region name - optional (OS_REGION_NAME). [$SWIFT_REGION]
   --swift-storage-policy value                 The storage policy to use when creating a new container. [$SWIFT_STORAGE_POLICY]
   --swift-storage-url value                    Storage URL - optional (OS_STORAGE_URL). [$SWIFT_STORAGE_URL]
   --swift-tenant value                         Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME). [$SWIFT_TENANT]
   --swift-tenant-domain value                  Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME). [$SWIFT_TENANT_DOMAIN]
   --swift-tenant-id value                      Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID). [$SWIFT_TENANT_ID]
   --swift-user value                           User name to log in (OS_USERNAME). [$SWIFT_USER]
   --swift-user-id value                        User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID). [$SWIFT_USER_ID]

```
