# 数据集管理

{% code fullWidth="true" %}
```
命令名称：
   singularity 数据集 - 数据集管理

用法：
   singularity 数据集 命令 [命令选项] [参数...]

命令列表：
   create         创建新的数据集
   list           列出所有数据集
   update         更新现有数据集
   remove         移除指定的数据集。这不会删除 CAR 文件。
   add-wallet     将钱包与数据集关联。钱包需要使用 `singularity 钱包导入` 命令先导入。
   list-wallet    列出与数据集关联的所有钱包
   remove-wallet  从数据集中移除关联的钱包
   add-piece      手动将一个部件（CAR 文件）注册到数据集，以进行交易
   list-pieces    列出所有对于交易可用的数据集的部件
   help, h        显示命令列表或某个命令的帮助信息

选项：
   --help, -h  显示帮助信息
```
{% endcode %}