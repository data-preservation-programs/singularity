# Microsoft OneDrive

{% code fullWidth="true" %}
```
名称：
   singularity datasource add onedrive - Microsoft OneDrive

用法：
   singularity datasource add onedrive [command options] <数据集名称> <源路径>

描述：
   --onedrive-access-scopes
      设置rclone请求的范围。
      
      选择或手动输入一个由空格分隔的自定义范围列表，rclone应该请求。
      

      例子：
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access | 所有资源的读写权限
         | Files.Read Files.Read.All Sites.Read.All offline_access                                     | 所有资源的只读权限
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All offline_access                | 所有资源的读写权限，不能浏览SharePoint站点。 
                                                                                                | 等同于将disable_site_permission设置为true

   --onedrive-auth-url
      授权服务器URL。
      
      留空以使用提供程序的默认值。

   --onedrive-chunk-size
      用于上传文件的块大小，必须是320k（327,680字节）的倍数。
      
      超过此大小将进行分块——必须是320k（327,680字节）的倍数，并且
      不应超过250M（262,144,000字节），否则可能遇到 \"Microsoft.SharePoint.Client.InvalidClientQueryException: The request message is too big.\"
      请注意，块将被缓冲到内存中。

   --onedrive-client-id
      OAuth客户端ID。
      
      通常留空。

   --onedrive-client-secret
      OAuth客户端密钥。
      
      通常留空。

   --onedrive-disable-site-permission
      禁用Sites.Read.All权限的请求。
      
      如果设置为true，则在
      配置驱动ID时将无法搜索SharePoint站点，因为rclone不会请求Sites.Read.All权限。
      如果您的组织没有为应用程序分配Sites.Read.All权限，并且您的组织不允许用户同意应用
      程序权限请求，请将其设置为true。

   --onedrive-drive-id
      要使用的驱动器的ID。

   --onedrive-drive-type
      驱动器的类型（个人 | 商务 | 文档库）。

   --onedrive-encoding
      后端的编码。
      
      有关详细信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --onedrive-expose-onenote-files
      设置使OneNote文件在目录列表中显示。
      
      默认情况下，rclone在目录列表中隐藏OneNote文件，因为
      “打开”和“更新”等操作在它们上不起作用。但是这
      行为也可能阻止您删除它们。如果您想
      删除OneNote文件或以其他方式希望它们显示在目录中
      列表，请设置此选项。

   --onedrive-hash-type
      指定后端中使用的哈希。
      
      这指定要使用的哈希类型。如果设置为“auto”，它将使用
      默认哈希类型是QuickXorHash。
      
      在rclone 1.62之前，Onedrive的默认哈希是SHA1哈希。（对于onedrive Business）。
      从rclone 1.62起，默认是对所有onedrive类型使用QuickXorHash。如果需要SHA1哈希，则相应地设置此选项。
      
      从2023年7月起，QuickXorHash将是唯一可用的哈希
      针对OneDrive for Business和OneDriver个人。
      
      这可以设置为“none”以不使用任何哈希。
      
      如果所请求的哈希在对象上不存在，则
      它将作为空字符串返回，rclone会将其视为缺少的哈希。
      

      例子：
         | auto     | rclone选择最佳哈希
         | quickxor | QuickXor
         | sha1     | SHA1
         | sha256   | SHA256
         | crc32    | CRC32
         | none     | 不使用任何哈希

   --onedrive-link-password
      设置由链接命令创建的链接的密码。
      
      撰写此文档时，这仅适用于付费的OneDrive个人帐户。
      

   --onedrive-link-scope
      设置由链接命令创建的链接的范围。

      例子：
         | anonymous    | 任何拥有链接的人都可访问，无需登录。
                        | 这可能包括组织以外的人。
                        | 管理员可能已禁用匿名链接支持。
         | organization | 任何登录到您的组织（租户）的人都可以使用该链接获取访问权限。
                        | 仅适用于OneDrive for Business和SharePoint。

   --onedrive-link-type
      设置由链接命令创建的链接的类型。

      例子：
         | view  | 创建一个只读链接指向该项目。
         | edit  | 创建一个读写链接指向该项目。
         | embed | 创建一个可嵌入的链接指向该项目。

   --onedrive-list-chunk
      列表块的大小。

   --onedrive-no-versions
      修改操作时删除所有版本。
      
      当rclone覆盖现有文件上传新文件时，Onedrive for Business会创建版本，并在设置更改修改时间时。
      
      这些版本会占用配额空间。
      
      此标志在文件上传和设置
      修改时间后检查版本，并删除除最后一个版本以外的所有版本。
      
      **NB** 目前Onedrive个人无法删除版本，因此请勿在此处使用此标志。

   --onedrive-region
      选择OneDrive的国家云区域。

      例子：
         | global | Microsoft全球云
         | us     | Microsoft美国政府云
         | de     | 德国Microsoft云
         | cn     | 在中国由Vnet Group运营的Azure和Office 365

   --onedrive-root-folder-id
      根目录的ID。
      
      通常不需要，但在特殊情况下，您可能会
      知道您希望访问的文件夹ID，但无法获得
      通过路径遍历到达那里。

   --onedrive-server-side-across-configs
      允许服务器端操作（例如复制）在不同的onedrive配置之间工作。
      
      这仅在您将文件从两个 OneDrive * Personal * 驱动器之间复制并且
      要复制的文件已经在它们之间共享时才有效。在其他情况下，rclone将
      退回到正常的复制（速度稍慢）。

   --onedrive-token
      OAuth访问令牌，以JSON形式提供。

   --onedrive-token-url
      令牌服务器URL。
      
      留空以使用提供程序的默认值。


选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作]在将数据集导出为CAR文件后，删除数据集的文件。（默认：false）
   --rescan-interval value  在上次成功扫描后，当此间隔时间过去时，自动重新扫描源目录（默认：禁用）
   --scanning-state value   设置初始扫描状态（默认：ready）

   onedrive的选项

   --onedrive-access-scopes value               设置rclone请求的范围。（默认：“Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access”） [$ONEDRIVE_ACCESS_SCOPES]
   --onedrive-auth-url value                    授权服务器URL。[$ONEDRIVE_AUTH_URL]
   --onedrive-chunk-size value                  用于上传文件的块大小，必须是320k（327,680字节）的倍数。（默认：“10Mi”）[$ONEDRIVE_CHUNK_SIZE]
   --onedrive-client-id value                   OAuth客户端ID。[$ONEDRIVE_CLIENT_ID]
   --onedrive-client-secret value               OAuth客户端密钥。[$ONEDRIVE_CLIENT_SECRET]
   --onedrive-drive-id value                    要使用的驱动器的ID。[$ONEDRIVE_DRIVE_ID]
   --onedrive-drive-type value                  驱动器的类型（个人 | 商务 | 文档库）。[$ONEDRIVE_DRIVE_TYPE]
   --onedrive-encoding value                    后端的编码。（默认：“Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot”）[$ONEDRIVE_ENCODING]
   --onedrive-expose-onenote-files value        设置使OneNote文件在目录列表中显示。（默认：“false”）[$ONEDRIVE_EXPOSE_ONENOTE_FILES]
   --onedrive-hash-type value                   指定后端中使用的哈希。（默认：“auto”）[$ONEDRIVE_HASH_TYPE]
   --onedrive-link-password value               设置由链接命令创建的链接的密码。[$ONEDRIVE_LINK_PASSWORD]
   --onedrive-link-scope value                  设置由链接命令创建的链接的范围。（默认：“anonymous”）[$ONEDRIVE_LINK_SCOPE]
   --onedrive-link-type value                   设置由链接命令创建的链接的类型。（默认：“view”）[$ONEDRIVE_LINK_TYPE]
   --onedrive-list-chunk value                  列表块的大小。（默认：“1000”）[$ONEDRIVE_LIST_CHUNK]
   --onedrive-no-versions value                 在修改操作中删除所有版本。（默认：“false”）[$ONEDRIVE_NO_VERSIONS]
   --onedrive-region value                      选择OneDrive的国家云区域。（默认：“global”）[$ONEDRIVE_REGION]
   --onedrive-root-folder-id value              根目录的ID。[$ONEDRIVE_ROOT_FOLDER_ID]
   --onedrive-server-side-across-configs value  允许服务器端操作（例如复制）在不同的onedrive配置之间工作。（默认：“false”）[$ONEDRIVE_SERVER_SIDE_ACROSS_CONFIGS]
   --onedrive-token value                       OAuth访问令牌，以JSON形式提供。[$ONEDRIVE_TOKEN]
   --onedrive-token-url value                   令牌服务器URL。[$ONEDRIVE_TOKEN_URL]

```
{% endcode %}