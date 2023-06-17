# premiumize.me（高级版）

{% code fullWidth="true" %}
```
名称：
   singularity datasource add premiumizeme - premiumize.me

用法：
   singularity datasource add premiumizeme [命令选项] <数据集名称> <源路径>

说明：
   --premiumizeme-api-key
      API密钥。

      此选项通常不使用 - 请改用oauth。

   --premiumizeme-encoding
      后端的编码方式。
      
      更多信息请参见[概述中的编码部分](/overview/#encoding)。


选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作]导出到CAR文件后删除数据集的文件。 （默认值：false）
   --rescan-interval value  当上次成功扫描后过去此间隔时，自动重新扫描源目录（默认值：禁用）

   针对 premiumizeme 的选项

   --premiumizeme-encoding value  后端的编码方式。 （默认值：“Slash，DoubleQuote，BackSlash，Del，Ctl，InvalidUtf8，Dot”）[$ PREMIUMIZEME_ENCODING]

```
{% endcode %}