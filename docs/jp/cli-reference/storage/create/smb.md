# SMB / CIFS

{% code fullWidth="true" %}
```
名前：
   singularity storage create smb - SMB / CIFS

使用方法：
   singularity storage create smb [コマンドオプション] [引数...]

説明：
   --host
      接続するSMBサーバーのホスト名です。
      
      例: "example.com".

   --user
      SMBのユーザー名です。

   --port
      SMBのポート番号です。

   --pass
      SMBのパスワードです。

   --domain
      NTLM認証のドメイン名です。

   --spn
      サービスプリンシパル名です。
      
      Rcloneはこの名前をサーバーに提示します。一部のサーバーでは、これをさらなる認証に使用することがあり、クラスタに設定する必要があります。例えば:
      
          cifs/remotehost:1020
      
      わからない場合は空白のままにしておきます。
      
   --idle-timeout
      アイドル接続を閉じる前の最大時間です。
      
      指定した時間内にコネクションプールに戻った接続がない場合、rcloneはコネクションプールを空にします。
      
      コネクションを無期限に保つには 0 を設定します。

   --hide-special-share
      ユーザーがアクセスできない特別な共有（例：print$）を非表示にします。

   --case-insensitive
      サーバーが大文字と小文字を区別しないように設定されているかどうかです。
      
      Windows共有の場合は常にtrueです。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、「[概要のエンコーディングセクション](/overview/#encoding)」を参照してください。

オプション：
   --domain value  NTLM認証のドメイン名です。 (default: "WORKGROUP") [$DOMAIN]
   --help, -h      ヘルプを表示します
   --host value    接続するSMBサーバーのホスト名です。 [$HOST]
   --pass value    SMBのパスワードです。 [$PASS]
   --port value    SMBのポート番号です。 (default: 445) [$PORT]
   --spn value     サービスプリンシパル名です。 [$SPN]
   --user value    SMBのユーザー名です。 (default: "$USER") [$USER]

   上級設定

   --case-insensitive    サーバーが大文字と小文字を区別しないように設定されているかどうかです。 (default: true) [$CASE_INSENSITIVE]
   --encoding value      バックエンドのエンコーディングです。 (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --hide-special-share  ユーザーがアクセスできない特別な共有（例：print$）を非表示にします。 (default: true) [$HIDE_SPECIAL_SHARE]
   --idle-timeout value  アイドル接続を閉じる前の最大時間です。 (default: "1m0s") [$IDLE_TIMEOUT]

   一般的な設定

   --name value  ストレージの名前 (default: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}