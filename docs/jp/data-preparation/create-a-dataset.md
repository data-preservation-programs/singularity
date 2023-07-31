---
description: データベースの初期化と新しいデータセットの作成から始めましょう
---

# データセットの作成

## データベースの初期化

デフォルトでは、`sqlite3` データベースバックエンドを使用し、データベースファイルを `$HOME/.singularity` に初期化します。

本番環境で異なるデータベースバックエンドを使用する場合は、[deploy-to-production.md](../installation/deploy-to-production.md "mention") を参照してください。

```sh
singularity admin init
```

## 新しいデータセットの作成

データセットは、単一のデータセットに関連するデータソースのコレクションです。データセットが作成されると、データソースを追加したり、Filecoin ウォレットアドレスを関連付けることができます。

```sh
singularity dataset create my_dataset
```

デフォルトでは、Singularity は「Inline Preparation」と呼ばれる技術を使用しており、CAR ファイルにエクスポートされません。これは、ほとんどのデータソースにおいて変更がなく、CAR ファイルは元のデータソースと同じ内容を保存しているからです。

## 次のステップ

[add-a-data-source.md](add-a-data-source.md "mention")