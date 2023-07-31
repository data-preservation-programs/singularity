# Jottacloud

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add jottacloud - Jottacloud

USAGE:
   singularity datasource add jottacloud [command options] <dataset_name> <source_path>

DESCRIPTION:
   --jottacloud-encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --jottacloud-hard-delete
      ファイルをゴミ箱に入れずに永久に削除します。

   --jottacloud-md5-memory-limit
      このサイズを超えるファイルは、必要に応じてMD5を計算するためにディスク上にキャッシュされます。

   --jottacloud-no-versions
      サーバーサイドのバージョニングを回避し、ファイルを上書きする代わりに削除して再作成します。

   --jottacloud-trashed-only
      ゴミ箱にあるファイルのみ表示します。
      
      これにより、元のディレクトリ構造でトレイされたファイルが表示されます。

   --jottacloud-upload-resume-limit
      このサイズを超えるファイルは、アップロードが失敗した場合に再開できます。


OPTIONS:
   --help, -h  ヘルプを表示

   データ準備オプション

   --delete-after-export    [危険] データセットのファイルをエクスポート後に削除する（デフォルト：無効）
   --rescan-interval value  最後の正常なスキャンからこの間隔が経過した場合、自動的にソースディレクトリを再スキャンします（デフォルト：無効）
   --scanning-state value   初期のスキャン状態を設定します（デフォルト：準備完了）

   Jottacloudオプション

   --jottacloud-encoding value             バックエンドのエンコーディングです。（デフォルト："Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot"）[$JOTTACLOUD_ENCODING]
   --jottacloud-hard-delete value          ファイルをゴミ箱に入れずに永久に削除します。（デフォルト："false"）[$JOTTACLOUD_HARD_DELETE]
   --jottacloud-md5-memory-limit value     このサイズを超えるファイルは、必要に応じてMD5を計算するためにディスク上にキャッシュされます。（デフォルト："10Mi"）[$JOTTACLOUD_MD5_MEMORY_LIMIT]
   --jottacloud-no-versions value          サーバーサイドのバージョニングを回避し、ファイルを上書きする代わりに削除して再作成します。（デフォルト："false"）[$JOTTACLOUD_NO_VERSIONS]
   --jottacloud-trashed-only value         ゴミ箱にあるファイルのみ表示します。（デフォルト："false"）[$JOTTACLOUD_TRASHED_ONLY]
   --jottacloud-upload-resume-limit value  このサイズを超えるファイルは、アップロードが失敗した場合に再開できます。（デフォルト："10Mi"）[$JOTTACLOUD_UPLOAD_RESUME_LIMIT]

```
{% endcode %}