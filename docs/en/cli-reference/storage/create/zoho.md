# Zoho

{% code fullWidth="true" %}
```
NAME:
   singularity storage create zoho - Zoho

USAGE:
   singularity storage create zoho [command options] <name> <path>

DESCRIPTION:
   --client_id
      OAuth Client Id.
      
      Leave blank normally.

   --client_secret
      OAuth Client Secret.
      
      Leave blank normally.

   --token
      OAuth Access Token as a JSON blob.

   --auth_url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --token_url
      Token server url.
      
      Leave blank to use the provider defaults.

   --region
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

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --client_id value      OAuth Client Id. [$CLIENT_ID]
   --client_secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help
   --region value         Zoho region to connect to. [$REGION]

   Advanced

   --auth_url value   Auth server URL. [$AUTH_URL]
   --encoding value   The encoding for the backend. (default: "Del,Ctl,InvalidUtf8") [$ENCODING]
   --token value      OAuth Access Token as a JSON blob. [$TOKEN]
   --token_url value  Token server url. [$TOKEN_URL]

```
{% endcode %}
