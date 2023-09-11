# Koofr, https://app.koofr.net/

{% code fullWidth="true" %}
```
名称:
   singularity storage update koofr koofr - Koofr, https://app.koofr.net/

用法:
   singularity storage update koofr koofr [命令选项] <名称|ID>

描述:
   --mountid
      要使用的挂载ID。
      
      如果省略，则使用主要挂载。

   --setmtime
      后端是否支持设置修改时间。
      
      如果您使用指向Dropbox或Amazon Drive后端的挂载ID，请将其设置为false。

   --user
      您的用户名。

   --password
      您在rclone上的密码（在https://app.koofr.net/app/admin/preferences/password生成）。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见概述中的[编码部分](/overview/#encoding)。


选项:
   --help, -h         显示帮助信息
   --password value   您在rclone上的密码（在https://app.koofr.net/app/admin/preferences/password生成）。 [$PASSWORD]
   --user value       您的用户名。 [$USER]

   高级选项

   --encoding value  后端的编码方式。（默认值："Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"） [$ENCODING]
   --mountid value    要使用的挂载ID。[$MOUNTID]
   --setmtime         后端是否支持设置修改时间。（默认值：true） [$SETMTIME]

```
{% endcode %}