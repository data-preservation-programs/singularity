# Hadoop分散ファイルシステム

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add hdfs - Hadoop分散ファイルシステム

USAGE:
   singularity datasource add hdfs [コマンドオプション] <データセット名> <ソースパス>

DESCRIPTION:
   --hdfs-data-transfer-protection
      Kerberosデータ転送保護: authentication|integrity|privacy.
      
      データノードとやり取りする際に、認証、データ署名整合性のチェック、
      ワイヤー暗号化が必要かどうかを指定します。
      使用可能な値は'authentication'、'integrity'、'privacy'です。
      KERBEROSが有効な場合にのみ使用されます。

      例:
         | privacy | 認証、整合性、暗号化が有効になっていることを保証する。

   --hdfs-encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --hdfs-namenode
      Hadoopネームノードとポートです。
      
      例：ホストnamenodeのポート8020に接続する場合は、"namenode:8020"とします。

   --hdfs-service-principal-name
      ネームノードのKerberosサービスプリンシパル名です。
      
      KERBEROS認証を有効にします。
      ネームノードのService Principal Name（SERVICE/FQDN）を指定します。
      例：FQDNが'namenode.hadoop.docker'でサービスが'hdfs'として実行される場合は、
      "hdfs/namenode.hadoop.docker"とします。

   --hdfs-username
      Hadoopユーザー名です。

      例:
         | root | rootとしてhdfsに接続する。

OPTIONS:
   --help, -h  ヘルプを表示

   データ準備オプション

   --delete-after-export    [危険] データセットのファイルをエクスポートした後、ファイルを削除します。  (デフォルト: false)
   --rescan-interval value  最後の成功したスキャンからこの間隔が経過すると自動的にソースディレクトリを再スキャンします (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します (デフォルト: ready)

   hdfs向けオプション

   --hdfs-data-transfer-protection value  Kerberosデータ転送保護: authentication|integrity|privacy. [$HDFS_DATA_TRANSFER_PROTECTION]
   --hdfs-encoding value                  バックエンドのエンコーディングです。 (デフォルト: "Slash,Colon,Del,Ctl,InvalidUtf8,Dot") [$HDFS_ENCODING]
   --hdfs-namenode value                  Hadoopネームノードとポートです。 [$HDFS_NAMENODE]
   --hdfs-service-principal-name value    ネームノードのKerberosサービスプリンシパル名です。 [$HDFS_SERVICE_PRINCIPAL_NAME]
   --hdfs-username value                  Hadoopユーザー名です。 [$HDFS_USERNAME]
```
{% endcode %}