# Box

{% code fullWidth="true" %}
```
名称：
   singularity datasource add box - Box

用法:
   singularity datasource add box [命令选项] <数据集名称> <源路径>
   
描述:
   --box-client-id
      OAuth客户端ID。
   
      通常留空。

   --box-auth-url
      认证服务器URL。
   
      留空以使用提供程序默认值。

   --box-box-config-file
      Box应用程序`config.json`位置。

      通常留空。

      前导`~`将扩展文件名，环境变量，如`${RCLONE_CONFIG_DIR}`也是。
      
   --box-encoding
      后端的编码方式。

      有关更多信息，请参见[总览](/overview/#encoding)中的编码部分。

   --box-owned-by
      仅显示由传入的登录（电子邮件地址）拥有的项。

   --box-upload-cutoff
      切换到分块上传的截止值（> = 50 MB）。

   --box-list-chunk
      显示清单块大小(1-1000)。

   --box-client-secret
      OAuth客户端秘钥。

      通常留空。

   --box-token
      OAuth访问令牌的JSON blob。

   --box-token-url
      令牌服务器的URL。

      留空以使用提供程序默认值。

   --box-root-folder-id
      填写以便rclone使用非根文件夹作为其起始点。

   --box-access-token
      Box应用程序主访问令牌。

      通常留空。

   --box-box-sub-type
      

      示例:
         |用户| Rclone应该代表一个用户。
         |企业| Rclone应该代表服务帐户。

   --box-commit-retries
      尝试提交多部分文件的最大次数。


选项：
   --help, -h  显示帮助
   
   数据准备选项

   --delete-after-export [危险] 导出到CAR文件后，删除数据集的文件。 （默认值：false）
   --rescan-interval value 当上一次扫描成功后经过该间隔时，自动重新扫描源目录（默认禁用）

   box选项

   --box-access-token value Box应用程序主访问令牌[$BOX_ACCESS_TOKEN]
   --box-auth-url value 认证服务器URL[$BOX_AUTH_URL]
   --box-box-config-file value Box应用程序`config.json`位置[$BOX_BOX_CONFIG_FILE]
   --box-box-sub-type value（默认值："user"）[$BOX_BOX_SUB_TYPE]
   --box-client-id value OAuth客户端ID[$BOX_CLIENT_ID]
   --box-client-secret value OAuth客户端秘钥[$BOX_CLIENT_SECRET]
   --box-commit-retries value 尝试提交多部分文件的最大次数。 （默认： "100"）[$BOX_COMMIT_RETRIES]
   --box-encoding value 后端的编码方式。(默认值："Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot")[$BOX_ENCODING]
   --box-list-chunk value 显示清单块大小(1-1000)。 （默认值： "1000"）[$BOX_LIST_CHUNK]
   --box-owned-by value 仅显示由传入的登录（电子邮件地址）拥有的项。[$BOX_OWNED_BY]
   --box-root-folder-id value 填写以便rclone使用非根文件夹作为其起始点。（默认值："0"）[$BOX_ROOT_FOLDER_ID]
   --box-token value OAuth访问令牌的JSON blob。[$BOX_TOKEN]
   --box-token-url value 令牌服务器的URL。[$BOX_TOKEN_URL]
   --box-upload-cutoff value 切换到分块上传的截止值（> = 50MB）。 （默认值： "50Mi"）[$BOX_UPLOAD_CUTOFF]

```
{% endcode %}