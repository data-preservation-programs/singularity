# Hadoop分散ファイルシステム

{% code fullWidth="true" %}
```
NAME:
   singularity storage create hdfs - Hadoop分散ファイルシステム

USAGE:
   singularity storage create hdfs [command options] [arguments...]

DESCRIPTION:
   --namenode
      Hadoop名ノードとポート。

      例: ポート8020でホスト名ノードに接続する場合は「namenode:8020」。

   --username
      Hadoopユーザ名。

      例:
         | root | rootとしてhdfsに接続する。

   --service-principal-name
      ネームノードのKerberosサービス主体名。

      KERBEROS認証を有効にします。サービスプリンシパル名（SERVICE/FQDN）を指定します。
      例: 「hdfs/namenode.hadoop.docker」は、サービスが「hdfs」でFQDNが「namenode.hadoop.docker」
      のネームノードを指します。

   --data-transfer-protection
      Kerberosデータ転送保護: authentication|integrity|privacy。

      データノード間の通信において、認証、データ署名の整合性チェック、ワイヤ暗号化が必要かどうかを指定します。
      可能な値は「authentication」、「integrity」および「privacy」です。KERBEROSが有効な場合のみ使用します。

      例:
         | privacy | 認証、整合性、暗号化が有効になっていることを保証します。

   --encoding
      バックエンドのエンコーディング。

      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --help, -h        ヘルプを表示
   --namenode value  Hadoop名ノードとポート。 [$NAMENODE]
   --username value  Hadoopユーザ名。 [$USERNAME]

   Advanced

   --data-transfer-protection value  Kerberosデータ転送保護: authentication|integrity|privacy。 [$DATA_TRANSFER_PROTECTION]
   --encoding value                  バックエンドのエンコーディング。 (デフォルト: "Slash,Colon,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --service-principal-name value    ネームノードのKerberosサービス主体名。 [$SERVICE_PRINCIPAL_NAME]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}