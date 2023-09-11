# Hadoop 分散ファイルシステム

{% code fullWidth="true" %}
```
NAME:
   singularity storage update hdfs - Hadoop 分散ファイルシステム

USAGE:
   singularity storage update hdfs [オプション] <名前|ID>

DESCRIPTION:
   --namenode
      Hadoop ネームノードとポート。

      例: "namenode:8020" は、ポート 8020 のホストネームノードに接続します。

   --username
      Hadoop ユーザー名。

      例:
         | root | root として HDFS に接続します。

   --service-principal-name
      ネームノードの Kerberos サービスプリンシパル名。

      KERBEROS 認証を有効にします。ネームノードのサービスプリンシパル名（SERVICE/FQDN）を指定します。
      例: "hdfs/namenode.hadoop.docker" は、'hdfs' という名前のサービスが FQDN 'namenode.hadoop.docker' で実行されるネームノードを指定します。

   --data-transfer-protection
      Kerberos データ転送保護: authentication|integrity|privacy。

      データノードと通信する際に、認証、データ署名の整合性チェック、ワイヤー暗号化が必要かどうかを指定します。
      可能な値は 'authentication'、'integrity'、'privacy' です。KERBEROS が有効な場合にのみ使用します。

      例:
         | privacy | 認証、整合性、暗号化が有効になっていることを確認します。

   --encoding
      バックエンドのエンコーディング。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

オプション:
   --help, -h        ヘルプを表示
   --namenode value  Hadoop ネームノードとポート。 [$NAMENODE]
   --username value  Hadoop ユーザー名。 [$USERNAME]

   Advanced

   --data-transfer-protection value  Kerberos データ転送保護: authentication|integrity|privacy。 [$DATA_TRANSFER_PROTECTION]
   --encoding value                  バックエンドのエンコーディング。 (デフォルト: "Slash,Colon,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --service-principal-name value    ネームノードの Kerberos サービスプリンシパル名。 [$SERVICE_PRINCIPAL_NAME]

```
{% endcode %}