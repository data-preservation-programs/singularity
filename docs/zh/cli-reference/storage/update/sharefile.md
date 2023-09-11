# Citrix Sharefile

{% code fullWidth="true" %}
```
名称：
   singularity storage update sharefile - Citrix Sharefile

用法：
   singularity storage update sharefile [命令选项] <名称|ID>

介绍：
   --upload-cutoff
      切换到多部分上传的截止点。

   --root-folder-id
      根文件夹的ID。
      
      留空以访问“个人文件夹”。您可以在此处使用标准值或任何文件夹ID（长十六进制数字ID）。

      例如：
         | <unset>    | 访问个人文件夹（默认）。
         | favorites  | 访问“收藏夹”文件夹。
         | allshared  | 访问所有共享文件夹。
         | connectors | 访问所有个别连接器。
         | top        | 访问主页、收藏夹、共享文件夹以及连接器。

   --chunk-size
      上传块大小。
      
      必须是大于等于256k的2的乘幂。
      
      增大这个值将提高性能，但请注意每个块都会在内存中缓冲一次传输。
      
      减小这个值将减少内存使用量，但降低性能。

   --endpoint
      API调用的端点。
      
      这通常在oauth过程中自动发现，但可以手动设置为如：https://XXX.sharefile.com

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

选项：
   --help, -h              显示帮助信息
   --root-folder-id value  根文件夹的ID。[$ROOT_FOLDER_ID]

   高级选项

   --chunk-size value     上传块大小。 (默认值："64Mi") [$CHUNK_SIZE]
   --encoding value       后端的编码方式。 (默认值："Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value       API调用的端点。[$ENDPOINT]
   --upload-cutoff value  切换到多部分上传的截止点。 (默认值："128Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}