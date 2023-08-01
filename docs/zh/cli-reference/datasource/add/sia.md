# Sia分散式云

{% code fullWidth="true" %}
```
命令名:
   singularity datasource add sia - Sia分散式云

使用方法:
   singularity datasource add sia [可选项] <数据集名称> <数据源路径>

描述:
   --sia-api-password
      Sia守护进程API密码。
      
      可在HOME/.sia/或守护进程目录中的apipassword文件中找到。

   --sia-api-url
      Sia守护进程API URL，格式为http://sia.daemon.host:9980。
      
      请注意，siad必须使用--disable-api-security参数运行，以便为其他主机开放API端口（不建议）。
      如果Sia守护进程在本地主机上运行，请保留默认设置。

   --sia-encoding
      后端的编码。
      
      有关更多信息，请参见[概览中的编码部分](/overview/#encoding)。

   --sia-user-agent
      Siad用户代理
      
      对于安全原因，默认情况下，Sia守护进程需要'Sia-Agent'用户代理。


选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险操作]导出数据集为CAR文件后，删除数据集的文件。(默认值: false)
   --rescan-interval value  当距离上次成功扫描已过去指定时间间隔后，自动重新扫描源目录(默认值: 禁用)
   --scanning-state value   设置初始扫描状态(默认值: ready)

   Sia选项

   --sia-api-password value  Sia守护进程API密码。[$SIA_API_PASSWORD]
   --sia-api-url value       Sia守护进程API URL，格式为http://sia.daemon.host:9980。(默认值: "http://127.0.0.1:9980") [$SIA_API_URL]
   --sia-encoding value      后端的编码。(默认值: "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot") [$SIA_ENCODING]
   --sia-user-agent value    Siad用户代理(默认值: "Sia-Agent") [$SIA_USER_AGENT]

```
{% endcode %}