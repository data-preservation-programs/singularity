# 更新现有数据集

{% code fullWidth="true" %}
```
命令：
   singularity dataset update - 更新现有数据集

用法：
   singularity dataset update [命令选项] <数据集名称>

选项：
   --help, -h  显示帮助信息

   加密

   --encryption-recipient value [ --encryption-recipient value ]  加密收件人的公钥
   --encryption-script value                                      自定义加密的EncryptionScript命令

   内联准备

   --output-dir value, -o value [ --output-dir value, -o value ]  CAR文件的输出目录（默认值：不需要）

   准备参数

   --max-size value, -M value    要创建的CAR文件的最大大小（默认值：“30GiB”）
   --piece-size value, -s value  用于碎片承诺计算的CAR文件的目标碎片大小（默认值：推断）
```
{% endcode %}