# 青云对象存储

```NAME:
   singularity datasource add qingstor - 青云对象存储

USAGE:
   singularity datasource add qingstor [command options] <dataset_name> <source_path>

DESCRIPTION:
   --qingstor-access-key-id
      青云存储的 Access Key ID。
      
      留空以匿名访问或在运行时获取凭证。

   --qingstor-chunk-size
      用于上传的块大小。
      
      当上传大于上传截止值的文件时，将使用此块大小进行分块上传。
      
      请注意，每个传输都会在内存中缓冲此块大小的"--qingstor-upload-concurrency"块。
      
      如果您在高速链接上传输大文件并且具有足够的内存，增加此值会加快传输速度。

   --qingstor-connection-retries
      连接重试次数。

   --qingstor-encoding
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --qingstor-endpoint
      输入一个端点URL以连接青云存储的 API。
      
      留空将使用默认值 "https://qingstor.com:443"。

   --qingstor-env-auth
      从运行时获取青云存储的凭证。
      
      只适用于访问密钥 ID 和密钥的情况下留空的情况。

      示例：
         | false | 在下一步中输入青云存储的凭证。
         | true  | 从环境（env vars 或 IAM）获取青云存储的凭证。

   --qingstor-secret-access-key
      青云存储的 Secret Access Key（密码）。
      
      留空以匿名访问或在运行时获取凭证。

   --qingstor-upload-concurrency
      分块上传的并发数。
      
      这是同时上传同一文件的块数。
      
      注意，如果将此值设置为 >1，则多部分上传的校验和将损坏（但上传本身不会损坏）。
      
      如果您在高速链接上上传少量大文件，并且这些上传未充分利用您的带宽，增加此值可能有助于加快传输速度。

   --qingstor-upload-cutoff
      切换到分块上传的截止值。
      
      大于此值的任何文件都会以块大小的分块上传。
      最小值为0，最大值为5 GB。

   --qingstor-zone
      要连接的区域。
      
      默认为“pek3a”。

      示例：
         | pek3a | 北京三区。
                 | 需要位置约束 pek3a。
         | sh1a  | 上海一区。
                 | 需要位置约束 sh1a。
         | gd2a  | 广东二区。
                 | 需要位置约束 gd2a。


OPTIONS:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险] 导出 CAR 文件后删除数据集的文件。（默认值：false）
   --rescan-interval value  当距离上次成功扫描的时间间隔达到该时间时，自动重新扫描源目录。（默认值：禁用）
   --scanning-state value   设置初始扫描状态。（默认值：准备就绪）

   qingstor选项

   --qingstor-access-key-id value       青云存储的 Access Key ID。[$QINGSTOR_ACCESS_KEY_ID]
   --qingstor-chunk-size value          用于上传的块大小。（默认值："4Mi"）[$QINGSTOR_CHUNK_SIZE]
   --qingstor-connection-retries value  连接重试次数。（默认值："3"）[$QINGSTOR_CONNECTION_RETRIES]
   --qingstor-encoding value            后端的编码。（默认值："Slash,Ctl,InvalidUtf8"）[$QINGSTOR_ENCODING]
   --qingstor-endpoint value            输入一个端点URL以连接青云存储的 API。[$QINGSTOR_ENDPOINT]
   --qingstor-env-auth value            从运行时获取青云存储的凭证。（默认值："false"）[$QINGSTOR_ENV_AUTH]
   --qingstor-secret-access-key value   青云存储的 Secret Access Key（密码）。[$QINGSTOR_SECRET_ACCESS_KEY]
   --qingstor-upload-concurrency value  分块上传的并发数。（默认值："1"）[$QINGSTOR_UPLOAD_CONCURRENCY]
   --qingstor-upload-cutoff value       切换到分块上传的截止值。（默认值："200Mi"）[$QINGSTOR_UPLOAD_CUTOFF]
   --qingstor-zone value                要连接的区域。[$QINGSTOR_ZONE]```

```