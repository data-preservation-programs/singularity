# Table of contents

## Overview <a href="#overview" id="overview"></a>

* [What is Singularity](README.md)
* [V1 or V2](overview/v1-or-v2.md)

## Installation <a href="#installation" id="installation"></a>

* [Download binaries](installation/download-binaries.md)
* [Install via docker](installation/install-from-docker.md)
* [Built from source](installation/install-from-source.md)
* [Deploy to production](installation/deploy-to-production.md)
* [Version upgrade](installation/upgrade.md)

## Data Preparation <a href="#data-preparation" id="data-preparation"></a>

* [Get Started](data-preparation/get-started.md)
* [Performance Tuning](data-preparation/performance-tuning.md)

## Content Distribution <a href="#content-distribution" id="content-distribution"></a>

* [Distribute CAR files](content-distribution/distribute-car-files.md)

## Deal Making <a href="#deal-making" id="deal-making"></a>

* [Create a deal schedule](deal-making/create-a-deal-schedule.md)
* [Deal Templates](deal-templates.md)

## Topics <a href="#topics" id="topics"></a>

* [Inline Preparation](topics/inline-preparation.md)
* [Benchmark](topics/benchmark.md)

## 💻 CLI Reference <a href="#cli-reference" id="cli-reference"></a>
<!-- cli begin -->

* [Menu](cli-reference/README.md)
* [Onboard](cli-reference/onboard.md)
* [Ez Prep](cli-reference/ez-prep.md)
* [Version](cli-reference/version.md)
* [Admin](cli-reference/admin/README.md)
  * [Init](cli-reference/admin/init.md)
  * [Reset](cli-reference/admin/reset.md)
  * [Migrate](cli-reference/admin/migrate/README.md)
    * [Up](cli-reference/admin/migrate/up.md)
    * [Down](cli-reference/admin/migrate/down.md)
    * [To](cli-reference/admin/migrate/to.md)
    * [Which](cli-reference/admin/migrate/which.md)
  * [Migrate Dataset](cli-reference/admin/migrate-dataset.md)
  * [Migrate Schedule](cli-reference/admin/migrate-schedule.md)
* [Download](cli-reference/download.md)
* [Extract Car](cli-reference/extract-car.md)
* [Deal](cli-reference/deal/README.md)
  * [Schedule](cli-reference/deal/schedule/README.md)
    * [Create](cli-reference/deal/schedule/create.md)
    * [List](cli-reference/deal/schedule/list.md)
    * [Update](cli-reference/deal/schedule/update.md)
    * [Pause](cli-reference/deal/schedule/pause.md)
    * [Resume](cli-reference/deal/schedule/resume.md)
    * [Remove](cli-reference/deal/schedule/remove.md)
  * [Send Manual](cli-reference/deal/send-manual.md)
  * [List](cli-reference/deal/list.md)
* [Deal Template](cli-reference/deal-template/README.md)
  * [Create](cli-reference/deal-template/create.md)
  * [List](cli-reference/deal-template/list.md)
  * [Get](cli-reference/deal-template/get.md)
  * [Delete](cli-reference/deal-template/delete.md)
* [Run](cli-reference/run/README.md)
  * [Api](cli-reference/run/api.md)
  * [Dataset Worker](cli-reference/run/dataset-worker.md)
  * [Content Provider](cli-reference/run/content-provider.md)
  * [Deal Tracker](cli-reference/run/deal-tracker.md)
  * [Deal Pusher](cli-reference/run/deal-pusher.md)
  * [Download Server](cli-reference/run/download-server.md)
  * [Unified](cli-reference/run/unified.md)
* [Wallet](cli-reference/wallet/README.md)
  * [Create](cli-reference/wallet/create.md)
  * [Import](cli-reference/wallet/import.md)
  * [Init](cli-reference/wallet/init.md)
  * [List](cli-reference/wallet/list.md)
  * [Remove](cli-reference/wallet/remove.md)
  * [Update](cli-reference/wallet/update.md)
* [Storage](cli-reference/storage/README.md)
  * [Create](cli-reference/storage/create/README.md)
    * [Acd](cli-reference/storage/create/acd.md)
    * [Azureblob](cli-reference/storage/create/azureblob.md)
    * [B2](cli-reference/storage/create/b2.md)
    * [Box](cli-reference/storage/create/box.md)
    * [Drive](cli-reference/storage/create/drive.md)
    * [Dropbox](cli-reference/storage/create/dropbox.md)
    * [Fichier](cli-reference/storage/create/fichier.md)
    * [Filefabric](cli-reference/storage/create/filefabric.md)
    * [Ftp](cli-reference/storage/create/ftp.md)
    * [Google Cloud Storage](cli-reference/storage/create/gcs.md)
    * [Gphotos](cli-reference/storage/create/gphotos.md)
    * [Hdfs](cli-reference/storage/create/hdfs.md)
    * [Hidrive](cli-reference/storage/create/hidrive.md)
    * [Http](cli-reference/storage/create/http.md)
    * [Internetarchive](cli-reference/storage/create/internetarchive.md)
    * [Jottacloud](cli-reference/storage/create/jottacloud.md)
    * [Koofr / Digi Storage](cli-reference/storage/create/koofr/README.md)
      * [Digistorage](cli-reference/storage/create/koofr/digistorage.md)
      * [Koofr / Digi Storage](cli-reference/storage/create/koofr/koofr.md)
      * [Other](cli-reference/storage/create/koofr/other.md)
    * [Local](cli-reference/storage/create/local.md)
    * [Mailru](cli-reference/storage/create/mailru.md)
    * [Mega](cli-reference/storage/create/mega.md)
    * [Netstorage](cli-reference/storage/create/netstorage.md)
    * [Onedrive](cli-reference/storage/create/onedrive.md)
    * [Oos](cli-reference/storage/create/oos/README.md)
      * [Env_auth](cli-reference/storage/create/oos/env_auth.md)
      * [Instance_principal_auth](cli-reference/storage/create/oos/instance_principal_auth.md)
      * [No_auth](cli-reference/storage/create/oos/no_auth.md)
      * [Resource_principal_auth](cli-reference/storage/create/oos/resource_principal_auth.md)
      * [User_principal_auth](cli-reference/storage/create/oos/user_principal_auth.md)
    * [Opendrive](cli-reference/storage/create/opendrive.md)
    * [Pcloud](cli-reference/storage/create/pcloud.md)
    * [Premiumizeme](cli-reference/storage/create/premiumizeme.md)
    * [Putio](cli-reference/storage/create/putio.md)
    * [Qingstor](cli-reference/storage/create/qingstor.md)
    * [AWS S3 and compliant](cli-reference/storage/create/s3/README.md)
      * [Aws](cli-reference/storage/create/s3/aws.md)
      * [Alibaba](cli-reference/storage/create/s3/alibaba.md)
      * [Arvancloud](cli-reference/storage/create/s3/arvancloud.md)
      * [Ceph](cli-reference/storage/create/s3/ceph.md)
      * [Chinamobile](cli-reference/storage/create/s3/chinamobile.md)
      * [Cloudflare](cli-reference/storage/create/s3/cloudflare.md)
      * [Digitalocean](cli-reference/storage/create/s3/digitalocean.md)
      * [Dreamhost](cli-reference/storage/create/s3/dreamhost.md)
      * [Huaweiobs](cli-reference/storage/create/s3/huaweiobs.md)
      * [Ibmcos](cli-reference/storage/create/s3/ibmcos.md)
      * [Idrive](cli-reference/storage/create/s3/idrive.md)
      * [Ionos](cli-reference/storage/create/s3/ionos.md)
      * [Liara](cli-reference/storage/create/s3/liara.md)
      * [Lyvecloud](cli-reference/storage/create/s3/lyvecloud.md)
      * [Minio](cli-reference/storage/create/s3/minio.md)
      * [Netease](cli-reference/storage/create/s3/netease.md)
      * [Other](cli-reference/storage/create/s3/other.md)
      * [Qiniu](cli-reference/storage/create/s3/qiniu.md)
      * [Rackcorp](cli-reference/storage/create/s3/rackcorp.md)
      * [Scaleway](cli-reference/storage/create/s3/scaleway.md)
      * [Seaweedfs](cli-reference/storage/create/s3/seaweedfs.md)
      * [Stackpath](cli-reference/storage/create/s3/stackpath.md)
      * [Storj](cli-reference/storage/create/s3/storj.md)
      * [Tencentcos](cli-reference/storage/create/s3/tencentcos.md)
      * [Wasabi](cli-reference/storage/create/s3/wasabi.md)
    * [Seafile](cli-reference/storage/create/seafile.md)
    * [Sftp](cli-reference/storage/create/sftp.md)
    * [Sharefile](cli-reference/storage/create/sharefile.md)
    * [Sia](cli-reference/storage/create/sia.md)
    * [Smb](cli-reference/storage/create/smb.md)
    * [Storj](cli-reference/storage/create/storj/README.md)
      * [Existing](cli-reference/storage/create/storj/existing.md)
      * [New](cli-reference/storage/create/storj/new.md)
    * [Sugarsync](cli-reference/storage/create/sugarsync.md)
    * [Swift](cli-reference/storage/create/swift.md)
    * [Union](cli-reference/storage/create/union.md)
    * [Uptobox](cli-reference/storage/create/uptobox.md)
    * [Webdav](cli-reference/storage/create/webdav.md)
    * [Yandex](cli-reference/storage/create/yandex.md)
    * [Zoho](cli-reference/storage/create/zoho.md)
  * [Explore](cli-reference/storage/explore.md)
  * [List](cli-reference/storage/list.md)
  * [Remove](cli-reference/storage/remove.md)
  * [Update](cli-reference/storage/update/README.md)
    * [Acd](cli-reference/storage/update/acd.md)
    * [Azureblob](cli-reference/storage/update/azureblob.md)
    * [B2](cli-reference/storage/update/b2.md)
    * [Box](cli-reference/storage/update/box.md)
    * [Drive](cli-reference/storage/update/drive.md)
    * [Dropbox](cli-reference/storage/update/dropbox.md)
    * [Fichier](cli-reference/storage/update/fichier.md)
    * [Filefabric](cli-reference/storage/update/filefabric.md)
    * [Ftp](cli-reference/storage/update/ftp.md)
    * [Google Cloud Storage](cli-reference/storage/update/gcs.md)
    * [Gphotos](cli-reference/storage/update/gphotos.md)
    * [Hdfs](cli-reference/storage/update/hdfs.md)
    * [Hidrive](cli-reference/storage/update/hidrive.md)
    * [Http](cli-reference/storage/update/http.md)
    * [Internetarchive](cli-reference/storage/update/internetarchive.md)
    * [Jottacloud](cli-reference/storage/update/jottacloud.md)
    * [Koofr / Digi Storage](cli-reference/storage/update/koofr/README.md)
      * [Digistorage](cli-reference/storage/update/koofr/digistorage.md)
      * [Koofr / Digi Storage](cli-reference/storage/update/koofr/koofr.md)
      * [Other](cli-reference/storage/update/koofr/other.md)
    * [Local](cli-reference/storage/update/local.md)
    * [Mailru](cli-reference/storage/update/mailru.md)
    * [Mega](cli-reference/storage/update/mega.md)
    * [Netstorage](cli-reference/storage/update/netstorage.md)
    * [Onedrive](cli-reference/storage/update/onedrive.md)
    * [Oos](cli-reference/storage/update/oos/README.md)
      * [Env_auth](cli-reference/storage/update/oos/env_auth.md)
      * [Instance_principal_auth](cli-reference/storage/update/oos/instance_principal_auth.md)
      * [No_auth](cli-reference/storage/update/oos/no_auth.md)
      * [Resource_principal_auth](cli-reference/storage/update/oos/resource_principal_auth.md)
      * [User_principal_auth](cli-reference/storage/update/oos/user_principal_auth.md)
    * [Opendrive](cli-reference/storage/update/opendrive.md)
    * [Pcloud](cli-reference/storage/update/pcloud.md)
    * [Premiumizeme](cli-reference/storage/update/premiumizeme.md)
    * [Putio](cli-reference/storage/update/putio.md)
    * [Qingstor](cli-reference/storage/update/qingstor.md)
    * [AWS S3 and compliant](cli-reference/storage/update/s3/README.md)
      * [Aws](cli-reference/storage/update/s3/aws.md)
      * [Alibaba](cli-reference/storage/update/s3/alibaba.md)
      * [Arvancloud](cli-reference/storage/update/s3/arvancloud.md)
      * [Ceph](cli-reference/storage/update/s3/ceph.md)
      * [Chinamobile](cli-reference/storage/update/s3/chinamobile.md)
      * [Cloudflare](cli-reference/storage/update/s3/cloudflare.md)
      * [Digitalocean](cli-reference/storage/update/s3/digitalocean.md)
      * [Dreamhost](cli-reference/storage/update/s3/dreamhost.md)
      * [Huaweiobs](cli-reference/storage/update/s3/huaweiobs.md)
      * [Ibmcos](cli-reference/storage/update/s3/ibmcos.md)
      * [Idrive](cli-reference/storage/update/s3/idrive.md)
      * [Ionos](cli-reference/storage/update/s3/ionos.md)
      * [Liara](cli-reference/storage/update/s3/liara.md)
      * [Lyvecloud](cli-reference/storage/update/s3/lyvecloud.md)
      * [Minio](cli-reference/storage/update/s3/minio.md)
      * [Netease](cli-reference/storage/update/s3/netease.md)
      * [Other](cli-reference/storage/update/s3/other.md)
      * [Qiniu](cli-reference/storage/update/s3/qiniu.md)
      * [Rackcorp](cli-reference/storage/update/s3/rackcorp.md)
      * [Scaleway](cli-reference/storage/update/s3/scaleway.md)
      * [Seaweedfs](cli-reference/storage/update/s3/seaweedfs.md)
      * [Stackpath](cli-reference/storage/update/s3/stackpath.md)
      * [Storj](cli-reference/storage/update/s3/storj.md)
      * [Tencentcos](cli-reference/storage/update/s3/tencentcos.md)
      * [Wasabi](cli-reference/storage/update/s3/wasabi.md)
    * [Seafile](cli-reference/storage/update/seafile.md)
    * [Sftp](cli-reference/storage/update/sftp.md)
    * [Sharefile](cli-reference/storage/update/sharefile.md)
    * [Sia](cli-reference/storage/update/sia.md)
    * [Smb](cli-reference/storage/update/smb.md)
    * [Storj](cli-reference/storage/update/storj/README.md)
      * [Existing](cli-reference/storage/update/storj/existing.md)
      * [New](cli-reference/storage/update/storj/new.md)
    * [Sugarsync](cli-reference/storage/update/sugarsync.md)
    * [Swift](cli-reference/storage/update/swift.md)
    * [Union](cli-reference/storage/update/union.md)
    * [Uptobox](cli-reference/storage/update/uptobox.md)
    * [Webdav](cli-reference/storage/update/webdav.md)
    * [Yandex](cli-reference/storage/update/yandex.md)
    * [Zoho](cli-reference/storage/update/zoho.md)
  * [Rename](cli-reference/storage/rename.md)
* [Prep](cli-reference/prep/README.md)
  * [Create](cli-reference/prep/create.md)
  * [List](cli-reference/prep/list.md)
  * [Status](cli-reference/prep/status.md)
  * [Rename](cli-reference/prep/rename.md)
  * [Attach Source](cli-reference/prep/attach-source.md)
  * [Attach Output](cli-reference/prep/attach-output.md)
  * [Detach Output](cli-reference/prep/detach-output.md)
  * [Start Scan](cli-reference/prep/start-scan.md)
  * [Pause Scan](cli-reference/prep/pause-scan.md)
  * [Start Pack](cli-reference/prep/start-pack.md)
  * [Pause Pack](cli-reference/prep/pause-pack.md)
  * [Start Daggen](cli-reference/prep/start-daggen.md)
  * [Pause Daggen](cli-reference/prep/pause-daggen.md)
  * [List Pieces](cli-reference/prep/list-pieces.md)
  * [Add Piece](cli-reference/prep/add-piece.md)
  * [Explore](cli-reference/prep/explore.md)
  * [Attach Wallet](cli-reference/prep/attach-wallet.md)
  * [List Wallets](cli-reference/prep/list-wallets.md)
  * [Detach Wallet](cli-reference/prep/detach-wallet.md)
  * [Remove](cli-reference/prep/remove.md)

<!-- cli end -->

## 🌐 Web API Reference <a href="#web-api-reference" id="web-api-reference"></a>
<!-- webapi begin -->

* [Admin](web-api-reference/admin.md)
* [Deal Schedule](web-api-reference/deal-schedule.md)
* [Deal](web-api-reference/deal.md)
* [File](web-api-reference/file.md)
* [Job](web-api-reference/job.md)
* [Piece](web-api-reference/piece.md)
* [Preparation](web-api-reference/preparation.md)
* [Storage](web-api-reference/storage.md)
* [Wallet Association](web-api-reference/wallet-association.md)
* [Wallet](web-api-reference/wallet.md)
* [Specification](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)

<!-- webapi end -->

## ❓ FAQ <a href="#faq" id="faq"></a>

* [Database is locked](faq/database-is-locked.md)
