# SMB / CIFS

{% code fullWidth="true" %}
```
命令名称:
   singularity storage update smb - SMB / CIFS

命令用法:
   singularity storage update smb [命令选项] <名称|ID>

命令描述:
   --host
      要连接的SMB服务器主机名。
      
      例如："example.com"。

   --user
      SMB用户名。

   --port
      SMB端口号。

   --pass
      SMB密码。

   --domain
      NTLM身份验证的域名。

   --spn
      服务主体名称。
      
      Rclone将此名称呈现给服务器。一些服务器将其用作进一步的身份验证，对于群集通常需要设置。例如：
      
          cifs/remotehost:1020
      
      如果不确定，请留空。
      

   --idle-timeout
      空闲连接关闭之前的最大时间。
      
      如果在给定的时间内没有将连接返回到连接池中，Rclone将清空连接池。
      
      设置为0以无限期保持连接。
      

   --hide-special-share
      隐藏用户无权访问的特殊共享（例如print$）。

   --case-insensitive
      服务器是否配置为大小写不敏感。
      
      在Windows共享中始终为true。

   --encoding
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。


选项:
   --domain value  NTLM身份验证的域名。（默认值：“WORKGROUP”）[$DOMAIN]
   --help, -h      显示帮助信息
   --host value    要连接的SMB服务器主机名。[$HOST]
   --pass value    SMB密码。[$PASS]
   --port value    SMB端口号。（默认值：445）[$PORT]
   --spn value     服务主体名称。[$SPN]
   --user value    SMB用户名。（默认值：“$USER”）[$USER]

   高级选项

   --case-insensitive    服务器是否配置为大小写不敏感。（默认值：true）[$CASE_INSENSITIVE]
   --encoding value      后端的编码。（默认值：“Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot”）[$ENCODING]
   --hide-special-share  隐藏用户无权访问的特殊共享（例如print$）。（默认值：true）[$HIDE_SPECIAL_SHARE]
   --idle-timeout value  空闲连接关闭之前的最大时间。（默认值：“1m0s”）[$IDLE_TIMEOUT]

```
{% endcode %}