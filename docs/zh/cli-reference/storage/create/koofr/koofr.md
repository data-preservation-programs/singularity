# Koofr, https://app.koofr.net/

{% code fullWidth="true" %}
```
命令：
   singularity storage create koofr koofr - Koofr，https://app.koofr.net/

用法：
   singularity storage create koofr koofr [命令选项] [参数...]

描述：
   --mountid
      要使用的挂载的挂载ID。
      
      如果省略，将使用主要挂载。

   --setmtime
      后端是否支持修改修改时间。
      
      如果使用指向Dropbox或Amazon Drive后端的挂载ID，请将此选项设置为false。

   --user
      您的用户名。

   --password
      rclone的密码（在https://app.koofr.net/app/admin/preferences/password生成）。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。


选项：
   --help, -h        显示帮助信息
   --password value  rclone的密码（在https://app.koofr.net/app/admin/preferences/password生成）。[$PASSWORD]
   --user value      您的用户名。[$USER]

   高级选项

   --encoding value  后端的编码方式。 (默认值：“Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot”) [$ENCODING]
   --mountid value   要使用的挂载的挂载ID。[$MOUNTID]
   --setmtime        后端是否支持修改修改时间。 (默认值:true) [$SETMTIME]

   常规选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}