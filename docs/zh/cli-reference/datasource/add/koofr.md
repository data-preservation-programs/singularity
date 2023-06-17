# Koofr，Digi Storage和其他与Koofr兼容的存储提供商

{% code fullWidth="true" %}
```
名称：
   singularity datasource add koofr - Koofr，Digi Storage和其他与Koofr兼容的存储提供商

用法：
   singularity datasource add koofr [命令选项] <dataset_name> <source_path>

说明：
   --koofr-mountid
      要使用的装载点ID。
      
      如果省略，使用主装载点。

   --koofr-setmtime
      后端是否支持设置修改时间。
      
      如果您使用指向Dropbox或Amazon Drive后端的装载点ID，请将其设置为false。

   --koofr-user
      您的用户名。

   --koofr-password
      [提供商] - koofr
         您在rclone中的密码（在https://app.koofr.net/app/admin/preferences/password上生成）。

      [提供商] - digistorage
         您在rclone中的密码（在https://storage.rcs-rds.ro/app/admin/preferences/password上生成）。

      [提供商] - other
         您在rclone中的密码（在服务的设置页面上生成）。

   --koofr-encoding
      后端的编码。
      
      有关详细信息，请参见概述中的[编码部分](/overview/#encoding)。

   --koofr-provider
      选择您的存储提供商。

      例如：
         | koofr       | Koofr，https://app.koofr.net/
         | digistorage | Digi Storage，https://storage.rcs-rds.ro/
         | other       | 任何其他与Koofr API兼容的存储服务

   --koofr-endpoint
      [提供商] - other
         要使用的Koofr API端点。


选项：
   --help，-h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险]将数据集导出成CAR文件后将数据集中的文件删除。 （默认值：false）
   --rescan-interval value  当距离上次成功扫描已过去此时间间隔时，自动重新扫描源目录（默认值：禁用）

   Koofr的选项

   --koofr-encoding value  后端的编码。（默认值：“Slash，BackSlash，Del，Ctl，InvalidUtf8，Dot”）[$KOOFR_ENCODING]
   --koofr-endpoint value  要使用的Koofr API端点。[$KOOFR_ENDPOINT]
   --koofr-mountid value   要使用的装载点ID。[$KOOFR_MOUNTID]
   --koofr-password value  您在rclone中的密码（在https://app.koofr.net/app/admin/preferences/password上生成）。[$KOOFR_PASSWORD]
   --koofr-provider value  选择您的存储提供商。[$KOOFR_PROVIDER]
   --koofr-setmtime value  后端是否支持设置修改时间。（默认值：“true”）[$KOOFR_SETMTIME]
   --koofr-user value      您的用户名。[$KOOFR_USER]

```
{% endcode %}