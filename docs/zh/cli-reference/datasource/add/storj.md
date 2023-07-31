# Storj 分散式云存储

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add storj - Storj 分散式云存储

USAGE:
   singularity datasource add storj [command options] <dataset_name> <source_path>

DESCRIPTION:
   --storj-access-grant
      [供应商] - 现有
         访问凭证。

   --storj-api-key
      [供应商] - 新建
         API 密钥。

   --storj-passphrase
      [供应商] - 新建
         加密密语。
         
         要访问现有对象，请输入用于上传的密语。

   --storj-provider
      选择身份验证方法。

      示例:
         | existing | 使用现有访问凭证。
         | new      | 从卫星地址、API 密钥和密语创建新的访问凭证。

   --storj-satellite-address
      [供应商] - 新建
         卫星地址。
         
         自定义的卫星地址应匹配格式: `<nodeid>@<address>:<port>`。

         示例:
            | us1.storj.io | 美国1
            | eu1.storj.io | 欧洲1
            | ap1.storj.io | 亚太1


OPTIONS:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作] 将数据集导出为 CAR 文件后删除数据集中的文件。  (默认: false)
   --rescan-interval value  当最后一次成功扫描后的时间间隔超过此值时，自动重新扫描源目录 (默认: 禁用)
   --scanning-state value   设置初始扫描状态 (默认: ready)

   Storj 选项

   --storj-access-grant value       访问凭证。 [$STORJ_ACCESS_GRANT]
   --storj-api-key value            API 密钥。 [$STORJ_API_KEY]
   --storj-passphrase value         加密密语。 [$STORJ_PASSPHRASE]
   --storj-provider value           选择身份验证方法。 (默认: "existing") [$STORJ_PROVIDER]
   --storj-satellite-address value  卫星地址。 (默认: "us1.storj.io") [$STORJ_SATELLITE_ADDRESS]

```
{% endcode %}