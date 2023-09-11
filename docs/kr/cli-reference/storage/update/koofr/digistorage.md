# Digi Storage, https://storage.rcs-rds.ro/

{% code fullWidth="true" %}
```
이름:
   singularity storage update koofr digistorage - Digi Storage, https://storage.rcs-rds.ro/

사용법:
   singularity storage update koofr digistorage [command options] <이름|아이디>

설명:
   --mountid
      사용할 마운트 ID입니다.
      
      생략하면 기본 마운트가 사용됩니다.

   --setmtime
      백엔드에서 수정 시간을 설정할 수 있는지 여부입니다.
      
      Dropbox나 Amazon Drive 백엔드를 가리키는 마운트 ID를 사용하는 경우 이 값을 false로 설정하세요.

   --user
      사용자 이름입니다.

   --password
      rclone의 비밀번호입니다 (https://storage.rcs-rds.ro/app/admin/preferences/password에서 생성).

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --help, -h        도움말 표시
   --password value  rclone의 비밀번호입니다 (https://storage.rcs-rds.ro/app/admin/preferences/password에서 생성). [$PASSWORD]
   --user value      사용자 이름입니다. [$USER]

   고급 옵션:

   --encoding value  백엔드에 대한 인코딩입니다. (기본값: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --mountid value   사용할 마운트 ID입니다. [$MOUNTID]
   --setmtime        백엔드에서 수정 시간을 설정할 수 있는지 여부입니다. (기본값: true) [$SETMTIME]

```
{% endcode %}