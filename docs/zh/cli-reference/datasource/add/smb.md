# SMB / CIFS

{% code fullWidth="true" %}
```
名称:
   singularity datasource add smb - SMB / CIFS

用法:
   singularity datasource add smb [命令选项] <数据集名称> <源路径>

说明:
   --smb-case-insensitive
      服务器是否被配置为不区分大小写。
      
      在Windows共享上始终为真。
      
   --smb-domain
      NTLM认证的域名。
      
   --smb-encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --smb-hide-special-share
      隐藏用户不能访问的特殊共享（例如：print$）。

   --smb-host
      要连接的SMB服务器主机名。
      
      例如："example.com"。

   --smb-idle-timeout
      闲置连接关闭之前的最长时间。
      
      如果在给定的时间内未返回连接到连接池中，rclone将清空连接池。
      
      将其设置为0以无限期保持连接。

   --smb-pass
      SMB密码。

   --smb-port
      SMB端口号。

   --smb-spn
      服务主体名称。
      
      Rclone将此名称呈现给服务器。某些服务器会将其用作进一步的身份验证，并且对于集群来说，通常需要设置。例如：
      
          cifs/remotehost:1020
      
      如果不确定，请保留为空。

   --smb-user
      SMB用户名。


选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作] 在将数据集导出为CAR文件后删除数据集的文件。（默认值：false）
   --rescan-interval value  当距离上次成功扫描已经过去该时间间隔时，自动重新扫描源目录（默认值：禁用）
   --scanning-state value   设置初始的扫描状态（默认值：准备好）

   smb选项

   --smb-case-insensitive value    服务器是否被配置为不区分大小写。（默认值："true"）[$SMB_CASE_INSENSITIVE]
   --smb-domain value              NTLM认证的域名。（默认值："WORKGROUP"）[$SMB_DOMAIN]
   --smb-encoding value            后端的编码方式。（默认值："Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot"）[$SMB_ENCODING]
   --smb-hide-special-share value  隐藏用户不能访问的特殊共享（例如：print$）。（默认值："true"）[$SMB_HIDE_SPECIAL_SHARE]
   --smb-host value                要连接的SMB服务器主机名。[$SMB_HOST]
   --smb-idle-timeout value        闲置连接关闭之前的最长时间。（默认值："1m0s"）[$SMB_IDLE_TIMEOUT]
   --smb-pass value                SMB密码。[$SMB_PASS]
   --smb-port value                SMB端口号。（默认值："445"）[$SMB_PORT]
   --smb-spn value                 服务主体名称。[$SMB_SPN]
   --smb-user value                SMB用户名。（默认值："$USER"）[$SMB_USER]
```
{% endcode %}