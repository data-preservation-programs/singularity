# Mega

{% code fullWidth="true" %}
```
命令名称:
   singularity datasource add mega - Mega

使用方法:
   singularity datasource add mega [命令选项] <数据集名称> <源路径>

描述:
   --mega-debug
      输出更多Mega的调试信息。
      
      如果设置了此标志（以及-vv），它将从Mega后端打印更多的调试信息。

   --mega-encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --mega-hard-delete
      永久删除文件，而不是将它们放入回收站。
      
      通常，Mega后端将所有删除操作放入回收站而不是永久删除它们。如果您指定此选项，rclone将永久删除对象。

   --mega-pass
      密码。

   --mega-use-https
      使用HTTPS进行传输。
      
      MEGA默认使用明文HTTP连接。一些ISP限制HTTP连接，这导致传输变得非常缓慢。启用此选项将强制MEGA使用HTTPS进行所有传输。由于所有数据已经加密，通常不需要使用HTTPS。启用它将增加CPU使用率并增加网络开销。

   --mega-user
      用户名。


选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作] 在将数据集导出到CAR文件后删除数据集的文件。 (默认: false)
   --rescan-interval value  上次成功扫描后，当此时间间隔过去时，自动重新扫描源目录（默认：禁用）
   --scanning-state value   设置初始扫描状态（默认：准备就绪）

   Mega选项

   --mega-debug value        输出更多Mega的调试信息。 (默认: "false") [$MEGA_DEBUG]
   --mega-encoding value     后端的编码方式。 (默认: "Slash,InvalidUtf8,Dot") [$MEGA_ENCODING]
   --mega-hard-delete value  永久删除文件，而不是将它们放入回收站。 (默认: "false") [$MEGA_HARD_DELETE]
   --mega-pass value         密码。 [$MEGA_PASS]
   --mega-use-https value    使用HTTPS进行传输。 (默认: "false") [$MEGA_USE_HTTPS]
   --mega-user value         用户名。 [$MEGA_USER]
```
{% endcode %}
