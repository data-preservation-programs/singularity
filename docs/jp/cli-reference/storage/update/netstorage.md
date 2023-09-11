# Akamai NetStorage

{% code fullWidth="true" %}
```
名称:
   singularity storage update netstorage - Akamai NetStorage

使用方法:
   singularity storage update netstorage [コマンドオプション] <名前|ID>

説明:
   --protocol
      HTTPまたはHTTPSプロトコルを選択します。

      ほとんどのユーザーはデフォルトのHTTPSを選択するべきです。
      HTTPは主にデバッグ目的で提供されます。

      例:
         | http  | HTTPプロトコル
         | https | HTTPSプロトコル

   --host
      接続するNetStorageホストのドメイン+パスを設定します。

      フォーマットは `<ドメイン>/<内部フォルダ>` の形式にする必要があります。

   --account
      NetStorageアカウント名を設定します。

   --secret
      認証のためのNetStorageアカウントの秘密キー/ G2Oキーを設定します。

      秘密キー設定の場合、'y'オプションを選択してから秘密キーを入力してください。


オプション:
   --account value  NetStorageアカウント名を設定します [$ACCOUNT]
   --help, -h       ヘルプを表示します
   --host value     接続するNetStorageホストのドメイン+パスを設定します [$HOST]
   --secret value   認証のためのNetStorageアカウントの秘密キー/ G2Oキーを設定します [$SECRET]

   高度なオプション:

   --protocol value  HTTPまたはHTTPSプロトコルを選択します。 (デフォルト: "https") [$PROTOCOL]
```
{% endcode %}