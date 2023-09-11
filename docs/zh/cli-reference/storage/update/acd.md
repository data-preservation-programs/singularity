# 亚马逊云盘

{% code fullWidth="true" %}
```
名称:
   singularity storage update acd - 亚马逊云盘

用法:
   singularity storage update acd [命令选项] <name|id>

描述:
   --client-id
      OAuth 客户端 ID。
      
      通常留空。

   --client-secret
      OAuth 客户端密钥。
      
      通常留空。

   --token
      作为 JSON 数据块的 OAuth 访问令牌。

   --auth-url
      认证服务器 URL。
      
      留空以使用提供者的默认值。

   --token-url
      令牌服务器 URL。
      
      留空以使用提供者的默认值。

   --checkpoint
      内部轮询的检查点（调试）。

   --upload-wait-per-gb
      完成上传失败后，每 GB 额外等待的时间，以查看是否出现完整文件。
      
      有时候，当一个文件已经完全上传时，亚马逊云盘会返回一个错误，但是文件在一段时间后仍然显示。对于大小超过 1GB 的文件，以及几乎每个大小超过 10GB 的文件，这种情况经常发生。此参数控制了 rclone 等待文件出现的时间。
      
      此参数的默认值为每 GB 3 分钟，因此默认情况下，rclone 将等待每上传一个 GB 3 分钟，以查看文件是否出现。
      
      您可以通过将此特性设置为 0 来禁用它。这可能会导致冲突错误，因为 rclone 会重试上传失败的文件，但是文件最终通常会正确显示。
      
      这些值是通过观察一系列不同文件大小的大文件上传而经验性地确定的。
      
      使用 "-v" 标志上传以查看 rclone 在此情况下的更多信息。

   --templink-threshold
      大小大于等于此值的文件将通过其 tempLink 下载。
      
      大小等于或大于此值的文件将通过其 "tempLink" 下载。这是为了解决亚马逊云盘的问题，该问题会阻止下载大约 10GB 或更大的文件。此参数的默认值为 9GB，通常不需要更改。
      
      要下载超过此阈值的文件，rclone 请求一个 "tempLink"，该链接会通过临时的 URL 直接从底层 S3 存储下载文件。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概览中的编码部分](/overview/#encoding)。


选项:
   --client-id value      OAuth 客户端 ID。[$CLIENT_ID]
   --client-secret value  OAuth 客户端密钥。[$CLIENT_SECRET]
   --help, -h             显示帮助信息

   高级选项

   --auth-url value            认证服务器 URL。[$AUTH_URL]
   --checkpoint value          内部轮询的检查点（调试）。[$CHECKPOINT]
   --encoding value            后端的编码方式。（默认值："Slash,InvalidUtf8,Dot"）[$ENCODING]
   --templink-threshold value  大小大于等于此值的文件将通过其 tempLink 下载。（默认值："9Gi"）[$TEMPLINK_THRESHOLD]
   --token value               作为 JSON 数据块的 OAuth 访问令牌。[$TOKEN]
   --token-url value           令牌服务器 URL。[$TOKEN_URL]
   --upload-wait-per-gb value  完成上传失败后，每 GB 额外等待的时间，以查看是否出现完整文件。（默认值："3m0s"）[$UPLOAD_WAIT_PER_GB]

```
{% endcode %}