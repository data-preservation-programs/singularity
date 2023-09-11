# 运行不同的Singularity组件

## 用法：
```
singularity run 命令 [命令选项] [参数...]
```

## 命令：
- api：运行Singularity API
- dataset-worker：启动一个数据集准备工作程序，用于处理数据集扫描和准备任务
- content-provider：启动一个内容提供程序，用于提供检索请求服务
- deal-tracker：启动一个交易追踪器，用于跟踪所有相关钱包的交易
- deal-pusher：启动一个交易推送程序，用于监控交易计划并将交易推送给存储提供程序
- help, h：显示命令列表或某个命令的帮助信息

## 选项：
- --help, -h：显示帮助信息