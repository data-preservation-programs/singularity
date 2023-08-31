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
  /** preparation ID filter */
  preparations?: number[];
  /** provider filter */
  providers?: string[];
  /** schedule id filter */
  schedules?: number[];
  /** source filter */
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

export enum FsDuration {
  ModTimeNotSupported = -9223372036854776000,
}

export interface FsTristate {
  valid?: boolean;
  value?: boolean;
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
  /** Preparation ID */
  preparationId?: number;
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

export type StorageCreateAcdStorageRequest = object;

export type StorageCreateAzureblobStorageRequest = object;

export type StorageCreateB2StorageRequest = object;

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

export type StorageCreateDriveStorageRequest = object;

export type StorageCreateDropboxStorageRequest = object;

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

export type StorageCreateFtpStorageRequest = object;

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

export type StorageCreateInternetarchiveStorageRequest = object;

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

export type StorageCreateOnedriveStorageRequest = object;

export type StorageCreateOosEnvAuthStorageRequest = object;

export type StorageCreateOosInstancePrincipalAuthStorageRequest = object;

export type StorageCreateOosNoAuthStorageRequest = object;

export type StorageCreateOosResourcePrincipalAuthStorageRequest = object;

export type StorageCreateOosUserPrincipalAuthStorageRequest = object;

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

export type StorageCreateS3AWSStorageRequest = object;

export type StorageCreateS3AlibabaStorageRequest = object;

export type StorageCreateS3ArvanCloudStorageRequest = object;

export type StorageCreateS3CephStorageRequest = object;

export type StorageCreateS3ChinaMobileStorageRequest = object;

export type StorageCreateS3CloudflareStorageRequest = object;

export type StorageCreateS3DigitalOceanStorageRequest = object;

export type StorageCreateS3DreamhostStorageRequest = object;

export type StorageCreateS3HuaweiOBSStorageRequest = object;

export type StorageCreateS3IBMCOSStorageRequest = object;

export type StorageCreateS3IDriveStorageRequest = object;

export type StorageCreateS3IONOSStorageRequest = object;

export type StorageCreateS3LiaraStorageRequest = object;

export type StorageCreateS3LyveCloudStorageRequest = object;

export type StorageCreateS3MinioStorageRequest = object;

export type StorageCreateS3NeteaseStorageRequest = object;

export type StorageCreateS3OtherStorageRequest = object;

export type StorageCreateS3QiniuStorageRequest = object;

export type StorageCreateS3RackCorpStorageRequest = object;

export type StorageCreateS3ScalewayStorageRequest = object;

export type StorageCreateS3SeaweedFSStorageRequest = object;

export type StorageCreateS3StackPathStorageRequest = object;

export type StorageCreateS3StorjStorageRequest = object;

export type StorageCreateS3TencentCOSStorageRequest = object;

export type StorageCreateS3WasabiStorageRequest = object;

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

export type StorageCreateSftpStorageRequest = object;

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

export type StorageCreateSmbStorageRequest = object;

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
  headers?: string[];
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
  headers?: string[];
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
