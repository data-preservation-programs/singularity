# Distributing CAR Files with Singularity

To ensure your data is easily accessible by storage providers, you need to distribute the CAR (Content Addressable Archive) files effectively.

## 1. Start the Content Provider Service

Begin by launching the content provider service. This service facilitates the download of pieces from the dataset you've prepared.

```sh
singularity run content-provider
```

## 2. Methods for CAR File Download

There are multiple ways for storage providers to download the CAR files:

### Direct HTTP Download

Providers can utilize the HTTP API exposed by the content provider service to directly download the CAR file:

```shell
wget http://127.0.0.1:7777/piece/bagaxxxxxxxxxxx
```
If you've specified an output directory during preparation, the CAR files will be sourced directly from there. However, if you used inline preparation or accidentally deleted the CAR files, the service will retrieve the content from the original data source and serve it.

### Singularity Download Utility
For providers seeking an alternative download method, especially when dealing with remote data sources like S3 or FTP, Singularity offers a dedicated download utility:
```shell
singularity download bagaxxxxxxxxxxx
```
This utility communicates with the content provider service to fetch metadata about the piece. Once obtained, it uses this metadata to reconstruct the piece directly from the original data source.
