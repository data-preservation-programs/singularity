# OpenDrive

{% code fullWidth="true" %}
```
名称：
   singularity datasource add opendrive - OpenDrive

使用：
   singularity datasource add opendrive [命令选项] <数据集名称> <源路径>

说明：
   --opendrive-username
      用户名。

   --opendrive-password
      密码。

   --opendrive-encoding
      后端的编码。
      
      有关详细信息，请参见[概述中的编码部分](/overview/#encoding)。

   --opendrive-chunk-size
      文件将按此大小分块上传。
      
      请注意，这些块在内存中缓冲，因此增加它们将增加内存使用量。


选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 将数据集导出为CAR文件后，删除数据集的文件。(默认值：false)
   --rescan-interval value  当此时间间隔从上次成功扫描开始后，自动重新扫描源目录 (默认值：禁用)

   OpenDrive的选项

   --opendrive-chunk-size value  文件将按此大小分块上传。(默认值："10Mi") [$OPENDRIVE_CHUNK_SIZE]
   --opendrive-encoding value    后端的编码。(默认值："Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot") [$OPENDRIVE_ENCODING]
   --opendrive-password value    密码。[$OPENDRIVE_PASSWORD]
   --opendrive-username value    用户名。[$OPENDRIVE_USERNAME]

```
{% endcode %}