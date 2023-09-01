# WebDAV

{% code fullWidth="true" %}
```
NAME:
   singularity storage update webdav - WebDAV

USAGE:
   singularity storage update webdav [command options] <name>

DESCRIPTION:
   --url
      URL of http host to connect to.
      
      E.g. https://example.com.

   --vendor
      Name of the WebDAV site/service/software you are using.

      Examples:
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | Sharepoint Online, authenticated by Microsoft account
         | sharepoint-ntlm | Sharepoint with NTLM authentication, usually self-hosted or on-premises
         | other           | Other site/service or software

   --user
      User name.
      
      In case NTLM authentication is used, the username should be in the format 'Domain\User'.

   --pass
      Password.

   --bearer-token
      Bearer token instead of user/pass (e.g. a Macaroon).

   --bearer-token-command
      Command to run to get a bearer token.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.
      
      Default encoding is Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Hash,Percent,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8 for sharepoint-ntlm or identity otherwise.

   --headers
      Set HTTP headers for all transactions.
      
      Use this to set additional HTTP headers for all transactions
      
      The input format is comma separated list of key,value pairs.  Standard
      [CSV encoding](https://godoc.org/encoding/csv) may be used.
      
      For example, to set a Cookie use 'Cookie,name=value', or '"Cookie","name=value"'.
      
      You can set multiple headers, e.g. '"Cookie","name=value","Authorization","xxx"'.
      


OPTIONS:
   --bearer-token value  Bearer token instead of user/pass (e.g. a Macaroon). [$BEARER_TOKEN]
   --help, -h            show help
   --pass value          Password. [$PASS]
   --url value           URL of http host to connect to. [$URL]
   --user value          User name. [$USER]
   --vendor value        Name of the WebDAV site/service/software you are using. [$VENDOR]

   Advanced

   --bearer-token-command value  Command to run to get a bearer token. [$BEARER_TOKEN_COMMAND]
   --encoding value              The encoding for the backend. [$ENCODING]
   --headers value               Set HTTP headers for all transactions. [$HEADERS]

```
{% endcode %}