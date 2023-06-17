# 从本地路径准备数据集

{% code fullWidth="true" %}
```
名称：
   singularity ez-prep  - 从本地路径准备数据集

用法：
   singularity ez-prep [命令选项] <路径>

类别：
   简便命令

描述：
   此命令可用于最小配置参数从本地路径准备数据集。有关更高级的用法，请使用“数据集”和“数据源”下的子命令。
   您还可以使用此命令进行内存数据库和内联准备的基准测试，即： 
     mkdir dataset
     truncate -s 1024G test.img
     singularity ez-prep --output-dir '' --database-file '' -j $(nproc) ./dataset

选项：
   --max-size value, -M value     创建CAR文件的最大大小（默认为“31.5GiB”）
   --output-dir value, -o value   CAR文件的输出目录。 要使用内联准备，请使用空字符串（默认为“./cars”）
   --concurrency value, -j value  打包并发度（默认值：1）
   --database-file value          用于存储元数据的数据库文件。 要使用内存数据库，请使用空字符串。 （默认值为./ezprep-<name> .db）
   --help, -h                     显示帮助信息
```
{% endcode %}