# 加密

## 概览

Singularity支持内置的加密解决方案，可使用提供的收件人（公钥）或甚至硬件PIV令牌（如YubiKeys）加密文件。您还可以提供自定义加密脚本，允许您与任何外部加密解决方案集成\[需要测试]。

## 内置加密

首先，为非对称加密创建公钥-私钥对。 Singularity使用的底层加密库称为[age](https://github.com/FiloSottile/age)。

```sh
go install filippo.io/age/cmd/...@latest
age-keygen -o key.txt
> 公钥: agexxxxxxxxxxxx
```

现在，我们可以设置数据集以使用生成的公钥加密每个文件

```sh
singularity dataset create --encryption-recipient agexxxxxxxxxxxx \
  --output-dir . test
```

由于我们无法内联准备，因为加密相同文件的第二次将产生不同的加密内容，因为在加密过程中引入了初始随机性。

然后，我们可以像以前一样添加数据源，继续我们的数据准备过程。请注意，文件夹结构将不会被加密，因此您可以选择为文件夹结构生成DAG，或者不运行“daggen”命令。在后一种情况下，文件夹结构仅可从Singularity数据库和命令中访问。

## 自定义加密

Singularity还通过提供自定义脚本来提供自定义加密以加密文件流。它可以潜在地与密钥管理服务和自定义加密算法或工具一起使用。