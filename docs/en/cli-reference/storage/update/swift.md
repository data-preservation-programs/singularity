# OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

{% code fullWidth="true" %}
```
NAME:
   singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity storage update swift - OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

USAGE:
   singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity storage update swift [command options] <name>

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

   --chunk-size
      Above this size files will be chunked into a _segments container.
      
      Above this size files will be chunked into a _segments container.  The
      default for this is 5 GiB which is its maximum value.

   --no-chunk
      Don't chunk files during streaming upload.
      
      When doing streaming uploads (e.g. using rcat or mount) setting this
      flag will cause the swift backend to not upload chunked files.
      
      This will limit the maximum upload size to 5 GiB. However non chunked
      files are easier to deal with and have an MD5SUM.
      
      Rclone will still chunk files bigger than chunk_size when doing normal
      copy operations.

   --no-large-objects
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
      

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


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

   --chunk-size value      Above this size files will be chunked into a _segments container. (default: "5Gi") [$CHUNK_SIZE]
   --encoding value        The encoding for the backend. (default: "Slash,InvalidUtf8") [$ENCODING]
   --leave-parts-on-error  If true avoid calling abort upload on a failure. (default: false) [$LEAVE_PARTS_ON_ERROR]
   --no-chunk              Don't chunk files during streaming upload. (default: false) [$NO_CHUNK]
   --no-large-objects      Disable support for static and dynamic large objects (default: false) [$NO_LARGE_OBJECTS]

```
{% endcode %}
