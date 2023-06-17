# Sia 分布式云

{% code fullWidth="true" %}
```
命令：
    singularity datasource add sia - Sia 分布式云

用法：
    singularity datasource add sia [命令选项] <数据集名称> <源路径>

描述：
    --sia-api-url
        Sia 守护进程 API URL，如 http://sia.daemon.host:9980。
        
        注意，如果要开放 API 端口供其他主机使用（不建议），则必须以 --disable-api-security 标志运行 siad。
        如果 Sia 守护进程运行在 localhost 上，请保持默认设置。

    --sia-api-password
        Sia 守护进程 API 密码。
        
        可在 HOME/.sia/ 或守护进程目录中的 apipassword 文件中找到。

    --sia-user-agent
        Siad 用户代理。
        
        为了安全起见，Sia 守护进程默认需要“Sia-Agent”用户代理。

    --sia-encoding
        后端编码。
        
        更多信息请参见[概述中的编码部分](/overview/#encoding)。

选项：
    --help, -h  显示帮助信息

    数据准备选项

    --delete-after-export    [危险操作] 将数据集导出到 CAR 文件后删除数据集中的文件。 (默认为 false)
    --rescan-interval value  在上一次成功扫描之后，自动重新扫描源目录的时间间隔。 (默认为禁用)

    Sia 的选项

    --sia-api-password value  Sia 守护进程 API 密码。[$SIA_API_PASSWORD]
    --sia-api-url value       Sia 守护进程 API URL，如 http://sia.daemon.host:9980。 (默认为 "http://127.0.0.1:9980") [$SIA_API_URL]
    --sia-encoding value      后端编码。（默认为“Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot”）[$SIA_ENCODING]
    --sia-user-agent value    Siad 用户代理（默认为“Sia-Agent”）。[$SIA_USER_AGENT]

```
{% endcode %}