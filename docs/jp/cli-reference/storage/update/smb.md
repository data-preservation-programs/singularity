# SMB / CIFS

{% code fullWidth="true" %}
```
名前:
   singularity storage update smb - SMB / CIFS

使用法:
   singularity storage update smb [コマンドオプション] <名前|ID>

説明:
   --host
      接続するSMBサーバーのホスト名。
      
      例: "example.com"。

   --user
      SMBユーザー名。

   --port
      SMBポート番号。

   --pass
      SMBパスワード。

   --domain
      NTLM認証のドメイン名。

   --spn
      サービスプリンシパル名。
      
      Rcloneはこの名前をサーバーに提供します。一部のサーバーはこれをさらなる認証として使用することがあり、クラスタに設定する必要があります。たとえば：

          cifs/remotehost:1020
      
      よくわからない場合は空白のままにしてください。
      

   --idle-timeout
      アイドル接続を終了する前の最大時間。
      
      指定された時間内に接続がコネクションプールに返されない場合、rcloneはコネクションプールを空にします。
      
      接続を無期限に保持するには、0に設定してください。
      

   --hide-special-share
      ユーザーがアクセスできない特別な共有（たとえばprint$）を非表示にする。

   --case-insensitive
      サーバーが大文字小文字を区別しないように設定されているかどうか。
      
      Windows共有の場合は常にtrueです。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --domain value  NTLM認証のドメイン名。 (デフォルト: "WORKGROUP") [$DOMAIN]
   --help, -h      ヘルプを表示する
   --host value    接続するSMBサーバーのホスト名。 [$HOST]
   --pass value    SMBパスワード。 [$PASS]
   --port value    SMBポート番号。 (デフォルト: 445) [$PORT]
   --spn value     サービスプリンシパル名。 [$SPN]
   --user value    SMBユーザー名。 (デフォルト: "$USER") [$USER]

   高度な設定

   --case-insensitive    サーバーが大文字小文字を区別しないように設定されているかどうか。 (デフォルト: true) [$CASE_INSENSITIVE]
   --encoding value      バックエンドのエンコーディング。 (デフォルト: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --hide-special-share  ユーザーがアクセスできない特別な共有（たとえばprint$）を非表示にする。 (デフォルト: true) [$HIDE_SPECIAL_SHARE]
   --idle-timeout value  アイドル接続を終了する前の最大時間。 (デフォルト: "1m0s") [$IDLE_TIMEOUT]

```
{% endcode %}