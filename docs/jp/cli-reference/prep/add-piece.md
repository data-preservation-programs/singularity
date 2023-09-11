# プリパレーションにピース情報を手動で追加する。これは外部ツールによって準備されたピースに便利です。

{% code fullWidth="true" %}
```
NAME:
   singularity prep add-piece - プリパレーションにピース情報を手動で追加する。これは外部ツールによって準備されたピースに便利です。

USAGE:
   singularity prep add-piece [コマンドオプション] <プリパレーションIDまたは名前>

CATEGORY:
   ピース管理

OPTIONS:
   --piece-cid value   ピースのCID
   --piece-size value  ピースのサイズ（デフォルト: "32GiB"）
   --file-path value   CARファイルのパス。ファイルサイズとルートCIDを決定するために使用されます。
   --root-cid value    CARファイルのルートCID
   --help, -h          ヘルプを表示
```
{% endcode %}