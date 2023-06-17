# Google Cloud 存储 (非 Google Drive)

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add gcs - Google Cloud 存储 (非 Google Drive)

USAGE:
   singularity datasource add gcs [command options] <dataset_name> <source_path>

DESCRIPTION:
   --gcs-endpoint
      服务的终端节点。
      
      通常情况下保持为空白。

   --gcs-encoding
      后端的编码方式。
      
      更多信息请参见[概述页面中的编码部分](/overview/#encoding)。

   --gcs-client-secret
      OAuth 客户端密钥。
      
      通常情况下保持为空白。

   --gcs-object-acl
      新对象的访问控制列表。

      例如:
         | authenticatedRead      | 对象所有者获得 OWNER 访问权限。
                                  | 所有已认证用户获得 READER 访问权限。
         | bucketOwnerFullControl | 对象所有者获得 OWNER 访问权限。
                                  | 项目团队所有者获得 OWNER 访问权限。
         | bucketOwnerRead        | 对象所有者获得 OWNER 访问权限。
                                  | 项目团队所有者获得 READER 访问权限。
         | private                | 对象所有者获得 OWNER 访问权限。
                                  | 如果未设置，则默认为此项。
         | projectPrivate         | 对象所有者获得 OWNER 访问权限。
                                  | 项目团队成员根据其角色获得访问权限。
         | publicRead             | 对象所有者获得 OWNER 访问权限。
                                  | 所有用户获得 READER 访问权限。

   --gcs-bucket-policy-only
      访问检查应使用存储桶级别的 IAM 策略。
      
      如果要向启用了存储桶策略而添加对象，则需要设置此参数。
      
      当它被设置时，rclone：
      
      - 忽略已设置在存储桶上的 ACL。
      - 忽略已设置在对象上的 ACL。
      - 创建启用了存储桶策略的存储桶。
      
      文档：https://cloud.google.com/storage/docs/bucket-policy-only
      

   --gcs-storage-class
      存储对象时使用的存储级别。

      例如:
         | <unset>                      | 默认值
         | MULTI_REGIONAL               | 多地域存储级别
         | REGIONAL                     | 区域性存储级别
         | NEARLINE                     | 近线性存储级别
         | COLDLINE                     | 冷线性存储级别
         | ARCHIVE                      | 存档存储级别
         | DURABLE_REDUCED_AVAILABILITY | 可靠性降低的存储级别

   --gcs-service-account-file
      服务账号凭据 JSON 文件路径。
      
      通常情况下保持为空白。
      只有当您想要使用 SA（而不是交互式登录）时才需要此参数。
      
      文件名中的'~'将被展开，其他环境变量(如`${RCLONE_CONFIG_DIR}`)也会以相应的值被替换。

   --gcs-location
      新创建的存储桶的位置。

      例如:
         | <unset>                 | 默认位置(美国)
         | asia                    | 亚洲多地域位置
         | eu                      | 欧洲多地域位置
         | us                      | 美国多地域位置
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
         | us-central1             | 爱荷华州
         | us-east1                | 南卡罗来纳州
         | us-east4                | 弗吉尼亚北部
         | us-west1                | 俄勒冈州
         | us-west2                | 加利福尼亚州
         | us-west3                | 盐湖城
         | us-west4                | 拉斯维加斯
         | northamerica-northeast1 | 蒙特利尔
         | northamerica-northeast2 | 多伦多
         | southamerica-east1      | 圣保罗
         | southamerica-west1      | 圣地亚哥
         | asia1                   | 双区域：亚洲东北和亚洲东北2。
         | eur4                    | 双区域：欧洲北部1和欧洲西部4。
         | nam4                    | 双区域：美国中部1和美国东部1。
         

   --gcs-env-auth
      从运行时获取 GCP IAM 凭据(从环境变量或实例元数据(如果没有设置环境变量)中获取)。
      
      仅在 service_account_file 和 service_account_credentials 为空时才应用。

      例如:
         | false | 在下一步中输入凭据。
         | true  | 从环境(环境变量或 IAM)中获取 GCP IAM 凭据。

   --gcs-token
      作为 JSON 块的 OAuth 访问令牌。

   --gcs-auth-url
      认证服务器 URL。
      
      保持为空白以使用提供程序的默认值。

   --gcs-token-url
      令牌服务器 URL。
      
      保持为空白以使用提供程序的默认值。

   --gcs-project-number
      项目编号。
      
      可选 - 仅在列表、创建或删除存储桶时需要 - 参见您的开发者控制台。

   --gcs-client-id
      OAuth 客户端 ID。
      
      通常情况下保持为空白。

   --gcs-service-account-credentials
      服务账号凭据 JSON 块。
      
      通常情况下保持为空白。
      只有当您想要使用 SA（而不是交互式登录）时才需要此参数。

   --gcs-no-check-bucket
      如果设置，不尝试检查存储桶是否存在或创建它。
      
      当您知道存储桶已经存在时，这可能有用，以尝试最小化 rclone 所执行的事务数量。

   --gcs-decompress
      如果设置，则解压缩 gzip 编码的对象。
      
      可以上传具有"Content-Encoding: gzip"的对象到 GCS。通常情况下，rclone 将下载这些文件作为已压缩的对象。
      
      如果设置了此标志，rclone 将使用接