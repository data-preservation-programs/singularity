# 创建新数据集

{% code fullWidth="true" %}
```
名称:
   singularity dataset create - 创建新数据集

使用方式:
   singularity dataset create [命令选项] <数据集名称>

描述:
   <数据集名称> 必须是一个唯一标识符，用于区分不同的数据集。
   数据集是一个顶级对象，用于区别不同的数据集。

选项:
   --help, -h  显示帮助信息

   加密

   --encryption-recipient value [ --encryption-recipient value ]  加密接收者的公钥
   --encryption-script value                                      [WIP] 运行自定义加密的 EncryptionScript 命令

   内联准备

   --output-dir value, -o value [ --output-dir value, -o value ]  用于 CAR 文件的输出目录 (默认: 不需要)

   准备参数

   --max-size value, -M value    要创建的 CAR 文件的最大大小 (默认: "31.5GiB")
   --piece-size value, -s value  用于碎片验证计算的 CAR 文件的目标碎片大小 (默认: 推断)
```
{% endcode %}