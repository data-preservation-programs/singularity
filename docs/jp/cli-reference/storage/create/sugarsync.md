# Sugarsync

{% code fullWidth="true" %}
```
NAME:
   singularity storage create sugarsync - Sugarsync

USAGE:
   singularity storage create sugarsync [command options] [arguments...]

DESCRIPTION:
   --app-id
      Sugarsync アプリID。
      
      空白の場合はrcloneのものが使用されます。

   --access-key-id
      Sugarsync Access Key ID。
      
      空白の場合はrcloneのものが使用されます。

   --private-access-key
      Sugarsync Private Access Key。
      
      空白の場合はrcloneのものが使用されます。

   --hard-delete
      trueの場合、ファイルを完全に削除します。
      そうでない場合は削除フォルダに移動します。

   --refresh-token
      Sugarsync リフレッシュトークン。
      
      通常は空白のままにしておきます。rcloneによって自動設定されます。

   --authorization
      Sugarsync 認証。
      
      通常は空白のままにしておきます。rcloneによって自動設定されます。

   --authorization-expiry
      Sugarsync 認証有効期限。
      
      通常は空白のままにしておきます。rcloneによって自動設定されます。

   --user
      Sugarsync ユーザー。
      
      通常は空白のままにしておきます。rcloneによって自動設定されます。

   --root-id
      Sugarsync ルートID。
      
      通常は空白のままにしておきます。rcloneによって自動設定されます。

   --deleted-id
      Sugarsync 削除されたフォルダのID。
      
      通常は空白のままにしておきます。rcloneによって自動設定されます。

   --encoding
      バックエンドのエンコーディング。
      
      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --access-key-id value       Sugarsync Access Key ID。 [$ACCESS_KEY_ID]
   --app-id value              Sugarsync App ID。 [$APP_ID]
   --hard-delete               trueの場合、ファイルを完全に削除します（デフォルト: false） [$HARD_DELETE]
   --help, -h                  ヘルプを表示
   --private-access-key value  Sugarsync Private Access Key。 [$PRIVATE_ACCESS_KEY]

   Advanced

   --authorization value         Sugarsync 認証。 [$AUTHORIZATION]
   --authorization-expiry value  Sugarsync 認証有効期限。 [$AUTHORIZATION_EXPIRY]
   --deleted-id value            Sugarsync 削除されたフォルダのID。 [$DELETED_ID]
   --encoding value              バックエンドのエンコーディング（デフォルト: "Slash,Ctl,InvalidUtf8,Dot"） [$ENCODING]
   --refresh-token value         Sugarsync リフレッシュトークン。 [$REFRESH_TOKEN]
   --root-id value               Sugarsync ルートID。 [$ROOT_ID]
   --user value                  Sugarsync ユーザー。 [$USER]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}