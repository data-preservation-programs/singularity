# Google Cloud Storage (this is not Google Drive)

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add gcs - Google Cloud Storage (this is not Google Drive)

USAGE:
   singularity datasource add gcs [command options] <dataset_name> <source_path>

DESCRIPTION:
   --gcs-anonymous
      Access public buckets and objects without credentials.
      
      Set to 'true' if you just want to download files and don't configure credentials.

   --gcs-auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --gcs-bucket-acl
      Access Control List for new buckets.

      Examples:
         | authenticatedRead | Project team owners get OWNER access.
                             | All Authenticated Users get READER access.
         | private           | Project team owners get OWNER access.
                             | Default if left blank.
         | projectPrivate    | Project team members get access according to their roles.
         | publicRead        | Project team owners get OWNER access.
                             | All Users get READER access.
         | publicReadWrite   | Project team owners get OWNER access.
                             | All Users get WRITER access.

   --gcs-bucket-policy-only
      Access checks should use bucket-level IAM policies.
      
      If you want to upload objects to a bucket with Bucket Policy Only set
      then you will need to set this.
      
      When it is set, rclone:
      
      - ignores ACLs set on buckets
      - ignores ACLs set on objects
      - creates buckets with Bucket Policy Only set
      
      Docs: https://cloud.google.com/storage/docs/bucket-policy-only
      

   --gcs-client-id
      OAuth Client Id.
      
      Leave blank normally.

   --gcs-client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --gcs-decompress
      If set this will decompress gzip encoded objects.
      
      It is possible to upload objects to GCS with "Content-Encoding: gzip"
      set. Normally rclone will download these files as compressed objects.
      
      If this flag is set then rclone will decompress these files with
      "Content-Encoding: gzip" as they are received. This means that rclone
      can't check the size and hash but the file contents will be decompressed.
      

   --gcs-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --gcs-endpoint
      Endpoint for the service.
      
      Leave blank normally.

   --gcs-env-auth
      Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
      
      Only applies if service_account_file and service_account_credentials is blank.

      Examples:
         | false | Enter credentials in the next step.
         | true  | Get GCP IAM credentials from the environment (env vars or IAM).

   --gcs-location
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

   --gcs-no-check-bucket
      If set, don't attempt to check the bucket exists or create it.
      
      This can be useful when trying to minimise the number of transactions
      rclone does if you know the bucket exists already.
      

   --gcs-object-acl
      Access Control List for new objects.

      Examples:
         | authenticatedRead      | Object owner gets OWNER access.
                                  | All Authenticated Users get READER access.
         | bucketOwnerFullControl | Object owner gets OWNER access.
                                  | Project team owners get OWNER access.
         | bucketOwnerRead        | Object owner gets OWNER access.
                                  | Project team owners get READER access.
         | private                | Object owner gets OWNER access.
                                  | Default if left blank.
         | projectPrivate         | Object owner gets OWNER access.
                                  | Project team members get access according to their roles.
         | publicRead             | Object owner gets OWNER access.
                                  | All Users get READER access.

   --gcs-project-number
      Project number.
      
      Optional - needed only for list/create/delete buckets - see your developer console.

   --gcs-service-account-credentials
      Service Account Credentials JSON blob.
      
      Leave blank normally.
      Needed only if you want use SA instead of interactive login.

   --gcs-service-account-file
      Service Account Credentials JSON file path.
      
      Leave blank normally.
      Needed only if you want use SA instead of interactive login.
      
      Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.

   --gcs-storage-class
      The storage class to use when storing objects in Google Cloud Storage.

      Examples:
         | <unset>                      | Default
         | MULTI_REGIONAL               | Multi-regional storage class
         | REGIONAL                     | Regional storage class
         | NEARLINE                     | Nearline storage class
         | COLDLINE                     | Coldline storage class
         | ARCHIVE                      | Archive storage class
         | DURABLE_REDUCED_AVAILABILITY | Durable reduced availability storage class

   --gcs-token
      OAuth Access Token as a JSON blob.

   --gcs-token-url
      Token server url.
      
      Leave blank to use the provider defaults.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for gcs

   --gcs-anonymous value             Access public buckets and objects without credentials. (default: "false") [$GCS_ANONYMOUS]
   --gcs-auth-url value              Auth server URL. [$GCS_AUTH_URL]
   --gcs-bucket-acl value            Access Control List for new buckets. [$GCS_BUCKET_ACL]
   --gcs-bucket-policy-only value    Access checks should use bucket-level IAM policies. (default: "false") [$GCS_BUCKET_POLICY_ONLY]
   --gcs-client-id value             OAuth Client Id. [$GCS_CLIENT_ID]
   --gcs-client-secret value         OAuth Client Secret. [$GCS_CLIENT_SECRET]
   --gcs-decompress value            If set this will decompress gzip encoded objects. (default: "false") [$GCS_DECOMPRESS]
   --gcs-encoding value              The encoding for the backend. (default: "Slash,CrLf,InvalidUtf8,Dot") [$GCS_ENCODING]
   --gcs-endpoint value              Endpoint for the service. [$GCS_ENDPOINT]
   --gcs-env-auth value              Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars). (default: "false") [$GCS_ENV_AUTH]
   --gcs-location value              Location for the newly created buckets. [$GCS_LOCATION]
   --gcs-no-check-bucket value       If set, don't attempt to check the bucket exists or create it. (default: "false") [$GCS_NO_CHECK_BUCKET]
   --gcs-object-acl value            Access Control List for new objects. [$GCS_OBJECT_ACL]
   --gcs-project-number value        Project number. [$GCS_PROJECT_NUMBER]
   --gcs-service-account-file value  Service Account Credentials JSON file path. [$GCS_SERVICE_ACCOUNT_FILE]
   --gcs-storage-class value         The storage class to use when storing objects in Google Cloud Storage. [$GCS_STORAGE_CLASS]
   --gcs-token value                 OAuth Access Token as a JSON blob. [$GCS_TOKEN]
   --gcs-token-url value             Token server url. [$GCS_TOKEN_URL]

```
{% endcode %}
