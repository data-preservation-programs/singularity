# Storj 去中心化云存储

{% code fullWidth="true" %}
```
名称:
   singularity datasource add storj - Storj 去中心化云存储

用法:
   singularity datasource add storj [命令选项] <数据集名称> <源路径>

说明:
   --storj-provider
      选择一种身份验证方法。

      示例:
         | existing | 使用现有的 Access Grant。
         | new      | 从卫星地址、API 密钥和密码短语创建一个新的 Access Grant。

   --storj-access-grant
      [Provider] - existing
         Access Grant。

   --storj-satellite-address
      [Provider] - new
         卫星地址。

         自定义卫星地址应与格式匹配: `<nodeid>@<address>:<port>`。

         示例:
            | us1.storj.io | 美国1区
            | eu1.storj.io | 欧洲1区
            | ap1.storj.io | 亚太1区

   --storj-api-key
      [Provider] - new
         API 密钥。

   --storj-passphrase
      [Provider] - new
         加密密码短语。

         要访问现有对象，请输入用于上传的密码短语。


选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [Dangerous] 导出为 CAR 文件并删除数据集文件。 (默认为 false)
   --rescan-interval VALUE  在此时间间隔内自动重新扫描源目录，若距上一次成功扫描时长超过此时间则进行新的扫描 (默认禁用)

   Storj 选项

   --storj-access-grant VALUE       Access Grant。[$STORJ_ACCESS_GRANT]
   --storj-api-key VALUE            API 密钥。[$STORJ_API_KEY]
   --storj-passphrase VALUE         加密密码短语。[$STORJ_PASSPHRASE]
   --storj-provider VALUE           选择一种身份验证方法。 (默认为 "existing") [$STORJ_PROVIDER]
   --storj-satellite-address VALUE  卫星地址。 (默认为 "us1.storj.io") [$STORJ_SATELLITE_ADDRESS]

```
{% endcode %}