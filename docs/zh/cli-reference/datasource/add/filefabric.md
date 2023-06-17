# 企业文件云
{% code fullWidth="true" %}
```
命令名:
   singularity datasource add filefabric - 企业文件云

用法:
   singularity datasource add filefabric [命令选项] <数据集名称> <源路径>

说明:
   --filefabric-token-expiry
      Token 过期时间。
      
      请不要设置该值，rclone 将自动设置.
      

   --filefabric-version
      文件云版本。
      
      请不要设置该值，rclone 将自动设置。

   --filefabric-encoding
      后端编码。
      
      更多信息请查看[概述](/overview/#encoding)中 "encoding" 部分。

   --filefabric-url
      连接企业文件云的 URL.

      举例:
         | https://storagemadeeasy.com       | Storage Made Easy 美国
         | https://eu.storagemadeeasy.com    | Storage Made Easy 欧洲
         | https://yourfabric.smestorage.com | 连接到您的企业文件云

   --filefabric-root-folder-id
      根目录的 ID。
      
      通常留空。
      
      如需使用指定文件夹作为起始位置，请填写对应文件夹的 ID。
      

   --filefabric-permanent-token
      永久认证 Token.
      
      永久认证 Token 可以在企业文件云的用户 Dashboard 中的 Security 里找到 "My Authentication Tokens" 条目，点击 Manage 按钮即可创建。
      
      这些 Token 通常有效期几年。

      更多信息请参考: https://docs.storagemadeeasy.com/organisationcloud/api-tokens
      

   --filefabric-token
      会话 Token。
      
      这是 rclone 缓存在配置文件中的会话令牌。通常情况下有效期为 1 小时。
      
      请不要设置该值，rclone 将自动设置。
      


选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作] 导出数据集到 CAR 文件后删除数据集文件。  (默认: false)
   --rescan-interval value  当成功扫描后，自动按一定间隔重新扫描源目录 (默认: disabled)

   用于文件云的选项

   --filefabric-encoding value         后端编码。 (默认:"Slash,Del,Ctl,InvalidUtf8,Dot") [$FILEFABRIC_ENCODING]
   --filefabric-permanent-token value  永久认证 Token. [$FILEFABRIC_PERMANENT_TOKEN]
   --filefabric-root-folder-id value   根目录的 ID. [$FILEFABRIC_ROOT_FOLDER_ID]
   --filefabric-token value            会话 Token. [$FILEFABRIC_TOKEN]
   --filefabric-token-expiry value     Token 过期时间. [$FILEFABRIC_TOKEN_EXPIRY]
   --filefabric-url value              连接企业文件云的 URL. [$FILEFABRIC_URL]
   --filefabric-version value          文件云版本。[$FILEFABRIC_VERSION]

```
{% endcode %}