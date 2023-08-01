# 创建交易计划

是时候与存储提供商进行一些交易了。首先运行交易服务。

```
singularity run dealmaker
```

## 一次发送所有交易

对于较小的数据集，您可以一次性将所有交易发送给存储提供商。为了实现这一点，您可以使用以下命令。

```sh
singularity deal schedule create dataset_name provider_id
```

然而，如果数据集很大，存储提供商在交易建议到期之前可能无法处理那么多的交易，因此您可以创建一个计划。

## 定时发送交易

使用相同的命令，您可以创建自己的计划来控制交易向存储提供商的发送速度和频率。

```
--schedule-deal-number value, --number value     触发计划的每个计划的最大交易数，例如30（默认：无限）
--schedule-deal-size value, --size value         触发计划的每个计划的最大交易大小，例如500GB（默认：无限）
--schedule-interval value, --every value         发送批量交易的Cron计划（默认：禁用）
--total-deal-number value, --total-number value  此请求的最大总交易数，例如1000（默认：无限）
```