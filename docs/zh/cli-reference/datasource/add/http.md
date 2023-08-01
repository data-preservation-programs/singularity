# HTTP

{% code fullWidth="true" %}
```
名称:
   singularity datasource add http - HTTP

用法:
   singularity datasource add http [命令选项] <数据集名> <源路径>

描述:
   --http-headers
      设置所有事务的HTTP头。
      
      使用此选项为所有事务设置额外的HTTP头。
      
      输入格式是以逗号分隔的键值对列表。可以使用标准的[CSV编码](https://godoc.org/encoding/csv)。
      
      例如，要设置一个Cookie，使用'Cookie,name=value'，或'"Cookie","name=value"'。
      
      您可以设置多个头，例如'"Cookie","name=value","Authorization","xxx"'。

   --http-no-head
      不使用HEAD请求。
      
      HEAD请求主要用于在目录列表中找到文件大小。
      如果您的站点加载非常慢，可以尝试此选项。
      通常，rclone会对目录列表中每个潜在文件执行HEAD请求，以执行以下操作：
      
      - 查找其大小
      - 检查它是否真的存在
      - 检查它是否为目录
      
      如果设置了此选项，则rclone不会执行HEAD请求。这意味着目录列表将会快得多，但rclone不会有任何文件的时间或大小，并且列表中可能包含一些不存在的文件。

   --http-no-slash
      如果站点不以/结尾，请设置此选项。
      
      如果目标网站不在目录末尾使用/，请使用此选项。
      
      /在路径末尾是rclone通常用来区分文件和目录的方法。如果设置了此标志，rclone将把所有Content-Type为text/html的文件视为目录，并从中读取URL而不是下载它们。
      
      请注意，这可能导致rclone将真正的HTML文件与目录混淆。

   --http-url
      要连接的HTTP主机的URL。
      
      例如："https://example.com"，或者"https://user:pass@example.com"以使用用户名和密码。

选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险] 将数据集导出为CAR文件后，删除数据集文件。 (默认: false)
   --rescan-interval value  当距离上次成功扫描已经过去指定的间隔时间时，自动重新扫描源目录 (默认: 禁用)
   --scanning-state value   设置初始扫描状态 (默认: 就绪)

   http选项

   --http-headers value   设置所有事务的HTTP头。[$HTTP_HEADERS]
   --http-no-head value   不使用HEAD请求。 (默认: "false") [$HTTP_NO_HEAD]
   --http-no-slash value  如果站点不以/结尾，请设置此选项。 (默认: "false") [$HTTP_NO_SLASH]
   --http-url value       要连接的HTTP主机的URL。[$HTTP_URL]

```
{% endcode %}