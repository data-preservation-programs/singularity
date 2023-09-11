# Create a deal schedule

Time to make some deals with storage providers. Start with running the deal pusher service

```
singularity run deal-pusher
```

## Send all deals at once

With smaller dataset, you could send all deals to your storage providers all at once. To achieve this, you can use below command

```sh
singularity deal schedule create <preparation> <provider_id>
```

However, if the dataset is large, it may be too much for storage providers to ingest that many deals before the deal proposal expiration, so you can create a schedule

## Send deals with schedule

With the same command, you can create your own schedule to control how fast and how often should the deals be made to storage providers
```sh
singularity deal schedule create -h
```
