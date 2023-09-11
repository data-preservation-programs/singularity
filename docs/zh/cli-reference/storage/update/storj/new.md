# 从卫星地址、API 密钥和口令创建一个新的访问授权

{% code fullWidth="true" %}
```
名称:
   singularity storage update storj new - 从卫星地址、API 密钥和口令创建一个新的访问授权

用法:
   singularity storage update storj new [命令选项] <名称或ID>

描述:
   --satellite-address
      卫星地址。
      
      自定义的卫星地址应与以下格式匹配: `<节点ID>@<地址>:<端口>`。

      例子:
         | us1.storj.io | 美国1
         | eu1.storj.io | 欧洲1
         | ap1.storj.io | 亚太1

   --api-key
      API 密钥。

   --passphrase
      加密口令。
      
      如需访问现有的对象，请输入用于上传的口令。


选项:
   --satellite-address value  卫星地址。 (默认值: "us1.storj.io") [$SATELLITE_ADDRESS]
   --api-key value            API 密钥。 [$API_KEY]
   --passphrase value         加密口令。 [$PASSPHRASE]
   --help, -h                 显示帮助
```
{% endcode %}