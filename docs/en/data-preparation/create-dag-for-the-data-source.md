# Create DAG for the data source

The DAG in this context contain all relevant folder information for a data source as well as how files are splitted across multiple chunks. If the CAR file for this DAG is sealed by storage providers, you will be able to lookup file using unixfs path with a single Root CID of the dataset.

To trigger the DAG generation process for a data source

```sh
# Assume there is a single datasourcesh
singularity datasource daggen 1
```

Now the job has been recorded in the database, you would need to rerun the dataset worker, or wait for it to pick up the job if the worker is already running

```sh
singularity run dataset-worker --exit-on-complete --exit-on-error
```

Once completed, you can check the relevant DAG

```
singularity datasource inspect dags 1
```

The CAR file for the DAG will be automatically included for deal making

## Next step

[distribute-car-files.md](../content-distribution/distribute-car-files.md "mention")

## Related resources

[Trigger a DAG generation](../cli-reference/datasource/daggen.md)

[Inspect DAG of a data source](../cli-reference/datasource/inspect/dags.md)
