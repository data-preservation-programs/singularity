# Digi Storage, https://storage.rcs-rds.ro/

{% code fullWidth="true" %}
```
指令名称:
   singularity storage update koofr digistorage - Digi Storage, https://storage.rcs-rds.ro/

用法:
   singularity storage update koofr digistorage [命令选项] <名称|ID>

说明:
   --mountid
      要使用的挂载ID。
      
      如果省略，将使用主挂载。

   --setmtime
      后端是否支持设置修改时间。
      
      如果您使用的是指向Dropbox或Amazon Drive后端的挂载ID，请将此选项设置为false。

   --user
      您的用户名。

   --password
      您的rclone密码（可在https://storage.rcs-rds.ro/app/admin/preferences/password上生成）。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

选项:
   --help, -h        显示帮助
   --password value  您的rclone密码（可在https://storage.rcs-rds.ro/app/admin/preferences/password上生成）。 [$PASSWORD]
   --user value      您的用户名。 [$USER]

   高级选项

   --encoding value  后端的编码方式。 (默认值: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --mountid value   要使用的挂载ID。[$MOUNTID]
   --setmtime        后端是否支持设置修改时间。 (默认值: true) [$SETMTIME]

```
{% endcode %}