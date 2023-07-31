# Mega

{% code fullWidth="true" %}
```
이름:
   singularity datasource add mega - Mega

사용법:
   singularity datasource add mega [옵션] <데이터셋_이름> <소스_경로>

설명:
   --mega-debug
      Mega로부터 추적 정보를 더 많이 출력합니다.
      
      이 플래그가 설정되면 (-vv와 함께 사용되어야 함) mega 백엔드에서 더 많은 디버그 정보를 인쇄합니다.

   --mega-encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --mega-hard-delete
      파일을 영구적으로 삭제합니다. 휴지통에 넣지 않고.
      
      일반적으로 mega 백엔드는 모든 삭제된 항목을 휴지통에 넣습니다.
      이 플래그를 지정하면 rclone은 대신 객체를 영구적으로 삭제합니다.

   --mega-pass
      비밀번호입니다.

   --mega-use-https
      전송에 HTTPS를 사용합니다.
      
      MEGA는 기본적으로 일반 텍스트 HTTP 연결을 사용합니다.
      일부 ISP는 HTTP 연결을 제한하기 때문에 전송이 매우 느려집니다.
      이 플래그를 활성화하면 MEGA가 모든 전송에 HTTPS를 사용하도록 강제합니다.
      HTTPS는 이미 모든 데이터가 암호화되어 있으므로 일반적으로 필요하지 않습니다.
      그러나 이를 활성화하면 CPU 사용량이 증가하고 네트워크 부하가 추가됩니다.

   --mega-user
      사용자 이름입니다.


옵션:
   --help, -h  도움말 보기

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보낸 후에 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터 이 시간 간격이 경과하면 소스 디렉토리를 자동으로 다시 스캔합니다. (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비됨)

   Mega 옵션

   --mega-debug value        Mega로부터 추적 정보를 더 많이 출력합니다. (기본값: "false") [$MEGA_DEBUG]
   --mega-encoding value     백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$MEGA_ENCODING]
   --mega-hard-delete value  파일을 영구적으로 삭제합니다. 휴지통에 넣지 않고. (기본값: "false") [$MEGA_HARD_DELETE]
   --mega-pass value         비밀번호입니다. [$MEGA_PASS]
   --mega-use-https value    전송에 HTTPS를 사용합니다. (기본값: "false") [$MEGA_USE_HTTPS]
   --mega-user value         사용자 이름입니다. [$MEGA_USER]

```
{% endcode %}