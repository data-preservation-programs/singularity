# 创建和管理数据集准备

{% code fullWidth="true" %}
```
名称：
   singularity prep - 创建和管理数据集准备

用法：
   singularity prep 命令 [命令选项] [参数...]

命令：
   create         创建新的准备
   list           列出所有准备
   status         获取准备的准备作业状态
   attach-source  将源存储附加到准备
   attach-output  将输出存储附加到准备
   detach-output  从准备中分离输出存储
   start-scan     开始扫描源存储
   pause-scan     暂停扫描作业
   start-pack     启动/重启所有打包作业或特定作业
   pause-pack     暂停所有打包作业或特定作业
   start-daggen   启动生成包含所有文件夹结构快照的DAG生成
   pause-daggen   暂停DAG生成作业
   list-pieces    列出准备的所有生成的片段
   add-piece      手动向准备中添加片段信息。这对于由外部工具准备的片段很有用。
   explore        通过路径浏览准备的源
   attach-wallet  将钱包附加到准备
   list-wallets   列出与准备相关的已附加钱包
   detach-wallet  从准备中分离钱包
   help, h        显示命令列表或单个命令的帮助

选项：
   --help, -h  显示帮助信息
```
{% endcode %}