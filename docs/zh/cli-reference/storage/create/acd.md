# 亚马逊云盘

{% code fullWidth="true" %}
```
名称:
   singularity storage create acd - 亚马逊云盘

用法:
   singularity storage create acd [选项] [参数...]

描述:
   --client-id
      OAuth 客户端 ID。
      
      一般情况下不需要填写。

   --client-secret
      OAuth 客户端密钥。
      
      一般情况下不需要填写。

   --token
      OAuth 访问令牌以 JSON 形式提供。

   --auth-url
      鉴权服务器 URL。
      
      如果要使用提供器的默认值，请留空。

   --token-url
      令牌服务器 URL。
      
      如果要使用提供器的默认值，请留空。

   --checkpoint
      内部轮询的检查点（调试用）。

   --upload-wait-per-gb
      在上传完整文件失败后，每 GiB 额外等待的时间，以查看文件是否出现。
      
      有时，亚马逊云盘会在文件完全上传后报错，但文件稍后仍然会出现。对于大于 1 GiB 大小的文件，这种情况有时会发生，而对于大于 10 GiB 的文件，几乎每次都会发生。此参数控制 rclone 等待文件出现的时间。
      
      此参数的默认值为每 GiB 等待 3 分钟，因此默认情况下，rclone 会等待每上传一个 GiB 的时间来查看文件是否出现。
      
      您可以将此功能禁用，方法是将其设置为 0。这可能会导致冲突错误，因为 rclone 会重试失败的上传，但文件最终可能会正确出现。
      
      这些值是通过观察大量不同文件大小的上传过程经验性确定的。
      
      若要查看有关 rclone 在此情况下的更多信息，请使用 "-v" 标志进行上传。

   --templink-threshold
      将下载文件大小大于此值的文件通过临时链接下载。
      
      大于此值的文件将通过其 "tempLink" 进行下载。这是为了解决亚马逊云盘阻止下载大于约 10 GiB 大小文件的问题。默认值为 9 GiB，一般无需更改。
      
      要下载超过此阈值的文件，rclone 请求一个 "tempLink"，该链接通过从底层的 S3 存储直接下载文件。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/#encoding)。

选项:
   --client-id value      OAuth 客户端 ID。[$CLIENT_ID]
   --client-secret value  OAuth 客户端密钥。[$CLIENT_SECRET]
   --help, -h             查看帮助

   高级选项

   --auth-url value            鉴权服务器 URL。[$AUTH_URL]
   --checkpoint value          内部轮询的检查点（调试用）。[$CHECKPOINT]
   --encoding value            后端的编码方式。 (默认值: "斜杠、无效的 UTF-8、点")[$ENCODING]
   --templink-threshold value  将下载文件大小大于此值的文件通过临时链接下载。 (默认值: "9Gi")[$TEMPLINK_THRESHOLD]
   --token value               OAuth 访问令牌以 JSON 形式提供。[$TOKEN]
   --token-url value           令牌服务器 URL。[$TOKEN_URL]
   --upload-wait-per-gb value  在上传完整文件失败后，每 GiB 额外等待的时间，以查看文件是否出现。 (默认值: "3m0s")[$UPLOAD_WAIT_PER_GB]

   常规选项

   --name value  存储的名称（默认值：自动生成的）
   --path value  存储的路径

```
{% endcode %}