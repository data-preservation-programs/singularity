---
description: Connect to a data source that needs to be prepared
---

# Add a data source

## Add a local file system data source

The most command data source is the local file system. To add a folder as a data source to the dataset:

```sh
singularity datasource add local my_dataset /mnt/dataset/folder
```

Next, you would want to [start-dataset-worker.md](start-dataset-worker.md "mention") or checkout [#other-data-source-types](add-a-data-source.md#other-data-source-types "mention")

## Other Data source types

Singularity is deeply integrated with [Rclone](https://rclone.org/overview/) and supports all backend that Rclone supports

<table data-view="cards"><thead><tr><th></th><th data-hidden data-card-target data-type="content-ref"></th></tr></thead><tbody><tr><td>Amazon Drive</td><td></td></tr><tr><td>Microsoft Azure Blob Storage</td><td></td></tr><tr><td>Backblaze B2</td><td></td></tr><tr><td>Box</td><td></td></tr><tr><td>Google Drive</td><td></td></tr><tr><td>Dropbox</td><td></td></tr><tr><td>1Fichier</td><td></td></tr><tr><td>Enterprise File Fabric</td><td></td></tr><tr><td>Google Cloud Storage</td><td></td></tr><tr><td>Google Photos</td><td></td></tr><tr><td>Hadoop distributed file system</td><td></td></tr><tr><td>HiDrive</td><td></td></tr><tr><td>HTTP</td><td></td></tr><tr><td>Internet Archive</td><td></td></tr><tr><td>Jottacloud</td><td></td></tr><tr><td>Koofr / Digi Storage</td><td></td></tr><tr><td>Local Disk</td><td><a href="../cli-reference/data-source/add-data-source/local-file-system.md">local-file-system.md</a></td></tr><tr><td>Mail.ru Cloud</td><td></td></tr><tr><td>Mega</td><td></td></tr><tr><td>Akamai NetStorage</td><td></td></tr><tr><td>Microsoft OneDrive</td><td></td></tr><tr><td>OpenDrive</td><td></td></tr><tr><td>Oracle Cloud Infrastructure Object Storage</td><td></td></tr><tr><td>Pcloud</td><td></td></tr><tr><td>premiumize.me</td><td></td></tr><tr><td>Put.io</td><td></td></tr><tr><td>QingCloud Object Storage</td><td></td></tr><tr><td>AWS / other S3</td><td></td></tr><tr><td>seafile</td><td></td></tr><tr><td>SSH/SFTP</td><td></td></tr><tr><td>Citrix Sharefile</td><td></td></tr><tr><td>Sia Decentralized Cloud</td><td></td></tr><tr><td>SMB / CIFS</td><td></td></tr><tr><td>Storj Decentralized Cloud Storage</td><td></td></tr><tr><td>Sugarsync</td><td></td></tr><tr><td>OpenStack Swift</td><td></td></tr><tr><td>Uptobox</td><td></td></tr><tr><td>WebDAV</td><td></td></tr><tr><td>Yandex Disk</td><td></td></tr><tr><td>Zoho</td><td></td></tr></tbody></table>
