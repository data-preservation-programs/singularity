# Add a piece to a preparation. If the piece exists in the database, metadata is copied. Otherwise, --piece-size is required.

{% code fullWidth="true" %}
```
NAME:
   singularity prep add-piece - Add a piece to a preparation. If the piece exists in the database, metadata is copied. Otherwise, --piece-size is required.

USAGE:
   singularity prep add-piece [command options] <preparation id|name>

CATEGORY:
   Piece Management

DESCRIPTION:
   Add a piece to a preparation by piece CID.

   If the piece CID already exists in the database (from a previous preparation),
   the metadata (size, root CID, etc.) is automatically copied. This is useful for
   reorganizing pieces between preparations (e.g., consolidating small pieces for
   batch deal scheduling).

   For external pieces not in the database, --piece-size must be provided.

   NOTE: This is an advanced feature. When overriding file-path for an existing
   piece, ensure the new file has matching content. File paths must be accessible
   to any workers or content providers that will serve this piece.

OPTIONS:
   --piece-cid value   CID of the piece
   --piece-size value  Size of the piece (e.g. 32GiB). Required only for external pieces not in database.
   --file-path value   Path to the CAR file, used to determine the size of the file and root CID
   --root-cid value    Root CID of the CAR file
   --file-size value   Size of the CAR file, this is required for boost online deal. If not set, it will be determined from the file path if provided. (default: 0)
   --help, -h          show help
```
{% endcode %}
