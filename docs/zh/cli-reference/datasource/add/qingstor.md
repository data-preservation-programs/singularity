# 青云对象存储

{% code fullWidth="true" %}
```
命令名称：
  singularity datasource add qingstor - 青云对象存储

用法：
  singularity datasource add qingstor [command options] < dataset_name > < source_path >

说明：
  --qingstor-endpoint
      输入连接青云API的端点URL。
      
      留空则使用默认值"https://qingstor.com:443"。

  --qingstor-encoding
      后端的编码。
      
      更多信息请参见[总览页面](/overview/#encoding)的“编码”部分。

  --qingstor-upload-concurrency
      多部分上传的并发。
      
      这是同一文件的多个块并行上传的数量。注意：如果将此设置为>1，则多部分上传的校验和将损坏（但上传本身不会损坏）。如果您通过高速链接上传少量大文件，并且这些上传未完全利用带宽，则增加这些上传可能有助于加速传输。

  --qingstor-env-auth
      从运行时获取青云凭证。
      
      仅适用于access_key_id和secret_access_key为空的情况。

      示例：
         | false | 在下一步中输入青云凭证。
         | true  | 从环境（env vars或IAM）获取青云凭证。

  --qingstor-access-key-id
      青云访问密钥ID。
      
      如果access_key_id和secret_access_key为空，则留空以获得匿名访问或运行时凭据。

  --qingstor-secret-access-key
      青云的Secret Access Key(密码)。
      
      如果access_key_id和secret_access_key为空，则留空以获得匿名访问或运行时凭据。

  --qingstor-zone
      连接的区域。
      
      默认值为"pek3a"。

      示例：
         | pek3a | 北京（中国）第三区域。
                 | 需要位置约束pek3a。
         | sh1a  | 上海（中国）第一区域。
                 | 需要位置约束sh1a。
         | gd2a  | 广东省第二个区域。
                 | 需要位置约束gd2a。

  --qingstor-connection-retries
      连接重试次数。

  --qingstor-upload-cutoff
      切换到分块上传的截止点。
      
      任何大于此值的文件将分块上传。最小值为0，最大值为5 GB。

  --qingstor-chunk-size
      用于上传的块大小。
      
      当上传大于`upload_cutoff`的文件时，将使用此块大小作为多部分上传。
      
      请注意，`--qingstor-upload-concurrency`每次传输会在内存中缓冲这个大小的块。如果您通过高速链接传输大文件并且有足够的内存，则增加这个大小会加速传输。


选项：
  --help，-h显示帮助

  数据准备选项

  --delete-after-export    [危险]导出CAR文件后删除数据集文件。 (default: false)
  --rescan-interval value  从上次成功扫描到现在经过指定的时间段后，自动重新扫描源目录（默认值：禁用）

  QingStor选项

  --qingstor-access-key-id value         青云的访问密钥ID。 [$QINGSTOR_ACCESS_KEY_ID]
  --qingstor-chunk-size value            用于上传的块大小。 (default: "4Mi") [$QINGSTOR_CHUNK_SIZE]
  --qingstor-connection-retries value   连接重试次数。 (default: "3") [$QINGSTOR_CONNECTION_RETRIES]
  --qingstor-encoding value              后端编码。 (default: "Slash,Ctl,InvalidUtf8") [$QINGSTOR_ENCODING]
  --qingstor-endpoint value              输入连接青云API的端点URL。 [$QINGSTOR_ENDPOINT]
  --qingstor-env-auth value              从运行时获取青云凭证。 (default: "false") [$QINGSTOR_ENV_AUTH]
  --qingstor-secret-access-key value     青云的Secret Access Key(密码)。 [$QINGSTOR_SECRET_ACCESS_KEY]
  --qingstor-upload-concurrency value    多部分上传的并发。 (default: "1") [$QINGSTOR_UPLOAD_CONCURRENCY]
  --qingstor-upload-cutoff value         切换到分块上传的截止点。 (default: "200Mi") [$QINGSTOR_UPLOAD_CUTOFF]
  --qingstor-zone value                  连接的区域。 [$QINGSTOR_ZONE]

```
{% endcode %}