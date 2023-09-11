# Oracle Cloud Infrastructure Object Storage

{% code fullWidth="true" %}
```
NAME:
   singularity storage create oos - Oracle Cloud Infrastructure オブジェクトストレージ

使用法:
   singularity storage create oos command [command options] [arguments...]

コマンド:
   env_auth                 ランタイム (環境) から自動的に認証情報を自動的に取得します。認証情報を最初に提供したものが優先されます。
   instance_principal_auth  インスタンスプリンシパルを使用して、インスタンスが API 呼び出しを行うことを認可します。
                            各インスタンスは独自のアイデンティティを持ち、インスタンスメタデータから読み取られる証明書を使用して認証します。
                            https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm
   no_auth                  認証情報は必要ありません。これは通常、パブリックバケットの読み取りに使用します。
   resource_principal_auth  リソースプリンシパルを使用して API 呼び出しを行います。
   user_principal_auth      OCI ユーザーと API キーを使用して認証します。
                            テナンシーの OCID、ユーザーの OCID、リージョン、API キーへのパス、指紋をコンフィグファイルに入力する必要があります。
                            https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm
   help, h                  コマンドのリストまたはコマンドのヘルプを表示します。

オプション:
   --help, -h  ヘルプを表示
```
{% endcode %}