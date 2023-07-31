# Enterprise File Fabric

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add filefabric - Enterprise File Fabric

USAGE:
   singularity datasource add filefabric [command options] <dataset_name> <source_path>

DESCRIPTION:
   --filefabric-encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --filefabric-permanent-token
      永続的な認証トークンです。
      
      永続的な認証トークンは、Enterprise File Fabricで作成できます。
      ユーザーダッシュボードのセキュリティの下に表示される、「マイ認証トークン」というエントリがあります。
      作成するには、管理ボタンをクリックしてください。
      
      これらのトークンは通常数年間有効です。
      
      詳細についてはこちらを参照してください：https://docs.storagemadeeasy.com/organisationcloud/api-tokens

   --filefabric-root-folder-id
      ルートフォルダのIDです。
      
      通常は空白のままにしておきます。
      
      特定のIDのディレクトリからrcloneを開始するには、それを入力してください。

   --filefabric-token
      セッショントークンです。
      
      rcloneはこのセッショントークンを設定ファイルにキャッシュします。
      通常は1時間有効です。
      
      この値を設定しないでください - rcloneは自動的に設定します。

   --filefabric-token-expiry
      トークンの有効期限です。
      
      この値を設定しないでください - rcloneは自動的に設定します。

   --filefabric-url
      接続するEnterprise File FabricのURLです。

      例：
         | https://storagemadeeasy.com       | Storage Made Easy US
         | https://eu.storagemadeeasy.com    | Storage Made Easy EU
         | https://yourfabric.smestorage.com | 自分のEnterprise File Fabricに接続する

   --filefabric-version
      ファイルファブリックから読み取られるバージョンです。
      
      この値を設定しないでください - rcloneは自動的に設定します。

OPTIONS:
   --help, -h  ヘルプを表示します。

   データの準備オプション

   --delete-after-export    [危険] データセットのファイルをエクスポート後に削除します。 (デフォルト: false)
   --rescan-interval value  前回のスキャンから指定された間隔が経過した場合に、ソースディレクトリを自動的に再スキャンします。 (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します。 (デフォルト: ready)

   filefabric用オプション

   --filefabric-encoding value         バックエンドのエンコーディングです。 (デフォルト: "Slash,Del,Ctl,InvalidUtf8,Dot") [$FILEFABRIC_ENCODING]
   --filefabric-permanent-token value  永続的な認証トークンです。 [$FILEFABRIC_PERMANENT_TOKEN]
   --filefabric-root-folder-id value   ルートフォルダのIDです。 [$FILEFABRIC_ROOT_FOLDER_ID]
   --filefabric-token value            セッショントークンです。 [$FILEFABRIC_TOKEN]
   --filefabric-token-expiry value     トークンの有効期限です。 [$FILEFABRIC_TOKEN_EXPIRY]
   --filefabric-url value              接続するEnterprise File FabricのURLです。 [$FILEFABRIC_URL]
   --filefabric-version value          ファイルファブリックから読み取られるバージョンです。 [$FILEFABRIC_VERSION]
```
{% endcode %}