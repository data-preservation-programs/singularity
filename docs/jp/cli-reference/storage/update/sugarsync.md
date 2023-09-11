# Sugarsync

{% code fullWidth="true" %}
```
NAME:
   singularity storage update sugarsync - Sugarsync

USAGE:
   singularity storage update sugarsync [コマンドオプション] <名前|ID>

DESCRIPTION:
   --app-id
      Sugarsync アプリID。
      
      rclone のものを使用する場合は空白のままにします。

   --access-key-id
      Sugarsync アクセスキーID。
      
      rclone のものを使用する場合は空白のままにします。

   --private-access-key
      Sugarsync プライベートアクセスキー。
      
      rclone のものを使用する場合は空白のままにします。

   --hard-delete
      true の場合、ファイルを永久に削除します。
      それ以外の場合は、削除されたファイルに配置します。

   --refresh-token
      Sugarsync リフレッシュトークン。
      
      通常空白のままにします。rclone によって自動的に設定されます。

   --authorization
      Sugarsync 認証情報。
      
      通常空白のままにします。rclone によって自動的に設定されます。

   --authorization-expiry
      Sugarsync 認証情報の有効期限。
      
      通常空白のままにします。rclone によって自動的に設定されます。

   --user
      Sugarsync ユーザー。
      
      通常空白のままにします。rclone によって自動的に設定されます。

   --root-id
      Sugarsync ルートID。
      
      通常空白のままにします。rclone によって自動的に設定されます。

   --deleted-id
      Sugarsync 削除済みフォルダのID。
      
      通常空白のままにします。rclone によって自動的に設定されます。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --access-key-id value       Sugarsync アクセスキーID。[$ACCESS_KEY_ID]
   --app-id value              Sugarsync アプリID。[$APP_ID]
   --hard-delete               true の場合、ファイルを永久に削除します (デフォルト: false) [$HARD_DELETE]
   --help, -h                  ヘルプを表示
   --private-access-key value  Sugarsync プライベートアクセスキー。[$PRIVATE_ACCESS_KEY]

   Advanced

   --authorization value         Sugarsync 認証情報。[$AUTHORIZATION]
   --authorization-expiry value  Sugarsync 認証情報の有効期限。[$AUTHORIZATION_EXPIRY]
   --deleted-id value            Sugarsync 削除済みフォルダのID。[$DELETED_ID]
   --encoding value              バックエンドのエンコーディング (デフォルト: "Slash,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --refresh-token value         Sugarsync リフレッシュトークン。[$REFRESH_TOKEN]
   --root-id value               Sugarsync ルートID。[$ROOT_ID]
   --user value                  Sugarsync ユーザー。[$USER]
```
{% endcode %}