# 网络档案馆

{% code fullWidth="true" %}
```
名称：
   singularity datasource add internetarchive - 网络档案馆

用法：
   singularity datasource add internetarchive [命令选项] <数据集名称> <源路径>

说明：
   --internetarchive-endpoint
      IAS3端点。
      
      留空以使用默认值。

   --internetarchive-front-endpoint
      网络档案馆前端主机。
      
      留空以使用默认值。

   --internetarchive-disable-checksum
      不要向服务器请求针对rclone计算的MD5校验和进行测试。
      通常，在上传之前，rclone将计算输入数据的MD5校验和，
      以便它可以要求服务器根据校验和检查对象。这对于数据完整性检查非常有用，
      但对于开始上传大文件可能会导致长时间的延迟。

   --internetarchive-wait-archive
      服务器处理任务（具体而言是存档和book_op）完成的等待时间限制。
      仅在需要进行写操作后保证结果才启用。0表示禁用等待。超时情况下不会抛出任何错误。

   --internetarchive-encoding
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --internetarchive-access-key-id
      IAS3访问密钥。
      
      留空以进行匿名访问。
      您可以在此处找到Key: https://archive.org/account/s3.php

   --internetarchive-secret-access-key
      IAS3秘密密钥（密码）。
      
      留空以进行匿名访问。


选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    【危险】导出CAR文件后删除数据集文件（默认值:false）
   --rescan-interval value  当从上次成功扫描以来经过此间隔时，自动重新扫描源目录（默认值:已禁用）

   网络档案馆选项

   --internetarchive-access-key-id value      IAS3访问密钥。[$INTERNETARCHIVE_ACCESS_KEY_ID]
   --internetarchive-disable-checksum value   不要向服务器请求针对rclone计算的MD5校验和进行测试。 (默认值:“true”) [$INTERNETARCHIVE_DISABLE_CHECKSUM]
   --internetarchive-encoding value           后端的编码。（默认值：“Slash，LtGt，CrLf，Del，Ctl，InvalidUtf8，Dot”）[$INTERNETARCHIVE_ENCODING]
   --internetarchive-endpoint value           IAS3端点。(默认值：“https://s3.us.archive.org”）[$INTERNETARCHIVE_ENDPOINT]
   --internetarchive-front-endpoint value     网络档案馆前端主机。（默认值：“https://archive.org”）[$INTERNETARCHIVE_FRONT_ENDPOINT]
   --internetarchive-secret-access-key value  IAS3秘密密钥（密码）。[$INTERNETARCHIVE_SECRET_ACCESS_KEY]
   --internetarchive-wait-archive value       服务器处理任务（具体而言是存档和book_op）完成的等待时间限制（默认值：“0s”）[$INTERNETARCHIVE_WAIT_ARCHIVE]

```
{% endcode %}