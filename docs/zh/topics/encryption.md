# 加密

## 概述

Singularity支持内置的加密解决方案，可以使用提供的接收方（公钥）或甚至硬件PIV令牌（如YubiKeys）加密文件。您还可以提供一个自定义的加密脚本，以便与任何外部加密解决方案集成\[需要测试]。

## 内置加密

首先创建一个用于非对称加密的公私钥对。Singularity使用的底层加密库名为[age](https://github.com/FiloSottile/age)。

```sh
go install filippo.io/age/cmd/...@latest
age-keygen -o key.txt
> 公钥: agexxxxxxxxxxxx
```

现在，我们可以设置一个数据集，使用生成的公钥对每个文件进行加密

```sh
singularity dataset create --encryption-recipient agexxxxxxxxxxxx \
  --output-dir . test
```

禁用内联准备，因为由于加密过程中引入的初始随机性，无法对相同的文件进行二次加密，结果会有所不同。

然后可以添加数据源并继续数据准备过程。请注意，文件夹结构将不会被加密，因此您可以选择为文件夹结构生成DAG，或者不运行`daggen`命令。在后一种情况下，文件夹结构只能从Singularity数据库和命令中访问。

## 自定义加密

Singularity还通过提供自定义脚本来实现自定义加密。它可以与密钥管理服务和自定义加密算法或工具一起使用。