# Microsoft OneDrive

## 概要

`singularity datasource add onedrive`コマンドは、OneDriveからデータをソースとして使用できるように設定するためのものです。

## 使用法

```
singularity datasource add onedrive [コマンドオプション] <データセット名> <ソースパス>
```

## 詳細

以下のオプションを使用して、OneDriveを構成します。

- `--onedrive-access-scopes`: rcloneがリクエストするスコープを設定します。
- `--onedrive-auth-url`: 認証サーバーのURLを設定します。
- `--onedrive-chunk-size`: ファイルをアップロードするためのチャンクサイズを設定します。
- `--onedrive-client-id`: OAuthクライアントIDを設定します。
- `--onedrive-client-secret`: OAuthクライアントシークレットを設定します。
- `--onedrive-disable-site-permission`: Sites.Read.All権限のリクエストを無効にします。
- `--onedrive-drive-id`: 使用するドライブのIDを設定します。
- `--onedrive-drive-type`: ドライブの種類（personal | business | documentLibrary）を設定します。
- `--onedrive-encoding`: バックエンドのエンコーディングを設定します。
- `--onedrive-expose-onenote-files`: OneNoteファイルをディレクトリリストに表示するように設定します。
- `--onedrive-hash-type`: バックエンドで使用するハッシュを指定します。
- `--onedrive-link-password`: リンクコマンドで作成されるリンクのパスワードを設定します。
- `--onedrive-link-scope`: リンクコマンドで作成されるリンクのスコープを設定します。
- `--onedrive-link-type`: リンクコマンドで作成されるリンクのタイプを設定します。
- `--onedrive-list-chunk`: リスティングチャンクのサイズを設定します。
- `--onedrive-no-versions`: 変更操作時にすべてのバージョンを削除します。
- `--onedrive-region`: OneDriveの国別クラウドリージョンを選択します。
- `--onedrive-root-folder-id`: ルートフォルダーのIDを設定します。
- `--onedrive-server-side-across-configs`: サーバーサイドの操作（コピーなど）が異なるonedrive構成間で動作するようにします。
- `--onedrive-token`: OAuthアクセストークンをJSONブロブで設定します。
- `--onedrive-token-url`: トークンサーバーのURLを設定します。