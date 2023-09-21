# Performance Tuning in Singularity

Singularity offers a range of configurations allowing users to optimize data preparation performance. This guide elucidates these configurations and provides instructions for tuning them effectively.

## Inline Preparation
* **Description**: Inline preparation eradicates the need for extra disk space to store CAR files. However, it incurs a minor overhead in database lookups and storage.
* **Implications**: The overhead is usually negligible but can become significant for datasets containing many small files.
* **Configuration**: To disable, Use `--no-inline` with `singularity prep create`.
* **Further Reading**: [Inline Preparation](../topics/inline-preparation.md)

## DAG Updates
* **Description**:
  During preparation, Singularity refreshes the DAG and CID for each directory, which is useful for real-time tracking of changes.
* **Implications**:
  This introduces a slight database overhead as directories get updated each time a CAR file is prepared.
* **Configuration**:
  To disable, use `--no-dag` with `singularity prep create`.

To turn off DAG updates, set `--no-dag` when creating a preparation using `singularity prep create`.

## Parallelism in Data Preparation

### Scanning
* **Description**: Scanning involves traversing the source storage to curate a file list. While fast on local storage, it might be sluggish for remote storage like S3.
* **Configuration**:
  * **Enable Parallelism**: Use `--client-scan-concurrency <number>` with `singularity storage create` or `singularity storage update`.
  * **Note**: Enabling can cause files to be processed in a non-deterministic order.

### Packing
* **Description**: Packing merges multiple files into a single CAR file, a both CPU-intensive and IO-intensive operation. For remote storage with network limitations, increasing parallelism is beneficial.
* **Configuration**:
    * **Adjust Parallelism**: Use `--concurrency <number>` with `singularity run dataset-worker`.

## Use Server's Last Modified Time
* **Description**: Some remote storages such as `AWS S3` offer custom `mtime` and server-side last modified time. By default, Singularity checks for custom `mtime` and uses it if available. Otherwise, it uses the server's last modified time.
* **Implication**: Skip checking custom `mtime` and directly use server's last modified time can reduce the number of requests to the remote storage.
* **Configuration**: To prioritize server's time and bypass object metadata fetching, use `--client-use-server-mod-time` with `singularity storage create` or `singularity storage update`.

## Retry Strategy
### Retry on Network Request
* **Description**: For failed remote folder listings or file openings, Singularity leverages RClone's retry mechanism.
* **Configuration**: To increase Retries, use `--client-low-level-retries <number>` with `singularity storage create` or `singularity storage update`.

## Retry on Network IO
* **Description**: Despite successful network requests, network IO can fail due to unstable network connections. Singularity supports retrying and resuming from the last successful point.
* **Configuration**: Use below flags with `singularity storage create` or `singularity storage update`.
```shell
 --client-retry-backoff value      # Delay backoff for retrying IO read errors (default: 1s)
 --client-retry-backoff-exp value  # Exponential delay backoff for retrying IO read errors (default: 1.0)
 --client-retry-delay value        # Initial delay before retrying IO read errors (default: 1s)
 --client-retry-max value          # Max number of retries for IO read errors (default: 10)
```

## Skip Inaccessible Files
* **Description**: Permissions might prevent accessing certain files from remote storage. These issues may only surface when attempting to open the file, causing the packing job to fail.
* **Configuration**: To skip inaccessible files, use `--client-skip-inaccessible-files` with `singularity storage create` or `singularity storage update`.
