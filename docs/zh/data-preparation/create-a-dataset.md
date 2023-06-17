---
description: 从初始化数据库开始创建新数据集
---

# 创建数据集

## 初始化数据库

默认情况下，它将使用 `sqlite3` 数据库后端，并在 `$HOME/.singularity` 中初始化数据库文件。

要在生产环境中使用不同的数据库后端，请查看 [deploy-to-production.md](../installation/deploy-to-production.md "mention")

```sh
singularity admin init
```

## 创建新的数据集

数据集是与单个数据集相关的数据源集合。创建了数据集后，您将能够添加数据源以及关联 Filecoin 钱包地址。

```sh
singularity dataset create my_dataset
```

默认情况下，singularity 使用一种称为 Inline Preparation 的技术，不会导出任何 CAR 文件。这是因为对于大多数数据源而言，其实际并未更改，CAR 文件基本上存储与原始数据源相同的内容。

## 下一步

[add-a-data-source.md](add-a-data-source.md "mention")