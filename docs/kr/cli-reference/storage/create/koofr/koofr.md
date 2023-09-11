# Koofr, https://app.koofr.net/

{% code fullWidth="true" %}
```
이름:
   singularity storage create koofr koofr - Koofr, https://app.koofr.net/

사용법:
   singularity storage create koofr koofr [command options] [arguments...]

설명:
   --mountid
      사용할 마운트의 마운트 ID입니다.
      
      지정하지 않으면 기본 마운트를 사용합니다.

   --setmtime
      백엔드에서 수정 시간을 설정할 수 있는지 여부입니다.
      
      Dropbox 또는 Amazon Drive 백엔드를 가리키는 마운트 ID를 사용하는 경우 false로 설정하세요.

   --user
      사용자 이름입니다.

   --password
      rclone에 대한 비밀번호입니다 (https://app.koofr.net/app/admin/preferences/password에서 생성).

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --help, -h        도움말 표시
   --password value  rclone에 대한 비밀번호입니다 (https://app.koofr.net/app/admin/preferences/password에서 생성). [$PASSWORD]
   --user value      사용자 이름입니다. [$USER]

   Advanced

   --encoding value  백엔드에 대한 인코딩입니다. (기본값: "슬래시,백슬래시,삭제,제어,무효한Utf8,점") [$ENCODING]
   --mountid value   사용할 마운트의 마운트 ID입니다. [$MOUNTID]
   --setmtime        백엔드에서 수정 시간을 설정할 수 있는지 여부입니다. (기본값: true) [$SETMTIME]

   General

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}