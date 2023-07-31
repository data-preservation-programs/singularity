# Mail.ru Cloud

{% code fullWidth="true" %}
```
名称:
   singularity datasource add mailru - Mail.ru云

用法:
   singularity datasource add mailru [命令选项] <数据集名称> <源路径>

描述:
   --mailru-check-hash
      当文件校验和不匹配或无效时，复制操作应该怎么处理。

      示例:
         | true  | 失败并抛出错误。
         | false | 忽略并继续。

   --mailru-encoding
      后端的编码方式。

      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --mailru-pass
      密码。

      它必须是一个应用程序密码，rclone不会使用您的普通密码。请查看文档中的配置部分，了解如何创建应用程序密码。

   --mailru-quirks
      逗号分隔的内部维护标志列表。

      普通用户不应使用此选项。它仅用于远程处理后端问题。标志的具体含义没有文档记录，并不能保证在版本之间保持一致。当后端变得稳定时，这些标志将被移除。
      支持的标志：atomicmkdir binlist unknowndirs

   --mailru-speedup-enable
      如果存在另一个具有相同数据哈希的文件，则跳过完整上传。

      此功能称为“speedup”或“put by hash”。在普通可用文件（如热门图书、视频或音频剪辑）中，该功能尤其高效，因为文件将根据哈希在所有Mail.ru用户的所有帐户中进行搜索。
      如果源文件是唯一的或加密的，这个功能毫无意义且无效。请注意，rclone可能需要本地内存和磁盘空间来预先计算内容哈希并决定是否需要完整上传。
      另外，如果rclone无法提前知道文件的大小（例如流式传输或部分上传），它甚至不会尝试此优化。

      示例:
         | true  | 启用
         | false | 禁用

   --mailru-speedup-file-patterns
      逗号分隔的文件名模式列表，用于“speedup”（按哈希方式上传）。
      
      模式不区分大小写，可以包含“*”或“?”元字符。

      示例:
         | <unset>                 | 空列表完全禁用“speedup”（按哈希方式上传）。
         | *                       | 所有文件都将尝试进行“speedup”。
         | *.mkv,*.avi,*.mp4,*.mp3 | 仅尝试对常见的音视频文件进行“speedup”。
         | *.zip,*.gz,*.rar,*.pdf  | 仅尝试对常见的存档文件或PDF图书进行“speedup”。

   --mailru-speedup-max-disk
      此选项允许您在大文件上禁用“speedup”（按哈希方式上传）。
      
      原因是预先计算哈希可能会耗尽您的RAM或磁盘空间。

      示例:
         | 0  | 完全禁用“speedup”（按哈希方式上传）。
         | 1G | 较大于1GB的文件将直接上传。
         | 3G | 如果本地磁盘上有小于3GB可用空间，请选择此选项。

   --mailru-speedup-max-memory
      大于给定大小的文件将始终在磁盘上进行哈希计算。

      示例:
         | 0    | 预先计算的哈希值将始终在临时磁盘位置上完成。
         | 32M  | 为预先计算的哈希值，不要超过32MB的RAM。
         | 256M | 你最多有256MB的可用RAM用于哈希计算。

   --mailru-user
      用户名（通常是电子邮件）。

   --mailru-user-agent
      客户端内部使用的HTTP用户代理。

      默认为"rclone/VERSION"，或者通过命令行提供"--user-agent"参数。

选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险操作] 在导出为CAR文件后删除数据集的文件。 (默认值：false)
   --rescan-interval value  当距离上次成功扫描已经过去指定时间间隔时，自动重新扫描源目录（默认值：禁用）
   --scanning-state value   设置初始扫描状态（默认值：准备就绪）

   Mail.ru的选项

   --mailru-check-hash value             当文件校验和不匹配或无效时，复制操作应该怎么处理。（默认值："true"）[$MAILRU_CHECK_HASH]
   --mailru-encoding value               后端的编码方式。（默认值："Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot"）[$MAILRU_ENCODING]
   --mailru-pass value                   密码。[$MAILRU_PASS]
   --mailru-speedup-enable value         如果存在另一个具有相同数据哈希的文件，则跳过完整上传。（默认值："true"）[$MAILRU_SPEEDUP_ENABLE]
   --mailru-speedup-file-patterns value  逗号分隔的文件名模式列表，用于“speedup”（按哈希方式上传）。 （默认值："*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf"）[$MAILRU_SPEEDUP_FILE_PATTERNS]
   --mailru-speedup-max-disk value       此选项允许您在大文件上禁用“speedup”（按哈希方式上传）。 （默认值："3Gi"）[$MAILRU_SPEEDUP_MAX_DISK]
   --mailru-speedup-max-memory value     大于给定大小的文件将始终在磁盘上进行哈希计算。（默认值："32Mi"）[$MAILRU_SPEEDUP_MAX_MEMORY]
   --mailru-user value                   用户名（通常是电子邮件）。[$MAILRU_USER]

```
{% endcode %}