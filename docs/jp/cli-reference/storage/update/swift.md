# OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

{% code fullWidth="true" %}
```
名称:
   singularity storage update swift - OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

使用法:
   singularity storage update swift [command options] <name|id>

说明:
   --env-auth
      从标准的OpenStack环境变量中获取Swift凭据。

      示例:
         | false | 在下一步中输入Swift凭据。
         | true  | 从环境变量中获取Swift凭据。
         |       | 如果使用此选项，其他字段请留空。

   --user
      登录用户名 (OS_USERNAME)。

   --key
      API密钥或密码 (OS_PASSWORD)。

   --auth
      服务器的身份验证URL (OS_AUTH_URL)。

      示例:
         | https://auth.api.rackspacecloud.com/v1.0     | Rackspace美国
         | https://lon.auth.api.rackspacecloud.com/v1.0 | Rackspace英国
         | https://identity.api.rackspacecloud.com/v2.0 | Rackspace v2
         | https://auth.storage.memset.com/v1.0         | Memset Memstore英国
         | https://auth.storage.memset.com/v2.0         | Memset Memstore英国v2
         | https://auth.cloud.ovh.net/v3                | OVH

   --user-id
      要登录的用户ID - 可选 - 大多数Swift系统使用用户名称，此字段留空 (v3认证) (OS_USER_ID)。

   --domain
      用户域 - 可选 (v3认证) (OS_USER_DOMAIN_NAME)。

   --tenant
      租户名称 - v1认证可选，否则需要此选项或tenant_id (OS_TENANT_NAME或OS_PROJECT_NAME)。

   --tenant-id
      租户ID - v1认证可选，否则需要此选项或tenant (OS_TENANT_ID)。

   --tenant-domain
      租户域 - 可选 (v3认证) (OS_PROJECT_DOMAIN_NAME)。

   --region
      区域名称 - 可选 (OS_REGION_NAME)。

   --storage-url
      存储URL - 可选 (OS_STORAGE_URL)。

   --auth-token
      来自替代身份验证的认证令牌 - 可选 (OS_AUTH_TOKEN)。

   --application-credential-id
      应用凭证ID (OS_APPLICATION_CREDENTIAL_ID)。

   --application-credential-name
      应用凭证名称 (OS_APPLICATION_CREDENTIAL_NAME)。

   --application-credential-secret
      应用凭证密钥 (OS_APPLICATION_CREDENTIAL_SECRET)。

   --auth-version
      身份验证版本 - 可选 - 如果您的身份验证URL没有版本信息，则设置为(1,2,3) (ST_AUTH_VERSION)。

   --endpoint-type
      选择服务目录中的终端类型 (OS_ENDPOINT_TYPE)。

      示例:
         | public   | 公共 (默认值，如果不确定请选择此项)
         | internal | 内部 (使用内部服务网络)
         | admin    | 管理员

   --leave-parts-on-error
      如果为true，则在失败时避免调用中止上传。
      
      对于在不同会话之间恢复上传，应将其设置为true。

   --storage-policy
      创建新容器时要使用的存储策略。
      
      在创建新容器时应用指定的存储策略。策略之后将无法更改。允许的配置值及其含义取决于Swift存储提供商。

      示例:
         | <unset> | 默认值
         | pcs     | OVH 公共云存储
         | pca     | OVH 公共云存储存档

   --chunk-size
      大小超过此值的文件将被分块存储到_segments容器中。
      
      大小超过此值的文件将被分块存储到_segments容器中。默认值为5 GiB，是其最大值。

   --no-chunk
      在进行流式上传时，不要对文件进行分块。
      
      在进行流式上传（例如使用rcat或mount）时，设置此标志将导致Swift后端不上传分块文件。
      
      这将将最大上传大小限制为5 GiB。然而，非分块文件更易处理并具有MD5SUM。
      
      在执行普通复制操作时，Rclone仍会对大于chunk_size的文件进行分块。

   --no-large-objects
      禁用对静态和动态大对象的支持。
      
      Swift无法透明地存储大于5 GiB的文件。有两种方案可以实现，即静态或动态大对象，但是API不允许rclone在不进行对象HEAD的情况下确定文件是静态还是动态大对象。由于这两种方案需要以不同的方式处理，因此rclone必须为对象发出HEAD请求，例如在读取校验和时。
      
      当设置`no_large_objects`时，rclone会假定没有存储静态或动态大对象。这意味着它可以停止执行额外的HEAD调用，这反过来极大地提高了性能，特别是在使用`--checksum`进行swift到swift传输时。
      
      设置此选项意味着`no_chunk`并且不会以分块上传任何文件，因此大于5 GiB的文件将无法上传并且会导致失败。
      
      如果设置此选项并且存在静态或动态大对象，那么对它们的哈希计算将是不正确的。下载将成功，但其他操作（例如删除和复制）将失败。

   --encoding
      后端的编码方式。
      
      有关详细信息，请参阅概述中的[编码部分](/overview/#encoding)。

选项:
   --application-credential-id value      应用凭证ID (OS_APPLICATION_CREDENTIAL_ID)。 [$APPLICATION_CREDENTIAL_ID]
   --application-credential-name value    应用凭证名称 (OS_APPLICATION_CREDENTIAL_NAME)。 [$APPLICATION_CREDENTIAL_NAME]
   --application-credential-secret value  应用凭证密钥 (OS_APPLICATION_CREDENTIAL_SECRET)。 [$APPLICATION_CREDENTIAL_SECRET]
   --auth value                           服务器的身份验证URL (OS_AUTH_URL)。 [$AUTH]
   --auth-token value                     来自替代身份验证的认证令牌 - 可选 (OS_AUTH_TOKEN)。 [$AUTH_TOKEN]
   --auth-version value                   身份验证版本 - 可选 - 如果您的身份验证URL没有版本信息，则设置为(1,2,3) (ST_AUTH_VERSION)。 (默认值: 0) [$AUTH_VERSION]
   --domain value                         用户域 - 可选 (v3认证) (OS_USER_DOMAIN_NAME) [$DOMAIN]
   --endpoint-type value                  选择服务目录中的终端类型 (OS_ENDPOINT_TYPE)。 (默认值: "public") [$ENDPOINT_TYPE]
   --env-auth                             从标准的OpenStack环境变量中获取Swift凭据。 (默认值: false) [$ENV_AUTH]
   --help, -h                             显示帮助信息
   --key value                            API密钥或密码 (OS_PASSWORD)。 [$KEY]
   --region value                         区域名称 - 可选 (OS_REGION_NAME)。 [$REGION]
   --storage-policy value                 创建新容器时要使用的存储策略。 [$STORAGE_POLICY]
   --storage-url value                    存储URL - 可选 (OS_STORAGE_URL)。 [$STORAGE_URL]
   --tenant value                         租户名称 - v1认证可选，否则需要此选项或tenant_id (OS_TENANT_NAME或OS_PROJECT_NAME)。 [$TENANT]
   --tenant-domain value                  租户域 - 可选 (v3认证) (OS_PROJECT_DOMAIN_NAME)。 [$TENANT_DOMAIN]
   --tenant-id value                      租户ID - v1认证可选，否则需要此选项或tenant (OS_TENANT_ID)。 [$TENANT_ID]
   --user value                           登录用户名 (OS_USERNAME)。 [$USER]
   --user-id value                        要登录的用户ID - 可选 - 大多数Swift系统使用用户名称，此字段留空 (v3认证) (OS_USER_ID)。 [$USER_ID]

   Advanced

   --chunk-size value      大小超过此值的文件将被分块存储到_segments容器中。 (默认值: "5Gi") [$CHUNK_SIZE]
   --encoding value        后端的编码方式。 (默认值: "Slash,InvalidUtf8") [$ENCODING]
   --leave-parts-on-error  如果为true，则在失败时避免调用中止上传。 (默认值: false) [$LEAVE_PARTS_ON_ERROR]
   --no-chunk              在进行流式上传时，不要对文件进行分块。 (默认值: false) [$NO_CHUNK]
   --no-large-objects      禁用对静态和动态大对象的支持 (默认值: false) [$NO_LARGE_OBJECTS]

```
{% endcode %}