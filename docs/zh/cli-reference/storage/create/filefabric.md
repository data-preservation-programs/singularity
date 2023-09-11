# 企业文件系统

{% code fullWidth="true" %}
```
命令:
   singularity storage create filefabric - 企业文件系统

用法:
   singularity storage create filefabric [命令选项] [参数...]

描述:
   --url
      连接到企业文件系统的 URL。

      示例:
         | https://storagemadeeasy.com       | Storage Made Easy US
         | https://eu.storagemadeeasy.com    | Storage Made Easy EU
         | https://yourfabric.smestorage.com | 连接到您的企业文件系统

   --root-folder-id
      根文件夹的 ID。
      
      通常留空。
      
      如果要让 rclone 启动时跳转到指定 ID 的目录，请填写具体 ID。
      

   --permanent-token
      永久认证令牌。
      
      永久认证令牌可以在企业文件系统的用户仪表盘下的安全性中创建。
      在那里，您会看到一个名为“我的认证令牌”的项目。
      单击“管理”按钮即可创建。
      
      这些令牌通常有效期几年。
      
      更多信息请参见：https://docs.storagemadeeasy.com/organisationcloud/api-tokens
      

   --token
      会话令牌。
      
      这是一个会话令牌，rclone 会缓存在配置文件中。
      通常有效期为 1 小时。
      
      不要设置此值 - rclone 会自动设置。
      

   --token-expiry
      令牌到期时间。
      
      不要设置此值 - rclone 会自动设置。
      

   --version
      从文件系统中读取的版本信息。
      
      不要设置此值 - rclone 会自动设置。
      

   --encoding
      后端的编码方式。
      
      有关详细信息，请参见[概述中的编码部分](/overview/#encoding)。


选项:
   --help, -h               显示帮助
   --permanent-token value  永久认证令牌。[$PERMANENT_TOKEN]
   --root-folder-id value   根文件夹的 ID。[$ROOT_FOLDER_ID]
   --url value              连接到企业文件系统的 URL。[$URL]

   高级选项

   --encoding value      后端的编码方式。 (默认值: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --token value         会话令牌。[$TOKEN]
   --token-expiry value  令牌到期时间。[$TOKEN_EXPIRY]
   --version value       从文件系统中读取的版本信息。[$VERSION]

   常规选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}