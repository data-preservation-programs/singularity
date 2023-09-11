# 互联网档案馆

```
名称：
   singularity storage create internetarchive - 互联网档案馆

用法：
   singularity storage create internetarchive [命令选项] [参数...]

描述：
   --access-key-id
      IAS3 访问密钥。

      留空以进行匿名访问。
      您可以在此处找到一个：https://archive.org/account/s3.php

   --secret-access-key
      IAS3 密钥（密码）。

      留空以进行匿名访问。

   --endpoint
      IAS3 终端。

      留空以使用默认值。

   --front-endpoint
      互联网档案馆前端的主机。

      留空以使用默认值。

   --disable-checksum
      不要要求服务器根据 rclone 计算的 MD5 校验和进行测试。
      通常，rclone 会在上传之前计算输入的 MD5 校验和，以便询问服务器根据校验和检查对象。
      这对于数据完整性检查很有用，但对于大文件启动上传可能会导致长时间延迟。

   --wait-archive
      等待服务器处理任务（特别是归档和 book_op）完成的超时时间。
      仅当您需要确保在写入操作后能够反映更改时才启用。
      设置为 0 表示禁用等待。超时时不会抛出错误。

   --encoding
      后端的编码方式。

      更多信息请参见[概述部分的编码说明](/overview/#encoding)。


选项：
   --access-key-id value      IAS3 访问密钥。[$ACCESS_KEY_ID]
   --help, -h                 显示帮助信息
   --secret-access-key value  IAS3 密钥（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --disable-checksum      不要要求服务器根据 rclone 计算的 MD5 校验和进行测试。（默认值：true）[$DISABLE_CHECKSUM]
   --encoding value        后端的编码方式。（默认值："Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot"）[$ENCODING]
   --endpoint value        IAS3 终端。（默认值："https://s3.us.archive.org"）[$ENDPOINT]
   --front-endpoint value  互联网档案馆前端的主机。（默认值："https://archive.org"）[$FRONT_ENDPOINT]
   --wait-archive value    等待服务器处理任务（特别是归档和 book_op）完成的超时时间。（默认值："0s"）[$WAIT_ARCHIVE]

   常规选项

   --name value  存储的名称（默认值：自动生成的）
   --path value  存储的路径
```