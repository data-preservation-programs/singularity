# Google Cloud Storage（不是Google Drive）

{% code fullWidth="true" %}
```
名称：
   singularity datasource add gcs - Google Cloud Storage（不是Google Drive）

用法：
   singularity datasource add gcs [命令选项] <数据集名称> <源路径>

描述：
   --gcs-anonymous
      访问公共存储桶和对象而无需凭据。
      
      如果您只想下载文件并且不需要配置凭据，请将其设置为“true”。

   --gcs-auth-url
      认证服务器URL。
      
      如果要使用提供程序的默认值，请将其留空。

   --gcs-bucket-acl
      新存储桶的访问控制列表。

      示例:
         | authenticatedRead | 项目团队所有者获得OWNER权限。
                             | 所有通过身份验证的用户获得READER权限。
         | private           | 项目团队所有者获得OWNER权限。
                             | 默认值，如果留空。
         | projectPrivate    | 项目团队成员根据其角色获得访问权限。
         | publicRead        | 项目团队所有者获得OWNER权限。
                             | 所有用户获得READER权限。
         | publicReadWrite   | 项目团队所有者获得OWNER权限。
                             | 所有用户获得WRITER权限。

   --gcs-bucket-policy-only
      访问检查应使用桶级别的IAM策略。
      
      如果要将对象上传到设置了“仅限桶策略”的存储桶中，
      则需要设置此选项。
      
      设置后，rclone将会：
      
      - 忽略设置在存储桶上的ACL
      - 忽略设置在对象上的ACL
      - 创建设置了“仅限桶策略”的存储桶
      
      文档：https://cloud.google.com/storage/docs/bucket-policy-only
      

   --gcs-client-id
      OAuth客户端ID。
      
      通常留空。

   --gcs-client-secret
      OAuth客户端密钥。
      
      通常留空。

   --gcs-decompress
      如果设置，将解压缩gzip编码的对象。
      
      可以使用“Content-Encoding: gzip”将对象上传到GCS。
      通常情况下，rclone会将这些文件作为压缩对象下载。
      
      如果设置了此标志，则rclone将在接收到这些文件时使用“Content-Encoding: gzip”进行解压缩。
      这意味着rclone无法检查大小和哈希值，但是文件内容将被解压缩。
      

   --gcs-encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --gcs-endpoint
      服务的终端。

      通常留空。

   --gcs-env-auth
      从运行时获取GCP IAM凭据（环境变量或实例元数据，如果没有环境变量）。
      
      仅适用于service_account_file和service_account_credentials为空的情况。

      示例:
         | false | 在下一步中输入凭据。
         | true  | 从环境中获取GCP IAM凭据（环境变量或IAM）。

   --gcs-location
      新创建存储桶的位置。

      示例:
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
         | asia1                   | 双区域：亚洲东北1和亚洲东北2。
         | eur4                    | 双区域：欧洲北1和欧洲西4。
         | nam4                    | 双区域：美国中部1和美国东部1。

   --gcs-no-check-bucket
      如果设置，不尝试检查存储桶是否存在或创建存储桶。
      
      如果已经知道存储桶已经存在，这样可以减少rclone的事务操作次数，
      这样做可能会有用处。
      

   --gcs-object-acl
      新对象的访问控制列表。

      示例:
         | authenticatedRead      | 对象所有者获得OWNER权限。
                                  | 所有通过身份验证的用户获得READER权限。
         | bucketOwnerFullControl | 对象所有者获得OWNER权限。
                                  | 项目团队所有者获得OWNER权限。
         | bucketOwnerRead        | 对象所有者获得OWNER权限。
                                  | 项目团队所有者获得READER权限。
         | private                | 对象所有者获得OWNER权限。
                                  | 默认值，如果留空。
         | projectPrivate         | 对象所有者获得OWNER权限。
                                  | 项目团队成员根据其角色获得访问权限。
         | publicRead             | 对象所有者获得OWNER权限。
                                  | 所有用户获得READER权限。

   --gcs-project-number
      项目编号。
      
      可选项 - 仅用于列举/创建/删除存储桶 - 查看您的开发者控制台。

   --gcs-service-account-credentials
      服务帐号凭据的JSON字符串。
      
      通常留空。
      仅在希望使用服务帐号而不是交互式登录时需要。

   --gcs-service-account-file
      服务帐号凭据的JSON文件路径。
      
      通常留空。
      仅在希望使用服务帐号而不是交互式登录时需要。
      
      文件名中的`~`将被展开，环境变量如`${RCLONE_CONFIG_DIR}`也将被展开。

   --gcs-storage-class
      在Google Cloud Storage中存储对象时使用的存储级别。

      示例:
         | <unset>                      | 默认值
         | MULTI_REGIONAL               | 多区域存储级别
         | REGIONAL                     | 区域存储级别
         | NEARLINE                     | 附近存储级别
         | COLDLINE                     | 冷线存储级别
         | ARCHIVE                      | 存档存储级别
         | DURABLE_REDUCED_AVAILABILITY | 可靠的降低可用性存储级别

   --gcs-token
      OAuth访问令牌的JSON字符串。

   --gcs-token-url
      令牌服务器URL。
      
      如果要使用提供程序的默认值，请将其留空。


选项：
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险操作] 在将数据集导出为CAR文件后删除数据集中的文件。  (默认值：false)
   --rescan-interval value  上次成功扫描后，当此时间间隔过去时，自动重新扫描源目录（默认值：禁用）
   --scanning-state value   设置初始扫描状态（默认值：准备就绪）

   gcs选项

   --gcs-anonymous value             访问公共存储桶和对象而无需凭据。 (默认值："false") [$GCS_ANONYMOUS]
   --gcs-auth-url value              认证服务器URL。 [$GCS_AUTH_URL]
   --gcs-bucket-acl value            新存储桶的访问控制列表。 [$GCS_BUCKET_ACL]
   --gcs-bucket-policy-only value    访问检查应使用桶级别的IAM策略。 (默认值："false") [$GCS_BUCKET_POLICY_ONLY]
   --gcs-client-id value             OAuth客户端ID。 [$GCS_CLIENT_ID]
   --gcs-client-secret value         OAuth客户端密钥。 [$GCS_CLIENT_SECRET]
   --gcs-decompress value            如果设置，将解压缩gzip编码的对象。 (默认值："false") [$GCS_DECOMPRESS]
   --gcs-encoding value              后端的编码方式。 (默认值："Slash,CrLf,InvalidUtf8,Dot") [$GCS_ENCODING]
   --gcs-endpoint value              服务的终端。 [$GCS_ENDPOINT]
   --gcs-env-auth value              从运行时获取GCP IAM凭据（环境变量或实例元数据，如果没有环境变量）。 (默认值："false") [$GCS_ENV_AUTH]
   --gcs-location value              新创建存储桶的位置。 [$GCS_LOCATION]
   --gcs-no-check-bucket value       如果设置，不尝试检查存储桶是否存在或创建存储桶。 (默认值："false") [$GCS_NO_CHECK_BUCKET]
   --gcs-object-acl value            新对象的访问控制列表。 [$GCS_OBJECT_ACL]
   --gcs-project-number value        项目编号。 [$GCS_PROJECT_NUMBER]
   --gcs-service-account-file value  服务帐号凭据的JSON文件路径。 [$GCS_SERVICE_ACCOUNT_FILE]
   --gcs-storage-class value         在Google Cloud Storage中存储对象时使用的存储级别。 [$GCS_STORAGE_CLASS]
   --gcs-token value                 OAuth访问令牌的JSON字符串。 [$GCS_TOKEN]
   --gcs-token-url value             令牌服务器URL。 [$GCS_TOKEN_URL]

```
{% endcode %}