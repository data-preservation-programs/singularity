# 기타 Koofr API 호환 스토리지 서비스

```shell
이름:
   singularity storage update koofr other - 기타 Koofr API 호환 스토리지 서비스

사용법:
   singularity storage update koofr other [command options] <이름|id>

설명:
   --endpoint
      사용할 Koofr API 엔드포인트입니다.

   --mountid
      사용할 마운트의 마운트 ID입니다.
      
      생략하면, 기본 마운트가 사용됩니다.

   --setmtime
      백엔드가 수정 시간을 설정하는 것을 지원하는지 여부입니다.
      
      Dropbox 또는 Amazon Drive 백엔드를 가리키는 마운트 ID를 사용하는 경우, 이 값을 false로 설정하십시오.

   --user
      사용자 이름입니다.

   --password
      rclone을 위한 비밀번호입니다 (서비스의 설정 페이지에서 생성).

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --endpoint value  사용할 Koofr API 엔드포인트입니다. [$ENDPOINT]
   --help, -h        도움말을 표시합니다
   --password value  rclone을 위한 비밀번호입니다 (서비스의 설정 페이지에서 생성). [$PASSWORD]
   --user value      사용자 이름입니다. [$USER]

   기능 확장

   --encoding value  백엔드의 인코딩입니다. (기본값: "슬래시,역 슬래시,문자 삭제,제어 문자,잘못된 UTF-8,마침표") [$ENCODING]
   --mountid value   사용할 마운트의 마운트 ID입니다. [$MOUNTID]
   --setmtime        백엔드가 수정 시간을 설정하는 것을 지원하는지 여부입니다. (기본값: true) [$SETMTIME]
```
