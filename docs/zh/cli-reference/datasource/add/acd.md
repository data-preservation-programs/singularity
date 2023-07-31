# 亚马逊云盘

{% code fullWidth="true" %}
```
命令名:
   singularity datasource add acd - 亚马逊云盘

使用方法:
   singularity datasource add acd [options] <数据集名> <源路径>

说明:
   --acd-auth-url
      认证服务器URL。
      
      留空以使用提供商默认值。

   --acd-checkpoint
      内部轮询检查点（调试）。

   --acd-client-id
      OAuth客户端ID。
      
      通常留空。

   --acd-client-secret
      OAuth客户端密钥。
      
      通常留空。

   --acd-encoding
      后端的编码方式。
      
      请参阅[概览中的编码部分](/overview/#encoding)以获取更多信息。

   --acd-templink-threshold
      文件的大小大于等于此值将通过其tempLink下载。
      
      文件的大小大于等于此值将通过其“tempLink”下载。这是为了解决亚马逊云盘的一个问题，该问题会阻止大约10 GiB以上的文件下载。
      此参数的默认值为9 GiB，一般无需更改。
      
      若要下载超过此阈值的文件，rclone会请求一个“tempLink”，该链接通过底层S3存储直接下载文件。

   --acd-token
      OAuth访问令牌的JSON数据。

   --acd-token-url
      令牌服务器URL。
      
      留空以使用提供商默认值。

   --acd-upload-wait-per-gb
      每GiB附加等待时间，用于在完全上传失败后查看文件是否出现。
      
      有时，当文件完全上传后，亚马逊云盘会出现错误，但是文件最终会在一段时间后显示出来。
      对于大小超过1 GiB的文件，会出现此情况几乎每次都会发生在大于10 GiB的文件上。
      此参数控制rclone等待文件出现的时间。
      
      此参数的默认值为每GiB等待3分钟，因此默认情况下，每上传1 GiB文件，rclone会等待3分钟查看文件是否出现。
      
      您可以通过将其设置为0来禁用此功能。这可能会导致冲突错误，因为rclone会重试上传失败的文件，但是文件最终可能会正确显示。
      
      这些值是通过观察多个不同文件大小的大文件上传而经验性确定的。
      
      使用"-v"标志上传以查看rclone在此情况下的更多信息。


选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作] 在将数据集导出到CAR文件后删除数据集中的文件。 (默认: false)
   --rescan-interval value  当上次成功扫描后经过此时间间隔后，自动重新扫描源目录（默认: 禁用）
   --scanning-state value   设置初始的扫描状态 (默认: ready)

   acd选项

   --acd-auth-url value            认证服务器URL。 [$ACD_AUTH_URL]
   --acd-client-id value           OAuth客户端ID。 [$ACD_CLIENT_ID]
   --acd-client-secret value       OAuth客户端密钥。 [$ACD_CLIENT_SECRET]
   --acd-encoding value            后端的编码方式。 (默认: "Slash,InvalidUtf8,Dot") [$ACD_ENCODING]
   --acd-templink-threshold value  文件的大小大于等于此值将通过其tempLink下载。 (默认: "9Gi") [$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               OAuth访问令牌的JSON数据。 [$ACD_TOKEN]
   --acd-token-url value           令牌服务器URL。 [$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  每GiB附加等待时间，用于在完全上传失败后查看文件是否出现。 (默认: "3m0s") [$ACD_UPLOAD_WAIT_PER_GB]

```
{% endcode %}