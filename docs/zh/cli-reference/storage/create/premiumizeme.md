# premiumize.me

{% code fullWidth="true" %}
```
命令：
   singularity storage create premiumizeme - premiumize.me

用法：
   singularity storage create premiumizeme [命令选项] [参数...]

描述：
   --api-key
      API密钥。
      
      通常不使用此选项 - 请改用oauth。
      

   --encoding
      后端的编码。
      
      参见[概览中的编码部分](/overview/#encoding)了解更多信息。


选项：
   --api-key value  API密钥。[$API_KEY]
   --help, -h       显示帮助信息

   高级选项

   --encoding value  后端的编码。（默认值："Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot"）[$ENCODING]

   常规选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}