# メタデータAPIからCARファイルをダウンロードする

{% code fullWidth="true" %}
```
NAME:
   singularity download - メタデータAPIからCARファイルをダウンロードする

USAGE:
   singularity download [command options] PIECE_CID

CATEGORY:
   ユーティリティ

OPTIONS:
   オプション一般

   --api value                    メタデータAPIのURL (デフォルト: "http://127.0.0.1:7777")
   --concurrency value, -j value  並行ダウンロード数 (デフォルト: 10)
   --out-dir value, -o value      CARファイルを書き込むディレクトリ (デフォルト: ".")

   acdのオプション

   --acd-auth-url value            認証サーバーのURL [$ACD_AUTH_URL]
   --acd-client-id value           OAuthクライアントID [$ACD_CLIENT_ID]
   --acd-client-secret value       OAuthクライアントシークレット [$ACD_CLIENT_SECRET]
   --acd-encoding value            バックエンドのエンコーディング (デフォルト: "Slash,InvalidUtf8,Dot") [$ACD_ENCODING]
   --acd-templink-threshold value  このサイズ以上のファイルはtempLinkを使用してダウンロードされます (デフォルト: "9Gi") [$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               OAuthアクセストークンのJSON blob [$ACD_TOKEN]
   --acd-token-url value           トークンサーバーのURL [$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  完了したアップロードの後に失敗した場合に、GiBごとに追加の待ち時間を設定します (デフォルト: "3m0s") [$ACD_UPLOAD_WAIT_PER_GB]

   azureblobのオプション

   --azureblob-access-tier value                    ブロブのアクセスレベル: hot、cool、またはarchive [$AZUREBLOB_ACCESS_TIER]
   --azu..以下略..
```
{% endcode %}