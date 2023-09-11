# OpenStack Swift（Rackspace Cloud Files、Memset Memstore、OVH）

{% code fullWidth="true" %}
```
NAME:
   singularity storage create swift - OpenStack Swift（Rackspace Cloud Files, Memset Memstore, OVH）

USAGE:
   singularity storage create swift [command options] [arguments...]

DESCRIPTION:
   --env-auth
      使用标准的OpenStack形式从环境变量获取Swift凭证。

      示例:
         | false | 在后续步骤中输入Swift凭证。
         | true  | 从环境变量获取Swift凭证。
         |       | 如果使用此选项，请将其他字段留空。

   --user
      登录用户名（OS_USERNAME）。

   --key
      API密钥或密码（OS_PASSWORD）。

   --auth
      服务器的身份验证URL（OS_AUTH_URL）。

      示例:
         | https://auth.api.rackspacecloud.com/v1.0     | Rackspace（美国）
         | https://lon.auth.api.rackspacecloud.com/v1.0 | Rackspace（英国）
         | https://identity.api.rackspacecloud.com/v2.0 | Rackspace v2
         | https://auth.storage.memset.com/v1.0         | Memset Memstore（英国）
         | https://auth.storage.memset.com/v2.0         | Memset Memstore（英国）v2
         | https://auth.cloud.ovh.net/v3                | OVH

   --user-id
      用户ID - 可选 - 大多数Swift系统使用用户，并将此字段留空（v3认证）（OS_USER_ID）。

   --domain
      用户域 - 可选（v3认证）（OS_USER_DOMAIN_NAME）。

   --tenant
      租户名称 - 可选（v1认证）；否则需要此字段或tenant_id（OS_TENANT_NAME或OS_PROJECT_NAME）。

   --tenant-id
      租户ID - 可选（v1认证）；否则需要此字段或租户（OS_TENANT_ID）。

   --tenant-domain
      租户域 - 可选（v3认证）（OS_PROJECT_DOMAIN_NAME）。

   --region
      区域名称 - 可选（OS_REGION_NAME）。

   --storage-url
      存储URL - 可选（OS_STORAGE_URL）。

   --auth-token
      来自备用身份验证的身份验证令牌 - 可选（OS_AUTH_TOKEN）。

   --application-credential-id
      应用凭证ID（OS_APPLICATION_CREDENTIAL_ID）。

   --application-credential-name
      应用凭证名称（OS_APPLICATION_CREDENTIAL_NAME）。

   --application-credential-secret
      应用凭证密钥（OS_APPLICATION_CREDENTIAL_SECRET）。

   --auth-version
      身份验证版本 - 可选 - 如果您的身份验证URL没有版本，则设置为（1,2,3）（ST_AUTH_VERSION）。

   --endpoint-type
      从服务目录中选择的端点类型（OS_ENDPOINT_TYPE）。

      示例:
         | public   | 公共（默认值，如果不确定，选择此选项）
         | internal | 内部（使用内部服务网络）
         | admin    | 管理员

   --leave-parts-on-error
      如果发生错误，则避免调用中止上传。

      对于在不同会话间恢复上传，应将其设置为true。

   --storage-policy
      创建新容器时要使用的存储策略。

      当创建新容器时，将应用指定的存储策略。之后无法更改策略。允许的配置值及其含义取决于您的Swift存储提供程序。

      示例:
         | <unset> | 默认值
         | pcs     | OVH公有云存储
         | pca     | OVH公有云存储档案

   --chunk-size
      将文件分块到_segments容器中的大小限制。

      将文件分块到_segments容器中的大小限制。默认值为5 GiB，这是其最大值。

   --no-chunk
      在流式传输上传期间不分块文件。

      在进行流式上传时（例如使用rcat或mount），设置此标志将导致Swift后端不上传分块文件。

      这将将最大上传大小限制为5 GiB。但是非分块文件更容易处理，并且有MD5SUM。

      Rclone在执行普通复制操作时仍然会对大于chunk_size的文件进行分块。

   --no-large-objects
      禁用对静态和动态大对象的支持。

      Swift无法透明地存储大于5 GiB的文件。这方面有两种方法，即静态或动态大对象，而API无法使rclone确定文件是静态大对象还是动态大对象，除非对该对象进行HEAD请求。因为这些对象需要以不同方式处理，所以这意味着rclone必须针对对象发出HEAD请求，例如在读取校验和时。

      当设置“no_large_objects”时，rclone将假设没有存储静态或动态大对象。这意味着它可以停止执行额外的HEAD调用，从而大大提高性能，特别是在使用`--checksum`进行swift到swift传输时。

      设置此选项意味着`no_chunk`以及不会以分块方式上传文件，因此大于5 GiB的文件将在上传时发生错误。

      如果设置此选项且存在静态或动态大对象，则其哈希值将不正确。下载将成功，但其他操作（如删除和复制）将失败。

   --encoding
      后端的编码形式。

      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。


OPTIONS:
   --application-credential-id value      应用凭证ID（OS_APPLICATION_CREDENTIAL_ID）。[$APPLICATION_CREDENTIAL_ID]
   --application-credential-name value    应用凭证名称（OS_APPLICATION_CREDENTIAL_NAME）。[$APPLICATION_CREDENTIAL_NAME]
   --application-credential-secret value  应用凭证密钥（OS_APPLICATION_CREDENTIAL_SECRET）。[$APPLICATION_CREDENTIAL_SECRET]
   --auth value                           服务器的身份验证URL（OS_AUTH_URL）。[$AUTH]
   --auth-token value                     来自备用身份验证的身份验证令牌 - 可选（OS_AUTH_TOKEN）。[$AUTH_TOKEN]
   --auth-version value                   身份验证版本 - 可选 - 如果您的身份验证URL没有版本，则设置为（1,2,3）（ST_AUTH_VERSION）。（默认值：0）[$AUTH_VERSION]
   --domain value                         用户域 - 可选（v3认证）（OS_USER_DOMAIN_NAME）[$DOMAIN]
   --endpoint-type value                  从服务目录中选择的端点类型（OS_ENDPOINT_TYPE）。（默认值："public"）[$ENDPOINT_TYPE]
   --env-auth                             使用标准的OpenStack形式从环境变量获取Swift凭证。[默认值：false] [$ENV_AUTH]
   --help, -h                             显示帮助
   --key value                            API密钥或密码（OS_PASSWORD）。[$KEY]
   --region value                         区域名称 - 可选（OS_REGION_NAME）。[$REGION]
   --storage-policy value                 创建新容器时要使用的存储策略。[$STORAGE_POLICY]
   --storage-url value                    存储URL - 可选（OS_STORAGE_URL）。[$STORAGE_URL]
   --tenant value                         租户名称 - 可选（v1认证）；否则需要此字段或tenant_id（OS_TENANT_NAME或OS_PROJECT_NAME）。[$TENANT]
   --tenant-domain value                  租户域 - 可选（v3认证）（OS_PROJECT_DOMAIN_NAME）。[$TENANT_DOMAIN]
   --tenant-id value                      租户ID - 可选（v1认证）；否则需要此字段或租户（OS_TENANT_ID）。[$TENANT_ID]
   --user value                           登录用户名（OS_USERNAME）。[$USER]
   --user-id value                        用户ID - 可选 - 大多数Swift系统使用用户，并将此字段留空（v3认证）（OS_USER_ID）。[$USER_ID]

   高级选项

   --chunk-size value      将文件分块到_segments容器中的大小限制。[默认值："5Gi"] [$CHUNK_SIZE]
   --encoding value        后端的编码形式。[默认值："Slash,InvalidUtf8"] [$ENCODING]
   --leave-parts-on-error  如果发生错误，则避免调用中止上传。[默认值：false] [$LEAVE_PARTS_ON_ERROR]
   --no-chunk              在流式传输上传期间不分块文件。[默认值：false] [$NO_CHUNK]
   --no-large-objects      禁用对静态和动态大对象的支持。[默认值：false] [$NO_LARGE_OBJECTS]

   通用选项

   --name value  存储名称（默认值：自动生成的）
   --path value  存储路径

```
{% endcode %}