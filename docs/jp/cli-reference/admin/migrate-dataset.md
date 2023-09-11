# 旧バージョンのSingularity MongoDBからデータセットを移行する

{% code fullWidth="true" %}
```
NAME:
   singularity admin migrate-dataset - 旧バージョンのSingularity MongoDBからデータセットを移行する

USAGE:
   singularity admin migrate-dataset [コマンドオプション] [引数...]

DESCRIPTION:
   Singularity V1からV2へのデータセットの移行を行います。以下の手順が含まれます：
     1. ソースストレージと出力ストレージを作成し、それらをV2のデータプリップにアタッチします。
     2. 新しいデータセットにすべてのフォルダの構造とファイルを作成します。
   注意事項：
     1. 作成されたデータプリップは新しいデータセットのワーカーと互換性がありません。
        したがって、データプリップを再開したり、移行されたデータセットに新しいファイルを追加したりしないでください。
        問題なくデータセットを取引したり、閲覧することはできます。
     2. フォルダのCIDは、複雑さのために生成されたり移行されたりしません。

OPTIONS:
   --mongo-connection-string value  MongoDBの接続文字列 (デフォルト: "mongodb://localhost:27017") [$MONGO_CONNECTION_STRING]
   --skip-files                     ファイルとフォルダの詳細情報の移行をスキップします。これにより移行が高速化されます。取引のみを行いたい場合に便利です。 (デフォルト: false)
   --help, -h                       ヘルプを表示します
```
{% endcode %}