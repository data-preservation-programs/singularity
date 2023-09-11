# 本地磁盘

{% code fullWidth="true" %}
```
命令:
   singularity storage update local - 本地磁盘

用法:
   singularity storage update local [命令选项] <名称|ID>

说明:
   --nounc
      关闭在Windows上的UNC（长路径名）转换。

      示例:
         | true | 关闭长文件名。

   --copy-links
      跟随符号链接并复制指向的项目。

   --links
      将符号链接与普通文件之间进行转换，并添加“.rclonelink”扩展名。

   --skip-links
      不要警告被跳过的符号链接。
      
      此标志关闭对跳过的符号链接或联接点的警告消息，您明确确认应该跳过它们。

   --zero-size-links
      假设符号链接的 Stat 大小为零（并读取它们）（已弃用）。
      
      Rclone 曾经使用符号链接的 Stat 大小作为链接的大小，但这在很多地方都失败了:
      
      - Windows
      - 一些虚拟文件系统（例如 LucidLink）
      - Android
      
      因此，现在的 rclone 总是读取符号链接。
      

   --unicode-normalization
      对路径和文件名应用 Unicode NFC 标准化。
      
      此标志可用于将从本地文件系统读取的文件名规范化为 Unicode NFC 形式。
      
      Rclone 通常不会更改从文件系统读取的文件名的编码。
      
      在使用 macOS 时，这可能很有用，因为它通常会提供分解的（NFD）Unicode，在某些操作系统上（例如韩语）不能正确显示。
      
      请注意，rclone 在同步过程中使用 Unicode 标准化比较文件名，因此通常不应使用此标志。

   --no-check-updated
      不检查上传过程中文件是否发生更改。
      
      通常情况下，rclone 会检查文件的大小和修改时间，如果在上传过程中文件发生更改，则会中止并显示一条以“can't copy - source file is being updated”开头的消息。
      
      但是在某些文件系统上，此修改时间检查可能失败（例如 [Glusterfs #2206](https://github.com/rclone/rclone/issues/2206)），因此可以使用此标志来禁用此检查。
      
      如果设置了此标志，rclone 将尽最大努力传输正在更新的文件。如果文件只是在追加内容（例如日志），则 rclone 将使用首次看到该文件时的大小传输日志文件。
      
      如果文件正在全程修改（而不仅仅是追加），则传输可能会因为哈希检查失败而失败。
      
      具体地说，一旦第一次调用了 stat() 函数获取文件的信息，我们将:
      
      - 只传输 stat 函数给出的大小
      - 只对 stat 函数给出的大小进行校验
      - 不更新文件的 stat 信息
      
      

   --one-file-system
      不穿越文件系统边界（仅适用于 Unix/macOS）。

   --case-sensitive
      强制文件系统报告自身为区分大小写。
      
      通常情况下，本地后端会在 Windows/macOS 上声明自身为大小写不敏感，在其他所有情况下声明自身为区分大小写。使用此标志覆盖默认选择。

   --case-insensitive
      强制文件系统报告自身为大小写不敏感。
      
      通常情况下，本地后端会在 Windows/macOS 上声明自身为大小写不敏感，在其他所有情况下声明自身为区分大小写。使用此标志覆盖默认选择。

   --no-preallocate
      禁用传输文件时的磁盘空间预分配。
      
      磁盘空间预分配有助于防止文件系统碎片化。但是，某些虚拟文件系统层（例如 Google Drive 文件流）可能会错误地将实际文件大小设置为等于预分配的空间大小，导致校验和和文件大小检查失败。使用此标志来禁用预分配。

   --no-sparse
      禁用多线程下载的稀疏文件。
      
      在 Windows 平台上，rclone 在进行多线程下载时会创建稀疏文件。这样可以避免在操作系统将文件清零时导致长时间的停顿。但稀疏文件可能不太理想，因为它们会导致磁盘碎片化，并且处理起来可能很慢。

   --no-set-modtime
      禁用设置修改时间。
      
      通常情况下，rclone 会在上传完成后更新文件的修改时间。这可能会导致在 Linux 平台上出现权限问题，因为 rclone 所运行的用户不拥有上传的文件，例如当复制到由另一个用户拥有的 CIFS 挂载点时。如果启用了此选项，rclone 将不再在复制文件后更新修改时间。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。


选项:
   --help, -h  显示帮助

   高级选项

   --case-insensitive    强制文件系统报告自身为大小写不敏感。 (默认值: false) [$CASE_INSENSITIVE]
   --case-sensitive      强制文件系统报告自身为区分大小写。 (默认值: false) [$CASE_SENSITIVE]
   --copy-links, -L      跟随符号链接并复制指向的项目。 (默认值: false) [$COPY_LINKS]
   --encoding value      后端的编码方式。 (默认值: "Slash,Dot") [$ENCODING]
   --links, -l           将符号链接与普通文件之间进行转换，并添加“.rclonelink”扩展名。 (默认值: false) [$LINKS]
   --no-check-updated    不检查上传过程中文件是否发生更改。 (默认值: false) [$NO_CHECK_UPDATED]
   --no-preallocate      禁用传输文件时的磁盘空间预分配。 (默认值: false) [$NO_PREALLOCATE]
   --no-set-modtime      禁用设置修改时间。 (默认值: false) [$NO_SET_MODTIME]
   --no-sparse           禁用多线程下载的稀疏文件。 (默认值: false) [$NO_SPARSE]
   --nounc               关闭在Windows上的UNC（长路径名）转换。 (默认值: false) [$NOUNC]
   --one-file-system, -x 不穿越文件系统边界（仅适用于 Unix/macOS）。 (默认值: false) [$ONE_FILE_SYSTEM]
   --skip-links          不要警告被跳过的符号链接。 (默认值: false) [$SKIP_LINKS]
   --unicode-normalization  对路径和文件名应用 Unicode NFC 标准化。 (默认值: false) [$UNICODE_NORMALIZATION]
   --zero-size-links     假设符号链接的 Stat 大小为零（并读取它们）（已弃用）。 (默认值: false) [$ZERO_SIZE_LINKS]

```
{% endcode %}