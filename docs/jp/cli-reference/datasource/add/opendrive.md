# OpenDrive

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add opendrive - OpenDrive

USAGE:
   singularity datasource add opendrive [コマンドオプション] <データセット名> <ソースパス>

DESCRIPTION:
   --opendrive-chunk-size
      ファイルはこのサイズのチャンクでアップロードされます。
      
      これらのチャンクはメモリにバッファされるため、増やすとメモリ使用量が増加します。

   --opendrive-encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --opendrive-password
      パスワードです。

   --opendrive-username
      ユーザー名です。


OPTIONS:
   --help, -h  ヘルプを表示します

   データ準備オプション

   --delete-after-export    [危険] データセットのファイルをCARファイルにエクスポートした後に削除します。 (デフォルト: false)
   --rescan-interval value  最後の成功したスキャンからこの間隔が経過した場合に、自動的にソースディレクトリを再スキャンします。 (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します。 (デフォルト: ready)

   OpenDriveのオプション

   --opendrive-chunk-size value  ファイルはこのサイズのチャンクでアップロードされます。 (デフォルト: "10Mi") [$OPENDRIVE_CHUNK_SIZE]
   --opendrive-encoding value    バックエンドのエンコーディングです。 (デフォルト: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot") [$OPENDRIVE_ENCODING]
   --opendrive-password value    パスワードです。 [$OPENDRIVE_PASSWORD]
   --opendrive-username value    ユーザー名です。 [$OPENDRIVE_USERNAME]

```
{% endcode %}