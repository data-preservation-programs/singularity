# OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

{% code fullWidth="true" %}
```
名称:
   singularity datasource add swift - OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

用法:
   singularity datasource add swift [命令选项] <dataset_name> <source_path>

描述:
   --swift-auth-version
      AuthVersion - 可选的 - 如果您的认证 URL 没有版本号 (ST_AUTH_VERSION) ，请设置为 (1,2,3)。

   --swift-endpoint-type
      选择服务目录中的端点类型 (OS_ENDPOINT_TYPE)。

      示例:
         | public   | 公共 (默认，如果不确定，请选择此选项)
         | internal | 内部 (使用内部服务网络)
         | admin    | 管理员

   --swift-encoding
      后端的编码格式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --swift-application-credential-secret
      应用凭据密钥 (OS_APPLICATION_CREDENTIAL_SECRET)。

   --swift-user
      用户名 (OS_USERNAME)。

   --swift-tenant
      租户名称 - 对于 v1 认证来说是可选项，否则需要此选项或 tenant_id (OS_TENANT_NAME 或 OS_PROJECT_NAME)。

   --swift-region
      区域名称 - 可选项 (OS_REGION_NAME)。

   --swift-application-credential-name
      应用凭据名称 (OS_APPLICATION_CREDENTIAL_NAME)。

   --swift-no-chunk
      流式上传文件时不将文件分块。
      
      当进行流式上传（例如使用 rcat 或 mount）时，将此标志设置将导致 Swift 后端不上传分块文件。
      
      这将限制最大上传大小为 5 GiB。但是非分块文件更容易处理，并且具有 MD5SUM。
      
      在普通复制操作时，当您将大于 chunk_size 的文件时，rclone 仍会将其分块。

   --swift-env-auth
      在标准 OpenStack 格式的环境变量中获取 swift 凭据。

      示例:
         | false | 在下一步中输入 swift 凭据。
         | true  | 从环境变量中获取 swift 凭据。
                 | 如果使用此方法，请将其他字段留空。

   --swift-tenant-domain
      租户域名 - 可选项 (v3 认证) (OS_PROJECT_DOMAIN_NAME)。

   --swift-storage-url
      存储 URL - 可选项 (OS_STORAGE_URL)。

   --swift-auth-token
      来自替代身份验证的 Auth Token - 可选项（OS_AUTH_TOKEN）。

   --swift-application-credential-id
      应用凭据 ID (OS_APPLICATION_CREDENTIAL_ID)。

   --swift-chunk-size
      文件大小超过该大小将被分块为 _segments 容器。
      
      文件大小超过此范围时，将把文件分块为 _segments 容器。
      其默认值为 5 GiB，这是最大值。

   --swift-no-large-objects
      禁用静态和动态大对象支持。
      
      Swift 无法透明地存储大于 5 GiB 的文件。有两种方案可以完成这一点，静态或动态大对象，而 API 没有允许 rclone 在不执行对象的 HEAD 的情况下确定文件是静态还是动态大对象。由于这些需要进行不同的处理，这意味着 rclone 必须针对例如读取校验和时发出对象的 HEAD 请求。
      
      当设置 `no_large_objects` 时，rclone 将假设没有存储静态或动态大对象。这意味着它可以停止执行额外的 HEAD 调用，这反过来会显著提高性能，特别是在使用 `--checksum` 进行 swift 到 swift 的转移时。
      
      设置此选项会意味着 `no_chunk` 以及不上传任何文件块，因此大于 5 GiB 的文件仅会上传失败。
      
      如果您设置此选项，而实际上存在静态或动态大对象，则这将为它们提供不正确的哈希值。下载将成功，但其他操作，例如“删除”和“复制”将失败。

   --swift-domain
      用户域 - 可选项 (v3 认证) (OS_USER_DOMAIN_NAME)

   --swift-auth
      服务器的身份验证 URL (OS_AUTH_URL)。

      示例:
         | https://auth.api.rackspacecloud.com/v1.0     | Rackspace US
         | https://lon.auth.api.rackspacecloud.com/v1.0 | Rackspace UK
         | https://identity.api.rackspacecloud.com/v2.0 | Rackspace v2
         | https://auth.storage.memset.com/v1.0         | Memset Memstore UK
         | https://auth.storage.memset.com/v2.0         | Memset Memstore UK v2
         | https://auth.cloud.ovh.net/v3                | OVH

   --swift-user-id
      用户 ID - 可选项 - 大多数 swift 系统使用用户并将其留空 (v3 认证) (OS_USER_ID)。

   --swift-tenant-id
      租户 ID - 对于 v1 认证是可选的，否则需要此选项或 tenant，否则 (OS_TENANT_ID)。

   --swift-leave-parts-on-error
      避免在失败时调用“中止上传”。
      
      对于跨不同会话恢复上传，必须将其设置为 true。

   --swift-storage-policy
      创建新容器时要使用的存储策略。
      
      在创建新容器时，应用指定的存储策略。之后无法更改此策略。允许的配置值及其含义取决于您的 Swift 存储提供者。

      示例:
         | <unset> | 默认值
         | pcs     | OVH Public Cloud Storage
         | pca     | OVH Public Cloud Archive

   --swift-key
      API 密钥或密码 (OS_PASSWORD)。


选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险操作] 导出数据集后删除数据集的文件。 (默认为 false)
   --rescan-interval value  当自上次成功扫描以来经过了此时间间隔时，自动重新扫描源目录 (默认: 禁用)

   Swift 的选项

   --swift-application-credential-id value      应用凭据 ID (OS_APPLICATION_CREDENTIAL_ID)。 [$SWIFT_APPLICATION_CREDENTIAL_ID]
   --swift-application-credential-name value    应用凭据名称 (OS_APPLICATION_CREDENTIAL_NAME)。 [$SWIFT_APPLICATION_CREDENTIAL_NAME]
   --swift-application-credential-secret value  应用凭据密钥 (OS_APPLICATION_CREDENTIAL_SECRET)。 [$SWIFT_APPLICATION_CREDENTIAL_SECRET]
   --swift-auth value                           服务器的身份验证 URL (OS_AUTH_URL)。 [$SWIFT_AUTH]
   --swift-auth-token value                     来自替代身份验证的 Auth Token - 可选项（OS_AUTH_TOKEN）。 [$SWIFT_AUTH_TOKEN]
   --swift-auth-version value                   AuthVersion - 可选的 - 如果您的认证 URL 没有版本号 (ST_AUTH_VERSION) ，请设置为 (1,2,3)。 (默认为 “0”) [$SWIFT_AUTH_VERSION]
   --swift-chunk-size value                     文件大小超过该大小将被分块为 _segments 容器。 (默认为 “5Gi”) [$SWIFT_CHUNK_SIZE]
   --swift-domain value                         用户域 - 可选项 (v3 认证) (OS_USER_DOMAIN_NAME) [$SWIFT_DOMAIN]
   --swift-encoding value                       后端的编码格式。 (默认为 “Slash,InvalidUtf