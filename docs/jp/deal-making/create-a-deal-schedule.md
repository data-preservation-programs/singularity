# 取引スケジュールを作成する

ストレージプロバイダーとの取引を行いましょう。まず、取引作成サービスを実行します。

```
singularity run dealmaker
```

## 一度にすべての取引を送信する

データセットが小さい場合、ストレージプロバイダーに一度にすべての取引を送信することができます。以下のコマンドを使用してください。

```sh
singularity deal schedule create dataset_name provider_id
```

ただし、データセットが大きい場合、取引提案の有効期限前にそれだけの数の取引をストレージプロバイダーが受け入れるのは大変かもしれません。そのため、スケジュールを作成することができます。

## スケジュール付きの取引を送信する

同じコマンドを使用して、取引がストレージプロバイダーにどのように迅速かつ頻繁に行われるかを制御するための独自のスケジュールを作成できます。

```
--schedule-deal-number value, --number value     1回のスケジュールでの最大取引数（デフォルト：無制限）
--schedule-deal-size value, --size value         1回のスケジュールでの最大取引サイズ（デフォルト：無制限）
--schedule-interval value, --every value         バッチ取引を送信するためのクロンスケジュール（デフォルト：無効）
--total-deal-number value, --total-number value  このリクエストの最大総取引数（デフォルト：無制限）
```