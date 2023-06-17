# 发送手动的交易提案来促进或传统市场

{% code fullWidth="true" %}
```
命令名称：
   singularity deal send-manual - 发送手动的交易提案来促进或传统市场

用法：
   singularity deal send-manual [命令选项] 客户端地址 提供者 ID  Piece_CID Piece_Size

选项:
   --help, -h  显示帮助

   仅限 Boost

   --file-size 值                            Boost 获取 CAR 文件时的文件大小（默认值：0）
   --http-header 值 [ --http-header 值 ]      http 头文件将随请求一起传递（即 键=值）
   --ipni                                     是否向 IPNI 公布交易提案（默认值：true）
   --url-template 值                          具有 PIECE_CID 占位符的 URL 模板供 Boost 获取 CAR 文件。 例如：http://127.0.0.1/piece/{PIECE_CID}.car

   交易提案

   --duration, -d 值            交易期限，以天为单位（默认值：535）
   --keep-unsealed               是否保留未封装的副本（默认值：true）
   --price, -p 值               FIL 币每 32GiB 交易期限内的价格（默认值：0）
   --root-cid 值                作为交易提案的一部分所需的根 CID 。如果为空，则将设置为空 CID（默认值：Empty CID）
   --start-delay, -s 值         交易开始的延迟时间（以天为单位）（默认值：3）
   --verified                  是否作为已验证的交易提案（默认值：true）

   Lotus

   --lotus-api 值      Lotus RPC API 终端节点，仅用于获取矿工信息（默认值："https://api.node.glif.io/rpc/v1"）
   --lotus-token 值    Lotus RPC API 令牌，仅用于获取矿工信息

```
{% endcode %}