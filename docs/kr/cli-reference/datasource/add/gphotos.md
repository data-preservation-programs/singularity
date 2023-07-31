# Google Photos

{% code fullWidth="true" %}
```
이름:
   singularity 데이터원 추가 gphotos - 구글 포토

사용법:
   singularity 데이터원 추가 gphotos [옵션] <데이터셋_이름> <원본_경로>

설명:
   --gphotos-auth-url
      인증 서버 URL.
      
      기본값을 사용하려면 비워 두세요.

   --gphotos-client-id
      OAuth 클라이언트 ID.
      
      보통 비워 둡니다.

   --gphotos-client-secret
      OAuth 클라이언트 비밀.
      
      보통 비워 둡니다.

   --gphotos-encoding
      백엔드의 인코딩 설정.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --gphotos-include-archived
      보관된 미디어도 표시하고 다운로드합니다.
      
      기본적으로, rclone은 보관된 미디어를 요청하지 않습니다. 따라서 동기화할 때
      디렉터리 목록이나 전송에 보존된 미디어가 표시되지 않습니다.
      
      앨범에 있는 미디어는 보관 상태에 관계없이 항상 표시되고 동기화됩니다.
      
      이 플래그를 사용하면 디렉터리 목록에 보관된 미디어가 항상 표시되고
      전송됩니다.
      
      이 플래그가 없으면 보관된 미디어는 디렉터리 목록에 표시되지 않으며
      전송되지 않습니다.

   --gphotos-read-only
      구글 포토 백엔드를 읽기 전용으로 설정합니다.
      
      읽기 전용으로 선택하면 rclone은 사진에 대해 읽기 전용 액세스만 요청합니다.
      그렇지 않으면 rclone은 전체 액세스를 요청합니다.

   --gphotos-read-size
      미디어 항목의 크기를 읽도록 설정합니다.
      
      일반적으로 rclone은 미디어 항목의 크기를 읽지 않습니다. 크기를 읽는
      작업은 다른 트랜잭션을 필요로 하기 때문입니다. 이는 동기화에는 필요하지
      않습니다. 그러나 rclone mount는 파일을 읽기 전에 파일의 크기를
      미리 알아야 하므로 rclone mount를 사용할 때 이 플래그를 설정하는 것이
      좋습니다.

   --gphotos-start-year
      다운로드할 사진을 주어진 연도 이후에 업로드된 사진으로 제한합니다.

   --gphotos-token
      OAuth 액세스 토큰으로 JSON blob 형식입니다.

   --gphotos-token-url
      토큰 서버 URL.
      
      기본값을 사용하려면 비워 두세요.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] 데이터셋을 CAR 파일로 내보낸 후 파일 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터 경과한 시간이 지나면 소스 디렉터리를 자동으로 다시 스캔합니다. (기본값: 사용 안 함)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비)

   gphotos용 옵션

   --gphotos-auth-url value          인증 서버 URL. [$GPHOTOS_AUTH_URL]
   --gphotos-client-id value         OAuth 클라이언트 ID. [$GPHOTOS_CLIENT_ID]
   --gphotos-client-secret value     OAuth 클라이언트 비밀. [$GPHOTOS_CLIENT_SECRET]
   --gphotos-encoding value          백엔드의 인코딩 설정. (기본값: "Slash,CrLf,InvalidUtf8,Dot") [$GPHOTOS_ENCODING]
   --gphotos-include-archived value  보관된 미디어도 표시하고 다운로드합니다. (기본값: "false") [$GPHOTOS_INCLUDE_ARCHIVED]
   --gphotos-read-only value         구글 포토 백엔드를 읽기 전용으로 설정합니다. (기본값: "false") [$GPHOTOS_READ_ONLY]
   --gphotos-read-size value         미디어 항목의 크기를 읽도록 설정합니다. (기본값: "false") [$GPHOTOS_READ_SIZE]
   --gphotos-start-year value        다운로드할 사진을 주어진 연도 이후에 업로드된 사진으로 제한합니다. (기본값: "2000") [$GPHOTOS_START_YEAR]
   --gphotos-token value             OAuth 액세스 토큰으로 JSON blob 형식입니다. [$GPHOTOS_TOKEN]
   --gphotos-token-url value         토큰 서버 URL. [$GPHOTOS_TOKEN_URL]

```
{% endcode %}