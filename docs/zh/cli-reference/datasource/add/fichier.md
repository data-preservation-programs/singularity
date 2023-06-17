# 1Fichier

{% code fullWidth="true" %}
```
名称：
   singularity 数据源添加 fichier - 1Fichier

用法：
   singularity datasource add fichier [command options] <dataset_name> <source_path>

描述：
   --fichier-api-key
      您的API密钥，请从 https://1fichier.com/console/params.pl 获取。

   --fichier-shared-folder
      如果您想下载共享文件夹，请添加此参数。

   --fichier-file-password
      如果要下载受密码保护的共享文件，请添加此参数。

   --fichier-folder-password
      如果要列出受密码保护的共享文件夹中的文件，请添加此参数。

   --fichier-encoding
      后端的编码方式。

      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。


选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 导出 CAR 文件后删除数据集文件。 (默认：false)
   --rescan-interval value  自动重新扫描源目录的时间间隔，以便上次扫描成功后经过一段时间。（默认：禁用）

   fichier选项

   --fichier-api-key value          您的API密钥，请从 https://1fichier.com/console/params.pl 获取。 [$FICHIER_API_KEY]
   --fichier-encoding value         后端的编码方式。（默认值为 "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"） [$FICHIER_ENCODING]
   --fichier-file-password value    如果要下载受密码保护的共享文件，请添加此参数。 [$FICHIER_FILE_PASSWORD]
   --fichier-folder-password value  如果要列出受密码保护的共享文件夹中的文件，请添加此参数。 [$FICHIER_FOLDER_PASSWORD]
   --fichier-shared-folder value    如果要下载共享文件夹，请添加此参数。 [$FICHIER_SHARED_FOLDER]

```
{% endcode %}