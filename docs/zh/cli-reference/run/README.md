# 运行不同的Singularity组件

{% code fullWidth="true" %}
```
NAME:
   singularity run - 运行不同的Singularity组件

USAGE:
   singularity run 命令 [命令选项] [参数...]

COMMANDS:
   api               运行Singularity API
   dataset-worker    启动数据集准备工作进程以处理数据集扫描和准备任务
   content-provider  启动内容提供者以服务检索请求
   dealmaker         启动交易处理/跟踪工人以处理交易制作
   spade-api         启动符合Spade规范的API，以提供自助存储供应商交易提议
   help, h           显示命令列表或某个命令的帮助信息

OPTIONS:
   --help, -h  显示帮助信息
```
{% endcode %}