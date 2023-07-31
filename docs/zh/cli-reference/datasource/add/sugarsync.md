# Sugarsync

{% code fullWidth="true" %}
```
名称:
   singularity datasource add sugarsync - Sugarsync

用法:
   singularity datasource add sugarsync [命令选项] <数据集名称> <源路径>

描述:
   --sugarsync-access-key-id
      Sugarsync访问密钥ID。
      
      不填则使用rclone的。

   --sugarsync-app-id
      Sugarsync应用程序ID。
      
      不填则使用rclone的。

   --sugarsync-authorization
      Sugarsync授权。
      
      通常留空，将由rclone自动配置。

   --sugarsync-authorization-expiry
      Sugarsync授权过期时间。
      
      通常留空，将由rclone自动配置。

   --sugarsync-deleted-id
      Sugarsync已删除文件夹的ID。
      
      通常留空，将由rclone自动配置。

   --sugarsync-encoding
      后端的编码方式。
      
      详见[概览中的编码章节](/overview/#encoding)以获取更多信息。

   --sugarsync-hard-delete
      如果为true，则永久删除文件；否则将文件放入已删除文件夹中。

   --sugarsync-private-access-key
      Sugarsync私有访问密钥。
      
      不填则使用rclone的。

   --sugarsync-refresh-token
      Sugarsync刷新令牌。
      
      通常留空，将由rclone自动配置。

   --sugarsync-root-id
      Sugarsync根目录的ID。
      
      通常留空，将由rclone自动配置。

   --sugarsync-user
      Sugarsync用户。
      
      通常留空，将由rclone自动配置。


选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险操作] 导出数据集到CAR文件后，删除数据集中的文件。 (default: false)
   --rescan-interval value  当距离上次成功扫描已经过去指定的时间间隔时，自动重新扫描源目录 (default: disabled)
   --scanning-state value   设置初始扫描状态 (default: ready)

   Sugarsync专用选项

   --sugarsync-access-key-id value         Sugarsync访问密钥ID。[$SUGARSYNC_ACCESS_KEY_ID]
   --sugarsync-app-id value                Sugarsync应用程序ID。[$SUGARSYNC_APP_ID]
   --sugarsync-authorization value         Sugarsync授权。[$SUGARSYNC_AUTHORIZATION]
   --sugarsync-authorization-expiry value  Sugarsync授权过期时间。[$SUGARSYNC_AUTHORIZATION_EXPIRY]
   --sugarsync-deleted-id value            Sugarsync已删除文件夹的ID。[$SUGARSYNC_DELETED_ID]
   --sugarsync-encoding value              后端的编码方式（默认值: "Slash,Ctl,InvalidUtf8,Dot"）。[$SUGARSYNC_ENCODING]
   --sugarsync-hard-delete value           如果为true，则永久删除文件（默认值: "false"）。[$SUGARSYNC_HARD_DELETE]
   --sugarsync-private-access-key value    Sugarsync私有访问密钥。[$SUGARSYNC_PRIVATE_ACCESS_KEY]
   --sugarsync-refresh-token value         Sugarsync刷新令牌。[$SUGARSYNC_REFRESH_TOKEN]
   --sugarsync-root-id value               Sugarsync根目录的ID。[$SUGARSYNC_ROOT_ID]
   --sugarsync-user value                  Sugarsync用户。[$SUGARSYNC_USER]
```
{% endcode %}