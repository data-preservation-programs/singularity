# OpenDrive

{% code fullWidth="true" %}
```
NAME:
   singularity storage create opendrive - OpenDrive

USAGE:
   singularity storage create opendrive [command options] <name> <path>

DESCRIPTION:
   --username
      Username.

   --password
      Password.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --chunk_size
      Files will be uploaded in chunks this size.
      
      Note that these chunks are buffered in memory so increasing them will
      increase memory use.


OPTIONS:
   --help, -h        show help
   --password value  Password. [$PASSWORD]
   --username value  Username. [$USERNAME]

   Advanced

   --chunk_size value  Files will be uploaded in chunks this size. (default: "10Mi") [$CHUNK_SIZE]
   --encoding value    The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot") [$ENCODING]

```
{% endcode %}
