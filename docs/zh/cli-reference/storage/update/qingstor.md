# QingCloud 对象存储

{% code fullWidth="true" %}
```
名称：
   singularity storage update qingstor - QingCloud 对象存储

用法：
   singularity storage update qingstor [命令选项] <名称|ID>

描述：
   --env-auth
      从运行时获取 QingStor 凭据。
      
      仅当 access_key_id 和 secret_access_key 为空时适用。

      示例：
         | false | 在下一步中输入 QingStor 凭据。
         | true  | 从环境中获取 QingStor 凭据（环境变量或 IAM）。

   --access-key-id
      QingStor Access Key ID。
      
      留空以进行匿名访问或使用运行时凭据。

   --secret-access-key
      QingStor Secret Access Key（密码）。
      
      留空以进行匿名访问或使用运行时凭据。

   --endpoint
      输入连接 QingStor API 的端点URL。
      
      留空将使用默认值 "https://qingstor.com:443"。

   --zone
      要连接的区域。
      
      默认值为 "pek3a"。

      示例：
         | pek3a | 北京三区。
         |       | 需要 location constraint pek3a。
         | sh1a  | 上海一区。
         |       | 需要 location constraint sh1a。
         | gd2a  | 广东二区。
         |       | 需要 location constraint gd2a。

   --connection-retries
      连接重试次数。

   --upload-cutoff
      切换到分块上传的界限。
      
      大于此大小的任何文件都将分块上传，分块大小为 chunk_size。
      最小值为 0，最大值为 5 GiB。

   --chunk-size
      用于上传的分块大小。
      
      当上传的文件大于 upload_cutoff 时，将使用该分块大小进行分块上传。
      
      请注意，"--qingstor-upload-concurrency" 每个传输中会缓冲此大小的分块。
      
      如果您通过高速链接传输大文件并且有足够的内存，增加此值将加快传输速度。

   --upload-concurrency
      分块上传的并发数。
      
      这是同时上传同一文件的分块数。
      
      如果将此值设置为 > 1，则分块上传的校验和会损坏（上传本身不会损坏）。
      
      如果您通过高速链接上传少量大文件，并且这些上传没有充分利用您的带宽，增加此值可能有助于加快传输速度。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。


选项：
   --access-key-id value      QingStor Access Key ID。[$ACCESS_KEY_ID]
   --endpoint value           输入连接 QingStor API 的端点URL。[$ENDPOINT]
   --env-auth                 从运行时获取 QingStor 凭据。（默认值：false）[$ENV_AUTH]
   --help, -h                 显示帮助
   --secret-access-key value  QingStor Secret Access Key（密码）。[$SECRET_ACCESS_KEY]
   --zone value               要连接的区域。[$ZONE]

   高级选项

   --chunk-size value          用于上传的分块大小。（默认值："4Mi"）[$CHUNK_SIZE]
   --connection-retries value  连接重试次数。（默认值：3）[$CONNECTION_RETRIES]
   --encoding value            后端的编码方式。（默认值："Slash,Ctl,InvalidUtf8"）[$ENCODING]
   --upload-concurrency value  分块上传的并发数。（默认值：1）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value       切换到分块上传的界限。（默认值："200Mi"）[$UPLOAD_CUTOFF]

```
{% endcode %}