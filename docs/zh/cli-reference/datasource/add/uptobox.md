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
      
      从 https://uptobox.com/my_account 获取。

  --uptobox-encoding
      后端编码。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。


选项：
  --help, -h  显示帮助

  数据准备选项

  --delete-after-export    【危险】将数据集导出为CAR文件后删除数据集文件。（默认值：false）
  --rescan-interval value  当上一次成功扫描之后经过此时间间隔时，自动重新扫描源目录。（默认值：disabled）
  --scanning-state value   设置初始扫描状态。（默认值：ready）

  Uptobox选项

  --uptobox-access-token value  您的访问令牌。[$UPTOBOX_ACCESS_TOKEN]
  --uptobox-encoding value      后端编码。（默认值："Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot"）[$UPTOBOX_ENCODING]

```
{% endcode %}