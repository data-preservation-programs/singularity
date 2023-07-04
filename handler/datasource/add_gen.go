// Code generated. DO NOT EDIT.
package datasource

type AcdRequest struct {
	SourcePath        string `validate:"required" json:"sourcePath"`           // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`    // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`       // Automatically rescan the source directory when this interval has passed from last successful scan
	AuthUrl           string `json:"authUrl"`                                  // Auth server URL.
	Checkpoint        string `json:"checkpoint"`                               // Checkpoint for internal polling (debug).
	ClientId          string `json:"clientId"`                                 // OAuth Client Id.
	ClientSecret      string `json:"clientSecret"`                             // OAuth Client Secret.
	Encoding          string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	TemplinkThreshold string `json:"templinkThreshold" default:"9Gi"`          // Files >= this size will be downloaded via their tempLink.
	Token             string `json:"token"`                                    // OAuth Access Token as a JSON blob.
	TokenUrl          string `json:"tokenUrl"`                                 // Token server url.
	UploadWaitPerGb   string `json:"uploadWaitPerGb" default:"3m0s"`           // Additional time per GiB to wait after a failed complete upload to see if it appears.
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
	SourcePath                 string `validate:"required" json:"sourcePath"`                                     // The path of the source to scan items
	DeleteAfterExport          bool   `validate:"required" json:"deleteAfterExport"`                              // Delete the source after exporting to CAR files
	RescanInterval             string `validate:"required" json:"rescanInterval"`                                 // Automatically rescan the source directory when this interval has passed from last successful scan
	AccessTier                 string `json:"accessTier"`                                                         // Access tier of blob: hot, cool or archive.
	Account                    string `json:"account"`                                                            // Azure Storage Account Name.
	ArchiveTierDelete          string `json:"archiveTierDelete" default:"false"`                                  // Delete archive tier blobs before overwriting.
	ChunkSize                  string `json:"chunkSize" default:"4Mi"`                                            // Upload chunk size.
	ClientCertificatePassword  string `json:"clientCertificatePassword"`                                          // Password for the certificate file (optional).
	ClientCertificatePath      string `json:"clientCertificatePath"`                                              // Path to a PEM or PKCS12 certificate file including the private key.
	ClientId                   string `json:"clientId"`                                                           // The ID of the client in use.
	ClientSecret               string `json:"clientSecret"`                                                       // One of the service principal's client secrets
	ClientSendCertificateChain string `json:"clientSendCertificateChain" default:"false"`                         // Send the certificate chain when using certificate auth.
	DisableChecksum            string `json:"disableChecksum" default:"false"`                                    // Don't store MD5 checksum with object metadata.
	Encoding                   string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"` // The encoding for the backend.
	Endpoint                   string `json:"endpoint"`                                                           // Endpoint for the service.
	EnvAuth                    string `json:"envAuth" default:"false"`                                            // Read credentials from runtime (environment variables, CLI or MSI).
	Key                        string `json:"key"`                                                                // Storage Account Shared Key.
	ListChunk                  string `json:"listChunk" default:"5000"`                                           // Size of blob list.
	MemoryPoolFlushTime        string `json:"memoryPoolFlushTime" default:"1m0s"`                                 // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap          string `json:"memoryPoolUseMmap" default:"false"`                                  // Whether to use mmap buffers in internal memory pool.
	MsiClientId                string `json:"msiClientId"`                                                        // Object ID of the user-assigned MSI to use, if any.
	MsiMiResId                 string `json:"msiMiResId"`                                                         // Azure resource ID of the user-assigned MSI to use, if any.
	MsiObjectId                string `json:"msiObjectId"`                                                        // Object ID of the user-assigned MSI to use, if any.
	NoCheckContainer           string `json:"noCheckContainer" default:"false"`                                   // If set, don't attempt to check the container exists or create it.
	NoHeadObject               string `json:"noHeadObject" default:"false"`                                       // If set, do not do HEAD before GET when getting objects.
	Password                   string `json:"password"`                                                           // The user's password
	PublicAccess               string `json:"publicAccess"`                                                       // Public access level of a container: blob or container.
	SasUrl                     string `json:"sasUrl"`                                                             // SAS URL for container level access only.
	ServicePrincipalFile       string `json:"servicePrincipalFile"`                                               // Path to file containing credentials for use with a service principal.
	Tenant                     string `json:"tenant"`                                                             // ID of the service principal's tenant. Also called its directory ID.
	UploadConcurrency          string `json:"uploadConcurrency" default:"16"`                                     // Concurrency for multipart uploads.
	UploadCutoff               string `json:"uploadCutoff"`                                                       // Cutoff for switching to chunked upload (<= 256 MiB) (deprecated).
	UseEmulator                string `json:"useEmulator" default:"false"`                                        // Uses local storage emulator if provided as 'true'.
	UseMsi                     string `json:"useMsi" default:"false"`                                             // Use a managed service identity to authenticate (only works in Azure).
	Username                   string `json:"username"`                                                           // User name (usually an email address)
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
	SourcePath           string `validate:"required" json:"sourcePath"`                             // The path of the source to scan items
	DeleteAfterExport    bool   `validate:"required" json:"deleteAfterExport"`                      // Delete the source after exporting to CAR files
	RescanInterval       string `validate:"required" json:"rescanInterval"`                         // Automatically rescan the source directory when this interval has passed from last successful scan
	Account              string `json:"account"`                                                    // Account ID or Application Key ID.
	ChunkSize            string `json:"chunkSize" default:"96Mi"`                                   // Upload chunk size.
	CopyCutoff           string `json:"copyCutoff" default:"4Gi"`                                   // Cutoff for switching to multipart copy.
	DisableChecksum      string `json:"disableChecksum" default:"false"`                            // Disable checksums for large (> upload cutoff) files.
	DownloadAuthDuration string `json:"downloadAuthDuration" default:"1w"`                          // Time before the authorization token will expire in s or suffix ms|s|m|h|d.
	DownloadUrl          string `json:"downloadUrl"`                                                // Custom endpoint for downloads.
	Encoding             string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	Endpoint             string `json:"endpoint"`                                                   // Endpoint for the service.
	HardDelete           string `json:"hardDelete" default:"false"`                                 // Permanently delete files on remote removal, otherwise hide files.
	Key                  string `json:"key"`                                                        // Application Key.
	MemoryPoolFlushTime  string `json:"memoryPoolFlushTime" default:"1m0s"`                         // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap    string `json:"memoryPoolUseMmap" default:"false"`                          // Whether to use mmap buffers in internal memory pool.
	TestMode             string `json:"testMode"`                                                   // A flag string for X-Bz-Test-Mode header for debugging.
	UploadCutoff         string `json:"uploadCutoff" default:"200Mi"`                               // Cutoff for switching to chunked upload.
	VersionAt            string `json:"versionAt" default:"off"`                                    // Show file versions as they were at the specified time.
	Versions             string `json:"versions" default:"false"`                                   // Include old versions in directory listings.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                                        // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                                 // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                                    // Automatically rescan the source directory when this interval has passed from last successful scan
	AccessToken       string `json:"accessToken"`                                                           // Box App Primary Access Token
	AuthUrl           string `json:"authUrl"`                                                               // Auth server URL.
	BoxConfigFile     string `json:"boxConfigFile"`                                                         // Box App config.json location
	BoxSubType        string `json:"boxSubType" default:"user"`                                             //
	ClientId          string `json:"clientId"`                                                              // OAuth Client Id.
	ClientSecret      string `json:"clientSecret"`                                                          // OAuth Client Secret.
	CommitRetries     string `json:"commitRetries" default:"100"`                                           // Max number of times to try committing a multipart file.
	Encoding          string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
	ListChunk         string `json:"listChunk" default:"1000"`                                              // Size of listing chunk 1-1000.
	OwnedBy           string `json:"ownedBy"`                                                               // Only show items owned by the login (email address) passed in.
	RootFolderId      string `json:"rootFolderId" default:"0"`                                              // Fill in for rclone to use a non root folder as its starting point.
	Token             string `json:"token"`                                                                 // OAuth Access Token as a JSON blob.
	TokenUrl          string `json:"tokenUrl"`                                                              // Token server url.
	UploadCutoff      string `json:"uploadCutoff" default:"50Mi"`                                           // Cutoff for switching to multipart upload (>= 50 MiB).
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

type DriveRequest struct {
	SourcePath                string `validate:"required" json:"sourcePath"`             // The path of the source to scan items
	DeleteAfterExport         bool   `validate:"required" json:"deleteAfterExport"`      // Delete the source after exporting to CAR files
	RescanInterval            string `validate:"required" json:"rescanInterval"`         // Automatically rescan the source directory when this interval has passed from last successful scan
	AcknowledgeAbuse          string `json:"acknowledgeAbuse" default:"false"`           // Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
	AllowImportNameChange     string `json:"allowImportNameChange" default:"false"`      // Allow the filetype to change when uploading Google docs.
	AlternateExport           string `json:"alternateExport" default:"false"`            // Deprecated: No longer needed.
	AuthOwnerOnly             string `json:"authOwnerOnly" default:"false"`              // Only consider files owned by the authenticated user.
	AuthUrl                   string `json:"authUrl"`                                    // Auth server URL.
	ChunkSize                 string `json:"chunkSize" default:"8Mi"`                    // Upload chunk size.
	ClientId                  string `json:"clientId"`                                   // Google Application Client Id
	ClientSecret              string `json:"clientSecret"`                               // OAuth Client Secret.
	CopyShortcutContent       string `json:"copyShortcutContent" default:"false"`        // Server side copy contents of shortcuts instead of the shortcut.
	DisableHttp2              string `json:"disableHttp2" default:"true"`                // Disable drive using http2.
	Encoding                  string `json:"encoding" default:"InvalidUtf8"`             // The encoding for the backend.
	ExportFormats             string `json:"exportFormats" default:"docx,xlsx,pptx,svg"` // Comma separated list of preferred formats for downloading Google docs.
	Formats                   string `json:"formats"`                                    // Deprecated: See export_formats.
	Impersonate               string `json:"impersonate"`                                // Impersonate this user when using a service account.
	ImportFormats             string `json:"importFormats"`                              // Comma separated list of preferred formats for uploading Google docs.
	KeepRevisionForever       string `json:"keepRevisionForever" default:"false"`        // Keep new head revision of each file forever.
	ListChunk                 string `json:"listChunk" default:"1000"`                   // Size of listing chunk 100-1000, 0 to disable.
	PacerBurst                string `json:"pacerBurst" default:"100"`                   // Number of API calls to allow without sleeping.
	PacerMinSleep             string `json:"pacerMinSleep" default:"100ms"`              // Minimum time to sleep between API calls.
	ResourceKey               string `json:"resourceKey"`                                // Resource key for accessing a link-shared file.
	RootFolderId              string `json:"rootFolderId"`                               // ID of the root folder.
	Scope                     string `json:"scope"`                                      // Scope that rclone should use when requesting access from drive.
	ServerSideAcrossConfigs   string `json:"serverSideAcrossConfigs" default:"false"`    // Allow server-side operations (e.g. copy) to work across different drive configs.
	ServiceAccountCredentials string `json:"serviceAccountCredentials"`                  // Service Account Credentials JSON blob.
	ServiceAccountFile        string `json:"serviceAccountFile"`                         // Service Account Credentials JSON file path.
	SharedWithMe              string `json:"sharedWithMe" default:"false"`               // Only show files that are shared with me.
	SizeAsQuota               string `json:"sizeAsQuota" default:"false"`                // Show sizes as storage quota usage, not actual size.
	SkipChecksumGphotos       string `json:"skipChecksumGphotos" default:"false"`        // Skip MD5 checksum on Google photos and videos only.
	SkipDanglingShortcuts     string `json:"skipDanglingShortcuts" default:"false"`      // If set skip dangling shortcut files.
	SkipGdocs                 string `json:"skipGdocs" default:"false"`                  // Skip google documents in all listings.
	SkipShortcuts             string `json:"skipShortcuts" default:"false"`              // If set skip shortcut files.
	StarredOnly               string `json:"starredOnly" default:"false"`                // Only show files that are starred.
	StopOnDownloadLimit       string `json:"stopOnDownloadLimit" default:"false"`        // Make download limit errors be fatal.
	StopOnUploadLimit         string `json:"stopOnUploadLimit" default:"false"`          // Make upload limit errors be fatal.
	TeamDrive                 string `json:"teamDrive"`                                  // ID of the Shared Drive (Team Drive).
	Token                     string `json:"token"`                                      // OAuth Access Token as a JSON blob.
	TokenUrl                  string `json:"tokenUrl"`                                   // Token server url.
	TrashedOnly               string `json:"trashedOnly" default:"false"`                // Only show files that are in the trash.
	UploadCutoff              string `json:"uploadCutoff" default:"8Mi"`                 // Cutoff for switching to chunked upload.
	UseCreatedDate            string `json:"useCreatedDate" default:"false"`             // Use file created date instead of modified date.
	UseSharedDate             string `json:"useSharedDate" default:"false"`              // Use date file was shared instead of modified date.
	UseTrash                  string `json:"useTrash" default:"true"`                    // Send files to the trash instead of deleting permanently.
	V2DownloadMinSize         string `json:"v2DownloadMinSize" default:"off"`            // If Object's are greater, use drive v2 API to download.
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
	SourcePath         string `validate:"required" json:"sourcePath"`                                    // The path of the source to scan items
	DeleteAfterExport  bool   `validate:"required" json:"deleteAfterExport"`                             // Delete the source after exporting to CAR files
	RescanInterval     string `validate:"required" json:"rescanInterval"`                                // Automatically rescan the source directory when this interval has passed from last successful scan
	AuthUrl            string `json:"authUrl"`                                                           // Auth server URL.
	BatchCommitTimeout string `json:"batchCommitTimeout" default:"10m0s"`                                // Max time to wait for a batch to finish committing
	BatchMode          string `json:"batchMode" default:"sync"`                                          // Upload file batching sync|async|off.
	BatchSize          string `json:"batchSize" default:"0"`                                             // Max number of files in upload batch.
	BatchTimeout       string `json:"batchTimeout" default:"0s"`                                         // Max time to allow an idle upload batch before uploading.
	ChunkSize          string `json:"chunkSize" default:"48Mi"`                                          // Upload chunk size (< 150Mi).
	ClientId           string `json:"clientId"`                                                          // OAuth Client Id.
	ClientSecret       string `json:"clientSecret"`                                                      // OAuth Client Secret.
	Encoding           string `json:"encoding" default:"Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
	Impersonate        string `json:"impersonate"`                                                       // Impersonate this user when using a business account.
	SharedFiles        string `json:"sharedFiles" default:"false"`                                       // Instructs rclone to work on individual shared files.
	SharedFolders      string `json:"sharedFolders" default:"false"`                                     // Instructs rclone to work on shared folders.
	Token              string `json:"token"`                                                             // OAuth Access Token as a JSON blob.
	TokenUrl           string `json:"tokenUrl"`                                                          // Token server url.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                                                                                                // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                                                                                         // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                                                                                            // Automatically rescan the source directory when this interval has passed from last successful scan
	ApiKey            string `json:"apiKey"`                                                                                                                        // Your API Key, get it from https://1fichier.com/console/params.pl.
	Encoding          string `json:"encoding" default:"Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
	FilePassword      string `json:"filePassword"`                                                                                                                  // If you want to download a shared file that is password protected, add this parameter.
	FolderPassword    string `json:"folderPassword"`                                                                                                                // If you want to list the files in a shared folder that is password protected, add this parameter.
	SharedFolder      string `json:"sharedFolder"`                                                                                                                  // If you want to download a shared folder, add this parameter.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                   // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`            // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`               // Automatically rescan the source directory when this interval has passed from last successful scan
	Encoding          string `json:"encoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	PermanentToken    string `json:"permanentToken"`                                   // Permanent Authentication Token.
	RootFolderId      string `json:"rootFolderId"`                                     // ID of the root folder.
	Token             string `json:"token"`                                            // Session Token.
	TokenExpiry       string `json:"tokenExpiry"`                                      // Token expiry time.
	Url               string `json:"url"`                                              // URL of the Enterprise File Fabric to connect to.
	Version           string `json:"version"`                                          // Version read from the file fabric.
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
	SourcePath         string `validate:"required" json:"sourcePath"`                  // The path of the source to scan items
	DeleteAfterExport  bool   `validate:"required" json:"deleteAfterExport"`           // Delete the source after exporting to CAR files
	RescanInterval     string `validate:"required" json:"rescanInterval"`              // Automatically rescan the source directory when this interval has passed from last successful scan
	AskPassword        string `json:"askPassword" default:"false"`                     // Allow asking for FTP password when needed.
	CloseTimeout       string `json:"closeTimeout" default:"1m0s"`                     // Maximum time to wait for a response to close.
	Concurrency        string `json:"concurrency" default:"0"`                         // Maximum number of FTP simultaneous connections, 0 for unlimited.
	DisableEpsv        string `json:"disableEpsv" default:"false"`                     // Disable using EPSV even if server advertises support.
	DisableMlsd        string `json:"disableMlsd" default:"false"`                     // Disable using MLSD even if server advertises support.
	DisableTls13       string `json:"disableTls13" default:"false"`                    // Disable TLS 1.3 (workaround for FTP servers with buggy TLS)
	DisableUtf8        string `json:"disableUtf8" default:"false"`                     // Disable using UTF-8 even if server advertises support.
	Encoding           string `json:"encoding" default:"Slash,Del,Ctl,RightSpace,Dot"` // The encoding for the backend.
	ExplicitTls        string `json:"explicitTls" default:"false"`                     // Use Explicit FTPS (FTP over TLS).
	ForceListHidden    string `json:"forceListHidden" default:"false"`                 // Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.
	Host               string `json:"host"`                                            // FTP host to connect to.
	IdleTimeout        string `json:"idleTimeout" default:"1m0s"`                      // Max time before closing idle connections.
	NoCheckCertificate string `json:"noCheckCertificate" default:"false"`              // Do not verify the TLS certificate of the server.
	Pass               string `json:"pass"`                                            // FTP password.
	Port               string `json:"port" default:"21"`                               // FTP port number.
	ShutTimeout        string `json:"shutTimeout" default:"1m0s"`                      // Maximum time to wait for data connection closing status.
	Tls                string `json:"tls" default:"false"`                             // Use Implicit FTPS (FTP over TLS).
	TlsCacheSize       string `json:"tlsCacheSize" default:"32"`                       // Size of TLS session cache for all control and data connections.
	User               string `json:"user" default:"shane"`                            // FTP username.
	WritingMdtm        string `json:"writingMdtm" default:"false"`                     // Use MDTM to set modification time (VsFtpd quirk)
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
	SourcePath                string `validate:"required" json:"sourcePath"`                // The path of the source to scan items
	DeleteAfterExport         bool   `validate:"required" json:"deleteAfterExport"`         // Delete the source after exporting to CAR files
	RescanInterval            string `validate:"required" json:"rescanInterval"`            // Automatically rescan the source directory when this interval has passed from last successful scan
	Anonymous                 string `json:"anonymous" default:"false"`                     // Access public buckets and objects without credentials.
	AuthUrl                   string `json:"authUrl"`                                       // Auth server URL.
	BucketAcl                 string `json:"bucketAcl"`                                     // Access Control List for new buckets.
	BucketPolicyOnly          string `json:"bucketPolicyOnly" default:"false"`              // Access checks should use bucket-level IAM policies.
	ClientId                  string `json:"clientId"`                                      // OAuth Client Id.
	ClientSecret              string `json:"clientSecret"`                                  // OAuth Client Secret.
	Decompress                string `json:"decompress" default:"false"`                    // If set this will decompress gzip encoded objects.
	Encoding                  string `json:"encoding" default:"Slash,CrLf,InvalidUtf8,Dot"` // The encoding for the backend.
	Endpoint                  string `json:"endpoint"`                                      // Endpoint for the service.
	EnvAuth                   string `json:"envAuth" default:"false"`                       // Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
	Location                  string `json:"location"`                                      // Location for the newly created buckets.
	NoCheckBucket             string `json:"noCheckBucket" default:"false"`                 // If set, don't attempt to check the bucket exists or create it.
	ObjectAcl                 string `json:"objectAcl"`                                     // Access Control List for new objects.
	ProjectNumber             string `json:"projectNumber"`                                 // Project number.
	ServiceAccountCredentials string `json:"serviceAccountCredentials"`                     // Service Account Credentials JSON blob.
	ServiceAccountFile        string `json:"serviceAccountFile"`                            // Service Account Credentials JSON file path.
	StorageClass              string `json:"storageClass"`                                  // The storage class to use when storing objects in Google Cloud Storage.
	Token                     string `json:"token"`                                         // OAuth Access Token as a JSON blob.
	TokenUrl                  string `json:"tokenUrl"`                                      // Token server url.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`         // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`            // Automatically rescan the source directory when this interval has passed from last successful scan
	AuthUrl           string `json:"authUrl"`                                       // Auth server URL.
	ClientId          string `json:"clientId"`                                      // OAuth Client Id.
	ClientSecret      string `json:"clientSecret"`                                  // OAuth Client Secret.
	Encoding          string `json:"encoding" default:"Slash,CrLf,InvalidUtf8,Dot"` // The encoding for the backend.
	IncludeArchived   string `json:"includeArchived" default:"false"`               // Also view and download archived media.
	ReadOnly          string `json:"readOnly" default:"false"`                      // Set to make the Google Photos backend read only.
	ReadSize          string `json:"readSize" default:"false"`                      // Set to read the size of media items.
	StartYear         string `json:"startYear" default:"2000"`                      // Year limits the photos to be downloaded to those which are uploaded after the given year.
	Token             string `json:"token"`                                         // OAuth Access Token as a JSON blob.
	TokenUrl          string `json:"tokenUrl"`                                      // Token server url.
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
	SourcePath             string `validate:"required" json:"sourcePath"`                         // The path of the source to scan items
	DeleteAfterExport      bool   `validate:"required" json:"deleteAfterExport"`                  // Delete the source after exporting to CAR files
	RescanInterval         string `validate:"required" json:"rescanInterval"`                     // Automatically rescan the source directory when this interval has passed from last successful scan
	DataTransferProtection string `json:"dataTransferProtection"`                                 // Kerberos data transfer protection: authentication|integrity|privacy.
	Encoding               string `json:"encoding" default:"Slash,Colon,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	Namenode               string `json:"namenode"`                                               // Hadoop name node and port.
	ServicePrincipalName   string `json:"servicePrincipalName"`                                   // Kerberos service principal name for the namenode.
	Username               string `json:"username"`                                               // Hadoop user name.
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
	SourcePath                 string `validate:"required" json:"sourcePath"`                        // The path of the source to scan items
	DeleteAfterExport          bool   `validate:"required" json:"deleteAfterExport"`                 // Delete the source after exporting to CAR files
	RescanInterval             string `validate:"required" json:"rescanInterval"`                    // Automatically rescan the source directory when this interval has passed from last successful scan
	AuthUrl                    string `json:"authUrl"`                                               // Auth server URL.
	ChunkSize                  string `json:"chunkSize" default:"48Mi"`                              // Chunksize for chunked uploads.
	ClientId                   string `json:"clientId"`                                              // OAuth Client Id.
	ClientSecret               string `json:"clientSecret"`                                          // OAuth Client Secret.
	DisableFetchingMemberCount string `json:"disableFetchingMemberCount" default:"false"`            // Do not fetch number of objects in directories unless it is absolutely necessary.
	Encoding                   string `json:"encoding" default:"Slash,Dot"`                          // The encoding for the backend.
	Endpoint                   string `json:"endpoint" default:"https://api.hidrive.strato.com/2.1"` // Endpoint for the service.
	RootPrefix                 string `json:"rootPrefix" default:"/"`                                // The root/parent folder for all paths.
	ScopeAccess                string `json:"scopeAccess" default:"rw"`                              // Access permissions that rclone should use when requesting access from HiDrive.
	ScopeRole                  string `json:"scopeRole" default:"user"`                              // User-level that rclone should use when requesting access from HiDrive.
	Token                      string `json:"token"`                                                 // OAuth Access Token as a JSON blob.
	TokenUrl                   string `json:"tokenUrl"`                                              // Token server url.
	UploadConcurrency          string `json:"uploadConcurrency" default:"4"`                         // Concurrency for chunked uploads.
	UploadCutoff               string `json:"uploadCutoff" default:"96Mi"`                           // Cutoff/Threshold for chunked uploads.
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
	SourcePath        string `validate:"required" json:"sourcePath"`        // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"` // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`    // Automatically rescan the source directory when this interval has passed from last successful scan
	Headers           string `json:"headers"`                               // Set HTTP headers for all transactions.
	NoHead            string `json:"noHead" default:"false"`                // Don't use HEAD requests.
	NoSlash           string `json:"noSlash" default:"false"`               // Set this if the site doesn't end directories with /.
	Url               string `json:"url"`                                   // URL of HTTP host to connect to.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                             // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                      // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                         // Automatically rescan the source directory when this interval has passed from last successful scan
	AccessKeyId       string `json:"accessKeyId"`                                                // IAS3 Access Key.
	DisableChecksum   string `json:"disableChecksum" default:"true"`                             // Don't ask the server to test against MD5 checksum calculated by rclone.
	Encoding          string `json:"encoding" default:"Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	Endpoint          string `json:"endpoint" default:"https://s3.us.archive.org"`               // IAS3 Endpoint.
	FrontEndpoint     string `json:"frontEndpoint" default:"https://archive.org"`                // Host of InternetArchive Frontend.
	SecretAccessKey   string `json:"secretAccessKey"`                                            // IAS3 Secret Key (password).
	WaitArchive       string `json:"waitArchive" default:"0s"`                                   // Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                                                                 // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                                                          // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                                                             // Automatically rescan the source directory when this interval has passed from last successful scan
	Encoding          string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	HardDelete        string `json:"hardDelete" default:"false"`                                                                     // Delete files permanently rather than putting them into the trash.
	Md5MemoryLimit    string `json:"md5MemoryLimit" default:"10Mi"`                                                                  // Files bigger than this will be cached on disk to calculate the MD5 if required.
	NoVersions        string `json:"noVersions" default:"false"`                                                                     // Avoid server side versioning by deleting files and recreating files instead of overwriting them.
	TrashedOnly       string `json:"trashedOnly" default:"false"`                                                                    // Only show files that are in the trash.
	UploadResumeLimit string `json:"uploadResumeLimit" default:"10Mi"`                                                               // Files bigger than this can be resumed if the upload fail's.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                             // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                      // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                         // Automatically rescan the source directory when this interval has passed from last successful scan
	Encoding          string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	Endpoint          string `json:"endpoint"`                                                   // The Koofr API endpoint to use.
	Mountid           string `json:"mountid"`                                                    // Mount ID of the mount to use.
	Password          string `json:"password"`                                                   // Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password).
	Provider          string `json:"provider"`                                                   // Choose your storage provider.
	Setmtime          string `json:"setmtime" default:"true"`                                    // Does the backend support setting modification time.
	User              string `json:"user"`                                                       // Your user name.
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
	SourcePath           string `validate:"required" json:"sourcePath"`        // The path of the source to scan items
	DeleteAfterExport    bool   `validate:"required" json:"deleteAfterExport"` // Delete the source after exporting to CAR files
	RescanInterval       string `validate:"required" json:"rescanInterval"`    // Automatically rescan the source directory when this interval has passed from last successful scan
	CaseInsensitive      string `json:"caseInsensitive" default:"false"`       // Force the filesystem to report itself as case insensitive.
	CaseSensitive        string `json:"caseSensitive" default:"false"`         // Force the filesystem to report itself as case sensitive.
	CopyLinks            string `json:"copyLinks" default:"false"`             // Follow symlinks and copy the pointed to item.
	Encoding             string `json:"encoding" default:"Slash,Dot"`          // The encoding for the backend.
	Links                string `json:"links" default:"false"`                 // Translate symlinks to/from regular files with a '.rclonelink' extension.
	NoCheckUpdated       string `json:"noCheckUpdated" default:"false"`        // Don't check to see if the files change during upload.
	NoPreallocate        string `json:"noPreallocate" default:"false"`         // Disable preallocation of disk space for transferred files.
	NoSetModtime         string `json:"noSetModtime" default:"false"`          // Disable setting modtime.
	NoSparse             string `json:"noSparse" default:"false"`              // Disable sparse files for multi-thread downloads.
	Nounc                string `json:"nounc" default:"false"`                 // Disable UNC (long path names) conversion on Windows.
	OneFileSystem        string `json:"oneFileSystem" default:"false"`         // Don't cross filesystem boundaries (unix/macOS only).
	SkipLinks            string `json:"skipLinks" default:"false"`             // Don't warn about skipped symlinks.
	UnicodeNormalization string `json:"unicodeNormalization" default:"false"`  // Apply unicode NFC normalization to paths and filenames.
	ZeroSizeLinks        string `json:"zeroSizeLinks" default:"false"`         // Assume the Stat size of links is zero (and read them instead) (deprecated).
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
	SourcePath          string `validate:"required" json:"sourcePath"`                                                                           // The path of the source to scan items
	DeleteAfterExport   bool   `validate:"required" json:"deleteAfterExport"`                                                                    // Delete the source after exporting to CAR files
	RescanInterval      string `validate:"required" json:"rescanInterval"`                                                                       // Automatically rescan the source directory when this interval has passed from last successful scan
	CheckHash           string `json:"checkHash" default:"true"`                                                                                 // What should copy do if file checksum is mismatched or invalid.
	Encoding            string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	Pass                string `json:"pass"`                                                                                                     // Password.
	Quirks              string `json:"quirks"`                                                                                                   // Comma separated list of internal maintenance flags.
	SpeedupEnable       string `json:"speedupEnable" default:"true"`                                                                             // Skip full upload if there is another file with same data hash.
	SpeedupFilePatterns string `json:"speedupFilePatterns" default:"*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf"`                             // Comma separated list of file name patterns eligible for speedup (put by hash).
	SpeedupMaxDisk      string `json:"speedupMaxDisk" default:"3Gi"`                                                                             // This option allows you to disable speedup (put by hash) for large files.
	SpeedupMaxMemory    string `json:"speedupMaxMemory" default:"32Mi"`                                                                          // Files larger than the size given below will always be hashed on disk.
	User                string `json:"user"`                                                                                                     // User name (usually email).
	UserAgent           string `json:"userAgent"`                                                                                                // HTTP user agent used internally by client.
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
	SourcePath        string `validate:"required" json:"sourcePath"`           // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`    // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`       // Automatically rescan the source directory when this interval has passed from last successful scan
	Debug             string `json:"debug" default:"false"`                    // Output more debug from Mega.
	Encoding          string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	HardDelete        string `json:"hardDelete" default:"false"`               // Delete files permanently rather than putting them into the trash.
	Pass              string `json:"pass"`                                     // Password.
	UseHttps          string `json:"useHttps" default:"false"`                 // Use HTTPS for transfers.
	User              string `json:"user"`                                     // User name.
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

type NetstorageRequest struct {
	SourcePath        string `validate:"required" json:"sourcePath"`        // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"` // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`    // Automatically rescan the source directory when this interval has passed from last successful scan
	Account           string `json:"account"`                               // Set the NetStorage account name
	Host              string `json:"host"`                                  // Domain+path of NetStorage host to connect to.
	Protocol          string `json:"protocol" default:"https"`              // Select between HTTP or HTTPS protocol.
	Secret            string `json:"secret"`                                // Set the NetStorage account secret/G2O key for authentication.
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
	SourcePath              string `validate:"required" json:"sourcePath"`                                                                                                                      // The path of the source to scan items
	DeleteAfterExport       bool   `validate:"required" json:"deleteAfterExport"`                                                                                                               // Delete the source after exporting to CAR files
	RescanInterval          string `validate:"required" json:"rescanInterval"`                                                                                                                  // Automatically rescan the source directory when this interval has passed from last successful scan
	AccessScopes            string `json:"accessScopes" default:"Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"`                                  // Set scopes to be requested by rclone.
	AuthUrl                 string `json:"authUrl"`                                                                                                                                             // Auth server URL.
	ChunkSize               string `json:"chunkSize" default:"10Mi"`                                                                                                                            // Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
	ClientId                string `json:"clientId"`                                                                                                                                            // OAuth Client Id.
	ClientSecret            string `json:"clientSecret"`                                                                                                                                        // OAuth Client Secret.
	DisableSitePermission   string `json:"disableSitePermission" default:"false"`                                                                                                               // Disable the request for Sites.Read.All permission.
	DriveId                 string `json:"driveId"`                                                                                                                                             // The ID of the drive to use.
	DriveType               string `json:"driveType"`                                                                                                                                           // The type of the drive (personal | business | documentLibrary).
	Encoding                string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
	ExposeOnenoteFiles      string `json:"exposeOnenoteFiles" default:"false"`                                                                                                                  // Set to make OneNote files show up in directory listings.
	HashType                string `json:"hashType" default:"auto"`                                                                                                                             // Specify the hash in use for the backend.
	LinkPassword            string `json:"linkPassword"`                                                                                                                                        // Set the password for links created by the link command.
	LinkScope               string `json:"linkScope" default:"anonymous"`                                                                                                                       // Set the scope of the links created by the link command.
	LinkType                string `json:"linkType" default:"view"`                                                                                                                             // Set the type of the links created by the link command.
	ListChunk               string `json:"listChunk" default:"1000"`                                                                                                                            // Size of listing chunk.
	NoVersions              string `json:"noVersions" default:"false"`                                                                                                                          // Remove all versions on modifying operations.
	Region                  string `json:"region" default:"global"`                                                                                                                             // Choose national cloud region for OneDrive.
	RootFolderId            string `json:"rootFolderId"`                                                                                                                                        // ID of the root folder.
	ServerSideAcrossConfigs string `json:"serverSideAcrossConfigs" default:"false"`                                                                                                             // Allow server-side operations (e.g. copy) to work across different onedrive configs.
	Token                   string `json:"token"`                                                                                                                                               // OAuth Access Token as a JSON blob.
	TokenUrl                string `json:"tokenUrl"`                                                                                                                                            // Token server url.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                                                                                                                   // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                                                                                                            // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                                                                                                               // Automatically rescan the source directory when this interval has passed from last successful scan
	ChunkSize         string `json:"chunkSize" default:"10Mi"`                                                                                                                         // Files will be uploaded in chunks this size.
	Encoding          string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot"` // The encoding for the backend.
	Password          string `json:"password"`                                                                                                                                         // Password.
	Username          string `json:"username"`                                                                                                                                         // Username.
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
	SourcePath           string `validate:"required" json:"sourcePath"`           // The path of the source to scan items
	DeleteAfterExport    bool   `validate:"required" json:"deleteAfterExport"`    // Delete the source after exporting to CAR files
	RescanInterval       string `validate:"required" json:"rescanInterval"`       // Automatically rescan the source directory when this interval has passed from last successful scan
	ChunkSize            string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	Compartment          string `json:"compartment"`                              // Object storage compartment OCID
	ConfigFile           string `json:"configFile" default:"~/.oci/config"`       // Path to OCI config file
	ConfigProfile        string `json:"configProfile" default:"Default"`          // Profile name inside the oci config file
	CopyCutoff           string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	CopyTimeout          string `json:"copyTimeout" default:"1m0s"`               // Timeout for copy.
	DisableChecksum      string `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	Encoding             string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	Endpoint             string `json:"endpoint"`                                 // Endpoint for Object storage API.
	LeavePartsOnError    string `json:"leavePartsOnError" default:"false"`        // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	Namespace            string `json:"namespace"`                                // Object storage namespace
	NoCheckBucket        string `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	Provider             string `json:"provider" default:"env_auth"`              // Choose your Auth Provider
	Region               string `json:"region"`                                   // Object storage Region
	SseCustomerAlgorithm string `json:"sseCustomerAlgorithm"`                     // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
	SseCustomerKey       string `json:"sseCustomerKey"`                           // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKeyFile   string `json:"sseCustomerKeyFile"`                       // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKeySha256 string `json:"sseCustomerKeySha256"`                     // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseKmsKeyId          string `json:"sseKmsKeyId"`                              // if using using your own master key in vault, this header specifies the
	StorageTier          string `json:"storageTier" default:"Standard"`           // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	UploadConcurrency    string `json:"uploadConcurrency" default:"10"`           // Concurrency for multipart uploads.
	UploadCutoff         string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                             // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                      // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                         // Automatically rescan the source directory when this interval has passed from last successful scan
	AuthUrl           string `json:"authUrl"`                                                    // Auth server URL.
	ClientId          string `json:"clientId"`                                                   // OAuth Client Id.
	ClientSecret      string `json:"clientSecret"`                                               // OAuth Client Secret.
	Encoding          string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	Hostname          string `json:"hostname" default:"api.pcloud.com"`                          // Hostname to connect to.
	Password          string `json:"password"`                                                   // Your pcloud password.
	RootFolderId      string `json:"rootFolderId" default:"d0"`                                  // Fill in for rclone to use a non root folder as its starting point.
	Token             string `json:"token"`                                                      // OAuth Access Token as a JSON blob.
	TokenUrl          string `json:"tokenUrl"`                                                   // Token server url.
	Username          string `json:"username"`                                                   // Your pcloud username.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                                         // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                                  // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                                     // Automatically rescan the source directory when this interval has passed from last successful scan
	ApiKey            string `json:"apiKey"`                                                                 // API Key.
	Encoding          string `json:"encoding" default:"Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                             // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                      // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                         // Automatically rescan the source directory when this interval has passed from last successful scan
	Encoding          string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
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
	SourcePath        string `validate:"required" json:"sourcePath"`           // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`    // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`       // Automatically rescan the source directory when this interval has passed from last successful scan
	AccessKeyId       string `json:"accessKeyId"`                              // QingStor Access Key ID.
	ChunkSize         string `json:"chunkSize" default:"4Mi"`                  // Chunk size to use for uploading.
	ConnectionRetries string `json:"connectionRetries" default:"3"`            // Number of connection retries.
	Encoding          string `json:"encoding" default:"Slash,Ctl,InvalidUtf8"` // The encoding for the backend.
	Endpoint          string `json:"endpoint"`                                 // Enter an endpoint URL to connection QingStor API.
	EnvAuth           string `json:"envAuth" default:"false"`                  // Get QingStor credentials from runtime.
	SecretAccessKey   string `json:"secretAccessKey"`                          // QingStor Secret Access Key (password).
	UploadConcurrency string `json:"uploadConcurrency" default:"1"`            // Concurrency for multipart uploads.
	UploadCutoff      string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	Zone              string `json:"zone"`                                     // Zone to connect to.
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
	SourcePath            string `validate:"required" json:"sourcePath"`           // The path of the source to scan items
	DeleteAfterExport     bool   `validate:"required" json:"deleteAfterExport"`    // Delete the source after exporting to CAR files
	RescanInterval        string `validate:"required" json:"rescanInterval"`       // Automatically rescan the source directory when this interval has passed from last successful scan
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	Acl                   string `json:"acl"`                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl"`                                // Canned ACL used when creating buckets.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	Decompress            string `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	DisableChecksum       string `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	DisableHttp2          string `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	Endpoint              string `json:"endpoint"`                                 // Endpoint for S3 API.
	EnvAuth               string `json:"envAuth" default:"false"`                  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	ForcePathStyle        string `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	LeavePartsOnError     string `json:"leavePartsOnError" default:"false"`        // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	ListChunk             string `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	ListVersion           string `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	LocationConstraint    string `json:"locationConstraint"`                       // Location constraint - must be set to match the Region.
	MaxUploadParts        string `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     string `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoCheckBucket         string `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                string `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          string `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	NoSystemMetadata      string `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	Provider              string `json:"provider"`                                 // Choose your S3 provider.
	Region                string `json:"region"`                                   // Region to connect to.
	RequesterPays         string `json:"requesterPays" default:"false"`            // Enables requester pays option when interacting with S3 bucket.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	ServerSideEncryption  string `json:"serverSideEncryption"`                     // The server-side encryption algorithm used when storing this object in S3.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	SseCustomerAlgorithm  string `json:"sseCustomerAlgorithm"`                     // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	SseCustomerKey        string `json:"sseCustomerKey"`                           // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	SseCustomerKeyBase64  string `json:"sseCustomerKeyBase64"`                     // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	SseCustomerKeyMd5     string `json:"sseCustomerKeyMd5"`                        // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	SseKmsKeyId           string `json:"sseKmsKeyId"`                              // If using KMS ID you must provide the ARN of Key.
	StorageClass          string `json:"storageClass"`                             // The storage class to use when storing new objects in S3.
	StsEndpoint           string `json:"stsEndpoint"`                              // Endpoint for STS.
	UploadConcurrency     string `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	UseAccelerateEndpoint string `json:"useAccelerateEndpoint" default:"false"`    // If true use the AWS S3 accelerated endpoint.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   string `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	V2Auth                string `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Versions              string `json:"versions" default:"false"`                 // Include old versions in directory listings.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                                 // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                          // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                             // Automatically rescan the source directory when this interval has passed from last successful scan
	TwoFA             string `json:"2fa" default:"false"`                                            // Two-factor authentication ('true' if the account has 2FA enabled).
	AuthToken         string `json:"authToken"`                                                      // Authentication token.
	CreateLibrary     string `json:"createLibrary" default:"false"`                                  // Should rclone create a library if it doesn't exist.
	Encoding          string `json:"encoding" default:"Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8"` // The encoding for the backend.
	Library           string `json:"library"`                                                        // Name of the library.
	LibraryKey        string `json:"libraryKey"`                                                     // Library password (for encrypted libraries only).
	Pass              string `json:"pass"`                                                           // Password.
	Url               string `json:"url"`                                                            // URL of seafile host to connect to.
	User              string `json:"user"`                                                           // User name (usually email address).
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
	SourcePath              string `validate:"required" json:"sourcePath"`          // The path of the source to scan items
	DeleteAfterExport       bool   `validate:"required" json:"deleteAfterExport"`   // Delete the source after exporting to CAR files
	RescanInterval          string `validate:"required" json:"rescanInterval"`      // Automatically rescan the source directory when this interval has passed from last successful scan
	AskPassword             string `json:"askPassword" default:"false"`             // Allow asking for SFTP password when needed.
	ChunkSize               string `json:"chunkSize" default:"32Ki"`                // Upload and download chunk size.
	Ciphers                 string `json:"ciphers"`                                 // Space separated list of ciphers to be used for session encryption, ordered by preference.
	Concurrency             string `json:"concurrency" default:"64"`                // The maximum number of outstanding requests for one file
	DisableConcurrentReads  string `json:"disableConcurrentReads" default:"false"`  // If set don't use concurrent reads.
	DisableConcurrentWrites string `json:"disableConcurrentWrites" default:"false"` // If set don't use concurrent writes.
	DisableHashcheck        string `json:"disableHashcheck" default:"false"`        // Disable the execution of SSH commands to determine if remote file hashing is available.
	Host                    string `json:"host"`                                    // SSH host to connect to.
	IdleTimeout             string `json:"idleTimeout" default:"1m0s"`              // Max time before closing idle connections.
	KeyExchange             string `json:"keyExchange"`                             // Space separated list of key exchange algorithms, ordered by preference.
	KeyFile                 string `json:"keyFile"`                                 // Path to PEM-encoded private key file.
	KeyFilePass             string `json:"keyFilePass"`                             // The passphrase to decrypt the PEM-encoded private key file.
	KeyPem                  string `json:"keyPem"`                                  // Raw PEM-encoded private key.
	KeyUseAgent             string `json:"keyUseAgent" default:"false"`             // When set forces the usage of the ssh-agent.
	KnownHostsFile          string `json:"knownHostsFile"`                          // Optional path to known_hosts file.
	Macs                    string `json:"macs"`                                    // Space separated list of MACs (message authentication code) algorithms, ordered by preference.
	Md5sumCommand           string `json:"md5sumCommand"`                           // The command used to read md5 hashes.
	Pass                    string `json:"pass"`                                    // SSH password, leave blank to use ssh-agent.
	PathOverride            string `json:"pathOverride"`                            // Override path used by SSH shell commands.
	Port                    string `json:"port" default:"22"`                       // SSH port number.
	PubkeyFile              string `json:"pubkeyFile"`                              // Optional path to public key file.
	ServerCommand           string `json:"serverCommand"`                           // Specifies the path or command to run a sftp server on the remote host.
	SetEnv                  string `json:"setEnv"`                                  // Environment variables to pass to sftp and commands
	SetModtime              string `json:"setModtime" default:"true"`               // Set the modified time on the remote if set.
	Sha1sumCommand          string `json:"sha1sumCommand"`                          // The command used to read sha1 hashes.
	ShellType               string `json:"shellType"`                               // The type of SSH shell on remote server, if any.
	SkipLinks               string `json:"skipLinks" default:"false"`               // Set to skip any symlinks and any other non regular files.
	Subsystem               string `json:"subsystem" default:"sftp"`                // Specifies the SSH2 subsystem on the remote host.
	UseFstat                string `json:"useFstat" default:"false"`                // If set use fstat instead of stat.
	UseInsecureCipher       string `json:"useInsecureCipher" default:"false"`       // Enable the use of insecure ciphers and key exchange methods.
	User                    string `json:"user" default:"shane"`                    // SSH username.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                                                                                                                   // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                                                                                                            // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                                                                                                               // Automatically rescan the source directory when this interval has passed from last successful scan
	ChunkSize         string `json:"chunkSize" default:"64Mi"`                                                                                                                         // Upload chunk size.
	Encoding          string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
	Endpoint          string `json:"endpoint"`                                                                                                                                         // Endpoint for API calls.
	RootFolderId      string `json:"rootFolderId"`                                                                                                                                     // ID of the root folder.
	UploadCutoff      string `json:"uploadCutoff" default:"128Mi"`                                                                                                                     // Cutoff for switching to multipart upload.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                                         // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                                  // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                                     // Automatically rescan the source directory when this interval has passed from last successful scan
	ApiPassword       string `json:"apiPassword"`                                                            // Sia Daemon API Password.
	ApiUrl            string `json:"apiUrl" default:"http://127.0.0.1:9980"`                                 // Sia daemon API URL, like http://sia.daemon.host:9980.
	Encoding          string `json:"encoding" default:"Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	UserAgent         string `json:"userAgent" default:"Sia-Agent"`                                          // Siad User Agent
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
	SourcePath        string `validate:"required" json:"sourcePath"`                                                                                              // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                                                                                       // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                                                                                          // Automatically rescan the source directory when this interval has passed from last successful scan
	CaseInsensitive   string `json:"caseInsensitive" default:"true"`                                                                                              // Whether the server is configured to be case-insensitive.
	Domain            string `json:"domain" default:"WORKGROUP"`                                                                                                  // Domain name for NTLM authentication.
	Encoding          string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
	HideSpecialShare  string `json:"hideSpecialShare" default:"true"`                                                                                             // Hide special shares (e.g. print$) which users aren't supposed to access.
	Host              string `json:"host"`                                                                                                                        // SMB server hostname to connect to.
	IdleTimeout       string `json:"idleTimeout" default:"1m0s"`                                                                                                  // Max time before closing idle connections.
	Pass              string `json:"pass"`                                                                                                                        // SMB password.
	Port              string `json:"port" default:"445"`                                                                                                          // SMB port number.
	Spn               string `json:"spn"`                                                                                                                         // Service principal name.
	User              string `json:"user" default:"shane"`                                                                                                        // SMB username.
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
	SourcePath        string `validate:"required" json:"sourcePath"`          // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`   // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`      // Automatically rescan the source directory when this interval has passed from last successful scan
	AccessGrant       string `json:"accessGrant"`                             // Access grant.
	ApiKey            string `json:"apiKey"`                                  // API key.
	Passphrase        string `json:"passphrase"`                              // Encryption passphrase.
	Provider          string `json:"provider" default:"existing"`             // Choose an authentication method.
	SatelliteAddress  string `json:"satelliteAddress" default:"us1.storj.io"` // Satellite address.
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

type SugarsyncRequest struct {
	SourcePath          string `validate:"required" json:"sourcePath"`               // The path of the source to scan items
	DeleteAfterExport   bool   `validate:"required" json:"deleteAfterExport"`        // Delete the source after exporting to CAR files
	RescanInterval      string `validate:"required" json:"rescanInterval"`           // Automatically rescan the source directory when this interval has passed from last successful scan
	AccessKeyId         string `json:"accessKeyId"`                                  // Sugarsync Access Key ID.
	AppId               string `json:"appId"`                                        // Sugarsync App ID.
	Authorization       string `json:"authorization"`                                // Sugarsync authorization.
	AuthorizationExpiry string `json:"authorizationExpiry"`                          // Sugarsync authorization expiry.
	DeletedId           string `json:"deletedId"`                                    // Sugarsync deleted folder id.
	Encoding            string `json:"encoding" default:"Slash,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	HardDelete          string `json:"hardDelete" default:"false"`                   // Permanently delete files if true
	PrivateAccessKey    string `json:"privateAccessKey"`                             // Sugarsync Private Access Key.
	RefreshToken        string `json:"refreshToken"`                                 // Sugarsync refresh token.
	RootId              string `json:"rootId"`                                       // Sugarsync root id.
	User                string `json:"user"`                                         // Sugarsync user.
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
	SourcePath                  string `validate:"required" json:"sourcePath"`        // The path of the source to scan items
	DeleteAfterExport           bool   `validate:"required" json:"deleteAfterExport"` // Delete the source after exporting to CAR files
	RescanInterval              string `validate:"required" json:"rescanInterval"`    // Automatically rescan the source directory when this interval has passed from last successful scan
	ApplicationCredentialId     string `json:"applicationCredentialId"`               // Application Credential ID (OS_APPLICATION_CREDENTIAL_ID).
	ApplicationCredentialName   string `json:"applicationCredentialName"`             // Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME).
	ApplicationCredentialSecret string `json:"applicationCredentialSecret"`           // Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET).
	Auth                        string `json:"auth"`                                  // Authentication URL for server (OS_AUTH_URL).
	AuthToken                   string `json:"authToken"`                             // Auth Token from alternate authentication - optional (OS_AUTH_TOKEN).
	AuthVersion                 string `json:"authVersion" default:"0"`               // AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).
	ChunkSize                   string `json:"chunkSize" default:"5Gi"`               // Above this size files will be chunked into a _segments container.
	Domain                      string `json:"domain"`                                // User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME)
	Encoding                    string `json:"encoding" default:"Slash,InvalidUtf8"`  // The encoding for the backend.
	EndpointType                string `json:"endpointType" default:"public"`         // Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).
	EnvAuth                     string `json:"envAuth" default:"false"`               // Get swift credentials from environment variables in standard OpenStack form.
	Key                         string `json:"key"`                                   // API key or password (OS_PASSWORD).
	LeavePartsOnError           string `json:"leavePartsOnError" default:"false"`     // If true avoid calling abort upload on a failure.
	NoChunk                     string `json:"noChunk" default:"false"`               // Don't chunk files during streaming upload.
	NoLargeObjects              string `json:"noLargeObjects" default:"false"`        // Disable support for static and dynamic large objects
	Region                      string `json:"region"`                                // Region name - optional (OS_REGION_NAME).
	StoragePolicy               string `json:"storagePolicy"`                         // The storage policy to use when creating a new container.
	StorageUrl                  string `json:"storageUrl"`                            // Storage URL - optional (OS_STORAGE_URL).
	Tenant                      string `json:"tenant"`                                // Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME).
	TenantDomain                string `json:"tenantDomain"`                          // Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME).
	TenantId                    string `json:"tenantId"`                              // Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID).
	User                        string `json:"user"`                                  // User name to log in (OS_USERNAME).
	UserId                      string `json:"userId"`                                // User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID).
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
	SourcePath        string `validate:"required" json:"sourcePath"`                                                        // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`                                                 // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`                                                    // Automatically rescan the source directory when this interval has passed from last successful scan
	AccessToken       string `json:"accessToken"`                                                                           // Your access token.
	Encoding          string `json:"encoding" default:"Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot"` // The encoding for the backend.
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
	SourcePath         string `validate:"required" json:"sourcePath"`        // The path of the source to scan items
	DeleteAfterExport  bool   `validate:"required" json:"deleteAfterExport"` // Delete the source after exporting to CAR files
	RescanInterval     string `validate:"required" json:"rescanInterval"`    // Automatically rescan the source directory when this interval has passed from last successful scan
	BearerToken        string `json:"bearerToken"`                           // Bearer token instead of user/pass (e.g. a Macaroon).
	BearerTokenCommand string `json:"bearerTokenCommand"`                    // Command to run to get a bearer token.
	Encoding           string `json:"encoding"`                              // The encoding for the backend.
	Headers            string `json:"headers"`                               // Set HTTP headers for all transactions.
	Pass               string `json:"pass"`                                  // Password.
	Url                string `json:"url"`                                   // URL of http host to connect to.
	User               string `json:"user"`                                  // User name.
	Vendor             string `json:"vendor"`                                // Name of the WebDAV site/service/software you are using.
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
	SourcePath        string `validate:"required" json:"sourcePath"`                   // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`            // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`               // Automatically rescan the source directory when this interval has passed from last successful scan
	AuthUrl           string `json:"authUrl"`                                          // Auth server URL.
	ClientId          string `json:"clientId"`                                         // OAuth Client Id.
	ClientSecret      string `json:"clientSecret"`                                     // OAuth Client Secret.
	Encoding          string `json:"encoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	HardDelete        string `json:"hardDelete" default:"false"`                       // Delete files permanently rather than putting them into the trash.
	Token             string `json:"token"`                                            // OAuth Access Token as a JSON blob.
	TokenUrl          string `json:"tokenUrl"`                                         // Token server url.
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
	SourcePath        string `validate:"required" json:"sourcePath"`         // The path of the source to scan items
	DeleteAfterExport bool   `validate:"required" json:"deleteAfterExport"`  // Delete the source after exporting to CAR files
	RescanInterval    string `validate:"required" json:"rescanInterval"`     // Automatically rescan the source directory when this interval has passed from last successful scan
	AuthUrl           string `json:"authUrl"`                                // Auth server URL.
	ClientId          string `json:"clientId"`                               // OAuth Client Id.
	ClientSecret      string `json:"clientSecret"`                           // OAuth Client Secret.
	Encoding          string `json:"encoding" default:"Del,Ctl,InvalidUtf8"` // The encoding for the backend.
	Region            string `json:"region"`                                 // Zoho region to connect to.
	Token             string `json:"token"`                                  // OAuth Access Token as a JSON blob.
	TokenUrl          string `json:"tokenUrl"`                               // Token server url.
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
	AcdAuthUrl                          string `json:"acdAuthUrl"`                                                                                                                                                  // Auth server URL.
	AcdCheckpoint                       string `json:"acdCheckpoint"`                                                                                                                                               // Checkpoint for internal polling (debug).
	AcdClientId                         string `json:"acdClientId"`                                                                                                                                                 // OAuth Client Id.
	AcdClientSecret                     string `json:"acdClientSecret"`                                                                                                                                             // OAuth Client Secret.
	AcdEncoding                         string `json:"acdEncoding" default:"Slash,InvalidUtf8,Dot"`                                                                                                                 // The encoding for the backend.
	AcdTemplinkThreshold                string `json:"acdTemplinkThreshold" default:"9Gi"`                                                                                                                          // Files >= this size will be downloaded via their tempLink.
	AcdToken                            string `json:"acdToken"`                                                                                                                                                    // OAuth Access Token as a JSON blob.
	AcdTokenUrl                         string `json:"acdTokenUrl"`                                                                                                                                                 // Token server url.
	AcdUploadWaitPerGb                  string `json:"acdUploadWaitPerGb" default:"3m0s"`                                                                                                                           // Additional time per GiB to wait after a failed complete upload to see if it appears.
	AzureblobAccessTier                 string `json:"azureblobAccessTier"`                                                                                                                                         // Access tier of blob: hot, cool or archive.
	AzureblobAccount                    string `json:"azureblobAccount"`                                                                                                                                            // Azure Storage Account Name.
	AzureblobArchiveTierDelete          string `json:"azureblobArchiveTierDelete" default:"false"`                                                                                                                  // Delete archive tier blobs before overwriting.
	AzureblobChunkSize                  string `json:"azureblobChunkSize" default:"4Mi"`                                                                                                                            // Upload chunk size.
	AzureblobClientCertificatePassword  string `json:"azureblobClientCertificatePassword"`                                                                                                                          // Password for the certificate file (optional).
	AzureblobClientCertificatePath      string `json:"azureblobClientCertificatePath"`                                                                                                                              // Path to a PEM or PKCS12 certificate file including the private key.
	AzureblobClientId                   string `json:"azureblobClientId"`                                                                                                                                           // The ID of the client in use.
	AzureblobClientSecret               string `json:"azureblobClientSecret"`                                                                                                                                       // One of the service principal's client secrets
	AzureblobClientSendCertificateChain string `json:"azureblobClientSendCertificateChain" default:"false"`                                                                                                         // Send the certificate chain when using certificate auth.
	AzureblobDisableChecksum            string `json:"azureblobDisableChecksum" default:"false"`                                                                                                                    // Don't store MD5 checksum with object metadata.
	AzureblobEncoding                   string `json:"azureblobEncoding" default:"Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"`                                                                                 // The encoding for the backend.
	AzureblobEndpoint                   string `json:"azureblobEndpoint"`                                                                                                                                           // Endpoint for the service.
	AzureblobEnvAuth                    string `json:"azureblobEnvAuth" default:"false"`                                                                                                                            // Read credentials from runtime (environment variables, CLI or MSI).
	AzureblobKey                        string `json:"azureblobKey"`                                                                                                                                                // Storage Account Shared Key.
	AzureblobListChunk                  string `json:"azureblobListChunk" default:"5000"`                                                                                                                           // Size of blob list.
	AzureblobMemoryPoolFlushTime        string `json:"azureblobMemoryPoolFlushTime" default:"1m0s"`                                                                                                                 // How often internal memory buffer pools will be flushed.
	AzureblobMemoryPoolUseMmap          string `json:"azureblobMemoryPoolUseMmap" default:"false"`                                                                                                                  // Whether to use mmap buffers in internal memory pool.
	AzureblobMsiClientId                string `json:"azureblobMsiClientId"`                                                                                                                                        // Object ID of the user-assigned MSI to use, if any.
	AzureblobMsiMiResId                 string `json:"azureblobMsiMiResId"`                                                                                                                                         // Azure resource ID of the user-assigned MSI to use, if any.
	AzureblobMsiObjectId                string `json:"azureblobMsiObjectId"`                                                                                                                                        // Object ID of the user-assigned MSI to use, if any.
	AzureblobNoCheckContainer           string `json:"azureblobNoCheckContainer" default:"false"`                                                                                                                   // If set, don't attempt to check the container exists or create it.
	AzureblobNoHeadObject               string `json:"azureblobNoHeadObject" default:"false"`                                                                                                                       // If set, do not do HEAD before GET when getting objects.
	AzureblobPassword                   string `json:"azureblobPassword"`                                                                                                                                           // The user's password
	AzureblobPublicAccess               string `json:"azureblobPublicAccess"`                                                                                                                                       // Public access level of a container: blob or container.
	AzureblobSasUrl                     string `json:"azureblobSasUrl"`                                                                                                                                             // SAS URL for container level access only.
	AzureblobServicePrincipalFile       string `json:"azureblobServicePrincipalFile"`                                                                                                                               // Path to file containing credentials for use with a service principal.
	AzureblobTenant                     string `json:"azureblobTenant"`                                                                                                                                             // ID of the service principal's tenant. Also called its directory ID.
	AzureblobUploadConcurrency          string `json:"azureblobUploadConcurrency" default:"16"`                                                                                                                     // Concurrency for multipart uploads.
	AzureblobUploadCutoff               string `json:"azureblobUploadCutoff"`                                                                                                                                       // Cutoff for switching to chunked upload (<= 256 MiB) (deprecated).
	AzureblobUseEmulator                string `json:"azureblobUseEmulator" default:"false"`                                                                                                                        // Uses local storage emulator if provided as 'true'.
	AzureblobUseMsi                     string `json:"azureblobUseMsi" default:"false"`                                                                                                                             // Use a managed service identity to authenticate (only works in Azure).
	AzureblobUsername                   string `json:"azureblobUsername"`                                                                                                                                           // User name (usually an email address)
	B2Account                           string `json:"b2Account"`                                                                                                                                                   // Account ID or Application Key ID.
	B2ChunkSize                         string `json:"b2ChunkSize" default:"96Mi"`                                                                                                                                  // Upload chunk size.
	B2CopyCutoff                        string `json:"b2CopyCutoff" default:"4Gi"`                                                                                                                                  // Cutoff for switching to multipart copy.
	B2DisableChecksum                   string `json:"b2DisableChecksum" default:"false"`                                                                                                                           // Disable checksums for large (> upload cutoff) files.
	B2DownloadAuthDuration              string `json:"b2DownloadAuthDuration" default:"1w"`                                                                                                                         // Time before the authorization token will expire in s or suffix ms|s|m|h|d.
	B2DownloadUrl                       string `json:"b2DownloadUrl"`                                                                                                                                               // Custom endpoint for downloads.
	B2Encoding                          string `json:"b2Encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"`                                                                                                // The encoding for the backend.
	B2Endpoint                          string `json:"b2Endpoint"`                                                                                                                                                  // Endpoint for the service.
	B2HardDelete                        string `json:"b2HardDelete" default:"false"`                                                                                                                                // Permanently delete files on remote removal, otherwise hide files.
	B2Key                               string `json:"b2Key"`                                                                                                                                                       // Application Key.
	B2MemoryPoolFlushTime               string `json:"b2MemoryPoolFlushTime" default:"1m0s"`                                                                                                                        // How often internal memory buffer pools will be flushed.
	B2MemoryPoolUseMmap                 string `json:"b2MemoryPoolUseMmap" default:"false"`                                                                                                                         // Whether to use mmap buffers in internal memory pool.
	B2TestMode                          string `json:"b2TestMode"`                                                                                                                                                  // A flag string for X-Bz-Test-Mode header for debugging.
	B2UploadCutoff                      string `json:"b2UploadCutoff" default:"200Mi"`                                                                                                                              // Cutoff for switching to chunked upload.
	B2VersionAt                         string `json:"b2VersionAt" default:"off"`                                                                                                                                   // Show file versions as they were at the specified time.
	B2Versions                          string `json:"b2Versions" default:"false"`                                                                                                                                  // Include old versions in directory listings.
	BoxAccessToken                      string `json:"boxAccessToken"`                                                                                                                                              // Box App Primary Access Token
	BoxAuthUrl                          string `json:"boxAuthUrl"`                                                                                                                                                  // Auth server URL.
	BoxBoxConfigFile                    string `json:"boxBoxConfigFile"`                                                                                                                                            // Box App config.json location
	BoxBoxSubType                       string `json:"boxBoxSubType" default:"user"`                                                                                                                                //
	BoxClientId                         string `json:"boxClientId"`                                                                                                                                                 // OAuth Client Id.
	BoxClientSecret                     string `json:"boxClientSecret"`                                                                                                                                             // OAuth Client Secret.
	BoxCommitRetries                    string `json:"boxCommitRetries" default:"100"`                                                                                                                              // Max number of times to try committing a multipart file.
	BoxEncoding                         string `json:"boxEncoding" default:"Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"`                                                                                    // The encoding for the backend.
	BoxListChunk                        string `json:"boxListChunk" default:"1000"`                                                                                                                                 // Size of listing chunk 1-1000.
	BoxOwnedBy                          string `json:"boxOwnedBy"`                                                                                                                                                  // Only show items owned by the login (email address) passed in.
	BoxRootFolderId                     string `json:"boxRootFolderId" default:"0"`                                                                                                                                 // Fill in for rclone to use a non root folder as its starting point.
	BoxToken                            string `json:"boxToken"`                                                                                                                                                    // OAuth Access Token as a JSON blob.
	BoxTokenUrl                         string `json:"boxTokenUrl"`                                                                                                                                                 // Token server url.
	BoxUploadCutoff                     string `json:"boxUploadCutoff" default:"50Mi"`                                                                                                                              // Cutoff for switching to multipart upload (>= 50 MiB).
	DriveAcknowledgeAbuse               string `json:"driveAcknowledgeAbuse" default:"false"`                                                                                                                       // Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
	DriveAllowImportNameChange          string `json:"driveAllowImportNameChange" default:"false"`                                                                                                                  // Allow the filetype to change when uploading Google docs.
	DriveAlternateExport                string `json:"driveAlternateExport" default:"false"`                                                                                                                        // Deprecated: No longer needed.
	DriveAuthOwnerOnly                  string `json:"driveAuthOwnerOnly" default:"false"`                                                                                                                          // Only consider files owned by the authenticated user.
	DriveAuthUrl                        string `json:"driveAuthUrl"`                                                                                                                                                // Auth server URL.
	DriveChunkSize                      string `json:"driveChunkSize" default:"8Mi"`                                                                                                                                // Upload chunk size.
	DriveClientId                       string `json:"driveClientId"`                                                                                                                                               // Google Application Client Id
	DriveClientSecret                   string `json:"driveClientSecret"`                                                                                                                                           // OAuth Client Secret.
	DriveCopyShortcutContent            string `json:"driveCopyShortcutContent" default:"false"`                                                                                                                    // Server side copy contents of shortcuts instead of the shortcut.
	DriveDisableHttp2                   string `json:"driveDisableHttp2" default:"true"`                                                                                                                            // Disable drive using http2.
	DriveEncoding                       string `json:"driveEncoding" default:"InvalidUtf8"`                                                                                                                         // The encoding for the backend.
	DriveExportFormats                  string `json:"driveExportFormats" default:"docx,xlsx,pptx,svg"`                                                                                                             // Comma separated list of preferred formats for downloading Google docs.
	DriveFormats                        string `json:"driveFormats"`                                                                                                                                                // Deprecated: See export_formats.
	DriveImpersonate                    string `json:"driveImpersonate"`                                                                                                                                            // Impersonate this user when using a service account.
	DriveImportFormats                  string `json:"driveImportFormats"`                                                                                                                                          // Comma separated list of preferred formats for uploading Google docs.
	DriveKeepRevisionForever            string `json:"driveKeepRevisionForever" default:"false"`                                                                                                                    // Keep new head revision of each file forever.
	DriveListChunk                      string `json:"driveListChunk" default:"1000"`                                                                                                                               // Size of listing chunk 100-1000, 0 to disable.
	DrivePacerBurst                     string `json:"drivePacerBurst" default:"100"`                                                                                                                               // Number of API calls to allow without sleeping.
	DrivePacerMinSleep                  string `json:"drivePacerMinSleep" default:"100ms"`                                                                                                                          // Minimum time to sleep between API calls.
	DriveResourceKey                    string `json:"driveResourceKey"`                                                                                                                                            // Resource key for accessing a link-shared file.
	DriveRootFolderId                   string `json:"driveRootFolderId"`                                                                                                                                           // ID of the root folder.
	DriveScope                          string `json:"driveScope"`                                                                                                                                                  // Scope that rclone should use when requesting access from drive.
	DriveServerSideAcrossConfigs        string `json:"driveServerSideAcrossConfigs" default:"false"`                                                                                                                // Allow server-side operations (e.g. copy) to work across different drive configs.
	DriveServiceAccountCredentials      string `json:"driveServiceAccountCredentials"`                                                                                                                              // Service Account Credentials JSON blob.
	DriveServiceAccountFile             string `json:"driveServiceAccountFile"`                                                                                                                                     // Service Account Credentials JSON file path.
	DriveSharedWithMe                   string `json:"driveSharedWithMe" default:"false"`                                                                                                                           // Only show files that are shared with me.
	DriveSizeAsQuota                    string `json:"driveSizeAsQuota" default:"false"`                                                                                                                            // Show sizes as storage quota usage, not actual size.
	DriveSkipChecksumGphotos            string `json:"driveSkipChecksumGphotos" default:"false"`                                                                                                                    // Skip MD5 checksum on Google photos and videos only.
	DriveSkipDanglingShortcuts          string `json:"driveSkipDanglingShortcuts" default:"false"`                                                                                                                  // If set skip dangling shortcut files.
	DriveSkipGdocs                      string `json:"driveSkipGdocs" default:"false"`                                                                                                                              // Skip google documents in all listings.
	DriveSkipShortcuts                  string `json:"driveSkipShortcuts" default:"false"`                                                                                                                          // If set skip shortcut files.
	DriveStarredOnly                    string `json:"driveStarredOnly" default:"false"`                                                                                                                            // Only show files that are starred.
	DriveStopOnDownloadLimit            string `json:"driveStopOnDownloadLimit" default:"false"`                                                                                                                    // Make download limit errors be fatal.
	DriveStopOnUploadLimit              string `json:"driveStopOnUploadLimit" default:"false"`                                                                                                                      // Make upload limit errors be fatal.
	DriveTeamDrive                      string `json:"driveTeamDrive"`                                                                                                                                              // ID of the Shared Drive (Team Drive).
	DriveToken                          string `json:"driveToken"`                                                                                                                                                  // OAuth Access Token as a JSON blob.
	DriveTokenUrl                       string `json:"driveTokenUrl"`                                                                                                                                               // Token server url.
	DriveTrashedOnly                    string `json:"driveTrashedOnly" default:"false"`                                                                                                                            // Only show files that are in the trash.
	DriveUploadCutoff                   string `json:"driveUploadCutoff" default:"8Mi"`                                                                                                                             // Cutoff for switching to chunked upload.
	DriveUseCreatedDate                 string `json:"driveUseCreatedDate" default:"false"`                                                                                                                         // Use file created date instead of modified date.
	DriveUseSharedDate                  string `json:"driveUseSharedDate" default:"false"`                                                                                                                          // Use date file was shared instead of modified date.
	DriveUseTrash                       string `json:"driveUseTrash" default:"true"`                                                                                                                                // Send files to the trash instead of deleting permanently.
	DriveV2DownloadMinSize              string `json:"driveV2DownloadMinSize" default:"off"`                                                                                                                        // If Object's are greater, use drive v2 API to download.
	DropboxAuthUrl                      string `json:"dropboxAuthUrl"`                                                                                                                                              // Auth server URL.
	DropboxBatchCommitTimeout           string `json:"dropboxBatchCommitTimeout" default:"10m0s"`                                                                                                                   // Max time to wait for a batch to finish committing
	DropboxBatchMode                    string `json:"dropboxBatchMode" default:"sync"`                                                                                                                             // Upload file batching sync|async|off.
	DropboxBatchSize                    string `json:"dropboxBatchSize" default:"0"`                                                                                                                                // Max number of files in upload batch.
	DropboxBatchTimeout                 string `json:"dropboxBatchTimeout" default:"0s"`                                                                                                                            // Max time to allow an idle upload batch before uploading.
	DropboxChunkSize                    string `json:"dropboxChunkSize" default:"48Mi"`                                                                                                                             // Upload chunk size (< 150Mi).
	DropboxClientId                     string `json:"dropboxClientId"`                                                                                                                                             // OAuth Client Id.
	DropboxClientSecret                 string `json:"dropboxClientSecret"`                                                                                                                                         // OAuth Client Secret.
	DropboxEncoding                     string `json:"dropboxEncoding" default:"Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"`                                                                                    // The encoding for the backend.
	DropboxImpersonate                  string `json:"dropboxImpersonate"`                                                                                                                                          // Impersonate this user when using a business account.
	DropboxSharedFiles                  string `json:"dropboxSharedFiles" default:"false"`                                                                                                                          // Instructs rclone to work on individual shared files.
	DropboxSharedFolders                string `json:"dropboxSharedFolders" default:"false"`                                                                                                                        // Instructs rclone to work on shared folders.
	DropboxToken                        string `json:"dropboxToken"`                                                                                                                                                // OAuth Access Token as a JSON blob.
	DropboxTokenUrl                     string `json:"dropboxTokenUrl"`                                                                                                                                             // Token server url.
	FichierApiKey                       string `json:"fichierApiKey"`                                                                                                                                               // Your API Key, get it from https://1fichier.com/console/params.pl.
	FichierEncoding                     string `json:"fichierEncoding" default:"Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"`                        // The encoding for the backend.
	FichierFilePassword                 string `json:"fichierFilePassword"`                                                                                                                                         // If you want to download a shared file that is password protected, add this parameter.
	FichierFolderPassword               string `json:"fichierFolderPassword"`                                                                                                                                       // If you want to list the files in a shared folder that is password protected, add this parameter.
	FichierSharedFolder                 string `json:"fichierSharedFolder"`                                                                                                                                         // If you want to download a shared folder, add this parameter.
	FilefabricEncoding                  string `json:"filefabricEncoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"`                                                                                                  // The encoding for the backend.
	FilefabricPermanentToken            string `json:"filefabricPermanentToken"`                                                                                                                                    // Permanent Authentication Token.
	FilefabricRootFolderId              string `json:"filefabricRootFolderId"`                                                                                                                                      // ID of the root folder.
	FilefabricToken                     string `json:"filefabricToken"`                                                                                                                                             // Session Token.
	FilefabricTokenExpiry               string `json:"filefabricTokenExpiry"`                                                                                                                                       // Token expiry time.
	FilefabricUrl                       string `json:"filefabricUrl"`                                                                                                                                               // URL of the Enterprise File Fabric to connect to.
	FilefabricVersion                   string `json:"filefabricVersion"`                                                                                                                                           // Version read from the file fabric.
	FtpAskPassword                      string `json:"ftpAskPassword" default:"false"`                                                                                                                              // Allow asking for FTP password when needed.
	FtpCloseTimeout                     string `json:"ftpCloseTimeout" default:"1m0s"`                                                                                                                              // Maximum time to wait for a response to close.
	FtpConcurrency                      string `json:"ftpConcurrency" default:"0"`                                                                                                                                  // Maximum number of FTP simultaneous connections, 0 for unlimited.
	FtpDisableEpsv                      string `json:"ftpDisableEpsv" default:"false"`                                                                                                                              // Disable using EPSV even if server advertises support.
	FtpDisableMlsd                      string `json:"ftpDisableMlsd" default:"false"`                                                                                                                              // Disable using MLSD even if server advertises support.
	FtpDisableTls13                     string `json:"ftpDisableTls13" default:"false"`                                                                                                                             // Disable TLS 1.3 (workaround for FTP servers with buggy TLS)
	FtpDisableUtf8                      string `json:"ftpDisableUtf8" default:"false"`                                                                                                                              // Disable using UTF-8 even if server advertises support.
	FtpEncoding                         string `json:"ftpEncoding" default:"Slash,Del,Ctl,RightSpace,Dot"`                                                                                                          // The encoding for the backend.
	FtpExplicitTls                      string `json:"ftpExplicitTls" default:"false"`                                                                                                                              // Use Explicit FTPS (FTP over TLS).
	FtpForceListHidden                  string `json:"ftpForceListHidden" default:"false"`                                                                                                                          // Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.
	FtpHost                             string `json:"ftpHost"`                                                                                                                                                     // FTP host to connect to.
	FtpIdleTimeout                      string `json:"ftpIdleTimeout" default:"1m0s"`                                                                                                                               // Max time before closing idle connections.
	FtpNoCheckCertificate               string `json:"ftpNoCheckCertificate" default:"false"`                                                                                                                       // Do not verify the TLS certificate of the server.
	FtpPass                             string `json:"ftpPass"`                                                                                                                                                     // FTP password.
	FtpPort                             string `json:"ftpPort" default:"21"`                                                                                                                                        // FTP port number.
	FtpShutTimeout                      string `json:"ftpShutTimeout" default:"1m0s"`                                                                                                                               // Maximum time to wait for data connection closing status.
	FtpTls                              string `json:"ftpTls" default:"false"`                                                                                                                                      // Use Implicit FTPS (FTP over TLS).
	FtpTlsCacheSize                     string `json:"ftpTlsCacheSize" default:"32"`                                                                                                                                // Size of TLS session cache for all control and data connections.
	FtpUser                             string `json:"ftpUser" default:"shane"`                                                                                                                                     // FTP username.
	FtpWritingMdtm                      string `json:"ftpWritingMdtm" default:"false"`                                                                                                                              // Use MDTM to set modification time (VsFtpd quirk)
	GcsAnonymous                        string `json:"gcsAnonymous" default:"false"`                                                                                                                                // Access public buckets and objects without credentials.
	GcsAuthUrl                          string `json:"gcsAuthUrl"`                                                                                                                                                  // Auth server URL.
	GcsBucketAcl                        string `json:"gcsBucketAcl"`                                                                                                                                                // Access Control List for new buckets.
	GcsBucketPolicyOnly                 string `json:"gcsBucketPolicyOnly" default:"false"`                                                                                                                         // Access checks should use bucket-level IAM policies.
	GcsClientId                         string `json:"gcsClientId"`                                                                                                                                                 // OAuth Client Id.
	GcsClientSecret                     string `json:"gcsClientSecret"`                                                                                                                                             // OAuth Client Secret.
	GcsDecompress                       string `json:"gcsDecompress" default:"false"`                                                                                                                               // If set this will decompress gzip encoded objects.
	GcsEncoding                         string `json:"gcsEncoding" default:"Slash,CrLf,InvalidUtf8,Dot"`                                                                                                            // The encoding for the backend.
	GcsEndpoint                         string `json:"gcsEndpoint"`                                                                                                                                                 // Endpoint for the service.
	GcsEnvAuth                          string `json:"gcsEnvAuth" default:"false"`                                                                                                                                  // Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
	GcsLocation                         string `json:"gcsLocation"`                                                                                                                                                 // Location for the newly created buckets.
	GcsNoCheckBucket                    string `json:"gcsNoCheckBucket" default:"false"`                                                                                                                            // If set, don't attempt to check the bucket exists or create it.
	GcsObjectAcl                        string `json:"gcsObjectAcl"`                                                                                                                                                // Access Control List for new objects.
	GcsProjectNumber                    string `json:"gcsProjectNumber"`                                                                                                                                            // Project number.
	GcsServiceAccountCredentials        string `json:"gcsServiceAccountCredentials"`                                                                                                                                // Service Account Credentials JSON blob.
	GcsServiceAccountFile               string `json:"gcsServiceAccountFile"`                                                                                                                                       // Service Account Credentials JSON file path.
	GcsStorageClass                     string `json:"gcsStorageClass"`                                                                                                                                             // The storage class to use when storing objects in Google Cloud Storage.
	GcsToken                            string `json:"gcsToken"`                                                                                                                                                    // OAuth Access Token as a JSON blob.
	GcsTokenUrl                         string `json:"gcsTokenUrl"`                                                                                                                                                 // Token server url.
	GphotosAuthUrl                      string `json:"gphotosAuthUrl"`                                                                                                                                              // Auth server URL.
	GphotosClientId                     string `json:"gphotosClientId"`                                                                                                                                             // OAuth Client Id.
	GphotosClientSecret                 string `json:"gphotosClientSecret"`                                                                                                                                         // OAuth Client Secret.
	GphotosEncoding                     string `json:"gphotosEncoding" default:"Slash,CrLf,InvalidUtf8,Dot"`                                                                                                        // The encoding for the backend.
	GphotosIncludeArchived              string `json:"gphotosIncludeArchived" default:"false"`                                                                                                                      // Also view and download archived media.
	GphotosReadOnly                     string `json:"gphotosReadOnly" default:"false"`                                                                                                                             // Set to make the Google Photos backend read only.
	GphotosReadSize                     string `json:"gphotosReadSize" default:"false"`                                                                                                                             // Set to read the size of media items.
	GphotosStartYear                    string `json:"gphotosStartYear" default:"2000"`                                                                                                                             // Year limits the photos to be downloaded to those which are uploaded after the given year.
	GphotosToken                        string `json:"gphotosToken"`                                                                                                                                                // OAuth Access Token as a JSON blob.
	GphotosTokenUrl                     string `json:"gphotosTokenUrl"`                                                                                                                                             // Token server url.
	HdfsDataTransferProtection          string `json:"hdfsDataTransferProtection"`                                                                                                                                  // Kerberos data transfer protection: authentication|integrity|privacy.
	HdfsEncoding                        string `json:"hdfsEncoding" default:"Slash,Colon,Del,Ctl,InvalidUtf8,Dot"`                                                                                                  // The encoding for the backend.
	HdfsNamenode                        string `json:"hdfsNamenode"`                                                                                                                                                // Hadoop name node and port.
	HdfsServicePrincipalName            string `json:"hdfsServicePrincipalName"`                                                                                                                                    // Kerberos service principal name for the namenode.
	HdfsUsername                        string `json:"hdfsUsername"`                                                                                                                                                // Hadoop user name.
	HidriveAuthUrl                      string `json:"hidriveAuthUrl"`                                                                                                                                              // Auth server URL.
	HidriveChunkSize                    string `json:"hidriveChunkSize" default:"48Mi"`                                                                                                                             // Chunksize for chunked uploads.
	HidriveClientId                     string `json:"hidriveClientId"`                                                                                                                                             // OAuth Client Id.
	HidriveClientSecret                 string `json:"hidriveClientSecret"`                                                                                                                                         // OAuth Client Secret.
	HidriveDisableFetchingMemberCount   string `json:"hidriveDisableFetchingMemberCount" default:"false"`                                                                                                           // Do not fetch number of objects in directories unless it is absolutely necessary.
	HidriveEncoding                     string `json:"hidriveEncoding" default:"Slash,Dot"`                                                                                                                         // The encoding for the backend.
	HidriveEndpoint                     string `json:"hidriveEndpoint" default:"https://api.hidrive.strato.com/2.1"`                                                                                                // Endpoint for the service.
	HidriveRootPrefix                   string `json:"hidriveRootPrefix" default:"/"`                                                                                                                               // The root/parent folder for all paths.
	HidriveScopeAccess                  string `json:"hidriveScopeAccess" default:"rw"`                                                                                                                             // Access permissions that rclone should use when requesting access from HiDrive.
	HidriveScopeRole                    string `json:"hidriveScopeRole" default:"user"`                                                                                                                             // User-level that rclone should use when requesting access from HiDrive.
	HidriveToken                        string `json:"hidriveToken"`                                                                                                                                                // OAuth Access Token as a JSON blob.
	HidriveTokenUrl                     string `json:"hidriveTokenUrl"`                                                                                                                                             // Token server url.
	HidriveUploadConcurrency            string `json:"hidriveUploadConcurrency" default:"4"`                                                                                                                        // Concurrency for chunked uploads.
	HidriveUploadCutoff                 string `json:"hidriveUploadCutoff" default:"96Mi"`                                                                                                                          // Cutoff/Threshold for chunked uploads.
	HttpHeaders                         string `json:"httpHeaders"`                                                                                                                                                 // Set HTTP headers for all transactions.
	HttpNoHead                          string `json:"httpNoHead" default:"false"`                                                                                                                                  // Don't use HEAD requests.
	HttpNoSlash                         string `json:"httpNoSlash" default:"false"`                                                                                                                                 // Set this if the site doesn't end directories with /.
	HttpUrl                             string `json:"httpUrl"`                                                                                                                                                     // URL of HTTP host to connect to.
	InternetarchiveAccessKeyId          string `json:"internetarchiveAccessKeyId"`                                                                                                                                  // IAS3 Access Key.
	InternetarchiveDisableChecksum      string `json:"internetarchiveDisableChecksum" default:"true"`                                                                                                               // Don't ask the server to test against MD5 checksum calculated by rclone.
	InternetarchiveEncoding             string `json:"internetarchiveEncoding" default:"Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot"`                                                                                   // The encoding for the backend.
	InternetarchiveEndpoint             string `json:"internetarchiveEndpoint" default:"https://s3.us.archive.org"`                                                                                                 // IAS3 Endpoint.
	InternetarchiveFrontEndpoint        string `json:"internetarchiveFrontEndpoint" default:"https://archive.org"`                                                                                                  // Host of InternetArchive Frontend.
	InternetarchiveSecretAccessKey      string `json:"internetarchiveSecretAccessKey"`                                                                                                                              // IAS3 Secret Key (password).
	InternetarchiveWaitArchive          string `json:"internetarchiveWaitArchive" default:"0s"`                                                                                                                     // Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
	JottacloudEncoding                  string `json:"jottacloudEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot"`                                                    // The encoding for the backend.
	JottacloudHardDelete                string `json:"jottacloudHardDelete" default:"false"`                                                                                                                        // Delete files permanently rather than putting them into the trash.
	JottacloudMd5MemoryLimit            string `json:"jottacloudMd5MemoryLimit" default:"10Mi"`                                                                                                                     // Files bigger than this will be cached on disk to calculate the MD5 if required.
	JottacloudNoVersions                string `json:"jottacloudNoVersions" default:"false"`                                                                                                                        // Avoid server side versioning by deleting files and recreating files instead of overwriting them.
	JottacloudTrashedOnly               string `json:"jottacloudTrashedOnly" default:"false"`                                                                                                                       // Only show files that are in the trash.
	JottacloudUploadResumeLimit         string `json:"jottacloudUploadResumeLimit" default:"10Mi"`                                                                                                                  // Files bigger than this can be resumed if the upload fail's.
	KoofrEncoding                       string `json:"koofrEncoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"`                                                                                             // The encoding for the backend.
	KoofrEndpoint                       string `json:"koofrEndpoint"`                                                                                                                                               // The Koofr API endpoint to use.
	KoofrMountid                        string `json:"koofrMountid"`                                                                                                                                                // Mount ID of the mount to use.
	KoofrPassword                       string `json:"koofrPassword"`                                                                                                                                               // Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password).
	KoofrProvider                       string `json:"koofrProvider"`                                                                                                                                               // Choose your storage provider.
	KoofrSetmtime                       string `json:"koofrSetmtime" default:"true"`                                                                                                                                // Does the backend support setting modification time.
	KoofrUser                           string `json:"koofrUser"`                                                                                                                                                   // Your user name.
	LocalCaseInsensitive                string `json:"localCaseInsensitive" default:"false"`                                                                                                                        // Force the filesystem to report itself as case insensitive.
	LocalCaseSensitive                  string `json:"localCaseSensitive" default:"false"`                                                                                                                          // Force the filesystem to report itself as case sensitive.
	LocalCopyLinks                      string `json:"localCopyLinks" default:"false"`                                                                                                                              // Follow symlinks and copy the pointed to item.
	LocalEncoding                       string `json:"localEncoding" default:"Slash,Dot"`                                                                                                                           // The encoding for the backend.
	LocalLinks                          string `json:"localLinks" default:"false"`                                                                                                                                  // Translate symlinks to/from regular files with a '.rclonelink' extension.
	LocalNoCheckUpdated                 string `json:"localNoCheckUpdated" default:"false"`                                                                                                                         // Don't check to see if the files change during upload.
	LocalNoPreallocate                  string `json:"localNoPreallocate" default:"false"`                                                                                                                          // Disable preallocation of disk space for transferred files.
	LocalNoSetModtime                   string `json:"localNoSetModtime" default:"false"`                                                                                                                           // Disable setting modtime.
	LocalNoSparse                       string `json:"localNoSparse" default:"false"`                                                                                                                               // Disable sparse files for multi-thread downloads.
	LocalNounc                          string `json:"localNounc" default:"false"`                                                                                                                                  // Disable UNC (long path names) conversion on Windows.
	LocalOneFileSystem                  string `json:"localOneFileSystem" default:"false"`                                                                                                                          // Don't cross filesystem boundaries (unix/macOS only).
	LocalSkipLinks                      string `json:"localSkipLinks" default:"false"`                                                                                                                              // Don't warn about skipped symlinks.
	LocalUnicodeNormalization           string `json:"localUnicodeNormalization" default:"false"`                                                                                                                   // Apply unicode NFC normalization to paths and filenames.
	LocalZeroSizeLinks                  string `json:"localZeroSizeLinks" default:"false"`                                                                                                                          // Assume the Stat size of links is zero (and read them instead) (deprecated).
	MailruCheckHash                     string `json:"mailruCheckHash" default:"true"`                                                                                                                              // What should copy do if file checksum is mismatched or invalid.
	MailruEncoding                      string `json:"mailruEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot"`                                              // The encoding for the backend.
	MailruPass                          string `json:"mailruPass"`                                                                                                                                                  // Password.
	MailruQuirks                        string `json:"mailruQuirks"`                                                                                                                                                // Comma separated list of internal maintenance flags.
	MailruSpeedupEnable                 string `json:"mailruSpeedupEnable" default:"true"`                                                                                                                          // Skip full upload if there is another file with same data hash.
	MailruSpeedupFilePatterns           string `json:"mailruSpeedupFilePatterns" default:"*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf"`                                                                          // Comma separated list of file name patterns eligible for speedup (put by hash).
	MailruSpeedupMaxDisk                string `json:"mailruSpeedupMaxDisk" default:"3Gi"`                                                                                                                          // This option allows you to disable speedup (put by hash) for large files.
	MailruSpeedupMaxMemory              string `json:"mailruSpeedupMaxMemory" default:"32Mi"`                                                                                                                       // Files larger than the size given below will always be hashed on disk.
	MailruUser                          string `json:"mailruUser"`                                                                                                                                                  // User name (usually email).
	MailruUserAgent                     string `json:"mailruUserAgent"`                                                                                                                                             // HTTP user agent used internally by client.
	MegaDebug                           string `json:"megaDebug" default:"false"`                                                                                                                                   // Output more debug from Mega.
	MegaEncoding                        string `json:"megaEncoding" default:"Slash,InvalidUtf8,Dot"`                                                                                                                // The encoding for the backend.
	MegaHardDelete                      string `json:"megaHardDelete" default:"false"`                                                                                                                              // Delete files permanently rather than putting them into the trash.
	MegaPass                            string `json:"megaPass"`                                                                                                                                                    // Password.
	MegaUseHttps                        string `json:"megaUseHttps" default:"false"`                                                                                                                                // Use HTTPS for transfers.
	MegaUser                            string `json:"megaUser"`                                                                                                                                                    // User name.
	NetstorageAccount                   string `json:"netstorageAccount"`                                                                                                                                           // Set the NetStorage account name
	NetstorageHost                      string `json:"netstorageHost"`                                                                                                                                              // Domain+path of NetStorage host to connect to.
	NetstorageProtocol                  string `json:"netstorageProtocol" default:"https"`                                                                                                                          // Select between HTTP or HTTPS protocol.
	NetstorageSecret                    string `json:"netstorageSecret"`                                                                                                                                            // Set the NetStorage account secret/G2O key for authentication.
	OnedriveAccessScopes                string `json:"onedriveAccessScopes" default:"Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"`                                  // Set scopes to be requested by rclone.
	OnedriveAuthUrl                     string `json:"onedriveAuthUrl"`                                                                                                                                             // Auth server URL.
	OnedriveChunkSize                   string `json:"onedriveChunkSize" default:"10Mi"`                                                                                                                            // Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
	OnedriveClientId                    string `json:"onedriveClientId"`                                                                                                                                            // OAuth Client Id.
	OnedriveClientSecret                string `json:"onedriveClientSecret"`                                                                                                                                        // OAuth Client Secret.
	OnedriveDisableSitePermission       string `json:"onedriveDisableSitePermission" default:"false"`                                                                                                               // Disable the request for Sites.Read.All permission.
	OnedriveDriveId                     string `json:"onedriveDriveId"`                                                                                                                                             // The ID of the drive to use.
	OnedriveDriveType                   string `json:"onedriveDriveType"`                                                                                                                                           // The type of the drive (personal | business | documentLibrary).
	OnedriveEncoding                    string `json:"onedriveEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
	OnedriveExposeOnenoteFiles          string `json:"onedriveExposeOnenoteFiles" default:"false"`                                                                                                                  // Set to make OneNote files show up in directory listings.
	OnedriveHashType                    string `json:"onedriveHashType" default:"auto"`                                                                                                                             // Specify the hash in use for the backend.
	OnedriveLinkPassword                string `json:"onedriveLinkPassword"`                                                                                                                                        // Set the password for links created by the link command.
	OnedriveLinkScope                   string `json:"onedriveLinkScope" default:"anonymous"`                                                                                                                       // Set the scope of the links created by the link command.
	OnedriveLinkType                    string `json:"onedriveLinkType" default:"view"`                                                                                                                             // Set the type of the links created by the link command.
	OnedriveListChunk                   string `json:"onedriveListChunk" default:"1000"`                                                                                                                            // Size of listing chunk.
	OnedriveNoVersions                  string `json:"onedriveNoVersions" default:"false"`                                                                                                                          // Remove all versions on modifying operations.
	OnedriveRegion                      string `json:"onedriveRegion" default:"global"`                                                                                                                             // Choose national cloud region for OneDrive.
	OnedriveRootFolderId                string `json:"onedriveRootFolderId"`                                                                                                                                        // ID of the root folder.
	OnedriveServerSideAcrossConfigs     string `json:"onedriveServerSideAcrossConfigs" default:"false"`                                                                                                             // Allow server-side operations (e.g. copy) to work across different onedrive configs.
	OnedriveToken                       string `json:"onedriveToken"`                                                                                                                                               // OAuth Access Token as a JSON blob.
	OnedriveTokenUrl                    string `json:"onedriveTokenUrl"`                                                                                                                                            // Token server url.
	OpendriveChunkSize                  string `json:"opendriveChunkSize" default:"10Mi"`                                                                                                                           // Files will be uploaded in chunks this size.
	OpendriveEncoding                   string `json:"opendriveEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot"`   // The encoding for the backend.
	OpendrivePassword                   string `json:"opendrivePassword"`                                                                                                                                           // Password.
	OpendriveUsername                   string `json:"opendriveUsername"`                                                                                                                                           // Username.
	OosChunkSize                        string `json:"oosChunkSize" default:"5Mi"`                                                                                                                                  // Chunk size to use for uploading.
	OosCompartment                      string `json:"oosCompartment"`                                                                                                                                              // Object storage compartment OCID
	OosConfigFile                       string `json:"oosConfigFile" default:"~/.oci/config"`                                                                                                                       // Path to OCI config file
	OosConfigProfile                    string `json:"oosConfigProfile" default:"Default"`                                                                                                                          // Profile name inside the oci config file
	OosCopyCutoff                       string `json:"oosCopyCutoff" default:"4.656Gi"`                                                                                                                             // Cutoff for switching to multipart copy.
	OosCopyTimeout                      string `json:"oosCopyTimeout" default:"1m0s"`                                                                                                                               // Timeout for copy.
	OosDisableChecksum                  string `json:"oosDisableChecksum" default:"false"`                                                                                                                          // Don't store MD5 checksum with object metadata.
	OosEncoding                         string `json:"oosEncoding" default:"Slash,InvalidUtf8,Dot"`                                                                                                                 // The encoding for the backend.
	OosEndpoint                         string `json:"oosEndpoint"`                                                                                                                                                 // Endpoint for Object storage API.
	OosLeavePartsOnError                string `json:"oosLeavePartsOnError" default:"false"`                                                                                                                        // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	OosNamespace                        string `json:"oosNamespace"`                                                                                                                                                // Object storage namespace
	OosNoCheckBucket                    string `json:"oosNoCheckBucket" default:"false"`                                                                                                                            // If set, don't attempt to check the bucket exists or create it.
	OosProvider                         string `json:"oosProvider" default:"env_auth"`                                                                                                                              // Choose your Auth Provider
	OosRegion                           string `json:"oosRegion"`                                                                                                                                                   // Object storage Region
	OosSseCustomerAlgorithm             string `json:"oosSseCustomerAlgorithm"`                                                                                                                                     // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
	OosSseCustomerKey                   string `json:"oosSseCustomerKey"`                                                                                                                                           // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	OosSseCustomerKeyFile               string `json:"oosSseCustomerKeyFile"`                                                                                                                                       // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	OosSseCustomerKeySha256             string `json:"oosSseCustomerKeySha256"`                                                                                                                                     // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	OosSseKmsKeyId                      string `json:"oosSseKmsKeyId"`                                                                                                                                              // if using using your own master key in vault, this header specifies the
	OosStorageTier                      string `json:"oosStorageTier" default:"Standard"`                                                                                                                           // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	OosUploadConcurrency                string `json:"oosUploadConcurrency" default:"10"`                                                                                                                           // Concurrency for multipart uploads.
	OosUploadCutoff                     string `json:"oosUploadCutoff" default:"200Mi"`                                                                                                                             // Cutoff for switching to chunked upload.
	PcloudAuthUrl                       string `json:"pcloudAuthUrl"`                                                                                                                                               // Auth server URL.
	PcloudClientId                      string `json:"pcloudClientId"`                                                                                                                                              // OAuth Client Id.
	PcloudClientSecret                  string `json:"pcloudClientSecret"`                                                                                                                                          // OAuth Client Secret.
	PcloudEncoding                      string `json:"pcloudEncoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"`                                                                                            // The encoding for the backend.
	PcloudHostname                      string `json:"pcloudHostname" default:"api.pcloud.com"`                                                                                                                     // Hostname to connect to.
	PcloudPassword                      string `json:"pcloudPassword"`                                                                                                                                              // Your pcloud password.
	PcloudRootFolderId                  string `json:"pcloudRootFolderId" default:"d0"`                                                                                                                             // Fill in for rclone to use a non root folder as its starting point.
	PcloudToken                         string `json:"pcloudToken"`                                                                                                                                                 // OAuth Access Token as a JSON blob.
	PcloudTokenUrl                      string `json:"pcloudTokenUrl"`                                                                                                                                              // Token server url.
	PcloudUsername                      string `json:"pcloudUsername"`                                                                                                                                              // Your pcloud username.
	PremiumizemeApiKey                  string `json:"premiumizemeApiKey"`                                                                                                                                          // API Key.
	PremiumizemeEncoding                string `json:"premiumizemeEncoding" default:"Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot"`                                                                          // The encoding for the backend.
	PutioEncoding                       string `json:"putioEncoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"`                                                                                             // The encoding for the backend.
	QingstorAccessKeyId                 string `json:"qingstorAccessKeyId"`                                                                                                                                         // QingStor Access Key ID.
	QingstorChunkSize                   string `json:"qingstorChunkSize" default:"4Mi"`                                                                                                                             // Chunk size to use for uploading.
	QingstorConnectionRetries           string `json:"qingstorConnectionRetries" default:"3"`                                                                                                                       // Number of connection retries.
	QingstorEncoding                    string `json:"qingstorEncoding" default:"Slash,Ctl,InvalidUtf8"`                                                                                                            // The encoding for the backend.
	QingstorEndpoint                    string `json:"qingstorEndpoint"`                                                                                                                                            // Enter an endpoint URL to connection QingStor API.
	QingstorEnvAuth                     string `json:"qingstorEnvAuth" default:"false"`                                                                                                                             // Get QingStor credentials from runtime.
	QingstorSecretAccessKey             string `json:"qingstorSecretAccessKey"`                                                                                                                                     // QingStor Secret Access Key (password).
	QingstorUploadConcurrency           string `json:"qingstorUploadConcurrency" default:"1"`                                                                                                                       // Concurrency for multipart uploads.
	QingstorUploadCutoff                string `json:"qingstorUploadCutoff" default:"200Mi"`                                                                                                                        // Cutoff for switching to chunked upload.
	QingstorZone                        string `json:"qingstorZone"`                                                                                                                                                // Zone to connect to.
	S3AccessKeyId                       string `json:"s3AccessKeyId"`                                                                                                                                               // AWS Access Key ID.
	S3Acl                               string `json:"s3Acl"`                                                                                                                                                       // Canned ACL used when creating buckets and storing or copying objects.
	S3BucketAcl                         string `json:"s3BucketAcl"`                                                                                                                                                 // Canned ACL used when creating buckets.
	S3ChunkSize                         string `json:"s3ChunkSize" default:"5Mi"`                                                                                                                                   // Chunk size to use for uploading.
	S3CopyCutoff                        string `json:"s3CopyCutoff" default:"4.656Gi"`                                                                                                                              // Cutoff for switching to multipart copy.
	S3Decompress                        string `json:"s3Decompress" default:"false"`                                                                                                                                // If set this will decompress gzip encoded objects.
	S3DisableChecksum                   string `json:"s3DisableChecksum" default:"false"`                                                                                                                           // Don't store MD5 checksum with object metadata.
	S3DisableHttp2                      string `json:"s3DisableHttp2" default:"false"`                                                                                                                              // Disable usage of http2 for S3 backends.
	S3DownloadUrl                       string `json:"s3DownloadUrl"`                                                                                                                                               // Custom endpoint for downloads.
	S3Encoding                          string `json:"s3Encoding" default:"Slash,InvalidUtf8,Dot"`                                                                                                                  // The encoding for the backend.
	S3Endpoint                          string `json:"s3Endpoint"`                                                                                                                                                  // Endpoint for S3 API.
	S3EnvAuth                           string `json:"s3EnvAuth" default:"false"`                                                                                                                                   // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	S3ForcePathStyle                    string `json:"s3ForcePathStyle" default:"true"`                                                                                                                             // If true use path style access if false use virtual hosted style.
	S3LeavePartsOnError                 string `json:"s3LeavePartsOnError" default:"false"`                                                                                                                         // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	S3ListChunk                         string `json:"s3ListChunk" default:"1000"`                                                                                                                                  // Size of listing chunk (response list for each ListObject S3 request).
	S3ListUrlEncode                     string `json:"s3ListUrlEncode" default:"unset"`                                                                                                                             // Whether to url encode listings: true/false/unset
	S3ListVersion                       string `json:"s3ListVersion" default:"0"`                                                                                                                                   // Version of ListObjects to use: 1,2 or 0 for auto.
	S3LocationConstraint                string `json:"s3LocationConstraint"`                                                                                                                                        // Location constraint - must be set to match the Region.
	S3MaxUploadParts                    string `json:"s3MaxUploadParts" default:"10000"`                                                                                                                            // Maximum number of parts in a multipart upload.
	S3MemoryPoolFlushTime               string `json:"s3MemoryPoolFlushTime" default:"1m0s"`                                                                                                                        // How often internal memory buffer pools will be flushed.
	S3MemoryPoolUseMmap                 string `json:"s3MemoryPoolUseMmap" default:"false"`                                                                                                                         // Whether to use mmap buffers in internal memory pool.
	S3MightGzip                         string `json:"s3MightGzip" default:"unset"`                                                                                                                                 // Set this if the backend might gzip objects.
	S3NoCheckBucket                     string `json:"s3NoCheckBucket" default:"false"`                                                                                                                             // If set, don't attempt to check the bucket exists or create it.
	S3NoHead                            string `json:"s3NoHead" default:"false"`                                                                                                                                    // If set, don't HEAD uploaded objects to check integrity.
	S3NoHeadObject                      string `json:"s3NoHeadObject" default:"false"`                                                                                                                              // If set, do not do HEAD before GET when getting objects.
	S3NoSystemMetadata                  string `json:"s3NoSystemMetadata" default:"false"`                                                                                                                          // Suppress setting and reading of system metadata
	S3Profile                           string `json:"s3Profile"`                                                                                                                                                   // Profile to use in the shared credentials file.
	S3Provider                          string `json:"s3Provider"`                                                                                                                                                  // Choose your S3 provider.
	S3Region                            string `json:"s3Region"`                                                                                                                                                    // Region to connect to.
	S3RequesterPays                     string `json:"s3RequesterPays" default:"false"`                                                                                                                             // Enables requester pays option when interacting with S3 bucket.
	S3SecretAccessKey                   string `json:"s3SecretAccessKey"`                                                                                                                                           // AWS Secret Access Key (password).
	S3ServerSideEncryption              string `json:"s3ServerSideEncryption"`                                                                                                                                      // The server-side encryption algorithm used when storing this object in S3.
	S3SessionToken                      string `json:"s3SessionToken"`                                                                                                                                              // An AWS session token.
	S3SharedCredentialsFile             string `json:"s3SharedCredentialsFile"`                                                                                                                                     // Path to the shared credentials file.
	S3SseCustomerAlgorithm              string `json:"s3SseCustomerAlgorithm"`                                                                                                                                      // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	S3SseCustomerKey                    string `json:"s3SseCustomerKey"`                                                                                                                                            // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	S3SseCustomerKeyBase64              string `json:"s3SseCustomerKeyBase64"`                                                                                                                                      // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	S3SseCustomerKeyMd5                 string `json:"s3SseCustomerKeyMd5"`                                                                                                                                         // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	S3SseKmsKeyId                       string `json:"s3SseKmsKeyId"`                                                                                                                                               // If using KMS ID you must provide the ARN of Key.
	S3StorageClass                      string `json:"s3StorageClass"`                                                                                                                                              // The storage class to use when storing new objects in S3.
	S3StsEndpoint                       string `json:"s3StsEndpoint"`                                                                                                                                               // Endpoint for STS.
	S3UploadConcurrency                 string `json:"s3UploadConcurrency" default:"4"`                                                                                                                             // Concurrency for multipart uploads.
	S3UploadCutoff                      string `json:"s3UploadCutoff" default:"200Mi"`                                                                                                                              // Cutoff for switching to chunked upload.
	S3UseAccelerateEndpoint             string `json:"s3UseAccelerateEndpoint" default:"false"`                                                                                                                     // If true use the AWS S3 accelerated endpoint.
	S3UseMultipartEtag                  string `json:"s3UseMultipartEtag" default:"unset"`                                                                                                                          // Whether to use ETag in multipart uploads for verification
	S3UsePresignedRequest               string `json:"s3UsePresignedRequest" default:"false"`                                                                                                                       // Whether to use a presigned request or PutObject for single part uploads
	S3V2Auth                            string `json:"s3V2Auth" default:"false"`                                                                                                                                    // If true use v2 authentication.
	S3VersionAt                         string `json:"s3VersionAt" default:"off"`                                                                                                                                   // Show file versions as they were at the specified time.
	S3Versions                          string `json:"s3Versions" default:"false"`                                                                                                                                  // Include old versions in directory listings.
	Seafile2fa                          string `json:"seafile2fa" default:"false"`                                                                                                                                  // Two-factor authentication ('true' if the account has 2FA enabled).
	SeafileAuthToken                    string `json:"seafileAuthToken"`                                                                                                                                            // Authentication token.
	SeafileCreateLibrary                string `json:"seafileCreateLibrary" default:"false"`                                                                                                                        // Should rclone create a library if it doesn't exist.
	SeafileEncoding                     string `json:"seafileEncoding" default:"Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8"`                                                                                       // The encoding for the backend.
	SeafileLibrary                      string `json:"seafileLibrary"`                                                                                                                                              // Name of the library.
	SeafileLibraryKey                   string `json:"seafileLibraryKey"`                                                                                                                                           // Library password (for encrypted libraries only).
	SeafilePass                         string `json:"seafilePass"`                                                                                                                                                 // Password.
	SeafileUrl                          string `json:"seafileUrl"`                                                                                                                                                  // URL of seafile host to connect to.
	SeafileUser                         string `json:"seafileUser"`                                                                                                                                                 // User name (usually email address).
	SftpAskPassword                     string `json:"sftpAskPassword" default:"false"`                                                                                                                             // Allow asking for SFTP password when needed.
	SftpChunkSize                       string `json:"sftpChunkSize" default:"32Ki"`                                                                                                                                // Upload and download chunk size.
	SftpCiphers                         string `json:"sftpCiphers"`                                                                                                                                                 // Space separated list of ciphers to be used for session encryption, ordered by preference.
	SftpConcurrency                     string `json:"sftpConcurrency" default:"64"`                                                                                                                                // The maximum number of outstanding requests for one file
	SftpDisableConcurrentReads          string `json:"sftpDisableConcurrentReads" default:"false"`                                                                                                                  // If set don't use concurrent reads.
	SftpDisableConcurrentWrites         string `json:"sftpDisableConcurrentWrites" default:"false"`                                                                                                                 // If set don't use concurrent writes.
	SftpDisableHashcheck                string `json:"sftpDisableHashcheck" default:"false"`                                                                                                                        // Disable the execution of SSH commands to determine if remote file hashing is available.
	SftpHost                            string `json:"sftpHost"`                                                                                                                                                    // SSH host to connect to.
	SftpIdleTimeout                     string `json:"sftpIdleTimeout" default:"1m0s"`                                                                                                                              // Max time before closing idle connections.
	SftpKeyExchange                     string `json:"sftpKeyExchange"`                                                                                                                                             // Space separated list of key exchange algorithms, ordered by preference.
	SftpKeyFile                         string `json:"sftpKeyFile"`                                                                                                                                                 // Path to PEM-encoded private key file.
	SftpKeyFilePass                     string `json:"sftpKeyFilePass"`                                                                                                                                             // The passphrase to decrypt the PEM-encoded private key file.
	SftpKeyPem                          string `json:"sftpKeyPem"`                                                                                                                                                  // Raw PEM-encoded private key.
	SftpKeyUseAgent                     string `json:"sftpKeyUseAgent" default:"false"`                                                                                                                             // When set forces the usage of the ssh-agent.
	SftpKnownHostsFile                  string `json:"sftpKnownHostsFile"`                                                                                                                                          // Optional path to known_hosts file.
	SftpMacs                            string `json:"sftpMacs"`                                                                                                                                                    // Space separated list of MACs (message authentication code) algorithms, ordered by preference.
	SftpMd5sumCommand                   string `json:"sftpMd5sumCommand"`                                                                                                                                           // The command used to read md5 hashes.
	SftpPass                            string `json:"sftpPass"`                                                                                                                                                    // SSH password, leave blank to use ssh-agent.
	SftpPathOverride                    string `json:"sftpPathOverride"`                                                                                                                                            // Override path used by SSH shell commands.
	SftpPort                            string `json:"sftpPort" default:"22"`                                                                                                                                       // SSH port number.
	SftpPubkeyFile                      string `json:"sftpPubkeyFile"`                                                                                                                                              // Optional path to public key file.
	SftpServerCommand                   string `json:"sftpServerCommand"`                                                                                                                                           // Specifies the path or command to run a sftp server on the remote host.
	SftpSetEnv                          string `json:"sftpSetEnv"`                                                                                                                                                  // Environment variables to pass to sftp and commands
	SftpSetModtime                      string `json:"sftpSetModtime" default:"true"`                                                                                                                               // Set the modified time on the remote if set.
	SftpSha1sumCommand                  string `json:"sftpSha1sumCommand"`                                                                                                                                          // The command used to read sha1 hashes.
	SftpShellType                       string `json:"sftpShellType"`                                                                                                                                               // The type of SSH shell on remote server, if any.
	SftpSkipLinks                       string `json:"sftpSkipLinks" default:"false"`                                                                                                                               // Set to skip any symlinks and any other non regular files.
	SftpSubsystem                       string `json:"sftpSubsystem" default:"sftp"`                                                                                                                                // Specifies the SSH2 subsystem on the remote host.
	SftpUseFstat                        string `json:"sftpUseFstat" default:"false"`                                                                                                                                // If set use fstat instead of stat.
	SftpUseInsecureCipher               string `json:"sftpUseInsecureCipher" default:"false"`                                                                                                                       // Enable the use of insecure ciphers and key exchange methods.
	SftpUser                            string `json:"sftpUser" default:"shane"`                                                                                                                                    // SSH username.
	SharefileChunkSize                  string `json:"sharefileChunkSize" default:"64Mi"`                                                                                                                           // Upload chunk size.
	SharefileEncoding                   string `json:"sharefileEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot"`   // The encoding for the backend.
	SharefileEndpoint                   string `json:"sharefileEndpoint"`                                                                                                                                           // Endpoint for API calls.
	SharefileRootFolderId               string `json:"sharefileRootFolderId"`                                                                                                                                       // ID of the root folder.
	SharefileUploadCutoff               string `json:"sharefileUploadCutoff" default:"128Mi"`                                                                                                                       // Cutoff for switching to multipart upload.
	SiaApiPassword                      string `json:"siaApiPassword"`                                                                                                                                              // Sia Daemon API Password.
	SiaApiUrl                           string `json:"siaApiUrl" default:"http://127.0.0.1:9980"`                                                                                                                   // Sia daemon API URL, like http://sia.daemon.host:9980.
	SiaEncoding                         string `json:"siaEncoding" default:"Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot"`                                                                                   // The encoding for the backend.
	SiaUserAgent                        string `json:"siaUserAgent" default:"Sia-Agent"`                                                                                                                            // Siad User Agent
	SmbCaseInsensitive                  string `json:"smbCaseInsensitive" default:"true"`                                                                                                                           // Whether the server is configured to be case-insensitive.
	SmbDomain                           string `json:"smbDomain" default:"WORKGROUP"`                                                                                                                               // Domain name for NTLM authentication.
	SmbEncoding                         string `json:"smbEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot"`                              // The encoding for the backend.
	SmbHideSpecialShare                 string `json:"smbHideSpecialShare" default:"true"`                                                                                                                          // Hide special shares (e.g. print$) which users aren't supposed to access.
	SmbHost                             string `json:"smbHost"`                                                                                                                                                     // SMB server hostname to connect to.
	SmbIdleTimeout                      string `json:"smbIdleTimeout" default:"1m0s"`                                                                                                                               // Max time before closing idle connections.
	SmbPass                             string `json:"smbPass"`                                                                                                                                                     // SMB password.
	SmbPort                             string `json:"smbPort" default:"445"`                                                                                                                                       // SMB port number.
	SmbSpn                              string `json:"smbSpn"`                                                                                                                                                      // Service principal name.
	SmbUser                             string `json:"smbUser" default:"shane"`                                                                                                                                     // SMB username.
	StorjAccessGrant                    string `json:"storjAccessGrant"`                                                                                                                                            // Access grant.
	StorjApiKey                         string `json:"storjApiKey"`                                                                                                                                                 // API key.
	StorjPassphrase                     string `json:"storjPassphrase"`                                                                                                                                             // Encryption passphrase.
	StorjProvider                       string `json:"storjProvider" default:"existing"`                                                                                                                            // Choose an authentication method.
	StorjSatelliteAddress               string `json:"storjSatelliteAddress" default:"us1.storj.io"`                                                                                                                // Satellite address.
	SugarsyncAccessKeyId                string `json:"sugarsyncAccessKeyId"`                                                                                                                                        // Sugarsync Access Key ID.
	SugarsyncAppId                      string `json:"sugarsyncAppId"`                                                                                                                                              // Sugarsync App ID.
	SugarsyncAuthorization              string `json:"sugarsyncAuthorization"`                                                                                                                                      // Sugarsync authorization.
	SugarsyncAuthorizationExpiry        string `json:"sugarsyncAuthorizationExpiry"`                                                                                                                                // Sugarsync authorization expiry.
	SugarsyncDeletedId                  string `json:"sugarsyncDeletedId"`                                                                                                                                          // Sugarsync deleted folder id.
	SugarsyncEncoding                   string `json:"sugarsyncEncoding" default:"Slash,Ctl,InvalidUtf8,Dot"`                                                                                                       // The encoding for the backend.
	SugarsyncHardDelete                 string `json:"sugarsyncHardDelete" default:"false"`                                                                                                                         // Permanently delete files if true
	SugarsyncPrivateAccessKey           string `json:"sugarsyncPrivateAccessKey"`                                                                                                                                   // Sugarsync Private Access Key.
	SugarsyncRefreshToken               string `json:"sugarsyncRefreshToken"`                                                                                                                                       // Sugarsync refresh token.
	SugarsyncRootId                     string `json:"sugarsyncRootId"`                                                                                                                                             // Sugarsync root id.
	SugarsyncUser                       string `json:"sugarsyncUser"`                                                                                                                                               // Sugarsync user.
	SwiftApplicationCredentialId        string `json:"swiftApplicationCredentialId"`                                                                                                                                // Application Credential ID (OS_APPLICATION_CREDENTIAL_ID).
	SwiftApplicationCredentialName      string `json:"swiftApplicationCredentialName"`                                                                                                                              // Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME).
	SwiftApplicationCredentialSecret    string `json:"swiftApplicationCredentialSecret"`                                                                                                                            // Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET).
	SwiftAuth                           string `json:"swiftAuth"`                                                                                                                                                   // Authentication URL for server (OS_AUTH_URL).
	SwiftAuthToken                      string `json:"swiftAuthToken"`                                                                                                                                              // Auth Token from alternate authentication - optional (OS_AUTH_TOKEN).
	SwiftAuthVersion                    string `json:"swiftAuthVersion" default:"0"`                                                                                                                                // AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).
	SwiftChunkSize                      string `json:"swiftChunkSize" default:"5Gi"`                                                                                                                                // Above this size files will be chunked into a _segments container.
	SwiftDomain                         string `json:"swiftDomain"`                                                                                                                                                 // User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME)
	SwiftEncoding                       string `json:"swiftEncoding" default:"Slash,InvalidUtf8"`                                                                                                                   // The encoding for the backend.
	SwiftEndpointType                   string `json:"swiftEndpointType" default:"public"`                                                                                                                          // Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).
	SwiftEnvAuth                        string `json:"swiftEnvAuth" default:"false"`                                                                                                                                // Get swift credentials from environment variables in standard OpenStack form.
	SwiftKey                            string `json:"swiftKey"`                                                                                                                                                    // API key or password (OS_PASSWORD).
	SwiftLeavePartsOnError              string `json:"swiftLeavePartsOnError" default:"false"`                                                                                                                      // If true avoid calling abort upload on a failure.
	SwiftNoChunk                        string `json:"swiftNoChunk" default:"false"`                                                                                                                                // Don't chunk files during streaming upload.
	SwiftNoLargeObjects                 string `json:"swiftNoLargeObjects" default:"false"`                                                                                                                         // Disable support for static and dynamic large objects
	SwiftRegion                         string `json:"swiftRegion"`                                                                                                                                                 // Region name - optional (OS_REGION_NAME).
	SwiftStoragePolicy                  string `json:"swiftStoragePolicy"`                                                                                                                                          // The storage policy to use when creating a new container.
	SwiftStorageUrl                     string `json:"swiftStorageUrl"`                                                                                                                                             // Storage URL - optional (OS_STORAGE_URL).
	SwiftTenant                         string `json:"swiftTenant"`                                                                                                                                                 // Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME).
	SwiftTenantDomain                   string `json:"swiftTenantDomain"`                                                                                                                                           // Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME).
	SwiftTenantId                       string `json:"swiftTenantId"`                                                                                                                                               // Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID).
	SwiftUser                           string `json:"swiftUser"`                                                                                                                                                   // User name to log in (OS_USERNAME).
	SwiftUserId                         string `json:"swiftUserId"`                                                                                                                                                 // User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID).
	UptoboxAccessToken                  string `json:"uptoboxAccessToken"`                                                                                                                                          // Your access token.
	UptoboxEncoding                     string `json:"uptoboxEncoding" default:"Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot"`                                                                // The encoding for the backend.
	WebdavBearerToken                   string `json:"webdavBearerToken"`                                                                                                                                           // Bearer token instead of user/pass (e.g. a Macaroon).
	WebdavBearerTokenCommand            string `json:"webdavBearerTokenCommand"`                                                                                                                                    // Command to run to get a bearer token.
	WebdavEncoding                      string `json:"webdavEncoding"`                                                                                                                                              // The encoding for the backend.
	WebdavHeaders                       string `json:"webdavHeaders"`                                                                                                                                               // Set HTTP headers for all transactions.
	WebdavPass                          string `json:"webdavPass"`                                                                                                                                                  // Password.
	WebdavUrl                           string `json:"webdavUrl"`                                                                                                                                                   // URL of http host to connect to.
	WebdavUser                          string `json:"webdavUser"`                                                                                                                                                  // User name.
	WebdavVendor                        string `json:"webdavVendor"`                                                                                                                                                // Name of the WebDAV site/service/software you are using.
	YandexAuthUrl                       string `json:"yandexAuthUrl"`                                                                                                                                               // Auth server URL.
	YandexClientId                      string `json:"yandexClientId"`                                                                                                                                              // OAuth Client Id.
	YandexClientSecret                  string `json:"yandexClientSecret"`                                                                                                                                          // OAuth Client Secret.
	YandexEncoding                      string `json:"yandexEncoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"`                                                                                                      // The encoding for the backend.
	YandexHardDelete                    string `json:"yandexHardDelete" default:"false"`                                                                                                                            // Delete files permanently rather than putting them into the trash.
	YandexToken                         string `json:"yandexToken"`                                                                                                                                                 // OAuth Access Token as a JSON blob.
	YandexTokenUrl                      string `json:"yandexTokenUrl"`                                                                                                                                              // Token server url.
	ZohoAuthUrl                         string `json:"zohoAuthUrl"`                                                                                                                                                 // Auth server URL.
	ZohoClientId                        string `json:"zohoClientId"`                                                                                                                                                // OAuth Client Id.
	ZohoClientSecret                    string `json:"zohoClientSecret"`                                                                                                                                            // OAuth Client Secret.
	ZohoEncoding                        string `json:"zohoEncoding" default:"Del,Ctl,InvalidUtf8"`                                                                                                                  // The encoding for the backend.
	ZohoRegion                          string `json:"zohoRegion"`                                                                                                                                                  // Zoho region to connect to.
	ZohoToken                           string `json:"zohoToken"`                                                                                                                                                   // OAuth Access Token as a JSON blob.
	ZohoTokenUrl                        string `json:"zohoTokenUrl"`                                                                                                                                                // Token server url.
}
