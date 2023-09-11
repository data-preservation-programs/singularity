# Citrix Sharefile

{% code fullWidth="true" %}
```
名称：
   singularity storage create sharefile - Citrix Sharefile

用法：
   singularity storage create sharefile [命令选项] [参数...]

描述：
   --upload-cutoff
      切换到多部分上传的截断点。

   --root-folder-id
      根目录的ID。
      
      留空以访问“个人文件夹”。 您可以在此处使用其中一个标准值或任何文件夹ID（长十六进制数字ID）。

      示例:
        | <unset>  | 访问个人文件夹（默认）。
        | favorites| 访问收藏夹。
        | allshared | 访问所有共享文件夹。
        | connectors | 访问所有单独的连接器。
        | top       | 访问主页，收藏夹，共享文件夹以及连接器。

   --chunk-size
      上传块大小。
      
      必须是2的幂并大于等于256k。
      
      增大此值将提高性能，但请注意每个块都会在传输期间缓冲在内存中。
      
      减小此值将减少内存使用但降低性能。

   --endpoint
      API调用的终端点。
      
      这通常是作为OAuth过程的一部分自动发现的，但也可以手动设置为类似于：https://XXX.sharefile.com 

   --encoding
      后端的编码。
      
      有关详细信息，请参阅[概览中的编码部分](/overview/#encoding)。


选项：
   --help, -h            显示帮助
   --root-folder-id value 根目录的ID [$ROOT_FOLDER_ID]

   高级功能

   --chunk-size value     上传块大小（默认值：“64Mi”） [$CHUNK_SIZE]
   --encoding value       后端的编码（默认值：“Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot”） [$ENCODING]
   --endpoint value       API调用的终端点 [$ENDPOINT]
   --upload-cutoff value  切换到多部分上传的截断点（默认值：“128Mi”） [$UPLOAD_CUTOFF]

   一般

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}