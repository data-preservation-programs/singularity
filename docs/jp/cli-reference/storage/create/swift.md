# OpenStack Swift（Rackspace Cloud Files、Memset Memstore、OVH）

{% code fullWidth="true" %}
```
名前:
   singularity storage create swift - OpenStack Swift（Rackspace Cloud Files、Memset Memstore、OVH）

使用法:
   singularity storage create swift [コマンドオプション] [引数...]

説明:
   --env-auth
      標準のOpenStack形式で環境変数からSwiftの認証情報を取得します。

      例:
         | false | 次のステップでSwiftの認証情報を入力します。
         | true  | 環境変数からSwiftの認証情報を取得します。
         |       | このオプションを使用する場合、他のフィールドは空白のままにします。

   --user
      ログインするユーザー名（OS_USERNAME）。

   --key
      APIキーまたはパスワード（OS_PASSWORD）。

   --auth
      サーバーの認証URL（OS_AUTH_URL）。

      例:
         | https://auth.api.rackspacecloud.com/v1.0     | Rackspace US
         | https://lon.auth.api.rackspacecloud.com/v1.0 | Rackspace UK
         | https://identity.api.rackspacecloud.com/v2.0 | Rackspace v2
         | https://auth.storage.memset.com/v1.0         | Memset Memstore UK
         | https://auth.storage.memset.com/v2.0         | Memset Memstore UK v2
         | https://auth.cloud.ovh.net/v3                | OVH

   --user-id
      ログインするユーザーID - オプション - ほとんどのSwiftシステムではユーザーを使用し、このフィールドには何も入力しない（v3認証）（OS_USER_ID）。

   --domain
      ユーザードメイン - オプション（v3認証）（OS_USER_DOMAIN_NAME）。

   --tenant
      テナント名 - v1認証のオプション - もしくはtenant_idが必要（OS_TENANT_NAMEまたはOS_PROJECT_NAME）。

   --tenant-id
      テナントID - v1認証のオプション - もしくはtenantが必要（OS_TENANT_ID）。

   --tenant-domain
      テナントドメイン - オプション（v3認証）（OS_PROJECT_DOMAIN_NAME）。

   --region
      リージョン名 - オプション（OS_REGION_NAME）。

   --storage-url
      ストレージURL - オプション（OS_STORAGE_URL）。

   --auth-token
      別の認証からの認証トークン - オプション（OS_AUTH_TOKEN）。

   --application-credential-id
      アプリケーション資格情報ID（OS_APPLICATION_CREDENTIAL_ID）。

   --application-credential-name
      アプリケーション資格情報名（OS_APPLICATION_CREDENTIAL_NAME）。

   --application-credential-secret
      アプリケーション資格情報のシークレット（OS_APPLICATION_CREDENTIAL_SECRET）。

   --auth-version
      AuthVersion - オプション - 認証URLにバージョンがない場合は（1,2,3）に設定します（ST_AUTH_VERSION）。

   --endpoint-type
      サービスカタログから選択するエンドポイントタイプ（OS_ENDPOINT_TYPE）。

      例:
         | public   | パブリック（デフォルト、よくわからない場合はこれを選択）
         | internal | インターナル（内部サービスネットを使用）
         | admin    | 管理者

   --leave-parts-on-error
      trueの場合、エラーが発生した場合にアップロードを中止しません。
      
      これは、異なるセッション間でアップロードを再開する場合にtrueに設定する必要があります。

   --storage-policy
      新しいコンテナを作成するときに使用するストレージポリシー。
      
      これは、新しいコンテナの作成時に指定したストレージポリシーを適用します。
      ポリシーはその後変更できません。
      許可される構成値とその意味は、Swiftストレージプロバイダによって異なります。

      例:
         | <unset> | デフォルト
         | pcs     | OVH Public Cloud Storage
         | pca     | OVH Public Cloud Archive

   --chunk-size
      このサイズ以上のファイルは_segmentsコンテナにチャンク分割されます。
      
      このサイズ以上のファイルは_segmentsコンテナにチャンク分割されます。デフォルトは5 GiBで、最大値です。

   --no-chunk
      ストリーミングアップロード中にファイルをチャンク分割しないでください。
      
      ストリーミングアップロード（例：rcatまたはマウントを使用して）を実行する場合、このフラグを設定すると、
      swiftバックエンドではチャンク分割されたファイルがアップロードされません。
      
      これにより、最大アップロードサイズが5 GiBに制限されます。
      ただし、チャンク分割されていないファイルは取り扱いが容易で、MD5SUMがあります。
      
      通常のコピー操作を実行する場合、Rcloneはまだchunk_sizeより大きなファイルをチャンク分割します。

   --no-large-objects
      静的および動的な大型オブジェクトのサポートを無効にします。
      
      Swiftは、5 GiBより大きなファイルを透過的に保存することができません。
      これを実現するためには、静的または動的な大型オブジェクトのどちらかの方法があり、
      APIからはオブジェクトが静的な大型オブジェクトか動的な大型オブジェクトかを判断することができません。
      これにより、例えばチェックサムを読み取る場合にHEADリクエストを発行する必要があります。
      
      `no_large_objects`が設定されている場合、rcloneは静的または動的な大型オブジェクトは保存されていないと仮定します。
      これにより、rcloneは余分なHEADコールを中止できるため、パフォーマンスが大幅に向上します。
      特に、`--checksum`が設定されたswift to swift転送を行う場合には特にそうです。
      
      このオプションを設定すると、`no_chunk`も意味し、チャンク分割されたファイルはアップロードされないため、
      5 GiBより大きなファイルはアップロードに失敗します。
      
      このオプションを設定し、実際に静的または動的な大型オブジェクトがある場合、それらのオブジェクトのハッシュが正しくなりません。
      ダウンロードは成功しますが、RemoveやCopyなどの他の操作は失敗します。
      

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --application-credential-id value      アプリケーション資格情報ID（OS_APPLICATION_CREDENTIAL_ID）。[$APPLICATION_CREDENTIAL_ID]
   --application-credential-name value    アプリケーション資格情報名（OS_APPLICATION_CREDENTIAL_NAME）。[$APPLICATION_CREDENTIAL_NAME]
   --application-credential-secret value  アプリケーション資格情報のシークレット（OS_APPLICATION_CREDENTIAL_SECRET）。[$APPLICATION_CREDENTIAL_SECRET]
   --auth value                           サーバーの認証URL（OS_AUTH_URL）。[$AUTH]
   --auth-token value                     別の認証からの認証トークン - オプション（OS_AUTH_TOKEN）。[$AUTH_TOKEN]
   --auth-version value                   AuthVersion - オプション - 認証URLにバージョンがない場合は（1,2,3）に設定します（ST_AUTH_VERSION）。 (デフォルト値: 0) [$AUTH_VERSION]
   --domain value                         ユーザードメイン - オプション（v3認証）（OS_USER_DOMAIN_NAME）[$DOMAIN]
   --endpoint-type value                  サービスカタログから選択するエンドポイントタイプ（OS_ENDPOINT_TYPE）。 (デフォルト値: "public") [$ENDPOINT_TYPE]
   --env-auth                             標準のOpenStack形式で環境変数からSwiftの認証情報を取得します。 (デフォルト値: false) [$ENV_AUTH]
   --help, -h                             ヘルプを表示
   --key value                            APIキーまたはパスワード（OS_PASSWORD）。[$KEY]
   --region value                         リージョン名 - オプション（OS_REGION_NAME）。[$REGION]
   --storage-policy value                 新しいコンテナを作成するときに使用するストレージポリシー。[$STORAGE_POLICY]
   --storage-url value                    ストレージURL - オプション（OS_STORAGE_URL）。[$STORAGE_URL]
   --tenant value                         テナント名 - v1認証のオプション - もしくはtenant_idが必要（OS_TENANT_NAMEまたはOS_PROJECT_NAME）。[$TENANT]
   --tenant-domain value                  テナントドメイン - オプション（v3認証）（OS_PROJECT_DOMAIN_NAME）。[$TENANT_DOMAIN]
   --tenant-id value                      テナントID - v1認証のオプション - もしくはtenantが必要（OS_TENANT_ID）。[$TENANT_ID]
   --user value                           ログインするユーザー名（OS_USERNAME）。[$USER]
   --user-id value                        ログインするユーザーID - オプション - ほとんどのSwiftシステムではユーザーを使用し、このフィールドには何も入力しない（v3認証）（OS_USER_ID）。[$USER_ID]

   上級

   --chunk-size value      このサイズ以上のファイルは_segmentsコンテナにチャンク分割されます。 (デフォルト値: "5Gi") [$CHUNK_SIZE]
   --encoding value        バックエンドのエンコーディング。 (デフォルト値: "Slash,InvalidUtf8") [$ENCODING]
   --leave-parts-on-error  trueの場合、エラーが発生した場合にアップロードを中止しません。 (デフォルト値: false) [$LEAVE_PARTS_ON_ERROR]
   --no-chunk              ストリーミングアップロード中にファイルをチャンク分割しないでください。 (デフォルト値: false) [$NO_CHUNK]
   --no-large-objects      静的および動的な大型オブジェクトのサポートを無効にします (デフォルト値: false) [$NO_LARGE_OBJECTS]

   一般

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス
```
{% endcode %}