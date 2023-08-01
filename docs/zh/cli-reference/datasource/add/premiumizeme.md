# premiumize.me

{% code fullWidth="true" %}
```
名称:
   singularity数据源添加premiumizeme - premiumize.me

用法:
   singularity数据源添加premiumizeme [命令选项] <数据集名称> <源路径>

说明:
   --premiumizeme-api-key
      API密钥。
      
      此选项通常不使用 - 请改用oauth。
      

   --premiumizeme-encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码方式部分](/overview/#encoding)。


选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险操作] 在将数据集导出为CAR文件后删除数据集的文件。  (默认值: false)
   --rescan-interval value  当距离上一次成功扫描的间隔超过此值时，自动重新扫描源目录 (默认值: 禁用)
   --scanning-state value   设置初始扫描状态 (默认值: 就绪)

   premiumizeme选项

   --premiumizeme-encoding value  后端的编码方式。 (默认值: "Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PREMIUMIZEME_ENCODING]

```
{% endcode %}