# premiumize.me

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add premiumizeme - premiumize.me

USAGE:
   singularity datasource add premiumizeme [command options] <dataset_name> <source_path>

DESCRIPTION:
   --premiumizeme-api-key
      API Key.

      This is not normally used - use oauth instead.


   --premiumizeme-encoding
      The encoding for the backend.

      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h  show help

   Advanced Options

   --premiumizeme-encoding value  The encoding for the backend. (default: "Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PREMIUMIZEME_ENCODING]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)

```
{% endcode %}
