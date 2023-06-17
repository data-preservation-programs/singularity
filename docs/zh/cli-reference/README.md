# CLI 参考手册

{% code fullWidth="true" %}
```
名称：
   singularity - 用于向 Filecoin 网络上载 PB 级别数据的大规模客户端工具

用法：
   singularity [全局选择项] 命令 [命令选择项] [参数...]

命令：
   help, h  显示命令列表或单个命令的帮助
   Daemons：
     run  运行不同的 singularity 组件
   简易命令：
     ez-prep  从本地路径准备数据集
   操作：
     admin       管理员命令
     deal        复制/交易管理
     dataset     数据集管理
     datasource  数据源管理
     wallet      钱包管理
   实用工具：
     download 从元数据 API 下载 CAR 文件

全局选择项:
   --database-connection-string CREATE DATABASE <dbname> DEFAULT CHARACTER SET ascii  连接到数据库的字符形式。
      支持的数据库：sqlite3，postgres，mysql
      postgres 示例 - postgres://user:pass@example.com:5432/dbname
      mysql 示例 - mysql://user:pass@tcp(localhost:3306)/dbname?charset=ascii&parseTime=true
        注意：需要使用 ASCII 字符集创建数据库：CREATE DATABASE <dbname> DEFAULT CHARACTER SET ascii
      sqlite3 示例- sqlite:/absolute/path/to/database.db
                或 - sqlite:relative/path/to/database.db
       (默认值：sqlite：/home/shane/.singularity/singularity.db) [$DATABASE_CONNECTION_STRING]
   --verbose   启用详细日志记录 ( 默认值：false )
   --json      启用 JSON 输出 ( 默认值：false )
   --help, -h  显示帮助
```
{% endcode %}