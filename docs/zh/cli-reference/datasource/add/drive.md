# Google Drive

{% code fullWidth="true" %}
```
名称：
   singularity datasource add drive - Google Drive

用法：
   singularity datasource add drive [命令选项] <数据集名称> <源路径>

描述：
   --drive-acknowledge-abuse
      设置为允许下载返回“cannotDownloadAbusiveFile”的文件。
      
      如果下载文件返回错误“This file has been identified as malware or spam and cannot be downloaded”，
      错误码为“cannotDownloadAbusiveFile”，请向rclone提供此标志以表明您知道下载文件的风险，rclone将继续下载。
      
      请注意，如果您使用的是服务帐号，则需要Manager权限（而不是内容管理器）才能使此标志有效。
      如果SA没有正确的权限，Google将直接忽略该标志。

   --drive-allow-import-name-change
      允许上传Google文档时文件类型更改。
      
      如果上传文件时，文件类型更改为其他格式（例如，file.doc转为file.docx），这会导致同步后每次重新上传。
      
   --drive-alternate-export
      已弃用：不再需要。

   --drive-auth-owner-only
      仅考虑经过身份验证的用户拥有的文件。

   --drive-auth-url
      用户认证服务器URL。
      
      留空以使用提供者的默认值。

   --drive-chunk-size
      上传块大小。
      
      必须是大于或等于256k的2的幂。
      
      将此值增大可提高性能，但请注意每个块都会在内存中缓冲一次数据。
      
      减小此值可减少内存使用量，但会降低性能。

   --drive-client-id
      Google应用程序客户端ID
      强烈建议您设置自己的客户端ID。
      请参阅https://rclone.org/drive/#making-your-own-client-id获取如何创建自己的ID。
      如果您将此值留空，则会使用低性能的内部密钥。

   --drive-client-secret
      OAuth客户端秘钥。
      
      通常将其留空。

   --drive-copy-shortcut-content
      使用服务器端复制快捷方式的内容，而不是复制快捷方式本身。
      
      当执行服务器端复制操作时，rclone通常会将快捷方式复制为快捷方式。
      
      如果使用此标志，则rclone在执行服务器端复制操作时将复制快捷方式的内容而不是快捷方式本身。

   --drive-disable-http2
      禁用Drive的HTTP/2支持。
      
      目前，Google Drive后端和HTTP/2存在一个未解决的问题。因此，默认情况下，drive后端禁用HTTP/2，但可以在此处重新启用。
      当问题解决后，此标志将被删除。
      
      参见：https://github.com/rclone/rclone/issues/3631。
      
      

   --drive-encoding
      后端的编码。
      
      参见概览中的[编码部分](/overview/#encoding)获取更多信息。

   --drive-export-formats
      用于下载Google文档的首选格式（逗号分隔）。

   --drive-formats
      已弃用：参见export_formats。

   --drive-impersonate
      使用服务帐号时模拟此用户。

   --drive-import-formats
      用于上传Google文档的首选格式（逗号分隔）。

   --drive-keep-revision-forever
      永久保存每个文件的新头部修订版。

   --drive-list-chunk
      列出的块的大小，100-1000，0表示禁用。

   --drive-pacer-burst
      允许在不休眠的情况下调用的API数量。

   --drive-pacer-min-sleep
      API调用之间的最小休眠时间。

   --drive-resource-key
      用于访问共享链接文件的资源密钥。
      
      如果您需要访问共享链接文件，例如
      
          https://drive.google.com/drive/folders/XXX?resourcekey=YYY&usp=sharing
      
      那么您需要将第一部分“XXX”作为“root_folder_id”，将第二部分“YYY”作为“resource_key”，否则在尝试访问目录时会收到404未找到错误。
      
      参见：https://developers.google.com/drive/api/guides/resource-keys
      
      此资源密钥要求仅适用于某些旧文件的子集。
      
      还请注意，在网络界面中打开文件夹一次（使用您在rclone进行身份验证的用户）似乎足以使资源密钥不再需要。
      

   --drive-root-folder-id
      根文件夹的ID。
      通常为空白。
      
      如需访问“计算机”文件夹（请参阅文档），或者希望rclone使用非根文件夹作为起点，可填入此值。
      

   --drive-scope
      rclone在请求Drive访问权限时应使用的范围。

      示例：
         | drive                   | 具有完全访问所有文件的权限，但不包括应用数据文件夹。
         | drive.readonly          | 对文件元数据和文件内容的只读访问权限。
         | drive.file              | 仅对rclone创建的文件进行访问。
                                   | 这些文件在Drive网站上可见。
                                   | 文件授权将在用户取消授权应用程序时被撤销。
         | drive.appfolder         | 允许对应用程序数据文件夹进行读写访问。
                                   | Drive网站上不显示此文件夹。
         | drive.metadata.readonly | 仅对文件元数据进行只读访问，但不允许访问文件内容。

   --drive-server-side-across-configs
      允许服务器端操作（如复制）在不同的Drive配置之间工作。
      
      如果您希望在两个不同的Google Drive之间进行服务器端复制，这可能会有所帮助。但请注意，默认情况下不启用该功能，因为很难确定它是否适用于任何两个配置。

   --drive-service-account-credentials
      服务帐号凭据的JSON数据。
      
      通常将其留空。
      只有在需要使用服务帐号而不是交互式登录时才需要。

   --drive-service-account-file
      服务帐号凭据的JSON文件路径。
      
      通常将其留空。
      只有在需要使用服务帐号而不是交互式登录时才需要。
      
      文件名中的主目录将扩展为“~”，环境变量如`${RCLONE_CONFIG_DIR}`也将被扩展。

   --drive-shared-with-me
      仅显示与我共享的文件。
      
      告知rclone在“与我共享”文件夹上操作（Google Drive允许您访问其他人与您共享的文件和文件夹）。
      
      无论是“list”命令（如lsd，lsl等）还是“copy”命令（如copy，sync等），以及所有其他命令，这都适用。

   --drive-size-as-quota
      以存储配额使用量而不是实际大小展示文件大小。
      
      将文件的大小显示为存储配额使用量。这是当前版本加上设置为永久保留的旧版本。
      
      **警告**：此标志可能会产生一些意外后果。
      
      不建议在配置中设置此标志，建议使用标志形式`--drive-size-as-quota`仅在执行rclone ls/lsl/lsf/lsjson等操作时使用。
      
      如果您确实在同步过程中使用此标志（不推荐），则还必须使用`--ignore size`。

   --drive-skip-checksum-gphotos
      仅跳过Google照片和视频的MD5校验和。
      
      如果在传输Google照片或视频时发生校验和错误，请使用此选项。
      
      设置此标志将导致Google照片和视频返回空白的MD5校验和。
      
      Google照片的特征是在“photos”空间中。
      
      校验和错误是由Google修改图像/视频文件但不更新校验和所造成的。

   --drive-skip-dangling-shortcuts
      如果设置，则跳过悬空的快捷方式文件。
      
      如果设置此标志，则rclone不会在列表中显示任何悬空的快捷方式文件。
      

   --drive-skip-gdocs
      在所有列表中跳过Google文档。
      
      如果设置此标志，则gdocs在rclone中实际上看不见。

   --drive-skip-shortcuts
      如果设置，则跳过快捷方式文件。
      
      通常rclone会对快捷方式文件进行解引用，使它们看起来像是原始文件（参见[快捷方式部分](#shortcuts)）。
      如果设置了此标志，则rclone将完全忽略快捷方式文件。
      

   --drive-starred-only
      仅显示已加星标识的文件。

   --drive-stop-on-download-limit
      使下载限制错误成为致命错误。
      
      在编写本文时，一天只能从Google Drive下载10TiB的数据（此限制未记录在案）。当达到此限制时，Google Drive会生成一个稍有不同的错误消息。如果设置了此标志，则会导致这些错误变为致命错误。这将停止正在进行的同步。
      
      请注意，此检测依赖于错误消息字符串，而Google并未对其进行文档化，因此在将来可能会失效。
      

   --drive-stop-on-upload-limit
      使上传限制错误成为致命错误。
      
      在编写本文时，一天只能将750 GiB的数据上传到Google Drive（此限制未记录在案）。当达到此限制时，Google Drive会生成一个稍有不同的错误消息。如果设置了此标志，则会导致这些错误变为致命错误。这将停止正在进行的同步。
      
      请注意，此检测依赖于错误消息字符串，而Google并未对其进行文档化，因此在将来可能会失效。
      
      参见：https://github.com/rclone/rclone/issues/3857。
      

   --drive-team-drive
      共享的驱动器（团队驱动器）的ID。

   --drive-token
      OAuth访问令牌，JSON形式。

   --drive-token-url
      令牌服务器URL。
      
      留空以使用提供者的默认值。

   --drive-trashed-only
      仅显示在垃圾箱中的文件。
      
      这将按原始目录结构显示垃圾箱中的文件。

   --drive-upload-cutoff
      切换到分块上传的临界值。

   --drive-use-created-date
      使用文件创建日期而不是修改日期。
      
      在下载数据且希望使用创建日期代替最后修改日期时很有用。
      
      **警告**：此标志可能会产生一些意外后果。
      
      当上传到Drive时，除非文件自创建后未被修改，否则所有文件都将被覆盖。而在下载时，相反的情况也会发生。您可以通过使用``--checksum``标志避免这种副作用。
      
      此功能的实现是为了在Google照片记录的照片拍摄日期保留的情况下使用。您首先需要在Google Drive设置中选中“创建一个Google照片文件夹”选项。然后您可以在本地复制或移动照片，并使用图像拍摄的日期（创建日期）作为修改日期。

   --drive-use-shared-date
      使用文件分享日期而不是修改日期。
      
      请注意，与`--drive-use-created-date`一样，此标志可能会在上传/下载文件时产生一些意外后果。
      
      如果`--drive-use-shared-date`和`--drive-use-created-date`都设置，则使用创建日期。

   --drive-use-trash
      将文件发送到垃圾箱而不是永久删除。
      
      默认情况下为true，即将文件发送到垃圾箱。
      使用`--drive-use-trash=false`以永久删除文件。

   --drive-v2-download-min-size
      如果对象更大，则使用Drive V2 API下载。

选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作] 在导出CAR文件后删除数据集的文件。  (default: false)
   --rescan-interval value  当距离上次成功扫描已经过去此间隔时，自动重新扫描源目录（默认：禁用）
   --scanning-state value   设置初始扫描状态（默认：准备就绪）

   drive选项

   --drive-acknowledge-abuse value            设置为允许下载返回“cannotDownloadAbusiveFile”的文件（默认值："false"）[$DRIVE_ACKNOWLEDGE_ABUSE]
   --drive-allow-import-name-change value     允许上传Google文档时文件类型更改（默认值："false"）[$DRIVE_ALLOW_IMPORT_NAME_CHANGE]
   --drive-auth-owner-only value              仅考虑经过身份验证的用户拥有的文件（默认值："false"）[$DRIVE_AUTH_OWNER_ONLY]
   --drive-auth-url value                     用户认证服务器URL [$DRIVE_AUTH_URL]
   --drive-chunk-size value                   上传块大小（默认值："8Mi"）[$DRIVE_CHUNK_SIZE]
   --drive-client-id value                    Google应用程序客户端ID [$DRIVE_CLIENT_ID]
   --drive-client-secret value                OAuth客户端秘钥 [$DRIVE_CLIENT_SECRET]
   --drive-copy-shortcut-content value        使用服务器端复制快捷方式的内容，而不是复制快捷方式本身（默认值："false"）[$DRIVE_COPY_SHORTCUT_CONTENT]
   --drive-disable-http2 value                禁用Drive使用HTTP/2（默认值："true"）[$DRIVE_DISABLE_HTTP2]
   --drive-encoding value                     后端的编码（默认值："InvalidUtf8"）[$DRIVE_ENCODING]
   --drive-export-formats value               用于下载Google文档的首选格式（逗号分隔， 默认值："docx,xlsx,pptx,svg"）[$DRIVE_EXPORT_FORMATS]
   --drive-formats value                      已弃用：参见export_formats [$DRIVE_FORMATS]
   --drive-impersonate value                  使用服务帐号时模拟此用户 [$DRIVE_IMPERSONATE]
   --drive-import-formats value               用于上传Google文档的首选格式（逗号分隔）[$DRIVE_IMPORT_FORMATS]
   --drive-keep-revision-forever value        永久保存每个文件的新头部修订版（默认值："false"）[$DRIVE_KEEP_REVISION_FOREVER]
   --drive-list-chunk value                   列出的块的大小，100-1000，0表示禁用（默认值："1000"）[$DRIVE_LIST_CHUNK]
   --drive-pacer-burst value                  允许在不休眠的情况下调用的API数量（默认值："100"）[$DRIVE_PACER_BURST]
   --drive-pacer-min-sleep value              API调用之间的最小休眠时间（默认值："100ms"）[$DRIVE_PACER_MIN_SLEEP]
   --drive-resource-key value                 用于访问共享链接文件的资源密钥 [$DRIVE_RESOURCE_KEY]
   --drive-root-folder-id value               根文件夹的ID [$DRIVE_ROOT_FOLDER_ID]
   --drive-scope value                        rclone在请求Drive访问权限时应使用的范围 [$DRIVE_SCOPE]
   --drive-server-side-across-configs value   允许服务器端操作（如复制）在不同的Drive配置之间工作（默认值："false"）[$DRIVE_SERVER_SIDE_ACROSS_CONFIGS]
   --drive-service-account-credentials value  服务帐号凭据的JSON数据 [$DRIVE_SERVICE_ACCOUNT_CREDENTIALS]
   --drive-service-account-file value         服务帐号凭据的JSON文件路径 [$DRIVE_SERVICE_ACCOUNT_FILE]
   --drive-shared-with-me value               仅显示与我共享的文件（默认值："false"）[$DRIVE_SHARED_WITH_ME]
   --drive-size-as-quota value                以存储配额使用量而不是实际大小展示文件大小（默认值："false"）[$DRIVE_SIZE_AS_QUOTA]
   --drive-skip-checksum-gphotos value        仅跳过Google照片和视频的MD5校验和（默认值："false"）[$DRIVE_SKIP_CHECKSUM_GPHOTOS]
   --drive-skip-dangling-shortcuts value      如果设置，则跳过悬空的快捷方式文件（默认值："false"）[$DRIVE_SKIP_DANGLING_SHORTCUTS]
   --drive-skip-gdocs value                   在所有列表中跳过Google文档（默认值："false"）[$DRIVE_SKIP_GDOCS]
   --drive-skip-shortcuts value               如果设置，则跳过快捷方式文件（默认值："false"）[$DRIVE_SKIP_SHORTCUTS]
   --drive-starred-only value                 仅显示已加星标识的文件（默认值："false"）[$DRIVE_STARRED_ONLY]
   --drive-stop-on-download-limit value       使下载限制错误成为致命错误（默认值："false"）[$DRIVE_STOP_ON_DOWNLOAD_LIMIT]
   --drive-stop-on-upload-limit value         使上传限制错误成为致命错误（默认值："false"）[$DRIVE_STOP_ON_UPLOAD_LIMIT]
   --drive-team-drive value                   共享驱动器（团队驱动器）的ID [$DRIVE_TEAM_DRIVE]
   --drive-token value                        OAuth访问令牌，JSON形式 [$DRIVE_TOKEN]
   --drive-token-url value                    令牌服务器URL [$DRIVE_TOKEN_URL]
   --drive-trashed-only value                 仅显示在垃圾箱中的文件（默认值："false"）[$DRIVE_TRASHED_ONLY]
   --drive-upload-cutoff value                切换到分块上传的临界值（默认值："8Mi"）[$DRIVE_UPLOAD_CUTOFF]
   --drive-use-created-date value             使用文件创建日期而不是修改日期（默认值："false"）[$DRIVE_USE_CREATED_DATE]
   --drive-use-shared-date value              使用文件分享日期而不是修改日期（默认值："false"）[$DRIVE_USE_SHARED_DATE]
   --drive-use-trash value                    将文件发送到垃圾箱而不是永久删除（默认值："true"）[$DRIVE_USE_TRASH]
   --drive-v2-download-min-size value         如果对象更大，则使用Drive V2 API下载（默认值："off"）[$DRIVE_V2_DOWNLOAD_MIN_SIZE]

```
{% endcode %}