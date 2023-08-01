# 启动数据集准备工作者来处理数据集扫描和准备任务

{% code fullWidth="true" %}
```
名称:
   singularity run dataset-worker - 启动数据集准备工作者来处理数据集扫描和准备任务

用法:
   singularity run dataset-worker [命令选项] [参数...]

选项:
   --concurrency 值  要运行的并发工作者数量（默认值：1）[$DATASET_WORKER_CONCURRENCY]
   --enable-scan        启用数据集扫描（默认值：true）[$DATASET_WORKER_ENABLE_SCAN]
   --enable-pack        启用数据集打包，计算 CIDs 并将其打包为 CAR 文件（默认值：true）[$DATASET_WORKER_ENABLE_PACK]
   --enable-dag         启用数据集生成 DAG，保持数据集的目录结构（默认值：true）[$DATASET_WORKER_ENABLE_DAG]
   --exit-on-complete   当没有更多任务可执行时退出工作者（默认值：false）[$DATASET_WORKER_EXIT_ON_COMPLETE]
   --exit-on-error      当有任何错误时退出工作者（默认值：false）[$DATASET_WORKER_EXIT_ON_ERROR]
   --help, -h           显示帮助信息
```
{% endcode %}