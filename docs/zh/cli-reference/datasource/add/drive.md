# Google云端硬盘

{% code fullWidth="true" %}
```
名称：
   singularity datasource add drive - Google Drive

用法：
   singularity datasource add drive [命令选项] <数据集名称> <源路径>

说明：
   --drive-token
      OAuth访问令牌的JSON代码块。

   --drive-root-folder-id
      根文件夹的ID。
      通常留空。
      
      填写以访问“计算机”文件夹（请参见文档）或为rclone使用
      非根文件夹作为其起始点。
      

   --drive-v2-download-min-size
      如果对象较大，则使用v2 API下载。

   --drive-stop-on-download-limit
      使下载限制错误成为致命错误。
      
      在撰写本文时，一天只能从Google Drive下载10 TiB的数据（这是一项未经记录的限制）。
      当达到此限制时，Google Drive会产生略有不同的错误消息。当设置了这个标志时，
      它会导致这些错误成为致命错误。这将停止正在进行的同步。
      
      请注意，此检测依赖于谷歌未记录的错误信息字符串，因此可能会在将来发生错误。
      

   --drive-formats
      弃用：请参见export_formats。

   --drive-import-formats
      逗号分隔的首选格式列表，用于上传Google文档。

   --drive-upload-cutoff
      切换到分块上传的截止点。

   --drive-team-drive
      共享驱动器（团队驱动器）的ID。

   --drive-auth-owner-only
      仅考虑已认证用户拥有的文件。

   --drive-copy-shortcut-content
      复制快捷方式的内容而不是快捷方式的服务器端复制。
      
      在进行服务器端副本时，通常情况下rclone将快捷方式作为快捷方式复制。
      
      如果使用此标志，则rclone在进行服务器端副本时将复制快捷方式的内容，
      而不是快捷方式本身。

   --drive-skip-gdocs
      跳过所有列表中的谷歌文档。
      
      如果给定，谷歌文档在rclone中实际上会变得不可见。

   --drive-shared-with-me
      仅显示与我共享的文件。
      
      指示rclone在“与我共享的文件”文件夹中操作（Google Drive允许您访问其他人与您共享的文件和文件夹）。
      
      这适用于“列表”（lsd、lsl等）和“复制”（copy、sync等）命令，以及所有其他命令。

   --drive-stop-on-upload-limit
      使上传限制错误成为致命错误。
      
      在撰写本文时，一天只能将750 GiB的数据上传到Google Drive（这是一项未经记录的限制）。
      当达到此限制时，Google Drive会产生略有不同的错误消息。当设置了这个标志时，它会导致这些错误成为致命错误。
      这将停止正在进行的同步。
      
      请注意，此检测依赖于谷歌未记录的错误信息字符串，因此可能会在将来发生错误。
      
      参见: https://github.com/rclone/rclone/issues/3857
      

   --drive-scope
      rclone在请求来自驱动程序的访问时应使用的范围。

      示例：
         | drive                   | 完全访问所有文件，不包括应用程序数据文件夹。
         | drive.readonly          | 对文件元数据和文件内容的只读访问。
         | drive.file              | 仅访问由rclone创建的文件。
                                   | 这些文件在驱动器网站上可见。
                                   | 文件授权在用户解除应用授权时被撤销。
         | drive.appfolder         | 允许对应用程序数据文件夹进行读写访问。
                                   | 这在驱动器网站上不可见。
         | drive.metadata.readonly | 仅允许对文件元数据进行只读访问，
                                   | 但不允许访问或下载文件内容。

   --drive-skip-dangling-shortcuts
      如果设置则跳过悬空快捷方式文件。
      
      如果设置了这个文件，则rclone将不会在列表中显示任何悬空的快捷方式。

   --drive-alternate-export
      弃用：不再需要。

   --drive-acknowledge-abuse
      设置以允许下载返回cannotDownloadAbusiveFile的文件。
      
      如果下载文件返回错误消息“This file has been identified as malware or spam and cannot be downloaded”，
      错误代码为“cannotDownloadAbusiveFile”，则向rclone提供此标志以指示您承认下载文件的风险，
      并且rclone将仍会下载该文件。
      
      请注意，如果您使用的是服务帐户，则需要具有管理员权限（不是内容管理员），才能使此标志起作用。
      如果SA没有正确的权限，Google将忽略该标志。

   --drive-keep-revision-forever
      永久保留每个文件的新主修订版。

   --drive-client-secret
      OAuth客户端秘密。
      
      通常情况下保留空白。

   --drive-token-url
      令牌服务器网址。
      
      在使用提供商默认设置不希望将其留空。

   --drive-service-account-file
      服务帐户凭据JSON文件路径。
      
      通常情况下保留空白。
      仅在要使用SA而不是交互式登录时需要。

      在文件名中以“~”开头将被扩展为环境变量，
      如`${RCLONE_CONFIG_DIR}`。

   --drive-use-trash
      发送文件到回收站而不是永久删除。
      
      默认为true，即将文件发送到垃圾箱。
      使用`--drive-use-trash=false`代替永久删除文件。

   --drive-skip-checksum-gphotos
      仅跳过谷歌照片和视频的MD5校验和。
      
      如果在传输Google照片或视频时出现校验和错误，请使用此选项。
      
      设置此标志将导致Google照片和视频返回一个空的MD5校验和。
      
      Google照片的标识方法是在“照片”空间中。
      
      损坏的校验和是由于Google修改了图像/视频，但未更新校验和造成的。

   --drive-pacer-min-sleep
      API调用之间的最小休眠时间。

   --drive-disable-http2
      禁用驱动使用http2。
      
      目前，谷歌驱动器后端和HTTP/2存在未解决的问题。
      因此，默认情况下禁用了驱动程序后端的HTTP/2，但可以在此处重新启用。当问题解决时，将删除该标志。
      
      参见: https://github.com/rclone/rclone/issues/3631
      
      

   --drive-skip-shortcuts
      如果设置则跳过快捷方式文件。
      
      通常情况下，rclone取消引用快捷方式文件，使它们看起来像原始文件一样（请参阅[快捷方式部分](#shortcuts)）。
      如果设置了此标志，则rclone将完全忽略快捷方式文件。

   --drive-service-account-credentials
      服务帐户凭据JSON代码
- --drive-size - 必须是2的幂，大于等于256k。增加这个值将提高性能，但是需要注意每个块需要缓冲一次。缩小这个值将减少内存使用，但是会降低性能。

- --drive-size-as-quota - 展示文件大小为存储配额，而不是实际大小。展示文件大小为存储配额使用情况。这是当前版本以及已设置为永久保留的旧版本。

   **警告**：这个标志可能会有一些意想不到的后果。不建议在配置中设置这个标志 - 推荐的用法是仅在做 rclone ls/lsl/lsf/lsjson等操作时使用标志形式 --drive-size-as-quota。如果您确实要将这个标志用于同步（不推荐），则需要同时使用 --ignore size 标志。

- --drive-client-id - Google 应用程序客户端 ID。建议您设置自己的 ID。详见 https://rclone.org/drive/#making-your-own-client-id 如何创建您自己的ID。如果您将其留空，则会使用一个性能较低的内部键。

- --drive-auth-url - 认证服务器 URL。如果留空，则使用提供者默认设置。

- --drive-trashed-only - 仅显示处于回收站中的文件。这将显示已删除的文件的原始目录结构。

- --drive-allow-import-name-change - 允许在上传 Google 文档时更改文件类型。例如，将 file.doc 更改为 file.docx。这将混淆同步并导致每次重新上传。

- --drive-list-chunk - 列出块的大小。范围在100-1000之间，0表示禁用。

- --drive-encoding - 后端的编码方式。请参见概述中的[编码部分](/overview/#encoding)。

- --drive-starred-only - 仅显示已加星标记的文件。

- --drive-export-formats - 以逗号分隔的首选格式列表，用于下载 Google 文档。

- --drive-use-shared-date - 使用共享文件的日期，而不是修改日期。请注意，与“ --drive-use-created-date”一样，这个标志可能会在上传/下载文件时产生意想不到的结果。如果同时设置了此标志和“--drive-use-created-date”，则使用已创建的日期。

- --drive-pacer-burst - 允许不休眠的 API 调用次数。