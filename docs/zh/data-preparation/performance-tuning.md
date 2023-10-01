# Singularity 中的性能调优

Singularity 提供了一系列的配置选项，允许用户优化数据准备的性能。本指南将阐述这些配置选项，并提供有效调优的指导说明。

## 内行准备
* **描述**：内行准备消除了需要额外磁盘空间来存储 CAR 文件的需求。然而，它会增加一定的数据库查询和存储开销。
* **影响**：这种开销通常是可以忽略不计的，但对于包含许多小文件的数据集来说，可能会变得显著。
* **配置**：要禁用内行准备，请在 `singularity prep create` 命令中使用 `--no-inline`。
* **进一步阅读**：[内行准备](../topics/inline-preparation.md)

## DAG 更新
* **描述**：在准备过程中，Singularity 会为每个目录刷新 DAG 和 CID，以便实时跟踪变化。
* **影响**：这会引入轻微的数据库开销，因为每次准备 CAR 文件时都会更新目录。
* **配置**：要禁用 DAG 更新，请在 `singularity prep create` 命令中使用 `--no-dag`。
  
## 数据准备中的并行处理

### 扫描
* **描述**：扫描涉及遍历源存储以整理文件列表。对于本地存储，速度较快，但对于像 S3 这样的远程存储来说可能较慢。
* **配置**：
  * **启用并行处理**：在 `singularity storage create` 或 `singularity storage update` 命令中使用 `--client-scan-concurrency` 参数。
  * **注意**：启用并行处理可能导致文件按非确定性方式处理。

### 打包
* **描述**：打包将多个文件合并为一个 CAR 文件，这是一个既消耗 CPU 又消耗 IO 的操作。对于网络受限的远程存储来说，增加并行处理是有益的。
* **配置**：
    * **调整并行处理**：在 `singularity run dataset-worker` 命令中使用 `--concurrency` 参数。

## 使用服务器的最后修改时间
* **描述**：某些远程存储（如 `AWS S3`）提供自定义的 `mtime` 和服务器端的最后修改时间。默认情况下，Singularity 会检查是否存在自定义的 `mtime`，如果有则使用它，否则使用服务器的最后修改时间。
* **影响**：跳过检查自定义的 `mtime` 并直接使用服务器的最后修改时间可以减少对远程存储的请求数量。
* **配置**：要优先使用服务器的时间并避免获取对象元数据，请在 `singularity storage create` 或 `singularity storage update` 命令中使用 `--client-use-server-mod-time`。

## 重试策略
### 网络请求重试
* **描述**：对于失败的远程文件夹列表或文件打开操作，Singularity 利用 RClone 的重试机制。
* **配置**：要增加重试次数，请在 `singularity storage create` 或 `singularity storage update` 命令中使用 `--client-low-level-retries` 参数。

### 网络 IO 重试
* **描述**：尽管网络请求成功，但由于不稳定的网络连接，网络 IO 可能失败。Singularity 支持重试和从上次成功点继续操作。
* **配置**：请在 `singularity storage create` 或 `singularity storage update` 命令中使用以下标志。
```shell
 --client-retry-backoff value      # 重试 IO 读取错误时的延迟后退（默认值：1s）
 --client-retry-backoff-exp value  # 在重试 IO 读取错误时指数级的延迟后退（默认值：1.0）
 --client-retry-delay value        # 在重试 IO 读取错误之前的初始延迟（默认值：1s）
 --client-retry-max value          # IO 读取错误的最大重试次数（默认值：10）
```

## 跳过无法访问的文件
* **描述**：权限可能会阻止从远程存储访问某些文件。只有在尝试打开文件时，这些问题才会显现，导致打包作业失败。
* **配置**：要跳过无法访问的文件，请在 `singularity storage create` 或 `singularity storage update` 命令中使用 `--client-skip-inaccessible-files`。