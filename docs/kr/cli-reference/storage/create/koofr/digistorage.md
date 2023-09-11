# Digi Storage, https://storage.rcs-rds.ro/

{% code fullWidth="true" %}
```
이름:
   singularity storage create koofr digistorage - Digi Storage, https://storage.rcs-rds.ro/

사용법:
   singularity storage create koofr digistorage [command options] [arguments...]

설명:
   --mountid
      사용할 마운트의 마운트 ID입니다.
      
      지정하지 않으면 기본 마운트가 사용됩니다.

   --setmtime
      백엔드가 수정 시간 설정을 지원하는지 여부입니다.
      
      Dropbox 또는 Amazon Drive 백엔드를 가리키는 마운트 ID를 사용하는 경우 이 값을 false로 설정합니다.

   --user
      사용자 이름입니다.

   --password
      rclone을 위한 비밀번호입니다. (https://storage.rcs-rds.ro/app/admin/preferences/password에서 생성하세요).

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --help, -h        도움말 표시
   --password value  rclone을 위한 비밀번호입니다. (https://storage.rcs-rds.ro/app/admin/preferences/password에서 생성하세요). [$PASSWORD]
   --user value      사용자 이름입니다. [$USER]

   고급

   --encoding value  백엔드에 대한 인코딩입니다. (기본값: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --mountid value   사용할 마운트의 마운트 ID입니다. [$MOUNTID]
   --setmtime        백엔드가 수정 시간 설정을 지원하는지 여부입니다. (기본값: true) [$SETMTIME]

   일반

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}