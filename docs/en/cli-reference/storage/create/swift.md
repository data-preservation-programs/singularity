# OpenStack Swift (Rackspace Cloud Files, Blomp Cloud Storage, Memset Memstore, OVH)

{% code fullWidth="true" %}
```
NAME:
   singularity storage create swift - OpenStack Swift (Rackspace Cloud Files, Blomp Cloud Storage, Memset Memstore, OVH)

USAGE:
   singularity storage create swift [command options]

DESCRIPTION:
   --env-auth
      Get swift credentials from environment variables in standard OpenStack form.

      Examples:
         | false | Enter swift credentials in the next step.
         | true  | Get swift credentials from environment vars.
         |       | Leave other fields blank if using this.

   --user
      User name to log in (OS_USERNAME).

   --key
      API key or password (OS_PASSWORD).

   --auth
      Authentication URL for server (OS_AUTH_URL).

      Examples:
         | https://auth.api.rackspacecloud.com/v1.0     | Rackspace US
         | https://lon.auth.api.rackspacecloud.com/v1.0 | Rackspace UK
         | https://identity.api.rackspacecloud.com/v2.0 | Rackspace v2
         | https://auth.storage.memset.com/v1.0         | Memset Memstore UK
         | https://auth.storage.memset.com/v2.0         | Memset Memstore UK v2
         | https://auth.cloud.ovh.net/v3                | OVH
         | https://authenticate.ain.net                 | Blomp Cloud Storage

   --user-id
      User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID).

   --domain
      User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME)

   --tenant
      Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME).

   --tenant-id
      Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID).

   --tenant-domain
      Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME).

   --region
      Region name - optional (OS_REGION_NAME).

   --storage-url
      Storage URL - optional (OS_STORAGE_URL).

   --auth-token
      Auth Token from alternate authentication - optional (OS_AUTH_TOKEN).

   --application-credential-id
      Application Credential ID (OS_APPLICATION_CREDENTIAL_ID).

   --application-credential-name
      Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME).

   --application-credential-secret
      Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET).

   --auth-version
      AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).

   --endpoint-type
      Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).

      Examples:
         | public   | Public (default, choose this if not sure)
         | internal | Internal (use internal service net)
         | admin    | Admin

   --leave-parts-on-error
      If true avoid calling abort upload on a failure.
      
      It should be set to true for resuming uploads across different sessions.

   --storage-policy
      The storage policy to use when creating a new container.
      
      This applies the specified storage policy when creating a new
      container. The policy cannot be changed afterwards. The allowed
      configuration values and their meaning depend on your Swift storage
      provider.

      Examples:
         | <unset> | Default
         | pcs     | OVH Public Cloud Storage
         | pca     | OVH Public Cloud Archive

   --fetch-until-empty-page
      When paginating, always fetch unless we received an empty page.
      
      Consider using this option if rclone listings show fewer objects
      than expected, or if repeated syncs copy unchanged objects.
      
      It is safe to enable this, but rclone may make more API calls than
      necessary.
      
      This is one of a pair of workarounds to handle implementations
      of the Swift API that do not implement pagination as expected.  See
      also "partial_page_fetch_threshold".

   --partial-page-fetch-threshold
      When paginating, fetch if the current page is within this percentage of the limit.
      
      Consider using this option if rclone listings show fewer objects
      than expected, or if repeated syncs copy unchanged objects.
      
      It is safe to enable this, but rclone may make more API calls than
      necessary.
      
      This is one of a pair of workarounds to handle implementations
      of the Swift API that do not implement pagination as expected.  See
      also "fetch_until_empty_page".

   --chunk-size
      Above this size files will be chunked.
      
      Above this size files will be chunked into a a `_segments` container
      or a `.file-segments` directory. (See the `use_segments_container` option
      for more info). Default for this is 5 GiB which is its maximum value, which
      means only files above this size will be chunked.
      
      Rclone uploads chunked files as dynamic large objects (DLO).
      

   --no-chunk
      Don't chunk files during streaming upload.
      
      When doing streaming uploads (e.g. using `rcat` or `mount` with
      `--vfs-cache-mode off`) setting this flag will cause the swift backend
      to not upload chunked files.
      
      This will limit the maximum streamed upload size to 5 GiB. This is
      useful because non chunked files are easier to deal with and have an
      MD5SUM.
      
      Rclone will still chunk files bigger than `chunk_size` when doing
      normal copy operations.

   --no-large-objects
      Disable support for static and dynamic large objects
      
      Swift cannot transparently store files bigger than 5 GiB. There are
      two schemes for chunking large files, static large objects (SLO) or
      dynamic large objects (DLO), and the API does not allow rclone to
      determine whether a file is a static or dynamic large object without
      doing a HEAD on the object. Since these need to be treated
      differently, this means rclone has to issue HEAD requests for objects
      for example when reading checksums.
      
      When `no_large_objects` is set, rclone will assume that there are no
      static or dynamic large objects stored. This means it can stop doing
      the extra HEAD calls which in turn increases performance greatly
      especially when doing a swift to swift transfer with `--checksum` set.
      
      Setting this option implies `no_chunk` and also that no files will be
      uploaded in chunks, so files bigger than 5 GiB will just fail on
      upload.
      
      If you set this option and there **are** static or dynamic large objects,
      then this will give incorrect hashes for them. Downloads will succeed,
      but other operations such as Remove and Copy will fail.
      

   --use-segments-container
      Choose destination for large object segments
      
      Swift cannot transparently store files bigger than 5 GiB and rclone
      will chunk files larger than `chunk_size` (default 5 GiB) in order to
      upload them.
      
      If this value is `true` the chunks will be stored in an additional
      container named the same as the destination container but with
      `_segments` appended. This means that there won't be any duplicated
      data in the original container but having another container may not be
      acceptable.
      
      If this value is `false` the chunks will be stored in a
      `.file-segments` directory in the root of the container. This
      directory will be omitted when listing the container. Some
      providers (eg Blomp) require this mode as creating additional
      containers isn't allowed. If it is desired to see the `.file-segments`
      directory in the root then this flag must be set to `true`.
      
      If this value is `unset` (the default), then rclone will choose the value
      to use. It will be `false` unless rclone detects any `auth_url`s that
      it knows need it to be `true`. In this case you'll see a message in
      the DEBUG log.
      

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --description
      Description of the remote.


OPTIONS:
   --application-credential-id value      Application Credential ID (OS_APPLICATION_CREDENTIAL_ID). [$APPLICATION_CREDENTIAL_ID]
   --application-credential-name value    Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME). [$APPLICATION_CREDENTIAL_NAME]
   --application-credential-secret value  Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET). [$APPLICATION_CREDENTIAL_SECRET]
   --auth value                           Authentication URL for server (OS_AUTH_URL). [$AUTH]
   --auth-token value                     Auth Token from alternate authentication - optional (OS_AUTH_TOKEN). [$AUTH_TOKEN]
   --auth-version value                   AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION). (default: 0) [$AUTH_VERSION]
   --domain value                         User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME) [$DOMAIN]
   --endpoint-type value                  Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE). (default: "public") [$ENDPOINT_TYPE]
   --env-auth                             Get swift credentials from environment variables in standard OpenStack form. (default: false) [$ENV_AUTH]
   --help, -h                             show help
   --key value                            API key or password (OS_PASSWORD). [$KEY]
   --region value                         Region name - optional (OS_REGION_NAME). [$REGION]
   --storage-policy value                 The storage policy to use when creating a new container. [$STORAGE_POLICY]
   --storage-url value                    Storage URL - optional (OS_STORAGE_URL). [$STORAGE_URL]
   --tenant value                         Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME). [$TENANT]
   --tenant-domain value                  Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME). [$TENANT_DOMAIN]
   --tenant-id value                      Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID). [$TENANT_ID]
   --user value                           User name to log in (OS_USERNAME). [$USER]
   --user-id value                        User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID). [$USER_ID]

   Advanced

   --chunk-size value                    Above this size files will be chunked. (default: "5Gi") [$CHUNK_SIZE]
   --description value                   Description of the remote. [$DESCRIPTION]
   --encoding value                      The encoding for the backend. (default: "Slash,InvalidUtf8") [$ENCODING]
   --fetch-until-empty-page              When paginating, always fetch unless we received an empty page. (default: false) [$FETCH_UNTIL_EMPTY_PAGE]
   --leave-parts-on-error                If true avoid calling abort upload on a failure. (default: false) [$LEAVE_PARTS_ON_ERROR]
   --no-chunk                            Don't chunk files during streaming upload. (default: false) [$NO_CHUNK]
   --no-large-objects                    Disable support for static and dynamic large objects (default: false) [$NO_LARGE_OBJECTS]
   --partial-page-fetch-threshold value  When paginating, fetch if the current page is within this percentage of the limit. (default: 0) [$PARTIAL_PAGE_FETCH_THRESHOLD]
   --use-segments-container value        Choose destination for large object segments (default: "unset") [$USE_SEGMENTS_CONTAINER]

   Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value)
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-scan-concurrency value                  Max number of concurrent listing requests when scanning data source (default: 1)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-use-server-mod-time                     Use server modified time if possible (default: false)
   --client-user-agent value                        Set the user-agent to a specified string (default: rclone default)

   General

   --name value  Name of the storage (default: Auto generated)
   --path value  Path of the storage

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
