# データセットにピース（CARファイル）を手動で登録し、取引を行う目的のために

{% code fullWidth="true" %}
```
NAME:
   singularity dataset add-piece - データセットにピース（CARファイル）を手動で登録し、取引を行う目的のために

USAGE:
   singularity dataset add-piece [command options] <dataset_name> <piece_cid> <piece_size>

DESCRIPTION:
   すでにCARファイルをお持ちの場合：
     singularity dataset add-piece -p <path_to_car_file> <dataset_name> <piece_cid> <piece_size>

   CARファイルを持っていないがRootCIDは知っている場合：
     singularity dataset add-piece -r <root_cid> <dataset_name> <piece_cid> <piece_size>

   どちらも持っていない場合：
     singularity dataset add-piece -r <root_cid> <dataset_name> <piece_cid> <piece_size>
   ただし、この場合、作成された取引のrootCIDが正しく設定されないため、取り出しのテストとの互換性が損なわれる可能性があります。

OPTIONS:
   --file-path value, -p value  CARファイルへのパス。ファイルのサイズとRootCIDを決定するために使用されます
   --root-cid value, -r value   CARファイルのRootCID。指定されていない場合、CARファイルのヘッダーから取得します。ストレージ取引のラベルフィールドに使用されます
   --help, -h                   ヘルプを表示します
```
{% endcode %}