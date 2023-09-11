# HTTP

{% code fullWidth="true" %}
```
名称:
   singularity存储更新http - HTTP

用法:
   singularity存储更新http [命令选项] <名称|编号>

说明:
   --url
      要连接的HTTP主机的URL。
      
      例如，"https://example.com"，或者"https://user:pass@example.com"以使用用户名和密码。

   --headers
      设置所有事务的HTTP头部。
      
      使用此选项来设置所有事务的附加HTTP头部。
      
      输入格式为用逗号分隔的键值对列表。可以使用标准的[CSV编码](https://godoc.org/encoding/csv)。
      
      例如，要设置Cookie，请使用'Cookie,name=value'或'"Cookie","name=value"'。
      
      您可以设置多个头部，例如，'"Cookie","name=value","Authorization","xxx"'。

   --no-slash
      如果站点没有在目录末尾使用/，请设置此选项。
      
      如果目标网站不在目录末尾使用/，请使用此选项。
      
      在路径末尾加上/是rclone通常用来区分文件和目录的方式。如果设置了此标志，
      则rclone将把所有具有Content-Type: text/html的文件视为目录，并从中读取URL而不是下载它们。
      
      请注意，这可能会导致rclone将真正的HTML文件与目录混淆。

   --no-head
      不使用HEAD请求。
      
      HEAD请求主要用于在目录列表中查找文件大小。
      如果您的网站加载非常慢，您可以尝试使用此选项。
      通常rclone对目录列表中的每个潜在文件执行HEAD请求来：
      
      - 查找其大小
      - 检查它是否真实存在
      - 检查它是否为目录
      
      如果您设置了此选项，rclone将不会执行HEAD请求。这意味着
      目录列表将更快，但rclone将没有任何文件的时间或大小，并且在列表中可能存在一些不存在的文件。


选项：
   --help, -h   显示帮助信息
   --url value  要连接的HTTP主机的URL。[$URL]

   高级选项

   --headers value  设置所有事务的HTTP头部。[$HEADERS]
   --no-head        不使用HEAD请求。（默认为false）[$NO_HEAD]
   --no-slash       如果站点没有在目录末尾使用/，请设置此选项。（默认为false）[$NO_SLASH]
```
{% endcode %}