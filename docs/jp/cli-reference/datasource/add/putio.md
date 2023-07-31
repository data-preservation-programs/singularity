# Put.io

{% code fullWidth="true" %}
```
名前：
   singularity datasource add putio - Put.io

使用方法：
   singularity datasource add putio [コマンドオプション] <データセット名> <ソースパス>

説明：
   --putio-encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション：
   --help, -h  ヘルプを表示する

   データの準備オプション

   --delete-after-export    [注意] データセットをCARファイルにエクスポートした後、ファイルを削除します。 (デフォルト: false)
   --rescan-interval value  最後の成功したスキャンからこの間隔が経過すると、自動的にソースディレクトリを再スキャンします (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します (デフォルト: ready)

   Put.ioのオプション

   --putio-encoding value  バックエンドのエンコーディングです (デフォルト: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PUTIO_ENCODING]

```
{% endcode %}