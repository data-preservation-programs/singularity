# Citrix Sharefile

{% code fullWidth="true" %}
```
이름:
   singularity storage create sharefile - Citrix Sharefile

사용법:
   singularity storage create sharefile [command options] [arguments...]

설명:
   --upload-cutoff
      멀티파트 업로드로 전환하는 기준 설정.

   --root-folder-id
      최상위 폴더의 ID입니다.
      
      "Personal Folders"에 액세스하기 위해 비워 둘 수 있습니다. 여기에는 표준 값 중 하나 또는 폴더 ID (16진수로 이루어진 긴 숫자 ID)를 사용할 수 있습니다.

      예시:
         | <미설정>   | Personal Folders에 액세스 (기본값).
         | favorites  | Favorites 폴더에 액세스.
         | allshared  | 모든 공유 폴더에 액세스.
         | connectors | 개별 커넥터에 액세스.
         | top        | 홈, Favorites, 공유 폴더 및 커넥터에 액세스.

   --chunk-size
      업로드 청크 크기입니다.
      
      256KB 이상인 2의 제곱수여야 합니다.
      
      이 값을 더 크게 설정하면 성능이 향상됩니다. 그러나 각 청크는 전송 당 하나씩 메모리에 버퍼링됩니다.
      
      이 값을 줄이면 메모리 사용량이 줄어들지만 성능이 감소합니다.

   --endpoint
      API 호출을 위한 엔드포인트입니다.
      
      보통 OAuth 프로세스 중에 자동으로 검색됩니다만, 수동으로 https://XXX.sharefile.com과 같은 값으로 설정할 수 있습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --help, -h              도움말 표시
   --root-folder-id value  최상위 폴더의 ID입니다. [$ROOT_FOLDER_ID]

   Advanced

   --chunk-size value     업로드 청크 크기입니다. (기본값: "64Mi") [$CHUNK_SIZE]
   --encoding value       백엔드의 인코딩입니다. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value       API 호출을 위한 엔드포인트입니다. [$ENDPOINT]
   --upload-cutoff value  멀티파트 업로드로 전환하는 기준입니다. (기본값: "128Mi") [$UPLOAD_CUTOFF]

   General

   --name value  스토리지의 이름(기본값: 자동으로 생성됨)
   --path value  스토리지의 경로

```
{% endcode %}