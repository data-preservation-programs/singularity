# Akamai NetStorage

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add netstorage - Akamai NetStorage

USAGE:
   singularity datasource add netstorage [command options] <dataset_name> <source_path>

DESCRIPTION:
   --netstorage-account
      NetStorageアカウント名を設定します。

   --netstorage-host
      接続するNetStorageホストのドメイン+パスを指定します。

      フォーマットは `<ドメイン>/<内部フォルダ>` となります。

   --netstorage-protocol
      HTTPまたはHTTPSプロトコルを選択します。

      多くのユーザーはデフォルトのHTTPSを選択するべきですが、HTTPは主にデバッグ目的で提供されています。

      例:
         | http  | HTTPプロトコル
         | https | HTTPSプロトコル

   --netstorage-secret
      認証のためにNetStorageアカウントのシークレット/ G2Oキーを設定します。

      パスワード設定するために 'y' オプションを選択し、シークレットを入力してください。


OPTIONS:
   --help, -h  ヘルプを表示します。

   データ準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、ファイルを削除します。 (デフォルト: false)
   --rescan-interval value  最後の成功したスキャンからの経過時間がこの間隔を超えると、自動的にソースディレクトリが再スキャンされます。 (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します。 (デフォルト: ready)

   NetStorage用のオプション

   --netstorage-account value   NetStorageアカウント名を設定します。 [$NETSTORAGE_ACCOUNT]
   --netstorage-host value      接続するNetStorageホストのドメイン+パスを指定します。 [$NETSTORAGE_HOST]
   --netstorage-protocol value  HTTPまたはHTTPSプロトコルを選択します。 (デフォルト: "https") [$NETSTORAGE_PROTOCOL]
   --netstorage-secret value    認証のためにNetStorageアカウントのシークレット/ G2Oキーを設定します。 [$NETSTORAGE_SECRET]
```
{% endcode %}