# Pcloud

{% code fullWidth="true" %}
```
名称：
   singularity storage create pcloud - Pcloud

用法：
   singularity storage create pcloud [命令选项] [参数...]

描述：
   --client-id
      OAuth客户端ID。
      
      通常保持空白。

   --client-secret
      OAuth客户端密钥。
      
      通常保持空白。

   --token
      OAuth访问令牌，以JSON格式。

   --auth-url
      认证服务器URL。
      
      保留空白以使用提供程序的默认值。

   --token-url
      令牌服务器URL。
      
      保留空白以使用提供程序的默认值。

   --encoding
      后端的编码。
      
      有关详细信息，请参阅[概览中的编码部分](/overview/#encoding)。

   --root-folder-id
      填写rclone所使用的非根文件夹作为起始点。

   --hostname
      连接的主机名。
      
      这通常在rclone最初进行oauth连接时设置，
      但如果您使用rclone授权进行远程配置，则需要手动设置它。
      

      示例：
         | api.pcloud.com  | 原始/美国地区
         | eapi.pcloud.com | 欧洲地区

   --username
      您的pcloud用户名。
            
      仅当您想要使用清理命令时才需要。由于pcloud API中所需的API不支持OAuth身份验证，
      所以我们必须依赖用户密码验证。

   --password
      您的pcloud密码。


选项：
   --client-id value      OAuth客户端ID。[$CLIENT_ID]
   --client-secret value  OAuth客户端密钥。[$CLIENT_SECRET]
   --help, -h             显示帮助信息

   高级选项

   --auth-url value        认证服务器URL。[$AUTH_URL]
   --encoding value        后端的编码。（默认值："Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"）[$ENCODING]
   --hostname value        连接的主机名。（默认值："api.pcloud.com"）[$HOSTNAME]
   --password value        您的pcloud密码。[$PASSWORD]
   --root-folder-id value  填写rclone所使用的非根文件夹作为起始点。（默认值："d0"）[$ROOT_FOLDER_ID]
   --token value           OAuth访问令牌，以JSON格式。[$TOKEN]
   --token-url value       令牌服务器URL。[$TOKEN_URL]
   --username value        您的pcloud用户名。[$USERNAME]

   常规选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}