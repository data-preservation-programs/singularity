# Zoho

{% code fullWidth="true" %}
```
命令名称：
   singularity storage create zoho - Zoho

使用方法：
   singularity storage create zoho [命令选项] [参数...]

描述：
   --client-id
      OAuth 客户端 ID。
   
      通常保留为空白。

   --client-secret
      OAuth 客户端秘钥。
   
      通常保留为空白。

   --token
      OAuth 访问令牌，使用 JSON 形式给出。

   --auth-url
      认证服务器 URL。
      
      保留为空白则使用提供者的默认配置。

   --token-url
      令牌服务器 URL。
      
      保留为空白则使用提供者的默认配置。

   --region
      要连接的 Zoho 区域。
      
      您需要使用您所属组织注册的区域。如果不确定，请使用与您在浏览器中连接的顶级域名相同的区域。

      例如：
         | com    | 美国/全球
         | eu     | 欧洲
         | in     | 印度
         | jp     | 日本
         | com.cn | 中国
         | com.au | 澳大利亚

   --encoding
      后端存储的编码。
      
      更多信息请参见[概述中的编码章节](/overview/#encoding)。


选项：
   --client-id value      OAuth 客户端 ID。[$CLIENT_ID]
   --client-secret value  OAuth 客户端秘钥。[$CLIENT_SECRET]
   --help, -h             显示帮助信息
   --region value         要连接的 Zoho 区域。[$REGION]

   进阶选项

   --auth-url value   认证服务器 URL。[$AUTH_URL]
   --encoding value   后端存储的编码。（默认值："Del,Ctl,InvalidUtf8"）[$ENCODING]
   --token value      OAuth 访问令牌，使用 JSON 形式给出。[$TOKEN]
   --token-url value  令牌服务器 URL。[$TOKEN_URL]

   常规选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径
```
{% endcode %}