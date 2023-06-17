# 手动向数据集注册一个碎片（CAR 文件），以进行交易目的

{% code fullWidth="true" %}
```
名称：
   singularity dataset add-piece - 手动向数据集注册一个碎片（CAR 文件），以进行交易目的

用法：
   singularity dataset add-piece [command options] <dataset_name> <piece_cid> <piece_size>

选项：
   --file-path value, -p value  CAR 文件的路径，用于确定文件的大小和根 CID
   --file-size value, -s value  CAR 文件的大小，如果未提供，将由 CAR 文件确定（默认值：0）
   --root-cid value, -r value   CAR 文件的根 CID，如果未提供，将由 CAR 文件头确定。用于填充存储交易的标签字段
   --help, -h                   显示帮助
```
{% endcode %}