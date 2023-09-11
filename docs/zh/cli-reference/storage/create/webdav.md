# WebDAV

{% code fullWidth="true" %}
```
名称：
   singularity storage create webdav - WebDAV

用法：
   singularity storage create webdav [命令选项] [参数...]

描述：
   --url
      要连接的 HTTP 主机的 URL。
      
      例如：https://example.com。

   --vendor
      您正在使用的 WebDAV 站点/服务/软件的名称。

      示例：
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | SharePoint Online，使用 Microsoft 帐户进行身份验证
         | sharepoint-ntlm | 使用 NTLM 身份验证的 Sharepoint，通常是自托管或本地部署
         | other           | 其他站点/服务或软件

   --user
      用户名。
      
      如果使用 NTLM 身份验证，则用户名应采用 '域\用户' 的格式。

   --pass
      密码。

   --bearer-token
      使用承载令牌而不是用户名/密码（例如 Macaroon）。

   --bearer-token-command
      用于获取承载令牌的命令。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。
      
      如果是 sharepoint-ntlm 或 identity，则默认编码为 Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Hash,Percent,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8。

   --headers
      为所有交易设置 HTTP 标头。
      
      可以使用此选项为所有交易设置附加的 HTTP 标头。
      
      输入格式为逗号分隔的键值对列表。可以使用标准的[CSV 编码](https://godoc.org/encoding/csv)。
      
      例如，要设置 Cookie，请使用 'Cookie,name=value' 或 '"Cookie","name=value"'。
      
      您可以设置多个标头，例如：'"Cookie","name=value","Authorization","xxx"'。
      


选项：
   --bearer-token value  使用承载令牌而不是用户名/密码（例如 Macaroon）。[$BEARER_TOKEN]
   --help, -h            显示帮助信息
   --pass value          密码。[$PASS]
   --url value           要连接的 HTTP 主机的 URL。[$URL]
   --user value          用户名。[$USER]
   --vendor value        您正在使用的 WebDAV 站点/服务/软件的名称。[$VENDOR]

   高级选项

   --bearer-token-command value  用于获取承载令牌的命令。[$BEARER_TOKEN_COMMAND]
   --encoding value              后端的编码方式。[$ENCODING]
   --headers value               为所有交易设置 HTTP 标头。[$HEADERS]

   一般选项

   --name value  存储的名称（默认为自动生成）
   --path value  存储的路径

```
{% endcode %}