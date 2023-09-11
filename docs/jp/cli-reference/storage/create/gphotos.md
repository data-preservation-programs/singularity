# Google フォト

{% code fullWidth="true" %}
```
NAME:
   singularity storage create gphotos - Google フォト

USAGE:
   singularity storage create gphotos [command options] [arguments...]

DESCRIPTION:
   --client-id
      OAuth Client Id.
      
      通常は空欄のままにしてください。

   --client-secret
      OAuth Client Secret.
      
      通常は空欄のままにしてください。

   --token
      JSON ブロブ形式の OAuth アクセストークン。

   --auth-url
      認証サーバーの URL。
      
      デフォルトのプロバイダー設定を使用する場合は空欄のままにしてください。

   --token-url
      トークンサーバーの URL。
      
      デフォルトのプロバイダー設定を使用する場合は空欄のままにしてください。

   --read-only
      Google フォトバックエンドを読み取り専用に設定します。
      
      読み取り専用に設定すると、rclone は写真への読み取り専用アクセスのみ要求しますが、
      読み取り専用にしない場合はフルアクセスを要求します。

   --read-size
      メディアアイテムのサイズを読み取るように設定します。
      
      通常、rclone はメディアアイテムのサイズを読み取りませんが、これには別のトランザクションが必要です。
      同期にはこれは必要ありませんが、rclone mount を使用する場合は、
      メディアのサイズを読み取るためにこのフラグを設定することをお勧めします。

   --start-year
      ダウンロードする写真の年を指定すると、指定した年以降にアップロードされた写真のみがダウンロードされます。

   --include-archived
      アーカイブされたメディアも表示してダウンロードします。
      
      デフォルトでは、rclone はアーカイブされたメディアをリクエストしません。
      そのため、同期時にアーカイブされたメディアはディレクトリリストや転送に表示されません。
      
      アルバム内のメディアは、アーカイブの状態にかかわらず常に表示および同期されます。
      
      このフラグを指定すると、ディレクトリリストや転送時にアーカイブされたメディアが常に表示されます。
      
      このフラグを指定しないと、ディレクトリリストにアーカイブされたメディアは表示されず、転送もされません。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --client-id value      OAuth Client Id. [$CLIENT_ID]
   --client-secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             ヘルプを表示
   --read-only            Google フォトバックエンドを読み取り専用に設定します。 (デフォルト: false) [$READ_ONLY]

   Advanced

   --auth-url value    認証サーバーの URL。 [$AUTH_URL]
   --encoding value    バックエンドのエンコーディング。 (デフォルト: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --include-archived  アーカイブされたメディアも表示してダウンロードします。 (デフォルト: false) [$INCLUDE_ARCHIVED]
   --read-size         メディアアイテムのサイズを読み取るように設定します。 (デフォルト: false) [$READ_SIZE]
   --start-year value  ダウンロードする写真の年を指定すると、指定した年以降にアップロードされた写真のみがダウンロードされます。 (デフォルト: 2000) [$START_YEAR]
   --token value       JSON ブロブ形式の OAuth アクセストークン。 [$TOKEN]
   --token-url value   トークンサーバーの URL。 [$TOKEN_URL]

   General

   --name value  ストレージの名前 (デフォルト: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}