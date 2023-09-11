# Google云存储（不是Google Drive）

{% code fullWidth="true" %}
```
名称：
   singularity storage create gcs - Google云存储（不是Google Drive）

用法：
   singularity storage create gcs [命令选项] [参数...]

描述：
   --client-id
      OAuth Client Id.
      
      通常情况下保持空白即可。

   --client-secret
      OAuth Client Secret.
      
      通常情况下保持空白即可。

   --token
      OAuth Access Token，以JSON格式提供。

   --auth-url
      Auth服务器URL。
      
      保持空白以使用提供程序的默认值。

   --token-url
      Token服务器URL。
      
      保持空白以使用提供程序的默认值。

   --project-number
      项目编号。
      
      可选 - 仅需要在列出/创建/删除存储桶时使用，请参阅您的开发者控制台。

   --service-account-file
      服务账号凭据的JSON文件路径。
      
      通常情况下保持空白。
      仅在希望使用服务账号代替交互式登录时需要。

   --service-account-credentials
      服务账号凭据的JSON blob。
      
      通常情况下保持空白。
      仅在希望使用服务账号代替交互式登录时需要。

   --anonymous
      使用公开的存储桶和对象进行无凭证访问。
      
      如果只想下载文件而不配置凭证，请将其设置为“true”。

   --object-acl
      新对象的访问控制列表。

      示例：
         | authenticatedRead      | 对象所有者获得OWNER权限。
         |                        | 所有经过身份验证的用户获得READER权限。
         | bucketOwnerFullControl | 对象所有者获得OWNER权限。
         |                        | 项目团队所有者获得OWNER权限。
         | bucketOwnerRead        | 对象所有者获得OWNER权限。
         |                        | 项目团队所有者获得READER权限。
         | private                | 对象所有者获得OWNER权限。
         |                        | 如果保留空白，则为默认值。
         | projectPrivate         | 对象所有者获得OWNER权限。
         |                        | 项目团队成员根据其角色获得权限。
         | publicRead             | 对象所有者获得OWNER权限。
         |                        | 所有用户获得READER权限。

   --bucket-acl
      新存储桶的访问控制列表。

      示例：
         | authenticatedRead | 项目团队所有者获得OWNER权限。
         |                   | 所有经过身份验证的用户获得READER权限。
         | private           | 项目团队所有者获得OWNER权限。
         |                   | 如果保留空白，则为默认值。
         | projectPrivate    | 项目团队成员根据其角色获得权限。
         | publicRead        | 项目团队所有者获得OWNER权限。
         |                   | 所有用户获得READER权限。
         | publicReadWrite   | 项目团队所有者获得OWNER权限。
         |                   | 所有用户获得WRITER权限。

   --bucket-policy-only
      访问检查应使用存储桶级别的IAM策略。
      
      如果要向配置了"Bucket Policy Only"的存储桶上传对象，
      则需要设置此选项。
      
      当设置时，rclone:
      
      - 忽略在存储桶上设置的ACL
      - 忽略在对象上设置的ACL
      - 创建带有"Bucket Policy Only"设置的存储桶
      
      文档：https://cloud.google.com/storage/docs/bucket-policy-only
      

   --location
      新创建的存储桶的位置。

      示例：
         | <unset>                 | 默认位置（美国）
         | asia                    | 亚洲的多地域位置
         | eu                      | 欧洲的多地域位置
         | us                      | 美国的多地域位置
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
         | asia1                   | 双重地域：亚洲-东京和亚洲-大阪.
         | eur4                    | 双重地域：欧洲-芬兰和欧洲-荷兰.
         | nam4                    | 双重地域：美国-爱荷华和美国-南卡罗来纳.

   --storage-class
      在Google云存储中存储对象时要使用的存储类别。

      示例：
         | <unset>                      | 默认存储类别
         | MULTI_REGIONAL               | 多地域存储类别
         | REGIONAL                     | 区域存储类别
         | NEARLINE                     | 近线存储类别
         | COLDLINE                     | 冷线存储类别
         | ARCHIVE                      | 归档存储类别
         | DURABLE_REDUCED_AVAILABILITY | 抗击灾难存储类别

   --no-check-bucket
      如果设置，则不尝试检查桶是否存在或创建桶。
      
      当您知道桶已经存在时，这可能对尽量减少rclone事务数量有帮助。

   --decompress
      如果设置，则将解压缩gzip编码的对象。
      
      可以使用"Content-Encoding: gzip"将对象上传到GCS。通常情况下，rclone会将这些文件下载为压缩对象。
      
      如果设置了此标志，则rclone将收到的文件使用"Content-Encoding: gzip"解压缩。这意味着rclone不能检查文件的大小和哈希值，但文件内容将被解压缩。

   --endpoint
      服务的终结点。
      
      通常情况下保持空白。

   --encoding
      后端使用的编码方式。
      
      详见[概述中的编码章节](/overview/#encoding)了解更多信息。

   --env-auth
      从运行时获取GCP IAM凭据（环境变量或实例元数据（如果没有环境变量））。
      
      仅当service_account_file和service_account_credentials为空时有效。

      示例：
         | false | 在下一步输入凭证。
         | true  | 从环境中获取GCP IAM凭据（环境变量或IAM）。

选项：
   --anonymous                          使用公开的存储桶和对象进行无凭证访问。 (默认值: false) [$ANONYMOUS]
   --bucket-acl value                   新存储桶的访问控制列表。 [$BUCKET_ACL]
   --bucket-policy-only                 访问检查应使用存储桶级别的IAM策略。 (默认值: false) [$BUCKET_POLICY_ONLY]
   --client-id value                    OAuth Client Id. [$CLIENT_ID]
   --client-secret value                OAuth Client Secret. [$CLIENT_SECRET]
   --env-auth                           从运行时获取GCP IAM凭据（环境变量或实例元数据（如果没有环境变量））。 (默认值: false) [$ENV_AUTH]
   --help, -h                           显示帮助
   --location value                     新创建的存储桶的位置。 [$LOCATION]
   --object-acl value                   新对象的访问控制列表。 [$OBJECT_ACL]
   --project-number value               项目编号。 [$PROJECT_NUMBER]
   --service-account-credentials value  服务账号凭据的JSON blob。 [$SERVICE_ACCOUNT_CREDENTIALS]
   --service-account-file value         服务账号凭据的JSON文件路径。 [$SERVICE_ACCOUNT_FILE]
   --storage-class value                在Google云存储中存储对象时要使用的存储类别。 [$STORAGE_CLASS]

   高级选项

   --auth-url value   Auth服务器URL。 [$AUTH_URL]
   --decompress       如果设置，则将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --encoding value   后端使用的编码方式。 (默认值: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value   服务的终结点。 [$ENDPOINT]
   --no-check-bucket  如果设置，则不尝试检查桶是否存在或创建桶。 (默认值: false) [$NO_CHECK_BUCKET]
   --token value      OAuth Access Token，以JSON格式提供。 [$TOKEN]
   --token-url value  Token服务器URL。 [$TOKEN_URL]

   通用选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}