# 启动一个提供检索请求的内容提供者

{% code fullWidth="true" %}
```
名称：
   singularity run content-provider - 启动一个提供检索请求的内容提供者

用法：
   singularity run content-provider [命令选项] [参数...]

选项：
   --help, -h  显示帮助信息

   Bitswap检索

   --enable-bitswap                                 启用Bitswap检索（默认：false）
   --libp2p-identity-key value                      libp2p对等节点的Base64编码的私钥（默认：自动生成）
   --libp2p-listen value [ --libp2p-listen value ]  用于libp2p连接的监听地址

   HTTP检索

   --enable-http      启用HTTP检索（默认：true）
   --http-bind value  绑定HTTP服务器的地址（默认："127.0.0.1:7777"）

```
{% endcode %}