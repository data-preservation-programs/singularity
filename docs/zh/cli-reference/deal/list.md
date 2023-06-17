# 列出所有交易

{% code fullWidth="true" %}
```
名称：
   singularity deal list - 列出所有交易

用法：
   singularity deal list [命令选项] [参数...]

选项：
   --dataset value [ --dataset value ]    通过数据集名称筛选交易
   --schedule value [ --schedule value ]  通过计划筛选交易
   --provider value [ --provider value ]  通过提供者筛选交易
   --state value [ --state value ]        通过状态筛选交易：proposed（已提出），published（已发布），active（已激活），expired（已过期），proposal_expired（提议过期），slashed（返还抵押品）
   --help, -h                             显示帮助
```
{% endcode %}