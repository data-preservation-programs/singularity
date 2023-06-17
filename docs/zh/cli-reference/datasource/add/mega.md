# Mega

{% code fullWidth="true" %}
```
名称：
  singularity datasource add mega - Mega

使用方式：
  singularity datasource add mega [命令选项] <数据集名称> <源路径>

描述：
  --mega-encoding
      后端的编码方式。
      
      有关详细信息，请参阅[概览中的编码部分](/overview/#encoding)。

  --mega-user
      用户名。

  --mega-pass
      密码。

  --mega-debug
      输出来自Mega的更多调试信息。
      
      如果设置了这个标志（以及-vv），rclone将从Mega后端打印更多的调试信息。

  --mega-hard-delete
      永久删除文件而不是将其放入回收站。
      
      通常，MEGA后端会将所有删除操作放在回收站内而不是永久删除它们。如果指定了这个标志，rclone将会永久删除文件。

  --mega-use-https
      使用HTTPS进行传输。
      
      MEGA默认使用纯文本HTTP连接。一些ISP会限制HTTP连接，并导致传输变得非常缓慢。启用此选项将会强制MEGA使用HTTPS进行所有传输。HTTPS通常是不必要的，因为所有数据都已经被加密了。启用此功能将会增加CPU使用率并增加网络开销。

选项：
  --help，-h  显示帮助

  数据准备选项

  --delete-after-export    [危险选项]导出数据集为CAR文件后删除数据集中的文件。(默认值: false)
  --rescan-interval value  上次成功扫描后，自动重新扫描源目录的时间间隔。(默认值: 禁用)

  Mega选项

  --mega-debug value        输出来自Mega的更多调试信息。(默认值: "false") [$MEGA_DEBUG]
  --mega-encoding value     后端的编码方式。(默认值: "Slash,InvalidUtf8,Dot") [$MEGA_ENCODING]
  --mega-hard-delete value  永久删除文件而不是将其放入回收站。(默认值: "false") [$MEGA_HARD_DELETE]
  --mega-pass value         密码。[$MEGA_PASS]
  --mega-use-https value    使用HTTPS进行传输。(默认值: "false") [$MEGA_USE_HTTPS]
  --mega-user value         用户名。[$MEGA_USER]
```
{% endcode %}