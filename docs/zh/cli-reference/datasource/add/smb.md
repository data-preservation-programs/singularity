# SMB / CIFS

{% code fullWidth="true" %}
```
名称:
    singularity datasource add smb - SMB / CIFS

用法:
    singularity datasource add smb [命令选项] <数据集名称> <源路径>

描述:
    --smb-encoding
        后端编码。

        有关更多信息，请参见[概述](/overview/#encoding)中的编码部分。

    --smb-user
        SMB用户名。

    --smb-port
        SMB端口号。

    --smb-domain
        用于NTLM身份验证的域名。

    --smb-hide-special-share
        隐藏不应由用户访问的特殊共享（例如print$）。

    --smb-case-insensitive
        服务器是否配置为不区分大小写。

        在Windows共享上始终为真。

    --smb-host
        要连接到的SMB服务器主机名。

        例如："example.com"。

    --smb-pass
        SMB密码。

    --smb-spn
        服务主体名称。

        Rclone将此名称呈现给服务器。某些服务器将其用作进一步的身份验证，并且通常需要设置集群。例如：

            cifs/remotehost:1020

         如果不确定，请保留为空。

    --smb-idle-timeout
        关闭空闲连接之前的最大时间。

        如果在给定的时间内没有将连接返回到连接池，则rclone将清空连接池。

        将其设置为0以无限制保留连接。


选项:
    --help，-h  显示帮助

    数据准备选项

    --delete-after-export    [危险]导出为CAR文件后删除数据集的文件。  (默认: false)
    --rescan-interval value  当此间隔从上次成功扫描过去时，自动重新扫描源目录（默认值: 禁用）

    smb选项

    --smb-case-insensitive value    服务器是否配置为不区分大小写。(默认值："true") [$SMB_CASE_INSENSITIVE]
    --smb-domain value              用于NTLM身份验证的域名。(默认值："WORKGROUP") [$SMB_DOMAIN]
    --smb-encoding value            后端编码。(默认值："Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot") [$SMB_ENCODING]
    --smb-hide-special-share value  隐藏不应由用户访问的特殊共享（例如print$）(默认值："true") [$SMB_HIDE_SPECIAL_SHARE]
    --smb-host value                要连接到的SMB服务器主机名。[$SMB_HOST]
    --smb-idle-timeout value        关闭空闲连接之前的最大时间。(默认值："1m0s") [$SMB_IDLE_TIMEOUT]
    --smb-pass value                SMB密码。[$SMB_PASS]
    --smb-port value                SMB端口号。 (默认值："445") [$SMB_PORT]
    --smb-spn value                 服务主体名称。 [$SMB_SPN]
    --smb-user value                SMB用户名。(默认值："shane") [$SMB_USER]

```
{% endcode %}