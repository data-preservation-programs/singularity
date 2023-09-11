# 向提升或传统市场发送手动交易提案

{% code fullWidth="true" %}
```
名称：
   singularity deal send-manual - 向提升或传统市场发送手动交易提案

用法：
   singularity deal send-manual [命令选项] <客户端> <提供者> <片段CID> <片段大小>

描述：
   向提升或传统市场发送手动交易提案
     示例：singularity deal send-manual f01234 f05678 bagaxxxx 32GiB
   注意事项：
     * 客户端地址必须使用“singularity wallet import”导入到钱包中
     * 交易提案将不会保存在数据库中，但将在交易跟踪器运行时进行跟踪
     * 可以通过将LOTUS_API和LOTUS_TOKEN设置为您自己的lotus节点来快速进行地址验证，使用GLIF API

选项：
   --help, -h       显示帮助信息
   --timeout value  交易提案的超时时间（默认值：1m）

   仅限提升

   --file-size value                            提供程序为提取CAR文件而获取的文件大小（默认值：0）
   --http-header value [ --http-header value ]  与请求一起传递的http标头（例如key=value）
   --ipni                                       是否向IPNI宣布交易（默认值：true）
   --url-template value                         带有PIECE_CID占位符的URL模板，供提升使用以获取CAR文件，例如 http://127.0.0.1/piece/{PIECE_CID}.car

   交易提案

   --client value                 发送交易的客户端地址
   --duration value, -d value     以时期格式（i.e. 1500000）或持续时间格式（i.e. 2400h）指定的交易时期（默认值：12840h[535天]）
   --keep-unsealed                是否保留未密封的副本（默认值：true）
   --piece-cid value              交易的片段CID
   --piece-size value             交易的片段大小（默认值："32GiB"）
   --price-per-deal value         单次交易的FIL价格（默认值：0）
   --price-per-gb value           单GB的FIL价格（默认值：0）
   --price-per-gb-epoch value     每个时期的每GB的FIL价格（默认值：0）
   --provider value               存储提供者ID，用于发送交易
   --root-cid value               必须作为交易提案的一部分的根CID，如果为空，则设置为空CID（默认值：空CID）
   --start-delay value, -s value  以时期格式（i.e. 1000）或持续时间格式（i.e. 72h）指定的交易开始延迟（默认值：72h[3天]）
   --verified                     是否提出已验证的交易提案（默认值：true）

```
{% endcode %}