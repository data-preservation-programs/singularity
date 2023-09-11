# 网络存档

{% code fullWidth="true" %}
```
名称:
   singularity storage update internetarchive - 网络存档

用法:
   singularity storage update internetarchive [command options] <name|id>

描述:
   --access-key-id
      IAS3 访问密钥。
      
      留空表示匿名访问。
      您可以在此处找到一个：https://archive.org/account/s3.php

   --secret-access-key
      IAS3 密钥（密码）。
      
      留空表示匿名访问。

   --endpoint
      IAS3 终端。
      
      留空以使用默认值。

   --front-endpoint
      InternetArchive 前端的主机。
      
      留空以使用默认值。

   --disable-checksum
      不要请求服务器对 rclone 计算的 MD5 校验和进行测试。
      通常，rclone 会在上传前计算输入的 MD5 校验和，以便请求服务器根据校验和检查对象。
      这对于数据完整性检查非常有用，但对于大文件可能导致长时间的上传启动延迟。

   --wait-archive
      等待服务器处理任务（特别是存档和 book_op）完成的超时时间。
      仅在需要写操作后保证被反映时启用。
      0 表示禁用等待。超时情况下不抛出错误。

   --encoding
      后端的编码方式。
      
      参见[概览中的编码部分](/overview/#encoding)了解更多信息。


选项:
   --access-key-id value      IAS3 访问密钥。[$ACCESS_KEY_ID]
   --help, -h                 显示帮助
   --secret-access-key value  IAS3 密钥（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --disable-checksum      不要请求服务器对 rclone 计算的 MD5 校验和进行测试。 (默认值: true) [$DISABLE_CHECKSUM]
   --encoding value        后端的编码方式。 (默认值: "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value        IAS3 终端。 (默认值: "https://s3.us.archive.org") [$ENDPOINT]
   --front-endpoint value  InternetArchive 前端的主机。 (默认值: "https://archive.org") [$FRONT_ENDPOINT]
   --wait-archive value    等待服务器处理任务（特别是存档和 book_op）完成的超时时间。 (默认值: "0s") [$WAIT_ARCHIVE]

```
{% endcode %}