# 发送手动交易提议以促进或传统市场

{% code fullWidth="true" %}
```
命令名称：
   singularity deal send-manual - 发送手动交易提议以促进或传统市场

用法：
   singularity deal send-manual [command options] CLIENT_ADDRESS PROVIDER_ID PIECE_CID PIECE_SIZE

选项：
   --help, -h       显示帮助
   --timeout value  交易提议的超时时间（默认值：1m）

   仅用于促进市场

   --file-size value                            促进市场获取CAR文件所需的文件大小，以字节为单位（默认值：0）
   --http-header value [ --http-header value ]  要随请求一起传递的HTTP头（例如：key=value）
   --ipni                                       是否向IPNI公示交易（默认值：true）
   --url-template value                         URL模板，其中包含PIECE_CID占位符，以便促进市场获取CAR文件，例如http://127.0.0.1/piece/{PIECE_CID}.car

   交易提议

   --duration value, -d value     持续时间，以周期或持续时间格式表示，例如1500000、2400h（默认值：12840h[535 days]）
   --keep-unsealed                是否保留未密封的副本（默认值：true）
   --price-per-deal value         交易的FIL价格（默认值：0）
   --price-per-gb value           每GiB的FIL价格（默认值：0）
   --price-per-gb-epoch value     每周期每GiB的FIL价格（默认值：0）
   --root-cid value               作为交易提议的一部分所需的Root CID，如果为空，则设置为空CID（默认值：空CID）
   --start-delay value, -s value  交易开始的延迟时间，以周期或持续时间格式表示，例如1000、72h（默认值：72h[3 days]）
   --verified                     是否提议以验证方式进行交易（默认值：true）

```
{% endcode %}