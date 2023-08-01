# Put.io

{% code fullWidth="true" %}
```
이름:
   singularity datasource add putio - Put.io

사용법:
   singularity datasource add putio [command options] <데이터셋_이름> <소스_경로>

설명:
   --putio-encoding
      백엔드에 대한 인코딩입니다.

      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] 데이터셋을 CAR 파일로 내보낸 후에 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔 이후에 소스 디렉토리를 자동으로 다시 스캔할 간격을 설정합니다. (기본값: 해제됨)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비)

   putio에 대한 옵션

   --putio-encoding value  백엔드의 인코딩입니다. (기본값: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PUTIO_ENCODING]

```
{% endcode %}