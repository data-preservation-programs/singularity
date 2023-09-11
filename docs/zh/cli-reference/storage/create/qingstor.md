# QingCloud对象存储

{% code fullWidth="true" %}
```
名称：
   singularity storage create qingstor - 青云对象存储

用法：
   singularity storage create qingstor [命令选项] [参数...]

描述：
   --env-auth
      从运行环境获取青云对象存储凭证。
      
      只在access_key_id和secret_access_key为空时生效。

      示例：
         | false | 在下一步输入青云对象存储凭证。
         | true  | 从环境变量或IAM获取青云对象存储凭证。

   --access-key-id
      青云对象存储的Access Key ID。
      
      若需匿名访问或运行时凭证，可将其留空。

   --secret-access-key
      青云对象存储的Secret Access Key（密码）。
      
      若需匿名访问或运行时凭证，可将其留空。

   --endpoint
      输入一个连接青云对象存储API的终端URL。
      
      若留空，则使用默认值"https://qingstor.com:443"。

   --zone
      要连接的区域。
      
      默认值为"pek3a"。

      示例：
         | pek3a | 北京第三区域（中国）。
         |       | 需要位置约束pek3a。
         | sh1a  | 上海第一区域（中国）。
         |       | 需要位置约束sh1a。
         | gd2a  | 广东第二区域（中国）。
         |       | 需要位置约束gd2a。

   --connection-retries
      连接重试次数。

   --upload-cutoff
      切换到分片上传的临界点。
      
      大于此大小的文件将以分片的方式上传。
      最小值为0，最大值为5 GB。

   --chunk-size
      用于上传的分片大小。
      
      当上传的文件大于upload_cutoff时，将使用此分片大小进行多部分上传。
      
      注意，每个传输中，"--qingstor-upload-concurrency"个此大小的分片被缓存在内存中。
      
      若在高速链路上传输大文件，并且有足够的内存，则增加此大小可加快传输速度。

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传同一文件的分片数。
      
      注意，如果将此值设置为大于1，则多部分上传的校验和将变得损坏（但上传本身不会损坏）。
      
      若在高速链路上上传少量大文件，并且这些上传未能充分利用带宽，则增加此值可能有助于加快传输速度。

   --encoding
      后端的编码方式。
      
      更多信息请参见[概述中的编码方式部分](/overview/#encoding)。


选项：
   --access-key-id value      青云对象存储的Access Key ID。[$ACCESS_KEY_ID]
   --endpoint value           输入一个连接青云对象存储API的终端URL。[$ENDPOINT]
   --env-auth                 从运行环境获取青云对象存储凭证。（默认值：false）[$ENV_AUTH]
   --help, -h                 显示帮助
   --secret-access-key value  青云对象存储的Secret Access Key（密码）。[$SECRET_ACCESS_KEY]
   --zone value               要连接的区域。[$ZONE]

   高级选项

   --chunk-size value          用于上传的分片大小。（默认值：“4Mi”）[$CHUNK_SIZE]
   --connection-retries value  连接重试次数。（默认值：3）[$CONNECTION_RETRIES]
   --encoding value            后端的编码方式。（默认值：“Slash,Ctl,InvalidUtf8”）[$ENCODING]
   --upload-concurrency value  多部分上传的并发数。（默认值：1）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value       切换到分片上传的临界点。（默认值：“200Mi”）[$UPLOAD_CUTOFF]

   通用选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}