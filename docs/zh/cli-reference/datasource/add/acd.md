# Amazon Drive

{% code fullWidth="true" %}
```
名称：
   singularity datasource add acd - 亚马逊云盘

用法：
   singularity datasource add acd [命令选项] <dataset_name> <source_path>

描述：
   --acd-client-id
      OAuth 客户端 ID。
      
      通常情况下留空。

   --acd-client-secret
      OAuth 客户端密钥。
      
      通常情况下留空。

   --acd-token
      OAuth 访问令牌为 JSON 块。

   --acd-auth-url
      认证服务器 URL。
      
      留空以使用提供程序默认值。

   --acd-checkpoint
      用于内部轮询的检查点（调试）。

   --acd-token-url
      令牌服务器 URL。
      
      留空以使用提供程序默认值。

   --acd-upload-wait-per-gb
      完成上传失败后等待每个 GiB 的额外时间以查看是否出现。
      
      有时亚马逊云盘会在文件完全上传后给出错误，但文件在一段时间后会出现。对于文件大小超过1 GiB的文件，这种情况有时会发生，对于超过10 GiB的文件，几乎每次都会发生。此参数控制rclone等待文件出现的时间。
      
      此参数的默认值为每个 GiB 等待3分钟，因此默认情况下，它会等待每个 GiB 上传3分钟以查看文件是否出现。
      
      您可以通过将其设置为0来禁用此功能。这可能会导致冲突错误，因为rclone会重试失败的上传，但文件最终很可能会出现正确。
      
      这些值是通过观察各种文件大小的大文件上传而经验性确定的。
      
      使用“-v”标志上传以查看有关rclone在此情况下正在执行的更多信息。

   --acd-templink-threshold
      将以该大小或以上的文件通过它们的 tempLink 进行下载。
      
      该大小或以上的文件将通过其“tempLink”进行下载。这是为了解决Amazon Drive的一个问题，该问题会阻止下载大约10 GiB的文件。这个参数的默认值为9 GiB，不应该被更改。
      
      要下载超过此阈值的文件，rclone将请求“tempLink”，该“tempLink”通过一个临时 URL 从底层的S3存储直接下载文件。

   --acd-encoding
      后端的编码。
      
      请参见[概述中的编码部分](/overview/#encoding)获取更多信息。


选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险]在将数据集导出为 CAR 文件后删除数据集的文件。（默认值：false）
   --rescan-interval value  当从上次成功扫描以来经过此间隔时，自动重新扫描源目录（默认值：禁用）

   acd 的选项

   --acd-auth-url value            认证服务器 URL。[$ACD_AUTH_URL]
   --acd-client-id value           OAuth 客户端 ID。[$ACD_CLIENT_ID]
   --acd-client-secret value       OAuth 客户端密钥。[$ACD_CLIENT_SECRET]
   --acd-encoding value            后端的编码。（默认值：“Slash，InvalidUtf8，Dot”）[$ACD_ENCODING]
   --acd-templink-threshold value  将以该大小或以上的文件通过它们的 tempLink 进行下载。（默认值：“9Gi”）[$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               OAuth 访问令牌为 JSON 块。[$ACD_TOKEN]
   --acd-token-url value           令牌服务器 URL。[$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  完成上传失败后等待每个 GiB 的额外时间以查看是否出现。（默认值：“3m0s”）[$ACD_UPLOAD_WAIT_PER_GB]

```
{% endcode %}