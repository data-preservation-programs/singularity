# 从卫星地址、API密钥和密码短语创建新的访问授权

{% code fullWidth="true" %}
```
NAME:
   singularity storage create storj new - 从卫星地址、API密钥和密码短语创建新的访问授权

使用方法:
   singularity storage create storj new [command options] [arguments...]

描述:
   --satellite-address
      卫星地址。
      
      自定义卫星地址应匹配格式: `<nodeid>@<address>:<port>`。

      示例:
         | us1.storj.io | US1
         | eu1.storj.io | EU1
         | ap1.storj.io | AP1

   --api-key
      API 密钥。

   --passphrase
      加密密码短语。
      
      要访问现有对象，请输入用于上传的密码短语。


选项:
   --api-key value            API 密钥。[$API_KEY]
   --help, -h                 显示帮助
   --passphrase value         加密密码短语。[$PASSPHRASE]
   --satellite-address value  卫星地址。 (默认值: "us1.storj.io") [$SATELLITE_ADDRESS]

   通用

   --name value  存储的名称 (默认值: 自动生成)
   --path value  存储的路径

```
{% endcode %}