# データセットの管理

{% code fullWidth="true" %}
```
NAME:
   singularity dataset - データセットの管理

USAGE:
   singularity dataset command [command options] [arguments...]

COMMANDS:
   create         新しいデータセットを作成します
   list           すべてのデータセットを表示します
   update         既存のデータセットを更新します
   remove         特定のデータセットを削除します。CAR ファイルは削除されません。
   add-wallet     データセットにウォレットを関連付けます。ウォレットは `singularity wallet import` コマンドを使用して事前にインポートする必要があります。
   list-wallet    データセットに関連付けられたすべてのウォレットを表示します
   remove-wallet  データセットから関連付けられたウォレットを削除します
   add-piece      データセットにピース（CAR ファイル）を手動で登録し、ディールのための目的で使用します
   list-pieces    データセットのディールのために使用可能なすべてのピースを表示します
   help, h        コマンドのリストを表示するか、特定のコマンドのヘルプを表示します

OPTIONS:
   --help, -h  ヘルプを表示する
```
{% endcode %}