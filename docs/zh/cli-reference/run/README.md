# 运行不同的 Singularity 组件

{% code fullWidth="true" %}
```
命令名称：
   singularity run - 运行不同的 Singularity 组件

使用方法：
   singularity run command [command options] [arguments...]

命令列表：
   api               运行 Singularity API
   dataset-worker    启动数据集准备工作器以处理数据集扫描和准备任务
   content-provider  启动内容提供器以提供检索请求
   deal-tracker      启动交易追踪器以追踪所有相关钱包的交易
   dealmaker         启动交易生成/追踪工作器以处理交易生成
   spade-api         启动符合 Spade 标准的 API，用于存储提供商交易提案自助服务
   help, h           显示命令列表或有关某个命令的帮助信息

选项：
   --help, -h  显示帮助信息
```
{% endcode %}