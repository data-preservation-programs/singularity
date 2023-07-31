# Put.io

{% code fullWidth="true" %}
```
名称:
   singularity datasource add putio - Put.io

用法:
   singularity datasource add putio [命令选项] <数据集名称> <源路径>

描述:
   --putio-encoding
      后端的编码。
      
      更多信息请参阅[概述中的编码部分](/overview/#encoding)。


选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 导出数据集为CAR文件后删除文件。(默认: false)
   --rescan-interval value  当上次成功扫描后经过此时间间隔后，自动重新扫描源目录（默认：禁用）
   --scanning-state value   设置初始扫描状态（默认：就绪）

   putio选项

   --putio-encoding value  后端的编码。(默认: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PUTIO_ENCODING]

```
{% endcode %}