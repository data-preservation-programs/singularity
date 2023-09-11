# 使用 OCI 用户和 API 密钥进行身份验证。
您需要在配置文件中输入你的租户 OCID、用户 OCID、区域、路径和 API 密钥的指纹。
[https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm)

## 用法
```
singularity storage update oos user_principal_auth [命令选项] <名称|ID>
```

## 描述
- `--namespace`：对象存储命名空间
- `--compartment`：对象存储区域 OCID
- `--region`：对象存储区域
- `--endpoint`：对象存储 API 的终端

  如果留空，则使用该区域的默认终端。

- `--config-file`：OCI 配置文件的路径

  示例：
  - `~/.oci/config`：OCI 配置文件的位置

- `--config-profile`：OCI 配置文件中的配置文件名

  示例：
  - `Default`：使用默认配置文件

- `--storage-tier`：存储新对象时要使用的存储类别

  参考：[https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm](https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm)

  示例：
  - `Standard`：标准存储类别，这是默认的存储类别
  - `InfrequentAccess`：低频访问存储类别
  - `Archive`：存档存储类别

- `--upload-cutoff`：切换到分块上传的截止大小

  大于此大小的文件将分块上传。最小值为 0，最大值为 5 GiB。

- `--chunk-size`：上传时使用的块大小

  对于大于 `upload_cutoff` 的文件，或者大小未知的文件（例如使用 `rclone rcat` 上传或使用 `rclone mount` 上传的文件，或者谷歌照片或谷歌文档），将使用此块大小进行分块上传。

  请注意，每个传输会在内存中缓冲 `upload_concurrency` 个此大小的块。如果您正在通过高速链路传输大文件并且拥有足够的内存，增加此值将加快传输速度。

  当上传已知大小的大文件以保持低于 10,000 个块的限制时，rclone 会自动增加块大小。

  大小未知的文件使用配置的块大小进行上传。默认块大小为 5 MiB，最多可有 10,000 个块，这意味着默认情况下您可以流式上传的文件的最大大小为 48 GiB。如果您要流式上传更大的文件，则需要增加块大小。

  增加块大小会降低使用 "-P" 标志显示的进度统计的准确性。

- `--upload-concurrency`：分块上传的并发数

  同时上传同一文件的块数量。

  如果您通过高速链路上传少量大文件，并且这些上传没有充分利用您的带宽，则增加此值可能有助于加快传输速度。

- `--copy-cutoff`：切换到分块复制的截止大小

  需要服务器端复制的大于此大小的文件将分块复制。

  最小值为 0，最大值为 5 GiB。

- `--copy-timeout`：复制超时时间

  复制是异步操作，设置超时时间以等待复制成功。

- `--disable-checksum`：不在对象元数据中存储 MD5 校验和

  通常，在上传之前，rclone 会计算输入的 MD5 校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件的开始上传可能会导致长时间的延迟。

- `--encoding`：后端的编码方式

  有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

- `--leave-parts-on-error`：在失败时避免调用中止上传，以便将所有成功上传的分块留在 S3 中以供手动恢复。

  对于在不同会话之间恢复上传，应将其设置为 true。

  警告：在对象存储上存储不完整的分块上传的部分将计入空间使用量，并且如果不进行清理，则会增加额外成本。

- `--no-check-bucket`：如果设置了此选项，则不会尝试检查 bucket 是否存在或创建 bucket。

  如果您知道 bucket 已经存在，那么这可能有助于最小化 rclone 的事务次数。

  如果使用的用户没有创建桶的权限，这也是必需的。

- `--sse-customer-key-file`：使用 SSE-C 时，将包含与对象关联的 AES-256 加密密钥的 Base64 编码字符串的文件

  示例：
  - `<unset>`：无

- `--sse-customer-key`：使用 SSE-C 时，可选头，指定用于加密或解密数据的 Base64 编码的 256 位加密密钥

  示例：
  - `<unset>`：无

- `--sse-customer-key-sha256`：如果使用 SSE-C，可选头，指定加密密钥的 Base64 编码 SHA256 哈希值

  示例：
  - `<unset>`：无

- `--sse-kms-key-id`：如果在库中使用您自己的主密钥，此头指定用于调用密钥管理服务来生成数据加密密钥或加密或解密数据加密密钥的主加密密钥的 OCID

  示例：
  - `<unset>`：无

- `--sse-customer-algorithm`：如果使用 SSE-C，可选头，指定 "AES256" 作为加密算法

  对象存储支持 "AES256" 作为加密算法。有关详细信息，请参阅[使用您自己的密钥进行服务器端加密](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

  示例：
  - `<unset>`：无
  - `AES256`：AES256

## 选项
- `--compartment`：对象存储区域 OCID
- `--config-file`：OCI 配置文件的路径（默认值：`~/.oci/config`）
- `--config-profile`：OCI 配置文件中的配置文件名（默认值："Default"）
- `--endpoint`：对象存储 API 的终端
- `--help, -h`：显示帮助
- `--namespace`：对象存储命名空间
- `--region`：对象存储区域

高级选项：
- `--chunk-size`：用于上传的块大小（默认值："5Mi"）
- `--copy-cutoff`：切换到分块复制的截止大小（默认值："4.656Gi"）
- `--copy-timeout`：复制超时时间（默认值："1m0s"）
- `--disable-checksum`：不在对象元数据中存储 MD5 校验和（默认值：false）
- `--encoding`：后端的编码方式（默认值："Slash,InvalidUtf8,Dot"）
- `--leave-parts-on-error`：在失败时避免调用中止上传，以便将所有成功上传的分块留在 S3 中以供手动恢复（默认值：false）
- `--no-check-bucket`：如果设置了此选项，则不会尝试检查 bucket 是否存在或创建 bucket（默认值：false）
- `--sse-customer-algorithm`：如果使用 SSE-C，可选头，指定 "AES256" 作为加密算法
- `--sse-customer-key`：要使用 SSE-C，可选头，指定用于加密或解密数据的 Base64 编码的 256 位加密密钥
- `--sse-customer-key-file`：要使用 SSE-C，将包含与对象关联的 AES-256 加密密钥的 Base64 编码字符串的文件
- `--sse-customer-key-sha256`：如果使用 SSE-C，可选头，指定加密密钥的 Base64 编码 SHA256 哈希值
- `--sse-kms-key-id`：如果在库中使用您自己的主密钥，此头指定用于调用密钥管理服务来生成数据加密密钥或加密或解密数据加密密钥的主加密密钥的 OCID
- `--storage-tier`：存储新对象时要使用的存储类别
- `--upload-concurrency`：分块上传的并发数（默认值：10）
- `--upload-cutoff`：切换到分块上传的截止大小（默认值："200Mi"）