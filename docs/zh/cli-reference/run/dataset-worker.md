# 启动数据集准备工作程序以处理数据集扫描和准备任务

{% code fullWidth="true" %}
```
命令名：
   singularity run dataset-worker - 启动数据集准备工作程序以处理数据集扫描和准备任务

用法：
   singularity run dataset-worker [命令选项] [参数...]

选项：
   --concurrency value  并发的工作程序数 (默认值: 1) [$DATASET_WORKER_CONCURRENCY]
   --enable-scan        启用数据集扫描 (默认值: true) [$DATASET_WORKER_ENABLE_SCAN]
   --enable-pack        启用数据集的打包，它计算CIDs并将它们打包到CAR文件中 (默认值: true) [$DATASET_WORKER_ENABLE_PACK]
   --enable-dag         启用数据集的 DAG 生成，它保持数据集的目录结构 (默认值: true) [$DATASET_WORKER_ENABLE_DAG]
   --exit-on-complete   当没有更多的工作要做时退出工作程序 (默认值: false) [$DATASET_WORKER_EXIT_ON_COMPLETE]
   --exit-on-error      当出现任何错误时退出工作程序 (默认值: false) [$DATASET_WORKER_EXIT_ON_ERROR]
   --help, -h           显示帮助信息
```
{% endcode %}