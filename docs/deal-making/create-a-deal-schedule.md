# Create a deal schedule

Time to make some deals with storage providers

## Send all deals at once

With smaller dataset, you could send all deals to your storage providers all at once. To achieve this, you can use below command

```sh
singularity deal schedule create dataset_name provider_id
```

However, if the dataset is large, it may be too much for storage providers to ingest that many deals before the deal proposal expiration, so you can create a schedule

## Send deals with schedule

With the same command, you can create your own schedule to control how fast and how often should the deals be made to storage providers

```
--schedule-deal-number value, --number value     Max deal number per triggered schedule, i.e. 30 (default: Unlimited)
--schedule-deal-size value, --size value         Max deal sizes per triggered schedule, i.e. 500GB (default: Unlimited)
--schedule-interval value, --every value         Cron schedule to send out batch deals (default: disabled)
--total-deal-number value, --total-number value  Max total deal number for this request, i.e. 1000 (default: Unlimited)
```
