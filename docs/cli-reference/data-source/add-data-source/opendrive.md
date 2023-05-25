# OpenDrive

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add opendrive - OpenDrive

USAGE:
   singularity datasource add opendrive [command options] <dataset_name> <source_path>

DESCRIPTION:
   --opendrive-username
      Username.

   --opendrive-password
      Password.

   --opendrive-encoding
      The encoding for the backend.

      See the [encoding section in the overview](/overview/#encoding) for more info.

   --opendrive-chunk-size
      Files will be uploaded in chunks this size.

      Note that these chunks are buffered in memory so increasing them will
      increase memory use.


OPTIONS:
   --help, -h                  show help
   --opendrive-password value  Password. [$OPENDRIVE_PASSWORD]
   --opendrive-username value  Username. [$OPENDRIVE_USERNAME]

   Advanced Options

   --opendrive-chunk-size value  Files will be uploaded in chunks this size. (default: "10Mi") [$OPENDRIVE_CHUNK_SIZE]
   --opendrive-encoding value    The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot") [$OPENDRIVE_ENCODING]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)

```
{% endcode %}
