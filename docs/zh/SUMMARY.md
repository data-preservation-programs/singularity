# 目录

## 概述

* [什么是Singularity](README.md)
* [版本1或版本2](overview/v1-or-v2.md)
* [当前状态](overview/current-status.md)
* [相关项目](overview/related-projects.md)

## 安装

* [从源代码安装](installation/install-from-source.md)
* [从Docker安装](installation/install-from-docker.md)
* [部署到生产环境](installation/deploy-to-production.md)

## 数据准备

* [创建数据集](data-preparation/create-a-dataset.md)
* [添加数据源](data-preparation/add-a-data-source.md)
* [启动数据集处理器](data-preparation/start-dataset-worker.md)
* [为数据源创建DAG](data-preparation/create-dag-for-the-data-source.md)

## 内容分发

* [分发CAR文件](content-distribution/distribute-car-files.md)
* [文件检索（暂存）](content-distribution/file-retrieval-staging.md)

## 交易处理

* [交易处理前提条件](deal-making/deal-making-prerequisite.md)
* [创建交易计划](deal-making/create-a-deal-schedule.md)
* [自助服务提供商（SP Self Service）](deal-making/sp-self-service.md)

## 检索

* [概述](retrieval/overview.md)

## 专题

* [加密](topics/encryption.md)
* [内联准备](topics/inline-preparation.md)
* [数据源重新扫描](topics/datasource-rescan.md)
* [推送和上传](topics/push-and-upload.md)
* [性能基准](topics/benchmark.md)

## 💻 CLI 参考
<!-- cli begin -->

* [菜单](cli-reference/README.md)
* [快速准备](cli-reference/ez-prep.md)
* [管理](cli-reference/admin/README.md)
  * [初始化](cli-reference/admin/init.md)
  * [重置](cli-reference/admin/reset.md)
  * [迁移](cli-reference/admin/migrate.md)
* [版本](cli-reference/version.md)
* [下载](cli-reference/download.md)
* [交易](cli-reference/deal/README.md)
  * [计划](cli-reference/deal/schedule/README.md)
    * [创建](cli-reference/deal/schedule/create.md)
    * [列表](cli-reference/deal/schedule/list.md)
    * [暂停](cli-reference/deal/schedule/pause.md)
    * [恢复](cli-reference/deal/schedule/resume.md)
  * [SPADE策略](cli-reference/deal/spade-policy/README.md)
    * [创建](cli-reference/deal/spade-policy/create.md)
    * [列表](cli-reference/deal/spade-policy/list.md)
    * [移除](cli-reference/deal/spade-policy/remove.md)
  * [手动发送](cli-reference/deal/send-manual.md)
  * [列表](cli-reference/deal/list.md)
* [运行](cli-reference/run/README.md)
  * [API](cli-reference/run/api.md)
  * [数据集处理器](cli-reference/run/dataset-worker.md)
  * [内容提供者](cli-reference/run/content-provider.md)
  * [交易追踪器](cli-reference/run/deal-tracker.md)
  * [交易商](cli-reference/run/dealmaker.md)
  * [SPADE API](cli-reference/run/spade-api.md)
* [数据集](cli-reference/dataset/README.md)
  * [创建](cli-reference/dataset/create.md)
  * [列表](cli-reference/dataset/list.md)
  * [更新](cli-reference/dataset/update.md)
  * [移除](cli-reference/dataset/remove.md)
  * [添加钱包](cli-reference/dataset/add-wallet.md)
  * [钱包列表](cli-reference/dataset/list-wallet.md)
  * [移除钱包](cli-reference/dataset/remove-wallet.md)
  * [添加片段](cli-reference/dataset/add-piece.md)
  * [片段列表](cli-reference/dataset/list-pieces.md)
* [数据源](cli-reference/datasource/README.md)
  * [添加](cli-reference/datasource/add/README.md)
    * [Amazon Drive](cli-reference/datasource/add/acd.md)
    * [Microsoft Azure Blob Storage](cli-reference/datasource/add/azureblob.md)
    * [Backblaze B2](cli-reference/datasource/add/b2.md)
    * [Box](cli-reference/datasource/add/box.md)
    * [Google Drive](cli-reference/datasource/add/drive.md)
    * [Dropbox](cli-reference/datasource/add/dropbox.md)
    * [1Fichier](cli-reference/datasource/add/fichier.md)
    * [Enterprise File Fabric](cli-reference/datasource/add/filefabric.md)
    * [FTP](cli-reference/datasource/add/ftp.md)
    * [Google Cloud Storage](cli-reference/datasource/add/gcs.md)
    * [Google Photos](cli-reference/datasource/add/gphotos.md)
    * [Hadoop分布式文件系统](cli-reference/datasource/add/hdfs.md)
    * [HiDrive](cli-reference/datasource/add/hidrive.md)
    * [HTTP](cli-reference/datasource/add/http.md)
    * [Internet Archive](cli-reference/datasource/add/internetarchive.md)
    * [Jottacloud](cli-reference/datasource/add/jottacloud.md)
    * [Koofr / Digi Storage](cli-reference/datasource/add/koofr.md)
    * [本地磁盘](cli-reference/datasource/add/local.md)
    * [Mail.ru Cloud](cli-reference/datasource/add/mailru.md)
    * [Mega](cli-reference/datasource/add/mega.md)
    * [Akamai NetStorage](cli-reference/datasource/add/netstorage.md)
    * [Microsoft OneDrive](cli-reference/datasource/add/onedrive.md)
    * [OpenDrive](cli-reference/datasource/add/opendrive.md)
    * [Oracle Cloud Infrastructure Object Storage](cli-reference/datasource/add/oos.md)
    * [Pcloud](cli-reference/datasource/add/pcloud.md)
    * [premiumize.me](cli-reference/datasource/add/premiumizeme.md)
    * [Put.io](cli-reference/datasource/add/putio.md)
    * [QingCloud Object Storage](cli-reference/datasource/add/qingstor.md)
    * [AWS S3 and compliant](cli-reference/datasource/add/s3.md)
    * [Seafile](cli-reference/datasource/add/seafile.md)
    * [SSH/SFTP](cli-reference/datasource/add/sftp.md)
    * [Citrix Sharefile](cli-reference/datasource/add/sharefile.md)
    * [Sia分布式云](cli-reference/datasource/add/sia.md)
    * [SMB / CIFS](cli-reference/datasource/add/smb.md)
    * [Storj分布式云存储](cli-reference/datasource/add/storj.md)
    * [Sugarsync](cli-reference/datasource/add/sugarsync.md)
    * [OpenStack Swift（Rackspace Cloud Files，Memset Memstore，OVH）](cli-reference/datasource/add/swift.md)
    * [Uptobox](cli-reference/datasource/add/uptobox.md)
    * [WebDAV](cli-reference/datasource/add/webdav.md)
    * [Yandex Disk](cli-reference/datasource/add/yandex.md)
    * [Zoho](cli-reference/datasource/add/zoho.md)
  * [列表](cli-reference/datasource/list.md)
  * [状态](cli-reference/datasource/status.md)
  * [移除](cli-reference/datasource/remove.md)
  * [检查](cli-reference/datasource/check.md)
  * [更新](cli-reference/datasource/update.md)
  * [重新扫描](cli-reference/datasource/rescan.md)
  * [DAG生成器](cli-reference/datasource/daggen.md)
  * [检查](cli-reference/datasource/inspect/README.md)
    * [数据块](cli-reference/datasource/inspect/chunks.md)
    * [数据项](cli-reference/datasource/inspect/items.md)
    * [DAGs](cli-reference/datasource/inspect/dags.md)
    * [块详情](cli-reference/datasource/inspect/chunkdetail.md)
    * [项详情](cli-reference/datasource/inspect/itemdetail.md)
    * [路径](cli-reference/datasource/inspect/path.md)
* [钱包](cli-reference/wallet/README.md)
  * [导入](cli-reference/wallet/import.md)
  * [列表](cli-reference/wallet/list.md)
  * [添加远程钱包](cli-reference/wallet/add-remote.md)
  * [移除](cli-reference/wallet/remove.md)
* [工具](cli-reference/tool/README.md)
  * [提取CAR文件](cli-reference/tool/extract-car.md)

<!-- cli end -->

## 🌐 Web API 参考
<!-- webapi begin -->

* [管理](web-api-reference/admin.md)
* [数据源](web-api-reference/data-source.md)
* [数据集](web-api-reference/dataset.md)
* [交易计划](web-api-reference/deal-schedule.md)
* [交易](web-api-reference/deal.md)
* [元数据](web-api-reference/metadata.md)
* [钱包关联](web-api-reference/wallet-association.md)
* [钱包](web-api-reference/wallet.md)
* [规范](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)

<!-- webapi end -->

## 技术架构（针对开发人员）

* [数据准备](technical-architecture/data-preparation.md)

## ❓ 常见问题

* [数据库已锁定](faq/database-is-locked.md)
* [文件删除](faq/file-deletion.md)