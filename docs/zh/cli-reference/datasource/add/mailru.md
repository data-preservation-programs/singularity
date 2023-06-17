# Mail.ru Cloud

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add mailru - Mail.ru Cloud

USAGE:
   singularity datasource add mailru [command options] <dataset_name> <source_path>

DESCRIPTION:
   --mailru-quirks
      内部维护标志的逗号分隔列表。

      该选项不应由普通用户使用，仅为远程故障排除而设计。标志的严格含义未记录，
      并不能保证在版本之间持久存在。在后端变得稳定之后，这些标志将被删除。
      支持的标志：atomicmkdirBinlist unknowndirs

   --mailru-user
      用户名（通常为电子邮件）。

   --mailru-pass
      应用程序密码。不应使用普通密码。

      您需要生成并使用应用程序密码。请查看文档中的配置部分了解如何生成应用程序密码。

   --mailru-speedup-enable
      如果存在另一个具有相同数据哈希的文件，将跳过完全上传。

      此功能称为“加速”或“哈希修改上传”。当处理普遍可用的文件时（如流行的书籍、视频或音频剪辑），
      受益尤其明显，因为文件在所有 mailru 用户的所有帐户中通过哈希搜索。
      如果源文件是唯一的或已加密，则此功能是没意义和无效的。
      请注意，rclone 可能需要在本地内存和磁盘上计算内容哈希，
      并决定是否需要完全上传。另外，如果 rclone 不知道先前文件的大小（例如，对于流式传输或部分上传），
      它根本不会尝试此优化。

      示例：
         | true  | 启用
         | false | 禁用

   --mailru-speedup-max-disk
      该选项允许您为大文件禁用加速（哈希修改上传）。

      原因是预处理哈希可能会消耗您的 RAM 或磁盘空间。

      示例：
         | 0  | 完全禁用加速（哈希修改上传）。
         | 1G | 大于 1GB 的文件将直接上传。
         | 3G | 如果您的计算机上有少于 3GB 的可用空间，请选择此选项。

   --mailru-speedup-max-memory
      大小超过下面给出的大小的文件将始终在磁盘上进行哈希计算。

      示例：
         | 0    | 哈希计算将始终在临时磁盘位置上完成。
         | 32M  | 不要分配超过 32Mb 的 RAM 用于哈希计算。
         | 256M | 您最多可以为哈希计算分配 256MB 的可用 RAM。

   --mailru-user-agent
      客户端内部使用的 HTTP 用户代理。

      默认为“rclone/VERSION”，或者用于命令行的“--user-agent”。

   --mailru-speedup-file-patterns
      逗号分隔的文件名模式列表，适用于加速（哈希修改上传）。

      模式不区分大小写，可以包含 '*' 或 '?' 特殊字符。

      示例：
         | <unset>                 | 空列表完全禁用加速（哈希修改上传）。
         | *                       | 所有文件都将尝试加速哈希算法处理。
         | *.mkv,*.avi,*.mp4,*.mp3 | 仅针对常见音视频文件，会尝试使用哈希修改上传算法进行处理。
         | *.zip,*.gz,*.rar,*.pdf  | 仅针对常见存档文件或 PDF 书籍，会尝试使用哈希修改上传算法进行处理。

   --mailru-check-hash
      如果文件校验和不匹配或无效，复制操作应该怎么做。

      示例：
         | true  | 报错并退出操作。
         | false | 忽略并继续操作。

   --mailru-encoding
      后端使用的编码方式。

      更多信息请参阅概述中的[编码部分](/overview/#encoding)。

OPTIONS:
   --help, -h  显示帮助信息。

   数据准备选项

   --delete-after-export    [危险操作] 导出 CAR 文件后删除数据集中的文件（默认值：false）。
   --rescan-interval value  当从上次成功扫描开始经过此时间间隔时，自动重新扫描源目录（默认值：禁用）

   Mail.ru Cloud 选项

   --mailru-check-hash value             如果文件校验和不匹配或无效，复制操作应该怎么做。 (默认值："true") [$MAILRU_CHECK_HASH]
   --mailru-encoding value               后端使用的编码方式。 (默认值："Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$MAILRU_ENCODING]
   --mailru-pass value                   应用程序密码。 [$MAILRU_PASS]
   --mailru-speedup-enable value         如果存在另一个具有相同数据哈希的文件，将跳过完全上传。 (默认值："true") [$MAILRU_SPEEDUP_ENABLE]
   --mailru-speedup-file-patterns value  逗号分隔的文件名模式列表，适用于加速（哈希修改上传）。 (默认值："*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf") [$MAILRU_SPEEDUP_FILE_PATTERNS]
   --mailru-speedup-max-disk value       该选项允许您为大文件禁用加速（哈希修改上传） (默认值："3Gi") [$MAILRU_SPEEDUP_MAX_DISK]
   --mailru-speedup-max-memory value     大于下面给出的大小的文件将始终在磁盘上进行哈希计算（默认值："32Mi"） [$MAILRU_SPEEDUP_MAX_MEMORY]
   --mailru-user value                   用户名（通常为电子邮件）。 [$MAILRU_USER]

```
{% endcode %}