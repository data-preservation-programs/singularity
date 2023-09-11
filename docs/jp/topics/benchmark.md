# Singularityを使用したベンチマーク

Singularityの`ez-prep`コマンドは、ベンチマークを行うための簡略化された方法を提供します。

## テストデータの準備

まず、ベンチマークのためのデータを生成する必要があります。ここでは、ベンチマークからディスクIO時間を除くために、スパースファイルを使用します。現在、SingularityはCIDの重複削除を実行していないため、これらのファイルはランダムなバイトとして処理されます。

```sh
mkdir dataset
truncate -s 1024G dataset/1T.bin
```

ベンチマークにディスクIO時間を含める場合は、次の方法でランダムファイルを作成してください。

```sh
dd if=/dev/urandom of=dataset/8G.bin bs=1M count=8192
```

## `ez-prep`の使用
`ez-prep`コマンドは、最小限の設定可能なオプションを使用して、ローカルフォルダからデータの準備を簡素化します。

### インライン準備によるベンチマーク

インライン準備により、CARファイルのエクスポートの必要がなくなり、メタデータを直接データベースに保存できます。

```sh
time singularity ez-prep --output-dir '' ./dataset
```

### インメモリデータベースを使用したベンチマーク

ディスクIOを最小限に抑えるために、インメモリデータベースを選択します。

```sh
time singularity ez-prep --output-dir '' --database-file '' ./dataset
```

### 複数のワーカーを使用したベンチマーク

最適なCPUコアの利用を実現するために、ベンチマークの並列処理を設定します。注意：各ワーカーは約4つのCPUコアを使用します。

```sh
time singularity ez-prep --output-dir '' -j $(($(nproc) / 4 + 1)) ./dataset
```

## 結果の解釈

典型的な出力は次のようになります：

```
real    0m20.379s
user    0m44.937s
sys     0m8.981s
```

* `real`：実経過時間。より多くのワーカーを使用すると、この時間が短縮されます。
* `user`：ユーザースペースで使用されるCPU時間。`user`を`real`で割ると、使用されるCPUコアの数が近似されます。
* `sys`：カーネルスペースで使用されるCPU時間（ディスクIOを示します）。

## 比較

以下は、ランダムな8Gファイルに対して実行されたベンチマークの結果です：

<table><thead><tr><th width="290">ツール</th><th width="178.33333333333331" data-type="number">クロック時間（秒）</th><th data-type="number">CPU時間（秒）</th><th data-type="number">メモリ（KB）</th></tr></thead><tbody><tr><td>Singularity（インライン準備あり）</td><td>15.66</td><td>51.82</td><td>99</td></tr><tr><td>Singularity（インライン準備なし）</td><td>19.13</td><td>51.51</td><td>99</td></tr><tr><td>go-fil-dataprep</td><td>16.39</td><td>43.94</td><td>83</td></tr><tr><td>generate-car</td><td>42.6</td><td>56.08</td><td>44</td></tr><tr><td>go-car + stream-commp</td><td>70.21</td><td>139.01</td><td>42</td></tr></tbody></table>