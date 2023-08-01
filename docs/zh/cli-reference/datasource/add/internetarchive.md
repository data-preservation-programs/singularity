# 互联网档案馆

{% code fullWidth="true" %}
```
名称:
   singularity datasource add internetarchive - 互联网档案馆

使用方法:
   singularity datasource add internetarchive [命令选项] <数据集名称> <源路径>

描述:
   --internetarchive-access-key-id
      IAS3访问密钥。
      
      匿名访问留空。
      您可以在此处找到：https://archive.org/account/s3.php

   --internetarchive-disable-checksum
      不要求服务器根据rclone计算的MD5校验和进行测试。
      通常，rclone会在上传之前计算输入的MD5校验和，以便可以要求服务器根据校验和检查对象。
      这对于数据完整性检查非常有用，但对于大文件启动上传可能会导致长时间延迟。

   --internetarchive-encoding
      后端的编码方式。
      
      有关更多信息，请参见[概览中的编码部分](/overview/#encoding)。

   --internetarchive-endpoint
      IAS3端点。
      
      留空使用默认值。

   --internetarchive-front-endpoint
      互联网档案馆前端的主机。
      
      留空使用默认值。

   --internetarchive-secret-access-key
      IAS3密钥（密码）。
      
      匿名访问留空。

   --internetarchive-wait-archive
      等待服务器处理任务（特别是存档和book_op）完成的超时时间。
      仅在需要在写操作后保证反映时启用。
      0禁用等待。在超时情况下不会抛出错误。


选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险操作] 在将数据集导出为CAR文件后删除数据集文件。 (默认值: false)
   --rescan-interval value  当上一次成功扫描后经过一段时间时，自动重新扫描源目录 (默认值: 禁用)
   --scanning-state value   设置初始扫描状态 (默认值: 准备就绪)

   互联网档案馆选项

   --internetarchive-access-key-id value      IAS3访问密钥。[$INTERNETARCHIVE_ACCESS_KEY_ID]
   --internetarchive-disable-checksum value   不要求服务器根据rclone计算的MD5校验和进行测试。 (默认值: "true") [$INTERNETARCHIVE_DISABLE_CHECKSUM]
   --internetarchive-encoding value           后端的编码方式。 (默认值: "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot") [$INTERNETARCHIVE_ENCODING]
   --internetarchive-endpoint value           IAS3端点。 (默认值: "https://s3.us.archive.org") [$INTERNETARCHIVE_ENDPOINT]
   --internetarchive-front-endpoint value     互联网档案馆前端的主机。 (默认值: "https://archive.org") [$INTERNETARCHIVE_FRONT_ENDPOINT]
   --internetarchive-secret-access-key value  IAS3密钥（密码）。[$INTERNETARCHIVE_SECRET_ACCESS_KEY]
   --internetarchive-wait-archive value       等待服务器处理任务（特别是存档和book_op）完成的超时时间。 (默认值: "0s") [$INTERNETARCHIVE_WAIT_ARCHIVE]

```
{% endcode %}