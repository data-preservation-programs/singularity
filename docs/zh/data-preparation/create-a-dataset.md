---
description: 开始从初始化数据库并创建一个新数据集
---

# 创建数据集

## 初始化数据库

默认情况下，它将使用 `sqlite3` 数据库后端并在 `$HOME/.singularity` 目录下初始化数据库文件。

要在生产环境中使用不同的数据库后端，请查看 [deploy-to-production.md](../installation/deploy-to-production.md "mention")

```sh
singularity admin init
```

## 创建新数据集

数据集是与单个数据集相关的数据源的集合。创建数据集后，您将能够添加数据源以及关联 Filecoin 钱包地址。

```sh
singularity dataset create my_dataset
```

默认情况下，Singularity 使用一种称为 Inline Preparation 的技术，它不会导出任何 CAR 文件。这是因为对于大多数数据源，它们并不会发生变化，CAR 文件实际上存储的内容与原始数据源相同。

## 下一步

[add-a-data-source.md](add-a-data-source.md "mention")