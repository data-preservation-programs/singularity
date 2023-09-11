# 所有其他兼容Koofr API的存储服务

{% code fullWidth="true" %}
```
名称：
   singularity storage create koofr other - 任何其他兼容Koofr API的存储服务

用法：
   singularity storage create koofr other [命令选项] [参数...]

描述：
   --endpoint
      要使用的Koofr API端点。

   --mountid
      要使用的挂载ID。
     
     如果忽略，将使用主挂载。

   --setmtime
      后端是否支持设置修改时间。
      
      如果您使用的是指向Dropbox或亚马逊云存储后端的挂载ID，请将其设置为false。

   --user
      您的用户名。

   --password
      您在rclone上的密码（在服务的设置页面上生成）。

   --encoding
      后端的编码。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。


选项：
   --endpoint value  要使用的Koofr API端点。 [$ENDPOINT]
   --help, -h        显示帮助
   --password value  您在rclone上的密码（在服务的设置页面上生成）。 [$PASSWORD]
   --user value      您的用户名。 [$USER]

   高级选项

   --encoding value  后端的编码。 （默认值：“Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot”）[$ENCODING]
   --mountid value   要使用的挂载ID。 [$MOUNTID]
   --setmtime        后端是否支持设置修改时间。 （默认值：true）[$SETMTIME]

   通用选项

   --name value  存储的名称（默认值：Auto generated）
   --path value  存储的路径

```
{% endcode %}