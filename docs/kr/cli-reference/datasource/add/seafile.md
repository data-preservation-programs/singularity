# seafile

{% code fullWidth="true" %}
```
이름:
   singularity 데이터 소스 추가 seafile - seafile

사용법:
   singularity 데이터 소스 추가 seafile [command options] <데이터셋_이름> <소스_경로>

설명:
   --seafile-2fa
      2단계 인증 활성화 여부 ('true'인 경우 2단계 인증 활성화).

   --seafile-auth-token
      인증 토큰.

   --seafile-create-library
      라이브러리가 존재하지 않을 경우 rclone이 라이브러리를 생성해야 하는지 여부.

   --seafile-encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --seafile-library
      라이브러리 이름.
      
      암호화되지 않은 모든 라이브러리에 액세스하려면 비워 두십시오.

   --seafile-library-key
      라이브러리 비밀번호 (암호화된 라이브러리 전용).
      
      명령행을 통해 전달하는 경우 비워 두십시오.

   --seafile-pass
      비밀번호.

   --seafile-url
      연결할 seafile 호스트의 URL입니다.

      예시:
         | https://cloud.seafile.com/ | cloud.seafile.com에 연결합니다.

   --seafile-user
      사용자 이름 (일반적으로 이메일 주소).


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] 데이터셋을 CAR 파일로 내보낸 후 해당 파일 삭제하기 (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔부터 이 간격이 경과하면 소스 디렉토리를 자동으로 다시 스캔합니다 (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태 설정 (기본값: 준비 완료)

   seafile 옵션

   --seafile-2fa value             2단계 인증 활성화 여부 ('true'인 경우 2단계 인증 활성화). (기본값: "false") [$SEAFILE_2FA]
   --seafile-create-library value  라이브러리가 존재하지 않을 경우 rclone이 라이브러리를 생성해야 하는지 여부. (기본값: "false") [$SEAFILE_CREATE_LIBRARY]
   --seafile-encoding value        백엔드의 인코딩. (기본값: "슬래시,따옴표,역슬래시,Ctl,유효하지 않은 UTF-8") [$SEAFILE_ENCODING]
   --seafile-library value         라이브러리 이름. [$SEAFILE_LIBRARY]
   --seafile-library-key value     라이브러리 비밀번호 (암호화된 라이브러리 전용). [$SEAFILE_LIBRARY_KEY]
   --seafile-pass value            비밀번호. [$SEAFILE_PASS]
   --seafile-url value             seafile 호스트의 URL. [$SEAFILE_URL]
   --seafile-user value            사용자 이름 (일반적으로 이메일 주소). [$SEAFILE_USER]

```
{% endcode %}