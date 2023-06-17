# Sugarsync

{% code fullWidth="true" %}
```
名称：
   singularity 数据源添加 sugarsync - Sugarsync

用法：
   singularity datasource add sugarsync [command options] <数据集名称> <源路径>

说明：
   --sugarsync-user
      Sugarsync用户。
      
      通常情况下留空，将由rclone自动配置。

   --sugarsync-encoding
      后端的编码方式。
      
      更多信息请参见 [概述中的编码章节](/overview/#encoding)。

   --sugarsync-access-key-id
      Sugarsync Access Key ID。
      
      留空以使用rclone的。

   --sugarsync-hard-delete
      如果为true，则永久删除文件；否则将它们放入已删除文件夹中。

   --sugarsync-refresh-token
      Sugarsync refresh token。
      
      通常情况下留空，将由rclone自动配置。

   --sugarsync-authorization
      Sugarsync authorization。
      
      通常情况下留空，将由rclone自动配置。

   --sugarsync-authorization-expiry
      Sugarsync authorization expiry。
      
      通常情况下留空，将由rclone自动配置。

   --sugarsync-root-id
      Sugarsync root id。
      
      通常情况下留空，将由rclone自动配置。

   --sugarsync-deleted-id
      Sugarsync deleted folder id。
      
      通常情况下留空，将由rclone自动配置。

   --sugarsync-app-id
      Sugarsync App ID。
      
      留空以使用rclone的。

   --sugarsync-private-access-key
      Sugarsync Private Access Key。
      
      留空以使用rclone的。


选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 将数据集导出为CAR文件后删除数据集的文件。  (默认值: false)
   --rescan-interval value  当上次成功扫描后经过了此间隔时，自动重新扫描源目录 (默认值: 禁用）

   Sugarsync 选项

   --sugarsync-access-key-id value         Sugarsync Access Key ID。 [$SUGARSYNC_ACCESS_KEY_ID]
   --sugarsync-app-id value                Sugarsync App ID。 [$SUGARSYNC_APP_ID]
   --sugarsync-authorization value         Sugarsync authorization。 [$SUGARSYNC_AUTHORIZATION]
   --sugarsync-authorization-expiry value  Sugarsync authorization expiry。 [$SUGARSYNC_AUTHORIZATION_EXPIRY]
   --sugarsync-deleted-id value            Sugarsync deleted folder id。 [$SUGARSYNC_DELETED_ID]
   --sugarsync-encoding value              后端的编码方式。（默认值："Slash,Ctl,InvalidUtf8,Dot"）[$SUGARSYNC_ENCODING]
   --sugarsync-hard-delete value           如果为true，则永久删除文件（默认值："false"）[$SUGARSYNC_HARD_DELETE]
   --sugarsync-private-access-key value    Sugarsync Private Access Key。 [$SUGARSYNC_PRIVATE_ACCESS_KEY]
   --sugarsync-refresh-token value         Sugarsync refresh token。[$SUGARSYNC_REFRESH_TOKEN]
   --sugarsync-root-id value               Sugarsync root id。 [$SUGARSYNC_ROOT_ID]
   --sugarsync-user value                  Sugarsync用户。 [$SUGARSYNC_USER]

```
{% endcode %}