# 本地磁盘

{% code fullWidth="true" %}
```
命令:
   singularity datasource add local - 本地磁盘

用法:
   singularity datasource add local [命令选项] <数据集名称> <源路径>

描述:
   --local-case-insensitive
      强制文件系统自声明为不区分大小写。
      
      通常本地后端在Windows/macOS上自声明为不区分大小写，而对于其他系统则自声明为区分大小写。使用此标志可以覆盖默认选择。

   --local-case-sensitive
      强制文件系统自声明为区分大小写。
      
      通常本地后端在Windows/macOS上自声明为不区分大小写，而对于其他系统则自声明为区分大小写。使用此标志可以覆盖默认选择。

   --local-copy-links
      跟随符号链接并复制指向的项目。

   --local-encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --local-links
      将符号链接转换成常规文件，并加上“.rclonelink”扩展名。

   --local-no-check-updated
      在上传过程中不检查文件是否更改。
      
      通常，在上传文件时，rclone会检查文件的大小和修改时间，如果文件在上传过程中发生更改，则会中止并显示一条消息，消息以“can't copy - source file is being updated”开头。
      
      但是在某些文件系统上，此修改时间检查可能失败（例如 [Glusterfs #2206](https://github.com/rclone/rclone/issues/2206)），因此可以使用此标志禁用此检查。
      
      如果设置了此标志，rclone将尽力传输正在更新的文件。如果文件只是追加内容（例如日志），那么rclone将传输拥有第一次看到的大小的日志文件。
      
      如果文件正在进行完全修改（而不仅仅是追加），则传输可能会因哈希检查失败而失败。
      
      具体来说，一旦文件被首次调用stat()：
      
      - 只传输stat给出的大小
      - 只对stat给出的大小进行校验和
      - 不更新文件的stat信息
      
      

   --local-no-preallocate
      禁用传输文件的磁盘空间预先分配。
      
      磁盘空间的预先分配有助于防止文件系统碎片化。但是，某些虚拟文件系统层（例如Google Drive File Stream）可能会错误地将实际文件大小设置为已预先分配的空间大小，导致校验和和文件大小检查失败。使用此标志禁用预先分配。

   --local-no-set-modtime
      禁用修改修改时间。
      
      通常，在文件上传完成后，rclone会更新文件的修改时间。当rclone正在运行的用户不拥有所上传的文件（例如，复制到由另一个用户拥有的CIFS挂载点）时，这可能会导致权限问题。如果启用了此选项，rclone将不再在复制文件后更新修改时间。

   --local-no-sparse
      禁用多线程下载的稀疏文件。
      
      在Windows平台上，rclone在进行多线程下载时会生成稀疏文件。这样可以避免在操作系统将文件清零时出现长时间的暂停。但稀疏文件可能不可取，因为它们会导致磁盘碎片化，并且使用起来可能很慢。

   --local-nounc
      在Windows上禁用UNC（长路径名）转换。

      示例：
         | true | 禁用长文件名。

   --local-one-file-system
      不跨文件系统边界（仅适用于Unix/macOS）。

   --local-skip-links
      不提醒跳过的符号链接。
      
      此标志禁用有关跳过的符号链接或结合点的警告消息，因为您明确确认它们应该被跳过。

   --local-unicode-normalization
      将路径和文件名应用于Unicode NFC规范化。
      
      可以使用此标志将从本地文件系统读取的文件名规范化为Unicode NFC形式。
      
      rclone通常不会更改从文件系统读取的文件名的编码。
      
      当在macOS上使用时，这可能很有用，因为它通常提供了分解的（NFD）Unicode，在某些语言（例如韩语）中在某些操作系统上无法正确显示。
      
      请注意，rclone在同步过程中使用Unicode规范化比较文件名，因此通常不应使用此标志。

   --local-zero-size-links
      假设链接的Stat大小为零（并读取它们）（已弃用）。
      
      Rclone之前使用链接的Stat大小作为链接大小，但在以下几个位置失败：
      
      - Windows
      - 某些虚拟文件系统（例如LucidLink）
      - Android
      
      因此，rclone现在总是读取链接。
      


选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险] 导出数据集为CAR文件后删除数据集的文件。 (默认值: false)
   --rescan-interval value  当距离上一次成功扫描的时间间隔过去时，自动重新扫描源目录（默认值: 禁用）
   --scanning-state value   设置初始扫描状态（默认值: 准备就绪）

   本地选项

   --local-case-insensitive value       强制文件系统自声明为不区分大小写。 (默认值: "false") [$LOCAL_CASE_INSENSITIVE]
   --local-case-sensitive value         强制文件系统自声明为区分大小写。 (默认值: "false") [$LOCAL_CASE_SENSITIVE]
   --local-copy-links value             跟随符号链接并复制指向的项目。 (默认值: "false") [$LOCAL_COPY_LINKS]
   --local-encoding value               后端的编码方式。 (默认值: "Slash,Dot") [$LOCAL_ENCODING]
   --local-links value                  将符号链接转换成常规文件，并加上“.rclonelink”扩展名。 (默认值: "false") [$LOCAL_LINKS]
   --local-no-check-updated value       在上传过程中不检查文件是否更改。 (默认值: "false") [$LOCAL_NO_CHECK_UPDATED]
   --local-no-preallocate value         禁用传输文件的磁盘空间预先分配。 (默认值: "false") [$LOCAL_NO_PREALLOCATE]
   --local-no-set-modtime value         禁用修改修改时间。 (默认值: "false") [$LOCAL_NO_SET_MODTIME]
   --local-no-sparse value              禁用多线程下载的稀疏文件。 (默认值: "false") [$LOCAL_NO_SPARSE]
   --local-nounc value                  在Windows上禁用UNC（长路径名）转换。 (默认值: "false") [$LOCAL_NOUNC]
   --local-one-file-system value        不跨文件系统边界（仅适用于Unix/macOS）。 (默认值: "false") [$LOCAL_ONE_FILE_SYSTEM]
   --local-skip-links value             不提醒跳过的符号链接。 (默认值: "false") [$LOCAL_SKIP_LINKS]
   --local-unicode-normalization value  将路径和文件名应用于Unicode NFC规范化。 (默认值: "false") [$LOCAL_UNICODE_NORMALIZATION]
   --local-zero-size-links value        假设链接的Stat大小为零（并读取它们）（已弃用）。 (默认值: "false") [$LOCAL_ZERO_SIZE_LINKS]

```
{% endcode %}