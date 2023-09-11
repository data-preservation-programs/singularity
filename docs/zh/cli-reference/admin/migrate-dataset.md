# 从旧的Singularity MongoDB迁移数据集

{% code fullWidth="true" %}
```
命令:
   singularity admin migrate-dataset - 从旧的Singularity MongoDB迁移数据集

使用方法:
   singularity admin migrate-dataset [命令选项] [参数]

描述:
   从Singularity V1迁移数据集到V2。步骤包括：
     1. 在V2中创建源存储和输出存储，并将它们附加到数据准备。
     2. 在新数据集中创建所有文件夹结构和文件。
   注意事项：
     1. 创建的数据准备与新数据集的工作程序不兼容。
        因此，请勿尝试恢复数据准备或将新文件推送到迁移的数据集。
        您可以无问题地查看或浏览数据集。
     2. 由于复杂性，文件夹CID将不会生成或迁移。

选项:
   --mongo-connection-string value  MongoDB连接字符串（默认值："mongodb://localhost:27017"）[$MONGO_CONNECTION_STRING]
   --skip-files                     跳过迁移文件和文件夹的详细信息。这将加快迁移速度。如果只想进行交易，则非常有用。（默认值：false）
   --help, -h                       显示帮助
```
{% endcode %}