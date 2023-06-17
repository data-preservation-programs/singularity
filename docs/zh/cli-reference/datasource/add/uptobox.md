# Uptobox

{% code fullWidth="true" %}
```
名称：
   singularity datasource add uptobox - Uptobox

用法：
   singularity datasource add uptobox [命令选项] <数据集名称> <源路径>

描述：
   --uptobox-access-token
      您的访问令牌。
      
      从https://uptobox.com/my_account获取。

   --uptobox-encoding
      后端使用的编码方式。
      
      关于“编码”细节，请参见[概览中的编码部分](/overview/#encoding)。

选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 将数据集导出为CAR文件后删除数据集中的文件。 (默认值：false)
   --rescan-interval value  上一次成功扫描后，等待重新扫描源目录的时间间隔（默认值：禁用）

   Uptobox选项

   --uptobox-access-token value  您的访问令牌。 [$UPTOBOX_ACCESS_TOKEN]
   --uptobox-encoding value      后端使用的编码方式。 (默认值： "Slash, LtGt, DoubleQuote, BackQuote, Del, Ctl, LeftSpace, InvalidUtf8, Dot") [$UPTOBOX_ENCODING]
```
{% endcode %}