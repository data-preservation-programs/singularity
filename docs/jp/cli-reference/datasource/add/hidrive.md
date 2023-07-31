# HiDrive

{% code fullWidth="true" %}
```
名前:
   singularity datasource add hidrive - HiDrive

使用法:
   singularity datasource add hidrive [command options] <dataset_name> <source_path>

説明:
   --hidrive-auth-url
      認証サーバーのURLです。

      プロバイダーのデフォルトを使用する場合は空白のままにしてください。

   --hidrive-chunk-size
      チャンクアップロードのチャンクサイズです。

      構成済みのカットオフサイズより大きい（または不明なサイズの）ファイルは、このサイズのチャンクでアップロードされます。

      これには最大2147483647バイト（約2.000Gi）の上限があります。
      これは、単一のアップロード操作でサポートされるバイト数の最大値です。
      上限を超えた値に設定したり、負の値に設定すると、アップロードが失敗します。

      この値を大きな値に設定すると、アップロードの速度が向上する一方、より多くのメモリを使用することになります。

      メモリを節約するために、この値を小さな値に設定することもできます。

   --hidrive-client-id
      OAuthクライアントIDです。

      通常は空白のままにしてください。

   --hidrive-client-secret
      OAuthクライアントシークレットです。

      通常は空白のままにしてください。

   --hidrive-disable-fetching-member-count
      ディレクトリ内のオブジェクトの数を必要な場合以外は取得しないでください。

      サブディレクトリ内のオブジェクトの数を取得しないと、リクエストの処理が高速化される場合があります。

   --hidrive-encoding
      バックエンドのエンコーディングです。

      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --hidrive-endpoint
      サービスのエンドポイントです。

      これはAPI呼び出しが行われるURLです。

   --hidrive-root-prefix
      すべてのパスのルート/親フォルダです。

      すべてのパスに指定されたフォルダを親として使用するように指定します。
      これにより、rcloneは任意のフォルダを起点として使用できます。

      例:
         | /       | rcloneでアクセス可能な最上位のディレクトリです。
                   | rcloneが通常のHiDriveユーザーアカウントを使用する場合、これは「root」と同じです。
         | root    | HiDriveユーザーアカウントの最上位ディレクトリ
         | <unset> | パスのルートプレフィックスがないことを指定します。
                   | これを使用する場合は、常に有効な親（例えば「remote:/path/to/dir」または「remote:root/path/to/dir」）を指定する必要があります。

   --hidrive-scope-access
      HiDriveからアクセスを要求する際に、rcloneが使用するアクセス許可です。

      例:
         | rw | リソースへの読み書きアクセス
         | ro | リソースへの読み取り専用アクセス

   --hidrive-scope-role
      HiDriveからアクセスを要求する際にrcloneが使用するユーザーレベルです。

      例:
         | user  | 管理権限へのユーザーレベルアクセス
                 | これはほとんどの場合で十分です。
         | admin | 管理権限への広範なアクセス
         | owner | 管理権限への完全なアクセス

   --hidrive-token
      OAuthアクセストークン（JSON形式）です。

   --hidrive-token-url
      トークンサーバーのURLです。

      プロバイダーのデフォルトを使用する場合は空白のままにしてください。

   --hidrive-upload-concurrency
      チャンクアップロードの同時実行数です。

      同じファイルに対する転送が同時に実行される最大数の上限です。
      1未満の値に設定すると、アップロードがデッドロックします。

      高速なリンクを介して大量の大きなファイルをアップロードし、これらのアップロードが帯域幅を十分に利用していない場合は、
      これを増やすことで転送を高速化することができます。

   --hidrive-upload-cutoff
      チャンクアップロードのカットオフ/閾値です。

      これより大きいファイルは、指定されたチャンクサイズのチャンクでアップロードされます。

      これには最大2147483647バイト（約2.000Gi）の上限があります。
      これは、単一のアップロード操作でサポートされるバイト数の最大値です。
      上限を超えた値に設定すると、アップロードが失敗します。


オプション:
   --help, -h  ヘルプを表示します

   データ準備オプション

   --delete-after-export    [危険] CARファイルにエクスポートしたデータセットのファイルを削除します。  (デフォルト: false)
   --rescan-interval value  最後の正常なスキャンからこの間隔が経過した場合、ソースディレクトリを自動的に再スキャンします。 (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します。 (デフォルト: ready)

   HiDrive用のオプション

   --hidrive-auth-url value                       認証サーバーのURLです。 [$HIDRIVE_AUTH_URL]
   --hidrive-chunk-size value                     チャンクアップロードのチャンクサイズです。 (デフォルト: "48Mi") [$HIDRIVE_CHUNK_SIZE]
   --hidrive-client-id value                      OAuthクライアントIDです。 [$HIDRIVE_CLIENT_ID]
   --hidrive-client-secret value                  OAuthクライアントシークレットです。 [$HIDRIVE_CLIENT_SECRET]
   --hidrive-disable-fetching-member-count value  ディレクトリ内のオブジェクトの数を必要な場合以外は取得しないでください。 (デフォルト: "false") [$HIDRIVE_DISABLE_FETCHING_MEMBER_COUNT]
   --hidrive-encoding value                       バックエンドのエンコーディングです。 (デフォルト: "Slash,Dot") [$HIDRIVE_ENCODING]
   --hidrive-endpoint value                       サービスのエンドポイントです。 (デフォルト: "https://api.hidrive.strato.com/2.1") [$HIDRIVE_ENDPOINT]
   --hidrive-root-prefix value                    すべてのパスのルート/親フォルダです。 (デフォルト: "/") [$HIDRIVE_ROOT_PREFIX]
   --hidrive-scope-access value                   HiDriveからアクセスを要求する際に、rcloneが使用するアクセス許可です。 (デフォルト: "rw") [$HIDRIVE_SCOPE_ACCESS]
   --hidrive-scope-role value                     HiDriveからアクセスを要求する際に、rcloneが使用するユーザーレベルです。 (デフォルト: "user") [$HIDRIVE_SCOPE_ROLE]
   --hidrive-token value                          OAuthアクセストークン（JSON形式）です。 [$HIDRIVE_TOKEN]
   --hidrive-token-url value                      トークンサーバーのURLです。 [$HIDRIVE_TOKEN_URL]
   --hidrive-upload-concurrency value             チャンクアップロードの同時実行数です。 (デフォルト: "4") [$HIDRIVE_UPLOAD_CONCURRENCY]
   --hidrive-upload-cutoff value                  チャンクアップロードのカットオフ/閾値です。 (デフォルト: "96Mi") [$HIDRIVE_UPLOAD_CUTOFF]

```
{% endcode %}