// Code generated. DO NOT EDIT.
//
//lint:file-ignore U1000 Ignore all unused code, it's generated
package storage

import "github.com/data-preservation-programs/singularity/model"

type hdfsConfig struct {
	Namenode               string `json:"namenode"`                                               // Hadoop name node and port.
	Username               string `json:"username" example:"root"`                                // Hadoop user name.
	ServicePrincipalName   string `json:"servicePrincipalName"`                                   // Kerberos service principal name for the namenode.
	DataTransferProtection string `json:"dataTransferProtection" example:"privacy"`               // Kerberos data transfer protection: authentication|integrity|privacy.
	Encoding               string `json:"encoding" default:"Slash,Colon,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createHdfsStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type localConfig struct {
	Nounc                bool   `json:"nounc" default:"false" example:"true"` // Disable UNC (long path names) conversion on Windows.
	CopyLinks            bool   `json:"copyLinks" default:"false"`            // Follow symlinks and copy the pointed to item.
	Links                bool   `json:"links" default:"false"`                // Translate symlinks to/from regular files with a '.rclonelink' extension.
	SkipLinks            bool   `json:"skipLinks" default:"false"`            // Don't warn about skipped symlinks.
	ZeroSizeLinks        bool   `json:"zeroSizeLinks" default:"false"`        // Assume the Stat size of links is zero (and read them instead) (deprecated).
	UnicodeNormalization bool   `json:"unicodeNormalization" default:"false"` // Apply unicode NFC normalization to paths and filenames.
	NoCheckUpdated       bool   `json:"noCheckUpdated" default:"false"`       // Don't check to see if the files change during upload.
	OneFileSystem        bool   `json:"oneFileSystem" default:"false"`        // Don't cross filesystem boundaries (unix/macOS only).
	CaseSensitive        bool   `json:"caseSensitive" default:"false"`        // Force the filesystem to report itself as case sensitive.
	CaseInsensitive      bool   `json:"caseInsensitive" default:"false"`      // Force the filesystem to report itself as case insensitive.
	NoPreallocate        bool   `json:"noPreallocate" default:"false"`        // Disable preallocation of disk space for transferred files.
	NoSparse             bool   `json:"noSparse" default:"false"`             // Disable sparse files for multi-thread downloads.
	NoSetModtime         bool   `json:"noSetModtime" default:"false"`         // Disable setting modtime.
	Encoding             string `json:"encoding" default:"Slash,Dot"`         // The encoding for the backend.
}

type createLocalStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type storjExistingConfig struct {
	AccessGrant string `json:"accessGrant"` // Access grant.
}

type createStorjExistingStorageRequest struct {
	Name         string              `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	SatelliteAddress string `json:"satelliteAddress" default:"us1.storj.io" example:"us1.storj.io"` // Satellite address.
	ApiKey           string `json:"apiKey"`                                                         // API key.
	Passphrase       string `json:"passphrase"`                                                     // Encryption passphrase.
}

type createStorjNewStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type azureblobConfig struct {
	Account                    string `json:"account"`                                                            // Azure Storage Account Name.
	EnvAuth                    bool   `json:"envAuth" default:"false"`                                            // Read credentials from runtime (environment variables, CLI or MSI).
	Key                        string `json:"key"`                                                                // Storage Account Shared Key.
	SasUrl                     string `json:"sasUrl"`                                                             // SAS URL for container level access only.
	Tenant                     string `json:"tenant"`                                                             // ID of the service principal's tenant. Also called its directory ID.
	ClientId                   string `json:"clientId"`                                                           // The ID of the client in use.
	ClientSecret               string `json:"clientSecret"`                                                       // One of the service principal's client secrets
	ClientCertificatePath      string `json:"clientCertificatePath"`                                              // Path to a PEM or PKCS12 certificate file including the private key.
	ClientCertificatePassword  string `json:"clientCertificatePassword"`                                          // Password for the certificate file (optional).
	ClientSendCertificateChain bool   `json:"clientSendCertificateChain" default:"false"`                         // Send the certificate chain when using certificate auth.
	Username                   string `json:"username"`                                                           // User name (usually an email address)
	Password                   string `json:"password"`                                                           // The user's password
	ServicePrincipalFile       string `json:"servicePrincipalFile"`                                               // Path to file containing credentials for use with a service principal.
	UseMsi                     bool   `json:"useMsi" default:"false"`                                             // Use a managed service identity to authenticate (only works in Azure).
	MsiObjectId                string `json:"msiObjectId"`                                                        // Object ID of the user-assigned MSI to use, if any.
	MsiClientId                string `json:"msiClientId"`                                                        // Object ID of the user-assigned MSI to use, if any.
	MsiMiResId                 string `json:"msiMiResId"`                                                         // Azure resource ID of the user-assigned MSI to use, if any.
	UseEmulator                bool   `json:"useEmulator" default:"false"`                                        // Uses local storage emulator if provided as 'true'.
	Endpoint                   string `json:"endpoint"`                                                           // Endpoint for the service.
	UploadCutoff               string `json:"uploadCutoff"`                                                       // Cutoff for switching to chunked upload (<= 256 MiB) (deprecated).
	ChunkSize                  string `json:"chunkSize" default:"4Mi"`                                            // Upload chunk size.
	UploadConcurrency          int    `json:"uploadConcurrency" default:"16"`                                     // Concurrency for multipart uploads.
	ListChunk                  int    `json:"listChunk" default:"5000"`                                           // Size of blob list.
	AccessTier                 string `json:"accessTier"`                                                         // Access tier of blob: hot, cool or archive.
	ArchiveTierDelete          bool   `json:"archiveTierDelete" default:"false"`                                  // Delete archive tier blobs before overwriting.
	DisableChecksum            bool   `json:"disableChecksum" default:"false"`                                    // Don't store MD5 checksum with object metadata.
	MemoryPoolFlushTime        string `json:"memoryPoolFlushTime" default:"1m0s"`                                 // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap          bool   `json:"memoryPoolUseMmap" default:"false"`                                  // Whether to use mmap buffers in internal memory pool.
	Encoding                   string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"` // The encoding for the backend.
	PublicAccess               string `json:"publicAccess" example:""`                                            // Public access level of a container: blob or container.
	NoCheckContainer           bool   `json:"noCheckContainer" default:"false"`                                   // If set, don't attempt to check the container exists or create it.
	NoHeadObject               bool   `json:"noHeadObject" default:"false"`                                       // If set, do not do HEAD before GET when getting objects.
}

type createAzureblobStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Account              string `json:"account"`                                                    // Account ID or Application Key ID.
	Key                  string `json:"key"`                                                        // Application Key.
	Endpoint             string `json:"endpoint"`                                                   // Endpoint for the service.
	TestMode             string `json:"testMode"`                                                   // A flag string for X-Bz-Test-Mode header for debugging.
	Versions             bool   `json:"versions" default:"false"`                                   // Include old versions in directory listings.
	VersionAt            string `json:"versionAt" default:"off"`                                    // Show file versions as they were at the specified time.
	HardDelete           bool   `json:"hardDelete" default:"false"`                                 // Permanently delete files on remote removal, otherwise hide files.
	UploadCutoff         string `json:"uploadCutoff" default:"200Mi"`                               // Cutoff for switching to chunked upload.
	CopyCutoff           string `json:"copyCutoff" default:"4Gi"`                                   // Cutoff for switching to multipart copy.
	ChunkSize            string `json:"chunkSize" default:"96Mi"`                                   // Upload chunk size.
	DisableChecksum      bool   `json:"disableChecksum" default:"false"`                            // Disable checksums for large (> upload cutoff) files.
	DownloadUrl          string `json:"downloadUrl"`                                                // Custom endpoint for downloads.
	DownloadAuthDuration string `json:"downloadAuthDuration" default:"1w"`                          // Time before the authorization token will expire in s or suffix ms|s|m|h|d.
	MemoryPoolFlushTime  string `json:"memoryPoolFlushTime" default:"1m0s"`                         // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap    bool   `json:"memoryPoolUseMmap" default:"false"`                          // Whether to use mmap buffers in internal memory pool.
	Encoding             string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createB2StorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type fichierConfig struct {
	ApiKey         string `json:"apiKey"`                                                                                                                        // Your API Key, get it from https://1fichier.com/console/params.pl.
	SharedFolder   string `json:"sharedFolder"`                                                                                                                  // If you want to download a shared folder, add this parameter.
	FilePassword   string `json:"filePassword"`                                                                                                                  // If you want to download a shared file that is password protected, add this parameter.
	FolderPassword string `json:"folderPassword"`                                                                                                                // If you want to list the files in a shared folder that is password protected, add this parameter.
	Encoding       string `json:"encoding" default:"Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createFichierStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type ftpConfig struct {
	Host               string `json:"host"`                                                                             // FTP host to connect to.
	User               string `json:"user" default:"$USER"`                                                             // FTP username.
	Port               int    `json:"port" default:"21"`                                                                // FTP port number.
	Pass               string `json:"pass"`                                                                             // FTP password.
	Tls                bool   `json:"tls" default:"false"`                                                              // Use Implicit FTPS (FTP over TLS).
	ExplicitTls        bool   `json:"explicitTls" default:"false"`                                                      // Use Explicit FTPS (FTP over TLS).
	Concurrency        int    `json:"concurrency" default:"0"`                                                          // Maximum number of FTP simultaneous connections, 0 for unlimited.
	NoCheckCertificate bool   `json:"noCheckCertificate" default:"false"`                                               // Do not verify the TLS certificate of the server.
	DisableEpsv        bool   `json:"disableEpsv" default:"false"`                                                      // Disable using EPSV even if server advertises support.
	DisableMlsd        bool   `json:"disableMlsd" default:"false"`                                                      // Disable using MLSD even if server advertises support.
	DisableUtf8        bool   `json:"disableUtf8" default:"false"`                                                      // Disable using UTF-8 even if server advertises support.
	WritingMdtm        bool   `json:"writingMdtm" default:"false"`                                                      // Use MDTM to set modification time (VsFtpd quirk)
	ForceListHidden    bool   `json:"forceListHidden" default:"false"`                                                  // Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.
	IdleTimeout        string `json:"idleTimeout" default:"1m0s"`                                                       // Max time before closing idle connections.
	CloseTimeout       string `json:"closeTimeout" default:"1m0s"`                                                      // Maximum time to wait for a response to close.
	TlsCacheSize       int    `json:"tlsCacheSize" default:"32"`                                                        // Size of TLS session cache for all control and data connections.
	DisableTls13       bool   `json:"disableTls13" default:"false"`                                                     // Disable TLS 1.3 (workaround for FTP servers with buggy TLS)
	ShutTimeout        string `json:"shutTimeout" default:"1m0s"`                                                       // Maximum time to wait for data connection closing status.
	AskPassword        bool   `json:"askPassword" default:"false"`                                                      // Allow asking for FTP password when needed.
	Encoding           string `json:"encoding" default:"Slash,Del,Ctl,RightSpace,Dot" example:"Asterisk,Ctl,Dot,Slash"` // The encoding for the backend.
}

type createFtpStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type httpConfig struct {
	Url     string `json:"url"`                     // URL of HTTP host to connect to.
	Headers string `json:"headers"`                 // Set HTTP headers for all transactions.
	NoSlash bool   `json:"noSlash" default:"false"` // Set this if the site doesn't end directories with /.
	NoHead  bool   `json:"noHead" default:"false"`  // Don't use HEAD requests.
}

type createHttpStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	AccessKeyId     string `json:"accessKeyId"`                                                // IAS3 Access Key.
	SecretAccessKey string `json:"secretAccessKey"`                                            // IAS3 Secret Key (password).
	Endpoint        string `json:"endpoint" default:"https://s3.us.archive.org"`               // IAS3 Endpoint.
	FrontEndpoint   string `json:"frontEndpoint" default:"https://archive.org"`                // Host of InternetArchive Frontend.
	DisableChecksum bool   `json:"disableChecksum" default:"true"`                             // Don't ask the server to test against MD5 checksum calculated by rclone.
	WaitArchive     string `json:"waitArchive" default:"0s"`                                   // Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
	Encoding        string `json:"encoding" default:"Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createInternetarchiveStorageRequest struct {
	Name         string                `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type koofrDigistorageConfig struct {
	Mountid  string `json:"mountid"`                                                    // Mount ID of the mount to use.
	Setmtime bool   `json:"setmtime" default:"true"`                                    // Does the backend support setting modification time.
	User     string `json:"user"`                                                       // Your user name.
	Password string `json:"password"`                                                   // Your password for rclone (generate one at https://storage.rcs-rds.ro/app/admin/preferences/password).
	Encoding string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createKoofrDigistorageStorageRequest struct {
	Name         string                 `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Setmtime bool   `json:"setmtime" default:"true"`                                    // Does the backend support setting modification time.
	User     string `json:"user"`                                                       // Your user name.
	Password string `json:"password"`                                                   // Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password).
	Encoding string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createKoofrKoofrStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Setmtime bool   `json:"setmtime" default:"true"`                                    // Does the backend support setting modification time.
	User     string `json:"user"`                                                       // Your user name.
	Password string `json:"password"`                                                   // Your password for rclone (generate one at your service's settings page).
	Encoding string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createKoofrOtherStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type megaConfig struct {
	User       string `json:"user"`                                     // User name.
	Pass       string `json:"pass"`                                     // Password.
	Debug      bool   `json:"debug" default:"false"`                    // Output more debug from Mega.
	HardDelete bool   `json:"hardDelete" default:"false"`               // Delete files permanently rather than putting them into the trash.
	UseHttps   bool   `json:"useHttps" default:"false"`                 // Use HTTPS for transfers.
	Encoding   string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createMegaStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Protocol string `json:"protocol" default:"https" example:"http"` // Select between HTTP or HTTPS protocol.
	Host     string `json:"host"`                                    // Domain+path of NetStorage host to connect to.
	Account  string `json:"account"`                                 // Set the NetStorage account name
	Secret   string `json:"secret"`                                  // Set the NetStorage account secret/G2O key for authentication.
}

type createNetstorageStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type opendriveConfig struct {
	Username  string `json:"username"`                                                                                                                                         // Username.
	Password  string `json:"password"`                                                                                                                                         // Password.
	Encoding  string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot"` // The encoding for the backend.
	ChunkSize string `json:"chunkSize" default:"10Mi"`                                                                                                                         // Files will be uploaded in chunks this size.
}

type createOpendriveStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type qingstorConfig struct {
	EnvAuth           bool   `json:"envAuth" default:"false" example:"false"`  // Get QingStor credentials from runtime.
	AccessKeyId       string `json:"accessKeyId"`                              // QingStor Access Key ID.
	SecretAccessKey   string `json:"secretAccessKey"`                          // QingStor Secret Access Key (password).
	Endpoint          string `json:"endpoint"`                                 // Enter an endpoint URL to connection QingStor API.
	Zone              string `json:"zone" example:"pek3a"`                     // Zone to connect to.
	ConnectionRetries int    `json:"connectionRetries" default:"3"`            // Number of connection retries.
	UploadCutoff      string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize         string `json:"chunkSize" default:"4Mi"`                  // Chunk size to use for uploading.
	UploadConcurrency int    `json:"uploadConcurrency" default:"1"`            // Concurrency for multipart uploads.
	Encoding          string `json:"encoding" default:"Slash,Ctl,InvalidUtf8"` // The encoding for the backend.
}

type createQingstorStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type seafileConfig struct {
	Url           string `json:"url" example:"https://cloud.seafile.com/"`                       // URL of seafile host to connect to.
	User          string `json:"user"`                                                           // User name (usually email address).
	Pass          string `json:"pass"`                                                           // Password.
	TwoFA         bool   `json:"2fa" default:"false"`                                            // Two-factor authentication ('true' if the account has 2FA enabled).
	Library       string `json:"library"`                                                        // Name of the library.
	LibraryKey    string `json:"libraryKey"`                                                     // Library password (for encrypted libraries only).
	CreateLibrary bool   `json:"createLibrary" default:"false"`                                  // Should rclone create a library if it doesn't exist.
	AuthToken     string `json:"authToken"`                                                      // Authentication token.
	Encoding      string `json:"encoding" default:"Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8"` // The encoding for the backend.
}

type createSeafileStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Host                    string `json:"host"`                                              // SSH host to connect to.
	User                    string `json:"user" default:"$USER"`                              // SSH username.
	Port                    int    `json:"port" default:"22"`                                 // SSH port number.
	Pass                    string `json:"pass"`                                              // SSH password, leave blank to use ssh-agent.
	KeyPem                  string `json:"keyPem"`                                            // Raw PEM-encoded private key.
	KeyFile                 string `json:"keyFile"`                                           // Path to PEM-encoded private key file.
	KeyFilePass             string `json:"keyFilePass"`                                       // The passphrase to decrypt the PEM-encoded private key file.
	PubkeyFile              string `json:"pubkeyFile"`                                        // Optional path to public key file.
	KnownHostsFile          string `json:"knownHostsFile" example:"~/.ssh/known_hosts"`       // Optional path to known_hosts file.
	KeyUseAgent             bool   `json:"keyUseAgent" default:"false"`                       // When set forces the usage of the ssh-agent.
	UseInsecureCipher       bool   `json:"useInsecureCipher" default:"false" example:"false"` // Enable the use of insecure ciphers and key exchange methods.
	DisableHashcheck        bool   `json:"disableHashcheck" default:"false"`                  // Disable the execution of SSH commands to determine if remote file hashing is available.
	AskPassword             bool   `json:"askPassword" default:"false"`                       // Allow asking for SFTP password when needed.
	PathOverride            string `json:"pathOverride"`                                      // Override path used by SSH shell commands.
	SetModtime              bool   `json:"setModtime" default:"true"`                         // Set the modified time on the remote if set.
	ShellType               string `json:"shellType" example:"none"`                          // The type of SSH shell on remote server, if any.
	Md5sumCommand           string `json:"md5sumCommand"`                                     // The command used to read md5 hashes.
	Sha1sumCommand          string `json:"sha1sumCommand"`                                    // The command used to read sha1 hashes.
	SkipLinks               bool   `json:"skipLinks" default:"false"`                         // Set to skip any symlinks and any other non regular files.
	Subsystem               string `json:"subsystem" default:"sftp"`                          // Specifies the SSH2 subsystem on the remote host.
	ServerCommand           string `json:"serverCommand"`                                     // Specifies the path or command to run a sftp server on the remote host.
	UseFstat                bool   `json:"useFstat" default:"false"`                          // If set use fstat instead of stat.
	DisableConcurrentReads  bool   `json:"disableConcurrentReads" default:"false"`            // If set don't use concurrent reads.
	DisableConcurrentWrites bool   `json:"disableConcurrentWrites" default:"false"`           // If set don't use concurrent writes.
	IdleTimeout             string `json:"idleTimeout" default:"1m0s"`                        // Max time before closing idle connections.
	ChunkSize               string `json:"chunkSize" default:"32Ki"`                          // Upload and download chunk size.
	Concurrency             int    `json:"concurrency" default:"64"`                          // The maximum number of outstanding requests for one file
	SetEnv                  string `json:"setEnv"`                                            // Environment variables to pass to sftp and commands
	Ciphers                 string `json:"ciphers"`                                           // Space separated list of ciphers to be used for session encryption, ordered by preference.
	KeyExchange             string `json:"keyExchange"`                                       // Space separated list of key exchange algorithms, ordered by preference.
	Macs                    string `json:"macs"`                                              // Space separated list of MACs (message authentication code) algorithms, ordered by preference.
}

type createSftpStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type siaConfig struct {
	ApiUrl      string `json:"apiUrl" default:"http://127.0.0.1:9980"`                                 // Sia daemon API URL, like http://sia.daemon.host:9980.
	ApiPassword string `json:"apiPassword"`                                                            // Sia Daemon API Password.
	UserAgent   string `json:"userAgent" default:"Sia-Agent"`                                          // Siad User Agent
	Encoding    string `json:"encoding" default:"Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createSiaStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Host             string `json:"host"`                                                                                                                        // SMB server hostname to connect to.
	User             string `json:"user" default:"$USER"`                                                                                                        // SMB username.
	Port             int    `json:"port" default:"445"`                                                                                                          // SMB port number.
	Pass             string `json:"pass"`                                                                                                                        // SMB password.
	Domain           string `json:"domain" default:"WORKGROUP"`                                                                                                  // Domain name for NTLM authentication.
	Spn              string `json:"spn"`                                                                                                                         // Service principal name.
	IdleTimeout      string `json:"idleTimeout" default:"1m0s"`                                                                                                  // Max time before closing idle connections.
	HideSpecialShare bool   `json:"hideSpecialShare" default:"true"`                                                                                             // Hide special shares (e.g. print$) which users aren't supposed to access.
	CaseInsensitive  bool   `json:"caseInsensitive" default:"true"`                                                                                              // Whether the server is configured to be case-insensitive.
	Encoding         string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createSmbStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type sugarsyncConfig struct {
	AppId               string `json:"appId"`                                        // Sugarsync App ID.
	AccessKeyId         string `json:"accessKeyId"`                                  // Sugarsync Access Key ID.
	PrivateAccessKey    string `json:"privateAccessKey"`                             // Sugarsync Private Access Key.
	HardDelete          bool   `json:"hardDelete" default:"false"`                   // Permanently delete files if true
	RefreshToken        string `json:"refreshToken"`                                 // Sugarsync refresh token.
	Authorization       string `json:"authorization"`                                // Sugarsync authorization.
	AuthorizationExpiry string `json:"authorizationExpiry"`                          // Sugarsync authorization expiry.
	User                string `json:"user"`                                         // Sugarsync user.
	RootId              string `json:"rootId"`                                       // Sugarsync root id.
	DeletedId           string `json:"deletedId"`                                    // Sugarsync deleted folder id.
	Encoding            string `json:"encoding" default:"Slash,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createSugarsyncStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type uptoboxConfig struct {
	AccessToken string `json:"accessToken"`                                                                           // Your access token.
	Encoding    string `json:"encoding" default:"Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createUptoboxStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Url                string `json:"url"`                        // URL of http host to connect to.
	Vendor             string `json:"vendor" example:"nextcloud"` // Name of the WebDAV site/service/software you are using.
	User               string `json:"user"`                       // User name.
	Pass               string `json:"pass"`                       // Password.
	BearerToken        string `json:"bearerToken"`                // Bearer token instead of user/pass (e.g. a Macaroon).
	BearerTokenCommand string `json:"bearerTokenCommand"`         // Command to run to get a bearer token.
	Encoding           string `json:"encoding"`                   // The encoding for the backend.
	Headers            string `json:"headers"`                    // Set HTTP headers for all transactions.
}

type createWebdavStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type filefabricConfig struct {
	Url            string `json:"url" example:"https://storagemadeeasy.com"`        // URL of the Enterprise File Fabric to connect to.
	RootFolderId   string `json:"rootFolderId"`                                     // ID of the root folder.
	PermanentToken string `json:"permanentToken"`                                   // Permanent Authentication Token.
	Token          string `json:"token"`                                            // Session Token.
	TokenExpiry    string `json:"tokenExpiry"`                                      // Token expiry time.
	Version        string `json:"version"`                                          // Version read from the file fabric.
	Encoding       string `json:"encoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createFilefabricStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
type oosEnv_authConfig struct {
	Namespace            string `json:"namespace"`                                         // Object storage namespace
	Compartment          string `json:"compartment"`                                       // Object storage compartment OCID
	Region               string `json:"region"`                                            // Object storage Region
	Endpoint             string `json:"endpoint"`                                          // Endpoint for Object storage API.
	StorageTier          string `json:"storageTier" default:"Standard" example:"Standard"` // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	UploadCutoff         string `json:"uploadCutoff" default:"200Mi"`                      // Cutoff for switching to chunked upload.
	ChunkSize            string `json:"chunkSize" default:"5Mi"`                           // Chunk size to use for uploading.
	UploadConcurrency    int    `json:"uploadConcurrency" default:"10"`                    // Concurrency for multipart uploads.
	CopyCutoff           string `json:"copyCutoff" default:"4.656Gi"`                      // Cutoff for switching to multipart copy.
	CopyTimeout          string `json:"copyTimeout" default:"1m0s"`                        // Timeout for copy.
	DisableChecksum      bool   `json:"disableChecksum" default:"false"`                   // Don't store MD5 checksum with object metadata.
	Encoding             string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`          // The encoding for the backend.
	LeavePartsOnError    bool   `json:"leavePartsOnError" default:"false"`                 // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	NoCheckBucket        bool   `json:"noCheckBucket" default:"false"`                     // If set, don't attempt to check the bucket exists or create it.
	SseCustomerKeyFile   string `json:"sseCustomerKeyFile" example:""`                     // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKey       string `json:"sseCustomerKey" example:""`                         // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKeySha256 string `json:"sseCustomerKeySha256" example:""`                   // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseKmsKeyId          string `json:"sseKmsKeyId" example:""`                            // if using using your own master key in vault, this header specifies the
	SseCustomerAlgorithm string `json:"sseCustomerAlgorithm" example:""`                   // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
}

type createOosEnv_authStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Namespace            string `json:"namespace"`                                         // Object storage namespace
	Compartment          string `json:"compartment"`                                       // Object storage compartment OCID
	Region               string `json:"region"`                                            // Object storage Region
	Endpoint             string `json:"endpoint"`                                          // Endpoint for Object storage API.
	StorageTier          string `json:"storageTier" default:"Standard" example:"Standard"` // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	UploadCutoff         string `json:"uploadCutoff" default:"200Mi"`                      // Cutoff for switching to chunked upload.
	ChunkSize            string `json:"chunkSize" default:"5Mi"`                           // Chunk size to use for uploading.
	UploadConcurrency    int    `json:"uploadConcurrency" default:"10"`                    // Concurrency for multipart uploads.
	CopyCutoff           string `json:"copyCutoff" default:"4.656Gi"`                      // Cutoff for switching to multipart copy.
	CopyTimeout          string `json:"copyTimeout" default:"1m0s"`                        // Timeout for copy.
	DisableChecksum      bool   `json:"disableChecksum" default:"false"`                   // Don't store MD5 checksum with object metadata.
	Encoding             string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`          // The encoding for the backend.
	LeavePartsOnError    bool   `json:"leavePartsOnError" default:"false"`                 // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	NoCheckBucket        bool   `json:"noCheckBucket" default:"false"`                     // If set, don't attempt to check the bucket exists or create it.
	SseCustomerKeyFile   string `json:"sseCustomerKeyFile" example:""`                     // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKey       string `json:"sseCustomerKey" example:""`                         // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKeySha256 string `json:"sseCustomerKeySha256" example:""`                   // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseKmsKeyId          string `json:"sseKmsKeyId" example:""`                            // if using using your own master key in vault, this header specifies the
	SseCustomerAlgorithm string `json:"sseCustomerAlgorithm" example:""`                   // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
}

type createOosInstance_principal_authStorageRequest struct {
	Name         string                           `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Namespace            string `json:"namespace"`                                         // Object storage namespace
	Region               string `json:"region"`                                            // Object storage Region
	Endpoint             string `json:"endpoint"`                                          // Endpoint for Object storage API.
	StorageTier          string `json:"storageTier" default:"Standard" example:"Standard"` // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	UploadCutoff         string `json:"uploadCutoff" default:"200Mi"`                      // Cutoff for switching to chunked upload.
	ChunkSize            string `json:"chunkSize" default:"5Mi"`                           // Chunk size to use for uploading.
	UploadConcurrency    int    `json:"uploadConcurrency" default:"10"`                    // Concurrency for multipart uploads.
	CopyCutoff           string `json:"copyCutoff" default:"4.656Gi"`                      // Cutoff for switching to multipart copy.
	CopyTimeout          string `json:"copyTimeout" default:"1m0s"`                        // Timeout for copy.
	DisableChecksum      bool   `json:"disableChecksum" default:"false"`                   // Don't store MD5 checksum with object metadata.
	Encoding             string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`          // The encoding for the backend.
	LeavePartsOnError    bool   `json:"leavePartsOnError" default:"false"`                 // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	NoCheckBucket        bool   `json:"noCheckBucket" default:"false"`                     // If set, don't attempt to check the bucket exists or create it.
	SseCustomerKeyFile   string `json:"sseCustomerKeyFile" example:""`                     // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKey       string `json:"sseCustomerKey" example:""`                         // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKeySha256 string `json:"sseCustomerKeySha256" example:""`                   // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseKmsKeyId          string `json:"sseKmsKeyId" example:""`                            // if using using your own master key in vault, this header specifies the
	SseCustomerAlgorithm string `json:"sseCustomerAlgorithm" example:""`                   // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
}

type createOosNo_authStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Namespace            string `json:"namespace"`                                         // Object storage namespace
	Compartment          string `json:"compartment"`                                       // Object storage compartment OCID
	Region               string `json:"region"`                                            // Object storage Region
	Endpoint             string `json:"endpoint"`                                          // Endpoint for Object storage API.
	StorageTier          string `json:"storageTier" default:"Standard" example:"Standard"` // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	UploadCutoff         string `json:"uploadCutoff" default:"200Mi"`                      // Cutoff for switching to chunked upload.
	ChunkSize            string `json:"chunkSize" default:"5Mi"`                           // Chunk size to use for uploading.
	UploadConcurrency    int    `json:"uploadConcurrency" default:"10"`                    // Concurrency for multipart uploads.
	CopyCutoff           string `json:"copyCutoff" default:"4.656Gi"`                      // Cutoff for switching to multipart copy.
	CopyTimeout          string `json:"copyTimeout" default:"1m0s"`                        // Timeout for copy.
	DisableChecksum      bool   `json:"disableChecksum" default:"false"`                   // Don't store MD5 checksum with object metadata.
	Encoding             string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`          // The encoding for the backend.
	LeavePartsOnError    bool   `json:"leavePartsOnError" default:"false"`                 // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	NoCheckBucket        bool   `json:"noCheckBucket" default:"false"`                     // If set, don't attempt to check the bucket exists or create it.
	SseCustomerKeyFile   string `json:"sseCustomerKeyFile" example:""`                     // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKey       string `json:"sseCustomerKey" example:""`                         // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKeySha256 string `json:"sseCustomerKeySha256" example:""`                   // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseKmsKeyId          string `json:"sseKmsKeyId" example:""`                            // if using using your own master key in vault, this header specifies the
	SseCustomerAlgorithm string `json:"sseCustomerAlgorithm" example:""`                   // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
}

type createOosResource_principal_authStorageRequest struct {
	Name         string                           `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Namespace            string `json:"namespace"`                                                  // Object storage namespace
	Compartment          string `json:"compartment"`                                                // Object storage compartment OCID
	Region               string `json:"region"`                                                     // Object storage Region
	Endpoint             string `json:"endpoint"`                                                   // Endpoint for Object storage API.
	ConfigFile           string `json:"configFile" default:"~/.oci/config" example:"~/.oci/config"` // Path to OCI config file
	ConfigProfile        string `json:"configProfile" default:"Default" example:"Default"`          // Profile name inside the oci config file
	StorageTier          string `json:"storageTier" default:"Standard" example:"Standard"`          // The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
	UploadCutoff         string `json:"uploadCutoff" default:"200Mi"`                               // Cutoff for switching to chunked upload.
	ChunkSize            string `json:"chunkSize" default:"5Mi"`                                    // Chunk size to use for uploading.
	UploadConcurrency    int    `json:"uploadConcurrency" default:"10"`                             // Concurrency for multipart uploads.
	CopyCutoff           string `json:"copyCutoff" default:"4.656Gi"`                               // Cutoff for switching to multipart copy.
	CopyTimeout          string `json:"copyTimeout" default:"1m0s"`                                 // Timeout for copy.
	DisableChecksum      bool   `json:"disableChecksum" default:"false"`                            // Don't store MD5 checksum with object metadata.
	Encoding             string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`                   // The encoding for the backend.
	LeavePartsOnError    bool   `json:"leavePartsOnError" default:"false"`                          // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	NoCheckBucket        bool   `json:"noCheckBucket" default:"false"`                              // If set, don't attempt to check the bucket exists or create it.
	SseCustomerKeyFile   string `json:"sseCustomerKeyFile" example:""`                              // To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
	SseCustomerKey       string `json:"sseCustomerKey" example:""`                                  // To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
	SseCustomerKeySha256 string `json:"sseCustomerKeySha256" example:""`                            // If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
	SseKmsKeyId          string `json:"sseKmsKeyId" example:""`                                     // if using using your own master key in vault, this header specifies the
	SseCustomerAlgorithm string `json:"sseCustomerAlgorithm" example:""`                            // If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
}

type createOosUser_principal_authStorageRequest struct {
	Name         string                       `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Username  string `json:"username"`                                                                                                                                         // Username.
	Password  string `json:"password"`                                                                                                                                         // Password.
	Encoding  string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot"` // The encoding for the backend.
	ChunkSize string `json:"chunkSize" default:"10Mi"`                                                                                                                         // Files will be uploaded in chunks this size.
}

type createOpendriveStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	ClientId     string `json:"clientId"`                                                   // OAuth Client Id.
	ClientSecret string `json:"clientSecret"`                                               // OAuth Client Secret.
	Token        string `json:"token"`                                                      // OAuth Access Token as a JSON blob.
	AuthUrl      string `json:"authUrl"`                                                    // Auth server URL.
	TokenUrl     string `json:"tokenUrl"`                                                   // Token server url.
	Encoding     string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	RootFolderId string `json:"rootFolderId" default:"d0"`                                  // Fill in for rclone to use a non root folder as its starting point.
	Hostname     string `json:"hostname" default:"api.pcloud.com" example:"api.pcloud.com"` // Hostname to connect to.
	Username     string `json:"username"`                                                   // Your pcloud username.
	Password     string `json:"password"`                                                   // Your pcloud password.
}

type createPcloudStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Encoding string `json:"encoding" default:"Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createPremiumizemeStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Encoding string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createPutioStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth           bool   `json:"envAuth" default:"false" example:"false"`  // Get QingStor credentials from runtime.
	AccessKeyId       string `json:"accessKeyId"`                              // QingStor Access Key ID.
	SecretAccessKey   string `json:"secretAccessKey"`                          // QingStor Secret Access Key (password).
	Endpoint          string `json:"endpoint"`                                 // Enter an endpoint URL to connection QingStor API.
	Zone              string `json:"zone" example:"pek3a"`                     // Zone to connect to.
	ConnectionRetries int    `json:"connectionRetries" default:"3"`            // Number of connection retries.
	UploadCutoff      string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize         string `json:"chunkSize" default:"4Mi"`                  // Chunk size to use for uploading.
	UploadConcurrency int    `json:"uploadConcurrency" default:"1"`            // Concurrency for multipart uploads.
	Encoding          string `json:"encoding" default:"Slash,Ctl,InvalidUtf8"` // The encoding for the backend.
}

type createQingstorStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Region                string `json:"region" example:"us-east-1"`               // Region to connect to.
	Endpoint              string `json:"endpoint"`                                 // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint" example:""`            // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`              // Canned ACL used when creating buckets.
	RequesterPays         bool   `json:"requesterPays" default:"false"`            // Enables requester pays option when interacting with S3 bucket.
	ServerSideEncryption  string `json:"serverSideEncryption" example:""`          // The server-side encryption algorithm used when storing this object in S3.
	SseCustomerAlgorithm  string `json:"sseCustomerAlgorithm" example:""`          // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	SseKmsKeyId           string `json:"sseKmsKeyId" example:""`                   // If using KMS ID you must provide the ARN of Key.
	SseCustomerKey        string `json:"sseCustomerKey" example:""`                // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	SseCustomerKeyBase64  string `json:"sseCustomerKeyBase64" example:""`          // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	SseCustomerKeyMd5     string `json:"sseCustomerKeyMd5" example:""`             // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	StorageClass          string `json:"storageClass" example:""`                  // The storage class to use when storing new objects in S3.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	UseAccelerateEndpoint bool   `json:"useAccelerateEndpoint" default:"false"`    // If true use the AWS S3 accelerated endpoint.
	LeavePartsOnError     bool   `json:"leavePartsOnError" default:"false"`        // If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
	ListChunk             int    `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                 // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
	StsEndpoint           string `json:"stsEndpoint"`                              // Endpoint for STS.
}

type createS3AWSStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`        // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                    // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                // AWS Secret Access Key (password).
	Endpoint              string `json:"endpoint" example:"oss-accelerate.aliyuncs.com"` // Endpoint for OSS API.
	Acl                   string `json:"acl"`                                            // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`                    // Canned ACL used when creating buckets.
	StorageClass          string `json:"storageClass" example:""`                        // The storage class to use when storing new objects in OSS.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`                   // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                        // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`                 // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`                   // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`                // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                          // Path to the shared credentials file.
	Profile               string `json:"profile"`                                        // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                   // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`                  // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                         // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                       // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`                  // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`                  // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                         // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`                   // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`       // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`             // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`              // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`                   // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                    // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`               // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`            // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                       // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                        // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                     // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                      // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`               // Suppress setting and reading of system metadata
}

type createS3AlibabaStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`           // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                       // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                   // AWS Secret Access Key (password).
	Endpoint              string `json:"endpoint" example:"s3.ir-thr-at1.arvanstorage.com"` // Endpoint for Arvan Cloud Object Storage (AOS) API.
	LocationConstraint    string `json:"locationConstraint" example:"ir-thr-at1"`           // Location constraint - must match endpoint.
	Acl                   string `json:"acl"`                                               // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`                       // Canned ACL used when creating buckets.
	StorageClass          string `json:"storageClass" example:"STANDARD"`                   // The storage class to use when storing new objects in ArvanCloud.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`                      // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                           // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`                    // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`                      // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`                   // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                             // Path to the shared credentials file.
	Profile               string `json:"profile"`                                           // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                      // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`                     // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                            // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                           // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`                     // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`                     // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                            // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`                      // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`          // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`                 // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`                      // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                       // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`                  // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`               // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                          // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                           // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                        // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                         // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`                  // Suppress setting and reading of system metadata
}

type createS3ArvanCloudStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Region                string `json:"region" example:""`                        // Region to connect to.
	Endpoint              string `json:"endpoint"`                                 // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                       // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`              // Canned ACL used when creating buckets.
	ServerSideEncryption  string `json:"serverSideEncryption" example:""`          // The server-side encryption algorithm used when storing this object in S3.
	SseCustomerAlgorithm  string `json:"sseCustomerAlgorithm" example:""`          // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	SseKmsKeyId           string `json:"sseKmsKeyId" example:""`                   // If using KMS ID you must provide the ARN of Key.
	SseCustomerKey        string `json:"sseCustomerKey" example:""`                // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	SseCustomerKeyBase64  string `json:"sseCustomerKeyBase64" example:""`          // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	SseCustomerKeyMd5     string `json:"sseCustomerKeyMd5" example:""`             // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                 // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
}

type createS3CephStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`   // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                               // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                           // AWS Secret Access Key (password).
	Endpoint              string `json:"endpoint" example:"eos-wuxi-1.cmecloud.cn"` // Endpoint for China Mobile Ecloud Elastic Object Storage (EOS) API.
	LocationConstraint    string `json:"locationConstraint" example:"wuxi1"`        // Location constraint - must match endpoint.
	Acl                   string `json:"acl"`                                       // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`               // Canned ACL used when creating buckets.
	ServerSideEncryption  string `json:"serverSideEncryption" example:""`           // The server-side encryption algorithm used when storing this object in S3.
	SseCustomerAlgorithm  string `json:"sseCustomerAlgorithm" example:""`           // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	SseCustomerKey        string `json:"sseCustomerKey" example:""`                 // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	SseCustomerKeyBase64  string `json:"sseCustomerKeyBase64" example:""`           // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	SseCustomerKeyMd5     string `json:"sseCustomerKeyMd5" example:""`              // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	StorageClass          string `json:"storageClass" example:""`                   // The storage class to use when storing new objects in ChinaMobile.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`              // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                   // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`            // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`              // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`           // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                     // Path to the shared credentials file.
	Profile               string `json:"profile"`                                   // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                              // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`             // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`             // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                    // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                  // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                   // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`             // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`             // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                    // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`              // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`  // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`        // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`         // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`              // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                               // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`          // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`       // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                  // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                   // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                 // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`          // Suppress setting and reading of system metadata
}

type createS3ChinaMobileStorageRequest struct {
	Name         string              `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Region                string `json:"region" example:"auto"`                    // Region to connect to.
	Endpoint              string `json:"endpoint"`                                 // Endpoint for S3 API.
	BucketAcl             string `json:"bucketAcl" example:"private"`              // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                 // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
}

type createS3CloudflareStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`        // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                    // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                // AWS Secret Access Key (password).
	Region                string `json:"region" example:""`                              // Region to connect to.
	Endpoint              string `json:"endpoint" example:"syd1.digitaloceanspaces.com"` // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                             // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                            // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`                    // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`                   // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                        // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`                 // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`                   // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`                // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                          // Path to the shared credentials file.
	Profile               string `json:"profile"`                                        // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                   // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`                  // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                         // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                       // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`                  // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`                  // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                         // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`                   // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`       // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`             // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`              // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`                   // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                    // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`               // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`            // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                       // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                        // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                     // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                      // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`               // Suppress setting and reading of system metadata
}

type createS3DigitalOceanStorageRequest struct {
	Name         string               `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`       // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                   // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                               // AWS Secret Access Key (password).
	Region                string `json:"region" example:""`                             // Region to connect to.
	Endpoint              string `json:"endpoint" example:"objects-us-east-1.dream.io"` // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                            // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                           // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`                   // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`                  // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                       // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`                // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`                  // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`               // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                         // Path to the shared credentials file.
	Profile               string `json:"profile"`                                       // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                  // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`                 // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`                 // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                        // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                      // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                       // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`                 // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`                 // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                        // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`                  // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`      // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`            // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`             // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`                  // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                   // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`              // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`           // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                      // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                       // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                    // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                     // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`              // Suppress setting and reading of system metadata
}

type createS3DreamhostStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`             // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                         // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                     // AWS Secret Access Key (password).
	Region                string `json:"region" example:"af-south-1"`                         // Region to connect to. - the location where your bucket will be created and your data stored. Need bo be same with your endpoint.
	Endpoint              string `json:"endpoint" example:"obs.af-south-1.myhuaweicloud.com"` // Endpoint for OBS API.
	Acl                   string `json:"acl"`                                                 // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`                         // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`                        // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                             // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`                      // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`                        // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`                     // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                               // Path to the shared credentials file.
	Profile               string `json:"profile"`                                             // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                        // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`                       // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`                       // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                              // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                            // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                             // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`                       // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`                       // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                              // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`                        // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`            // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`                  // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`                   // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`                        // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                         // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`                    // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`                 // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                            // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                             // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                          // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                           // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`                    // Suppress setting and reading of system metadata
}

type createS3HuaweiOBSStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`                       // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                                   // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                               // AWS Secret Access Key (password).
	Region                string `json:"region" example:""`                                             // Region to connect to.
	Endpoint              string `json:"endpoint" example:"s3.us.cloud-object-storage.appdomain.cloud"` // Endpoint for IBM COS S3 API.
	LocationConstraint    string `json:"locationConstraint" example:"us-standard"`                      // Location constraint - must match endpoint when using IBM Cloud Public.
	Acl                   string `json:"acl" example:"private"`                                         // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`                                   // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`                                  // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                                       // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`                                // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`                                  // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`                               // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                         // Path to the shared credentials file.
	Profile               string `json:"profile"`                                                       // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                                  // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`                                 // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`                                 // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                                        // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                                      // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                                       // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`                                 // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`                                 // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                                        // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`                                  // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`                      // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`                            // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`                             // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`                                  // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                                   // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`                              // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`                           // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                                      // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                                       // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                                    // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                                     // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`                              // Suppress setting and reading of system metadata
}

type createS3IBMCOSStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Acl                   string `json:"acl"`                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`              // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                 // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
}

type createS3IDriveStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`           // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                       // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                   // AWS Secret Access Key (password).
	Region                string `json:"region" example:"de"`                               // Region where your bucket will be created and your data stored.
	Endpoint              string `json:"endpoint" example:"s3-eu-central-1.ionoscloud.com"` // Endpoint for IONOS S3 Object Storage.
	Acl                   string `json:"acl"`                                               // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`                       // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`                      // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                           // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`                    // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`                      // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`                   // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                             // Path to the shared credentials file.
	Profile               string `json:"profile"`                                           // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                      // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`                     // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`                     // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                            // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                          // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                           // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`                     // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`                     // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                            // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`                      // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`          // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`                // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`                 // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`                      // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                       // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`                  // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`               // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                          // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                           // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                        // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                         // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`                  // Suppress setting and reading of system metadata
}

type createS3IONOSStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`     // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                 // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                             // AWS Secret Access Key (password).
	Endpoint              string `json:"endpoint" example:"storage.iran.liara.space"` // Endpoint for Liara Object Storage API.
	Acl                   string `json:"acl"`                                         // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`                 // Canned ACL used when creating buckets.
	StorageClass          string `json:"storageClass" example:"STANDARD"`             // The storage class to use when storing new objects in Liara
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`                // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                     // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`              // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`                // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`             // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                       // Path to the shared credentials file.
	Profile               string `json:"profile"`                                     // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`               // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`               // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                      // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                    // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                     // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`               // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`               // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                      // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`                // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`    // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`          // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`           // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`                // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                 // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`            // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`         // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                    // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                     // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                  // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                   // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`            // Suppress setting and reading of system metadata
}

type createS3LiaraStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`               // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                           // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                       // AWS Secret Access Key (password).
	Region                string `json:"region" example:""`                                     // Region to connect to.
	Endpoint              string `json:"endpoint" example:"s3.us-east-1.lyvecloud.seagate.com"` // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                                    // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                                   // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`                           // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`                          // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                               // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`                        // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`                          // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`                       // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                 // Path to the shared credentials file.
	Profile               string `json:"profile"`                                               // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                          // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`                         // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`                         // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                                // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                              // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                               // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`                         // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`                         // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                                // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`                          // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`              // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`                    // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`                     // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`                          // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                           // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`                      // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`                   // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                              // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                               // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                            // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                             // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`                      // Suppress setting and reading of system metadata
}

type createS3LyveCloudStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Region                string `json:"region" example:""`                        // Region to connect to.
	Endpoint              string `json:"endpoint"`                                 // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                       // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`              // Canned ACL used when creating buckets.
	ServerSideEncryption  string `json:"serverSideEncryption" example:""`          // The server-side encryption algorithm used when storing this object in S3.
	SseCustomerAlgorithm  string `json:"sseCustomerAlgorithm" example:""`          // If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
	SseKmsKeyId           string `json:"sseKmsKeyId" example:""`                   // If using KMS ID you must provide the ARN of Key.
	SseCustomerKey        string `json:"sseCustomerKey" example:""`                // To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
	SseCustomerKeyBase64  string `json:"sseCustomerKeyBase64" example:""`          // If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
	SseCustomerKeyMd5     string `json:"sseCustomerKeyMd5" example:""`             // If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                 // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
}

type createS3MinioStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Region                string `json:"region" example:""`                        // Region to connect to.
	Endpoint              string `json:"endpoint"`                                 // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                       // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`              // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                 // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
}

type createS3NeteaseStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Region                string `json:"region" example:""`                        // Region to connect to.
	Endpoint              string `json:"endpoint"`                                 // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                       // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`              // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                 // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
}

type createS3OtherStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`     // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                 // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                             // AWS Secret Access Key (password).
	Region                string `json:"region" example:"cn-east-1"`                  // Region to connect to.
	Endpoint              string `json:"endpoint" example:"s3-cn-east-1.qiniucs.com"` // Endpoint for Qiniu Object Storage.
	LocationConstraint    string `json:"locationConstraint" example:"cn-east-1"`      // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                         // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`                 // Canned ACL used when creating buckets.
	StorageClass          string `json:"storageClass" example:"STANDARD"`             // The storage class to use when storing new objects in Qiniu.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`                // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                     // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`              // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`                // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`             // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                       // Path to the shared credentials file.
	Profile               string `json:"profile"`                                     // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`               // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`               // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                      // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                    // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                     // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`               // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`               // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                      // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`                // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`    // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`          // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`           // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`                // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                 // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`            // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`         // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                    // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                     // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                  // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                   // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`            // Suppress setting and reading of system metadata
}

type createS3QiniuStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Region                string `json:"region" example:"global"`                  // region - the location where your bucket will be created and your data stored.
	Endpoint              string `json:"endpoint" example:"s3.rackcorp.com"`       // Endpoint for RackCorp Object Storage.
	LocationConstraint    string `json:"locationConstraint" example:"global"`      // Location constraint - the location where your bucket will be located and your data stored.
	Acl                   string `json:"acl"`                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`              // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                 // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
}

type createS3RackCorpStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Region                string `json:"region" example:"nl-ams"`                  // Region to connect to.
	Endpoint              string `json:"endpoint" example:"s3.nl-ams.scw.cloud"`   // Endpoint for Scaleway Object Storage.
	Acl                   string `json:"acl"`                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`              // Canned ACL used when creating buckets.
	StorageClass          string `json:"storageClass" example:""`                  // The storage class to use when storing new objects in S3.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                 // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
}

type createS3ScalewayStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Region                string `json:"region" example:""`                        // Region to connect to.
	Endpoint              string `json:"endpoint" example:"localhost:8333"`        // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                       // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`              // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                 // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
}

type createS3SeaweedFSStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`              // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                          // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                      // AWS Secret Access Key (password).
	Region                string `json:"region" example:""`                                    // Region to connect to.
	Endpoint              string `json:"endpoint" example:"s3.us-east-2.stackpathstorage.com"` // Endpoint for StackPath Object Storage.
	Acl                   string `json:"acl"`                                                  // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`                          // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`                         // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                              // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`                       // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`                         // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`                      // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                                // Path to the shared credentials file.
	Profile               string `json:"profile"`                                              // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                         // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`                        // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`                        // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                               // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                             // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                              // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`                        // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`                        // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                               // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`                         // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`             // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`                   // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`                    // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`                         // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                          // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`                     // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`                  // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                             // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                              // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                           // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                            // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`                     // Suppress setting and reading of system metadata
}

type createS3StackPathStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Endpoint              string `json:"endpoint" example:"gateway.storjshare.io"` // Endpoint for Storj Gateway.
	BucketAcl             string `json:"bucketAcl" example:"private"`              // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                 // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
}

type createS3StorjStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`        // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                                    // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                                // AWS Secret Access Key (password).
	Endpoint              string `json:"endpoint" example:"cos.ap-beijing.myqcloud.com"` // Endpoint for Tencent COS API.
	Acl                   string `json:"acl" example:"default"`                          // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`                    // Canned ACL used when creating buckets.
	StorageClass          string `json:"storageClass" example:""`                        // The storage class to use when storing new objects in Tencent COS.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`                   // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                        // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`                 // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`                   // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`                // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                          // Path to the shared credentials file.
	Profile               string `json:"profile"`                                        // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                                   // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`                  // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`                  // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                         // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                       // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                        // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`                  // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`                  // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                         // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`                   // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"`       // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`             // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`              // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`                   // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                                    // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`               // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`            // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                       // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                        // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`                     // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                      // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`               // Suppress setting and reading of system metadata
}

type createS3TencentCOSStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	EnvAuth               bool   `json:"envAuth" default:"false" example:"false"`  // Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
	AccessKeyId           string `json:"accessKeyId"`                              // AWS Access Key ID.
	SecretAccessKey       string `json:"secretAccessKey"`                          // AWS Secret Access Key (password).
	Region                string `json:"region" example:""`                        // Region to connect to.
	Endpoint              string `json:"endpoint" example:"s3.wasabisys.com"`      // Endpoint for S3 API.
	LocationConstraint    string `json:"locationConstraint"`                       // Location constraint - must be set to match the Region.
	Acl                   string `json:"acl"`                                      // Canned ACL used when creating buckets and storing or copying objects.
	BucketAcl             string `json:"bucketAcl" example:"private"`              // Canned ACL used when creating buckets.
	UploadCutoff          string `json:"uploadCutoff" default:"200Mi"`             // Cutoff for switching to chunked upload.
	ChunkSize             string `json:"chunkSize" default:"5Mi"`                  // Chunk size to use for uploading.
	MaxUploadParts        int    `json:"maxUploadParts" default:"10000"`           // Maximum number of parts in a multipart upload.
	CopyCutoff            string `json:"copyCutoff" default:"4.656Gi"`             // Cutoff for switching to multipart copy.
	DisableChecksum       bool   `json:"disableChecksum" default:"false"`          // Don't store MD5 checksum with object metadata.
	SharedCredentialsFile string `json:"sharedCredentialsFile"`                    // Path to the shared credentials file.
	Profile               string `json:"profile"`                                  // Profile to use in the shared credentials file.
	SessionToken          string `json:"sessionToken"`                             // An AWS session token.
	UploadConcurrency     int    `json:"uploadConcurrency" default:"4"`            // Concurrency for multipart uploads.
	ForcePathStyle        bool   `json:"forcePathStyle" default:"true"`            // If true use path style access if false use virtual hosted style.
	V2Auth                bool   `json:"v2Auth" default:"false"`                   // If true use v2 authentication.
	ListChunk             int    `json:"listChunk" default:"1000"`                 // Size of listing chunk (response list for each ListObject S3 request).
	ListVersion           int    `json:"listVersion" default:"0"`                  // Version of ListObjects to use: 1,2 or 0 for auto.
	ListUrlEncode         string `json:"listUrlEncode" default:"unset"`            // Whether to url encode listings: true/false/unset
	NoCheckBucket         bool   `json:"noCheckBucket" default:"false"`            // If set, don't attempt to check the bucket exists or create it.
	NoHead                bool   `json:"noHead" default:"false"`                   // If set, don't HEAD uploaded objects to check integrity.
	NoHeadObject          bool   `json:"noHeadObject" default:"false"`             // If set, do not do HEAD before GET when getting objects.
	Encoding              string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
	MemoryPoolFlushTime   string `json:"memoryPoolFlushTime" default:"1m0s"`       // How often internal memory buffer pools will be flushed.
	MemoryPoolUseMmap     bool   `json:"memoryPoolUseMmap" default:"false"`        // Whether to use mmap buffers in internal memory pool.
	DisableHttp2          bool   `json:"disableHttp2" default:"false"`             // Disable usage of http2 for S3 backends.
	DownloadUrl           string `json:"downloadUrl"`                              // Custom endpoint for downloads.
	UseMultipartEtag      string `json:"useMultipartEtag" default:"unset"`         // Whether to use ETag in multipart uploads for verification
	UsePresignedRequest   bool   `json:"usePresignedRequest" default:"false"`      // Whether to use a presigned request or PutObject for single part uploads
	Versions              bool   `json:"versions" default:"false"`                 // Include old versions in directory listings.
	VersionAt             string `json:"versionAt" default:"off"`                  // Show file versions as they were at the specified time.
	Decompress            bool   `json:"decompress" default:"false"`               // If set this will decompress gzip encoded objects.
	MightGzip             string `json:"mightGzip" default:"unset"`                // Set this if the backend might gzip objects.
	NoSystemMetadata      bool   `json:"noSystemMetadata" default:"false"`         // Suppress setting and reading of system metadata
}

type createS3WasabiStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type swiftConfig struct {
	EnvAuth                     bool   `json:"envAuth" default:"false" example:"false"`                 // Get swift credentials from environment variables in standard OpenStack form.
	User                        string `json:"user"`                                                    // User name to log in (OS_USERNAME).
	Key                         string `json:"key"`                                                     // API key or password (OS_PASSWORD).
	Auth                        string `json:"auth" example:"https://auth.api.rackspacecloud.com/v1.0"` // Authentication URL for server (OS_AUTH_URL).
	UserId                      string `json:"userId"`                                                  // User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID).
	Domain                      string `json:"domain"`                                                  // User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME)
	Tenant                      string `json:"tenant"`                                                  // Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME).
	TenantId                    string `json:"tenantId"`                                                // Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID).
	TenantDomain                string `json:"tenantDomain"`                                            // Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME).
	Region                      string `json:"region"`                                                  // Region name - optional (OS_REGION_NAME).
	StorageUrl                  string `json:"storageUrl"`                                              // Storage URL - optional (OS_STORAGE_URL).
	AuthToken                   string `json:"authToken"`                                               // Auth Token from alternate authentication - optional (OS_AUTH_TOKEN).
	ApplicationCredentialId     string `json:"applicationCredentialId"`                                 // Application Credential ID (OS_APPLICATION_CREDENTIAL_ID).
	ApplicationCredentialName   string `json:"applicationCredentialName"`                               // Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME).
	ApplicationCredentialSecret string `json:"applicationCredentialSecret"`                             // Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET).
	AuthVersion                 int    `json:"authVersion" default:"0"`                                 // AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).
	EndpointType                string `json:"endpointType" default:"public" example:"public"`          // Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).
	LeavePartsOnError           bool   `json:"leavePartsOnError" default:"false"`                       // If true avoid calling abort upload on a failure.
	StoragePolicy               string `json:"storagePolicy" example:""`                                // The storage policy to use when creating a new container.
	ChunkSize                   string `json:"chunkSize" default:"5Gi"`                                 // Above this size files will be chunked into a _segments container.
	NoChunk                     bool   `json:"noChunk" default:"false"`                                 // Don't chunk files during streaming upload.
	NoLargeObjects              bool   `json:"noLargeObjects" default:"false"`                          // Disable support for static and dynamic large objects
	Encoding                    string `json:"encoding" default:"Slash,InvalidUtf8"`                    // The encoding for the backend.
}

type createSwiftStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type acdConfig struct {
	ClientId          string `json:"clientId"`                                 // OAuth Client Id.
	ClientSecret      string `json:"clientSecret"`                             // OAuth Client Secret.
	Token             string `json:"token"`                                    // OAuth Access Token as a JSON blob.
	AuthUrl           string `json:"authUrl"`                                  // Auth server URL.
	TokenUrl          string `json:"tokenUrl"`                                 // Token server url.
	Checkpoint        string `json:"checkpoint"`                               // Checkpoint for internal polling (debug).
	UploadWaitPerGb   string `json:"uploadWaitPerGb" default:"3m0s"`           // Additional time per GiB to wait after a failed complete upload to see if it appears.
	TemplinkThreshold string `json:"templinkThreshold" default:"9Gi"`          // Files >= this size will be downloaded via their tempLink.
	Encoding          string `json:"encoding" default:"Slash,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createAcdStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type dropboxConfig struct {
	ClientId           string `json:"clientId"`                                                          // OAuth Client Id.
	ClientSecret       string `json:"clientSecret"`                                                      // OAuth Client Secret.
	Token              string `json:"token"`                                                             // OAuth Access Token as a JSON blob.
	AuthUrl            string `json:"authUrl"`                                                           // Auth server URL.
	TokenUrl           string `json:"tokenUrl"`                                                          // Token server url.
	ChunkSize          string `json:"chunkSize" default:"48Mi"`                                          // Upload chunk size (< 150Mi).
	Impersonate        string `json:"impersonate"`                                                       // Impersonate this user when using a business account.
	SharedFiles        bool   `json:"sharedFiles" default:"false"`                                       // Instructs rclone to work on individual shared files.
	SharedFolders      bool   `json:"sharedFolders" default:"false"`                                     // Instructs rclone to work on shared folders.
	BatchMode          string `json:"batchMode" default:"sync"`                                          // Upload file batching sync|async|off.
	BatchSize          int    `json:"batchSize" default:"0"`                                             // Max number of files in upload batch.
	BatchTimeout       string `json:"batchTimeout" default:"0s"`                                         // Max time to allow an idle upload batch before uploading.
	BatchCommitTimeout string `json:"batchCommitTimeout" default:"10m0s"`                                // Max time to wait for a batch to finish committing
	Encoding           string `json:"encoding" default:"Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createDropboxStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type gphotosConfig struct {
	ClientId        string `json:"clientId"`                                      // OAuth Client Id.
	ClientSecret    string `json:"clientSecret"`                                  // OAuth Client Secret.
	Token           string `json:"token"`                                         // OAuth Access Token as a JSON blob.
	AuthUrl         string `json:"authUrl"`                                       // Auth server URL.
	TokenUrl        string `json:"tokenUrl"`                                      // Token server url.
	ReadOnly        bool   `json:"readOnly" default:"false"`                      // Set to make the Google Photos backend read only.
	ReadSize        bool   `json:"readSize" default:"false"`                      // Set to read the size of media items.
	StartYear       int    `json:"startYear" default:"2000"`                      // Year limits the photos to be downloaded to those which are uploaded after the given year.
	IncludeArchived bool   `json:"includeArchived" default:"false"`               // Also view and download archived media.
	Encoding        string `json:"encoding" default:"Slash,CrLf,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createGphotosStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type hidriveConfig struct {
	ClientId                   string `json:"clientId"`                                              // OAuth Client Id.
	ClientSecret               string `json:"clientSecret"`                                          // OAuth Client Secret.
	Token                      string `json:"token"`                                                 // OAuth Access Token as a JSON blob.
	AuthUrl                    string `json:"authUrl"`                                               // Auth server URL.
	TokenUrl                   string `json:"tokenUrl"`                                              // Token server url.
	ScopeAccess                string `json:"scopeAccess" default:"rw" example:"rw"`                 // Access permissions that rclone should use when requesting access from HiDrive.
	ScopeRole                  string `json:"scopeRole" default:"user" example:"user"`               // User-level that rclone should use when requesting access from HiDrive.
	RootPrefix                 string `json:"rootPrefix" default:"/" example:"/"`                    // The root/parent folder for all paths.
	Endpoint                   string `json:"endpoint" default:"https://api.hidrive.strato.com/2.1"` // Endpoint for the service.
	DisableFetchingMemberCount bool   `json:"disableFetchingMemberCount" default:"false"`            // Do not fetch number of objects in directories unless it is absolutely necessary.
	ChunkSize                  string `json:"chunkSize" default:"48Mi"`                              // Chunksize for chunked uploads.
	UploadCutoff               string `json:"uploadCutoff" default:"96Mi"`                           // Cutoff/Threshold for chunked uploads.
	UploadConcurrency          int    `json:"uploadConcurrency" default:"4"`                         // Concurrency for chunked uploads.
	Encoding                   string `json:"encoding" default:"Slash,Dot"`                          // The encoding for the backend.
}

type createHidriveStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type jottacloudConfig struct {
	Md5MemoryLimit    string `json:"md5MemoryLimit" default:"10Mi"`                                                                  // Files bigger than this will be cached on disk to calculate the MD5 if required.
	TrashedOnly       bool   `json:"trashedOnly" default:"false"`                                                                    // Only show files that are in the trash.
	HardDelete        bool   `json:"hardDelete" default:"false"`                                                                     // Delete files permanently rather than putting them into the trash.
	UploadResumeLimit string `json:"uploadResumeLimit" default:"10Mi"`                                                               // Files bigger than this can be resumed if the upload fail's.
	NoVersions        bool   `json:"noVersions" default:"false"`                                                                     // Avoid server side versioning by deleting files and recreating files instead of overwriting them.
	Encoding          string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createJottacloudStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type mailruConfig struct {
	User                string `json:"user"`                                                                                                     // User name (usually email).
	Pass                string `json:"pass"`                                                                                                     // Password.
	SpeedupEnable       bool   `json:"speedupEnable" default:"true" example:"true"`                                                              // Skip full upload if there is another file with same data hash.
	SpeedupFilePatterns string `json:"speedupFilePatterns" default:"*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf" example:""`                  // Comma separated list of file name patterns eligible for speedup (put by hash).
	SpeedupMaxDisk      string `json:"speedupMaxDisk" default:"3Gi" example:"0"`                                                                 // This option allows you to disable speedup (put by hash) for large files.
	SpeedupMaxMemory    string `json:"speedupMaxMemory" default:"32Mi" example:"0"`                                                              // Files larger than the size given below will always be hashed on disk.
	CheckHash           bool   `json:"checkHash" default:"true" example:"true"`                                                                  // What should copy do if file checksum is mismatched or invalid.
	UserAgent           string `json:"userAgent"`                                                                                                // HTTP user agent used internally by client.
	Quirks              string `json:"quirks"`                                                                                                   // Comma separated list of internal maintenance flags.
	Encoding            string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createMailruStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type onedriveConfig struct {
	ClientId                string `json:"clientId"`                                                                                                                                                                                                                 // OAuth Client Id.
	ClientSecret            string `json:"clientSecret"`                                                                                                                                                                                                             // OAuth Client Secret.
	Token                   string `json:"token"`                                                                                                                                                                                                                    // OAuth Access Token as a JSON blob.
	AuthUrl                 string `json:"authUrl"`                                                                                                                                                                                                                  // Auth server URL.
	TokenUrl                string `json:"tokenUrl"`                                                                                                                                                                                                                 // Token server url.
	Region                  string `json:"region" default:"global" example:"global"`                                                                                                                                                                                 // Choose national cloud region for OneDrive.
	ChunkSize               string `json:"chunkSize" default:"10Mi"`                                                                                                                                                                                                 // Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
	DriveId                 string `json:"driveId"`                                                                                                                                                                                                                  // The ID of the drive to use.
	DriveType               string `json:"driveType"`                                                                                                                                                                                                                // The type of the drive (personal | business | documentLibrary).
	RootFolderId            string `json:"rootFolderId"`                                                                                                                                                                                                             // ID of the root folder.
	AccessScopes            string `json:"accessScopes" default:"Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access" example:"Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"` // Set scopes to be requested by rclone.
	DisableSitePermission   bool   `json:"disableSitePermission" default:"false"`                                                                                                                                                                                    // Disable the request for Sites.Read.All permission.
	ExposeOnenoteFiles      bool   `json:"exposeOnenoteFiles" default:"false"`                                                                                                                                                                                       // Set to make OneNote files show up in directory listings.
	ServerSideAcrossConfigs bool   `json:"serverSideAcrossConfigs" default:"false"`                                                                                                                                                                                  // Allow server-side operations (e.g. copy) to work across different onedrive configs.
	ListChunk               int    `json:"listChunk" default:"1000"`                                                                                                                                                                                                 // Size of listing chunk.
	NoVersions              bool   `json:"noVersions" default:"false"`                                                                                                                                                                                               // Remove all versions on modifying operations.
	LinkScope               string `json:"linkScope" default:"anonymous" example:"anonymous"`                                                                                                                                                                        // Set the scope of the links created by the link command.
	LinkType                string `json:"linkType" default:"view" example:"view"`                                                                                                                                                                                   // Set the type of the links created by the link command.
	LinkPassword            string `json:"linkPassword"`                                                                                                                                                                                                             // Set the password for links created by the link command.
	HashType                string `json:"hashType" default:"auto" example:"auto"`                                                                                                                                                                                   // Specify the hash in use for the backend.
	Encoding                string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot"`                                                                      // The encoding for the backend.
}

type createOnedriveStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type pcloudConfig struct {
	ClientId     string `json:"clientId"`                                                   // OAuth Client Id.
	ClientSecret string `json:"clientSecret"`                                               // OAuth Client Secret.
	Token        string `json:"token"`                                                      // OAuth Access Token as a JSON blob.
	AuthUrl      string `json:"authUrl"`                                                    // Auth server URL.
	TokenUrl     string `json:"tokenUrl"`                                                   // Token server url.
	Encoding     string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
	RootFolderId string `json:"rootFolderId" default:"d0"`                                  // Fill in for rclone to use a non root folder as its starting point.
	Hostname     string `json:"hostname" default:"api.pcloud.com" example:"api.pcloud.com"` // Hostname to connect to.
	Username     string `json:"username"`                                                   // Your pcloud username.
	Password     string `json:"password"`                                                   // Your pcloud password.
}

type createPcloudStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Encoding string `json:"encoding" default:"Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createPremiumizemeStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Encoding string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createPutioStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type sharefileConfig struct {
	UploadCutoff string `json:"uploadCutoff" default:"128Mi"`                                                                                                                     // Cutoff for switching to multipart upload.
	RootFolderId string `json:"rootFolderId" example:""`                                                                                                                          // ID of the root folder.
	ChunkSize    string `json:"chunkSize" default:"64Mi"`                                                                                                                         // Upload chunk size.
	Endpoint     string `json:"endpoint"`                                                                                                                                         // Endpoint for API calls.
	Encoding     string `json:"encoding" default:"Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createSharefileStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type yandexConfig struct {
	ClientId     string `json:"clientId"`                                         // OAuth Client Id.
	ClientSecret string `json:"clientSecret"`                                     // OAuth Client Secret.
	Token        string `json:"token"`                                            // OAuth Access Token as a JSON blob.
	AuthUrl      string `json:"authUrl"`                                          // Auth server URL.
	TokenUrl     string `json:"tokenUrl"`                                         // Token server url.
	HardDelete   bool   `json:"hardDelete" default:"false"`                       // Delete files permanently rather than putting them into the trash.
	Encoding     string `json:"encoding" default:"Slash,Del,Ctl,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createYandexStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	Region       string `json:"region" example:"com"`                   // Zoho region to connect to.
	Encoding     string `json:"encoding" default:"Del,Ctl,InvalidUtf8"` // The encoding for the backend.
}

type createZohoStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type boxConfig struct {
	ClientId      string `json:"clientId"`                                                              // OAuth Client Id.
	ClientSecret  string `json:"clientSecret"`                                                          // OAuth Client Secret.
	Token         string `json:"token"`                                                                 // OAuth Access Token as a JSON blob.
	AuthUrl       string `json:"authUrl"`                                                               // Auth server URL.
	TokenUrl      string `json:"tokenUrl"`                                                              // Token server url.
	RootFolderId  string `json:"rootFolderId" default:"0"`                                              // Fill in for rclone to use a non root folder as its starting point.
	BoxConfigFile string `json:"boxConfigFile"`                                                         // Box App config.json location
	AccessToken   string `json:"accessToken"`                                                           // Box App Primary Access Token
	BoxSubType    string `json:"boxSubType" default:"user" example:"user"`                              //
	UploadCutoff  string `json:"uploadCutoff" default:"50Mi"`                                           // Cutoff for switching to multipart upload (>= 50 MiB).
	CommitRetries int    `json:"commitRetries" default:"100"`                                           // Max number of times to try committing a multipart file.
	ListChunk     int    `json:"listChunk" default:"1000"`                                              // Size of listing chunk 1-1000.
	OwnedBy       string `json:"ownedBy"`                                                               // Only show items owned by the login (email address) passed in.
	Encoding      string `json:"encoding" default:"Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"` // The encoding for the backend.
}

type createBoxStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
	ClientId                  string `json:"clientId"`                                   // Google Application Client Id
	ClientSecret              string `json:"clientSecret"`                               // OAuth Client Secret.
	Token                     string `json:"token"`                                      // OAuth Access Token as a JSON blob.
	AuthUrl                   string `json:"authUrl"`                                    // Auth server URL.
	TokenUrl                  string `json:"tokenUrl"`                                   // Token server url.
	Scope                     string `json:"scope" example:"drive"`                      // Scope that rclone should use when requesting access from drive.
	RootFolderId              string `json:"rootFolderId"`                               // ID of the root folder.
	ServiceAccountFile        string `json:"serviceAccountFile"`                         // Service Account Credentials JSON file path.
	ServiceAccountCredentials string `json:"serviceAccountCredentials"`                  // Service Account Credentials JSON blob.
	TeamDrive                 string `json:"teamDrive"`                                  // ID of the Shared Drive (Team Drive).
	AuthOwnerOnly             bool   `json:"authOwnerOnly" default:"false"`              // Only consider files owned by the authenticated user.
	UseTrash                  bool   `json:"useTrash" default:"true"`                    // Send files to the trash instead of deleting permanently.
	CopyShortcutContent       bool   `json:"copyShortcutContent" default:"false"`        // Server side copy contents of shortcuts instead of the shortcut.
	SkipGdocs                 bool   `json:"skipGdocs" default:"false"`                  // Skip google documents in all listings.
	SkipChecksumGphotos       bool   `json:"skipChecksumGphotos" default:"false"`        // Skip MD5 checksum on Google photos and videos only.
	SharedWithMe              bool   `json:"sharedWithMe" default:"false"`               // Only show files that are shared with me.
	TrashedOnly               bool   `json:"trashedOnly" default:"false"`                // Only show files that are in the trash.
	StarredOnly               bool   `json:"starredOnly" default:"false"`                // Only show files that are starred.
	Formats                   string `json:"formats"`                                    // Deprecated: See export_formats.
	ExportFormats             string `json:"exportFormats" default:"docx,xlsx,pptx,svg"` // Comma separated list of preferred formats for downloading Google docs.
	ImportFormats             string `json:"importFormats"`                              // Comma separated list of preferred formats for uploading Google docs.
	AllowImportNameChange     bool   `json:"allowImportNameChange" default:"false"`      // Allow the filetype to change when uploading Google docs.
	UseCreatedDate            bool   `json:"useCreatedDate" default:"false"`             // Use file created date instead of modified date.
	UseSharedDate             bool   `json:"useSharedDate" default:"false"`              // Use date file was shared instead of modified date.
	ListChunk                 int    `json:"listChunk" default:"1000"`                   // Size of listing chunk 100-1000, 0 to disable.
	Impersonate               string `json:"impersonate"`                                // Impersonate this user when using a service account.
	AlternateExport           bool   `json:"alternateExport" default:"false"`            // Deprecated: No longer needed.
	UploadCutoff              string `json:"uploadCutoff" default:"8Mi"`                 // Cutoff for switching to chunked upload.
	ChunkSize                 string `json:"chunkSize" default:"8Mi"`                    // Upload chunk size.
	AcknowledgeAbuse          bool   `json:"acknowledgeAbuse" default:"false"`           // Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
	KeepRevisionForever       bool   `json:"keepRevisionForever" default:"false"`        // Keep new head revision of each file forever.
	SizeAsQuota               bool   `json:"sizeAsQuota" default:"false"`                // Show sizes as storage quota usage, not actual size.
	V2DownloadMinSize         string `json:"v2DownloadMinSize" default:"off"`            // If Object's are greater, use drive v2 API to download.
	PacerMinSleep             string `json:"pacerMinSleep" default:"100ms"`              // Minimum time to sleep between API calls.
	PacerBurst                int    `json:"pacerBurst" default:"100"`                   // Number of API calls to allow without sleeping.
	ServerSideAcrossConfigs   bool   `json:"serverSideAcrossConfigs" default:"false"`    // Allow server-side operations (e.g. copy) to work across different drive configs.
	DisableHttp2              bool   `json:"disableHttp2" default:"true"`                // Disable drive using http2.
	StopOnUploadLimit         bool   `json:"stopOnUploadLimit" default:"false"`          // Make upload limit errors be fatal.
	StopOnDownloadLimit       bool   `json:"stopOnDownloadLimit" default:"false"`        // Make download limit errors be fatal.
	SkipShortcuts             bool   `json:"skipShortcuts" default:"false"`              // If set skip shortcut files.
	SkipDanglingShortcuts     bool   `json:"skipDanglingShortcuts" default:"false"`      // If set skip dangling shortcut files.
	ResourceKey               string `json:"resourceKey"`                                // Resource key for accessing a link-shared file.
	Encoding                  string `json:"encoding" default:"InvalidUtf8"`             // The encoding for the backend.
}

type createDriveStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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

type gcsConfig struct {
	ClientId                  string `json:"clientId"`                                      // OAuth Client Id.
	ClientSecret              string `json:"clientSecret"`                                  // OAuth Client Secret.
	Token                     string `json:"token"`                                         // OAuth Access Token as a JSON blob.
	AuthUrl                   string `json:"authUrl"`                                       // Auth server URL.
	TokenUrl                  string `json:"tokenUrl"`                                      // Token server url.
	ProjectNumber             string `json:"projectNumber"`                                 // Project number.
	ServiceAccountFile        string `json:"serviceAccountFile"`                            // Service Account Credentials JSON file path.
	ServiceAccountCredentials string `json:"serviceAccountCredentials"`                     // Service Account Credentials JSON blob.
	Anonymous                 bool   `json:"anonymous" default:"false"`                     // Access public buckets and objects without credentials.
	ObjectAcl                 string `json:"objectAcl" example:"authenticatedRead"`         // Access Control List for new objects.
	BucketAcl                 string `json:"bucketAcl" example:"authenticatedRead"`         // Access Control List for new buckets.
	BucketPolicyOnly          bool   `json:"bucketPolicyOnly" default:"false"`              // Access checks should use bucket-level IAM policies.
	Location                  string `json:"location" example:""`                           // Location for the newly created buckets.
	StorageClass              string `json:"storageClass" example:""`                       // The storage class to use when storing objects in Google Cloud Storage.
	NoCheckBucket             bool   `json:"noCheckBucket" default:"false"`                 // If set, don't attempt to check the bucket exists or create it.
	Decompress                bool   `json:"decompress" default:"false"`                    // If set this will decompress gzip encoded objects.
	Endpoint                  string `json:"endpoint"`                                      // Endpoint for the service.
	Encoding                  string `json:"encoding" default:"Slash,CrLf,InvalidUtf8,Dot"` // The encoding for the backend.
	EnvAuth                   bool   `json:"envAuth" default:"false" example:"false"`       // Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
}

type createGcsStorageRequest struct {
	Name         string             `json:"name" example:"my-storage"` // Name of the storage, must be unique
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
