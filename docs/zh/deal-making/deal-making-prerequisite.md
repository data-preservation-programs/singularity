# 交易前提条件

## 找到存储提供者

目前，Singularity 不会帮助您找到接受您交易的存储提供者。您可以使用许多资源来寻找高质量的存储提供者，比如：

* TODO

## 创建 Filecoin 钱包

在进行交易之前，您必须创建 Filecoin 钱包。您不能使用带账本的钱包或交易所钱包。要创建 Filecoin 钱包，您可以运行下面的命令：

```
singularity wallet create
```

这将生成一个钱包地址以及与之关联的私钥。此钱包目前还不能用于交易，因为 blockchain 尚未识别它。现在是一个很好的时间将 0 FIL 转入此钱包，这样所有人都会知道它。

一旦此钱包在链上记录，上面的命令将完成，您就可以准备进行交易了。

如果您已经有现有的钱包，您可以使用下面的命令导入：

```
singularity wallet import xxx
```

## 【可选】获得 [datacap](https://docs.filecoin.io/basics/how-storage-works/filecoin-plus/#datacap)

在当前的市场状况下，大多数存储提供者更愿意接受[验证交易](https://docs.filecoin.io/storage-provider/filecoin-deals/verified-deals/)而不是普通交易。如果您的数据集大于几 TiB，最好向 [Filplus 治理团队和公证人](https://github.com/filecoin-project/notary-governance)申请 datacap。

## 下一步

[创建交易时间表](create-a-deal-schedule.md)