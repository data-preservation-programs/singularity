# File deletion

Currently there is no good way to handle File deletion with the current Filecoin protocol. It is extremely expensive to terminate sectors for storage provider.

Singularity ignores file deletion during datasource rescan, and will include deleted file in the final data source folder hierarchy.&#x20;
