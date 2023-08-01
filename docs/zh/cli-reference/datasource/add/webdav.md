# WebDAV

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add webdav - WebDAV

USAGE:
   singularity datasource add webdav [command options] <dataset_name> <source_path>

DESCRIPTION:
   --webdav-bearer-token
      使用 Bearer token 而不是用户名/密码（例如 Macaroon）。

   --webdav-bearer-token-command
      运行命令以获取 Bearer token。

   --webdav-encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。
      
      默认编码方式为 Slash, LtGt, DoubleQuote, Colon, Question, Asterisk, Pipe, Hash, Percent, BackSlash, Del, Ctl, LeftSpace, LeftTilde, RightSpace, RightPeriod, InvalidUtf8（用于 sharepoint-ntlm 或 identity）。

   --webdav-headers
      设置所有事务的 HTTP 头。
      
      可以使用此选项来设置所有事务的额外 HTTP 头。
      
      输入格式为逗号分隔的键值对列表。 可以使用标准的[CSV编码](https://godoc.org/encoding/csv)。
      
      例如，要设置一个 Cookie，请使用'Cookie,name=value' 或者 '"Cookie","name=value"'。
      
      您可以设置多个头，例如 '"Cookie","name=value","Authorization","xxx"'。

   --webdav-pass
      密码。

   --webdav-url
      要连接的 http 主机的 URL。
      
      例如 https://example.com。

   --webdav-user
      用户名。
      
      如果使用 NTLM 身份验证，用户名应采用 'Domain\User' 格式。

   --webdav-vendor
      您正在使用的 WebDAV 站点/服务/软件的名称。

      示例：
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | Sharepoint Online，由 Microsoft 帐户进行身份验证
         | sharepoint-ntlm | 带有 NTLM 身份验证的 Sharepoint，通常是自托管或本地部署的
         | other           | 其他站点/服务或软件


OPTIONS:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作] 在将数据集导出到 CAR 文件后删除数据集的文件。  (default: false)
   --rescan-interval value  当距离上次成功扫描已过去此间隔时，自动重新扫描源目录（default: disabled）
   --scanning-state value   设置初始扫描状态（default: ready）

   WebDAV 选项

   --webdav-bearer-token value          使用 Bearer token 而不是用户名/密码（例如 Macaroon）。[$WEBDAV_BEARER_TOKEN]
   --webdav-bearer-token-command value  运行命令以获取 Bearer token。[$WEBDAV_BEARER_TOKEN_COMMAND]
   --webdav-encoding value              后端的编码方式。[$WEBDAV_ENCODING]
   --webdav-headers value               设置所有事务的 HTTP 头。[$WEBDAV_HEADERS]
   --webdav-pass value                  密码。[$WEBDAV_PASS]
   --webdav-url value                   要连接的 http 主机的 URL。[$WEBDAV_URL]
   --webdav-user value                  用户名。[$WEBDAV_USER]
   --webdav-vendor value                您正在使用的 WebDAV 站点/服务/软件的名称。[$WEBDAV_VENDOR]

```
{% endcode %}