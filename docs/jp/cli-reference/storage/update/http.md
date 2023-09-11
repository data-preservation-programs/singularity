# HTTP

{% code fullWidth="true" %}
```
名前:
   singularity storage update http - HTTP

使用法:
   singularity storage update http [コマンドオプション] <名前|ID>

説明:
   --url
      接続するHTTPホストのURLです。
      
      例：「https://example.com」、または「https://user:pass@example.com」はユーザー名とパスワードを使用します。

   --headers
      すべてのトランザクションに対するHTTPヘッダーを設定します。
      
      追加のHTTPヘッダーを設定するために使用します。
      
      入力形式は、キー、値のペアのコンマ区切りリストです。[CSV エンコーディング](https://godoc.org/encoding/csv)を使用できます。
      
      たとえば、Cookieを設定するには「Cookie,name=value」または「"Cookie","name=value"」とします。
      
      複数のヘッダーを設定できます。たとえば、「"Cookie","name=value","Authorization","xxx"」です。

   --no-slash
      サイトがディレクトリの末尾に/を使用していない場合に設定します。
      
      対象のウェブサイトがディレクトリの末尾に/を使用しない場合に使用します。
      
      パスの末尾に/があることは、通常、rcloneがファイルとディレクトリを区別する方法です。このフラグが設定されている場合、rcloneはすべてのContent-Type: text/htmlをディレクトリとみなし、ダウンロードする代わりにそこからURLを読み取ります。
      
      ただし、これにより、rcloneは本物のHTMLファイルをディレクトリと間違えることがあります。

   --no-head
      HEADリクエストを使用しない。
      
      HEADリクエストは、ディレクトリリスト内のファイルサイズを検索するために主に使用されます。
      サイトの読み込みが非常に遅い場合、このオプションを試すことができます。
      通常、rcloneはディレクトリリスト内の各潜在的なファイルに対してHEADリクエストを行って以下を確認します：
      
      - サイズを取得する
      - 実際に存在するか確認する
      - ディレクトリであるかどうかを確認する
      
      このオプションを設定すると、rcloneはHEADリクエストを行いません。これによりディレクトリリストがはるかに速くなりますが、rcloneはファイルの時刻またはサイズを持たず、存在しないファイルがリストに含まれる場合があります。


オプション:
   --help, -h   ヘルプを表示
   --url value  接続するHTTPホストのURL [$URL]

   Advanced

   --headers value  すべてのトランザクションに対するHTTPヘッダーを設定 [$HEADERS]
   --no-head        HEADリクエストを使用しない (デフォルト: false) [$NO_HEAD]
   --no-slash       サイトがディレクトリの末尾に/を使用しない場合に設定 (デフォルト: false) [$NO_SLASH]
```
{% endcode %}