# Start dataset worker

To start a dataset worker to prepare a dataset, run the following command

```sh
singularity run dataset-worker --exit-on-complete --exit-on-error
```

By default, it will spawn up a single worker thread that works on scanning, packing and dagnifing for the dataset. The process will exit upon completion or any encountered error. In Production, you would want it to keep running.

You can also configure some concurrency value with flag`--concurrency value`

After preparation completes, you can examine the prepared data using some of following commands

```sh
# List all data sources added
singularity datasource list

# Give an overview of scanning and packing result
singularity datasource status 1

# Check the CID of each file for the root folder
singularity datasource inspect dir 1

# Check all CAR files generated
singularity datasource inspect chunks

# Check all items that are prepared
singularity datasource inspect items
```

## Next step

[create-dag-for-the-data-source.md](create-dag-for-the-data-source.md "mention")

## Related Resources

[List all datasources](../cli-reference/datasource/list.md)

[Check datasource preparation status](../cli-reference/datasource/status.md)

[Check all items of a datasource](../cli-reference/datasource/inspect/items.md)

[Check all chunks of a datasource](../cli-reference/datasource/inspect/chunks.md)

[Check directory of a datasource](../cli-reference/datasource/inspect/dir.md)
