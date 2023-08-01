# Sia 탈중앙 클라우드

{% code fullWidth="true" %}
```
이름:
   singularity datasource add sia - Sia 탈중앙 클라우드

사용법:
   singularity datasource add sia [command options] <dataset_name> <source_path>

설명:
   --sia-api-password
      Sia 데몬 API 비밀번호.
      
      홈 디렉토리의 `.sia/` 또는 데몬 디렉토리에 있는 `apipassword` 파일에서 찾을 수 있습니다.

   --sia-api-url
      Sia 데몬 API URL, 예시: http://sia.daemon.host:9980.

      다른 호스트를 위해 API 포트를 열기 위해 siad가 `--disable-api-security` 옵션으로 실행된 것을 유의하세요 (비추천).
      Sia 데몬이 localhost에서 실행 중인 경우에는 기본값을 유지하세요.

   --sia-encoding
      백엔드의 인코딩 방식입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --sia-user-agent
      Siad 사용자 에이전트
      
      Sia 데몬은 보안을 위해 기본적으로 'Sia-Agent' 사용자 에이전트를 요구합니다.


옵션:
   --help, -h  도움말 보기

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보낸 후 데이터 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터 경과한 시간이 이 간격을 넘을 경우, 소스 디렉토리를 자동으로 다시 스캔합니다 (기본값: 사용 안 함)
   --scanning-state value   초기 스캔 상태를 설정합니다 (기본값: 준비됨)

   Sia 옵션

   --sia-api-password value  Sia 데몬 API 비밀번호. [$SIA_API_PASSWORD]
   --sia-api-url value       Sia 데몬 API URL, 예시: http://sia.daemon.host:9980. (기본값: "http://127.0.0.1:9980") [$SIA_API_URL]
   --sia-encoding value      백엔드의 인코딩 방식 (기본값: "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot") [$SIA_ENCODING]
   --sia-user-agent value    Siad 사용자 에이전트 (기본값: "Sia-Agent") [$SIA_USER_AGENT]

```
{% endcode %}