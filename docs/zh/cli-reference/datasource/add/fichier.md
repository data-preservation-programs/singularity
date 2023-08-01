# 1Fichier

{% code fullWidth="true" %}
```
名称:
   singularity datasource add fichier - 1Fichier

用法:
   singularity datasource add fichier [命令选项] <数据集名称> <源路径>

描述:
   --fichier-api-key
      您的 API 密钥，从 https://1fichier.com/console/params.pl 获取。
      
   --fichier-encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --fichier-file-password
      如果您要下载一个受密码保护的共享文件，请添加此参数。

   --fichier-folder-password
      如果您要列出一个受密码保护的共享文件夹中的文件，请添加此参数。

   --fichier-shared-folder
      如果您要下载一个共享文件夹，请添加此参数。


选项:
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险] 将数据集导出为 CAR 文件后，删除数据集文件。  (默认值: false)
   --rescan-interval value  当最后一次成功扫描后的时间间隔大于此间隔时，自动重新扫描源目录（默认值：禁用）
   --scanning-state value   设置初始扫描状态（默认值：ready）

   fichier选项

   --fichier-api-key value          您的 API 密钥，从 https://1fichier.com/console/params.pl 获取。 [$FICHIER_API_KEY]
   --fichier-encoding value         后端的编码方式。（默认值: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"） [$FICHIER_ENCODING]
   --fichier-file-password value    如果您要下载一个受密码保护的共享文件，请添加此参数。 [$FICHIER_FILE_PASSWORD]
   --fichier-folder-password value  如果您要列出一个受密码保护的共享文件夹中的文件，请添加此参数。 [$FICHIER_FOLDER_PASSWORD]
   --fichier-shared-folder value    如果您要下载一个共享文件夹，请添加此参数。 [$FICHIER_SHARED_FOLDER]

```
{% endcode %}