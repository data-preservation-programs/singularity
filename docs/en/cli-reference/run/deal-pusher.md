# Start a deal pusher that monitors deal schedules and pushes deals to storage providers

{% code fullWidth="true" %}
```
NAME:
   singularity run deal-pusher - Start a deal pusher that monitors deal schedules and pushes deals to storage providers

USAGE:
   singularity run deal-pusher [command options]

OPTIONS:
   --deal-attempts value, -d value           Number of times to attempt a deal before giving up (default: 3)
   --max-replication-factor value, -M value  Max number of replicas for each individual PieceCID across all clients and providers (default: Unlimited)
   --help, -h                                show help
```
{% endcode %}
