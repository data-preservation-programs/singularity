# 本地磁盘

{% code fullWidth="true" %}
```
名称:
   singularity storage create local - 本地磁盘

用法:
   singularity storage create local [命令选项] [参数...]

描述:
   --nounc
      在 Windows 上禁用 UNC（长路径名）转换。

      示例：
         | true | 禁用长文件名。

   --copy-links
      跟随符号链接并复制指向的项目。

   --links
      将符号链接转换为普通文件或从普通文件转换为带有“.rclonelink”扩展名的符号链接。

   --skip-links
      不会警告有关跳过的符号链接。
      
      此标志禁止在跳过符号链接或联接点时显示警告消息，因为您明确承认应该跳过它们。

   --zero-size-links
      假设链接的 Stat 大小为零（并读取它们）（已弃用）。
      
      Rclone 曾经使用链接的 Stat 大小作为链接大小，但在以下许多情况下失败：
      
      - Windows
      - 在某些虚拟文件系统（如 LucidLink）
      - Android
      
      因此，rclone 现在总是读取链接。

   --unicode-normalization
      对路径和文件名应用 Unicode NFC 标准化。
      
      可以使用此标志将从本地文件系统读取的文件名规范化为 Unicode NFC 格式。
      
      Rclone 通常不会更改从文件系统读取的文件名的编码。
      
      当使用 macOS 时，这可能很有用，因为它通常提供分解（NFD）的 Unicode，在某些操作系统上（例如韩国）无法正确显示。
      
      请注意，rclone 在同步过程中使用 Unicode 标准化来比较文件名，因此通常不应使用此标志。

   --no-check-updated
      不检查文件在上传期间是否更改。
      
      通常，在上传文件时，rclone 会检查文件的大小和修改时间，并在文件在上传过程中更改时以“无法复制 - 源文件正在更新”开头的消息中中止。
      
      但是，在某些文件系统上，此修改时间检查可能失败（例如 [Glusterfs#2206](https://github.com/rclone/rclone/issues/2206)），因此可以使用此标志禁用此检查。
      
      如果设置了此标志，则 rclone 将尽最大努力传输正在更新的文件。如果文件仅在其后附加了一些内容（例如日志），那么 rclone 将使用它第一次看到时的大小传输日志文件。
      
      如果文件在整个过程中被修改（不仅是附加的内容），则传输可能会由于哈希检查失败而失败。
      
      具体来说，一旦首次调用文件的 stat()，我们会：
      
      - 仅传输 stat 给出的大小
      - 仅对 stat 给出的大小进行校验和
      - 不更新文件的 stat 信息
      
      

   --one-file-system
      不会跨越文件系统边界（仅适用于 Unix/macOS）。

   --case-sensitive
      强制文件系统将其自身报告为区分大小写。
      
      通常，本地后端在 Windows/macOS 上声明自身为不区分大小写，对其他所有系统区分大小写。使用此标志覆盖默认选择。

   --case-insensitive
      强制文件系统将其自身报告为不区分大小写。
      
      通常，本地后端在 Windows/macOS 上声明自身为不区分大小写，对其他所有系统区分大小写。使用此标志覆盖默认选择。

   --no-preallocate
      禁用将磁盘空间预分配给传输文件。
      
      磁盘空间的预分配有助于防止文件系统碎片化。然而，某些虚拟文件系统层（如 Google Drive File Stream）可能会错误地将实际文件大小设置为预分配空间的大小，导致校验和和文件大小检查失败。使用此标志禁用预分配。

   --no-sparse
      禁用多线程下载时的稀疏文件。
      
      在 Windows 平台上，rclone 在执行多线程下载时会产生稀疏文件。这样可以避免在文件较大且操作系统将文件置零时出现较长的暂停。然而，稀疏文件可能不可取，因为它们会导致磁盘碎片化并且使用起来可能很慢。

   --no-set-modtime
      禁用设置修改时间。
      
      通常，rclone 在完成上传后会更新文件的修改时间。这可能会在 Linux 平台上出现权限问题，特别是当 rclone 运行的用户不拥有已上传的文件时，例如在复制到另一个用户所拥有的 CIFS 挂载点时。如果启用此选项，rclone 将不再在复制文件后更新修改时间。

   --encoding
      后端的编码。
      
      有关更多信息，请参见概述中的[编码部分](/overview/#encoding)。

选项:
   --help, -h  显示帮助信息

   高级选项

   --case-insensitive       强制文件系统将其自身报告为不区分大小写。 (默认值：false) [$CASE_INSENSITIVE]
   --case-sensitive         强制文件系统将其自身报告为区分大小写。 (默认值：false) [$CASE_SENSITIVE]
   --copy-links, -L         跟随符号链接并复制指向的项目。 (默认值：false) [$COPY_LINKS]
   --encoding value         后端的编码。 (默认值："Slash,Dot") [$ENCODING]
   --links, -l              将符号链接转换为普通文件或从普通文件转换为带有“.rclonelink”扩展名的符号链接。 (默认值：false) [$LINKS]
   --no-check-updated       不检查文件在上传期间是否更改。 (默认值：false) [$NO_CHECK_UPDATED]
   --no-preallocate         禁用将磁盘空间预分配给传输文件。 (默认值：false) [$NO_PREALLOCATE]
   --no-set-modtime         禁用设置修改时间。 (默认值：false) [$NO_SET_MODTIME]
   --no-sparse              禁用多线程下载时的稀疏文件。 (默认值：false) [$NO_SPARSE]
   --nounc                  在 Windows 上禁用 UNC（长路径名）转换。 (默认值：false) [$NOUNC]
   --one-file-system, -x    不会跨越文件系统边界（仅适用于 Unix/macOS）。 (默认值：false) [$ONE_FILE_SYSTEM]
   --skip-links             不会警告有关跳过的符号链接。 (默认值：false) [$SKIP_LINKS]
   --unicode-normalization  对路径和文件名应用 Unicode NFC 标准化。 (默认值：false) [$UNICODE_NORMALIZATION]
   --zero-size-links        假设链接的 Stat 大小为零（并读取它们）（已弃用）。 (默认值：false) [$ZERO_SIZE_LINKS]

   常规选项

   --name value  存储的名称 (默认值：自动生成)
   --path value  存储的路径

```
{% endcode %}