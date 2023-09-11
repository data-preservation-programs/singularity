# 从本地路径准备数据集

{% code fullWidth="true" %}
```
名称：
   singularity ez-prep - 从本地路径准备数据集

用法：
   singularity ez-prep [命令选项] <路径>

类别：
   实用工具

描述：
   此命令可用于使用最少的可配置参数从本地路径准备数据集。
   对于更高级的用法，请使用`storage`和`data-prep`下的子命令。
   您还可以使用此命令进行内存数据库和内联准备的基准测试，例如：
     mkdir dataset
     truncate -s 1024G dataset/1T.bin
     singularity ez-prep --output-dir '' --database-file '' -j $(($(nproc) / 4 + 1)) ./dataset

选项：
   --max-size value, -M value       创建CAR文件的最大大小（默认值："31.5GiB"）
   --output-dir value, -o value     CAR文件的输出目录。要使用内联准备，请使用空字符串（默认值："./cars"）
   --concurrency value, -j value    打包的并发数（默认值：1）
   --database-file value, -f value  存储元数据的数据库文件。要使用内存数据库，请使用空字符串（默认值：./ezprep-<name>.db）
   --help, -h                       显示帮助信息
```
{% endcode %}