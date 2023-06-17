# Citrix Sharefile

{% code fullWidth="true" %}
```
名称：
   singularity 数据源添加 sharefile - Citrix Sharefile

用法：
   singularity datasource add sharefile [command options] <dataset_name> <source_path>

描述：
   --sharefile-upload-cutoff
      切换到多部分上传的截止点。

   --sharefile-root-folder-id
      根文件夹的 ID。
      
      留空以访问“个人文件夹”。您可以在此处使用标准值或任何文件夹 ID（长十六进制数字 ID）。

      示例：
         | <unset>    | 访问个人文件夹（默认）。
         | favorites  | 访问“收藏夹”文件夹。
         | allshared  | 访问所有共享文件夹。
         | connectors | 访问所有个别连接器。
         | top        | 访问主页、收藏夹、共享文件夹以及连接器。

   --sharefile-chunk-size
      上传块的大小。
      
      必须是 2 的幂 >= 256k。
      
      将其设置为更大将提高性能，但请注意每个块都会在传输期间缓冲。
      
      将其设置更小将减少内存使用量但会降低性能。

   --sharefile-endpoint
      API 调用的端点。
      
      这通常是在 oauth 过程中自动发现的，但也可以手动设置为类似于：https://XXX.sharefile.com

   --sharefile-encoding
      后端的编码方式。 
      
      更多信息请参见[“概览”中的编码部分](/overview/#encoding)。

选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作] 导出 CAR 文件后删除数据集的文件。  (默认值：false)
   --rescan-interval value  当从上次成功扫描后删除此间隔时，将自动重新扫描源目录 (默认值：未启用)

   sharefile 的选项

   --sharefile-chunk-size value      上传块的大小。 (默认值： "64Mi") [$SHAREFILE_CHUNK_SIZE]
   --sharefile-encoding value        后端的编码方式。 (默认值： "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot") [$SHAREFILE_ENCODING]
   --sharefile-endpoint value        API 调用的端点。 [$SHAREFILE_ENDPOINT]
   --sharefile-root-folder-id value  根文件夹的 ID。 [$SHAREFILE_ROOT_FOLDER_ID]
   --sharefile-upload-cutoff value   切换到多部分上传的截止点。 (默认值： "128Mi") [$SHAREFILE_UPLOAD_CUTOFF]
```
{% endcode %}