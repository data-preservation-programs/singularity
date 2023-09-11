# Jottacloud

{% code fullWidth="true" %}
```
NAME:
   singularity storage update jottacloud - Jottacloud

USAGE:
   singularity storage update jottacloud [オプション] <名前|ID>

DESCRIPTION:
   --md5-memory-limit
      このサイズより大きいファイルは、MD5を計算するためにディスクにキャッシュされます。

   --trashed-only
      ゴミ箱にあるファイルのみ表示します。
      
      オリジナルのディレクトリ構造でゴミ箱のファイルを表示します。

   --hard-delete
      ファイルをごみ箱に入れる代わりに、永久に削除します。

   --upload-resume-limit
      このサイズより大きいファイルは、アップロードが失敗した場合に再開することができます。

   --no-versions
      ファイルの上書きではなく、ファイルを削除して再作成することで、サーバーサイドのバージョニングを回避します。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --help, -h  ヘルプを表示します

   Advanced

   --encoding value             バックエンドのエンコーディングです。 (デフォルト: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete                ファイルをごみ箱に入れる代わりに、永久に削除します。 (デフォルト: false) [$HARD_DELETE]
   --md5-memory-limit value     このサイズより大きいファイルは、MD5を計算するためにディスクにキャッシュされます。 (デフォルト: "10Mi") [$MD5_MEMORY_LIMIT]
   --no-versions                ファイルの上書きではなく、ファイルを削除して再作成することで、サーバーサイドのバージョニングを回避します。 (デフォルト: false) [$NO_VERSIONS]
   --trashed-only               ゴミ箱にあるファイルのみ表示します。 (デフォルト: false) [$TRASHED_ONLY]
   --upload-resume-limit value  このサイズより大きいファイルは、アップロードが失敗した場合に再開することができます。 (デフォルト: "10Mi") [$UPLOAD_RESUME_LIMIT]
```
{% endcode %}