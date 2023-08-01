# 成交前提条件

## 寻找存储提供商

目前，Singularity并不能帮助您寻找接受您的交易的存储提供商。您可以使用许多资源来寻找高质量的存储提供商，例如：

- TODO

## 创建Filecoin钱包

在进行交易之前，您必须创建一个Filecoin钱包。您不能使用硬件钱包或交易所钱包。要创建一个Filecoin钱包，您可以运行以下命令：

```sh
singularity wallet create
```

这将生成一个钱包地址以及与钱包相关联的私钥。这个钱包目前还不能用于交易，因为它尚未被区块链识别到。现在是时候将0 FIL转入此钱包，以便大家都知道它的存在。

一旦这个钱包被记录在链上，上述命令就会完成，您就可以开始进行交易了。

或者，如果您已经拥有现有的钱包，可以使用以下命令导入：

```sh
singularity wallet import xxx
```

## [可选] 获取[数据容量(datacap)](https://docs.filecoin.io/basics/how-storage-works/filecoin-plus/#datacap)

在当前市场情况下，大多数存储提供商更喜欢[验证的交易(verified deals)](https://docs.filecoin.io/storage-provider/filecoin-deals/verified-deals/)，而不是普通交易。如果您的数据集超过几TiB，最好向[Filplus治理团队和公证人](https://github.com/filecoin-project/notary-governance)申请数据容量。

## 下一步

[创建交易计划](create-a-deal-schedule.md)