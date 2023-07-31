# Pcloud

{% code fullWidth="true" %}
```
命令名称：
   singularity datasource add pcloud - Pcloud

用法：
   singularity datasource add pcloud [命令选项] <数据集名称> <源路径>

描述：
   --pcloud-auth-url
      授权服务器 URL。
      
      留空以使用提供程序默认值。

   --pcloud-client-id
      OAuth 客户端 ID。
      
      通常留空。

   --pcloud-client-secret
      OAuth 客户端秘钥。
      
      通常留空。

   --pcloud-encoding
      后端编码。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --pcloud-hostname
      要连接的主机名。
      
      当 rclone 最初执行 OAuth 连接时，通常会自动设置，
      但如果您正在使用带有 rclone 授权的远程配置，则需手动设置。
      

      示例：
         | api.pcloud.com  | 原始/美国地区
         | eapi.pcloud.com | 欧洲地区

   --pcloud-password
      您的 pcloud 密码。

   --pcloud-root-folder-id
      填写以使用非根文件夹作为其起始点。

   --pcloud-token
      OAuth 访问令牌（JSON 格式）。

   --pcloud-token-url
      令牌服务器 URL。
      
      留空以使用提供程序默认值。

   --pcloud-username
      您的 pcloud 用户名。
            
      仅在想要使用清理命令时才需要。由于 pcloud API 的一个错误，
      必需的 API 不支持 OAuth 验证，因此我们必须依靠用户密码验证。

选项：
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险] 导出 CAR 文件后删除数据集的文件。（默认值：false）
   --rescan-interval value  源目录最后一次成功扫描后，自动重新扫描的间隔时间。（默认值：禁用）
   --scanning-state value   设置初始扫描状态（默认值：ready）

   pcloud 选项

   --pcloud-auth-url value        授权服务器 URL。[$PCLOUD_AUTH_URL]
   --pcloud-client-id value       OAuth 客户端 ID。[$PCLOUD_CLIENT_ID]
   --pcloud-client-secret value   OAuth 客户端秘钥。[$PCLOUD_CLIENT_SECRET]
   --pcloud-encoding value        后端编码。（默认值："Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"）[$PCLOUD_ENCODING]
   --pcloud-hostname value        要连接的主机名。（默认值："api.pcloud.com"）[$PCLOUD_HOSTNAME]
   --pcloud-password value        您的 pcloud 密码。[$PCLOUD_PASSWORD]
   --pcloud-root-folder-id value  填写以使用非根文件夹作为其起始点。（默认值："d0"）[$PCLOUD_ROOT_FOLDER_ID]
   --pcloud-token value           OAuth 访问令牌（JSON 格式）。[$PCLOUD_TOKEN]
   --pcloud-token-url value       令牌服务器 URL。[$PCLOUD_TOKEN_URL]
   --pcloud-username value        您的 pcloud 用户名。[$PCLOUD_USERNAME]

```
{% endcode %}