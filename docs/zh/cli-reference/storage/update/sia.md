# Sia 分散化云存储

{% code fullWidth="true" %}
```
命令名称:
   singularity storage update sia - Sia 分散化云存储

使用方法:
   singularity storage update sia [command options] <名字|ID>

描述:
   --api-url
      Sia 守护进程的 API URL，例如 http://sia.daemon.host:9980。
      
      注意，Sia 守护进程必须以 --disable-api-security 选项启动，才能为其他主机开放API端口（不推荐）。
      如果 Sia 守护进程运行在本地主机上，请使用默认值。

   --api-password
      Sia 守护进程的 API 密码。
      
      可以在 HOME/.sia/ 或守护进程目录中的 apipassword 文件中找到。

   --user-agent
      Siad 用户代理
      
      为了安全起见，默认情况下，Sia 守护进程需要 'Sia-Agent' 用户代理

   --encoding
      后端的编码方式。
      
      详见[概述中的编码部分](/overview/#encoding)了解更多信息。


选项:
   --api-password value  Sia 守护进程的 API 密码。 [$API_PASSWORD]
   --api-url value       Sia 守护进程的 API URL，例如 http://sia.daemon.host:9980。 (默认值: "http://127.0.0.1:9980") [$API_URL]
   --help, -h            显示帮助

   高级选项

   --encoding value    后端的编码方式。 (默认值: "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --user-agent value  Siad 用户代理 (默认值: "Sia-Agent") [$USER_AGENT]

```
{% endcode %}