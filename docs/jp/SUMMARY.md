# 目次

## 概要 <a href="#overview" id="overview"></a>

* [Singularityとは](README.md)
* [V1またはV2](overview/v1-or-v2.md)
* [現在の状況](overview/current-status.md)
* [関連プロジェクト](overview/related-projects.md)

## インストール <a href="#installation" id="installation"></a>

* [ソースからのインストール](installation/install-from-source.md)
* [Dockerからのインストール](installation/install-from-docker.md)
* [本番環境への展開](installation/deploy-to-production.md)

## データの準備 <a href="#data-preparation" id="data-preparation"></a>

* [データセットの作成](data-preparation/create-a-dataset.md)
* [データソースの追加](data-preparation/add-a-data-source.md)
* [データセットワーカーの起動](data-preparation/start-dataset-worker.md)
* [データソース用のDAGの作成](data-preparation/create-dag-for-the-data-source.md)

## コンテンツの配布 <a href="#content-distribution" id="content-distribution"></a>

* [CARファイルの配布](content-distribution/distribute-car-files.md)
* [ファイルの取得（ステージング）](content-distribution/file-retrieval-staging.md)

## 取引の実行 <a href="#deal-making" id="deal-making"></a>

* [取引の実行の前提条件](deal-making/deal-making-prerequisite.md)
* [取引スケジュールの作成](deal-making/create-a-deal-schedule.md)
* [SPセルフサービス](deal-making/sp-self-service.md)

## 取得 <a href="#retrieval" id="retrieval"></a>

* [概要](retrieval/overview.md)

## トピックス <a href="#topics" id="topics"></a>

* [暗号化](topics/encryption.md)
* [インライン準備](topics/inline-preparation.md)
* [データソースの再スキャン](topics/datasource-rescan.md)
* [プッシュとアップロード](topics/push-and-upload.md)
* [ベンチマーク](topics/benchmark.md)

## 💻 CLIリファレンス <a href="#cli-reference" id="cli-reference"></a>
<!-- cli begin -->

* [メニュー](cli-reference/README.md)
* [Ez Prep](cli-reference/ez-prep.md)
* [管理者](cli-reference/admin/README.md)
  * [初期化](cli-reference/admin/init.md)
  * [リセット](cli-reference/admin/reset.md)
  * [マイグレーション](cli-reference/admin/migrate.md)
* [バージョン](cli-reference/version.md)
* [ダウンロード](cli-reference/download.md)
* [取引](cli-reference/deal/README.md)
  * [スケジュール](cli-reference/deal/schedule/README.md)
    * [作成](cli-reference/deal/schedule/create.md)
    * [一覧](cli-reference/deal/schedule/list.md)
    * [一時停止](cli-reference/deal/schedule/pause.md)
    * [再開](cli-reference/deal/schedule/resume.md)
  * [Spadeポリシー](cli-reference/deal/spade-policy/README.md)
    * [作成](cli-reference/deal/spade-policy/create.md)
    * [一覧](cli-reference/deal/spade-policy/list.md)
    * [削除](cli-reference/deal/spade-policy/remove.md)
  * [手動送信](cli-reference/deal/send-manual.md)
  * [一覧](cli-reference/deal/list.md)
* [実行](cli-reference/run/README.md)
  * [API](cli-reference/run/api.md)
  * [データセットワーカー](cli-reference/run/dataset-worker.md)
  * [コンテンツプロバイダー](cli-reference/run/content-provider.md)
  * [取引トラッカー](cli-reference/run/deal-tracker.md)
  * [ディールメーカー](cli-reference/run/dealmaker.md)
  * [Spade API](cli-reference/run/spade-api.md)
* [データセット](cli-reference/dataset/README.md)
  * [作成](cli-reference/dataset/create.md)
  * [一覧](cli-reference/dataset/list.md)
  * [更新](cli-reference/dataset/update.md)
  * [削除](cli-reference/dataset/remove.md)
  * [ウォレットの追加](cli-reference/dataset/add-wallet.md)
  * [ウォレットの一覧](cli-reference/dataset/list-wallet.md)
  * [ウォレットの削除](cli-reference/dataset/remove-wallet.md)
  * [ピースの追加](cli-reference/dataset/add-piece.md)
  * [ピースの一覧](cli-reference/dataset/list-pieces.md)
* [データソース](cli-reference/datasource/README.md)
  * [追加](cli-reference/datasource/add/README.md)
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
    * [Hadoop distributed file system](cli-reference/datasource/add/hdfs.md)
    * [HiDrive](cli-reference/datasource/add/hidrive.md)
    * [HTTP](cli-reference/datasource/add/http.md)
    * [Internet Archive](cli-reference/datasource/add/internetarchive.md)
    * [Jottacloud](cli-reference/datasource/add/jottacloud.md)
    * [Koofr / Digi Storage](cli-reference/datasource/add/koofr.md)
    * [ローカルディスク](cli-reference/datasource/add/local.md)
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
    * [seafile](cli-reference/datasource/add/seafile.md)
    * [SSH/SFTP](cli-reference/datasource/add/sftp.md)
    * [Citrix Sharefile](cli-reference/datasource/add/sharefile.md)
    * [Sia Decentralized Cloud](cli-reference/datasource/add/sia.md)
    * [SMB / CIFS](cli-reference/datasource/add/smb.md)
    * [Storj Decentralized Cloud Storage](cli-reference/datasource/add/storj.md)
    * [Sugarsync](cli-reference/datasource/add/sugarsync.md)
    * [OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)](cli-reference/datasource/add/swift.md)
    * [Uptobox](cli-reference/datasource/add/uptobox.md)
    * [WebDAV](cli-reference/datasource/add/webdav.md)
    * [Yandex Disk](cli-reference/datasource/add/yandex.md)
    * [Zoho](cli-reference/datasource/add/zoho.md)
  * [一覧](cli-reference/datasource/list.md)
  * [ステータス](cli-reference/datasource/status.md)
  * [削除](cli-reference/datasource/remove.md)
  * [チェック](cli-reference/datasource/check.md)
  * [更新](cli-reference/datasource/update.md)
  * [再スキャン](cli-reference/datasource/rescan.md)
  * [Daggen](cli-reference/datasource/daggen.md)
  * [調査](cli-reference/datasource/inspect/README.md)
    * [チャンク](cli-reference/datasource/inspect/chunks.md)
    * [アイテム](cli-reference/datasource/inspect/items.md)
    * [ダグ](cli-reference/datasource/inspect/dags.md)
    * [チャンクの詳細](cli-reference/datasource/inspect/chunkdetail.md)
    * [アイテムの詳細](cli-reference/datasource/inspect/itemdetail.md)
    * [パス](cli-reference/datasource/inspect/path.md)
* [ウォレット](cli-reference/wallet/README.md)
  * [インポート](cli-reference/wallet/import.md)
  * [一覧](cli-reference/wallet/list.md)
  * [リモートの追加](cli-reference/wallet/add-remote.md)
  * [削除](cli-reference/wallet/remove.md)
* [ツール](cli-reference/tool/README.md)
  * [Carの抽出](cli-reference/tool/extract-car.md)

<!-- cli end -->

## 🌐 Web APIリファレンス <a href="#web-api-reference" id="web-api-reference"></a>
<!-- webapi begin -->

* [管理者](web-api-reference/admin.md)
* [データソース](web-api-reference/data-source.md)
* [データセット](web-api-reference/dataset.md)
* [取引スケジュール](web-api-reference/deal-schedule.md)
* [取引](web-api-reference/deal.md)
* [メタデータ](web-api-reference/metadata.md)
* [ウォレット関連](web-api-reference/wallet-association.md)
* [ウォレット](web-api-reference/wallet.md)
* [仕様](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)

<!-- webapi end -->

## 開発者向け技術アーキテクチャ <a href="#technical-architecture" id="technical-architecture"></a>

* [データの準備](technical-architecture/data-preparation.md)

## ❓ FAQ <a href="#faq" id="faq"></a>

* [データベースがロックされています](faq/database-is-locked.md)
* [ファイルの削除](faq/file-deletion.md)
