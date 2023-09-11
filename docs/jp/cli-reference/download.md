# メタデータAPIからCARファイルをダウンロードする

{% code fullWidth="true" %}
```
NAME:
   singularity download - メタデータAPIからCARファイルをダウンロードする

使用法:
    singularity download [コマンドオプション] <piece_cid>

カテゴリ:
    ユーティリティ

オプション:
    1Fichier

    --fichier-api-key value           APIキー、https://1fichier.com/console/params.plから取得します。 [$FICHIER_API_KEY]
    --fichier-encoding value          バックエンドのエンコーディング。 (デフォルト: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$FICHIER_ENCODING]
    --fichier-file-password value     パスワードで保護された共有ファイルをダウンロードする場合、このパラメータを追加します。 [$FICHIER_FILE_PASSWORD]
    --fichier-folder-password value   パスワードで保護された共有フォルダ内のファイルをリストする場合、このパラメータを追加します。 [$FICHIER_FOLDER_PASSWORD]
    --fichier-shared-folder value     共有フォルダをダウンロードする場合、このパラメータを追加します。 [$FICHIER_SHARED_FOLDER]

    Akamai NetStorage

    --netstorage-account value   NetStorageアカウント名を設定します。 [$NETSTORAGE_ACCOUNT]
    --netstorage-host value      接続するNetStorageホストのドメイン+パスを設定します。 [$NETSTORAGE_HOST]
    --netstorage-protocol value  HTTPまたはHTTPSプロトコルを選択します。 (デフォルト: "https") [$NETSTORAGE_PROTOCOL]
    --netstorage-secret value    認証のためのNetStorageアカウントのsecret/G2Oキーを設定します。 [$NETSTORAGE_SECRET]

    Amazon Drive

    --acd-auth-url value    AuthサーバーのURL。 [$ACD_AUTH_URL]
    --acd-checkpoint value  内部ポーリングのチェックポイント（デバッグ用）。 [$ACD_CHECKPOINT]
    --acd-client-id value   OAuthクライアントID。 [$ACD_CLIENT_ID]
    --acd-client-secret value  OAuthクライアントシークレット [$ACD_CLIENT_SECRET]
    --acd-encoding value  バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ACD_ENCODING]
    --acd-templink-threshold value  このサイズ以上のファイルはtempLinkを使用してダウンロードされます。 (デフォルト: "9Gi") [$ACD_TEMPLINK_THRESHOLD]
    --acd-token value  OAuthアクセストークン（JSON形式） [$ACD_TOKEN]
    --acd-token-url value  TokenサーバーのURL。[$ACD_TOKEN_URL]
    --acd-upload-wait-per-gb value  失敗した完全なアップロードの後に待機するための1GiBごとの追加時間を設定します。 (デフォルト: "3m0s") [$ACD_UPLOAD_WAIT_PER_GB]

    Amazon S3に準拠したストレージプロバイダ（AWS、Alibaba、Ceph、China Mobile、Cloudflare、ArvanCloud、DigitalOcean、Dreamhost、Huawei OBS、IBM COS、IDrive e2、IONOS Cloud、Liara、Lyve Cloud、Minio、Netease、RackCorp、Scaleway、SeaweedFS、StackPath、Storj、Tencent COS、Qiniu、Wasabiを含む）

    --s3-access-key-id value  AWS Access Key ID。 [$S3_ACCESS_KEY_ID]
    --s3-acl value  バケットとオブジェクトの作成時に使用するCanned ACL。 [$S3_ACL]
    --s3-bucket-acl value  バケットの作成時に使用するCanned ACL。 [$S3_BUCKET_ACL]
    --s3-chunk-size value  アップロードに使用するチャンクサイズ。 (デフォルト: "5Mi") [$S3_CHUNK_SIZE]
    --s3-copy-cutoff value  マルチパートコピーに切り替えるためのカットオフ。 (デフォルト: "4.656Gi") [$S3_COPY_CUTOFF]
    --s3-decompress  設定した場合、gzipエンコードされたオブジェクトを展開します。 (デフォルト: false) [$S3_DECOMPRESS]
    --s3-disable-checksum  オブジェクトメタデータにMD5チェックサムを格納しな