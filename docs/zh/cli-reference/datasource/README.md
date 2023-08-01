# 数据源管理

{% code fullWidth="true" %}
```
命令：singularity datasource - 数据源管理

使用方式：
   singularity datasource command [命令选项] [参数...]

命令列表：
   add      添加一个新的数据源到数据集
   list     列出所有的数据源
   status   获取数据源的数据准备概要
   remove   删除一个数据源
   check    通过列出数据源的条目来检查数据源连接。这不是列出准备好的项目，而是直接列出数据源以验证数据源连接。
   update   更新数据源的配置选项
   rescan   重新扫描数据源
   daggen   生成并导出代表数据源完整文件夹结构的DAG
   inspect  获取数据源的数据准备状态
   help, h  显示命令列表或有关某个命令的帮助信息

选项：
   --help, -h  显示帮助信息
```
{% endcode %}