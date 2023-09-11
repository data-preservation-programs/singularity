# 1Fichier

{% code fullWidth="true" %}
```
名称：
   singularity storage update fichier - 1Fichier

用法：
   singularity storage update fichier [命令选项] <名称|ID>

描述：
   --api-key
      您的API密钥，从https://1fichier.com/console/params.pl获取。

   --shared-folder
      如果您想下载共享文件夹，请添加此参数。

   --file-password
      如果您想下载受密码保护的共享文件，请添加此参数。

   --folder-password
      如果您想列出受密码保护的共享文件夹中的文件，请添加此参数。

   --encoding
      后端的编码。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

选项：
   --api-key value  您的API密钥，从https://1fichier.com/console/params.pl获取。 [$API_KEY]
   --help, -h       显示帮助信息

   高级选项

   --encoding value         后端的编码。（默认值："Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"）[$ENCODING]
   --file-password value    如果您想下载受密码保护的共享文件，请添加此参数。[$FILE_PASSWORD]
   --folder-password value  如果您想列出受密码保护的共享文件夹中的文件，请添加此参数。[$FOLDER_PASSWORD]
   --shared-folder value    如果您想下载共享文件夹，请添加此参数。[$SHARED_FOLDER]

```
{% endcode %}