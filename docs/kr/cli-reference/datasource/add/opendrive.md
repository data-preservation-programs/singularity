# OpenDrive

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add opendrive - OpenDrive

사용법:
   singularity datasource add opendrive [command options] <데이터셋_이름> <소스_경로>

설명:
   --opendrive-chunk-size
      파일은 이 크기로 청크 단위로 업로드됩니다.
      
      이 청크들은 메모리에 버퍼링되므로, 크기를 늘리면 메모리 사용량도 증가합니다.

   --opendrive-encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요 문서의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --opendrive-password
      비밀번호입니다.

   --opendrive-username
      사용자 이름입니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보낸 후 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막으로 성공한 스캔으로부터 지정된 시간이 경과하면 소스 디렉토리를 자동으로 스캔합니다. (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비 완료)

   OpenDrive 옵션

   --opendrive-chunk-size value  파일은 이 크기로 청크 단위로 업로드됩니다. (기본값: "10Mi") [$OPENDRIVE_CHUNK_SIZE]
   --opendrive-encoding value    백엔드에 대한 인코딩입니다. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot") [$OPENDRIVE_ENCODING]
   --opendrive-password value    비밀번호입니다. [$OPENDRIVE_PASSWORD]
   --opendrive-username value    사용자 이름입니다. [$OPENDRIVE_USERNAME]

```
{% endcode %}