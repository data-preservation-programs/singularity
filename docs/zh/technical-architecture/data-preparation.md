---
description: 为了帮助开发者理解 Singularity 中的数据准备流程，本文提供了技术概览。
---

# Singularity 数据准备架构

![Singularity 数据准备模型](data-prep-model.jpg)

# 数据集和数据源

数据集是一组具有复制策略和一个或多个关联 Filecoin 钱包的数据。

每个数据集可以有一个或多个数据源。数据源只是指向一个数据文件夹的指针，可以存储在[RClone](https://github.com/rclone/rclone)支持的存储上，包括本地文件存储。

数据准备在 Singularity 用户执行以下步骤时开始：
1. 创建数据集
2. 向数据集添加数据源
3. 启动数据准备工作器

# 扫描

数据准备的第一阶段是扫描数据源，构建包含文件和文件夹的目录树结构，以及将文件切分成 CAR 文件的计划。

用于表示数据源目录树结构的模型是 Directory 和 Item（其中 Directory 是文件夹，Item 是文件）。在扫描过程中，使用 RClone 来映射数据源的目录结构，然后从这个目录结构中组装 Directory 和 Item 模型，从数据源的根目录开始创建。

对于数据源中的每个 Item，扫描过程还会创建 ItemParts，表示文件的连续部分，每个部分最多为 1GB（可以配置此大小）。此外，将一组 ItemParts 添加到 Chunk 中，Chunk 只是一个包含足够存储在 Filecoin 的 32GB 分片上的单个 CAR 文件的 ItemParts 列表。

在扫描过程结束时，我们得到了一个 Directory 树，每个 Directory 都有一个 Item 列表，每个 Item 又被切分为表示每个文件的部分的 ItemParts，每个部分最多为 1GB。此外，我们还有一组 Chunks，每个 Chunk 链接到一组 ItemParts（足够填满一个分片），我们将它们打包成一个单独的 CAR 文件。

请注意，由于数据源可能会发生变化，因此随着文件和文件夹的添加、更改和删除，数据源也可能会进行重新扫描。

# 打包

一旦数据源被扫描，就可以将其打包为 CAR 文件。打包是将 Chunks 转换为实际写入了每个块的 CAR 文件的过程。

要打包 CAR 文件，需要读取 Chunk 中的每个 ItemPart，并将其分块为指定块大小的 IPLD 原始块，然后将每个原始块写入 CAR 文件。在所有原始块写入完成后，假设 ItemPart 包含多个原始块，则组装并写入 UnixFS 中间节点块的树结构，将原始块链接在一起并生成 ItemPart 的根 CID。完成此过程后，我们就有一个包含了所有 Chunk 中 ItemParts 的原始块和 UnixFS 中间节点块的 CAR 文件。

在打包过程结束时，Singularity 还会将 Car 模型写入其数据库，以表示 CAR 文件，以及为 CAR 中的每个块写入一个 CarBlock。

完成每个 CAR 的写入后，我们回到我们的 Directories 和 Items。对于每个所有 ItemParts 都已写入的 Item，我们构建一个额外的 UnixFS 中间节点树，将一个 Item 中的所有 ItemParts 连接在一起，形成一个单独的 UnixFS 文件。我们还为每个 Directory 组装并更新 UnixFS 目录节点。这些数据暂时存储在数据库中，与 Directory 对象关联。

当打包过程完成时，我们就有了存储数据源中每个 ItemPart 的 CAR 文件。但是，请注意，在此时，虽然我们还组装了一个表示来自数据源的目录（Directories）和文件（Items）的 UnixFS DAG，但我们尚未将其序列化为自己的 CAR 文件，以存储在 Filecoin 上。

# Daggen

因为文件和目录结构随时间变化，将该结构的快照存储到 Filecoin 是一个手动的准备步骤，称为 Daggen。当用户单独启动此步骤时，Pack 过程中组装的 UnixFS DAG 树将被序列化为一个 CAR 文件，并存储到 Filecoin 上。完成此操作后，如果我们将数据准备过程中编写的每个 CAR 都存储到 Filecoin 上，我们将已经存储了从 Filecoin 检索整个数据源快照所需的所有内容。