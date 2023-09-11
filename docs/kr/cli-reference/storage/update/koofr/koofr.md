# Koofr, https://app.koofr.net/

{% code fullWidth="true" %}
```
이름:
   singularity storage update koofr koofr - Koofr, https://app.koofr.net/

사용법:
   singularity storage update koofr koofr [옵션] <이름|아이디>

설명:
   --mountid
      사용할 마운트 ID입니다.
      
      지정하지 않으면 기본 마운트를 사용합니다.

   --setmtime
      백엔드가 수정 날짜를 설정하는 기능을 지원하는지 여부입니다.
      
      Dropbox 또는 Amazon Drive 백엔드를 가리키는 마운트 ID를 사용하는 경우 false로 설정합니다.

   --user
      사용자 이름입니다.

   --password
      rclone을 위한 비밀번호입니다. (https://app.koofr.net/app/admin/preferences/password에서 생성할 수 있습니다.)

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --help, -h        도움말 표시
   --password value  rclone을 위한 비밀번호입니다. (https://app.koofr.net/app/admin/preferences/password에서 생성할 수 있습니다.) [$PASSWORD]
   --user value      사용자 이름입니다. [$USER]

   고급

   --encoding value  백엔드에 대한 인코딩입니다. (기본값: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --mountid value   사용할 마운트 ID입니다. [$MOUNTID]
   --setmtime        백엔드가 수정 날짜를 설정하는 기능을 지원하는지 여부입니다. (기본값: true) [$SETMTIME]
```
{% endcode %}