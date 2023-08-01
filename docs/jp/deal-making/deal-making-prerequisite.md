# ディールを行うための事前準備

## ストレージプロバイダーを探す

現時点では、Singularityはディールを受け入れるストレージプロバイダーを見つける手助けをしてくれません。良質なストレージプロバイダーを見つけるためには、以下のリソースを利用することができます。

* TODO

## Filecoinウォレットを作成する

ディールを行う前に、Filecoinウォレットを作成する必要があります。Ledgerウォレットや取引所ウォレットは使用できません。Filecoinウォレットを作成するには、以下のコマンドを実行します。

```sh
singularity wallet create
```

これにより、ウォレットに関連付けられたウォレットアドレスとプライベートキーが生成されます。このウォレットはまだブロックチェーンに認識されていないため、ディールを行うことはできません。今が適切なタイミングですので、このウォレットに0 FILを送金しておくことで皆が認識できるようにしましょう。

このウォレットがチェーン上に記録されたら、上記のコマンドが完了し、ディールの準備が整います。

また、すでに既存のウォレットをお持ちの場合は、以下のコマンドを使用してウォレットをインポートすることもできます。

```sh
singularity wallet import xxx
```

## \[オプション] [データキャップ](https://docs.filecoin.io/basics/how-storage-works/filecoin-plus/#datacap)の取得

現在の市場状況では、多くのストレージプロバイダーは通常のディールではなく、[検証済みディール](https://docs.filecoin.io/storage-provider/filecoin-deals/verified-deals/)を好む傾向にあります。もしデータセットが数テラバイト以上の場合は、[Filplusガバナンスチームとノテアリー](https://github.com/filecoin-project/notary-governance)にデータキャップの申請をするべきです。

## 次のステップ

[create-a-deal-schedule.md](create-a-deal-schedule.md "mention")を参照してください。