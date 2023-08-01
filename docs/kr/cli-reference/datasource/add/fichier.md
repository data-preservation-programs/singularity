# 1Fichier

{% code fullWidth="true" %}
```
이름:
   singularity datasource add fichier - 1Fichier

사용법:
   singularity datasource add fichier [command options] <데이터셋_이름> <소스_경로>

설명:
   --fichier-api-key
      API 키입니다. https://1fichier.com/console/params.pl에서 얻으실 수 있습니다.

   --fichier-encoding
      백엔드의 인코딩 방식입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --fichier-file-password
      공유 파일을 다운로드하려면 이 매개변수를 추가합니다.

   --fichier-folder-password
      공유 폴더의 파일 목록을 확인하려면 이 매개변수를 추가합니다.

   --fichier-shared-folder
      공유 폴더를 다운로드하려면 이 매개변수를 추가합니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] CAR 파일로 내보낸 후 데이터셋의 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터 이 간격이 경과하면 자동으로 소스 디렉토리를 다시 스캔합니다. (기본값: 비활성화됨)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비)

   fichier을 위한 옵션

   --fichier-api-key value          API 키입니다. https://1fichier.com/console/params.pl에서 얻으실 수 있습니다. [$FICHIER_API_KEY]
   --fichier-encoding value         백엔드의 인코딩 방식입니다. (기본값: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$FICHIER_ENCODING]
   --fichier-file-password value    공유 파일을 다운로드하려면 이 매개변수를 추가합니다. [$FICHIER_FILE_PASSWORD]
   --fichier-folder-password value  공유 폴더의 파일 목록을 확인하려면 이 매개변수를 추가합니다. [$FICHIER_FOLDER_PASSWORD]
   --fichier-shared-folder value    공유 폴더를 다운로드하려면 이 매개변수를 추가합니다. [$FICHIER_SHARED_FOLDER]

```
{% endcode %}