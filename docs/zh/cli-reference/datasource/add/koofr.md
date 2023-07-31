# Koofr, Digi Storage和其他兼容Koofr的存储提供商

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add koofr - Koofr, Digi Storage和其他兼容Koofr的存储提供商

USAGE:
   singularity datasource add koofr [command options] <dataset_name> <source_path>

DESCRIPTION:
   --koofr-encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --koofr-endpoint
      [提供商] - 其他
         要使用的Koofr API端点。

   --koofr-mountid
      要使用的挂载的挂载ID。
      
      如果省略，将使用主要挂载。

   --koofr-password
      [提供商] - koofr
         您用于rclone的密码（在https://app.koofr.net/app/admin/preferences/password 上生成）。

      [提供商] - digistorage
         您用于rclone的密码（在https://storage.rcs-rds.ro/app/admin/preferences/password 上生成）。

      [提供商] - 其他
         您用于rclone的密码（在您的服务的设置页面上生成）。

   --koofr-provider
      选择您的存储提供商。

      示例:
         | koofr       | Koofr, https://app.koofr.net/
         | digistorage | Digi Storage, https://storage.rcs-rds.ro/
         | other       | 任何其他兼容Koofr API的存储服务

   --koofr-setmtime
      后端是否支持设置修改时间。
      
      如果您使用的是指向Dropbox或Amazon Drive后端的挂载ID，请将其设置为false。

   --koofr-user
      您的用户名。


OPTIONS:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 将数据集导出为CAR文件后删除数据集的文件。 (默认: false)
   --rescan-interval value  当距离上次成功扫描的时间间隔超过此值时，自动重新扫描源目录（默认: 禁用）
   --scanning-state value   设置初始扫描状态 (默认: 就绪)

   Koofr选项

   --koofr-encoding value  后端的编码方式。 (默认: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$KOOFR_ENCODING]
   --koofr-endpoint value  要使用的Koofr API端点。 [$KOOFR_ENDPOINT]
   --koofr-mountid value   要使用的挂载的挂载ID。 [$KOOFR_MOUNTID]
   --koofr-password value  您用于rclone的密码（在https://app.koofr.net/app/admin/preferences/password 上生成）。 [$KOOFR_PASSWORD]
   --koofr-provider value  选择您的存储提供商。 [$KOOFR_PROVIDER]
   --koofr-setmtime value  后端是否支持设置修改时间。 (默认: "true") [$KOOFR_SETMTIME]
   --koofr-user value      您的用户名。 [$KOOFR_USER]

```
{% endcode %}