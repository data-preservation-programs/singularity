# seafile

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add seafile - seafile

USAGE:
   singularity datasource add seafile [command options] <dataset_name> <source_path>

DESCRIPTION:
   --seafile-encoding
      The encoding for the backend.

      See the [encoding section in the overview](/overview/#encoding) for more info.

   --seafile-user
      User name (usually email address).

   --seafile-library
      Name of the library.

      Leave blank to access all non-encrypted libraries.

   --seafile-library-key
      Library password (for encrypted libraries only).

      Leave blank if you pass it through the command line.

   --seafile-create-library
      Should rclone create a library if it doesn't exist.

   --seafile-url
      URL of seafile host to connect to.

      Examples:
         | https://cloud.seafile.com/ | Connect to cloud.seafile.com.

   --seafile-pass
      Password.

   --seafile-2fa
      Two-factor authentication ('true' if the account has 2FA enabled).

   --seafile-auth-token
      Authentication token.


OPTIONS:
   --help, -h                   show help
   --seafile-2fa value          Two-factor authentication ('true' if the account has 2FA enabled). (default: "false") [$SEAFILE_2FA]
   --seafile-library value      Name of the library. [$SEAFILE_LIBRARY]
   --seafile-library-key value  Library password (for encrypted libraries only). [$SEAFILE_LIBRARY_KEY]
   --seafile-pass value         Password. [$SEAFILE_PASS]
   --seafile-url value          URL of seafile host to connect to. [$SEAFILE_URL]
   --seafile-user value         User name (usually email address). [$SEAFILE_USER]

   Advanced Options

   --seafile-create-library value  Should rclone create a library if it doesn't exist. (default: "false") [$SEAFILE_CREATE_LIBRARY]
   --seafile-encoding value        The encoding for the backend. (default: "Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8") [$SEAFILE_ENCODING]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)

```
{% endcode %}
