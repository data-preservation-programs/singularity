# Google 사진

{% code fullWidth="true" %}
```
이름:
   싱귤래리티 저장소 생성 gphotos - Google 사진

사용법:
   singularity storage create gphotos [command options] [arguments...]

설명:
   --client-id
      OAuth 클라이언트 ID입니다.
      
      보통 비워둡니다.

   --client-secret
      OAuth 클라이언트 비밀입니다.
      
      보통 비워둡니다.

   --token
      JSON blob 형태의 OAuth 액세스 토큰입니다.

   --auth-url
      인증 서버 URL입니다.
      
      제공자 기본값을 사용하려면 비워둡니다.

   --token-url
      토큰 서버 URL입니다.
      
      제공자 기본값을 사용하려면 비워둡니다.

   --read-only
      Google 사진 백엔드를 읽기 전용으로 설정합니다.
      
      읽기 전용으로 선택하면 rclone은 사진에 대해 읽기 전용 액세스만 요청하며,
      그렇지 않으면 rclone은 전체 액세스를 요청합니다.

   --read-size
      미디어 항목의 크기를 읽을 수 있도록 설정합니다.
      
      보통 rclone은 미디어 항목의 크기를 읽지 않습니다.
      이 작업을 수행하면 다른 트랜잭션이 필요하기 때문입니다.
      동기화에는 이것이 필요하지 않습니다.
      그러나 rclone mount는 파일의 크기를 미리 알아야 하므로,
      rclone mount를 사용하는 경우 이 플래그를 설정하는 것이 좋습니다.

   --start-year
      다운로드할 사진을 지정된 연도 이후에 업로드된 사진으로 제한합니다.

   --include-archived
      보관된 미디어를 봅니다.
      
      기본적으로 rclone은 보관된 미디어를 요청하지 않습니다.
      따라서 동기화할 때 디렉토리 목록 또는 전송에서 보관된 미디어는 보이지 않습니다.
      
      앨범 내 미디어는 보관 상태에 관계없이 항상 보이고 동기화됩니다.
      
      이 플래그를 사용하면 디렉토리 목록에 보관된 미디어가 항상 표시되고 전송됩니다.
      
      이 플래그를 사용하지 않으면 디렉토리 목록에서 보관된 미디어가 표시되지 않고 전송되지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --client-id value      OAuth 클라이언트 ID입니다. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 비밀입니다. [$CLIENT_SECRET]
   --help, -h             도움말 표시
   --read-only            Google 사진 백엔드를 읽기 전용으로 설정합니다. (기본값: false) [$READ_ONLY]

   고급

   --auth-url value    인증 서버 URL입니다. [$AUTH_URL]
   --encoding value    백엔드의 인코딩입니다. (기본값: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --include-archived  보관된 미디어도 보고 다운로드합니다. (기본값: false) [$INCLUDE_ARCHIVED]
   --read-size         미디어 항목의 크기를 읽을 수 있도록 설정합니다. (기본값: false) [$READ_SIZE]
   --start-year value  다운로드할 사진을 지정된 연도 이후에 업로드된 사진으로 제한합니다. (기본값: 2000) [$START_YEAR]
   --token value       JSON blob 형태의 OAuth 액세스 토큰입니다. [$TOKEN]
   --token-url value   토큰 서버 URL입니다. [$TOKEN_URL]

   기본

   --name value  저장소의 이름 (기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}