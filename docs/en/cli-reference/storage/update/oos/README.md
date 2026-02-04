# Oracle Cloud Infrastructure Object Storage

{% code fullWidth="true" %}
```
NAME:
   singularity storage update oos - Oracle Cloud Infrastructure Object Storage

USAGE:
   singularity storage update oos command [command options]

COMMANDS:
   env_auth                 automatically pickup the credentials from runtime(env), first one to provide auth wins
   instance_principal_auth  use instance principals to authorize an instance to make API calls. 
                            each instance has its own identity, and authenticates using the certificates that are read from instance metadata. 
                            https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm
   no_auth                  no credentials needed, this is typically for reading public buckets
   resource_principal_auth  use resource principals to make API calls
   user_principal_auth      use an OCI user and an API key for authentication.
                            you’ll need to put in a config file your tenancy OCID, user OCID, region, the path, fingerprint to an API key.
                            https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm
   workload_identity_auth   use workload identity to grant OCI Container Engine for Kubernetes workloads policy-driven access to OCI resources using OCI Identity and Access Management (IAM).
                            https://docs.oracle.com/en-us/iaas/Content/ContEng/Tasks/contenggrantingworkloadaccesstoresources.htm
   help, h                  Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```
{% endcode %}
