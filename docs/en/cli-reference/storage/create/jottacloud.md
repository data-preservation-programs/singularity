# Jottacloud

{% code fullWidth="true" %}
```
NAME:
   singularity storage create jottacloud - Jottacloud

USAGE:
   singularity storage create jottacloud [command options] <name> <path>

DESCRIPTION:
   --md5_memory_limit
      Files bigger than this will be cached on disk to calculate the MD5 if required.

   --trashed_only
      Only show files that are in the trash.
      
      This will show trashed files in their original directory structure.

   --hard_delete
      Delete files permanently rather than putting them into the trash.

   --upload_resume_limit
      Files bigger than this can be resumed if the upload fail's.

   --no_versions
      Avoid server side versioning by deleting files and recreating files instead of overwriting them.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h  show help

   Advanced

   --encoding value             The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hard_delete                Delete files permanently rather than putting them into the trash. (default: false) [$HARD_DELETE]
   --md5_memory_limit value     Files bigger than this will be cached on disk to calculate the MD5 if required. (default: "10Mi") [$MD5_MEMORY_LIMIT]
   --no_versions                Avoid server side versioning by deleting files and recreating files instead of overwriting them. (default: false) [$NO_VERSIONS]
   --trashed_only               Only show files that are in the trash. (default: false) [$TRASHED_ONLY]
   --upload_resume_limit value  Files bigger than this can be resumed if the upload fail's. (default: "10Mi") [$UPLOAD_RESUME_LIMIT]

```
{% endcode %}
