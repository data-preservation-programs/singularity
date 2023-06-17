# 数据集管理

{% code fullWidth="true" %}
```
命令名称:
   singularity dataset - 数据集管理

用法:
   singularity dataset 命令 [命令选项] [参数...]

命令:
   create         创建一个新的数据集
   list           列出所有数据集
   update         更新现有的数据集
   remove         删除指定的数据集。这不会删除CAR文件。
   add-wallet     为数据集关联钱包。钱包需要使用“singularity wallet import”命令先进行导入。
   list-wallet    列出与数据集关联的所有钱包。
   remove-wallet  从数据集中删除关联的钱包。
   add-piece      为交易目的手动向数据集注册一个片段（CAR 文件）。
   list-pieces    列出所有可用于交易的数据集片段。
   help, h        显示命令列表或一个命令的帮助信息。

选项:
   --help, -h  显示帮助信息
```
{% endcode %}