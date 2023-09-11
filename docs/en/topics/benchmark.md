# Benchmarking with Singularity

The `ez-prep` command in Singularity provides a streamlined approach to benchmarking.

## Preparing Test Data

Initially, you need to generate data for benchmarking. Sparse files are used here to remove disk IO time from the benchmark. Currently, Singularity does not perform CID deduplication, so it processes these files as random bytes.

```sh
mkdir dataset
truncate -s 1024G dataset/1T.bin
```

If you aim to include disk IO time in your benchmark, use the following method to create a random file:

```sh
dd if=/dev/urandom of=dataset/8G.bin bs=1M count=8192
```

## Using `ez-prep`
The `ez-prep` command streamlines data preparation from a local folder with minimal configurable options.

### Benchmarking Inline Preparation
Inline preparation negates the need for exporting CAR files, saving metadata directly to the database:

```sh
time singularity ez-prep --output-dir '' ./dataset
```

### Benchmarking with In-Memory Database

To minimize disk IO, opt for an in-memory database:

```sh
time singularity ez-prep --output-dir '' --database-file '' ./dataset
```

### Benchmarking with Multiple Workers

For optimal CPU core utilization, set concurrency for the benchmark. Note: each worker uses approximately 4 CPU cores:

```sh
time singularity ez-prep --output-dir '' -j $(($(nproc) / 4 + 1)) ./dataset
```

## Interpreting Results

Typical output will resemble:

```
real    0m20.379s
user    0m44.937s
sys     0m8.981s
```

* `real`: Actual elapsed time. Using more workers should reduce this time.
* `user`: CPU time used in user space. Dividing `user` by `real` approximates the number of CPU cores used.
* `sys`: CPU time used in kernel space (represents disk IO).

## Comparison

The following benchmarks were conducted on a random 8G file:

<table><thead><tr><th width="290">Tool</th><th width="178.33333333333331" data-type="number">clock time (sec)</th><th data-type="number">cpu time (sec)</th><th data-type="number">memory (KB)</th></tr></thead><tbody><tr><td>Singularity w/ inline prep</td><td>15.66</td><td>51.82</td><td>99</td></tr><tr><td>Singularity w/o inline prep</td><td>19.13</td><td>51.51</td><td>99</td></tr><tr><td>go-fil-dataprep</td><td>16.39</td><td>43.94</td><td>83</td></tr><tr><td>generate-car</td><td>42.6</td><td>56.08</td><td>44</td></tr><tr><td>go-car + stream-commp</td><td>70.21</td><td>139.01</td><td>42</td></tr></tbody></table>
