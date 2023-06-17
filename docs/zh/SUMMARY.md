# 目录

## 概览

* [Singularity 是什么](README.md)
* [V1 还是 V2](overview/v1-or-v2.md)
* [当前状态](overview/current-status.md)
* [相关项目](overview/related-projects.md)

## 安装

* [从源代码安装](installation/install-from-source.md)
* [从 Docker 安装](installation/install-from-docker.md)
* [部署到生产环境](installation/deploy-to-production.md)

## 数据准备

* [创建数据集](data-preparation/create-a-dataset.md)
* [添加数据源](data-preparation/add-a-data-source.md)
* [启动数据集 worker](data-preparation/start-dataset-worker.md)
* [为数据源创建 DAG](data-preparation/create-dag-for-the-data-source.md)

## 内容分发

* [分发 CAR 文件](content-distribution/distribute-car-files.md)
* [文件检索（暂存）](content-distribution/file-retrieval-staging.md)

## 交易达成

* [交易达成先决条件](deal-making/deal-making-prerequisite.md)
* [创建交易时间表](deal-making/create-a-deal-schedule.md)
* [SP 自服务](deal-making/sp-self-service.md)

## 数据检索

* [概览](retrieval/overview.md)

## 主题

* [加密](topics/encryption.md)
* [内联准备](topics/inline-preparation.md)
* [数据源重新扫描](topics/datasource-rescan.md)
* [推送和上传](topics/push-and-upload.md)
* [基准测试](topics/benchmark.md)

## 💻 CLI 参考手册

* [Ez Prep](cli-reference/ez-prep.md)
* [管理员](cli-reference/admin/README.md)
  * [初始化](cli-reference/admin/init.md)
  * [重置](cli-reference/admin/reset.md)
  * [迁移](cli-reference/admin/migrate.md)
* [下载](cli-reference/download.md)
* [交易](cli-reference/deal/README.md)
  * [时间表](cli-reference/deal/schedule/README.md)
    * [创建](cli-reference/deal/schedule/create.md)
    * [列表](cli-reference/deal/schedule/list.md)
    * [暂停](cli-reference/deal/schedule/pause.md)
    * [恢复](cli-reference/deal/schedule/resume.md)
  * [Spade 策略](cli-reference/deal/spade-policy/README.md)
    * [创建](cli-reference/deal/spade-policy/create.md)
    * [列表](cli-reference/deal/spade-policy/list.md)
    * [删除](cli-reference/deal/spade-policy/remove.md)
  * [手动发送](cli-reference/deal/send-manual.md)
  * [列表](cli-reference/deal/list.md)
* [运行](cli-reference/run/README.md)
  * [API](cli-reference/run/api.md)
  * [数据集 worker](cli-reference/run/dataset-worker.md)
  * [内容提供者](cli-reference/run/content-provider.md)
  * [交易达成](cli-reference/run/dealmaker.md)
  * [Spade API](cli-reference/run/spade-api.md)
* [数据集](cli-reference/dataset/README.md)
  * [创建](cli-reference/dataset/create.md)
  * [列表](cli-reference/dataset/list.md)
  * [更新](cli-reference/dataset/update.md)
  * [删除](cli-reference/dataset/remove.md)
  * [添加钱包](cli-reference/dataset/add-wallet.md)
  * [列出钱包](cli-reference/dataset/list-wallet.md)
  * [删除钱包](cli-reference/dataset/remove-wallet.md)
  * [添加条目](cli-reference/dataset/add-piece.md)
  * [列出片段](cli-reference/dataset/list-pieces.md)
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
    * [分布式文件系统 Hadoop](cli-reference/datasource/add/hdfs.md)
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
    * [青云对象存储](cli-reference/datasource/add/qingstor.md)
    * [AWS S3 和兼容](cli-reference/datasource/add/s3.md)
    * [seafile](cli-reference/datasource/add/seafile.md)
    * [SSH/ SFTP](cli-reference/datasource/add/sftp.md)
    * [Citrix Sharefile](cli-reference/datasource/add/sharefile.md)
    * [Sia Decentralized Cloud](cli-reference/datasource/add/sia.md)
    * [SMB / CIFS](cli-reference/datasource/add/smb.md)
    * [\[Storj Decentralized Cloud Storage](cli-reference/datasource/add/storjgo-install-github.com-anjor-go-fil-dataprep/cmd/data-prep@latest.md)\]
    * [Sugarsync](cli-reference/datasource/add/sugarsync.md)
    * [OpenStack Swift（Rackspace Cloud 文件，Memset Memstore，OVH）](cli-reference/datasource/add/swift.md)
    * [Uptobox](cli-reference/datasource/add/uptobox.md)
    * [WebDAV](cli-reference/datasource/add/webdav.md)
    * [Yandex Disk](cli-reference/datasource/add/yandex.md)
    * [Zoho](cli-reference/datasource/add/zoho.md)
  * [列表](cli-reference/datasource/list.md)
  * [状态](cli-reference/datasource/status.md)
  * [删除](cli-reference/datasource/remove.md)
  * [检查](cli-reference/datasource/check.md)
  * [更新](cli-reference/datasource/update.md)
  * [重新扫描](cli-reference/datasource/rescan.md)
  * [Daggen](cli-reference/datasource/daggen.md)
  * [检查](cli-reference/datasource/inspect/README.md)
    * [块](cli-reference/datasource/inspect/chunks.md)
    * [项目](cli-reference/datasource/inspect/items.md)
    * [DAG](cli-reference/datasource/inspect/dags.md)
    * [块详情](cli-reference/datasource/inspect/chunkdetail.md)
    * [项目