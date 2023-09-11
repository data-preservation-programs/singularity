# Sugarsync

{% code fullWidth="true" %}
```
命令名称:
   singularity storage create sugarsync - Sugarsync

使用方法:
   singularity storage create sugarsync [command options] [arguments...]

说明:
   --app-id
      Sugarsync 应用程序 ID。
      
      留空以使用 rclone 的默认值。

   --access-key-id
      Sugarsync 访问密钥 ID。
      
      留空以使用 rclone 的默认值。

   --private-access-key
      Sugarsync 私有访问密钥。
      
      留空以使用 rclone 的默认值。

   --hard-delete
      如果为 true，则永久删除文件；
      否则将文件放入已删除文件夹中。

   --refresh-token
      Sugarsync 刷新令牌。
      
      通常留空，将由 rclone 自动配置。

   --authorization
      Sugarsync 授权信息。
      
      通常留空，将由 rclone 自动配置。

   --authorization-expiry
      Sugarsync 授权信息到期时间。
      
      通常留空，将由 rclone 自动配置。

   --user
      Sugarsync 用户。
      
      通常留空，将由 rclone 自动配置。

   --root-id
      Sugarsync 根 ID。
      
      通常留空，将由 rclone 自动配置。

   --deleted-id
      Sugarsync 已删除文件夹 ID。
      
      通常留空，将由 rclone 自动配置。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

选项:
   --access-key-id value       Sugarsync 访问密钥 ID。[$ACCESS_KEY_ID]
   --app-id value              Sugarsync 应用程序 ID。[$APP_ID]
   --hard-delete               如果为 true，则永久删除文件（默认值：false）[$HARD_DELETE]
   --help, -h                  显示帮助
   --private-access-key value  Sugarsync 私有访问密钥。[$PRIVATE_ACCESS_KEY]

   高级选项

   --authorization value         Sugarsync 授权信息。[$AUTHORIZATION]
   --authorization-expiry value  Sugarsync 授权信息到期时间。[$AUTHORIZATION_EXPIRY]
   --deleted-id value            Sugarsync 已删除文件夹 ID。[$DELETED_ID]
   --encoding value              后端的编码方式。（默认值："Slash,Ctl,InvalidUtf8,Dot"）[$ENCODING]
   --refresh-token value         Sugarsync 刷新令牌。[$REFRESH_TOKEN]
   --root-id value               Sugarsync 根 ID。[$ROOT_ID]
   --user value                  Sugarsync 用户。[$USER]

   常规选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}