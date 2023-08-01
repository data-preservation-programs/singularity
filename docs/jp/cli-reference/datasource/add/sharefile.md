# Citrix Sharefile

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add sharefile - Citrix Sharefile

USAGE:
   singularity datasource add sharefile [command options] <dataset_name> <source_path>

DESCRIPTION:
   --sharefile-chunk-size
      アップロードのチャンクサイズです。
      
      256k以上で、2のべき乗でなければなりません。
      
      これを大きくするとパフォーマンスが向上しますが、各チャンクは転送ごとにメモリにバッファされます。
      
      これを小さくするとメモリ使用量が減りますが、パフォーマンスは低下します。

   --sharefile-encoding
      バックエンドのエンコーディングです。
      
      詳細については[エンコーディングセクション](/overview/#encoding)を参照してください。

   --sharefile-endpoint
      APIの呼び出し用のエンドポイントです。
      
      通常はOAuthプロセスの一部として自動的に検出されますが、`https://XXX.sharefile.com`のように手動で設定することもできます。
      

   --sharefile-root-folder-id
      ルートフォルダのIDです。
      
      空白のままにすると、「個人フォルダ」にアクセスします。ここでは、標準の値を使用するか、任意のフォルダID（長い16進数のID）を使用できます。

      例:
         | <unset>    | パーソナルフォルダへのアクセス（デフォルト）。
         | favorites  | お気に入りのフォルダへのアクセス。
         | allshared  | 共有フォルダへのアクセス。
         | connectors | 個々のコネクタへのアクセス。
         | top        | ホーム、お気に入り、共有フォルダとコネクタへのアクセス。

   --sharefile-upload-cutoff
      マルチパートアップロードに切り替えるためのカットオフ値です。


OPTIONS:
   --help, -h  ヘルプを表示する

   データの準備オプション

   --delete-after-export    [注意] データセットのファイルをCARファイルにエクスポートした後に削除する (デフォルト: false)
   --rescan-interval value  最後の正常なスキャンからこの間隔が経過した場合に自動的にソースディレクトリを再スキャンする (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定する (デフォルト: ready)

   Sharefileのオプション

   --sharefile-chunk-size value      アップロードのチャンクサイズ (デフォルト: "64Mi") [$SHAREFILE_CHUNK_SIZE]
   --sharefile-encoding value        バックエンドのエンコーディング (デフォルト: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot") [$SHAREFILE_ENCODING]
   --sharefile-endpoint value        APIの呼び出し用のエンドポイント [$SHAREFILE_ENDPOINT]
   --sharefile-root-folder-id value  ルートフォルダのID [$SHAREFILE_ROOT_FOLDER_ID]
   --sharefile-upload-cutoff value   マルチパートアップロードに切り替えるためのカットオフ値 (デフォルト: "128Mi") [$SHAREFILE_UPLOAD_CUTOFF]

```
{% endcode %}