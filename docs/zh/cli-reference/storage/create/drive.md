# Google Drive

{% code fullWidth="true" %}
```
名称：
   singularity storage create drive - Google Drive

用法：
   singularity storage create drive [命令选项] [参数...]

描述：
   --client-id
      Google 应用程序客户端 ID
      建议您设置自己的客户端 ID。
      请参阅 https://rclone.org/drive/#making-your-own-client-id 了解如何创建自己的客户端 ID。
      如果您将此留空，它将使用一个性能较低的内部密钥。

   --client-secret
      OAuth 客户端密钥。
      
      通常留空。

   --token
      OAuth 访问令牌（JSON 格式）。

   --auth-url
      认证服务器 URL。
      
      留空以使用提供程序默认值。

   --token-url
      令牌服务器 URL。
      
      留空以使用提供程序默认值。

   --scope
      rclone 在请求从 Google Drive 获取访问权限时应使用的范围。

      示例：
         | drive                  | 具有完全访问所有文件权限，但不包括应用数据文件夹。
         | drive.readonly         | 仅具有对文件元数据和文件内容的只读访问权限。
         | drive.file             | 仅允许对 rclone 创建的文件访问。
         |                        | 这些文件在「Drive」网站上可见。
         |                        | 用户停用应用时，文件授权将被撤销。
         | drive.appfolder        | 允许对「应用数据」文件夹进行读写访问。
         |                        | 这些文件在「Drive」网站上不可见。
         | drive.metadata.readonly| 允许对文件元数据进行只读访问，
         |                        | 但不允许任何读取或下载文件内容的权限。

   --root-folder-id
      根文件夹的 ID。
      通常留空。
      
      如果要访问「计算机」文件夹（请参阅文档），或让 rclone 使用
      非根文件夹作为起始点，请填写该值。
      

   --service-account-file
      服务账号凭据的 JSON 文件路径。
      
      通常留空。
      仅在您想要使用服务账号代替交互登录时才需要提供。
      
      您可以使用 "~" 来扩展文件名，也可以使用环境变量，例如 `${RCLONE_CONFIG_DIR}`。

   --service-account-credentials
      服务账号凭据的 JSON blob。
      
      通常留空。
      仅在您想要使用服务账号代替交互登录时才需要提供。

   --team-drive
      共享驱动器（Team Drive）的 ID。

   --auth-owner-only
      只考虑经过身份验证用户拥有的文件。

   --use-trash
      将文件发送到回收站而不是永久删除。
      
      默认为 true，即将文件发送到回收站。
      使用 `--drive-use-trash=false` 以永久删除文件。

   --copy-shortcut-content
      服务器端复制快捷方式的内容，而不是复制快捷方式本身。
      
      在 doing server side copies 时，rclone 通常会将快捷方式复制为快捷方式。
      
      如果使用此标志，则在进行服务器端复制时，rclone 将复制快捷方式的内容
      而不是复制快捷方式本身。

   --skip-gdocs
      在所有列表中跳过 Google 文档。
      
      如果提供此参数，则 Google 文档在 rclone 中实际上会变得不可见。

   --skip-checksum-gphotos
      仅跳过 Google 照片和视频的 MD5 校验和。
      
      如果传输 Google 照片或视频时遇到校验和错误，请使用此标志。
      
      设置此标志将导致 Google 照片和视频返回空的 MD5 校验和。
      
      通过处于「照片」空间中的文件识别 Google 照片。
      
      校验和错误是由 Google 修改图像/视频但未更新校验和引起的。

   --shared-with-me
      仅显示与我共享的文件。
      
      让 rclone 操作您的「与我共享」文件夹（Google Drive 允许您访问其他人与您共享的文件和文件夹）。
      
      这适用于 "list"（lsd、lsl 等）和 "copy"（copy、sync 等）命令，
      以及所有其他命令。

   --trashed-only
      仅显示在垃圾箱中的文件。
      
      这将以原始目录结构显示已删除的文件。

   --starred-only
      仅显示被标记为星标的文件。

   --formats
      不推荐使用：请参阅 export_formats。

   --export-formats
      逗号分隔的首选格式列表，用于下载 Google 文档。

   --import-formats
      逗号分隔的首选格式列表，用于上传 Google 文档。

   --allow-import-name-change
      允许上传 Google 文档时更改文件类型。
      
      例如，可以将 file.doc 更改为 file.docx。这样每次同步时都会造成混淆和重新上传。

   --use-created-date
      使用文件创建日期而不是修改日期。
      
      在下载数据并希望使用创建日期而不是最后修改日期时非常有用。
      
      **警告**：此标志可能会产生一些意想不到的结果。
      
      在上传到您的 Drive 时，除非文件自创建以来未经修改，
      否则所有文件都将被覆盖。而在下载时，将发生相反的情况。
      可以使用 "--checksum" 标志避免出现这种副作用。
      
      此功能是为了保留由 Google 照片记录的照片拍摄日期而实施的。
      首先，您需要在 Google Drive 设置中勾选 "创建 Google 照片文件夹" 选项。
      然后，您可以将照片复制或移动到本地，并使用图像拍摄的日期
      （创建日期）设置为修改日期。

   --use-shared-date
      使用文件共享日期而不是修改日期。
      
      请注意，与 "--drive-use-created-date" 一样，
      此标志在上传/下载文件时可能会产生意想不到的结果。
      
      如果同时设置了此标志和 "--drive-use-created-date"，
      则将使用创建日期。

   --list-chunk
      列表块的大小，范围为 100-1000，0 表示禁用。

   --impersonate
      使用服务账号时，模拟此用户。

   --alternate-export
      不推荐使用：不再需要。

   --upload-cutoff
      切换到分块上传的大小。

   --chunk-size
      上传分块大小。
      
      必须为大于等于 256k 的 2 的幂次方。
      
      增大该值可以提高性能，但请注意每个分块都会在内存中缓冲一次。
      
      减小该值会减少内存使用量，但会降低性能。

   --acknowledge-abuse
      设置以允许下载返回 "cannotDownloadAbusiveFile" 错误的文件。
      
      如果下载文件返回错误 "This file has been identified
      as malware or spam and cannot be downloaded"，
      并显示错误代码 "cannotDownloadAbusiveFile"，
      则向 rclone 提供此标志以表示您承认下载该文件的风险，
      并且 rclone 将继续下载该文件。
      
      请注意，如果您使用的是服务账号，则需要 Manager
      权限（而不是 Content Manager）才能使此标志有效。
      如果服务账号没有正确的权限，Google 将会忽略此标志。

   --keep-revision-forever
      永久保留每个文件的新头修订版。

   --size-as-quota
      显示存储配额使用情况，而不是实际大小作为文件大小。
      
      显示文件的大小作为使用的存储配额。这是
      当前版本加上所有已设置为永久保存的旧版本。
      
      **警告**：该标志可能会产生一些意想不到的结果。
      
      不推荐在配置中设置该标志，建议在只进行
      rclone ls/lsl/lsf/lsjson 等操作时使用 --drive-size-as-quota 标志。
      
      如果您确实使用该标志进行同步（不推荐），
      那么您还需要使用 --ignore size。

   --v2-download-min-size
      如果对象较大，则使用 Drive v2 API 下载。

   --pacer-min-sleep
      API 调用之间的最小休眠时间。

   --pacer-burst
      允许的连续 API 调用次数。

   --server-side-across-configs
      允许对不同的 Drive 配置进行服务器端操作（如复制）。
      
      如果要在两个不同的 Google Drive 之间进行服务器端复制，这可能会很有用。
      请注意，默认情况下未启用此功能，因为它对两个配置中是否能够正常工作很难判断。

   --disable-http2
      禁用使用 http2 进行传输。
      
      Google Drive 后端目前存在一个未解决的问题，
      因此默认情况下禁用了 drive 的 HTTP/2 支持。
      但您可以在此启用它。在问题解决后，将删除此标志。
      
      请参阅：https://github.com/rclone/rclone/issues/3631
      
      

   --stop-on-upload-limit
      使上传限制错误成为致命错误。
      
      目前，每天只能向 Google Drive 上传 750 GiB 的数据（这是一个未记录的限制）。
      达到此限制时，Google Drive 会生成略有不同的错误消息。
      如果设置了此标志，它将导致这些错误成为致命错误。这些错误将停止进行中的同步。
      
      请注意，此检测依赖于 Google 不记录的错误消息，
      因此未来可能会失效。
      
      参阅：https://github.com/rclone/rclone/issues/3857
      

   --stop-on-download-limit
      使下载限制错误成为致命错误。
      
      目前，每天只能从 Google Drive 下载 10 TiB 的数据（这是一个未记录的限制）。
      达到此限制时，Google Drive 会生成略有不同的错误消息。
      如果设置了此标志，它将导致这些错误成为致命错误。这些错误将停止进行中的同步。
      
      请注意，此检测依赖于 Google 不记录的错误消息，
      因此未来可能会失效。
      

   --skip-shortcuts
      如果设置，则跳过快捷方式文件。
      
      通常，rclone 会解引用快捷方式文件，使其看起来像原始文件（请参阅「快捷方式」）。。
      如果设置了此标志，则 rclone 将完全忽略快捷方式文件。
      

   --skip-dangling-shortcuts
      如果设置，则跳过悬空的快捷方式文件。
      
      如果设置了此标志，则 rclone 在列表中将不显示任何悬空的快捷方式。
      

   --resource-key
      用于访问共享链接文件的资源密钥。
      
      如果您需要访问使用链接共享的文件，例如
      
          https://drive.google.com/drive/folders/XXX?resourcekey=YYY&usp=sharing
      
      那么您需要使用第一部分 "XXX" 作为 "root_folder_id"，
      并使用第二部分 "YYY" 作为 "resource_key"，否则在访问目录时会收到 404 找不到的错误。
      
      请参阅：https://developers.google.com/drive/api/guides/resource-keys
      
      此资源密钥要求仅适用于某些旧文件子集。
      
      还要注意，只需要在 Web 界面中打开该文件夹一次
      （使用您已使用 rclone 进行身份验证的用户）即可，此时则无需资源密钥。
      

   --encoding
      后端的编码方式。
      
      请参阅[概述中的编码部分](/overview/#encoding)以获取更多信息。


选项：
   --alternate-export                         不推荐使用：不再需要。 (默认值：false) [$ALTERNATE_EXPORT]
   --client-id value                          Google 应用程序客户端 ID [$CLIENT_ID]
   --client-secret value                      OAuth 客户端密钥 [$CLIENT_SECRET]
   --help, -h                                 显示帮助
   --scope value                              rclone 在请求从 Google Drive 获取访问权限时应使用的范围 [$SCOPE]
   --service-account-file value               服务账号凭据的 JSON 文件路径 [$SERVICE_ACCOUNT_FILE]

   高级选项

   --acknowledge-abuse                        设置以允许下载返回 "cannotDownloadAbusiveFile" 错误的文件。 (默认值：false) [$ACKNOWLEDGE_ABUSE]
   --allow-import-name-change                 允许上传 Google 文档时更改文件类型。 (默认值：false) [$ALLOW_IMPORT_NAME_CHANGE]
   --auth-owner-only                          只考虑经过身份验证用户拥有的文件。 (默认值：false) [$AUTH_OWNER_ONLY]
   --auth-url value                           认证服务器 URL [$AUTH_URL]
   --chunk-size value                         上传分块大小。（默认值："8Mi"）[$CHUNK_SIZE]
   --copy-shortcut-content                    服务器端复制快捷方式的内容，而不是复制快捷方式本身。 (默认值：false) [$COPY_SHORTCUT_CONTENT]
   --disable-http2                            禁用使用 http2 进行传输。 (默认值：true) [$DISABLE_HTTP2]
   --encoding value                           后端的编码方式。 (默认值："InvalidUtf8") [$ENCODING]
   --export-formats value                     逗号分隔的首选格式列表，用于下载 Google 文档。 (默认值："docx,xlsx,pptx,svg") [$EXPORT_FORMATS]
   --formats value                            不推荐使用：请参阅 export_formats。 [$FORMATS]
   --impersonate value                        使用服务账号时，模拟此用户。 [$IMPERSONATE]
   --import-formats value                     逗号分隔的首选格式列表，用于上传 Google 文档。 [$IMPORT_FORMATS]
   --keep-revision-forever                    永久保留每个文件的新头修订版。 (默认值：false) [$KEEP_REVISION_FOREVER]
   --list-chunk value                         列表块的大小，范围为 100-1000，0 表示禁用。 (默认值：1000) [$LIST_CHUNK]
   --pacer-burst value                        允许的连续 API 调用次数。 (默认值：100) [$PACER_BURST]
   --pacer-min-sleep value                    API 调用之间的最小休眠时间。 (默认值："100ms") [$PACER_MIN_SLEEP]
   --resource-key value                       用于访问共享链接文件的资源密钥。 [$RESOURCE_KEY]
   --root-folder-id value                     根文件夹的 ID。 [$ROOT_FOLDER_ID]
   --server-side-across-configs               允许对不同的 Drive 配置进行服务器端操作。 (默认值：false) [$SERVER_SIDE_ACROSS_CONFIGS]
   --service-account-credentials value        服务账号凭据的 JSON blob。 [$SERVICE_ACCOUNT_CREDENTIALS]
   --shared-with-me                           仅显示与我共享的文件。 (默认值：false) [$SHARED_WITH_ME]
   --size-as-quota                            显示存储配额使用情况，而不是实际大小作为文件大小。 (默认值：false) [$SIZE_AS_QUOTA]
   --skip-checksum-gphotos                    仅跳过 Google 照片和视频的 MD5 校验和。 (默认值：false) [$SKIP_CHECKSUM_GPHOTOS]
   --skip-dangling-shortcuts                  如果设置，则跳过悬空的快捷方式文件。 (默认值：false) [$SKIP_DANGLING_SHORTCUTS]
   --skip-gdocs                               在所有列表中跳过 Google 文档。 (默认值：false) [$SKIP_GDOCS]
   --skip-shortcuts                           如果设置，则跳过快捷方式文件。 (默认值：false) [$SKIP_SHORTCUTS]
   --starred-only                             仅显示被标记为星标的文件。 (默认值：false) [$STARRED_ONLY]
   --stop-on-download-limit                   使下载限制错误成为致命错误。 (默认值：false) [$STOP_ON_DOWNLOAD_LIMIT]
   --stop-on-upload-limit                     使上传限制错误成为致命错误。 (默认值：false) [$STOP_ON_UPLOAD_LIMIT]
   --team-drive value                         共享驱动器（Team Drive）的 ID。 [$TEAM_DRIVE]
   --token value                              OAuth 访问令牌（JSON 格式）。 [$TOKEN]
   --token-url value                          令牌服务器 URL。 [$TOKEN_URL]
   --trashed-only                             仅显示在垃圾箱中的文件。 (默认值：false) [$TRASHED_ONLY]
   --upload-cutoff value                      切换到分块上传的大小。（默认值："8Mi"）[$UPLOAD_CUTOFF]
   --use-created-date                         使用文件创建日期而不是修改日期。 (默认值：false) [$USE_CREATED_DATE]
   --use-shared-date                          使用文件共享日期而不是修改日期。 (默认值：false) [$USE_SHARED_DATE]
   --use-trash                                将文件发送到回收站而不是永久删除。 (默认值：true) [$USE_TRASH]
   --v2-download-min-size value               如果对象较大，则使用 Drive v2 API 下载。 (默认值："off") [$V2_DOWNLOAD_MIN_SIZE]

   常规选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}