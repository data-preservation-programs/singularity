# HTTP

{% code fullWidth="true" %}
```
名称:
   singularity storage create http - HTTP

使用方法:
   singularity storage create http [命令选项] [参数...]

描述:
   --url
      要连接的HTTP主机的URL。
      
      例如："https://example.com"，或者 "https://user:pass@example.com" 使用用户名和密码。

   --headers
      为所有事务设置HTTP标头。
      
      使用此选项为所有事务设置额外的HTTP标头。
      
      输入格式是逗号分隔的键值对列表。可以使用标准的[CSV编码](https://godoc.org/encoding/csv)。
      
      例如，要设置Cookie，请使用'Cookie,name=value'，或者'"Cookie","name=value"'。
      
      你可以设置多个标头，例如'"Cookie","name=value","Authorization","xxx"'。

   --no-slash
      如果站点不在目录名称后面加斜杠，请设置此选项。
      
      如果你的目标网站在目录名称后面不使用/，请使用此选项。
      
      在路径的末尾加上/是rclone通常用来区分文件和目录的方法。如果设置了此标志，rclone将把Content-Type: text/html的所有文件都视为目录，并从中读取URL，而不是下载文件。
      
      请注意，这可能会导致rclone将真实的HTML文件与目录混淆。

   --no-head
      不使用HEAD请求。
      
      HEAD请求主要用于在目录列表中查找文件大小。
      如果您的网站加载速度非常慢，您可以尝试此选项。
      通常情况下，rclone会对目录列表中的每个潜在文件进行一个HEAD请求，以便：
      
      - 找到文件的大小
      - 检查它是否真实存在
      - 检查它是否是目录
      
      如果设置了此选项，rclone将不会执行HEAD请求。这意味着目录列表加载速度更快，但rclone不会有任何文件的时间或大小，并且可能会在目录列表中包含一些不存在的文件。


选项:
   --help, -h   显示帮助信息
   --url value  要连接的HTTP主机的URL。[$URL]

   高级选项

   --headers value  为所有事务设置HTTP标头。[$HEADERS]
   --no-head        不使用HEAD请求。（默认值: false）[$NO_HEAD]
   --no-slash       如果站点不在目录名称后面加斜杠，请设置此选项。（默认值: false）[$NO_SLASH]

   通用选项

   --name value  存储的名称（默认值: 自动生成）
   --path value  存储的路径

```
{% endcode %}