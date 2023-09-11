# Microsoft OneDrive

{% code fullWidth="true" %}
```
名称:
   singularity storage create onedrive - Microsoft OneDrive

用法:
   singularity storage create onedrive [命令选项] [参数...]

描述:
   --client-id
      OAuth 客户端 ID。
      
      通常留空。

   --client-secret
      OAuth 客户端密钥。
      
      通常留空。

   --token
      OAuth 访问令牌的 JSON 字符串。

   --auth-url
      认证服务器 URL。
      
      为空表示使用提供程序的默认值。

   --token-url
      令牌服务器 URL。
      
      为空表示使用提供程序的默认值。

   --region
      选择 OneDrive 的国家云区域。

      示例:
         | global | Microsoft 云全球版
         | us     | Microsoft 云 - 美国政府
         | de     | Microsoft 云 - 德国
         | cn     | Azure 和 Office 365 在中国的 Vnet Group 运营

   --chunk-size
      上传文件的块大小 - 必须是 320k 的倍数 (327,680 字节)。
      
      超过此大小的文件将被分块上传 - 块大小必须是 320k 的倍数 (327,680 字节)，
      不应超过 250M (262,144,000 字节)，否则可能会遇到 "Microsoft.SharePoint.Client.InvalidClientQueryException: The request message is too big." 错误。
      注意，块将被缓冲到内存中。

   --drive-id
      要使用的驱动器的 ID。

   --drive-type
      驱动器的类型 (个人 | 商业 | 文档库)。

   --root-folder-id
      根文件夹的 ID。
      
      通常不需要，但在特殊情况下，您可能知道要访问的文件夹 ID，但无法通过路径遍历到达。

   --access-scopes
      设置 rclone 请求的范围。
      
      选择或手动输入自定义范围的空格分隔列表，rclone应请求这些范围。

      示例:
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access | 对所有资源的读写访问权限
         | Files.Read Files.Read.All Sites.Read.All offline_access                                     | 对所有资源的只读访问权限
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All offline_access                | 对所有资源的读写访问权限，但无法浏览 SharePoint 站点
         |                                                                                             | 禁用站点权限与将 disable_site_permission 设置为 true 相同

   --disable-site-permission
      禁用 Sites.Read.All 权限的请求。
      
      如果设置为 true，则在配置驱动器 ID 时将无法搜索 SharePoint 站点，
      因为 rclone 不会请求 Sites.Read.All 权限。
      如果您的组织未向应用程序授予 Sites.Read.All 权限，并且您的组织不允许用户自行同意应用程序权限请求，
      请将其设置为 true。

   --expose-onenote-files
      设置以在目录列表中显示 OneNote 文件。
      
      默认情况下，rclone 在目录列表中隐藏 OneNote 文件，因为
      "打开" 和 "更新" 等操作无法在这些文件上执行。但是，
      这种行为也可能阻止您删除它们。如果要
      删除 OneNote 文件或以其他方式希望它们在目录中显示，请设置此选项。

   --server-side-across-configs
      允许服务器端操作 (如复制) 在不同的 onedrive 配置中工作。
      
      只有在复制两个 OneDrive *个人* 驱动器之间且要复制的文件已在它们之间共享的情况下，此功能才有效。
      否则，rclone 将退回到正常的复制操作 (速度稍慢)。

   --list-chunk
      列出块的大小。

   --no-versions
      修改操作时删除所有版本。
      
      当 rclone 上传新文件覆盖现有文件并设置修改时间时，OneDrive for Business 会创建版本。
      
      这些版本将占用配额空间。
      
      此标志在文件上传和设置修改时间后检查版本，并删除除最新版本之外的所有版本。
      
      **注意**，OneDrive 个人目前无法删除版本，因此请勿在个人版中使用此标志。
      

   --link-scope
      设置 link 命令创建的链接的范围。

      示例:
         | anonymous    | 允许具有链接的任何人访问，无需登录。
         |              | 这可能包括您组织之外的人。
         |              | 管理员可能禁用匿名链接支持。
         | organization | 允许您组织 (租户) 已登录的任何人使用链接获取访问权限。
         |              | 仅适用于 OneDrive for Business 和 SharePoint。

   --link-type
      设置 link 命令创建的链接的类型。

      示例:
         | view  | 创建一个只读链接。
         | edit  | 创建一个读写链接。
         | embed | 创建一个可嵌入的链接。

   --link-password
      设置 link 命令创建的链接的密码。
      
      撰写本文时，此功能仅适用于付费的 OneDrive 个人帐户。
      

   --hash-type
      指定后端使用的哈希类型。
      
      这指定正在使用的哈希类型。如果设置为 "auto"，则将使用
      默认的哈希类型，即 QuickXorHash。
      
      在 rclone 1.62 及之前的版本中，默认使用 SHA1 哈希用于 Onedrive
      个人版。对于 1.62 及更高版本，默认使用 QuickXorHash 用于所有类型的
      onedrive。如果需要 SHA1 哈希，则相应设置此选项。
      
      从 2023 年 7 月开始，QuickXorHash 将是
      OneDrive for Business 和 OneDriver 个人版的唯一可用哈希。
      
      可以将其设置为 "none"，以不使用任何哈希。
      
      如果请求的哈希在对象上不存在，则会将其
      返回为空字符串，rclone 将其视为丢失的哈希。
      

      示例:
         | auto     | rclone 选择最佳哈希
         | quickxor | QuickXor
         | sha1     | SHA1
         | sha256   | SHA256
         | crc32    | CRC32
         | none     | 无 - 不使用任何哈希

   --encoding
      后端的编码。
      
      详见[概述中的编码部分](/overview/#encoding)了解更多信息。


选项:
   --client-id value      OAuth 客户端 ID。[$CLIENT_ID]
   --client-secret value  OAuth 客户端密钥。[$CLIENT_SECRET]
   --help, -h             显示帮助信息
   --region value         选择 OneDrive 的国家云区域。 (默认值: "global") [$REGION]

   高级

   --access-scopes value         设置 rclone 请求的范围。 (默认值: "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access") [$ACCESS_SCOPES]
   --auth-url value              认证服务器 URL。[$AUTH_URL]
   --chunk-size value            上传文件的块大小 - 必须是 320k 的倍数 (327,680 字节)。 (默认值: "10Mi") [$CHUNK_SIZE]
   --disable-site-permission     禁用 Sites.Read.All 权限的请求。 (默认值: false) [$DISABLE_SITE_PERMISSION]
   --drive-id value              要使用的驱动器的 ID。[$DRIVE_ID]
   --drive-type value            驱动器的类型 (personal | business | documentLibrary)。[$DRIVE_TYPE]
   --encoding value              后端的编码。 (默认值: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --expose-onenote-files        设置以在目录列表中显示 OneNote 文件。 (默认值: false) [$EXPOSE_ONENOTE_FILES]
   --hash-type value             指定后端使用的哈希类型。 (默认值: "auto") [$HASH_TYPE]
   --link-password value         设置 link 命令创建的链接的密码。[$LINK_PASSWORD]
   --link-scope value            设置 link 命令创建的链接的范围。 (默认值: "anonymous") [$LINK_SCOPE]
   --link-type value             设置 link 命令创建的链接的类型。 (默认值: "view") [$LINK_TYPE]
   --list-chunk value            列出块的大小。 (默认值: 1000) [$LIST_CHUNK]
   --no-versions                 修改操作时删除所有版本。 (默认值: false) [$NO_VERSIONS]
   --root-folder-id value        根文件夹的 ID。[$ROOT_FOLDER_ID]
   --server-side-across-configs  允许服务器端操作 (如复制) 在不同的 onedrive 配置中工作。 (默认值: false) [$SERVER_SIDE_ACROSS_CONFIGS]
   --token value                 OAuth 访问令牌的 JSON 字符串。[$TOKEN]
   --token-url value             令牌服务器 URL。[$TOKEN_URL]

   通用

   --name value  存储的名称 (默认值: 自动生成)
   --path value  存储的路径

```
{% endcode %}