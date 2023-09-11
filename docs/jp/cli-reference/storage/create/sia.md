# Sia 分布式云存储

{% code fullWidth="true" %}
```
NAME:
   singularity storage create sia - Sia 分布式云存储

USAGE:
   singularity storage create sia [command options] [arguments...]

DESCRIPTION:
   --api-url
      Sia守护程序的API URL，例如http://sia.daemon.host:9980。
      
      请注意，siad必须以--disable-api-security运行，以打开API端口供其他主机使用（不建议）。
      如果Sia守护程序在本地主机上运行，请保持默认值。

   --api-password
      Sia守护程序的API密码。
      
      可以在HOME/.sia/或守护程序目录中的apipassword文件中找到。

   --user-agent
      Siad用户代理
      
      Sia守护程序默认要求使用“Sia-Agent”用户代理以提高安全性。

   --encoding
      后端的编码方式。
      
      更多信息，请参考[概述中的编码章节](/overview/#encoding)。


OPTIONS:
   --api-password value  Sia守护程序的API密码。[$API_PASSWORD]
   --api-url value       Sia守护程序的API URL，例如http://sia.daemon.host:9980。（默认值："http://127.0.0.1:9980"）[$API_URL]
   --help, -h            显示帮助

   高级选项

   --encoding value    后端的编码方式（默认值："Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot"）[$ENCODING]
   --user-agent value  Siad用户代理（默认值："Sia-Agent"）[$USER_AGENT]

   通用选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}