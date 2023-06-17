# 创建自我交易提案的 SPADE 策略

{% code fullWidth="true" %}
```
名称：
   singularity deal spade-policy create - 为自我交易提案创建 SPADE 策略

用法：
   singularity deal spade-policy create [命令选项] DATASET_NAME [...PROVIDER_ID]

选项：
   --min-delay 值      交易开始时期的最小延迟天数（默认值：3）
   --max-delay 值      交易开始时期的最大延迟天数（默认值：3）
   --min-duration 值   交易开始时期的最短持续时间（默认值：535）
   --max-duration 值   交易开始时期的最长持续时间（默认值：535）
   --verified          是否将提案作为已验证的提出（默认值：true）
   --price 值           整个持续时间内每32GiB的价格（默认值：0）
   --help，-h          显示帮助
```
{% endcode %}