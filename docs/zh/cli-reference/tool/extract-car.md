# 从一个包含CAR文件的文件夹中提取文件夹或文件到本地目录

{% code fullWidth="true" %}
```
USAGE:
   singularity工具 extract-car [command options] [arguments...]

OPTIONS:
   --input-dir value, -i value  包含CAR文件的输入文件夹。将会递归扫描此文件夹内的文件
   --output value, -o value     要提取到的输出文件夹或文件。如果不存在，将会被创建（默认为当前目录）
   --cid value, -c value        要提取的文件夹或文件的CID
   --help, -h                   显示帮助信息
```
{% endcode %}