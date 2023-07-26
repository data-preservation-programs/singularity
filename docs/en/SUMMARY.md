# Table of contents

## Overview

* [What is Singularity](README.md)
* [V1 or V2](overview/v1-or-v2.md)
* [Current Status](overview/current-status.md)
* [Related Projects](overview/related-projects.md)

## Installation

* [Install from source](installation/install-from-source.md)
* [Install from docker](installation/install-from-docker.md)
* [Deploy to production](installation/deploy-to-production.md)

## Data Preparation

* [Create a dataset](data-preparation/create-a-dataset.md)
* [Add a data source](data-preparation/add-a-data-source.md)
* [Start dataset worker](data-preparation/start-dataset-worker.md)
* [Create DAG for the data source](data-preparation/create-dag-for-the-data-source.md)

## Content Distribution

* [Distribute CAR files](content-distribution/distribute-car-files.md)
* [File Retrieval (Staging)](content-distribution/file-retrieval-staging.md)

## Deal Making

* [Deal Making Prerequisite](deal-making/deal-making-prerequisite.md)
* [Create a deal schedule](deal-making/create-a-deal-schedule.md)
* [SP Self Service](deal-making/sp-self-service.md)

## Retrieval

* [Overview](retrieval/overview.md)

## Topics

* [Encryption](topics/encryption.md)
* [Inline Preparation](topics/inline-preparation.md)
* [Datasource rescan](topics/datasource-rescan.md)
* [Push and Upload](topics/push-and-upload.md)
* [Benchmark](topics/benchmark.md)

## 💻 CLI Reference
<!-- cli begin -->

* [Ez Prep](cli-reference/ez-prep.md)
* [Admin](cli-reference/admin/README.md)
  * [Init](cli-reference/admin/init.md)
  * [Reset](cli-reference/admin/reset.md)
  * [Migrate](cli-reference/admin/migrate.md)
* [Download](cli-reference/download.md)
* [Deal](cli-reference/deal/README.md)
  * [Schedule](cli-reference/deal/schedule/README.md)
    * [Create](cli-reference/deal/schedule/create.md)
    * [List](cli-reference/deal/schedule/list.md)
    * [Pause](cli-reference/deal/schedule/pause.md)
    * [Resume](cli-reference/deal/schedule/resume.md)
  * [Spade Policy](cli-reference/deal/spade-policy/README.md)
    * [Create](cli-reference/deal/spade-policy/create.md)
    * [List](cli-reference/deal/spade-policy/list.md)
    * [Remove](cli-reference/deal/spade-policy/remove.md)
  * [Send Manual](cli-reference/deal/send-manual.md)
  * [List](cli-reference/deal/list.md)
* [Run](cli-reference/run/README.md)
  * [Api](cli-reference/run/api.md)
  * [Dataset Worker](cli-reference/run/dataset-worker.md)
  * [Content Provider](cli-reference/run/content-provider.md)
  * [Deal Tracker](cli-reference/run/deal-tracker.md)
  * [Dealmaker](cli-reference/run/dealmaker.md)
  * [Spade Api](cli-reference/run/spade-api.md)
* [Dataset](cli-reference/dataset/README.md)
  * [Create](cli-reference/dataset/create.md)
  * [List](cli-reference/dataset/list.md)
  * [Update](cli-reference/dataset/update.md)
  * [Remove](cli-reference/dataset/remove.md)
  * [Add Wallet](cli-reference/dataset/add-wallet.md)
  * [List Wallet](cli-reference/dataset/list-wallet.md)
  * [Remove Wallet](cli-reference/dataset/remove-wallet.md)
  * [Add Piece](cli-reference/dataset/add-piece.md)
  * [List Pieces](cli-reference/dataset/list-pieces.md)
* [Datasource](cli-reference/datasource/README.md)
  * [Add](cli-reference/datasource/add/README.md)
    * [Acd](cli-reference/datasource/add/acd.md)
    * [Azureblob](cli-reference/datasource/add/azureblob.md)
    * [B2](cli-reference/datasource/add/b2.md)
    * [Box](cli-reference/datasource/add/box.md)
    * [Drive](cli-reference/datasource/add/drive.md)
    * [Dropbox](cli-reference/datasource/add/dropbox.md)
    * [Fichier](cli-reference/datasource/add/fichier.md)
    * [Filefabric](cli-reference/datasource/add/filefabric.md)
    * [Ftp](cli-reference/datasource/add/ftp.md)
    * [Google Cloud Storage](cli-reference/datasource/add/gcs.md)
    * [Gphotos](cli-reference/datasource/add/gphotos.md)
    * [Hdfs](cli-reference/datasource/add/hdfs.md)
    * [Hidrive](cli-reference/datasource/add/hidrive.md)
    * [Http](cli-reference/datasource/add/http.md)
    * [Internetarchive](cli-reference/datasource/add/internetarchive.md)
    * [Jottacloud](cli-reference/datasource/add/jottacloud.md)
    * [Koofr / Digi Storage](cli-reference/datasource/add/koofr.md)
    * [Local](cli-reference/datasource/add/local.md)
    * [Mailru](cli-reference/datasource/add/mailru.md)
    * [Mega](cli-reference/datasource/add/mega.md)
    * [Netstorage](cli-reference/datasource/add/netstorage.md)
    * [Onedrive](cli-reference/datasource/add/onedrive.md)
    * [Opendrive](cli-reference/datasource/add/opendrive.md)
    * [Oos](cli-reference/datasource/add/oos.md)
    * [Pcloud](cli-reference/datasource/add/pcloud.md)
    * [Premiumizeme](cli-reference/datasource/add/premiumizeme.md)
    * [Putio](cli-reference/datasource/add/putio.md)
    * [Qingstor](cli-reference/datasource/add/qingstor.md)
    * [AWS S3 and compliant](cli-reference/datasource/add/s3.md)
    * [Seafile](cli-reference/datasource/add/seafile.md)
    * [Sftp](cli-reference/datasource/add/sftp.md)
    * [Sharefile](cli-reference/datasource/add/sharefile.md)
    * [Sia](cli-reference/datasource/add/sia.md)
    * [Smb](cli-reference/datasource/add/smb.md)
    * [Storj](cli-reference/datasource/add/storj.md)
    * [Sugarsync](cli-reference/datasource/add/sugarsync.md)
    * [Swift](cli-reference/datasource/add/swift.md)
    * [Uptobox](cli-reference/datasource/add/uptobox.md)
    * [Webdav](cli-reference/datasource/add/webdav.md)
    * [Yandex](cli-reference/datasource/add/yandex.md)
    * [Zoho](cli-reference/datasource/add/zoho.md)
  * [List](cli-reference/datasource/list.md)
  * [Status](cli-reference/datasource/status.md)
  * [Remove](cli-reference/datasource/remove.md)
  * [Check](cli-reference/datasource/check.md)
  * [Update](cli-reference/datasource/update.md)
  * [Rescan](cli-reference/datasource/rescan.md)
  * [Daggen](cli-reference/datasource/daggen.md)
  * [Inspect](cli-reference/datasource/inspect/README.md)
    * [Chunks](cli-reference/datasource/inspect/chunks.md)
    * [Items](cli-reference/datasource/inspect/items.md)
    * [Dags](cli-reference/datasource/inspect/dags.md)
    * [Chunkdetail](cli-reference/datasource/inspect/chunkdetail.md)
    * [Itemdetail](cli-reference/datasource/inspect/itemdetail.md)
    * [Path](cli-reference/datasource/inspect/path.md)
* [Wallet](cli-reference/wallet/README.md)
  * [Import](cli-reference/wallet/import.md)
  * [List](cli-reference/wallet/list.md)
  * [Add Remote](cli-reference/wallet/add-remote.md)
  * [Remove](cli-reference/wallet/remove.md)
* [Tool](cli-reference/tool/README.md)
  * [Extract Car](cli-reference/tool/extract-car.md)

<!-- cli end -->

## 🌐 Web API Reference
<!-- webapi begin -->

* [Admin](web-api-reference/admin.md)
* [Data Source](web-api-reference/data-source.md)
* [Dataset](web-api-reference/dataset.md)
* [Deal Schedule](web-api-reference/deal-schedule.md)
* [Deal](web-api-reference/deal.md)
* [Metadata](web-api-reference/metadata.md)
* [Wallet Association](web-api-reference/wallet-association.md)
* [Wallet](web-api-reference/wallet.md)
* [Specification](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)

<!-- webapi end -->

## Technical Architecture (for developers)

* [Data Preparation](technical-architecture/data-preparation.md)

## ❓ FAQ

* [Database is locked](faq/database-is-locked.md)
* [File deletion](faq/file-deletion.md)
