# Zoho

{% code fullWidth="true" %}
```
命令:
   singularity storage update zoho - Zoho

用法:
   singularity storage update zoho [命令选项] <名称|ID>

描述:
   --client-id <客户端ID>
      OAuth客户端ID。
      
      通常为空白。

   --client-secret <客户端秘钥>
      OAuth客户端秘钥。
      
      通常为空白。

   --token <访问令牌>
      以JSON格式的OAuth访问令牌。

   --auth-url <认证服务器URL>
      认证服务器URL。
      
      若为空白，则使用默认提供者。

   --token-url <令牌服务器URL>
      令牌服务器URL。
      
      若为空白，则使用默认提供者。

   --region <Zoho区域>
      要连接的Zoho区域。
      
      必须使用组织注册的区域。如果不确定，请使用与浏览器中连接的顶级域名相同的区域。

      例如:
         | com    | 美国 / 全球
         | eu     | 欧洲
         | in     | 印度
         | jp     | 日本
         | com.cn | 中国
         | com.au | 澳大利亚

   --encoding <后端编码>
      后端的编码格式。
      
      更多信息请参见[概述中的编码部分](/overview/#encoding)。


选项:
   --client-id value      OAuth客户端ID。[$CLIENT_ID]
   --client-secret value  OAuth客户端秘钥。[$CLIENT_SECRET]
   --help, -h             显示帮助
   --region value         要连接的Zoho区域。[$REGION]

   高级选项:

   --auth-url value   认证服务器URL。[$AUTH_URL]
   --encoding value   后端编码格式。（默认值："Del,Ctl,InvalidUtf8"）[$ENCODING]
   --token value      以JSON格式的OAuth访问令牌。[$TOKEN]
   --token-url value  令牌服务器URL。[$TOKEN_URL]

```
{% endcode %}