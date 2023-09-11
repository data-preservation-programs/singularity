# 新しい準備の作成

{% code fullWidth="true" %}
```
名前:
  singularity prep create - 新しい準備の作成

使用方法:
  singularity prep create [コマンドオプション] [引数...]

カテゴリ:
  準備管理

オプション:
  --delete-after-export  CAR ファイルへのエクスポート後にソースファイルを削除するかどうか（デフォルト: false）
  --help, -h             ヘルプの表示
  --max-size value       1 つの CAR ファイルの最大サイズ（デフォルト: "31.5GiB"）
  --name value           準備の名前（デフォルト: 自動生成）
  --output value         準備に使用する出力ストレージの ID または名前
  --piece-size value     ピース検証計算に使用する CAR ファイルのターゲットピースサイズ（デフォルト: --max-size で決定）
  --source value         準備に使用するソースストレージの ID または名前

  ローカルの出力パスを使用したクイック作成

  --local-output value   準備に使用するローカルの出力パス。これは提供されたパスで出力ストレージを作成する便利なフラグです。

  ローカルのソースパスを使用したクイック作成

  --local-source value   準備に使用するローカルのソースパス。これは提供されたパスでソースストレージを作成する便利なフラグです。
```
{% endcode %}