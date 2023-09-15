# seafile

{% code fullWidth="true" %}
```
NAME:
   singularity storage create seafile - seafile

USAGE:
   singularity storage create seafile [command options] [arguments...]

DESCRIPTION:
   --url
      URL of seafile host to connect to.

      Examples:
         | https://cloud.seafile.com/ | Connect to cloud.seafile.com.

   --user
      User name (usually email address).

   --pass
      Password.

   --2fa
      Two-factor authentication ('true' if the account has 2FA enabled).

   --library
      Name of the library.
      
      Leave blank to access all non-encrypted libraries.

   --library-key
      Library password (for encrypted libraries only).
      
      Leave blank if you pass it through the command line.

   --create-library
      Should rclone create a library if it doesn't exist.

   --auth-token
      Authentication token.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --2fa                Two-factor authentication ('true' if the account has 2FA enabled). (default: false) [$2FA]
   --auth-token value   Authentication token. [$AUTH_TOKEN]
   --help, -h           show help
   --library value      Name of the library. [$LIBRARY]
   --library-key value  Library password (for encrypted libraries only). [$LIBRARY_KEY]
   --pass value         Password. [$PASS]
   --url value          URL of seafile host to connect to. [$URL]
   --user value         User name (usually email address). [$USER]

   Advanced

   --create-library  Should rclone create a library if it doesn't exist. (default: false) [$CREATE_LIBRARY]
   --encoding value  The encoding for the backend. (default: "Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8") [$ENCODING]

   General

   --name value  Name of the storage (default: Auto generated)
   --path value  Path of the storage

   HTTP Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value)
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-user-agent value                        Set the user-agent to a specified string (default: rclone/v1.62.2-DEV)

```
{% endcode %}
