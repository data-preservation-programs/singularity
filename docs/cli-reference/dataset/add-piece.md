# Manually register a piece (CAR file) with the dataset for deal making purpose

{% code fullWidth="true" %}
```
NAME:
   singularity dataset add-piece - Manually register a piece (CAR file) with the dataset for deal making purpose

USAGE:
   singularity dataset add-piece [command options] <dataset_name> <piece_cid> <piece_size>

OPTIONS:
   --file-path value, -p value  Path to the CAR file, used to determine the size of the file and root CID
   --file-size value, -s value  Size of the CAR file, if not provided, will be determined by the CAR file (default: 0)
   --root-cid value, -r value   Root CID of the CAR file, if not provided, will be determined by the CAR file header. Used to populate the label field of storage deal
   --help, -h                   show help
```
{% endcode %}
