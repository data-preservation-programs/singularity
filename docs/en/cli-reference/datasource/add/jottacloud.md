# Jottacloud

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add jottacloud - Jottacloud

USAGE:
   singularity datasource add jottacloud [command options] <dataset_name> <source_path>

DESCRIPTION:
   --jottacloud-trashed-only
      Only show files that are in the trash.
      
      This will show trashed files in their original directory structure.

   --jottacloud-hard-delete
      Delete files permanently rather than putting them into the trash.

   --jottacloud-upload-resume-limit
      Files bigger than this can be resumed if the upload fail's.

   --jottacloud-no-versions
      Avoid server side versioning by deleting files and recreating files instead of overwriting them.

   --jottacloud-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --jottacloud-md5-memory-limit
      Files bigger than this will be cached on disk to calculate the MD5 if required.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for jottacloud

   --jottacloud-encoding value             The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot") [$JOTTACLOUD_ENCODING]
   --jottacloud-hard-delete value          Delete files permanently rather than putting them into the trash. (default: "false") [$JOTTACLOUD_HARD_DELETE]
   --jottacloud-md5-memory-limit value     Files bigger than this will be cached on disk to calculate the MD5 if required. (default: "10Mi") [$JOTTACLOUD_MD5_MEMORY_LIMIT]
   --jottacloud-no-versions value          Avoid server side versioning by deleting files and recreating files instead of overwriting them. (default: "false") [$JOTTACLOUD_NO_VERSIONS]
   --jottacloud-trashed-only value         Only show files that are in the trash. (default: "false") [$JOTTACLOUD_TRASHED_ONLY]
   --jottacloud-upload-resume-limit value  Files bigger than this can be resumed if the upload fail's. (default: "10Mi") [$JOTTACLOUD_UPLOAD_RESUME_LIMIT]

```
{% endcode %}
