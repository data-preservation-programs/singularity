// Code generated. DO NOT EDIT.
package datasource

type AcdRequest struct {
	SourcePath        string `json:"sourcePath"`                               // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                        // Delete the source after exporting to CAR files
	ClientId          string `json:"clientId"`                                 // OAuth Client Id.
	Checkpoint        string `json:"checkpoint"`                               // Checkpoint for internal polling (debug).
	UploadWaitPerGb   string `json:"uploadWaitPerGb" default:"3m0s"`           // Additional time per GiB to wait after a failed complete upload to see if it appears.
	ClientSecret      string `json:"clientSecret"`                             // OAuth Client Secret.
	Token             string `json:"token"`                                    // OAuth Access Token as a JSON blob.
	AuthUrl           string `json:"authUrl"`                                  // Auth server URL.
	TokenUrl          string `json:"tokenUrl"`                                 // Token server url.
	TemplinkThreshold string `json:"templinkThreshold" default:"9Gi"`          // Files >= this size will be downloaded via their tempLink.
	Encoding          string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
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
	SourcePath                 string `json:"sourcePath"`                                                         // The path of the source to scan items
	DeleteAfterExport          string `json:"deleteAfterExport"`                                                  // Delete the source after exporting to CAR files
	ListChunk                  string `json:"listChunk" default:"5000"`                                           // Size of blob list.
	PublicAccess               string `json:"publicAccess"`                                                       // Public access level of a container: blob or container.
	NoHeadObject               string `json:"noHeadObject" default:"false"`                                       // If set, do not do HEAD before GET when getting objects.
	ClientCertificatePassword  string `json:"clientCertificatePassword"`                                          // Password for the certificate file (optional).
	ClientSendCertificateChain string `json:"clientSendCertificateChain" default:"false"`                         // Send the certificate chain when using certificate auth.
	UseMsi                     string `json:"useMsi" default:"false"`                                             // Use a managed service identity to authenticate (only works in Azure).
	MsiMiResId                 string `json:"msiMiResId"`                                                         // Azure resource ID of the user-assigned MSI to use, if any.
	ClientSecret               string `json:"clientSecret"`                                                       // One of the service principal's client secrets
	Password                   string `json:"password"`                                                           // The user's password
	Key                        string `json:"key"`                                                                // Storage Account Shared Key.
	ArchiveTierDelete          string `json:"archiveTierDelete" default:"false"`                                  // Delete archive tier blobs before overwriting.
	MemoryPoolFlushTime        string `json:"memoryPoolFlushTime" default:"1m0s"`                                 // How often internal memory buffer pools will be flushed.
	UploadCutoff               string `json:"uploadCutoff"`                                                       // Cutoff for switching to chunked upload (<= 256 MiB) (deprecated).
	ChunkSize                  string `json:"chunkSize" default:"4Mi"`                                            // Upload chunk size.
	Tenant                     string `json:"tenant"`                                                             // ID of the service principal's tenant. Also called its directory ID.
	ClientId                   string `json:"clientId"`                                                           // The ID of the client in use.
	ClientCertificatePath      string `json:"clientCertificatePath"`                                              // Path to a PEM or PKCS12 certificate file including the private key.
	Encoding                   string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"` // The encoding for the backend.
	Account                    string `json:"account"`                                                            // Azure Storage Account Name.
	Username                   string `json:"username"`                                                           // User name (usually an email address)
	Endpoint                   string `json:"endpoint"`                                                           // Endpoint for the service.
	NoCheckContainer           string `json:"noCheckContainer" default:"false"`                                   // If set, don't attempt to check the container exists or create it.
	EnvAuth                    string `json:"envAuth" default:"false"`                                            // Read credentials from runtime (environment variables, CLI or MSI).
	ServicePrincipalFile       string `json:"servicePrincipalFile"`                                               // Path to file containing credentials for use with a service principal.
	UploadConcurrency          string `json:"uploadConcurrency" default:"16"`                                     // Concurrency for multipart uploads.
	AccessTier                 string `json:"accessTier"`                                                         // Access tier of blob: hot, cool or archive.
	DisableChecksum            string `json:"disableChecksum" default:"false"`                                    // Don't store MD5 checksum with object metadata.
	MemoryPoolUseMmap          string `json:"memoryPoolUseMmap" default:"false"`                                  // Whether to use mmap buffers in internal memory pool.
	SasUrl                     string `json:"sasUrl"`                                                             // SAS URL for container level access only.
	MsiObjectId                string `json:"msiObjectId"`                                                        // Object ID of the user-assigned MSI to use, if any.
	MsiClientId                string `json:"msiClientId"`                                                        // Object ID of the user-assigned MSI to use, if any.
	UseEmulator                string `json:"useEmulator" default:"false"`                                        // Uses local storage emulator if provided as 'true'.
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
	SourcePath           string `json:"sourcePath"`                                                 // The path of the source to scan items
	DeleteAfterExport    string `json:"deleteAfterExport"`                                          // Delete the source after exporting to CAR files
	VersionAt            string `json:"versionAt" default:"off"`                                    // Show file versions as they were at the specified time.
	UploadCutoff         string `json:"uploadCutoff" default:"200Mi"`                               // Cutoff for switching to chunked upload.
	CopyCutoff           string `json:"copyCutoff" default:"4Gi"`                                   // Cutoff for switching to multipart copy.
	ChunkSize            string `json:"chunkSize" default:"96Mi"`                                   // Upload chunk size.
	Account              string `json:"account"`                                                    // Account ID or Application Key ID.
	Key                  string `json:"key"`                                                        // Application Key.
	HardDelete           string `json:"hardDelete" default:"false"`                                 // Permanently delete files on remote removal, otherwise hide files.
	Endpoint             string `json:"endpoint"`                                                   // Endpoint for the service.
	Versions             string `json:"versions" default:"false"`                                   // Include old versions in directory listings.
	DisableChecksum      string `json:"disableChecksum" default:"false"`                            // Disable checksums for large (> upload cutoff) files.
	DownloadAuthDuration string `json:"downloadAuthDuration" default:"1w"`                          // Time before the authorization token will expire in s or suffix ms|s|m|h|d.
	MemoryPoolFlushTime  string `json:"memoryPoolFlushTime" default:"1m0s"`                         // How often internal memory buffer pools will be flushed.
	Encoding             string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	TestMode             string `json:"testMode"`                                                   // A flag string for X-Bz-Test-Mode header for debugging.
	DownloadUrl          string `json:"downloadUrl"`                                                // Custom endpoint for downloads.
	MemoryPoolUseMmap    string `json:"memoryPoolUseMmap" default:"false"`                          // Whether to use mmap buffers in internal memory pool.
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
	SourcePath        string `json:"sourcePath"`                                                            // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                                     // Delete the source after exporting to CAR files
	ClientId          string `json:"clientId"`                                                              // OAuth Client Id.
	TokenUrl          string `json:"tokenUrl"`                                                              // Token server url.
	RootFolderId      string `json:"rootFolderId" default:"0"`                                              // Fill in for rclone to use a non root folder as its starting point.
	ListChunk         string `json:"listChunk" default:"1000"`                                              // Size of listing chunk 1-1000.
	Encoding          string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
	ClientSecret      string `json:"clientSecret"`                                                          // OAuth Client Secret.
	BoxConfigFile     string `json:"boxConfigFile"`                                                         // Box App config.json location
	AccessToken       string `json:"accessToken"`                                                           // Box App Primary Access Token
	UploadCutoff      string `json:"uploadCutoff" default:"50Mi"`                                           // Cutoff for switching to multipart upload (>= 50 MiB).
	OwnedBy           string `json:"ownedBy"`                                                               // Only show items owned by the login (email address) passed in.
	Token             string `json:"token"`                                                                 // OAuth Access Token as a JSON blob.
	AuthUrl           string `json:"authUrl"`                                                               // Auth server URL.
	BoxSubType        string `json:"boxSubType" default:"user"`                                             //
	CommitRetries     string `json:"commitRetries" default:"100"`                                           // Max number of times to try committing a multipart file.
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
	SourcePath              string `json:"sourcePath"`                              // The path of the source to scan items
	DeleteAfterExport       string `json:"deleteAfterExport"`                       // Delete the source after exporting to CAR files
	FilenameEncryption      string `json:"filenameEncryption" default:"standard"`   // How to encrypt the filenames.
	Password                string `json:"password"`                                // Password or pass phrase for encryption.
	NoDataEncryption        string `json:"noDataEncryption" default:"false"`        // Option to either encrypt file data or leave it unencrypted.
	FilenameEncoding        string `json:"filenameEncoding" default:"base32"`       // How to encode the encrypted filename to text string.
	ShowMapping             string `json:"showMapping" default:"false"`             // For all files listed show how the names encrypt.
	Remote                  string `json:"remote"`                                  // Remote to encrypt/decrypt.
	DirectoryNameEncryption string `json:"directoryNameEncryption" default:"true"`  // Option to either encrypt directory names or leave them intact.
	Password2               string `json:"password2"`                               // Password or pass phrase for salt.
	ServerSideAcrossConfigs string `json:"serverSideAcrossConfigs" default:"false"` // Allow server-side operations (e.g. copy) to work across different crypt configs.
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
	SourcePath                string `json:"sourcePath"`                                 // The path of the source to scan items
	DeleteAfterExport         string `json:"deleteAfterExport"`                          // Delete the source after exporting to CAR files
	ClientSecret              string `json:"clientSecret"`                               // OAuth Client Secret.
	SkipChecksumGphotos       string `json:"skipChecksumGphotos" default:"false"`        // Skip MD5 checksum on Google photos and videos only.
	Formats                   string `json:"formats"`                                    // Deprecated: See export_formats.
	DisableHttp2              string `json:"disableHttp2" default:"true"`                // Disable drive using http2.
	StopOnUploadLimit         string `json:"stopOnUploadLimit" default:"false"`          // Make upload limit errors be fatal.
	ClientId                  string `json:"clientId"`                                   // Google Application Client Id
	TeamDrive                 string `json:"teamDrive"`                                  // ID of the Shared Drive (Team Drive).
	AuthOwnerOnly             string `json:"authOwnerOnly" default:"false"`              // Only consider files owned by the authenticated user.
	Impersonate               string `json:"impersonate"`                                // Impersonate this user when using a service account.
	AlternateExport           string `json:"alternateExport" default:"false"`            // Deprecated: No longer needed.
	AcknowledgeAbuse          string `json:"acknowledgeAbuse" default:"false"`           // Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
	PacerBurst                string `json:"pacerBurst" default:"100"`                   // Number of API calls to allow without sleeping.
	StopOnDownloadLimit       string `json:"stopOnDownloadLimit" default:"false"`        // Make download limit errors be fatal.
	Token                     string `json:"token"`                                      // OAuth Access Token as a JSON blob.
	SkipGdocs                 string `json:"skipGdocs" default:"false"`                  // Skip google documents in all listings.
	ExportFormats             string `json:"exportFormats" default:"docx,xlsx,pptx,svg"` // Comma separated list of preferred formats for downloading Google docs.
	UseSharedDate             string `json:"useSharedDate" default:"false"`              // Use date file was shared instead of modified date.
	ChunkSize                 string `json:"chunkSize" default:"8Mi"`                    // Upload chunk size.
	V2DownloadMinSize         string `json:"v2DownloadMinSize" default:"off"`            // If Object's are greater, use drive v2 API to download.
	CopyShortcutContent       string `json:"copyShortcutContent" default:"false"`        // Server side copy contents of shortcuts instead of the shortcut.
	SkipShortcuts             string `json:"skipShortcuts" default:"false"`              // If set skip shortcut files.
	ImportFormats             string `json:"importFormats"`                              // Comma separated list of preferred formats for uploading Google docs.
	AllowImportNameChange     string `json:"allowImportNameChange" default:"false"`      // Allow the filetype to change when uploading Google docs.
	PacerMinSleep             string `json:"pacerMinSleep" default:"100ms"`              // Minimum time to sleep between API calls.
	ResourceKey               string `json:"resourceKey"`                                // Resource key for accessing a link-shared file.
	TrashedOnly               string `json:"trashedOnly" default:"false"`                // Only show files that are in the trash.
	SharedWithMe              string `json:"sharedWithMe" default:"false"`               // Only show files that are shared with me.
	UploadCutoff              string `json:"uploadCutoff" default:"8Mi"`                 // Cutoff for switching to chunked upload.
	ServerSideAcrossConfigs   string `json:"serverSideAcrossConfigs" default:"false"`    // Allow server-side operations (e.g. copy) to work across different drive configs.
	Scope                     string `json:"scope"`                                      // Scope that rclone should use when requesting access from drive.
	RootFolderId              string `json:"rootFolderId"`                               // ID of the root folder.
	ServiceAccountCredentials string `json:"serviceAccountCredentials"`                  // Service Account Credentials JSON blob.
	UseCreatedDate            string `json:"useCreatedDate" default:"false"`             // Use file created date instead of modified date.
	KeepRevisionForever       string `json:"keepRevisionForever" default:"false"`        // Keep new head revision of each file forever.
	SizeAsQuota               string `json:"sizeAsQuota" default:"false"`                // Show sizes as storage quota usage, not actual size.
	AuthUrl                   string `json:"authUrl"`                                    // Auth server URL.
	ServiceAccountFile        string `json:"serviceAccountFile"`                         // Service Account Credentials JSON file path.
	UseTrash                  string `json:"useTrash" default:"true"`                    // Send files to the trash instead of deleting permanently.
	StarredOnly               string `json:"starredOnly" default:"false"`                // Only show files that are starred.
	ListChunk                 string `json:"listChunk" default:"1000"`                   // Size of listing chunk 100-1000, 0 to disable.
	SkipDanglingShortcuts     string `json:"skipDanglingShortcuts" default:"false"`      // If set skip dangling shortcut files.
	Encoding                  string `json:"encoding" default:"InvalidUtf8"`             // The encoding for the backend.
	TokenUrl                  string `json:"tokenUrl"`                                   // Token server url.
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
	SourcePath         string `json:"sourcePath"`                                                        // The path of the source to scan items
	DeleteAfterExport  string `json:"deleteAfterExport"`                                                 // Delete the source after exporting to CAR files
	SharedFolders      string `json:"sharedFolders" default:"false"`                                     // Instructs rclone to work on shared folders.
	BatchSize          string `json:"batchSize" default:"0"`                                             // Max number of files in upload batch.
	BatchCommitTimeout string `json:"batchCommitTimeout" default:"10m0s"`                                // Max time to wait for a batch to finish committing
	SharedFiles        string `json:"sharedFiles" default:"false"`                                       // Instructs rclone to work on individual shared files.
	ClientSecret       string `json:"clientSecret"`                                                      // OAuth Client Secret.
	AuthUrl            string `json:"authUrl"`                                                           // Auth server URL.
	TokenUrl           string `json:"tokenUrl"`                                                          // Token server url.
	Impersonate        string `json:"impersonate"`                                                       // Impersonate this user when using a business account.
	ClientId           string `json:"clientId"`                                                          // OAuth Client Id.
	ChunkSize          string `json:"chunkSize" default:"48Mi"`                                          // Upload chunk size (< 150Mi).
	BatchMode          string `json:"batchMode" default:"sync"`                                          // Upload file batching sync|async|off.
	Encoding           string `json:"encoding" default:"Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
	Token              string `json:"token"`                                                             // OAuth Access Token as a JSON blob.
	BatchTimeout       string `json:"batchTimeout" default:"0s"`                                         // Max time to allow an idle upload batch before uploading.
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
	SourcePath        string `json:"sourcePath"`                                                                                                                    // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                                                                                             // Delete the source after exporting to CAR files
	FolderPassword    string `json:"folderPassword"`                                                                                                                // If you want to list the files in a shared folder that is password protected, add this parameter.
	Encoding          string `json:"encoding" default:"Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
	ApiKey            string `json:"apiKey"`                                                                                                                        // Your API Key, get it from https://1fichier.com/console/params.pl.
	SharedFolder      string `json:"sharedFolder"`                                                                                                                  // If you want to download a shared folder, add this parameter.
	FilePassword      string `json:"filePassword"`                                                                                                                  // If you want to download a shared file that is password protected, add this parameter.
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
	SourcePath        string `json:"sourcePath"`                                       // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                // Delete the source after exporting to CAR files
	Encoding          string `json:"encoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	Url               string `json:"url"`                                              // URL of the Enterprise File Fabric to connect to.
	RootFolderId      string `json:"rootFolderId"`                                     // ID of the root folder.
	PermanentToken    string `json:"permanentToken"`                                   // Permanent Authentication Token.
	Token             string `json:"token"`                                            // Session Token.
	TokenExpiry       string `json:"tokenExpiry"`                                      // Token expiry time.
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
	SourcePath         string `json:"sourcePath"`                                      // The path of the source to scan items
	DeleteAfterExport  string `json:"deleteAfterExport"`                               // Delete the source after exporting to CAR files
	Encoding           string `json:"encoding" default:"Slash,Del,Ctl,RightSpace,Dot"` // The encoding for the backend.
	ForceListHidden    string `json:"forceListHidden" default:"false"`                 // Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.
	ShutTimeout        string `json:"shutTimeout" default:"1m0s"`                      // Maximum time to wait for data connection closing status.
	DisableEpsv        string `json:"disableEpsv" default:"false"`                     // Disable using EPSV even if server advertises support.
	TlsCacheSize       string `json:"tlsCacheSize" default:"32"`                       // Size of TLS session cache for all control and data connections.
	DisableTls13       string `json:"disableTls13" default:"false"`                    // Disable TLS 1.3 (workaround for FTP servers with buggy TLS)
	AskPassword        string `json:"askPassword" default:"false"`                     // Allow asking for FTP password when needed.
	Host               string `json:"host"`                                            // FTP host to connect to.
	ExplicitTls        string `json:"explicitTls" default:"false"`                     // Use Explicit FTPS (FTP over TLS).
	DisableMlsd        string `json:"disableMlsd" default:"false"`                     // Disable using MLSD even if server advertises support.
	IdleTimeout        string `json:"idleTimeout" default:"1m0s"`                      // Max time before closing idle connections.
	Pass               string `json:"pass"`                                            // FTP password.
	Tls                string `json:"tls" default:"false"`                             // Use Implicit FTPS (FTP over TLS).
	Concurrency        string `json:"concurrency" default:"0"`                         // Maximum number of FTP simultaneous connections, 0 for unlimited.
	NoCheckCertificate string `json:"noCheckCertificate" default:"false"`              // Do not verify the TLS certificate of the server.
	DisableUtf8        string `json:"disableUtf8" default:"false"`                     // Disable using UTF-8 even if server advertises support.
	WritingMdtm        string `json:"writingMdtm" default:"false"`                     // Use MDTM to set modification time (VsFtpd quirk)
	CloseTimeout       string `json:"closeTimeout" default:"1m0s"`                     // Maximum time to wait for a response to close.
	User               string `json:"user" default:"shane"`                            // FTP username.
	Port               string `json:"port" default:"21"`                               // FTP port number.
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
	SourcePath                string `json:"sourcePath"`                                    // The path of the source to scan items
	DeleteAfterExport         string `json:"deleteAfterExport"`                             // Delete the source after exporting to CAR files
	EnvAuth                   string `json:"envAuth" default:"false"`                       // Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
	Token                     string `json:"token"`                                         // OAuth Access Token as a JSON blob.
	Location                  string `json:"location"`                                      // Location for the newly created buckets.
	StorageClass              string `json:"storageClass"`                                  // The storage class to use when storing objects in Google Cloud Storage.
	Decompress                string `json:"decompress" default:"false"`                    // If set this will decompress gzip encoded objects.
	Endpoint                  string `json:"endpoint"`                                      // Endpoint for the service.
	AuthUrl                   string `json:"authUrl"`                                       // Auth server URL.
	TokenUrl                  string `json:"tokenUrl"`                                      // Token server url.
	ServiceAccountFile        string `json:"serviceAccountFile"`                            // Service Account Credentials JSON file path.
	ObjectAcl                 string `json:"objectAcl"`                                     // Access Control List for new objects.
	Encoding                  string `json:"encoding" default:"Slash,CrLf,InvalidUtf8,Dot"` // The encoding for the backend.
	NoCheckBucket             string `json:"noCheckBucket" default:"false"`                 // If set, don't attempt to check the bucket exists or create it.
	ClientId                  string `json:"clientId"`                                      // OAuth Client Id.
	ServiceAccountCredentials string `json:"serviceAccountCredentials"`                     // Service Account Credentials JSON blob.
	Anonymous                 string `json:"anonymous" default:"false"`                     // Access public buckets and objects without credentials.
	BucketAcl                 string `json:"bucketAcl"`                                     // Access Control List for new buckets.
	BucketPolicyOnly          string `json:"bucketPolicyOnly" default:"false"`              // Access checks should use bucket-level IAM policies.
	ClientSecret              string `json:"clientSecret"`                                  // OAuth Client Secret.
	ProjectNumber             string `json:"projectNumber"`                                 // Project number.
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
	SourcePath        string `json:"sourcePath"`                                    // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                             // Delete the source after exporting to CAR files
	ClientSecret      string `json:"clientSecret"`                                  // OAuth Client Secret.
	AuthUrl           string `json:"authUrl"`                                       // Auth server URL.
	TokenUrl          string `json:"tokenUrl"`                                      // Token server url.
	ReadSize          string `json:"readSize" default:"false"`                      // Set to read the size of media items.
	Encoding          string `json:"encoding" default:"Slash,CrLf,InvalidUtf8,Dot"` // The encoding for the backend.
	ClientId          string `json:"clientId"`                                      // OAuth Client Id.
	Token             string `json:"token"`                                         // OAuth Access Token as a JSON blob.
	ReadOnly          string `json:"readOnly" default:"false"`                      // Set to make the Google Photos backend read only.
	StartYear         string `json:"startYear" default:"2000"`                      // Year limits the photos to be downloaded to those which are uploaded after the given year.
	IncludeArchived   string `json:"includeArchived" default:"false"`               // Also view and download archived media.
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
	SourcePath             string `json:"sourcePath"`                                             // The path of the source to scan items
	DeleteAfterExport      string `json:"deleteAfterExport"`                                      // Delete the source after exporting to CAR files
	DataTransferProtection string `json:"dataTransferProtection"`                                 // Kerberos data transfer protection: authentication|integrity|privacy.
	Encoding               string `json:"encoding" default:"Slash,Colon,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	Namenode               string `json:"namenode"`                                               // Hadoop name node and port.
	Username               string `json:"username"`                                               // Hadoop user name.
	ServicePrincipalName   string `json:"servicePrincipalName"`                                   // Kerberos service principal name for the namenode.
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
	SourcePath                 string `json:"sourcePath"`                                            // The path of the source to scan items
	DeleteAfterExport          string `json:"deleteAfterExport"`                                     // Delete the source after exporting to CAR files
	TokenUrl                   string `json:"tokenUrl"`                                              // Token server url.
	RootPrefix                 string `json:"rootPrefix" default:"/"`                                // The root/parent folder for all paths.
	UploadConcurrency          string `json:"uploadConcurrency" default:"4"`                         // Concurrency for chunked uploads.
	Encoding                   string `json:"encoding" default:"Slash,Dot"`                          // The encoding for the backend.
	Token                      string `json:"token"`                                                 // OAuth Access Token as a JSON blob.
	ScopeRole                  string `json:"scopeRole" default:"user"`                              // User-level that rclone should use when requesting access from HiDrive.
	Endpoint                   string `json:"endpoint" default:"https://api.hidrive.strato.com/2.1"` // Endpoint for the service.
	UploadCutoff               string `json:"uploadCutoff" default:"96Mi"`                           // Cutoff/Threshold for chunked uploads.
	ClientId                   string `json:"clientId"`                                              // OAuth Client Id.
	ClientSecret               string `json:"clientSecret"`                                          // OAuth Client Secret.
	DisableFetchingMemberCount string `json:"disableFetchingMemberCount" default:"false"`            // Do not fetch number of objects in directories unless it is absolutely necessary.
	ChunkSize                  string `json:"chunkSize" default:"48Mi"`                              // Chunksize for chunked uploads.
	AuthUrl                    string `json:"authUrl"`                                               // Auth server URL.
	ScopeAccess                string `json:"scopeAccess" default:"rw"`                              // Access permissions that rclone should use when requesting access from HiDrive.
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
	SourcePath        string `json:"sourcePath"`              // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`       // Delete the source after exporting to CAR files
	Headers           string `json:"headers"`                 // Set HTTP headers for all transactions.
	NoSlash           string `json:"noSlash" default:"false"` // Set this if the site doesn't end directories with /.
	NoHead            string `json:"noHead" default:"false"`  // Don't use HEAD requests.
	Url               string `json:"url"`                     // URL of HTTP host to connect to.
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
	SourcePath        string `json:"sourcePath"`                                                 // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                          // Delete the source after exporting to CAR files
	DisableChecksum   string `json:"disableChecksum" default:"true"`                             // Don't ask the server to test against MD5 checksum calculated by rclone.
	WaitArchive       string `json:"waitArchive" default:"0s"`                                   // Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
	Encoding          string `json:"encoding" default:"Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	AccessKeyId       string `json:"accessKeyId"`                                                // IAS3 Access Key.
	SecretAccessKey   string `json:"secretAccessKey"`                                            // IAS3 Secret Key (password).
	Endpoint          string `json:"endpoint" default:"https://s3.us.archive.org"`               // IAS3 Endpoint.
	FrontEndpoint     string `json:"frontEndpoint" default:"https://archive.org"`                // Host of InternetArchive Frontend.
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
	SourcePath        string `json:"sourcePath"`                                                                                     // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                                                              // Delete the source after exporting to CAR files
	Md5MemoryLimit    string `json:"md5MemoryLimit" default:"10Mi"`                                                                  // Files bigger than this will be cached on disk to calculate the MD5 if required.
	TrashedOnly       string `json:"trashedOnly" default:"false"`                                                                    // Only show files that are in the trash.
	HardDelete        string `json:"hardDelete" default:"false"`                                                                     // Delete files permanently rather than putting them into the trash.
	UploadResumeLimit string `json:"uploadResumeLimit" default:"10Mi"`                                                               // Files bigger than this can be resumed if the upload fail's.
	NoVersions        string `json:"noVersions" default:"false"`                                                                     // Avoid server side versioning by deleting files and recreating files instead of overwriting them.
	Encoding          string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
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
	SourcePath        string `json:"sourcePath"`                                                 // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                          // Delete the source after exporting to CAR files
	Endpoint          string `json:"endpoint"`                                                   // The Koofr API endpoint to use.
	Mountid           string `json:"mountid"`                                                    // Mount ID of the mount to use.
	Setmtime          string `json:"setmtime" default:"true"`                                    // Does the backend support setting modification time.
	User              string `json:"user"`                                                       // Your user name.
	Password          string `json:"password"`                                                   // Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password).
	Encoding          string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	Provider          string `json:"provider"`                                                   // Choose your storage provider.
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
	SourcePath           string `json:"sourcePath"`                           // The path of the source to scan items
	DeleteAfterExport    string `json:"deleteAfterExport"`                    // Delete the source after exporting to CAR files
	OneFileSystem        string `json:"oneFileSystem" default:"false"`        // Don't cross filesystem boundaries (unix/macOS only).
	CaseInsensitive      string `json:"caseInsensitive" default:"false"`      // Force the filesystem to report itself as case insensitive.
	CopyLinks            string `json:"copyLinks" default:"false"`            // Follow symlinks and copy the pointed to item.
	Encoding             string `json:"encoding" default:"Slash,Dot"`         // The encoding for the backend.
	Nounc                string `json:"nounc" default:"false"`                // Disable UNC (long path names) conversion on Windows.
	Links                string `json:"links" default:"false"`                // Translate symlinks to/from regular files with a '.rclonelink' extension.
	SkipLinks            string `json:"skipLinks" default:"false"`            // Don't warn about skipped symlinks.
	UnicodeNormalization string `json:"unicodeNormalization" default:"false"` // Apply unicode NFC normalization to paths and filenames.
	NoCheckUpdated       string `json:"noCheckUpdated" default:"false"`       // Don't check to see if the files change during upload.
	NoSparse             string `json:"noSparse" default:"false"`             // Disable sparse files for multi-thread downloads.
	NoSetModtime         string `json:"noSetModtime" default:"false"`         // Disable setting modtime.
	ZeroSizeLinks        string `json:"zeroSizeLinks" default:"false"`        // Assume the Stat size of links is zero (and read them instead) (deprecated).
	CaseSensitive        string `json:"caseSensitive" default:"false"`        // Force the filesystem to report itself as case sensitive.
	NoPreallocate        string `json:"noPreallocate" default:"false"`        // Disable preallocation of disk space for transferred files.
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
	SourcePath          string `json:"sourcePath"`                                                                                               // The path of the source to scan items
	DeleteAfterExport   string `json:"deleteAfterExport"`                                                                                        // Delete the source after exporting to CAR files
	SpeedupMaxDisk      string `json:"speedupMaxDisk" default:"3Gi"`                                                                             // This option allows you to disable speedup (put by hash) for large files.
	CheckHash           string `json:"checkHash" default:"true"`                                                                                 // What should copy do if file checksum is mismatched or invalid.
	SpeedupEnable       string `json:"speedupEnable" default:"true"`                                                                             // Skip full upload if there is another file with same data hash.
	Pass                string `json:"pass"`                                                                                                     // Password.
	SpeedupFilePatterns string `json:"speedupFilePatterns" default:"*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf"`                             // Comma separated list of file name patterns eligible for speedup (put by hash).
	SpeedupMaxMemory    string `json:"speedupMaxMemory" default:"32Mi"`                                                                          // Files larger than the size given below will always be hashed on disk.
	UserAgent           string `json:"userAgent"`                                                                                                // HTTP user agent used internally by client.
	Quirks              string `json:"quirks"`                                                                                                   // Comma separated list of internal maintenance flags.
	Encoding            string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	User                string `json:"user"`                                                                                                     // User name (usually email).
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
	SourcePath        string `json:"sourcePath"`                               // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                        // Delete the source after exporting to CAR files
	User              string `json:"user"`                                     // User name.
	Pass              string `json:"pass"`                                     // Password.
	Debug             string `json:"debug" default:"false"`                    // Output more debug from Mega.
	HardDelete        string `json:"hardDelete" default:"false"`               // Delete files permanently rather than putting them into the trash.
	UseHttps          string `json:"useHttps" default:"false"`                 // Use HTTPS for transfers.
	Encoding          string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
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
	SourcePath        string `json:"sourcePath"`        // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"` // Delete the source after exporting to CAR files
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
	SourcePath        string `json:"sourcePath"`               // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`        // Delete the source after exporting to CAR files
	Protocol          string `json:"protocol" default:"https"` // Select between HTTP or HTTPS protocol.
	Host              string `json:"host"`                     // Domain+path of NetStorage host to connect to.
	Account           string `json:"account"`                  // Set the NetStorage account name
	Secret            string `json:"secret"`                   // Set the NetStorage account secret/G2O key for authentication.
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
	SourcePath              string `json:"sourcePath"`                                                                                                                                          // The path of the source to scan items
	DeleteAfterExport       string `json:"deleteAfterExport"`                                                                                                                                   // Delete the source after exporting to CAR files
	ChunkSize               string `json:"chunkSize" default:"10Mi"`                                                                                                                            // Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
	DriveType               string `json:"driveType"`                                                                                                                                           // The type of the drive (personal | business | documentLibrary).
	RootFolderId            string `json:"rootFolderId"`                                                                                                                                        // ID of the root folder.
	ServerSideAcrossConfigs string `json:"serverSideAcrossConfigs" default:"false"`                                                                                                             // Allow server-side operations (e.g. copy) to work across different onedrive configs.
	NoVersions              string `json:"noVersions" default:"false"`                                                                                                                          // Remove all versions on modifying operations.
	Token                   string `json:"token"`                                                                                                                                               // OAuth Access Token as a JSON blob.
	AuthUrl                 string `json:"authUrl"`                                                                                                                                             // Auth server URL.
	AccessScopes            string `json:"accessScopes" default:"Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"`                                  // Set scopes to be requested by rclone.
	ExposeOnenoteFiles      string `json:"exposeOnenoteFiles" default:"false"`                                                                                                                  // Set to make OneNote files show up in directory listings.
	LinkScope               string `json:"linkScope" default:"anonymous"`                                                                                                                       // Set the scope of the links created by the link command.
	ClientSecret            string `json:"clientSecret"`                                                                                                                                        // OAuth Client Secret.
	DisableSitePermission   string `json:"disableSitePermission" default:"false"`                                                                                                               // Disable the request for Sites.Read.All permission.
	LinkType                string `json:"linkType" default:"view"`                                                                                                                             // Set the type of the links created by the link command.
	ClientId                string `json:"clientId"`                                                                                                                                            // OAuth Client Id.
	Region                  string `json:"region" default:"global"`                                                                                                                             // Choose national cloud region for OneDrive.
	DriveId                 string `json:"driveId"`                                                                                                                                             // The ID of the drive to use.
	ListChunk               string `json:"listChunk" default:"1000"`                                                                                                                            // Size of listing chunk.
	LinkPassword            string `json:"linkPassword"`                                                                                                                                        // Set the password for links created by the link command.
	HashType                string `json:"hashType" default:"auto"`                                                                                                                             // Specify the hash in use for the backend.
	Encoding                string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
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
	SourcePath        string `json:"sourcePath"`                                                                                                                                       // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                                                                                                                // Delete the source after exporting to CAR files
	Username          string `json:"username"`                                                                                                                                         // Username.
	Password          string `json:"password"`                                                                                                                                         // Password.
	Encoding          string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot"` // The encoding for the backend.
	ChunkSize         string `json:"chunkSize" default:"10Mi"`                                                                                                                         // Files will be uploaded in chunks this size.
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
	SourcePath           string `json:"sourcePath"`                               // The path of the source to scan items
	DeleteAfterExport    string `json:"deleteAfterExport"`                        // Delete the source after exporting to CAR files
	DisableChecksum      string `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	Encoding             string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	NoCheckBucket        string `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	SseCustomerKey       string `json:"sseCustomerKey"`                           // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKeySha256 string `json:"sseCustomerKeySha256"`                     // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	UploadConcurrency    string `json:"uploadConcurrency" default:"10"`           // Concurrency for multipart uploads.
	Compartment          string `json:"compartment"`                              // Object storage compartment OCID
	Region               string `json:"region"`                                   // Object storage Region
	UploadCutoff         string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize            string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	Namespace            string `json:"namespace"`                                // Object storage namespace
	ConfigFile           string `json:"configFile" default:"~/.oci/config"`       // Path to OCI config file
	ConfigProfile        string `json:"configProfile" default:"Default"`          // Profile name inside the oci config file
	StorageTier          string `json:"storageTier" default:"Standard"`           // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	CopyTimeout          string `json:"copyTimeout" default:"1m0s"`               // Timeout for copy.
	SseCustomerKeyFile   string `json:"sseCustomerKeyFile"`                       // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerAlgorithm string `json:"sseCustomerAlgorithm"`                     // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
	Provider             string `json:"provider" default:"env_auth"`              // Choose your Auth Provider
	CopyCutoff           string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	LeavePartsOnError    string `json:"leavePartsOnError" default:"false"`        // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	SseKmsKeyId          string `json:"sseKmsKeyId"`                              // if using using your own master key in vault, this header specifies the
	Endpoint             string `json:"endpoint"`                                 // Endpoint for Object storage API.
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
	SourcePath        string `json:"sourcePath"`                                                 // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                          // Delete the source after exporting to CAR files
	TokenUrl          string `json:"tokenUrl"`                                                   // Token server url.
	Hostname          string `json:"hostname" default:"api.pcloud.com"`                          // Hostname to connect to.
	Password          string `json:"password"`                                                   // Your pcloud password.
	ClientSecret      string `json:"clientSecret"`                                               // OAuth Client Secret.
	AuthUrl           string `json:"authUrl"`                                                    // Auth server URL.
	Encoding          string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	RootFolderId      string `json:"rootFolderId" default:"d0"`                                  // Fill in for rclone to use a non root folder as its starting point.
	Username          string `json:"username"`                                                   // Your pcloud username.
	ClientId          string `json:"clientId"`                                                   // OAuth Client Id.
	Token             string `json:"token"`                                                      // OAuth Access Token as a JSON blob.
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
	SourcePath        string `json:"sourcePath"`                                                             // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                                      // Delete the source after exporting to CAR files
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
	SourcePath        string `json:"sourcePath"`                                                 // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                          // Delete the source after exporting to CAR files
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
	SourcePath        string `json:"sourcePath"`                               // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                        // Delete the source after exporting to CAR files
	EnvAuth           string `json:"envAuth" default:"false"`                  // Get QingStor credentials from runtime.
	AccessKeyId       string `json:"accessKeyId"`                              // QingStor Access Key ID.
	ConnectionRetries string `json:"connectionRetries" default:"3"`            // Number of connection retries.
	Encoding          string `json:"encoding" default:"Slash,Ctl,InvalidUtf8"` // The encoding for the backend.
	UploadConcurrency string `json:"uploadConcurrency" default:"1"`            // Concurrency for multipart uploads.
	SecretAccessKey   string `json:"secretAccessKey"`                          // QingStor Secret Access Key (password).
	Endpoint          string `json:"endpoint"`                                 // Enter an endpoint URL to connection QingStor API.
	Zone              string `json:"zone"`                                     // Zone to connect to.
	UploadCutoff      string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize         string `json:"chunkSize" default:"4Mi"`                  // Chunk size to use for uploading.
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
	SourcePath            string `json:"sourcePath"`                               // The path of the source to scan items
	DeleteAfterExport     string `json:"deleteAfterExport"`                        // Delete the source after exporting to CAR files
	Provider              string `json:"provider"`                                 // Choose your S3 provider.
	StorageClass          string `json:"storageClass"`                             // The storage class to use when storing new objects in S3.
	MaxUploadParts        string `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	ListChunk             string `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	NoHead                string `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	DisableHttp2          string `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	SseCustomerAlgorithm  string `json:"sseCustomerAlgorithm"`                     // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	SseCustomerKeyBase64  string `json:"sseCustomerKeyBase64"`                     // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	SseCustomerKeyMd5     string `json:"sseCustomerKeyMd5"`                        // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	DisableChecksum       string `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	Decompress            string `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	ListVersion           string `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	EnvAuth               string `json:"envAuth" default:"false"`                  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	BucketAcl             string `json:"bucketAcl"`                                // Canned ACL used when creating buckets.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	V2Auth                string `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	UsePresignedRequest   string `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	LocationConstraint    string `json:"locationConstraint"`                       // Location constraint - must be set to match the Region.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         string `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHeadObject          string `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	MemoryPoolUseMmap     string `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Endpoint              string `json:"endpoint"`                                 // Endpoint for S3 API.
	Acl                   string `json:"acl"`                                      // Canned ACL used when creating buckets and storing or copying objects.
	UseAccelerateEndpoint string `json:"useAccelerateEndpoint" default:"false"`    // If true use the AWS S3 accelerated endpoint.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	ServerSideEncryption  string `json:"serverSideEncryption"`                     // The server-side encryption algorithm used when storing this object in S3.
	SseKmsKeyId           string `json:"sseKmsKeyId"`                              // If using KMS ID you must provide the ARN of Key.
	UploadConcurrency     string `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      string `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	StsEndpoint           string `json:"stsEndpoint"`                              // Endpoint for STS.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	Versions              string `json:"versions" default:"false"`                 // Include old versions in directory listings.
	Region                string `json:"region"`                                   // Region to connect to.
	RequesterPays         string `json:"requesterPays" default:"false"`            // Enables requester pays option when interacting with S3 bucket.
	SseCustomerKey        string `json:"sseCustomerKey"`                           // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	ForcePathStyle        string `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	LeavePartsOnError     string `json:"leavePartsOnError" default:"false"`        // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
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
	SourcePath        string `json:"sourcePath"`                                                     // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                              // Delete the source after exporting to CAR files
	CreateLibrary     string `json:"createLibrary" default:"false"`                                  // Should rclone create a library if it doesn't exist.
	AuthToken         string `json:"authToken"`                                                      // Authentication token.
	Url               string `json:"url"`                                                            // URL of seafile host to connect to.
	User              string `json:"user"`                                                           // User name (usually email address).
	Pass              string `json:"pass"`                                                           // Password.
	TwoFA             string `json:"2fa" default:"false"`                                            // Two-factor authentication ('true' if the account has 2FA enabled).
	Library           string `json:"library"`                                                        // Name of the library.
	LibraryKey        string `json:"libraryKey"`                                                     // Library password (for encrypted libraries only).
	Encoding          string `json:"encoding" default:"Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8"` // The encoding for the backend.
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
	SourcePath              string `json:"sourcePath"`                              // The path of the source to scan items
	DeleteAfterExport       string `json:"deleteAfterExport"`                       // Delete the source after exporting to CAR files
	Concurrency             string `json:"concurrency" default:"64"`                // The maximum number of outstanding requests for one file
	KnownHostsFile          string `json:"knownHostsFile"`                          // Optional path to known_hosts file.
	UseInsecureCipher       string `json:"useInsecureCipher" default:"false"`       // Enable the use of insecure ciphers and key exchange methods.
	ShellType               string `json:"shellType"`                               // The type of SSH shell on remote server, if any.
	Pass                    string `json:"pass"`                                    // SSH password, leave blank to use ssh-agent.
	ChunkSize               string `json:"chunkSize" default:"32Ki"`                // Upload and download chunk size.
	Ciphers                 string `json:"ciphers"`                                 // Space separated list of ciphers to be used for session encryption, ordered by preference.
	IdleTimeout             string `json:"idleTimeout" default:"1m0s"`              // Max time before closing idle connections.
	KeyPem                  string `json:"keyPem"`                                  // Raw PEM-encoded private key.
	SkipLinks               string `json:"skipLinks" default:"false"`               // Set to skip any symlinks and any other non regular files.
	DisableConcurrentWrites string `json:"disableConcurrentWrites" default:"false"` // If set don't use concurrent writes.
	Md5sumCommand           string `json:"md5sumCommand"`                           // The command used to read md5 hashes.
	Subsystem               string `json:"subsystem" default:"sftp"`                // Specifies the SSH2 subsystem on the remote host.
	Host                    string `json:"host"`                                    // SSH host to connect to.
	AskPassword             string `json:"askPassword" default:"false"`             // Allow asking for SFTP password when needed.
	SetModtime              string `json:"setModtime" default:"true"`               // Set the modified time on the remote if set.
	KeyFile                 string `json:"keyFile"`                                 // Path to PEM-encoded private key file.
	KeyExchange             string `json:"keyExchange"`                             // Space separated list of key exchange algorithms, ordered by preference.
	SetEnv                  string `json:"setEnv"`                                  // Environment variables to pass to sftp and commands
	PubkeyFile              string `json:"pubkeyFile"`                              // Optional path to public key file.
	ServerCommand           string `json:"serverCommand"`                           // Specifies the path or command to run a sftp server on the remote host.
	DisableConcurrentReads  string `json:"disableConcurrentReads" default:"false"`  // If set don't use concurrent reads.
	UseFstat                string `json:"useFstat" default:"false"`                // If set use fstat instead of stat.
	Macs                    string `json:"macs"`                                    // Space separated list of MACs (message authentication code) algorithms, ordered by preference.
	User                    string `json:"user" default:"shane"`                    // SSH username.
	KeyUseAgent             string `json:"keyUseAgent" default:"false"`             // When set forces the usage of the ssh-agent.
	DisableHashcheck        string `json:"disableHashcheck" default:"false"`        // Disable the execution of SSH commands to determine if remote file hashing is available.
	Sha1sumCommand          string `json:"sha1sumCommand"`                          // The command used to read sha1 hashes.
	Port                    string `json:"port" default:"22"`                       // SSH port number.
	KeyFilePass             string `json:"keyFilePass"`                             // The passphrase to decrypt the PEM-encoded private key file.
	PathOverride            string `json:"pathOverride"`                            // Override path used by SSH shell commands.
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
	SourcePath        string `json:"sourcePath"`                                                                                                                                       // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                                                                                                                // Delete the source after exporting to CAR files
	UploadCutoff      string `json:"uploadCutoff" default:"128Mi"`                                                                                                                     // Cutoff for switching to multipart upload.
	RootFolderId      string `json:"rootFolderId"`                                                                                                                                     // ID of the root folder.
	ChunkSize         string `json:"chunkSize" default:"64Mi"`                                                                                                                         // Upload chunk size.
	Endpoint          string `json:"endpoint"`                                                                                                                                         // Endpoint for API calls.
	Encoding          string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
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
	SourcePath        string `json:"sourcePath"`                                                             // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                                      // Delete the source after exporting to CAR files
	ApiUrl            string `json:"apiUrl" default:"http://127.0.0.1:9980"`                                 // Sia daemon API URL, like http://sia.daemon.host:9980.
	ApiPassword       string `json:"apiPassword"`                                                            // Sia Daemon API Password.
	UserAgent         string `json:"userAgent" default:"Sia-Agent"`                                          // Siad User Agent
	Encoding          string `json:"encoding" default:"Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
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
	SourcePath        string `json:"sourcePath"`                                                                                                                  // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                                                                                           // Delete the source after exporting to CAR files
	Spn               string `json:"spn"`                                                                                                                         // Service principal name.
	IdleTimeout       string `json:"idleTimeout" default:"1m0s"`                                                                                                  // Max time before closing idle connections.
	HideSpecialShare  string `json:"hideSpecialShare" default:"true"`                                                                                             // Hide special shares (e.g. print$) which users aren't supposed to access.
	CaseInsensitive   string `json:"caseInsensitive" default:"true"`                                                                                              // Whether the server is configured to be case-insensitive.
	Encoding          string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
	Host              string `json:"host"`                                                                                                                        // SMB server hostname to connect to.
	Port              string `json:"port" default:"445"`                                                                                                          // SMB port number.
	Pass              string `json:"pass"`                                                                                                                        // SMB password.
	Domain            string `json:"domain" default:"WORKGROUP"`                                                                                                  // Domain name for NTLM authentication.
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
	SourcePath        string `json:"sourcePath"`                              // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                       // Delete the source after exporting to CAR files
	AccessGrant       string `json:"accessGrant"`                             // Access grant.
	SatelliteAddress  string `json:"satelliteAddress" default:"us1.storj.io"` // Satellite address.
	ApiKey            string `json:"apiKey"`                                  // API key.
	Passphrase        string `json:"passphrase"`                              // Encryption passphrase.
	Provider          string `json:"provider" default:"existing"`             // Choose an authentication method.
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
	SourcePath        string `json:"sourcePath"`                              // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                       // Delete the source after exporting to CAR files
	AccessGrant       string `json:"accessGrant"`                             // Access grant.
	SatelliteAddress  string `json:"satelliteAddress" default:"us1.storj.io"` // Satellite address.
	ApiKey            string `json:"apiKey"`                                  // API key.
	Passphrase        string `json:"passphrase"`                              // Encryption passphrase.
	Provider          string `json:"provider" default:"existing"`             // Choose an authentication method.
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
	SourcePath          string `json:"sourcePath"`                                   // The path of the source to scan items
	DeleteAfterExport   string `json:"deleteAfterExport"`                            // Delete the source after exporting to CAR files
	AuthorizationExpiry string `json:"authorizationExpiry"`                          // Sugarsync authorization expiry.
	DeletedId           string `json:"deletedId"`                                    // Sugarsync deleted folder id.
	Encoding            string `json:"encoding" default:"Slash,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	AppId               string `json:"appId"`                                        // Sugarsync App ID.
	AccessKeyId         string `json:"accessKeyId"`                                  // Sugarsync Access Key ID.
	PrivateAccessKey    string `json:"privateAccessKey"`                             // Sugarsync Private Access Key.
	Authorization       string `json:"authorization"`                                // Sugarsync authorization.
	HardDelete          string `json:"hardDelete" default:"false"`                   // Permanently delete files if true
	RefreshToken        string `json:"refreshToken"`                                 // Sugarsync refresh token.
	User                string `json:"user"`                                         // Sugarsync user.
	RootId              string `json:"rootId"`                                       // Sugarsync root id.
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
	SourcePath                  string `json:"sourcePath"`                           // The path of the source to scan items
	DeleteAfterExport           string `json:"deleteAfterExport"`                    // Delete the source after exporting to CAR files
	TenantDomain                string `json:"tenantDomain"`                         // Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME).
	Region                      string `json:"region"`                               // Region name - optional (OS_REGION_NAME).
	AuthVersion                 string `json:"authVersion" default:"0"`              // AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).
	Encoding                    string `json:"encoding" default:"Slash,InvalidUtf8"` // The encoding for the backend.
	Domain                      string `json:"domain"`                               // User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME)
	Tenant                      string `json:"tenant"`                               // Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME).
	TenantId                    string `json:"tenantId"`                             // Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID).
	StorageUrl                  string `json:"storageUrl"`                           // Storage URL - optional (OS_STORAGE_URL).
	AuthToken                   string `json:"authToken"`                            // Auth Token from alternate authentication - optional (OS_AUTH_TOKEN).
	ApplicationCredentialId     string `json:"applicationCredentialId"`              // Application Credential ID (OS_APPLICATION_CREDENTIAL_ID).
	LeavePartsOnError           string `json:"leavePartsOnError" default:"false"`    // If true avoid calling abort upload on a failure.
	StoragePolicy               string `json:"storagePolicy"`                        // The storage policy to use when creating a new container.
	UserId                      string `json:"userId"`                               // User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID).
	NoChunk                     string `json:"noChunk" default:"false"`              // Don't chunk files during streaming upload.
	Auth                        string `json:"auth"`                                 // Authentication URL for server (OS_AUTH_URL).
	EndpointType                string `json:"endpointType" default:"public"`        // Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).
	NoLargeObjects              string `json:"noLargeObjects" default:"false"`       // Disable support for static and dynamic large objects
	User                        string `json:"user"`                                 // User name to log in (OS_USERNAME).
	Key                         string `json:"key"`                                  // API key or password (OS_PASSWORD).
	ApplicationCredentialName   string `json:"applicationCredentialName"`            // Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME).
	ApplicationCredentialSecret string `json:"applicationCredentialSecret"`          // Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET).
	ChunkSize                   string `json:"chunkSize" default:"5Gi"`              // Above this size files will be chunked into a _segments container.
	EnvAuth                     string `json:"envAuth" default:"false"`              // Get swift credentials from environment variables in standard OpenStack form.
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
	SourcePath        string `json:"sourcePath"`                                                                            // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                                                     // Delete the source after exporting to CAR files
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
	SourcePath         string `json:"sourcePath"`         // The path of the source to scan items
	DeleteAfterExport  string `json:"deleteAfterExport"`  // Delete the source after exporting to CAR files
	Encoding           string `json:"encoding"`           // The encoding for the backend.
	Headers            string `json:"headers"`            // Set HTTP headers for all transactions.
	Url                string `json:"url"`                // URL of http host to connect to.
	Vendor             string `json:"vendor"`             // Name of the WebDAV site/service/software you are using.
	User               string `json:"user"`               // User name.
	Pass               string `json:"pass"`               // Password.
	BearerToken        string `json:"bearerToken"`        // Bearer token instead of user/pass (e.g. a Macaroon).
	BearerTokenCommand string `json:"bearerTokenCommand"` // Command to run to get a bearer token.
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
	SourcePath        string `json:"sourcePath"`                                       // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                                // Delete the source after exporting to CAR files
	ClientId          string `json:"clientId"`                                         // OAuth Client Id.
	ClientSecret      string `json:"clientSecret"`                                     // OAuth Client Secret.
	Token             string `json:"token"`                                            // OAuth Access Token as a JSON blob.
	AuthUrl           string `json:"authUrl"`                                          // Auth server URL.
	TokenUrl          string `json:"tokenUrl"`                                         // Token server url.
	HardDelete        string `json:"hardDelete" default:"false"`                       // Delete files permanently rather than putting them into the trash.
	Encoding          string `json:"encoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
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
	SourcePath        string `json:"sourcePath"`                             // The path of the source to scan items
	DeleteAfterExport string `json:"deleteAfterExport"`                      // Delete the source after exporting to CAR files
	TokenUrl          string `json:"tokenUrl"`                               // Token server url.
	Region            string `json:"region"`                                 // Zoho region to connect to.
	Encoding          string `json:"encoding" default:"Del,Ctl,InvalidUtf8"` // The encoding for the backend.
	ClientId          string `json:"clientId"`                               // OAuth Client Id.
	ClientSecret      string `json:"clientSecret"`                           // OAuth Client Secret.
	Token             string `json:"token"`                                  // OAuth Access Token as a JSON blob.
	AuthUrl           string `json:"authUrl"`                                // Auth server URL.
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
	AcdClientId                         string `json:"acdClientId"`                                                                                                                                                 // OAuth Client Id.
	AcdCheckpoint                       string `json:"acdCheckpoint"`                                                                                                                                               // Checkpoint for internal polling (debug).
	AcdUploadWaitPerGb                  string `json:"acdUploadWaitPerGb" default:"3m0s"`                                                                                                                           // Additional time per GiB to wait after a failed complete upload to see if it appears.
	AcdClientSecret                     string `json:"acdClientSecret"`                                                                                                                                             // OAuth Client Secret.
	AcdToken                            string `json:"acdToken"`                                                                                                                                                    // OAuth Access Token as a JSON blob.
	AcdAuthUrl                          string `json:"acdAuthUrl"`                                                                                                                                                  // Auth server URL.
	AcdTokenUrl                         string `json:"acdTokenUrl"`                                                                                                                                                 // Token server url.
	AcdTemplinkThreshold                string `json:"acdTemplinkThreshold" default:"9Gi"`                                                                                                                          // Files >= this size will be downloaded via their tempLink.
	AcdEncoding                         string `json:"acdEncoding" default:"Slash,InvalidUtf8,Dot"`                                                                                                                 // The encoding for the backend.
	AzureblobListChunk                  string `json:"azureblobListChunk" default:"5000"`                                                                                                                           // Size of blob list.
	AzureblobPublicAccess               string `json:"azureblobPublicAccess"`                                                                                                                                       // Public access level of a container: blob or container.
	AzureblobNoHeadObject               string `json:"azureblobNoHeadObject" default:"false"`                                                                                                                       // If set, do not do HEAD before GET when getting objects.
	AzureblobClientCertificatePassword  string `json:"azureblobClientCertificatePassword"`                                                                                                                          // Password for the certificate file (optional).
	AzureblobClientSendCertificateChain string `json:"azureblobClientSendCertificateChain" default:"false"`                                                                                                         // Send the certificate chain when using certificate auth.
	AzureblobUseMsi                     string `json:"azureblobUseMsi" default:"false"`                                                                                                                             // Use a managed service identity to authenticate (only works in Azure).
	AzureblobMsiMiResId                 string `json:"azureblobMsiMiResId"`                                                                                                                                         // Azure resource ID of the user-assigned MSI to use, if any.
	AzureblobClientSecret               string `json:"azureblobClientSecret"`                                                                                                                                       // One of the service principal's client secrets
	AzureblobPassword                   string `json:"azureblobPassword"`                                                                                                                                           // The user's password
	AzureblobKey                        string `json:"azureblobKey"`                                                                                                                                                // Storage Account Shared Key.
	AzureblobArchiveTierDelete          string `json:"azureblobArchiveTierDelete" default:"false"`                                                                                                                  // Delete archive tier blobs before overwriting.
	AzureblobMemoryPoolFlushTime        string `json:"azureblobMemoryPoolFlushTime" default:"1m0s"`                                                                                                                 // How often internal memory buffer pools will be flushed.
	AzureblobUploadCutoff               string `json:"azureblobUploadCutoff"`                                                                                                                                       // Cutoff for switching to chunked upload (<= 256 MiB) (deprecated).
	AzureblobChunkSize                  string `json:"azureblobChunkSize" default:"4Mi"`                                                                                                                            // Upload chunk size.
	AzureblobTenant                     string `json:"azureblobTenant"`                                                                                                                                             // ID of the service principal's tenant. Also called its directory ID.
	AzureblobClientId                   string `json:"azureblobClientId"`                                                                                                                                           // The ID of the client in use.
	AzureblobClientCertificatePath      string `json:"azureblobClientCertificatePath"`                                                                                                                              // Path to a PEM or PKCS12 certificate file including the private key.
	AzureblobEncoding                   string `json:"azureblobEncoding" default:"Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"`                                                                                 // The encoding for the backend.
	AzureblobAccount                    string `json:"azureblobAccount"`                                                                                                                                            // Azure Storage Account Name.
	AzureblobUsername                   string `json:"azureblobUsername"`                                                                                                                                           // User name (usually an email address)
	AzureblobEndpoint                   string `json:"azureblobEndpoint"`                                                                                                                                           // Endpoint for the service.
	AzureblobNoCheckContainer           string `json:"azureblobNoCheckContainer" default:"false"`                                                                                                                   // If set, don't attempt to check the container exists or create it.
	AzureblobEnvAuth                    string `json:"azureblobEnvAuth" default:"false"`                                                                                                                            // Read credentials from runtime (environment variables, CLI or MSI).
	AzureblobServicePrincipalFile       string `json:"azureblobServicePrincipalFile"`                                                                                                                               // Path to file containing credentials for use with a service principal.
	AzureblobUploadConcurrency          string `json:"azureblobUploadConcurrency" default:"16"`                                                                                                                     // Concurrency for multipart uploads.
	AzureblobAccessTier                 string `json:"azureblobAccessTier"`                                                                                                                                         // Access tier of blob: hot, cool or archive.
	AzureblobDisableChecksum            string `json:"azureblobDisableChecksum" default:"false"`                                                                                                                    // Don't store MD5 checksum with object metadata.
	AzureblobMemoryPoolUseMmap          string `json:"azureblobMemoryPoolUseMmap" default:"false"`                                                                                                                  // Whether to use mmap buffers in internal memory pool.
	AzureblobSasUrl                     string `json:"azureblobSasUrl"`                                                                                                                                             // SAS URL for container level access only.
	AzureblobMsiObjectId                string `json:"azureblobMsiObjectId"`                                                                                                                                        // Object ID of the user-assigned MSI to use, if any.
	AzureblobMsiClientId                string `json:"azureblobMsiClientId"`                                                                                                                                        // Object ID of the user-assigned MSI to use, if any.
	AzureblobUseEmulator                string `json:"azureblobUseEmulator" default:"false"`                                                                                                                        // Uses local storage emulator if provided as 'true'.
	B2VersionAt                         string `json:"b2VersionAt" default:"off"`                                                                                                                                   // Show file versions as they were at the specified time.
	B2UploadCutoff                      string `json:"b2UploadCutoff" default:"200Mi"`                                                                                                                              // Cutoff for switching to chunked upload.
	B2CopyCutoff                        string `json:"b2CopyCutoff" default:"4Gi"`                                                                                                                                  // Cutoff for switching to multipart copy.
	B2ChunkSize                         string `json:"b2ChunkSize" default:"96Mi"`                                                                                                                                  // Upload chunk size.
	B2Account                           string `json:"b2Account"`                                                                                                                                                   // Account ID or Application Key ID.
	B2Key                               string `json:"b2Key"`                                                                                                                                                       // Application Key.
	B2HardDelete                        string `json:"b2HardDelete" default:"false"`                                                                                                                                // Permanently delete files on remote removal, otherwise hide files.
	B2Endpoint                          string `json:"b2Endpoint"`                                                                                                                                                  // Endpoint for the service.
	B2Versions                          string `json:"b2Versions" default:"false"`                                                                                                                                  // Include old versions in directory listings.
	B2DisableChecksum                   string `json:"b2DisableChecksum" default:"false"`                                                                                                                           // Disable checksums for large (> upload cutoff) files.
	B2DownloadAuthDuration              string `json:"b2DownloadAuthDuration" default:"1w"`                                                                                                                         // Time before the authorization token will expire in s or suffix ms|s|m|h|d.
	B2MemoryPoolFlushTime               string `json:"b2MemoryPoolFlushTime" default:"1m0s"`                                                                                                                        // How often internal memory buffer pools will be flushed.
	B2Encoding                          string `json:"b2Encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"`                                                                                                // The encoding for the backend.
	B2TestMode                          string `json:"b2TestMode"`                                                                                                                                                  // A flag string for X-Bz-Test-Mode header for debugging.
	B2DownloadUrl                       string `json:"b2DownloadUrl"`                                                                                                                                               // Custom endpoint for downloads.
	B2MemoryPoolUseMmap                 string `json:"b2MemoryPoolUseMmap" default:"false"`                                                                                                                         // Whether to use mmap buffers in internal memory pool.
	BoxClientId                         string `json:"boxClientId"`                                                                                                                                                 // OAuth Client Id.
	BoxTokenUrl                         string `json:"boxTokenUrl"`                                                                                                                                                 // Token server url.
	BoxRootFolderId                     string `json:"boxRootFolderId" default:"0"`                                                                                                                                 // Fill in for rclone to use a non root folder as its starting point.
	BoxListChunk                        string `json:"boxListChunk" default:"1000"`                                                                                                                                 // Size of listing chunk 1-1000.
	BoxEncoding                         string `json:"boxEncoding" default:"Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"`                                                                                    // The encoding for the backend.
	BoxClientSecret                     string `json:"boxClientSecret"`                                                                                                                                             // OAuth Client Secret.
	BoxBoxConfigFile                    string `json:"boxBoxConfigFile"`                                                                                                                                            // Box App config.json location
	BoxAccessToken                      string `json:"boxAccessToken"`                                                                                                                                              // Box App Primary Access Token
	BoxUploadCutoff                     string `json:"boxUploadCutoff" default:"50Mi"`                                                                                                                              // Cutoff for switching to multipart upload (>= 50 MiB).
	BoxOwnedBy                          string `json:"boxOwnedBy"`                                                                                                                                                  // Only show items owned by the login (email address) passed in.
	BoxToken                            string `json:"boxToken"`                                                                                                                                                    // OAuth Access Token as a JSON blob.
	BoxAuthUrl                          string `json:"boxAuthUrl"`                                                                                                                                                  // Auth server URL.
	BoxBoxSubType                       string `json:"boxBoxSubType" default:"user"`                                                                                                                                //
	BoxCommitRetries                    string `json:"boxCommitRetries" default:"100"`                                                                                                                              // Max number of times to try committing a multipart file.
	CryptFilenameEncryption             string `json:"cryptFilenameEncryption" default:"standard"`                                                                                                                  // How to encrypt the filenames.
	CryptPassword                       string `json:"cryptPassword"`                                                                                                                                               // Password or pass phrase for encryption.
	CryptNoDataEncryption               string `json:"cryptNoDataEncryption" default:"false"`                                                                                                                       // Option to either encrypt file data or leave it unencrypted.
	CryptFilenameEncoding               string `json:"cryptFilenameEncoding" default:"base32"`                                                                                                                      // How to encode the encrypted filename to text string.
	CryptShowMapping                    string `json:"cryptShowMapping" default:"false"`                                                                                                                            // For all files listed show how the names encrypt.
	CryptRemote                         string `json:"cryptRemote"`                                                                                                                                                 // Remote to encrypt/decrypt.
	CryptDirectoryNameEncryption        string `json:"cryptDirectoryNameEncryption" default:"true"`                                                                                                                 // Option to either encrypt directory names or leave them intact.
	CryptPassword2                      string `json:"cryptPassword2"`                                                                                                                                              // Password or pass phrase for salt.
	CryptServerSideAcrossConfigs        string `json:"cryptServerSideAcrossConfigs" default:"false"`                                                                                                                // Allow server-side operations (e.g. copy) to work across different crypt configs.
	DriveClientSecret                   string `json:"driveClientSecret"`                                                                                                                                           // OAuth Client Secret.
	DriveSkipChecksumGphotos            string `json:"driveSkipChecksumGphotos" default:"false"`                                                                                                                    // Skip MD5 checksum on Google photos and videos only.
	DriveFormats                        string `json:"driveFormats"`                                                                                                                                                // Deprecated: See export_formats.
	DriveDisableHttp2                   string `json:"driveDisableHttp2" default:"true"`                                                                                                                            // Disable drive using http2.
	DriveStopOnUploadLimit              string `json:"driveStopOnUploadLimit" default:"false"`                                                                                                                      // Make upload limit errors be fatal.
	DriveClientId                       string `json:"driveClientId"`                                                                                                                                               // Google Application Client Id
	DriveTeamDrive                      string `json:"driveTeamDrive"`                                                                                                                                              // ID of the Shared Drive (Team Drive).
	DriveAuthOwnerOnly                  string `json:"driveAuthOwnerOnly" default:"false"`                                                                                                                          // Only consider files owned by the authenticated user.
	DriveImpersonate                    string `json:"driveImpersonate"`                                                                                                                                            // Impersonate this user when using a service account.
	DriveAlternateExport                string `json:"driveAlternateExport" default:"false"`                                                                                                                        // Deprecated: No longer needed.
	DriveAcknowledgeAbuse               string `json:"driveAcknowledgeAbuse" default:"false"`                                                                                                                       // Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
	DrivePacerBurst                     string `json:"drivePacerBurst" default:"100"`                                                                                                                               // Number of API calls to allow without sleeping.
	DriveStopOnDownloadLimit            string `json:"driveStopOnDownloadLimit" default:"false"`                                                                                                                    // Make download limit errors be fatal.
	DriveToken                          string `json:"driveToken"`                                                                                                                                                  // OAuth Access Token as a JSON blob.
	DriveSkipGdocs                      string `json:"driveSkipGdocs" default:"false"`                                                                                                                              // Skip google documents in all listings.
	DriveExportFormats                  string `json:"driveExportFormats" default:"docx,xlsx,pptx,svg"`                                                                                                             // Comma separated list of preferred formats for downloading Google docs.
	DriveUseSharedDate                  string `json:"driveUseSharedDate" default:"false"`                                                                                                                          // Use date file was shared instead of modified date.
	DriveChunkSize                      string `json:"driveChunkSize" default:"8Mi"`                                                                                                                                // Upload chunk size.
	DriveV2DownloadMinSize              string `json:"driveV2DownloadMinSize" default:"off"`                                                                                                                        // If Object's are greater, use drive v2 API to download.
	DriveCopyShortcutContent            string `json:"driveCopyShortcutContent" default:"false"`                                                                                                                    // Server side copy contents of shortcuts instead of the shortcut.
	DriveSkipShortcuts                  string `json:"driveSkipShortcuts" default:"false"`                                                                                                                          // If set skip shortcut files.
	DriveImportFormats                  string `json:"driveImportFormats"`                                                                                                                                          // Comma separated list of preferred formats for uploading Google docs.
	DriveAllowImportNameChange          string `json:"driveAllowImportNameChange" default:"false"`                                                                                                                  // Allow the filetype to change when uploading Google docs.
	DrivePacerMinSleep                  string `json:"drivePacerMinSleep" default:"100ms"`                                                                                                                          // Minimum time to sleep between API calls.
	DriveResourceKey                    string `json:"driveResourceKey"`                                                                                                                                            // Resource key for accessing a link-shared file.
	DriveTrashedOnly                    string `json:"driveTrashedOnly" default:"false"`                                                                                                                            // Only show files that are in the trash.
	DriveSharedWithMe                   string `json:"driveSharedWithMe" default:"false"`                                                                                                                           // Only show files that are shared with me.
	DriveUploadCutoff                   string `json:"driveUploadCutoff" default:"8Mi"`                                                                                                                             // Cutoff for switching to chunked upload.
	DriveServerSideAcrossConfigs        string `json:"driveServerSideAcrossConfigs" default:"false"`                                                                                                                // Allow server-side operations (e.g. copy) to work across different drive configs.
	DriveScope                          string `json:"driveScope"`                                                                                                                                                  // Scope that rclone should use when requesting access from drive.
	DriveRootFolderId                   string `json:"driveRootFolderId"`                                                                                                                                           // ID of the root folder.
	DriveServiceAccountCredentials      string `json:"driveServiceAccountCredentials"`                                                                                                                              // Service Account Credentials JSON blob.
	DriveUseCreatedDate                 string `json:"driveUseCreatedDate" default:"false"`                                                                                                                         // Use file created date instead of modified date.
	DriveKeepRevisionForever            string `json:"driveKeepRevisionForever" default:"false"`                                                                                                                    // Keep new head revision of each file forever.
	DriveSizeAsQuota                    string `json:"driveSizeAsQuota" default:"false"`                                                                                                                            // Show sizes as storage quota usage, not actual size.
	DriveAuthUrl                        string `json:"driveAuthUrl"`                                                                                                                                                // Auth server URL.
	DriveServiceAccountFile             string `json:"driveServiceAccountFile"`                                                                                                                                     // Service Account Credentials JSON file path.
	DriveUseTrash                       string `json:"driveUseTrash" default:"true"`                                                                                                                                // Send files to the trash instead of deleting permanently.
	DriveStarredOnly                    string `json:"driveStarredOnly" default:"false"`                                                                                                                            // Only show files that are starred.
	DriveListChunk                      string `json:"driveListChunk" default:"1000"`                                                                                                                               // Size of listing chunk 100-1000, 0 to disable.
	DriveSkipDanglingShortcuts          string `json:"driveSkipDanglingShortcuts" default:"false"`                                                                                                                  // If set skip dangling shortcut files.
	DriveEncoding                       string `json:"driveEncoding" default:"InvalidUtf8"`                                                                                                                         // The encoding for the backend.
	DriveTokenUrl                       string `json:"driveTokenUrl"`                                                                                                                                               // Token server url.
	DropboxSharedFolders                string `json:"dropboxSharedFolders" default:"false"`                                                                                                                        // Instructs rclone to work on shared folders.
	DropboxBatchSize                    string `json:"dropboxBatchSize" default:"0"`                                                                                                                                // Max number of files in upload batch.
	DropboxBatchCommitTimeout           string `json:"dropboxBatchCommitTimeout" default:"10m0s"`                                                                                                                   // Max time to wait for a batch to finish committing
	DropboxSharedFiles                  string `json:"dropboxSharedFiles" default:"false"`                                                                                                                          // Instructs rclone to work on individual shared files.
	DropboxClientSecret                 string `json:"dropboxClientSecret"`                                                                                                                                         // OAuth Client Secret.
	DropboxAuthUrl                      string `json:"dropboxAuthUrl"`                                                                                                                                              // Auth server URL.
	DropboxTokenUrl                     string `json:"dropboxTokenUrl"`                                                                                                                                             // Token server url.
	DropboxImpersonate                  string `json:"dropboxImpersonate"`                                                                                                                                          // Impersonate this user when using a business account.
	DropboxClientId                     string `json:"dropboxClientId"`                                                                                                                                             // OAuth Client Id.
	DropboxChunkSize                    string `json:"dropboxChunkSize" default:"48Mi"`                                                                                                                             // Upload chunk size (< 150Mi).
	DropboxBatchMode                    string `json:"dropboxBatchMode" default:"sync"`                                                                                                                             // Upload file batching sync|async|off.
	DropboxEncoding                     string `json:"dropboxEncoding" default:"Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"`                                                                                    // The encoding for the backend.
	DropboxToken                        string `json:"dropboxToken"`                                                                                                                                                // OAuth Access Token as a JSON blob.
	DropboxBatchTimeout                 string `json:"dropboxBatchTimeout" default:"0s"`                                                                                                                            // Max time to allow an idle upload batch before uploading.
	FichierFolderPassword               string `json:"fichierFolderPassword"`                                                                                                                                       // If you want to list the files in a shared folder that is password protected, add this parameter.
	FichierEncoding                     string `json:"fichierEncoding" default:"Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"`                        // The encoding for the backend.
	FichierApiKey                       string `json:"fichierApiKey"`                                                                                                                                               // Your API Key, get it from https://1fichier.com/console/params.pl.
	FichierSharedFolder                 string `json:"fichierSharedFolder"`                                                                                                                                         // If you want to download a shared folder, add this parameter.
	FichierFilePassword                 string `json:"fichierFilePassword"`                                                                                                                                         // If you want to download a shared file that is password protected, add this parameter.
	FilefabricEncoding                  string `json:"filefabricEncoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"`                                                                                                  // The encoding for the backend.
	FilefabricUrl                       string `json:"filefabricUrl"`                                                                                                                                               // URL of the Enterprise File Fabric to connect to.
	FilefabricRootFolderId              string `json:"filefabricRootFolderId"`                                                                                                                                      // ID of the root folder.
	FilefabricPermanentToken            string `json:"filefabricPermanentToken"`                                                                                                                                    // Permanent Authentication Token.
	FilefabricToken                     string `json:"filefabricToken"`                                                                                                                                             // Session Token.
	FilefabricTokenExpiry               string `json:"filefabricTokenExpiry"`                                                                                                                                       // Token expiry time.
	FilefabricVersion                   string `json:"filefabricVersion"`                                                                                                                                           // Version read from the file fabric.
	FtpEncoding                         string `json:"ftpEncoding" default:"Slash,Del,Ctl,RightSpace,Dot"`                                                                                                          // The encoding for the backend.
	FtpForceListHidden                  string `json:"ftpForceListHidden" default:"false"`                                                                                                                          // Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.
	FtpShutTimeout                      string `json:"ftpShutTimeout" default:"1m0s"`                                                                                                                               // Maximum time to wait for data connection closing status.
	FtpDisableEpsv                      string `json:"ftpDisableEpsv" default:"false"`                                                                                                                              // Disable using EPSV even if server advertises support.
	FtpTlsCacheSize                     string `json:"ftpTlsCacheSize" default:"32"`                                                                                                                                // Size of TLS session cache for all control and data connections.
	FtpDisableTls13                     string `json:"ftpDisableTls13" default:"false"`                                                                                                                             // Disable TLS 1.3 (workaround for FTP servers with buggy TLS)
	FtpAskPassword                      string `json:"ftpAskPassword" default:"false"`                                                                                                                              // Allow asking for FTP password when needed.
	FtpHost                             string `json:"ftpHost"`                                                                                                                                                     // FTP host to connect to.
	FtpExplicitTls                      string `json:"ftpExplicitTls" default:"false"`                                                                                                                              // Use Explicit FTPS (FTP over TLS).
	FtpDisableMlsd                      string `json:"ftpDisableMlsd" default:"false"`                                                                                                                              // Disable using MLSD even if server advertises support.
	FtpIdleTimeout                      string `json:"ftpIdleTimeout" default:"1m0s"`                                                                                                                               // Max time before closing idle connections.
	FtpPass                             string `json:"ftpPass"`                                                                                                                                                     // FTP password.
	FtpTls                              string `json:"ftpTls" default:"false"`                                                                                                                                      // Use Implicit FTPS (FTP over TLS).
	FtpConcurrency                      string `json:"ftpConcurrency" default:"0"`                                                                                                                                  // Maximum number of FTP simultaneous connections, 0 for unlimited.
	FtpNoCheckCertificate               string `json:"ftpNoCheckCertificate" default:"false"`                                                                                                                       // Do not verify the TLS certificate of the server.
	FtpDisableUtf8                      string `json:"ftpDisableUtf8" default:"false"`                                                                                                                              // Disable using UTF-8 even if server advertises support.
	FtpWritingMdtm                      string `json:"ftpWritingMdtm" default:"false"`                                                                                                                              // Use MDTM to set modification time (VsFtpd quirk)
	FtpCloseTimeout                     string `json:"ftpCloseTimeout" default:"1m0s"`                                                                                                                              // Maximum time to wait for a response to close.
	FtpUser                             string `json:"ftpUser" default:"shane"`                                                                                                                                     // FTP username.
	FtpPort                             string `json:"ftpPort" default:"21"`                                                                                                                                        // FTP port number.
	GcsEnvAuth                          string `json:"gcsEnvAuth" default:"false"`                                                                                                                                  // Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
	GcsToken                            string `json:"gcsToken"`                                                                                                                                                    // OAuth Access Token as a JSON blob.
	GcsLocation                         string `json:"gcsLocation"`                                                                                                                                                 // Location for the newly created buckets.
	GcsStorageClass                     string `json:"gcsStorageClass"`                                                                                                                                             // The storage class to use when storing objects in Google Cloud Storage.
	GcsDecompress                       string `json:"gcsDecompress" default:"false"`                                                                                                                               // If set this will decompress gzip encoded objects.
	GcsEndpoint                         string `json:"gcsEndpoint"`                                                                                                                                                 // Endpoint for the service.
	GcsAuthUrl                          string `json:"gcsAuthUrl"`                                                                                                                                                  // Auth server URL.
	GcsTokenUrl                         string `json:"gcsTokenUrl"`                                                                                                                                                 // Token server url.
	GcsServiceAccountFile               string `json:"gcsServiceAccountFile"`                                                                                                                                       // Service Account Credentials JSON file path.
	GcsObjectAcl                        string `json:"gcsObjectAcl"`                                                                                                                                                // Access Control List for new objects.
	GcsEncoding                         string `json:"gcsEncoding" default:"Slash,CrLf,InvalidUtf8,Dot"`                                                                                                            // The encoding for the backend.
	GcsNoCheckBucket                    string `json:"gcsNoCheckBucket" default:"false"`                                                                                                                            // If set, don't attempt to check the bucket exists or create it.
	GcsClientId                         string `json:"gcsClientId"`                                                                                                                                                 // OAuth Client Id.
	GcsServiceAccountCredentials        string `json:"gcsServiceAccountCredentials"`                                                                                                                                // Service Account Credentials JSON blob.
	GcsAnonymous                        string `json:"gcsAnonymous" default:"false"`                                                                                                                                // Access public buckets and objects without credentials.
	GcsBucketAcl                        string `json:"gcsBucketAcl"`                                                                                                                                                // Access Control List for new buckets.
	GcsBucketPolicyOnly                 string `json:"gcsBucketPolicyOnly" default:"false"`                                                                                                                         // Access checks should use bucket-level IAM policies.
	GcsClientSecret                     string `json:"gcsClientSecret"`                                                                                                                                             // OAuth Client Secret.
	GcsProjectNumber                    string `json:"gcsProjectNumber"`                                                                                                                                            // Project number.
	GphotosClientSecret                 string `json:"gphotosClientSecret"`                                                                                                                                         // OAuth Client Secret.
	GphotosAuthUrl                      string `json:"gphotosAuthUrl"`                                                                                                                                              // Auth server URL.
	GphotosTokenUrl                     string `json:"gphotosTokenUrl"`                                                                                                                                             // Token server url.
	GphotosReadSize                     string `json:"gphotosReadSize" default:"false"`                                                                                                                             // Set to read the size of media items.
	GphotosEncoding                     string `json:"gphotosEncoding" default:"Slash,CrLf,InvalidUtf8,Dot"`                                                                                                        // The encoding for the backend.
	GphotosClientId                     string `json:"gphotosClientId"`                                                                                                                                             // OAuth Client Id.
	GphotosToken                        string `json:"gphotosToken"`                                                                                                                                                // OAuth Access Token as a JSON blob.
	GphotosReadOnly                     string `json:"gphotosReadOnly" default:"false"`                                                                                                                             // Set to make the Google Photos backend read only.
	GphotosStartYear                    string `json:"gphotosStartYear" default:"2000"`                                                                                                                             // Year limits the photos to be downloaded to those which are uploaded after the given year.
	GphotosIncludeArchived              string `json:"gphotosIncludeArchived" default:"false"`                                                                                                                      // Also view and download archived media.
	HdfsDataTransferProtection          string `json:"hdfsDataTransferProtection"`                                                                                                                                  // Kerberos data transfer protection: authentication|integrity|privacy.
	HdfsEncoding                        string `json:"hdfsEncoding" default:"Slash,Colon,Del,Ctl,InvalidUtf8,Dot"`                                                                                                  // The encoding for the backend.
	HdfsNamenode                        string `json:"hdfsNamenode"`                                                                                                                                                // Hadoop name node and port.
	HdfsUsername                        string `json:"hdfsUsername"`                                                                                                                                                // Hadoop user name.
	HdfsServicePrincipalName            string `json:"hdfsServicePrincipalName"`                                                                                                                                    // Kerberos service principal name for the namenode.
	HidriveTokenUrl                     string `json:"hidriveTokenUrl"`                                                                                                                                             // Token server url.
	HidriveRootPrefix                   string `json:"hidriveRootPrefix" default:"/"`                                                                                                                               // The root/parent folder for all paths.
	HidriveUploadConcurrency            string `json:"hidriveUploadConcurrency" default:"4"`                                                                                                                        // Concurrency for chunked uploads.
	HidriveEncoding                     string `json:"hidriveEncoding" default:"Slash,Dot"`                                                                                                                         // The encoding for the backend.
	HidriveToken                        string `json:"hidriveToken"`                                                                                                                                                // OAuth Access Token as a JSON blob.
	HidriveScopeRole                    string `json:"hidriveScopeRole" default:"user"`                                                                                                                             // User-level that rclone should use when requesting access from HiDrive.
	HidriveEndpoint                     string `json:"hidriveEndpoint" default:"https://api.hidrive.strato.com/2.1"`                                                                                                // Endpoint for the service.
	HidriveUploadCutoff                 string `json:"hidriveUploadCutoff" default:"96Mi"`                                                                                                                          // Cutoff/Threshold for chunked uploads.
	HidriveClientId                     string `json:"hidriveClientId"`                                                                                                                                             // OAuth Client Id.
	HidriveClientSecret                 string `json:"hidriveClientSecret"`                                                                                                                                         // OAuth Client Secret.
	HidriveDisableFetchingMemberCount   string `json:"hidriveDisableFetchingMemberCount" default:"false"`                                                                                                           // Do not fetch number of objects in directories unless it is absolutely necessary.
	HidriveChunkSize                    string `json:"hidriveChunkSize" default:"48Mi"`                                                                                                                             // Chunksize for chunked uploads.
	HidriveAuthUrl                      string `json:"hidriveAuthUrl"`                                                                                                                                              // Auth server URL.
	HidriveScopeAccess                  string `json:"hidriveScopeAccess" default:"rw"`                                                                                                                             // Access permissions that rclone should use when requesting access from HiDrive.
	HttpHeaders                         string `json:"httpHeaders"`                                                                                                                                                 // Set HTTP headers for all transactions.
	HttpNoSlash                         string `json:"httpNoSlash" default:"false"`                                                                                                                                 // Set this if the site doesn't end directories with /.
	HttpNoHead                          string `json:"httpNoHead" default:"false"`                                                                                                                                  // Don't use HEAD requests.
	HttpUrl                             string `json:"httpUrl"`                                                                                                                                                     // URL of HTTP host to connect to.
	InternetarchiveDisableChecksum      string `json:"internetarchiveDisableChecksum" default:"true"`                                                                                                               // Don't ask the server to test against MD5 checksum calculated by rclone.
	InternetarchiveWaitArchive          string `json:"internetarchiveWaitArchive" default:"0s"`                                                                                                                     // Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
	InternetarchiveEncoding             string `json:"internetarchiveEncoding" default:"Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot"`                                                                                   // The encoding for the backend.
	InternetarchiveAccessKeyId          string `json:"internetarchiveAccessKeyId"`                                                                                                                                  // IAS3 Access Key.
	InternetarchiveSecretAccessKey      string `json:"internetarchiveSecretAccessKey"`                                                                                                                              // IAS3 Secret Key (password).
	InternetarchiveEndpoint             string `json:"internetarchiveEndpoint" default:"https://s3.us.archive.org"`                                                                                                 // IAS3 Endpoint.
	InternetarchiveFrontEndpoint        string `json:"internetarchiveFrontEndpoint" default:"https://archive.org"`                                                                                                  // Host of InternetArchive Frontend.
	JottacloudMd5MemoryLimit            string `json:"jottacloudMd5MemoryLimit" default:"10Mi"`                                                                                                                     // Files bigger than this will be cached on disk to calculate the MD5 if required.
	JottacloudTrashedOnly               string `json:"jottacloudTrashedOnly" default:"false"`                                                                                                                       // Only show files that are in the trash.
	JottacloudHardDelete                string `json:"jottacloudHardDelete" default:"false"`                                                                                                                        // Delete files permanently rather than putting them into the trash.
	JottacloudUploadResumeLimit         string `json:"jottacloudUploadResumeLimit" default:"10Mi"`                                                                                                                  // Files bigger than this can be resumed if the upload fail's.
	JottacloudNoVersions                string `json:"jottacloudNoVersions" default:"false"`                                                                                                                        // Avoid server side versioning by deleting files and recreating files instead of overwriting them.
	JottacloudEncoding                  string `json:"jottacloudEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot"`                                                    // The encoding for the backend.
	KoofrEndpoint                       string `json:"koofrEndpoint"`                                                                                                                                               // The Koofr API endpoint to use.
	KoofrMountid                        string `json:"koofrMountid"`                                                                                                                                                // Mount ID of the mount to use.
	KoofrSetmtime                       string `json:"koofrSetmtime" default:"true"`                                                                                                                                // Does the backend support setting modification time.
	KoofrUser                           string `json:"koofrUser"`                                                                                                                                                   // Your user name.
	KoofrPassword                       string `json:"koofrPassword"`                                                                                                                                               // Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password).
	KoofrEncoding                       string `json:"koofrEncoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"`                                                                                             // The encoding for the backend.
	KoofrProvider                       string `json:"koofrProvider"`                                                                                                                                               // Choose your storage provider.
	LocalOneFileSystem                  string `json:"localOneFileSystem" default:"false"`                                                                                                                          // Don't cross filesystem boundaries (unix/macOS only).
	LocalCaseInsensitive                string `json:"localCaseInsensitive" default:"false"`                                                                                                                        // Force the filesystem to report itself as case insensitive.
	LocalCopyLinks                      string `json:"localCopyLinks" default:"false"`                                                                                                                              // Follow symlinks and copy the pointed to item.
	LocalEncoding                       string `json:"localEncoding" default:"Slash,Dot"`                                                                                                                           // The encoding for the backend.
	LocalNounc                          string `json:"localNounc" default:"false"`                                                                                                                                  // Disable UNC (long path names) conversion on Windows.
	LocalLinks                          string `json:"localLinks" default:"false"`                                                                                                                                  // Translate symlinks to/from regular files with a '.rclonelink' extension.
	LocalSkipLinks                      string `json:"localSkipLinks" default:"false"`                                                                                                                              // Don't warn about skipped symlinks.
	LocalUnicodeNormalization           string `json:"localUnicodeNormalization" default:"false"`                                                                                                                   // Apply unicode NFC normalization to paths and filenames.
	LocalNoCheckUpdated                 string `json:"localNoCheckUpdated" default:"false"`                                                                                                                         // Don't check to see if the files change during upload.
	LocalNoSparse                       string `json:"localNoSparse" default:"false"`                                                                                                                               // Disable sparse files for multi-thread downloads.
	LocalNoSetModtime                   string `json:"localNoSetModtime" default:"false"`                                                                                                                           // Disable setting modtime.
	LocalZeroSizeLinks                  string `json:"localZeroSizeLinks" default:"false"`                                                                                                                          // Assume the Stat size of links is zero (and read them instead) (deprecated).
	LocalCaseSensitive                  string `json:"localCaseSensitive" default:"false"`                                                                                                                          // Force the filesystem to report itself as case sensitive.
	LocalNoPreallocate                  string `json:"localNoPreallocate" default:"false"`                                                                                                                          // Disable preallocation of disk space for transferred files.
	MailruSpeedupMaxDisk                string `json:"mailruSpeedupMaxDisk" default:"3Gi"`                                                                                                                          // This option allows you to disable speedup (put by hash) for large files.
	MailruCheckHash                     string `json:"mailruCheckHash" default:"true"`                                                                                                                              // What should copy do if file checksum is mismatched or invalid.
	MailruSpeedupEnable                 string `json:"mailruSpeedupEnable" default:"true"`                                                                                                                          // Skip full upload if there is another file with same data hash.
	MailruPass                          string `json:"mailruPass"`                                                                                                                                                  // Password.
	MailruSpeedupFilePatterns           string `json:"mailruSpeedupFilePatterns" default:"*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf"`                                                                          // Comma separated list of file name patterns eligible for speedup (put by hash).
	MailruSpeedupMaxMemory              string `json:"mailruSpeedupMaxMemory" default:"32Mi"`                                                                                                                       // Files larger than the size given below will always be hashed on disk.
	MailruUserAgent                     string `json:"mailruUserAgent"`                                                                                                                                             // HTTP user agent used internally by client.
	MailruQuirks                        string `json:"mailruQuirks"`                                                                                                                                                // Comma separated list of internal maintenance flags.
	MailruEncoding                      string `json:"mailruEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot"`                                              // The encoding for the backend.
	MailruUser                          string `json:"mailruUser"`                                                                                                                                                  // User name (usually email).
	MegaUser                            string `json:"megaUser"`                                                                                                                                                    // User name.
	MegaPass                            string `json:"megaPass"`                                                                                                                                                    // Password.
	MegaDebug                           string `json:"megaDebug" default:"false"`                                                                                                                                   // Output more debug from Mega.
	MegaHardDelete                      string `json:"megaHardDelete" default:"false"`                                                                                                                              // Delete files permanently rather than putting them into the trash.
	MegaUseHttps                        string `json:"megaUseHttps" default:"false"`                                                                                                                                // Use HTTPS for transfers.
	MegaEncoding                        string `json:"megaEncoding" default:"Slash,InvalidUtf8,Dot"`                                                                                                                // The encoding for the backend.
	NetstorageProtocol                  string `json:"netstorageProtocol" default:"https"`                                                                                                                          // Select between HTTP or HTTPS protocol.
	NetstorageHost                      string `json:"netstorageHost"`                                                                                                                                              // Domain+path of NetStorage host to connect to.
	NetstorageAccount                   string `json:"netstorageAccount"`                                                                                                                                           // Set the NetStorage account name
	NetstorageSecret                    string `json:"netstorageSecret"`                                                                                                                                            // Set the NetStorage account secret/G2O key for authentication.
	OnedriveChunkSize                   string `json:"onedriveChunkSize" default:"10Mi"`                                                                                                                            // Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
	OnedriveDriveType                   string `json:"onedriveDriveType"`                                                                                                                                           // The type of the drive (personal | business | documentLibrary).
	OnedriveRootFolderId                string `json:"onedriveRootFolderId"`                                                                                                                                        // ID of the root folder.
	OnedriveServerSideAcrossConfigs     string `json:"onedriveServerSideAcrossConfigs" default:"false"`                                                                                                             // Allow server-side operations (e.g. copy) to work across different onedrive configs.
	OnedriveNoVersions                  string `json:"onedriveNoVersions" default:"false"`                                                                                                                          // Remove all versions on modifying operations.
	OnedriveToken                       string `json:"onedriveToken"`                                                                                                                                               // OAuth Access Token as a JSON blob.
	OnedriveAuthUrl                     string `json:"onedriveAuthUrl"`                                                                                                                                             // Auth server URL.
	OnedriveAccessScopes                string `json:"onedriveAccessScopes" default:"Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"`                                  // Set scopes to be requested by rclone.
	OnedriveExposeOnenoteFiles          string `json:"onedriveExposeOnenoteFiles" default:"false"`                                                                                                                  // Set to make OneNote files show up in directory listings.
	OnedriveLinkScope                   string `json:"onedriveLinkScope" default:"anonymous"`                                                                                                                       // Set the scope of the links created by the link command.
	OnedriveClientSecret                string `json:"onedriveClientSecret"`                                                                                                                                        // OAuth Client Secret.
	OnedriveDisableSitePermission       string `json:"onedriveDisableSitePermission" default:"false"`                                                                                                               // Disable the request for Sites.Read.All permission.
	OnedriveLinkType                    string `json:"onedriveLinkType" default:"view"`                                                                                                                             // Set the type of the links created by the link command.
	OnedriveClientId                    string `json:"onedriveClientId"`                                                                                                                                            // OAuth Client Id.
	OnedriveRegion                      string `json:"onedriveRegion" default:"global"`                                                                                                                             // Choose national cloud region for OneDrive.
	OnedriveDriveId                     string `json:"onedriveDriveId"`                                                                                                                                             // The ID of the drive to use.
	OnedriveListChunk                   string `json:"onedriveListChunk" default:"1000"`                                                                                                                            // Size of listing chunk.
	OnedriveLinkPassword                string `json:"onedriveLinkPassword"`                                                                                                                                        // Set the password for links created by the link command.
	OnedriveHashType                    string `json:"onedriveHashType" default:"auto"`                                                                                                                             // Specify the hash in use for the backend.
	OnedriveEncoding                    string `json:"onedriveEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
	OnedriveTokenUrl                    string `json:"onedriveTokenUrl"`                                                                                                                                            // Token server url.
	OpendriveUsername                   string `json:"opendriveUsername"`                                                                                                                                           // Username.
	OpendrivePassword                   string `json:"opendrivePassword"`                                                                                                                                           // Password.
	OpendriveEncoding                   string `json:"opendriveEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot"`   // The encoding for the backend.
	OpendriveChunkSize                  string `json:"opendriveChunkSize" default:"10Mi"`                                                                                                                           // Files will be uploaded in chunks this size.
	OosDisableChecksum                  string `json:"oosDisableChecksum" default:"false"`                                                                                                                          // Don't store MD5 checksum with object metadata.
	OosEncoding                         string `json:"oosEncoding" default:"Slash,InvalidUtf8,Dot"`                                                                                                                 // The encoding for the backend.
	OosNoCheckBucket                    string `json:"oosNoCheckBucket" default:"false"`                                                                                                                            // If set, don't attempt to check the bucket exists or create it.
	OosSseCustomerKey                   string `json:"oosSseCustomerKey"`                                                                                                                                           // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	OosSseCustomerKeySha256             string `json:"oosSseCustomerKeySha256"`                                                                                                                                     // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	OosUploadConcurrency                string `json:"oosUploadConcurrency" default:"10"`                                                                                                                           // Concurrency for multipart uploads.
	OosCompartment                      string `json:"oosCompartment"`                                                                                                                                              // Object storage compartment OCID
	OosRegion                           string `json:"oosRegion"`                                                                                                                                                   // Object storage Region
	OosUploadCutoff                     string `json:"oosUploadCutoff" default:"200Mi"`                                                                                                                             // Cutoff for switching to chunked upload.
	OosChunkSize                        string `json:"oosChunkSize" default:"5Mi"`                                                                                                                                  // Chunk size to use for uploading.
	OosNamespace                        string `json:"oosNamespace"`                                                                                                                                                // Object storage namespace
	OosConfigFile                       string `json:"oosConfigFile" default:"~/.oci/config"`                                                                                                                       // Path to OCI config file
	OosConfigProfile                    string `json:"oosConfigProfile" default:"Default"`                                                                                                                          // Profile name inside the oci config file
	OosStorageTier                      string `json:"oosStorageTier" default:"Standard"`                                                                                                                           // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	OosCopyTimeout                      string `json:"oosCopyTimeout" default:"1m0s"`                                                                                                                               // Timeout for copy.
	OosSseCustomerKeyFile               string `json:"oosSseCustomerKeyFile"`                                                                                                                                       // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	OosSseCustomerAlgorithm             string `json:"oosSseCustomerAlgorithm"`                                                                                                                                     // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
	OosProvider                         string `json:"oosProvider" default:"env_auth"`                                                                                                                              // Choose your Auth Provider
	OosCopyCutoff                       string `json:"oosCopyCutoff" default:"4.656Gi"`                                                                                                                             // Cutoff for switching to multipart copy.
	OosLeavePartsOnError                string `json:"oosLeavePartsOnError" default:"false"`                                                                                                                        // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	OosSseKmsKeyId                      string `json:"oosSseKmsKeyId"`                                                                                                                                              // if using using your own master key in vault, this header specifies the
	OosEndpoint                         string `json:"oosEndpoint"`                                                                                                                                                 // Endpoint for Object storage API.
	PcloudTokenUrl                      string `json:"pcloudTokenUrl"`                                                                                                                                              // Token server url.
	PcloudHostname                      string `json:"pcloudHostname" default:"api.pcloud.com"`                                                                                                                     // Hostname to connect to.
	PcloudPassword                      string `json:"pcloudPassword"`                                                                                                                                              // Your pcloud password.
	PcloudClientSecret                  string `json:"pcloudClientSecret"`                                                                                                                                          // OAuth Client Secret.
	PcloudAuthUrl                       string `json:"pcloudAuthUrl"`                                                                                                                                               // Auth server URL.
	PcloudEncoding                      string `json:"pcloudEncoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"`                                                                                            // The encoding for the backend.
	PcloudRootFolderId                  string `json:"pcloudRootFolderId" default:"d0"`                                                                                                                             // Fill in for rclone to use a non root folder as its starting point.
	PcloudUsername                      string `json:"pcloudUsername"`                                                                                                                                              // Your pcloud username.
	PcloudClientId                      string `json:"pcloudClientId"`                                                                                                                                              // OAuth Client Id.
	PcloudToken                         string `json:"pcloudToken"`                                                                                                                                                 // OAuth Access Token as a JSON blob.
	PremiumizemeApiKey                  string `json:"premiumizemeApiKey"`                                                                                                                                          // API Key.
	PremiumizemeEncoding                string `json:"premiumizemeEncoding" default:"Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot"`                                                                          // The encoding for the backend.
	PutioEncoding                       string `json:"putioEncoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"`                                                                                             // The encoding for the backend.
	QingstorEnvAuth                     string `json:"qingstorEnvAuth" default:"false"`                                                                                                                             // Get QingStor credentials from runtime.
	QingstorAccessKeyId                 string `json:"qingstorAccessKeyId"`                                                                                                                                         // QingStor Access Key ID.
	QingstorConnectionRetries           string `json:"qingstorConnectionRetries" default:"3"`                                                                                                                       // Number of connection retries.
	QingstorEncoding                    string `json:"qingstorEncoding" default:"Slash,Ctl,InvalidUtf8"`                                                                                                            // The encoding for the backend.
	QingstorUploadConcurrency           string `json:"qingstorUploadConcurrency" default:"1"`                                                                                                                       // Concurrency for multipart uploads.
	QingstorSecretAccessKey             string `json:"qingstorSecretAccessKey"`                                                                                                                                     // QingStor Secret Access Key (password).
	QingstorEndpoint                    string `json:"qingstorEndpoint"`                                                                                                                                            // Enter an endpoint URL to connection QingStor API.
	QingstorZone                        string `json:"qingstorZone"`                                                                                                                                                // Zone to connect to.
	QingstorUploadCutoff                string `json:"qingstorUploadCutoff" default:"200Mi"`                                                                                                                        // Cutoff for switching to chunked upload.
	QingstorChunkSize                   string `json:"qingstorChunkSize" default:"4Mi"`                                                                                                                             // Chunk size to use for uploading.
	S3Provider                          string `json:"s3Provider"`                                                                                                                                                  // Choose your S3 provider.
	S3StorageClass                      string `json:"s3StorageClass"`                                                                                                                                              // The storage class to use when storing new objects in S3.
	S3MaxUploadParts                    string `json:"s3MaxUploadParts" default:"10000"`                                                                                                                            // Maximum number of parts in a multipart upload.
	S3ListChunk                         string `json:"s3ListChunk" default:"1000"`                                                                                                                                  // Size of listing chunk (response list for each ListObject S3 request).
	S3NoHead                            string `json:"s3NoHead" default:"false"`                                                                                                                                    // If set, don't HEAD uploaded objects to check integrity.
	S3SharedCredentialsFile             string `json:"s3SharedCredentialsFile"`                                                                                                                                     // Path to the shared credentials file.
	S3Profile                           string `json:"s3Profile"`                                                                                                                                                   // Profile to use in the shared credentials file.
	S3DisableHttp2                      string `json:"s3DisableHttp2" default:"false"`                                                                                                                              // Disable usage of http2 for S3 backends.
	S3SseCustomerAlgorithm              string `json:"s3SseCustomerAlgorithm"`                                                                                                                                      // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	S3SseCustomerKeyBase64              string `json:"s3SseCustomerKeyBase64"`                                                                                                                                      // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	S3SseCustomerKeyMd5                 string `json:"s3SseCustomerKeyMd5"`                                                                                                                                         // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	S3ChunkSize                         string `json:"s3ChunkSize" default:"5Mi"`                                                                                                                                   // Chunk size to use for uploading.
	S3DisableChecksum                   string `json:"s3DisableChecksum" default:"false"`                                                                                                                           // Don't store MD5 checksum with object metadata.
	S3UseMultipartEtag                  string `json:"s3UseMultipartEtag" default:"unset"`                                                                                                                          // Whether to use ETag in multipart uploads for verification
	S3Decompress                        string `json:"s3Decompress" default:"false"`                                                                                                                                // If set this will decompress gzip encoded objects.
	S3ListVersion                       string `json:"s3ListVersion" default:"0"`                                                                                                                                   // Version of ListObjects to use: 1,2 or 0 for auto.
	S3EnvAuth                           string `json:"s3EnvAuth" default:"false"`                                                                                                                                   // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	S3BucketAcl                         string `json:"s3BucketAcl"`                                                                                                                                                 // Canned ACL used when creating buckets.
	S3CopyCutoff                        string `json:"s3CopyCutoff" default:"4.656Gi"`                                                                                                                              // Cutoff for switching to multipart copy.
	S3SessionToken                      string `json:"s3SessionToken"`                                                                                                                                              // An AWS session token.
	S3V2Auth                            string `json:"s3V2Auth" default:"false"`                                                                                                                                    // If true use v2 authentication.
	S3Encoding                          string `json:"s3Encoding" default:"Slash,InvalidUtf8,Dot"`                                                                                                                  // The encoding for the backend.
	S3UsePresignedRequest               string `json:"s3UsePresignedRequest" default:"false"`                                                                                                                       // Whether to use a presigned request or PutObject for single part uploads
	S3LocationConstraint                string `json:"s3LocationConstraint"`                                                                                                                                        // Location constraint - must be set to match the Region.
	S3UploadCutoff                      string `json:"s3UploadCutoff" default:"200Mi"`                                                                                                                              // Cutoff for switching to chunked upload.
	S3ListUrlEncode                     string `json:"s3ListUrlEncode" default:"unset"`                                                                                                                             // Whether to url encode listings: true/false/unset
	S3NoCheckBucket                     string `json:"s3NoCheckBucket" default:"false"`                                                                                                                             // If set, don't attempt to check the bucket exists or create it.
	S3NoHeadObject                      string `json:"s3NoHeadObject" default:"false"`                                                                                                                              // If set, do not do HEAD before GET when getting objects.
	S3MemoryPoolUseMmap                 string `json:"s3MemoryPoolUseMmap" default:"false"`                                                                                                                         // Whether to use mmap buffers in internal memory pool.
	S3VersionAt                         string `json:"s3VersionAt" default:"off"`                                                                                                                                   // Show file versions as they were at the specified time.
	S3SecretAccessKey                   string `json:"s3SecretAccessKey"`                                                                                                                                           // AWS Secret Access Key (password).
	S3Endpoint                          string `json:"s3Endpoint"`                                                                                                                                                  // Endpoint for S3 API.
	S3Acl                               string `json:"s3Acl"`                                                                                                                                                       // Canned ACL used when creating buckets and storing or copying objects.
	S3UseAccelerateEndpoint             string `json:"s3UseAccelerateEndpoint" default:"false"`                                                                                                                     // If true use the AWS S3 accelerated endpoint.
	S3MemoryPoolFlushTime               string `json:"s3MemoryPoolFlushTime" default:"1m0s"`                                                                                                                        // How often internal memory buffer pools will be flushed.
	S3ServerSideEncryption              string `json:"s3ServerSideEncryption"`                                                                                                                                      // The server-side encryption algorithm used when storing this object in S3.
	S3SseKmsKeyId                       string `json:"s3SseKmsKeyId"`                                                                                                                                               // If using KMS ID you must provide the ARN of Key.
	S3UploadConcurrency                 string `json:"s3UploadConcurrency" default:"4"`                                                                                                                             // Concurrency for multipart uploads.
	S3MightGzip                         string `json:"s3MightGzip" default:"unset"`                                                                                                                                 // Set this if the backend might gzip objects.
	S3NoSystemMetadata                  string `json:"s3NoSystemMetadata" default:"false"`                                                                                                                          // Suppress setting and reading of system metadata
	S3AccessKeyId                       string `json:"s3AccessKeyId"`                                                                                                                                               // AWS Access Key ID.
	S3StsEndpoint                       string `json:"s3StsEndpoint"`                                                                                                                                               // Endpoint for STS.
	S3DownloadUrl                       string `json:"s3DownloadUrl"`                                                                                                                                               // Custom endpoint for downloads.
	S3Versions                          string `json:"s3Versions" default:"false"`                                                                                                                                  // Include old versions in directory listings.
	S3Region                            string `json:"s3Region"`                                                                                                                                                    // Region to connect to.
	S3RequesterPays                     string `json:"s3RequesterPays" default:"false"`                                                                                                                             // Enables requester pays option when interacting with S3 bucket.
	S3SseCustomerKey                    string `json:"s3SseCustomerKey"`                                                                                                                                            // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	S3ForcePathStyle                    string `json:"s3ForcePathStyle" default:"true"`                                                                                                                             // If true use path style access if false use virtual hosted style.
	S3LeavePartsOnError                 string `json:"s3LeavePartsOnError" default:"false"`                                                                                                                         // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	SeafileCreateLibrary                string `json:"seafileCreateLibrary" default:"false"`                                                                                                                        // Should rclone create a library if it doesn't exist.
	SeafileAuthToken                    string `json:"seafileAuthToken"`                                                                                                                                            // Authentication token.
	SeafileUrl                          string `json:"seafileUrl"`                                                                                                                                                  // URL of seafile host to connect to.
	SeafileUser                         string `json:"seafileUser"`                                                                                                                                                 // User name (usually email address).
	SeafilePass                         string `json:"seafilePass"`                                                                                                                                                 // Password.
	Seafile2fa                          string `json:"seafile2fa" default:"false"`                                                                                                                                  // Two-factor authentication ('true' if the account has 2FA enabled).
	SeafileLibrary                      string `json:"seafileLibrary"`                                                                                                                                              // Name of the library.
	SeafileLibraryKey                   string `json:"seafileLibraryKey"`                                                                                                                                           // Library password (for encrypted libraries only).
	SeafileEncoding                     string `json:"seafileEncoding" default:"Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8"`                                                                                       // The encoding for the backend.
	SftpConcurrency                     string `json:"sftpConcurrency" default:"64"`                                                                                                                                // The maximum number of outstanding requests for one file
	SftpKnownHostsFile                  string `json:"sftpKnownHostsFile"`                                                                                                                                          // Optional path to known_hosts file.
	SftpUseInsecureCipher               string `json:"sftpUseInsecureCipher" default:"false"`                                                                                                                       // Enable the use of insecure ciphers and key exchange methods.
	SftpShellType                       string `json:"sftpShellType"`                                                                                                                                               // The type of SSH shell on remote server, if any.
	SftpPass                            string `json:"sftpPass"`                                                                                                                                                    // SSH password, leave blank to use ssh-agent.
	SftpChunkSize                       string `json:"sftpChunkSize" default:"32Ki"`                                                                                                                                // Upload and download chunk size.
	SftpCiphers                         string `json:"sftpCiphers"`                                                                                                                                                 // Space separated list of ciphers to be used for session encryption, ordered by preference.
	SftpIdleTimeout                     string `json:"sftpIdleTimeout" default:"1m0s"`                                                                                                                              // Max time before closing idle connections.
	SftpKeyPem                          string `json:"sftpKeyPem"`                                                                                                                                                  // Raw PEM-encoded private key.
	SftpSkipLinks                       string `json:"sftpSkipLinks" default:"false"`                                                                                                                               // Set to skip any symlinks and any other non regular files.
	SftpDisableConcurrentWrites         string `json:"sftpDisableConcurrentWrites" default:"false"`                                                                                                                 // If set don't use concurrent writes.
	SftpMd5sumCommand                   string `json:"sftpMd5sumCommand"`                                                                                                                                           // The command used to read md5 hashes.
	SftpSubsystem                       string `json:"sftpSubsystem" default:"sftp"`                                                                                                                                // Specifies the SSH2 subsystem on the remote host.
	SftpHost                            string `json:"sftpHost"`                                                                                                                                                    // SSH host to connect to.
	SftpAskPassword                     string `json:"sftpAskPassword" default:"false"`                                                                                                                             // Allow asking for SFTP password when needed.
	SftpSetModtime                      string `json:"sftpSetModtime" default:"true"`                                                                                                                               // Set the modified time on the remote if set.
	SftpKeyFile                         string `json:"sftpKeyFile"`                                                                                                                                                 // Path to PEM-encoded private key file.
	SftpKeyExchange                     string `json:"sftpKeyExchange"`                                                                                                                                             // Space separated list of key exchange algorithms, ordered by preference.
	SftpSetEnv                          string `json:"sftpSetEnv"`                                                                                                                                                  // Environment variables to pass to sftp and commands
	SftpPubkeyFile                      string `json:"sftpPubkeyFile"`                                                                                                                                              // Optional path to public key file.
	SftpServerCommand                   string `json:"sftpServerCommand"`                                                                                                                                           // Specifies the path or command to run a sftp server on the remote host.
	SftpDisableConcurrentReads          string `json:"sftpDisableConcurrentReads" default:"false"`                                                                                                                  // If set don't use concurrent reads.
	SftpUseFstat                        string `json:"sftpUseFstat" default:"false"`                                                                                                                                // If set use fstat instead of stat.
	SftpMacs                            string `json:"sftpMacs"`                                                                                                                                                    // Space separated list of MACs (message authentication code) algorithms, ordered by preference.
	SftpUser                            string `json:"sftpUser" default:"shane"`                                                                                                                                    // SSH username.
	SftpKeyUseAgent                     string `json:"sftpKeyUseAgent" default:"false"`                                                                                                                             // When set forces the usage of the ssh-agent.
	SftpDisableHashcheck                string `json:"sftpDisableHashcheck" default:"false"`                                                                                                                        // Disable the execution of SSH commands to determine if remote file hashing is available.
	SftpSha1sumCommand                  string `json:"sftpSha1sumCommand"`                                                                                                                                          // The command used to read sha1 hashes.
	SftpPort                            string `json:"sftpPort" default:"22"`                                                                                                                                       // SSH port number.
	SftpKeyFilePass                     string `json:"sftpKeyFilePass"`                                                                                                                                             // The passphrase to decrypt the PEM-encoded private key file.
	SftpPathOverride                    string `json:"sftpPathOverride"`                                                                                                                                            // Override path used by SSH shell commands.
	SharefileUploadCutoff               string `json:"sharefileUploadCutoff" default:"128Mi"`                                                                                                                       // Cutoff for switching to multipart upload.
	SharefileRootFolderId               string `json:"sharefileRootFolderId"`                                                                                                                                       // ID of the root folder.
	SharefileChunkSize                  string `json:"sharefileChunkSize" default:"64Mi"`                                                                                                                           // Upload chunk size.
	SharefileEndpoint                   string `json:"sharefileEndpoint"`                                                                                                                                           // Endpoint for API calls.
	SharefileEncoding                   string `json:"sharefileEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot"`   // The encoding for the backend.
	SiaApiUrl                           string `json:"siaApiUrl" default:"http://127.0.0.1:9980"`                                                                                                                   // Sia daemon API URL, like http://sia.daemon.host:9980.
	SiaApiPassword                      string `json:"siaApiPassword"`                                                                                                                                              // Sia Daemon API Password.
	SiaUserAgent                        string `json:"siaUserAgent" default:"Sia-Agent"`                                                                                                                            // Siad User Agent
	SiaEncoding                         string `json:"siaEncoding" default:"Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot"`                                                                                   // The encoding for the backend.
	SmbSpn                              string `json:"smbSpn"`                                                                                                                                                      // Service principal name.
	SmbIdleTimeout                      string `json:"smbIdleTimeout" default:"1m0s"`                                                                                                                               // Max time before closing idle connections.
	SmbHideSpecialShare                 string `json:"smbHideSpecialShare" default:"true"`                                                                                                                          // Hide special shares (e.g. print$) which users aren't supposed to access.
	SmbCaseInsensitive                  string `json:"smbCaseInsensitive" default:"true"`                                                                                                                           // Whether the server is configured to be case-insensitive.
	SmbEncoding                         string `json:"smbEncoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot"`                              // The encoding for the backend.
	SmbHost                             string `json:"smbHost"`                                                                                                                                                     // SMB server hostname to connect to.
	SmbPort                             string `json:"smbPort" default:"445"`                                                                                                                                       // SMB port number.
	SmbPass                             string `json:"smbPass"`                                                                                                                                                     // SMB password.
	SmbDomain                           string `json:"smbDomain" default:"WORKGROUP"`                                                                                                                               // Domain name for NTLM authentication.
	SmbUser                             string `json:"smbUser" default:"shane"`                                                                                                                                     // SMB username.
	StorjAccessGrant                    string `json:"storjAccessGrant"`                                                                                                                                            // Access grant.
	StorjSatelliteAddress               string `json:"storjSatelliteAddress" default:"us1.storj.io"`                                                                                                                // Satellite address.
	StorjApiKey                         string `json:"storjApiKey"`                                                                                                                                                 // API key.
	StorjPassphrase                     string `json:"storjPassphrase"`                                                                                                                                             // Encryption passphrase.
	StorjProvider                       string `json:"storjProvider" default:"existing"`                                                                                                                            // Choose an authentication method.
	TardigradeAccessGrant               string `json:"tardigradeAccessGrant"`                                                                                                                                       // Access grant.
	TardigradeSatelliteAddress          string `json:"tardigradeSatelliteAddress" default:"us1.storj.io"`                                                                                                           // Satellite address.
	TardigradeApiKey                    string `json:"tardigradeApiKey"`                                                                                                                                            // API key.
	TardigradePassphrase                string `json:"tardigradePassphrase"`                                                                                                                                        // Encryption passphrase.
	TardigradeProvider                  string `json:"tardigradeProvider" default:"existing"`                                                                                                                       // Choose an authentication method.
	SugarsyncAuthorizationExpiry        string `json:"sugarsyncAuthorizationExpiry"`                                                                                                                                // Sugarsync authorization expiry.
	SugarsyncDeletedId                  string `json:"sugarsyncDeletedId"`                                                                                                                                          // Sugarsync deleted folder id.
	SugarsyncEncoding                   string `json:"sugarsyncEncoding" default:"Slash,Ctl,InvalidUtf8,Dot"`                                                                                                       // The encoding for the backend.
	SugarsyncAppId                      string `json:"sugarsyncAppId"`                                                                                                                                              // Sugarsync App ID.
	SugarsyncAccessKeyId                string `json:"sugarsyncAccessKeyId"`                                                                                                                                        // Sugarsync Access Key ID.
	SugarsyncPrivateAccessKey           string `json:"sugarsyncPrivateAccessKey"`                                                                                                                                   // Sugarsync Private Access Key.
	SugarsyncAuthorization              string `json:"sugarsyncAuthorization"`                                                                                                                                      // Sugarsync authorization.
	SugarsyncHardDelete                 string `json:"sugarsyncHardDelete" default:"false"`                                                                                                                         // Permanently delete files if true
	SugarsyncRefreshToken               string `json:"sugarsyncRefreshToken"`                                                                                                                                       // Sugarsync refresh token.
	SugarsyncUser                       string `json:"sugarsyncUser"`                                                                                                                                               // Sugarsync user.
	SugarsyncRootId                     string `json:"sugarsyncRootId"`                                                                                                                                             // Sugarsync root id.
	SwiftTenantDomain                   string `json:"swiftTenantDomain"`                                                                                                                                           // Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME).
	SwiftRegion                         string `json:"swiftRegion"`                                                                                                                                                 // Region name - optional (OS_REGION_NAME).
	SwiftAuthVersion                    string `json:"swiftAuthVersion" default:"0"`                                                                                                                                // AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).
	SwiftEncoding                       string `json:"swiftEncoding" default:"Slash,InvalidUtf8"`                                                                                                                   // The encoding for the backend.
	SwiftDomain                         string `json:"swiftDomain"`                                                                                                                                                 // User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME)
	SwiftTenant                         string `json:"swiftTenant"`                                                                                                                                                 // Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME).
	SwiftTenantId                       string `json:"swiftTenantId"`                                                                                                                                               // Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID).
	SwiftStorageUrl                     string `json:"swiftStorageUrl"`                                                                                                                                             // Storage URL - optional (OS_STORAGE_URL).
	SwiftAuthToken                      string `json:"swiftAuthToken"`                                                                                                                                              // Auth Token from alternate authentication - optional (OS_AUTH_TOKEN).
	SwiftApplicationCredentialId        string `json:"swiftApplicationCredentialId"`                                                                                                                                // Application Credential ID (OS_APPLICATION_CREDENTIAL_ID).
	SwiftLeavePartsOnError              string `json:"swiftLeavePartsOnError" default:"false"`                                                                                                                      // If true avoid calling abort upload on a failure.
	SwiftStoragePolicy                  string `json:"swiftStoragePolicy"`                                                                                                                                          // The storage policy to use when creating a new container.
	SwiftUserId                         string `json:"swiftUserId"`                                                                                                                                                 // User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID).
	SwiftNoChunk                        string `json:"swiftNoChunk" default:"false"`                                                                                                                                // Don't chunk files during streaming upload.
	SwiftAuth                           string `json:"swiftAuth"`                                                                                                                                                   // Authentication URL for server (OS_AUTH_URL).
	SwiftEndpointType                   string `json:"swiftEndpointType" default:"public"`                                                                                                                          // Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).
	SwiftNoLargeObjects                 string `json:"swiftNoLargeObjects" default:"false"`                                                                                                                         // Disable support for static and dynamic large objects
	SwiftUser                           string `json:"swiftUser"`                                                                                                                                                   // User name to log in (OS_USERNAME).
	SwiftKey                            string `json:"swiftKey"`                                                                                                                                                    // API key or password (OS_PASSWORD).
	SwiftApplicationCredentialName      string `json:"swiftApplicationCredentialName"`                                                                                                                              // Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME).
	SwiftApplicationCredentialSecret    string `json:"swiftApplicationCredentialSecret"`                                                                                                                            // Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET).
	SwiftChunkSize                      string `json:"swiftChunkSize" default:"5Gi"`                                                                                                                                // Above this size files will be chunked into a _segments container.
	SwiftEnvAuth                        string `json:"swiftEnvAuth" default:"false"`                                                                                                                                // Get swift credentials from environment variables in standard OpenStack form.
	UptoboxAccessToken                  string `json:"uptoboxAccessToken"`                                                                                                                                          // Your access token.
	UptoboxEncoding                     string `json:"uptoboxEncoding" default:"Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot"`                                                                // The encoding for the backend.
	WebdavEncoding                      string `json:"webdavEncoding"`                                                                                                                                              // The encoding for the backend.
	WebdavHeaders                       string `json:"webdavHeaders"`                                                                                                                                               // Set HTTP headers for all transactions.
	WebdavUrl                           string `json:"webdavUrl"`                                                                                                                                                   // URL of http host to connect to.
	WebdavVendor                        string `json:"webdavVendor"`                                                                                                                                                // Name of the WebDAV site/service/software you are using.
	WebdavUser                          string `json:"webdavUser"`                                                                                                                                                  // User name.
	WebdavPass                          string `json:"webdavPass"`                                                                                                                                                  // Password.
	WebdavBearerToken                   string `json:"webdavBearerToken"`                                                                                                                                           // Bearer token instead of user/pass (e.g. a Macaroon).
	WebdavBearerTokenCommand            string `json:"webdavBearerTokenCommand"`                                                                                                                                    // Command to run to get a bearer token.
	YandexClientId                      string `json:"yandexClientId"`                                                                                                                                              // OAuth Client Id.
	YandexClientSecret                  string `json:"yandexClientSecret"`                                                                                                                                          // OAuth Client Secret.
	YandexToken                         string `json:"yandexToken"`                                                                                                                                                 // OAuth Access Token as a JSON blob.
	YandexAuthUrl                       string `json:"yandexAuthUrl"`                                                                                                                                               // Auth server URL.
	YandexTokenUrl                      string `json:"yandexTokenUrl"`                                                                                                                                              // Token server url.
	YandexHardDelete                    string `json:"yandexHardDelete" default:"false"`                                                                                                                            // Delete files permanently rather than putting them into the trash.
	YandexEncoding                      string `json:"yandexEncoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"`                                                                                                      // The encoding for the backend.
	ZohoTokenUrl                        string `json:"zohoTokenUrl"`                                                                                                                                                // Token server url.
	ZohoRegion                          string `json:"zohoRegion"`                                                                                                                                                  // Zoho region to connect to.
	ZohoEncoding                        string `json:"zohoEncoding" default:"Del,Ctl,InvalidUtf8"`                                                                                                                  // The encoding for the backend.
	ZohoClientId                        string `json:"zohoClientId"`                                                                                                                                                // OAuth Client Id.
	ZohoClientSecret                    string `json:"zohoClientSecret"`                                                                                                                                            // OAuth Client Secret.
	ZohoToken                           string `json:"zohoToken"`                                                                                                                                                   // OAuth Access Token as a JSON blob.
	ZohoAuthUrl                         string `json:"zohoAuthUrl"`                                                                                                                                                 // Auth server URL.
}
