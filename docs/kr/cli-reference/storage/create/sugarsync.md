# Sugarsync

{% code fullWidth="true" %}
```
이름:
   singularity storage create sugarsync - Sugarsync

사용법:
   singularity storage create sugarsync [command options] [arguments...]

설명:
   --app-id
      Sugarsync 앱 ID입니다.
      
      rclone의 것을 사용하려면 비워두세요.

   --access-key-id
      Sugarsync 액세스 키 ID입니다.
      
      rclone의 것을 사용하려면 비워두세요.

   --private-access-key
      Sugarsync 개인 액세스 키입니다.
      
      rclone의 것을 사용하려면 비워두세요.

   --hard-delete
      true로 설정하면 파일을 영구적으로 삭제합니다.
      그렇지 않으면 삭제된 파일을 삭제된 파일 폴더에 저장합니다.

   --refresh-token
      Sugarsync 갱신 토큰입니다.
      
      보통 비워두며, rclone이 자동으로 구성합니다.

   --authorization
      Sugarsync 인증입니다.
      
      보통 비워두며, rclone이 자동으로 구성합니다.

   --authorization-expiry
      Sugarsync 인증 만료 기간입니다.
      
      보통 비워두며, rclone이 자동으로 구성합니다.

   --user
      Sugarsync 사용자입니다.
      
      보통 비워두며, rclone이 자동으로 구성합니다.

   --root-id
      Sugarsync 루트 ID입니다.
      
      보통 비워두며, rclone이 자동으로 구성합니다.

   --deleted-id
      Sugarsync 삭제된 폴더 ID입니다.
      
      보통 비워두며, rclone이 자동으로 구성합니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --access-key-id value       Sugarsync 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --app-id value              Sugarsync 앱 ID입니다. [$APP_ID]
   --hard-delete               파일을 영구적으로 삭제하려면 true (기본값: false) [$HARD_DELETE]
   --help, -h                  도움말 표시
   --private-access-key value  Sugarsync 개인 액세스 키 입니다. [$PRIVATE_ACCESS_KEY]

   고급

   --authorization value         Sugarsync 인증입니다. [$AUTHORIZATION]
   --authorization-expiry value  Sugarsync 인증 만료 기간입니다. [$AUTHORIZATION_EXPIRY]
   --deleted-id value            Sugarsync 삭제된 폴더 ID입니다. [$DELETED_ID]
   --encoding value              백엔드의 인코딩입니다. (기본값: "Slash,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --refresh-token value         Sugarsync 갱신 토큰입니다. [$REFRESH_TOKEN]
   --root-id value               Sugarsync 루트 ID입니다. [$ROOT_ID]
   --user value                  Sugarsync 사용자입니다. [$USER]

   일반

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}