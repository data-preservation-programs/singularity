# 创建一个计划以将交易发送给存储提供商

{% code fullWidth="true" %}
```
名称：
   singularity deal schedule create - 创建一个计划以将交易发送给存储提供商

用法：
   singularity deal schedule create [命令选项] [参数...]

描述：
   CRON 模式 '--schedule-cron': CRON 模式可以是描述符，也可以是带有可选秒字段的标准 CRON 模式
     标准 CRON：
       ┌───────────── 分钟 (0 - 59)
       │ ┌───────────── 小时 (0 - 23)
       │ │ ┌───────────── 月份中的日期 (1 - 31)
       │ │ │ ┌───────────── 月份 (1 - 12)
       │ │ │ │ ┌───────────── 星期中的日期 (0 - 6)（星期日到星期六）
       │ │ │ │ │                                   
       │ │ │ │ │
       │ │ │ │ │
       * * * * *

     可选的秒字段：
       ┌─────────────  秒 (0 - 59)
       │ ┌─────────────  分钟 (0 - 59)
       │ │ ┌─────────────  小时 (0 - 23)
       │ │ │ ┌─────────────  月份中的日期 (1 - 31)
       │ │ │ │ ┌─────────────  月份 (1 - 12)
       │ │ │ │ │ ┌─────────────  星期中的日期 (0 - 6)（星期日到星期六）
       │ │ │ │ │ │
       │ │ │ │ │ │
       * * * * * *

     描述符：
       @yearly, @annually - 等同于 0 0 1 1 *
       @monthly           - 等同于 0 0 1 * *
       @weekly            - 等同于 0 0 * * 0
       @daily,  @midnight - 等同于 0 0 * * *
       @hourly            - 等同于 0 * * * *

选项：
   --help, -h           显示帮助
   --preparation value  准备工作 ID 或名称
   --provider value     发送交易的存储提供商 ID

   仅限 Boost

   --http-header value, -H value [ --http-header value, -H value ]  要传递给请求的 HTTP 标头（键值对形式）
   --ipni                                                           是否将交易通知 IPNI（默认值：true）
   --url-template value, -u value                                   包含 PIECE_CID 占位符用于 Boost 获取 CAR 文件的 URL 模板，例如 http://127.0.0.1/piece/{PIECE_CID}.car

   交易提案

   --duration value, -d value     持续时间，可以是 epoch 或持续时间格式，例如 1500000, 2400h（默认值：12840h[535 天]）
   --keep-unsealed                是否保留未密封的副本（默认值：true）
   --price-per-deal value         单笔交易的 FIL 价格（默认值：0）
   --price-per-gb value           每 GiB 的 FIL 价格（默认值：0）
   --price-per-gb-epoch value     每个时期每 GiB 的 FIL 价格（默认值：0）
   --start-delay value, -s value  交易开始延迟时间，可以是 epoch 或持续时间格式，例如 1000, 72h（默认值：72h[3 天]）
   --verified                     是否将交易提案标记为已验证（默认值：true）

   限制条件

   --allowed-piece-cid value, --piece-cid value [ --allowed-piece-cid value, --piece-cid value ]                      允许使用的 Piece CID 列表（默认值：任意）
   --allowed-piece-cid-file value, --piece-cid-file value [ --allowed-piece-cid-file value, --piece-cid-file value ]  包含一系列允许的 Piece CID 的文件列表
   --max-pending-deal-number value, --pending-number value                                                            此请求中的最大待处理交易数量，例如 100TiB（默认值：无限）
   --max-pending-deal-size value, --pending-size value                                                                此请求中的最大待处理交易大小，例如 1000（默认值：无限）
   --total-deal-number value, --total-number value                                                                    此请求中的最大交易总数，例如 1000（默认值：无限）
   --total-deal-size value, --total-size value                                                                        此请求中的最大交易总大小，例如 100TiB（默认值：无限）

   计划安排

   --schedule-cron value, --cron value           触发批量交易的 Cron 计划（默认值：禁用）
   --schedule-deal-number value, --number value  每个触发计划的最大交易数量，例如 30（默认值：无限）
   --schedule-deal-size value, --size value      每个触发计划的最大交易大小，例如 500GiB（默认值：无限）

   跟踪

   --notes value, -n value  用于跟踪目的的任何说明或标签

```
{% endcode %}