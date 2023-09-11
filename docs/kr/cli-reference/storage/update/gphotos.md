# Google Photos

{% code fullWidth="true" %}
```
이름:
   singularity storage update gphotos - Google 사진

사용법:
   singularity storage update gphotos [command options] <name|id>

설명:
   --client-id
      OAuth 클라이언트 ID입니다.
      
      보통 비워두십시오.

   --client-secret
      OAuth 클라이언트 비밀입니다.
      
      보통 비워두십시오.

   --token
      JSON blob 형태의 OAuth 액세스 토큰입니다.

   --auth-url
      인증 서버 URL입니다.
      
      공급자 기본값을 사용하려면 비워두십시오.

   --token-url
      토큰 서버 URL입니다.
      
      공급자 기본값을 사용하려면 비워두십시오.

   --read-only
      Google 사진 백엔드를 읽기 전용으로 설정합니다.
      
      읽기 전용으로 설정하면 rclone은 사진에 대해 읽기 전용 액세스만 요청하고,
      그렇지 않으면 rclone은 전체 액세스를 요청합니다.

   --read-size
      미디어 항목의 크기를 읽도록 설정합니다.
      
      일반적으로 rclone은 미디어 항목의 크기를 읽지 않습니다. 이는 별도의 트랜잭션을
      수행하기 때문입니다. 이는 동기화에 필요하지 않습니다. 그러나 rclone mount를 사용할 때
      파일의 크기를 미리 읽어야 하므로 이 플래그를 설정하는 것이 권장됩니다.

   --start-year
      년도를 지정하여 해당 년도 이후에 업로드된 사진만 다운로드합니다.

   --include-archived
      아카이브된 미디어도 표시하고 다운로드합니다.
      
      기본적으로 rclone은 아카이브된 미디어를 요청하지 않습니다. 따라서 동기화할 때,
      아카이브된 미디어는 디렉터리 목록이나 전송에서 표시되지 않습니다.
      
      앨범의 미디어는 아카이브에 상관없이 항상 표시되고 동기화됩니다.
      
      이 플래그가 있는 경우, 아카이브된 미디어는 항상 디렉터리 목록에 표시되고 전송됩니다.
      
      이 플래그가 없는 경우, 아카이브된 미디어는 디렉터리 목록에 표시되지 않고 전송되지 않습니다.

   --encoding
      백엔드의 인코딩 방식입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --client-id 값      OAuth 클라이언트 ID입니다. [$CLIENT_ID]
   --client-secret 값  OAuth 클라이언트 비밀입니다. [$CLIENT_SECRET]
   --help, -h             도움말 표시
   --read-only            Google 사진 백엔드를 읽기 전용으로 설정합니다. (기본값: false) [$READ_ONLY]

   고급

   --auth-url 값    인증 서버 URL입니다. [$AUTH_URL]
   --encoding 값    백엔드의 인코딩 방식입니다. (기본값: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --include-archived  아카이브된 미디어도 표시하고 다운로드합니다. (기본값: false) [$INCLUDE_ARCHIVED]
   --read-size         미디어 항목의 크기를 읽도록 설정합니다. (기본값: false) [$READ_SIZE]
   --start-year 값  년도를 지정하여 해당 년도 이후에 업로드된 사진만 다운로드합니다. (기본값: 2000) [$START_YEAR]
   --token 값       JSON blob 형태의 OAuth 액세스 토큰입니다. [$TOKEN]
   --token-url 값   토큰 서버 URL입니다. [$TOKEN_URL]

```
{% endcode %}