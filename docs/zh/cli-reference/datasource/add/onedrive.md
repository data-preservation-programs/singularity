# Microsoft OneDrive

{% code fullWidth="true" %}
```

名称:
   singularity 数据源添加 onedrive - Microsoft OneDrive

用法:
   singularity datasource add onedrive [命令选项] <数据集名称> <源路径>

说明:
   --onedrive-expose-onenote-files
      设置此标记以在目录列表中显示 OneNote 文件。

      默认情况下，rclone 不会在目录列表中显示 OneNote 文件，因为像“打开”和“更新”之类的操作在其上不起作用。但是，此行为也可能阻止您删除它们。如果要删除 OneNote 文件或以其他方式使其出现在目录列表中，请设置此选项。

   --onedrive-server-side-across-configs
      允许跨不同 onedrive 配置（Personal、Business）使用服务端操作（例如，复制）。

      此仅在从一个 OneDrive *Personal* 驱动器复制到另一个 OneDrive 驱动器，并且要复制的文件已经在它们之间共享的情况下起作用。在其他情况下，rclone 将退回到普通的复制（稍慢一些）。

   --onedrive-no-versions
      在修改操作时删除所有版本。

      当 rclone 通过覆盖现有文件上传新文件并设置修改时间时，Onedrive for business 会创建版本。

      这些版本占用了配额的空间。

      此标记在文件上传和设置修改时间后检查版本，并删除除最后一个版本之外的所有版本。

      **注意**：Onedrive personal 当前不能删除版本，因此请不要在那里使用此标记。

   --onedrive-client-id
      OAuth 客户端 ID。

      通常留空。

   --onedrive-client-secret
      OAuth 客户端密钥。

      通常留空。

   --onedrive-token-url
      令牌服务器 URL。

      留空以使用提供程序的默认值。

   --onedrive-chunk-size
      用于上传文件的分块大小，必须是 320k（327,680 字节）的倍数。

      在此大小以上的文件将被分块，必须是 320k（327,680 字节）的倍数，并且不应超过 250M（262,144,000 字节），否则您可能会遇到 "Microsoft.SharePoint.Client.InvalidClientQueryException: The request message is too big." 问题。

      请注意，块将被缓冲到内存中。

   --onedrive-token
      OAuth 访问令牌，格式为 JSON。

   --onedrive-drive-id
      要使用的驱动器的 ID。

   --onedrive-hash-type
      指定后端使用的哈希。

      此选项指定正在使用的哈希类型。如果设置为“auto”，它将使用默认哈希，即 QuickXorHash。

      在 rclone 1.62 之前，Onedrive Personal 默认使用 SHA1 哈希。自 1.62 起，默认情况下对所有 onedrive 类型使用 QuickXorHash。如果需要 SHA1 哈希，则相应地设置此选项。

      从 2023 年 7 月起，QuickXorHash 将是 OneDrive for Business 和 OneDriver Personal 的唯一可用哈希。

      可以将此设置为“none”以不使用任何哈希。

      如果请求的哈希在对象上不存在，则将其作为空字符串返回，rclone 将其视为缺少哈希。

      示例：

      | auto     | 让 rclone 选择最佳哈希
      | quickxor | QuickXor
      | sha1     | SHA1
      | sha256   | SHA256
      | crc32    | CRC32
      | none     | 不使用任何哈希

   --onedrive-list-chunk
      列表块的大小。

   --onedrive-link-scope
      设置“link”命令创建的链接的范围。

      示例：

      | anonymous    | 拥有链接的人都可以访问，无需登录。 
                     | 这可能包括组织外的人。 
                     | 管理员可能会禁用匿名链接支持。
      | organization | 任何登录到您的组织（租户）中的人都可以使用链接访问。 
                     | 仅在 OneDrive for Business 和 SharePoint 中可用。

   --onedrive-link-type
      设置“link”命令创建的链接的类型。

      示例：

      | view  | 创建到项目的只读链接。
      | edit  | 创建到项目的读写链接。
      | embed | 创建到项目的可嵌入链接。

   --onedrive-region
      选择 OneDrive 的国家云区域。

      示例：

      | global | Microsoft Cloud 全球版。
      | us     | 为美国政府的 Microsoft Cloud。
      | de     | 德国的 Microsoft Cloud。
      | cn     | 在中国由 Vnet Group 经营的 Azure 和 Office 365。

   --onedrive-root-folder-id
      根文件夹的 ID。

      通常不需要，但在特殊情况下，您可能知道要访问的文件夹 ID，但无法通过路径遍历到达那里。

   --onedrive-access-scopes
      设置 rclone 应请求的作用域。

      选择或手动输入空格分隔的自定义列表，其中包含所有 rclone 应请求的作用域。

      示例：

      | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access | 对所有资源进行读写访问。
      | Files.Read Files.Read.All Sites.Read.All offline_access                                     | 对所有资源进行只读访问。
      | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All offline_access                | 对所有资源进行读写访问，但不能浏览 SharePoint 网站。
                                                                                                  | 与禁用站点权限相同。

   --onedrive-disable-site-permission
      禁用 Sites.Read.All 权限的请求。

      如果设置为 true，则在配置驱动器 ID 时，您将不再能够搜索 SharePoint 站点，因为 rclone 不会请求 Sites.Read.All 权限。如果组织未向应用程序分配 Sites.Read.All 权限，并且组织不允许用户自行同意应用程序权限请求，则将其设置为 true。

   --onedrive-auth-url
      认证服务器 URL。

      留空以使用提供程序的默认值。

   --onedrive-drive-type
      驱动器的类型（personal、business、documentLibrary）。

   --onedrive-link-password
      设置“link”命令创建的链接的密码。

      在编写本文时，仅适用于 OneDrive 个人付费账户。

   --onedrive-encoding
      后端的编码。

      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

选项:
   --help, -h  显示帮助

   数据准