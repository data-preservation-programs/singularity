# WebDAV

{% code fullWidth="true" %}
```
名称:
   singularity数据源添加webdav——WebDAV

用法:
   singularity datasource add webdav [命令选项] <数据集名称> <源路径>

描述:
   --webdav-bearer-token-command
      运行的获取访问令牌的命令。

   --webdav-encoding
      后端的编码方式。

      有关更多信息，请参阅[总览中的编码部分](/overview/#encoding)。

      当使用sharepoint-ntlm或身份验证时，默认编码方式为Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Hash,Percent,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8。

   --webdav-headers
      为所有交易设置HTTP头。

      使用此选项设置所有事务的其他HTTP头。

      输入格式为以逗号分隔的键值对列表。可以使用标准的[CSV编码](https://godoc.org/encoding/csv)。

      例如，要设置一个Cookie，请使用'Cookie,name=value'，或'"Cookie","name=value"'。

      可以设置多个头，例如'"Cookie","name=value","Authorization","xxx"'。

   --webdav-url
      要连接的http主机的URL。

      例如：https://example.com。

   --webdav-vendor
      您使用的WebDAV网站/服务/软件的名称。

      示例：
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | 由Microsoft帐户身份验证的Sharepoint Online
         | sharepoint-ntlm | 具有NTLM身份验证的Sharepoint，通常是自托管或在企业内部
         | other           | 其他站点/服务或软件

   --webdav-user
      用户名。

      如使用NTLM身份验证，则用户名应采用"域/用户"格式。

   --webdav-pass
      密码。

   --webdav-bearer-token
      用户名和密码之外的访问令牌（例如Macaroon）。

选项:
   --help，-h  显示帮助

   数据准备选项

   --delete-after-export    【危险操作】在将数据集导出到CAR文件后删除数据集文件。 （默认值：false）
   --rescan-interval value  在上次成功扫描后的这个时间段之后自动重新扫描源目录（默认值：禁用）

   webdav选项

   --webdav-bearer-token value          用户名和密码之外的访问令牌（例如Macaroon）。[$WEBDAV_BEARER_TOKEN]
   --webdav-bearer-token-command value  运行的获取访问令牌的命令。[$WEBDAV_BEARER_TOKEN_COMMAND]
   --webdav-encoding value              后端的编码方式。[$WEBDAV_ENCODING]
   --webdav-headers value               为所有交易设置HTTP头。[$WEBDAV_HEADERS]
   --webdav-pass value                  密码。[$WEBDAV_PASS]
   --webdav-url value                   要连接的http主机的URL。[$WEBDAV_URL]
   --webdav-user value                  用户名。[$WEBDAV_USER]
   --webdav-vendor value                您使用的WebDAV网站/服务/软件的名称。[$WEBDAV_VENDOR]

```
{% endcode %}