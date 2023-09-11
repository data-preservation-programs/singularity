# FTP

{% code fullWidth="true" %}
```
名称:
   singularity storage update ftp - FTP

用法:
   singularity storage update ftp [命令选项] <名称|ID>

说明:
   --host
      要连接的FTP主机。
      
      例如："ftp.example.com"。

   --user
      FTP用户名。

   --port
      FTP端口号。

   --pass
      FTP密码。

   --tls
      使用隐式FTPS（FTP通过TLS）。
      
      使用隐式FTPS时，客户端从一开始就使用TLS进行连接，这会导致与不支持TLS的服务器的兼容性问题。
      通常在端口990上提供而不是端口21.  不能与显式FTPS结合使用。

   --explicit-tls
      使用显式FTPS（FTP通过TLS）。
      
      使用显式FTPS时，客户端显式请求服务器的安全性，以将明文连接升级为加密连接。
      不能与隐式FTPS结合使用。

   --concurrency
      FTP同时连接的最大数量，0表示无限制。
      
      请注意，设置这个值很有可能导致死锁，因此应小心使用。
      
      如果您正在执行同步或复制操作，请确保并发度比 `--transfers` 和 `--checkers` 的总和多1。
      
      如果使用了 `--check-first`，那么它只需要比 `--checkers` 和 `--transfers` 的最大值多1。
      
      因此，对于 `concurrency 3`，您可以使用 `--checkers 2 --transfers 2 --check-first` 或 `--checkers 1 --transfers 1`。

   --no-check-certificate
      不验证服务器的TLS证书。

   --disable-epsv
      禁用EPSV，即使服务器宣传支持。

   --disable-mlsd
      禁用MLSD，即使服务器宣传支持。

   --disable-utf8
      禁用UTF-8，即使服务器宣传支持。

   --writing-mdtm
      使用MDTM来设置修改时间（VsFtpd的怪癖）。

   --force-list-hidden
      使用LIST -a强制列出隐藏的文件和文件夹。这将禁用MLSD。

   --idle-timeout
      闲置连接关闭之前的最长时间。
      
      如果在给定的时间内没有将连接返回到连接池，则rclone将清空连接池。
      
      设置为0表示无限制。

   --close-timeout
      关闭连接等待响应的最长时间。

   --tls-cache-size
      所有控制和数据连接的TLS会话缓存的大小。
      
      TLS缓存允许恢复TLS会话并在连接之间复用PSK。
      如果默认大小不足以产生TLS恢复错误，请增加此值。
      默认情况下启用。使用0禁用。

   --disable-tls13
      禁用TLS 1.3（用于带有有缺陷TLS的FTP服务器的解决方法）。

   --shut-timeout
      等待数据连接关闭状态的最长时间。

   --ask-password
      在需要时允许询问FTP密码。
      
      如果设置了这个选项且没有提供密码，则rclone将要求输入密码。

   --encoding
      后端的编码。
      
      有关更多信息，请参阅[概览中的编码部分](/overview/#encoding)。

      例如:
         | Asterisk,Ctl,Dot,Slash                               | ProFTPd不能处理文件名中的'*'
         | BackSlash,Ctl,Del,Dot,RightSpace,Slash,SquareBracket | PureFTPd不能处理文件名中的'[]'或'*'
         | Ctl,LeftPeriod,Slash                                 | VsFTPd不能以'.'开头的文件名


选项:
   --explicit-tls  使用显式FTPS（FTP通过TLS）。 (默认: false) [$EXPLICIT_TLS]
   --help, -h      显示帮助
   --host value    要连接的FTP主机。 [$HOST]
   --pass value    FTP密码。 [$PASS]
   --port value    FTP端口号。 (默认: 21) [$PORT]
   --tls           使用隐式FTPS（FTP通过TLS）。 (默认: false) [$TLS]
   --user value    FTP用户名。 (默认: "$USER") [$USER]

   高级选项

   --ask-password          允许在需要时询问FTP密码。 (默认: false) [$ASK_PASSWORD]
   --close-timeout value   关闭连接等待响应的最长时间。 (默认: "1m0s") [$CLOSE_TIMEOUT]
   --concurrency value     FTP同时连接的最大数量，0表示无限制。 (默认: 0) [$CONCURRENCY]
   --disable-epsv          禁用EPSV，即使服务器宣传支持。 (默认: false) [$DISABLE_EPSV]
   --disable-mlsd          禁用MLSD，即使服务器宣传支持。 (默认: false) [$DISABLE_MLSD]
   --disable-tls13         禁用TLS 1.3（用于带有有缺陷TLS的FTP服务器的解决方法）。 (默认: false) [$DISABLE_TLS13]
   --disable-utf8          禁用UTF-8，即使服务器宣传支持。 (默认: false) [$DISABLE_UTF8]
   --encoding value        后端的编码。 (默认: "Slash,Del,Ctl,RightSpace,Dot") [$ENCODING]
   --force-list-hidden     使用LIST -a强制列出隐藏的文件和文件夹。这将禁用MLSD。 (默认: false) [$FORCE_LIST_HIDDEN]
   --idle-timeout value    闲置连接关闭之前的最长时间。 (默认: "1m0s") [$IDLE_TIMEOUT]
   --no-check-certificate  不验证服务器的TLS证书。 (默认: false) [$NO_CHECK_CERTIFICATE]
   --shut-timeout value    最长时间等待数据连接关闭状态。 (默认: "1m0s") [$SHUT_TIMEOUT]
   --tls-cache-size value  所有控制和数据连接的TLS会话缓存的大小。 (默认: 32) [$TLS_CACHE_SIZE]
   --writing-mdtm          使用MDTM来设置修改时间（VsFtpd的怪癖）。 (默认: false) [$WRITING_MDTM]

```
{% endcode %}