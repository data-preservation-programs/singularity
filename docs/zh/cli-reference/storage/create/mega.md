# Mega

{% code fullWidth="true" %}
```
名称：
   singularity storage create mega - Mega

用法：
   singularity storage create mega [命令选项] [参数...]

描述：
   --user
      用户名。

   --pass
      密码。

   --debug
      输出更多Mega的调试信息。
      
      如果设置了该标志（与-vv一起），它将从Mega后端打印更多的调试信息。

   --hard-delete
      永久删除文件而不是将其放入回收站。
      
      通常，Mega后端会将所有删除的文件放入回收站而不是永久删除它们。如果您指定了这个选项，rclone将永久删除对象。

   --use-https
      使用HTTPS进行传输。
      
      MEGA默认使用明文HTTP连接。一些互联网服务提供商会限制HTTP连接，这会导致传输速度变得非常慢。启用此选项将强制MEGA在所有传输中使用HTTPS。由于所有数据已经加密，通常不需要HTTPS。启用它将增加CPU使用和网络开销。

   --encoding
      后端的编码方式。
      
      有关详细信息，请参见[概述中的编码部分](/overview/#encoding)。

选项：
   --help, -h    显示帮助
   --pass value  密码。[$PASS]
   --user value  用户名。[$USER]

   高级选项

   --debug           输出更多Mega的调试信息。（默认值：false）[$DEBUG]
   --encoding value  后端的编码方式。（默认值："Slash,InvalidUtf8,Dot"）[$ENCODING]
   --hard-delete     永久删除文件而不是将其放入回收站。（默认值：false）[$HARD_DELETE]
   --use-https       使用HTTPS进行传输。（默认值：false）[$USE_HTTPS]

   基本选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}