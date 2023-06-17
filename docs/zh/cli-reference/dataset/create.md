# 创建一个新的数据集

{% code fullWidth="true" %}
```
命令名称:
   singularity dataset create - 创建一个新的数据集

用法:
   singularity dataset create [命令选项] <数据集名称>

描述:
   <数据集名称>必须是一个唯一的数据集标识符
   数据集是顶层对象，用于区分不同的数据集。

选项：
   --help, -h  显示帮助

   加密

   --encryption-recipient value [ --encryption-recipient value ]  加密接收者的公钥
   --encryption-script value                                      [WIP] 运行自定义加密的EncryptionScript命令

   内联准备

   --output-dir value, -o value [ --output-dir value, -o value ]  用于CAR文件的输出目录（默认值：不需要）

   准备参数

   --max-size value, -M value    要创建的CAR文件的最大大小（默认值："31.5GiB"）
   --piece-size value, -s value  用于片段承诺计算的CAR文件的目标片段大小（默认值：推断）

```
{% endcode %}