# 企业文件系统

{% code fullWidth="true" %}
```
命令名:
   singularity storage update filefabric - 企业文件系统

使用方法:
   singularity storage update filefabric [command options] <name|id>

说明:
   --url
      要连接的企业文件系统的URL。
      
      示例:
         | https://storagemadeeasy.com       | Storage Made Easy US（美国）
         | https://eu.storagemadeeasy.com    | Storage Made Easy EU（欧洲）
         | https://yourfabric.smestorage.com | 连接到你的企业文件系统

   --root-folder-id
      根文件夹的ID。
      
      通常不需要填写。
      
      如果指定了该参数，rclone会以给定ID的目录作为初始目录。

   --permanent-token
      永久认证令牌。
      
      永久认证令牌可以在企业文件系统中创建，在用户仪表板的安全选项卡中存在一个名为"我的认证令牌"的条目。点击"管理"按钮来创建一个。
      
      这些令牌通常有效期为数年。
      
      更多信息请参阅: https://docs.storagemadeeasy.com/organisationcloud/api-tokens

   --token
      会话令牌。
      
      这是rclone缓存在配置文件中的会话令牌。通常有效期为1个小时。
      
      请勿设置该值 - rclone会自动设置。

   --token-expiry
      令牌的过期时间。
      
      请勿设置该值 - rclone会自动设置。

   --version
      文件系统的版本。
      
      请勿设置该值 - rclone会自动设置。

   --encoding
      后端的编码方式。
      
      更多信息请参阅[概述中的编码方式章节](/overview/#encoding)。


选项:
   --help, -h               显示帮助
   --permanent-token value  永久认证令牌。 [$PERMANENT_TOKEN]
   --root-folder-id value   根文件夹的ID。 [$ROOT_FOLDER_ID]
   --url value              要连接的企业文件系统的URL。 [$URL]

   高级选项:

   --encoding value      后端的编码方式。 (默认: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --token value         会话令牌。 [$TOKEN]
   --token-expiry value  令牌的过期时间。 [$TOKEN_EXPIRY]
   --version value       文件系统的版本。 [$VERSION]

```
{% endcode %}