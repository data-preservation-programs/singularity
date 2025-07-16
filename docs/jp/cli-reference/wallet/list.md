# インポートされたすべてのウォレットをリストします

{% code fullWidth="true" %}
```
NAME:
   singularity wallet list - インポートされたすべてのウォレットをリストする

使用法:
   singularity wallet list [コマンドオプション] [引数...]

オプション:
   --with-balance  Lotusからライブウォレット残高を取得して表示する
   --help, -h      ヘルプを表示する
```
{% endcode %}

## 例: 残高付きでウォレットをリスト

```
singularity wallet list --with-balance --lotus-api <API> --lotus-token <TOKEN>

ADDRESS                                 BALANCE        DATACAP
f1abcde7zd3lfsv43aj2kb454ymaqw7debhumjxyz  1.000000 FIL   0
...
```

`--with-balance` フラグを指定すると、各ウォレットのFIL残高とFIL+データキャップが表示されます。