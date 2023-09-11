# Digi Storage, https://storage.rcs-rds.ro/

{% code fullWidth="true" %}
```
名称：
   singularity storage create koofr digistorage - Digi Storage, https://storage.rcs-rds.ro/

用法：
   singularity storage create koofr digistorage [命令选项] [参数…]

说明：
   --mountid
      要使用的挂载的挂载 ID 。
      
      如果省略，则使用主挂载。

   --setmtime
      后端是否支持设置修改时间。
      
      如果使用指向 Dropbox 或 Amazon Drive 后端的挂载 ID，请将此选项设置为 false。

   --user
      您的用户名。

   --password
      您在 rclone 上的密码（在 https://storage.rcs-rds.ro/app/admin/preferences/password 生成）。

   --encoding
      后端的编码。
   
      有关更多详细信息，请参阅[概览中的编码部分](/overview/#encoding)。


选项：
   --help, -h        显示帮助信息
   --password value  您在 rclone 上的密码（在 https://storage.rcs-rds.ro/app/admin/preferences/password 生成）。[$PASSWORD]
   --user value      您的用户名。[$USER]

   高级选项

   --encoding value  后端的编码。（默认值："Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"）[$ENCODING]
   --mountid value   要使用的挂载的挂载 ID。[$MOUNTID]
   --setmtime        后端是否支持设置修改时间。（默认值：true）[$SETMTIME]

   常规选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}