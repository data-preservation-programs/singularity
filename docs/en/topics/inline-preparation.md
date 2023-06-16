# Inline Preparation

## Overview

Tranditional approach to data preparation is to convert original data source, most likely a folder in a local file system to a bunch of CAR files each less than 32GiB. This creates a new issue for data preparers. If they are to prepare one PiB of dataset, they need to provision another PiB of storage servers to store those CAR files.

Inline preparation solves this problem by mapping blocks of those CAR files back to the original data source so that there is no need to store the exported CAR files.

## How CAR retrieval works

With inline preparation, the CAR files can be served via HTTP using the metadata database and the original data source because it knows how to map certain bytes range of the CAR file back to the original data source.&#x20;

To serve CAR files via HTTP, simply start the content provider and optionally put it behind a reverse proxy

```sh
singularity run content-provider
```

This creates another challenge, if the data source is already a remote storage system, i.e. S3 or FTP, the file content is still proxied through the singularity content provider to the storage provider.

We have a solution to solve that using singularity metadata API and singularity download utility

```bash
singularity run api
singularity download <piece_cid>
```

Singualrity metadata API will return a plan for how to assemble the CAR files from original data source and the singularity download utility will interprete this plan and stream from original data into a local CAR file. There is no conversion or assembling happening in the middle, everything works as streams.

The metadata API does not return any credential required to access the data from original data source so storage provider needs to get their own access to the data source and supply such credential to the singularity download command.

## Overhead

Inline preparation comes with a very small overhead.

Since we need to store a database row for each block of data, we need 100 bytes for each 1MiB block of data we prepared. For 1PiB dataset, the database will 10TiB of disk space for storing such mapping metadata. This is typically not a problem but for dataset with large number of small files, the disk overhead would be larger.

Later when we regenerates CAR files on the fly from the original data source, we will need to check those mappings in the database. This is typically not a problem because a bandwidth of 1GB / sec translates to 1000 entry lookup in the database which is well below bottleneck of all supported database backend and there could be future optimization that reduces such overhead.

## Enable Inline Preparation

Inline preparation is always enabled for all dataset that does not need encryption. When dataset is created with specifying the output directory, CAR files will be exported to those directories. In such case, CAR retrieval request will be served from those directories first and fall back to original data source if those CAR files are deleted by the user.

## Related resources

[download.md](../cli-reference/download.md "mention")
