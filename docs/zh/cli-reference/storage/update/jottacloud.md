# Jottacloud

{% code fullWidth="true" %}
```
名称:
   singularity storage update jottacloud - Jottacloud

用法:
   singularity storage update jottacloud [命令选项] <名称|id>

描述:
   --md5-memory-limit
      大于该大小的文件将被缓存在磁盘上以计算MD5。

   --trashed-only
      仅显示在回收站中的文件。
      
      这将按原始目录结构显示已被删除的文件。

   --hard-delete
      永久删除文件而不是将其放入回收站。

   --upload-resume-limit
      大于该大小的文件如果上传失败可以继续上传。

   --no-versions
      通过删除文件并重新创建文件来避免服务器端版本控制，而不是覆盖它们。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。


选项:
   --help, -h  显示帮助

   高级选项

   --encoding value             后端的编码方式。 (默认: "Slash, LtGt, DoubleQuote, Colon, Question, Asterisk, Pipe, Del, Ctl, InvalidUtf8, Dot") [$ENCODING]
   --hard-delete                永久删除文件而不是将其放入回收站。 (默认: false) [$HARD_DELETE]
   --md5-memory-limit value     大于该大小的文件将被缓存在磁盘上以计算MD5。 (默认: "10Mi") [$MD5_MEMORY_LIMIT]
   --no-versions                通过删除文件并重新创建文件来避免服务器端版本控制，而不是覆盖它们。 (默认: false) [$NO_VERSIONS]
   --trashed-only               仅显示在回收站中的文件。 (默认: false) [$TRASHED_ONLY]
   --upload-resume-limit value  大于该大小的文件如果上传失败可以继续上传。 (默认: "10Mi") [$UPLOAD_RESUME_LIMIT]

```
{% endcode %}