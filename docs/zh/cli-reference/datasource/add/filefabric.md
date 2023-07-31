# 企业文件搬运机

{% code fullWidth="true" %}
```
名称:
   singularity datasource add filefabric - 企业文件搬运机

使用方法:
   singularity datasource add filefabric [命令选项] <数据集名称> <源路径>

描述:
   --filefabric-encoding
      后端的编码。
      
      请参阅[概述中的编码部分](/overview/#encoding)获取更多信息。

   --filefabric-permanent-token
      永久认证令牌。
      
      永久认证令牌可以在企业文件搬运机中创建，可以在用户仪表盘的安全性下找到一个名为“我的认证令牌”的条目，点击管理按钮进行创建。
      
      这些令牌通常有效期为数年。
      
      更多信息请参阅: https://docs.storagemadeeasy.com/organisationcloud/api-tokens
      

   --filefabric-root-folder-id
      根文件夹的ID。
      
      通常保留为空。
      
      如果有特定的ID来指定文件夹，填入此处的ID，会使rclone以该文件夹为起始目录。
      

   --filefabric-token
      会话令牌。
      
      这是会话令牌，rclone会将其缓存在配置文件中。它通常有效期为1小时。
      
      无需设置此值 - rclone会自动设置。
      

   --filefabric-token-expiry
      令牌过期时间。
      
      无需设置此值 - rclone会自动设置。
      

   --filefabric-url
      企业文件搬运机的连接URL。

      示例:
         | https://storagemadeeasy.com       | Storage Made Easy US
         | https://eu.storagemadeeasy.com    | Storage Made Easy EU
         | https://yourfabric.smestorage.com | 连接到您的企业文件搬运机

   --filefabric-version
      从文件搬运机中读取的版本。
      
      无需设置此值 - rclone会自动设置。
      


选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险操作] 在将数据集导出到CAR文件后删除文件。(默认值：false)
   --rescan-interval value  当上一次成功扫描后，如果到了这个间隔时间，自动重新扫描源目录（默认值：禁用）
   --scanning-state value   设置初始扫描状态（默认值：ready）

   企业文件搬运机选项

   --filefabric-encoding value         后端的编码。(默认值："Slash,Del,Ctl,InvalidUtf8,Dot") [$FILEFABRIC_ENCODING]
   --filefabric-permanent-token value  永久认证令牌。[$FILEFABRIC_PERMANENT_TOKEN]
   --filefabric-root-folder-id value   根文件夹的ID。[$FILEFABRIC_ROOT_FOLDER_ID]
   --filefabric-token value            会话令牌。[$FILEFABRIC_TOKEN]
   --filefabric-token-expiry value     令牌过期时间。[$FILEFABRIC_TOKEN_EXPIRY]
   --filefabric-url value              企业文件搬运机的连接URL。[$FILEFABRIC_URL]
   --filefabric-version value          从文件搬运机中读取的版本。[$FILEFABRIC_VERSION]

```
{% endcode %}