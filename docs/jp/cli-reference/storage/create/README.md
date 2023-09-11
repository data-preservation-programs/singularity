# 新しいストレージを作成し、ソースまたは出力として使用できます

{% code fullWidth="true" %}
```
NAME:
   singularity storage create - 新しいストレージを作成し、ソースまたは出力として使用できます

使用法:
   singularity storage create command [command options] [arguments...]

コマンド:
   acd              Amazon Drive
   azureblob        Microsoft Azure Blob Storage
   b2               Backblaze B2
   box              Box
   drive            Google Drive
   dropbox          Dropbox
   fichier          1Fichier
   filefabric       Enterprise File Fabric
   ftp              FTP
   gcs              Google Cloud Storage（これはGoogle Driveではありません）
   gphotos          Google Photos
   hdfs             Hadoop分散ファイルシステム
   hidrive          HiDrive
   http             HTTP
   internetarchive  Internet Archive
   jottacloud       Jottacloud
   koofr            Koofr、Digi Storage、およびその他のKoofr互換ストレージプロバイダー
   local            ローカルディスク
   mailru           Mail.ru Cloud
   mega             Mega
   netstorage       Akamai NetStorage
   onedrive         Microsoft OneDrive
   opendrive        OpenDrive
   oos              Oracle Cloud Infrastructure Object Storage
   pcloud           Pcloud
   premiumizeme     premiumize.me
   putio            Put.io
   qingstor         QingCloud Object Storage
   s3               Amazon S3準拠のストレージプロバイダー（AWS、Alibaba、Ceph、China Mobile、Cloudflare、ArvanCloud、DigitalOcean、Dreamhost、Huawei OBS、IBM COS、IDrive e2、IONOS Cloud、Liara、Lyve Cloud、Minio、Netease、RackCorp、Scaleway、SeaweedFS、StackPath、Storj、Tencent COS、Qiniu、Wasabiを含む）
   seafile          seafile
   sftp             SSH/SFTP
   sharefile        Citrix Sharefile
   sia              Sia分散クラウド
   smb              SMB / CIFS
   storj            Storj分散クラウドストレージ
   sugarsync        Sugarsync
   swift            OpenStack Swift（Rackspace Cloud Files、Memset Memstore、OVH）
   uptobox          Uptobox
   webdav           WebDAV
   yandex           Yandex Disk
   zoho             Zoho
   help, h          コマンドのリストを表示またはコマンドのヘルプを表示

オプション:
   --help, -h  ヘルプを表示する
```
{% endcode %}