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

export interface DataprepAddPieceRequest {
  /** Path to the CAR file, used to determine the size of the file and root CID */
  filePath?: string;
  /** CID of the piece */
  pieceCid?: string;
  /** Size of the piece */
  pieceSize?: string;
  /** Root CID of the CAR file, if not provided, will be determined by the CAR file header. Used to populate the label field of storage deal */
  rootCid?: string;
}

export interface DataprepCreateRequest {
  /**
   * Whether to delete the source files after export
   * @default false
   */
  deleteAfterExport?: boolean;
  /**
   * Maximum size of the CAR files to be created
   * @default "31.5GiB"
   */
  maxSize: string;
  /** Name of the preparation */
  name: string;
  /** Name of Output storage systems to be used for the output */
  outputStorages?: string[];
  /** Target piece size of the CAR files used for piece commitment calculation */
  pieceSize?: string;
  /** Name of Source storage systems to be used for the source */
  sourceStorages: string[];
}

export interface DataprepDirEntry {
  cid?: string;
  fileVersions?: DataprepVersion[];
  isDir?: boolean;
  path?: string;
}

export interface DataprepExploreResult {
  cid?: string;
  path?: string;
  subEntries?: DataprepDirEntry[];
}

export interface DataprepPieceList {
  attachmentId?: number;
  pieces?: ModelCar[];
  source?: ModelStorage;
  storageId?: number;
}

export interface DataprepSourceStatus {
  attachmentId?: number;
  jobs?: ModelJob[];
  source?: ModelStorage;
  storageId?: number;
}

export interface DataprepVersion {
  cid?: string;
  hash?: string;
  id?: number;
  lastModified?: string;
  size?: number;
}

export interface DealListDealRequest {
  /** preparation ID or name filter */
  preparations?: string[];
  /** provider filter */
  providers?: string[];
  /** schedule id filter */
  schedules?: number[];
  /** source ID or name filter */
  sources?: string[];
  /** state filter */
  states?: ModelDealState[];
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

export type ModelCID = object;

export interface ModelCar {
  attachmentId?: number;
  createdAt?: string;
  fileSize?: number;
  id?: number;
  jobId?: number;
  pieceCid?: ModelCID;
  pieceSize?: number;
  /** Association */
  preparationId?: number;
  rootCid?: ModelCID;
  storageId?: number;
  /** StoragePath is the path to the CAR file inside the storage. If the StorageID is nil but StoragePath is not empty, it means the CAR file is stored at the local absolute path. */
  storagePath?: string;
}

export type ModelConfigMap = Record<string, string>;

export interface ModelDeal {
  clientId?: string;
  createdAt?: string;
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
  /** Associations */
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

export interface ModelJob {
  attachmentId?: number;
  errorMessage?: string;
  errorStackTrace?: string;
  id?: number;
  state?: ModelJobState;
  type?: ModelJobType;
  /** Associations */
  workerId?: string;
}

export enum ModelJobState {
  Created = "created",
  Ready = "ready",
  Paused = "paused",
  Processing = "processing",
  Complete = "complete",
  Error = "error",
}

export enum ModelJobType {
  Scan = "scan",
  Pack = "pack",
  DagGen = "daggen",
}

export interface ModelPreparation {
  createdAt?: string;
  /** DeleteAfterExport is a flag that indicates whether the source files should be deleted after export. */
  deleteAfterExport?: boolean;
  id?: number;
  maxSize?: number;
  name?: string;
  pieceSize?: number;
  updatedAt?: string;
}

export interface ModelSchedule {
  allowedPieceCids?: string[];
  announceToIpni?: boolean;
  createdAt?: string;
  duration?: number;
  errorMessage?: string;
  httpHeaders?: string[];
  id?: number;
  keepUnsealed?: boolean;
  maxPendingDealNumber?: number;
  maxPendingDealSize?: number;
  notes?: string;
  /** Associations */
  preparationId?: number;
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

export interface ModelStorage {
  /** Config is a map of key-value pairs that can be used to store RClone options. */
  config?: ModelConfigMap;
  createdAt?: string;
  id?: number;
  name?: string;
  /** Path is the path to the storage root. */
  path?: string;
  type?: string;
  updatedAt?: string;
}

export interface ModelWallet {
  /** Address is the Filecoin full address of the wallet */
  address?: string;
  /** ID is the short ID of the wallet */
  id?: string;
  /** PrivateKey is the private key of the wallet */
  privateKey?: string;
}

export interface ScheduleCreateRequest {
  /** Allowed piece CIDs in this schedule */
  allowedPieceCids?: string[];
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
  /** Preparation ID or name */
  preparation?: string;
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

export interface StorageAcdConfig {
  /** Auth server URL. */
  authUrl?: string;
  /** Checkpoint for internal polling (debug). */
  checkpoint?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
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

export interface StorageAzureblobConfig {
  /** Access tier of blob: hot, cool or archive. */
  accessTier?: string;
  /** Azure Storage Account Name. */
  account?: string;
  /**
   * Delete archive tier blobs before overwriting.
   * @default false
   */
  archiveTierDelete?: boolean;
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
   * @default false
   */
  clientSendCertificateChain?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"
   */
  encoding?: string;
  /** Endpoint for the service. */
  endpoint?: string;
  /**
   * Read credentials from runtime (environment variables, CLI or MSI).
   * @default false
   */
  envAuth?: boolean;
  /** Storage Account Shared Key. */
  key?: string;
  /**
   * Size of blob list.
   * @default 5000
   */
  listChunk?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /** Object ID of the user-assigned MSI to use, if any. */
  msiClientId?: string;
  /** Azure resource ID of the user-assigned MSI to use, if any. */
  msiMiResId?: string;
  /** Object ID of the user-assigned MSI to use, if any. */
  msiObjectId?: string;
  /**
   * If set, don't attempt to check the container exists or create it.
   * @default false
   */
  noCheckContainer?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /** The user's password */
  password?: string;
  /**
   * Public access level of a container: blob or container.
   * @example ""
   */
  publicAccess?: string;
  /** SAS URL for container level access only. */
  sasUrl?: string;
  /** Path to file containing credentials for use with a service principal. */
  servicePrincipalFile?: string;
  /** ID of the service principal's tenant. Also called its directory ID. */
  tenant?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 16
   */
  uploadConcurrency?: number;
  /** Cutoff for switching to chunked upload (<= 256 MiB) (deprecated). */
  uploadCutoff?: string;
  /**
   * Uses local storage emulator if provided as 'true'.
   * @default false
   */
  useEmulator?: boolean;
  /**
   * Use a managed service identity to authenticate (only works in Azure).
   * @default false
   */
  useMsi?: boolean;
  /** User name (usually an email address) */
  username?: string;
}

export interface StorageB2Config {
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
  /**
   * Disable checksums for large (> upload cutoff) files.
   * @default false
   */
  disableChecksum?: boolean;
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
   * @default false
   */
  hardDelete?: boolean;
  /** Application Key. */
  key?: string;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
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
   * @default false
   */
  versions?: boolean;
}

export interface StorageBoxConfig {
  /** Box App Primary Access Token */
  accessToken?: string;
  /** Auth server URL. */
  authUrl?: string;
  /** Box App config.json location */
  boxConfigFile?: string;
  /**
   * @default "user"
   * @example "user"
   */
  boxSubType?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /**
   * Max number of times to try committing a multipart file.
   * @default 100
   */
  commitRetries?: number;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Size of listing chunk 1-1000.
   * @default 1000
   */
  listChunk?: number;
  /** Only show items owned by the login (email address) passed in. */
  ownedBy?: string;
  /**
   * Fill in for rclone to use a non root folder as its starting point.
   * @default "0"
   */
  rootFolderId?: string;
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

export interface StorageCreateAcdStorageRequest {
  config?: StorageAcdConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateAzureblobStorageRequest {
  config?: StorageAzureblobConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateB2StorageRequest {
  config?: StorageB2Config;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateBoxStorageRequest {
  config?: StorageBoxConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateDriveStorageRequest {
  config?: StorageDriveConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateDropboxStorageRequest {
  config?: StorageDropboxConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateFichierStorageRequest {
  config?: StorageFichierConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateFilefabricStorageRequest {
  config?: StorageFilefabricConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateFtpStorageRequest {
  config?: StorageFtpConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateGcsStorageRequest {
  config?: StorageGcsConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateGphotosStorageRequest {
  config?: StorageGphotosConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateHdfsStorageRequest {
  config?: StorageHdfsConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateHidriveStorageRequest {
  config?: StorageHidriveConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateHttpStorageRequest {
  config?: StorageHttpConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateInternetarchiveStorageRequest {
  config?: StorageInternetarchiveConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateJottacloudStorageRequest {
  config?: StorageJottacloudConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateKoofrDigistorageStorageRequest {
  config?: StorageKoofrDigistorageConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateKoofrKoofrStorageRequest {
  config?: StorageKoofrKoofrConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateKoofrOtherStorageRequest {
  config?: StorageKoofrOtherConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateLocalStorageRequest {
  config?: StorageLocalConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateMailruStorageRequest {
  config?: StorageMailruConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateMegaStorageRequest {
  config?: StorageMegaConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateNetstorageStorageRequest {
  config?: StorageNetstorageConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateOnedriveStorageRequest {
  config?: StorageOnedriveConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateOosEnvAuthStorageRequest {
  config?: StorageOosEnvAuthConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateOosInstancePrincipalAuthStorageRequest {
  config?: StorageOosInstancePrincipalAuthConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateOosNoAuthStorageRequest {
  config?: StorageOosNoAuthConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateOosResourcePrincipalAuthStorageRequest {
  config?: StorageOosResourcePrincipalAuthConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateOosUserPrincipalAuthStorageRequest {
  config?: StorageOosUserPrincipalAuthConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateOpendriveStorageRequest {
  config?: StorageOpendriveConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreatePcloudStorageRequest {
  config?: StoragePcloudConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreatePremiumizemeStorageRequest {
  config?: StoragePremiumizemeConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreatePutioStorageRequest {
  config?: StoragePutioConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateQingstorStorageRequest {
  config?: StorageQingstorConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateRequest {
  config?: Record<string, string>;
  name: string;
  path: string;
  provider?: string;
}

export interface StorageCreateS3AWSStorageRequest {
  config?: StorageS3AWSConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3AlibabaStorageRequest {
  config?: StorageS3AlibabaConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3ArvanCloudStorageRequest {
  config?: StorageS3ArvanCloudConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3CephStorageRequest {
  config?: StorageS3CephConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3ChinaMobileStorageRequest {
  config?: StorageS3ChinaMobileConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3CloudflareStorageRequest {
  config?: StorageS3CloudflareConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3DigitalOceanStorageRequest {
  config?: StorageS3DigitalOceanConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3DreamhostStorageRequest {
  config?: StorageS3DreamhostConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3HuaweiOBSStorageRequest {
  config?: StorageS3HuaweiOBSConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3IBMCOSStorageRequest {
  config?: StorageS3IBMCOSConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3IDriveStorageRequest {
  config?: StorageS3IDriveConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3IONOSStorageRequest {
  config?: StorageS3IONOSConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3LiaraStorageRequest {
  config?: StorageS3LiaraConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3LyveCloudStorageRequest {
  config?: StorageS3LyveCloudConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3MinioStorageRequest {
  config?: StorageS3MinioConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3NeteaseStorageRequest {
  config?: StorageS3NeteaseConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3OtherStorageRequest {
  config?: StorageS3OtherConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3QiniuStorageRequest {
  config?: StorageS3QiniuConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3RackCorpStorageRequest {
  config?: StorageS3RackCorpConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3ScalewayStorageRequest {
  config?: StorageS3ScalewayConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3SeaweedFSStorageRequest {
  config?: StorageS3SeaweedFSConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3StackPathStorageRequest {
  config?: StorageS3StackPathConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3StorjStorageRequest {
  config?: StorageS3StorjConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3TencentCOSStorageRequest {
  config?: StorageS3TencentCOSConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateS3WasabiStorageRequest {
  config?: StorageS3WasabiConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateSeafileStorageRequest {
  config?: StorageSeafileConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateSftpStorageRequest {
  config?: StorageSftpConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateSharefileStorageRequest {
  config?: StorageSharefileConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateSiaStorageRequest {
  config?: StorageSiaConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateSmbStorageRequest {
  config?: StorageSmbConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateStorjExistingStorageRequest {
  config?: StorageStorjExistingConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateStorjNewStorageRequest {
  config?: StorageStorjNewConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateSugarsyncStorageRequest {
  config?: StorageSugarsyncConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateSwiftStorageRequest {
  config?: StorageSwiftConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateUptoboxStorageRequest {
  config?: StorageUptoboxConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateWebdavStorageRequest {
  config?: StorageWebdavConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateYandexStorageRequest {
  config?: StorageYandexConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageCreateZohoStorageRequest {
  config?: StorageZohoConfig;
  /**
   * Name of the storage, must be unique
   * @example "my-storage"
   */
  name?: string;
  /** Path of the storage */
  path?: string;
}

export interface StorageDirEntry {
  dirId?: string;
  hash?: string;
  isDir?: boolean;
  lastModified?: string;
  numItems?: number;
  path?: string;
  size?: number;
}

export interface StorageDriveConfig {
  /**
   * Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
   * @default false
   */
  acknowledgeAbuse?: boolean;
  /**
   * Allow the filetype to change when uploading Google docs.
   * @default false
   */
  allowImportNameChange?: boolean;
  /**
   * Deprecated: No longer needed.
   * @default false
   */
  alternateExport?: boolean;
  /**
   * Only consider files owned by the authenticated user.
   * @default false
   */
  authOwnerOnly?: boolean;
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
   * @default false
   */
  copyShortcutContent?: boolean;
  /**
   * Disable drive using http2.
   * @default true
   */
  disableHttp2?: boolean;
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
   * @default false
   */
  keepRevisionForever?: boolean;
  /**
   * Size of listing chunk 100-1000, 0 to disable.
   * @default 1000
   */
  listChunk?: number;
  /**
   * Number of API calls to allow without sleeping.
   * @default 100
   */
  pacerBurst?: number;
  /**
   * Minimum time to sleep between API calls.
   * @default "100ms"
   */
  pacerMinSleep?: string;
  /** Resource key for accessing a link-shared file. */
  resourceKey?: string;
  /** ID of the root folder. */
  rootFolderId?: string;
  /**
   * Scope that rclone should use when requesting access from drive.
   * @example "drive"
   */
  scope?: string;
  /**
   * Allow server-side operations (e.g. copy) to work across different drive configs.
   * @default false
   */
  serverSideAcrossConfigs?: boolean;
  /** Service Account Credentials JSON blob. */
  serviceAccountCredentials?: string;
  /** Service Account Credentials JSON file path. */
  serviceAccountFile?: string;
  /**
   * Only show files that are shared with me.
   * @default false
   */
  sharedWithMe?: boolean;
  /**
   * Show sizes as storage quota usage, not actual size.
   * @default false
   */
  sizeAsQuota?: boolean;
  /**
   * Skip MD5 checksum on Google photos and videos only.
   * @default false
   */
  skipChecksumGphotos?: boolean;
  /**
   * If set skip dangling shortcut files.
   * @default false
   */
  skipDanglingShortcuts?: boolean;
  /**
   * Skip google documents in all listings.
   * @default false
   */
  skipGdocs?: boolean;
  /**
   * If set skip shortcut files.
   * @default false
   */
  skipShortcuts?: boolean;
  /**
   * Only show files that are starred.
   * @default false
   */
  starredOnly?: boolean;
  /**
   * Make download limit errors be fatal.
   * @default false
   */
  stopOnDownloadLimit?: boolean;
  /**
   * Make upload limit errors be fatal.
   * @default false
   */
  stopOnUploadLimit?: boolean;
  /** ID of the Shared Drive (Team Drive). */
  teamDrive?: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
  /**
   * Only show files that are in the trash.
   * @default false
   */
  trashedOnly?: boolean;
  /**
   * Cutoff for switching to chunked upload.
   * @default "8Mi"
   */
  uploadCutoff?: string;
  /**
   * Use file created date instead of modified date.
   * @default false
   */
  useCreatedDate?: boolean;
  /**
   * Use date file was shared instead of modified date.
   * @default false
   */
  useSharedDate?: boolean;
  /**
   * Send files to the trash instead of deleting permanently.
   * @default true
   */
  useTrash?: boolean;
  /**
   * If Object's are greater, use drive v2 API to download.
   * @default "off"
   */
  v2DownloadMinSize?: string;
}

export interface StorageDropboxConfig {
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
   * @default 0
   */
  batchSize?: number;
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
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Impersonate this user when using a business account. */
  impersonate?: string;
  /**
   * Instructs rclone to work on individual shared files.
   * @default false
   */
  sharedFiles?: boolean;
  /**
   * Instructs rclone to work on shared folders.
   * @default false
   */
  sharedFolders?: boolean;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
}

export interface StorageFichierConfig {
  /** Your API Key, get it from https://1fichier.com/console/params.pl. */
  apiKey?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** If you want to download a shared file that is password protected, add this parameter. */
  filePassword?: string;
  /** If you want to list the files in a shared folder that is password protected, add this parameter. */
  folderPassword?: string;
  /** If you want to download a shared folder, add this parameter. */
  sharedFolder?: string;
}

export interface StorageFilefabricConfig {
  /**
   * The encoding for the backend.
   * @default "Slash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Permanent Authentication Token. */
  permanentToken?: string;
  /** ID of the root folder. */
  rootFolderId?: string;
  /** Session Token. */
  token?: string;
  /** Token expiry time. */
  tokenExpiry?: string;
  /**
   * URL of the Enterprise File Fabric to connect to.
   * @example "https://storagemadeeasy.com"
   */
  url?: string;
  /** Version read from the file fabric. */
  version?: string;
}

export interface StorageFtpConfig {
  /**
   * Allow asking for FTP password when needed.
   * @default false
   */
  askPassword?: boolean;
  /**
   * Maximum time to wait for a response to close.
   * @default "1m0s"
   */
  closeTimeout?: string;
  /**
   * Maximum number of FTP simultaneous connections, 0 for unlimited.
   * @default 0
   */
  concurrency?: number;
  /**
   * Disable using EPSV even if server advertises support.
   * @default false
   */
  disableEpsv?: boolean;
  /**
   * Disable using MLSD even if server advertises support.
   * @default false
   */
  disableMlsd?: boolean;
  /**
   * Disable TLS 1.3 (workaround for FTP servers with buggy TLS)
   * @default false
   */
  disableTls13?: boolean;
  /**
   * Disable using UTF-8 even if server advertises support.
   * @default false
   */
  disableUtf8?: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,Del,Ctl,RightSpace,Dot"
   * @example "Asterisk,Ctl,Dot,Slash"
   */
  encoding?: string;
  /**
   * Use Explicit FTPS (FTP over TLS).
   * @default false
   */
  explicitTls?: boolean;
  /**
   * Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD.
   * @default false
   */
  forceListHidden?: boolean;
  /** FTP host to connect to. */
  host?: string;
  /**
   * Max time before closing idle connections.
   * @default "1m0s"
   */
  idleTimeout?: string;
  /**
   * Do not verify the TLS certificate of the server.
   * @default false
   */
  noCheckCertificate?: boolean;
  /** FTP password. */
  pass?: string;
  /**
   * FTP port number.
   * @default 21
   */
  port?: number;
  /**
   * Maximum time to wait for data connection closing status.
   * @default "1m0s"
   */
  shutTimeout?: string;
  /**
   * Use Implicit FTPS (FTP over TLS).
   * @default false
   */
  tls?: boolean;
  /**
   * Size of TLS session cache for all control and data connections.
   * @default 32
   */
  tlsCacheSize?: number;
  /**
   * FTP username.
   * @default "$USER"
   */
  user?: string;
  /**
   * Use MDTM to set modification time (VsFtpd quirk)
   * @default false
   */
  writingMdtm?: boolean;
}

export interface StorageGcsConfig {
  /**
   * Access public buckets and objects without credentials.
   * @default false
   */
  anonymous?: boolean;
  /** Auth server URL. */
  authUrl?: string;
  /**
   * Access Control List for new buckets.
   * @example "authenticatedRead"
   */
  bucketAcl?: string;
  /**
   * Access checks should use bucket-level IAM policies.
   * @default false
   */
  bucketPolicyOnly?: boolean;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /**
   * If set this will decompress gzip encoded objects.
   * @default false
   */
  decompress?: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,CrLf,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Endpoint for the service. */
  endpoint?: string;
  /**
   * Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * Location for the newly created buckets.
   * @example ""
   */
  location?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * Access Control List for new objects.
   * @example "authenticatedRead"
   */
  objectAcl?: string;
  /** Project number. */
  projectNumber?: string;
  /** Service Account Credentials JSON blob. */
  serviceAccountCredentials?: string;
  /** Service Account Credentials JSON file path. */
  serviceAccountFile?: string;
  /**
   * The storage class to use when storing objects in Google Cloud Storage.
   * @example ""
   */
  storageClass?: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
}

export interface StorageGphotosConfig {
  /** Auth server URL. */
  authUrl?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,CrLf,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Also view and download archived media.
   * @default false
   */
  includeArchived?: boolean;
  /**
   * Set to make the Google Photos backend read only.
   * @default false
   */
  readOnly?: boolean;
  /**
   * Set to read the size of media items.
   * @default false
   */
  readSize?: boolean;
  /**
   * Year limits the photos to be downloaded to those which are uploaded after the given year.
   * @default 2000
   */
  startYear?: number;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
}

export interface StorageHdfsConfig {
  /**
   * Kerberos data transfer protection: authentication|integrity|privacy.
   * @example "privacy"
   */
  dataTransferProtection?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Colon,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Hadoop name node and port. */
  namenode?: string;
  /** Kerberos service principal name for the namenode. */
  servicePrincipalName?: string;
  /**
   * Hadoop user name.
   * @example "root"
   */
  username?: string;
}

export interface StorageHidriveConfig {
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
  /**
   * Do not fetch number of objects in directories unless it is absolutely necessary.
   * @default false
   */
  disableFetchingMemberCount?: boolean;
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
  /**
   * The root/parent folder for all paths.
   * @default "/"
   * @example "/"
   */
  rootPrefix?: string;
  /**
   * Access permissions that rclone should use when requesting access from HiDrive.
   * @default "rw"
   * @example "rw"
   */
  scopeAccess?: string;
  /**
   * User-level that rclone should use when requesting access from HiDrive.
   * @default "user"
   * @example "user"
   */
  scopeRole?: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
  /**
   * Concurrency for chunked uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff/Threshold for chunked uploads.
   * @default "96Mi"
   */
  uploadCutoff?: string;
}

export interface StorageHttpConfig {
  /** Set HTTP headers for all transactions. */
  headers?: string;
  /**
   * Don't use HEAD requests.
   * @default false
   */
  noHead?: boolean;
  /**
   * Set this if the site doesn't end directories with /.
   * @default false
   */
  noSlash?: boolean;
  /** URL of HTTP host to connect to. */
  url?: string;
}

export interface StorageInternetarchiveConfig {
  /** IAS3 Access Key. */
  accessKeyId?: string;
  /**
   * Don't ask the server to test against MD5 checksum calculated by rclone.
   * @default true
   */
  disableChecksum?: boolean;
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
  /** IAS3 Secret Key (password). */
  secretAccessKey?: string;
  /**
   * Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish.
   * @default "0s"
   */
  waitArchive?: string;
}

export interface StorageJottacloudConfig {
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Delete files permanently rather than putting them into the trash.
   * @default false
   */
  hardDelete?: boolean;
  /**
   * Files bigger than this will be cached on disk to calculate the MD5 if required.
   * @default "10Mi"
   */
  md5MemoryLimit?: string;
  /**
   * Avoid server side versioning by deleting files and recreating files instead of overwriting them.
   * @default false
   */
  noVersions?: boolean;
  /**
   * Only show files that are in the trash.
   * @default false
   */
  trashedOnly?: boolean;
  /**
   * Files bigger than this can be resumed if the upload fail's.
   * @default "10Mi"
   */
  uploadResumeLimit?: string;
}

export interface StorageKoofrDigistorageConfig {
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Mount ID of the mount to use. */
  mountid?: string;
  /** Your password for rclone (generate one at https://storage.rcs-rds.ro/app/admin/preferences/password). */
  password?: string;
  /**
   * Does the backend support setting modification time.
   * @default true
   */
  setmtime?: boolean;
  /** Your user name. */
  user?: string;
}

export interface StorageKoofrKoofrConfig {
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Mount ID of the mount to use. */
  mountid?: string;
  /** Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password). */
  password?: string;
  /**
   * Does the backend support setting modification time.
   * @default true
   */
  setmtime?: boolean;
  /** Your user name. */
  user?: string;
}

export interface StorageKoofrOtherConfig {
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** The Koofr API endpoint to use. */
  endpoint?: string;
  /** Mount ID of the mount to use. */
  mountid?: string;
  /** Your password for rclone (generate one at your service's settings page). */
  password?: string;
  /**
   * Does the backend support setting modification time.
   * @default true
   */
  setmtime?: boolean;
  /** Your user name. */
  user?: string;
}

export interface StorageLocalConfig {
  /**
   * Force the filesystem to report itself as case insensitive.
   * @default false
   */
  caseInsensitive?: boolean;
  /**
   * Force the filesystem to report itself as case sensitive.
   * @default false
   */
  caseSensitive?: boolean;
  /**
   * Follow symlinks and copy the pointed to item.
   * @default false
   */
  copyLinks?: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,Dot"
   */
  encoding?: string;
  /**
   * Translate symlinks to/from regular files with a '.rclonelink' extension.
   * @default false
   */
  links?: boolean;
  /**
   * Don't check to see if the files change during upload.
   * @default false
   */
  noCheckUpdated?: boolean;
  /**
   * Disable preallocation of disk space for transferred files.
   * @default false
   */
  noPreallocate?: boolean;
  /**
   * Disable setting modtime.
   * @default false
   */
  noSetModtime?: boolean;
  /**
   * Disable sparse files for multi-thread downloads.
   * @default false
   */
  noSparse?: boolean;
  /**
   * Disable UNC (long path names) conversion on Windows.
   * @default false
   * @example true
   */
  nounc?: boolean;
  /**
   * Don't cross filesystem boundaries (unix/macOS only).
   * @default false
   */
  oneFileSystem?: boolean;
  /**
   * Don't warn about skipped symlinks.
   * @default false
   */
  skipLinks?: boolean;
  /**
   * Apply unicode NFC normalization to paths and filenames.
   * @default false
   */
  unicodeNormalization?: boolean;
  /**
   * Assume the Stat size of links is zero (and read them instead) (deprecated).
   * @default false
   */
  zeroSizeLinks?: boolean;
}

export interface StorageMailruConfig {
  /**
   * What should copy do if file checksum is mismatched or invalid.
   * @default true
   * @example true
   */
  checkHash?: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Password. */
  pass?: string;
  /** Comma separated list of internal maintenance flags. */
  quirks?: string;
  /**
   * Skip full upload if there is another file with same data hash.
   * @default true
   * @example true
   */
  speedupEnable?: boolean;
  /**
   * Comma separated list of file name patterns eligible for speedup (put by hash).
   * @default "*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf"
   * @example ""
   */
  speedupFilePatterns?: string;
  /**
   * This option allows you to disable speedup (put by hash) for large files.
   * @default "3Gi"
   * @example "0"
   */
  speedupMaxDisk?: string;
  /**
   * Files larger than the size given below will always be hashed on disk.
   * @default "32Mi"
   * @example "0"
   */
  speedupMaxMemory?: string;
  /** User name (usually email). */
  user?: string;
  /** HTTP user agent used internally by client. */
  userAgent?: string;
}

export interface StorageMegaConfig {
  /**
   * Output more debug from Mega.
   * @default false
   */
  debug?: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Delete files permanently rather than putting them into the trash.
   * @default false
   */
  hardDelete?: boolean;
  /** Password. */
  pass?: string;
  /**
   * Use HTTPS for transfers.
   * @default false
   */
  useHttps?: boolean;
  /** User name. */
  user?: string;
}

export interface StorageNetstorageConfig {
  /** Set the NetStorage account name */
  account?: string;
  /** Domain+path of NetStorage host to connect to. */
  host?: string;
  /**
   * Select between HTTP or HTTPS protocol.
   * @default "https"
   * @example "http"
   */
  protocol?: string;
  /** Set the NetStorage account secret/G2O key for authentication. */
  secret?: string;
}

export interface StorageOnedriveConfig {
  /**
   * Set scopes to be requested by rclone.
   * @default "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"
   * @example "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"
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
  /**
   * Disable the request for Sites.Read.All permission.
   * @default false
   */
  disableSitePermission?: boolean;
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
   * @default false
   */
  exposeOnenoteFiles?: boolean;
  /**
   * Specify the hash in use for the backend.
   * @default "auto"
   * @example "auto"
   */
  hashType?: string;
  /** Set the password for links created by the link command. */
  linkPassword?: string;
  /**
   * Set the scope of the links created by the link command.
   * @default "anonymous"
   * @example "anonymous"
   */
  linkScope?: string;
  /**
   * Set the type of the links created by the link command.
   * @default "view"
   * @example "view"
   */
  linkType?: string;
  /**
   * Size of listing chunk.
   * @default 1000
   */
  listChunk?: number;
  /**
   * Remove all versions on modifying operations.
   * @default false
   */
  noVersions?: boolean;
  /**
   * Choose national cloud region for OneDrive.
   * @default "global"
   * @example "global"
   */
  region?: string;
  /** ID of the root folder. */
  rootFolderId?: string;
  /**
   * Allow server-side operations (e.g. copy) to work across different onedrive configs.
   * @default false
   */
  serverSideAcrossConfigs?: boolean;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
}

export interface StorageOosEnvAuthConfig {
  /**
   * Chunk size to use for uploading.
   * @default "5Mi"
   */
  chunkSize?: string;
  /** Object storage compartment OCID */
  compartment?: string;
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
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Endpoint for Object storage API. */
  endpoint?: string;
  /**
   * If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
   * @default false
   */
  leavePartsOnError?: boolean;
  /** Object storage namespace */
  namespace?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /** Object storage Region */
  region?: string;
  /**
   * If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
   * @example ""
   */
  sseCustomerAlgorithm?: string;
  /**
   * To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
   * @example ""
   */
  sseCustomerKey?: string;
  /**
   * To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
   * @example ""
   */
  sseCustomerKeyFile?: string;
  /**
   * If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
   * @example ""
   */
  sseCustomerKeySha256?: string;
  /**
   * if using using your own master key in vault, this header specifies the
   * @example ""
   */
  sseKmsKeyId?: string;
  /**
   * The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
   * @default "Standard"
   * @example "Standard"
   */
  storageTier?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 10
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
}

export interface StorageOosInstancePrincipalAuthConfig {
  /**
   * Chunk size to use for uploading.
   * @default "5Mi"
   */
  chunkSize?: string;
  /** Object storage compartment OCID */
  compartment?: string;
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
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Endpoint for Object storage API. */
  endpoint?: string;
  /**
   * If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
   * @default false
   */
  leavePartsOnError?: boolean;
  /** Object storage namespace */
  namespace?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /** Object storage Region */
  region?: string;
  /**
   * If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
   * @example ""
   */
  sseCustomerAlgorithm?: string;
  /**
   * To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
   * @example ""
   */
  sseCustomerKey?: string;
  /**
   * To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
   * @example ""
   */
  sseCustomerKeyFile?: string;
  /**
   * If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
   * @example ""
   */
  sseCustomerKeySha256?: string;
  /**
   * if using using your own master key in vault, this header specifies the
   * @example ""
   */
  sseKmsKeyId?: string;
  /**
   * The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
   * @default "Standard"
   * @example "Standard"
   */
  storageTier?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 10
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
}

export interface StorageOosNoAuthConfig {
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
   * Timeout for copy.
   * @default "1m0s"
   */
  copyTimeout?: string;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Endpoint for Object storage API. */
  endpoint?: string;
  /**
   * If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
   * @default false
   */
  leavePartsOnError?: boolean;
  /** Object storage namespace */
  namespace?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /** Object storage Region */
  region?: string;
  /**
   * If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
   * @example ""
   */
  sseCustomerAlgorithm?: string;
  /**
   * To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
   * @example ""
   */
  sseCustomerKey?: string;
  /**
   * To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
   * @example ""
   */
  sseCustomerKeyFile?: string;
  /**
   * If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
   * @example ""
   */
  sseCustomerKeySha256?: string;
  /**
   * if using using your own master key in vault, this header specifies the
   * @example ""
   */
  sseKmsKeyId?: string;
  /**
   * The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
   * @default "Standard"
   * @example "Standard"
   */
  storageTier?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 10
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
}

export interface StorageOosResourcePrincipalAuthConfig {
  /**
   * Chunk size to use for uploading.
   * @default "5Mi"
   */
  chunkSize?: string;
  /** Object storage compartment OCID */
  compartment?: string;
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
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Endpoint for Object storage API. */
  endpoint?: string;
  /**
   * If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
   * @default false
   */
  leavePartsOnError?: boolean;
  /** Object storage namespace */
  namespace?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /** Object storage Region */
  region?: string;
  /**
   * If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
   * @example ""
   */
  sseCustomerAlgorithm?: string;
  /**
   * To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
   * @example ""
   */
  sseCustomerKey?: string;
  /**
   * To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
   * @example ""
   */
  sseCustomerKeyFile?: string;
  /**
   * If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
   * @example ""
   */
  sseCustomerKeySha256?: string;
  /**
   * if using using your own master key in vault, this header specifies the
   * @example ""
   */
  sseKmsKeyId?: string;
  /**
   * The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
   * @default "Standard"
   * @example "Standard"
   */
  storageTier?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 10
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
}

export interface StorageOosUserPrincipalAuthConfig {
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
   * @example "~/.oci/config"
   */
  configFile?: string;
  /**
   * Profile name inside the oci config file
   * @default "Default"
   * @example "Default"
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
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Endpoint for Object storage API. */
  endpoint?: string;
  /**
   * If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
   * @default false
   */
  leavePartsOnError?: boolean;
  /** Object storage namespace */
  namespace?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /** Object storage Region */
  region?: string;
  /**
   * If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm.
   * @example ""
   */
  sseCustomerAlgorithm?: string;
  /**
   * To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to
   * @example ""
   */
  sseCustomerKey?: string;
  /**
   * To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated
   * @example ""
   */
  sseCustomerKeyFile?: string;
  /**
   * If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption
   * @example ""
   */
  sseCustomerKeySha256?: string;
  /**
   * if using using your own master key in vault, this header specifies the
   * @example ""
   */
  sseKmsKeyId?: string;
  /**
   * The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm
   * @default "Standard"
   * @example "Standard"
   */
  storageTier?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 10
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
}

export interface StorageOpendriveConfig {
  /**
   * Files will be uploaded in chunks this size.
   * @default "10Mi"
   */
  chunkSize?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Password. */
  password?: string;
  /** Username. */
  username?: string;
}

export interface StoragePcloudConfig {
  /** Auth server URL. */
  authUrl?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Hostname to connect to.
   * @default "api.pcloud.com"
   * @example "api.pcloud.com"
   */
  hostname?: string;
  /** Your pcloud password. */
  password?: string;
  /**
   * Fill in for rclone to use a non root folder as its starting point.
   * @default "d0"
   */
  rootFolderId?: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
  /** Your pcloud username. */
  username?: string;
}

export interface StoragePremiumizemeConfig {
  /** API Key. */
  apiKey?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
}

export interface StoragePutioConfig {
  /**
   * The encoding for the backend.
   * @default "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
}

export interface StorageQingstorConfig {
  /** QingStor Access Key ID. */
  accessKeyId?: string;
  /**
   * Chunk size to use for uploading.
   * @default "4Mi"
   */
  chunkSize?: string;
  /**
   * Number of connection retries.
   * @default 3
   */
  connectionRetries?: number;
  /**
   * The encoding for the backend.
   * @default "Slash,Ctl,InvalidUtf8"
   */
  encoding?: string;
  /** Enter an endpoint URL to connection QingStor API. */
  endpoint?: string;
  /**
   * Get QingStor credentials from runtime.
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /** QingStor Secret Access Key (password). */
  secretAccessKey?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 1
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Zone to connect to.
   * @example "pek3a"
   */
  zone?: string;
}

export interface StorageS3AWSConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
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
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery.
   * @default false
   */
  leavePartsOnError?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Location constraint - must be set to match the Region.
   * @example ""
   */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example "us-east-1"
   */
  region?: string;
  /**
   * Enables requester pays option when interacting with S3 bucket.
   * @default false
   */
  requesterPays?: boolean;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /**
   * The server-side encryption algorithm used when storing this object in S3.
   * @example ""
   */
  serverSideEncryption?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
   * @example ""
   */
  sseCustomerAlgorithm?: string;
  /**
   * To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
   * @example ""
   */
  sseCustomerKey?: string;
  /**
   * If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
   * @example ""
   */
  sseCustomerKeyBase64?: string;
  /**
   * If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
   * @example ""
   */
  sseCustomerKeyMd5?: string;
  /**
   * If using KMS ID you must provide the ARN of Key.
   * @example ""
   */
  sseKmsKeyId?: string;
  /**
   * The storage class to use when storing new objects in S3.
   * @example ""
   */
  storageClass?: string;
  /** Endpoint for STS. */
  stsEndpoint?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * If true use the AWS S3 accelerated endpoint.
   * @default false
   */
  useAccelerateEndpoint?: boolean;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3AlibabaConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for OSS API.
   * @example "oss-accelerate.aliyuncs.com"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * The storage class to use when storing new objects in OSS.
   * @example ""
   */
  storageClass?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3ArvanCloudConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for Arvan Cloud Object Storage (AOS) API.
   * @example "s3.ir-thr-at1.arvanstorage.com"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Location constraint - must match endpoint.
   * @example "ir-thr-at1"
   */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * The storage class to use when storing new objects in ArvanCloud.
   * @example "STANDARD"
   */
  storageClass?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3CephConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
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
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /** Location constraint - must be set to match the Region. */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example ""
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /**
   * The server-side encryption algorithm used when storing this object in S3.
   * @example ""
   */
  serverSideEncryption?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
   * @example ""
   */
  sseCustomerAlgorithm?: string;
  /**
   * To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
   * @example ""
   */
  sseCustomerKey?: string;
  /**
   * If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
   * @example ""
   */
  sseCustomerKeyBase64?: string;
  /**
   * If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
   * @example ""
   */
  sseCustomerKeyMd5?: string;
  /**
   * If using KMS ID you must provide the ARN of Key.
   * @example ""
   */
  sseKmsKeyId?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3ChinaMobileConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for China Mobile Ecloud Elastic Object Storage (EOS) API.
   * @example "eos-wuxi-1.cmecloud.cn"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Location constraint - must match endpoint.
   * @example "wuxi1"
   */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /**
   * The server-side encryption algorithm used when storing this object in S3.
   * @example ""
   */
  serverSideEncryption?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
   * @example ""
   */
  sseCustomerAlgorithm?: string;
  /**
   * To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
   * @example ""
   */
  sseCustomerKey?: string;
  /**
   * If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
   * @example ""
   */
  sseCustomerKeyBase64?: string;
  /**
   * If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
   * @example ""
   */
  sseCustomerKeyMd5?: string;
  /**
   * The storage class to use when storing new objects in ChinaMobile.
   * @example ""
   */
  storageClass?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3CloudflareConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
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
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example "auto"
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3DigitalOceanConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for S3 API.
   * @example "syd1.digitaloceanspaces.com"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /** Location constraint - must be set to match the Region. */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example ""
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3DreamhostConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for S3 API.
   * @example "objects-us-east-1.dream.io"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /** Location constraint - must be set to match the Region. */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example ""
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3HuaweiOBSConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for OBS API.
   * @example "obs.af-south-1.myhuaweicloud.com"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to. - the location where your bucket will be created and your data stored. Need bo be same with your endpoint.
   * @example "af-south-1"
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3IBMCOSConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /**
   * Canned ACL used when creating buckets and storing or copying objects.
   * @example "private"
   */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for IBM COS S3 API.
   * @example "s3.us.cloud-object-storage.appdomain.cloud"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Location constraint - must match endpoint when using IBM Cloud Public.
   * @example "us-standard"
   */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example ""
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3IDriveConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3IONOSConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for IONOS S3 Object Storage.
   * @example "s3-eu-central-1.ionoscloud.com"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region where your bucket will be created and your data stored.
   * @example "de"
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3LiaraConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for Liara Object Storage API.
   * @example "storage.iran.liara.space"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * The storage class to use when storing new objects in Liara
   * @example "STANDARD"
   */
  storageClass?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3LyveCloudConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for S3 API.
   * @example "s3.us-east-1.lyvecloud.seagate.com"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /** Location constraint - must be set to match the Region. */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example ""
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3MinioConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
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
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /** Location constraint - must be set to match the Region. */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example ""
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /**
   * The server-side encryption algorithm used when storing this object in S3.
   * @example ""
   */
  serverSideEncryption?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * If using SSE-C, the server-side encryption algorithm used when storing this object in S3.
   * @example ""
   */
  sseCustomerAlgorithm?: string;
  /**
   * To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data.
   * @example ""
   */
  sseCustomerKey?: string;
  /**
   * If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data.
   * @example ""
   */
  sseCustomerKeyBase64?: string;
  /**
   * If using SSE-C you may provide the secret encryption key MD5 checksum (optional).
   * @example ""
   */
  sseCustomerKeyMd5?: string;
  /**
   * If using KMS ID you must provide the ARN of Key.
   * @example ""
   */
  sseKmsKeyId?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3NeteaseConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
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
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /** Location constraint - must be set to match the Region. */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example ""
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3OtherConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
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
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /** Location constraint - must be set to match the Region. */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example ""
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3QiniuConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for Qiniu Object Storage.
   * @example "s3-cn-east-1.qiniucs.com"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Location constraint - must be set to match the Region.
   * @example "cn-east-1"
   */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example "cn-east-1"
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * The storage class to use when storing new objects in Qiniu.
   * @example "STANDARD"
   */
  storageClass?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3RackCorpConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for RackCorp Object Storage.
   * @example "s3.rackcorp.com"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Location constraint - the location where your bucket will be located and your data stored.
   * @example "global"
   */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * region - the location where your bucket will be created and your data stored.
   * @example "global"
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3ScalewayConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for Scaleway Object Storage.
   * @example "s3.nl-ams.scw.cloud"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example "nl-ams"
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * The storage class to use when storing new objects in S3.
   * @example ""
   */
  storageClass?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3SeaweedFSConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for S3 API.
   * @example "localhost:8333"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /** Location constraint - must be set to match the Region. */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example ""
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3StackPathConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for StackPath Object Storage.
   * @example "s3.us-east-2.stackpathstorage.com"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example ""
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3StorjConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for Storj Gateway.
   * @example "gateway.storjshare.io"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3TencentCOSConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /**
   * Canned ACL used when creating buckets and storing or copying objects.
   * @example "default"
   */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for Tencent COS API.
   * @example "cos.ap-beijing.myqcloud.com"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * The storage class to use when storing new objects in Tencent COS.
   * @example ""
   */
  storageClass?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageS3WasabiConfig {
  /** AWS Access Key ID. */
  accessKeyId?: string;
  /** Canned ACL used when creating buckets and storing or copying objects. */
  acl?: string;
  /**
   * Canned ACL used when creating buckets.
   * @example "private"
   */
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
   * @default false
   */
  decompress?: boolean;
  /**
   * Don't store MD5 checksum with object metadata.
   * @default false
   */
  disableChecksum?: boolean;
  /**
   * Disable usage of http2 for S3 backends.
   * @default false
   */
  disableHttp2?: boolean;
  /** Custom endpoint for downloads. */
  downloadUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Endpoint for S3 API.
   * @example "s3.wasabisys.com"
   */
  endpoint?: string;
  /**
   * Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars).
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /**
   * If true use path style access if false use virtual hosted style.
   * @default true
   */
  forcePathStyle?: boolean;
  /**
   * Size of listing chunk (response list for each ListObject S3 request).
   * @default 1000
   */
  listChunk?: number;
  /**
   * Whether to url encode listings: true/false/unset
   * @default "unset"
   */
  listUrlEncode?: string;
  /**
   * Version of ListObjects to use: 1,2 or 0 for auto.
   * @default 0
   */
  listVersion?: number;
  /** Location constraint - must be set to match the Region. */
  locationConstraint?: string;
  /**
   * Maximum number of parts in a multipart upload.
   * @default 10000
   */
  maxUploadParts?: number;
  /**
   * How often internal memory buffer pools will be flushed.
   * @default "1m0s"
   */
  memoryPoolFlushTime?: string;
  /**
   * Whether to use mmap buffers in internal memory pool.
   * @default false
   */
  memoryPoolUseMmap?: boolean;
  /**
   * Set this if the backend might gzip objects.
   * @default "unset"
   */
  mightGzip?: string;
  /**
   * If set, don't attempt to check the bucket exists or create it.
   * @default false
   */
  noCheckBucket?: boolean;
  /**
   * If set, don't HEAD uploaded objects to check integrity.
   * @default false
   */
  noHead?: boolean;
  /**
   * If set, do not do HEAD before GET when getting objects.
   * @default false
   */
  noHeadObject?: boolean;
  /**
   * Suppress setting and reading of system metadata
   * @default false
   */
  noSystemMetadata?: boolean;
  /** Profile to use in the shared credentials file. */
  profile?: string;
  /**
   * Region to connect to.
   * @example ""
   */
  region?: string;
  /** AWS Secret Access Key (password). */
  secretAccessKey?: string;
  /** An AWS session token. */
  sessionToken?: string;
  /** Path to the shared credentials file. */
  sharedCredentialsFile?: string;
  /**
   * Concurrency for multipart uploads.
   * @default 4
   */
  uploadConcurrency?: number;
  /**
   * Cutoff for switching to chunked upload.
   * @default "200Mi"
   */
  uploadCutoff?: string;
  /**
   * Whether to use ETag in multipart uploads for verification
   * @default "unset"
   */
  useMultipartEtag?: string;
  /**
   * Whether to use a presigned request or PutObject for single part uploads
   * @default false
   */
  usePresignedRequest?: boolean;
  /**
   * If true use v2 authentication.
   * @default false
   */
  v2Auth?: boolean;
  /**
   * Show file versions as they were at the specified time.
   * @default "off"
   */
  versionAt?: string;
  /**
   * Include old versions in directory listings.
   * @default false
   */
  versions?: boolean;
}

export interface StorageSeafileConfig {
  /**
   * Two-factor authentication ('true' if the account has 2FA enabled).
   * @default false
   */
  "2fa"?: boolean;
  /** Authentication token. */
  authToken?: string;
  /**
   * Should rclone create a library if it doesn't exist.
   * @default false
   */
  createLibrary?: boolean;
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
  /**
   * URL of seafile host to connect to.
   * @example "https://cloud.seafile.com/"
   */
  url?: string;
  /** User name (usually email address). */
  user?: string;
}

export interface StorageSftpConfig {
  /**
   * Allow asking for SFTP password when needed.
   * @default false
   */
  askPassword?: boolean;
  /**
   * Upload and download chunk size.
   * @default "32Ki"
   */
  chunkSize?: string;
  /** Space separated list of ciphers to be used for session encryption, ordered by preference. */
  ciphers?: string;
  /**
   * The maximum number of outstanding requests for one file
   * @default 64
   */
  concurrency?: number;
  /**
   * If set don't use concurrent reads.
   * @default false
   */
  disableConcurrentReads?: boolean;
  /**
   * If set don't use concurrent writes.
   * @default false
   */
  disableConcurrentWrites?: boolean;
  /**
   * Disable the execution of SSH commands to determine if remote file hashing is available.
   * @default false
   */
  disableHashcheck?: boolean;
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
   * @default false
   */
  keyUseAgent?: boolean;
  /**
   * Optional path to known_hosts file.
   * @example "~/.ssh/known_hosts"
   */
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
   * @default 22
   */
  port?: number;
  /** Optional path to public key file. */
  pubkeyFile?: string;
  /** Specifies the path or command to run a sftp server on the remote host. */
  serverCommand?: string;
  /** Environment variables to pass to sftp and commands */
  setEnv?: string;
  /**
   * Set the modified time on the remote if set.
   * @default true
   */
  setModtime?: boolean;
  /** The command used to read sha1 hashes. */
  sha1sumCommand?: string;
  /**
   * The type of SSH shell on remote server, if any.
   * @example "none"
   */
  shellType?: string;
  /**
   * Set to skip any symlinks and any other non regular files.
   * @default false
   */
  skipLinks?: boolean;
  /**
   * Specifies the SSH2 subsystem on the remote host.
   * @default "sftp"
   */
  subsystem?: string;
  /**
   * If set use fstat instead of stat.
   * @default false
   */
  useFstat?: boolean;
  /**
   * Enable the use of insecure ciphers and key exchange methods.
   * @default false
   * @example false
   */
  useInsecureCipher?: boolean;
  /**
   * SSH username.
   * @default "$USER"
   */
  user?: string;
}

export interface StorageSharefileConfig {
  /**
   * Upload chunk size.
   * @default "64Mi"
   */
  chunkSize?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot"
   */
  encoding?: string;
  /** Endpoint for API calls. */
  endpoint?: string;
  /**
   * ID of the root folder.
   * @example ""
   */
  rootFolderId?: string;
  /**
   * Cutoff for switching to multipart upload.
   * @default "128Mi"
   */
  uploadCutoff?: string;
}

export interface StorageSiaConfig {
  /** Sia Daemon API Password. */
  apiPassword?: string;
  /**
   * Sia daemon API URL, like http://sia.daemon.host:9980.
   * @default "http://127.0.0.1:9980"
   */
  apiUrl?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Siad User Agent
   * @default "Sia-Agent"
   */
  userAgent?: string;
}

export interface StorageSmbConfig {
  /**
   * Whether the server is configured to be case-insensitive.
   * @default true
   */
  caseInsensitive?: boolean;
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
   * @default true
   */
  hideSpecialShare?: boolean;
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
   * @default 445
   */
  port?: number;
  /** Service principal name. */
  spn?: string;
  /**
   * SMB username.
   * @default "$USER"
   */
  user?: string;
}

export interface StorageStorjExistingConfig {
  /** Access grant. */
  accessGrant?: string;
}

export interface StorageStorjNewConfig {
  /** API key. */
  apiKey?: string;
  /** Encryption passphrase. */
  passphrase?: string;
  /**
   * Satellite address.
   * @default "us1.storj.io"
   * @example "us1.storj.io"
   */
  satelliteAddress?: string;
}

export interface StorageSugarsyncConfig {
  /** Sugarsync Access Key ID. */
  accessKeyId?: string;
  /** Sugarsync App ID. */
  appId?: string;
  /** Sugarsync authorization. */
  authorization?: string;
  /** Sugarsync authorization expiry. */
  authorizationExpiry?: string;
  /** Sugarsync deleted folder id. */
  deletedId?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Permanently delete files if true
   * @default false
   */
  hardDelete?: boolean;
  /** Sugarsync Private Access Key. */
  privateAccessKey?: string;
  /** Sugarsync refresh token. */
  refreshToken?: string;
  /** Sugarsync root id. */
  rootId?: string;
  /** Sugarsync user. */
  user?: string;
}

export interface StorageSwiftConfig {
  /** Application Credential ID (OS_APPLICATION_CREDENTIAL_ID). */
  applicationCredentialId?: string;
  /** Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME). */
  applicationCredentialName?: string;
  /** Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET). */
  applicationCredentialSecret?: string;
  /**
   * Authentication URL for server (OS_AUTH_URL).
   * @example "https://auth.api.rackspacecloud.com/v1.0"
   */
  auth?: string;
  /** Auth Token from alternate authentication - optional (OS_AUTH_TOKEN). */
  authToken?: string;
  /**
   * AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION).
   * @default 0
   */
  authVersion?: number;
  /**
   * Above this size files will be chunked into a _segments container.
   * @default "5Gi"
   */
  chunkSize?: string;
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
   * @example "public"
   */
  endpointType?: string;
  /**
   * Get swift credentials from environment variables in standard OpenStack form.
   * @default false
   * @example false
   */
  envAuth?: boolean;
  /** API key or password (OS_PASSWORD). */
  key?: string;
  /**
   * If true avoid calling abort upload on a failure.
   * @default false
   */
  leavePartsOnError?: boolean;
  /**
   * Don't chunk files during streaming upload.
   * @default false
   */
  noChunk?: boolean;
  /**
   * Disable support for static and dynamic large objects
   * @default false
   */
  noLargeObjects?: boolean;
  /** Region name - optional (OS_REGION_NAME). */
  region?: string;
  /**
   * The storage policy to use when creating a new container.
   * @example ""
   */
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

export interface StorageUptoboxConfig {
  /** Your access token. */
  accessToken?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot"
   */
  encoding?: string;
}

export interface StorageWebdavConfig {
  /** Bearer token instead of user/pass (e.g. a Macaroon). */
  bearerToken?: string;
  /** Command to run to get a bearer token. */
  bearerTokenCommand?: string;
  /** The encoding for the backend. */
  encoding?: string;
  /** Set HTTP headers for all transactions. */
  headers?: string;
  /** Password. */
  pass?: string;
  /** URL of http host to connect to. */
  url?: string;
  /** User name. */
  user?: string;
  /**
   * Name of the WebDAV site/service/software you are using.
   * @example "nextcloud"
   */
  vendor?: string;
}

export interface StorageYandexConfig {
  /** Auth server URL. */
  authUrl?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /**
   * The encoding for the backend.
   * @default "Slash,Del,Ctl,InvalidUtf8,Dot"
   */
  encoding?: string;
  /**
   * Delete files permanently rather than putting them into the trash.
   * @default false
   */
  hardDelete?: boolean;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
}

export interface StorageZohoConfig {
  /** Auth server URL. */
  authUrl?: string;
  /** OAuth Client Id. */
  clientId?: string;
  /** OAuth Client Secret. */
  clientSecret?: string;
  /**
   * The encoding for the backend.
   * @default "Del,Ctl,InvalidUtf8"
   */
  encoding?: string;
  /**
   * Zoho region to connect to.
   * @example "com"
   */
  region?: string;
  /** OAuth Access Token as a JSON blob. */
  token?: string;
  /** Token server url. */
  tokenUrl?: string;
}

export type StorePieceReader = object;

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
  piece = {
    /**
     * @description Get metadata for a piece for how it may be reassembled from the data source
     *
     * @tags Piece
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
  preparation = {
    /**
     * No description
     *
     * @tags Preparation
     * @name PreparationList
     * @summary List all preparations
     * @request GET:/preparation
     */
    preparationList: (params: RequestParams = {}) =>
      this.request<ModelPreparation[], ApiHTTPError>({
        path: `/preparation`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Preparation
     * @name PreparationCreate
     * @summary Create a new preparation
     * @request POST:/preparation
     */
    preparationCreate: (request: DataprepCreateRequest, params: RequestParams = {}) =>
      this.request<ModelPreparation, ApiHTTPError>({
        path: `/preparation`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Preparation
     * @name PreparationDetail
     * @summary Get the status of a preparation
     * @request GET:/preparation/{id}
     */
    preparationDetail: (id: number, params: RequestParams = {}) =>
      this.request<DataprepSourceStatus[], ApiHTTPError>({
        path: `/preparation/${id}`,
        method: "GET",
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Preparation
     * @name OutputCreate
     * @summary Attach an output storage with a preparation
     * @request POST:/preparation/{id}/output/{name}
     */
    outputCreate: (id: number, name: string, params: RequestParams = {}) =>
      this.request<ModelPreparation, ApiHTTPError>({
        path: `/preparation/${id}/output/${name}`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Preparation
     * @name OutputDelete
     * @summary Detach an output storage from a preparation
     * @request DELETE:/preparation/{id}/output/{name}
     */
    outputDelete: (id: number, name: string, params: RequestParams = {}) =>
      this.request<ModelPreparation, ApiHTTPError>({
        path: `/preparation/${id}/output/${name}`,
        method: "DELETE",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Piece
     * @name PieceDetail
     * @summary List all prepared pieces for a preparation
     * @request GET:/preparation/{id}/piece
     */
    pieceDetail: (id: number, params: RequestParams = {}) =>
      this.request<DataprepPieceList[], ApiHTTPError>({
        path: `/preparation/${id}/piece`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Piece
     * @name PieceCreate
     * @summary Add a piece to a preparation
     * @request POST:/preparation/{id}/piece
     */
    pieceCreate: (id: number, request: DataprepAddPieceRequest, params: RequestParams = {}) =>
      this.request<ModelCar, ApiHTTPError>({
        path: `/preparation/${id}/piece`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Preparation
     * @name SourceCreate
     * @summary Attach a source storage with a preparation
     * @request POST:/preparation/{id}/source/{name}
     */
    sourceCreate: (id: number, name: string, params: RequestParams = {}) =>
      this.request<ModelPreparation, ApiHTTPError>({
        path: `/preparation/${id}/source/${name}`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Preparation
     * @name SourceExploreDetail
     * @summary Explore a directory in a prepared source storage
     * @request GET:/preparation/{id}/source/{name}/explore/{path}
     */
    sourceExploreDetail: (id: number, name: string, path: string, params: RequestParams = {}) =>
      this.request<DataprepExploreResult, ApiHTTPError>({
        path: `/preparation/${id}/source/${name}/explore/${path}`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Job
     * @name SourcePauseDaggenCreate
     * @summary Pause an ongoing DAG generation job
     * @request POST:/preparation/{id}/source/{name}/pause-daggen
     */
    sourcePauseDaggenCreate: (id: number, name: string, params: RequestParams = {}) =>
      this.request<ModelJob, ApiHTTPError>({
        path: `/preparation/${id}/source/${name}/pause-daggen`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Job
     * @name SourcePausePackCreate
     * @summary Pause all packing job
     * @request POST:/preparation/{id}/source/{name}/pause-pack
     */
    sourcePausePackCreate: (id: number, name: string, params: RequestParams = {}) =>
      this.request<ModelJob, ApiHTTPError>({
        path: `/preparation/${id}/source/${name}/pause-pack`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Job
     * @name SourcePausePackCreate2
     * @summary Pause a specific packing job
     * @request POST:/preparation/{id}/source/{name}/pause-pack/{job_id}
     * @originalName sourcePausePackCreate
     * @duplicate
     */
    sourcePausePackCreate2: (id: number, name: string, jobId: number, params: RequestParams = {}) =>
      this.request<ModelJob, ApiHTTPError>({
        path: `/preparation/${id}/source/${name}/pause-pack/${jobId}`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Job
     * @name SourcePauseScanCreate
     * @summary Pause an ongoing scanning job
     * @request POST:/preparation/{id}/source/{name}/pause-scan
     */
    sourcePauseScanCreate: (id: number, name: string, params: RequestParams = {}) =>
      this.request<ModelJob, ApiHTTPError>({
        path: `/preparation/${id}/source/${name}/pause-scan`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Job
     * @name SourceStartDaggenCreate
     * @summary Start a new DAG generation job
     * @request POST:/preparation/{id}/source/{name}/start-daggen
     */
    sourceStartDaggenCreate: (id: number, name: string, params: RequestParams = {}) =>
      this.request<ModelJob, ApiHTTPError>({
        path: `/preparation/${id}/source/${name}/start-daggen`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Job
     * @name SourceStartPackCreate
     * @summary Start or restart all packing job
     * @request POST:/preparation/{id}/source/{name}/start-pack
     */
    sourceStartPackCreate: (id: number, name: string, params: RequestParams = {}) =>
      this.request<ModelJob, ApiHTTPError>({
        path: `/preparation/${id}/source/${name}/start-pack`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Job
     * @name SourceStartPackCreate2
     * @summary Start or restart a specific packing job
     * @request POST:/preparation/{id}/source/{name}/start-pack/{job_id}
     * @originalName sourceStartPackCreate
     * @duplicate
     */
    sourceStartPackCreate2: (id: number, name: string, jobId: number, params: RequestParams = {}) =>
      this.request<ModelJob, ApiHTTPError>({
        path: `/preparation/${id}/source/${name}/start-pack/${jobId}`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Job
     * @name SourceStartScanCreate
     * @summary Start a new scanning job
     * @request POST:/preparation/{id}/source/{name}/start-scan
     */
    sourceStartScanCreate: (id: number, name: string, params: RequestParams = {}) =>
      this.request<ModelJob, ApiHTTPError>({
        path: `/preparation/${id}/source/${name}/start-scan`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Wallet Association
     * @name WalletCreate
     * @summary List all wallets of a preparation.
     * @request POST:/preparation/{id}/wallet
     */
    walletCreate: (id: number, params: RequestParams = {}) =>
      this.request<ModelWallet, ApiHTTPError>({
        path: `/preparation/${id}/wallet`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Wallet Association
     * @name WalletCreate2
     * @summary Attach a new wallet with a preparation
     * @request POST:/preparation/{id}/wallet/{wallet}
     * @originalName walletCreate
     * @duplicate
     */
    walletCreate2: (id: number, wallet: string, params: RequestParams = {}) =>
      this.request<ModelPreparation, ApiHTTPError>({
        path: `/preparation/${id}/wallet/${wallet}`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Wallet Association
     * @name WalletDelete
     * @summary Detach a new wallet from a preparation
     * @request DELETE:/preparation/{id}/wallet/{wallet}
     */
    walletDelete: (id: number, wallet: string, params: RequestParams = {}) =>
      this.request<ModelPreparation, ApiHTTPError>({
        path: `/preparation/${id}/wallet/${wallet}`,
        method: "DELETE",
        type: ContentType.Json,
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
  storage = {
    /**
     * No description
     *
     * @tags Storage
     * @name StorageList
     * @summary List all storages
     * @request GET:/storage
     */
    storageList: (params: RequestParams = {}) =>
      this.request<ModelStorage[], ApiHTTPError>({
        path: `/storage`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name PostStorage
     * @summary Create Acd storage
     * @request POST:/storage/acd
     */
    postStorage: (request: StorageCreateAcdStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/acd`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name AzureblobCreate
     * @summary Create Azureblob storage
     * @request POST:/storage/azureblob
     */
    azureblobCreate: (request: StorageCreateAzureblobStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/azureblob`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name PostStorage2
     * @summary Create B2 storage
     * @request POST:/storage/b2
     * @originalName postStorage
     * @duplicate
     */
    postStorage2: (request: StorageCreateB2StorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/b2`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name PostStorage3
     * @summary Create Box storage
     * @request POST:/storage/box
     * @originalName postStorage
     * @duplicate
     */
    postStorage3: (request: StorageCreateBoxStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/box`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name DriveCreate
     * @summary Create Drive storage
     * @request POST:/storage/drive
     */
    driveCreate: (request: StorageCreateDriveStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/drive`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name DropboxCreate
     * @summary Create Dropbox storage
     * @request POST:/storage/dropbox
     */
    dropboxCreate: (request: StorageCreateDropboxStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/dropbox`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name FichierCreate
     * @summary Create Fichier storage
     * @request POST:/storage/fichier
     */
    fichierCreate: (request: StorageCreateFichierStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/fichier`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name FilefabricCreate
     * @summary Create Filefabric storage
     * @request POST:/storage/filefabric
     */
    filefabricCreate: (request: StorageCreateFilefabricStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/filefabric`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name PostStorage4
     * @summary Create Ftp storage
     * @request POST:/storage/ftp
     * @originalName postStorage
     * @duplicate
     */
    postStorage4: (request: StorageCreateFtpStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/ftp`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name PostStorage5
     * @summary Create Gcs storage
     * @request POST:/storage/gcs
     * @originalName postStorage
     * @duplicate
     */
    postStorage5: (request: StorageCreateGcsStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/gcs`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name GphotosCreate
     * @summary Create Gphotos storage
     * @request POST:/storage/gphotos
     */
    gphotosCreate: (request: StorageCreateGphotosStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/gphotos`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name HdfsCreate
     * @summary Create Hdfs storage
     * @request POST:/storage/hdfs
     */
    hdfsCreate: (request: StorageCreateHdfsStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/hdfs`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name HidriveCreate
     * @summary Create Hidrive storage
     * @request POST:/storage/hidrive
     */
    hidriveCreate: (request: StorageCreateHidriveStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/hidrive`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name HttpCreate
     * @summary Create Http storage
     * @request POST:/storage/http
     */
    httpCreate: (request: StorageCreateHttpStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/http`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name InternetarchiveCreate
     * @summary Create Internetarchive storage
     * @request POST:/storage/internetarchive
     */
    internetarchiveCreate: (request: StorageCreateInternetarchiveStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/internetarchive`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name JottacloudCreate
     * @summary Create Jottacloud storage
     * @request POST:/storage/jottacloud
     */
    jottacloudCreate: (request: StorageCreateJottacloudStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/jottacloud`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name KoofrDigistorageCreate
     * @summary Create Koofr storage with digistorage - Digi Storage, https://storage.rcs-rds.ro/
     * @request POST:/storage/koofr/digistorage
     */
    koofrDigistorageCreate: (request: StorageCreateKoofrDigistorageStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/koofr/digistorage`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name KoofrKoofrCreate
     * @summary Create Koofr storage with koofr - Koofr, https://app.koofr.net/
     * @request POST:/storage/koofr/koofr
     */
    koofrKoofrCreate: (request: StorageCreateKoofrKoofrStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/koofr/koofr`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name KoofrOtherCreate
     * @summary Create Koofr storage with other - Any other Koofr API compatible storage service
     * @request POST:/storage/koofr/other
     */
    koofrOtherCreate: (request: StorageCreateKoofrOtherStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/koofr/other`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name LocalCreate
     * @summary Create Local storage
     * @request POST:/storage/local
     */
    localCreate: (request: StorageCreateLocalStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/local`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name MailruCreate
     * @summary Create Mailru storage
     * @request POST:/storage/mailru
     */
    mailruCreate: (request: StorageCreateMailruStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/mailru`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name MegaCreate
     * @summary Create Mega storage
     * @request POST:/storage/mega
     */
    megaCreate: (request: StorageCreateMegaStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/mega`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name NetstorageCreate
     * @summary Create Netstorage storage
     * @request POST:/storage/netstorage
     */
    netstorageCreate: (request: StorageCreateNetstorageStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/netstorage`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name OnedriveCreate
     * @summary Create Onedrive storage
     * @request POST:/storage/onedrive
     */
    onedriveCreate: (request: StorageCreateOnedriveStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/onedrive`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name OosEnvAuthCreate
     * @summary Create Oos storage with env_auth - automatically pickup the credentials from runtime(env), first one to provide auth wins
     * @request POST:/storage/oos/env_auth
     */
    oosEnvAuthCreate: (request: StorageCreateOosEnvAuthStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/oos/env_auth`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name OosInstancePrincipalAuthCreate
     * @summary Create Oos storage with instance_principal_auth - use instance principals to authorize an instance to make API calls.
     * @request POST:/storage/oos/instance_principal_auth
     */
    oosInstancePrincipalAuthCreate: (
      request: StorageCreateOosInstancePrincipalAuthStorageRequest,
      params: RequestParams = {},
    ) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/oos/instance_principal_auth`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name OosNoAuthCreate
     * @summary Create Oos storage with no_auth - no credentials needed, this is typically for reading public buckets
     * @request POST:/storage/oos/no_auth
     */
    oosNoAuthCreate: (request: StorageCreateOosNoAuthStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/oos/no_auth`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name OosResourcePrincipalAuthCreate
     * @summary Create Oos storage with resource_principal_auth - use resource principals to make API calls
     * @request POST:/storage/oos/resource_principal_auth
     */
    oosResourcePrincipalAuthCreate: (
      request: StorageCreateOosResourcePrincipalAuthStorageRequest,
      params: RequestParams = {},
    ) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/oos/resource_principal_auth`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name OosUserPrincipalAuthCreate
     * @summary Create Oos storage with user_principal_auth - use an OCI user and an API key for authentication.
     * @request POST:/storage/oos/user_principal_auth
     */
    oosUserPrincipalAuthCreate: (
      request: StorageCreateOosUserPrincipalAuthStorageRequest,
      params: RequestParams = {},
    ) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/oos/user_principal_auth`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name OpendriveCreate
     * @summary Create Opendrive storage
     * @request POST:/storage/opendrive
     */
    opendriveCreate: (request: StorageCreateOpendriveStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/opendrive`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name PcloudCreate
     * @summary Create Pcloud storage
     * @request POST:/storage/pcloud
     */
    pcloudCreate: (request: StorageCreatePcloudStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/pcloud`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name PremiumizemeCreate
     * @summary Create Premiumizeme storage
     * @request POST:/storage/premiumizeme
     */
    premiumizemeCreate: (request: StorageCreatePremiumizemeStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/premiumizeme`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name PutioCreate
     * @summary Create Putio storage
     * @request POST:/storage/putio
     */
    putioCreate: (request: StorageCreatePutioStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/putio`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name QingstorCreate
     * @summary Create Qingstor storage
     * @request POST:/storage/qingstor
     */
    qingstorCreate: (request: StorageCreateQingstorStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/qingstor`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3AlibabaCreate
     * @summary Create S3 storage with Alibaba - Alibaba Cloud Object Storage System (OSS) formerly Aliyun
     * @request POST:/storage/s3/alibaba
     */
    s3AlibabaCreate: (request: StorageCreateS3AlibabaStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/alibaba`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3ArvancloudCreate
     * @summary Create S3 storage with ArvanCloud - Arvan Cloud Object Storage (AOS)
     * @request POST:/storage/s3/arvancloud
     */
    s3ArvancloudCreate: (request: StorageCreateS3ArvanCloudStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/arvancloud`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3AwsCreate
     * @summary Create S3 storage with AWS - Amazon Web Services (AWS) S3
     * @request POST:/storage/s3/aws
     */
    s3AwsCreate: (request: StorageCreateS3AWSStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/aws`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3CephCreate
     * @summary Create S3 storage with Ceph - Ceph Object Storage
     * @request POST:/storage/s3/ceph
     */
    s3CephCreate: (request: StorageCreateS3CephStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/ceph`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3ChinamobileCreate
     * @summary Create S3 storage with ChinaMobile - China Mobile Ecloud Elastic Object Storage (EOS)
     * @request POST:/storage/s3/chinamobile
     */
    s3ChinamobileCreate: (request: StorageCreateS3ChinaMobileStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/chinamobile`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3CloudflareCreate
     * @summary Create S3 storage with Cloudflare - Cloudflare R2 Storage
     * @request POST:/storage/s3/cloudflare
     */
    s3CloudflareCreate: (request: StorageCreateS3CloudflareStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/cloudflare`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3DigitaloceanCreate
     * @summary Create S3 storage with DigitalOcean - DigitalOcean Spaces
     * @request POST:/storage/s3/digitalocean
     */
    s3DigitaloceanCreate: (request: StorageCreateS3DigitalOceanStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/digitalocean`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3DreamhostCreate
     * @summary Create S3 storage with Dreamhost - Dreamhost DreamObjects
     * @request POST:/storage/s3/dreamhost
     */
    s3DreamhostCreate: (request: StorageCreateS3DreamhostStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/dreamhost`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3HuaweiobsCreate
     * @summary Create S3 storage with HuaweiOBS - Huawei Object Storage Service
     * @request POST:/storage/s3/huaweiobs
     */
    s3HuaweiobsCreate: (request: StorageCreateS3HuaweiOBSStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/huaweiobs`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3IbmcosCreate
     * @summary Create S3 storage with IBMCOS - IBM COS S3
     * @request POST:/storage/s3/ibmcos
     */
    s3IbmcosCreate: (request: StorageCreateS3IBMCOSStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/ibmcos`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3IdriveCreate
     * @summary Create S3 storage with IDrive - IDrive e2
     * @request POST:/storage/s3/idrive
     */
    s3IdriveCreate: (request: StorageCreateS3IDriveStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/idrive`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3IonosCreate
     * @summary Create S3 storage with IONOS - IONOS Cloud
     * @request POST:/storage/s3/ionos
     */
    s3IonosCreate: (request: StorageCreateS3IONOSStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/ionos`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3LiaraCreate
     * @summary Create S3 storage with Liara - Liara Object Storage
     * @request POST:/storage/s3/liara
     */
    s3LiaraCreate: (request: StorageCreateS3LiaraStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/liara`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3LyvecloudCreate
     * @summary Create S3 storage with LyveCloud - Seagate Lyve Cloud
     * @request POST:/storage/s3/lyvecloud
     */
    s3LyvecloudCreate: (request: StorageCreateS3LyveCloudStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/lyvecloud`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3MinioCreate
     * @summary Create S3 storage with Minio - Minio Object Storage
     * @request POST:/storage/s3/minio
     */
    s3MinioCreate: (request: StorageCreateS3MinioStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/minio`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3NeteaseCreate
     * @summary Create S3 storage with Netease - Netease Object Storage (NOS)
     * @request POST:/storage/s3/netease
     */
    s3NeteaseCreate: (request: StorageCreateS3NeteaseStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/netease`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3OtherCreate
     * @summary Create S3 storage with Other - Any other S3 compatible provider
     * @request POST:/storage/s3/other
     */
    s3OtherCreate: (request: StorageCreateS3OtherStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/other`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3QiniuCreate
     * @summary Create S3 storage with Qiniu - Qiniu Object Storage (Kodo)
     * @request POST:/storage/s3/qiniu
     */
    s3QiniuCreate: (request: StorageCreateS3QiniuStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/qiniu`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3RackcorpCreate
     * @summary Create S3 storage with RackCorp - RackCorp Object Storage
     * @request POST:/storage/s3/rackcorp
     */
    s3RackcorpCreate: (request: StorageCreateS3RackCorpStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/rackcorp`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3ScalewayCreate
     * @summary Create S3 storage with Scaleway - Scaleway Object Storage
     * @request POST:/storage/s3/scaleway
     */
    s3ScalewayCreate: (request: StorageCreateS3ScalewayStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/scaleway`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3SeaweedfsCreate
     * @summary Create S3 storage with SeaweedFS - SeaweedFS S3
     * @request POST:/storage/s3/seaweedfs
     */
    s3SeaweedfsCreate: (request: StorageCreateS3SeaweedFSStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/seaweedfs`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3StackpathCreate
     * @summary Create S3 storage with StackPath - StackPath Object Storage
     * @request POST:/storage/s3/stackpath
     */
    s3StackpathCreate: (request: StorageCreateS3StackPathStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/stackpath`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3StorjCreate
     * @summary Create S3 storage with Storj - Storj (S3 Compatible Gateway)
     * @request POST:/storage/s3/storj
     */
    s3StorjCreate: (request: StorageCreateS3StorjStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/storj`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3TencentcosCreate
     * @summary Create S3 storage with TencentCOS - Tencent Cloud Object Storage (COS)
     * @request POST:/storage/s3/tencentcos
     */
    s3TencentcosCreate: (request: StorageCreateS3TencentCOSStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/tencentcos`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name S3WasabiCreate
     * @summary Create S3 storage with Wasabi - Wasabi Object Storage
     * @request POST:/storage/s3/wasabi
     */
    s3WasabiCreate: (request: StorageCreateS3WasabiStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/s3/wasabi`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name SeafileCreate
     * @summary Create Seafile storage
     * @request POST:/storage/seafile
     */
    seafileCreate: (request: StorageCreateSeafileStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/seafile`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name SftpCreate
     * @summary Create Sftp storage
     * @request POST:/storage/sftp
     */
    sftpCreate: (request: StorageCreateSftpStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/sftp`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name SharefileCreate
     * @summary Create Sharefile storage
     * @request POST:/storage/sharefile
     */
    sharefileCreate: (request: StorageCreateSharefileStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/sharefile`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name PostStorage6
     * @summary Create Sia storage
     * @request POST:/storage/sia
     * @originalName postStorage
     * @duplicate
     */
    postStorage6: (request: StorageCreateSiaStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/sia`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name PostStorage7
     * @summary Create Smb storage
     * @request POST:/storage/smb
     * @originalName postStorage
     * @duplicate
     */
    postStorage7: (request: StorageCreateSmbStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/smb`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name StorjExistingCreate
     * @summary Create Storj storage with existing - Use an existing access grant.
     * @request POST:/storage/storj/existing
     */
    storjExistingCreate: (request: StorageCreateStorjExistingStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/storj/existing`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name StorjNewCreate
     * @summary Create Storj storage with new - Create a new access grant from satellite address, API key, and passphrase.
     * @request POST:/storage/storj/new
     */
    storjNewCreate: (request: StorageCreateStorjNewStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/storj/new`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name SugarsyncCreate
     * @summary Create Sugarsync storage
     * @request POST:/storage/sugarsync
     */
    sugarsyncCreate: (request: StorageCreateSugarsyncStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/sugarsync`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name SwiftCreate
     * @summary Create Swift storage
     * @request POST:/storage/swift
     */
    swiftCreate: (request: StorageCreateSwiftStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/swift`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name UptoboxCreate
     * @summary Create Uptobox storage
     * @request POST:/storage/uptobox
     */
    uptoboxCreate: (request: StorageCreateUptoboxStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/uptobox`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name WebdavCreate
     * @summary Create Webdav storage
     * @request POST:/storage/webdav
     */
    webdavCreate: (request: StorageCreateWebdavStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/webdav`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name YandexCreate
     * @summary Create Yandex storage
     * @request POST:/storage/yandex
     */
    yandexCreate: (request: StorageCreateYandexStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/yandex`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name ZohoCreate
     * @summary Create Zoho storage
     * @request POST:/storage/zoho
     */
    zohoCreate: (request: StorageCreateZohoStorageRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/zoho`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name StorageDelete
     * @summary Remove a storage
     * @request DELETE:/storage/{name}
     */
    storageDelete: (name: string, params: RequestParams = {}) =>
      this.request<void, ApiHTTPError>({
        path: `/storage/${name}`,
        method: "DELETE",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name StoragePartialUpdate
     * @summary Update a storage connection
     * @request PATCH:/storage/{name}
     */
    storagePartialUpdate: (name: string, config: ModelCID, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/${name}`,
        method: "PATCH",
        body: config,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name ExploreDetail
     * @summary Explore directory entries in a storage system
     * @request GET:/storage/{name}/explore/{path}
     */
    exploreDetail: (name: string, path: string, params: RequestParams = {}) =>
      this.request<StorageDirEntry[], ApiHTTPError>({
        path: `/storage/${name}/explore/${path}`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * No description
     *
     * @tags Storage
     * @name StorageCreate
     * @summary Create a new storage
     * @request POST:/storage/{storageType}
     */
    storageCreate: (storageType: string, body: StorageCreateRequest, params: RequestParams = {}) =>
      this.request<ModelStorage, ApiHTTPError>({
        path: `/storage/${storageType}`,
        method: "POST",
        body: body,
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
