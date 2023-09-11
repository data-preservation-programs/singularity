# 其他兼容 Koofr API 的存储服务

{% code %}
```
名称:
   singularity storage update koofr other - 其他兼容 Koofr API 的存储服务

用法:
   singularity storage update koofr other [命令选项] <名称|ID>

描述:
   --endpoint
      要使用的 Koofr API 地址。

   --mountid
      要使用的挂载 ID。
      
      如果省略，将使用主挂载。

   --setmtime
      后端是否支持设置修改时间。
      
      如果使用指向 Dropbox 或 Amazon Drive 后端的挂载 ID，请将此项设置为 false。

   --user
      您的用户名。

   --password
      您在 rclone 上的密码（在服务的设置页面生成一个密码）。

   --encoding
      后端的编码。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。


选项:
   --endpoint value  要使用的 Koofr API 地址。 [$ENDPOINT]
   --help, -h        显示帮助
   --password value  您在 rclone 上的密码（在服务的设置页面生成一个密码）。 [$PASSWORD]
   --user value      您的用户名。 [$USER]

   高级选项

   --encoding value  后端的编码。（默认值: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"）[$ENCODING]
   --mountid value   要使用的挂载 ID。[$MOUNTID]
   --setmtime        后端是否支持设置修改时间。（默认值: true）[$SETMTIME]

```
{% endcode %}