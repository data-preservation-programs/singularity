# FTP

{% code fullWidth="true" %}
```
命令名：
   singularity datasource add ftp - FTP

用法：
   singularity datasource add ftp [命令选项] <数据集名称> <源路径>

描述：
   --ftp-ask-password
      当需要时允许询问FTP密码。
      
      如果设置了该选项且未提供密码，则rclone将要求输入密码。
      

   --ftp-close-timeout
      关闭连接等待响应的最长时间。

   --ftp-concurrency
      最大FTP并发连接数，0表示无限制。
      
      请注意，设置此选项很可能导致死锁，因此必须小心使用。
      
      如果进行同步或复制操作，请确保并发连接数比`--transfers`和`--checkers`的总和多1。
      
      如果使用`--check-first`选项，则只需要比`--checkers`和`--transfers`的最大值多1个。
      
      因此，对于`concurrency 3`，您可以使用`--checkers 2 --transfers 2 --check-first`或`--checkers 1 --transfers 1`。

   --ftp-disable-epsv
      禁用EPSV，即使服务器支持。

   --ftp-disable-mlsd
      禁用MLSD，即使服务器支持。

   --ftp-disable-tls13
      禁用TLS 1.3（用于修复具有错误的TLS的FTP服务器）。

   --ftp-disable-utf8
      禁用UTF-8，即使服务器支持。

   --ftp-encoding
      后端的编码方式。
      
      有关详细信息，请参见[概述中的编码章节](/overview/#encoding)。

      示例：
         | Asterisk,Ctl,Dot,Slash                               | ProFTPd无法处理文件名中的'*'
         | BackSlash,Ctl,Del,Dot,RightSpace,Slash,SquareBracket | PureFTPd无法处理文件名中的'[]'或'*'
         | Ctl,LeftPeriod,Slash                                 | VsFTPd无法处理以点开头的文件名

   --ftp-explicit-tls
      使用显式FTP over TLS (FTPS)。
      
      在使用显式FTP over TLS时，客户端明确请求服务器的安全性，以将明文连接升级为加密连接。不能与隐式FTPS同时使用。

   --ftp-force-list-hidden
      使用LIST -a来强制列出隐藏文件和文件夹，这将禁用使用MLSD。

   --ftp-host
      FTP主机名。
      
      例如："ftp.example.com"。

   --ftp-idle-timeout
      闲置连接关闭前的最大时间。
      
      如果在给定的时间内没有将连接返回到连接池中，rclone将清空连接池。
      
      设置为0以保持连接持续。

   --ftp-no-check-certificate
      不验证服务器的TLS证书。

   --ftp-pass
      FTP密码。

   --ftp-port
      FTP端口号。

   --ftp-shut-timeout
      关闭数据连接状态的最长等待时间。

   --ftp-tls
      使用隐式FTP over TLS (FTPS)。
      
      在使用隐式FTP over TLS时，客户端从一开始就使用TLS连接，这会导致无法与非TLS感知的服务器兼容。通常使用的端口号为990，而不是21。不能与显式FTPS同时使用。

   --ftp-tls-cache-size
      所有控制和数据连接的TLS会话缓存大小。
      
      TLS缓存允许恢复TLS会话并在连接之间重用PSK。如果默认大小不足以满足需求，会导致TLS恢复出错，请增加该值。默认情况下启用。使用0表示禁用。

   --ftp-user
      FTP用户名。

   --ftp-writing-mdtm
      使用MDTM设置修改时间（VsFtpd的怪癖）


选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 在将数据集导出到CAR文件后删除数据集的文件。  (默认值：false)
   --rescan-interval value  当距离上次成功扫描已过去指定时间间隔时，自动重新扫描源目录（默认值：禁用）
   --scanning-state value   设置初始扫描状态（默认值：ready）

   FTP选项

   --ftp-ask-password value          当需要时允许询问FTP密码（默认值："false"）[$FTP_ASK_PASSWORD]
   --ftp-close-timeout value         关闭连接等待响应的最长时间（默认值："1m0s"）[$FTP_CLOSE_TIMEOUT]
   --ftp-concurrency value           最大FTP并发连接数，0表示无限制（默认值："0"）[$FTP_CONCURRENCY]
   --ftp-disable-epsv value          禁用EPSV，即使服务器支持（默认值："false"）[$FTP_DISABLE_EPSV]
   --ftp-disable-mlsd value          禁用MLSD，即使服务器支持（默认值："false"）[$FTP_DISABLE_MLSD]
   --ftp-disable-tls13 value         禁用TLS 1.3（用于修复具有错误的TLS的FTP服务器）（默认值："false"）[$FTP_DISABLE_TLS13]
   --ftp-disable-utf8 value          禁用UTF-8，即使服务器支持（默认值："false"）[$FTP_DISABLE_UTF8]
   --ftp-encoding value              后端的编码方式（默认值："Slash,Del,Ctl,RightSpace,Dot"）[$FTP_ENCODING]
   --ftp-explicit-tls value          使用显式FTP over TLS (FTPS)（默认值："false"）[$FTP_EXPLICIT_TLS]
   --ftp-force-list-hidden value     使用LIST -a来强制列出隐藏文件和文件夹，这将禁用使用MLSD（默认值："false"）[$FTP_FORCE_LIST_HIDDEN]
   --ftp-host value                  FTP主机名 [$FTP_HOST]
   --ftp-idle-timeout value          闲置连接关闭前的最大时间（默认值："1m0s"）[$FTP_IDLE_TIMEOUT]
   --ftp-no-check-certificate value  不验证服务器的TLS证书（默认值："false"）[$FTP_NO_CHECK_CERTIFICATE]
   --ftp-pass value                  FTP密码 [$FTP_PASS]
   --ftp-port value                  FTP端口号（默认值："21"）[$FTP_PORT]
   --ftp-shut-timeout value          关闭数据连接状态的最长等待时间（默认值："1m0s"）[$FTP_SHUT_TIMEOUT]
   --ftp-tls value                   使用隐式FTP over TLS (FTPS)（默认值："false"）[$FTP_TLS]
   --ftp-tls-cache-size value        所有控制和数据连接的TLS会话缓存大小（默认值："32"）[$FTP_TLS_CACHE_SIZE]
   --ftp-user value                  FTP用户名（默认值："$USER"）[$FTP_USER]
   --ftp-writing-mdtm value          使用MDTM设置修改时间（VsFtpd的怪癖）（默认值："false"）[$FTP_WRITING_MDTM]

```
{% endcode %}