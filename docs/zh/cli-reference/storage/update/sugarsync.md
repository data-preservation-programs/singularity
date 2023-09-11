# Sugarsync

{% code fullWidth="true" %}
```
命令名称:
   singularity storage update sugarsync - Sugarsync

用法:
   singularity storage update sugarsync [命令选项] <名称|Id>

描述：
   --app-id
      Sugarsync 应用程序 ID。
      
      留空以使用 rclone 的应用程序 ID。

   --access-key-id
      Sugarsync 访问密钥 ID。
      
      留空以使用 rclone 的访问密钥 ID。

   --private-access-key
      Sugarsync 私有访问密钥。
      
      留空以使用 rclone 的私有访问密钥。

   --hard-delete
      如果设置为 true，则永久删除文件；否则将文件移到已删除文件夹中。

   --refresh-token
      Sugarsync 刷新令牌。
      
      通常留空，rclone 将自动配置它。

   --authorization
      Sugarsync 授权信息。
      
      通常留空，rclone 将自动配置它。

   --authorization-expiry
      Sugarsync 授权信息的过期时间。
      
      通常留空，rclone 将自动配置它。

   --user
      Sugarsync 用户。
      
      通常留空，rclone 将自动配置它。

   --root-id
      Sugarsync 根目录 ID。
      
      通常留空，rclone 将自动配置它。

   --deleted-id
      Sugarsync 已删除文件夹 ID。
      
      通常留空，rclone 将自动配置它。

   --encoding
      后端的编码方式。
      
      更多信息请参阅[概述中的编码部分](/overview/#encoding)。


选项:
   --access-key-id value       Sugarsync 访问密钥 ID。[$ACCESS_KEY_ID]
   --app-id value              Sugarsync 应用程序 ID。[$APP_ID]
   --hard-delete               如果设置为 true，则永久删除文件（默认值：false）[$HARD_DELETE]
   --help, -h                  显示帮助信息
   --private-access-key value  Sugarsync 私有访问密钥。[$PRIVATE_ACCESS_KEY]

   高级选项:

   --authorization value         Sugarsync 授权信息。[$AUTHORIZATION]
   --authorization-expiry value  Sugarsync 授权信息的过期时间。[$AUTHORIZATION_EXPIRY]
   --deleted-id value            Sugarsync 已删除文件夹 ID。[$DELETED_ID]
   --encoding value              后端的编码方式（默认值："Slash,Ctl,InvalidUtf8,Dot"）[$ENCODING]
   --refresh-token value         Sugarsync 刷新令牌。[$REFRESH_TOKEN]
   --root-id value               Sugarsync 根目录 ID。[$ROOT_ID]
   --user value                  Sugarsync 用户。[$USER]

```
{% endcode %}