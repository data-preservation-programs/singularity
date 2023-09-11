# Jottacloud

{% code fullWidth="true" %}
```
名称：
   singularity storage create jottacloud - Jottacloud

使用方法：
   singularity storage create jottacloud [命令选项] [参数...]

描述：
   --md5-memory-limit
      文件大小超过此值时，将在磁盘上缓存以计算MD5。

   --trashed-only
      仅显示在回收站中的文件。
      
      这将按原始目录结构显示被回收的文件。

   --hard-delete
      永久删除文件，而不是将其放入回收站。

   --upload-resume-limit
      大于此值的文件，如果上传失败，则可以进行断点续传。

   --no-versions
      避免服务器端版本控制，通过删除文件并重新创建文件而不是覆写文件来实现。

   --encoding
      后端编码。
      
      详情请见[概览中的编码部分](/overview/#encoding)。

选项：
   --help, -h  显示帮助

   高级选项

   --encoding value             后端编码（默认："Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot"）[$ENCODING]
   --hard-delete                永久删除文件，而不是将其放入回收站（默认：false）[$HARD_DELETE]
   --md5-memory-limit value     文件大小超过此值时，将在磁盘上缓存以计算MD5（默认："10Mi"）[$MD5_MEMORY_LIMIT]
   --no-versions                避免服务器端版本控制，通过删除文件并重新创建文件而不是覆写文件来实现（默认：false）[$NO_VERSIONS]
   --trashed-only               仅显示在回收站中的文件（默认：false）[$TRASHED_ONLY]
   --upload-resume-limit value  大于此值的文件，如果上传失败，则可以进行断点续传（默认："10Mi"）[$UPLOAD_RESUME_LIMIT]

   通用选项

   --name value  存储的名称（默认：自动生成）
   --path value  存储的路径

```
{% endcode %}