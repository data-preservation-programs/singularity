# 1Fichier

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add fichier - 1Fichier

USAGE:
   singularity datasource add fichier [command options] <dataset_name> <source_path>

DESCRIPTION:
   --fichier-api-key
      Your API Key, get it from https://1fichier.com/console/params.pl.

   --fichier-shared-folder
      If you want to download a shared folder, add this parameter.

   --fichier-file-password
      If you want to download a shared file that is password protected, add this parameter.

   --fichier-folder-password
      If you want to list the files in a shared folder that is password protected, add this parameter.

   --fichier-encoding
      The encoding for the backend.

      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --fichier-api-key value  Your API Key, get it from https://1fichier.com/console/params.pl. [$FICHIER_API_KEY]
   --help, -h               show help

   Advanced Options

   --fichier-encoding value         The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$FICHIER_ENCODING]
   --fichier-file-password value    If you want to download a shared file that is password protected, add this parameter. [$FICHIER_FILE_PASSWORD]
   --fichier-folder-password value  If you want to list the files in a shared folder that is password protected, add this parameter. [$FICHIER_FOLDER_PASSWORD]
   --fichier-shared-folder value    If you want to download a shared folder, add this parameter. [$FICHIER_SHARED_FOLDER]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)

```
{% endcode %}
