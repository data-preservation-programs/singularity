# HiDrive

{% code fullWidth="true" %}
```
NAME:
   singularity storage update hidrive - HiDrive

USAGE:
   singularity storage update hidrive [command options] <name|id>

DESCRIPTION:
   --client-id
      OAuth Client Id.
      
      通常、空白のままにしてください。

   --client-secret
      OAuth Client Secret.
      
      通常、空白のままにしてください。

   --token
      OAuth Access TokenをJSON形式の情報として指定してください。

   --auth-url
      認証サーバーのURL。
      
      プロバイダのデフォルト値を使いたい場合は、空白のままにしてください。

   --token-url
      トークンサーバーのURL。
      
      プロバイダのデフォルト値を使いたい場合は、空白のままにしてください。

   --scope-access
      HiDriveへのアクセス権限を指定します。

      例:
         | rw | リソースへの読み書きアクセス権限。
         | ro | リソースへの読み取り専用アクセス権限。

   --scope-role
      HiDriveへのアクセス時のユーザーレベルを指定します。

      例:
         | user  | 管理権限へのユーザーレベルアクセス。
         |       | 通常の場合これで十分です。
         | admin | 管理権限への広範なアクセス。
         | owner | 管理権限への完全なアクセス。

   --root-prefix
      すべてのパスのルート/親フォルダを指定します。

      特定のフォルダをすべてのパスの親に使用する場合は、ここに記入してください。
      これにより、rcloneは任意のフォルダを起点として使用できるようになります。

      例:
         | /       | rcloneでアクセスできる最上位のディレクトリ。
         |         | rcloneが通常のHiDriveユーザーアカウントを使用する場合、これは"root"と同等です。
         | root    | HiDriveユーザーアカウントの最上位ディレクトリ
         | <unset> | あなたのパスにルートプレフィックスがないことを指定します。
         |         | これを使用する場合、常に有効な親としてパスを指定する必要があります。例: "remote:/path/to/dir" または "remote:root/path/to/dir"。

   --endpoint
      サービスのエンドポイントURL。

      APIの呼び出しに使用されるURLです。

   --disable-fetching-member-count
      ディレクトリ内のオブジェクトの数を最低限必要な場合を除き、取得しないようにします。

      サブディレクトリ内のオブジェクトの数を取得しない場合、リクエストが高速になる場合があります。

   --chunk-size
      チャンクアップロードのチャンクサイズ。

      構成済みのカットオフより大きいファイル（またはサイズ不明のファイル）は、このサイズのチャンクでアップロードされます。

      この値の上限は2147483647バイト（約2.0Giバイト）です。
      これは、1つのアップロード操作でサポートされる最大バイト数です。
      この上限を超える値に設定した場合、アップロードは失敗します。

      この値を大きくすると、アップロード速度が速くなりますが、メモリをより多く使用することになります。

      メモリを節約するために、この値を小さく設定することもできます。

   --upload-cutoff
      チャンクアップロードの閾値。

      この値より大きいファイルは、構成済みのチャンクサイズのチャンクでアップロードされます。

      この値の上限は2147483647バイト（約2.0Giバイト）です。
      これは、1つのアップロード操作でサポートされる最大バイト数です。
      この上限を超える値に設定すると、アップロードが失敗します。

   --upload-concurrency
      チャンクアップロードの並行性。

      同じファイルに対して実行される転送の最大数です。
      1より小さい値に設定すると、アップロードがデッドロックになります。

      ハイスピードリンクを介して少数の大きなファイルをアップロードし、
      これらのアップロードが帯域幅を十分に利用しない場合は、この値を増やすと転送速度が向上するかもしれません。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --client-id value      OAuth Client Id. [$CLIENT_ID]
   --client-secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             ヘルプを表示
   --scope-access value   HiDriveへのアクセス権限を指定します。 (default: "rw") [$SCOPE_ACCESS]

   Advanced

   --auth-url value                 認証サーバーのURL。 [$AUTH_URL]
   --chunk-size value               チャンクアップロードのチャンクサイズ。 (default: "48Mi") [$CHUNK_SIZE]
   --disable-fetching-member-count  ディレクトリ内のオブジェクトの数を最低限必要な場合を除き、取得しないようにします。 (default: false) [$DISABLE_FETCHING_MEMBER_COUNT]
   --encoding value                 バックエンドのエンコーディング。 (default: "Slash,Dot") [$ENCODING]
   --endpoint value                 サービスのエンドポイントURL。 (default: "https://api.hidrive.strato.com/2.1") [$ENDPOINT]
   --root-prefix value              すべてのパスのルート/親フォルダを指定します。 (default: "/") [$ROOT_PREFIX]
   --scope-role value               HiDriveへのアクセス時のユーザーレベルを指定します。 (default: "user") [$SCOPE_ROLE]
   --token value                    OAuth Access TokenをJSON形式の情報として指定してください。 [$TOKEN]
   --token-url value                トークンサーバーのURL。 [$TOKEN_URL]
   --upload-concurrency value       チャンクアップロードの並行性。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードの閾値。 (default: "96Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}