# Zoho

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add zoho - Zoho

USAGE:
   singularity datasource add zoho [command options] <dataset_name> <source_path>

DESCRIPTION:
   --zoho-auth-url
      认证服务器URL。
      
      留空以使用提供程序的默认设置。

   --zoho-client-id
      OAuth客户端ID。
      
      通常留空。

   --zoho-client-secret
      OAuth客户端密钥。
      
      通常留空。

   --zoho-encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --zoho-region
      要连接的Zoho地区。
      
      您必须使用您所在组织注册的地区。如果不确定，请使用与您在浏览器中连接的顶级域名相同的地区。

      示例：
         | com    | 美国 / 全球
         | eu     | 欧洲
         | in     | 印度
         | jp     | 日本
         | com.cn | 中国
         | com.au | 澳大利亚

   --zoho-token
      OAuth访问令牌（以JSON格式）。

   --zoho-token-url
      令牌服务器URL。
      
      留空以使用提供程序的默认设置。


OPTIONS:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 将数据集导出为CAR文件后删除数据集的文件。（默认: false）
   --rescan-interval value  如果距离上次成功扫描已过去该间隔，则自动重新扫描源目录（默认: 禁用）
   --scanning-state value   设置初始的扫描状态（默认: 准备就绪）

   Zoho选项

   --zoho-auth-url value       认证服务器URL。[$ZOHO_AUTH_URL]
   --zoho-client-id value      OAuth客户端ID。[$ZOHO_CLIENT_ID]
   --zoho-client-secret value  OAuth客户端密钥。[$ZOHO_CLIENT_SECRET]
   --zoho-encoding value       后端的编码方式。（默认: "Del,Ctl,InvalidUtf8"）[$ZOHO_ENCODING]
   --zoho-region value         要连接的Zoho地区。[$ZOHO_REGION]
   --zoho-token value          OAuth访问令牌（以JSON格式）。[$ZOHO_TOKEN]
   --zoho-token-url value      令牌服务器URL。[$ZOHO_TOKEN_URL]

```
{% endcode %}