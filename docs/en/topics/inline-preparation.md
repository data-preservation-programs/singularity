# Inline Preparation

## Overview

The traditional method for data preparation involves converting the original data source, typically a folder on a local file system, into a collection of CAR files, each smaller than 32GiB. This method necessitates that Data Preparers possess twice the storage capacity, which can be very expensive. For instance, preparing a 1 PiB dataset would also require another 1 PiB of storage for the CAR files, resulting in a total requirement of 2 PiB of storage space.

Inline preparation solves this problem by mapping the blocks of CAR files back to the original data source so that there is no need to store the exported CAR files.

<div align="center">

<img src="https://github.com/data-preservation-programs/singularity/assets/12418265/4292faf1-9f01-4b7c-b79f-67b0bc1e2acc" alt="Traditional data prep diagram" width="500">

 

<img src="https://github.com/data-preservation-programs/singularity/assets/12418265/f5cfc209-5e38-4bb9-8cd9-f1aeffaf284d" alt="Inline data prep diagram" width="500">

</div>

## How CAR retrieval works

With inline preparation, the CAR files can be served via HTTP using the metadata database and the original data source because it knows how to map byte ranges of the CAR file back to the original data source.

To serve CAR files via HTTP, simply start the content provider

```sh
singularity run content-provider
```

> Note: This command will run a local HTTP server. If you intend to make it accessible over the internet, you may want to put it behind a reverse proxy such as nginx.

This creates a potential bottleneck if the data source is already a remote storage system (i.e. S3 or FTP), as the file content would be proxied through the Singularity content provider to the Storage Provider.

We have a solution to this challenge using Singularity metadata API and Singularity download utility

To run the Singularity Metadata API

```sh
singularity run api
```

Then, to use the Singularity download utility (on the Storage Provider)

```sh
singularity download <piece_cid>
```

The Singularity metadata API will return a plan for how to assemble the CAR files from original data source and the Singularity download utility will interpret this plan and stream data from original data into a local CAR file. There is no conversion or assembling happening in the middle, everything works as streams.

The metadata API does not return any credentials required to access the data from original data source. The Storage Provider needs to get their own access to the data source and supply such credentials to the `singularity download` command.

## Overhead

Inline preparation introduces a minimal overhead, primarily in terms of the storage space required. Additionally, computational and bandwidth overhead is also minimal.

Metadata for each block of data is stored as a database row, requiring 100 bytes for every 1MiB block of prepared data. For a 1PiB dataset, this translates to a requisite of 10TiB of disk space to store the mapping metadata. While this isn't typically an issue, datasets with a large number of small files may result in a significantly higher disk overhead.

Later, when CAR files are dynamically regenerated from the original data source, it's necessary to cross-reference these mappings in the database. However, this is generally not a concern. A bandwidth of 1GB/sec equates to 1,000 database entry lookups, which is far from the bottleneck capabilities of all supported database backends. Additionally, future optimizations may further reduce this overhead.

## Enable Inline Preparation

Inline preparation is automatically enabled for datasets that don't require encryption. Upon dataset creation, when an output directory is designated, CAR files are exported to that location. CAR retrieval requests prioritize these directories. If the CAR files are removed by the user, the system reverts to fetching from the original data source.
