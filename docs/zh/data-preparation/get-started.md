# 使用 Singularity 入门

按照以下步骤设置和开始使用 Singularity。

## 1. 初始化数据库

如果您是第一次使用 Singularity，您需要初始化数据库。这个步骤只需要执行一次。

```sh
singularity admin init
```

## 2. 连接存储系统
Singularity 与 RClone 合作，为超过 40 个不同的存储系统提供无缝集成。这些存储系统可以扮演两个主要角色：
* **源存储**: 这是数据集当前存储的位置，也是 Singularity 用来准备数据的地方。
* **输出存储**: 这是 Singularity 处理后将 CAR（内容寻址存档）文件存储的目的地。
选择一个适合您需求的存储系统，并将其与 Singularity 连接以开始准备数据集。

### 2a. 添加本地文件系统

最常用的存储系统是本地文件系统。要将文件夹添加为 Singularity 的源存储，请执行以下命令：

```sh
singularity storage create local --name "my-source" --path "/mnt/dataset/folder"
```

### 2b. 添加 S3 数据源

任何兼容 S3 的存储系统都可以使用，包括 AWS S3、MinIO 等。下面是一个公共数据集的示例：

```sh
singularity storage create s3 aws --name "my-source" --path "public-dataset-test"
```

## 3. 创建一个准备任务
```sh
singularity prep create --source "my-source" --name "my-prep"
```

## 4. 运行准备工作程序
```sh
singularity prep start-scan my-prep my-source
singularity run dataset-worker
```

## 5. 检查准备任务状态和结果
```sh
singularity prep status my-prep
singularity prep list-pieces my-prep
```