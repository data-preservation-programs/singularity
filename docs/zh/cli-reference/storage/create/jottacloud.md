# Jottacloud

{% code fullWidth="true" %}
```bash
命令名:
   singularity storage create jottacloud - Jottacloud

用法:
   singularity storage create jottacloud [命令选项] [参数...]

描述:
   --md5-memory-limit
      如果需要计算MD5，大于此大小的文件将在磁盘上缓存。

   --trashed-only
      仅显示在回收站中的文件。
      
      这将按其原始目录结构显示已删除的文件。

   --hard-delete
      永久删除文件，而不是将它们放入回收站。

   --upload-resume-limit
      大于此大小的文件在上传失败时可以恢复续传。

   --no-versions
      通过删除文件并重新创建文件而不是覆盖文件，避免服务器端版本控制。

   --encoding
      后端编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#编码)。


选项:
   --help, -h  显示帮助信息

   高级选项

   --encoding value             后端编码方式。 (默认值: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete                永久删除文件，而不是将它们放入回收站。 (默认值: false) [$HARD_DELETE]
   --md5-memory-limit value     如果需要计算MD5，大于此大小的文件将在磁盘上缓存。 (默认值: "10Mi") [$MD5_MEMORY_LIMIT]
   --no-versions                通过删除文件并重新创建文件而不是覆盖文件，避免服务器端版本控制。 (默认值: false) [$NO_VERSIONS]
   --trashed-only               仅显示在回收站中的文件。 (默认值: false) [$TRASHED_ONLY]
   --upload-resume-limit value  大于此大小的文件在上传失败时可以恢复续传。 (默认值: "10Mi") [$UPLOAD_RESUME_LIMIT]

   通用选项

   --name value  存储的名称 (默认值: 自动生成)
   --path value  存储的路径

```
{% endcode %}