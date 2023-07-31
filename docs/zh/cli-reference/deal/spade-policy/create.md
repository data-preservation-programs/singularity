# 创建自交易提议的 SPADE 策略

{% code fullWidth="true" %}
```
名称：
   singularity deal spade-policy create - 创建自交易提议的 SPADE 策略

用法：
   singularity deal spade-policy create [命令选项] DATASET_NAME [...PROVIDER_ID]

选项：
   --min-delay value     提议交易开始日期的最小延迟天数（默认值：3）
   --max-delay value     提议交易开始日期的最大延迟天数（默认值：3）
   --min-duration value  提议交易持续时间的最小天数（默认值：535）
   --max-duration value  提议交易持续时间的最大天数（默认值：535）
   --verified            提议交易是否为已验证（默认值：true）
   --price value         交易价格以每32GiB为单位计算整个持续时间的费用（默认值：0）
   --help, -h            显示帮助信息
```
{% endcode %}