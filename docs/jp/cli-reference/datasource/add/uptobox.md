# Uptobox

{% code fullWidth="true" %}
```
名前：
   singularity datasource add uptobox - Uptobox

使用法：
   singularity datasource add uptobox [コマンドオプション] <データセット名> <ソースパス>

説明：
   --uptobox-access-token
      アクセストークンです。
      
      https://uptobox.com/my_account で取得できます。

   --uptobox-encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション：
   --help, -h  ヘルプを表示する

   データの準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、データのファイルを削除します。 (デフォルト: false)
   --rescan-interval value  最後のスキャンからこのインターバルが経過すると、自動的にソースディレクトリを再スキャンします。 (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します。 (デフォルト: ready)

   uptoboxのオプション

   --uptobox-access-token value  アクセストークンです。 [$UPTOBOX_ACCESS_TOKEN]
   --uptobox-encoding value      バックエンドのエンコーディングです。 (デフォルト: "Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot") [$UPTOBOX_ENCODING]

```
{% endcode %}