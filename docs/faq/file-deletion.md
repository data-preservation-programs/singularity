# File deletion

Currently there is no good way to handle File deletion with the current Filecoin protocol. It is extremely expensive to terminate sectors for storage provider.

Singularity ignores file deletion during datasource rescan, and will include deleted file in the final data source folder hierarchy. It's not technically difficult to exclude deleted file in such structure, let us know if this is an important feature to you.
