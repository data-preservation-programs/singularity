# 从包含CAR文件的文件夹中提取文件夹或文件到本地目录

{% code fullWidth="true" %}
```
名称:
   singularity extract-car - 从包含CAR文件的文件夹中提取文件夹或文件到本地目录

用法:
   singularity extract-car [命令选项] [参数...]

类别:
   实用程序

选项:
   --input-dir value, -i value  包含CAR文件的输入文件夹。该文件夹将被递归扫描
   --output value, -o value     提取到的输出文件夹或文件。如果不存在将被创建（默认值: "."）
   --cid value, -c value        要提取的文件夹或文件的CID
   --help, -h                   显示帮助信息
```
{% endcode %}