// Code generated. DO NOT EDIT.
//
//lint:file-ignore U1000 Ignore all unused code, it's generated
package storage

import "github.com/data-preservation-programs/singularity/model"

type acdConfig struct {
	ClientId          string `json:"clientId"`                                          // OAuth Client Id.
	ClientSecret      string `json:"clientSecret"`                                      // OAuth Client Secret.
	Token             string `json:"token"`                                             // OAuth Access Token as a JSON blob.
	AuthUrl           string `json:"authUrl"`                                           // Auth server URL.
	TokenUrl          string `json:"tokenUrl"`                                          // Token server url.
	Checkpoint        string `json:"checkpoint"`                                        // Checkpoint for internal polling (debug).
	UploadWaitPerGb   string `default:"3m0s"                  json:"uploadWaitPerGb"`   // Additional time per GiB to wait after a failed complete upload to see if it appears.
	TemplinkThreshold string `default:"9Gi"                   json:"templinkThreshold"` // Files >= this size will be downloaded via their tempLink.
	Encoding          string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`          // The encoding for the backend.
}

type createAcdStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       acdConfig          `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateAcdStorage
// @Summary Create Acd storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createAcdStorageRequest true "Request body"
// @Router /storage/acd [post]
func createAcdStorage() {}

type azureblobConfig struct {
	Account                    string `json:"account"`                                                                              // Azure Storage Account Name.
	EnvAuth                    bool   `default:"false"                                           json:"envAuth"`                    // Read credentials from runtime (environment variables, CLI or MSI).
	Key                        string `json:"key"`                                                                                  // Storage Account Shared Key.
	SasUrl                     string `json:"sasUrl"`                                                                               // SAS URL for container level access only.
	Tenant                     string `json:"tenant"`                                                                               // ID of the service principal's tenant. Also called its directory ID.
	ClientId                   string `json:"clientId"`                                                                             // The ID of the client in use.
	ClientSecret               string `json:"clientSecret"`                                                                         // One of the service principal's client secrets
	ClientCertificatePath      string `json:"clientCertificatePath"`                                                                // Path to a PEM or PKCS12 certificate file including the private key.
	ClientCertificatePassword  string `json:"clientCertificatePassword"`                                                            // Password for the certificate file (optional).
	ClientSendCertificateChain bool   `default:"false"                                           json:"clientSendCertificateChain"` // Send the certificate chain when using certificate auth.
	Username                   string `json:"username"`                                                                             // User name (usually an email address)
	Password                   string `json:"password"`                                                                             // The user's password
	ServicePrincipalFile       string `json:"servicePrincipalFile"`                                                                 // Path to file containing credentials for use with a service principal.
	UseMsi                     bool   `default:"false"                                           json:"useMsi"`                     // Use a managed service identity to authenticate (only works in Azure).
	MsiObjectId                string `json:"msiObjectId"`                                                                          // Object ID of the user-assigned MSI to use, if any.
	MsiClientId                string `json:"msiClientId"`                                                                          // Object ID of the user-assigned MSI to use, if any.
	MsiMiResId                 string `json:"msiMiResId"`                                                                           // Azure resource ID of the user-assigned MSI to use, if any.
	UseEmulator                bool   `default:"false"                                           json:"useEmulator"`                // Uses local storage emulator if provided as 'true'.
	Endpoint                   string `json:"endpoint"`                                                                             // Endpoint for the service.
	UploadCutoff               string `json:"uploadCutoff"`                                                                         // Cutoff for switching to chunked upload (<= 256 MiB) (deprecated).
	ChunkSize                  string `default:"4Mi"                                             json:"chunkSize"`                  // Upload chunk size.
	UploadConcurrency          int    `default:"16"                                              json:"uploadConcurrency"`          // Concurrency for multipart uploads.
	ListChunk                  int    `default:"5000"                                            json:"listChunk"`                  // Size of blob list.
	AccessTier                 string `json:"accessTier"`                                                                           // Access tier of blob: hot, cool or archive.
	ArchiveTierDelete          bool   `default:"false"                                           json:"archiveTierDelete"`          // Delete archive tier blobs before overwriting.
	DisableChecksum            bool   `default:"false"                                           json:"disableChecksum"`            // Don't store MD5 checksum with object metadata.
	MemoryPoolFlushTime        string `default:"1m0s"                                            json:"memoryPoolFlushTime"`        // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap          bool   `default:"false"                                           json:"memoryPoolUseMmap"`          // Whether to use mmap buffers in internal memory pool.
	Encoding                   string `default:"Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8" json:"encoding"`                   // The encoding for the backend.
	PublicAccess               string `example:""                                                json:"publicAccess"`               // Public access level of a container: blob or container.
	NoCheckContainer           bool   `default:"false"                                           json:"noCheckContainer"`           // If set, don't attempt to check the container exists or create it.
	NoHeadObject               bool   `default:"false"                                           json:"noHeadObject"`               // If set, do not do HEAD before GET when getting objects.
}

type createAzureblobStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       azureblobConfig    `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateAzureblobStorage
// @Summary Create Azureblob storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createAzureblobStorageRequest true "Request body"
// @Router /storage/azureblob [post]
func createAzureblobStorage() {}

type b2Config struct {
	Account              string `json:"account"`                                                                // Account ID or Application Key ID.
	Key                  string `json:"key"`                                                                    // Application Key.
	Endpoint             string `json:"endpoint"`                                                               // Endpoint for the service.
	TestMode             string `json:"testMode"`                                                               // A flag string for X-Bz-Test-Mode header for debugging.
	Versions             bool   `default:"false"                                   json:"versions"`             // Include old versions in directory listings.
	VersionAt            string `default:"off"                                     json:"versionAt"`            // Show file versions as they were at the specified time.
	HardDelete           bool   `default:"false"                                   json:"hardDelete"`           // Permanently delete files on remote removal, otherwise hide files.
	UploadCutoff         string `default:"200Mi"                                   json:"uploadCutoff"`         // Cutoff for switching to chunked upload.
	CopyCutoff           string `default:"4Gi"                                     json:"copyCutoff"`           // Cutoff for switching to multipart copy.
	ChunkSize            string `default:"96Mi"                                    json:"chunkSize"`            // Upload chunk size.
	DisableChecksum      bool   `default:"false"                                   json:"disableChecksum"`      // Disable checksums for large (> upload cutoff) files.
	DownloadUrl          string `json:"downloadUrl"`                                                            // Custom endpoint for downloads.
	DownloadAuthDuration string `default:"1w"                                      json:"downloadAuthDuration"` // Time before the authorization token will expire in s or suffix ms|s|m|h|d.
	MemoryPoolFlushTime  string `default:"1m0s"                                    json:"memoryPoolFlushTime"`  // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap    bool   `default:"false"                                   json:"memoryPoolUseMmap"`    // Whether to use mmap buffers in internal memory pool.
	Encoding             string `default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot" json:"encoding"`             // The encoding for the backend.
}

type createB2StorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       b2Config           `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateB2Storage
// @Summary Create B2 storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createB2StorageRequest true "Request body"
// @Router /storage/b2 [post]
func createB2Storage() {}

type boxConfig struct {
	ClientId      string `json:"clientId"`                                                                                     // OAuth Client Id.
	ClientSecret  string `json:"clientSecret"`                                                                                 // OAuth Client Secret.
	Token         string `json:"token"`                                                                                        // OAuth Access Token as a JSON blob.
	AuthUrl       string `json:"authUrl"`                                                                                      // Auth server URL.
	TokenUrl      string `json:"tokenUrl"`                                                                                     // Token server url.
	RootFolderId  string `default:"0"                                                  json:"rootFolderId"`                    // Fill in for rclone to use a non root folder as its starting point.
	BoxConfigFile string `json:"boxConfigFile"`                                                                                // Box App config.json location
	AccessToken   string `json:"accessToken"`                                                                                  // Box App Primary Access Token
	BoxSubType    string `default:"user"                                               example:"user"       json:"boxSubType"` //
	UploadCutoff  string `default:"50Mi"                                               json:"uploadCutoff"`                    // Cutoff for switching to multipart upload (>= 50 MiB).
	CommitRetries int    `default:"100"                                                json:"commitRetries"`                   // Max number of times to try committing a multipart file.
	ListChunk     int    `default:"1000"                                               json:"listChunk"`                       // Size of listing chunk 1-1000.
	OwnedBy       string `json:"ownedBy"`                                                                                      // Only show items owned by the login (email address) passed in.
	Encoding      string `default:"Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot" json:"encoding"`                        // The encoding for the backend.
}

type createBoxStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       boxConfig          `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateBoxStorage
// @Summary Create Box storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createBoxStorageRequest true "Request body"
// @Router /storage/box [post]
func createBoxStorage() {}

type driveConfig struct {
	ClientId                  string `json:"clientId"`                                                 // Google Application Client Id
	ClientSecret              string `json:"clientSecret"`                                             // OAuth Client Secret.
	Token                     string `json:"token"`                                                    // OAuth Access Token as a JSON blob.
	AuthUrl                   string `json:"authUrl"`                                                  // Auth server URL.
	TokenUrl                  string `json:"tokenUrl"`                                                 // Token server url.
	Scope                     string `example:"drive"                  json:"scope"`                   // Scope that rclone should use when requesting access from drive.
	RootFolderId              string `json:"rootFolderId"`                                             // ID of the root folder.
	ServiceAccountFile        string `json:"serviceAccountFile"`                                       // Service Account Credentials JSON file path.
	ServiceAccountCredentials string `json:"serviceAccountCredentials"`                                // Service Account Credentials JSON blob.
	TeamDrive                 string `json:"teamDrive"`                                                // ID of the Shared Drive (Team Drive).
	AuthOwnerOnly             bool   `default:"false"                  json:"authOwnerOnly"`           // Only consider files owned by the authenticated user.
	UseTrash                  bool   `default:"true"                   json:"useTrash"`                // Send files to the trash instead of deleting permanently.
	CopyShortcutContent       bool   `default:"false"                  json:"copyShortcutContent"`     // Server side copy contents of shortcuts instead of the shortcut.
	SkipGdocs                 bool   `default:"false"                  json:"skipGdocs"`               // Skip google documents in all listings.
	SkipChecksumGphotos       bool   `default:"false"                  json:"skipChecksumGphotos"`     // Skip MD5 checksum on Google photos and videos only.
	SharedWithMe              bool   `default:"false"                  json:"sharedWithMe"`            // Only show files that are shared with me.
	TrashedOnly               bool   `default:"false"                  json:"trashedOnly"`             // Only show files that are in the trash.
	StarredOnly               bool   `default:"false"                  json:"starredOnly"`             // Only show files that are starred.
	Formats                   string `json:"formats"`                                                  // Deprecated: See export_formats.
	ExportFormats             string `default:"docx,xlsx,pptx,svg"     json:"exportFormats"`           // Comma separated list of preferred formats for downloading Google docs.
	ImportFormats             string `json:"importFormats"`                                            // Comma separated list of preferred formats for uploading Google docs.
	AllowImportNameChange     bool   `default:"false"                  json:"allowImportNameChange"`   // Allow the filetype to change when uploading Google docs.
	UseCreatedDate            bool   `default:"false"                  json:"useCreatedDate"`          // Use file created date instead of modified date.
	UseSharedDate             bool   `default:"false"                  json:"useSharedDate"`           // Use date file was shared instead of modified date.
	ListChunk                 int    `default:"1000"                   json:"listChunk"`               // Size of listing chunk 100-1000, 0 to disable.
	Impersonate               string `json:"impersonate"`                                              // Impersonate this user when using a service account.
	AlternateExport           bool   `default:"false"                  json:"alternateExport"`         // Deprecated: No longer needed.
	UploadCutoff              string `default:"8Mi"                    json:"uploadCutoff"`            // Cutoff for switching to chunked upload.
	ChunkSize                 string `default:"8Mi"                    json:"chunkSize"`               // Upload chunk size.
	AcknowledgeAbuse          bool   `default:"false"                  json:"acknowledgeAbuse"`        // Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
	KeepRevisionForever       bool   `default:"false"                  json:"keepRevisionForever"`     // Keep new head revision of each file forever.
	SizeAsQuota               bool   `default:"false"                  json:"sizeAsQuota"`             // Show sizes as storage quota usage, not actual size.
	V2DownloadMinSize         string `default:"off"                    json:"v2DownloadMinSize"`       // If Object's are greater, use drive v2 API to download.
	PacerMinSleep             string `default:"100ms"                  json:"pacerMinSleep"`           // Minimum time to sleep between API calls.
	PacerBurst                int    `default:"100"                    json:"pacerBurst"`              // Number of API calls to allow without sleeping.
	ServerSideAcrossConfigs   bool   `default:"false"                  json:"serverSideAcrossConfigs"` // Allow server-side operations (e.g. copy) to work across different drive configs.
	DisableHttp2              bool   `default:"true"                   json:"disableHttp2"`            // Disable drive using http2.
	StopOnUploadLimit         bool   `default:"false"                  json:"stopOnUploadLimit"`       // Make upload limit errors be fatal.
	StopOnDownloadLimit       bool   `default:"false"                  json:"stopOnDownloadLimit"`     // Make download limit errors be fatal.
	SkipShortcuts             bool   `default:"false"                  json:"skipShortcuts"`           // If set skip shortcut files.
	SkipDanglingShortcuts     bool   `default:"false"                  json:"skipDanglingShortcuts"`   // If set skip dangling shortcut files.
	ResourceKey               string `json:"resourceKey"`                                              // Resource key for accessing a link-shared file.
	Encoding                  string `default:"InvalidUtf8"            json:"encoding"`                // The encoding for the backend.
}

type createDriveStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       driveConfig        `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateDriveStorage
// @Summary Create Drive storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createDriveStorageRequest true "Request body"
// @Router /storage/drive [post]
func createDriveStorage() {}

type dropboxConfig struct {
	ClientId           string `json:"clientId"`                                                                    // OAuth Client Id.
	ClientSecret       string `json:"clientSecret"`                                                                // OAuth Client Secret.
	Token              string `json:"token"`                                                                       // OAuth Access Token as a JSON blob.
	AuthUrl            string `json:"authUrl"`                                                                     // Auth server URL.
	TokenUrl           string `json:"tokenUrl"`                                                                    // Token server url.
	ChunkSize          string `default:"48Mi"                                           json:"chunkSize"`          // Upload chunk size (< 150Mi).
	Impersonate        string `json:"impersonate"`                                                                 // Impersonate this user when using a business account.
	SharedFiles        bool   `default:"false"                                          json:"sharedFiles"`        // Instructs rclone to work on individual shared files.
	SharedFolders      bool   `default:"false"                                          json:"sharedFolders"`      // Instructs rclone to work on shared folders.
	BatchMode          string `default:"sync"                                           json:"batchMode"`          // Upload file batching sync|async|off.
	BatchSize          int    `default:"0"                                              json:"batchSize"`          // Max number of files in upload batch.
	BatchTimeout       string `default:"0s"                                             json:"batchTimeout"`       // Max time to allow an idle upload batch before uploading.
	BatchCommitTimeout string `default:"10m0s"                                          json:"batchCommitTimeout"` // Max time to wait for a batch to finish committing
	Encoding           string `default:"Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot" json:"encoding"`           // The encoding for the backend.
}

type createDropboxStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       dropboxConfig      `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateDropboxStorage
// @Summary Create Dropbox storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createDropboxStorageRequest true "Request body"
// @Router /storage/dropbox [post]
func createDropboxStorage() {}

type fichierConfig struct {
	ApiKey         string `json:"apiKey"`                                                                                                                        // Your API Key, get it from https://1fichier.com/console/params.pl.
	SharedFolder   string `json:"sharedFolder"`                                                                                                                  // If you want to download a shared folder, add this parameter.
	FilePassword   string `json:"filePassword"`                                                                                                                  // If you want to download a shared file that is password protected, add this parameter.
	FolderPassword string `json:"folderPassword"`                                                                                                                // If you want to list the files in a shared folder that is password protected, add this parameter.
	Encoding       string `default:"Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot" json:"encoding"` // The encoding for the backend.
}

type createFichierStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       fichierConfig      `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateFichierStorage
// @Summary Create Fichier storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createFichierStorageRequest true "Request body"
// @Router /storage/fichier [post]
func createFichierStorage() {}

type filefabricConfig struct {
	Url            string `example:"https://storagemadeeasy.com"   json:"url"`      // URL of the Enterprise File Fabric to connect to.
	RootFolderId   string `json:"rootFolderId"`                                     // ID of the root folder.
	PermanentToken string `json:"permanentToken"`                                   // Permanent Authentication Token.
	Token          string `json:"token"`                                            // Session Token.
	TokenExpiry    string `json:"tokenExpiry"`                                      // Token expiry time.
	Version        string `json:"version"`                                          // Version read from the file fabric.
	Encoding       string `default:"Slash,Del,Ctl,InvalidUtf8,Dot" json:"encoding"` // The encoding for the backend.
}

type createFilefabricStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       filefabricConfig   `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateFilefabricStorage
// @Summary Create Filefabric storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createFilefabricStorageRequest true "Request body"
// @Router /storage/filefabric [post]
func createFilefabricStorage() {}

type ftpConfig struct {
	Host               string `json:"host"`                                                                             // FTP host to connect to.
	User               string `default:"$USER"                        json:"user"`                                      // FTP username.
	Port               int    `default:"21"                           json:"port"`                                      // FTP port number.
	Pass               string `json:"pass"`                                                                             // FTP password.
	Tls                bool   `default:"false"                        json:"tls"`                                       // Use Implicit FTPS (FTP over TLS).
	ExplicitTls        bool   `default:"false"                        json:"explicitTls"`                               // Use Explicit FTPS (FTP over TLS).
	Concurrency        int    `default:"0"                            json:"concurrency"`                               // Maximum number of FTP simultaneous connections, 0 for unlimited.
	NoCheckCertificate bool   `default:"false"                        json:"noCheckCertificate"`                        // Do not verify the TLS certificate of the server.
	DisableEpsv        bool   `default:"false"                        json:"disableEpsv"`                               // Disable using EPSV even if server advertises support.
	DisableMlsd        bool   `default:"false"                        json:"disableMlsd"`                               // Disable using MLSD even if server advertises support.
	DisableUtf8        bool   `default:"false"                        json:"disableUtf8"`                               // Disable using UTF-8 even if server advertises support.
	WritingMdtm        bool   `default:"false"                        json:"writingMdtm"`                               // Use MDTM to set modification time (VsFtpd quirk)
	ForceListHidden    bool   `default:"false"                        json:"forceListHidden"`                           // Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.
	IdleTimeout        string `default:"1m0s"                         json:"idleTimeout"`                               // Max time before closing idle connections.
	CloseTimeout       string `default:"1m0s"                         json:"closeTimeout"`                              // Maximum time to wait for a response to close.
	TlsCacheSize       int    `default:"32"                           json:"tlsCacheSize"`                              // Size of TLS session cache for all control and data connections.
	DisableTls13       bool   `default:"false"                        json:"disableTls13"`                              // Disable TLS 1.3 (workaround for FTP servers with buggy TLS)
	ShutTimeout        string `default:"1m0s"                         json:"shutTimeout"`                               // Maximum time to wait for data connection closing status.
	AskPassword        bool   `default:"false"                        json:"askPassword"`                               // Allow asking for FTP password when needed.
	Encoding           string `default:"Slash,Del,Ctl,RightSpace,Dot" example:"Asterisk,Ctl,Dot,Slash" json:"encoding"` // The encoding for the backend.
}

type createFtpStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       ftpConfig          `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateFtpStorage
// @Summary Create Ftp storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createFtpStorageRequest true "Request body"
// @Router /storage/ftp [post]
func createFtpStorage() {}

type gcsConfig struct {
	ClientId                  string `json:"clientId"`                                                             // OAuth Client Id.
	ClientSecret              string `json:"clientSecret"`                                                         // OAuth Client Secret.
	Token                     string `json:"token"`                                                                // OAuth Access Token as a JSON blob.
	AuthUrl                   string `json:"authUrl"`                                                              // Auth server URL.
	TokenUrl                  string `json:"tokenUrl"`                                                             // Token server url.
	ProjectNumber             string `json:"projectNumber"`                                                        // Project number.
	ServiceAccountFile        string `json:"serviceAccountFile"`                                                   // Service Account Credentials JSON file path.
	ServiceAccountCredentials string `json:"serviceAccountCredentials"`                                            // Service Account Credentials JSON blob.
	Anonymous                 bool   `default:"false"                      json:"anonymous"`                       // Access public buckets and objects without credentials.
	ObjectAcl                 string `example:"authenticatedRead"          json:"objectAcl"`                       // Access Control List for new objects.
	BucketAcl                 string `example:"authenticatedRead"          json:"bucketAcl"`                       // Access Control List for new buckets.
	BucketPolicyOnly          bool   `default:"false"                      json:"bucketPolicyOnly"`                // Access checks should use bucket-level IAM policies.
	Location                  string `example:""                           json:"location"`                        // Location for the newly created buckets.
	StorageClass              string `example:""                           json:"storageClass"`                    // The storage class to use when storing objects in Google Cloud Storage.
	NoCheckBucket             bool   `default:"false"                      json:"noCheckBucket"`                   // If set, don't attempt to check the bucket exists or create it.
	Decompress                bool   `default:"false"                      json:"decompress"`                      // If set this will decompress gzip encoded objects.
	Endpoint                  string `json:"endpoint"`                                                             // Endpoint for the service.
	Encoding                  string `default:"Slash,CrLf,InvalidUtf8,Dot" json:"encoding"`                        // The encoding for the backend.
	EnvAuth                   bool   `default:"false"                      example:"false"         json:"envAuth"` // Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
}

type createGcsStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       gcsConfig          `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateGcsStorage
// @Summary Create Gcs storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createGcsStorageRequest true "Request body"
// @Router /storage/gcs [post]
func createGcsStorage() {}

type gphotosConfig struct {
	ClientId        string `json:"clientId"`                                             // OAuth Client Id.
	ClientSecret    string `json:"clientSecret"`                                         // OAuth Client Secret.
	Token           string `json:"token"`                                                // OAuth Access Token as a JSON blob.
	AuthUrl         string `json:"authUrl"`                                              // Auth server URL.
	TokenUrl        string `json:"tokenUrl"`                                             // Token server url.
	ReadOnly        bool   `default:"false"                      json:"readOnly"`        // Set to make the Google Photos backend read only.
	ReadSize        bool   `default:"false"                      json:"readSize"`        // Set to read the size of media items.
	StartYear       int    `default:"2000"                       json:"startYear"`       // Year limits the photos to be downloaded to those which are uploaded after the given year.
	IncludeArchived bool   `default:"false"                      json:"includeArchived"` // Also view and download archived media.
	Encoding        string `default:"Slash,CrLf,InvalidUtf8,Dot" json:"encoding"`        // The encoding for the backend.
}

type createGphotosStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       gphotosConfig      `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateGphotosStorage
// @Summary Create Gphotos storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createGphotosStorageRequest true "Request body"
// @Router /storage/gphotos [post]
func createGphotosStorage() {}

type hdfsConfig struct {
	Namenode               string `json:"namenode"`                                                             // Hadoop name node and port.
	Username               string `example:"root"                                json:"username"`               // Hadoop user name.
	ServicePrincipalName   string `json:"servicePrincipalName"`                                                 // Kerberos service principal name for the namenode.
	DataTransferProtection string `example:"privacy"                             json:"dataTransferProtection"` // Kerberos data transfer protection: authentication|integrity|privacy.
	Encoding               string `default:"Slash,Colon,Del,Ctl,InvalidUtf8,Dot" json:"encoding"`               // The encoding for the backend.
}

type createHdfsStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       hdfsConfig         `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateHdfsStorage
// @Summary Create Hdfs storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createHdfsStorageRequest true "Request body"
// @Router /storage/hdfs [post]
func createHdfsStorage() {}

type hidriveConfig struct {
	ClientId                   string `json:"clientId"`                                                                                   // OAuth Client Id.
	ClientSecret               string `json:"clientSecret"`                                                                               // OAuth Client Secret.
	Token                      string `json:"token"`                                                                                      // OAuth Access Token as a JSON blob.
	AuthUrl                    string `json:"authUrl"`                                                                                    // Auth server URL.
	TokenUrl                   string `json:"tokenUrl"`                                                                                   // Token server url.
	ScopeAccess                string `default:"rw"                                 example:"rw"                      json:"scopeAccess"` // Access permissions that rclone should use when requesting access from HiDrive.
	ScopeRole                  string `default:"user"                               example:"user"                    json:"scopeRole"`   // User-level that rclone should use when requesting access from HiDrive.
	RootPrefix                 string `default:"/"                                  example:"/"                       json:"rootPrefix"`  // The root/parent folder for all paths.
	Endpoint                   string `default:"https://api.hidrive.strato.com/2.1" json:"endpoint"`                                      // Endpoint for the service.
	DisableFetchingMemberCount bool   `default:"false"                              json:"disableFetchingMemberCount"`                    // Do not fetch number of objects in directories unless it is absolutely necessary.
	ChunkSize                  string `default:"48Mi"                               json:"chunkSize"`                                     // Chunksize for chunked uploads.
	UploadCutoff               string `default:"96Mi"                               json:"uploadCutoff"`                                  // Cutoff/Threshold for chunked uploads.
	UploadConcurrency          int    `default:"4"                                  json:"uploadConcurrency"`                             // Concurrency for chunked uploads.
	Encoding                   string `default:"Slash,Dot"                          json:"encoding"`                                      // The encoding for the backend.
}

type createHidriveStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       hidriveConfig      `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateHidriveStorage
// @Summary Create Hidrive storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createHidriveStorageRequest true "Request body"
// @Router /storage/hidrive [post]
func createHidriveStorage() {}

type httpConfig struct {
	Url     string `json:"url"`                     // URL of HTTP host to connect to.
	Headers string `json:"headers"`                 // Set HTTP headers for all transactions.
	NoSlash bool   `default:"false" json:"noSlash"` // Set this if the site doesn't end directories with /.
	NoHead  bool   `default:"false" json:"noHead"`  // Don't use HEAD requests.
}

type createHttpStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       httpConfig         `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateHttpStorage
// @Summary Create Http storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createHttpStorageRequest true "Request body"
// @Router /storage/http [post]
func createHttpStorage() {}

type internetarchiveConfig struct {
	AccessKeyId     string `json:"accessKeyId"`                                                       // IAS3 Access Key.
	SecretAccessKey string `json:"secretAccessKey"`                                                   // IAS3 Secret Key (password).
	Endpoint        string `default:"https://s3.us.archive.org"               json:"endpoint"`        // IAS3 Endpoint.
	FrontEndpoint   string `default:"https://archive.org"                     json:"frontEndpoint"`   // Host of InternetArchive Frontend.
	DisableChecksum bool   `default:"true"                                    json:"disableChecksum"` // Don't ask the server to test against MD5 checksum calculated by rclone.
	WaitArchive     string `default:"0s"                                      json:"waitArchive"`     // Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
	Encoding        string `default:"Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot" json:"encoding"`        // The encoding for the backend.
}

type createInternetarchiveStorageRequest struct {
	Name         string                `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string                `json:"path"`                      // Path of the storage
	Config       internetarchiveConfig `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig    `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateInternetarchiveStorage
// @Summary Create Internetarchive storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createInternetarchiveStorageRequest true "Request body"
// @Router /storage/internetarchive [post]
func createInternetarchiveStorage() {}

type jottacloudConfig struct {
	Md5MemoryLimit    string `default:"10Mi"                                                                        json:"md5MemoryLimit"`    // Files bigger than this will be cached on disk to calculate the MD5 if required.
	TrashedOnly       bool   `default:"false"                                                                       json:"trashedOnly"`       // Only show files that are in the trash.
	HardDelete        bool   `default:"false"                                                                       json:"hardDelete"`        // Delete files permanently rather than putting them into the trash.
	UploadResumeLimit string `default:"10Mi"                                                                        json:"uploadResumeLimit"` // Files bigger than this can be resumed if the upload fail's.
	NoVersions        bool   `default:"false"                                                                       json:"noVersions"`        // Avoid server side versioning by deleting files and recreating files instead of overwriting them.
	Encoding          string `default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot" json:"encoding"`          // The encoding for the backend.
}

type createJottacloudStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       jottacloudConfig   `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateJottacloudStorage
// @Summary Create Jottacloud storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createJottacloudStorageRequest true "Request body"
// @Router /storage/jottacloud [post]
func createJottacloudStorage() {}

type koofrDigistorageConfig struct {
	Mountid  string `json:"mountid"`                                                    // Mount ID of the mount to use.
	Setmtime bool   `default:"true"                                    json:"setmtime"` // Does the backend support setting modification time.
	User     string `json:"user"`                                                       // Your user name.
	Password string `json:"password"`                                                   // Your password for rclone (generate one at https://storage.rcs-rds.ro/app/admin/preferences/password).
	Encoding string `default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot" json:"encoding"` // The encoding for the backend.
}

type createKoofrDigistorageStorageRequest struct {
	Name         string                 `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string                 `json:"path"`                      // Path of the storage
	Config       koofrDigistorageConfig `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig     `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateKoofrDigistorageStorage
// @Summary Create Koofr storage with digistorage - Digi Storage, https://storage.rcs-rds.ro/
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createKoofrDigistorageStorageRequest true "Request body"
// @Router /storage/koofr/digistorage [post]
func createKoofrDigistorageStorage() {}

type koofrKoofrConfig struct {
	Mountid  string `json:"mountid"`                                                    // Mount ID of the mount to use.
	Setmtime bool   `default:"true"                                    json:"setmtime"` // Does the backend support setting modification time.
	User     string `json:"user"`                                                       // Your user name.
	Password string `json:"password"`                                                   // Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password).
	Encoding string `default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot" json:"encoding"` // The encoding for the backend.
}

type createKoofrKoofrStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       koofrKoofrConfig   `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateKoofrKoofrStorage
// @Summary Create Koofr storage with koofr - Koofr, https://app.koofr.net/
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createKoofrKoofrStorageRequest true "Request body"
// @Router /storage/koofr/koofr [post]
func createKoofrKoofrStorage() {}

type koofrOtherConfig struct {
	Endpoint string `json:"endpoint"`                                                   // The Koofr API endpoint to use.
	Mountid  string `json:"mountid"`                                                    // Mount ID of the mount to use.
	Setmtime bool   `default:"true"                                    json:"setmtime"` // Does the backend support setting modification time.
	User     string `json:"user"`                                                       // Your user name.
	Password string `json:"password"`                                                   // Your password for rclone (generate one at your service's settings page).
	Encoding string `default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot" json:"encoding"` // The encoding for the backend.
}

type createKoofrOtherStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       koofrOtherConfig   `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateKoofrOtherStorage
// @Summary Create Koofr storage with other - Any other Koofr API compatible storage service
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createKoofrOtherStorageRequest true "Request body"
// @Router /storage/koofr/other [post]
func createKoofrOtherStorage() {}

type localConfig struct {
	Nounc                bool   `default:"false"     example:"true"              json:"nounc"` // Disable UNC (long path names) conversion on Windows.
	CopyLinks            bool   `default:"false"     json:"copyLinks"`                         // Follow symlinks and copy the pointed to item.
	Links                bool   `default:"false"     json:"links"`                             // Translate symlinks to/from regular files with a '.rclonelink' extension.
	SkipLinks            bool   `default:"false"     json:"skipLinks"`                         // Don't warn about skipped symlinks.
	ZeroSizeLinks        bool   `default:"false"     json:"zeroSizeLinks"`                     // Assume the Stat size of links is zero (and read them instead) (deprecated).
	UnicodeNormalization bool   `default:"false"     json:"unicodeNormalization"`              // Apply unicode NFC normalization to paths and filenames.
	NoCheckUpdated       bool   `default:"false"     json:"noCheckUpdated"`                    // Don't check to see if the files change during upload.
	OneFileSystem        bool   `default:"false"     json:"oneFileSystem"`                     // Don't cross filesystem boundaries (unix/macOS only).
	CaseSensitive        bool   `default:"false"     json:"caseSensitive"`                     // Force the filesystem to report itself as case sensitive.
	CaseInsensitive      bool   `default:"false"     json:"caseInsensitive"`                   // Force the filesystem to report itself as case insensitive.
	NoPreallocate        bool   `default:"false"     json:"noPreallocate"`                     // Disable preallocation of disk space for transferred files.
	NoSparse             bool   `default:"false"     json:"noSparse"`                          // Disable sparse files for multi-thread downloads.
	NoSetModtime         bool   `default:"false"     json:"noSetModtime"`                      // Disable setting modtime.
	Encoding             string `default:"Slash,Dot" json:"encoding"`                          // The encoding for the backend.
}

type createLocalStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       localConfig        `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateLocalStorage
// @Summary Create Local storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createLocalStorageRequest true "Request body"
// @Router /storage/local [post]
func createLocalStorage() {}

type mailruConfig struct {
	User                string `json:"user"`                                                                                                                                // User name (usually email).
	Pass                string `json:"pass"`                                                                                                                                // Password.
	SpeedupEnable       bool   `default:"true"                                                                                  example:"true"  json:"speedupEnable"`       // Skip full upload if there is another file with same data hash.
	SpeedupFilePatterns string `default:"*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf"                                        example:""      json:"speedupFilePatterns"` // Comma separated list of file name patterns eligible for speedup (put by hash).
	SpeedupMaxDisk      string `default:"3Gi"                                                                                   example:"0"     json:"speedupMaxDisk"`      // This option allows you to disable speedup (put by hash) for large files.
	SpeedupMaxMemory    string `default:"32Mi"                                                                                  example:"0"     json:"speedupMaxMemory"`    // Files larger than the size given below will always be hashed on disk.
	CheckHash           bool   `default:"true"                                                                                  example:"true"  json:"checkHash"`           // What should copy do if file checksum is mismatched or invalid.
	UserAgent           string `json:"userAgent"`                                                                                                                           // HTTP user agent used internally by client.
	Quirks              string `json:"quirks"`                                                                                                                              // Comma separated list of internal maintenance flags.
	Encoding            string `default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot" json:"encoding"`                            // The encoding for the backend.
}

type createMailruStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       mailruConfig       `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateMailruStorage
// @Summary Create Mailru storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createMailruStorageRequest true "Request body"
// @Router /storage/mailru [post]
func createMailruStorage() {}

type megaConfig struct {
	User       string `json:"user"`                                       // User name.
	Pass       string `json:"pass"`                                       // Password.
	Debug      bool   `default:"false"                 json:"debug"`      // Output more debug from Mega.
	HardDelete bool   `default:"false"                 json:"hardDelete"` // Delete files permanently rather than putting them into the trash.
	UseHttps   bool   `default:"false"                 json:"useHttps"`   // Use HTTPS for transfers.
	Encoding   string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`   // The encoding for the backend.
}

type createMegaStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       megaConfig         `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateMegaStorage
// @Summary Create Mega storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createMegaStorageRequest true "Request body"
// @Router /storage/mega [post]
func createMegaStorage() {}

type netstorageConfig struct {
	Protocol string `default:"https" example:"http" json:"protocol"` // Select between HTTP or HTTPS protocol.
	Host     string `json:"host"`                                    // Domain+path of NetStorage host to connect to.
	Account  string `json:"account"`                                 // Set the NetStorage account name
	Secret   string `json:"secret"`                                  // Set the NetStorage account secret/G2O key for authentication.
}

type createNetstorageStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       netstorageConfig   `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateNetstorageStorage
// @Summary Create Netstorage storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createNetstorageStorageRequest true "Request body"
// @Router /storage/netstorage [post]
func createNetstorageStorage() {}

type onedriveConfig struct {
	ClientId                string `json:"clientId"`                                                                                                                                                                                                                                                      // OAuth Client Id.
	ClientSecret            string `json:"clientSecret"`                                                                                                                                                                                                                                                  // OAuth Client Secret.
	Token                   string `json:"token"`                                                                                                                                                                                                                                                         // OAuth Access Token as a JSON blob.
	AuthUrl                 string `json:"authUrl"`                                                                                                                                                                                                                                                       // Auth server URL.
	TokenUrl                string `json:"tokenUrl"`                                                                                                                                                                                                                                                      // Token server url.
	Region                  string `default:"global"                                                                                                                           example:"global"                                                                                      json:"region"`       // Choose national cloud region for OneDrive.
	ChunkSize               string `default:"10Mi"                                                                                                                             json:"chunkSize"`                                                                                                          // Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
	DriveId                 string `json:"driveId"`                                                                                                                                                                                                                                                       // The ID of the drive to use.
	DriveType               string `json:"driveType"`                                                                                                                                                                                                                                                     // The type of the drive (personal | business | documentLibrary).
	RootFolderId            string `json:"rootFolderId"`                                                                                                                                                                                                                                                  // ID of the root folder.
	AccessScopes            string `default:"Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"                                      example:"Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access" json:"accessScopes"` // Set scopes to be requested by rclone.
	DisableSitePermission   bool   `default:"false"                                                                                                                            json:"disableSitePermission"`                                                                                              // Disable the request for Sites.Read.All permission.
	ExposeOnenoteFiles      bool   `default:"false"                                                                                                                            json:"exposeOnenoteFiles"`                                                                                                 // Set to make OneNote files show up in directory listings.
	ServerSideAcrossConfigs bool   `default:"false"                                                                                                                            json:"serverSideAcrossConfigs"`                                                                                            // Allow server-side operations (e.g. copy) to work across different onedrive configs.
	ListChunk               int    `default:"1000"                                                                                                                             json:"listChunk"`                                                                                                          // Size of listing chunk.
	NoVersions              bool   `default:"false"                                                                                                                            json:"noVersions"`                                                                                                         // Remove all versions on modifying operations.
	LinkScope               string `default:"anonymous"                                                                                                                        example:"anonymous"                                                                                   json:"linkScope"`    // Set the scope of the links created by the link command.
	LinkType                string `default:"view"                                                                                                                             example:"view"                                                                                        json:"linkType"`     // Set the type of the links created by the link command.
	LinkPassword            string `json:"linkPassword"`                                                                                                                                                                                                                                                  // Set the password for links created by the link command.
	HashType                string `default:"auto"                                                                                                                             example:"auto"                                                                                        json:"hashType"`     // Specify the hash in use for the backend.
	Encoding                string `default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot" json:"encoding"`                                                                                                           // The encoding for the backend.
}

type createOnedriveStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       onedriveConfig     `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateOnedriveStorage
// @Summary Create Onedrive storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createOnedriveStorageRequest true "Request body"
// @Router /storage/onedrive [post]
func createOnedriveStorage() {}

type oosEnv_authConfig struct {
	Namespace            string `json:"namespace"`                                                               // Object storage namespace
	Compartment          string `json:"compartment"`                                                             // Object storage compartment OCID
	Region               string `json:"region"`                                                                  // Object storage Region
	Endpoint             string `json:"endpoint"`                                                                // Endpoint for Object storage API.
	StorageTier          string `default:"Standard"              example:"Standard"          json:"storageTier"` // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	UploadCutoff         string `default:"200Mi"                 json:"uploadCutoff"`                            // Cutoff for switching to chunked upload.
	ChunkSize            string `default:"5Mi"                   json:"chunkSize"`                               // Chunk size to use for uploading.
	UploadConcurrency    int    `default:"10"                    json:"uploadConcurrency"`                       // Concurrency for multipart uploads.
	CopyCutoff           string `default:"4.656Gi"               json:"copyCutoff"`                              // Cutoff for switching to multipart copy.
	CopyTimeout          string `default:"1m0s"                  json:"copyTimeout"`                             // Timeout for copy.
	DisableChecksum      bool   `default:"false"                 json:"disableChecksum"`                         // Don't store MD5 checksum with object metadata.
	Encoding             string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                                // The encoding for the backend.
	LeavePartsOnError    bool   `default:"false"                 json:"leavePartsOnError"`                       // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	NoCheckBucket        bool   `default:"false"                 json:"noCheckBucket"`                           // If set, don't attempt to check the bucket exists or create it.
	SseCustomerKeyFile   string `example:""                      json:"sseCustomerKeyFile"`                      // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKey       string `example:""                      json:"sseCustomerKey"`                          // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKeySha256 string `example:""                      json:"sseCustomerKeySha256"`                    // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseKmsKeyId          string `example:""                      json:"sseKmsKeyId"`                             // if using using your own master key in vault, this header specifies the
	SseCustomerAlgorithm string `example:""                      json:"sseCustomerAlgorithm"`                    // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
}

type createOosEnv_authStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       oosEnv_authConfig  `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateOosEnv_authStorage
// @Summary Create Oos storage with env_auth - automatically pickup the credentials from runtime(env), first one to provide auth wins
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createOosEnv_authStorageRequest true "Request body"
// @Router /storage/oos/env_auth [post]
func createOosEnv_authStorage() {}

type oosInstance_principal_authConfig struct {
	Namespace            string `json:"namespace"`                                                               // Object storage namespace
	Compartment          string `json:"compartment"`                                                             // Object storage compartment OCID
	Region               string `json:"region"`                                                                  // Object storage Region
	Endpoint             string `json:"endpoint"`                                                                // Endpoint for Object storage API.
	StorageTier          string `default:"Standard"              example:"Standard"          json:"storageTier"` // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	UploadCutoff         string `default:"200Mi"                 json:"uploadCutoff"`                            // Cutoff for switching to chunked upload.
	ChunkSize            string `default:"5Mi"                   json:"chunkSize"`                               // Chunk size to use for uploading.
	UploadConcurrency    int    `default:"10"                    json:"uploadConcurrency"`                       // Concurrency for multipart uploads.
	CopyCutoff           string `default:"4.656Gi"               json:"copyCutoff"`                              // Cutoff for switching to multipart copy.
	CopyTimeout          string `default:"1m0s"                  json:"copyTimeout"`                             // Timeout for copy.
	DisableChecksum      bool   `default:"false"                 json:"disableChecksum"`                         // Don't store MD5 checksum with object metadata.
	Encoding             string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                                // The encoding for the backend.
	LeavePartsOnError    bool   `default:"false"                 json:"leavePartsOnError"`                       // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	NoCheckBucket        bool   `default:"false"                 json:"noCheckBucket"`                           // If set, don't attempt to check the bucket exists or create it.
	SseCustomerKeyFile   string `example:""                      json:"sseCustomerKeyFile"`                      // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKey       string `example:""                      json:"sseCustomerKey"`                          // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKeySha256 string `example:""                      json:"sseCustomerKeySha256"`                    // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseKmsKeyId          string `example:""                      json:"sseKmsKeyId"`                             // if using using your own master key in vault, this header specifies the
	SseCustomerAlgorithm string `example:""                      json:"sseCustomerAlgorithm"`                    // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
}

type createOosInstance_principal_authStorageRequest struct {
	Name         string                           `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string                           `json:"path"`                      // Path of the storage
	Config       oosInstance_principal_authConfig `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig               `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateOosInstance_principal_authStorage
// @Summary Create Oos storage with instance_principal_auth - use instance principals to authorize an instance to make API calls.
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createOosInstance_principal_authStorageRequest true "Request body"
// @Router /storage/oos/instance_principal_auth [post]
func createOosInstance_principal_authStorage() {}

type oosNo_authConfig struct {
	Namespace            string `json:"namespace"`                                                               // Object storage namespace
	Region               string `json:"region"`                                                                  // Object storage Region
	Endpoint             string `json:"endpoint"`                                                                // Endpoint for Object storage API.
	StorageTier          string `default:"Standard"              example:"Standard"          json:"storageTier"` // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	UploadCutoff         string `default:"200Mi"                 json:"uploadCutoff"`                            // Cutoff for switching to chunked upload.
	ChunkSize            string `default:"5Mi"                   json:"chunkSize"`                               // Chunk size to use for uploading.
	UploadConcurrency    int    `default:"10"                    json:"uploadConcurrency"`                       // Concurrency for multipart uploads.
	CopyCutoff           string `default:"4.656Gi"               json:"copyCutoff"`                              // Cutoff for switching to multipart copy.
	CopyTimeout          string `default:"1m0s"                  json:"copyTimeout"`                             // Timeout for copy.
	DisableChecksum      bool   `default:"false"                 json:"disableChecksum"`                         // Don't store MD5 checksum with object metadata.
	Encoding             string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                                // The encoding for the backend.
	LeavePartsOnError    bool   `default:"false"                 json:"leavePartsOnError"`                       // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	NoCheckBucket        bool   `default:"false"                 json:"noCheckBucket"`                           // If set, don't attempt to check the bucket exists or create it.
	SseCustomerKeyFile   string `example:""                      json:"sseCustomerKeyFile"`                      // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKey       string `example:""                      json:"sseCustomerKey"`                          // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKeySha256 string `example:""                      json:"sseCustomerKeySha256"`                    // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseKmsKeyId          string `example:""                      json:"sseKmsKeyId"`                             // if using using your own master key in vault, this header specifies the
	SseCustomerAlgorithm string `example:""                      json:"sseCustomerAlgorithm"`                    // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
}

type createOosNo_authStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       oosNo_authConfig   `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateOosNo_authStorage
// @Summary Create Oos storage with no_auth - no credentials needed, this is typically for reading public buckets
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createOosNo_authStorageRequest true "Request body"
// @Router /storage/oos/no_auth [post]
func createOosNo_authStorage() {}

type oosResource_principal_authConfig struct {
	Namespace            string `json:"namespace"`                                                               // Object storage namespace
	Compartment          string `json:"compartment"`                                                             // Object storage compartment OCID
	Region               string `json:"region"`                                                                  // Object storage Region
	Endpoint             string `json:"endpoint"`                                                                // Endpoint for Object storage API.
	StorageTier          string `default:"Standard"              example:"Standard"          json:"storageTier"` // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	UploadCutoff         string `default:"200Mi"                 json:"uploadCutoff"`                            // Cutoff for switching to chunked upload.
	ChunkSize            string `default:"5Mi"                   json:"chunkSize"`                               // Chunk size to use for uploading.
	UploadConcurrency    int    `default:"10"                    json:"uploadConcurrency"`                       // Concurrency for multipart uploads.
	CopyCutoff           string `default:"4.656Gi"               json:"copyCutoff"`                              // Cutoff for switching to multipart copy.
	CopyTimeout          string `default:"1m0s"                  json:"copyTimeout"`                             // Timeout for copy.
	DisableChecksum      bool   `default:"false"                 json:"disableChecksum"`                         // Don't store MD5 checksum with object metadata.
	Encoding             string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                                // The encoding for the backend.
	LeavePartsOnError    bool   `default:"false"                 json:"leavePartsOnError"`                       // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	NoCheckBucket        bool   `default:"false"                 json:"noCheckBucket"`                           // If set, don't attempt to check the bucket exists or create it.
	SseCustomerKeyFile   string `example:""                      json:"sseCustomerKeyFile"`                      // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKey       string `example:""                      json:"sseCustomerKey"`                          // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKeySha256 string `example:""                      json:"sseCustomerKeySha256"`                    // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseKmsKeyId          string `example:""                      json:"sseKmsKeyId"`                             // if using using your own master key in vault, this header specifies the
	SseCustomerAlgorithm string `example:""                      json:"sseCustomerAlgorithm"`                    // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
}

type createOosResource_principal_authStorageRequest struct {
	Name         string                           `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string                           `json:"path"`                      // Path of the storage
	Config       oosResource_principal_authConfig `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig               `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateOosResource_principal_authStorage
// @Summary Create Oos storage with resource_principal_auth - use resource principals to make API calls
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createOosResource_principal_authStorageRequest true "Request body"
// @Router /storage/oos/resource_principal_auth [post]
func createOosResource_principal_authStorage() {}

type oosUser_principal_authConfig struct {
	Namespace            string `json:"namespace"`                                                                 // Object storage namespace
	Compartment          string `json:"compartment"`                                                               // Object storage compartment OCID
	Region               string `json:"region"`                                                                    // Object storage Region
	Endpoint             string `json:"endpoint"`                                                                  // Endpoint for Object storage API.
	ConfigFile           string `default:"~/.oci/config"         example:"~/.oci/config"     json:"configFile"`    // Path to OCI config file
	ConfigProfile        string `default:"Default"               example:"Default"           json:"configProfile"` // Profile name inside the oci config file
	StorageTier          string `default:"Standard"              example:"Standard"          json:"storageTier"`   // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	UploadCutoff         string `default:"200Mi"                 json:"uploadCutoff"`                              // Cutoff for switching to chunked upload.
	ChunkSize            string `default:"5Mi"                   json:"chunkSize"`                                 // Chunk size to use for uploading.
	UploadConcurrency    int    `default:"10"                    json:"uploadConcurrency"`                         // Concurrency for multipart uploads.
	CopyCutoff           string `default:"4.656Gi"               json:"copyCutoff"`                                // Cutoff for switching to multipart copy.
	CopyTimeout          string `default:"1m0s"                  json:"copyTimeout"`                               // Timeout for copy.
	DisableChecksum      bool   `default:"false"                 json:"disableChecksum"`                           // Don't store MD5 checksum with object metadata.
	Encoding             string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                                  // The encoding for the backend.
	LeavePartsOnError    bool   `default:"false"                 json:"leavePartsOnError"`                         // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	NoCheckBucket        bool   `default:"false"                 json:"noCheckBucket"`                             // If set, don't attempt to check the bucket exists or create it.
	SseCustomerKeyFile   string `example:""                      json:"sseCustomerKeyFile"`                        // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKey       string `example:""                      json:"sseCustomerKey"`                            // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKeySha256 string `example:""                      json:"sseCustomerKeySha256"`                      // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseKmsKeyId          string `example:""                      json:"sseKmsKeyId"`                               // if using using your own master key in vault, this header specifies the
	SseCustomerAlgorithm string `example:""                      json:"sseCustomerAlgorithm"`                      // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
}

type createOosUser_principal_authStorageRequest struct {
	Name         string                       `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string                       `json:"path"`                      // Path of the storage
	Config       oosUser_principal_authConfig `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig           `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateOosUser_principal_authStorage
// @Summary Create Oos storage with user_principal_auth - use an OCI user and an API key for authentication.
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createOosUser_principal_authStorageRequest true "Request body"
// @Router /storage/oos/user_principal_auth [post]
func createOosUser_principal_authStorage() {}

type opendriveConfig struct {
	Username  string `json:"username"`                                                                                                                                          // Username.
	Password  string `json:"password"`                                                                                                                                          // Password.
	Encoding  string `default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot" json:"encoding"`  // The encoding for the backend.
	ChunkSize string `default:"10Mi"                                                                                                                          json:"chunkSize"` // Files will be uploaded in chunks this size.
}

type createOpendriveStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       opendriveConfig    `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateOpendriveStorage
// @Summary Create Opendrive storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createOpendriveStorageRequest true "Request body"
// @Router /storage/opendrive [post]
func createOpendriveStorage() {}

type pcloudConfig struct {
	ClientId     string `json:"clientId"`                                                                            // OAuth Client Id.
	ClientSecret string `json:"clientSecret"`                                                                        // OAuth Client Secret.
	Token        string `json:"token"`                                                                               // OAuth Access Token as a JSON blob.
	AuthUrl      string `json:"authUrl"`                                                                             // Auth server URL.
	TokenUrl     string `json:"tokenUrl"`                                                                            // Token server url.
	Encoding     string `default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot" json:"encoding"`                          // The encoding for the backend.
	RootFolderId string `default:"d0"                                      json:"rootFolderId"`                      // Fill in for rclone to use a non root folder as its starting point.
	Hostname     string `default:"api.pcloud.com"                          example:"api.pcloud.com" json:"hostname"` // Hostname to connect to.
	Username     string `json:"username"`                                                                            // Your pcloud username.
	Password     string `json:"password"`                                                                            // Your pcloud password.
}

type createPcloudStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       pcloudConfig       `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreatePcloudStorage
// @Summary Create Pcloud storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createPcloudStorageRequest true "Request body"
// @Router /storage/pcloud [post]
func createPcloudStorage() {}

type premiumizemeConfig struct {
	ApiKey   string `json:"apiKey"`                                                                 // API Key.
	Encoding string `default:"Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot" json:"encoding"` // The encoding for the backend.
}

type createPremiumizemeStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       premiumizemeConfig `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreatePremiumizemeStorage
// @Summary Create Premiumizeme storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createPremiumizemeStorageRequest true "Request body"
// @Router /storage/premiumizeme [post]
func createPremiumizemeStorage() {}

type putioConfig struct {
	Encoding string `default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot" json:"encoding"` // The encoding for the backend.
}

type createPutioStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       putioConfig        `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreatePutioStorage
// @Summary Create Putio storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createPutioStorageRequest true "Request body"
// @Router /storage/putio [post]
func createPutioStorage() {}

type qingstorConfig struct {
	EnvAuth           bool   `default:"false"                 example:"false"          json:"envAuth"` // Get QingStor credentials from runtime.
	AccessKeyId       string `json:"accessKeyId"`                                                      // QingStor Access Key ID.
	SecretAccessKey   string `json:"secretAccessKey"`                                                  // QingStor Secret Access Key (password).
	Endpoint          string `json:"endpoint"`                                                         // Enter an endpoint URL to connection QingStor API.
	Zone              string `example:"pek3a"                 json:"zone"`                             // Zone to connect to.
	ConnectionRetries int    `default:"3"                     json:"connectionRetries"`                // Number of connection retries.
	UploadCutoff      string `default:"200Mi"                 json:"uploadCutoff"`                     // Cutoff for switching to chunked upload.
	ChunkSize         string `default:"4Mi"                   json:"chunkSize"`                        // Chunk size to use for uploading.
	UploadConcurrency int    `default:"1"                     json:"uploadConcurrency"`                // Concurrency for multipart uploads.
	Encoding          string `default:"Slash,Ctl,InvalidUtf8" json:"encoding"`                         // The encoding for the backend.
}

type createQingstorStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       qingstorConfig     `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateQingstorStorage
// @Summary Create Qingstor storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createQingstorStorageRequest true "Request body"
// @Router /storage/qingstor [post]
func createQingstorStorage() {}

type s3AWSConfig struct {
	EnvAuth               bool   `default:"false"                 example:"false"              json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                          // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                      // AWS Secret Access Key (password).
	Region                string `example:"us-east-1"             json:"region"`                               // Region to connect to.
	Endpoint              string `json:"endpoint"`                                                             // Endpoint for S3 API.
	LocationConstraint    string `example:""                      json:"locationConstraint"`                   // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                                                  // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"               json:"bucketAcl"`                            // Canned ACL used when creating buckets.
	RequesterPays         bool   `default:"false"                 json:"requesterPays"`                        // Enables requester pays option when interacting with S3 bucket.
	ServerSideEncryption  string `example:""                      json:"serverSideEncryption"`                 // The server-side encryption algorithm used when storing this object in S3.
	SseCustomerAlgorithm  string `example:""                      json:"sseCustomerAlgorithm"`                 // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	SseKmsKeyId           string `example:""                      json:"sseKmsKeyId"`                          // If using KMS ID you must provide the ARN of Key.
	SseCustomerKey        string `example:""                      json:"sseCustomerKey"`                       // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	SseCustomerKeyBase64  string `example:""                      json:"sseCustomerKeyBase64"`                 // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	SseCustomerKeyMd5     string `example:""                      json:"sseCustomerKeyMd5"`                    // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	StorageClass          string `example:""                      json:"storageClass"`                         // The storage class to use when storing new objects in S3.
	UploadCutoff          string `default:"200Mi"                 json:"uploadCutoff"`                         // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                   json:"chunkSize"`                            // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                 json:"maxUploadParts"`                       // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"               json:"copyCutoff"`                           // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                 json:"disableChecksum"`                      // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                              // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                         // An AWS session token.
	UploadConcurrency     int    `default:"4"                     json:"uploadConcurrency"`                    // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                  json:"forcePathStyle"`                       // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                 json:"v2Auth"`                               // If true use v2 authentication.
	UseAccelerateEndpoint bool   `default:"false"                 json:"useAccelerateEndpoint"`                // If true use the AWS S3 accelerated endpoint.
	LeavePartsOnError     bool   `default:"false"                 json:"leavePartsOnError"`                    // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	ListChunk             int    `default:"1000"                  json:"listChunk"`                            // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                     json:"listVersion"`                          // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                 json:"listUrlEncode"`                        // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                 json:"noCheckBucket"`                        // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                 json:"noHead"`                               // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                 json:"noHeadObject"`                         // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                             // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                  json:"memoryPoolFlushTime"`                  // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                 json:"memoryPoolUseMmap"`                    // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                 json:"disableHttp2"`                         // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                          // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                 json:"useMultipartEtag"`                     // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                 json:"usePresignedRequest"`                  // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                 json:"versions"`                             // Include old versions in directory listings.
	VersionAt             string `default:"off"                   json:"versionAt"`                            // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                 json:"decompress"`                           // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                 json:"mightGzip"`                            // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                 json:"noSystemMetadata"`                     // Suppress setting and reading of system metadata
	StsEndpoint           string `json:"stsEndpoint"`                                                          // Endpoint for STS.
}

type createS3AWSStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3AWSConfig        `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3AWSStorage
// @Summary Create S3 storage with AWS - Amazon Web Services (AWS) S3
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3AWSStorageRequest true "Request body"
// @Router /storage/s3/aws [post]
func createS3AWSStorage() {}

type s3AlibabaConfig struct {
	EnvAuth               bool   `default:"false"                       example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                          // AWS Secret Access Key (password).
	Endpoint              string `example:"oss-accelerate.aliyuncs.com" json:"endpoint"`                           // Endpoint for OSS API.
	Acl                   string `json:"acl"`                                                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                     json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	StorageClass          string `example:""                            json:"storageClass"`                       // The storage class to use when storing new objects in OSS.
	UploadCutoff          string `default:"200Mi"                       json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                         json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                       json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                     json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                       json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                             // An AWS session token.
	UploadConcurrency     int    `default:"4"                           json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                        json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                       json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                        json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                           json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                       json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                       json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                       json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                       json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"       json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                        json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                       json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                       json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                              // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                       json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                       json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                       json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                         json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                       json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                       json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                       json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3AlibabaStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3AlibabaConfig    `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3AlibabaStorage
// @Summary Create S3 storage with Alibaba - Alibaba Cloud Object Storage System (OSS) formerly Aliyun
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3AlibabaStorageRequest true "Request body"
// @Router /storage/s3/alibaba [post]
func createS3AlibabaStorage() {}

type s3ArvanCloudConfig struct {
	EnvAuth               bool   `default:"false"                          example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                                 // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                             // AWS Secret Access Key (password).
	Endpoint              string `example:"s3.ir-thr-at1.arvanstorage.com" json:"endpoint"`                           // Endpoint for Arvan Cloud Object Storage (AOS) API.
	LocationConstraint    string `example:"ir-thr-at1"                     json:"locationConstraint"`                 // Location constraint - must match endpoint.
	Acl                   string `json:"acl"`                                                                         // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                        json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	StorageClass          string `example:"STANDARD"                       json:"storageClass"`                       // The storage class to use when storing new objects in ArvanCloud.
	UploadCutoff          string `default:"200Mi"                          json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                            json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                          json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                        json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                          json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                       // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                                     // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                                // An AWS session token.
	UploadConcurrency     int    `default:"4"                              json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                           json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                          json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                           json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                              json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                          json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                          json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                          json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                          json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"          json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                           json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                          json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                          json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                                 // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                          json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                          json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                          json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                            json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                          json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                          json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                          json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3ArvanCloudStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3ArvanCloudConfig `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3ArvanCloudStorage
// @Summary Create S3 storage with ArvanCloud - Arvan Cloud Object Storage (AOS)
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3ArvanCloudStorageRequest true "Request body"
// @Router /storage/s3/arvancloud [post]
func createS3ArvanCloudStorage() {}

type s3CephConfig struct {
	EnvAuth               bool   `default:"false"                 example:"false"             json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                         // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                     // AWS Secret Access Key (password).
	Region                string `example:""                      json:"region"`                              // Region to connect to.
	Endpoint              string `json:"endpoint"`                                                            // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                                                  // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                                                 // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"               json:"bucketAcl"`                           // Canned ACL used when creating buckets.
	ServerSideEncryption  string `example:""                      json:"serverSideEncryption"`                // The server-side encryption algorithm used when storing this object in S3.
	SseCustomerAlgorithm  string `example:""                      json:"sseCustomerAlgorithm"`                // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	SseKmsKeyId           string `example:""                      json:"sseKmsKeyId"`                         // If using KMS ID you must provide the ARN of Key.
	SseCustomerKey        string `example:""                      json:"sseCustomerKey"`                      // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	SseCustomerKeyBase64  string `example:""                      json:"sseCustomerKeyBase64"`                // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	SseCustomerKeyMd5     string `example:""                      json:"sseCustomerKeyMd5"`                   // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	UploadCutoff          string `default:"200Mi"                 json:"uploadCutoff"`                        // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                   json:"chunkSize"`                           // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                 json:"maxUploadParts"`                      // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"               json:"copyCutoff"`                          // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                 json:"disableChecksum"`                     // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                               // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                             // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                        // An AWS session token.
	UploadConcurrency     int    `default:"4"                     json:"uploadConcurrency"`                   // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                  json:"forcePathStyle"`                      // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                 json:"v2Auth"`                              // If true use v2 authentication.
	ListChunk             int    `default:"1000"                  json:"listChunk"`                           // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                     json:"listVersion"`                         // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                 json:"listUrlEncode"`                       // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                 json:"noCheckBucket"`                       // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                 json:"noHead"`                              // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                 json:"noHeadObject"`                        // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                            // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                  json:"memoryPoolFlushTime"`                 // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                 json:"memoryPoolUseMmap"`                   // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                 json:"disableHttp2"`                        // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                         // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                 json:"useMultipartEtag"`                    // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                 json:"usePresignedRequest"`                 // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                 json:"versions"`                            // Include old versions in directory listings.
	VersionAt             string `default:"off"                   json:"versionAt"`                           // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                 json:"decompress"`                          // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                 json:"mightGzip"`                           // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                 json:"noSystemMetadata"`                    // Suppress setting and reading of system metadata
}

type createS3CephStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3CephConfig       `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3CephStorage
// @Summary Create S3 storage with Ceph - Ceph Object Storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3CephStorageRequest true "Request body"
// @Router /storage/s3/ceph [post]
func createS3CephStorage() {}

type s3ChinaMobileConfig struct {
	EnvAuth               bool   `default:"false"                  example:"false"             json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                          // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                      // AWS Secret Access Key (password).
	Endpoint              string `example:"eos-wuxi-1.cmecloud.cn" json:"endpoint"`                            // Endpoint for China Mobile Ecloud Elastic Object Storage (EOS) API.
	LocationConstraint    string `example:"wuxi1"                  json:"locationConstraint"`                  // Location constraint - must match endpoint.
	Acl                   string `json:"acl"`                                                                  // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                json:"bucketAcl"`                           // Canned ACL used when creating buckets.
	ServerSideEncryption  string `example:""                       json:"serverSideEncryption"`                // The server-side encryption algorithm used when storing this object in S3.
	SseCustomerAlgorithm  string `example:""                       json:"sseCustomerAlgorithm"`                // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	SseCustomerKey        string `example:""                       json:"sseCustomerKey"`                      // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	SseCustomerKeyBase64  string `example:""                       json:"sseCustomerKeyBase64"`                // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	SseCustomerKeyMd5     string `example:""                       json:"sseCustomerKeyMd5"`                   // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	StorageClass          string `example:""                       json:"storageClass"`                        // The storage class to use when storing new objects in ChinaMobile.
	UploadCutoff          string `default:"200Mi"                  json:"uploadCutoff"`                        // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                    json:"chunkSize"`                           // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                  json:"maxUploadParts"`                      // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                json:"copyCutoff"`                          // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                  json:"disableChecksum"`                     // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                              // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                         // An AWS session token.
	UploadConcurrency     int    `default:"4"                      json:"uploadConcurrency"`                   // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                   json:"forcePathStyle"`                      // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                  json:"v2Auth"`                              // If true use v2 authentication.
	ListChunk             int    `default:"1000"                   json:"listChunk"`                           // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                      json:"listVersion"`                         // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                  json:"listUrlEncode"`                       // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                  json:"noCheckBucket"`                       // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                  json:"noHead"`                              // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                  json:"noHeadObject"`                        // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"  json:"encoding"`                            // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                   json:"memoryPoolFlushTime"`                 // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                  json:"memoryPoolUseMmap"`                   // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                  json:"disableHttp2"`                        // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                          // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                  json:"useMultipartEtag"`                    // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                  json:"usePresignedRequest"`                 // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                  json:"versions"`                            // Include old versions in directory listings.
	VersionAt             string `default:"off"                    json:"versionAt"`                           // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                  json:"decompress"`                          // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                  json:"mightGzip"`                           // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                  json:"noSystemMetadata"`                    // Suppress setting and reading of system metadata
}

type createS3ChinaMobileStorageRequest struct {
	Name         string              `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string              `json:"path"`                      // Path of the storage
	Config       s3ChinaMobileConfig `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig  `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3ChinaMobileStorage
// @Summary Create S3 storage with ChinaMobile - China Mobile Ecloud Elastic Object Storage (EOS)
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3ChinaMobileStorageRequest true "Request body"
// @Router /storage/s3/chinamobile [post]
func createS3ChinaMobileStorage() {}

type s3CloudflareConfig struct {
	EnvAuth               bool   `default:"false"                 example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                        // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                    // AWS Secret Access Key (password).
	Region                string `example:"auto"                  json:"region"`                             // Region to connect to.
	Endpoint              string `json:"endpoint"`                                                           // Endpoint for S3 API.
	BucketAcl             string `example:"private"               json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                 json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                   json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                 json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"               json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                 json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                              // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                            // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                       // An AWS session token.
	UploadConcurrency     int    `default:"4"                     json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                  json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                 json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                  json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                     json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                 json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                 json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                 json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                 json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                  json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                 json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                 json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                        // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                 json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                 json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                 json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                   json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                 json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                 json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                 json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3CloudflareStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3CloudflareConfig `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3CloudflareStorage
// @Summary Create S3 storage with Cloudflare - Cloudflare R2 Storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3CloudflareStorageRequest true "Request body"
// @Router /storage/s3/cloudflare [post]
func createS3CloudflareStorage() {}

type s3DigitalOceanConfig struct {
	EnvAuth               bool   `default:"false"                       example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                          // AWS Secret Access Key (password).
	Region                string `example:""                            json:"region"`                             // Region to connect to.
	Endpoint              string `example:"syd1.digitaloceanspaces.com" json:"endpoint"`                           // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                                                       // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                     json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                       json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                         json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                       json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                     json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                       json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                             // An AWS session token.
	UploadConcurrency     int    `default:"4"                           json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                        json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                       json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                        json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                           json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                       json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                       json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                       json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                       json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"       json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                        json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                       json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                       json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                              // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                       json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                       json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                       json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                         json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                       json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                       json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                       json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3DigitalOceanStorageRequest struct {
	Name         string               `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string               `json:"path"`                      // Path of the storage
	Config       s3DigitalOceanConfig `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig   `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3DigitalOceanStorage
// @Summary Create S3 storage with DigitalOcean - DigitalOcean Spaces
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3DigitalOceanStorageRequest true "Request body"
// @Router /storage/s3/digitalocean [post]
func createS3DigitalOceanStorage() {}

type s3DreamhostConfig struct {
	EnvAuth               bool   `default:"false"                      example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                             // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                         // AWS Secret Access Key (password).
	Region                string `example:""                           json:"region"`                             // Region to connect to.
	Endpoint              string `example:"objects-us-east-1.dream.io" json:"endpoint"`                           // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                                                      // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                                                     // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                    json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                      json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                        json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                      json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                    json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                      json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                   // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                                 // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                            // An AWS session token.
	UploadConcurrency     int    `default:"4"                          json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                       json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                      json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                       json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                          json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                      json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                      json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                      json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                      json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"      json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                       json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                      json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                      json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                             // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                      json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                      json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                      json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                        json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                      json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                      json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                      json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3DreamhostStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3DreamhostConfig  `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3DreamhostStorage
// @Summary Create S3 storage with Dreamhost - Dreamhost DreamObjects
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3DreamhostStorageRequest true "Request body"
// @Router /storage/s3/dreamhost [post]
func createS3DreamhostStorage() {}

type s3HuaweiOBSConfig struct {
	EnvAuth               bool   `default:"false"                            example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                                   // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                               // AWS Secret Access Key (password).
	Region                string `example:"af-south-1"                       json:"region"`                             // Region to connect to. - the location where your bucket will be created and your data stored. Need bo be same with your endpoint.
	Endpoint              string `example:"obs.af-south-1.myhuaweicloud.com" json:"endpoint"`                           // Endpoint for OBS API.
	Acl                   string `json:"acl"`                                                                           // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                          json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                            json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                              json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                            json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                          json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                            json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                         // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                                       // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                                  // An AWS session token.
	UploadConcurrency     int    `default:"4"                                json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                             json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                            json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                             json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                                json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                            json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                            json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                            json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                            json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"            json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                             json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                            json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                            json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                                   // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                            json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                            json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                            json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                              json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                            json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                            json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                            json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3HuaweiOBSStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3HuaweiOBSConfig  `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3HuaweiOBSStorage
// @Summary Create S3 storage with HuaweiOBS - Huawei Object Storage Service
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3HuaweiOBSStorageRequest true "Request body"
// @Router /storage/s3/huaweiobs [post]
func createS3HuaweiOBSStorage() {}

type s3IBMCOSConfig struct {
	EnvAuth               bool   `default:"false"                                      example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                                             // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                                         // AWS Secret Access Key (password).
	Region                string `example:""                                           json:"region"`                             // Region to connect to.
	Endpoint              string `example:"s3.us.cloud-object-storage.appdomain.cloud" json:"endpoint"`                           // Endpoint for IBM COS S3 API.
	LocationConstraint    string `example:"us-standard"                                json:"locationConstraint"`                 // Location constraint - must match endpoint when using IBM Cloud Public.
	Acl                   string `example:"private"                                    json:"acl"`                                // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                                    json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                                      json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                                        json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                                      json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                                    json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                                      json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                                   // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                                                 // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                                            // An AWS session token.
	UploadConcurrency     int    `default:"4"                                          json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                                       json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                                      json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                                       json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                                          json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                                      json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                                      json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                                      json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                                      json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"                      json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                                       json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                                      json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                                      json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                                             // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                                      json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                                      json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                                      json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                                        json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                                      json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                                      json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                                      json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3IBMCOSStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3IBMCOSConfig     `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3IBMCOSStorage
// @Summary Create S3 storage with IBMCOS - IBM COS S3
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3IBMCOSStorageRequest true "Request body"
// @Router /storage/s3/ibmcos [post]
func createS3IBMCOSStorage() {}

type s3IDriveConfig struct {
	EnvAuth               bool   `default:"false"                 example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                        // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                    // AWS Secret Access Key (password).
	Acl                   string `json:"acl"`                                                                // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"               json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                 json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                   json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                 json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"               json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                 json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                              // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                            // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                       // An AWS session token.
	UploadConcurrency     int    `default:"4"                     json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                  json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                 json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                  json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                     json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                 json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                 json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                 json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                 json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                  json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                 json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                 json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                        // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                 json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                 json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                 json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                   json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                 json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                 json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                 json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3IDriveStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3IDriveConfig     `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3IDriveStorage
// @Summary Create S3 storage with IDrive - IDrive e2
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3IDriveStorageRequest true "Request body"
// @Router /storage/s3/idrive [post]
func createS3IDriveStorage() {}

type s3IONOSConfig struct {
	EnvAuth               bool   `default:"false"                          example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                                 // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                             // AWS Secret Access Key (password).
	Region                string `example:"de"                             json:"region"`                             // Region where your bucket will be created and your data stored.
	Endpoint              string `example:"s3-eu-central-1.ionoscloud.com" json:"endpoint"`                           // Endpoint for IONOS S3 Object Storage.
	Acl                   string `json:"acl"`                                                                         // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                        json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                          json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                            json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                          json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                        json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                          json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                       // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                                     // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                                // An AWS session token.
	UploadConcurrency     int    `default:"4"                              json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                           json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                          json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                           json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                              json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                          json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                          json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                          json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                          json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"          json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                           json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                          json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                          json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                                 // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                          json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                          json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                          json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                            json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                          json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                          json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                          json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3IONOSStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3IONOSConfig      `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3IONOSStorage
// @Summary Create S3 storage with IONOS - IONOS Cloud
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3IONOSStorageRequest true "Request body"
// @Router /storage/s3/ionos [post]
func createS3IONOSStorage() {}

type s3LiaraConfig struct {
	EnvAuth               bool   `default:"false"                    example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                           // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                       // AWS Secret Access Key (password).
	Endpoint              string `example:"storage.iran.liara.space" json:"endpoint"`                           // Endpoint for Liara Object Storage API.
	Acl                   string `json:"acl"`                                                                   // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                  json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	StorageClass          string `example:"STANDARD"                 json:"storageClass"`                       // The storage class to use when storing new objects in Liara
	UploadCutoff          string `default:"200Mi"                    json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                      json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                    json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                  json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                    json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                 // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                               // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                          // An AWS session token.
	UploadConcurrency     int    `default:"4"                        json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                     json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                    json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                     json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                        json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                    json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                    json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                    json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                    json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"    json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                     json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                    json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                    json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                           // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                    json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                    json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                    json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                      json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                    json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                    json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                    json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3LiaraStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3LiaraConfig      `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3LiaraStorage
// @Summary Create S3 storage with Liara - Liara Object Storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3LiaraStorageRequest true "Request body"
// @Router /storage/s3/liara [post]
func createS3LiaraStorage() {}

type s3LyveCloudConfig struct {
	EnvAuth               bool   `default:"false"                              example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                                     // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                                 // AWS Secret Access Key (password).
	Region                string `example:""                                   json:"region"`                             // Region to connect to.
	Endpoint              string `example:"s3.us-east-1.lyvecloud.seagate.com" json:"endpoint"`                           // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                                                              // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                                                             // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                            json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                              json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                                json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                              json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                            json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                              json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                           // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                                         // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                                    // An AWS session token.
	UploadConcurrency     int    `default:"4"                                  json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                               json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                              json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                               json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                                  json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                              json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                              json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                              json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                              json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"              json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                               json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                              json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                              json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                                     // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                              json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                              json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                              json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                                json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                              json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                              json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                              json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3LyveCloudStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3LyveCloudConfig  `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3LyveCloudStorage
// @Summary Create S3 storage with LyveCloud - Seagate Lyve Cloud
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3LyveCloudStorageRequest true "Request body"
// @Router /storage/s3/lyvecloud [post]
func createS3LyveCloudStorage() {}

type s3MinioConfig struct {
	EnvAuth               bool   `default:"false"                 example:"false"             json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                         // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                     // AWS Secret Access Key (password).
	Region                string `example:""                      json:"region"`                              // Region to connect to.
	Endpoint              string `json:"endpoint"`                                                            // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                                                  // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                                                 // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"               json:"bucketAcl"`                           // Canned ACL used when creating buckets.
	ServerSideEncryption  string `example:""                      json:"serverSideEncryption"`                // The server-side encryption algorithm used when storing this object in S3.
	SseCustomerAlgorithm  string `example:""                      json:"sseCustomerAlgorithm"`                // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	SseKmsKeyId           string `example:""                      json:"sseKmsKeyId"`                         // If using KMS ID you must provide the ARN of Key.
	SseCustomerKey        string `example:""                      json:"sseCustomerKey"`                      // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	SseCustomerKeyBase64  string `example:""                      json:"sseCustomerKeyBase64"`                // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	SseCustomerKeyMd5     string `example:""                      json:"sseCustomerKeyMd5"`                   // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	UploadCutoff          string `default:"200Mi"                 json:"uploadCutoff"`                        // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                   json:"chunkSize"`                           // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                 json:"maxUploadParts"`                      // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"               json:"copyCutoff"`                          // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                 json:"disableChecksum"`                     // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                               // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                             // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                        // An AWS session token.
	UploadConcurrency     int    `default:"4"                     json:"uploadConcurrency"`                   // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                  json:"forcePathStyle"`                      // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                 json:"v2Auth"`                              // If true use v2 authentication.
	ListChunk             int    `default:"1000"                  json:"listChunk"`                           // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                     json:"listVersion"`                         // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                 json:"listUrlEncode"`                       // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                 json:"noCheckBucket"`                       // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                 json:"noHead"`                              // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                 json:"noHeadObject"`                        // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                            // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                  json:"memoryPoolFlushTime"`                 // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                 json:"memoryPoolUseMmap"`                   // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                 json:"disableHttp2"`                        // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                         // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                 json:"useMultipartEtag"`                    // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                 json:"usePresignedRequest"`                 // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                 json:"versions"`                            // Include old versions in directory listings.
	VersionAt             string `default:"off"                   json:"versionAt"`                           // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                 json:"decompress"`                          // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                 json:"mightGzip"`                           // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                 json:"noSystemMetadata"`                    // Suppress setting and reading of system metadata
}

type createS3MinioStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3MinioConfig      `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3MinioStorage
// @Summary Create S3 storage with Minio - Minio Object Storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3MinioStorageRequest true "Request body"
// @Router /storage/s3/minio [post]
func createS3MinioStorage() {}

type s3NeteaseConfig struct {
	EnvAuth               bool   `default:"false"                 example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                        // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                    // AWS Secret Access Key (password).
	Region                string `example:""                      json:"region"`                             // Region to connect to.
	Endpoint              string `json:"endpoint"`                                                           // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                                                 // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                                                // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"               json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                 json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                   json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                 json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"               json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                 json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                              // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                            // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                       // An AWS session token.
	UploadConcurrency     int    `default:"4"                     json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                  json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                 json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                  json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                     json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                 json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                 json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                 json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                 json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                  json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                 json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                 json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                        // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                 json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                 json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                 json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                   json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                 json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                 json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                 json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3NeteaseStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3NeteaseConfig    `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3NeteaseStorage
// @Summary Create S3 storage with Netease - Netease Object Storage (NOS)
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3NeteaseStorageRequest true "Request body"
// @Router /storage/s3/netease [post]
func createS3NeteaseStorage() {}

type s3OtherConfig struct {
	EnvAuth               bool   `default:"false"                 example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                        // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                    // AWS Secret Access Key (password).
	Region                string `example:""                      json:"region"`                             // Region to connect to.
	Endpoint              string `json:"endpoint"`                                                           // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                                                 // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                                                // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"               json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                 json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                   json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                 json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"               json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                 json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                              // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                            // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                       // An AWS session token.
	UploadConcurrency     int    `default:"4"                     json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                  json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                 json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                  json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                     json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                 json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                 json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                 json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                 json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                  json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                 json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                 json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                        // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                 json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                 json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                 json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                   json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                 json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                 json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                 json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3OtherStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3OtherConfig      `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3OtherStorage
// @Summary Create S3 storage with Other - Any other S3 compatible provider
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3OtherStorageRequest true "Request body"
// @Router /storage/s3/other [post]
func createS3OtherStorage() {}

type s3QiniuConfig struct {
	EnvAuth               bool   `default:"false"                    example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                           // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                       // AWS Secret Access Key (password).
	Region                string `example:"cn-east-1"                json:"region"`                             // Region to connect to.
	Endpoint              string `example:"s3-cn-east-1.qiniucs.com" json:"endpoint"`                           // Endpoint for Qiniu Object Storage.
	LocationConstraint    string `example:"cn-east-1"                json:"locationConstraint"`                 // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                                                   // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                  json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	StorageClass          string `example:"STANDARD"                 json:"storageClass"`                       // The storage class to use when storing new objects in Qiniu.
	UploadCutoff          string `default:"200Mi"                    json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                      json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                    json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                  json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                    json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                 // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                               // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                          // An AWS session token.
	UploadConcurrency     int    `default:"4"                        json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                     json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                    json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                     json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                        json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                    json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                    json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                    json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                    json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"    json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                     json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                    json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                    json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                           // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                    json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                    json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                    json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                      json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                    json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                    json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                    json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3QiniuStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3QiniuConfig      `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3QiniuStorage
// @Summary Create S3 storage with Qiniu - Qiniu Object Storage (Kodo)
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3QiniuStorageRequest true "Request body"
// @Router /storage/s3/qiniu [post]
func createS3QiniuStorage() {}

type s3RackCorpConfig struct {
	EnvAuth               bool   `default:"false"                 example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                        // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                    // AWS Secret Access Key (password).
	Region                string `example:"global"                json:"region"`                             // region - the location where your bucket will be created and your data stored.
	Endpoint              string `example:"s3.rackcorp.com"       json:"endpoint"`                           // Endpoint for RackCorp Object Storage.
	LocationConstraint    string `example:"global"                json:"locationConstraint"`                 // Location constraint - the location where your bucket will be located and your data stored.
	Acl                   string `json:"acl"`                                                                // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"               json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                 json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                   json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                 json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"               json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                 json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                              // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                            // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                       // An AWS session token.
	UploadConcurrency     int    `default:"4"                     json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                  json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                 json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                  json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                     json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                 json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                 json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                 json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                 json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                  json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                 json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                 json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                        // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                 json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                 json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                 json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                   json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                 json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                 json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                 json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3RackCorpStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3RackCorpConfig   `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3RackCorpStorage
// @Summary Create S3 storage with RackCorp - RackCorp Object Storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3RackCorpStorageRequest true "Request body"
// @Router /storage/s3/rackcorp [post]
func createS3RackCorpStorage() {}

type s3ScalewayConfig struct {
	EnvAuth               bool   `default:"false"                 example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                        // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                    // AWS Secret Access Key (password).
	Region                string `example:"nl-ams"                json:"region"`                             // Region to connect to.
	Endpoint              string `example:"s3.nl-ams.scw.cloud"   json:"endpoint"`                           // Endpoint for Scaleway Object Storage.
	Acl                   string `json:"acl"`                                                                // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"               json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	StorageClass          string `example:""                      json:"storageClass"`                       // The storage class to use when storing new objects in S3.
	UploadCutoff          string `default:"200Mi"                 json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                   json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                 json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"               json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                 json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                              // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                            // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                       // An AWS session token.
	UploadConcurrency     int    `default:"4"                     json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                  json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                 json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                  json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                     json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                 json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                 json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                 json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                 json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                  json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                 json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                 json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                        // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                 json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                 json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                 json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                   json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                 json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                 json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                 json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3ScalewayStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3ScalewayConfig   `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3ScalewayStorage
// @Summary Create S3 storage with Scaleway - Scaleway Object Storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3ScalewayStorageRequest true "Request body"
// @Router /storage/s3/scaleway [post]
func createS3ScalewayStorage() {}

type s3SeaweedFSConfig struct {
	EnvAuth               bool   `default:"false"                 example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                        // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                    // AWS Secret Access Key (password).
	Region                string `example:""                      json:"region"`                             // Region to connect to.
	Endpoint              string `example:"localhost:8333"        json:"endpoint"`                           // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                                                 // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                                                // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"               json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                 json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                   json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                 json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"               json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                 json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                              // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                            // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                       // An AWS session token.
	UploadConcurrency     int    `default:"4"                     json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                  json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                 json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                  json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                     json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                 json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                 json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                 json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                 json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                  json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                 json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                 json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                        // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                 json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                 json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                 json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                   json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                 json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                 json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                 json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3SeaweedFSStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3SeaweedFSConfig  `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3SeaweedFSStorage
// @Summary Create S3 storage with SeaweedFS - SeaweedFS S3
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3SeaweedFSStorageRequest true "Request body"
// @Router /storage/s3/seaweedfs [post]
func createS3SeaweedFSStorage() {}

type s3StackPathConfig struct {
	EnvAuth               bool   `default:"false"                             example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                                    // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                                // AWS Secret Access Key (password).
	Region                string `example:""                                  json:"region"`                             // Region to connect to.
	Endpoint              string `example:"s3.us-east-2.stackpathstorage.com" json:"endpoint"`                           // Endpoint for StackPath Object Storage.
	Acl                   string `json:"acl"`                                                                            // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                           json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                             json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                               json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                             json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                           json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                             json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                          // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                                        // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                                   // An AWS session token.
	UploadConcurrency     int    `default:"4"                                 json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                              json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                             json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                              json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                                 json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                             json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                             json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                             json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                             json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"             json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                              json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                             json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                             json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                                    // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                             json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                             json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                             json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                               json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                             json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                             json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                             json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3StackPathStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3StackPathConfig  `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3StackPathStorage
// @Summary Create S3 storage with StackPath - StackPath Object Storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3StackPathStorageRequest true "Request body"
// @Router /storage/s3/stackpath [post]
func createS3StackPathStorage() {}

type s3StorjConfig struct {
	EnvAuth               bool   `default:"false"                 example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                        // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                    // AWS Secret Access Key (password).
	Endpoint              string `example:"gateway.storjshare.io" json:"endpoint"`                           // Endpoint for Storj Gateway.
	BucketAcl             string `example:"private"               json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                 json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                   json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                 json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"               json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                 json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                              // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                            // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                       // An AWS session token.
	UploadConcurrency     int    `default:"4"                     json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                  json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                 json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                  json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                     json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                 json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                 json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                 json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                 json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                  json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                 json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                 json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                        // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                 json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                 json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                 json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                   json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                 json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                 json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                 json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3StorjStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3StorjConfig      `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3StorjStorage
// @Summary Create S3 storage with Storj - Storj (S3 Compatible Gateway)
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3StorjStorageRequest true "Request body"
// @Router /storage/s3/storj [post]
func createS3StorjStorage() {}

type s3TencentCOSConfig struct {
	EnvAuth               bool   `default:"false"                       example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                          // AWS Secret Access Key (password).
	Endpoint              string `example:"cos.ap-beijing.myqcloud.com" json:"endpoint"`                           // Endpoint for Tencent COS API.
	Acl                   string `example:"default"                     json:"acl"`                                // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"                     json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	StorageClass          string `example:""                            json:"storageClass"`                       // The storage class to use when storing new objects in Tencent COS.
	UploadCutoff          string `default:"200Mi"                       json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                         json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                       json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"                     json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                       json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                             // An AWS session token.
	UploadConcurrency     int    `default:"4"                           json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                        json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                       json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                        json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                           json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                       json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                       json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                       json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                       json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot"       json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                        json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                       json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                       json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                              // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                       json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                       json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                       json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                         json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                       json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                       json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                       json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3TencentCOSStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3TencentCOSConfig `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3TencentCOSStorage
// @Summary Create S3 storage with TencentCOS - Tencent Cloud Object Storage (COS)
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3TencentCOSStorageRequest true "Request body"
// @Router /storage/s3/tencentcos [post]
func createS3TencentCOSStorage() {}

type s3WasabiConfig struct {
	EnvAuth               bool   `default:"false"                 example:"false"            json:"envAuth"` // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                        // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                                    // AWS Secret Access Key (password).
	Region                string `example:""                      json:"region"`                             // Region to connect to.
	Endpoint              string `example:"s3.wasabisys.com"      json:"endpoint"`                           // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                                                 // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                                                // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `example:"private"               json:"bucketAcl"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `default:"200Mi"                 json:"uploadCutoff"`                       // Cutoff for switching to chunked upload.
	ChunkSize             string `default:"5Mi"                   json:"chunkSize"`                          // Chunk size to use for uploading.
	MaxUploadParts        int    `default:"10000"                 json:"maxUploadParts"`                     // Maximum number of parts in a multipart upload.
	CopyCutoff            string `default:"4.656Gi"               json:"copyCutoff"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `default:"false"                 json:"disableChecksum"`                    // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                              // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                            // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                       // An AWS session token.
	UploadConcurrency     int    `default:"4"                     json:"uploadConcurrency"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `default:"true"                  json:"forcePathStyle"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `default:"false"                 json:"v2Auth"`                             // If true use v2 authentication.
	ListChunk             int    `default:"1000"                  json:"listChunk"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `default:"0"                     json:"listVersion"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `default:"unset"                 json:"listUrlEncode"`                      // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `default:"false"                 json:"noCheckBucket"`                      // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `default:"false"                 json:"noHead"`                             // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `default:"false"                 json:"noHeadObject"`                       // If set, do not do HEAD before GET when getting objects.
	Encoding              string `default:"Slash,InvalidUtf8,Dot" json:"encoding"`                           // The encoding for the backend.
	MemoryPoolFlushTime   string `default:"1m0s"                  json:"memoryPoolFlushTime"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `default:"false"                 json:"memoryPoolUseMmap"`                  // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `default:"false"                 json:"disableHttp2"`                       // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                        // Custom endpoint for downloads.
	UseMultipartEtag      string `default:"unset"                 json:"useMultipartEtag"`                   // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `default:"false"                 json:"usePresignedRequest"`                // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `default:"false"                 json:"versions"`                           // Include old versions in directory listings.
	VersionAt             string `default:"off"                   json:"versionAt"`                          // Show file versions as they were at the specified time.
	Decompress            bool   `default:"false"                 json:"decompress"`                         // If set this will decompress gzip encoded objects.
	MightGzip             string `default:"unset"                 json:"mightGzip"`                          // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `default:"false"                 json:"noSystemMetadata"`                   // Suppress setting and reading of system metadata
}

type createS3WasabiStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       s3WasabiConfig     `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateS3WasabiStorage
// @Summary Create S3 storage with Wasabi - Wasabi Object Storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createS3WasabiStorageRequest true "Request body"
// @Router /storage/s3/wasabi [post]
func createS3WasabiStorage() {}

type seafileConfig struct {
	Url           string `example:"https://cloud.seafile.com/"                  json:"url"`           // URL of seafile host to connect to.
	User          string `json:"user"`                                                                // User name (usually email address).
	Pass          string `json:"pass"`                                                                // Password.
	TwoFA         bool   `default:"false"                                       json:"2fa"`           // Two-factor authentication ('true' if the account has 2FA enabled).
	Library       string `json:"library"`                                                             // Name of the library.
	LibraryKey    string `json:"libraryKey"`                                                          // Library password (for encrypted libraries only).
	CreateLibrary bool   `default:"false"                                       json:"createLibrary"` // Should rclone create a library if it doesn't exist.
	AuthToken     string `json:"authToken"`                                                           // Authentication token.
	Encoding      string `default:"Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8" json:"encoding"`      // The encoding for the backend.
}

type createSeafileStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       seafileConfig      `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateSeafileStorage
// @Summary Create Seafile storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createSeafileStorageRequest true "Request body"
// @Router /storage/seafile [post]
func createSeafileStorage() {}

type sftpConfig struct {
	Host                    string `json:"host"`                                                                          // SSH host to connect to.
	User                    string `default:"$USER"              json:"user"`                                             // SSH username.
	Port                    int    `default:"22"                 json:"port"`                                             // SSH port number.
	Pass                    string `json:"pass"`                                                                          // SSH password, leave blank to use ssh-agent.
	KeyPem                  string `json:"keyPem"`                                                                        // Raw PEM-encoded private key.
	KeyFile                 string `json:"keyFile"`                                                                       // Path to PEM-encoded private key file.
	KeyFilePass             string `json:"keyFilePass"`                                                                   // The passphrase to decrypt the PEM-encoded private key file.
	PubkeyFile              string `json:"pubkeyFile"`                                                                    // Optional path to public key file.
	KnownHostsFile          string `example:"~/.ssh/known_hosts" json:"knownHostsFile"`                                   // Optional path to known_hosts file.
	KeyUseAgent             bool   `default:"false"              json:"keyUseAgent"`                                      // When set forces the usage of the ssh-agent.
	UseInsecureCipher       bool   `default:"false"              example:"false"                json:"useInsecureCipher"` // Enable the use of insecure ciphers and key exchange methods.
	DisableHashcheck        bool   `default:"false"              json:"disableHashcheck"`                                 // Disable the execution of SSH commands to determine if remote file hashing is available.
	AskPassword             bool   `default:"false"              json:"askPassword"`                                      // Allow asking for SFTP password when needed.
	PathOverride            string `json:"pathOverride"`                                                                  // Override path used by SSH shell commands.
	SetModtime              bool   `default:"true"               json:"setModtime"`                                       // Set the modified time on the remote if set.
	ShellType               string `example:"none"               json:"shellType"`                                        // The type of SSH shell on remote server, if any.
	Md5sumCommand           string `json:"md5sumCommand"`                                                                 // The command used to read md5 hashes.
	Sha1sumCommand          string `json:"sha1sumCommand"`                                                                // The command used to read sha1 hashes.
	SkipLinks               bool   `default:"false"              json:"skipLinks"`                                        // Set to skip any symlinks and any other non regular files.
	Subsystem               string `default:"sftp"               json:"subsystem"`                                        // Specifies the SSH2 subsystem on the remote host.
	ServerCommand           string `json:"serverCommand"`                                                                 // Specifies the path or command to run a sftp server on the remote host.
	UseFstat                bool   `default:"false"              json:"useFstat"`                                         // If set use fstat instead of stat.
	DisableConcurrentReads  bool   `default:"false"              json:"disableConcurrentReads"`                           // If set don't use concurrent reads.
	DisableConcurrentWrites bool   `default:"false"              json:"disableConcurrentWrites"`                          // If set don't use concurrent writes.
	IdleTimeout             string `default:"1m0s"               json:"idleTimeout"`                                      // Max time before closing idle connections.
	ChunkSize               string `default:"32Ki"               json:"chunkSize"`                                        // Upload and download chunk size.
	Concurrency             int    `default:"64"                 json:"concurrency"`                                      // The maximum number of outstanding requests for one file
	SetEnv                  string `json:"setEnv"`                                                                        // Environment variables to pass to sftp and commands
	Ciphers                 string `json:"ciphers"`                                                                       // Space separated list of ciphers to be used for session encryption, ordered by preference.
	KeyExchange             string `json:"keyExchange"`                                                                   // Space separated list of key exchange algorithms, ordered by preference.
	Macs                    string `json:"macs"`                                                                          // Space separated list of MACs (message authentication code) algorithms, ordered by preference.
}

type createSftpStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       sftpConfig         `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateSftpStorage
// @Summary Create Sftp storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createSftpStorageRequest true "Request body"
// @Router /storage/sftp [post]
func createSftpStorage() {}

type sharefileConfig struct {
	UploadCutoff string `default:"128Mi"                                                                                                                         json:"uploadCutoff"` // Cutoff for switching to multipart upload.
	RootFolderId string `example:""                                                                                                                              json:"rootFolderId"` // ID of the root folder.
	ChunkSize    string `default:"64Mi"                                                                                                                          json:"chunkSize"`    // Upload chunk size.
	Endpoint     string `json:"endpoint"`                                                                                                                                             // Endpoint for API calls.
	Encoding     string `default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot" json:"encoding"`     // The encoding for the backend.
}

type createSharefileStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       sharefileConfig    `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateSharefileStorage
// @Summary Create Sharefile storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createSharefileStorageRequest true "Request body"
// @Router /storage/sharefile [post]
func createSharefileStorage() {}

type siaConfig struct {
	ApiUrl      string `default:"http://127.0.0.1:9980"                               json:"apiUrl"`    // Sia daemon API URL, like http://sia.daemon.host:9980.
	ApiPassword string `json:"apiPassword"`                                                             // Sia Daemon API Password.
	UserAgent   string `default:"Sia-Agent"                                           json:"userAgent"` // Siad User Agent
	Encoding    string `default:"Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot" json:"encoding"`  // The encoding for the backend.
}

type createSiaStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       siaConfig          `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateSiaStorage
// @Summary Create Sia storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createSiaStorageRequest true "Request body"
// @Router /storage/sia [post]
func createSiaStorage() {}

type smbConfig struct {
	Host             string `json:"host"`                                                                                                                                // SMB server hostname to connect to.
	User             string `default:"$USER"                                                                                                    json:"user"`             // SMB username.
	Port             int    `default:"445"                                                                                                      json:"port"`             // SMB port number.
	Pass             string `json:"pass"`                                                                                                                                // SMB password.
	Domain           string `default:"WORKGROUP"                                                                                                json:"domain"`           // Domain name for NTLM authentication.
	Spn              string `json:"spn"`                                                                                                                                 // Service principal name.
	IdleTimeout      string `default:"1m0s"                                                                                                     json:"idleTimeout"`      // Max time before closing idle connections.
	HideSpecialShare bool   `default:"true"                                                                                                     json:"hideSpecialShare"` // Hide special shares (e.g. print$) which users aren't supposed to access.
	CaseInsensitive  bool   `default:"true"                                                                                                     json:"caseInsensitive"`  // Whether the server is configured to be case-insensitive.
	Encoding         string `default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot" json:"encoding"`         // The encoding for the backend.
}

type createSmbStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       smbConfig          `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateSmbStorage
// @Summary Create Smb storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createSmbStorageRequest true "Request body"
// @Router /storage/smb [post]
func createSmbStorage() {}

type storjExistingConfig struct {
	AccessGrant string `json:"accessGrant"` // Access grant.
}

type createStorjExistingStorageRequest struct {
	Name         string              `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string              `json:"path"`                      // Path of the storage
	Config       storjExistingConfig `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig  `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateStorjExistingStorage
// @Summary Create Storj storage with existing - Use an existing access grant.
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createStorjExistingStorageRequest true "Request body"
// @Router /storage/storj/existing [post]
func createStorjExistingStorage() {}

type storjNewConfig struct {
	SatelliteAddress string `default:"us1.storj.io" example:"us1.storj.io" json:"satelliteAddress"` // Satellite address.
	ApiKey           string `json:"apiKey"`                                                         // API key.
	Passphrase       string `json:"passphrase"`                                                     // Encryption passphrase.
}

type createStorjNewStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       storjNewConfig     `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateStorjNewStorage
// @Summary Create Storj storage with new - Create a new access grant from satellite address, API key, and passphrase.
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createStorjNewStorageRequest true "Request body"
// @Router /storage/storj/new [post]
func createStorjNewStorage() {}

type sugarsyncConfig struct {
	AppId               string `json:"appId"`                                          // Sugarsync App ID.
	AccessKeyId         string `json:"accessKeyId"`                                    // Sugarsync Access Key ID.
	PrivateAccessKey    string `json:"privateAccessKey"`                               // Sugarsync Private Access Key.
	HardDelete          bool   `default:"false"                     json:"hardDelete"` // Permanently delete files if true
	RefreshToken        string `json:"refreshToken"`                                   // Sugarsync refresh token.
	Authorization       string `json:"authorization"`                                  // Sugarsync authorization.
	AuthorizationExpiry string `json:"authorizationExpiry"`                            // Sugarsync authorization expiry.
	User                string `json:"user"`                                           // Sugarsync user.
	RootId              string `json:"rootId"`                                         // Sugarsync root id.
	DeletedId           string `json:"deletedId"`                                      // Sugarsync deleted folder id.
	Encoding            string `default:"Slash,Ctl,InvalidUtf8,Dot" json:"encoding"`   // The encoding for the backend.
}

type createSugarsyncStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       sugarsyncConfig    `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateSugarsyncStorage
// @Summary Create Sugarsync storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createSugarsyncStorageRequest true "Request body"
// @Router /storage/sugarsync [post]
func createSugarsyncStorage() {}

type swiftConfig struct {
	EnvAuth                     bool   `default:"false"                                    example:"false"          json:"envAuth"`      // Get swift credentials from environment variables in standard OpenStack form.
	User                        string `json:"user"`                                                                                     // User name to log in (OS_USERNAME).
	Key                         string `json:"key"`                                                                                      // API key or password (OS_PASSWORD).
	Auth                        string `example:"https://auth.api.rackspacecloud.com/v1.0" json:"auth"`                                  // Authentication URL for server (OS_AUTH_URL).
	UserId                      string `json:"userId"`                                                                                   // User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID).
	Domain                      string `json:"domain"`                                                                                   // User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME)
	Tenant                      string `json:"tenant"`                                                                                   // Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME).
	TenantId                    string `json:"tenantId"`                                                                                 // Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID).
	TenantDomain                string `json:"tenantDomain"`                                                                             // Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME).
	Region                      string `json:"region"`                                                                                   // Region name - optional (OS_REGION_NAME).
	StorageUrl                  string `json:"storageUrl"`                                                                               // Storage URL - optional (OS_STORAGE_URL).
	AuthToken                   string `json:"authToken"`                                                                                // Auth Token from alternate authentication - optional (OS_AUTH_TOKEN).
	ApplicationCredentialId     string `json:"applicationCredentialId"`                                                                  // Application Credential ID (OS_APPLICATION_CREDENTIAL_ID).
	ApplicationCredentialName   string `json:"applicationCredentialName"`                                                                // Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME).
	ApplicationCredentialSecret string `json:"applicationCredentialSecret"`                                                              // Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET).
	AuthVersion                 int    `default:"0"                                        json:"authVersion"`                           // AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).
	EndpointType                string `default:"public"                                   example:"public"         json:"endpointType"` // Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).
	LeavePartsOnError           bool   `default:"false"                                    json:"leavePartsOnError"`                     // If true avoid calling abort upload on a failure.
	StoragePolicy               string `example:""                                         json:"storagePolicy"`                         // The storage policy to use when creating a new container.
	ChunkSize                   string `default:"5Gi"                                      json:"chunkSize"`                             // Above this size files will be chunked into a _segments container.
	NoChunk                     bool   `default:"false"                                    json:"noChunk"`                               // Don't chunk files during streaming upload.
	NoLargeObjects              bool   `default:"false"                                    json:"noLargeObjects"`                        // Disable support for static and dynamic large objects
	Encoding                    string `default:"Slash,InvalidUtf8"                        json:"encoding"`                              // The encoding for the backend.
}

type createSwiftStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       swiftConfig        `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateSwiftStorage
// @Summary Create Swift storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createSwiftStorageRequest true "Request body"
// @Router /storage/swift [post]
func createSwiftStorage() {}

type unionConfig struct {
	Upstreams    string `json:"upstreams"`                     // List of space separated upstreams.
	ActionPolicy string `default:"epall"  json:"actionPolicy"` // Policy to choose upstream on ACTION category.
	CreatePolicy string `default:"epmfs"  json:"createPolicy"` // Policy to choose upstream on CREATE category.
	SearchPolicy string `default:"ff"     json:"searchPolicy"` // Policy to choose upstream on SEARCH category.
	CacheTime    int    `default:"120"    json:"cacheTime"`    // Cache time of usage and free space (in seconds).
	MinFreeSpace string `default:"1Gi"    json:"minFreeSpace"` // Minimum viable free space for lfs/eplfs policies.
}

type createUnionStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       unionConfig        `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateUnionStorage
// @Summary Create Union storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createUnionStorageRequest true "Request body"
// @Router /storage/union [post]
func createUnionStorage() {}

type uptoboxConfig struct {
	AccessToken string `json:"accessToken"`                                                                           // Your access token.
	Encoding    string `default:"Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot" json:"encoding"` // The encoding for the backend.
}

type createUptoboxStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       uptoboxConfig      `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateUptoboxStorage
// @Summary Create Uptobox storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createUptoboxStorageRequest true "Request body"
// @Router /storage/uptobox [post]
func createUptoboxStorage() {}

type webdavConfig struct {
	Url                string `json:"url"`                              // URL of http host to connect to.
	Vendor             string `example:"nextcloud"       json:"vendor"` // Name of the WebDAV site/service/software you are using.
	User               string `json:"user"`                             // User name.
	Pass               string `json:"pass"`                             // Password.
	BearerToken        string `json:"bearerToken"`                      // Bearer token instead of user/pass (e.g. a Macaroon).
	BearerTokenCommand string `json:"bearerTokenCommand"`               // Command to run to get a bearer token.
	Encoding           string `json:"encoding"`                         // The encoding for the backend.
	Headers            string `json:"headers"`                          // Set HTTP headers for all transactions.
}

type createWebdavStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       webdavConfig       `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateWebdavStorage
// @Summary Create Webdav storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createWebdavStorageRequest true "Request body"
// @Router /storage/webdav [post]
func createWebdavStorage() {}

type yandexConfig struct {
	ClientId     string `json:"clientId"`                                           // OAuth Client Id.
	ClientSecret string `json:"clientSecret"`                                       // OAuth Client Secret.
	Token        string `json:"token"`                                              // OAuth Access Token as a JSON blob.
	AuthUrl      string `json:"authUrl"`                                            // Auth server URL.
	TokenUrl     string `json:"tokenUrl"`                                           // Token server url.
	HardDelete   bool   `default:"false"                         json:"hardDelete"` // Delete files permanently rather than putting them into the trash.
	Encoding     string `default:"Slash,Del,Ctl,InvalidUtf8,Dot" json:"encoding"`   // The encoding for the backend.
}

type createYandexStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       yandexConfig       `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateYandexStorage
// @Summary Create Yandex storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createYandexStorageRequest true "Request body"
// @Router /storage/yandex [post]
func createYandexStorage() {}

type zohoConfig struct {
	ClientId     string `json:"clientId"`                               // OAuth Client Id.
	ClientSecret string `json:"clientSecret"`                           // OAuth Client Secret.
	Token        string `json:"token"`                                  // OAuth Access Token as a JSON blob.
	AuthUrl      string `json:"authUrl"`                                // Auth server URL.
	TokenUrl     string `json:"tokenUrl"`                               // Token server url.
	Region       string `example:"com"                 json:"region"`   // Zoho region to connect to.
	Encoding     string `default:"Del,Ctl,InvalidUtf8" json:"encoding"` // The encoding for the backend.
}

type createZohoStorageRequest struct {
	Name         string             `example:"my-storage" json:"name"` // Name of the storage, must be unique
	Path         string             `json:"path"`                      // Path of the storage
	Config       zohoConfig         `json:"config"`                    // config for the storage
	ClientConfig model.ClientConfig `json:"clientConfig"`              // config for underlying HTTP client
}

// @ID CreateZohoStorage
// @Summary Create Zoho storage
// @Tags Storage
// @Accept json
// @Produce json
// @Success 200 {object} model.Storage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Param request body createZohoStorageRequest true "Request body"
// @Router /storage/zoho [post]
func createZohoStorage() {}
