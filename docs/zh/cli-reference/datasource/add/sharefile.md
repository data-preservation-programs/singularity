# Citrix Sharefile

{% code fullWidth="true" %}
```
名称：
   singularity datasource add sharefile - Citrix Sharefile

用法：
   singularity datasource add sharefile [命令选项] <数据集名称> <源路径>

描述：
   --sharefile-chunk-size
      上传块大小。
      
      必须是大于等于 256KB 的 2 的整数倍。
      
      增大块大小会提高性能，但请注意每个块都会占用一定的内存空间。
      
      减小块大小会减少内存使用，但会降低性能。

   --sharefile-encoding
      后端的编码格式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --sharefile-endpoint
      API 调用的终点。
      
      这通常是在 OAuth 过程中自动发现的，但也可以手动设置为类似于：https://XXX.sharefile.com 的值。
      

   --sharefile-root-folder-id
      根文件夹的 ID。
      
      留空以访问“个人文件夹”。您可以在此处使用其中一个标准值，也可以使用任何文件夹的 ID（长十六进制数字 ID）。

      示例：
         | <unset>    | 访问个人文件夹（默认）。
         | favorites  | 访问收藏夹。
         | allshared  | 访问所有共享文件夹。
         | connectors | 访问所有个别连接器。
         | top        | 访问主页、收藏夹、共享文件夹以及连接器。

   --sharefile-upload-cutoff
      切换到多部分上传的截止点。

选项：
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险操作] 导出数据集为 CAR 文件后删除数据集中的文件。 (默认值：false)
   --rescan-interval value  最后一次成功扫描后，自动重新扫描源目录的时间间隔 (默认值：禁用)
   --scanning-state value   设置初始扫描状态 (默认值：就绪)

   Sharefile 选项

   --sharefile-chunk-size value      上传块大小 (默认值： "64Mi") [$SHAREFILE_CHUNK_SIZE]
   --sharefile-encoding value        后端的编码格式 (默认值： "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot") [$SHAREFILE_ENCODING]
   --sharefile-endpoint value        API 调用的终点 [$SHAREFILE_ENDPOINT]
   --sharefile-root-folder-id value  根文件夹的 ID [$SHAREFILE_ROOT_FOLDER_ID]
   --sharefile-upload-cutoff value   切换到多部分上传的截止点 (默认值： "128Mi") [$SHAREFILE_UPLOAD_CUTOFF]

```
{% endcode %}