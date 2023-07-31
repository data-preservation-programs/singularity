# 更新现有数据集

{% code fullWidth="true" %}
```
命令名：
   singularity dataset update - 更新现有数据集

用法：
   singularity dataset update [命令选项] <数据集名称>

选项：
   --help, -h  显示帮助信息

   加密选项

   --encryption-recipient value [ --encryption-recipient value ]  加密接收者的公钥
   --encryption-script value                                      自定义加密的EncryptionScript命令

   内联准备选项

   --output-dir value, -o value [ --output-dir value, -o value ]  CAR文件的输出目录（默认值：不需要）

   准备参数

   --max-size value, -M value    要创建的CAR文件的最大大小（默认值："30GiB"）
   --piece-size value, -s value  用于计算片段承诺的CAR文件的目标片段大小（默认值：推测）

```
{% endcode %}