# Zoho

{% code fullWidth="true" %}
```
名称:
   singularity datasource add zoho - Zoho

用法:
   singularity datasource add zoho [命令选项] <数据集名称> <源路径>

描述:
   --zoho-token
      OAuth访问令牌作为JSON代码块。

   --zoho-auth-url
      认证服务器URL。
      
      留空以使用提供程序的默认设置。

   --zoho-token-url
      令牌服务器URL。
      
      留空以使用提供程序的默认设置。

   --zoho-region
      要连接到的Zoho区域。
      
      您必须使用您的组织所注册的区域。如果不确定，请使用与您在浏览器中连接到的同一顶级域相同的域。

      示例:
         | com    |美国/全球
         | eu     |欧洲
         | in     |印度
         | jp     |日本
         | com.cn |中国
         | com.au |澳大利亚

   --zoho-encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --zoho-client-id
      OAuth客户端ID。
      
      通常不填。

   --zoho-client-secret
      OAuth客户端密码。
      
      通常不填。


选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作]导出到CAR文件后删除数据集中的文件。(默认值: 假)
   --rescan-interval value  当已过去此间隔自上次成功扫描后，自动重新扫描源目录 (默认值: 禁用)

   Zoho选项

   --zoho-auth-url value       认证服务器URL。[$ZOHO_AUTH_URL]
   --zoho-client-id value      OAuth客户端ID。[$ZOHO_CLIENT_ID]
   --zoho-client-secret value  OAuth客户端密码。[$ZOHO_CLIENT_SECRET]
   --zoho-encoding value       后端的编码方式。(默认值: "Del,Ctl,InvalidUtf8") [$ZOHO_ENCODING]
   --zoho-region value         要连接到的Zoho区域。[$ZOHO_REGION]
   --zoho-token value          OAuth访问令牌作为JSON代码块。[$ZOHO_TOKEN]
   --zoho-token-url value      令牌服务器URL。[$ZOHO_TOKEN_URL]

```
{% endcode %}