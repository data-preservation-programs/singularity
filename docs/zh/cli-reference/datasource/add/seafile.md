# seafile

{% code fullWidth="true" %}
```
名称:
   singularity datasource add seafile - seafile

用法:
   singularity datasource add seafile [命令选项] <dataset_name> <source_path>

描述:
   --seafile-2fa
      双因素身份验证（如果帐户启用了2FA，则为 'true'）。

   --seafile-auth-token
      认证令牌。

   --seafile-create-library
      如果库不存在，是否创建一个库。

   --seafile-encoding
      后端的编码。
      
      有关详细信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --seafile-library
      库的名称。
      
      如果要访问所有非加密库，则留空。

   --seafile-library-key
      库密码（仅适用于加密库）。
      
      如果通过命令行传递密码，则留空。

   --seafile-pass
      密码。

   --seafile-url
      要连接的 seafile 主机的 URL。

      示例：
         | https://cloud.seafile.com/ | 连接到 cloud.seafile.com。

   --seafile-user
      用户名（通常是电子邮件地址）。

选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 将数据集导出为 CAR 文件后删除数据集文件。 (默认值: false)
   --rescan-interval value  当经过此时间间隔的最后一次成功扫描后，自动重新扫描源目录 (默认值: 禁用)
   --scanning-state value   设置初始扫描状态 (默认值: 就绪)

   seafile 选项

   --seafile-2fa value             双因素身份验证（如果帐户启用了2FA，则为 'true'）。 (默认值: "false") [$SEAFILE_2FA]
   --seafile-create-library value  如果库不存在，是否创建一个库。 (默认值: "false") [$SEAFILE_CREATE_LIBRARY]
   --seafile-encoding value        后端的编码。 (默认值: "Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8") [$SEAFILE_ENCODING]
   --seafile-library value         库的名称。 [$SEAFILE_LIBRARY]
   --seafile-library-key value     库密码（仅适用于加密库）。 [$SEAFILE_LIBRARY_KEY]
   --seafile-pass value            密码。 [$SEAFILE_PASS]
   --seafile-url value             要连接的 seafile 主机的 URL。 [$SEAFILE_URL]
   --seafile-user value            用户名（通常是电子邮件地址）。 [$SEAFILE_USER]
```
{% endcode %}