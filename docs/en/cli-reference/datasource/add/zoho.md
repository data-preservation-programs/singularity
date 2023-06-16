# Zoho

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add zoho - Zoho

USAGE:
   singularity datasource add zoho [command options] <dataset_name> <source_path>

DESCRIPTION:
   --zoho-token
      OAuth Access Token as a JSON blob.

   --zoho-auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --zoho-token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --zoho-region
      Zoho region to connect to.
      
      You'll have to use the region your organization is registered in. If
      not sure use the same top level domain as you connect to in your
      browser.

      Examples:
         | com    | United states / Global
         | eu     | Europe
         | in     | India
         | jp     | Japan
         | com.cn | China
         | com.au | Australia

   --zoho-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --zoho-client-id
      OAuth Client Id.
      
      Leave blank normally.

   --zoho-client-secret
      OAuth Client Secret.
      
      Leave blank normally.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for zoho

   --zoho-auth-url value       Auth server URL. [$ZOHO_AUTH_URL]
   --zoho-client-id value      OAuth Client Id. [$ZOHO_CLIENT_ID]
   --zoho-client-secret value  OAuth Client Secret. [$ZOHO_CLIENT_SECRET]
   --zoho-encoding value       The encoding for the backend. (default: "Del,Ctl,InvalidUtf8") [$ZOHO_ENCODING]
   --zoho-region value         Zoho region to connect to. [$ZOHO_REGION]
   --zoho-token value          OAuth Access Token as a JSON blob. [$ZOHO_TOKEN]
   --zoho-token-url value      Token server url. [$ZOHO_TOKEN_URL]

```
{% endcode %}
