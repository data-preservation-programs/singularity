# WebDAV

{% code fullWidth="true" %}
```
名称:
   singularity存储更新webdav - WebDAV

用法:
   singularity存储更新webdav [命令选项] <名称|ID>

描述:
   --url
      要连接的http主机的URL。
      
      例如：https://example.com。

   --vendor
      您使用的WebDAV站点/服务/软件的名称。

      示例：
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | Sharepoint Online，使用Microsoft帐户进行身份验证
         | sharepoint-ntlm | 使用NTLM身份验证的Sharepoint，通常是自托管或本地部署
         | other           | 其他站点/服务或软件

   --user
      用户名。
      
      如果使用NTLM身份验证，则用户名应以'Domain\User'格式。

   --pass
      密码。

   --bearer-token
      使用令牌（如Macaroon）而不是用户名/密码。

   --bearer-token-command
      获取令牌的命令。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概览中的编码部分](/overview/#encoding)。
      
      对于sharepoint-ntlm，默认的编码方式为Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Hash,Percent,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8；对于其他情况，则是identity。

   --headers
      为所有交易设置HTTP标头。
      
      可以使用此选项为所有交易设置其他HTTP标头。
      
      输入格式是由逗号分隔的键值对列表。可以使用标准的[CSV编码](https://godoc.org/encoding/csv)。
      
      例如，要设置Cookie，请使用'Cookie,name=value'或'"Cookie","name=value"'。
      
      您可以设置多个标头，例如'"Cookie","name=value","Authorization","xxx"'。
      

选项:
   --bearer-token value  使用令牌（如Macaroon）而不是用户名/密码。[$BEARER_TOKEN]
   --help, -h            显示帮助信息
   --pass value          密码。[$PASS]
   --url value           要连接的http主机的URL。[$URL]
   --user value          用户名。[$USER]
   --vendor value        您使用的WebDAV站点/服务/软件的名称。[$VENDOR]

   高级选项

   --bearer-token-command value  获取令牌的命令。[$BEARER_TOKEN_COMMAND]
   --encoding value              后端的编码方式。[$ENCODING]
   --headers value               为所有交易设置HTTP标头。[$HEADERS]

```
{% endcode %}