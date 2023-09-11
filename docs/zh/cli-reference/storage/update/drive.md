# Google Drive

{% code fullWidth="true" %}
```
名称:
   singularity storage update drive - Google Drive

用法:
   singularity storage update drive [命令选项] <名称|id>

说明:
   --client-id
      Google应用程序客户端ID
      强烈建议您设置自己的客户端ID。
      有关如何创建自己的ID，请参见https://rclone.org/drive/#making-your-own-client-id。
      如果留空，将使用内部密钥，性能较低。

   --client-secret
      OAuth客户端秘钥。
      
      通常留空即可。

   --token
      OAuth访问令牌（作为JSON blob）。

   --auth-url
      授权服务器URL。
      
      若要使用提供程序的默认设置，请留空。

   --token-url
      令牌服务器URL。
      
      若要使用提供程序的默认设置，请留空。

   --scope
      rclone在请求访问Drive时应使用的范围。

      示例：
         | drive                   | 完全访问所有文件，不包括应用数据文件夹。
         | drive.readonly          | 只读访问文件元数据和文件内容。
         | drive.file              | 仅访问由rclone创建的文件。
         |                         | 这些文件在Drive网站上可见。
         |                         | 当用户取消对应用的授权时，文件授权将被撤销。
         | drive.appfolder         | 允许对应用数据文件夹进行读写操作。
         |                         | 这在Drive网站上不可见。
         | drive.metadata.readonly | 仅允许对文件元数据进行只读访问，
         |                         | 不允许交互式阅读或下载文件内容。

   --root-folder-id
      根文件夹的ID。
      通常留空。
      若要访问“计算机”文件夹（请参见文档），或者要使rclone使用
      非根文件夹作为起始点，请填写此字段。
      

   --service-account-file
      服务账号凭据JSON文件路径。
      
      通常留空。
      只有在希望使用服务账号而不是交互式登录时才需要。

   --service-account-credentials
      服务账号凭据JSON blob。
      
      通常留空。
      只有在希望使用服务账号而不是交互式登录时才需要。

   --team-drive
      共享Drive的ID。

   --auth-owner-only
      仅考虑归您拥有的文件。

   --use-trash
      将文件发送到回收站而非永久删除。
      
      默认为true，即将文件发送到回收站。
      使用`--drive-use-trash=false`可永久删除文件。

   --copy-shortcut-content
      服务器端复制快捷方式的内容，而不是复制快捷方式本身。
      
      在进行服务器端复制时，rclone通常会将快捷方式复制为快捷方式。
      
      如果使用此标志，则在进行服务器端复制时，rclone将复制快捷方式的内容
      而不是快捷方式本身。

   --skip-gdocs
      在所有列表中跳过Google文档。
      
      如果提供此标志，则rclone实际上无法看到gdocs。

   --skip-checksum-gphotos
      仅在Google照片和视频上跳过MD5校验和。
      
      如果在传输Google照片或视频时出现校验和错误，请使用此选项。
      
      设置此标志将导致Google照片和视频返回空白的MD5校验和。
      
      Google照片通过位于“照片”空间中。
      
      损坏的校验和是由于Google对图像/视频进行修改，
      但未更新校验和而导致的。

   --shared-with-me
      仅显示与我共享的文件。
      
      指示rclone操作您的“与我共享”文件夹（Google Drive上允许您访问其他人与您共享的文件和文件夹的位置）。
      
      这适用于“list”（lsd、lsl等）和“copy”命令（copy、sync等），
      以及所有其他命令。

   --trashed-only
      仅显示处于回收站中的文件。
      
      这将显示回收站中文件的原始目录结构。

   --starred-only
      仅显示标记为星标的文件。

   --formats
      已弃用：请参见export_formats。

   --export-formats
      逗号分隔的首选格式列表，用于下载Google文档。

   --import-formats
      逗号分隔的首选格式列表，用于上传Google文档。

   --allow-import-name-change
      允许上传Google文档时更改文件类型。
      
      例如，将file.doc更改为file.docx。每次同步，都会产生混乱并重新上传。

   --use-created-date
      使用文件的创建日期，而不是修改日期。
      
      在下载数据且希望使用创建日期代替最后修改日期时有用。
      
      **警告**：此标志可能会产生一些意外后果。
      
      当上传到您的Drive时，除非文件自创建以来未被修改，否则将覆盖所有文件。
      下载时将发生相反的情况。通过使用"--checksum"标志可以避免这种副作用。
      
      此功能用于保留由Google照片记录的照片捕获日期。
      首先，您需要在google drive设置中选中“创建一个Google照片文件夹”选项。
      然后，您可以在本地复制或移动照片，并将图片拍摄日期（创建日期）设置为修改日期。

   --use-shared-date
      使用文件共享日期而不是修改日期。
      
      请注意，与"--drive-use-created-date"一样，此标志可能会产生意外后果
      上传/下载文件时。
      
      如果同时设置此标志和"--drive-use-created-date"，则使用创建日期。

   --list-chunk
      列表块的大小，100-1000，0表示禁用。

   --impersonate
      使用服务账号时，模拟此用户。

   --alternate-export
      已弃用：不再需要。

   --upload-cutoff
      切换到分块上传的阈值。

   --chunk-size
      上传块的大小。
      
      必须是大于等于256 KB的2的幂。
      
      增加此值将改善性能，但请注意，每个块都会在内存中缓冲一次。
      
      减小此值将减少内存使用，但会降低性能。

   --acknowledge-abuse
      设置为允许下载返回"cannotDownloadAbusiveFile"错误的文件。
      
      如果下载文件返回错误消息"This file has been identified
      as malware or spam and cannot be downloaded"，错误代码为
      "cannotDownloadAbusiveFile"，请使用此标志告知rclone你意识到
      下载该文件的风险，rclone将继续下载该文件。
      
      请注意，如果您使用的是服务账号，则需要Manager
      权限（而不是Content Manager）才能使此标志起作用。如果SA
      没有正确的权限，Google将忽略此标志。

   --keep-revision-forever
      永久保留每个文件的新版副本。

   --size-as-quota
      将大小显示为存储配额使用情况，而不是实际大小。
      
      将文件的大小显示为已使用的存储配额。这是
      当前版本和已设置为永久保留的任何旧版本。

      **警告**：此标志可能会产生一些意外后果。
      
      不建议在配置文件中设置此标志-建议仅在进行rclone ls/lsl/lsf/lsjson等操作时使用
      `--drive-size-as-quota`标志。

      如果要对同步操作使用此标志（不建议），您将需要使用--ignore大小字符串。

   --v2-download-min-size
      如果对象较大，则使用drive v2 API进行下载。

   --pacer-min-sleep
      API调用之间的最小休眠时间。

   --pacer-burst
      允许的API调用次数而无需休眠。

   --server-side-across-configs
      允许服务器端操作（例如复制）在不同的drive配置之间工作。
      
      如果希望在两个不同的Google Drive之间进行服务器端复制，
      则此选项非常有效。请注意，此选项默认情况下是未启用的，
      因为无法轻松判断它是否适用于任何两个配置。

   --disable-http2
      禁止使用http2进行drive。
      
      目前，Google Drive的后端存在无法解决的问题与HTTP/2。
      因此，默认情况下禁用了Drive后端的HTTP/2，但可以在此处重新启用。
      解决此问题后，将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/3631

   --stop-on-upload-limit
      使上传限制错误成为致命错误。
      
      在撰写本文时，每天仅能上传750 GiB的数据到Google Drive（这是一个未记录的限制）。
      达到此限制后，Google Drive会生成稍有不同的错误消息。
      当设置此标志时，将导致这些错误变成致命错误。
      这些错误会停止进行中的同步。
      
      请注意，此检测依赖于Google未记录的错误消息字符串，
      因此可能会在将来发生错误。
      
      参见：https://github.com/rclone/rclone/issues/3857

   --stop-on-download-limit
      使下载限制错误成为致命错误。
      
      在撰写本文时，每天只能从Google Drive下载10 TiB的数据（这是一个未记录的限制）。
      达到此限制后，Google Drive会生成稍有不同的错误消息。
      当设置此标志时，将导致这些错误变成致命错误。
      这些错误会停止进行中的同步。
      
      请注意，此检测依赖于Google未记录的错误消息字符串，
      因此可能会在将来发生错误。

   --skip-shortcuts
      如果设置，跳过快捷方式文件。
      
      通常情况下，rclone会解引用快捷方式文件，使其看起来像原始文件（请参阅[快捷方式部分](#shortcuts)）。
      如果设置此标志，则rclone完全会忽略快捷方式文件。

   --skip-dangling-shortcuts
      如果设置，请跳过挂起的快捷方式文件。
      
      如果设置此标志，则rclone不会在列表中显示任何挂起的快捷方式。

   --resource-key
      访问共享链接文件的资源密钥。
      
      如果您需要访问通过链接共享的文件，请使用以下一部分作为“root_folder_id”的值
      
          https://drive.google.com/drive/folders/XXX?resourcekey=YYY&usp=sharing
      
      然后，您将需要使用第一部分“XXX”作为“root_folder_id”，使用第二部分“YYY”作为“resource_key”。
      否则，当尝试访问目录时，您将收到404错误。
      
      参见：https://developers.google.com/drive/api/guides/resource-keys
      
      此资源密钥要求仅适用于旧文件的子集。
      
      还要注意，只需在Web界面中打开该文件夹一次（使用已通过rclone进行身份验证的用户），
      看起来就足够，无需添加资源密钥。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概览中的编码部分](/overview/#encoding)。


选项:
   --alternate-export            已弃用：不再需要。（默认：false）[$ALTERNATE_EXPORT]
   --client-id value             Google应用程序客户端ID[$CLIENT_ID]
   --client-secret value         OAuth客户端秘钥[$CLIENT_SECRET]
   --help, -h                    显示帮助
   --scope value                 rclone在请求Drive访问权限时应使用的范围[$SCOPE]
   --service-account-file value  服务账号凭据JSON文件路径[$SERVICE_ACCOUNT_FILE]

   高级选项

   --acknowledge-abuse                  设置为允许下载返回cannotDownloadAbusiveFile的文件。（默认：false）[$ACKNOWLEDGE_ABUSE]
   --allow-import-name-change           允许上传Google文档时更改文件类型。（默认：false）[$ALLOW_IMPORT_NAME_CHANGE]
   --auth-owner-only                    仅考虑归您拥有的文件。（默认：false）[$AUTH_OWNER_ONLY]
   --auth-url value                     授权服务器URL[$AUTH_URL]
   --chunk-size value                   上传块的大小。（默认："8Mi"）[$CHUNK_SIZE]
   --copy-shortcut-content              服务器端复制快捷方式的内容，而不是复制快捷方式本身。（默认：false）[$COPY_SHORTCUT_CONTENT]
   --disable-http2                      禁止使用http2进行drive。（默认：true）[$DISABLE_HTTP2]
   --encoding value                     后端的编码方式。（默认："InvalidUtf8"）[$ENCODING]
   --export-formats value               逗号分隔的首选格式列表，用于下载Google文档。（默认："docx,xlsx,pptx,svg"）[$EXPORT_FORMATS]
   --formats value                      已弃用：请参见export_formats。[$FORMATS]
   --impersonate value                  使用服务账号时，模拟此用户[$IMPERSONATE]
   --import-formats value               逗号分隔的首选格式列表，用于上传Google文档[$IMPORT_FORMATS]
   --keep-revision-forever              永久保留每个文件的新版副本。（默认：false）[$KEEP_REVISION_FOREVER]
   --list-chunk value                   列表块的大小，100-1000，0表示禁用。（默认：1000）[$LIST_CHUNK]
   --pacer-burst value                  允许的API调用次数而无需休眠。（默认：100）[$PACER_BURST]
   --pacer-min-sleep value              API调用之间的最小休眠时间。（默认："100ms"）[$PACER_MIN_SLEEP]
   --resource-key value                 访问通过链接共享的文件的资源密钥[$RESOURCE_KEY]
   --root-folder-id value               根文件夹的ID[$ROOT_FOLDER_ID]
   --server-side-across-configs         允许服务器端操作（例如复制）在不同的drive配置之间工作。（默认：false）[$SERVER_SIDE_ACROSS_CONFIGS]
   --service-account-credentials value  服务账号凭据JSON blob[$SERVICE_ACCOUNT_CREDENTIALS]
   --shared-with-me                     仅显示与我共享的文件。（默认：false）[$SHARED_WITH_ME]
   --size-as-quota                      将大小显示为存储配额使用情况，而不是实际大小。（默认：false）[$SIZE_AS_QUOTA]
   --skip-checksum-gphotos              仅在Google照片和视频上跳过MD5校验和。（默认：false）[$SKIP_CHECKSUM_GPHOTOS]
   --skip-dangling-shortcuts            如果设置，请跳过挂起的快捷方式文件。（默认：false）[$SKIP_DANGLING_SHORTCUTS]
   --skip-gdocs                         在所有列表中跳过Google文档。（默认：false）[$SKIP_GDOCS]
   --skip-shortcuts                     如果设置，请跳过快捷方式文件。（默认：false）[$SKIP_SHORTCUTS]
   --starred-only                       仅显示标记为星标的文件。（默认：false）[$STARRED_ONLY]
   --stop-on-download-limit             使下载限制错误成为致命错误。（默认：false）[$STOP_ON_DOWNLOAD_LIMIT]
   --stop-on-upload-limit               使上传限制错误成为致命错误。（默认：false）[$STOP_ON_UPLOAD_LIMIT]
   --team-drive value                   共享Drive的ID[$TEAM_DRIVE]
   --token value                        OAuth访问令牌（作为JSON blob）[$TOKEN]
   --token-url value                    令牌服务器URL[$TOKEN_URL]
   --trashed-only                       仅显示处于回收站中的文件。（默认：false）[$TRASHED_ONLY]
   --upload-cutoff value                切换到分块上传的阈值。（默认："8Mi"）[$UPLOAD_CUTOFF]
   --use-created-date                   使用文件的创建日期，而不是修改日期。（默认：false）[$USE_CREATED_DATE]
   --use-shared-date                    使用文件共享日期而不是修改日期。（默认：false）[$USE_SHARED_DATE]
   --use-trash                          将文件发送到回收站而非永久删除。（默认：true）[$USE_TRASH]
   --v2-download-min-size value         如果对象较大，则使用drive v2 API进行下载。（默认："off"）[$V2_DOWNLOAD_MIN_SIZE]

```
{% endcode %}