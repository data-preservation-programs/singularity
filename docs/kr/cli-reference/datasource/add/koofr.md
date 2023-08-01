# Koofr, Digi Storage 및 다른 호환되는 저장 공급 업체

{% code fullWidth="true" %}
```
이름:
   singularity datasource add koofr - Koofr, Digi Storage 및 다른 Koofr 호환 저장 공급 업체

사용법:
   singularity datasource add koofr [command options] <dataset_name> <source_path>

설명:
   --koofr-encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --koofr-endpoint
      [공급 업체] - 기타
         사용할 Koofr API 엔드포인트입니다.

   --koofr-mountid
      사용할 마운트의 마운트 ID입니다.
      
      지정하지 않으면 기본 마운트를 사용합니다.

   --koofr-password
      [공급 업체] - koofr
         rclone을 위한 비밀번호입니다 (https://app.koofr.net/app/admin/preferences/password에서 생성하십시오).

      [공급 업체] - digistorage
         rclone을 위한 비밀번호입니다 (https://storage.rcs-rds.ro/app/admin/preferences/password에서 생성하십시오).

      [공급 업체] - 기타
         rclone을 위한 비밀번호입니다 (서비스의 설정 페이지에서 생성하십시오).

   --koofr-provider
      사용할 저장 공급 업체를 선택합니다.

      예시:
         | koofr       | Koofr, https://app.koofr.net/
         | digistorage | Digi Storage, https://storage.rcs-rds.ro/
         | other       | 기타 Koofr API 호환 저장 서비스

   --koofr-setmtime
      백엔드에서 수정 시간 설정을 지원하는지 여부입니다.
      
      Dropbox 또는 Amazon Drive 백엔드를 가리키는 마운트 ID를 사용하는 경우 false로 설정하세요.

   --koofr-user
      사용자 이름입니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보낸 후 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터 일정 시간이 경과하면 소스 디렉토리를 자동으로 다시 스캔합니다. (기본값: 사용 안 함)
   --scanning-state value   초기 스캐닝 상태를 설정합니다. (기본값: 준비 완료)

   koofr 옵션

   --koofr-encoding value  백엔드의 인코딩입니다. (기본값: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$KOOFR_ENCODING]
   --koofr-endpoint value  사용할 Koofr API 엔드포인트입니다. [$KOOFR_ENDPOINT]
   --koofr-mountid value   사용할 마운트의 마운트 ID입니다. [$KOOFR_MOUNTID]
   --koofr-password value  rclone을 위한 비밀번호입니다 (https://app.koofr.net/app/admin/preferences/password에서 생성하십시오). [$KOOFR_PASSWORD]
   --koofr-provider value  사용할 저장 공급 업체를 선택합니다. [$KOOFR_PROVIDER]
   --koofr-setmtime value  백엔드에서 수정 시간 설정을 지원하는지 여부입니다. (기본값: "true") [$KOOFR_SETMTIME]
   --koofr-user value      사용자 이름입니다. [$KOOFR_USER]

```
{% endcode %}