# Sia 分散式云存储

{% code fullWidth="true" %}
```
名称：
   singularity storage create sia - Sia 分散式云存储

用法：
   singularity storage create sia [命令选项] [参数...]

描述：
   --api-url
      Sia 守护进程 API 的 URL，例如 http://sia.daemon.host:9980。
      
      请注意，siad 必须以 --disable-api-security 启动，才能为其他主机打开 API 端口（不建议）。
      如果 Sia 守护进程在本地运行，请保持默认设置。

   --api-password
      Sia 守护进程 API 密码。
      
      可在位于 HOME/.sia/ 或守护进程目录中的 apipassword 文件中找到。

   --user-agent
      Siad 用户代理
      
      Sia 守护进程默认需要 'Sia-Agent' 用户代理以提高安全性

   --encoding
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。


选项：
   --api-password value  Sia 守护进程 API 密码。[$API_PASSWORD]
   --api-url value       Sia 守护进程 API 的 URL，例如 http://sia.daemon.host:9980。 (default: "http://127.0.0.1:9980") [$API_URL]
   --help, -h            显示帮助

   高级选项

   --encoding value    后端的编码。 (default: "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --user-agent value  Siad 用户代理 (default: "Sia-Agent") [$USER_AGENT]

   常规选项

   --name value  存储的名称（默认为自动生成）
   --path value  存储的路径

```
{% endcode %}