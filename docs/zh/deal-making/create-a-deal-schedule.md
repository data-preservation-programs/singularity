# 创建交易计划

现在是时候与存储提供商进行一些交易了。首先运行交易推送服务。

```
singularity run deal-pusher
```

## 一次性发送所有交易

对于较小的数据集，你可以一次性将所有交易发送给存储提供商。为了实现这个目标，可以使用以下命令。

```sh
singularity deal schedule create <preparation> <provider_id>
```

然而，如果数据集很大，在交易提案到期之前，存储提供商可能无法接收那么多交易。因此，你可以创建一个交易计划。

## 定时发送交易

使用相同的命令，你可以创建自己的交易计划，以控制交易发送的速度和频率。

```sh
singularity deal schedule create -h
```