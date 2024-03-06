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

### Singularity Download Server
If the data source is remote, i.e. comes from S3 or FTP, and the client has enabled inline preparation, then we can save the egress by assembling the CAR files directly from the original data source.
Singularity download server allows streaming the CAR file directly from original source by first querying the content provider for the CAR file metadata which tells how to assemble the CAR file from the original files.
Since it is a local HTTP server, you may choose to use any other HTTP client to download the CAR file, even with multithreading.
```shell
singularity run download-server --metadata-api "http://content-provider:7777" --bind "127.0.0.1:8888"
wget http://127.0.0.1:8888/piece/bagaxxxxxxxxxxx
```
This method also works with boost online deal, as the client can make a boost online deal proposal to SP using below command:
```shell
boost deal --http-url "http://127.0.0.1:8888/piece/bagaxxxx"
```
Or, when using Singularity to make deals, you can create a deal schedule with below command:
```shell
deal schedule create --url-template "http://127.0.0.1:8888/piece/{PIECE_CID}"
```

### Singularity Download Utility
Similar to the download server, the download utility serves the same purpose without standing up a local HTTP server.  
```shell
singularity download bagaxxxxxxxxxxx
```
This utility communicates with the content provider service to fetch metadata about the piece. Once obtained, it uses this metadata to reconstruct the piece directly from the original data source.
