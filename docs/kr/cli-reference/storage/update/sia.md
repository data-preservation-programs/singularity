# Sia 분산 클라우드

{% code fullWidth="true" %}
```
이름:
   singularity storage update sia - Sia 분산 클라우드

사용법:
   singularity storage update sia [command options] <name|id>

설명:
   --api-url
      Sia 데몬 API URL, 예: http://sia.daemon.host:9980.

      참고: 다른 호스트를 위해 API 포트를 열려면 siad가 --disable-api-security 옵션으로 실행되어야 합니다. (권장하지 않음).
      Sia 데몬이 localhost에서 실행 중인 경우 기본값을 유지하십시오.

   --api-password
      Sia 데몬 API 비밀번호.
      
      HOME/.sia/ 또는 데몬 디렉터리에 있는 apipassword 파일에서 찾을 수 있습니다.

   --user-agent
      Siad 사용자 에이전트
      
      Sia 데몬은 기본적으로 보안을 위해 'Sia-Agent' 사용자 에이전트를 요구합니다.

   --encoding
      백엔드에 대한 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --api-password value  Sia 데몬 API 비밀번호. [$API_PASSWORD]
   --api-url value       Sia 데몬 API URL, 예: http://sia.daemon.host:9980. (기본값: "http://127.0.0.1:9980") [$API_URL]
   --help, -h            도움말 표시

   고급

   --encoding value    백엔드에 대한 인코딩. (기본값: "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --user-agent value  Siad 사용자 에이전트 (기본값: "Sia-Agent") [$USER_AGENT]

```
{% endcode %}