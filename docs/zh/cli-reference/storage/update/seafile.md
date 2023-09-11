# seafile

{% code fullWidth="true" %}
```
命令名称：
   singularity storage update seafile - Seafile
   
使用方法：
   singularity storage update seafile [命令选项] <名称|ID>

描述：
   --url
      要连接的 Seafile 主机的 URL。

      示例：
         | https://cloud.seafile.com/ | 连接到 cloud.seafile.com。

   --user
      用户名（通常为电子邮件地址）。

   --pass
      密码。

   --2fa
      双重身份验证（如果该账户已启用双重身份验证，则为 'true'）。

   --library
      文库的名称。
      
      如果要访问所有非加密的文库，则留空。

   --library-key
      用于加密文库的密码（仅适用于加密的文库）。
      
      如果您通过命令行传递密码，则留空。

   --create-library
      如果库不存在，是否应该创建一个。

   --auth-token
      认证令牌。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。


选项：
   --2fa                双重身份验证（如果该账户已启用双重身份验证，则为 'true'）。 （默认值：false）[$2FA]
   --auth-token 值       认证令牌。 [$AUTH_TOKEN]
   --help, -h           显示帮助
   --library 值          文库的名称。 [$LIBRARY]
   --library-key 值     用于加密文库的密码。 （仅适用于加密的文库）。 [$LIBRARY_KEY]
   --pass 值             密码。 [$PASS]
   --url 值              要连接的 Seafile 主机的 URL。 [$URL]
   --user 值            用户名（通常为电子邮件地址）。 [$USER]

   高级选项

   --create-library  如果库不存在，是否应该创建一个。 （默认值：false）[$CREATE_LIBRARY]
   --encoding 值      后端的编码方式。 （默认值："Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8"）[$ENCODING]

```
{% endcode %}