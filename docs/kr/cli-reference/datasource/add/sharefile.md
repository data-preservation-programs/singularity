# Citrix Sharefile

{% code fullWidth="true" %}
```
이름:
   singularity datasource add sharefile - Citrix Sharefile

사용법:
   singularity datasource add sharefile [command options] <데이터셋_이름> <소스_경로>

설명:
   --sharefile-chunk-size
       업로드 청크 크기입니다.
       
       256k 이상의 2의 거듭제곱이어야 합니다.
       
       이 값을 크게 설정하면 성능이 향상되지만, 각 청크는 한 번의 전송에 대해 메모리에 버퍼링됩니다.
       
       이 값을 줄이면 메모리 사용량은 줄어들지만 성능은 감소합니다.

   --sharefile-encoding
       백엔드의 인코딩입니다.
       
       자세한 정보는 [개요의 encoding 섹션](/overview/#encoding)을 참조하십시오.

   --sharefile-endpoint
       API 호출를 위한 엔드포인트입니다.
       
       일반적으로 OAuth 프로세스의 일부로 자동으로 검색되지만, https://XXX.sharefile.com과 같이 수동으로 설정할 수도 있습니다.

   --sharefile-root-folder-id
       루트 폴더의 ID입니다.
       
       "Personal Folders"에 액세스하려면 비워 두세요. 여기에서는 표준 값 중 하나이거나 (긴 16진수 ID) 아무 폴더 ID를 사용할 수 있습니다.

       예시:
          | <unset>    | "Personal Folders"에 액세스합니다 (기본값).
          | favorites  | Favorites 폴더에 액세스합니다.
          | allshared  | 모든 공유 폴더에 액세스합니다.
          | connectors | 각각의 연결자에 액세스합니다.
          | top        | 홈, Favorites 및 공유 폴더뿐만 아니라 연결자에 액세스합니다.

   --sharefile-upload-cutoff
       멀티파트 업로드로 전환하는 임계값입니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보낸 후 해당 파일 삭제 (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔 후 지정된 시간이 경과하면 소스 디렉토리를 자동으로 재스캔합니다 (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태 설정 (기본값: 준비)

   Sharefile 옵션

   --sharefile-chunk-size value      업로드 청크 크기 (기본값: "64Mi") [$SHAREFILE_CHUNK_SIZE]
   --sharefile-encoding value        백엔드의 인코딩 (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot") [$SHAREFILE_ENCODING]
   --sharefile-endpoint value        API 호출을 위한 엔드포인트 [$SHAREFILE_ENDPOINT]
   --sharefile-root-folder-id value  루트 폴더의 ID [$SHAREFILE_ROOT_FOLDER_ID]
   --sharefile-upload-cutoff value   멀티파트 업로드로 전환하는 임계값 (기본값: "128Mi") [$SHAREFILE_UPLOAD_CUTOFF]

```
{% endcode %}