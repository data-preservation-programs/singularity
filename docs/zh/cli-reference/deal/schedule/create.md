# 创建时间表以将交易发送给存储提供者

{% code fullWidth="true" %}
```
名称:
   singularity deal schedule create - 创建时间表以将交易发送给存储提供者

用法:
   singularity deal schedule create [command options] DATASET_NAME PROVIDER_ID

选项：
   --help, -h  显示帮助

   只加速

   --http-header value, -H value    要与请求一起传递的 HTTP 标头（即 key=value）
   --ipni                            是否向 IPNI 公布交易（默认：true）
   --url-template value, -u value    URL 模板，其中包含用于从 boost 中提取 CAR 文件的 PIECE_CID 占位符，例如 http://127.0.0.1/piece/{PIECE_CID}.car

   交易提案

   --duration value, -d value        持续时间（以时代或时间格式），例如 1500000、2400h（默认："12840h"）
   --keep-unsealed                   是否保留未封装副本（默认：true）
   --price-per-deal value            每笔交易的 FIL 价格（默认：0）
   --price-per-gb value              每 GiB 的 FIL 价格（默认：0）
   --price-per-gb-epoch value        每个时代的每 GiB 的 FIL 价格（默认：0）
   --start-delay value, -s value     交易启动延迟（以时代或时间格式），例如 1000、72h（默认："72h"）
   --verified                        是否将交易提议为已验证（默认：true）

   限制

   --allowed-piece-cid value, --piece-cid value [ --allowed-piece-cid value, --piece-cid value ]                      此时间表中允许的 piece CID 列表（默认：Any）
   --allowed-piece-cid-file value, --piece-cid-file value [ --allowed-piece-cid-file value, --piece-cid-file value ]  包含一组允许的 piece CID 列表的文件
   --max-pending-deal-number value, --pending-number value                                                            此请求的全部待处理交易数量上限（默认：无限制）
   --max-pending-deal-size value, --pending-size value                                                                此请求的全部待处理交易大小上限（默认：无限制）
   --total-deal-size value, --total-size value                                                                        此请求的最大总交易大小，例如 100TB（默认：无限制）

   安排

   --schedule-cron value, --cron value              定时发送批量交易的 cron 安排（默认：禁用）
   --schedule-deal-number value, --number value     触发安排的最大交易数量，例如 30（默认：无限制）
   --schedule-deal-size value, --size value         触发安排的最大交易大小，例如 500GB（默认：无限制）
   --total-deal-number value, --total-number value  此请求的最大总交易数量，例如 1000（默认：无限制）

   跟踪

   --notes value, -n value  用于存储在请求中的任何注释或标签，以用于跟踪目的

```
{% endcode %}