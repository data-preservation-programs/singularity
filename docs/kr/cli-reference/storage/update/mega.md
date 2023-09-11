# Mega

{% code fullWidth="true" %}
```
이름:
   singularity storage update mega - Mega

사용법:
   singularity storage update mega [옵션] <이름|아이디>

설명:
   --user
      사용자 이름.

   --pass
      비밀번호.

   --debug
      Mega에서 더 많은 디버그 출력.
      
      이 플래그가 설정되면(-vv와 함께 설정되면) 더 많은 디버깅 정보가 mega 백엔드에서 출력됩니다.

   --hard-delete
      파일을 영구적으로 삭제하고 휴지통에 넣지 않습니다.
      
      일반적으로 mega 백엔드는 모든 삭제를 휴지통에 넣는 대신 영구적으로 삭제합니다. 이 옵션을 지정하면 rclone은 대신 객체를 영구적으로 삭제합니다.

   --use-https
      전송에 HTTPS 사용.
      
      MEGA는 기본적으로 일반 텍스트 HTTP 연결을 사용합니다.
      일부 ISP는 HTTP 연결을 제한하여 전송 속도가 매우 느려질 수 있습니다.
      이 옵션을 활성화하면 MEGA가 모든 전송에 HTTPS를 사용하도록 강제합니다.
      HTTPS는 이미 모든 데이터가 암호화되어 있으므로 일반적으로 필요하지 않습니다.
      이를 활성화하면 CPU 사용량이 증가하고 네트워크 오버헤드가 추가됩니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --help, -h    도움말 표시
   --pass value  비밀번호. [$PASS]
   --user value  사용자 이름. [$USER]

   Advanced

   --debug           Mega에서 더 많은 디버그 출력. (기본값: false) [$DEBUG]
   --encoding value  백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete     파일을 영구적으로 삭제하고 휴지통에 넣지 않습니다. (기본값: false) [$HARD_DELETE]
   --use-https       전송에 HTTPS 사용. (기본값: false) [$USE_HTTPS]

```
{% endcode %}