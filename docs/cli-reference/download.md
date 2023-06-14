# Download a CAR file from the metadata API

```
NAME:
   singularity download - Download a CAR file from the metadata API

USAGE:
   singularity download [command options] PIECE_CID

CATEGORY:
   Utility

OPTIONS:
   --api value  URL of the metadata API (default: "http://127.0.0.1:9090")

   HTTP data source

   --http-header value, -H value [ --http-header value, -H value ]  http headers to be passed with the request (i.e. key=value). The value shoud not be encoded [$HTTP_HEADER]

   S3 data source

   --s3-access-key-id value      IAM access key ID [$AWS_ACCESS_KEY_ID]
   --s3-endpoint value           Custom S3 endpoint [$S3_ENDPOINT]
   --s3-region value             S3 region to use with AWS S3 [$S3_REGION]
   --s3-secret-access-key value  IAM secret access key [$AWS_SECRET_ACCESS_KEY]

```
