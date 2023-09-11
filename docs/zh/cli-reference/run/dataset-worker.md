# 启动数据集准备工作者进行数据集扫描和准备任务处理

{% code fullWidth="true" %}
```
名称:
   singularity run dataset-worker - 启动数据集准备工作者进行数据集扫描和准备任务处理

使用方法:
   singularity run dataset-worker [命令选项] [参数...]

选项:
   --concurrency value  并发工作者的数量 (默认值: 1)
   --enable-scan        启用数据集扫描 (默认值: true)
   --enable-pack        启用数据集打包，计算CIDs，并将其打包为CAR文件 (默认值: true)
   --enable-dag         启用数据集的DAG生成，保持数据集的目录结构 (默认值: true)
   --exit-on-complete   当没有更多任务需要处理时退出工作者 (默认值: false)
   --exit-on-error      当出现任何错误时退出工作者 (默认值: false)
   --help, -h           显示帮助信息
```
{% endcode %}