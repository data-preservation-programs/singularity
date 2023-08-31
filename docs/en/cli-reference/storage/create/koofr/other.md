# Any other Koofr API compatible storage service

{% code fullWidth="true" %}
```
NAME:
   singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity storage create koofr other - Any other Koofr API compatible storage service

USAGE:
   singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity storage create koofr other [command options] <name> <path>

DESCRIPTION:
   --endpoint
      The Koofr API endpoint to use.

   --mountid
      Mount ID of the mount to use.
      
      If omitted, the primary mount is used.

   --setmtime
      Does the backend support setting modification time.
      
      Set this to false if you use a mount ID that points to a Dropbox or Amazon Drive backend.

   --user
      Your user name.

   --password
      Your password for rclone (generate one at your service's settings page).

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --endpoint value  The Koofr API endpoint to use. [$ENDPOINT]
   --help, -h        show help
   --password value  Your password for rclone (generate one at your service's settings page). [$PASSWORD]
   --user value      Your user name. [$USER]

   Advanced

   --encoding value  The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --mountid value   Mount ID of the mount to use. [$MOUNTID]
   --setmtime        Does the backend support setting modification time. (default: true) [$SETMTIME]

```
{% endcode %}
