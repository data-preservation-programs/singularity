# Mega

{% code fullWidth="true" %}
```
NAME:
   singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity storage create mega - Mega

USAGE:
   singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity singularity storage create mega [command options] <name> <path>

DESCRIPTION:
   --user
      User name.

   --pass
      Password.

   --debug
      Output more debug from Mega.
      
      If this flag is set (along with -vv) it will print further debugging
      information from the mega backend.

   --hard-delete
      Delete files permanently rather than putting them into the trash.
      
      Normally the mega backend will put all deletions into the trash rather
      than permanently deleting them.  If you specify this then rclone will
      permanently delete objects instead.

   --use-https
      Use HTTPS for transfers.
      
      MEGA uses plain text HTTP connections by default.
      Some ISPs throttle HTTP connections, this causes transfers to become very slow.
      Enabling this will force MEGA to use HTTPS for all transfers.
      HTTPS is normally not necesary since all data is already encrypted anyway.
      Enabling it will increase CPU usage and add network overhead.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h    show help
   --pass value  Password. [$PASS]
   --user value  User name. [$USER]

   Advanced

   --debug           Output more debug from Mega. (default: false) [$DEBUG]
   --encoding value  The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete     Delete files permanently rather than putting them into the trash. (default: false) [$HARD_DELETE]
   --use-https       Use HTTPS for transfers. (default: false) [$USE_HTTPS]

```
{% endcode %}
