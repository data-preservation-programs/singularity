# 기타 Koofr API 호환 스토리지 서비스

{% code fullWidth="true" %}
```
명령어:
   singularity storage create koofr other - 기타 Koofr API 호환 스토리지 서비스 생성

사용법:
   singularity storage create koofr other [command 옵션] [인자...]

설명:
   --endpoint
      사용할 Koofr API 엔드포인트입니다.

   --mountid
      사용할 마운트의 마운트 ID입니다.
      
      생략할 경우, 기본 마운트를 사용합니다.

   --setmtime
      백엔드에서 수정 시간을 설정할 수 있는지 여부입니다.
      
      Dropbox나 Amazon Drive 백엔드로 가리키는 마운트 ID를 사용하는 경우 이 값을 false로 설정합니다.

   --user
      사용자 이름입니다.

   --password
      rclone에서 사용할 비밀번호입니다 (서비스의 설정 페이지에서 생성).

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --endpoint value  사용할 Koofr API 엔드포인트입니다. [$ENDPOINT]
   --help, -h        도움말 표시
   --password value  rclone에서 사용할 비밀번호입니다 (서비스의 설정 페이지에서 생성). [$PASSWORD]
   --user value      사용자 이름입니다. [$USER]

   고급 옵션

   --encoding value  백엔드의 인코딩입니다. (기본값: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --mountid value   사용할 마운트의 마운트 ID입니다. [$MOUNTID]
   --setmtime        백엔드에서 수정 시간을 설정할 수 있는지 여부입니다. (기본값: true) [$SETMTIME]

   일반 옵션

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}