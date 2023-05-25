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

[local-file-system.md](../cli-reference/data-source/add-data-source/local-file-system.md "mention")

[aws-other-s3.md](../cli-reference/data-source/add-data-source/aws-other-s3.md "mention")

[add-data-source](../cli-reference/data-source/add-data-source/ "mention")
