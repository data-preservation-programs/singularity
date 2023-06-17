# 创建成交时间表

现在可以和存储供应商达成一些成交了。首先要运行成交服务。

```
singularity run dealmaker
```

## 一次性发送所有成交

如果数据集较小，您可以一次性将所有成交发送给您的存储供应商。您可以使用以下命令来实现：

```sh
singularity deal schedule create dataset_name provider_id
```

然而，如果数据集较大，存储提供商在成交建议到期之前可能无法处理那么多成交。所以您可以创建一个时间表。

## 按时间表发送成交

使用相同的命令，您可以创建自己的时间表来控制成交向存储供应商的生成速度和频率。

```
--schedule-deal-number 值，--number 值          每个触发计划的最大成交量，例如30（默认值：无限制）
--schedule-deal-size 值，--size 值              每个触发计划的最大成交大小，例如500GB（默认值：无限制）
--schedule-interval 值，--every 值              Cron调度发送批次成交（默认值：未启用）
--total-deal-number 值，--total-number 值       此请求的最大总成交量，例如1000（默认值：无限制）
```