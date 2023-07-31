# OpenStack Swift（Rackspace Cloud Files、Memset Memstore、OVH）

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add swift - OpenStack Swift（Rackspace Cloud Files、Memset Memstore、OVH）

USAGE:
   singularity datasource add swift [command options] <dataset_name> <source_path>

DESCRIPTION:
   --swift-application-credential-id
      アプリケーション資格情報のID（OS_APPLICATION_CREDENTIAL_ID）。

   --swift-application-credential-name
      アプリケーション資格情報の名前（OS_APPLICATION_CREDENTIAL_NAME）。

   --swift-application-credential-secret
      アプリケーション資格情報のシークレット（OS_APPLICATION_CREDENTIAL_SECRET）。

   --swift-auth
      サーバーの認証URL（OS_AUTH_URL）。

      例:
         | https://auth.api.rackspacecloud.com/v1.0     | Rackspace US
         | https://lon.auth.api.rackspacecloud.com/v1.0 | Rackspace UK
         | https://identity.api.rackspacecloud.com/v2.0 | Rackspace v2
         | https://auth.storage.memset.com/v1.0         | Memset Memstore UK
         | https://auth.storage.memset.com/v2.0         | Memset Memstore UK v2
         | https://auth.cloud.ovh.net/v3                | OVH

   --swift-auth-token
      代替認証の認証トークン（OS_AUTH_TOKEN）。

   --swift-auth-version
      AuthVersion。認証URLにバージョンがない場合は（1,2,3）に設定します（ST_AUTH_VERSION）。

   --swift-chunk-size
      このサイズ以上のファイルは_segmentsコンテナにチャンク化されます。

      このサイズ以上のファイルは_segmentsコンテナにチャンク化されます。デフォルトは最大値である5 GiBです。

   --swift-domain
      ユーザードメイン - オプション（v3認証）（OS_USER_DOMAIN_NAME）

   --swift-encoding
      バックエンドのエンコーディング。
      
      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --swift-endpoint-type
      サービスカタログから選択するエンドポイントタイプ（OS_ENDPOINT_TYPE）。

      例:
         | public   | Public（デフォルト、分からない場合はこれを選択）
         | internal | インターナル（内部サービスネットを使用）
         | admin    | Admin

   --swift-env-auth
      オープンスタックの標準形式の環境変数からswiftの資格情報を取得します。

      例:
         | false | 次のステップでswiftの資格情報を入力します。
         | true  | 環境変数からswiftの資格情報を取得します。これを使用する場合は他のフィールドは空白のままにしてください。

   --swift-key
      APIキーまたはパスワード（OS_PASSWORD）。

   --swift-leave-parts-on-error
      エラー発生時にアップロードを中止しない場合はtrueに設定します。

      これはセッションを超えたアップロードの再開時にtrueに設定する必要があります。

   --swift-no-chunk
      ストリーミングアップロード中にファイルをチャンク化しない場合は指定します。

      ストリーミングアップロード（たとえば、rcatまたはmountを使用する）の場合、このフラグを設定すると、
      swiftバックエンドはチャンク化されたファイルをアップロードしなくなります。

      これにより、最大アップロードサイズが5 GiBに制限されますが、
      チャンク化されていないファイルの取り扱いが容易でMD5SUMがあります。

      通常のコピー操作を行う場合、rcloneは引き続きchunk_sizeより大きいファイルをチャンク化します。

   --swift-no-large-objects
      静的および動的な大きなオブジェクトのサポートを無効にします。

      Swiftは5 GiBより大きいファイルを透過的に保存できません。それには2つの方法があり、
      静的または動的な大きなオブジェクトを使用しますが、APIではオブジェクトが静的なのか動的なのかを
      HEADリクエストなしでrcloneが判断できないため、たとえばチェックサムを読み取る場合にHEADリクエストを
      発行する必要があります。

      no_large_objectsが設定されていると、rcloneは静的または動的な大きなオブジェクトが
      格納されていないものと仮定します。これにより、余分なHEADリクエストを行わずにパフォーマンスが大幅に
      向上します。特に`--checksum`が設定されたswiftからswiftへの転送時にその効果があります。

      このオプションを設定すると、no_chunkおよびファイルを5 GiBより大きい場合にアップロードできない
      という意味も含まれます。そのため、5 GiBより大きいファイルはアップロードに失敗します。

      このオプションを設定し、実際に静的または動的な大きなオブジェクトが存在する場合、
      これによりこれらのオブジェクトのハッシュが正しくなりません。ダウンロードは成功しますが、
      削除やコピーなどのその他の操作は失敗します。

   --swift-region
      リージョン名 - オプション（OS_REGION_NAME）。

   --swift-storage-policy
      新しいコンテナの作成時に使用するストレージポリシー。

      新しいコンテナの作成時に指定したストレージポリシーが適用されます。
      このポリシーは後から変更できません。許可される構成値とその意味は、
      使用しているSwiftストレージプロバイダによって異なります。

      例:
         | <unset> | デフォルト
         | pcs     | OVH Public Cloud Storage
         | pca     | OVH Public Cloud Archive

   --swift-storage-url
      ストレージURL - オプション（OS_STORAGE_URL）。

   --swift-tenant
      テナント名 - v1認証の場合はオプション、それ以外の場合はこれまたはtenant_idが必要です（OS_TENANT_NAMEまたはOS_PROJECT_NAME）。

   --swift-tenant-domain
      テナントドメイン - オプション（v3認証）（OS_PROJECT_DOMAIN_NAME）。

   --swift-tenant-id
      テナントID - v1認証の場合はオプション、それ以外の場合はtenantが必要です（OS_TENANT_ID）。

   --swift-user
      ログインするユーザー名（OS_USERNAME）。

   --swift-user-id
      ログインするユーザーID - オプション - ほとんどのSwiftシステムはユーザーを使用し、空白のままにします（v3認証）（OS_USER_ID）。


OPTIONS:
   --help, -h  ヘルプを表示

   データ準備オプション

   --delete-after-export    [危険] データセットのファイルをCARファイルにエクスポートした後、削除します。  (デフォルト: false)
   --rescan-interval value  最後の成功したスキャンからこの間隔が経過すると、自動的にソースディレクトリを再スキャンします（デフォルト: 無効）
   --scanning-state value   初期のスキャン状態を設定します（デフォルト: ready）

   swift向けのオプション

   --swift-application-credential-id value      アプリケーション資格情報のID（OS_APPLICATION_CREDENTIAL_ID）。[$SWIFT_APPLICATION_CREDENTIAL_ID]
   --swift-application-credential-name value    アプリケーション資格情報の名前（OS_APPLICATION_CREDENTIAL_NAME）。[$SWIFT_APPLICATION_CREDENTIAL_NAME]
   --swift-application-credential-secret value  アプリケーション資格情報のシークレット（OS_APPLICATION_CREDENTIAL_SECRET）。[$SWIFT_APPLICATION_CREDENTIAL_SECRET]
   --swift-auth value                           サーバーの認証URL（OS_AUTH_URL）。[$SWIFT_AUTH]
   --swift-auth-token value                     代替認証の認証トークン - オプション（OS_AUTH_TOKEN）。[$SWIFT_AUTH_TOKEN]
   --swift-auth-version value                   AuthVersion。認証URLにバージョンがない場合は（1,2,3）に設定します（ST_AUTH_VERSION）。 (デフォルト: "0") [$SWIFT_AUTH_VERSION]
   --swift-chunk-size value                     このサイズ以上のファイルは_segmentsコンテナにチャンク化されます（デフォルト: "5Gi"）[$SWIFT_CHUNK_SIZE]
   --swift-domain value                         ユーザードメイン - オプション（v3認証）（OS_USER_DOMAIN_NAME） [$SWIFT_DOMAIN]
   --swift-encoding value                       バックエンドのエンコーディング（デフォルト: "Slash,InvalidUtf8"）[$SWIFT_ENCODING]
   --swift-endpoint-type value                  サービスカタログから選択するエンドポイントタイプ（OS_ENDPOINT_TYPE）。 (デフォルト: "public") [$SWIFT_ENDPOINT_TYPE]
   --swift-env-auth value                       オープンスタックの標準形式の環境変数からswiftの資格情報を取得します（デフォルト: "false"） [$SWIFT_ENV_AUTH]
   --swift-key value                            APIキーまたはパスワード（OS_PASSWORD）。[$SWIFT_KEY]
   --swift-leave-parts-on-error value           エラー発生時にアップロードを中止しない場合はtrueに設定します（デフォルト: "false"）[$SWIFT_LEAVE_PARTS_ON_ERROR]
   --swift-no-chunk value                       ストリーミングアップロード中にファイルをチャンク化しない場合は指定します（デフォルト: "false"）[$SWIFT_NO_CHUNK]
   --swift-no-large-objects value               静的および動的な大きなオブジェクトのサポートを無効にします（デフォルト: "false"）[$SWIFT_NO_LARGE_OBJECTS]
   --swift-region value                         リージョン名 - オプション（OS_REGION_NAME）[$SWIFT_REGION]
   --swift-storage-policy value                 新しいコンテナの作成時に使用するストレージポリシー。[$SWIFT_STORAGE_POLICY]
   --swift-storage-url value                    ストレージURL - オプション（OS_STORAGE_URL）[$SWIFT_STORAGE_URL]
   --swift-tenant value                         テナント名 - v1認証の場合はオプション、それ以外の場合はこれまたはtenant_idが必要です（OS_TENANT_NAMEまたはOS_PROJECT_NAME）[$SWIFT_TENANT]
   --swift-tenant-domain value                  テナントドメイン - オプション（v3認証）（OS_PROJECT_DOMAIN_NAME）[$SWIFT_TENANT_DOMAIN]
   --swift-tenant-id value                      テナントID - v1認証の場合はオプション、それ以外の場合はtenantが必要です（OS_TENANT_ID）[$SWIFT_TENANT_ID]
   --swift-user value                           ログインするユーザー名（OS_USERNAME）[$SWIFT_USER]
   --swift-user-id value                        ログインするユーザーID - オプション - ほとんどのSwiftシステムはユーザーを使用し、空白のままにします（v3認証）（OS_USER_ID）[$SWIFT_USER_ID]

```
{% endcode %}