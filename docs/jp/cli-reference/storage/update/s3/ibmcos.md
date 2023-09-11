# IBM COS S3

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 ibmcos - IBM COS S3

USAGE:
   singularity storage update s3 ibmcos [command options] <name|id>

DESCRIPTION:
   --env-auth
      AWS認証情報をランタイム（環境変数またはEC2/ECSメタデータ、環境変数がない場合）から取得します。
      
      access_key_idとsecret_access_keyが空の場合にのみ適用されます。

      例:
         | false | 次の手順でAWS認証情報を入力します。
         | true  | 環境（環境変数またはIAM）からAWS認証情報を取得します。

   --access-key-id
      AWSアクセスキーID。
      
      匿名アクセスまたはランタイム認証情報の場合は、空のままにします。

   --secret-access-key
      AWSシークレットアクセスキー（パスワード）。
      
      匿名アクセスまたはランタイム認証情報の場合は、空のままにします。

   --region
      接続するリージョン。
      
      S3クローンを使用してリージョンがない場合は、空のままにします。

      例:
         | <unset>            | 確認が取れない場合に使用します。
         |                    | v4シグネチャと空のリージョンを使用します。
         | other-v2-signature | v4シグネチャが機能しない場合にのみ使用します。
         |                    | 例：Jewel/v10 CEPHより前。
