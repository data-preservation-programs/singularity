# premiumize.me

{% code fullWidth="true" %}
```
이름:
   singularity datasource add premiumizeme - premiumize.me

사용법:
   singularity datasource add premiumizeme [command options] <dataset_name> <source_path>

설명:
   --premiumizeme-api-key
      API 키입니다.
      
      이는 일반적으로 사용되지 않으며, 대신에 oauth를 사용하세요.
      

   --premiumizeme-encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보낸 후 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터 이 간격이 지나면 소스 디렉토리를 자동으로 다시 스캔합니다 (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다 (기본값: 준비 완료)

   premiumizeme을 위한 옵션

   --premiumizeme-encoding value  백엔드에 대한 인코딩입니다. (기본값: "Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PREMIUMIZEME_ENCODING]

```
{% endcode %}