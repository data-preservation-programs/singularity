# HTTP

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add http - HTTP

USAGE:
   singularity datasource add http [command options] <dataset_name> <source_path>

DESCRIPTION:
   --http-headers
      すべてのトランザクションに対してHTTPヘッダーを設定します。
      
      これを使用して、すべてのトランザクションに追加のHTTPヘッダーを設定できます。
      
      入力形式は、キーと値のペアのカンマ区切りリストです。
      [CSVエンコーディング](https://godoc.org/encoding/csv) を使用できます。
      
      たとえば、Cookieを設定する場合は 'Cookie,name=value' または '"Cookie","name=value"' を使用します。
      
      複数のヘッダーを設定することもできます。例えば、'"Cookie","name=value","Authorization","xxx"' のようにしてください。

   --http-no-head
      HEADリクエストを行わないようにします。
      
      HEADリクエストは、ディレクトリリスト内のファイルサイズを検索するために主に使用されます。
      サイトの読み込みが非常に遅い場合は、このオプションを試してみてください。
      通常、rcloneはディレクトリリスト内の各潜在的なファイルに対してHEADリクエストを行って、以下のことを確認します:
      
      - サイズを取得する
      - 実際に存在するかどうかを確認する
      - ディレクトリであるかどうかを確認する
      
      このオプションを設定すると、rcloneはHEADリクエストを行わなくなります。これにより、ディレクトリリストの読み込みがはるかに高速化されますが、rcloneにはファイルの時刻やサイズがなくなり、リスト内に存在しないファイルがあるかもしれません。

   --http-no-slash
      サイトがディレクトリを/で終了していない場合に設定します。
      
      ターゲットのウェブサイトがディレクトリの末尾に/を使用していない場合に使用します。
      
      パスの末尾に/があると、rcloneは通常、ファイルとディレクトリの違いを伝えるために使用します。
      このフラグが設定されている場合、rcloneはContent-Type: text/htmlを持つすべてのファイルをディレクトリと見なし、それらからURLを読み取るようになります。

      注意: これにより、rcloneが本物のHTMLファイルをディレクトリと混同する可能性があります。

   --http-url
      接続するHTTPホストのURLです。
      
      例: "https://example.com"、または "https://user:pass@example.com" (ユーザー名とパスワードを使用する場合)。

OPTIONS:
   --help, -h  ヘルプを表示

   データ準備オプション
   
   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、データセットのファイルを削除します。 (デフォルト: false)
   --rescan-interval value  前回の成功したスキャンから一定の間隔が経過すると、ソースディレクトリを自動的に再スキャンします (デフォルト: 無効)
   --scanning-state value   初期スキャン状態を設定します (デフォルト: ready)

   httpのオプション

   --http-headers value   すべてのトランザクションに対してHTTPヘッダーを設定します。 [$HTTP_HEADERS]
   --http-no-head value   HEADリクエストを行わないようにします。 (デフォルト: "false") [$HTTP_NO_HEAD]
   --http-no-slash value  サイトがディレクトリを/で終了していない場合に設定します。 (デフォルト: "false") [$HTTP_NO_SLASH]
   --http-url value       接続するHTTPホストのURLです。 [$HTTP_URL]

```
{% endcode %}