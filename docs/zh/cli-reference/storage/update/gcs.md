# Google Cloud Storage（不是Google Drive）

{% code fullWidth="true" %}
```
名称：
   singularity storage update gcs - Google Cloud Storage（不是Google Drive）

用法：
   singularity storage update gcs [命令选项] <名称|id>

描述：
   --client-id
      OAuth客户端ID。
      
      通常留空。

   --client-secret
      OAuth客户端密钥。
      
      通常留空。

   --token
      OAuth访问令牌作为JSON数据块。

   --auth-url
      授权服务器URL。
      
      留空以使用提供商的默认值。

   --token-url
      令牌服务器URL。
      
      留空以使用提供商的默认值。

   --project-number
      项目编号。
      
      可选 - 仅在列出/创建/删除存储区时需要 - 参见您的开发者控制台。

   --service-account-file
      服务帐号凭据JSON文件路径。
      
      通常留空。
      仅在希望使用SA而非交互式登录时需要。
      
      文件名中的 `~` 将扩展为完整路径，环境变量例如 `${RCLONE_CONFIG_DIR}` 也会被扩展。

   --service-account-credentials
      服务帐号凭据JSON数据块。
      
      通常留空。
      仅在希望使用SA而非交互式登录时需要。

   --anonymous
      无凭据访问公有存储区和对象。
      
      如果只想下载文件且不配置凭据，则设置为 'true'。

   --object-acl
      新对象的访问控制列表。

      示例：
         | authenticatedRead      | 对象所有者获得OWNER访问权限。
         |                        | 所有认证用户获得READER访问权限。
         | bucketOwnerFullControl | 对象所有者获得OWNER访问权限。
         |                        | 项目团队所有者获得OWNER访问权限。
         | bucketOwnerRead        | 对象所有者获得OWNER访问权限。
         |                        | 项目团队所有者获得READER访问权限。
         | private                | 对象所有者获得OWNER访问权限。
         |                        | 如果留空，则默认为此选项。
         | projectPrivate         | 对象所有者获得OWNER访问权限。
         |                        | 与成员角色相对应的项目团队成员获得访问权限。
         | publicRead             | 对象所有者获得OWNER访问权限。
         |                        | 所有用户获得READER访问权限。

   --bucket-acl
      新存储区的访问控制列表。

      示例：
         | authenticatedRead | 项目团队所有者获得OWNER访问权限。
         |                   | 所有认证用户获得READER访问权限。
         | private           | 项目团队所有者获得OWNER访问权限。
         |                   | 如果留空，则默认为此选项。
         | projectPrivate    | 与成员角色相对应的项目团队成员获得访问权限。
         | publicRead        | 项目团队所有者获得OWNER访问权限。
         |                   | 所有用户获得READER访问权限。
         | publicReadWrite   | 项目团队所有者获得OWNER访问权限。
         |                   | 所有用户获得WRITER访问权限。

   --bucket-policy-only
      存储区级别IAM策略应用访问检查。
      
      如果要上传对象到设置了“仅限存储区策略”的存储区，则需要设置此选项。
      
      当设置此选项时，rclone：
      
      - 忽略设置在存储区上的ACL
      - 忽略设置在对象上的ACL
      - 创建设置了“仅限存储区策略”的存储区
      
      文档：https://cloud.google.com/storage/docs/bucket-policy-only
      

   --location
      新创建的存储区的位置。

      示例：
         | <unset>                 | 默认位置（美国）
         | asia                    | 亚洲的多区域位置
         | eu                      | 欧洲的多区域位置
         | us                      | 美国的多区域位置
         | asia-east1              | 台湾
         | asia-east2              | 香港
         | asia-northeast1         | 东京
         | asia-northeast2         | 大阪
         | asia-northeast3         | 首尔
         | asia-south1             | 孟买
         | asia-south2             | 德里
         | asia-southeast1         | 新加坡
         | asia-southeast2         | 雅加达
         | australia-southeast1    | 悉尼
         | australia-southeast2    | 墨尔本
         | europe-north1           | 芬兰
         | europe-west1            | 比利时
         | europe-west2            | 伦敦
         | europe-west3            | 法兰克福
         | europe-west4            | 荷兰
         | europe-west6            | 苏黎世
         | europe-central2         | 华沙
         | us-central1             | 爱荷华
         | us-east1                | 南卡罗来纳
         | us-east4                | 弗吉尼亚北部
         | us-west1                | 俄勒冈
         | us-west2                | 加利福尼亚
         | us-west3                | 盐湖城
         | us-west4                | 拉斯维加斯
         | northamerica-northeast1 | 蒙特利尔
         | northamerica-northeast2 | 多伦多
         | southamerica-east1      | 圣保罗
         | southamerica-west1      | 圣地亚哥
         | asia1                   | 双区域：亚洲东北1和亚洲东北2
         | eur4                    | 双区域：欧洲北部1和欧洲西部4
         | nam4                    | 双区域：美国中部1和美国东部1

   --storage-class
      在Google Cloud Storage中存储对象时使用的存储类别。

      示例：
         | <unset>                      | 默认
         | MULTI_REGIONAL               | 多区域存储类别
         | REGIONAL                     | 区域存储类别
         | NEARLINE                     | 近线存储类别
         | COLDLINE                     | 冷线存储类别
         | ARCHIVE                      | 存档存储类别
         | DURABLE_REDUCED_AVAILABILITY | 持久化降低可用性存储类别

   --no-check-bucket
      如果设置了，则不尝试检查存储区是否存在或创建存储区。
      
      如果已知存储区已存在，可以使用此选项来尝试最小化rclone的事务数量。

   --decompress
      如果设置了，则会解压缩gzip编码的对象。
      
      可以使用设置了 “Content-Encoding: gzip” 的对象上传到GCS。通常，rclone会将这些文件下载为压缩对象。
      
      如果设置了此标记，则rclone将在接收到这些文件时解压缩带有 “Content-Encoding: gzip” 的文件。这意味着rclone不能检查大小和哈希，但文件内容将被解压缩。

   --endpoint
      服务的终端节点。
      
      通常留空。

   --encoding
      后端的编码。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --env-auth
      从运行时获取GCP IAM凭据（环境变量或实例元数据，如果没有环境变量）。
      
      仅当service_account_file和service_account_credentials为空时适用。

      示例：
         | false | 在下一步中输入凭据。
         | true  | 从环境（环境变量或IAM）获取GCP IAM凭证。


选项：
   --anonymous                          无凭据访问公有存储区和对象。 (default: false) [$ANONYMOUS]
   --bucket-acl value                   新存储区的访问控制列表。 [$BUCKET_ACL]
   --bucket-policy-only                 存储区级别IAM策略应用访问检查。 (default: false) [$BUCKET_POLICY_ONLY]
   --client-id value                    OAuth客户端ID。 [$CLIENT_ID]
   --client-secret value                OAuth客户端密钥。 [$CLIENT_SECRET]
   --env-auth                           从运行时获取GCP IAM凭据（环境变量或实例元数据，如果没有环境变量）。 (default: false) [$ENV_AUTH]
   --help, -h                           显示帮助信息
   --location value                     新创建的存储区的位置。 [$LOCATION]
   --object-acl value                   新对象的访问控制列表。 [$OBJECT_ACL]
   --project-number value               项目编号。 [$PROJECT_NUMBER]
   --service-account-credentials value  服务帐号凭据JSON数据块。 [$SERVICE_ACCOUNT_CREDENTIALS]
   --service-account-file value         服务帐号凭据JSON文件路径。 [$SERVICE_ACCOUNT_FILE]
   --storage-class value                在Google Cloud Storage中存储对象时使用的存储类别。 [$STORAGE_CLASS]

   高级选项

   --auth-url value   授权服务器URL。 [$AUTH_URL]
   --decompress       如果设置了，则会解压缩gzip编码的对象。 (default: false) [$DECOMPRESS]
   --encoding value   后端的编码。 (default: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value   服务的终端节点。 [$ENDPOINT]
   --no-check-bucket  如果设置了，则不尝试检查存储区是否存在或创建存储区。 (default: false) [$NO_CHECK_BUCKET]
   --token value      OAuth访问令牌作为JSON数据块。 [$TOKEN]
   --token-url value  令牌服务器URL。 [$TOKEN_URL]

```
{% endcode %}