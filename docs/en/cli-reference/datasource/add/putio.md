# Put.io

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add putio - Put.io

USAGE:
   singularity datasource add putio [command options] <dataset_name> <source_path>

DESCRIPTION:
   --putio-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for putio

   --putio-encoding value  The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PUTIO_ENCODING]

```
{% endcode %}
