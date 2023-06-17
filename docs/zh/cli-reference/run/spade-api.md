# 启动一个Spade兼容的API，用于存储提供商交易提议自助服务

{% code fullWidth="true" %}
```
命令名称：
   singularity run spade-api - 启动一个Spade兼容的API，用于存储提供商交易提议自助服务

用法：
   singularity run spade-api [命令选项] [参数...]

选项：
   --bind value        API服务器的绑定地址（默认值：“127.0.0.1:9091”）
   --help，-h          显示帮助信息

   Lotus

   --lotus-api value    Lotus RPC API端点，仅用于获取矿工信息（默认值：“https://api.node.glif.io/rpc/v1”）
   --lotus-token value  Lotus RPC API令牌，仅用于获取矿工信息

```
{% endcode %}