# 盒子

{% code fullWidth="true" %}
```
命令：
   singularity datasource add box - 盒子

使用方法：
   singularity datasource add box [命令选项] <数据集名称> <源路径>

描述：
   --box-access-token
      盒子应用主访问令牌
      
      通常留空。

   --box-auth-url
      授权服务器URL。
      
      留空以使用提供者的默认值。

   --box-box-config-file
      盒子应用config.json文件位置
      
      通常留空。
      
      文件名中自动展开`~`，同时环境变量如`${RCLONE_CONFIG_DIR}`也会被展开。

   --box-box-sub-type
      

      示例：
         | user       | Rclone应以用户身份使用。
         | enterprise | Rclone应以服务帐户身份使用。

   --box-client-id
      OAuth客户端ID。
      
      通常留空。

   --box-client-secret
      OAuth客户端密钥。
      
      通常留空。

   --box-commit-retries
      尝试提交分段文件的最大次数。

   --box-encoding
      后端的编码。
      
      更多信息请参见[概述中的编码部分](/overview/#encoding)。

   --box-list-chunk
      列表块的大小，范围为1-1000。

   --box-owned-by
      仅显示由传递的登录（电子邮件地址）拥有的项目。

   --box-root-folder-id
      填写以使rclone将非根文件夹作为起点。

   --box-token
      OAuth访问令牌的JSON数据块。

   --box-token-url
      令牌服务器URL。
      
      留空以使用提供者的默认值。

   --box-upload-cutoff
      切换到分段上传的临界点（>= 50 MiB）。

选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 将数据集导出为CAR文件后删除数据集中的文件。（默认：false）
   --rescan-interval value  当从上次成功扫描后经过指定的时间间隔时，自动重新扫描源目录（默认：禁用）
   --scanning-state value   设置初始扫描状态（默认：ready）

   盒子选项

   --box-access-token value     盒子应用主访问令牌 [$BOX_ACCESS_TOKEN]
   --box-auth-url value         授权服务器URL [$BOX_AUTH_URL]
   --box-box-config-file value  盒子应用config.json文件位置 [$BOX_BOX_CONFIG_FILE]
   --box-box-sub-type value     （默认："user"）[$BOX_BOX_SUB_TYPE]
   --box-client-id value        OAuth客户端ID [$BOX_CLIENT_ID]
   --box-client-secret value    OAuth客户端密钥 [$BOX_CLIENT_SECRET]
   --box-commit-retries value   尝试提交分段文件的最大次数（默认："100"）[$BOX_COMMIT_RETRIES]
   --box-encoding value         后端的编码（默认："Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"）[$BOX_ENCODING]
   --box-list-chunk value       列表块的大小1-1000（默认："1000"）[$BOX_LIST_CHUNK]
   --box-owned-by value         仅显示由传递的登录（电子邮件地址）拥有的项目 [$BOX_OWNED_BY]
   --box-root-folder-id value   填写以使rclone将非根文件夹作为起点（默认："0"）[$BOX_ROOT_FOLDER_ID]
   --box-token value            OAuth访问令牌的JSON数据块 [$BOX_TOKEN]
   --box-token-url value        令牌服务器URL [$BOX_TOKEN_URL]
   --box-upload-cutoff value    切换到分段上传的临界点（>= 50 MiB）（默认："50Mi"）[$BOX_UPLOAD_CUTOFF]

```
{% endcode %}