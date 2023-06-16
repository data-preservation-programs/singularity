# Benchmark

The EZ preparation command offers a simple way to perform benchmark.

## Prepare test data

First, prepare some data for benchmarking. We are utilizing sparse file to avoid taking into account disk IO time. As of now, we don't do CID deduplication so singularity treat them as if those are random bytes.

```sh
mkdir dataset
truncate -s 1024G dataset/1T.bin
```

If you want to include disk IO time as part of the benchmark, you can use your own favorate way to create a random file, i.e.

```
dd if=/dev/urandom of=dataset/8G.bin bs=1M count=8192
```

## Run ez-prep

EZ prep command is a simple command that runs a few internal commands to prepare a local folder with very few customizable settings.

#### Benchmark with inline preparation&#x20;

Inline preparation eliminates the need to export CAR files and save the necessary metadata in the database

```sh
time singularity ez-prep --output-dir '' ./dataset
```

#### Benchmark with in-memory database

To further reduce disk IO, you can also choose to use in-memory database

```sh
time singularity ez-prep --output-dir '' --database-file '' ./dataset
```

#### Bechmark with multiple workers

To utilize all CPU cores, you can set a concurrency flag for the benchmark. Note that each worker consumes about 4 CPU cores so you would want to set it correctly

```sh
time singularity ez-prep --output-dir '' -j $(($(nproc) / 4 + 1)) ./dataset
```

## Interprete the result

You will see something like below

```
real    0m20.379s
user    0m44.937s
sys     0m8.981s
```

`real` means the actual clock time. Using more worker concurrency will likely reduce this number.

`user` means the CPU time spent on user space. If you divide `user` with `real`, it's roughly how many CPU cores the program has used. Using more concurrnecy will likely not have a significant impact on this number since the work needs to be done does not change.

`sys` means the CPU time spent on kernel space which is the disk IO

## Comparison

The below test is performed on a random 8G file

<table><thead><tr><th width="290">Tool</th><th width="178.33333333333331" data-type="number">clock time (sec)</th><th data-type="number">cpu time (sec)</th><th data-type="number">memory (KB)</th></tr></thead><tbody><tr><td>Singularity w/ inline prep</td><td>15.66</td><td>51.82</td><td>99</td></tr><tr><td>Singularity w/o inline prep</td><td>19.13</td><td>51.51</td><td>99</td></tr><tr><td>go-fil-dataprep</td><td>16.39</td><td>43.94</td><td>83</td></tr><tr><td>generate-car</td><td>42.6</td><td>56.08</td><td>44</td></tr><tr><td>go-car + stream-commp</td><td>70.21</td><td>139.01</td><td>42</td></tr></tbody></table>
