# Box

{% code fullWidth="true" %}
```
命令:
   singularity storage create box - Box

用法:
   singularity storage create box [命令选项] [参数...]

描述:
   --client-id
      OAuth客户端ID。
      
      通常留空。

   --client-secret
      OAuth客户端秘钥。
      
      通常留空。

   --token
      OAuth访问令牌，以JSON格式。
   
   --auth-url
      认证服务器URL。
      
      通常留空以使用默认提供商。

   --token-url
      令牌服务器URL。
      
      通常留空以使用默认提供商。

   --root-folder-id
      填写rclone要使用的非根目录作为起始点。

   --box-config-file
      Box App config.json文件位置。
      
      通常留空。
      
      `~`将被扩展为文件名，环境变量如`${RCLONE_CONFIG_DIR}`也将被扩展。

   --access-token
      Box App主要访问令牌。
      
      通常留空。

   --box-sub-type
      

      示例:
         | user       | Rclone应代表用户操作。
         | enterprise | Rclone应代表服务帐户操作。

   --upload-cutoff
      切换到多部分上传的截止值 (>= 50 MiB)。

   --commit-retries
      尝试提交多部分文件的最大次数。

   --list-chunk
      列出块的大小 1-1000。

   --owned-by
      仅显示由登录名（电子邮件地址）拥有的项目。

   --encoding
      后端的编码。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。


选项:
   --access-token value     Box App主要访问令牌 [$ACCESS_TOKEN]
   --box-config-file value  Box App config.json文件位置 [$BOX_CONFIG_FILE]
   --box-sub-type value     (默认: "user") [$BOX_SUB_TYPE]
   --client-id value        OAuth客户端ID [$CLIENT_ID]
   --client-secret value    OAuth客户端秘钥 [$CLIENT_SECRET]
   --help, -h               显示帮助

   高级选项

   --auth-url value        认证服务器URL [$AUTH_URL]
   --commit-retries value  尝试提交多部分文件的最大次数 (默认: 100) [$COMMIT_RETRIES]
   --encoding value        后端的编码 (默认: "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --list-chunk value      列出块的大小 1-1000 (默认: 1000) [$LIST_CHUNK]
   --owned-by value        仅显示由登录名（电子邮件地址）拥有的项目 [$OWNED_BY]
   --root-folder-id value  填写rclone要使用的非根目录作为起始点 (默认: "0") [$ROOT_FOLDER_ID]
   --token value           OAuth访问令牌，以JSON格式 [$TOKEN]
   --token-url value       令牌服务器URL [$TOKEN_URL]
   --upload-cutoff value   切换到多部分上传的截止值 (>= 50 MiB) (默认: "50Mi") [$UPLOAD_CUTOFF]

   一般选项

   --name value  存储名称 (默认: 自动生成)
   --path value  存储路径

```
{% endcode %}