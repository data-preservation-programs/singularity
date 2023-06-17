# HTTP

{% code fullWidth="true" %}
```
名称：
   singularity 数据源添加 http - HTTP

用法:
   singularity datasource add http [命令选项] <数据集名称> <源路径>

说明：
   --http-url
      连接到 HTTP 主机的 URL。
      
      例如 "https://example.com" 或者 "https://user:pass@example.com" （使用用户名和密码）.

   --http-headers
      设置所有事务的 HTTP 标头。
      
      可以用这个选项为所有事务设置额外的 HTTP 标头。
      
      输入格式为逗号分隔的键值对列表。可以使用标准的 CSV 编码。
      
      例如，要设置一个 Cookie，可以使用 'Cookie,name=value' 或者 '"Cookie","name=value"'。
      
      您可以设置多个标头，例如：'"Cookie","name=value","Authorization","xxx"'。

   --http-no-slash
      如果该站点不能在目录结尾使用 /，则设置此项。
      
      如果你的目标站点没有在目录结尾使用 /，请使用此选项。
      
      在路径末尾加上一个 / 通常是 rclone 用于区分文件和目录的方式。如果设置了此标志，则 rclone 将会将所有具有 Content-Type: text/html 的文件视为目录，并从它们读取 URL，而不是下载它们。
      
      请注意，这可能会导致 rclone 将真正的 HTML 文件与目录混淆。

   --http-no-head
      不使用 HEAD 请求。
      
      HEAD 请求主要用于在目录列表中查找文件大小。
      如果您的站点加载速度非常慢，您可以尝试此选项。
      通常，rclone 对目录列表中的每个潜在文件执行 HEAD 请求以：
      
      - 查找其大小
      - 检查它是否真实存在
      - 检查是否为目录
      
      如果您设置了此选项，则 rclone 将不会执行 HEAD 请求。这将意味着目录列表加载速度更快，但 rclone 没有任何文件的时间或大小，并且某些不存在的文件可能会出现在列表中。

选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 在将数据集导出到 CAR 文件后删除数据集的文件。  (默认值：false)
   --rescan-interval value  当自上次成功扫描以来经过此间隔后，自动重新扫描源目录（默认值：已禁用）

   http 选项

   --http-headers value   设置所有事务的 HTTP 标头。[$HTTP_HEADERS]
   --http-no-head value   不使用 HEAD 请求。（默认值：false）[$HTTP_NO_HEAD]
   --http-no-slash value  如果该站点不能在目录结尾使用 /，则设置此项。（默认值：false）[$HTTP_NO_SLASH]
   --http-url value       连接到 HTTP 主机的 URL。[$HTTP_URL]

```
{% endcode %}