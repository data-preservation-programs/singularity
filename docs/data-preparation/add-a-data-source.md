---
description: Connect to a data source that needs to be prepared
---

# Add a data source

## Add a local file system data source

The most command data source is the local file system. To add a folder as a data source to the dataset:

```sh
singularity datasource add local my_dataset /mnt/dataset/folder
```

## Add a public S3 data source

To demonstrate how you can add S3 data source to a dataset, let's use a public dataset called [Foldingathome COVID-19 Datasets](https://registry.opendata.aws/foldingathome-covid19/)

```
singularity datasource add s3 my_dataset fah-public-data-covid19-cryptic-pocketst 
```

## Next step

[start-dataset-worker.md](start-dataset-worker.md "mention")

## Related resources

[All data source types](../cli-reference/datasource/add/)
