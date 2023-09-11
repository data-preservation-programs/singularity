# CLI参考

{% code fullWidth="true" %}
```
名称：
   singularity - 一个用于将PB级数据上载到Filecoin网络的大规模客户端工具

使用：
   singularity [全局选项] 命令 [命令选项] [参数...]

描述：
   数据库后端支持：
     Singularity支持多个数据库后端：sqlite3，postgres，mysql5.7+
     使用'--database-connection-string'或$DATABASE_CONNECTION_STRING来指定数据库连接字符串。
       示例-表示postgres  - postgres://user:pass@example.com:5432/dbname
       示例-表示mysql     - mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true
       示例-表示sqlite3   - sqlite:/absolute/path/to/database.db
                   或       - sqlite:relative/path/to/database.db

   网络支持：
     SIngularity的默认设置适用于Mainnet。您可以使用以下环境变量设置其他网络：
       对于Calibration网络：
         * 设置LOTUS_API为https://api.calibration.node.glif.io/rpc/v1
         * 设置MARKET_DEAL_URL为https://marketdeals-calibration.s3.amazonaws.com/StateMarketDeals.json.zst
         * 设置LOTUS_TEST为1
       对于其他所有网络：
         * 设置LOTUS_API为您网络的Lotus API端点
         * 设置MARKET_DEAL_URL为空字符串
         * 设置LOTUS_TEST为0或1，根据网络地址是否以'f'或't'开头来决定
       不建议在同一数据库实例之间切换不同网络。

命令：
   version, v  打印版本信息
   help, h     显示命令列表或一个命令的帮助
   守护进程(daemons)：
     run  运行不同的singularity组件
   操作(operations)：
     admin    管理命令
     deal     复制/成交管理
     wallet   钱包管理
     storage  创建和管理存储系统连接
     prep     创建和管理数据集准备工作
   实用工具(utility)：
     ez-prep      从本地路径准备数据集
     download     从元数据API下载CAR文件
     extract-car  从CAR文件文件夹中提取文件夹或文件到本地目录

全局选项：
   --database-connection-string value  数据库连接字符串（默认值：sqlite:./singularity.db）[$DATABASE_CONNECTION_STRING]
   --help, -h                          显示帮助信息
   --json                              启用JSON输出（默认值：false）
   --verbose                           启用详细输出。这将打印更多结果的列以及完整的错误跟踪（默认值：false）

   Lotus

   --lotus-api value    Lotus RPC API端点（默认值："https://api.node.glif.io/rpc/v1"）[$LOTUS_API]
   --lotus-test         当前环境是否使用Testnet（默认值：false）[$LOTUS_TEST]
   --lotus-token value  Lotus RPC API令牌[$LOTUS_TOKEN]

```
{% endcode %}