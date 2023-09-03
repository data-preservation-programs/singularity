# Koofr, https://app.koofr.net/

{% code fullWidth="true" %}
```
NAME:
   singularity storage update koofr koofr - Koofr, https://app.koofr.net/

USAGE:
   singularity storage update koofr koofr [command options] <name|id>

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
      Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password).

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h        show help
   --password value  Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password). [$PASSWORD]
   --user value      Your user name. [$USER]

   Advanced

   --encoding value  The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --mountid value   Mount ID of the mount to use. [$MOUNTID]
   --setmtime        Does the backend support setting modification time. (default: true) [$SETMTIME]

```
{% endcode %}
