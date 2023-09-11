# OpenStack Swift（Rackspace Cloud Files、Memset Memstore、OVH）

{% code fullWidth="true" %}
```
名称:
   singularity storage update swift - OpenStack Swift（Rackspace Cloud Files、Memset Memstore、OVH）

用法:
   singularity storage update swift [命令选项] <名称|ID>

说明:
   --env-auth
      从标准的OpenStack环境变量中获取Swift凭据。

      示例:
         | false | 在下一步输入Swift凭据。
         | true  | 从环境变量获取Swift凭据。
         |       | 如果使用此选项，请留空其他字段。

   --user
      登录用户名（OS_USERNAME）。

   --key
      API密钥或密码（OS_PASSWORD）。

   --auth
      服务器的身份验证URL（OS_AUTH_URL）。

      示例:
         | https://auth.api.rackspacecloud.com/v1.0     | Rackspace US
         | https://lon.auth.api.rackspacecloud.com/v1.0 | Rackspace UK
         | https://identity.api.rackspacecloud.com/v2.0 | Rackspace v2
         | https://auth.storage.memset.com/v1.0         | Memset Memstore UK
         | https://auth.storage.memset.com/v2.0         | Memset Memstore UK v2
         | https://auth.cloud.ovh.net/v3                | OVH

   --user-id
      用户ID（v3身份验证可选）- 大多数Swift系统使用用户，并将此字段留空（OS_USER_ID）。

   --domain
      用户域（v3身份验证可选）（OS_USER_DOMAIN_NAME）。

   --tenant
      租户名称（v1身份验证可选，否则需要此选项或租户ID）（OS_TENANT_NAME或OS_PROJECT_NAME）。

   --tenant-id
      租户ID（v1身份验证可选，否则需要此选项或租户）（OS_TENANT_ID）。

   --tenant-domain
      租户域（v3身份验证可选）（OS_PROJECT_DOMAIN_NAME）。

   --region
      区域名称（OS_REGION_NAME）。

   --storage-url
      存储URL（OS_STORAGE_URL）。

   --auth-token
      替代身份验证的身份验证令牌（可选）（OS_AUTH_TOKEN）。

   --application-credential-id
      应用凭据ID（OS_APPLICATION_CREDENTIAL_ID）。

   --application-credential-name
      应用凭据名称（OS_APPLICATION_CREDENTIAL_NAME）。

   --application-credential-secret
      应用凭据密码（OS_APPLICATION_CREDENTIAL_SECRET）。

   --auth-version
      身份验证版本（可选）- 如果身份验证URL没有版本，则设置为（1,2,3）（ST_AUTH_VERSION）。

   --endpoint-type
      选择服务目录中的端点类型（OS_ENDPOINT_TYPE）。

      示例:
         | public   | 公共（默认值，如果不确定，请选择此项）
         | internal | 内部（使用内部服务网络）
         | admin    | 管理员

   --leave-parts-on-error
      如果为true，则在失败时避免调用中止上传。
      
      对于不同会话之间的恢复上传操作，应设置为true。

   --storage-policy
      创建新容器时要使用的存储策略。
      
      创建新容器时应用指定的存储策略。此策略之后无法更改。允许的配置值及其含义取决于您的Swift存储供应商。

      示例:
         | <未设置> | 默认值
         | pcs     | OVH公共云存储
         | pca     | OVH公共云归档

   --chunk-size
      文件大小超过此值时，文件将被分片成_segemnts容器。
      
      文件大小超过此值时，文件将被分片成_segemnts容器。默认值为5 GiB，这是其最大值。

   --no-chunk
      在流式上传期间不要对文件进行分片。
      
      在进行流式上传（例如使用rcat或mount）时，设置此标志将导致Swift后端不上传分片文件。
      
      这将限制最大上传大小为5 GiB。但是，非分片文件更易处理且具有MD5SUM哈希值。
      
      在执行常规复制操作时，rclone仍会将大于chunk_size的文件分片。

   --no-large-objects
      禁用静态和动态大对象的支持。
      
      Swift无法透明地存储大于5 GiB的文件。有两种方式可以实现，即静态大对象或动态大对象，但API不允许rclone在不进行对象的HEAD请求的情况下确定文件是静态还是动态大对象。由于这些对象需要以不同的方式处理，这意味着rclone必须对对象发出HEAD请求，例如在读取校验和时。
      
      当设置`no_large_objects`时，rclone将假定没有存储静态或动态大对象。这意味着它可以停止执行额外的HEAD调用，这反过来极大地增加了性能，尤其是在设置了`--checksum`的情况下执行的Swift到Swift传输。
      
      设置此选项意味着`no_chunk`，并且不会上传以分片方式创建的文件，因此大于5 GiB的文件在上传时将失败。
      
      如果设置此选项，而实际上存在静态或动态大对象，则这些对象的哈希值将不正确。下载会成功，但是其他操作（例如删除和复制）将失败。
      

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

选项:
   --application-credential-id value      应用凭据ID（OS_APPLICATION_CREDENTIAL_ID）。[$APPLICATION_CREDENTIAL_ID]
   --application-credential-name value    应用凭据名称（OS_APPLICATION_CREDENTIAL_NAME）。[$APPLICATION_CREDENTIAL_NAME]
   --application-credential-secret value  应用凭据密码（OS_APPLICATION_CREDENTIAL_SECRET）。[$APPLICATION_CREDENTIAL_SECRET]
   --auth value                           服务器的身份验证URL（OS_AUTH_URL）。[$AUTH]
   --auth-token value                     替代身份验证的身份验证令牌（OS_AUTH_TOKEN）（可选）。[$AUTH_TOKEN]
   --auth-version value                   身份验证版本（可选）- 如果身份验证URL没有版本，则设置为（1,2,3）（ST_AUTH_VERSION）。 (默认值: 0) [$AUTH_VERSION]
   --domain value                         用户域（v3身份验证可选）（OS_USER_DOMAIN_NAME）[$DOMAIN]
   --endpoint-type value                  选择服务目录中的端点类型（OS_ENDPOINT_TYPE）。 (默认值: "public") [$ENDPOINT_TYPE]
   --env-auth                             从标准的OpenStack环境变量中获取Swift凭据（默认值: false）[$ENV_AUTH]
   --help, -h                             显示帮助信息
   --key value                            API密钥或密码（OS_PASSWORD）。[$KEY]
   --region value                         区域名称（OS_REGION_NAME）。[$REGION]
   --storage-policy value                 创建新容器时要使用的存储策略。[$STORAGE_POLICY]
   --storage-url value                    存储URL（OS_STORAGE_URL）。[$STORAGE_URL]
   --tenant value                         租户名称（v1身份验证可选，否则需要此选项或租户ID）（OS_TENANT_NAME或OS_PROJECT_NAME）。[$TENANT]
   --tenant-domain value                  租户域（v3身份验证可选）（OS_PROJECT_DOMAIN_NAME）。[$TENANT_DOMAIN]
   --tenant-id value                      租户ID（v1身份验证可选，否则需要此选项或租户）（OS_TENANT_ID）。[$TENANT_ID]
   --user value                           登录用户名（OS_USERNAME）。[$USER]
   --user-id value                        用户ID（v3身份验证可选）- 大多数Swift系统使用用户，并将此字段留空（OS_USER_ID）。[$USER_ID]

   高级选项

   --chunk-size value      文件大小超过此值时，文件将被分片成_segemnts容器。 (默认值: "5Gi") [$CHUNK_SIZE]
   --encoding value        后端的编码方式。 (默认值: "Slash,InvalidUtf8") [$ENCODING]
   --leave-parts-on-error  如果为true，则在失败时避免调用中止上传。 (默认值: false) [$LEAVE_PARTS_ON_ERROR]
   --no-chunk              在流式上传期间不要对文件进行分片。 (默认值: false) [$NO_CHUNK]
   --no-large-objects      禁用静态和动态大对象的支持 (默认值: false) [$NO_LARGE_OBJECTS]

```
{% endcode %}