# SMB / CIFS

{% code fullWidth="true" %}
```
名称:
   singularity存储 创建 [smb] - SMB / CIFS

用法:
   singularity存储 创建 smb [命令选项] [参数...]

描述:
   --主机
      连接的SMB服务器主机名。
      
      例如："example.com"。

   --用户
      SMB用户名。

   --端口
      SMB端口号。

   --密码
      SMB密码。

   --域
      NTLM认证的域名。

   --SPN
      服务主体名称。
      
      Rclone将此名称呈现给服务器。某些服务器使用此作为进一步的认证，并且通常需要设置为群集。例如：
      
          cifs/remotehost:1020
      
      如果不确定，请留空。
      

   --空闲超时
      关闭空闲连接之前的最长时间。
      
      如果在给定时间内没有将连接返回到连接池中，Rclone将清空连接池。
      
      设置为0以无限期保持连接。
      

   --隐藏特殊共享
      隐藏用户不能访问的特殊共享（例如print$）。

   --不区分大小写
      服务器是否配置为不区分大小写。
      
      Windows共享始终为true。

   --编码
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。


选项:
   --域值    NTLM认证的域名。 (默认值: "WORKGROUP") [$域]
   --help, -h             显示帮助
   --主机值    连接的SMB服务器主机名。 [$主机]
   --密码值    SMB密码。 [$密码]
   --端口值    SMB端口号。 (默认值: 445) [$端口]
   --SPN值     服务主体名称. [$SPN]
   --用户值    SMB用户名. (默认值: "$USER") [$USER]

   高级选项

   --不区分大小写             服务器是否配置为不区分大小写。 (默认值: true) [$CASE_INSENSITIVE]
   --编码值                  后端的编码。 (默认值: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --隐藏特殊共享             隐藏用户不能访问的特殊共享（例如print$）。 (默认值: true) [$HIDE_SPECIAL_SHARE]
   --空闲超时值  关闭空闲连接之前的最长时间。 (默认值: "1m0s") [$IDLE_TIMEOUT]

   常规选项

   --名称值    存储的名称 (默认值: 自动生成的)
   --路径值    存储的路径

```
{% endcode %}