# seafile

{% code fullWidth="true" %}
```
名称:
   singularity 数据源添加 seafile - seafile

用法:
   singularity datasource add seafile [命令选项] <数据集名称> <源路径>

说明:
   --seafile-library
      库的名称。
      
      如果是非加密的库可以留空。

   --seafile-create-library
      如果库不存在，是否创建。

   --seafile-encoding
      后端的编码。
      
      参见[概述中的编码部分](/overview/#encoding) 了解更多信息。

   --seafile-url
      Seafile主机地址。

      例如:
         | https://cloud.seafile.com/ | 连接到 cloud.seafile.com。

   --seafile-pass
      密码。

   --seafile-2fa
      两步验证 ('true' 如果账户启用了2FA).

   --seafile-library-key
      库密码 (仅适用于加密库)。
      
      如果是通过命令行传递密码可以留空。

   --seafile-auth-token
      认证令牌。

   --seafile-user
      用户名(通常是电子邮箱)。


选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    在将数据集导出到CAR文件后删除数据集中的文件 [危险操作]。  (默认: false)
   --rescan-interval value  当这个时间段自上次成功扫描后过去时，自动重新扫描源目录 (defaultValue: disabled)

   Seafile选项

   --seafile-2fa value             两步验证 ('true' 如果账户启用了2FA)。  (默认: "false") [$SEAFILE_2FA]
   --seafile-create-library value  如果库不存在，是否创建。 (默认: "false") [$SEAFILE_CREATE_LIBRARY]
   --seafile-encoding value        后端编码。 (默认: "Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8") [$SEAFILE_ENCODING]
   --seafile-library value         库的名称。 [$SEAFILE_LIBRARY]
   --seafile-library-key value     库密码 (仅适用于加密库)。 [$SEAFILE_LIBRARY_KEY]
   --seafile-pass value            密码。 [$SEAFILE_PASS]
   --seafile-url value             Seafile主机地址。 [$SEAFILE_URL]
   --seafile-user value            用户名(通常是电子邮箱)。 [$SEAFILE_USER]

```
{% endcode %}