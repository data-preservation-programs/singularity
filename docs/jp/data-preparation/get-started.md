# Singularityのはじめ方

Singularityをセットアップして使用を開始するために、以下の手順に従ってください。

## 1. データベースの初期化

初めてSingularityを使用する場合、データベースの初期化が必要です。この手順は一度だけ必要です。

```sh
singularity admin init
```

## 2. ストレージシステムへの接続
Singularityは、40以上のさまざまなストレージシステムとシームレスに統合するためにRCloneと提携しています。これらのストレージシステムは、2つの主要な役割を果たすことができます：
* **ソースストレージ**: データセットが現在格納されている場所であり、Singularityが準備のためにデータを取得する場所です。
* **出力ストレージ**: Singularityが処理後のCAR（Content Addressable Archive）ファイルを格納する宛先です。
ニーズに適したストレージシステムを選択し、それをSingularityに接続してデータセットの準備を開始しましょう。

### 2a. ローカルファイルシステムの追加

最も一般的なストレージシステムはローカルファイルシステムです。以下の方法でフォルダをソースストレージとしてsingularityに追加します：

```sh
singularity storage create local --name "my-source" --path "/mnt/dataset/folder"
```

### 2b. S3データソースの追加

AWS S3やMinIOなどを含む、S3互換のストレージシステムを利用することができます。以下は公開データセットの例です。

```sh
singularity storage create s3 aws --name "my-source" --path "public-dataset-test"
```

## 3. 準備の作成
```sh
singularity prep create --source "my-source" --name "my-prep"
```

## 4. 準備ワーカーの実行
```sh
singularity prep start-scan my-prep my-source
singularity run dataset-worker
```

## 5. 準備のステータスと結果の確認
```sh
singularity prep status my-prep
singularity prep list-pieces my-prep
```