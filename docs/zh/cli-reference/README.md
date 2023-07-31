# CLI 参考指南

{% code fullWidth="true" %}
```
命令：
   singularity - 用于将 PB 级别的数据批量上传至 Filecoin 网络的大规模客户端工具

用法：
   singularity [全局选项] 命令 [命令选项] [参数...]

描述：
   数据库后端支持：
     Singularity 支持多种数据库后端：sqlite3、postgres、mysql5.7+
     使用 '--database-connection-string' 或 $DATABASE_CONNECTION_STRING 来指定数据库连接字符串。
       例如，对于 postgres  - postgres://user:pass@example.com:5432/dbname
       例如，对于 mysql     - mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true
       例如，对于 sqlite3   - sqlite:/绝对路径/to/database.db
                   或者    - sqlite:相对路径/to/database.db

   网络支持：
     Singularity 的默认设置是针对 Mainnet 的。您可以设置以下环境变量来使用其他网络：
       对于 Calibration 网络：
         * 将 LOTUS_API 设置为 https://api.calibration.node.glif.io/rpc/v1
         * 将 MARKET_DEAL_URL 设置为 https://marketdeals-calibration.s3.amazonaws.com/StateMarketDeals.json.zst
       对于其他所有网络：
         * 将 LOTUS_API 设置为您网络的 Lotus API 终点
         * 将 MARKET_DEAL_URL 设置为空字符串
       不建议在同一个数据库实例中切换不同的网络。

命令：
   version, v  打印版本信息
   help, h     显示命令列表或某个命令的帮助信息
   守护进程：
     run  运行不同的 singularity 组件
   简易命令：
     ez-prep  从本地路径准备数据集
   操作：
     admin       管理命令
     deal        复制/交易管理
     dataset     数据集管理
     datasource  数据源管理
     wallet      钱包管理
   工具：
     tool  开发和调试工具
   实用工具：
     download  从元数据 API 下载 CAR 文件

全局选项：
   --database-connection-string value  数据库连接字符串 (默认值：sqlite:./singularity.db) [$DATABASE_CONNECTION_STRING]
   --help, -h                          显示帮助信息
   --json                              启用 JSON 输出 (默认值：false)

   Lotus

   --lotus-api value    Lotus RPC API 终点 (默认值：https://api.node.glif.io/rpc/v1) [$LOTUS_API]
   --lotus-token value  Lotus RPC API 令牌 [$LOTUS_TOKEN]

```
{% endcode %}