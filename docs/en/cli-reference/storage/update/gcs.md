# Google Cloud Storage (this is not Google Drive)

{% code fullWidth="true" %}
```
NAME:
   singularity storage update gcs - Google Cloud Storage (this is not Google Drive)

USAGE:
   singularity storage update gcs [command options] <name|id>

DESCRIPTION:
   --client-id
      OAuth Client Id.
      
      Leave blank normally.

   --client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --token
      OAuth Access Token as a JSON blob.

   --auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --project-number
      Project number.
      
      Optional - needed only for list/create/delete buckets - see your developer console.

   --service-account-file
      Service Account Credentials JSON file path.
      
      Leave blank normally.
      Needed only if you want use SA instead of interactive login.
      
      Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.

   --service-account-credentials
      Service Account Credentials JSON blob.
      
      Leave blank normally.
      Needed only if you want use SA instead of interactive login.

   --anonymous
      Access public buckets and objects without credentials.
      
      Set to 'true' if you just want to download files and don't configure credentials.

   --object-acl
      Access Control List for new objects.

      Examples:
         | authenticatedRead      | Object owner gets OWNER access.
         |                        | All Authenticated Users get READER access.
         | bucketOwnerFullControl | Object owner gets OWNER access.
         |                        | Project team owners get OWNER access.
         | bucketOwnerRead        | Object owner gets OWNER access.
         |                        | Project team owners get READER access.
         | private                | Object owner gets OWNER access.
         |                        | Default if left blank.
         | projectPrivate         | Object owner gets OWNER access.
         |                        | Project team members get access according to their roles.
         | publicRead             | Object owner gets OWNER access.
         |                        | All Users get READER access.

   --bucket-acl
      Access Control List for new buckets.

      Examples:
         | authenticatedRead | Project team owners get OWNER access.
         |                   | All Authenticated Users get READER access.
         | private           | Project team owners get OWNER access.
         |                   | Default if left blank.
         | projectPrivate    | Project team members get access according to their roles.
         | publicRead        | Project team owners get OWNER access.
         |                   | All Users get READER access.
         | publicReadWrite   | Project team owners get OWNER access.
         |                   | All Users get WRITER access.

   --bucket-policy-only
      Access checks should use bucket-level IAM policies.
      
      If you want to upload objects to a bucket with Bucket Policy Only set
      then you will need to set this.
      
      When it is set, rclone:
      
      - ignores ACLs set on buckets
      - ignores ACLs set on objects
      - creates buckets with Bucket Policy Only set
      
      Docs: https://cloud.google.com/storage/docs/bucket-policy-only
      

   --location
      Location for the newly created buckets.

      Examples:
         | <unset>                 | Empty for default location (US)
         | asia                    | Multi-regional location for Asia
         | eu                      | Multi-regional location for Europe
         | us                      | Multi-regional location for United States
         | asia-east1              | Taiwan
         | asia-east2              | Hong Kong
         | asia-northeast1         | Tokyo
         | asia-northeast2         | Osaka
         | asia-northeast3         | Seoul
         | asia-south1             | Mumbai
         | asia-south2             | Delhi
         | asia-southeast1         | Singapore
         | asia-southeast2         | Jakarta
         | australia-southeast1    | Sydney
         | australia-southeast2    | Melbourne
         | europe-north1           | Finland
         | europe-west1            | Belgium
         | europe-west2            | London
         | europe-west3            | Frankfurt
         | europe-west4            | Netherlands
         | europe-west6            | Zürich
         | europe-central2         | Warsaw
         | us-central1             | Iowa
         | us-east1                | South Carolina
         | us-east4                | Northern Virginia
         | us-west1                | Oregon
         | us-west2                | California
         | us-west3                | Salt Lake City
         | us-west4                | Las Vegas
         | northamerica-northeast1 | Montréal
         | northamerica-northeast2 | Toronto
         | southamerica-east1      | São Paulo
         | southamerica-west1      | Santiago
         | asia1                   | Dual region: asia-northeast1 and asia-northeast2.
         | eur4                    | Dual region: europe-north1 and europe-west4.
         | nam4                    | Dual region: us-central1 and us-east1.

   --storage-class
      The storage class to use when storing objects in Google Cloud Storage.

      Examples:
         | <unset>                      | Default
         | MULTI_REGIONAL               | Multi-regional storage class
         | REGIONAL                     | Regional storage class
         | NEARLINE                     | Nearline storage class
         | COLDLINE                     | Coldline storage class
         | ARCHIVE                      | Archive storage class
         | DURABLE_REDUCED_AVAILABILITY | Durable reduced availability storage class

   --no-check-bucket
      If set, don't attempt to check the bucket exists or create it.
      
      This can be useful when trying to minimise the number of transactions
      rclone does if you know the bucket exists already.
      

   --decompress
      If set this will decompress gzip encoded objects.
      
      It is possible to upload objects to GCS with "Content-Encoding: gzip"
      set. Normally rclone will download these files as compressed objects.
      
      If this flag is set then rclone will decompress these files with
      "Content-Encoding: gzip" as they are received. This means that rclone
      can't check the size and hash but the file contents will be decompressed.
      

   --endpoint
      Endpoint for the service.
      
      Leave blank normally.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --env-auth
      Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
      
      Only applies if service_account_file and service_account_credentials is blank.

      Examples:
         | false | Enter credentials in the next step.
         | true  | Get GCP IAM credentials from the environment (env vars or IAM).


OPTIONS:
   --anonymous                          Access public buckets and objects without credentials. (default: false) [$ANONYMOUS]
   --bucket-acl value                   Access Control List for new buckets. [$BUCKET_ACL]
   --bucket-policy-only                 Access checks should use bucket-level IAM policies. (default: false) [$BUCKET_POLICY_ONLY]
   --client-id value                    OAuth Client Id. [$CLIENT_ID]
   --client-secret value                OAuth Client Secret. [$CLIENT_SECRET]
   --env-auth                           Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars). (default: false) [$ENV_AUTH]
   --help, -h                           show help
   --location value                     Location for the newly created buckets. [$LOCATION]
   --object-acl value                   Access Control List for new objects. [$OBJECT_ACL]
   --project-number value               Project number. [$PROJECT_NUMBER]
   --service-account-credentials value  Service Account Credentials JSON blob. [$SERVICE_ACCOUNT_CREDENTIALS]
   --service-account-file value         Service Account Credentials JSON file path. [$SERVICE_ACCOUNT_FILE]
   --storage-class value                The storage class to use when storing objects in Google Cloud Storage. [$STORAGE_CLASS]

   Advanced

   --auth-url value   Auth server URL. [$AUTH_URL]
   --decompress       If set this will decompress gzip encoded objects. (default: false) [$DECOMPRESS]
   --encoding value   The encoding for the backend. (default: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value   Endpoint for the service. [$ENDPOINT]
   --no-check-bucket  If set, don't attempt to check the bucket exists or create it. (default: false) [$NO_CHECK_BUCKET]
   --token value      OAuth Access Token as a JSON blob. [$TOKEN]
   --token-url value  Token server url. [$TOKEN_URL]

```
{% endcode %}
