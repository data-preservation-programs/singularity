# Sugarsync

{% code fullWidth="true" %}
```
이름:
   singularity datasource add sugarsync - Sugarsync

사용법:
   singularity datasource add sugarsync [command options] <데이터셋_이름> <소스_경로>

설명:
   --sugarsync-access-key-id
      Sugarsync 액세스 키 ID.
      
      비워 두면 rclone의 것을 사용합니다.

   --sugarsync-app-id
      Sugarsync 앱 ID.
      
      비워 두면 rclone의 것을 사용합니다.

   --sugarsync-authorization
      Sugarsync 인증.
      
      일반적으로 비워둡니다. rclone에 의해 자동으로 구성됩니다.

   --sugarsync-authorization-expiry
      Sugarsync 인증 만료.
      
      일반적으로 비워둡니다. rclone에 의해 자동으로 구성됩니다.

   --sugarsync-deleted-id
      Sugarsync 삭제된 폴더 ID.
      
      일반적으로 비워둡니다. rclone에 의해 자동으로 구성됩니다.

   --sugarsync-encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --sugarsync-hard-delete
      true이면 파일을 영구적으로 삭제합니다.
      그렇지 않으면 삭제된 파일이 있는 폴더에 저장합니다.

   --sugarsync-private-access-key
      Sugarsync 개인 액세스 키.
      
      비워 두면 rclone의 것을 사용합니다.

   --sugarsync-refresh-token
      Sugarsync 리프레시 토큰.
      
      일반적으로 비워둡니다. rclone에 의해 자동으로 구성됩니다.

   --sugarsync-root-id
      Sugarsync 루트 ID.
      
      일반적으로 비워둡니다. rclone에 의해 자동으로 구성됩니다.

   --sugarsync-user
      Sugarsync 사용자.
      
      일반적으로 비워둡니다. rclone에 의해 자동으로 구성됩니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] 데이터셋을 CAR 파일로 내보낸 후 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공한 스캔으로부터 지정된 간격만큼 지난 후 소스 디렉토리를 자동으로 다시 스캔합니다. (기본값: disabled)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: ready)

   SugarSync 옵션

   --sugarsync-access-key-id value         Sugarsync 액세스 키 ID. [$SUGARSYNC_ACCESS_KEY_ID]
   --sugarsync-app-id value                Sugarsync 앱 ID. [$SUGARSYNC_APP_ID]
   --sugarsync-authorization value         Sugarsync 인증. [$SUGARSYNC_AUTHORIZATION]
   --sugarsync-authorization-expiry value  Sugarsync 인증 만료. [$SUGARSYNC_AUTHORIZATION_EXPIRY]
   --sugarsync-deleted-id value            Sugarsync 삭제된 폴더 ID. [$SUGARSYNC_DELETED_ID]
   --sugarsync-encoding value              백엔드의 인코딩. (기본값: "Slash,Ctl,InvalidUtf8,Dot") [$SUGARSYNC_ENCODING]
   --sugarsync-hard-delete value           true이면 파일을 영구적으로 삭제합니다. (기본값: "false") [$SUGARSYNC_HARD_DELETE]
   --sugarsync-private-access-key value    Sugarsync 개인 액세스 키. [$SUGARSYNC_PRIVATE_ACCESS_KEY]
   --sugarsync-refresh-token value         Sugarsync 리프레시 토큰. [$SUGARSYNC_REFRESH_TOKEN]
   --sugarsync-root-id value               Sugarsync 루트 ID. [$SUGARSYNC_ROOT_ID]
   --sugarsync-user value                  Sugarsync 사용자. [$SUGARSYNC_USER]
```
{% endcode %}