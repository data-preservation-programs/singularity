# Manually register a piece (CAR file) with the dataset for deal making purpose

{% code fullWidth="true" %}
```
NAME:
   singularity dataset add-piece - Manually register a piece (CAR file) with the dataset for deal making purpose

USAGE:
   singularity dataset add-piece [command options] <dataset_name> <piece_cid> <piece_size>

DESCRIPTION:
   If you already have the CAR file:
     singularity dataset add-piece -p <path_to_car_file> <dataset_name> <piece_cid> <piece_size>

   If you don't have the CAR file but you know the RootCID:
     singularity dataset add-piece -r <root_cid> <dataset_name> <piece_cid> <piece_size>

   If you don't have either:
     singularity dataset add-piece -r <root_cid> <dataset_name> <piece_cid> <piece_size>
   However in this case, deals made will not have rootCID set correctly so it may not work well with retrieval testing.

OPTIONS:
   --file-path value, -p value  Path to the CAR file, used to determine the size of the file and root CID
   --root-cid value, -r value   Root CID of the CAR file, if not provided, will be determined by the CAR file header. Used to populate the label field of storage deal
   --help, -h                   show help
```
{% endcode %}
