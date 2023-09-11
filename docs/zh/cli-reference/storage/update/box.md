# Box

{% code fullWidth="true" %}
```
名称:
   singularity storage update box - Box

用法:
   singularity storage update box [命令选项] <名称|ID>

描述:
   --client-id
      OAuth Client Id.
      
      正常情况下留空。

   --client-secret
      OAuth Client Secret.
      
      正常情况下留空。

   --token
      OAuth Access Token 作为一个 JSON 对象。

   --auth-url
      认证服务器 URL。
      
      正常情况下留空以使用提供者默认。

   --token-url
      令牌服务器 URL。
      
      正常情况下留空以使用提供者默认。

   --root-folder-id
      填写以使 rclone 使用非根文件夹作为其起始点。

   --box-config-file
      Box App config.json 位置
      
      正常情况下留空。
      
      以及环境变量或以前备份的`${RCLONE_CONFIG_DIR}`等将在文件名中扩展。

   --access-token
      Box App 主要访问令牌
      
      正常情况下留空。

   --box-sub-type
      

      示例:
         | user       | Rclone 应以用户的身份操作。
         | enterprise | Rclone 应以服务帐号的身份操作。

   --upload-cutoff
      切换到分块上传的截止值（>= 50 MiB）。

   --commit-retries
      尝试提交分块文件的最大次数。

   --list-chunk
      列出块的大小 1-1000。

   --owned-by
      只显示由传入的登录（电子邮件地址）拥有的项目。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概览中的 encoding 部分](/overview/#encoding)。


选项:
   --access-token value     Box App 主要访问令牌 [$ACCESS_TOKEN]
   --box-config-file value  Box App config.json 位置 [$BOX_CONFIG_FILE]
   --box-sub-type value     （默认值："user"） [$BOX_SUB_TYPE]
   --client-id value        OAuth Client Id。 [$CLIENT_ID]
   --client-secret value    OAuth Client Secret。 [$CLIENT_SECRET]
   --help, -h               显示帮助

   高级选项

   --auth-url value        认证服务器 URL。[$AUTH_URL]
   --commit-retries value  尝试提交分块文件的最大次数。（默认值：100）[$COMMIT_RETRIES]
   --encoding value        后端的编码方式。（默认值："Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"）[$ENCODING]
   --list-chunk value      列出块的大小 1-1000。（默认值：1000）[$LIST_CHUNK]
   --owned-by value        只显示由传入的登录（电子邮件地址）拥有的项目。[$OWNED_BY]
   --root-folder-id value  填写以使 rclone 使用非根文件夹作为其起始点。（默认值："0"）[$ROOT_FOLDER_ID]
   --token value           OAuth Access Token 作为一个 JSON 对象。[$TOKEN]
   --token-url value       令牌服务器 URL。[$TOKEN_URL]
   --upload-cutoff value   切换到分块上传的截止值（>= 50 MiB）。（默认值："50Mi"）[$UPLOAD_CUTOFF]

```
{% endcode %}