# 创建一个定时任务，向存储提供者发送交易

{% code fullWidth="true" %}
```
命令名称：
   singularity deal schedule create - 创建一个定时任务，向存储提供者发送交易

用法：
   singularity deal schedule create [命令选项] 数据集名称 提供者ID

选项：  
   --help，-h  显示帮助

   仅限于增强Boost

   --http-header value, -H value [ --http-header value, -H value ]  客户端与服务端之间的HTTP请求头（例如key=value）
   --ipni                                                           是否要将交易告知IPNI（默认值：true）
   --url-template value, -u value                                   带有PIECE_CID占位符的URL模板，用于增强提取CAR文件，例如：http://127.0.0.1/piece/{PIECE_CID}.car

   交易提议

   --duration value, -d value     交易时长，单位为天数（默认值：530）
   --keep-unsealed                是否保留未封装的复制品（默认值：true）
   --price value, -p value        32GiB交易期间价格，以Fil为单位（默认值：0）
   --start-delay value, -s value  交易开始前的延迟天数（默认值：3）
   --verified                     是否作为已验证的交易提议（默认值：true）

   限制

   --allowed-piece-cid value, --piece-cid value [ --allowed-piece-cid value, --piece-cid value ]                      允许在此计划中使用的piece CID列表（默认值：Any）
   --allowed-piece-cid-file value, --piece-cid-file value [ --allowed-piece-cid-file value, --piece-cid-file value ]  包含允许使用的piece CID列表的文件列表
   --max-pending-deal-number value, --pending-number value                                                            整个请求中最大挂起交易数量（默认值：无限制）
   --max-pending-deal-size value, --pending-size value                                                                整个请求中最大挂起交易总大小（默认值：无限制）
   --total-deal-size value, --total-size value                                                                        此请求中的最大交易总大小，例如100TB（默认值：无限制）

   调度

   --schedule-deal-number value, --number value     触发计划表中最大的交易数量，例如30（默认值：无限制）
   --schedule-deal-size value, --size value         触发计划表中的最大交易总大小，例如500GB（默认值：无限制）
   --schedule-interval value, --every value         触发批量交易发送的定时任务规则（默认值：禁用）
   --total-deal-number value, --total-number value  此请求中的最大交易总数，例如1000（默认值：无限制）

   跟踪

   --notes value, -n value  用于跟踪的任何说明或标记。

```
{% endcode %}