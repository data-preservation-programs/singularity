# 手动向一个准备工作添加片段信息。这对于由外部工具准备的片段非常有用。

{% code fullWidth="true" %}
```
NAME:
   singularity prep add-piece - 手动向一个准备工作添加片段信息。这对于由外部工具准备的片段非常有用。

USAGE:
   singularity prep add-piece [command options] <preparation id|name>

CATEGORY:
   片段管理

OPTIONS:
   --piece-cid value   片段的CID
   --piece-size value  片段的大小（默认值："32GiB"）
   --file-path value   CAR文件路径，用于确定文件的大小和根CID
   --root-cid value    CAR文件的根CID
   --help, -h          显示帮助
```
{% endcode %}