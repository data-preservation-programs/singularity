# OpenStack Swift（Rackspace Cloud Files，Memset Memstore，OVH）

```sh
名称:
   singularity datasource add swift - OpenStack Swift（Rackspace Cloud Files，Memset Memstore，OVH）

用法:
   singularity datasource add swift [命令选项] <数据集名称> <源路径>

描述:
   --swift-application-credential-id
      应用凭据ID（OS_APPLICATION_CREDENTIAL_ID）。

   --swift-application-credential-name
      应用凭据名称（OS_APPLICATION_CREDENTIAL_NAME）。

   --swift-application-credential-secret
      应用凭据密钥（OS_APPLICATION_CREDENTIAL_SECRET）。

   --swift-auth
      服务器的认证URL（OS_AUTH_URL）。

      示例:
         | https://auth.api.rackspacecloud.com/v1.0     | Rackspace美国
         | https://lon.auth.api.rackspacecloud.com/v1.0 | Rackspace英国
         | https://identity.api.rackspacecloud.com/v2.0 | Rackspace v2
         | https://auth.storage.memset.com/v1.0         | Memset Memstore英国
         | https://auth.storage.memset.com/v2.0         | Memset Memstore英国v2
         | https://auth.cloud.ovh.net/v3                | OVH

   --swift-auth-token
      来自备用认证的认证令牌-可选（OS_AUTH_TOKEN）。

   --swift-auth-version
      认证版本-可选-如果您的认证URL没有版本，则设置为（1,2,3）（ST_AUTH_VERSION）。

   --swift-chunk-size
      文件大小超过此大小时，将将文件切割成_segments容器。
      
      文件大小超过此大小时，将文件切割成_segments容器。默认值为5 GiB，即其最大值。

   --swift-domain
      用户域名-可选（v3认证）（OS_USER_DOMAIN_NAME）

   --swift-encoding
      后端的编码。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --swift-endpoint-type
      从服务目录中选择的终端点类型（OS_ENDPOINT_TYPE）。

      示例:
         | public   | 公共（默认值，如果不确定，请选择此选项）
         | internal | 内部（使用内部服务网络）
         | admin    | 管理员

   --swift-env-auth
      从标准OpenStack形式的环境变量中获取Swift凭据。

      示例:
         | false | 在下一步中输入Swift凭据。
         | true  | 从环境变量获取Swift凭据。
                 | 如果使用此选项，请将其他字段留空。

   --swift-key
      API密钥或密码（OS_PASSWORD）。

   --swift-leave-parts-on-error
      如果为true，则在失败时避免调用中止上传。
      
      对于跨不同会话恢复上传，应将其设置为true。

   --swift-no-chunk
      在流式上传期间不分块文件。
      
      在进行流式上传时（例如使用rcat或mount），将此标志设置为true将导致Swift后端不上传分块文件。
      
      这将限制最大上传大小为5 GiB。但是，非分块文件更容易处理并且有MD5SUM。
      
      在执行正常的复制操作时，Rclone仍然会对大于chunk_size的文件进行分块。

   --swift-no-large-objects
      禁用对静态和动态大对象的支持
      
      Swift无法透明地存储大小超过5 GiB的文件。有两种方案可实现这一点，静态大对象和动态大对象，并且API允许rclone在不进行对象的HEAD请求的情况下确定文件是静态大对象还是动态大对象。由于它们需要以不同的方式处理，这意味着rclone必须为对象发出HEAD请求，例如用于读取校验和时。
      
      当设置`no_large_objects`时，rclone将假设没有存储任何静态或动态大对象。这意味着它可以停止执行额外的HEAD调用，从而大大提高性能，特别是在使用`--checksum`进行swift到swift传输时。
      
      设置此选项意味着`no_chunk`并且不会上传任何文件的分块，因此大于5 GiB的文件将在上传时失败。
      
      如果设置了此选项，并且存在静态或动态大对象，则此选项将为它们提供不正确的哈希值。下载将成功，但其他操作（如Remove和Copy）将失败。
      

   --swift-region
      区域名称-可选（OS_REGION_NAME）。

   --swift-storage-policy
      创建新容器时要使用的存储策略。
      
      创建新容器时，应用指定的存储策略。之后无法更改策略。允许的配置值及其含义取决于您的Swift存储提供商。

      示例:
         | <unset> | 默认值
         | pcs     | OVH公共云存储
         | pca     | OVH公共云存储档案

   --swift-storage-url
      存储URL-可选（OS_STORAGE_URL）。

   --swift-tenant
      租户名称-对于v1认证，可选，否则此项或租户id为必填项（OS_TENANT_NAME或OS_PROJECT_NAME）。

   --swift-tenant-domain
      租户域名-可选（v3认证）（OS_PROJECT_DOMAIN_NAME）。

   --swift-tenant-id
      租户id-对于v1认证，选填，否则此项或租户为必填项（OS_TENANT_ID）。

   --swift-user
      登录用户名（OS_USERNAME）。

   --swift-user-id
      登录用户id-可选-大多数Swift系统使用用户，并留空此项（v3认证）（OS_USER_ID）。


选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险]将数据集导出为CAR文件后删除数据集的文件。 (default: false)
   --rescan-interval value  最后一次成功扫描后，当经过此间隔时自动重新扫描源目录（default: 禁用）
   --scanning-state value   设置初始扫描状态（default: 准备就绪）

   swift的选项

   --swift-application-credential-id value      应用凭据ID（OS_APPLICATION_CREDENTIAL_ID）。[$SWIFT_APPLICATION_CREDENTIAL_ID]
   --swift-application-credential-name value    应用凭据名称（OS_APPLICATION_CREDENTIAL_NAME）。[$SWIFT_APPLICATION_CREDENTIAL_NAME]
   --swift-application-credential-secret value  应用凭据密钥（OS_APPLICATION_CREDENTIAL_SECRET）。[$SWIFT_APPLICATION_CREDENTIAL_SECRET]
   --swift-auth value                           服务器的认证URL（OS_AUTH_URL）。[$SWIFT_AUTH]
   --swift-auth-token value                     来自备用认证的认证令牌-可选（OS_AUTH_TOKEN）。[$SWIFT_AUTH_TOKEN]
   --swift-auth-version value                   认证版本-可选-如果您的认证URL没有版本，则设置为（1,2,3）（ST_AUTH_VERSION）。 (default: "0") [$SWIFT_AUTH_VERSION]
   --swift-chunk-size value                     文件大小超过此大小时，将将文件切割成_segments容器。 (default: "5Gi") [$SWIFT_CHUNK_SIZE]
   --swift-domain value                         用户域名-可选（v3认证）（OS_USER_DOMAIN_NAME）[$SWIFT_DOMAIN]
   --swift-encoding value                       后端的编码。 (default: "Slash,InvalidUtf8") [$SWIFT_ENCODING]
   --swift-endpoint-type value                  从服务目录中选择的终端点类型（OS_ENDPOINT_TYPE）。 (default: "public") [$SWIFT_ENDPOINT_TYPE]
   --swift-env-auth value                       从标准OpenStack形式的环境变量中获取Swift凭据。 (default: "false") [$SWIFT_ENV_AUTH]
   --swift-key value                            API密钥或密码（OS_PASSWORD）。[$SWIFT_KEY]
   --swift-leave-parts-on-error value           如果为true，则在失败时避免调用中止上传。 (default: "false") [$SWIFT_LEAVE_PARTS_ON_ERROR]
   --swift-no-chunk value                       在流式上传期间不分块文件。 (default: "false") [$SWIFT_NO_CHUNK]
   --swift-no-large-objects value               禁用对静态和动态大对象的支持 (default: "false") [$SWIFT_NO_LARGE_OBJECTS]
   --swift-region value                         区域名称-可选（OS_REGION_NAME）。[$SWIFT_REGION]
   --swift-storage-policy value                 创建新容器时要使用的存储策略。[$SWIFT_STORAGE_POLICY]
   --swift-storage-url value                    存储URL-可选（OS_STORAGE_URL）。[$SWIFT_STORAGE_URL]
   --swift-tenant value                         租户名称-对于v1认证，可选，否则此项或租户id为必填项（OS_TENANT_NAME或OS_PROJECT_NAME）。[$SWIFT_TENANT]
   --swift-tenant-domain value                  租户域名-可选（v3认证）（OS_PROJECT_DOMAIN_NAME）。[$SWIFT_TENANT_DOMAIN]
   --swift-tenant-id value                      租户id-对于v1认证，选填，否则此项或租户为必填项（OS_TENANT_ID）。[$SWIFT_TENANT_ID]
   --swift-user value                           登录用户名（OS_USERNAME）。[$SWIFT_USER]
   --swift-user-id value                        登录用户id-可选-大多数Swift系统使用用户，并留空此项（v3认证）（OS_USER_ID）。[$SWIFT_USER_ID]

```
