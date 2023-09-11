# seafile

{% code fullWidth="true" %}
```
名称：
   singularity storage create seafile - seafile

用法：
   singularity storage create seafile [命令选项] [参数...]

描述：
   --url
      要连接的 seafile 主机的 URL。

      示例：
         | https://cloud.seafile.com/ | 连接到 cloud.seafile.com。

   --user
      用户名（通常是电子邮箱地址）。

   --pass
      密码。

   --2fa
      两步验证（如果帐户启用了 2FA，则为 'true'）。

   --library
      库的名称。
      
      留空以访问所有未加密的库。

   --library-key
      库的密码（仅适用于加密的库）。
      
      如果通过命令行传递，则留空。

   --create-library
      如果库不存在，是否要创建一个库。

   --auth-token
      身份验证令牌。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

选项：
   --2fa                两步验证（如果帐户启用了 2FA，则为 'true'）。 (默认值: false) [$2FA]
   --auth-token value   身份验证令牌。 [$AUTH_TOKEN]
   --help, -h           显示帮助信息
   --library value      库的名称。 [$LIBRARY]
   --library-key value  库的密码（仅适用于加密的库）。 [$LIBRARY_KEY]
   --pass value         密码。 [$PASS]
   --url value          要连接的 seafile 主机的 URL。 [$URL]
   --user value         用户名（通常是电子邮箱地址）。 [$USER]

   高级选项

   --create-library  如果库不存在，rclone 是否应创建一个库。 (默认值: false) [$CREATE_LIBRARY]
   --encoding value  后端的编码方式。 (默认值: "Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8") [$ENCODING]

   常规选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}