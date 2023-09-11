# Uptobox

{% code fullWidth="true" %}
```
名前:
   singularity storage create uptobox - Uptobox

使用法:
   singularity storage create uptobox [コマンドオプション] [引数...]

説明:
   --access-token
      アクセストークンです。
      
      https://uptobox.com/my_account から取得してください。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --access-token value  アクセストークンです。 [$ACCESS_TOKEN]
   --help, -h            ヘルプを表示します

   高度な設定

   --encoding value  バックエンドのエンコーディングです。 (デフォルト: "Slash, LtGt, DoubleQuote, BackQuote, Del, Ctl, LeftSpace, InvalidUtf8, Dot") [$ENCODING]

   共通設定

   --name value  ストレージの名前です (デフォルト: 自動生成)
   --path value  ストレージのパスです

```
{% endcode %}