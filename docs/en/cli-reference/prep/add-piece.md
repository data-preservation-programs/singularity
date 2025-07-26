# Manually add piece info to a preparation. This is useful for pieces prepared by external tools.

{% code fullWidth="true" %}
```
NAME:
   singularity prep add-piece - Manually add piece info to a preparation. This is useful for pieces prepared by external tools.

USAGE:
   singularity prep add-piece [command options] <preparation id|name>

CATEGORY:
   Piece Management

OPTIONS:
   --piece-cid value   CID of the piece
   --piece-size value  Size of the piece (default: "32GiB")
   --file-path value   Path to the CAR file, used to determine the size of the file and root CID
   --root-cid value    Root CID of the CAR file
   --file-size value   Size of the CAR file, this is required for boost online deal. If not set, it will be determined from the file path if provided. (default: 0)
   --help, -h          show help
```
{% endcode %}
