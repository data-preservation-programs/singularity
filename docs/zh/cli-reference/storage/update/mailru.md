# Mail.ru Cloud

{% code fullWidth="true" %}
```
名称：
   singularity storage update mailru - Mail.ru Cloud

用法：
   singularity storage update mailru [命令选项] <名称|id>

描述：
   --user
      用户名（通常是电子邮件）。

   --pass
      密码。
      
      这必须是一个应用密码- rclone将无法使用您的普通密码。请参阅文档中的配置部分以了解如何创建应用密码。

   --speedup-enable
      如果存在具有相同数据哈希的其他文件，则跳过完整上传。
      
      此功能称为“加速”或“哈希上载”。对于通常可用的文件（如常见的书籍、视频或音频剪辑），这是特别高效的，因为文件通过哈希在所有mailru用户的所有帐户中搜索。
      如果源文件是唯一的或加密的，则此功能是没有意义和无效的。
      请注意，rclone可能需要本地内存和磁盘空间提前计算内容哈希并决定是否需要进行完整的上传。
      此外，在rclone无法提前知道文件大小的情况下（例如，在流式传输或部分上传的情况下），它甚至不会尝试进行此优化。

      示例：
         | true  | 启用
         | false | 禁用

   --speedup-file-patterns
      允许加速（哈希上载）的文件名模式的逗号分隔列表。
      
      模式不区分大小写，可以包含'*'或'?'元字符。

      示例：
         | <unset>                 | 空列表完全禁用加速（哈希上载）。
         | *                       | 将尝试加速所有文件。
         | *.mkv,*.avi,*.mp4,*.mp3 | 尝试加速常见的音频/视频文件。
         | *.zip,*.gz,*.rar,*.pdf  | 尝试加速常见的归档文件或PDF书籍。

   --speedup-max-disk
      此选项允许您禁用大文件的加速（哈希上载）。
      
      原因是预先计算的哈希可能会耗尽您的RAM或磁盘空间。

      示例：
         | 0  | 完全禁用加速（哈希上载）。
         | 1G | 大于1GB的文件将直接上传。
         | 3G | 如果本地磁盘上剩余空间小于3GB，请选择此选项。

   --speedup-max-memory
      大于下面给出的大小的文件将始终在磁盘上进行哈希计算。

      示例：
         | 0    | 始终在临时磁盘位置上进行哈希计算。
         | 32M  | 磁盘上的哈希计算不应占用超过32MB的内存。
         | 256M | 用于哈希计算的RAM最多可用256MB。

   --check-hash
      如果文件的校验和不匹配或无效，副本应该怎么办。

      示例：
         | true  | 报错并失败。
         | false | 忽略并继续。

   --user-agent
      客户端内部使用的HTTP用户代理。
      
      默认为“rclone/VERSION”或命令行中提供的“--user-agent”。

   --quirks
      内部维护标志的逗号分隔列表。
      
      普通用户不可以使用此选项。它仅用于方便远程解决问题。标志的严格含义未有文档记录并且不能保证在发布间保持一致。
      后端稳定后，这些标志将被移除。
      支持的标志：atomicmkdir binlist unknowndirs

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。


选项：
   --help, -h        显示帮助
   --pass value      密码。[$PASS]
   --speedup-enable  如果存在具有相同数据哈希的其他文件，则跳过完整上传。 (default: true) [$SPEEDUP_ENABLE]
   --user value      用户名（通常是电子邮件）。[$USER]

   高级选项

   --check-hash                   如果文件的校验和不匹配或无效，副本应该怎么办。 (default: true) [$CHECK_HASH]
   --encoding value               后端的编码方式。 (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --quirks value                 内部维护标志的逗号分隔列表。[$QUIRKS]
   --speedup-file-patterns value  允许加速（哈希上载）的文件名模式的逗号分隔列表。 (default: "*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf") [$SPEEDUP_FILE_PATTERNS]
   --speedup-max-disk value       此选项允许您禁用大文件的加速（哈希上载）。 (default: "3Gi") [$SPEEDUP_MAX_DISK]
   --speedup-max-memory value     大于下面给出的大小的文件将始终在磁盘上进行哈希计算。 (default: "32Mi") [$SPEEDUP_MAX_MEMORY]
   --user-agent value             客户端内部使用的HTTP用户代理。[$USER_AGENT]

```
{% endcode %}