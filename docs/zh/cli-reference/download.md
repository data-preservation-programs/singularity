# 从元数据 API 下载 CAR 文件

{% code fullWidth="true" %}
```
NAME:
   singularity download - 从元数据 API 下载 CAR 文件

用法:
   singularity download [command options] <piece_cid>

类别:
   实用工具

选项:
   1Fichier

   --fichier-api-key value          您的 API 密钥，从 https://1fichier.com/console/params.pl 获取。 [$FICHIER_API_KEY]
   --fichier-encoding value         后端的编码。 (默认值: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$FICHIER_ENCODING]
   --fichier-file-password value    如果要下载受密码保护的共享文件，请添加此参数。 [$FICHIER_FILE_PASSWORD]
   --fichier-folder-password value  如果要列出受密码保护的共享文件夹中的文件，请添加此参数。 [$FICHIER_FOLDER_PASSWORD]
   --fichier-shared-folder value    如果要下载共享文件夹，请添加此参数。 [$FICHIER_SHARED_FOLDER]

   Akamai NetStorage

   --netstorage-account value   设置 NetStorage 帐户名 [$NETSTORAGE_ACCOUNT]
   --netstorage-host value      要连接到的 NetStorage 主机的域名+路径。 [$NETSTORAGE_HOST]
   --netstorage-protocol value  选择 HTTP 或 HTTPS 协议。 (默认值: "https") [$NETSTORAGE_PROTOCOL]
   --netstorage-secret value    设置 NetStorage 帐户密码/G2O 密钥用于身份验证。 [$NETSTORAGE_SECRET]

   Amazon Drive

   --acd-auth-url value            Auth 服务器 URL。 [$ACD_AUTH_URL]
   --acd-checkpoint value          用于内部轮询的检查点 (用于调试)。 [$ACD_CHECKPOINT]
   --acd-client-id value           OAuth 客户端 ID。 [$ACD_CLIENT_ID]
   --acd-client-secret value       OAuth 客户端密码。 [$ACD_CLIENT_SECRET]
   --acd-encoding value            后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ACD_ENCODING]
   --acd-templink-threshold value  大于等于此大小的文件将通过临时链接下载。 (默认值: "9Gi") [$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               OAuth 访问令牌作为 JSON 字符串。 [$ACD_TOKEN]
   --acd-token-url value           令牌服务器的 URL。 [$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  完整上传失败后，在等待 GiB 时的附加时间。 (默认值: "3m0s") [$ACD_UPLOAD_WAIT_PER_GB]

   符合 Amazon S3 储存提供商，包括 AWS、阿里巴巴、Ceph、中国移动、Cloudflare、ArvanCloud、DigitalOcean、Dreamhost、华为 OBS、IBM COS、IDrive e2、IONOS Cloud、Liara、Lyve Cloud、Minio、Netease、RackCorp、Scaleway、SeaweedFS、StackPath、Storj、腾讯云 COS、七牛云和 Wasabi

   --s3-access-key-id value            AWS 访问密钥 ID。 [$S3_ACCESS_KEY_ID]
   --s3-acl value                      创建存储桶和存储或复制对象时使用的访问控制列表。 [$S3_ACL]
   --s3-bucket-acl value               创建存储桶时使用的访问控制列表。 [$S3_BUCKET_ACL]
   --s3-chunk-size value               用于上传的分块大小。 (默认值: "5Mi") [$S3_CHUNK_SIZE]
   --s3-copy-cutoff value              切换到分块复制的阈值。 (默认值: "4.656Gi") [$S3_COPY_CUTOFF]
   --s3-decompress                     如果设置则解压缩使用 gzip 编码的对象。 (默认值: false) [$S3_DECOMPRESS]
   --s3-disable-checksum               不要将 MD5 校验和与对象元数据一起存储。 (默认值: false) [$S3_DISABLE_CHECKSUM]
   --s3-disable-http2                  禁用 S3 后端的 http2。 (默认值: false) [$S3_DISABLE_HTTP2]
   --s3-download-url value             自定义下载的终端节点。 [$S3_DOWNLOAD_URL]
   --s3-encoding value                 后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$S3_ENCODING]
   --s3-endpoint value                 S3 API 的终端节点。 [$S3_ENDPOINT]
   --s3-env-auth                       从运行时获取 AWS 凭据 (环境变量或 EC2/ECS 元数据，如果没有 env vars)。 (默认值: false) [$S3_ENV_AUTH]
   --s3-force-path-style               如果为 true，则使用路径样式访问，否则使用虚拟主机样式访问。 (默认