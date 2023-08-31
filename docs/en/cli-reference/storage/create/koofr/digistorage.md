# Digi Storage, https://storage.rcs-rds.ro/

{% code fullWidth="true" %}
```
NAME:
   singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity storage create koofr digistorage - Digi Storage, https://storage.rcs-rds.ro/

USAGE:
   singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity storage create koofr digistorage [command options] <name> <path>

DESCRIPTION:
   --mountid
      Mount ID of the mount to use.
      
      If omitted, the primary mount is used.

   --setmtime
      Does the backend support setting modification time.
      
      Set this to false if you use a mount ID that points to a Dropbox or Amazon Drive backend.

   --user
      Your user name.

   --password
      Your password for rclone (generate one at https://storage.rcs-rds.ro/app/admin/preferences/password).

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h        show help
   --password value  Your password for rclone (generate one at https://storage.rcs-rds.ro/app/admin/preferences/password). [$PASSWORD]
   --user value      Your user name. [$USER]

   Advanced

   --encoding value  The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --mountid value   Mount ID of the mount to use. [$MOUNTID]
   --setmtime        Does the backend support setting modification time. (default: true) [$SETMTIME]

```
{% endcode %}
