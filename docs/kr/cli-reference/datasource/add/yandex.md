# Yandex Disk

{% code fullWidth="true" %}
```
이름:
   singularity datasource add yandex - Yandex Disk

사용법:
   singularity datasource add yandex [command options] <dataset_name> <source_path>

설명:
   --yandex-auth-url
      인증 서버 URL입니다.
      
      제공자 기본값을 사용하려면 비워 둡니다.

   --yandex-client-id
      OAuth 클라이언트 ID입니다.
      
      보통 비워 둡니다.

   --yandex-client-secret
      OAuth 클라이언트 비밀입니다.
      
      보통 비워 둡니다.

   --yandex-encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --yandex-hard-delete
      파일을 휴지통이 아니라 영구적으로 삭제합니다.

   --yandex-token
      OAuth 액세스 토큰입니다(JSON 형식).

   --yandex-token-url
      토큰 서버 URL입니다.
      
      제공자 기본값을 사용하려면 비워 둡니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] 데이터셋을 CAR 파일로 내보낸 후 데이터 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터 이 간격 만큼 경과한 경우 소스 디렉터리를 자동으로 다시 스캔합니다 (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다 (기본값: 준비 완료)

   Yandex을 위한 옵션

   --yandex-auth-url value       인증 서버 URL. [$YANDEX_AUTH_URL]
   --yandex-client-id value      OAuth 클라이언트 ID. [$YANDEX_CLIENT_ID]
   --yandex-client-secret value  OAuth 클라이언트 비밀. [$YANDEX_CLIENT_SECRET]
   --yandex-encoding value       백엔드의 인코딩. (기본값: "Slash,Del,Ctl,InvalidUtf8,Dot") [$YANDEX_ENCODING]
   --yandex-hard-delete value    파일을 휴지통이 아니라 영구적으로 삭제합니다. (기본값: "false") [$YANDEX_HARD_DELETE]
   --yandex-token value          OAuth 액세스 토큰입니다(JSON 형식). [$YANDEX_TOKEN]
   --yandex-token-url value      토큰 서버 URL입니다. [$YANDEX_TOKEN_URL]

```
{% endcode %}