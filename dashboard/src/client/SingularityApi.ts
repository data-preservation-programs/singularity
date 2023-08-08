/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ApiHTTPError {
  err?: string;
}

export interface DatasetAddPieceRequest {
  /** Path to the CAR file, used to determine the size of the file and root CID */
  filePath?: string;
  /** CID of the piece */
  pieceCid?: string;
  /** Size of the piece */
  pieceSize?: string;
  /** Root CID of the CAR file, if not provided, will be determined by the CAR file header. Used to populate the label field of storage deal */
  rootCid?: string;
}

export interface DatasetCreateRequest {
  /** Public key of the encryption recipient */
  encryptionRecipients?: string[];
  /** EncryptionScript command to run for custom encryption */
  encryptionScript?: string;
  /**
   * Maximum size of the CAR files to be created
   * @default "31.5GiB"
   */
  maxSize: string;
  /** Name must be a unique identifier for a dataset */
  name: string;
  /** Output directory for CAR files. Do not set if using inline preparation */
  outputDirs?: string[];
  /** Target piece size of the CAR files used for piece commitment calculation */
  pieceSize?: string;
}

export interface DatasetUpdateRequest {
  /** Public key of the encryption recipient */
  encryptionRecipients?: string[];
  /** EncryptionScript command to run for custom encryption */
  encryptionScript?: string;
  /**
   * Maximum size of the CAR files to be created
   * @default "31.5GiB"
   */
  maxSize?: string;
  /** Output directory for CAR files. Do not set if using inline preparation */
  outputDirs?: string[];
  /** Target piece size of the CAR files used for piece commitment calculation */
  pieceSize?: string;
}

export interface DatasourceAcdRequest {
  /** Auth server URL. */
  authUrl?: string;
  /** Checkpoint for internal polling (debug). */
  checkpoint?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Files >= this size will be downloaded via their tempLink.
   * @default "9Gi"
   */
  templinkThreshold?: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
  /**
   * Additional time per GiB to wait after a failed complete upload to see if it appears.
   * @default "3m0s"
   */
  uploadWaitPerGb?: string;
}

export interface DatasourceAllConfig {
  /** Auth server URL. */
  acdAuthUrl?: string;
  /** Checkpoint for internal polling (debug). */
  acdCheckpoint?: string;
  /** OAuth Client Id. */
  acdClientId?: string;
  /** OAuth Client Secret. */
  acdClientSecret?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  acdEncoding?: string;
  /**
   * Files >= this size will be downloaded via their tempLink.
   * @default "9Gi"
   */
  acdTemplinkThreshold?: string;
  /** OAuth Access Token as a JSON blob. */
  acdToken?: string;
  /** Token server url. */
  acdTokenUrl?: string;
  /**
   * Additional time per GiB to wait after a failed complete upload to see if it appears.
   * @default "3m0s"
   */
  acdUploadWaitPerGb?: string;
  /** Access tier of blob: hot, cool or archive. */
  azureblobAccessTier?: string;
  /** Azure Storage Account Name. */
  azureblobAccount?: string;
  /**
   * Delete archive tier blobs before overwriting.
   * @default "false"
   */
  azureblobArchiveTierDelete?: string;
  /**
   * Upload chunk size.
   * @default "4Mi"
   */
  azureblobChunkSize?: string;
  /** Password for the certificate file (optional). */
  azureblobClientCertificatePassword?: string;
  /** Path to a PEM or PKCS12 certificate file including the private key. */
  azureblobClientCertificatePath?: string;
  /** The ID of the client in use. */
  azureblobClientId?: string;
  /** One of the service principal's client secrets */
  azureblobClientSecret?: string;
  /**
   * Send the certificate chain when using certificate auth.
   * @default "false"
   */
  azureblobClientSendCertificateChain?: string;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default "false"
   */
  azureblobDisableChecksum?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"
   */
  azureblobEncoding?: string;
  /** Endpoint for the service. */
  azureblobEndpoint?: string;
  /**
   * Read credentials from runtime (environment variables, CLI or MSI).
   * @default "false"
   */
  azureblobEnvAuth?: string;
  /** Storage Account Shared Key. */
  azureblobKey?: string;
  /**
   * Size of blob list.
   * @default "5000"
   */
  azureblobListChunk?: string;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  azureblobMemoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default "false"
   */
  azureblobMemoryPoolUseMmap?: string;
  /** Object ID of the user-assigned MSI to use, if any. */
  azureblobMsiClientId?: string;
  /** Azure resource ID of the user-assigned MSI to use, if any. */
  azureblobMsiMiResId?: string;
  /** Object ID of the user-assigned MSI to use, if any. */
  azureblobMsiObjectId?: string;
  /**
   * If set, don't attempt to check the container exists or create it.
   * @default "false"
   */
  azureblobNoCheckContainer?: string;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default "false"
   */
  azureblobNoHeadObject?: string;
  /** The user's password */
  azureblobPassword?: string;
  /** Public access level of a container: blob or container. */
  azureblobPublicAccess?: string;
  /** SAS URL for container level access only. */
  azureblobSasUrl?: string;
  /** Path to file containing credentials for use with a service principal. */
  azureblobServicePrincipalFile?: string;
  /** ID of the service principal's tenant. Also called its directory ID. */
  azureblobTenant?: string;
  /**
   * Concurrency for multipart uploads.
   * @default "16"
   */
  azureblobUploadConcurrency?: string;
  /** Cutoff for switching to chunked upload (<= 256 MiB) (deprecated). */
  azureblobUploadCutoff?: string;
  /**
   * Uses local storage emulator if provided as 'true'.
   * @default "false"
   */
  azureblobUseEmulator?: string;
  /**
   * Use a managed service identity to authenticate (only works in Azure).
   * @default "false"
   */
  azureblobUseMsi?: string;
  /** User name (usually an email address) */
  azureblobUsername?: string;
  /** Account ID or Application Key ID. */
  b2Account?: string;
  /**
   * Upload chunk size.
   * @default "96Mi"
   */
  b2ChunkSize?: string;
  /**
   * Cutoff for switching to multipart copy.
   * @default "4Gi"
   */
  b2CopyCutoff?: string;
  /**
   * Disable checksums for large (> upload cutoff) files.
   * @default "false"
   */
  b2DisableChecksum?: string;
  /**
   * Time before the authorization token will expire in s or suffix ms|s|m|h|d.
   * @default "1w"
   */
  b2DownloadAuthDuration?: string;
  /** Custom endpoint for downloads. */
  b2DownloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  b2Encoding?: string;
  /** Endpoint for the service. */
  b2Endpoint?: string;
  /**
   * Permanently delete files on remote removal, otherwise hide files.
   * @default "false"
   */
  b2HardDelete?: string;
  /** Application Key. */
  b2Key?: string;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  b2MemoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default "false"
   */
  b2MemoryPoolUseMmap?: string;
  /** A flag string for X-Bz-Test-Mode header for debugging. */
  b2TestMode?: string;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  b2UploadCutoff?: string;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  b2VersionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default "false"
   */
  b2Versions?: string;
  /** Box App Primary Access Token */
  boxAccessToken?: string;
  /** Auth server URL. */
  boxAuthUrl?: string;
  /** Box App config.json location */
  boxBoxConfigFile?: string;
  /** @default "user" */
  boxBoxSubType?: string;
  /** OAuth Client Id. */
  boxClientId?: string;
  /** OAuth Client Secret. */
  boxClientSecret?: string;
  /**
   * Max number of times to try committing a multipart file.
   * @default "100"
   */
  boxCommitRetries?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"
   */
  boxEncoding?: string;
  /**
   * Size of listing chunk 1-1000.
   * @default "1000"
   */
  boxListChunk?: string;
  /** Only show items owned by the login (email address) passed in. */
  boxOwnedBy?: string;
  /**
   * Fill in for rclone to use a non root folder as its starting point.
   * @default "0"
   */
  boxRootFolderId?: string;
  /** OAuth Access Token as a JSON blob. */
  boxToken?: string;
  /** Token server url. */
  boxTokenUrl?: string;
  /**
   * Cutoff for switching to multipart upload (>= 50 MiB).
   * @default "50Mi"
   */
  boxUploadCutoff?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport?: boolean;
  /**
   * Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
   * @default "false"
   */
  driveAcknowledgeAbuse?: string;
  /**
   * Allow the filetype to change when uploading Google docs.
   * @default "false"
   */
  driveAllowImportNameChange?: string;
  /**
   * Deprecated: No longer needed.
   * @default "false"
   */
  driveAlternateExport?: string;
  /**
   * Only consider files owned by the authenticated user.
   * @default "false"
   */
  driveAuthOwnerOnly?: string;
  /** Auth server URL. */
  driveAuthUrl?: string;
  /**
   * Upload chunk size.
   * @default "8Mi"
   */
  driveChunkSize?: string;
  /** Google Application Client Id */
  driveClientId?: string;
  /** OAuth Client Secret. */
  driveClientSecret?: string;
  /**
   * Server side copy contents of shortcuts instead of the shortcut.
   * @default "false"
   */
  driveCopyShortcutContent?: string;
  /**
   * Disable drive using http2.
   * @default "true"
   */
  driveDisableHttp2?: string;
  /**
   * The encoding for the backend.
   * @default "InvalidUtf8"
   */
  driveEncoding?: string;
  /**
   * Comma separated list of preferred formats for downloading Google docs.
   * @default "docx,xlsx,pptx,svg"
   */
  driveExportFormats?: string;
  /** Deprecated: See export_formats. */
  driveFormats?: string;
  /** Impersonate this user when using a service account. */
  driveImpersonate?: string;
  /** Comma separated list of preferred formats for uploading Google docs. */
  driveImportFormats?: string;
  /**
   * Keep new head revision of each file forever.
   * @default "false"
   */
  driveKeepRevisionForever?: string;
  /**
   * Size of listing chunk 100-1000, 0 to disable.
   * @default "1000"
   */
  driveListChunk?: string;
  /**
   * Number of API calls to allow without sleeping.
   * @default "100"
   */
  drivePacerBurst?: string;
  /**
   * Minimum time to sleep between API calls.
   * @default "100ms"
   */
  drivePacerMinSleep?: string;
  /** Resource key for accessing a link-shared file. */
  driveResourceKey?: string;
  /** ID of the root folder. */
  driveRootFolderId?: string;
  /** Scope that rclone should use when requesting access from drive. */
  driveScope?: string;
  /**
   * Allow server-side operations (e.g. copy) to work across different drive configs.
   * @default "false"
   */
  driveServerSideAcrossConfigs?: string;
  /** Service Account Credentials JSON blob. */
  driveServiceAccountCredentials?: string;
  /** Service Account Credentials JSON file path. */
  driveServiceAccountFile?: string;
  /**
   * Only show files that are shared with me.
   * @default "false"
   */
  driveSharedWithMe?: string;
  /**
   * Show sizes as storage quota usage, not actual size.
   * @default "false"
   */
  driveSizeAsQuota?: string;
  /**
   * Skip MD5 checksum on Google photos and videos only.
   * @default "false"
   */
  driveSkipChecksumGphotos?: string;
  /**
   * If set skip dangling shortcut files.
   * @default "false"
   */
  driveSkipDanglingShortcuts?: string;
  /**
   * Skip google documents in all listings.
   * @default "false"
   */
  driveSkipGdocs?: string;
  /**
   * If set skip shortcut files.
   * @default "false"
   */
  driveSkipShortcuts?: string;
  /**
   * Only show files that are starred.
   * @default "false"
   */
  driveStarredOnly?: string;
  /**
   * Make download limit errors be fatal.
   * @default "false"
   */
  driveStopOnDownloadLimit?: string;
  /**
   * Make upload limit errors be fatal.
   * @default "false"
   */
  driveStopOnUploadLimit?: string;
  /** ID of the Shared Drive (Team Drive). */
  driveTeamDrive?: string;
  /** OAuth Access Token as a JSON blob. */
  driveToken?: string;
  /** Token server url. */
  driveTokenUrl?: string;
  /**
   * Only show files that are in the trash.
   * @default "false"
   */
  driveTrashedOnly?: string;
  /**
   * Cutoff for switching to chunked upload.
   * @default "8Mi"
   */
  driveUploadCutoff?: string;
  /**
   * Use file created date instead of modified date.
   * @default "false"
   */
  driveUseCreatedDate?: string;
  /**
   * Use date file was shared instead of modified date.
   * @default "false"
   */
  driveUseSharedDate?: string;
  /**
   * Send files to the trash instead of deleting permanently.
   * @default "true"
   */
  driveUseTrash?: string;
  /**
   * If Object's are greater, use drive v2 API to download.
   * @default "off"
   */
  driveV2DownloadMinSize?: string;
  /** Auth server URL. */
  dropboxAuthUrl?: string;
  /**
   * Max time to wait for a batch to finish committing
   * @default "10m0s"
   */
  dropboxBatchCommitTimeout?: string;
  /**
   * Upload file batching sync|async|off.
   * @default "sync"
   */
  dropboxBatchMode?: string;
  /**
   * Max number of files in upload batch.
   * @default "0"
   */
  dropboxBatchSize?: string;
  /**
   * Max time to allow an idle upload batch before uploading.
   * @default "0s"
   */
  dropboxBatchTimeout?: string;
  /**
   * Upload chunk size (< 150Mi).
   * @default "48Mi"
   */
  dropboxChunkSize?: string;
  /** OAuth Client Id. */
  dropboxClientId?: string;
  /** OAuth Client Secret. */
  dropboxClientSecret?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"
   */
  dropboxEncoding?: string;
  /** Impersonate this user when using a business account. */
  dropboxImpersonate?: string;
  /**
   * Instructs rclone to work on individual shared files.
   * @default "false"
   */
  dropboxSharedFiles?: string;
  /**
   * Instructs rclone to work on shared folders.
   * @default "false"
   */
  dropboxSharedFolders?: string;
  /** OAuth Access Token as a JSON blob. */
  dropboxToken?: string;
  /** Token server url. */
  dropboxTokenUrl?: string;
  /** Your API Key, get it from https://1fichier.com/console/params.pl. */
  fichierApiKey?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"
   */
  fichierEncoding?: string;
  /** If you want to download a shared file that is password protected, add this parameter. */
  fichierFilePassword?: string;
  /** If you want to list the files in a shared folder that is password protected, add this parameter. */
  fichierFolderPassword?: string;
  /** If you want to download a shared folder, add this parameter. */
  fichierSharedFolder?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Del,Ctl,InvalidUtf8,Dot"
   */
  filefabricEncoding?: string;
  /** Permanent Authentication Token. */
  filefabricPermanentToken?: string;
  /** ID of the root folder. */
  filefabricRootFolderId?: string;
  /** Session Token. */
  filefabricToken?: string;
  /** Token expiry time. */
  filefabricTokenExpiry?: string;
  /** URL of the Enterprise File Fabric to connect to. */
  filefabricUrl?: string;
  /** Version read from the file fabric. */
  filefabricVersion?: string;
  /**
   * Allow asking for FTP password when needed.
   * @default "false"
   */
  ftpAskPassword?: string;
  /**
   * Maximum time to wait for a response to close.
   * @default "1m0s"
   */
  ftpCloseTimeout?: string;
  /**
   * Maximum number of FTP simultaneous connections, 0 for unlimited.
   * @default "0"
   */
  ftpConcurrency?: string;
  /**
   * Disable using EPSV even if server advertises support.
   * @default "false"
   */
  ftpDisableEpsv?: string;
  /**
   * Disable using MLSD even if server advertises support.
   * @default "false"
   */
  ftpDisableMlsd?: string;
  /**
   * Disable TLS 1.3 (workaround for FTP servers with buggy TLS)
   * @default "false"
   */
  ftpDisableTls13?: string;
  /**
   * Disable using UTF-8 even if server advertises support.
   * @default "false"
   */
  ftpDisableUtf8?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Del,Ctl,RightSpace,Dot"
   */
  ftpEncoding?: string;
  /**
   * Use Explicit FTPS (FTP over TLS).
   * @default "false"
   */
  ftpExplicitTls?: string;
  /**
   * Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.
   * @default "false"
   */
  ftpForceListHidden?: string;
  /** FTP host to connect to. */
  ftpHost?: string;
  /**
   * Max time before closing idle connections.
   * @default "1m0s"
   */
  ftpIdleTimeout?: string;
  /**
   * Do not verify the TLS certificate of the server.
   * @default "false"
   */
  ftpNoCheckCertificate?: string;
  /** FTP password. */
  ftpPass?: string;
  /**
   * FTP port number.
   * @default "21"
   */
  ftpPort?: string;
  /**
   * Maximum time to wait for data connection closing status.
   * @default "1m0s"
   */
  ftpShutTimeout?: string;
  /**
   * Use Implicit FTPS (FTP over TLS).
   * @default "false"
   */
  ftpTls?: string;
  /**
   * Size of TLS session cache for all control and data connections.
   * @default "32"
   */
  ftpTlsCacheSize?: string;
  /**
   * FTP username.
   * @default "$USER"
   */
  ftpUser?: string;
  /**
   * Use MDTM to set modification time (VsFtpd quirk)
   * @default "false"
   */
  ftpWritingMdtm?: string;
  /**
   * Access public buckets and objects without credentials.
   * @default "false"
   */
  gcsAnonymous?: string;
  /** Auth server URL. */
  gcsAuthUrl?: string;
  /** Access Control List for new buckets. */
  gcsBucketAcl?: string;
  /**
   * Access checks should use bucket-level IAM policies.
   * @default "false"
   */
  gcsBucketPolicyOnly?: string;
  /** OAuth Client Id. */
  gcsClientId?: string;
  /** OAuth Client Secret. */
  gcsClientSecret?: string;
  /**
   * If set this will decompress gzip encoded objects.
   * @default "false"
   */
  gcsDecompress?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,CrLf,InvalidUtf8,Dot"
   */
  gcsEncoding?: string;
  /** Endpoint for the service. */
  gcsEndpoint?: string;
  /**
   * Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
   * @default "false"
   */
  gcsEnvAuth?: string;
  /** Location for the newly created buckets. */
  gcsLocation?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default "false"
   */
  gcsNoCheckBucket?: string;
  /** Access Control List for new objects. */
  gcsObjectAcl?: string;
  /** Project number. */
  gcsProjectNumber?: string;
  /** Service Account Credentials JSON blob. */
  gcsServiceAccountCredentials?: string;
  /** Service Account Credentials JSON file path. */
  gcsServiceAccountFile?: string;
  /** The storage class to use when storing objects in Google Cloud Storage. */
  gcsStorageClass?: string;
  /** OAuth Access Token as a JSON blob. */
  gcsToken?: string;
  /** Token server url. */
  gcsTokenUrl?: string;
  /** Auth server URL. */
  gphotosAuthUrl?: string;
  /** OAuth Client Id. */
  gphotosClientId?: string;
  /** OAuth Client Secret. */
  gphotosClientSecret?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,CrLf,InvalidUtf8,Dot"
   */
  gphotosEncoding?: string;
  /**
   * Also view and download archived media.
   * @default "false"
   */
  gphotosIncludeArchived?: string;
  /**
   * Set to make the Google Photos backend read only.
   * @default "false"
   */
  gphotosReadOnly?: string;
  /**
   * Set to read the size of media items.
   * @default "false"
   */
  gphotosReadSize?: string;
  /**
   * Year limits the photos to be downloaded to those which are uploaded after the given year.
   * @default "2000"
   */
  gphotosStartYear?: string;
  /** OAuth Access Token as a JSON blob. */
  gphotosToken?: string;
  /** Token server url. */
  gphotosTokenUrl?: string;
  /** Kerberos data transfer protection: authentication|integrity|privacy. */
  hdfsDataTransferProtection?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Colon,Del,Ctl,InvalidUtf8,Dot"
   */
  hdfsEncoding?: string;
  /** Hadoop name node and port. */
  hdfsNamenode?: string;
  /** Kerberos service principal name for the namenode. */
  hdfsServicePrincipalName?: string;
  /** Hadoop user name. */
  hdfsUsername?: string;
  /** Auth server URL. */
  hidriveAuthUrl?: string;
  /**
   * Chunksize for chunked uploads.
   * @default "48Mi"
   */
  hidriveChunkSize?: string;
  /** OAuth Client Id. */
  hidriveClientId?: string;
  /** OAuth Client Secret. */
  hidriveClientSecret?: string;
  /**
   * Do not fetch number of objects in directories unless it is absolutely necessary.
   * @default "false"
   */
  hidriveDisableFetchingMemberCount?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Dot"
   */
  hidriveEncoding?: string;
  /**
   * Endpoint for the service.
   * @default "https://api.hidrive.strato.com/2.1"
   */
  hidriveEndpoint?: string;
  /**
   * The root/parent folder for all paths.
   * @default "/"
   */
  hidriveRootPrefix?: string;
  /**
   * Access permissions that rclone should use when requesting access from HiDrive.
   * @default "rw"
   */
  hidriveScopeAccess?: string;
  /**
   * User-level that rclone should use when requesting access from HiDrive.
   * @default "user"
   */
  hidriveScopeRole?: string;
  /** OAuth Access Token as a JSON blob. */
  hidriveToken?: string;
  /** Token server url. */
  hidriveTokenUrl?: string;
  /**
   * Concurrency for chunked uploads.
   * @default "4"
   */
  hidriveUploadConcurrency?: string;
  /**
   * Cutoff/Threshold for chunked uploads.
   * @default "96Mi"
   */
  hidriveUploadCutoff?: string;
  /** Set HTTP headers for all transactions. */
  httpHeaders?: string;
  /**
   * Don't use HEAD requests.
   * @default "false"
   */
  httpNoHead?: string;
  /**
   * Set this if the site doesn't end directories with /.
   * @default "false"
   */
  httpNoSlash?: string;
  /** URL of HTTP host to connect to. */
  httpUrl?: string;
  /** IAS3 Access Key. */
  internetarchiveAccessKeyId?: string;
  /**
   * Don't ask the server to test against MD5 checksum calculated by rclone.
   * @default "true"
   */
  internetarchiveDisableChecksum?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot"
   */
  internetarchiveEncoding?: string;
  /**
   * IAS3 Endpoint.
   * @default "https://s3.us.archive.org"
   */
  internetarchiveEndpoint?: string;
  /**
   * Host of InternetArchive Frontend.
   * @default "https://archive.org"
   */
  internetarchiveFrontEndpoint?: string;
  /** IAS3 Secret Key (password). */
  internetarchiveSecretAccessKey?: string;
  /**
   * Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
   * @default "0s"
   */
  internetarchiveWaitArchive?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot"
   */
  jottacloudEncoding?: string;
  /**
   * Delete files permanently rather than putting them into the trash.
   * @default "false"
   */
  jottacloudHardDelete?: string;
  /**
   * Files bigger than this will be cached on disk to calculate the MD5 if required.
   * @default "10Mi"
   */
  jottacloudMd5MemoryLimit?: string;
  /**
   * Avoid server side versioning by deleting files and recreating files instead of overwriting them.
   * @default "false"
   */
  jottacloudNoVersions?: string;
  /**
   * Only show files that are in the trash.
   * @default "false"
   */
  jottacloudTrashedOnly?: string;
  /**
   * Files bigger than this can be resumed if the upload fail's.
   * @default "10Mi"
   */
  jottacloudUploadResumeLimit?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  koofrEncoding?: string;
  /** The Koofr API endpoint to use. */
  koofrEndpoint?: string;
  /** Mount ID of the mount to use. */
  koofrMountid?: string;
  /** Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password). */
  koofrPassword?: string;
  /** Choose your storage provider. */
  koofrProvider?: string;
  /**
   * Does the backend support setting modification time.
   * @default "true"
   */
  koofrSetmtime?: string;
  /** Your user name. */
  koofrUser?: string;
  /**
   * Force the filesystem to report itself as case insensitive.
   * @default "false"
   */
  localCaseInsensitive?: string;
  /**
   * Force the filesystem to report itself as case sensitive.
   * @default "false"
   */
  localCaseSensitive?: string;
  /**
   * Follow symlinks and copy the pointed to item.
   * @default "false"
   */
  localCopyLinks?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Dot"
   */
  localEncoding?: string;
  /**
   * Translate symlinks to/from regular files with a '.rclonelink' extension.
   * @default "false"
   */
  localLinks?: string;
  /**
   * Don't check to see if the files change during upload.
   * @default "false"
   */
  localNoCheckUpdated?: string;
  /**
   * Disable preallocation of disk space for transferred files.
   * @default "false"
   */
  localNoPreallocate?: string;
  /**
   * Disable setting modtime.
   * @default "false"
   */
  localNoSetModtime?: string;
  /**
   * Disable sparse files for multi-thread downloads.
   * @default "false"
   */
  localNoSparse?: string;
  /**
   * Disable UNC (long path names) conversion on Windows.
   * @default "false"
   */
  localNounc?: string;
  /**
   * Don't cross filesystem boundaries (unix/macOS only).
   * @default "false"
   */
  localOneFileSystem?: string;
  /**
   * Don't warn about skipped symlinks.
   * @default "false"
   */
  localSkipLinks?: string;
  /**
   * Apply unicode NFC normalization to paths and filenames.
   * @default "false"
   */
  localUnicodeNormalization?: string;
  /**
   * Assume the Stat size of links is zero (and read them instead) (deprecated).
   * @default "false"
   */
  localZeroSizeLinks?: string;
  /**
   * What should copy do if file checksum is mismatched or invalid.
   * @default "true"
   */
  mailruCheckHash?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  mailruEncoding?: string;
  /** Password. */
  mailruPass?: string;
  /** Comma separated list of internal maintenance flags. */
  mailruQuirks?: string;
  /**
   * Skip full upload if there is another file with same data hash.
   * @default "true"
   */
  mailruSpeedupEnable?: string;
  /**
   * Comma separated list of file name patterns eligible for speedup (put by hash).
   * @default "*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf"
   */
  mailruSpeedupFilePatterns?: string;
  /**
   * This option allows you to disable speedup (put by hash) for large files.
   * @default "3Gi"
   */
  mailruSpeedupMaxDisk?: string;
  /**
   * Files larger than the size given below will always be hashed on disk.
   * @default "32Mi"
   */
  mailruSpeedupMaxMemory?: string;
  /** User name (usually email). */
  mailruUser?: string;
  /** HTTP user agent used internally by client. */
  mailruUserAgent?: string;
  /**
   * Output more debug from Mega.
   * @default "false"
   */
  megaDebug?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  megaEncoding?: string;
  /**
   * Delete files permanently rather than putting them into the trash.
   * @default "false"
   */
  megaHardDelete?: string;
  /** Password. */
  megaPass?: string;
  /**
   * Use HTTPS for transfers.
   * @default "false"
   */
  megaUseHttps?: string;
  /** User name. */
  megaUser?: string;
  /** Set the NetStorage account name */
  netstorageAccount?: string;
  /** Domain+path of NetStorage host to connect to. */
  netstorageHost?: string;
  /**
   * Select between HTTP or HTTPS protocol.
   * @default "https"
   */
  netstorageProtocol?: string;
  /** Set the NetStorage account secret/G2O key for authentication. */
  netstorageSecret?: string;
  /**
   * Set scopes to be requested by rclone.
   * @default "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"
   */
  onedriveAccessScopes?: string;
  /** Auth server URL. */
  onedriveAuthUrl?: string;
  /**
   * Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
   * @default "10Mi"
   */
  onedriveChunkSize?: string;
  /** OAuth Client Id. */
  onedriveClientId?: string;
  /** OAuth Client Secret. */
  onedriveClientSecret?: string;
  /**
   * Disable the request for Sites.Read.All permission.
   * @default "false"
   */
  onedriveDisableSitePermission?: string;
  /** The ID of the drive to use. */
  onedriveDriveId?: string;
  /** The type of the drive (personal | business | documentLibrary). */
  onedriveDriveType?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot"
   */
  onedriveEncoding?: string;
  /**
   * Set to make OneNote files show up in directory listings.
   * @default "false"
   */
  onedriveExposeOnenoteFiles?: string;
  /**
   * Specify the hash in use for the backend.
   * @default "auto"
   */
  onedriveHashType?: string;
  /** Set the password for links created by the link command. */
  onedriveLinkPassword?: string;
  /**
   * Set the scope of the links created by the link command.
   * @default "anonymous"
   */
  onedriveLinkScope?: string;
  /**
   * Set the type of the links created by the link command.
   * @default "view"
   */
  onedriveLinkType?: string;
  /**
   * Size of listing chunk.
   * @default "1000"
   */
  onedriveListChunk?: string;
  /**
   * Remove all versions on modifying operations.
   * @default "false"
   */
  onedriveNoVersions?: string;
  /**
   * Choose national cloud region for OneDrive.
   * @default "global"
   */
  onedriveRegion?: string;
  /** ID of the root folder. */
  onedriveRootFolderId?: string;
  /**
   * Allow server-side operations (e.g. copy) to work across different onedrive configs.
   * @default "false"
   */
  onedriveServerSideAcrossConfigs?: string;
  /** OAuth Access Token as a JSON blob. */
  onedriveToken?: string;
  /** Token server url. */
  onedriveTokenUrl?: string;
  /**
   * Chunk size to use for uploading.
   * @default "5Mi"
   */
  oosChunkSize?: string;
  /** Object storage compartment OCID */
  oosCompartment?: string;
  /**
   * Path to OCI config file
   * @default "~/.oci/config"
   */
  oosConfigFile?: string;
  /**
   * Profile name inside the oci config file
   * @default "Default"
   */
  oosConfigProfile?: string;
  /**
   * Cutoff for switching to multipart copy.
   * @default "4.656Gi"
   */
  oosCopyCutoff?: string;
  /**
   * Timeout for copy.
   * @default "1m0s"
   */
  oosCopyTimeout?: string;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default "false"
   */
  oosDisableChecksum?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  oosEncoding?: string;
  /** Endpoint for Object storage API. */
  oosEndpoint?: string;
  /**
   * If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
   * @default "false"
   */
  oosLeavePartsOnError?: string;
  /** Object storage namespace */
  oosNamespace?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default "false"
   */
  oosNoCheckBucket?: string;
  /**
   * Choose your Auth Provider
   * @default "env_auth"
   */
  oosProvider?: string;
  /** Object storage Region */
  oosRegion?: string;
  /** If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm. */
  oosSseCustomerAlgorithm?: string;
  /** To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to */
  oosSseCustomerKey?: string;
  /** To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated */
  oosSseCustomerKeyFile?: string;
  /** If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption */
  oosSseCustomerKeySha256?: string;
  /** if using using your own master key in vault, this header specifies the */
  oosSseKmsKeyId?: string;
  /**
   * The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
   * @default "Standard"
   */
  oosStorageTier?: string;
  /**
   * Concurrency for multipart uploads.
   * @default "10"
   */
  oosUploadConcurrency?: string;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  oosUploadCutoff?: string;
  /**
   * Files will be uploaded in chunks this size.
   * @default "10Mi"
   */
  opendriveChunkSize?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot"
   */
  opendriveEncoding?: string;
  /** Password. */
  opendrivePassword?: string;
  /** Username. */
  opendriveUsername?: string;
  /** Auth server URL. */
  pcloudAuthUrl?: string;
  /** OAuth Client Id. */
  pcloudClientId?: string;
  /** OAuth Client Secret. */
  pcloudClientSecret?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  pcloudEncoding?: string;
  /**
   * Hostname to connect to.
   * @default "api.pcloud.com"
   */
  pcloudHostname?: string;
  /** Your pcloud password. */
  pcloudPassword?: string;
  /**
   * Fill in for rclone to use a non root folder as its starting point.
   * @default "d0"
   */
  pcloudRootFolderId?: string;
  /** OAuth Access Token as a JSON blob. */
  pcloudToken?: string;
  /** Token server url. */
  pcloudTokenUrl?: string;
  /** Your pcloud username. */
  pcloudUsername?: string;
  /** API Key. */
  premiumizemeApiKey?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  premiumizemeEncoding?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  putioEncoding?: string;
  /** QingStor Access Key ID. */
  qingstorAccessKeyId?: string;
  /**
   * Chunk size to use for uploading.
   * @default "4Mi"
   */
  qingstorChunkSize?: string;
  /**
   * Number of connection retries.
   * @default "3"
   */
  qingstorConnectionRetries?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Ctl,InvalidUtf8"
   */
  qingstorEncoding?: string;
  /** Enter an endpoint URL to connection QingStor API. */
  qingstorEndpoint?: string;
  /**
   * Get QingStor credentials from runtime.
   * @default "false"
   */
  qingstorEnvAuth?: string;
  /** QingStor Secret Access Key (password). */
  qingstorSecretAccessKey?: string;
  /**
   * Concurrency for multipart uploads.
   * @default "1"
   */
  qingstorUploadConcurrency?: string;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  qingstorUploadCutoff?: string;
  /** Zone to connect to. */
  qingstorZone?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval?: string;
  /** AWS Access Key ID. */
  s3AccessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  s3Acl?: string;
  /** Canned ACL used when creating buckets. */
  s3BucketAcl?: string;
  /**
   * Chunk size to use for uploading.
   * @default "5Mi"
   */
  s3ChunkSize?: string;
  /**
   * Cutoff for switching to multipart copy.
   * @default "4.656Gi"
   */
  s3CopyCutoff?: string;
  /**
   * If set this will decompress gzip encoded objects.
   * @default "false"
   */
  s3Decompress?: string;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default "false"
   */
  s3DisableChecksum?: string;
  /**
   * Disable usage of http2 for S3 backends.
   * @default "false"
   */
  s3DisableHttp2?: string;
  /** Custom endpoint for downloads. */
  s3DownloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  s3Encoding?: string;
  /** Endpoint for S3 API. */
  s3Endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default "false"
   */
  s3EnvAuth?: string;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default "true"
   */
  s3ForcePathStyle?: string;
  /**
   * If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
   * @default "false"
   */
  s3LeavePartsOnError?: string;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default "1000"
   */
  s3ListChunk?: string;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  s3ListUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default "0"
   */
  s3ListVersion?: string;
  /** Location constraint - must be set to match the Region. */
  s3LocationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default "10000"
   */
  s3MaxUploadParts?: string;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  s3MemoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default "false"
   */
  s3MemoryPoolUseMmap?: string;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  s3MightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default "false"
   */
  s3NoCheckBucket?: string;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default "false"
   */
  s3NoHead?: string;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default "false"
   */
  s3NoHeadObject?: string;
  /**
   * Suppress setting and reading of system metadata
   * @default "false"
   */
  s3NoSystemMetadata?: string;
  /** Profile to use in the shared credentials file. */
  s3Profile?: string;
  /** Choose your S3 provider. */
  s3Provider?: string;
  /** Region to connect to. */
  s3Region?: string;
  /**
   * Enables requester pays option when interacting with S3 bucket.
   * @default "false"
   */
  s3RequesterPays?: string;
  /** AWS Secret Access Key (password). */
  s3SecretAccessKey?: string;
  /** The server-side encryption algorithm used when storing this object in S3. */
  s3ServerSideEncryption?: string;
  /** An AWS session token. */
  s3SessionToken?: string;
  /** Path to the shared credentials file. */
  s3SharedCredentialsFile?: string;
  /** If using SSE-C, the server-side encryption algorithm used when storing this object in S3. */
  s3SseCustomerAlgorithm?: string;
  /** To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data. */
  s3SseCustomerKey?: string;
  /** If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data. */
  s3SseCustomerKeyBase64?: string;
  /** If using SSE-C you may provide the secret encryption key MD5 checksum (optional). */
  s3SseCustomerKeyMd5?: string;
  /** If using KMS ID you must provide the ARN of Key. */
  s3SseKmsKeyId?: string;
  /** The storage class to use when storing new objects in S3. */
  s3StorageClass?: string;
  /** Endpoint for STS. */
  s3StsEndpoint?: string;
  /**
   * Concurrency for multipart uploads.
   * @default "4"
   */
  s3UploadConcurrency?: string;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  s3UploadCutoff?: string;
  /**
   * If true use the AWS S3 accelerated endpoint.
   * @default "false"
   */
  s3UseAccelerateEndpoint?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  s3UseMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default "false"
   */
  s3UsePresignedRequest?: string;
  /**
   * If true use v2 authentication.
   * @default "false"
   */
  s3V2Auth?: string;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  s3VersionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default "false"
   */
  s3Versions?: string;
  /** Starting state for scanning */
  scanningState?: ModelWorkState;
  /**
   * Two-factor authentication ('true' if the account has 2FA enabled).
   * @default "false"
   */
  seafile2fa?: string;
  /** Authentication token. */
  seafileAuthToken?: string;
  /**
   * Should rclone create a library if it doesn't exist.
   * @default "false"
   */
  seafileCreateLibrary?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8"
   */
  seafileEncoding?: string;
  /** Name of the library. */
  seafileLibrary?: string;
  /** Library password (for encrypted libraries only). */
  seafileLibraryKey?: string;
  /** Password. */
  seafilePass?: string;
  /** URL of seafile host to connect to. */
  seafileUrl?: string;
  /** User name (usually email address). */
  seafileUser?: string;
  /**
   * Allow asking for SFTP password when needed.
   * @default "false"
   */
  sftpAskPassword?: string;
  /**
   * Upload and download chunk size.
   * @default "32Ki"
   */
  sftpChunkSize?: string;
  /** Space separated list of ciphers to be used for session encryption, ordered by preference. */
  sftpCiphers?: string;
  /**
   * The maximum number of outstanding requests for one file
   * @default "64"
   */
  sftpConcurrency?: string;
  /**
   * If set don't use concurrent reads.
   * @default "false"
   */
  sftpDisableConcurrentReads?: string;
  /**
   * If set don't use concurrent writes.
   * @default "false"
   */
  sftpDisableConcurrentWrites?: string;
  /**
   * Disable the execution of SSH commands to determine if remote file hashing is available.
   * @default "false"
   */
  sftpDisableHashcheck?: string;
  /** SSH host to connect to. */
  sftpHost?: string;
  /**
   * Max time before closing idle connections.
   * @default "1m0s"
   */
  sftpIdleTimeout?: string;
  /** Space separated list of key exchange algorithms, ordered by preference. */
  sftpKeyExchange?: string;
  /** Path to PEM-encoded private key file. */
  sftpKeyFile?: string;
  /** The passphrase to decrypt the PEM-encoded private key file. */
  sftpKeyFilePass?: string;
  /** Raw PEM-encoded private key. */
  sftpKeyPem?: string;
  /**
   * When set forces the usage of the ssh-agent.
   * @default "false"
   */
  sftpKeyUseAgent?: string;
  /** Optional path to known_hosts file. */
  sftpKnownHostsFile?: string;
  /** Space separated list of MACs (message authentication code) algorithms, ordered by preference. */
  sftpMacs?: string;
  /** The command used to read md5 hashes. */
  sftpMd5sumCommand?: string;
  /** SSH password, leave blank to use ssh-agent. */
  sftpPass?: string;
  /** Override path used by SSH shell commands. */
  sftpPathOverride?: string;
  /**
   * SSH port number.
   * @default "22"
   */
  sftpPort?: string;
  /** Optional path to public key file. */
  sftpPubkeyFile?: string;
  /** Specifies the path or command to run a sftp server on the remote host. */
  sftpServerCommand?: string;
  /** Environment variables to pass to sftp and commands */
  sftpSetEnv?: string;
  /**
   * Set the modified time on the remote if set.
   * @default "true"
   */
  sftpSetModtime?: string;
  /** The command used to read sha1 hashes. */
  sftpSha1sumCommand?: string;
  /** The type of SSH shell on remote server, if any. */
  sftpShellType?: string;
  /**
   * Set to skip any symlinks and any other non regular files.
   * @default "false"
   */
  sftpSkipLinks?: string;
  /**
   * Specifies the SSH2 subsystem on the remote host.
   * @default "sftp"
   */
  sftpSubsystem?: string;
  /**
   * If set use fstat instead of stat.
   * @default "false"
   */
  sftpUseFstat?: string;
  /**
   * Enable the use of insecure ciphers and key exchange methods.
   * @default "false"
   */
  sftpUseInsecureCipher?: string;
  /**
   * SSH username.
   * @default "$USER"
   */
  sftpUser?: string;
  /**
   * Upload chunk size.
   * @default "64Mi"
   */
  sharefileChunkSize?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot"
   */
  sharefileEncoding?: string;
  /** Endpoint for API calls. */
  sharefileEndpoint?: string;
  /** ID of the root folder. */
  sharefileRootFolderId?: string;
  /**
   * Cutoff for switching to multipart upload.
   * @default "128Mi"
   */
  sharefileUploadCutoff?: string;
  /** Sia Daemon API Password. */
  siaApiPassword?: string;
  /**
   * Sia daemon API URL, like http://sia.daemon.host:9980.
   * @default "http://127.0.0.1:9980"
   */
  siaApiUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot"
   */
  siaEncoding?: string;
  /**
   * Siad User Agent
   * @default "Sia-Agent"
   */
  siaUserAgent?: string;
  /**
   * Whether the server is configured to be case-insensitive.
   * @default "true"
   */
  smbCaseInsensitive?: string;
  /**
   * Domain name for NTLM authentication.
   * @default "WORKGROUP"
   */
  smbDomain?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot"
   */
  smbEncoding?: string;
  /**
   * Hide special shares (e.g. print$) which users aren't supposed to access.
   * @default "true"
   */
  smbHideSpecialShare?: string;
  /** SMB server hostname to connect to. */
  smbHost?: string;
  /**
   * Max time before closing idle connections.
   * @default "1m0s"
   */
  smbIdleTimeout?: string;
  /** SMB password. */
  smbPass?: string;
  /**
   * SMB port number.
   * @default "445"
   */
  smbPort?: string;
  /** Service principal name. */
  smbSpn?: string;
  /**
   * SMB username.
   * @default "$USER"
   */
  smbUser?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /** Access grant. */
  storjAccessGrant?: string;
  /** API key. */
  storjApiKey?: string;
  /** Encryption passphrase. */
  storjPassphrase?: string;
  /**
   * Choose an authentication method.
   * @default "existing"
   */
  storjProvider?: string;
  /**
   * Satellite address.
   * @default "us1.storj.io"
   */
  storjSatelliteAddress?: string;
  /** Sugarsync Access Key ID. */
  sugarsyncAccessKeyId?: string;
  /** Sugarsync App ID. */
  sugarsyncAppId?: string;
  /** Sugarsync authorization. */
  sugarsyncAuthorization?: string;
  /** Sugarsync authorization expiry. */
  sugarsyncAuthorizationExpiry?: string;
  /** Sugarsync deleted folder id. */
  sugarsyncDeletedId?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Ctl,InvalidUtf8,Dot"
   */
  sugarsyncEncoding?: string;
  /**
   * Permanently delete files if true
   * @default "false"
   */
  sugarsyncHardDelete?: string;
  /** Sugarsync Private Access Key. */
  sugarsyncPrivateAccessKey?: string;
  /** Sugarsync refresh token. */
  sugarsyncRefreshToken?: string;
  /** Sugarsync root id. */
  sugarsyncRootId?: string;
  /** Sugarsync user. */
  sugarsyncUser?: string;
  /** Application Credential ID (OS_APPLICATION_CREDENTIAL_ID). */
  swiftApplicationCredentialId?: string;
  /** Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME). */
  swiftApplicationCredentialName?: string;
  /** Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET). */
  swiftApplicationCredentialSecret?: string;
  /** Authentication URL for server (OS_AUTH_URL). */
  swiftAuth?: string;
  /** Auth Token from alternate authentication - optional (OS_AUTH_TOKEN). */
  swiftAuthToken?: string;
  /**
   * AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).
   * @default "0"
   */
  swiftAuthVersion?: string;
  /**
   * Above this size files will be chunked into a _segments container.
   * @default "5Gi"
   */
  swiftChunkSize?: string;
  /** User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME) */
  swiftDomain?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8"
   */
  swiftEncoding?: string;
  /**
   * Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).
   * @default "public"
   */
  swiftEndpointType?: string;
  /**
   * Get swift credentials from environment variables in standard OpenStack form.
   * @default "false"
   */
  swiftEnvAuth?: string;
  /** API key or password (OS_PASSWORD). */
  swiftKey?: string;
  /**
   * If true avoid calling abort upload on a failure.
   * @default "false"
   */
  swiftLeavePartsOnError?: string;
  /**
   * Don't chunk files during streaming upload.
   * @default "false"
   */
  swiftNoChunk?: string;
  /**
   * Disable support for static and dynamic large objects
   * @default "false"
   */
  swiftNoLargeObjects?: string;
  /** Region name - optional (OS_REGION_NAME). */
  swiftRegion?: string;
  /** The storage policy to use when creating a new container. */
  swiftStoragePolicy?: string;
  /** Storage URL - optional (OS_STORAGE_URL). */
  swiftStorageUrl?: string;
  /** Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME). */
  swiftTenant?: string;
  /** Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME). */
  swiftTenantDomain?: string;
  /** Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID). */
  swiftTenantId?: string;
  /** User name to log in (OS_USERNAME). */
  swiftUser?: string;
  /** User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID). */
  swiftUserId?: string;
  /** Your access token. */
  uptoboxAccessToken?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot"
   */
  uptoboxEncoding?: string;
  /** Bearer token instead of user/pass (e.g. a Macaroon). */
  webdavBearerToken?: string;
  /** Command to run to get a bearer token. */
  webdavBearerTokenCommand?: string;
  /** The encoding for the backend. */
  webdavEncoding?: string;
  /** Set HTTP headers for all transactions. */
  webdavHeaders?: string;
  /** Password. */
  webdavPass?: string;
  /** URL of http host to connect to. */
  webdavUrl?: string;
  /** User name. */
  webdavUser?: string;
  /** Name of the WebDAV site/service/software you are using. */
  webdavVendor?: string;
  /** Auth server URL. */
  yandexAuthUrl?: string;
  /** OAuth Client Id. */
  yandexClientId?: string;
  /** OAuth Client Secret. */
  yandexClientSecret?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Del,Ctl,InvalidUtf8,Dot"
   */
  yandexEncoding?: string;
  /**
   * Delete files permanently rather than putting them into the trash.
   * @default "false"
   */
  yandexHardDelete?: string;
  /** OAuth Access Token as a JSON blob. */
  yandexToken?: string;
  /** Token server url. */
  yandexTokenUrl?: string;
  /** Auth server URL. */
  zohoAuthUrl?: string;
  /** OAuth Client Id. */
  zohoClientId?: string;
  /** OAuth Client Secret. */
  zohoClientSecret?: string;
  /**
   * The encoding for the backend.
   * @default "Del,Ctl,InvalidUtf8"
   */
  zohoEncoding?: string;
  /** Zoho region to connect to. */
  zohoRegion?: string;
  /** OAuth Access Token as a JSON blob. */
  zohoToken?: string;
  /** Token server url. */
  zohoTokenUrl?: string;
}

export interface DatasourceAzureblobRequest {
  /** Access tier of blob: hot, cool or archive. */
  accessTier?: string;
  /** Azure Storage Account Name. */
  account?: string;
  /**
   * Delete archive tier blobs before overwriting.
   * @default "false"
   */
  archiveTierDelete?: string;
  /**
   * Upload chunk size.
   * @default "4Mi"
   */
  chunkSize?: string;
  /** Password for the certificate file (optional). */
  clientCertificatePassword?: string;
  /** Path to a PEM or PKCS12 certificate file including the private key. */
  clientCertificatePath?: string;
  /** The ID of the client in use. */
  clientId?: string;
  /** One of the service principal's client secrets */
  clientSecret?: string;
  /**
   * Send the certificate chain when using certificate auth.
   * @default "false"
   */
  clientSendCertificateChain?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default "false"
   */
  disableChecksum?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"
   */
  encoding?: string;
  /** Endpoint for the service. */
  endpoint?: string;
  /**
   * Read credentials from runtime (environment variables, CLI or MSI).
   * @default "false"
   */
  envAuth?: string;
  /** Storage Account Shared Key. */
  key?: string;
  /**
   * Size of blob list.
   * @default "5000"
   */
  listChunk?: string;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default "false"
   */
  memoryPoolUseMmap?: string;
  /** Object ID of the user-assigned MSI to use, if any. */
  msiClientId?: string;
  /** Azure resource ID of the user-assigned MSI to use, if any. */
  msiMiResId?: string;
  /** Object ID of the user-assigned MSI to use, if any. */
  msiObjectId?: string;
  /**
   * If set, don't attempt to check the container exists or create it.
   * @default "false"
   */
  noCheckContainer?: string;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default "false"
   */
  noHeadObject?: string;
  /** The user's password */
  password?: string;
  /** Public access level of a container: blob or container. */
  publicAccess?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** SAS URL for container level access only. */
  sasUrl?: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** Path to file containing credentials for use with a service principal. */
  servicePrincipalFile?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /** ID of the service principal's tenant. Also called its directory ID. */
  tenant?: string;
  /**
   * Concurrency for multipart uploads.
   * @default "16"
   */
  uploadConcurrency?: string;
  /** Cutoff for switching to chunked upload (<= 256 MiB) (deprecated). */
  uploadCutoff?: string;
  /**
   * Uses local storage emulator if provided as 'true'.
   * @default "false"
   */
  useEmulator?: string;
  /**
   * Use a managed service identity to authenticate (only works in Azure).
   * @default "false"
   */
  useMsi?: string;
  /** User name (usually an email address) */
  username?: string;
}

export interface DatasourceB2Request {
  /** Account ID or Application Key ID. */
  account?: string;
  /**
   * Upload chunk size.
   * @default "96Mi"
   */
  chunkSize?: string;
  /**
   * Cutoff for switching to multipart copy.
   * @default "4Gi"
   */
  copyCutoff?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * Disable checksums for large (> upload cutoff) files.
   * @default "false"
   */
  disableChecksum?: string;
  /**
   * Time before the authorization token will expire in s or suffix ms|s|m|h|d.
   * @default "1w"
   */
  downloadAuthDuration?: string;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Endpoint for the service. */
  endpoint?: string;
  /**
   * Permanently delete files on remote removal, otherwise hide files.
   * @default "false"
   */
  hardDelete?: string;
  /** Application Key. */
  key?: string;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default "false"
   */
  memoryPoolUseMmap?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** A flag string for X-Bz-Test-Mode header for debugging. */
  testMode?: string;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default "false"
   */
  versions?: string;
}

export interface DatasourceBoxRequest {
  /** Box App Primary Access Token */
  accessToken?: string;
  /** Auth server URL. */
  authUrl?: string;
  /** Box App config.json location */
  boxConfigFile?: string;
  /** @default "user" */
  boxSubType?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /**
   * Max number of times to try committing a multipart file.
   * @default "100"
   */
  commitRetries?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Size of listing chunk 1-1000.
   * @default "1000"
   */
  listChunk?: string;
  /** Only show items owned by the login (email address) passed in. */
  ownedBy?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /**
   * Fill in for rclone to use a non root folder as its starting point.
   * @default "0"
   */
  rootFolderId?: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
  /**
   * Cutoff for switching to multipart upload (>= 50 MiB).
   * @default "50Mi"
   */
  uploadCutoff?: string;
}

export interface DatasourceCheckSourceRequest {
  /** Path relative to the data source root */
  path?: string;
}

export interface DatasourceChunksByState {
  /** number of chunks in this state */
  count?: number;
  /** the state of the chunks */
  state?: ModelWorkState;
}

export interface DatasourceDriveRequest {
  /**
   * Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
   * @default "false"
   */
  acknowledgeAbuse?: string;
  /**
   * Allow the filetype to change when uploading Google docs.
   * @default "false"
   */
  allowImportNameChange?: string;
  /**
   * Deprecated: No longer needed.
   * @default "false"
   */
  alternateExport?: string;
  /**
   * Only consider files owned by the authenticated user.
   * @default "false"
   */
  authOwnerOnly?: string;
  /** Auth server URL. */
  authUrl?: string;
  /**
   * Upload chunk size.
   * @default "8Mi"
   */
  chunkSize?: string;
  /** Google Application Client Id */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /**
   * Server side copy contents of shortcuts instead of the shortcut.
   * @default "false"
   */
  copyShortcutContent?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * Disable drive using http2.
   * @default "true"
   */
  disableHttp2?: string;
  /**
   * The encoding for the backend.
   * @default "InvalidUtf8"
   */
  encoding?: string;
  /**
   * Comma separated list of preferred formats for downloading Google docs.
   * @default "docx,xlsx,pptx,svg"
   */
  exportFormats?: string;
  /** Deprecated: See export_formats. */
  formats?: string;
  /** Impersonate this user when using a service account. */
  impersonate?: string;
  /** Comma separated list of preferred formats for uploading Google docs. */
  importFormats?: string;
  /**
   * Keep new head revision of each file forever.
   * @default "false"
   */
  keepRevisionForever?: string;
  /**
   * Size of listing chunk 100-1000, 0 to disable.
   * @default "1000"
   */
  listChunk?: string;
  /**
   * Number of API calls to allow without sleeping.
   * @default "100"
   */
  pacerBurst?: string;
  /**
   * Minimum time to sleep between API calls.
   * @default "100ms"
   */
  pacerMinSleep?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Resource key for accessing a link-shared file. */
  resourceKey?: string;
  /** ID of the root folder. */
  rootFolderId?: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** Scope that rclone should use when requesting access from drive. */
  scope?: string;
  /**
   * Allow server-side operations (e.g. copy) to work across different drive configs.
   * @default "false"
   */
  serverSideAcrossConfigs?: string;
  /** Service Account Credentials JSON blob. */
  serviceAccountCredentials?: string;
  /** Service Account Credentials JSON file path. */
  serviceAccountFile?: string;
  /**
   * Only show files that are shared with me.
   * @default "false"
   */
  sharedWithMe?: string;
  /**
   * Show sizes as storage quota usage, not actual size.
   * @default "false"
   */
  sizeAsQuota?: string;
  /**
   * Skip MD5 checksum on Google photos and videos only.
   * @default "false"
   */
  skipChecksumGphotos?: string;
  /**
   * If set skip dangling shortcut files.
   * @default "false"
   */
  skipDanglingShortcuts?: string;
  /**
   * Skip google documents in all listings.
   * @default "false"
   */
  skipGdocs?: string;
  /**
   * If set skip shortcut files.
   * @default "false"
   */
  skipShortcuts?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Only show files that are starred.
   * @default "false"
   */
  starredOnly?: string;
  /**
   * Make download limit errors be fatal.
   * @default "false"
   */
  stopOnDownloadLimit?: string;
  /**
   * Make upload limit errors be fatal.
   * @default "false"
   */
  stopOnUploadLimit?: string;
  /** ID of the Shared Drive (Team Drive). */
  teamDrive?: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
  /**
   * Only show files that are in the trash.
   * @default "false"
   */
  trashedOnly?: string;
  /**
   * Cutoff for switching to chunked upload.
   * @default "8Mi"
   */
  uploadCutoff?: string;
  /**
   * Use file created date instead of modified date.
   * @default "false"
   */
  useCreatedDate?: string;
  /**
   * Use date file was shared instead of modified date.
   * @default "false"
   */
  useSharedDate?: string;
  /**
   * Send files to the trash instead of deleting permanently.
   * @default "true"
   */
  useTrash?: string;
  /**
   * If Object's are greater, use drive v2 API to download.
   * @default "off"
   */
  v2DownloadMinSize?: string;
}

export interface DatasourceDropboxRequest {
  /** Auth server URL. */
  authUrl?: string;
  /**
   * Max time to wait for a batch to finish committing
   * @default "10m0s"
   */
  batchCommitTimeout?: string;
  /**
   * Upload file batching sync|async|off.
   * @default "sync"
   */
  batchMode?: string;
  /**
   * Max number of files in upload batch.
   * @default "0"
   */
  batchSize?: string;
  /**
   * Max time to allow an idle upload batch before uploading.
   * @default "0s"
   */
  batchTimeout?: string;
  /**
   * Upload chunk size (< 150Mi).
   * @default "48Mi"
   */
  chunkSize?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Impersonate this user when using a business account. */
  impersonate?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /**
   * Instructs rclone to work on individual shared files.
   * @default "false"
   */
  sharedFiles?: string;
  /**
   * Instructs rclone to work on shared folders.
   * @default "false"
   */
  sharedFolders?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
}

export interface DatasourceFichierRequest {
  /** Your API Key, get it from https://1fichier.com/console/params.pl. */
  apiKey?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** If you want to download a shared file that is password protected, add this parameter. */
  filePassword?: string;
  /** If you want to list the files in a shared folder that is password protected, add this parameter. */
  folderPassword?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** If you want to download a shared folder, add this parameter. */
  sharedFolder?: string;
  /** The path of the source to scan items */
  sourcePath: string;
}

export interface DatasourceFilefabricRequest {
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Permanent Authentication Token. */
  permanentToken?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** ID of the root folder. */
  rootFolderId?: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** Session Token. */
  token?: string;
  /** Token expiry time. */
  tokenExpiry?: string;
  /** URL of the Enterprise File Fabric to connect to. */
  url?: string;
  /** Version read from the file fabric. */
  version?: string;
}

export interface DatasourceFtpRequest {
  /**
   * Allow asking for FTP password when needed.
   * @default "false"
   */
  askPassword?: string;
  /**
   * Maximum time to wait for a response to close.
   * @default "1m0s"
   */
  closeTimeout?: string;
  /**
   * Maximum number of FTP simultaneous connections, 0 for unlimited.
   * @default "0"
   */
  concurrency?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * Disable using EPSV even if server advertises support.
   * @default "false"
   */
  disableEpsv?: string;
  /**
   * Disable using MLSD even if server advertises support.
   * @default "false"
   */
  disableMlsd?: string;
  /**
   * Disable TLS 1.3 (workaround for FTP servers with buggy TLS)
   * @default "false"
   */
  disableTls13?: string;
  /**
   * Disable using UTF-8 even if server advertises support.
   * @default "false"
   */
  disableUtf8?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Del,Ctl,RightSpace,Dot"
   */
  encoding?: string;
  /**
   * Use Explicit FTPS (FTP over TLS).
   * @default "false"
   */
  explicitTls?: string;
  /**
   * Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.
   * @default "false"
   */
  forceListHidden?: string;
  /** FTP host to connect to. */
  host?: string;
  /**
   * Max time before closing idle connections.
   * @default "1m0s"
   */
  idleTimeout?: string;
  /**
   * Do not verify the TLS certificate of the server.
   * @default "false"
   */
  noCheckCertificate?: string;
  /** FTP password. */
  pass?: string;
  /**
   * FTP port number.
   * @default "21"
   */
  port?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /**
   * Maximum time to wait for data connection closing status.
   * @default "1m0s"
   */
  shutTimeout?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Use Implicit FTPS (FTP over TLS).
   * @default "false"
   */
  tls?: string;
  /**
   * Size of TLS session cache for all control and data connections.
   * @default "32"
   */
  tlsCacheSize?: string;
  /**
   * FTP username.
   * @default "$USER"
   */
  user?: string;
  /**
   * Use MDTM to set modification time (VsFtpd quirk)
   * @default "false"
   */
  writingMdtm?: string;
}

export interface DatasourceGcsRequest {
  /**
   * Access public buckets and objects without credentials.
   * @default "false"
   */
  anonymous?: string;
  /** Auth server URL. */
  authUrl?: string;
  /** Access Control List for new buckets. */
  bucketAcl?: string;
  /**
   * Access checks should use bucket-level IAM policies.
   * @default "false"
   */
  bucketPolicyOnly?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /**
   * If set this will decompress gzip encoded objects.
   * @default "false"
   */
  decompress?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,CrLf,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Endpoint for the service. */
  endpoint?: string;
  /**
   * Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
   * @default "false"
   */
  envAuth?: string;
  /** Location for the newly created buckets. */
  location?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default "false"
   */
  noCheckBucket?: string;
  /** Access Control List for new objects. */
  objectAcl?: string;
  /** Project number. */
  projectNumber?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** Service Account Credentials JSON blob. */
  serviceAccountCredentials?: string;
  /** Service Account Credentials JSON file path. */
  serviceAccountFile?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /** The storage class to use when storing objects in Google Cloud Storage. */
  storageClass?: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
}

export interface DatasourceGphotosRequest {
  /** Auth server URL. */
  authUrl?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,CrLf,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Also view and download archived media.
   * @default "false"
   */
  includeArchived?: string;
  /**
   * Set to make the Google Photos backend read only.
   * @default "false"
   */
  readOnly?: string;
  /**
   * Set to read the size of media items.
   * @default "false"
   */
  readSize?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Year limits the photos to be downloaded to those which are uploaded after the given year.
   * @default "2000"
   */
  startYear?: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
}

export interface DatasourceHdfsRequest {
  /** Kerberos data transfer protection: authentication|integrity|privacy. */
  dataTransferProtection?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,Colon,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Hadoop name node and port. */
  namenode?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** Kerberos service principal name for the namenode. */
  servicePrincipalName?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /** Hadoop user name. */
  username?: string;
}

export interface DatasourceHidriveRequest {
  /** Auth server URL. */
  authUrl?: string;
  /**
   * Chunksize for chunked uploads.
   * @default "48Mi"
   */
  chunkSize?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * Do not fetch number of objects in directories unless it is absolutely necessary.
   * @default "false"
   */
  disableFetchingMemberCount?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for the service.
   * @default "https://api.hidrive.strato.com/2.1"
   */
  endpoint?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /**
   * The root/parent folder for all paths.
   * @default "/"
   */
  rootPrefix?: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /**
   * Access permissions that rclone should use when requesting access from HiDrive.
   * @default "rw"
   */
  scopeAccess?: string;
  /**
   * User-level that rclone should use when requesting access from HiDrive.
   * @default "user"
   */
  scopeRole?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
  /**
   * Concurrency for chunked uploads.
   * @default "4"
   */
  uploadConcurrency?: string;
  /**
   * Cutoff/Threshold for chunked uploads.
   * @default "96Mi"
   */
  uploadCutoff?: string;
}

export interface DatasourceHttpRequest {
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /** Set HTTP headers for all transactions. */
  headers?: string;
  /**
   * Don't use HEAD requests.
   * @default "false"
   */
  noHead?: string;
  /**
   * Set this if the site doesn't end directories with /.
   * @default "false"
   */
  noSlash?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** URL of HTTP host to connect to. */
  url?: string;
}

export interface DatasourceInternetarchiveRequest {
  /** IAS3 Access Key. */
  accessKeyId?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * Don't ask the server to test against MD5 checksum calculated by rclone.
   * @default "true"
   */
  disableChecksum?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * IAS3 Endpoint.
   * @default "https://s3.us.archive.org"
   */
  endpoint?: string;
  /**
   * Host of InternetArchive Frontend.
   * @default "https://archive.org"
   */
  frontEndpoint?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** IAS3 Secret Key (password). */
  secretAccessKey?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
   * @default "0s"
   */
  waitArchive?: string;
}

export interface DatasourceItemInfo {
  /** Path to the new item, relative to the source */
  path?: string;
}

export interface DatasourceJottacloudRequest {
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Delete files permanently rather than putting them into the trash.
   * @default "false"
   */
  hardDelete?: string;
  /**
   * Files bigger than this will be cached on disk to calculate the MD5 if required.
   * @default "10Mi"
   */
  md5MemoryLimit?: string;
  /**
   * Avoid server side versioning by deleting files and recreating files instead of overwriting them.
   * @default "false"
   */
  noVersions?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Only show files that are in the trash.
   * @default "false"
   */
  trashedOnly?: string;
  /**
   * Files bigger than this can be resumed if the upload fail's.
   * @default "10Mi"
   */
  uploadResumeLimit?: string;
}

export interface DatasourceKoofrRequest {
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** The Koofr API endpoint to use. */
  endpoint?: string;
  /** Mount ID of the mount to use. */
  mountid?: string;
  /** Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password). */
  password?: string;
  /** Choose your storage provider. */
  provider?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /**
   * Does the backend support setting modification time.
   * @default "true"
   */
  setmtime?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /** Your user name. */
  user?: string;
}

export interface DatasourceLocalRequest {
  /**
   * Force the filesystem to report itself as case insensitive.
   * @default "false"
   */
  caseInsensitive?: string;
  /**
   * Force the filesystem to report itself as case sensitive.
   * @default "false"
   */
  caseSensitive?: string;
  /**
   * Follow symlinks and copy the pointed to item.
   * @default "false"
   */
  copyLinks?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,Dot"
   */
  encoding?: string;
  /**
   * Translate symlinks to/from regular files with a '.rclonelink' extension.
   * @default "false"
   */
  links?: string;
  /**
   * Don't check to see if the files change during upload.
   * @default "false"
   */
  noCheckUpdated?: string;
  /**
   * Disable preallocation of disk space for transferred files.
   * @default "false"
   */
  noPreallocate?: string;
  /**
   * Disable setting modtime.
   * @default "false"
   */
  noSetModtime?: string;
  /**
   * Disable sparse files for multi-thread downloads.
   * @default "false"
   */
  noSparse?: string;
  /**
   * Disable UNC (long path names) conversion on Windows.
   * @default "false"
   */
  nounc?: string;
  /**
   * Don't cross filesystem boundaries (unix/macOS only).
   * @default "false"
   */
  oneFileSystem?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /**
   * Don't warn about skipped symlinks.
   * @default "false"
   */
  skipLinks?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Apply unicode NFC normalization to paths and filenames.
   * @default "false"
   */
  unicodeNormalization?: string;
  /**
   * Assume the Stat size of links is zero (and read them instead) (deprecated).
   * @default "false"
   */
  zeroSizeLinks?: string;
}

export interface DatasourceMailruRequest {
  /**
   * What should copy do if file checksum is mismatched or invalid.
   * @default "true"
   */
  checkHash?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Password. */
  pass?: string;
  /** Comma separated list of internal maintenance flags. */
  quirks?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Skip full upload if there is another file with same data hash.
   * @default "true"
   */
  speedupEnable?: string;
  /**
   * Comma separated list of file name patterns eligible for speedup (put by hash).
   * @default "*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf"
   */
  speedupFilePatterns?: string;
  /**
   * This option allows you to disable speedup (put by hash) for large files.
   * @default "3Gi"
   */
  speedupMaxDisk?: string;
  /**
   * Files larger than the size given below will always be hashed on disk.
   * @default "32Mi"
   */
  speedupMaxMemory?: string;
  /** User name (usually email). */
  user?: string;
  /** HTTP user agent used internally by client. */
  userAgent?: string;
}

export interface DatasourceMegaRequest {
  /**
   * Output more debug from Mega.
   * @default "false"
   */
  debug?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Delete files permanently rather than putting them into the trash.
   * @default "false"
   */
  hardDelete?: string;
  /** Password. */
  pass?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Use HTTPS for transfers.
   * @default "false"
   */
  useHttps?: string;
  /** User name. */
  user?: string;
}

export interface DatasourceNetstorageRequest {
  /** Set the NetStorage account name */
  account?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /** Domain+path of NetStorage host to connect to. */
  host?: string;
  /**
   * Select between HTTP or HTTPS protocol.
   * @default "https"
   */
  protocol?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** Set the NetStorage account secret/G2O key for authentication. */
  secret?: string;
  /** The path of the source to scan items */
  sourcePath: string;
}

export interface DatasourceOnedriveRequest {
  /**
   * Set scopes to be requested by rclone.
   * @default "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"
   */
  accessScopes?: string;
  /** Auth server URL. */
  authUrl?: string;
  /**
   * Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
   * @default "10Mi"
   */
  chunkSize?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * Disable the request for Sites.Read.All permission.
   * @default "false"
   */
  disableSitePermission?: string;
  /** The ID of the drive to use. */
  driveId?: string;
  /** The type of the drive (personal | business | documentLibrary). */
  driveType?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Set to make OneNote files show up in directory listings.
   * @default "false"
   */
  exposeOnenoteFiles?: string;
  /**
   * Specify the hash in use for the backend.
   * @default "auto"
   */
  hashType?: string;
  /** Set the password for links created by the link command. */
  linkPassword?: string;
  /**
   * Set the scope of the links created by the link command.
   * @default "anonymous"
   */
  linkScope?: string;
  /**
   * Set the type of the links created by the link command.
   * @default "view"
   */
  linkType?: string;
  /**
   * Size of listing chunk.
   * @default "1000"
   */
  listChunk?: string;
  /**
   * Remove all versions on modifying operations.
   * @default "false"
   */
  noVersions?: string;
  /**
   * Choose national cloud region for OneDrive.
   * @default "global"
   */
  region?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** ID of the root folder. */
  rootFolderId?: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /**
   * Allow server-side operations (e.g. copy) to work across different onedrive configs.
   * @default "false"
   */
  serverSideAcrossConfigs?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
}

export interface DatasourceOosRequest {
  /**
   * Chunk size to use for uploading.
   * @default "5Mi"
   */
  chunkSize?: string;
  /** Object storage compartment OCID */
  compartment?: string;
  /**
   * Path to OCI config file
   * @default "~/.oci/config"
   */
  configFile?: string;
  /**
   * Profile name inside the oci config file
   * @default "Default"
   */
  configProfile?: string;
  /**
   * Cutoff for switching to multipart copy.
   * @default "4.656Gi"
   */
  copyCutoff?: string;
  /**
   * Timeout for copy.
   * @default "1m0s"
   */
  copyTimeout?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default "false"
   */
  disableChecksum?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Endpoint for Object storage API. */
  endpoint?: string;
  /**
   * If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
   * @default "false"
   */
  leavePartsOnError?: string;
  /** Object storage namespace */
  namespace?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default "false"
   */
  noCheckBucket?: string;
  /**
   * Choose your Auth Provider
   * @default "env_auth"
   */
  provider?: string;
  /** Object storage Region */
  region?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm. */
  sseCustomerAlgorithm?: string;
  /** To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to */
  sseCustomerKey?: string;
  /** To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated */
  sseCustomerKeyFile?: string;
  /** If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption */
  sseCustomerKeySha256?: string;
  /** if using using your own master key in vault, this header specifies the */
  sseKmsKeyId?: string;
  /**
   * The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
   * @default "Standard"
   */
  storageTier?: string;
  /**
   * Concurrency for multipart uploads.
   * @default "10"
   */
  uploadConcurrency?: string;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
}

export interface DatasourceOpendriveRequest {
  /**
   * Files will be uploaded in chunks this size.
   * @default "10Mi"
   */
  chunkSize?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Password. */
  password?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** Username. */
  username?: string;
}

export interface DatasourcePcloudRequest {
  /** Auth server URL. */
  authUrl?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Hostname to connect to.
   * @default "api.pcloud.com"
   */
  hostname?: string;
  /** Your pcloud password. */
  password?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /**
   * Fill in for rclone to use a non root folder as its starting point.
   * @default "d0"
   */
  rootFolderId?: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
  /** Your pcloud username. */
  username?: string;
}

export interface DatasourcePremiumizemeRequest {
  /** API Key. */
  apiKey?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
}

export interface DatasourcePutioRequest {
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
}

export interface DatasourceQingstorRequest {
  /** QingStor Access Key ID. */
  accessKeyId?: string;
  /**
   * Chunk size to use for uploading.
   * @default "4Mi"
   */
  chunkSize?: string;
  /**
   * Number of connection retries.
   * @default "3"
   */
  connectionRetries?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,Ctl,InvalidUtf8"
   */
  encoding?: string;
  /** Enter an endpoint URL to connection QingStor API. */
  endpoint?: string;
  /**
   * Get QingStor credentials from runtime.
   * @default "false"
   */
  envAuth?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** QingStor Secret Access Key (password). */
  secretAccessKey?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Concurrency for multipart uploads.
   * @default "1"
   */
  uploadConcurrency?: string;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /** Zone to connect to. */
  zone?: string;
}

export interface DatasourceS3Request {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /** Canned ACL used when creating buckets. */
  bucketAcl?: string;
  /**
   * Chunk size to use for uploading.
   * @default "5Mi"
   */
  chunkSize?: string;
  /**
   * Cutoff for switching to multipart copy.
   * @default "4.656Gi"
   */
  copyCutoff?: string;
  /**
   * If set this will decompress gzip encoded objects.
   * @default "false"
   */
  decompress?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default "false"
   */
  disableChecksum?: string;
  /**
   * Disable usage of http2 for S3 backends.
   * @default "false"
   */
  disableHttp2?: string;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Endpoint for S3 API. */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default "false"
   */
  envAuth?: string;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default "true"
   */
  forcePathStyle?: string;
  /**
   * If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
   * @default "false"
   */
  leavePartsOnError?: string;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default "1000"
   */
  listChunk?: string;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default "0"
   */
  listVersion?: string;
  /** Location constraint - must be set to match the Region. */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default "10000"
   */
  maxUploadParts?: string;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default "false"
   */
  memoryPoolUseMmap?: string;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default "false"
   */
  noCheckBucket?: string;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default "false"
   */
  noHead?: string;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default "false"
   */
  noHeadObject?: string;
  /**
   * Suppress setting and reading of system metadata
   * @default "false"
   */
  noSystemMetadata?: string;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /** Choose your S3 provider. */
  provider?: string;
  /** Region to connect to. */
  region?: string;
  /**
   * Enables requester pays option when interacting with S3 bucket.
   * @default "false"
   */
  requesterPays?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** The server-side encryption algorithm used when storing this object in S3. */
  serverSideEncryption?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /** If using SSE-C, the server-side encryption algorithm used when storing this object in S3. */
  sseCustomerAlgorithm?: string;
  /** To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data. */
  sseCustomerKey?: string;
  /** If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data. */
  sseCustomerKeyBase64?: string;
  /** If using SSE-C you may provide the secret encryption key MD5 checksum (optional). */
  sseCustomerKeyMd5?: string;
  /** If using KMS ID you must provide the ARN of Key. */
  sseKmsKeyId?: string;
  /** The storage class to use when storing new objects in S3. */
  storageClass?: string;
  /** Endpoint for STS. */
  stsEndpoint?: string;
  /**
   * Concurrency for multipart uploads.
   * @default "4"
   */
  uploadConcurrency?: string;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * If true use the AWS S3 accelerated endpoint.
   * @default "false"
   */
  useAccelerateEndpoint?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default "false"
   */
  usePresignedRequest?: string;
  /**
   * If true use v2 authentication.
   * @default "false"
   */
  v2Auth?: string;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default "false"
   */
  versions?: string;
}

export interface DatasourceSeafileRequest {
  /**
   * Two-factor authentication ('true' if the account has 2FA enabled).
   * @default "false"
   */
  "2fa"?: string;
  /** Authentication token. */
  authToken?: string;
  /**
   * Should rclone create a library if it doesn't exist.
   * @default "false"
   */
  createLibrary?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8"
   */
  encoding?: string;
  /** Name of the library. */
  library?: string;
  /** Library password (for encrypted libraries only). */
  libraryKey?: string;
  /** Password. */
  pass?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** URL of seafile host to connect to. */
  url?: string;
  /** User name (usually email address). */
  user?: string;
}

export interface DatasourceSftpRequest {
  /**
   * Allow asking for SFTP password when needed.
   * @default "false"
   */
  askPassword?: string;
  /**
   * Upload and download chunk size.
   * @default "32Ki"
   */
  chunkSize?: string;
  /** Space separated list of ciphers to be used for session encryption, ordered by preference. */
  ciphers?: string;
  /**
   * The maximum number of outstanding requests for one file
   * @default "64"
   */
  concurrency?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * If set don't use concurrent reads.
   * @default "false"
   */
  disableConcurrentReads?: string;
  /**
   * If set don't use concurrent writes.
   * @default "false"
   */
  disableConcurrentWrites?: string;
  /**
   * Disable the execution of SSH commands to determine if remote file hashing is available.
   * @default "false"
   */
  disableHashcheck?: string;
  /** SSH host to connect to. */
  host?: string;
  /**
   * Max time before closing idle connections.
   * @default "1m0s"
   */
  idleTimeout?: string;
  /** Space separated list of key exchange algorithms, ordered by preference. */
  keyExchange?: string;
  /** Path to PEM-encoded private key file. */
  keyFile?: string;
  /** The passphrase to decrypt the PEM-encoded private key file. */
  keyFilePass?: string;
  /** Raw PEM-encoded private key. */
  keyPem?: string;
  /**
   * When set forces the usage of the ssh-agent.
   * @default "false"
   */
  keyUseAgent?: string;
  /** Optional path to known_hosts file. */
  knownHostsFile?: string;
  /** Space separated list of MACs (message authentication code) algorithms, ordered by preference. */
  macs?: string;
  /** The command used to read md5 hashes. */
  md5sumCommand?: string;
  /** SSH password, leave blank to use ssh-agent. */
  pass?: string;
  /** Override path used by SSH shell commands. */
  pathOverride?: string;
  /**
   * SSH port number.
   * @default "22"
   */
  port?: string;
  /** Optional path to public key file. */
  pubkeyFile?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** Specifies the path or command to run a sftp server on the remote host. */
  serverCommand?: string;
  /** Environment variables to pass to sftp and commands */
  setEnv?: string;
  /**
   * Set the modified time on the remote if set.
   * @default "true"
   */
  setModtime?: string;
  /** The command used to read sha1 hashes. */
  sha1sumCommand?: string;
  /** The type of SSH shell on remote server, if any. */
  shellType?: string;
  /**
   * Set to skip any symlinks and any other non regular files.
   * @default "false"
   */
  skipLinks?: string;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Specifies the SSH2 subsystem on the remote host.
   * @default "sftp"
   */
  subsystem?: string;
  /**
   * If set use fstat instead of stat.
   * @default "false"
   */
  useFstat?: string;
  /**
   * Enable the use of insecure ciphers and key exchange methods.
   * @default "false"
   */
  useInsecureCipher?: string;
  /**
   * SSH username.
   * @default "$USER"
   */
  user?: string;
}

export interface DatasourceSharefileRequest {
  /**
   * Upload chunk size.
   * @default "64Mi"
   */
  chunkSize?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Endpoint for API calls. */
  endpoint?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** ID of the root folder. */
  rootFolderId?: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Cutoff for switching to multipart upload.
   * @default "128Mi"
   */
  uploadCutoff?: string;
}

export interface DatasourceSiaRequest {
  /** Sia Daemon API Password. */
  apiPassword?: string;
  /**
   * Sia daemon API URL, like http://sia.daemon.host:9980.
   * @default "http://127.0.0.1:9980"
   */
  apiUrl?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /**
   * Siad User Agent
   * @default "Sia-Agent"
   */
  userAgent?: string;
}

export interface DatasourceSmbRequest {
  /**
   * Whether the server is configured to be case-insensitive.
   * @default "true"
   */
  caseInsensitive?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * Domain name for NTLM authentication.
   * @default "WORKGROUP"
   */
  domain?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Hide special shares (e.g. print$) which users aren't supposed to access.
   * @default "true"
   */
  hideSpecialShare?: string;
  /** SMB server hostname to connect to. */
  host?: string;
  /**
   * Max time before closing idle connections.
   * @default "1m0s"
   */
  idleTimeout?: string;
  /** SMB password. */
  pass?: string;
  /**
   * SMB port number.
   * @default "445"
   */
  port?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** Service principal name. */
  spn?: string;
  /**
   * SMB username.
   * @default "$USER"
   */
  user?: string;
}

export interface DatasourceStorjRequest {
  /** Access grant. */
  accessGrant?: string;
  /** API key. */
  apiKey?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /** Encryption passphrase. */
  passphrase?: string;
  /**
   * Choose an authentication method.
   * @default "existing"
   */
  provider?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /**
   * Satellite address.
   * @default "us1.storj.io"
   */
  satelliteAddress?: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
}

export interface DatasourceSugarsyncRequest {
  /** Sugarsync Access Key ID. */
  accessKeyId?: string;
  /** Sugarsync App ID. */
  appId?: string;
  /** Sugarsync authorization. */
  authorization?: string;
  /** Sugarsync authorization expiry. */
  authorizationExpiry?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /** Sugarsync deleted folder id. */
  deletedId?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Permanently delete files if true
   * @default "false"
   */
  hardDelete?: string;
  /** Sugarsync Private Access Key. */
  privateAccessKey?: string;
  /** Sugarsync refresh token. */
  refreshToken?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Sugarsync root id. */
  rootId?: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** Sugarsync user. */
  user?: string;
}

export interface DatasourceSwiftRequest {
  /** Application Credential ID (OS_APPLICATION_CREDENTIAL_ID). */
  applicationCredentialId?: string;
  /** Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME). */
  applicationCredentialName?: string;
  /** Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET). */
  applicationCredentialSecret?: string;
  /** Authentication URL for server (OS_AUTH_URL). */
  auth?: string;
  /** Auth Token from alternate authentication - optional (OS_AUTH_TOKEN). */
  authToken?: string;
  /**
   * AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).
   * @default "0"
   */
  authVersion?: string;
  /**
   * Above this size files will be chunked into a _segments container.
   * @default "5Gi"
   */
  chunkSize?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /** User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME) */
  domain?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8"
   */
  encoding?: string;
  /**
   * Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE).
   * @default "public"
   */
  endpointType?: string;
  /**
   * Get swift credentials from environment variables in standard OpenStack form.
   * @default "false"
   */
  envAuth?: string;
  /** API key or password (OS_PASSWORD). */
  key?: string;
  /**
   * If true avoid calling abort upload on a failure.
   * @default "false"
   */
  leavePartsOnError?: string;
  /**
   * Don't chunk files during streaming upload.
   * @default "false"
   */
  noChunk?: string;
  /**
   * Disable support for static and dynamic large objects
   * @default "false"
   */
  noLargeObjects?: string;
  /** Region name - optional (OS_REGION_NAME). */
  region?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** The storage policy to use when creating a new container. */
  storagePolicy?: string;
  /** Storage URL - optional (OS_STORAGE_URL). */
  storageUrl?: string;
  /** Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME). */
  tenant?: string;
  /** Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME). */
  tenantDomain?: string;
  /** Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID). */
  tenantId?: string;
  /** User name to log in (OS_USERNAME). */
  user?: string;
  /** User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID). */
  userId?: string;
}

export interface DatasourceUptoboxRequest {
  /** Your access token. */
  accessToken?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
}

export interface DatasourceWebdavRequest {
  /** Bearer token instead of user/pass (e.g. a Macaroon). */
  bearerToken?: string;
  /** Command to run to get a bearer token. */
  bearerTokenCommand?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /** The encoding for the backend. */
  encoding?: string;
  /** Set HTTP headers for all transactions. */
  headers?: string;
  /** Password. */
  pass?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** URL of http host to connect to. */
  url?: string;
  /** User name. */
  user?: string;
  /** Name of the WebDAV site/service/software you are using. */
  vendor?: string;
}

export interface DatasourceYandexRequest {
  /** Auth server URL. */
  authUrl?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Delete files permanently rather than putting them into the trash.
   * @default "false"
   */
  hardDelete?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
}

export interface DatasourceZohoRequest {
  /** Auth server URL. */
  authUrl?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /** Delete the source after exporting to CAR files */
  deleteAfterExport: boolean;
  /**
   * The encoding for the backend.
   * @default "Del,Ctl,InvalidUtf8"
   */
  encoding?: string;
  /** Zoho region to connect to. */
  region?: string;
  /** Automatically rescan the source directory when this interval has passed from last successful scan */
  rescanInterval: string;
  /** Starting state for scanning */
  scanningState: ModelWorkState;
  /** The path of the source to scan items */
  sourcePath: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
}

export interface DealListDealRequest {
  /** dataset name filter */
  datasets?: string[];
  /** provider filter */
  providers?: string[];
  /** schedule id filter */
  schedules?: number[];
  /** state filter */
  states?: string[];
}

export interface DealProposal {
  /** Client address */
  clientAddress?: string;
  /**
   * Duration in epoch or in duration format, i.e. 1500000, 2400h
   * @default "12740h"
   */
  duration?: string;
  /** File size in bytes for boost to fetch the CAR file */
  fileSize?: number;
  /** http headers to be passed with the request (i.e. key=value) */
  httpHeaders?: string[];
  /**
   * Whether the deal should be IPNI
   * @default true
   */
  ipni?: boolean;
  /**
   * Whether the deal should be kept unsealed
   * @default true
   */
  keepUnsealed?: boolean;
  /** Piece CID */
  pieceCid?: string;
  /** Piece size */
  pieceSize?: string;
  /**
   * Price in FIL per deal
   * @default 0
   */
  pricePerDeal?: number;
  /**
   * Price in FIL  per GiB
   * @default 0
   */
  pricePerGb?: number;
  /**
   * Price in FIL per GiB per epoch
   * @default 0
   */
  pricePerGbEpoch?: number;
  /** Provider ID */
  providerId?: string;
  /**
   * Root CID that is required as part of the deal proposal, if empty, will be set to empty CID
   * @default "bafkqaaa"
   */
  rootCid?: string;
  /**
   * Deal start delay in epoch or in duration format, i.e. 1000, 72h
   * @default "72h"
   */
  startDelay?: string;
  /** URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car */
  urlTemplate?: string;
  /**
   * Whether the deal should be verified
   * @default true
   */
  verified?: boolean;
}

export interface GithubComDataPreservationProgramsSingularityHandlerDatasourceEntry {
  isDir?: boolean;
  lastModified?: string;
  path?: string;
  size?: number;
}

export interface InspectDirDetail {
  current?: ModelDirectory;
  dirs?: ModelDirectory[];
  items?: ModelItem[];
}

export interface InspectGetPathRequest {
  path?: string;
}

export type ModelCID = object;

export interface ModelCar {
  chunkId?: number;
  createdAt?: string;
  datasetId?: number;
  filePath?: string;
  fileSize?: number;
  header?: number[];
  id?: number;
  pieceCid?: ModelCID;
  pieceSize?: number;
  rootCid?: ModelCID;
  sourceId?: number;
}

export interface ModelChunk {
  cars?: ModelCar[];
  createdAt?: string;
  errorMessage?: string;
  id?: number;
  itemParts?: ModelItemPart[];
  packingState?: ModelWorkState;
  packingWorkerId?: string;
  sourceId?: number;
}

export interface ModelDataset {
  createdAt?: string;
  encryptionRecipients?: string[];
  encryptionScript?: string;
  id?: number;
  maxSize?: number;
  metadata?: ModelMetadata;
  name?: string;
  outputDirs?: string[];
  pieceSize?: number;
  updatedAt?: string;
}

export interface ModelDeal {
  clientId?: string;
  createdAt?: string;
  datasetId?: number;
  dealId?: number;
  endEpoch?: number;
  errorMessage?: string;
  id?: number;
  label?: string;
  pieceCid?: ModelCID;
  pieceSize?: number;
  price?: string;
  proposalId?: string;
  provider?: string;
  scheduleId?: number;
  sectorStartEpoch?: number;
  startEpoch?: number;
  state?: ModelDealState;
  updatedAt?: string;
  verified?: boolean;
}

export enum ModelDealState {
  DealProposed = "proposed",
  DealPublished = "published",
  DealActive = "active",
  DealExpired = "expired",
  DealProposalExpired = "proposal_expired",
  DealRejected = "rejected",
  DealSlashed = "slashed",
  DealErrored = "error",
}

export interface ModelDirectory {
  cid?: ModelCID;
  exported?: boolean;
  id?: number;
  name?: string;
  parentId?: number;
  sourceId?: number;
  updatedAt?: string;
}

export interface ModelItem {
  cid?: ModelCID;
  createdAt?: string;
  directoryId?: number;
  hash?: string;
  id?: number;
  itemParts?: ModelItemPart[];
  lastModified?: number;
  path?: string;
  size?: number;
  sourceId?: number;
}

export interface ModelItemPart {
  chunkId?: number;
  cid?: ModelCID;
  id?: number;
  item?: ModelItem;
  itemId?: number;
  length?: number;
  offset?: number;
}

export type ModelMetadata = Record<string, string>;

export interface ModelSchedule {
  allowedPieceCids?: string[];
  announceToIpni?: boolean;
  createdAt?: string;
  datasetId?: number;
  duration?: number;
  errorMessage?: string;
  httpHeaders?: string[];
  id?: number;
  keepUnsealed?: boolean;
  maxPendingDealNumber?: number;
  maxPendingDealSize?: number;
  notes?: string;
  pricePerDeal?: number;
  pricePerGb?: number;
  pricePerGbEpoch?: number;
  provider?: string;
  scheduleCron?: string;
  scheduleDealNumber?: number;
  scheduleDealSize?: number;
  startDelay?: number;
  state?: ModelScheduleState;
  totalDealNumber?: number;
  totalDealSize?: number;
  updatedAt?: string;
  urlTemplate?: string;
  verified?: boolean;
}

export enum ModelScheduleState {
  ScheduleActive = "active",
  SchedulePaused = "paused",
  ScheduleError = "error",
  ScheduleCompleted = "completed",
}

export interface ModelSource {
  createdAt?: string;
  dagGenErrorMessage?: string;
  dagGenState?: ModelWorkState;
  dagGenWorkerId?: string;
  datasetId?: number;
  deleteAfterExport?: boolean;
  errorMessage?: string;
  id?: number;
  lastScannedPath?: string;
  lastScannedTimestamp?: number;
  metadata?: ModelMetadata;
  path?: string;
  scanIntervalSeconds?: number;
  scanningState?: ModelWorkState;
  scanningWorkerId?: string;
  type?: ModelSourceType;
  updatedAt?: string;
}

export enum ModelSourceType {
  Local = "local",
  Upload = "upload",
}

export interface ModelWallet {
  /** Address is the Filecoin full address of the wallet */
  address?: string;
  /** ID is the short ID of the wallet */
  id?: string;
  /** PrivateKey is the private key of the wallet */
  privateKey?: string;
  /** RemotePeer is the remote peer ID of the wallet, for remote signing purpose */
  remotePeer?: string;
}

export interface ModelWalletAssignment {
  datasetId?: number;
  id?: number;
  walletId?: string;
}

export enum ModelWorkState {
  Created = "",
  Ready = "ready",
  Processing = "processing",
  Complete = "complete",
  Error = "error",
}

export interface ScheduleCreateRequest {
  /** Allowed piece CIDs in this schedule */
  allowedPieceCids?: string[];
  /** Dataset name */
  datasetName?: string;
  /**
   * Duration in epoch or in duration format, i.e. 1500000, 2400h
   * @default "12840h"
   */
  duration?: string;
  /** http headers to be passed with the request (i.e. key=value) */
  httpHeaders?: string[];
  /**
   * Whether the deal should be IPNI
   * @default true
   */
  ipni?: boolean;
  /**
   * Whether the deal should be kept unsealed
   * @default true
   */
  keepUnsealed?: boolean;
  /** Max pending deal number */
  maxPendingDealNumber?: number;
  /** Max pending deal size in human readable format, i.e. 100 TiB */
  maxPendingDealSize?: string;
  /** Notes */
  notes?: string;
  /**
   * Price in FIL per deal
   * @default 0
   */
  pricePerDeal?: number;
  /**
   * Price in FIL  per GiB
   * @default 0
   */
  pricePerGb?: number;
  /**
   * Price in FIL per GiB per epoch
   * @default 0
   */
  pricePerGbEpoch?: number;
  /** Provider */
  provider?: string;
  /** Schedule cron patter */
  scheduleCron?: string;
  /** Number of deals per scheduled time */
  scheduleDealNumber?: number;
  /** Size of deals per schedule trigger in human readable format, i.e. 100 TiB */
  scheduleDealSize?: string;
  /**
   * Deal start delay in epoch or in duration format, i.e. 1000, 72h
   * @default "72h"
   */
  startDelay?: string;
  /** Total number of deals */
  totalDealNumber?: number;
  /** Total size of deals in human readable format, i.e. 100 TiB */
  totalDealSize?: string;
  /** URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car */
  urlTemplate?: string;
  /**
   * Whether the deal should be verified
   * @default true
   */
  verified?: boolean;
}

export type StorePieceReader = object;

export interface WalletAddRemoteRequest {
  /** Address is the Filecoin full address of the wallet */
  address?: string;
  /** RemotePeer is the remote peer ID of the wallet, for remote signing purpose */
  remotePeer?: string;
}

export interface WalletImportRequest {
  /** This is the exported private key from lotus wallet export */
  privateKey?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseFormat;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType | null) => Promise<RequestParams | void> | RequestParams | void;
  customFetch?: typeof fetch;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
  Text = "text/plain",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "//localhost:9090/api";
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private abortControllers = new Map<CancelToken, AbortController>();
  private customFetch = (...fetchParams: Parameters<typeof fetch>) => fetch(...fetchParams);

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  protected encodeQueryParam(key: string, value: any) {
    const encodedKey = encodeURIComponent(key);
    return `${encodedKey}=${encodeURIComponent(typeof value === "number" ? value : `${value}`)}`;
  }

  protected addQueryParam(query: QueryParamsType, key: string) {
    return this.encodeQueryParam(key, query[key]);
  }

  protected addArrayQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];
    return value.map((v: any) => this.encodeQueryParam(key, v)).join("&");
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) => (Array.isArray(query[key]) ? this.addArrayQueryParam(query, key) : this.addQueryParam(query, key)))
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.Text]: (input: any) => (input !== null && typeof input !== "string" ? JSON.stringify(input) : input),
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((formData, key) => {
        const property = input[key];
        formData.append(
          key,
          property instanceof Blob
            ? property
            : typeof property === "object" && property !== null
            ? JSON.stringify(property)
            : `${property}`,
        );
        return formData;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  protected mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  protected createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = async <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format,
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.baseApiParams.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];
    const responseFormat = format || requestParams.format;

    return this.customFetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(requestParams.headers || {}),
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
      },
      signal: (cancelToken ? this.createAbortSignal(cancelToken) : requestParams.signal) || null,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = null as unknown as T;
      r.error = null as unknown as E;

      const data = !responseFormat
        ? r
        : await response[responseFormat]()
            .then((data) => {
              if (r.ok) {
                r.data = data;
              } else {
                r.error = data;
              }
              return r;
            })
            .catch((e) => {
              r.error = e;
              return r;
            });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title Singularity API
 * @version beta
 * @license MIT + Apache 2.0 (https://github.com/data-preservation-programs/singularity/blob/main/LICENSE)
 * @baseUrl //localhost:9090/api
 * @externalDocs https://swagger.io/resources/open-api/
 * @contact Xinan Xu (https://github.com/data-preservation-programs/singularity/issues)
 *
 * This is the API for Singularity, a tool for large-scale clients with PB-scale data onboarding to Filecoin network.
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  admin = {
    /**
     * No description
     *
     * @tags Admin
     * @name InitCreate
     * @summary Initialize the database
     * @request POST:/admin/init
     */
    initCreate: (params: RequestParams = {}) =>
      this.request<void, ApiHTTPError>({
        path: `/admin/init`,
        method: "POST",
        ...params,
      }),

    /**
     * @description This will drop all tables and recreate them.
     *
     * @tags Admin
     * @name ResetCreate
     * @summary Reset the database
     * @request POST:/admin/reset
     */
    resetCreate: (params: RequestParams = {}) =>
      this.request<void, ApiHTTPError>({
        path: `/admin/reset`,
        method: "POST",
        ...params,
      }),
  };
  chunk = {
    /**
     * No description
     *
     * @tags Data Source
     * @name ChunkDetail
     * @summary Get detail of a specific chunk
     * @request GET:/chunk/{id}
     */
    chunkDetail: (id: string, params: RequestParams = {}) =>
      this.request<ModelChunk, ApiHTTPError>({
        path: `/chunk/${id}`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  dataset = {
    /**
     * No description
     *
     * @tags Dataset
     * @name DatasetList
     * @summary List all datasets
     * @request GET:/dataset
     */
    datasetList: (params: RequestParams = {}) =>
      this.request<ModelDataset[], ApiHTTPError>({
        path: `/dataset`,
        method: "GET",
        format: "json",
        ...params,
      }),

    /**
     * @description The dataset is a top level object to distinguish different dataset.
     *
     * @tags Dataset
     * @name DatasetCreate
     * @summary Create a new dataset
     * @request POST:/dataset
     */
    datasetCreate: (request: DatasetCreateRequest, params: RequestParams = {}) =>
      this.request<ModelDataset, ApiHTTPError>({
        path: `/dataset`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description Important! If the dataset is large, this command will take some time to remove all relevant data.
     *
     * @tags Dataset
     * @name DatasetDelete
     * @summary Remove a specific dataset. This will not remove the CAR files.
     * @request DELETE:/dataset/{datasetName}
     */
    datasetDelete: (datasetName: string, params: RequestParams = {}) =>
      this.request<void, ApiHTTPError>({
        path: `/dataset/${datasetName}`,
        method: "DELETE",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Dataset
     * @name DatasetPartialUpdate
     * @summary Update a dataset
     * @request PATCH:/dataset/{datasetName}
     */
    datasetPartialUpdate: (datasetName: string, request: DatasetUpdateRequest, params: RequestParams = {}) =>
      this.request<ModelDataset, ApiHTTPError>({
        path: `/dataset/${datasetName}`,
        method: "PATCH",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Dataset
     * @name PieceDetail
     * @summary List all pieces for the dataset that are available for deal making
     * @request GET:/dataset/{datasetName}/piece
     */
    pieceDetail: (datasetName: string, params: RequestParams = {}) =>
      this.request<ModelCar[], ApiHTTPError>({
        path: `/dataset/${datasetName}/piece`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Dataset
     * @name PieceCreate
     * @summary Manually register a piece (CAR file) with the dataset for deal making purpose
     * @request POST:/dataset/{datasetName}/piece
     */
    pieceCreate: (datasetName: string, request: DatasetAddPieceRequest, params: RequestParams = {}) =>
      this.request<ModelCar, ApiHTTPError>({
        path: `/dataset/${datasetName}/piece`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Wallet
     * @name WalletDetail
     * @summary List all wallets of a dataset.
     * @request GET:/dataset/{datasetName}/wallet
     */
    walletDetail: (datasetName: string, params: RequestParams = {}) =>
      this.request<ModelWallet[], ApiHTTPError>({
        path: `/dataset/${datasetName}/wallet`,
        method: "GET",
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Wallet Association
     * @name WalletCreate
     * @summary Associate a new wallet with a dataset
     * @request POST:/dataset/{datasetName}/wallet/{wallet}
     */
    walletCreate: (datasetName: string, wallet: string, params: RequestParams = {}) =>
      this.request<ModelWalletAssignment, ApiHTTPError>({
        path: `/dataset/${datasetName}/wallet/${wallet}`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Wallet
     * @name WalletDelete
     * @summary Remove an associated wallet from a dataset
     * @request DELETE:/dataset/{datasetName}/wallet/{wallet}
     */
    walletDelete: (datasetName: string, wallet: string, params: RequestParams = {}) =>
      this.request<void, ApiHTTPError>({
        path: `/dataset/${datasetName}/wallet/${wallet}`,
        method: "DELETE",
        ...params,
      }),
  };
  deal = {
    /**
     * @description List all deals
     *
     * @tags Deal
     * @name DealCreate
     * @summary List all deals
     * @request POST:/deal
     */
    dealCreate: (request: DealListDealRequest, params: RequestParams = {}) =>
      this.request<ModelDeal[], ApiHTTPError>({
        path: `/deal`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  item = {
    /**
     * No description
     *
     * @tags Data Source
     * @name ItemDetail
     * @summary Get details about an item
     * @request GET:/item/{id}
     */
    itemDetail: (id: string, params: RequestParams = {}) =>
      this.request<ModelItem, ApiHTTPError>({
        path: `/item/${id}`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  piece = {
    /**
     * @description Get metadata for a piece for how it may be reassembled from the data source
     *
     * @tags Metadata
     * @name MetadataDetail
     * @summary Get metadata for a piece
     * @request GET:/piece/{id}/metadata
     */
    metadataDetail: (id: string, params: RequestParams = {}) =>
      this.request<StorePieceReader, string>({
        path: `/piece/${id}/metadata`,
        method: "GET",
        format: "json",
        ...params,
      }),
  };
  schedule = {
    /**
     * @description Create a new schedule
     *
     * @tags Deal Schedule
     * @name ScheduleCreate
     * @summary Create a new schedule
     * @request POST:/schedule
     */
    scheduleCreate: (schedule: ScheduleCreateRequest, params: RequestParams = {}) =>
      this.request<ModelSchedule, ApiHTTPError>({
        path: `/schedule`,
        method: "POST",
        body: schedule,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Deal Schedule
     * @name PauseCreate
     * @summary Pause a specific schedule
     * @request POST:/schedule/{id}/pause
     */
    pauseCreate: (id: string, params: RequestParams = {}) =>
      this.request<ModelSchedule, ApiHTTPError>({
        path: `/schedule/${id}/pause`,
        method: "POST",
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Deal Schedule
     * @name ResumeCreate
     * @summary Resume a specific schedule
     * @request POST:/schedule/{id}/resume
     */
    resumeCreate: (id: string, params: RequestParams = {}) =>
      this.request<ModelSchedule, ApiHTTPError>({
        path: `/schedule/${id}/resume`,
        method: "POST",
        format: "json",
        ...params,
      }),
  };
  schedules = {
    /**
     * No description
     *
     * @tags Deal Schedule
     * @name SchedulesList
     * @summary List all deal making schedules
     * @request GET:/schedules
     */
    schedulesList: (params: RequestParams = {}) =>
      this.request<ModelSchedule[], ApiHTTPError>({
        path: `/schedules`,
        method: "GET",
        format: "json",
        ...params,
      }),
  };
  sendDeal = {
    /**
     * @description Send a manual deal proposal
     *
     * @tags Deal
     * @name SendDealCreate
     * @summary Send a manual deal proposal
     * @request POST:/send_deal
     */
    sendDealCreate: (proposal: DealProposal, params: RequestParams = {}) =>
      this.request<ModelDeal, ApiHTTPError>({
        path: `/send_deal`,
        method: "POST",
        body: proposal,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  source = {
    /**
     * No description
     *
     * @tags Data Source
     * @name SourceList
     * @summary List all sources for a dataset
     * @request GET:/source
     */
    sourceList: (
      query?: {
        /** Dataset name */
        dataset?: string;
      },
      params: RequestParams = {},
    ) =>
      this.request<ModelSource[], ApiHTTPError>({
        path: `/source`,
        method: "GET",
        query: query,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name AcdDatasetCreate
     * @summary Add acd source for a dataset
     * @request POST:/source/acd/dataset/{datasetName}
     */
    acdDatasetCreate: (datasetName: string, request: DatasourceAcdRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/acd/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name AzureblobDatasetCreate
     * @summary Add azureblob source for a dataset
     * @request POST:/source/azureblob/dataset/{datasetName}
     */
    azureblobDatasetCreate: (datasetName: string, request: DatasourceAzureblobRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/azureblob/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name B2DatasetCreate
     * @summary Add b2 source for a dataset
     * @request POST:/source/b2/dataset/{datasetName}
     */
    b2DatasetCreate: (datasetName: string, request: DatasourceB2Request, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/b2/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name BoxDatasetCreate
     * @summary Add box source for a dataset
     * @request POST:/source/box/dataset/{datasetName}
     */
    boxDatasetCreate: (datasetName: string, request: DatasourceBoxRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/box/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name DriveDatasetCreate
     * @summary Add drive source for a dataset
     * @request POST:/source/drive/dataset/{datasetName}
     */
    driveDatasetCreate: (datasetName: string, request: DatasourceDriveRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/drive/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name DropboxDatasetCreate
     * @summary Add dropbox source for a dataset
     * @request POST:/source/dropbox/dataset/{datasetName}
     */
    dropboxDatasetCreate: (datasetName: string, request: DatasourceDropboxRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/dropbox/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name FichierDatasetCreate
     * @summary Add fichier source for a dataset
     * @request POST:/source/fichier/dataset/{datasetName}
     */
    fichierDatasetCreate: (datasetName: string, request: DatasourceFichierRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/fichier/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name FilefabricDatasetCreate
     * @summary Add filefabric source for a dataset
     * @request POST:/source/filefabric/dataset/{datasetName}
     */
    filefabricDatasetCreate: (datasetName: string, request: DatasourceFilefabricRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/filefabric/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name FtpDatasetCreate
     * @summary Add ftp source for a dataset
     * @request POST:/source/ftp/dataset/{datasetName}
     */
    ftpDatasetCreate: (datasetName: string, request: DatasourceFtpRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/ftp/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name GcsDatasetCreate
     * @summary Add gcs source for a dataset
     * @request POST:/source/gcs/dataset/{datasetName}
     */
    gcsDatasetCreate: (datasetName: string, request: DatasourceGcsRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/gcs/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name GphotosDatasetCreate
     * @summary Add gphotos source for a dataset
     * @request POST:/source/gphotos/dataset/{datasetName}
     */
    gphotosDatasetCreate: (datasetName: string, request: DatasourceGphotosRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/gphotos/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name HdfsDatasetCreate
     * @summary Add hdfs source for a dataset
     * @request POST:/source/hdfs/dataset/{datasetName}
     */
    hdfsDatasetCreate: (datasetName: string, request: DatasourceHdfsRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/hdfs/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name HidriveDatasetCreate
     * @summary Add hidrive source for a dataset
     * @request POST:/source/hidrive/dataset/{datasetName}
     */
    hidriveDatasetCreate: (datasetName: string, request: DatasourceHidriveRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/hidrive/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name HttpDatasetCreate
     * @summary Add http source for a dataset
     * @request POST:/source/http/dataset/{datasetName}
     */
    httpDatasetCreate: (datasetName: string, request: DatasourceHttpRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/http/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name InternetarchiveDatasetCreate
     * @summary Add internetarchive source for a dataset
     * @request POST:/source/internetarchive/dataset/{datasetName}
     */
    internetarchiveDatasetCreate: (
      datasetName: string,
      request: DatasourceInternetarchiveRequest,
      params: RequestParams = {},
    ) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/internetarchive/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name JottacloudDatasetCreate
     * @summary Add jottacloud source for a dataset
     * @request POST:/source/jottacloud/dataset/{datasetName}
     */
    jottacloudDatasetCreate: (datasetName: string, request: DatasourceJottacloudRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/jottacloud/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name KoofrDatasetCreate
     * @summary Add koofr source for a dataset
     * @request POST:/source/koofr/dataset/{datasetName}
     */
    koofrDatasetCreate: (datasetName: string, request: DatasourceKoofrRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/koofr/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name LocalDatasetCreate
     * @summary Add local source for a dataset
     * @request POST:/source/local/dataset/{datasetName}
     */
    localDatasetCreate: (datasetName: string, request: DatasourceLocalRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/local/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name MailruDatasetCreate
     * @summary Add mailru source for a dataset
     * @request POST:/source/mailru/dataset/{datasetName}
     */
    mailruDatasetCreate: (datasetName: string, request: DatasourceMailruRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/mailru/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name MegaDatasetCreate
     * @summary Add mega source for a dataset
     * @request POST:/source/mega/dataset/{datasetName}
     */
    megaDatasetCreate: (datasetName: string, request: DatasourceMegaRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/mega/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name NetstorageDatasetCreate
     * @summary Add netstorage source for a dataset
     * @request POST:/source/netstorage/dataset/{datasetName}
     */
    netstorageDatasetCreate: (datasetName: string, request: DatasourceNetstorageRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/netstorage/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name OnedriveDatasetCreate
     * @summary Add onedrive source for a dataset
     * @request POST:/source/onedrive/dataset/{datasetName}
     */
    onedriveDatasetCreate: (datasetName: string, request: DatasourceOnedriveRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/onedrive/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name OosDatasetCreate
     * @summary Add oos source for a dataset
     * @request POST:/source/oos/dataset/{datasetName}
     */
    oosDatasetCreate: (datasetName: string, request: DatasourceOosRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/oos/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name OpendriveDatasetCreate
     * @summary Add opendrive source for a dataset
     * @request POST:/source/opendrive/dataset/{datasetName}
     */
    opendriveDatasetCreate: (datasetName: string, request: DatasourceOpendriveRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/opendrive/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name PcloudDatasetCreate
     * @summary Add pcloud source for a dataset
     * @request POST:/source/pcloud/dataset/{datasetName}
     */
    pcloudDatasetCreate: (datasetName: string, request: DatasourcePcloudRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/pcloud/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name PremiumizemeDatasetCreate
     * @summary Add premiumizeme source for a dataset
     * @request POST:/source/premiumizeme/dataset/{datasetName}
     */
    premiumizemeDatasetCreate: (
      datasetName: string,
      request: DatasourcePremiumizemeRequest,
      params: RequestParams = {},
    ) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/premiumizeme/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name PutioDatasetCreate
     * @summary Add putio source for a dataset
     * @request POST:/source/putio/dataset/{datasetName}
     */
    putioDatasetCreate: (datasetName: string, request: DatasourcePutioRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/putio/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name QingstorDatasetCreate
     * @summary Add qingstor source for a dataset
     * @request POST:/source/qingstor/dataset/{datasetName}
     */
    qingstorDatasetCreate: (datasetName: string, request: DatasourceQingstorRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/qingstor/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name S3DatasetCreate
     * @summary Add s3 source for a dataset
     * @request POST:/source/s3/dataset/{datasetName}
     */
    s3DatasetCreate: (datasetName: string, request: DatasourceS3Request, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/s3/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name SeafileDatasetCreate
     * @summary Add seafile source for a dataset
     * @request POST:/source/seafile/dataset/{datasetName}
     */
    seafileDatasetCreate: (datasetName: string, request: DatasourceSeafileRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/seafile/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name SftpDatasetCreate
     * @summary Add sftp source for a dataset
     * @request POST:/source/sftp/dataset/{datasetName}
     */
    sftpDatasetCreate: (datasetName: string, request: DatasourceSftpRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/sftp/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name SharefileDatasetCreate
     * @summary Add sharefile source for a dataset
     * @request POST:/source/sharefile/dataset/{datasetName}
     */
    sharefileDatasetCreate: (datasetName: string, request: DatasourceSharefileRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/sharefile/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name SiaDatasetCreate
     * @summary Add sia source for a dataset
     * @request POST:/source/sia/dataset/{datasetName}
     */
    siaDatasetCreate: (datasetName: string, request: DatasourceSiaRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/sia/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name SmbDatasetCreate
     * @summary Add smb source for a dataset
     * @request POST:/source/smb/dataset/{datasetName}
     */
    smbDatasetCreate: (datasetName: string, request: DatasourceSmbRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/smb/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name StorjDatasetCreate
     * @summary Add storj source for a dataset
     * @request POST:/source/storj/dataset/{datasetName}
     */
    storjDatasetCreate: (datasetName: string, request: DatasourceStorjRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/storj/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name SugarsyncDatasetCreate
     * @summary Add sugarsync source for a dataset
     * @request POST:/source/sugarsync/dataset/{datasetName}
     */
    sugarsyncDatasetCreate: (datasetName: string, request: DatasourceSugarsyncRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/sugarsync/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name SwiftDatasetCreate
     * @summary Add swift source for a dataset
     * @request POST:/source/swift/dataset/{datasetName}
     */
    swiftDatasetCreate: (datasetName: string, request: DatasourceSwiftRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/swift/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name UptoboxDatasetCreate
     * @summary Add uptobox source for a dataset
     * @request POST:/source/uptobox/dataset/{datasetName}
     */
    uptoboxDatasetCreate: (datasetName: string, request: DatasourceUptoboxRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/uptobox/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name WebdavDatasetCreate
     * @summary Add webdav source for a dataset
     * @request POST:/source/webdav/dataset/{datasetName}
     */
    webdavDatasetCreate: (datasetName: string, request: DatasourceWebdavRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/webdav/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name YandexDatasetCreate
     * @summary Add yandex source for a dataset
     * @request POST:/source/yandex/dataset/{datasetName}
     */
    yandexDatasetCreate: (datasetName: string, request: DatasourceYandexRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/yandex/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name ZohoDatasetCreate
     * @summary Add zoho source for a dataset
     * @request POST:/source/zoho/dataset/{datasetName}
     */
    zohoDatasetCreate: (datasetName: string, request: DatasourceZohoRequest, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/zoho/dataset/${datasetName}`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name SourceDelete
     * @summary Remove a source
     * @request DELETE:/source/{id}
     */
    sourceDelete: (id: string, params: RequestParams = {}) =>
      this.request<void, ApiHTTPError>({
        path: `/source/${id}`,
        method: "DELETE",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name SourcePartialUpdate
     * @summary Update the config options of a source
     * @request PATCH:/source/{id}
     */
    sourcePartialUpdate: (id: string, config: DatasourceAllConfig, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/${id}`,
        method: "PATCH",
        body: config,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name CheckCreate
     * @summary Check the connection of the data source by listing a path
     * @request POST:/source/{id}/check
     */
    checkCreate: (id: string, request: DatasourceCheckSourceRequest, params: RequestParams = {}) =>
      this.request<GithubComDataPreservationProgramsSingularityHandlerDatasourceEntry[], ApiHTTPError>({
        path: `/source/${id}/check`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name ChunksDetail
     * @summary Get all dag details of a data source
     * @request GET:/source/{id}/chunks
     */
    chunksDetail: (id: string, params: RequestParams = {}) =>
      this.request<ModelCar[], ApiHTTPError>({
        path: `/source/${id}/chunks`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name DaggenCreate
     * @summary Mark a source as ready for DAG generation
     * @request POST:/source/{id}/daggen
     */
    daggenCreate: (id: string, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/${id}/daggen`,
        method: "POST",
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name ItemsDetail
     * @summary Get all item details of a data source
     * @request GET:/source/{id}/items
     */
    itemsDetail: (id: string, params: RequestParams = {}) =>
      this.request<ModelItem[], ApiHTTPError>({
        path: `/source/${id}/items`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name PathDetail
     * @summary Get all item details inside a data source path
     * @request GET:/source/{id}/path
     */
    pathDetail: (id: string, request: InspectGetPathRequest, params: RequestParams = {}) =>
      this.request<InspectDirDetail, ApiHTTPError>({
        path: `/source/${id}/path`,
        method: "GET",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description Tells Singularity that something is ready to be grabbed for data preparation
     *
     * @tags Data Source
     * @name PushCreate
     * @summary Push an item to be queued
     * @request POST:/source/{id}/push
     */
    pushCreate: (id: string, item: DatasourceItemInfo, params: RequestParams = {}) =>
      this.request<ModelItem, string>({
        path: `/source/${id}/push`,
        method: "POST",
        body: item,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name RescanCreate
     * @summary Trigger a rescan of a data source
     * @request POST:/source/{id}/rescan
     */
    rescanCreate: (id: string, params: RequestParams = {}) =>
      this.request<ModelSource, ApiHTTPError>({
        path: `/source/${id}/rescan`,
        method: "POST",
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Data Source
     * @name SummaryDetail
     * @summary Get the data preparation summary of a data source
     * @request GET:/source/{id}/summary
     */
    summaryDetail: (id: string, params: RequestParams = {}) =>
      this.request<DatasourceChunksByState, ApiHTTPError>({
        path: `/source/${id}/summary`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  wallet = {
    /**
     * No description
     *
     * @tags Wallet
     * @name WalletList
     * @summary List all imported wallets
     * @request GET:/wallet
     */
    walletList: (params: RequestParams = {}) =>
      this.request<ModelWallet[], ApiHTTPError>({
        path: `/wallet`,
        method: "GET",
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Wallet
     * @name WalletCreate
     * @summary Import a private key
     * @request POST:/wallet
     */
    walletCreate: (request: WalletImportRequest, params: RequestParams = {}) =>
      this.request<ModelWallet, ApiHTTPError>({
        path: `/wallet`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Wallet
     * @name RemoteCreate
     * @summary Add a remote wallet
     * @request POST:/wallet/remote
     */
    remoteCreate: (request: WalletAddRemoteRequest, params: RequestParams = {}) =>
      this.request<ModelWallet, ApiHTTPError>({
        path: `/wallet/remote`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Wallet
     * @name WalletDelete
     * @summary Remove a wallet
     * @request DELETE:/wallet/{address}
     */
    walletDelete: (address: string, params: RequestParams = {}) =>
      this.request<void, ApiHTTPError>({
        path: `/wallet/${address}`,
        method: "DELETE",
        ...params,
      }),
  };
}
