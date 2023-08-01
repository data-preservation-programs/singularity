---
description: 为了帮助开发者了解 Singularity 中的数据准备过程，本文档提供了技术概述。
---

# Singularity 数据准备架构

![Singularity 数据准备模型](data-prep-model.jpg)

# 数据集和数据源

数据集是具有复制共享策略和一个或多个关联的 Filecoin 钱包的数据集合。

每个数据集可以有一个或多个数据源。数据源只是指向存储在 [RClone](https://github.com/rclone/rclone) 支持的存储中（包括本地文件存储）的数据文件夹的指针。

数据准备开始于 Singularity 用户：
1. 创建一个数据集
2. 将数据源添加到数据集中
3. 启动数据准备工作器

# 扫描

数据准备的第一个阶段是扫描数据源，以构建包含文件和文件夹的目录树结构，以及将文件划分为 CAR 文件的计划。

用于表示数据源目录树结构的模型是 Directory 和 Item（其中 Directory 是文件夹，Item 是文件）。在扫描过程中，使用 RClone 来映射数据源的目录结构，然后从这个目录结构中组装 Directory 和 Item 模型，从数据源的根目录开始创建的。

对于数据源中的每个 Item，扫描过程还会创建 ItemParts，用于表示文件的连续部分，每个部分的大小最大为 1GB（这个大小可以进行配置）。此外，一组 ItemParts 也会被添加到一个 Chunk 中，Chunk 是一个包含足够大的 ItemParts 的列表，用于存储在 Filecoin 的 32GB 份额中。

在扫描过程结束时，我们得到了一个目录树结构，其中每个目录都有一组 Item，每个 Item 被划分为最多 1GB 的 ItemParts。此外，我们还有一组 Chunks，每个 Chunk 链接到一组 ItemParts（足够组成一个片段），我们将这些 ItemParts 组装成一个单独的 CAR 文件。

请注意，由于数据源可能发生变化，因此可以对其进行重新扫描，以捕捉文件和文件夹的添加、更改和删除情况。

# 打包

一旦源数据被扫描，就可以开始将其打包成 CAR 文件。打包是将 Chunks 转换为实际的写入 CAR 文件的过程。

为了打包一个 CAR 文件，需要读取 Chunk 中的每个 ItemPart，并将其分成指定块大小的 IPLD Raw 块，然后将每个块写入 CAR 文件中。在所有的原始块都写入完成后，假设 ItemPart 包含多个原始块，我们会组装 UnixFS 中间节点块的树状结构，并将其连接起来，生成一个 ItemPart 的根 CID。完成此过程后，我们将得到一个包含所有 ItemParts 的原始块和 UnixFS 中间节点块的 CAR 文件。

在打包过程结束时，Singularity 还会将 Car 模型写入其数据库，以表示 Car 文件，以及为 CAR 中的每个块写入一个 CarBlock。

在完成每个 CAR 的写入时，我们回到目录和文件上。对于所有 ItemParts 都已写入的 Item，我们构建了另一个 UnixFS 中间节点树，将 Item 中的所有 ItemParts 连接成一个单独的 UnixFS 文件。我们还为每个 Directory 组装和更新 UnixFS 目录节点。这些数据被临时存储在数据库中，与 Directory 对象链接在一起。

当打包过程完成后，我们就有了用于存储数据源中每个 ItemPart 的 CAR 文件。但是请注意，此时，虽然我们还组装了表示数据源的目录（Directories）和文件（Items）的 UnixFS DAG，但我们还没有将其序列化为自己的 CAR 文件，以存储在 Filecoin 上。

# Daggen

由于文件和目录结构可能随时间变化，将此结构的快照存储到 Filecoin 中是一个手动的准备步骤，称为 Daggen。当用户单独启动此步骤时，Pack 过程期间组装的 UnixFS DAG 树将被序列化为 CAR 文件，以存储在 Filecoin 上。完成此步骤后，如果我们将数据准备过程中编写的每个 CAR 都存储到 Filecoin 上，我们将存储了从 Filecoin 检索整个数据源快照所需的所有内容。