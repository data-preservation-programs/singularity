# 本地磁盘

{% code fullWidth="true" %}
```
名称:
   singularity datasource add local - 本地磁盘

用法:
   singularity datasource add local [命令选项] <数据集名称> <数据源路径>

说明:
   --local-links
      将符号链接转换为普通文件并加上“.rclonelink”扩展名。

   --local-case-sensitive
      强制文件系统报告其区分大小写。
      
      通常，本地后端在 Windows/macOS 上声明为不区分大小写，而在其他系统上即区分大小写。使用此标志覆盖默认选择。

   --local-no-preallocate
      禁用已转移文件的磁盘空间的预分配。
      
      预先分配磁盘空间有助于防止文件系统碎片。但是，某些虚拟文件系统层（例如 Google Drive File Stream）可能会错误地将实际文件大小设置为分配的空间大小，导致校验和和文件大小检查失败。使用此标志来禁用预分配。

   --local-no-set-modtime
      禁用设置修改时间。
      
      通常，rclone在上传完成后会更新文件的修改时间。当 rclone 运行的用户不拥有上传的文件（例如在拷贝到由另一个用户拥有的 CIFS 挂载时）时，这可能会引起权限问题。如果启用此选项，则 rclone 更不会在拷贝文件后更新修改时间。

   --local-copy-links
      对符号链接进行跟随并复制指向的项目。

   --local-skip-links
      不会警告已跳过的符号链接。
      
      此标志禁用有关跳过的符号链接或联接点的警告消息，因此您明确知道应跳过它们。

   --local-no-check-updated
      不检查上传期间文件是否更改。
      
      通常，rclone 会在上传时检查文件的大小和修改时间，并在文件上传期间更改时中止并显示一条消息开头为“can't copy -  source file is being updated”的消息。
      
      但是，在某些文件系统中，此修改时间检查可能会失败（例如 [Glusterfs #2206](https://github.com/rclone/rclone/issues/2206)），因此可以使用此标志禁用此检查。
      
      如果设置了此标志，则 rclone 将尽其所能传输正在更新的文件。如果文件只是在其中添加内容（例如日志），那么 rclone 将传输第一次看到它时具有的大小的日志文件。
      
      如果该文件正在全面修改（而不仅仅是添加到其中），则传输可能会失败，显示散列校验失败信息。
      
      详细来说，一旦对该文件第一次调用 stat()，我们就会：
      
      - 仅传输 stat 指定的大小
      - 仅核对 stat 指定的大小
      - 不要更新文件的 stat 信息
      
      

   --local-no-sparse
      禁用多线程下载的稀疏文件。
      
      在 Windows 平台上，rclone 在进行多线程下载时会制作稀疏文件。这可以避免 OS 长时间在文件中放置零。但是稀疏文件可能不可取，因为它们会导致磁盘碎片，并且使用起来可能很慢。

   --local-encoding
      后端的编码。
      
      有关详细信息，请参见 [总览中的编码部分](/overview/#encoding)。

   --local-nounc
      在 Windows 上禁用 UNC（长文件名）转换。

      示例:
         | true | 禁用长文件名。

   --local-unicode-normalization
      对路径和文件名应用 Unicode NFC 标准化。
      
      可以使用此标志将从本地文件系统读取的文件名标准化为 Unicode NFC 形式。
      
      rclone 通常不会触及从文件系统中读取的文件名的编码。
      
      在 macOS 上使用它可以非常有用，因为它通常提供了已分解的（NFD）Unicode，这在某些语言（例如韩文）上在某些操作系统上无法正确显示。
      
      请注意，rclone 在同步过程中比较文件名时使用 Unicode 标准化，因此通常不应使用此标志。

   --local-one-file-system
      不要跨越文件系统边界（仅限 Unix/macOS）。

   --local-zero-size-links
      假设链接的 Stat 大小为零（并读取它们）（已废弃）。
      
      rclone 曾经使用链接的 Stat 大小作为链接大小，但它在相当多的地方失败了：
      
      - Windows
      - 在某些虚拟文件系统上（例如 LucidLink）
      - Android
      
      因此，rclone 现在总是读取链接。
      

   --local-case-insensitive
      强制文件系统报告其区分大小写。
      
      通常，本地后端在 Windows/macOS 上声明为不区分大小写，而在其他系统上即区分大小写。使用此标志覆盖默认选择。


选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 在将数据集导出为 CAR 文件后删除其文件。  (默认值: false)
   --rescan-interval value  当此间隔从上次成功扫描后经过时，自动重新扫描源目录（默认: 禁用）

   本地选项

   --local-case-insensitive value           强制文件系统报告其区分大小写。 (默认值: "false") [$LOCAL_CASE_INSENSITIVE]
   --local-case-sensitive value             强制文件系统报告其区分大小写。 (默认值: "false") [$LOCAL_CASE_SENSITIVE]
   --local-copy-links value, -L value       对符号链接进行跟随并复制指向的项目。 (默认值: "false") [$LOCAL_COPY_LINKS]
   --local-encoding value                   后端的编码。 (默认值: "Slash,Dot") [$LOCAL_ENCODING]
   --local-links value, -l value            将符号链接转换为普通文件并加上“.rclonelink”扩展名。 (默认值: "false") [$LOCAL_LINKS]
   --local-no-check-updated value           不检查上传期间文件是否更改。 (默认值: "false") [$LOCAL_NO_CHECK_UPDATED]
   --local-no-preallocate value             禁用已转移文件的磁盘空间的预分配。 (默认值: "false") [$LOCAL_NO_PREALLOCATE]
   --local-no-set-modtime value             禁用设置修改时间。 (默认值: "false") [$LOCAL_NO_SET_MODTIME]
   --local-no-sparse value                  禁用多线程下载的稀疏文件。 (默认值: "false") [$LOCAL_NO_SPARSE]
   --local-nounc value                      在 Windows 上禁用 UNC（长文件名）转换。 (默认值: "false") [$LOCAL_NOUNC]
   --local-one-file-system value, -x value  不要跨越文件系统边界（仅限 Unix/macOS）。 (默认值: "false") [$LOCAL_ONE_FILE_SYSTEM]
   --local-skip-links value                 不会警告已跳过的符号链接。 (默认值: "false") [$LOCAL_SKIP_LINKS]
   --local-unicode-normalization value      对路径和文件名应用 Unicode NFC 标准化。 (默认值: "false") [$LOCAL_UNICODE_NORMALIZATION]
   --local-zero-size-links value            假设链接的 Stat 大小为零（并读取它们）（已废弃）。 (默认值: "false") [$LOCAL_ZERO_SIZE_LINKS]

```
{% endcode %}