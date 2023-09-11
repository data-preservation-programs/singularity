# 1Fichier

{% code fullWidth="true" %}
```
命令名称：
   singularity storage create fichier - 1Fichier

用法：
   singularity storage create fichier [命令选项] [参数...]

描述：
   --api-key
      你的 API 密钥，请从 https://1fichier.com/console/params.pl 获取。

   --shared-folder
      如果你想下载一个共享文件夹，请添加此参数。

   --file-password
      如果你想下载一个受密码保护的共享文件，请添加此参数。

   --folder-password
      如果你想列出一个受密码保护的共享文件夹中的文件，请添加此参数。

   --encoding
      后端的编码方式。
      
      更多信息请参阅 [概述中的编码部分](/overview/#encoding)。


选项：
   --api-key 值  你的 API 密钥，请从 https://1fichier.com/console/params.pl 获取。[$API_KEY]
   --help, -h    显示帮助

   高级选项

   --encoding 值                后端的编码方式。（默认值："Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot"）[$ENCODING]
   --file-password 值           如果你想下载一个受密码保护的共享文件，请添加此参数。[$FILE_PASSWORD]
   --folder-password 值         如果你想列出一个受密码保护的共享文件夹中的文件，请添加此参数。[$FOLDER_PASSWORD]
   --shared-folder 值           如果你想下载一个共享文件夹，请添加此参数。[$SHARED_FOLDER]

   常规选项

   --name 值  存储的名称（默认值：自动生成）
   --path 值  存储的路径

```
{% endcode %}