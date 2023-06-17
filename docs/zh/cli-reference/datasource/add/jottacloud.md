# Jottacloud

{% code fullWidth="true" %}
```
名称：
   singularity DataSource add jottacloud - Jottacloud

用法：
   singularity DataSource add jottacloud [command options] <dataset_name> <source_path>

说明：
   --jottacloud-upload-resume-limit
      如果上传失败，大小大于此文件的文件可以恢复上传。

   --jottacloud-no-versions
      删除文件并重新创建文件，而不是覆盖它们，以避免服务器端版本控制。

   --jottacloud-encoding
      后端编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --jottacloud-md5-memory-limit
      如果需要，大于此文件大小的文件将被缓存在磁盘上以计算MD5。

   --jottacloud-trashed-only
      仅显示在垃圾桶中的文件。
      
      这将在其原始目录结构中显示已删除的文件。

   --jottacloud-hard-delete
      永久删除文件而不是将它们放入垃圾桶。

选项：
   --help，-h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险]将数据集导出为CAR文件后，删除其中的文件。 (默认值：false)
   --rescan-interval value  当上一次成功扫描后经过此间隔时间时，自动重新扫描源目录 (默认：禁用)

   Jottacloud选项

   --jottacloud-encoding value             后端编码。 (默认值："Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot") [$JOTTACLOUD_ENCODING]
   --jottacloud-hard-delete value          永久删除文件而不是将它们放入垃圾桶。 (默认值："false") [$JOTTACLOUD_HARD_DELETE]
   --jottacloud-md5-memory-limit value     如果需要，大于此文件大小的文件将被缓存在磁盘上以计算MD5。 (默认值："10Mi") [$JOTTACLOUD_MD5_MEMORY_LIMIT]
   --jottacloud-no-versions value          删除文件并重新创建文件，而不是覆盖它们，以避免服务器端版本控制。 (默认值："false") [$JOTTACLOUD_NO_VERSIONS]
   --jottacloud-trashed-only value         仅显示在垃圾桶中的文件。 (默认值："false") [$JOTTACLOUD_TRASHED_ONLY]
   --jottacloud-upload-resume-limit value  如果上传失败，大小大于此文件的文件可以恢复上传。 (默认值："10Mi") [$JOTTACLOUD_UPLOAD_RESUME_LIMIT]

```
{% endcode %}