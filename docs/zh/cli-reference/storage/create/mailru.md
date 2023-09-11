#  Mail.ru云

{% code fullWidth="true" %}
```
命令:
   singularity storage create mailru - Mail.ru云

用法:
   singularity storage create mailru [命令选项] [参数...]

说明:
   --user
      用户名（通常是邮箱）。

   --pass
      密码。
      
      这必须是一个应用程序密码- rclone无法使用您的常规密码。请参阅文档中的配置部分，了解如何生成应用程序密码。

   --speedup-enable
      如果存在具有相同数据哈希的其他文件，则跳过完整上传。
      
      此功能称为“加速”或“按哈希放置”。对于常见文件（如热门书籍、视频或音频剪辑），它特别高效，因为在所有mailru用户的所有帐户中按哈希搜索文件。如果源文件是唯一的或加密的，则该功能无意义且低效。请注意，rclone可能需要本地内存和磁盘空间来提前计算内容哈希并决定是否需要进行完整上传。另外，如果rclone无法提前获取文件大小（例如流式传输或部分上传），它甚至不会尝试此优化。

      示例:
         | true  | 启用
         | false | 禁用

   --speedup-file-patterns
      适用于加速（按哈希放置）的文件名模式的逗号分隔列表。
      
      模式不区分大小写，可以包含 '*' 或 '?' 元字符。

      示例:
         | <未设置>                | 空列表会完全禁用加速（按哈希放置）。
         | *                      | 所有文件都会尝试加速。
         | *.mkv,*.avi,*.mp4,*.mp3 | 仅尝试常见的音频/视频文件进行按哈希放置。
         | *.zip,*.gz,*.rar,*.pdf  | 仅尝试常见的存档文件或PDF书籍进行加速。

   --speedup-max-disk
      此选项允许您禁用较大文件的加速（按哈希放置）。
      
      原因是预处理哈希可能耗尽您的RAM或磁盘空间。

      示例:
         | 0   | 完全禁用加速（按哈希放置）。
         | 1G  | 文件大于1Gb将直接上传。
         | 3G  | 如果本地磁盘上的可用空间小于3Gb，请选择此选项。

   --speedup-max-memory
      大于下面给定大小的文件将始终在磁盘上进行哈希处理。

      示例:
         | 0    | 预处理哈希将始终在临时磁盘位置上执行。
         | 32M  | 不要为预处理哈希分配超过32Mb的RAM。
         | 256M | 最多可以使用256Mb的空闲RAM进行哈希计算。

   --check-hash
      如果文件校验和不匹配或无效，复制应该如何处理。

      示例:
         | true  | 失败并显示错误。
         | false | 忽略并继续。

   --user-agent
      客户端内部使用的HTTP用户代理。
      
      默认为“rclone / VERSION”，或者使用命令行中提供的“--user-agent”。

   --quirks
      内部维护标志的逗号分隔列表。
      
      普通用户不应使用此选项。它仅用于便于远程排除后端问题。标志的严格含义未记录，也不能保证在版本之间保持一致。后端稳定后将删除此选项。支持的quirks: atomicmkdir binlist unknowndirs

   --encoding
      后端的编码方式。
      
      有关详细信息，请参见概述中的[编码部分](/overview/#encoding)。

选项:
   --help, -h        显示帮助
   --pass value      密码。[$ PASS]
   --speedup-enable  如果存在具有相同数据哈希的其他文件，则跳过完整上传。 (默认值: true) [$ SPEEDUP_ENABLE]
   --user value      用户名（通常是邮箱）。[$ USER]

高级选项

   --check-hash                   如果文件校验和不匹配或无效，复制应该如何处理。 (默认值: true) [$ CHECK_HASH]
   --encoding value               后端的编码方式。 (默认值: "Slash, LtGt, DoubleQuote, Colon, Question, Asterisk, Pipe, BackSlash, Del, Ctl, InvalidUtf8, Dot") [$ ENCODING]
   --quirks value                 内部维护标志的逗号分隔列表。[$ QUIRKS]
   --speedup-file-patterns value 适用于加速（按哈希放置）的文件名模式的逗号分隔列表。 (默认值: "*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf") [$ SPEEDUP_FILE_PATTERNS]
   --speedup-max-disk value       此选项允许您禁用较大文件的加速（按哈希放置）。 (默认值: "3Gi") [$ SPEEDUP_MAX_DISK]
   --speedup-max-memory value     大于下面给定大小的文件将始终在磁盘上进行哈希处理。 (默认值: "32Mi") [$ SPEEDUP_MAX_MEMORY]
   --user-agent value             客户端内部使用的HTTP用户代理。[$ USER_AGENT]

通用选项

   --name value  存储的名称（默认值: 自动生成）
   --path value  存储的路径

```
{% endcode %}