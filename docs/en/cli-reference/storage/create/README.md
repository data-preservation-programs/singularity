# Create a new storage which can be used as source or output

{% code fullWidth="true" %}
```
NAME:
   singularity storage create - Create a new storage which can be used as source or output

USAGE:
   singularity storage create command [command options]

COMMANDS:
   azureblob        Microsoft Azure Blob Storage
   b2               Backblaze B2
   box              Box
   drive            Google Drive
   dropbox          Dropbox
   fichier          1Fichier
   filefabric       Enterprise File Fabric
   ftp              FTP
   gcs              Google Cloud Storage (this is not Google Drive)
   gphotos          Google Photos
   hdfs             Hadoop distributed file system
   hidrive          HiDrive
   http             HTTP
   internetarchive  Internet Archive
   jottacloud       Jottacloud
   koofr            Koofr, Digi Storage and other Koofr-compatible storage providers
   local            Local Disk
   mailru           Mail.ru Cloud
   mega             Mega
   netstorage       Akamai NetStorage
   onedrive         Microsoft OneDrive
   oos              Oracle Cloud Infrastructure Object Storage
   opendrive        OpenDrive
   pcloud           Pcloud
   premiumizeme     premiumize.me
   putio            Put.io
   qingstor         QingCloud Object Storage
   s3               Amazon S3 Compliant Storage Providers including AWS, Alibaba, ArvanCloud, BizflyCloud, Ceph, ChinaMobile, Cloudflare, Cubbit, DigitalOcean, Dreamhost, Exaba, FileLu, FlashBlade, GCS, Hetzner, HuaweiOBS, IBMCOS, IDrive, Intercolo, IONOS, Leviia, Liara, Linode, LyveCloud, Magalu, Mega, Minio, Netease, Outscale, OVHcloud, Petabox, Qiniu, Rabata, RackCorp, Rclone, Scaleway, SeaweedFS, Selectel, Servercore, SpectraLogic, StackPath, Storj, Synology, TencentCOS, Wasabi, Zata, Other
   seafile          seafile
   sftp             SSH/SFTP
   sharefile        Citrix Sharefile
   sia              Sia Decentralized Cloud
   smb              SMB / CIFS
   storj            Storj Decentralized Cloud Storage
   sugarsync        Sugarsync
   swift            OpenStack Swift (Rackspace Cloud Files, Blomp Cloud Storage, Memset Memstore, OVH)
   union            Union merges the contents of several upstream fs
   webdav           WebDAV
   yandex           Yandex Disk
   zoho             Zoho
   help, h          Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```
{% endcode %}
