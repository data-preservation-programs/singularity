# OpenDrive

{% code fullWidth="true" %}
```
名称:
   singularity storage update opendrive - OpenDrive

用法:
   singularity storage update opendrive [命令选项] <名称|ID>

说明:
   --username
      用户名。

   --password
      密码。

   --encoding
      后端的编码。
      
      更多信息，请参考[概览中的编码部分](/overview/#encoding)。

   --chunk-size
      文件将以此大小分块上传。
      
      注意，这些分块将缓存在内存中，因此增加分块大小将增加内存使用量。


选项:
   --help, -h        显示帮助
   --password value  密码。[$PASSWORD]
   --username value  用户名。[$USERNAME]

   高级选项:

   --chunk-size value  文件将以此大小分块上传。(默认: "10Mi") [$CHUNK_SIZE]
   --encoding value    后端的编码。(默认: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot") [$ENCODING]

```
{% endcode %}