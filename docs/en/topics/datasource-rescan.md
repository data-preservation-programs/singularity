# Datasource rescan

By default, the data source is only scanned once, however Singularity provides an option to rescan the data source after a set interval.

```sh
singularity datasource add <type> --rescan-interval value
```

Alternatively, you may also trigger the rescan manually

```sh
singularity datasource rescan datasource_id
```

During a rescan, all new files will be queued up for preparation. All deleted files will be ignored.

#### File Versioning

For files that have changed, a new version of the file will be queued up for preparation and the directory CID will be updated to use the latest version of files.

The logic of whether a file has changed is determined by below steps:

1. If the data source provides a hash value of the file (i.e. Etag), then a new version will be created if the hash value has changed
2. Otherwise, if the data source provides the last modified time of the file, then a new version will be created if such value has changed or the file size has changed
3. Otherwise, use file size to determine whether a new version should be created

It's still possible to miss some file versions if the same file is overriden multiple times between rescan. To ensure all file versions are captured, user should use the [push-and-upload.md](push-and-upload.md "mention") to let singularity know each time a file should be updated.



