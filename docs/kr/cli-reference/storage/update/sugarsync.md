# Sugarsync

{% code fullWidth="true" %}
```
이름:
   singularity storage update sugarsync - Sugarsync

사용법:
   singularity storage update sugarsync [command options] <이름|아이디>

설명:
   --app-id
      Sugarsync 어플리케이션 ID입니다.
      
      비어 두면 rclone의 기본값을 사용합니다.

   --access-key-id
      Sugarsync 접근 키 ID입니다.
      
      비어 두면 rclone의 기본값을 사용합니다.

   --private-access-key
      Sugarsync 개인 접근 키입니다.
      
      비어 두면 rclone의 기본값을 사용합니다.

   --hard-delete
      파일을 영구적으로 삭제하려면 true로 설정하세요.
      그렇지 않으면 삭제된 파일이 삭제 폴더에 저장됩니다.

   --refresh-token
      Sugarsync 갱신 토큰입니다.
      
      일반적으로 비워두고, rclone에 의해 자동으로 구성됩니다.

   --authorization
      Sugarsync 인증입니다.
      
      일반적으로 비워두고, rclone에 의해 자동으로 구성됩니다.

   --authorization-expiry
      Sugarsync 인증 만료일입니다.
      
      일반적으로 비워두고, rclone에 의해 자동으로 구성됩니다.

   --user
      Sugarsync 사용자입니다.
      
      일반적으로 비워두고, rclone에 의해 자동으로 구성됩니다.

   --root-id
      Sugarsync 루트 ID입니다.
      
      일반적으로 비워두고, rclone에 의해 자동으로 구성됩니다.

   --deleted-id
      Sugarsync 삭제된 폴더 ID입니다.
      
      일반적으로 비워두고, rclone에 의해 자동으로 구성됩니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --access-key-id value       Sugarsync 접근 키 ID입니다. [$ACCESS_KEY_ID]
   --app-id value              Sugarsync 어플리케이션 ID입니다. [$APP_ID]
   --hard-delete               파일을 영구적으로 삭제하려면 true (기본값: false) [$HARD_DELETE]
   --help, -h                  도움말 표시
   --private-access-key value  Sugarsync 개인 접근 키입니다. [$PRIVATE_ACCESS_KEY]

   Advanced

   --authorization value         Sugarsync 인증입니다. [$AUTHORIZATION]
   --authorization-expiry value  Sugarsync 인증 만료일입니다. [$AUTHORIZATION_EXPIRY]
   --deleted-id value            Sugarsync 삭제된 폴더 ID입니다. [$DELETED_ID]
   --encoding value              백엔드의 인코딩입니다. (기본값: "Slash,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --refresh-token value         Sugarsync 갱신 토큰입니다. [$REFRESH_TOKEN]
   --root-id value               Sugarsync 루트 ID입니다. [$ROOT_ID]
   --user value                  Sugarsync 사용자입니다. [$USER]

```
{% endcode %}