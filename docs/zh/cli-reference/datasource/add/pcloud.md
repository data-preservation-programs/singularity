# Pcloud

{% code fullWidth="true" %}
```
名称：
   singularity datasource add pcloud - Pcloud
   
用法：
   singularity datasource add pcloud [命令选项] <数据集名称> <源路径>
   
描述：
   --pcloud-token
      OAuth访问令牌的JSON格式。
   
   --pcloud-auth-url
      Auth服务器的URL。
      
      保留为空以使用提供程序默认设置。
   
   --pcloud-encoding
      后端编码。
      
      更多信息请参见[概述中的编码部分](/overview/#encoding)。

   --pcloud-hostname
      连接的主机名。
      
      通常在rclone初始进行oauth连接时设置，但是如果使用rclone授权的远程配置，您需要手动设置。
      

      例如：
        | api.pcloud.com  | 原版/美国地区
        | eapi.pcloud.com | 欧盟地区

   --pcloud-username
      您的pcloud用户名。
      
      只有在想要使用清理（cleanup）命令时才需要。由于pcloud API中的一个错误，所需的API不支持OAuth认证，因此我们必须依靠用户密码身份验证。

   --pcloud-password
      您的pcloud密码。

   --pcloud-client-id
      OAuth客户端ID。
      
      通常保留为空。

   --pcloud-client-secret
      OAuth客户端密码。
      
      通常保留为空。

   --pcloud-token-url
      Token服务器的URL。
      
      保留为空以使用提供程序默认设置。

   --pcloud-root-folder-id
      用rclone使用非根目录作为其起始点时填写。
      
      
选项：
   --help，-h  显示帮助
   
   数据准备选项
   
   --delete-after-export    [危险]导出CAR文件后，删除数据集的文件。  (默认值：false)
   --rescan-interval value  自上次成功扫描之后，当此间隔经过时自动重新扫描源目录（默认值：已禁用）

   pcloud选项

   --pcloud-auth-url value        Auth服务器的URL。[$PCLOUD_AUTH_URL]
   --pcloud-client-id value       OAuth客户端ID。[$PCLOUD_CLIENT_ID]
   --pcloud-client-secret value   OAuth客户端密码。[$PCLOUD_CLIENT_SECRET]
   --pcloud-encoding value        后端编码。默认值："Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot" [$PCLOUD_ENCODING]
   --pcloud-hostname value        主机名以连接到。默认值："api.pcloud.com" [$PCLOUD_HOSTNAME]
   --pcloud-password value        您的pcloud密码。[$PCLOUD_PASSWORD]
   --pcloud-root-folder-id value  用rclone使用非根目录作为其起始点时填写。默认值："d0" [$PCLOUD_ROOT_FOLDER_ID]
   --pcloud-token value           OAuth访问令牌的JSON格式。[$PCLOUD_TOKEN]
   --pcloud-token-url value       Token服务器的URL。[$PCLOUD_TOKEN_URL]
   --pcloud-username value        您的pcloud用户名。[$PCLOUD_USERNAME]


```
{% endcode %}