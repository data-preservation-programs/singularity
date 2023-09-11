# Mega

{% code fullWidth="true" %}
```
命令:
   singularity storage update mega - 大型云存储

用法:
   singularity storage update mega [命令选项] <名称|ID>

说明:
   --user
      用户名。

   --pass
      密码。

   --debug
      输出来自大型云存储的更多调试信息。
      
      如果设置了该标志（以及-vv），将从大型云存储输出进一步的调试信息。

   --hard-delete
      永久删除文件而不将其放入回收站。
      
      通常，大型云存储将所有删除的文件放入回收站而不是永久删除它们。
      如果指定了该标志，则rclone将永久删除对象。

   --use-https
      使用HTTPS进行传输。
      
      大型云存储默认使用明文HTTP连接。
      一些ISP会限制HTTP连接，这会导致传输变得非常缓慢。
      启用此选项将强制大型云存储在所有传输中使用HTTPS。
      HTTPS通常不是必需的，因为所有数据已经加密。
      启用它将增加CPU使用率并增加网络开销。

   --encoding
      后端的编码。
      
      有关详细信息，请参阅[概述中的编码部分](/overview/#encoding)。


选项:
   --help, -h    显示帮助信息
   --pass value  密码。[$PASS]
   --user value  用户名。[$USER]

   高级选项

   --debug           输出来自大型云存储的更多调试信息。 (默认值: false) [$DEBUG]
   --encoding value  后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete     永久删除文件而不将其放入回收站。 (默认值: false) [$HARD_DELETE]
   --use-https       使用HTTPS进行传输。 (默认值: false) [$USE_HTTPS]
```
{% endcode %}