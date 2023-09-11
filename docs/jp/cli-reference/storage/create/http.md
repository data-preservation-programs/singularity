# HTTP

{% code fullWidth="true" %}
```
名前:
   singularity storage create http - HTTP

使用法:
   singularity storage create http [コマンドオプション] [引数...]

説明:
   --url
      接続するHTTPホストのURLです。
      
      例： "https://example.com"、または "https://user:pass@example.com" でユーザー名とパスワードを使用する場合。

   --headers
      すべてのトランザクションのためのHTTPヘッダーを設定します。
      
      これを使用して、すべてのトランザクションのための追加のHTTPヘッダーを設定します。
      
      入力形式はキー、値のペアのカンマ区切りリストです。標準の[CSVエンコーディング](https://godoc.org/encoding/csv)が使用できます。
      
      たとえば、Cookieを設定する場合は 'Cookie,name=value' や '"Cookie","name=value"' を使用します。
      
      複数のヘッダーを設定することもできます。たとえば、'"Cookie","name=value","Authorization","xxx"' のようにします。

   --no-slash
      サイトがディレクトリを / で終えない場合に設定します。
      
      ターゲットのウェブサイトがディレクトリの末尾に / を使用しない場合に使用します。
      
      パスの末尾の / は、通常のrcloneがファイルとディレクトリの違いを伝える方法です。このフラグが設定されている場合、rcloneはContent-Type: text/htmlを持つすべてのファイルをディレクトリと見なし、ダウンロードではなくそれらからURLを読み込みます。
      
      ただし、これによりrcloneが正当なHTMLファイルをディレクトリと混同する可能性があります。

   --no-head
      HEADリクエストを使用しないでください。
      
      HEADリクエストは、ディレクトリリストでファイルサイズを見つけるために主に使用されます。
      サイトの読み込みが非常に遅い場合は、このオプションを試してみることができます。
      通常、rcloneはディレクトリリスト内の各潜在的なファイルに対して次のことを行います。
      
      - サイズを見つける
      - 実際に存在するかチェックする
      - ディレクトリであるかをチェックする
      
      このオプションを設定すると、rcloneはHEADリクエストを実行しません。これにより、ディレクトリリストの作成速度が向上しますが、rcloneは各ファイルの時間やサイズを持たず、存在しないファイルがリストに含まれる場合があります。


オプション:
   --help, -h   ヘルプを表示する
   --url value  接続するHTTPホストのURLです。 [$URL]

   Advanced

   --headers value  すべてのトランザクションのためのHTTPヘッダーを設定します。 [$HEADERS]
   --no-head        HEADリクエストを使用しないでください。 (デフォルト: false) [$NO_HEAD]
   --no-slash       サイトがディレクトリを / で終えない場合に設定します。 (デフォルト: false) [$NO_SLASH]

   General

   --name value  ストレージの名前 (デフォルト: 自動生成)
   --path value  ストレージのパス
```
{% endcode %}