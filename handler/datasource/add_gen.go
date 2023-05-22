
// Code generated. DO NOT EDIT.
package datasource

type AcdRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    AuthUrl string `json:"authUrl"` // Auth server URL.
    TokenUrl string `json:"tokenUrl"` // Token server url.
    Checkpoint string `json:"checkpoint"` // Checkpoint for internal polling (debug).
    UploadWaitPerGb string `json:"uploadWaitPerGb" default:"3m0s"` // Additional time per GiB to wait after a failed complete upload to see if it appears.
    TemplinkThreshold string `json:"templinkThreshold" default:"9Gi"` // Files >= this size will be downloaded via their tempLink.
    Encoding string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
    ClientId string `json:"clientId"` // OAuth Client Id.
    ClientSecret string `json:"clientSecret"` // OAuth Client Secret.
    Token string `json:"token"` // OAuth Access Token as a JSON blob.
}

// HandleAcd godoc
// @Summary Add acd source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body AcdRequest true "Request body"
// @Router /dataset/{datasetName}/source/acd [post]
func HandleAcd() {}


type AzureblobRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Account string `json:"account"` // Azure Storage Account Name.
    ClientCertificatePassword string `json:"clientCertificatePassword"` // Password for the certificate file (optional).
    MsiObjectId string `json:"msiObjectId"` // Object ID of the user-assigned MSI to use, if any.
    ChunkSize string `json:"chunkSize" default:"4Mi"` // Upload chunk size.
    UploadConcurrency string `json:"uploadConcurrency" default:"16"` // Concurrency for multipart uploads.
    ClientSendCertificateChain string `json:"clientSendCertificateChain" default:"false"` // Send the certificate chain when using certificate auth.
    Password string `json:"password"` // The user's password
    Endpoint string `json:"endpoint"` // Endpoint for the service.
    UploadCutoff string `json:"uploadCutoff"` // Cutoff for switching to chunked upload (<= 256 MiB) (deprecated).
    NoCheckContainer string `json:"noCheckContainer" default:"false"` // If set, don't attempt to check the container exists or create it.
    Key string `json:"key"` // Storage Account Shared Key.
    ClientId string `json:"clientId"` // The ID of the client in use.
    UseMsi string `json:"useMsi" default:"false"` // Use a managed service identity to authenticate (only works in Azure).
    MsiClientId string `json:"msiClientId"` // Object ID of the user-assigned MSI to use, if any.
    ClientCertificatePath string `json:"clientCertificatePath"` // Path to a PEM or PKCS12 certificate file including the private key.
    ArchiveTierDelete string `json:"archiveTierDelete" default:"false"` // Delete archive tier blobs before overwriting.
    MemoryPoolUseMmap string `json:"memoryPoolUseMmap" default:"false"` // Whether to use mmap buffers in internal memory pool.
    PublicAccess string `json:"publicAccess"` // Public access level of a container: blob or container.
    SasUrl string `json:"sasUrl"` // SAS URL for container level access only.
    DisableChecksum string `json:"disableChecksum" default:"false"` // Don't store MD5 checksum with object metadata.
    Encoding string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"` // The encoding for the backend.
    NoHeadObject string `json:"noHeadObject" default:"false"` // If set, do not do HEAD before GET when getting objects.
    MsiMiResId string `json:"msiMiResId"` // Azure resource ID of the user-assigned MSI to use, if any.
    UseEmulator string `json:"useEmulator" default:"false"` // Uses local storage emulator if provided as 'true'.
    EnvAuth string `json:"envAuth" default:"false"` // Read credentials from runtime (environment variables, CLI or MSI).
    Tenant string `json:"tenant"` // ID of the service principal's tenant. Also called its directory ID.
    ServicePrincipalFile string `json:"servicePrincipalFile"` // Path to file containing credentials for use with a service principal.
    ListChunk string `json:"listChunk" default:"5000"` // Size of blob list.
    MemoryPoolFlushTime string `json:"memoryPoolFlushTime" default:"1m0s"` // How often internal memory buffer pools will be flushed.
    ClientSecret string `json:"clientSecret"` // One of the service principal's client secrets
    Username string `json:"username"` // User name (usually an email address)
    AccessTier string `json:"accessTier"` // Access tier of blob: hot, cool or archive.
}

// HandleAzureblob godoc
// @Summary Add azureblob source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body AzureblobRequest true "Request body"
// @Router /dataset/{datasetName}/source/azureblob [post]
func HandleAzureblob() {}


type B2Request struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    DisableChecksum string `json:"disableChecksum" default:"false"` // Disable checksums for large (> upload cutoff) files.
    MemoryPoolUseMmap string `json:"memoryPoolUseMmap" default:"false"` // Whether to use mmap buffers in internal memory pool.
    Key string `json:"key"` // Application Key.
    Endpoint string `json:"endpoint"` // Endpoint for the service.
    TestMode string `json:"testMode"` // A flag string for X-Bz-Test-Mode header for debugging.
    Encoding string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    Account string `json:"account"` // Account ID or Application Key ID.
    ChunkSize string `json:"chunkSize" default:"96Mi"` // Upload chunk size.
    HardDelete string `json:"hardDelete" default:"false"` // Permanently delete files on remote removal, otherwise hide files.
    UploadCutoff string `json:"uploadCutoff" default:"200Mi"` // Cutoff for switching to chunked upload.
    CopyCutoff string `json:"copyCutoff" default:"4Gi"` // Cutoff for switching to multipart copy.
    DownloadUrl string `json:"downloadUrl"` // Custom endpoint for downloads.
    DownloadAuthDuration string `json:"downloadAuthDuration" default:"1w"` // Time before the authorization token will expire in s or suffix ms|s|m|h|d.
    MemoryPoolFlushTime string `json:"memoryPoolFlushTime" default:"1m0s"` // How often internal memory buffer pools will be flushed.
    Versions string `json:"versions" default:"false"` // Include old versions in directory listings.
    VersionAt string `json:"versionAt" default:"off"` // Show file versions as they were at the specified time.
}

// HandleB2 godoc
// @Summary Add b2 source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body B2Request true "Request body"
// @Router /dataset/{datasetName}/source/b2 [post]
func HandleB2() {}


type BoxRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    ClientSecret string `json:"clientSecret"` // OAuth Client Secret.
    ListChunk string `json:"listChunk" default:"1000"` // Size of listing chunk 1-1000.
    TokenUrl string `json:"tokenUrl"` // Token server url.
    BoxSubType string `json:"boxSubType" default:"user"` // 
    UploadCutoff string `json:"uploadCutoff" default:"50Mi"` // Cutoff for switching to multipart upload (>= 50 MiB).
    OwnedBy string `json:"ownedBy"` // Only show items owned by the login (email address) passed in.
    Encoding string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
    AuthUrl string `json:"authUrl"` // Auth server URL.
    BoxConfigFile string `json:"boxConfigFile"` // Box App config.json location
    CommitRetries string `json:"commitRetries" default:"100"` // Max number of times to try committing a multipart file.
    ClientId string `json:"clientId"` // OAuth Client Id.
    Token string `json:"token"` // OAuth Access Token as a JSON blob.
    RootFolderId string `json:"rootFolderId" default:"0"` // Fill in for rclone to use a non root folder as its starting point.
    AccessToken string `json:"accessToken"` // Box App Primary Access Token
}

// HandleBox godoc
// @Summary Add box source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body BoxRequest true "Request body"
// @Router /dataset/{datasetName}/source/box [post]
func HandleBox() {}


type CryptRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Remote string `json:"remote"` // Remote to encrypt/decrypt.
    FilenameEncryption string `json:"filenameEncryption" default:"standard"` // How to encrypt the filenames.
    Password string `json:"password"` // Password or pass phrase for encryption.
    Password2 string `json:"password2"` // Password or pass phrase for salt.
    ShowMapping string `json:"showMapping" default:"false"` // For all files listed show how the names encrypt.
    FilenameEncoding string `json:"filenameEncoding" default:"base32"` // How to encode the encrypted filename to text string.
    DirectoryNameEncryption string `json:"directoryNameEncryption" default:"true"` // Option to either encrypt directory names or leave them intact.
    ServerSideAcrossConfigs string `json:"serverSideAcrossConfigs" default:"false"` // Allow server-side operations (e.g. copy) to work across different crypt configs.
    NoDataEncryption string `json:"noDataEncryption" default:"false"` // Option to either encrypt file data or leave it unencrypted.
}

// HandleCrypt godoc
// @Summary Add crypt source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body CryptRequest true "Request body"
// @Router /dataset/{datasetName}/source/crypt [post]
func HandleCrypt() {}


type DriveRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Scope string `json:"scope"` // Scope that rclone should use when requesting access from drive.
    ServiceAccountCredentials string `json:"serviceAccountCredentials"` // Service Account Credentials JSON blob.
    AuthOwnerOnly string `json:"authOwnerOnly" default:"false"` // Only consider files owned by the authenticated user.
    StarredOnly string `json:"starredOnly" default:"false"` // Only show files that are starred.
    SkipDanglingShortcuts string `json:"skipDanglingShortcuts" default:"false"` // If set skip dangling shortcut files.
    ClientSecret string `json:"clientSecret"` // OAuth Client Secret.
    CopyShortcutContent string `json:"copyShortcutContent" default:"false"` // Server side copy contents of shortcuts instead of the shortcut.
    Impersonate string `json:"impersonate"` // Impersonate this user when using a service account.
    StopOnUploadLimit string `json:"stopOnUploadLimit" default:"false"` // Make upload limit errors be fatal.
    ChunkSize string `json:"chunkSize" default:"8Mi"` // Upload chunk size.
    DisableHttp2 string `json:"disableHttp2" default:"true"` // Disable drive using http2.
    TokenUrl string `json:"tokenUrl"` // Token server url.
    ServiceAccountFile string `json:"serviceAccountFile"` // Service Account Credentials JSON file path.
    AllowImportNameChange string `json:"allowImportNameChange" default:"false"` // Allow the filetype to change when uploading Google docs.
    ListChunk string `json:"listChunk" default:"1000"` // Size of listing chunk 100-1000, 0 to disable.
    AlternateExport string `json:"alternateExport" default:"false"` // Deprecated: No longer needed.
    UploadCutoff string `json:"uploadCutoff" default:"8Mi"` // Cutoff for switching to chunked upload.
    ResourceKey string `json:"resourceKey"` // Resource key for accessing a link-shared file.
    Token string `json:"token"` // OAuth Access Token as a JSON blob.
    RootFolderId string `json:"rootFolderId"` // ID of the root folder.
    SharedWithMe string `json:"sharedWithMe" default:"false"` // Only show files that are shared with me.
    AcknowledgeAbuse string `json:"acknowledgeAbuse" default:"false"` // Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
    SizeAsQuota string `json:"sizeAsQuota" default:"false"` // Show sizes as storage quota usage, not actual size.
    SkipShortcuts string `json:"skipShortcuts" default:"false"` // If set skip shortcut files.
    ClientId string `json:"clientId"` // Google Application Client Id
    UseTrash string `json:"useTrash" default:"true"` // Send files to the trash instead of deleting permanently.
    UseSharedDate string `json:"useSharedDate" default:"false"` // Use date file was shared instead of modified date.
    AuthUrl string `json:"authUrl"` // Auth server URL.
    TrashedOnly string `json:"trashedOnly" default:"false"` // Only show files that are in the trash.
    KeepRevisionForever string `json:"keepRevisionForever" default:"false"` // Keep new head revision of each file forever.
    ServerSideAcrossConfigs string `json:"serverSideAcrossConfigs" default:"false"` // Allow server-side operations (e.g. copy) to work across different drive configs.
    ExportFormats string `json:"exportFormats" default:"docx,xlsx,pptx,svg"` // Comma separated list of preferred formats for downloading Google docs.
    ImportFormats string `json:"importFormats"` // Comma separated list of preferred formats for uploading Google docs.
    V2DownloadMinSize string `json:"v2DownloadMinSize" default:"off"` // If Object's are greater, use drive v2 API to download.
    PacerBurst string `json:"pacerBurst" default:"100"` // Number of API calls to allow without sleeping.
    Encoding string `json:"encoding" default:"InvalidUtf8"` // The encoding for the backend.
    StopOnDownloadLimit string `json:"stopOnDownloadLimit" default:"false"` // Make download limit errors be fatal.
    TeamDrive string `json:"teamDrive"` // ID of the Shared Drive (Team Drive).
    SkipGdocs string `json:"skipGdocs" default:"false"` // Skip google documents in all listings.
    SkipChecksumGphotos string `json:"skipChecksumGphotos" default:"false"` // Skip MD5 checksum on Google photos and videos only.
    Formats string `json:"formats"` // Deprecated: See export_formats.
    UseCreatedDate string `json:"useCreatedDate" default:"false"` // Use file created date instead of modified date.
    PacerMinSleep string `json:"pacerMinSleep" default:"100ms"` // Minimum time to sleep between API calls.
}

// HandleDrive godoc
// @Summary Add drive source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body DriveRequest true "Request body"
// @Router /dataset/{datasetName}/source/drive [post]
func HandleDrive() {}


type DropboxRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    SharedFiles string `json:"sharedFiles" default:"false"` // Instructs rclone to work on individual shared files.
    SharedFolders string `json:"sharedFolders" default:"false"` // Instructs rclone to work on shared folders.
    BatchMode string `json:"batchMode" default:"sync"` // Upload file batching sync|async|off.
    Encoding string `json:"encoding" default:"Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
    ClientId string `json:"clientId"` // OAuth Client Id.
    Token string `json:"token"` // OAuth Access Token as a JSON blob.
    AuthUrl string `json:"authUrl"` // Auth server URL.
    TokenUrl string `json:"tokenUrl"` // Token server url.
    ClientSecret string `json:"clientSecret"` // OAuth Client Secret.
    Impersonate string `json:"impersonate"` // Impersonate this user when using a business account.
    BatchTimeout string `json:"batchTimeout" default:"0s"` // Max time to allow an idle upload batch before uploading.
    BatchCommitTimeout string `json:"batchCommitTimeout" default:"10m0s"` // Max time to wait for a batch to finish committing
    ChunkSize string `json:"chunkSize" default:"48Mi"` // Upload chunk size (< 150Mi).
    BatchSize string `json:"batchSize" default:"0"` // Max number of files in upload batch.
}

// HandleDropbox godoc
// @Summary Add dropbox source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body DropboxRequest true "Request body"
// @Router /dataset/{datasetName}/source/dropbox [post]
func HandleDropbox() {}


type FichierRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    ApiKey string `json:"apiKey"` // Your API Key, get it from https://1fichier.com/console/params.pl.
    SharedFolder string `json:"sharedFolder"` // If you want to download a shared folder, add this parameter.
    FilePassword string `json:"filePassword"` // If you want to download a shared file that is password protected, add this parameter.
    FolderPassword string `json:"folderPassword"` // If you want to list the files in a shared folder that is password protected, add this parameter.
    Encoding string `json:"encoding" default:"Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
}

// HandleFichier godoc
// @Summary Add fichier source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body FichierRequest true "Request body"
// @Router /dataset/{datasetName}/source/fichier [post]
func HandleFichier() {}


type FilefabricRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    PermanentToken string `json:"permanentToken"` // Permanent Authentication Token.
    Token string `json:"token"` // Session Token.
    TokenExpiry string `json:"tokenExpiry"` // Token expiry time.
    Version string `json:"version"` // Version read from the file fabric.
    Encoding string `json:"encoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    Url string `json:"url"` // URL of the Enterprise File Fabric to connect to.
    RootFolderId string `json:"rootFolderId"` // ID of the root folder.
}

// HandleFilefabric godoc
// @Summary Add filefabric source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body FilefabricRequest true "Request body"
// @Router /dataset/{datasetName}/source/filefabric [post]
func HandleFilefabric() {}


type FtpRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Pass string `json:"pass"` // FTP password.
    Tls string `json:"tls" default:"false"` // Use Implicit FTPS (FTP over TLS).
    ExplicitTls string `json:"explicitTls" default:"false"` // Use Explicit FTPS (FTP over TLS).
    DisableEpsv string `json:"disableEpsv" default:"false"` // Disable using EPSV even if server advertises support.
    IdleTimeout string `json:"idleTimeout" default:"1m0s"` // Max time before closing idle connections.
    CloseTimeout string `json:"closeTimeout" default:"1m0s"` // Maximum time to wait for a response to close.
    ShutTimeout string `json:"shutTimeout" default:"1m0s"` // Maximum time to wait for data connection closing status.
    User string `json:"user" default:"shane"` // FTP username.
    Encoding string `json:"encoding" default:"Slash,Del,Ctl,RightSpace,Dot"` // The encoding for the backend.
    DisableUtf8 string `json:"disableUtf8" default:"false"` // Disable using UTF-8 even if server advertises support.
    ForceListHidden string `json:"forceListHidden" default:"false"` // Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.
    AskPassword string `json:"askPassword" default:"false"` // Allow asking for FTP password when needed.
    NoCheckCertificate string `json:"noCheckCertificate" default:"false"` // Do not verify the TLS certificate of the server.
    Port string `json:"port" default:"21"` // FTP port number.
    TlsCacheSize string `json:"tlsCacheSize" default:"32"` // Size of TLS session cache for all control and data connections.
    Host string `json:"host"` // FTP host to connect to.
    DisableMlsd string `json:"disableMlsd" default:"false"` // Disable using MLSD even if server advertises support.
    WritingMdtm string `json:"writingMdtm" default:"false"` // Use MDTM to set modification time (VsFtpd quirk)
    DisableTls13 string `json:"disableTls13" default:"false"` // Disable TLS 1.3 (workaround for FTP servers with buggy TLS)
    Concurrency string `json:"concurrency" default:"0"` // Maximum number of FTP simultaneous connections, 0 for unlimited.
}

// HandleFtp godoc
// @Summary Add ftp source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body FtpRequest true "Request body"
// @Router /dataset/{datasetName}/source/ftp [post]
func HandleFtp() {}


type GcsRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    BucketAcl string `json:"bucketAcl"` // Access Control List for new buckets.
    BucketPolicyOnly string `json:"bucketPolicyOnly" default:"false"` // Access checks should use bucket-level IAM policies.
    Endpoint string `json:"endpoint"` // Endpoint for the service.
    EnvAuth string `json:"envAuth" default:"false"` // Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
    ClientId string `json:"clientId"` // OAuth Client Id.
    ServiceAccountFile string `json:"serviceAccountFile"` // Service Account Credentials JSON file path.
    ObjectAcl string `json:"objectAcl"` // Access Control List for new objects.
    Location string `json:"location"` // Location for the newly created buckets.
    AuthUrl string `json:"authUrl"` // Auth server URL.
    TokenUrl string `json:"tokenUrl"` // Token server url.
    ProjectNumber string `json:"projectNumber"` // Project number.
    Anonymous string `json:"anonymous" default:"false"` // Access public buckets and objects without credentials.
    Decompress string `json:"decompress" default:"false"` // If set this will decompress gzip encoded objects.
    Encoding string `json:"encoding" default:"Slash,CrLf,InvalidUtf8,Dot"` // The encoding for the backend.
    ClientSecret string `json:"clientSecret"` // OAuth Client Secret.
    Token string `json:"token"` // OAuth Access Token as a JSON blob.
    ServiceAccountCredentials string `json:"serviceAccountCredentials"` // Service Account Credentials JSON blob.
    StorageClass string `json:"storageClass"` // The storage class to use when storing objects in Google Cloud Storage.
    NoCheckBucket string `json:"noCheckBucket" default:"false"` // If set, don't attempt to check the bucket exists or create it.
}

// HandleGcs godoc
// @Summary Add gcs source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body GcsRequest true "Request body"
// @Router /dataset/{datasetName}/source/gcs [post]
func HandleGcs() {}


type GphotosRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    IncludeArchived string `json:"includeArchived" default:"false"` // Also view and download archived media.
    Encoding string `json:"encoding" default:"Slash,CrLf,InvalidUtf8,Dot"` // The encoding for the backend.
    ClientSecret string `json:"clientSecret"` // OAuth Client Secret.
    AuthUrl string `json:"authUrl"` // Auth server URL.
    ReadOnly string `json:"readOnly" default:"false"` // Set to make the Google Photos backend read only.
    ReadSize string `json:"readSize" default:"false"` // Set to read the size of media items.
    StartYear string `json:"startYear" default:"2000"` // Year limits the photos to be downloaded to those which are uploaded after the given year.
    ClientId string `json:"clientId"` // OAuth Client Id.
    Token string `json:"token"` // OAuth Access Token as a JSON blob.
    TokenUrl string `json:"tokenUrl"` // Token server url.
}

// HandleGphotos godoc
// @Summary Add gphotos source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body GphotosRequest true "Request body"
// @Router /dataset/{datasetName}/source/gphotos [post]
func HandleGphotos() {}


type HdfsRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    DataTransferProtection string `json:"dataTransferProtection"` // Kerberos data transfer protection: authentication|integrity|privacy.
    Encoding string `json:"encoding" default:"Slash,Colon,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    Namenode string `json:"namenode"` // Hadoop name node and port.
    Username string `json:"username"` // Hadoop user name.
    ServicePrincipalName string `json:"servicePrincipalName"` // Kerberos service principal name for the namenode.
}

// HandleHdfs godoc
// @Summary Add hdfs source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body HdfsRequest true "Request body"
// @Router /dataset/{datasetName}/source/hdfs [post]
func HandleHdfs() {}


type HidriveRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    RootPrefix string `json:"rootPrefix" default:"/"` // The root/parent folder for all paths.
    UploadConcurrency string `json:"uploadConcurrency" default:"4"` // Concurrency for chunked uploads.
    Token string `json:"token"` // OAuth Access Token as a JSON blob.
    AuthUrl string `json:"authUrl"` // Auth server URL.
    ScopeRole string `json:"scopeRole" default:"user"` // User-level that rclone should use when requesting access from HiDrive.
    ChunkSize string `json:"chunkSize" default:"48Mi"` // Chunksize for chunked uploads.
    UploadCutoff string `json:"uploadCutoff" default:"96Mi"` // Cutoff/Threshold for chunked uploads.
    Encoding string `json:"encoding" default:"Slash,Dot"` // The encoding for the backend.
    ClientId string `json:"clientId"` // OAuth Client Id.
    ClientSecret string `json:"clientSecret"` // OAuth Client Secret.
    DisableFetchingMemberCount string `json:"disableFetchingMemberCount" default:"false"` // Do not fetch number of objects in directories unless it is absolutely necessary.
    TokenUrl string `json:"tokenUrl"` // Token server url.
    Endpoint string `json:"endpoint" default:"https://api.hidrive.strato.com/2.1"` // Endpoint for the service.
    ScopeAccess string `json:"scopeAccess" default:"rw"` // Access permissions that rclone should use when requesting access from HiDrive.
}

// HandleHidrive godoc
// @Summary Add hidrive source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body HidriveRequest true "Request body"
// @Router /dataset/{datasetName}/source/hidrive [post]
func HandleHidrive() {}


type HttpRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Url string `json:"url"` // URL of HTTP host to connect to.
    Headers string `json:"headers"` // Set HTTP headers for all transactions.
    NoSlash string `json:"noSlash" default:"false"` // Set this if the site doesn't end directories with /.
    NoHead string `json:"noHead" default:"false"` // Don't use HEAD requests.
}

// HandleHttp godoc
// @Summary Add http source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body HttpRequest true "Request body"
// @Router /dataset/{datasetName}/source/http [post]
func HandleHttp() {}


type InternetarchiveRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Encoding string `json:"encoding" default:"Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    AccessKeyId string `json:"accessKeyId"` // IAS3 Access Key.
    SecretAccessKey string `json:"secretAccessKey"` // IAS3 Secret Key (password).
    Endpoint string `json:"endpoint" default:"https://s3.us.archive.org"` // IAS3 Endpoint.
    FrontEndpoint string `json:"frontEndpoint" default:"https://archive.org"` // Host of InternetArchive Frontend.
    DisableChecksum string `json:"disableChecksum" default:"true"` // Don't ask the server to test against MD5 checksum calculated by rclone.
    WaitArchive string `json:"waitArchive" default:"0s"` // Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
}

// HandleInternetarchive godoc
// @Summary Add internetarchive source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body InternetarchiveRequest true "Request body"
// @Router /dataset/{datasetName}/source/internetarchive [post]
func HandleInternetarchive() {}


type JottacloudRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Encoding string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    Md5MemoryLimit string `json:"md5MemoryLimit" default:"10Mi"` // Files bigger than this will be cached on disk to calculate the MD5 if required.
    TrashedOnly string `json:"trashedOnly" default:"false"` // Only show files that are in the trash.
    HardDelete string `json:"hardDelete" default:"false"` // Delete files permanently rather than putting them into the trash.
    UploadResumeLimit string `json:"uploadResumeLimit" default:"10Mi"` // Files bigger than this can be resumed if the upload fail's.
    NoVersions string `json:"noVersions" default:"false"` // Avoid server side versioning by deleting files and recreating files instead of overwriting them.
}

// HandleJottacloud godoc
// @Summary Add jottacloud source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body JottacloudRequest true "Request body"
// @Router /dataset/{datasetName}/source/jottacloud [post]
func HandleJottacloud() {}


type KoofrRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Endpoint string `json:"endpoint"` // The Koofr API endpoint to use.
    Mountid string `json:"mountid"` // Mount ID of the mount to use.
    Setmtime string `json:"setmtime" default:"true"` // Does the backend support setting modification time.
    User string `json:"user"` // Your user name.
    Password string `json:"password"` // Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password).
    Encoding string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    Provider string `json:"provider"` // Choose your storage provider.
}

// HandleKoofr godoc
// @Summary Add koofr source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body KoofrRequest true "Request body"
// @Router /dataset/{datasetName}/source/koofr [post]
func HandleKoofr() {}


type LocalRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    ZeroSizeLinks string `json:"zeroSizeLinks" default:"false"` // Assume the Stat size of links is zero (and read them instead) (deprecated).
    NoSparse string `json:"noSparse" default:"false"` // Disable sparse files for multi-thread downloads.
    NoSetModtime string `json:"noSetModtime" default:"false"` // Disable setting modtime.
    CopyLinks string `json:"copyLinks" default:"false"` // Follow symlinks and copy the pointed to item.
    SkipLinks string `json:"skipLinks" default:"false"` // Don't warn about skipped symlinks.
    NoCheckUpdated string `json:"noCheckUpdated" default:"false"` // Don't check to see if the files change during upload.
    OneFileSystem string `json:"oneFileSystem" default:"false"` // Don't cross filesystem boundaries (unix/macOS only).
    CaseSensitive string `json:"caseSensitive" default:"false"` // Force the filesystem to report itself as case sensitive.
    NoPreallocate string `json:"noPreallocate" default:"false"` // Disable preallocation of disk space for transferred files.
    Nounc string `json:"nounc" default:"false"` // Disable UNC (long path names) conversion on Windows.
    Links string `json:"links" default:"false"` // Translate symlinks to/from regular files with a '.rclonelink' extension.
    UnicodeNormalization string `json:"unicodeNormalization" default:"false"` // Apply unicode NFC normalization to paths and filenames.
    CaseInsensitive string `json:"caseInsensitive" default:"false"` // Force the filesystem to report itself as case insensitive.
    Encoding string `json:"encoding" default:"Slash,Dot"` // The encoding for the backend.
}

// HandleLocal godoc
// @Summary Add local source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body LocalRequest true "Request body"
// @Router /dataset/{datasetName}/source/local [post]
func HandleLocal() {}


type MailruRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    SpeedupMaxMemory string `json:"speedupMaxMemory" default:"32Mi"` // Files larger than the size given below will always be hashed on disk.
    Encoding string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    Pass string `json:"pass"` // Password.
    SpeedupMaxDisk string `json:"speedupMaxDisk" default:"3Gi"` // This option allows you to disable speedup (put by hash) for large files.
    SpeedupFilePatterns string `json:"speedupFilePatterns" default:"*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf"` // Comma separated list of file name patterns eligible for speedup (put by hash).
    CheckHash string `json:"checkHash" default:"true"` // What should copy do if file checksum is mismatched or invalid.
    UserAgent string `json:"userAgent"` // HTTP user agent used internally by client.
    Quirks string `json:"quirks"` // Comma separated list of internal maintenance flags.
    User string `json:"user"` // User name (usually email).
    SpeedupEnable string `json:"speedupEnable" default:"true"` // Skip full upload if there is another file with same data hash.
}

// HandleMailru godoc
// @Summary Add mailru source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body MailruRequest true "Request body"
// @Router /dataset/{datasetName}/source/mailru [post]
func HandleMailru() {}


type MegaRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    User string `json:"user"` // User name.
    Pass string `json:"pass"` // Password.
    Debug string `json:"debug" default:"false"` // Output more debug from Mega.
    HardDelete string `json:"hardDelete" default:"false"` // Delete files permanently rather than putting them into the trash.
    UseHttps string `json:"useHttps" default:"false"` // Use HTTPS for transfers.
    Encoding string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
}

// HandleMega godoc
// @Summary Add mega source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body MegaRequest true "Request body"
// @Router /dataset/{datasetName}/source/mega [post]
func HandleMega() {}


type MemoryRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
}

// HandleMemory godoc
// @Summary Add memory source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body MemoryRequest true "Request body"
// @Router /dataset/{datasetName}/source/memory [post]
func HandleMemory() {}


type NetstorageRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Protocol string `json:"protocol" default:"https"` // Select between HTTP or HTTPS protocol.
    Host string `json:"host"` // Domain+path of NetStorage host to connect to.
    Account string `json:"account"` // Set the NetStorage account name
    Secret string `json:"secret"` // Set the NetStorage account secret/G2O key for authentication.
}

// HandleNetstorage godoc
// @Summary Add netstorage source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body NetstorageRequest true "Request body"
// @Router /dataset/{datasetName}/source/netstorage [post]
func HandleNetstorage() {}


type OnedriveRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    ClientSecret string `json:"clientSecret"` // OAuth Client Secret.
    Token string `json:"token"` // OAuth Access Token as a JSON blob.
    ServerSideAcrossConfigs string `json:"serverSideAcrossConfigs" default:"false"` // Allow server-side operations (e.g. copy) to work across different onedrive configs.
    Encoding string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
    TokenUrl string `json:"tokenUrl"` // Token server url.
    DriveType string `json:"driveType"` // The type of the drive (personal | business | documentLibrary).
    DisableSitePermission string `json:"disableSitePermission" default:"false"` // Disable the request for Sites.Read.All permission.
    LinkPassword string `json:"linkPassword"` // Set the password for links created by the link command.
    LinkType string `json:"linkType" default:"view"` // Set the type of the links created by the link command.
    AuthUrl string `json:"authUrl"` // Auth server URL.
    ChunkSize string `json:"chunkSize" default:"10Mi"` // Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
    ExposeOnenoteFiles string `json:"exposeOnenoteFiles" default:"false"` // Set to make OneNote files show up in directory listings.
    ListChunk string `json:"listChunk" default:"1000"` // Size of listing chunk.
    AccessScopes string `json:"accessScopes" default:"Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"` // Set scopes to be requested by rclone.
    NoVersions string `json:"noVersions" default:"false"` // Remove all versions on modifying operations.
    LinkScope string `json:"linkScope" default:"anonymous"` // Set the scope of the links created by the link command.
    HashType string `json:"hashType" default:"auto"` // Specify the hash in use for the backend.
    ClientId string `json:"clientId"` // OAuth Client Id.
    Region string `json:"region" default:"global"` // Choose national cloud region for OneDrive.
    DriveId string `json:"driveId"` // The ID of the drive to use.
    RootFolderId string `json:"rootFolderId"` // ID of the root folder.
}

// HandleOnedrive godoc
// @Summary Add onedrive source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body OnedriveRequest true "Request body"
// @Router /dataset/{datasetName}/source/onedrive [post]
func HandleOnedrive() {}


type OpendriveRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Username string `json:"username"` // Username.
    Password string `json:"password"` // Password.
    Encoding string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot"` // The encoding for the backend.
    ChunkSize string `json:"chunkSize" default:"10Mi"` // Files will be uploaded in chunks this size.
}

// HandleOpendrive godoc
// @Summary Add opendrive source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body OpendriveRequest true "Request body"
// @Router /dataset/{datasetName}/source/opendrive [post]
func HandleOpendrive() {}


type OosRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Provider string `json:"provider" default:"env_auth"` // Choose your Auth Provider
    Namespace string `json:"namespace"` // Object storage namespace
    ConfigProfile string `json:"configProfile" default:"Default"` // Profile name inside the oci config file
    UploadConcurrency string `json:"uploadConcurrency" default:"10"` // Concurrency for multipart uploads.
    SseCustomerKey string `json:"sseCustomerKey"` // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
    SseKmsKeyId string `json:"sseKmsKeyId"` // if using using your own master key in vault, this header specifies the 
    Endpoint string `json:"endpoint"` // Endpoint for Object storage API.
    StorageTier string `json:"storageTier" default:"Standard"` // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
    ChunkSize string `json:"chunkSize" default:"5Mi"` // Chunk size to use for uploading.
    CopyTimeout string `json:"copyTimeout" default:"1m0s"` // Timeout for copy.
    LeavePartsOnError string `json:"leavePartsOnError" default:"false"` // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
    Encoding string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
    NoCheckBucket string `json:"noCheckBucket" default:"false"` // If set, don't attempt to check the bucket exists or create it.
    Compartment string `json:"compartment"` // Object storage compartment OCID
    Region string `json:"region"` // Object storage Region
    ConfigFile string `json:"configFile" default:"~/.oci/config"` // Path to OCI config file
    UploadCutoff string `json:"uploadCutoff" default:"200Mi"` // Cutoff for switching to chunked upload.
    CopyCutoff string `json:"copyCutoff" default:"4.656Gi"` // Cutoff for switching to multipart copy.
    DisableChecksum string `json:"disableChecksum" default:"false"` // Don't store MD5 checksum with object metadata.
    SseCustomerKeyFile string `json:"sseCustomerKeyFile"` // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
    SseCustomerKeySha256 string `json:"sseCustomerKeySha256"` // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
    SseCustomerAlgorithm string `json:"sseCustomerAlgorithm"` // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
}

// HandleOos godoc
// @Summary Add oos source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body OosRequest true "Request body"
// @Router /dataset/{datasetName}/source/oos [post]
func HandleOos() {}


type PcloudRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Token string `json:"token"` // OAuth Access Token as a JSON blob.
    TokenUrl string `json:"tokenUrl"` // Token server url.
    Hostname string `json:"hostname" default:"api.pcloud.com"` // Hostname to connect to.
    Password string `json:"password"` // Your pcloud password.
    RootFolderId string `json:"rootFolderId" default:"d0"` // Fill in for rclone to use a non root folder as its starting point.
    Username string `json:"username"` // Your pcloud username.
    ClientId string `json:"clientId"` // OAuth Client Id.
    ClientSecret string `json:"clientSecret"` // OAuth Client Secret.
    AuthUrl string `json:"authUrl"` // Auth server URL.
    Encoding string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

// HandlePcloud godoc
// @Summary Add pcloud source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body PcloudRequest true "Request body"
// @Router /dataset/{datasetName}/source/pcloud [post]
func HandlePcloud() {}


type PremiumizemeRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    ApiKey string `json:"apiKey"` // API Key.
    Encoding string `json:"encoding" default:"Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

// HandlePremiumizeme godoc
// @Summary Add premiumizeme source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body PremiumizemeRequest true "Request body"
// @Router /dataset/{datasetName}/source/premiumizeme [post]
func HandlePremiumizeme() {}


type PutioRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Encoding string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

// HandlePutio godoc
// @Summary Add putio source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body PutioRequest true "Request body"
// @Router /dataset/{datasetName}/source/putio [post]
func HandlePutio() {}


type QingstorRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    EnvAuth string `json:"envAuth" default:"false"` // Get QingStor credentials from runtime.
    AccessKeyId string `json:"accessKeyId"` // QingStor Access Key ID.
    Endpoint string `json:"endpoint"` // Enter an endpoint URL to connection QingStor API.
    Zone string `json:"zone"` // Zone to connect to.
    UploadCutoff string `json:"uploadCutoff" default:"200Mi"` // Cutoff for switching to chunked upload.
    SecretAccessKey string `json:"secretAccessKey"` // QingStor Secret Access Key (password).
    ConnectionRetries string `json:"connectionRetries" default:"3"` // Number of connection retries.
    ChunkSize string `json:"chunkSize" default:"4Mi"` // Chunk size to use for uploading.
    UploadConcurrency string `json:"uploadConcurrency" default:"1"` // Concurrency for multipart uploads.
    Encoding string `json:"encoding" default:"Slash,Ctl,InvalidUtf8"` // The encoding for the backend.
}

// HandleQingstor godoc
// @Summary Add qingstor source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body QingstorRequest true "Request body"
// @Router /dataset/{datasetName}/source/qingstor [post]
func HandleQingstor() {}


type S3Request struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    UploadCutoff string `json:"uploadCutoff" default:"200Mi"` // Cutoff for switching to chunked upload.
    CopyCutoff string `json:"copyCutoff" default:"4.656Gi"` // Cutoff for switching to multipart copy.
    ForcePathStyle string `json:"forcePathStyle" default:"true"` // If true use path style access if false use virtual hosted style.
    UseAccelerateEndpoint string `json:"useAccelerateEndpoint" default:"false"` // If true use the AWS S3 accelerated endpoint.
    ListChunk string `json:"listChunk" default:"1000"` // Size of listing chunk (response list for each ListObject S3 request).
    AccessKeyId string `json:"accessKeyId"` // AWS Access Key ID.
    Region string `json:"region"` // Region to connect to.
    SseCustomerKeyBase64 string `json:"sseCustomerKeyBase64"` // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
    StsEndpoint string `json:"stsEndpoint"` // Endpoint for STS.
    ListVersion string `json:"listVersion" default:"0"` // Version of ListObjects to use: 1,2 or 0 for auto.
    Encoding string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
    NoSystemMetadata string `json:"noSystemMetadata" default:"false"` // Suppress setting and reading of system metadata
    SseKmsKeyId string `json:"sseKmsKeyId"` // If using KMS ID you must provide the ARN of Key.
    DisableChecksum string `json:"disableChecksum" default:"false"` // Don't store MD5 checksum with object metadata.
    NoCheckBucket string `json:"noCheckBucket" default:"false"` // If set, don't attempt to check the bucket exists or create it.
    UsePresignedRequest string `json:"usePresignedRequest" default:"false"` // Whether to use a presigned request or PutObject for single part uploads
    Endpoint string `json:"endpoint"` // Endpoint for S3 API.
    LocationConstraint string `json:"locationConstraint"` // Location constraint - must be set to match the Region.
    BucketAcl string `json:"bucketAcl"` // Canned ACL used when creating buckets.
    ListUrlEncode string `json:"listUrlEncode" default:"unset"` // Whether to url encode listings: true/false/unset
    MemoryPoolUseMmap string `json:"memoryPoolUseMmap" default:"false"` // Whether to use mmap buffers in internal memory pool.
    VersionAt string `json:"versionAt" default:"off"` // Show file versions as they were at the specified time.
    ChunkSize string `json:"chunkSize" default:"5Mi"` // Chunk size to use for uploading.
    SessionToken string `json:"sessionToken"` // An AWS session token.
    UploadConcurrency string `json:"uploadConcurrency" default:"4"` // Concurrency for multipart uploads.
    Decompress string `json:"decompress" default:"false"` // If set this will decompress gzip encoded objects.
    Acl string `json:"acl"` // Canned ACL used when creating buckets and storing or copying objects.
    MaxUploadParts string `json:"maxUploadParts" default:"10000"` // Maximum number of parts in a multipart upload.
    NoHeadObject string `json:"noHeadObject" default:"false"` // If set, do not do HEAD before GET when getting objects.
    NoHead string `json:"noHead" default:"false"` // If set, don't HEAD uploaded objects to check integrity.
    DownloadUrl string `json:"downloadUrl"` // Custom endpoint for downloads.
    MightGzip string `json:"mightGzip" default:"unset"` // Set this if the backend might gzip objects.
    Provider string `json:"provider"` // Choose your S3 provider.
    SecretAccessKey string `json:"secretAccessKey"` // AWS Secret Access Key (password).
    SseCustomerAlgorithm string `json:"sseCustomerAlgorithm"` // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
    V2Auth string `json:"v2Auth" default:"false"` // If true use v2 authentication.
    LeavePartsOnError string `json:"leavePartsOnError" default:"false"` // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
    DisableHttp2 string `json:"disableHttp2" default:"false"` // Disable usage of http2 for S3 backends.
    EnvAuth string `json:"envAuth" default:"false"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
    StorageClass string `json:"storageClass"` // The storage class to use when storing new objects in S3.
    Profile string `json:"profile"` // Profile to use in the shared credentials file.
    UseMultipartEtag string `json:"useMultipartEtag" default:"unset"` // Whether to use ETag in multipart uploads for verification
    ServerSideEncryption string `json:"serverSideEncryption"` // The server-side encryption algorithm used when storing this object in S3.
    SharedCredentialsFile string `json:"sharedCredentialsFile"` // Path to the shared credentials file.
    MemoryPoolFlushTime string `json:"memoryPoolFlushTime" default:"1m0s"` // How often internal memory buffer pools will be flushed.
    Versions string `json:"versions" default:"false"` // Include old versions in directory listings.
    RequesterPays string `json:"requesterPays" default:"false"` // Enables requester pays option when interacting with S3 bucket.
    SseCustomerKey string `json:"sseCustomerKey"` // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
    SseCustomerKeyMd5 string `json:"sseCustomerKeyMd5"` // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
}

// HandleS3 godoc
// @Summary Add s3 source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body S3Request true "Request body"
// @Router /dataset/{datasetName}/source/s3 [post]
func HandleS3() {}


type SeafileRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Url string `json:"url"` // URL of seafile host to connect to.
    User string `json:"user"` // User name (usually email address).
    Library string `json:"library"` // Name of the library.
    CreateLibrary string `json:"createLibrary" default:"false"` // Should rclone create a library if it doesn't exist.
    AuthToken string `json:"authToken"` // Authentication token.
    Pass string `json:"pass"` // Password.
    TwoFA string `json:"2fa" default:"false"` // Two-factor authentication ('true' if the account has 2FA enabled).
    LibraryKey string `json:"libraryKey"` // Library password (for encrypted libraries only).
    Encoding string `json:"encoding" default:"Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8"` // The encoding for the backend.
}

// HandleSeafile godoc
// @Summary Add seafile source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body SeafileRequest true "Request body"
// @Router /dataset/{datasetName}/source/seafile [post]
func HandleSeafile() {}


type SftpRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Subsystem string `json:"subsystem" default:"sftp"` // Specifies the SSH2 subsystem on the remote host.
    KeyPem string `json:"keyPem"` // Raw PEM-encoded private key.
    KeyFilePass string `json:"keyFilePass"` // The passphrase to decrypt the PEM-encoded private key file.
    SetModtime string `json:"setModtime" default:"true"` // Set the modified time on the remote if set.
    Host string `json:"host"` // SSH host to connect to.
    DisableHashcheck string `json:"disableHashcheck" default:"false"` // Disable the execution of SSH commands to determine if remote file hashing is available.
    PathOverride string `json:"pathOverride"` // Override path used by SSH shell commands.
    Sha1sumCommand string `json:"sha1sumCommand"` // The command used to read sha1 hashes.
    Port string `json:"port" default:"22"` // SSH port number.
    KnownHostsFile string `json:"knownHostsFile"` // Optional path to known_hosts file.
    UseInsecureCipher string `json:"useInsecureCipher" default:"false"` // Enable the use of insecure ciphers and key exchange methods.
    Md5sumCommand string `json:"md5sumCommand"` // The command used to read md5 hashes.
    SkipLinks string `json:"skipLinks" default:"false"` // Set to skip any symlinks and any other non regular files.
    SetEnv string `json:"setEnv"` // Environment variables to pass to sftp and commands
    Ciphers string `json:"ciphers"` // Space separated list of ciphers to be used for session encryption, ordered by preference.
    PubkeyFile string `json:"pubkeyFile"` // Optional path to public key file.
    DisableConcurrentReads string `json:"disableConcurrentReads" default:"false"` // If set don't use concurrent reads.
    IdleTimeout string `json:"idleTimeout" default:"1m0s"` // Max time before closing idle connections.
    KeyFile string `json:"keyFile"` // Path to PEM-encoded private key file.
    Macs string `json:"macs"` // Space separated list of MACs (message authentication code) algorithms, ordered by preference.
    AskPassword string `json:"askPassword" default:"false"` // Allow asking for SFTP password when needed.
    ServerCommand string `json:"serverCommand"` // Specifies the path or command to run a sftp server on the remote host.
    Concurrency string `json:"concurrency" default:"64"` // The maximum number of outstanding requests for one file
    KeyExchange string `json:"keyExchange"` // Space separated list of key exchange algorithms, ordered by preference.
    User string `json:"user" default:"shane"` // SSH username.
    KeyUseAgent string `json:"keyUseAgent" default:"false"` // When set forces the usage of the ssh-agent.
    ShellType string `json:"shellType"` // The type of SSH shell on remote server, if any.
    UseFstat string `json:"useFstat" default:"false"` // If set use fstat instead of stat.
    DisableConcurrentWrites string `json:"disableConcurrentWrites" default:"false"` // If set don't use concurrent writes.
    ChunkSize string `json:"chunkSize" default:"32Ki"` // Upload and download chunk size.
    Pass string `json:"pass"` // SSH password, leave blank to use ssh-agent.
}

// HandleSftp godoc
// @Summary Add sftp source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body SftpRequest true "Request body"
// @Router /dataset/{datasetName}/source/sftp [post]
func HandleSftp() {}


type SharefileRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    UploadCutoff string `json:"uploadCutoff" default:"128Mi"` // Cutoff for switching to multipart upload.
    RootFolderId string `json:"rootFolderId"` // ID of the root folder.
    ChunkSize string `json:"chunkSize" default:"64Mi"` // Upload chunk size.
    Endpoint string `json:"endpoint"` // Endpoint for API calls.
    Encoding string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
}

// HandleSharefile godoc
// @Summary Add sharefile source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body SharefileRequest true "Request body"
// @Router /dataset/{datasetName}/source/sharefile [post]
func HandleSharefile() {}


type SiaRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    ApiUrl string `json:"apiUrl" default:"http://127.0.0.1:9980"` // Sia daemon API URL, like http://sia.daemon.host:9980.
    ApiPassword string `json:"apiPassword"` // Sia Daemon API Password.
    UserAgent string `json:"userAgent" default:"Sia-Agent"` // Siad User Agent
    Encoding string `json:"encoding" default:"Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

// HandleSia godoc
// @Summary Add sia source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body SiaRequest true "Request body"
// @Router /dataset/{datasetName}/source/sia [post]
func HandleSia() {}


type SmbRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    IdleTimeout string `json:"idleTimeout" default:"1m0s"` // Max time before closing idle connections.
    HideSpecialShare string `json:"hideSpecialShare" default:"true"` // Hide special shares (e.g. print$) which users aren't supposed to access.
    CaseInsensitive string `json:"caseInsensitive" default:"true"` // Whether the server is configured to be case-insensitive.
    Encoding string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
    Host string `json:"host"` // SMB server hostname to connect to.
    User string `json:"user" default:"shane"` // SMB username.
    Domain string `json:"domain" default:"WORKGROUP"` // Domain name for NTLM authentication.
    Spn string `json:"spn"` // Service principal name.
    Port string `json:"port" default:"445"` // SMB port number.
    Pass string `json:"pass"` // SMB password.
}

// HandleSmb godoc
// @Summary Add smb source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body SmbRequest true "Request body"
// @Router /dataset/{datasetName}/source/smb [post]
func HandleSmb() {}


type StorjRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    Provider string `json:"provider" default:"existing"` // Choose an authentication method.
    AccessGrant string `json:"accessGrant"` // Access grant.
    SatelliteAddress string `json:"satelliteAddress" default:"us1.storj.io"` // Satellite address.
    ApiKey string `json:"apiKey"` // API key.
    Passphrase string `json:"passphrase"` // Encryption passphrase.
}

// HandleStorj godoc
// @Summary Add storj source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body StorjRequest true "Request body"
// @Router /dataset/{datasetName}/source/storj [post]
func HandleStorj() {}


type TardigradeRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    ApiKey string `json:"apiKey"` // API key.
    Passphrase string `json:"passphrase"` // Encryption passphrase.
    Provider string `json:"provider" default:"existing"` // Choose an authentication method.
    AccessGrant string `json:"accessGrant"` // Access grant.
    SatelliteAddress string `json:"satelliteAddress" default:"us1.storj.io"` // Satellite address.
}

// HandleTardigrade godoc
// @Summary Add tardigrade source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body TardigradeRequest true "Request body"
// @Router /dataset/{datasetName}/source/tardigrade [post]
func HandleTardigrade() {}


type SugarsyncRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    PrivateAccessKey string `json:"privateAccessKey"` // Sugarsync Private Access Key.
    HardDelete string `json:"hardDelete" default:"false"` // Permanently delete files if true
    AuthorizationExpiry string `json:"authorizationExpiry"` // Sugarsync authorization expiry.
    User string `json:"user"` // Sugarsync user.
    AppId string `json:"appId"` // Sugarsync App ID.
    RefreshToken string `json:"refreshToken"` // Sugarsync refresh token.
    Authorization string `json:"authorization"` // Sugarsync authorization.
    RootId string `json:"rootId"` // Sugarsync root id.
    DeletedId string `json:"deletedId"` // Sugarsync deleted folder id.
    Encoding string `json:"encoding" default:"Slash,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    AccessKeyId string `json:"accessKeyId"` // Sugarsync Access Key ID.
}

// HandleSugarsync godoc
// @Summary Add sugarsync source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body SugarsyncRequest true "Request body"
// @Router /dataset/{datasetName}/source/sugarsync [post]
func HandleSugarsync() {}


type SwiftRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    NoLargeObjects string `json:"noLargeObjects" default:"false"` // Disable support for static and dynamic large objects
    EnvAuth string `json:"envAuth" default:"false"` // Get swift credentials from environment variables in standard OpenStack form.
    AuthToken string `json:"authToken"` // Auth Token from alternate authentication - optional (OS_AUTH_TOKEN).
    EndpointType string `json:"endpointType" default:"public"` // Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).
    NoChunk string `json:"noChunk" default:"false"` // Don't chunk files during streaming upload.
    ApplicationCredentialName string `json:"applicationCredentialName"` // Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME).
    AuthVersion string `json:"authVersion" default:"0"` // AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).
    User string `json:"user"` // User name to log in (OS_USERNAME).
    Tenant string `json:"tenant"` // Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME).
    Region string `json:"region"` // Region name - optional (OS_REGION_NAME).
    StorageUrl string `json:"storageUrl"` // Storage URL - optional (OS_STORAGE_URL).
    ApplicationCredentialSecret string `json:"applicationCredentialSecret"` // Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET).
    LeavePartsOnError string `json:"leavePartsOnError" default:"false"` // If true avoid calling abort upload on a failure.
    Encoding string `json:"encoding" default:"Slash,InvalidUtf8"` // The encoding for the backend.
    Auth string `json:"auth"` // Authentication URL for server (OS_AUTH_URL).
    UserId string `json:"userId"` // User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID).
    TenantId string `json:"tenantId"` // Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID).
    ApplicationCredentialId string `json:"applicationCredentialId"` // Application Credential ID (OS_APPLICATION_CREDENTIAL_ID).
    ChunkSize string `json:"chunkSize" default:"5Gi"` // Above this size files will be chunked into a _segments container.
    Key string `json:"key"` // API key or password (OS_PASSWORD).
    Domain string `json:"domain"` // User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME)
    TenantDomain string `json:"tenantDomain"` // Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME).
    StoragePolicy string `json:"storagePolicy"` // The storage policy to use when creating a new container.
}

// HandleSwift godoc
// @Summary Add swift source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body SwiftRequest true "Request body"
// @Router /dataset/{datasetName}/source/swift [post]
func HandleSwift() {}


type UptoboxRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    AccessToken string `json:"accessToken"` // Your access token.
    Encoding string `json:"encoding" default:"Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot"` // The encoding for the backend.
}

// HandleUptobox godoc
// @Summary Add uptobox source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body UptoboxRequest true "Request body"
// @Router /dataset/{datasetName}/source/uptobox [post]
func HandleUptobox() {}


type WebdavRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    User string `json:"user"` // User name.
    Pass string `json:"pass"` // Password.
    BearerToken string `json:"bearerToken"` // Bearer token instead of user/pass (e.g. a Macaroon).
    BearerTokenCommand string `json:"bearerTokenCommand"` // Command to run to get a bearer token.
    Encoding string `json:"encoding"` // The encoding for the backend.
    Headers string `json:"headers"` // Set HTTP headers for all transactions.
    Url string `json:"url"` // URL of http host to connect to.
    Vendor string `json:"vendor"` // Name of the WebDAV site/service/software you are using.
}

// HandleWebdav godoc
// @Summary Add webdav source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body WebdavRequest true "Request body"
// @Router /dataset/{datasetName}/source/webdav [post]
func HandleWebdav() {}


type YandexRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    ClientSecret string `json:"clientSecret"` // OAuth Client Secret.
    Token string `json:"token"` // OAuth Access Token as a JSON blob.
    AuthUrl string `json:"authUrl"` // Auth server URL.
    TokenUrl string `json:"tokenUrl"` // Token server url.
    HardDelete string `json:"hardDelete" default:"false"` // Delete files permanently rather than putting them into the trash.
    Encoding string `json:"encoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    ClientId string `json:"clientId"` // OAuth Client Id.
}

// HandleYandex godoc
// @Summary Add yandex source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body YandexRequest true "Request body"
// @Router /dataset/{datasetName}/source/yandex [post]
func HandleYandex() {}


type ZohoRequest struct {
    SourcePath string `json:"sourcePath"`// The path of the source to scan items
    TokenUrl string `json:"tokenUrl"` // Token server url.
    Region string `json:"region"` // Zoho region to connect to.
    Encoding string `json:"encoding" default:"Del,Ctl,InvalidUtf8"` // The encoding for the backend.
    ClientId string `json:"clientId"` // OAuth Client Id.
    ClientSecret string `json:"clientSecret"` // OAuth Client Secret.
    Token string `json:"token"` // OAuth Access Token as a JSON blob.
    AuthUrl string `json:"authUrl"` // Auth server URL.
}

// HandleZoho godoc
// @Summary Add zoho source for a dataset
// @Tags New Data Source
// @Accept json
// @Produce json
// @Param datasetName path string true "Dataset name"
// @Success 200 {array} model.Source
// @Param request body ZohoRequest true "Request body"
// @Router /dataset/{datasetName}/source/zoho [post]
func HandleZoho() {}


type AllConfig struct {
    AcdAuthUrl string `json:"acdAuthUrl"` // Auth server URL.
    AcdTokenUrl string `json:"acdTokenUrl"` // Token server url.
    AcdCheckpoint string `json:"acdCheckpoint"` // Checkpoint for internal polling (debug).
    AcdUploadWaitPerGb string `json:"acdUploadWaitPerGb" default:"3m0s"` // Additional time per GiB to wait after a failed complete upload to see if it appears.
    AcdTemplinkThreshold string `json:"acdTemplinkThreshold" default:"9Gi"` // Files >= this size will be downloaded via their tempLink.
    AcdEncoding string `json:"acdEncoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
    AcdClientId string `json:"acdClientId"` // OAuth Client Id.
    AcdClientSecret string `json:"acdClientSecret"` // OAuth Client Secret.
    AcdToken string `json:"acdToken"` // OAuth Access Token as a JSON blob.
    AzureblobAccount string `json:"azureblobAccount"` // Azure Storage Account Name.
    AzureblobClientCertificatePassword string `json:"azureblobClientCertificatePassword"` // Password for the certificate file (optional).
    AzureblobMsiObjectId string `json:"azureblobMsiObjectId"` // Object ID of the user-assigned MSI to use, if any.
    AzureblobChunkSize string `json:"azureblobChunkSize" default:"4Mi"` // Upload chunk size.
    AzureblobUploadConcurrency string `json:"azureblobUploadConcurrency" default:"16"` // Concurrency for multipart uploads.
    AzureblobClientSendCertificateChain string `json:"azureblobClientSendCertificateChain" default:"false"` // Send the certificate chain when using certificate auth.
    AzureblobPassword string `json:"azureblobPassword"` // The user's password
    AzureblobEndpoint string `json:"azureblobEndpoint"` // Endpoint for the service.
    AzureblobUploadCutoff string `json:"azureblobUploadCutoff"` // Cutoff for switching to chunked upload (<= 256 MiB) (deprecated).
    AzureblobNoCheckContainer string `json:"azureblobNoCheckContainer" default:"false"` // If set, don't attempt to check the container exists or create it.
    AzureblobKey string `json:"azureblobKey"` // Storage Account Shared Key.
    AzureblobClientId string `json:"azureblobClientId"` // The ID of the client in use.
    AzureblobUseMsi string `json:"azureblobUseMsi" default:"false"` // Use a managed service identity to authenticate (only works in Azure).
    AzureblobMsiClientId string `json:"azureblobMsiClientId"` // Object ID of the user-assigned MSI to use, if any.
    AzureblobClientCertificatePath string `json:"azureblobClientCertificatePath"` // Path to a PEM or PKCS12 certificate file including the private key.
    AzureblobArchiveTierDelete string `json:"azureblobArchiveTierDelete" default:"false"` // Delete archive tier blobs before overwriting.
    AzureblobMemoryPoolUseMmap string `json:"azureblobMemoryPoolUseMmap" default:"false"` // Whether to use mmap buffers in internal memory pool.
    AzureblobPublicAccess string `json:"azureblobPublicAccess"` // Public access level of a container: blob or container.
    AzureblobSasUrl string `json:"azureblobSasUrl"` // SAS URL for container level access only.
    AzureblobDisableChecksum string `json:"azureblobDisableChecksum" default:"false"` // Don't store MD5 checksum with object metadata.
    AzureblobEncoding string `json:"azureblobEncoding" default:"Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"` // The encoding for the backend.
    AzureblobNoHeadObject string `json:"azureblobNoHeadObject" default:"false"` // If set, do not do HEAD before GET when getting objects.
    AzureblobMsiMiResId string `json:"azureblobMsiMiResId"` // Azure resource ID of the user-assigned MSI to use, if any.
    AzureblobUseEmulator string `json:"azureblobUseEmulator" default:"false"` // Uses local storage emulator if provided as 'true'.
    AzureblobEnvAuth string `json:"azureblobEnvAuth" default:"false"` // Read credentials from runtime (environment variables, CLI or MSI).
    AzureblobTenant string `json:"azureblobTenant"` // ID of the service principal's tenant. Also called its directory ID.
    AzureblobServicePrincipalFile string `json:"azureblobServicePrincipalFile"` // Path to file containing credentials for use with a service principal.
    AzureblobListChunk string `json:"azureblobListChunk" default:"5000"` // Size of blob list.
    AzureblobMemoryPoolFlushTime string `json:"azureblobMemoryPoolFlushTime" default:"1m0s"` // How often internal memory buffer pools will be flushed.
    AzureblobClientSecret string `json:"azureblobClientSecret"` // One of the service principal's client secrets
    AzureblobUsername string `json:"azureblobUsername"` // User name (usually an email address)
    AzureblobAccessTier string `json:"azureblobAccessTier"` // Access tier of blob: hot, cool or archive.
    B2DisableChecksum string `json:"b2DisableChecksum" default:"false"` // Disable checksums for large (> upload cutoff) files.
    B2MemoryPoolUseMmap string `json:"b2MemoryPoolUseMmap" default:"false"` // Whether to use mmap buffers in internal memory pool.
    B2Key string `json:"b2Key"` // Application Key.
    B2Endpoint string `json:"b2Endpoint"` // Endpoint for the service.
    B2TestMode string `json:"b2TestMode"` // A flag string for X-Bz-Test-Mode header for debugging.
    B2Encoding string `json:"b2Encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    B2Account string `json:"b2Account"` // Account ID or Application Key ID.
    B2ChunkSize string `json:"b2ChunkSize" default:"96Mi"` // Upload chunk size.
    B2HardDelete string `json:"b2HardDelete" default:"false"` // Permanently delete files on remote removal, otherwise hide files.
    B2UploadCutoff string `json:"b2UploadCutoff" default:"200Mi"` // Cutoff for switching to chunked upload.
    B2CopyCutoff string `json:"b2CopyCutoff" default:"4Gi"` // Cutoff for switching to multipart copy.
    B2DownloadUrl string `json:"b2DownloadUrl"` // Custom endpoint for downloads.
    B2DownloadAuthDuration string `json:"b2DownloadAuthDuration" default:"1w"` // Time before the authorization token will expire in s or suffix ms|s|m|h|d.
    B2MemoryPoolFlushTime string `json:"b2MemoryPoolFlushTime" default:"1m0s"` // How often internal memory buffer pools will be flushed.
    B2Versions string `json:"b2Versions" default:"false"` // Include old versions in directory listings.
    B2VersionAt string `json:"b2VersionAt" default:"off"` // Show file versions as they were at the specified time.
    BoxClientSecret string `json:"boxClientSecret"` // OAuth Client Secret.
    BoxListChunk string `json:"boxListChunk" default:"1000"` // Size of listing chunk 1-1000.
    BoxTokenUrl string `json:"boxTokenUrl"` // Token server url.
    BoxBoxSubType string `json:"boxBoxSubType" default:"user"` // 
    BoxUploadCutoff string `json:"boxUploadCutoff" default:"50Mi"` // Cutoff for switching to multipart upload (>= 50 MiB).
    BoxOwnedBy string `json:"boxOwnedBy"` // Only show items owned by the login (email address) passed in.
    BoxEncoding string `json:"boxEncoding" default:"Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
    BoxAuthUrl string `json:"boxAuthUrl"` // Auth server URL.
    BoxBoxConfigFile string `json:"boxBoxConfigFile"` // Box App config.json location
    BoxCommitRetries string `json:"boxCommitRetries" default:"100"` // Max number of times to try committing a multipart file.
    BoxClientId string `json:"boxClientId"` // OAuth Client Id.
    BoxToken string `json:"boxToken"` // OAuth Access Token as a JSON blob.
    BoxRootFolderId string `json:"boxRootFolderId" default:"0"` // Fill in for rclone to use a non root folder as its starting point.
    BoxAccessToken string `json:"boxAccessToken"` // Box App Primary Access Token
    CryptRemote string `json:"cryptRemote"` // Remote to encrypt/decrypt.
    CryptFilenameEncryption string `json:"cryptFilenameEncryption" default:"standard"` // How to encrypt the filenames.
    CryptPassword string `json:"cryptPassword"` // Password or pass phrase for encryption.
    CryptPassword2 string `json:"cryptPassword2"` // Password or pass phrase for salt.
    CryptShowMapping string `json:"cryptShowMapping" default:"false"` // For all files listed show how the names encrypt.
    CryptFilenameEncoding string `json:"cryptFilenameEncoding" default:"base32"` // How to encode the encrypted filename to text string.
    CryptDirectoryNameEncryption string `json:"cryptDirectoryNameEncryption" default:"true"` // Option to either encrypt directory names or leave them intact.
    CryptServerSideAcrossConfigs string `json:"cryptServerSideAcrossConfigs" default:"false"` // Allow server-side operations (e.g. copy) to work across different crypt configs.
    CryptNoDataEncryption string `json:"cryptNoDataEncryption" default:"false"` // Option to either encrypt file data or leave it unencrypted.
    DriveScope string `json:"driveScope"` // Scope that rclone should use when requesting access from drive.
    DriveServiceAccountCredentials string `json:"driveServiceAccountCredentials"` // Service Account Credentials JSON blob.
    DriveAuthOwnerOnly string `json:"driveAuthOwnerOnly" default:"false"` // Only consider files owned by the authenticated user.
    DriveStarredOnly string `json:"driveStarredOnly" default:"false"` // Only show files that are starred.
    DriveSkipDanglingShortcuts string `json:"driveSkipDanglingShortcuts" default:"false"` // If set skip dangling shortcut files.
    DriveClientSecret string `json:"driveClientSecret"` // OAuth Client Secret.
    DriveCopyShortcutContent string `json:"driveCopyShortcutContent" default:"false"` // Server side copy contents of shortcuts instead of the shortcut.
    DriveImpersonate string `json:"driveImpersonate"` // Impersonate this user when using a service account.
    DriveStopOnUploadLimit string `json:"driveStopOnUploadLimit" default:"false"` // Make upload limit errors be fatal.
    DriveChunkSize string `json:"driveChunkSize" default:"8Mi"` // Upload chunk size.
    DriveDisableHttp2 string `json:"driveDisableHttp2" default:"true"` // Disable drive using http2.
    DriveTokenUrl string `json:"driveTokenUrl"` // Token server url.
    DriveServiceAccountFile string `json:"driveServiceAccountFile"` // Service Account Credentials JSON file path.
    DriveAllowImportNameChange string `json:"driveAllowImportNameChange" default:"false"` // Allow the filetype to change when uploading Google docs.
    DriveListChunk string `json:"driveListChunk" default:"1000"` // Size of listing chunk 100-1000, 0 to disable.
    DriveAlternateExport string `json:"driveAlternateExport" default:"false"` // Deprecated: No longer needed.
    DriveUploadCutoff string `json:"driveUploadCutoff" default:"8Mi"` // Cutoff for switching to chunked upload.
    DriveResourceKey string `json:"driveResourceKey"` // Resource key for accessing a link-shared file.
    DriveToken string `json:"driveToken"` // OAuth Access Token as a JSON blob.
    DriveRootFolderId string `json:"driveRootFolderId"` // ID of the root folder.
    DriveSharedWithMe string `json:"driveSharedWithMe" default:"false"` // Only show files that are shared with me.
    DriveAcknowledgeAbuse string `json:"driveAcknowledgeAbuse" default:"false"` // Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
    DriveSizeAsQuota string `json:"driveSizeAsQuota" default:"false"` // Show sizes as storage quota usage, not actual size.
    DriveSkipShortcuts string `json:"driveSkipShortcuts" default:"false"` // If set skip shortcut files.
    DriveClientId string `json:"driveClientId"` // Google Application Client Id
    DriveUseTrash string `json:"driveUseTrash" default:"true"` // Send files to the trash instead of deleting permanently.
    DriveUseSharedDate string `json:"driveUseSharedDate" default:"false"` // Use date file was shared instead of modified date.
    DriveAuthUrl string `json:"driveAuthUrl"` // Auth server URL.
    DriveTrashedOnly string `json:"driveTrashedOnly" default:"false"` // Only show files that are in the trash.
    DriveKeepRevisionForever string `json:"driveKeepRevisionForever" default:"false"` // Keep new head revision of each file forever.
    DriveServerSideAcrossConfigs string `json:"driveServerSideAcrossConfigs" default:"false"` // Allow server-side operations (e.g. copy) to work across different drive configs.
    DriveExportFormats string `json:"driveExportFormats" default:"docx,xlsx,pptx,svg"` // Comma separated list of preferred formats for downloading Google docs.
    DriveImportFormats string `json:"driveImportFormats"` // Comma separated list of preferred formats for uploading Google docs.
    DriveV2DownloadMinSize string `json:"driveV2DownloadMinSize" default:"off"` // If Object's are greater, use drive v2 API to download.
    DrivePacerBurst string `json:"drivePacerBurst" default:"100"` // Number of API calls to allow without sleeping.
    DriveEncoding string `json:"driveEncoding" default:"InvalidUtf8"` // The encoding for the backend.
    DriveStopOnDownloadLimit string `json:"driveStopOnDownloadLimit" default:"false"` // Make download limit errors be fatal.
    DriveTeamDrive string `json:"driveTeamDrive"` // ID of the Shared Drive (Team Drive).
    DriveSkipGdocs string `json:"driveSkipGdocs" default:"false"` // Skip google documents in all listings.
    DriveSkipChecksumGphotos string `json:"driveSkipChecksumGphotos" default:"false"` // Skip MD5 checksum on Google photos and videos only.
    DriveFormats string `json:"driveFormats"` // Deprecated: See export_formats.
    DriveUseCreatedDate string `json:"driveUseCreatedDate" default:"false"` // Use file created date instead of modified date.
    DrivePacerMinSleep string `json:"drivePacerMinSleep" default:"100ms"` // Minimum time to sleep between API calls.
    DropboxSharedFiles string `json:"dropboxSharedFiles" default:"false"` // Instructs rclone to work on individual shared files.
    DropboxSharedFolders string `json:"dropboxSharedFolders" default:"false"` // Instructs rclone to work on shared folders.
    DropboxBatchMode string `json:"dropboxBatchMode" default:"sync"` // Upload file batching sync|async|off.
    DropboxEncoding string `json:"dropboxEncoding" default:"Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
    DropboxClientId string `json:"dropboxClientId"` // OAuth Client Id.
    DropboxToken string `json:"dropboxToken"` // OAuth Access Token as a JSON blob.
    DropboxAuthUrl string `json:"dropboxAuthUrl"` // Auth server URL.
    DropboxTokenUrl string `json:"dropboxTokenUrl"` // Token server url.
    DropboxClientSecret string `json:"dropboxClientSecret"` // OAuth Client Secret.
    DropboxImpersonate string `json:"dropboxImpersonate"` // Impersonate this user when using a business account.
    DropboxBatchTimeout string `json:"dropboxBatchTimeout" default:"0s"` // Max time to allow an idle upload batch before uploading.
    DropboxBatchCommitTimeout string `json:"dropboxBatchCommitTimeout" default:"10m0s"` // Max time to wait for a batch to finish committing
    DropboxChunkSize string `json:"dropboxChunkSize" default:"48Mi"` // Upload chunk size (< 150Mi).
    DropboxBatchSize string `json:"dropboxBatchSize" default:"0"` // Max number of files in upload batch.
    FichierApiKey string `json:"fichierApiKey"` // Your API Key, get it from https://1fichier.com/console/params.pl.
    FichierSharedFolder string `json:"fichierSharedFolder"` // If you want to download a shared folder, add this parameter.
    FichierFilePassword string `json:"fichierFilePassword"` // If you want to download a shared file that is password protected, add this parameter.
    FichierFolderPassword string `json:"fichierFolderPassword"` // If you want to list the files in a shared folder that is password protected, add this parameter.
    FichierEncoding string `json:"fichierEncoding" default:"Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
    FilefabricPermanentToken string `json:"filefabricPermanentToken"` // Permanent Authentication Token.
    FilefabricToken string `json:"filefabricToken"` // Session Token.
    FilefabricTokenExpiry string `json:"filefabricTokenExpiry"` // Token expiry time.
    FilefabricVersion string `json:"filefabricVersion"` // Version read from the file fabric.
    FilefabricEncoding string `json:"filefabricEncoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    FilefabricUrl string `json:"filefabricUrl"` // URL of the Enterprise File Fabric to connect to.
    FilefabricRootFolderId string `json:"filefabricRootFolderId"` // ID of the root folder.
    FtpPass string `json:"ftpPass"` // FTP password.
    FtpTls string `json:"ftpTls" default:"false"` // Use Implicit FTPS (FTP over TLS).
    FtpExplicitTls string `json:"ftpExplicitTls" default:"false"` // Use Explicit FTPS (FTP over TLS).
    FtpDisableEpsv string `json:"ftpDisableEpsv" default:"false"` // Disable using EPSV even if server advertises support.
    FtpIdleTimeout string `json:"ftpIdleTimeout" default:"1m0s"` // Max time before closing idle connections.
    FtpCloseTimeout string `json:"ftpCloseTimeout" default:"1m0s"` // Maximum time to wait for a response to close.
    FtpShutTimeout string `json:"ftpShutTimeout" default:"1m0s"` // Maximum time to wait for data connection closing status.
    FtpUser string `json:"ftpUser" default:"shane"` // FTP username.
    FtpEncoding string `json:"ftpEncoding" default:"Slash,Del,Ctl,RightSpace,Dot"` // The encoding for the backend.
    FtpDisableUtf8 string `json:"ftpDisableUtf8" default:"false"` // Disable using UTF-8 even if server advertises support.
    FtpForceListHidden string `json:"ftpForceListHidden" default:"false"` // Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.
    FtpAskPassword string `json:"ftpAskPassword" default:"false"` // Allow asking for FTP password when needed.
    FtpNoCheckCertificate string `json:"ftpNoCheckCertificate" default:"false"` // Do not verify the TLS certificate of the server.
    FtpPort string `json:"ftpPort" default:"21"` // FTP port number.
    FtpTlsCacheSize string `json:"ftpTlsCacheSize" default:"32"` // Size of TLS session cache for all control and data connections.
    FtpHost string `json:"ftpHost"` // FTP host to connect to.
    FtpDisableMlsd string `json:"ftpDisableMlsd" default:"false"` // Disable using MLSD even if server advertises support.
    FtpWritingMdtm string `json:"ftpWritingMdtm" default:"false"` // Use MDTM to set modification time (VsFtpd quirk)
    FtpDisableTls13 string `json:"ftpDisableTls13" default:"false"` // Disable TLS 1.3 (workaround for FTP servers with buggy TLS)
    FtpConcurrency string `json:"ftpConcurrency" default:"0"` // Maximum number of FTP simultaneous connections, 0 for unlimited.
    GcsBucketAcl string `json:"gcsBucketAcl"` // Access Control List for new buckets.
    GcsBucketPolicyOnly string `json:"gcsBucketPolicyOnly" default:"false"` // Access checks should use bucket-level IAM policies.
    GcsEndpoint string `json:"gcsEndpoint"` // Endpoint for the service.
    GcsEnvAuth string `json:"gcsEnvAuth" default:"false"` // Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
    GcsClientId string `json:"gcsClientId"` // OAuth Client Id.
    GcsServiceAccountFile string `json:"gcsServiceAccountFile"` // Service Account Credentials JSON file path.
    GcsObjectAcl string `json:"gcsObjectAcl"` // Access Control List for new objects.
    GcsLocation string `json:"gcsLocation"` // Location for the newly created buckets.
    GcsAuthUrl string `json:"gcsAuthUrl"` // Auth server URL.
    GcsTokenUrl string `json:"gcsTokenUrl"` // Token server url.
    GcsProjectNumber string `json:"gcsProjectNumber"` // Project number.
    GcsAnonymous string `json:"gcsAnonymous" default:"false"` // Access public buckets and objects without credentials.
    GcsDecompress string `json:"gcsDecompress" default:"false"` // If set this will decompress gzip encoded objects.
    GcsEncoding string `json:"gcsEncoding" default:"Slash,CrLf,InvalidUtf8,Dot"` // The encoding for the backend.
    GcsClientSecret string `json:"gcsClientSecret"` // OAuth Client Secret.
    GcsToken string `json:"gcsToken"` // OAuth Access Token as a JSON blob.
    GcsServiceAccountCredentials string `json:"gcsServiceAccountCredentials"` // Service Account Credentials JSON blob.
    GcsStorageClass string `json:"gcsStorageClass"` // The storage class to use when storing objects in Google Cloud Storage.
    GcsNoCheckBucket string `json:"gcsNoCheckBucket" default:"false"` // If set, don't attempt to check the bucket exists or create it.
    GphotosIncludeArchived string `json:"gphotosIncludeArchived" default:"false"` // Also view and download archived media.
    GphotosEncoding string `json:"gphotosEncoding" default:"Slash,CrLf,InvalidUtf8,Dot"` // The encoding for the backend.
    GphotosClientSecret string `json:"gphotosClientSecret"` // OAuth Client Secret.
    GphotosAuthUrl string `json:"gphotosAuthUrl"` // Auth server URL.
    GphotosReadOnly string `json:"gphotosReadOnly" default:"false"` // Set to make the Google Photos backend read only.
    GphotosReadSize string `json:"gphotosReadSize" default:"false"` // Set to read the size of media items.
    GphotosStartYear string `json:"gphotosStartYear" default:"2000"` // Year limits the photos to be downloaded to those which are uploaded after the given year.
    GphotosClientId string `json:"gphotosClientId"` // OAuth Client Id.
    GphotosToken string `json:"gphotosToken"` // OAuth Access Token as a JSON blob.
    GphotosTokenUrl string `json:"gphotosTokenUrl"` // Token server url.
    HdfsDataTransferProtection string `json:"hdfsDataTransferProtection"` // Kerberos data transfer protection: authentication|integrity|privacy.
    HdfsEncoding string `json:"hdfsEncoding" default:"Slash,Colon,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    HdfsNamenode string `json:"hdfsNamenode"` // Hadoop name node and port.
    HdfsUsername string `json:"hdfsUsername"` // Hadoop user name.
    HdfsServicePrincipalName string `json:"hdfsServicePrincipalName"` // Kerberos service principal name for the namenode.
    HidriveRootPrefix string `json:"hidriveRootPrefix" default:"/"` // The root/parent folder for all paths.
    HidriveUploadConcurrency string `json:"hidriveUploadConcurrency" default:"4"` // Concurrency for chunked uploads.
    HidriveToken string `json:"hidriveToken"` // OAuth Access Token as a JSON blob.
    HidriveAuthUrl string `json:"hidriveAuthUrl"` // Auth server URL.
    HidriveScopeRole string `json:"hidriveScopeRole" default:"user"` // User-level that rclone should use when requesting access from HiDrive.
    HidriveChunkSize string `json:"hidriveChunkSize" default:"48Mi"` // Chunksize for chunked uploads.
    HidriveUploadCutoff string `json:"hidriveUploadCutoff" default:"96Mi"` // Cutoff/Threshold for chunked uploads.
    HidriveEncoding string `json:"hidriveEncoding" default:"Slash,Dot"` // The encoding for the backend.
    HidriveClientId string `json:"hidriveClientId"` // OAuth Client Id.
    HidriveClientSecret string `json:"hidriveClientSecret"` // OAuth Client Secret.
    HidriveDisableFetchingMemberCount string `json:"hidriveDisableFetchingMemberCount" default:"false"` // Do not fetch number of objects in directories unless it is absolutely necessary.
    HidriveTokenUrl string `json:"hidriveTokenUrl"` // Token server url.
    HidriveEndpoint string `json:"hidriveEndpoint" default:"https://api.hidrive.strato.com/2.1"` // Endpoint for the service.
    HidriveScopeAccess string `json:"hidriveScopeAccess" default:"rw"` // Access permissions that rclone should use when requesting access from HiDrive.
    HttpUrl string `json:"httpUrl"` // URL of HTTP host to connect to.
    HttpHeaders string `json:"httpHeaders"` // Set HTTP headers for all transactions.
    HttpNoSlash string `json:"httpNoSlash" default:"false"` // Set this if the site doesn't end directories with /.
    HttpNoHead string `json:"httpNoHead" default:"false"` // Don't use HEAD requests.
    InternetarchiveEncoding string `json:"internetarchiveEncoding" default:"Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    InternetarchiveAccessKeyId string `json:"internetarchiveAccessKeyId"` // IAS3 Access Key.
    InternetarchiveSecretAccessKey string `json:"internetarchiveSecretAccessKey"` // IAS3 Secret Key (password).
    InternetarchiveEndpoint string `json:"internetarchiveEndpoint" default:"https://s3.us.archive.org"` // IAS3 Endpoint.
    InternetarchiveFrontEndpoint string `json:"internetarchiveFrontEndpoint" default:"https://archive.org"` // Host of InternetArchive Frontend.
    InternetarchiveDisableChecksum string `json:"internetarchiveDisableChecksum" default:"true"` // Don't ask the server to test against MD5 checksum calculated by rclone.
    InternetarchiveWaitArchive string `json:"internetarchiveWaitArchive" default:"0s"` // Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
    JottacloudEncoding string `json:"jottacloudEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    JottacloudMd5MemoryLimit string `json:"jottacloudMd5MemoryLimit" default:"10Mi"` // Files bigger than this will be cached on disk to calculate the MD5 if required.
    JottacloudTrashedOnly string `json:"jottacloudTrashedOnly" default:"false"` // Only show files that are in the trash.
    JottacloudHardDelete string `json:"jottacloudHardDelete" default:"false"` // Delete files permanently rather than putting them into the trash.
    JottacloudUploadResumeLimit string `json:"jottacloudUploadResumeLimit" default:"10Mi"` // Files bigger than this can be resumed if the upload fail's.
    JottacloudNoVersions string `json:"jottacloudNoVersions" default:"false"` // Avoid server side versioning by deleting files and recreating files instead of overwriting them.
    KoofrEndpoint string `json:"koofrEndpoint"` // The Koofr API endpoint to use.
    KoofrMountid string `json:"koofrMountid"` // Mount ID of the mount to use.
    KoofrSetmtime string `json:"koofrSetmtime" default:"true"` // Does the backend support setting modification time.
    KoofrUser string `json:"koofrUser"` // Your user name.
    KoofrPassword string `json:"koofrPassword"` // Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password).
    KoofrEncoding string `json:"koofrEncoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    KoofrProvider string `json:"koofrProvider"` // Choose your storage provider.
    LocalZeroSizeLinks string `json:"localZeroSizeLinks" default:"false"` // Assume the Stat size of links is zero (and read them instead) (deprecated).
    LocalNoSparse string `json:"localNoSparse" default:"false"` // Disable sparse files for multi-thread downloads.
    LocalNoSetModtime string `json:"localNoSetModtime" default:"false"` // Disable setting modtime.
    LocalCopyLinks string `json:"localCopyLinks" default:"false"` // Follow symlinks and copy the pointed to item.
    LocalSkipLinks string `json:"localSkipLinks" default:"false"` // Don't warn about skipped symlinks.
    LocalNoCheckUpdated string `json:"localNoCheckUpdated" default:"false"` // Don't check to see if the files change during upload.
    LocalOneFileSystem string `json:"localOneFileSystem" default:"false"` // Don't cross filesystem boundaries (unix/macOS only).
    LocalCaseSensitive string `json:"localCaseSensitive" default:"false"` // Force the filesystem to report itself as case sensitive.
    LocalNoPreallocate string `json:"localNoPreallocate" default:"false"` // Disable preallocation of disk space for transferred files.
    LocalNounc string `json:"localNounc" default:"false"` // Disable UNC (long path names) conversion on Windows.
    LocalLinks string `json:"localLinks" default:"false"` // Translate symlinks to/from regular files with a '.rclonelink' extension.
    LocalUnicodeNormalization string `json:"localUnicodeNormalization" default:"false"` // Apply unicode NFC normalization to paths and filenames.
    LocalCaseInsensitive string `json:"localCaseInsensitive" default:"false"` // Force the filesystem to report itself as case insensitive.
    LocalEncoding string `json:"localEncoding" default:"Slash,Dot"` // The encoding for the backend.
    MailruSpeedupMaxMemory string `json:"mailruSpeedupMaxMemory" default:"32Mi"` // Files larger than the size given below will always be hashed on disk.
    MailruEncoding string `json:"mailruEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    MailruPass string `json:"mailruPass"` // Password.
    MailruSpeedupMaxDisk string `json:"mailruSpeedupMaxDisk" default:"3Gi"` // This option allows you to disable speedup (put by hash) for large files.
    MailruSpeedupFilePatterns string `json:"mailruSpeedupFilePatterns" default:"*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf"` // Comma separated list of file name patterns eligible for speedup (put by hash).
    MailruCheckHash string `json:"mailruCheckHash" default:"true"` // What should copy do if file checksum is mismatched or invalid.
    MailruUserAgent string `json:"mailruUserAgent"` // HTTP user agent used internally by client.
    MailruQuirks string `json:"mailruQuirks"` // Comma separated list of internal maintenance flags.
    MailruUser string `json:"mailruUser"` // User name (usually email).
    MailruSpeedupEnable string `json:"mailruSpeedupEnable" default:"true"` // Skip full upload if there is another file with same data hash.
    MegaUser string `json:"megaUser"` // User name.
    MegaPass string `json:"megaPass"` // Password.
    MegaDebug string `json:"megaDebug" default:"false"` // Output more debug from Mega.
    MegaHardDelete string `json:"megaHardDelete" default:"false"` // Delete files permanently rather than putting them into the trash.
    MegaUseHttps string `json:"megaUseHttps" default:"false"` // Use HTTPS for transfers.
    MegaEncoding string `json:"megaEncoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
    NetstorageProtocol string `json:"netstorageProtocol" default:"https"` // Select between HTTP or HTTPS protocol.
    NetstorageHost string `json:"netstorageHost"` // Domain+path of NetStorage host to connect to.
    NetstorageAccount string `json:"netstorageAccount"` // Set the NetStorage account name
    NetstorageSecret string `json:"netstorageSecret"` // Set the NetStorage account secret/G2O key for authentication.
    OnedriveClientSecret string `json:"onedriveClientSecret"` // OAuth Client Secret.
    OnedriveToken string `json:"onedriveToken"` // OAuth Access Token as a JSON blob.
    OnedriveServerSideAcrossConfigs string `json:"onedriveServerSideAcrossConfigs" default:"false"` // Allow server-side operations (e.g. copy) to work across different onedrive configs.
    OnedriveEncoding string `json:"onedriveEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
    OnedriveTokenUrl string `json:"onedriveTokenUrl"` // Token server url.
    OnedriveDriveType string `json:"onedriveDriveType"` // The type of the drive (personal | business | documentLibrary).
    OnedriveDisableSitePermission string `json:"onedriveDisableSitePermission" default:"false"` // Disable the request for Sites.Read.All permission.
    OnedriveLinkPassword string `json:"onedriveLinkPassword"` // Set the password for links created by the link command.
    OnedriveLinkType string `json:"onedriveLinkType" default:"view"` // Set the type of the links created by the link command.
    OnedriveAuthUrl string `json:"onedriveAuthUrl"` // Auth server URL.
    OnedriveChunkSize string `json:"onedriveChunkSize" default:"10Mi"` // Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
    OnedriveExposeOnenoteFiles string `json:"onedriveExposeOnenoteFiles" default:"false"` // Set to make OneNote files show up in directory listings.
    OnedriveListChunk string `json:"onedriveListChunk" default:"1000"` // Size of listing chunk.
    OnedriveAccessScopes string `json:"onedriveAccessScopes" default:"Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"` // Set scopes to be requested by rclone.
    OnedriveNoVersions string `json:"onedriveNoVersions" default:"false"` // Remove all versions on modifying operations.
    OnedriveLinkScope string `json:"onedriveLinkScope" default:"anonymous"` // Set the scope of the links created by the link command.
    OnedriveHashType string `json:"onedriveHashType" default:"auto"` // Specify the hash in use for the backend.
    OnedriveClientId string `json:"onedriveClientId"` // OAuth Client Id.
    OnedriveRegion string `json:"onedriveRegion" default:"global"` // Choose national cloud region for OneDrive.
    OnedriveDriveId string `json:"onedriveDriveId"` // The ID of the drive to use.
    OnedriveRootFolderId string `json:"onedriveRootFolderId"` // ID of the root folder.
    OpendriveUsername string `json:"opendriveUsername"` // Username.
    OpendrivePassword string `json:"opendrivePassword"` // Password.
    OpendriveEncoding string `json:"opendriveEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot"` // The encoding for the backend.
    OpendriveChunkSize string `json:"opendriveChunkSize" default:"10Mi"` // Files will be uploaded in chunks this size.
    OosProvider string `json:"oosProvider" default:"env_auth"` // Choose your Auth Provider
    OosNamespace string `json:"oosNamespace"` // Object storage namespace
    OosConfigProfile string `json:"oosConfigProfile" default:"Default"` // Profile name inside the oci config file
    OosUploadConcurrency string `json:"oosUploadConcurrency" default:"10"` // Concurrency for multipart uploads.
    OosSseCustomerKey string `json:"oosSseCustomerKey"` // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
    OosSseKmsKeyId string `json:"oosSseKmsKeyId"` // if using using your own master key in vault, this header specifies the 
    OosEndpoint string `json:"oosEndpoint"` // Endpoint for Object storage API.
    OosStorageTier string `json:"oosStorageTier" default:"Standard"` // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
    OosChunkSize string `json:"oosChunkSize" default:"5Mi"` // Chunk size to use for uploading.
    OosCopyTimeout string `json:"oosCopyTimeout" default:"1m0s"` // Timeout for copy.
    OosLeavePartsOnError string `json:"oosLeavePartsOnError" default:"false"` // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
    OosEncoding string `json:"oosEncoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
    OosNoCheckBucket string `json:"oosNoCheckBucket" default:"false"` // If set, don't attempt to check the bucket exists or create it.
    OosCompartment string `json:"oosCompartment"` // Object storage compartment OCID
    OosRegion string `json:"oosRegion"` // Object storage Region
    OosConfigFile string `json:"oosConfigFile" default:"~/.oci/config"` // Path to OCI config file
    OosUploadCutoff string `json:"oosUploadCutoff" default:"200Mi"` // Cutoff for switching to chunked upload.
    OosCopyCutoff string `json:"oosCopyCutoff" default:"4.656Gi"` // Cutoff for switching to multipart copy.
    OosDisableChecksum string `json:"oosDisableChecksum" default:"false"` // Don't store MD5 checksum with object metadata.
    OosSseCustomerKeyFile string `json:"oosSseCustomerKeyFile"` // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
    OosSseCustomerKeySha256 string `json:"oosSseCustomerKeySha256"` // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
    OosSseCustomerAlgorithm string `json:"oosSseCustomerAlgorithm"` // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
    PcloudToken string `json:"pcloudToken"` // OAuth Access Token as a JSON blob.
    PcloudTokenUrl string `json:"pcloudTokenUrl"` // Token server url.
    PcloudHostname string `json:"pcloudHostname" default:"api.pcloud.com"` // Hostname to connect to.
    PcloudPassword string `json:"pcloudPassword"` // Your pcloud password.
    PcloudRootFolderId string `json:"pcloudRootFolderId" default:"d0"` // Fill in for rclone to use a non root folder as its starting point.
    PcloudUsername string `json:"pcloudUsername"` // Your pcloud username.
    PcloudClientId string `json:"pcloudClientId"` // OAuth Client Id.
    PcloudClientSecret string `json:"pcloudClientSecret"` // OAuth Client Secret.
    PcloudAuthUrl string `json:"pcloudAuthUrl"` // Auth server URL.
    PcloudEncoding string `json:"pcloudEncoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    PremiumizemeApiKey string `json:"premiumizemeApiKey"` // API Key.
    PremiumizemeEncoding string `json:"premiumizemeEncoding" default:"Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    PutioEncoding string `json:"putioEncoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    QingstorEnvAuth string `json:"qingstorEnvAuth" default:"false"` // Get QingStor credentials from runtime.
    QingstorAccessKeyId string `json:"qingstorAccessKeyId"` // QingStor Access Key ID.
    QingstorEndpoint string `json:"qingstorEndpoint"` // Enter an endpoint URL to connection QingStor API.
    QingstorZone string `json:"qingstorZone"` // Zone to connect to.
    QingstorUploadCutoff string `json:"qingstorUploadCutoff" default:"200Mi"` // Cutoff for switching to chunked upload.
    QingstorSecretAccessKey string `json:"qingstorSecretAccessKey"` // QingStor Secret Access Key (password).
    QingstorConnectionRetries string `json:"qingstorConnectionRetries" default:"3"` // Number of connection retries.
    QingstorChunkSize string `json:"qingstorChunkSize" default:"4Mi"` // Chunk size to use for uploading.
    QingstorUploadConcurrency string `json:"qingstorUploadConcurrency" default:"1"` // Concurrency for multipart uploads.
    QingstorEncoding string `json:"qingstorEncoding" default:"Slash,Ctl,InvalidUtf8"` // The encoding for the backend.
    S3UploadCutoff string `json:"s3UploadCutoff" default:"200Mi"` // Cutoff for switching to chunked upload.
    S3CopyCutoff string `json:"s3CopyCutoff" default:"4.656Gi"` // Cutoff for switching to multipart copy.
    S3ForcePathStyle string `json:"s3ForcePathStyle" default:"true"` // If true use path style access if false use virtual hosted style.
    S3UseAccelerateEndpoint string `json:"s3UseAccelerateEndpoint" default:"false"` // If true use the AWS S3 accelerated endpoint.
    S3ListChunk string `json:"s3ListChunk" default:"1000"` // Size of listing chunk (response list for each ListObject S3 request).
    S3AccessKeyId string `json:"s3AccessKeyId"` // AWS Access Key ID.
    S3Region string `json:"s3Region"` // Region to connect to.
    S3SseCustomerKeyBase64 string `json:"s3SseCustomerKeyBase64"` // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
    S3StsEndpoint string `json:"s3StsEndpoint"` // Endpoint for STS.
    S3ListVersion string `json:"s3ListVersion" default:"0"` // Version of ListObjects to use: 1,2 or 0 for auto.
    S3Encoding string `json:"s3Encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
    S3NoSystemMetadata string `json:"s3NoSystemMetadata" default:"false"` // Suppress setting and reading of system metadata
    S3SseKmsKeyId string `json:"s3SseKmsKeyId"` // If using KMS ID you must provide the ARN of Key.
    S3DisableChecksum string `json:"s3DisableChecksum" default:"false"` // Don't store MD5 checksum with object metadata.
    S3NoCheckBucket string `json:"s3NoCheckBucket" default:"false"` // If set, don't attempt to check the bucket exists or create it.
    S3UsePresignedRequest string `json:"s3UsePresignedRequest" default:"false"` // Whether to use a presigned request or PutObject for single part uploads
    S3Endpoint string `json:"s3Endpoint"` // Endpoint for S3 API.
    S3LocationConstraint string `json:"s3LocationConstraint"` // Location constraint - must be set to match the Region.
    S3BucketAcl string `json:"s3BucketAcl"` // Canned ACL used when creating buckets.
    S3ListUrlEncode string `json:"s3ListUrlEncode" default:"unset"` // Whether to url encode listings: true/false/unset
    S3MemoryPoolUseMmap string `json:"s3MemoryPoolUseMmap" default:"false"` // Whether to use mmap buffers in internal memory pool.
    S3VersionAt string `json:"s3VersionAt" default:"off"` // Show file versions as they were at the specified time.
    S3ChunkSize string `json:"s3ChunkSize" default:"5Mi"` // Chunk size to use for uploading.
    S3SessionToken string `json:"s3SessionToken"` // An AWS session token.
    S3UploadConcurrency string `json:"s3UploadConcurrency" default:"4"` // Concurrency for multipart uploads.
    S3Decompress string `json:"s3Decompress" default:"false"` // If set this will decompress gzip encoded objects.
    S3Acl string `json:"s3Acl"` // Canned ACL used when creating buckets and storing or copying objects.
    S3MaxUploadParts string `json:"s3MaxUploadParts" default:"10000"` // Maximum number of parts in a multipart upload.
    S3NoHeadObject string `json:"s3NoHeadObject" default:"false"` // If set, do not do HEAD before GET when getting objects.
    S3NoHead string `json:"s3NoHead" default:"false"` // If set, don't HEAD uploaded objects to check integrity.
    S3DownloadUrl string `json:"s3DownloadUrl"` // Custom endpoint for downloads.
    S3MightGzip string `json:"s3MightGzip" default:"unset"` // Set this if the backend might gzip objects.
    S3Provider string `json:"s3Provider"` // Choose your S3 provider.
    S3SecretAccessKey string `json:"s3SecretAccessKey"` // AWS Secret Access Key (password).
    S3SseCustomerAlgorithm string `json:"s3SseCustomerAlgorithm"` // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
    S3V2Auth string `json:"s3V2Auth" default:"false"` // If true use v2 authentication.
    S3LeavePartsOnError string `json:"s3LeavePartsOnError" default:"false"` // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
    S3DisableHttp2 string `json:"s3DisableHttp2" default:"false"` // Disable usage of http2 for S3 backends.
    S3EnvAuth string `json:"s3EnvAuth" default:"false"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
    S3StorageClass string `json:"s3StorageClass"` // The storage class to use when storing new objects in S3.
    S3Profile string `json:"s3Profile"` // Profile to use in the shared credentials file.
    S3UseMultipartEtag string `json:"s3UseMultipartEtag" default:"unset"` // Whether to use ETag in multipart uploads for verification
    S3ServerSideEncryption string `json:"s3ServerSideEncryption"` // The server-side encryption algorithm used when storing this object in S3.
    S3SharedCredentialsFile string `json:"s3SharedCredentialsFile"` // Path to the shared credentials file.
    S3MemoryPoolFlushTime string `json:"s3MemoryPoolFlushTime" default:"1m0s"` // How often internal memory buffer pools will be flushed.
    S3Versions string `json:"s3Versions" default:"false"` // Include old versions in directory listings.
    S3RequesterPays string `json:"s3RequesterPays" default:"false"` // Enables requester pays option when interacting with S3 bucket.
    S3SseCustomerKey string `json:"s3SseCustomerKey"` // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
    S3SseCustomerKeyMd5 string `json:"s3SseCustomerKeyMd5"` // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
    SeafileUrl string `json:"seafileUrl"` // URL of seafile host to connect to.
    SeafileUser string `json:"seafileUser"` // User name (usually email address).
    SeafileLibrary string `json:"seafileLibrary"` // Name of the library.
    SeafileCreateLibrary string `json:"seafileCreateLibrary" default:"false"` // Should rclone create a library if it doesn't exist.
    SeafileAuthToken string `json:"seafileAuthToken"` // Authentication token.
    SeafilePass string `json:"seafilePass"` // Password.
    Seafile2fa string `json:"seafile2fa" default:"false"` // Two-factor authentication ('true' if the account has 2FA enabled).
    SeafileLibraryKey string `json:"seafileLibraryKey"` // Library password (for encrypted libraries only).
    SeafileEncoding string `json:"seafileEncoding" default:"Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8"` // The encoding for the backend.
    SftpSubsystem string `json:"sftpSubsystem" default:"sftp"` // Specifies the SSH2 subsystem on the remote host.
    SftpKeyPem string `json:"sftpKeyPem"` // Raw PEM-encoded private key.
    SftpKeyFilePass string `json:"sftpKeyFilePass"` // The passphrase to decrypt the PEM-encoded private key file.
    SftpSetModtime string `json:"sftpSetModtime" default:"true"` // Set the modified time on the remote if set.
    SftpHost string `json:"sftpHost"` // SSH host to connect to.
    SftpDisableHashcheck string `json:"sftpDisableHashcheck" default:"false"` // Disable the execution of SSH commands to determine if remote file hashing is available.
    SftpPathOverride string `json:"sftpPathOverride"` // Override path used by SSH shell commands.
    SftpSha1sumCommand string `json:"sftpSha1sumCommand"` // The command used to read sha1 hashes.
    SftpPort string `json:"sftpPort" default:"22"` // SSH port number.
    SftpKnownHostsFile string `json:"sftpKnownHostsFile"` // Optional path to known_hosts file.
    SftpUseInsecureCipher string `json:"sftpUseInsecureCipher" default:"false"` // Enable the use of insecure ciphers and key exchange methods.
    SftpMd5sumCommand string `json:"sftpMd5sumCommand"` // The command used to read md5 hashes.
    SftpSkipLinks string `json:"sftpSkipLinks" default:"false"` // Set to skip any symlinks and any other non regular files.
    SftpSetEnv string `json:"sftpSetEnv"` // Environment variables to pass to sftp and commands
    SftpCiphers string `json:"sftpCiphers"` // Space separated list of ciphers to be used for session encryption, ordered by preference.
    SftpPubkeyFile string `json:"sftpPubkeyFile"` // Optional path to public key file.
    SftpDisableConcurrentReads string `json:"sftpDisableConcurrentReads" default:"false"` // If set don't use concurrent reads.
    SftpIdleTimeout string `json:"sftpIdleTimeout" default:"1m0s"` // Max time before closing idle connections.
    SftpKeyFile string `json:"sftpKeyFile"` // Path to PEM-encoded private key file.
    SftpMacs string `json:"sftpMacs"` // Space separated list of MACs (message authentication code) algorithms, ordered by preference.
    SftpAskPassword string `json:"sftpAskPassword" default:"false"` // Allow asking for SFTP password when needed.
    SftpServerCommand string `json:"sftpServerCommand"` // Specifies the path or command to run a sftp server on the remote host.
    SftpConcurrency string `json:"sftpConcurrency" default:"64"` // The maximum number of outstanding requests for one file
    SftpKeyExchange string `json:"sftpKeyExchange"` // Space separated list of key exchange algorithms, ordered by preference.
    SftpUser string `json:"sftpUser" default:"shane"` // SSH username.
    SftpKeyUseAgent string `json:"sftpKeyUseAgent" default:"false"` // When set forces the usage of the ssh-agent.
    SftpShellType string `json:"sftpShellType"` // The type of SSH shell on remote server, if any.
    SftpUseFstat string `json:"sftpUseFstat" default:"false"` // If set use fstat instead of stat.
    SftpDisableConcurrentWrites string `json:"sftpDisableConcurrentWrites" default:"false"` // If set don't use concurrent writes.
    SftpChunkSize string `json:"sftpChunkSize" default:"32Ki"` // Upload and download chunk size.
    SftpPass string `json:"sftpPass"` // SSH password, leave blank to use ssh-agent.
    SharefileUploadCutoff string `json:"sharefileUploadCutoff" default:"128Mi"` // Cutoff for switching to multipart upload.
    SharefileRootFolderId string `json:"sharefileRootFolderId"` // ID of the root folder.
    SharefileChunkSize string `json:"sharefileChunkSize" default:"64Mi"` // Upload chunk size.
    SharefileEndpoint string `json:"sharefileEndpoint"` // Endpoint for API calls.
    SharefileEncoding string `json:"sharefileEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
    SiaApiUrl string `json:"siaApiUrl" default:"http://127.0.0.1:9980"` // Sia daemon API URL, like http://sia.daemon.host:9980.
    SiaApiPassword string `json:"siaApiPassword"` // Sia Daemon API Password.
    SiaUserAgent string `json:"siaUserAgent" default:"Sia-Agent"` // Siad User Agent
    SiaEncoding string `json:"siaEncoding" default:"Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    SmbIdleTimeout string `json:"smbIdleTimeout" default:"1m0s"` // Max time before closing idle connections.
    SmbHideSpecialShare string `json:"smbHideSpecialShare" default:"true"` // Hide special shares (e.g. print$) which users aren't supposed to access.
    SmbCaseInsensitive string `json:"smbCaseInsensitive" default:"true"` // Whether the server is configured to be case-insensitive.
    SmbEncoding string `json:"smbEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
    SmbHost string `json:"smbHost"` // SMB server hostname to connect to.
    SmbUser string `json:"smbUser" default:"shane"` // SMB username.
    SmbDomain string `json:"smbDomain" default:"WORKGROUP"` // Domain name for NTLM authentication.
    SmbSpn string `json:"smbSpn"` // Service principal name.
    SmbPort string `json:"smbPort" default:"445"` // SMB port number.
    SmbPass string `json:"smbPass"` // SMB password.
    StorjProvider string `json:"storjProvider" default:"existing"` // Choose an authentication method.
    StorjAccessGrant string `json:"storjAccessGrant"` // Access grant.
    StorjSatelliteAddress string `json:"storjSatelliteAddress" default:"us1.storj.io"` // Satellite address.
    StorjApiKey string `json:"storjApiKey"` // API key.
    StorjPassphrase string `json:"storjPassphrase"` // Encryption passphrase.
    TardigradeApiKey string `json:"tardigradeApiKey"` // API key.
    TardigradePassphrase string `json:"tardigradePassphrase"` // Encryption passphrase.
    TardigradeProvider string `json:"tardigradeProvider" default:"existing"` // Choose an authentication method.
    TardigradeAccessGrant string `json:"tardigradeAccessGrant"` // Access grant.
    TardigradeSatelliteAddress string `json:"tardigradeSatelliteAddress" default:"us1.storj.io"` // Satellite address.
    SugarsyncPrivateAccessKey string `json:"sugarsyncPrivateAccessKey"` // Sugarsync Private Access Key.
    SugarsyncHardDelete string `json:"sugarsyncHardDelete" default:"false"` // Permanently delete files if true
    SugarsyncAuthorizationExpiry string `json:"sugarsyncAuthorizationExpiry"` // Sugarsync authorization expiry.
    SugarsyncUser string `json:"sugarsyncUser"` // Sugarsync user.
    SugarsyncAppId string `json:"sugarsyncAppId"` // Sugarsync App ID.
    SugarsyncRefreshToken string `json:"sugarsyncRefreshToken"` // Sugarsync refresh token.
    SugarsyncAuthorization string `json:"sugarsyncAuthorization"` // Sugarsync authorization.
    SugarsyncRootId string `json:"sugarsyncRootId"` // Sugarsync root id.
    SugarsyncDeletedId string `json:"sugarsyncDeletedId"` // Sugarsync deleted folder id.
    SugarsyncEncoding string `json:"sugarsyncEncoding" default:"Slash,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    SugarsyncAccessKeyId string `json:"sugarsyncAccessKeyId"` // Sugarsync Access Key ID.
    SwiftNoLargeObjects string `json:"swiftNoLargeObjects" default:"false"` // Disable support for static and dynamic large objects
    SwiftEnvAuth string `json:"swiftEnvAuth" default:"false"` // Get swift credentials from environment variables in standard OpenStack form.
    SwiftAuthToken string `json:"swiftAuthToken"` // Auth Token from alternate authentication - optional (OS_AUTH_TOKEN).
    SwiftEndpointType string `json:"swiftEndpointType" default:"public"` // Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).
    SwiftNoChunk string `json:"swiftNoChunk" default:"false"` // Don't chunk files during streaming upload.
    SwiftApplicationCredentialName string `json:"swiftApplicationCredentialName"` // Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME).
    SwiftAuthVersion string `json:"swiftAuthVersion" default:"0"` // AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).
    SwiftUser string `json:"swiftUser"` // User name to log in (OS_USERNAME).
    SwiftTenant string `json:"swiftTenant"` // Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME).
    SwiftRegion string `json:"swiftRegion"` // Region name - optional (OS_REGION_NAME).
    SwiftStorageUrl string `json:"swiftStorageUrl"` // Storage URL - optional (OS_STORAGE_URL).
    SwiftApplicationCredentialSecret string `json:"swiftApplicationCredentialSecret"` // Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET).
    SwiftLeavePartsOnError string `json:"swiftLeavePartsOnError" default:"false"` // If true avoid calling abort upload on a failure.
    SwiftEncoding string `json:"swiftEncoding" default:"Slash,InvalidUtf8"` // The encoding for the backend.
    SwiftAuth string `json:"swiftAuth"` // Authentication URL for server (OS_AUTH_URL).
    SwiftUserId string `json:"swiftUserId"` // User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID).
    SwiftTenantId string `json:"swiftTenantId"` // Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID).
    SwiftApplicationCredentialId string `json:"swiftApplicationCredentialId"` // Application Credential ID (OS_APPLICATION_CREDENTIAL_ID).
    SwiftChunkSize string `json:"swiftChunkSize" default:"5Gi"` // Above this size files will be chunked into a _segments container.
    SwiftKey string `json:"swiftKey"` // API key or password (OS_PASSWORD).
    SwiftDomain string `json:"swiftDomain"` // User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME)
    SwiftTenantDomain string `json:"swiftTenantDomain"` // Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME).
    SwiftStoragePolicy string `json:"swiftStoragePolicy"` // The storage policy to use when creating a new container.
    UptoboxAccessToken string `json:"uptoboxAccessToken"` // Your access token.
    UptoboxEncoding string `json:"uptoboxEncoding" default:"Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot"` // The encoding for the backend.
    WebdavUser string `json:"webdavUser"` // User name.
    WebdavPass string `json:"webdavPass"` // Password.
    WebdavBearerToken string `json:"webdavBearerToken"` // Bearer token instead of user/pass (e.g. a Macaroon).
    WebdavBearerTokenCommand string `json:"webdavBearerTokenCommand"` // Command to run to get a bearer token.
    WebdavEncoding string `json:"webdavEncoding"` // The encoding for the backend.
    WebdavHeaders string `json:"webdavHeaders"` // Set HTTP headers for all transactions.
    WebdavUrl string `json:"webdavUrl"` // URL of http host to connect to.
    WebdavVendor string `json:"webdavVendor"` // Name of the WebDAV site/service/software you are using.
    YandexClientSecret string `json:"yandexClientSecret"` // OAuth Client Secret.
    YandexToken string `json:"yandexToken"` // OAuth Access Token as a JSON blob.
    YandexAuthUrl string `json:"yandexAuthUrl"` // Auth server URL.
    YandexTokenUrl string `json:"yandexTokenUrl"` // Token server url.
    YandexHardDelete string `json:"yandexHardDelete" default:"false"` // Delete files permanently rather than putting them into the trash.
    YandexEncoding string `json:"yandexEncoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
    YandexClientId string `json:"yandexClientId"` // OAuth Client Id.
    ZohoTokenUrl string `json:"zohoTokenUrl"` // Token server url.
    ZohoRegion string `json:"zohoRegion"` // Zoho region to connect to.
    ZohoEncoding string `json:"zohoEncoding" default:"Del,Ctl,InvalidUtf8"` // The encoding for the backend.
    ZohoClientId string `json:"zohoClientId"` // OAuth Client Id.
    ZohoClientSecret string `json:"zohoClientSecret"` // OAuth Client Secret.
    ZohoToken string `json:"zohoToken"` // OAuth Access Token as a JSON blob.
    ZohoAuthUrl string `json:"zohoAuthUrl"` // Auth server URL.
}
