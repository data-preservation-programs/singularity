# FTP

{% code fullWidth="true" %}
```
名称：
   singularity storage create ftp - FTP

用法：
   singularity storage create ftp [命令选项] [参数...]

描述：
   --host
      连接的FTP主机。
      
      例如："ftp.example.com"。

   --user
      FTP用户名。

   --port
      FTP端口号。

   --pass
      FTP密码。

   --tls
      使用隐式FTPS（FTP over TLS）。
      
      使用隐式FTPS时，客户端从一开始就使用TLS连接，这与不支持TLS的服务器不兼容。
      隐式FTPS通常在端口990上提供，而不是端口21。不能与显示FTPS组合使用。

   --explicit-tls
      使用显示FTPS（FTP over TLS）。
      
      使用显示FTPS时，客户端明确要求服务器提供安全性，以将纯文本连接升级为加密连接。
      不能与隐式FTPS组合使用。

   --concurrency
      FTP同时连接的最大数量，设置为0表示无限制。
      
      请注意，设置此选项很有可能导致死锁，因此应谨慎使用。
      
      如果正在进行同步或复制操作，请确保并发数比`--transfers`和`--checkers`之和多1。
      
      如果使用了`--check-first`，只需要比`--checkers`和`--transfers`中的最大值多1。
      
      因此，对于`concurrency 3`，可以使用`--checkers 2 --transfers 2 --check-first`或`--checkers 1 --transfers 1`。

   --no-check-certificate
      禁止验证服务器的TLS证书。

   --disable-epsv
      禁用EPSV，即使服务器支持。

   --disable-mlsd
      禁用MLSD，即使服务器支持。

   --disable-utf8
      禁用UTF-8，即使服务器支持。

   --writing-mdtm
      使用MDTM设置修改时间（VsFtpd的特殊情况）

   --force-list-hidden
      使用LIST -a来强制列出隐藏文件和文件夹。将禁用MLSD。

   --idle-timeout
      闲置连接关闭之前的最长时间。
      
      如果在给定的时间内没有连接返回到连接池，则rclone将清空连接池。
      
      设置为0以无限期保持连接。

   --close-timeout
      关闭响应的最长等待时间。

   --tls-cache-size
      所有控制和数据连接的TLS会话缓存大小。
      
      TLS缓存允许恢复TLS会话并在连接之间重用PSK。
      如果默认大小不够，可增加此值以避免TLS恢复错误。
      默认启用。使用0来禁用。

   --disable-tls13
      禁用TLS 1.3（用于具有错误的TLS的FTP服务器的解决方法）

   --shut-timeout
      等待数据连接关闭状态的最长时间。

   --ask-password
      允许在需要时询问FTP密码。
      
      如果设置了此选项且没有提供密码，则rclone将要求输入密码

   --encoding
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

      示例:
         | Asterisk,Ctl,Dot,Slash                               | ProFTPd不能处理文件名中的'*'
         | BackSlash,Ctl,Del,Dot,RightSpace,Slash,SquareBracket | PureFTPd不能处理文件名中的'[]'或'*'
         | Ctl,LeftPeriod,Slash                                 | VsFTPd不能处理以点开头的文件名


选项：
   --explicit-tls  使用显示FTPS（FTP over TLS）。(默认值: false) [$EXPLICIT_TLS]
   --help, -h      显示帮助
   --host value    连接的FTP主机。[$HOST]
   --pass value    FTP密码。[$PASS]
   --port value    FTP端口号。(默认值: 21) [$PORT]
   --tls           使用隐式FTPS（FTP over TLS）。(默认值: false) [$TLS]
   --user value    FTP用户名。(默认值: "$USER") [$USER]

   高级选项

   --ask-password          允许在需要时询问FTP密码。(默认值: false) [$ASK_PASSWORD]
   --close-timeout value   最长等待响应的时间来关闭连接。(默认值: "1m0s") [$CLOSE_TIMEOUT]
   --concurrency value     FTP同时连接的最大数量，设置为0表示无限制。(默认值: 0) [$CONCURRENCY]
   --disable-epsv          禁用EPSV即使服务器支持。(默认值: false) [$DISABLE_EPSV]
   --disable-mlsd          禁用MLSD即使服务器支持。(默认值: false) [$DISABLE_MLSD]
   --disable-tls13         禁用TLS 1.3（用于具有错误的TLS的FTP服务器的解决方法）(默认值: false) [$DISABLE_TLS13]
   --disable-utf8          禁用UTF-8即使服务器支持。(默认值: false) [$DISABLE_UTF8]
   --encoding value        后端的编码。(默认值: "Slash,Del,Ctl,RightSpace,Dot") [$ENCODING]
   --force-list-hidden     使用LIST -a来强制列出隐藏文件和文件夹。将禁用MLSD。(默认值: false) [$FORCE_LIST_HIDDEN]
   --idle-timeout value    闲置连接关闭之前的最长时间。(默认值: "1m0s") [$IDLE_TIMEOUT]
   --no-check-certificate  禁止验证服务器的TLS证书。(默认值: false) [$NO_CHECK_CERTIFICATE]
   --shut-timeout value    等待数据连接关闭状态的最长时间。(默认值: "1m0s") [$SHUT_TIMEOUT]
   --tls-cache-size value  所有控制和数据连接的TLS会话缓存大小。(默认值: 32) [$TLS_CACHE_SIZE]
   --writing-mdtm          使用MDTM设置修改时间（VsFtpd的特殊情况）(默认值: false) [$WRITING_MDTM]

   常规选项

   --name value  存储的名称(默认值: 自动生成)
   --path value  存储的路径

```
{% endcode %}