# 创建一个新的准备

{% code fullWidth="true" %}
```
名称：
   singularity prep create - 创建一个新的准备

用法：
   singularity prep create [命令选项] [参数...]

类别：
   准备管理

选项：
   --delete-after-export              是否在导出为 CAR 文件后删除源文件（默认：false）
   --help, -h                         显示帮助
   --max-size value                   每个 CAR 文件的最大大小（默认："31.5GiB"）
   --name value                       准备的名称（默认：自动生成）
   --output value [ --output value ]  准备使用的输出存储的 ID 或名称
   --piece-size value                 用于计算分块承诺的 CAR 文件的目标分块大小（默认：由 --max-size 决定）
   --source value [ --source value ]  准备使用的源存储的 ID 或名称

   通过本地输出路径进行快速创建

   --local-output value [ --local-output value ]  准备使用的本地输出路径。这是一个方便的标志，将创建带有提供的路径的输出存储

   通过本地源路径进行快速创建

   --local-source value [ --local-source value ]  准备使用的本地源路径。这是一个方便的标志，将创建带有提供的路径的源存储

```
{% endcode %}