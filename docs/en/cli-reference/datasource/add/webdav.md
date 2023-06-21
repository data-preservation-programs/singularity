# WebDAV

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add webdav - WebDAV

USAGE:
   singularity datasource add webdav [command options] <dataset_name> <source_path>

DESCRIPTION:
   --webdav-bearer-token
      Bearer token instead of user/pass (e.g. a Macaroon).

   --webdav-bearer-token-command
      Command to run to get a bearer token.

   --webdav-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.
      
      Default encoding is Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Hash,Percent,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8 for sharepoint-ntlm or identity otherwise.

   --webdav-headers
      Set HTTP headers for all transactions.
      
      Use this to set additional HTTP headers for all transactions
      
      The input format is comma separated list of key,value pairs.  Standard
      [CSV encoding](https://godoc.org/encoding/csv) may be used.
      
      For example, to set a Cookie use 'Cookie,name=value', or '"Cookie","name=value"'.
      
      You can set multiple headers, e.g. '"Cookie","name=value","Authorization","xxx"'.
      

   --webdav-pass
      Password.

   --webdav-url
      URL of http host to connect to.
      
      E.g. https://example.com.

   --webdav-user
      User name.
      
      In case NTLM authentication is used, the username should be in the format 'Domain\User'.

   --webdav-vendor
      Name of the WebDAV site/service/software you are using.

      Examples:
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | Sharepoint Online, authenticated by Microsoft account
         | sharepoint-ntlm | Sharepoint with NTLM authentication, usually self-hosted or on-premises
         | other           | Other site/service or software


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for webdav

   --webdav-bearer-token value          Bearer token instead of user/pass (e.g. a Macaroon). [$WEBDAV_BEARER_TOKEN]
   --webdav-bearer-token-command value  Command to run to get a bearer token. [$WEBDAV_BEARER_TOKEN_COMMAND]
   --webdav-encoding value              The encoding for the backend. [$WEBDAV_ENCODING]
   --webdav-headers value               Set HTTP headers for all transactions. [$WEBDAV_HEADERS]
   --webdav-pass value                  Password. [$WEBDAV_PASS]
   --webdav-url value                   URL of http host to connect to. [$WEBDAV_URL]
   --webdav-user value                  User name. [$WEBDAV_USER]
   --webdav-vendor value                Name of the WebDAV site/service/software you are using. [$WEBDAV_VENDOR]

```
{% endcode %}
