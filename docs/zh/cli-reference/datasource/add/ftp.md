# FTP

{% code fullWidth="true" %}
```
名称：
   singularity datasource add ftp - FTP

用法：
   singularity datasource add ftp [命令选项] <数据集名称> <来源路径>

说明：
   --ftp-tls-cache-size
      所有控制和数据连接的TLS会话缓存大小。
      
      TLS缓存允许恢复TLS会话和在连接之间复用PSK。
      如果默认大小不够，会导致TLS恢复错误，则应增加此参数。
      默认情况下启用。使用0来禁用。

   --ftp-disable-tls13
      禁用TLS 1.3（解决具有错误TLS的FTP服务器）

   --ftp-port
      FTP端口号。

   --ftp-tls
      使用隐含的FTPS（FTP over TLS）。
      
      当使用隐式的FTP over TLS时，客户端会从一开始就使用TLS，无法兼容
      不支持TLS的服务器。通常通过端口990提供，而不是端口21。
      不能与显式FTPS组合使用。

   --ftp-disable-epsv
      禁用EPSV，即使服务器宣传支持也是如此。

   --ftp-disable-utf8
      禁用UTF-8，即使服务器宣传支持也是如此。

   --ftp-writing-mdtm
      使用MDTM设置修改时间（VsFtpd习惯用法）

   --ftp-force-list-hidden
      通过使用LIST -a来强制列出隐藏文件和文件夹。这将禁用MLSD。

   --ftp-encoding
      后端的编码。
      
      请参阅[概述中的编码部分](/overview/#encoding)获取更多信息。

      示例：
         | Asterisk,Ctl,Dot,Slash                               | ProFTPd无法处理带*的文件名
         | BackSlash,Ctl,Del,Dot,RightSpace,Slash,SquareBracket | PureFTPd无法处理 "[]" 或带*的文件名
         | Ctl,LeftPeriod,Slash                                 | VsFTPd无法处理以点开头的文件名

   --ftp-host
      要连接到的FTP主机。
      
      例如：“ftp.example.com”。

   --ftp-user
      FTP用户名。

   --ftp-no-check-certificate
      不验证服务器的TLS证书。

   --ftp-disable-mlsd
      禁用MLSD，即使服务器宣传支持也是如此。

   --ftp-idle-timeout
      关闭闲置连接之前的最大时间。
      
      如果在给定的时间内没有将连接返回到连接池中，rclone将清空连接池。
      
      设置为0以无限期保留连接。
      

   --ftp-close-timeout
      等待关闭的响应的最大时间。

   --ftp-shut-timeout
      等待数据连接关闭状态的最长时间。

   --ftp-ask-password
      允许在需要密码时要求输入FTP密码。
      
      如果设置了此选项且没有提供密码，则rclone将要求输入密码
      

   --ftp-pass
      FTP密码。

   --ftp-explicit-tls
      使用显式的FTPS（FTP over TLS）。
      
      当使用显式FTP over TLS时，客户机明确请求
      从一个纯文本连接升级到一个加密连接来安全地请求服务器。
      不能与隐式FTPS组合使用。

   --ftp-concurrency
      FTP同时连接的最大数量，0表示无限制。
      
      请注意，设置此值很可能会导致死锁，因此应谨慎使用。
      
      如果您正在进行同步或复制操作，请确保并发性比“ --transfers”和“ --checkers”的总和多1。
      
      如果您使用“ --check-first”，则它只需要比“ --checkers”和“ --transfers”的最大值多1。
      
      因此，对于“并发性3”，您将使用“ --checkers 2 --transfers 2 --check-first”或“--checkers 1 --transfers 1”。

OPTIONS：
   --help，-h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 将数据集导出至CAR文件后，删除数据集文件。  (default: false)
   --rescan-interval value  当距离上次成功扫描的间隔时间已过时，自动重新扫描源目录（默认：禁用）

   FTP选项

   --ftp-ask-password value          允许在需要密码时要求输入FTP密码。 (default: "false") [$FTP_ASK_PASSWORD]
   --ftp-close-timeout value         等待关闭的响应的最大时间。 (default: "1m0s") [$FTP_CLOSE_TIMEOUT]
   --ftp-concurrency value           FTP同时连接的最大数量，0表示无限制。 (default: "0") [$FTP_CONCURRENCY]
   --ftp-disable-epsv value          禁用EPSV，即使服务器宣传支持也是如此。 (default: "false") [$FTP_DISABLE_EPSV]
   --ftp-disable-mlsd value          禁用MLSD，即使服务器宣传支持也是如此。 (default: "false") [$FTP_DISABLE_MLSD]
   --ftp-disable-tls13 value         禁用TLS 1.3（解决具有错误TLS的FTP服务器） (default: "false") [$FTP_DISABLE_TLS13]
   --ftp-disable-utf8 value          禁用UTF-8，即使服务器宣传支持也是如此。 (default: "false") [$FTP_DISABLE_UTF8]
   --ftp-encoding value              后端的编码。 (default: "Slash,Del,Ctl,RightSpace,Dot") [$FTP_ENCODING]
   --ftp-explicit-tls value          使用显式的FTPS（FTP over TLS）。 (default: "false") [$FTP_EXPLICIT_TLS]
   --ftp-force-list-hidden value     通过使用LIST -a来强制列出隐藏文件和文件夹。这将禁用MLSD。 (default: "false") [$FTP_FORCE_LIST_HIDDEN]
   --ftp-host value                  要连接到的FTP主机。 [$FTP_HOST]
   --ftp-idle-timeout value          关闭闲置连接之前的最大时间。 (default: "1m0s") [$FTP_IDLE_TIMEOUT]
   --ftp-no-check-certificate value  不验证服务器的TLS证书。 (default: "false") [$FTP_NO_CHECK_CERTIFICATE]
   --ftp-pass value                  FTP密码。 [$FTP_PASS]
   --ftp-port value                  FTP端口号。 (default: "21") [$FTP_PORT]
   --ftp-shut-timeout value          等待数据连接关闭状态的最大时间。 (default: "1m0s") [$FTP_SHUT_TIMEOUT]
   --ftp-tls value                   使用隐含的FTPS（FTP over TLS）。 (default: "false") [$FTP_TLS]
   --ftp-tls-cache-size value        所有控制和数据连接的TLS会话缓存大小。 (default: "32") [$FTP_TLS_CACHE_SIZE]
   --ftp-user value                  FTP用户名。 (default: "shane") [$FTP_USER]
   --ftp-writing-mdtm value          使用MDTM设置修改时间（VsFtpd习惯用法） (default: "false") [$FTP_WRITING_MDTM]

```
{% endcode %}