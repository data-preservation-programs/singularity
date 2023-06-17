# Put.io

{% code fullWidth="true" %}
```
名称：
   singularity datasource add putio - Put.io

用法：
   singularity datasource add putio [命令选项] <数据集名称> <源路径>

说明：
   --putio-encoding
      后端的编码形式。
      
      更多信息请参见[概述中的编码部分](/overview/#encoding)。


选项：
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险] 导出数据集为CAR文件后删除数据集中的文件。（默认值：false）
   --rescan-interval value  自动重新扫描源目录，当这个时间间隔从上次成功扫描后到达时。（默认值：禁用）

   Put.io选项

   --putio-encoding value  后端的编码形式。（默认值："Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"）[$PUTIO_ENCODING]
```
{% endcode %}