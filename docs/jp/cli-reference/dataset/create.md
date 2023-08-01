# 新しいデータセットの作成

{% code fullWidth="true" %}
```
NAME:
   singularity dataset create - 新しいデータセットの作成

USAGE:
   singularity dataset create [command options] <dataset_name>

DESCRIPTION:
   <dataset_name>はデータセットの一意の識別子である必要があります
   データセットは異なるデータセットを区別するためのトップレベルのオブジェクトです。

OPTIONS:
   --help, -h  ヘルプを表示する

   暗号化

   --encryption-recipient value [ --encryption-recipient value ]  暗号化受信者の公開鍵
   --encryption-script value                                      [WIP] カスタム暗号化の実行に使用するEncryptionScriptコマンド

   インライン準備

   --output-dir value, -o value [ --output-dir value, -o value ]  CARファイルの出力ディレクトリ（デフォルト：必要なし）

   準備パラメータ

   --max-size value, -M value    作成されるCARファイルの最大サイズ（デフォルト："31.5GiB"）
   --piece-size value, -s value  ピースの確認値の計算に使用されるCARファイルのターゲットピースサイズ（デフォルト：推測）

```
{% endcode %}