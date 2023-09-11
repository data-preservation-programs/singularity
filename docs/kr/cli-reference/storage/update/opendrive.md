# OpenDrive

{% code fullWidth="true" %}
```
NAME:
   singularity storage update opendrive - OpenDrive

사용법:
   singularity storage update opendrive [command options] <name|id>

설명:
   --username
      사용자 이름.

   --password
      비밀번호.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --chunk-size
      파일이 이 크기로 청크 단위로 업로드됩니다.
      
      이러한 청크는 메모리에 버퍼링되므로 증가시키면
      메모리 사용량이 증가합니다.


옵션:
   --help, -h        도움말 표시
   --password value  비밀번호. [$PASSWORD]
   --username value  사용자 이름. [$USERNAME]

   Advanced

   --chunk-size value  파일이 이 크기로 청크 단위로 업로드됩니다. (기본값: "10Mi") [$CHUNK_SIZE]
   --encoding value    백엔드의 인코딩입니다. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot") [$ENCODING]

```
{% endcode %}