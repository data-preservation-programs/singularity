# Sugarsync

{% code fullWidth="true" %}
```
名前:
   singularityデータソースの追加 sugarsync - Sugarsync

使用法:
   singularityデータソースの追加 sugarsync [コマンドオプション] <データセット名> <ソースパス>

説明:
   --sugarsync-access-key-id
      SugarsyncのアクセスキーIDです。
      
      空欄の場合はrcloneのものを使用します。

   --sugarsync-app-id
      SugarsyncのアプリIDです。
      
      空欄の場合はrcloneのものを使用します。

   --sugarsync-authorization
      Sugarsyncの認証です。
      
      通常は空欄で、rcloneによって自動設定されます。

   --sugarsync-authorization-expiry
      Sugarsyncの認証の有効期限です。
      
      通常は空欄で、rcloneによって自動設定されます。

   --sugarsync-deleted-id
      Sugarsyncの削除済みフォルダIDです。
      
      通常は空欄で、rcloneによって自動設定されます。

   --sugarsync-encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --sugarsync-hard-delete
      trueの場合、ファイルを永久に削除します。
      そうでなければ、削除ファイルに配置されます。

   --sugarsync-private-access-key
      Sugarsyncのプライベートアクセスキーです。
      
      空欄の場合はrcloneのものを使用します。

   --sugarsync-refresh-token
      Sugarsyncのリフレッシュトークンです。
      
      通常は空欄で、rcloneによって自動設定されます。

   --sugarsync-root-id
      SugarsyncのルートIDです。
      
      通常は空欄で、rcloneによって自動設定されます。

   --sugarsync-user
      Sugarsyncのユーザーです。
      
      通常は空欄で、rcloneによって自動設定されます。


オプション:
   --help, -h  ヘルプを表示する

   データの準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、データのファイルを削除します。 (デフォルト: false)
   --rescan-interval value  前回のスキャンからこの間隔が経過した場合、自動的にソースディレクトリを再スキャンします (デフォルト: 無効)
   --scanning-state value   初期スキャン状態を設定します (デフォルト: ready)

   Sugarsync用オプション

   --sugarsync-access-key-id value         SugarsyncのアクセスキーIDです。 [$SUGARSYNC_ACCESS_KEY_ID]
   --sugarsync-app-id value                SugarsyncのアプリIDです。 [$SUGARSYNC_APP_ID]
   --sugarsync-authorization value         Sugarsyncの認証です。 [$SUGARSYNC_AUTHORIZATION]
   --sugarsync-authorization-expiry value  Sugarsyncの認証の有効期限です。 [$SUGARSYNC_AUTHORIZATION_EXPIRY]
   --sugarsync-deleted-id value            Sugarsyncの削除済みフォルダIDです。 [$SUGARSYNC_DELETED_ID]
   --sugarsync-encoding value              バックエンドのエンコーディングです。 (デフォルト: "Slash,Ctl,InvalidUtf8,Dot") [$SUGARSYNC_ENCODING]
   --sugarsync-hard-delete value           trueの場合、ファイルを永久に削除します。 (デフォルト: "false") [$SUGARSYNC_HARD_DELETE]
   --sugarsync-private-access-key value    Sugarsyncのプライベートアクセスキーです。 [$SUGARSYNC_PRIVATE_ACCESS_KEY]
   --sugarsync-refresh-token value         Sugarsyncのリフレッシュトークンです。 [$SUGARSYNC_REFRESH_TOKEN]
   --sugarsync-root-id value               SugarsyncのルートIDです。 [$SUGARSYNC_ROOT_ID]
   --sugarsync-user value                  Sugarsyncのユーザーです。 [$SUGARSYNC_USER]

```
{% endcode %}