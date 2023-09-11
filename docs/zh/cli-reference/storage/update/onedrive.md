# Microsoft OneDrive

{% code fullWidth="true" %}
```
名称:
   singularity存储更新onedrive - 微软OneDrive

用法:
   singularity存储更新onedrive [命令选项] <名称|ID>

说明:
   --client-id
      OAuth客户端ID。
      
      通常为空。

   --client-secret
      OAuth客户端密钥。
      
      通常为空。

   --token
      OAuth访问令牌，以JSON格式提供。

   --auth-url
      认证服务器URL。
      
      如果不填写，默认使用提供者的默认值。

   --token-url
      令牌服务器URL。
      
      如果不填写，默认使用提供者的默认值。

   --region
      选择OneDrive的国家云区域。

      示例:
         | global | 微软全球云
         | us     | 美国政府用微软云
         | de     | 微软德国云
         | cn     | 在中国由微软Vnet Group运营的Azure和Office 365

   --chunk-size
      上传文件的块大小 - 必须是320k（327,680字节）的倍数。
      
      超过此大小的文件将被分块处理 - 必须是320k（327,680字节）的倍数，并且不应超过250M（262,144,000字节），否则可能会遇到 "Microsoft.SharePoint.Client.InvalidClientQueryException: The request message is too big."。
      请注意，分块数据将被缓存到内存中。

   --drive-id
      要使用的驱动器的ID。

   --drive-type
      驱动器类型（personal | business | documentLibrary）。

   --root-folder-id
      根文件夹的ID。
      
      通常情况下不需要，但在特殊情况下，您可能知道要访问的文件夹的ID，但无法通过路径遍历访问。
      

   --access-scopes
      设置rclone请求的范围。
      
      选择或手动输入一个自定义的以空格分隔的范围列表，以便rclone请求。
      

      示例:
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access | 对所有资源进行读写访问
         | Files.Read Files.Read.All Sites.Read.All offline_access                                     | 只读访问所有资源
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All offline_access                | 对所有资源进行读写访问，但无法浏览SharePoint站点。 
         |                                                                                             | 相当于将disable_site_permission设置为true时的操作

   --disable-site-permission
      禁用对Sites.Read.All权限的请求。
      
      如果设置为true，您将无法在配置驱动器ID时搜索SharePoint站点，因为rclone将不会请求Sites.Read.All权限。
      如果您的组织未将Sites.Read.All权限分配给应用程序，并且您的组织禁止用户自行同意应用程序权限请求，请将其设置为true。

   --expose-onenote-files
      设置以使OneNote文件显示在目录列表中。
      
      默认情况下，rclone在目录列表中隐藏OneNote文件，因为无法对它们执行“打开”和“更新”等操作。但是，此行为可能也会防止您删除它们。如果要删除OneNote文件或以其他方式希望它们出现在目录列表中，请设置此选项。

   --server-side-across-configs
      允许在不同的OneDrive配置之间进行服务器端操作（例如复制）。
      
      这仅适用于在两个OneDrive *个人*驱动器之间进行复制，并且要复制的文件已在它们之间共享。在其他情况下，rclone将回退到正常的复制（速度稍慢一些）。

   --list-chunk
      列出块的大小。

   --no-versions
      在修改操作中删除所有版本。
      
      当rclone覆盖现有文件并设置修改时间时，OneDrive for Business会创建版本。
      
      这些版本将占用配额空间。
      
      此标志在文件上传和设置修改时间后检查版本，并删除除最后一个版本之外的所有版本。
      
      **注意** OneDrive个人目前无法删除版本，因此请勿在此处使用此标志。
      

   --link-scope
      设置link命令创建的链接范围。

      示例:
         | anonymous    | 链接对任何人都可访问，无需登录。
         |              | 这可能包括组织外的人员。
         |              | 管理员可能已禁用匿名链接支持。
         | organization | 组织（租户）中已登录的任何人都可以使用该链接访问。
         |              | 仅在OneDrive for Business和SharePoint中可用。

   --link-type
      设置link命令创建的链接类型。

      示例:
         | view  | 创建到项目的只读链接。
         | edit  | 创建到项目的读写链接。
         | embed | 创建到项目的可嵌入链接。

   --link-password
      设置link命令创建的链接的密码。
      
      在编写本文时，此功能仅适用于OneDrive个人付费帐户。
      

   --hash-type
      指定后端使用的哈希。
      
      这指定了正在使用的哈希类型。如果设置为"auto"，它将使用默认哈希，即QuickXorHash。
      
      在rclone 1.62之前，默认情况下，OneDrive个人使用SHA1哈希。对于1.62及更高版本，默认情况下所有OneDrive类型都使用QuickXorHash。如果需要SHA1哈希，请相应地设置此选项。
      
      从2023年7月起，QuickXorHash将是OneDrive for Business和OneDriver Personal的唯一可用哈希。
      
      可以将其设置为"none"以不使用任何哈希。
      
      如果所请求的哈希在对象上不存在，它将返回一个空字符串，rclone将将其视为缺少的哈希。
      

      示例:
         | auto     | Rclone选择最佳哈希
         | quickxor | QuickXor
         | sha1     | SHA1
         | sha256   | SHA256
         | crc32    | CRC32
         | none     | 无 - 不使用任何哈希

   --encoding
      后端的编码。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。


选项:
   --client-id value      OAuth客户端ID。[$CLIENT_ID]
   --client-secret value  OAuth客户端密钥。[$CLIENT_SECRET]
   --help, -h             显示帮助信息
   --region value         选择OneDrive的国家云区域。（默认值: "global"）[$REGION]

   高级选项

   --access-scopes value         设置rclone请求的范围。（默认值: "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access"）[$ACCESS_SCOPES]
   --auth-url value              认证服务器URL。[$AUTH_URL]
   --chunk-size value            上传文件的块大小 - 必须是320k（327,680字节）的倍数。（默认值: "10Mi"）[$CHUNK_SIZE]
   --disable-site-permission     禁用对Sites.Read.All权限的请求。（默认值: false）[$DISABLE_SITE_PERMISSION]
   --drive-id value              要使用的驱动器的ID。[$DRIVE_ID]
   --drive-type value            驱动器类型（personal | business | documentLibrary）。[$DRIVE_TYPE]
   --encoding value              后端的编码。（默认值: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot"）[$ENCODING]
   --expose-onenote-files        设置以使OneNote文件显示在目录列表中。（默认值: false）[$EXPOSE_ONENOTE_FILES]
   --hash-type value             指定后端使用的哈希。（默认值: "auto"）[$HASH_TYPE]
   --link-password value         设置link命令创建的链接的密码。[$LINK_PASSWORD]
   --link-scope value            设置link命令创建的链接的范围。（默认值: "anonymous"）[$LINK_SCOPE]
   --link-type value             设置link命令创建的链接的类型。（默认值: "view"）[$LINK_TYPE]
   --list-chunk value            列出块的大小。（默认值: 1000）[$LIST_CHUNK]
   --no-versions                 在修改操作中删除所有版本。（默认值: false）[$NO_VERSIONS]
   --root-folder-id value        根文件夹的ID。[$ROOT_FOLDER_ID]
   --server-side-across-configs  允许在不同的OneDrive配置之间进行服务器端操作（例如复制）。（默认值: false）[$SERVER_SIDE_ACROSS_CONFIGS]
   --token value                 OAuth访问令牌，以JSON格式提供。[$TOKEN]
   --token-url value             令牌服务器URL。[$TOKEN_URL]

```
{% endcode %}