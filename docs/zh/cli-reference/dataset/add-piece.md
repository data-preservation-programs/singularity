# 手动注册一个（CAR文件）到数据集中以进行交易目的

{% code fullWidth="true" %}
```
名称:
   singularity dataset add-piece - 手动注册一个（CAR文件）到数据集中以进行交易目的

用法:
   singularity dataset add-piece [命令选项] <数据集名称> <片段CID> <片段大小>

描述:
   如果您已经有CAR文件：
     singularity dataset add-piece -p <CAR文件路径> <数据集名称> <片段CID> <片段大小>

   如果您没有CAR文件但您知道RootCID：
     singularity dataset add-piece -r <根CID> <数据集名称> <片段CID> <片段大小>

   如果您既没有CAR文件也不知道RootCID：
     singularity dataset add-piece -r <根CID> <数据集名称> <片段CID> <片段大小>
   但在这种情况下，交易将不会正确设置rootCID，因此可能无法与检索测试良好配合使用。

选项:
   --file-path value, -p value  CAR文件的路径，用于确定文件的大小和根CID
   --root-cid value, -r value   CAR文件的根CID，如果未提供，将由CAR文件头确定。用于填充存储交易的标签字段
   --help, -h                   显示帮助
```
{% endcode %}