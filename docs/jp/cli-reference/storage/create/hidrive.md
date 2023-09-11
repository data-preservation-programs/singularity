# HiDrive

{% code fullWidth="true" %}
```
名前:
   singularity storage create hidrive - HiDrive

使用法:
   singularity storage create hidrive [コマンドオプション] [引数...]

説明:
   --client-id
      OAuthクライアントID。
      
      通常は空白のままにしてください。

   --client-secret
      OAuthクライアントシークレット。
      
      通常は空白のままにしてください。

   --token
      JSON形式のOAuthアクセストークン。

   --auth-url
      認証サーバーのURL。
      
      デフォルトのプロバイダーを使用するために空白のままにしてください。

   --token-url
      トークンサーバーのURL。
      
      デフォルトのプロバイダーを使用するために空白のままにしてください。

   --scope-access
      HiDriveへのアクセス権限。
      
      例:
         | rw | 読み書きアクセス権限。
         | ro | 読み取り専用アクセス権限。

   --scope-role
      HiDriveへのアクセスパーミッションのユーザーレベル。
      
      例:
         | user  | 管理権限へのユーザーレベルアクセス。
         |       | ほとんどの場合、これで十分です。
         | admin | 広範な管理権限へのアクセス。
         | owner | 全ての管理権限へのフルアクセス。

   --root-prefix
      すべてのパスのルート/親フォルダです。
      
      指定したフォルダを、リモートに与えられたすべてのパスの親として使用するために入力します。
      これにより、rcloneは任意のフォルダを開始点として使用できます。

      例:
         | /       | rcloneからアクセス可能な一番上のディレクトリ。
         |         | 通常、rcloneが通常のHiDriveユーザーアカウントを使用している場合、これは「root」と同等です。
         | root    | HiDriveユーザーアカウントの一番上のディレクトリ
         | <unset> | パスに対してルートプレフィックスがないことを指定します。
         |         | この場合、常に有効な親を持つパスを指定する必要があります。 例: "remote:/path/to/dir" または "remote:root/path/to/dir"。

   --endpoint
      サービスのエンドポイント。
      
      これはAPIコールが行われるURLです。

   --disable-fetching-member-count
      必要な場合以外はディレクトリ内のオブジェクト数を取得しないようにします。
      
      サブディレクトリ内のオブジェクト数を取得しない方がリクエストが速くなる場合があります。

   --chunk-size
      チャンクアップロードのチャンクサイズです。
      
      設定されたカットオフよりも大きいファイル（または不明なサイズのファイル）は、このサイズのチャンクでアップロードされます。
      
      この値の上限は2147483647バイト（約2.000Gi）です。
      これは単一のアップロード操作がサポートすることができるバイト数の最大値です。
      この値を上限を超えたり、負の値に設定すると、アップロードが失敗します。
      
      この値を大きく設定すると、メモリをより多く使用する代わりにアップロード速度が向上する場合があります。
      メモリを節約するために、この値を小さく設定することもできます。

   --upload-cutoff
      チャンクアップロードのカットオフ/閾値です。
      
      この値よりも大きいファイルは、設定されたチャンクサイズのチャンクでアップロードされます。
      
      この値の上限は2147483647バイト（約2.000Gi）です。
      これは単一のアップロード操作がサポートすることができるバイト数の最大値です。
      この値を上限を超えると、アップロードが失敗します。

   --upload-concurrency
      チャンクアップロードの同時実行数です。
      
      同じファイルに対する転送が同時に実行される上限です。
      1よりも小さい値に設定すると、アップロードがデッドロックします。
      
      高速リンク上で小数の大きなファイルをアップロードし、これらのアップロードが帯域幅を十分に利用していない場合、
      これを増やすことで転送速度を向上させることができます。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --client-id value      OAuthクライアントID。 [$CLIENT_ID]
   --client-secret value  OAuthクライアントシークレット。 [$CLIENT_SECRET]
   --help, -h             ヘルプを表示
   --scope-access value   HiDriveのアクセス権限。 (デフォルト: "rw") [$SCOPE_ACCESS]

   高度なオプション

   --auth-url value                 認証サーバーのURL。 [$AUTH_URL]
   --chunk-size value               チャンクアップロードのチャンクサイズ。 (デフォルト: "48Mi") [$CHUNK_SIZE]
   --disable-fetching-member-count  必要な場合以外はディレクトリ内のオブジェクト数を取得しないようにします。 (デフォルト: false) [$DISABLE_FETCHING_MEMBER_COUNT]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト: "Slash,Dot") [$ENCODING]
   --endpoint value                 サービスのエンドポイント。 (デフォルト: "https://api.hidrive.strato.com/2.1") [$ENDPOINT]
   --root-prefix value              すべてのパスのルート/親フォルダ。 (デフォルト: "/") [$ROOT_PREFIX]
   --scope-role value               HiDriveへのアクセスパーミッションのユーザーレベル。 (デフォルト: "user") [$SCOPE_ROLE]
   --token value                    JSON形式のOAuthアクセストークン。 [$TOKEN]
   --token-url value                トークンサーバーのURL。 [$TOKEN_URL]
   --upload-concurrency value       チャンクアップロードの同時実行数。 (デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードのカットオフ/閾値。 (デフォルト: "96Mi") [$UPLOAD_CUTOFF]

   一般

   --name value  ストレージの名前 (デフォルト: 自動生成される)
   --path value  ストレージのパス

```
{% endcode %}