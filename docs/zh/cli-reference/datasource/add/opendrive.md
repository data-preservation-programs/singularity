# OpenDrive

{% code fullWidth="true" %}
```
名称:
   singularity datasource add opendrive - OpenDrive

用法:
   singularity datasource add opendrive [命令选项] <数据集名称> <源路径>

描述:
   --opendrive-chunk-size
      文件将按此大小分块上传。
      
      注意，这些块是在内存中缓冲的，增加它们将增加内存使用量。

   --opendrive-encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --opendrive-password
      密码。

   --opendrive-username
      用户名。


选项:
   --help, -h  显示帮助
   
   数据准备选项

   --delete-after-export    [危险操作] 在将文件导出为CAR文件后，删除数据集的文件。 (默认值: false)
   --rescan-interval value  当从上次成功扫描已经过去这段时间后，自动重新扫描源目录 (默认值: 禁用)
   --scanning-state value   设置初始扫描状态 (默认值: 准备就绪)

   OpenDrive选项

   --opendrive-chunk-size value  文件将按此大小分块上传。 (默认值: "10Mi") [$OPENDRIVE_CHUNK_SIZE]
   --opendrive-encoding value    后端的编码方式。 (默认值: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot") [$OPENDRIVE_ENCODING]
   --opendrive-password value    密码。 [$OPENDRIVE_PASSWORD]
   --opendrive-username value    用户名。 [$OPENDRIVE_USERNAME]

```
{% endcode %}