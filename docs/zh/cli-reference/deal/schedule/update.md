# 更新现有的调度

{% code fullWidth="true" %}
```
名称:
   singularity deal schedule update - 更新现有的调度

用法:
   singularity deal schedule update [命令选项] <schedule_id>

描述:
   CRON模式 '--schedule-cron': CRON模式可以是描述符或带有可选秒字段的标准CRON模式
     标准CRON模式:
       ┌───────────── 分钟 (0 - 59)
       │ ┌───────────── 小时 (0 - 23)
       │ │ ┌───────────── 每月的日期 (1 - 31)
       │ │ │ ┌───────────── 月份 (1 - 12)
       │ │ │ │ ┌───────────── 星期几 (0 - 6) (星期天到星期六)
       │ │ │ │ │                                   
       │ │ │ │ │
       │ │ │ │ │
       * * * * *

     可选的秒字段:
       ┌─────────────  秒 (0 - 59)
       │ ┌─────────────  分钟 (0 - 59)
       │ │ ┌─────────────  小时 (0 - 23)
       │ │ │ ┌─────────────  每月的日期 (1 - 31)
       │ │ │ │ ┌─────────────  月份 (1 - 12)
       │ │ │ │ │ ┌─────────────  星期几 (0 - 6) (星期天到星期六)
       │ │ │ │ │ │
       │ │ │ │ │ │
       * * * * * *

     描述符:
       @yearly, @annually - 等同于 0 0 1 1 *
       @monthly           - 等同于 0 0 1 * *
       @weekly            - 等同于 0 0 * * 0
       @daily,  @midnight -等同于 0 0 * * *
       @hourly            - 等同于 0 * * * *

选项:
   --help, -h  显示帮助

   仅提升

   --http-header value, -H value [ --http-header value, -H value ]  要随请求传递的HTTP标头（即 key=value）。这将替换现有的标头值。要删除标头，请使用 --http-header "key="。要删除所有标头，请使用 --http-header ""
   --ipni                                                           是否要将交易公布给IPNI（默认值：true）
   --url-template value, -u value                                   包含用于提升获取CAR文件的PIECE_CID占位符的URL模板，例如http://127.0.0.1/piece/{PIECE_CID}.car

   交易建议

   --duration value, -d value     持续时间，以纪元或持续时间格式表示，例如 1500000，2400h
   --keep-unsealed                是否保留未封存的副本（默认值：true）
   --price-per-deal value         每次交易的FIL价格（默认值：0）
   --price-per-gb value           每GiB的FIL价格（默认值：0）
   --price-per-gb-epoch value     每个纪元每GiB的FIL价格（默认值：0）
   --start-delay value, -s value  交易启动的延迟时间，以纪元或持续时间格式表示，例如 1000，72h
   --verified                     是否将交易建议作为经过验证的交易（默认值：true）

   限制条件

   --allowed-piece-cid value, --piece-cid value [ --allowed-piece-cid value, --piece-cid value ]                      此调度中允许的piece CID列表。仅追加。
   --allowed-piece-cid-file value, --piece-cid-file value [ --allowed-piece-cid-file value, --piece-cid-file value ]  包含要允许的piece CID列表的文件列表。仅追加。
   --max-pending-deal-number value, --pending-number value                                                            对于这个请求，最大待处理交易数量，例如100TiB（默认值：0）
   --max-pending-deal-size value, --pending-size value                                                                对于这个请求，最大待处理交易大小，例如1000
   --total-deal-number value, --total-number value                                                                    这个请求的最大总交易数量，例如1000（默认值：0）
   --total-deal-size value, --total-size value                                                                        这个请求的最大总交易大小，例如100TiB

   调度

   --schedule-cron value, --cron value           触发批量交易的Cron调度
   --schedule-deal-number value, --number value  触发调度的最大交易数量，例如30（默认值：0）
   --schedule-deal-size value, --size value      触发调度的最大交易大小，例如500GiB

   跟踪

   --notes value, -n value  与请求一起存储的任何注释或标记，用于跟踪目的

```
{% endcode %}