# Oracle Cloud Infrastructure オブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity storage update oos - Oracle Cloud Infrastructure オブジェクトストレージ

USAGE:
   singularity storage update oos command [command options] [arguments...]

COMMANDS:
   env_auth                 実行時の認証情報（env）から自動的にクレデンシャルを取得します。最初に認証情報が提供されたものが優先されます。
   instance_principal_auth  インスタンスプリンシパルを使用して、インスタンスがAPI呼び出しを行うための認証を行います。
                            各インスタンスには独自のアイデンティティがあり、インスタンスメタデータから読み取られる証明書を使用して認証します。 
                            詳細については、https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm を参照してください。
   no_auth                  クレデンシャルは必要ありません。通常、公開バケットを読み取るために使用されます。
   resource_principal_auth  リソースプリンシパルを使用してAPI呼び出しを行います。
   user_principal_auth      OCIユーザーとAPIキーを使用して認証します。
                            テナンシーOCID、ユーザーOCID、リージョン、パス、APIキーのフィンガープリントを設定ファイルに入力する必要があります。
                            詳細については、https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm を参照してください。
   help, h                  コマンドのリストを表示するか、特定のコマンドのヘルプを表示します。

OPTIONS:
   --help, -h  ヘルプの表示
```
{% endcode %}