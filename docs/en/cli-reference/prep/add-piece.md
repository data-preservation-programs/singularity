# Manually add piece info to a preparation. This is useful for pieces prepared by external tools.

{% code fullWidth="true" %}
```
NAME:
   singularity prep add-piece - Manually add piece info to a preparation. This is useful for pieces prepared by external tools.

USAGE:
   singularity prep add-piece [command options] <preparation_id>

CATEGORY:
   Piece Management

OPTIONS:
   --piece-cid value   CID of the piece
   --piece-size value  Size of the piece (default: "32GiB")
   --file-path value   Path to the CAR file, used to determine the size of the file and root CID
   --root-cid value    Root CID of the CAR file
   --help, -h          show help
```
{% endcode %}
