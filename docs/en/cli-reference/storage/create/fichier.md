# 1Fichier

{% code fullWidth="true" %}
```
NAME:
   singularity storage create fichier - 1Fichier

USAGE:
   singularity storage create fichier [command options] <name> <path>

DESCRIPTION:
   --api-key
      Your API Key, get it from https://1fichier.com/console/params.pl.

   --shared-folder
      If you want to download a shared folder, add this parameter.

   --file-password
      If you want to download a shared file that is password protected, add this parameter.

   --folder-password
      If you want to list the files in a shared folder that is password protected, add this parameter.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --api-key value  Your API Key, get it from https://1fichier.com/console/params.pl. [$API_KEY]
   --help, -h       show help

   Advanced

   --encoding value         The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --file-password value    If you want to download a shared file that is password protected, add this parameter. [$FILE_PASSWORD]
   --folder-password value  If you want to list the files in a shared folder that is password protected, add this parameter. [$FOLDER_PASSWORD]
   --shared-folder value    If you want to download a shared folder, add this parameter. [$SHARED_FOLDER]

```
{% endcode %}