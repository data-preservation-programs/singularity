# SMB / CIFS

{% code fullWidth="true" %}
```
名前:
   singularity datasource add smb - SMB / CIFS

使用方法:
   singularity datasource add smb [command options] <dataset_name> <source_path>

説明:
   --smb-case-insensitive
      サーバーが大文字と小文字を区別しないように設定されているかどうか。

      Windows共有では常にtrueです。

   --smb-domain
      NTLM認証のドメイン名。

   --smb-encoding
      バックエンドのエンコーディング。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --smb-hide-special-share
      アクセスが許可されていない特殊共有（例：print$）を非表示にします。

   --smb-host
      接続するSMBサーバーのホスト名。

      例: "example.com"。

   --smb-idle-timeout
      アイドル接続を閉じるまでの最大時間。

      指定された時間内に接続がコネクションプールに返されない場合、rcloneはコネクションプールを空にします。

      接続を無期限に保持するには、0に設定します。

   --smb-pass
      SMBのパスワード。

   --smb-port
      SMBポート番号。

   --smb-spn
      サービスプリンシパル名。

      Rcloneはこの名前をサーバーに提示します。一部のサーバーではこれをさらなる認証として使用し、クラスタによく設定する必要があります。例：

          cifs/remotehost:1020

      わからない場合は空白のままにしてください。

   --smb-user
      SMBのユーザー名。

オプション:
   --help, -h  ヘルプを表示する

   データの準備オプション

   --delete-after-export    [危険] データセットのファイルをエクスポートした後、CARファイルにエクスポートする前に削除します。  (デフォルト: false)
   --rescan-interval value  最後の成功したスキャンからこの間隔が経過すると、ソースディレクトリを自動的に再スキャンします (デフォルト: disabled)
   --scanning-state value   初期のスキャン状態を設定します (デフォルト: ready)

   smbのオプション

   --smb-case-insensitive value    サーバーが大文字と小文字を区別しないように設定されているかどうか。 (デフォルト: "true") [$SMB_CASE_INSENSITIVE]
   --smb-domain value              NTLM認証のドメイン名。 (デフォルト: "WORKGROUP") [$SMB_DOMAIN]
   --smb-encoding value            バックエンドのエンコーディング。 (デフォルト: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot") [$SMB_ENCODING]
   --smb-hide-special-share value  アクセスが許可されていない特殊共有（例：print$）を非表示にします。 (デフォルト: "true") [$SMB_HIDE_SPECIAL_SHARE]
   --smb-host value                接続するSMBサーバーのホスト名。 [$SMB_HOST]
   --smb-idle-timeout value        アイドル接続を閉じるまでの最大時間。 (デフォルト: "1m0s") [$SMB_IDLE_TIMEOUT]
   --smb-pass value                SMBのパスワード。 [$SMB_PASS]
   --smb-port value                SMBポート番号。 (デフォルト: "445") [$SMB_PORT]
   --smb-spn value                 サービスプリンシパル名。 [$SMB_SPN]
   --smb-user value                SMBのユーザー名。 (デフォルト: "$USER") [$SMB_USER]

```
{% endcode %}