# Jottacloud

{% code fullWidth="true" %}
```
命令名称:
   singularity datasource add jottacloud - Jottacloud

使用方法:
   singularity datasource add jottacloud [命令选项] <数据集名称> <源路径>

描述:
   --jottacloud-encoding
      后端的编码方式。
      
      参见[概述中的编码章节](/overview/#encoding)获取更多信息。

   --jottacloud-hard-delete
      永久删除文件，而不是将其放入回收站。

   --jottacloud-md5-memory-limit
      如果需要，大于此大小的文件将被缓存到磁盘以计算MD5。

   --jottacloud-no-versions
      避免服务器端版本控制，通过删除文件并重新创建文件来覆盖文件。

   --jottacloud-trashed-only
      仅显示在回收站中的文件。
      
      这将以它们原始的目录结构显示回收站中的文件。

   --jottacloud-upload-resume-limit
      如果上传失败，大于此大小的文件可以进行断点续传。


选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export   [危险操作] 导出数据集到CAR文件后删除数据集中的文件。 (默认值: false)
   --rescan-interval value 当离上次成功扫描间隔超过此时间时，自动重新扫描源目录（默认值: 禁用）
   --scanning-state value  设置初始的扫描状态（默认值: 就绪）

   Jottacloud选项

   --jottacloud-encoding value             后端的编码方式。 (默认值: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot") [$JOTTACLOUD_ENCODING]
   --jottacloud-hard-delete value          永久删除文件，而不是将其放入回收站。 (默认值: "false") [$JOTTACLOUD_HARD_DELETE]
   --jottacloud-md5-memory-limit value     如果需要，大于此大小的文件将被缓存到磁盘以计算MD5。 (默认值: "10Mi") [$JOTTACLOUD_MD5_MEMORY_LIMIT]
   --jottacloud-no-versions value          避免服务器端版本控制，通过删除文件并重新创建文件来覆盖文件。 (默认值: "false") [$JOTTACLOUD_NO_VERSIONS]
   --jottacloud-trashed-only value         仅显示在回收站中的文件。 (默认值: "false") [$JOTTACLOUD_TRASHED_ONLY]
   --jottacloud-upload-resume-limit value  如果上传失败，大于此大小的文件可以进行断点续传。 (默认值: "10Mi") [$JOTTACLOUD_UPLOAD_RESUME_LIMIT]

```
{% endcode %}