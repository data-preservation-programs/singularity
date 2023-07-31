# Google Photos

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add gphotos - Google フォト

USAGE:
   singularity datasource add gphotos [コマンドオプション] <データセット名> <ソースパス>

説明:
   --gphotos-auth-url
      認証サーバーのURL。
      
      プロバイダーデフォルトを使用するには空白にしておきます。

   --gphotos-client-id
      OAuth クライアント ID。
      
      通常は空白のままにしておきます。

   --gphotos-client-secret
      OAuth クライアントシークレット。
      
      通常は空白のままにしておきます。

   --gphotos-encoding
      バックエンドのエンコーディング。
      
      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --gphotos-include-archived
      アーカイブされたメディアも表示およびダウンロードします。
      
      デフォルトでは、rcloneはアーカイブされたメディアを要求しません。
      したがって、同期中にはディレクトリリストや転送中にアーカイブされたメディアは表示されません。
      
      アルバム内のメディアは、アーカイブの状態に関係なく常に表示および同期されます。
      
      このフラグを使用すると、アーカイブされたメディアは常にディレクトリリストに表示され、転送されます。
      
      このフラグを使用しない場合、アーカイブされたメディアはディレクトリリストに表示されず、転送されません。

   --gphotos-read-only
      Google フォトバックエンドを読み取り専用に設定します。
      
      読み取り専用を選択すると、rcloneは写真の読み取り専用アクセスしか要求しません。
      それ以外の場合、rcloneは完全なアクセスを要求します。

   --gphotos-read-size
      メディアアイテムのサイズを読み取るように設定します。
      
      通常、rcloneはメディアアイテムのサイズを読み取りません。これには別のトランザクションが必要です。
      これは同期には必要ありません。ただし、rclone mountは読み取る前にファイルのサイズを事前に知る必要があるため、
      rclone mountを使用する場合はこのフラグを設定することをおすすめします。

   --gphotos-start-year
      指定された年以降にアップロードされた写真にダウンロードを制限します。

   --gphotos-token
      JSON ブロブとしての OAuth アクセストークン。

   --gphotos-token-url
      トークンサーバーの URL。
      
      プロバイダーデフォルトを使用するには空白にしておきます。


オプション:
   --help, -h  ヘルプを表示します

   データの準備オプション

   --delete-after-export    [危険] データセットのファイルをエクスポート後に削除します。 (デフォルト: false)
   --rescan-interval value  最後の成功したスキャンからこの間隔が経過した場合に自動的にソースディレクトリを再スキャンします (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します (デフォルト: ready)

   gphotos用オプション

   --gphotos-auth-url value          認証サーバーのURL。 [$GPHOTOS_AUTH_URL]
   --gphotos-client-id value         OAuth クライアント ID。 [$GPHOTOS_CLIENT_ID]
   --gphotos-client-secret value     OAuth クライアントシークレット。 [$GPHOTOS_CLIENT_SECRET]
   --gphotos-encoding value          バックエンドのエンコーディング。 (デフォルト: "Slash,CrLf,InvalidUtf8,Dot") [$GPHOTOS_ENCODING]
   --gphotos-include-archived value  アーカイブされたメディアも表示およびダウンロードします。 (デフォルト: "false") [$GPHOTOS_INCLUDE_ARCHIVED]
   --gphotos-read-only value         Google フォトバックエンドを読み取り専用に設定します。 (デフォルト: "false") [$GPHOTOS_READ_ONLY]
   --gphotos-read-size value         メディアアイテムのサイズを読み取ります。 (デフォルト: "false") [$GPHOTOS_READ_SIZE]
   --gphotos-start-year value        指定された年以降にアップロードされた写真にダウンロードを制限します。 (デフォルト: "2000") [$GPHOTOS_START_YEAR]
   --gphotos-token value             JSON ブロブとしての OAuth アクセストークン。 [$GPHOTOS_TOKEN]
   --gphotos-token-url value         トークンサーバーの URL。 [$GPHOTOS_TOKEN_URL]
```
{% endcode %}