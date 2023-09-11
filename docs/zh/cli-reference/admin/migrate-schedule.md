# 从旧的Singularity MongoDB迁移调度

{% code fullWidth="true" %}
```
NAME:
   singularity admin migrate-schedule - 从旧的Singularity MongoDB迁移调度

用法:
   singularity admin migrate-schedule [命令选项] [参数...]

描述:
   从Singularity V1迁移调度到V2。请注意：
     1. 您必须先完成数据集迁移
     2. 所有新的调度将被创建为“已暂停”状态
     3. 交易状态不会被迁移，因为它将自动由交易跟踪器填充
     4. --output-csv不再支持。我们将在未来提供一个新的工具
     5. 副本数不再作为调度的一部分支持。我们将在未来将其作为可配置策略
     6. --force不再支持。我们可能会在未来添加类似的支持来忽略所有策略限制
     7. --offline不再支持。如果配置了URL模板，传统市场将始终是离线交易，增量市场将始终是在线交易

选项:
   --mongo-connection-string value  MongoDB连接字符串 (默认: "mongodb://localhost:27017") [$MONGO_CONNECTION_STRING]
   --help, -h                       显示帮助
```
{% endcode %}