# 既存のデータセットを更新する

{% code fullWidth="true" %}
```
NAME:
   singularity dataset update - 既存のデータセットを更新する

USAGE:
   singularity dataset update [command options] <dataset_name>

OPTIONS:
   --help, -h  ヘルプを表示する

   暗号化

   --encryption-recipient value [ --encryption-recipient value ]  暗号化の受信側の公開鍵
   --encryption-script value                                      カスタム暗号化のために実行するEncryptionScriptコマンド

   インライン準備

   --output-dir value, -o value [ --output-dir value, -o value ]  CARファイルの出力ディレクトリ（デフォルト：必要ありません）

   準備パラメータ

   --max-size value, -M value    作成されるCARファイルの最大サイズ（デフォルト："30GiB"）
   --piece-size value, -s value  ピース確定計算に使用されるCARファイルの目標ピースサイズ（デフォルト：推測）
```
{% endcode %}