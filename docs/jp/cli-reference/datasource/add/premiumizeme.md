# premiumize.me

{% code fullWidth="true" %}
```
名前:
   singularity datasource add premiumizeme - premiumize.me

使用法:
   singularity datasource add premiumizeme [コマンドオプション] <データセット名> <ソースパス>

説明:
   --premiumizeme-api-key
      APIキー。
      
      通常は使用されません - 代わりにOAuthを使用してください。
      

  --premiumizeme-encoding
      バックエンドのエンコーディング。
      
      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --help, -h  ヘルプを表示

   データ準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、ファイルを削除します。  (デフォルト: false)
   --rescan-interval 値     前回の成功したスキャンからこの間隔が経過したら、ソースディレクトリを自動的に再スキャンします。(デフォルト: 無効)
   --scanning-state 値     初期のスキャン状態を設定します。(デフォルト: ready)

   premiumizeme用オプション

   --premiumizeme-encoding 値  バックエンドのエンコーディングです。(デフォルト: "Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PREMIUMIZEME_ENCODING]

```
{% endcode %}