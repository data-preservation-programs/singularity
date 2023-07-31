# CARファイルの配布

さて、CARファイルをストレージプロバイダーに配布して、彼らがそれを自分の側にインポートできるようにしましょう。まず、コンテンツプロバイダーサービスを実行して、準備したデータセットのピースをダウンロードします。

```sh
singularity run content-provider
wget 127.0.0.1:8088/piece/bagaxxxx
```

もし以前にCARのエクスポート先ディレクトリを指定していた場合（インラインの準備が無効になります）、そのCARファイルは直接それらのCARファイルから提供されます。さもなければ、インラインの準備を使用していたか、それらのCARファイルを誤って削除した場合、元のデータソースから直接提供されます。

## 次のステップ

[deal-making-prerequisite.md](../deal-making/deal-making-prerequisite.md "mention")