# Akamai NetStorage

{% code fullWidth="true" %}
```
NAME:
   singularity storage create netstorage - Akamai NetStorage

USAGE:
   singularity storage create netstorage [command options] [arguments...]

DESCRIPTION:
   --protocol
      HTTPまたはHTTPSプロトコルを選択します。
      
      ほとんどのユーザはHTTPSを選択する必要があります（デフォルトです）。
      HTTPは主にデバッグ目的で提供されています。

      例:
         | http  | HTTPプロトコル
         | https | HTTPSプロトコル

   --host
      接続するNetStorageホストのドメイン+パスです。
      
      形式は `<ドメイン>/<内部フォルダ>` としてください。

   --account
      NetStorageアカウント名を設定します。

   --secret
      認証のためのNetStorageアカウントのシークレット/G2Oキーを設定します。
      
      シークレットを設定するために、'y'オプションを選択し、シークレットを入力してください。


OPTIONS:
   --account value  NetStorageアカウント名を設定します [$ACCOUNT]
   --help, -h       ヘルプを表示します
   --host value     接続するNetStorageホストのドメイン+パスです [$HOST]
   --secret value   認証のためのNetStorageアカウントのシークレット/G2Oキーを設定します [$SECRET]

   Advanced

   --protocol value  HTTPまたはHTTPSプロトコルを選択します（デフォルト: "https"） [$PROTOCOL]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}