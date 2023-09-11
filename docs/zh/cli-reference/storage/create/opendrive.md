# OpenDrive

{% code fullWidth="true" %}
```
名称:
   singularity storage create opendrive - OpenDrive

使用方法:
   singularity storage create opendrive [命令选项] [参数...]

描述:
   --username
      用户名。

   --password
      密码。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --chunk-size
      文件将按照此大小分块上传。
      
      注意，这些分块缓存在内存中，增加它们的大小将增加内存使用量。


选项:
   --help, -h        显示帮助信息
   --password value  密码。 [$PASSWORD]
   --username value  用户名。 [$USERNAME]

   进阶

   --chunk-size value  文件将按照此大小分块上传。 (默认值: "10Mi") [$CHUNK_SIZE]
   --encoding value    后端的编码方式。 (默认值: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot") [$ENCODING]

   常规

   --name value  存储的名称 (默认值：自动生成)
   --path value  存储的路径

```
{% endcode %}