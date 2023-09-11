# Citrix Sharefile

{% code fullWidth="true" %}
```
이름:
   singularity storage update sharefile - Citrix Sharefile

사용법:
   singularity storage update sharefile [command options] <이름|ID>

설명:
   --upload-cutoff
      멀티파트 업로드로 전환되는 커트오프.

   --root-folder-id
      루트 폴더의 ID.
      
      "개인 폴더"에 액세스하려면 비워 두십시오. 여기에는 표준 값을 사용하거나 (긴 16진수 ID) 아무 폴더 ID를 사용할 수 있습니다.

      예:
         | <미설정>    | 개인 폴더에 액세스 (기본값).
         | 즐겨찾기    | 즐겨찾기 폴더에 액세스.
         | 모두공유    | 모든 공유 폴더에 액세스.
         | 커넥터      | 개별 커넥터에 액세스.
         | 상위        | 홈, 즐겨찾기, 공유 폴더 및 커넥터에 액세스.

   --chunk-size
      업로드 청크 크기.
      
      256k 이상이고 2의 거듭제곱인 값이어야 합니다.
      
      이 값을 더 크게 설정하면 성능이 향상되지만, 각 청크는 전송 당 한 번씩 메모리에 버퍼링됩니다.
      
      이 값을 줄이면 메모리 사용량이 감소하지만 성능이 감소합니다.

   --endpoint
      API 호출을 위한 엔드포인트.
      
      일반적으로 oauth 프로세스의 일부로 자동으로 검색되지만, https://XXX.sharefile.com과 같이 수동으로 설정할 수도 있습니다.
      

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --help, -h              도움말 표시
   --root-folder-id 값     루트 폴더의 ID. [$ROOT_FOLDER_ID]

   고급 옵션

   --chunk-size 값      업로드 청크 크기. (기본값: "64Mi") [$CHUNK_SIZE]
   --encoding 값        백엔드의 인코딩. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --endpoint 값        API 호출을 위한 엔드포인트. [$ENDPOINT]
   --upload-cutoff 값   멀티파트 업로드로 전환되는 커트오프. (기본값: "128Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}