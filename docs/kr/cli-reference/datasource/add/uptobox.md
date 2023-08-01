# Uptobox

{% code fullWidth="true" %}
```
이름:
   singularity datasource add uptobox - Uptobox

사용법:
   singularity datasource add uptobox [옵션] <데이터셋_이름> <소스_경로>

설명:
   --uptobox-access-token
      액세스 토큰을 입력하세요.
      
      https://uptobox.com/my_account에서 확인하실 수 있습니다.

   --uptobox-encoding
      백엔드용 인코딩을 입력하세요.
      
      자세한 내용은 [개요 섹션의 인코딩](/overview/#encoding)을 참조하세요.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] CAR 파일로 내보낸 후 데이터셋의 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 스캔 이후 이 시간 간격이 지나면 소스 디렉토리를 자동으로 다시 스캔합니다. (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비됨)

   uptobox용 옵션

   --uptobox-access-token value  액세스 토큰을 입력하세요. [$UPTOBOX_ACCESS_TOKEN]
   --uptobox-encoding value      백엔드용 인코딩을 입력하세요. (기본값: "Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot") [$UPTOBOX_ENCODING]

```
{% endcode %}