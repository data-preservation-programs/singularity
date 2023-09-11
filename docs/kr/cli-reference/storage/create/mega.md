# Mega

{% code fullWidth="true" %}
```
이름:
   singularity storage create mega - Mega

사용법:
   singularity storage create mega [command options] [arguments...]

설명:
   --user
      사용자 이름입니다.

   --pass
      비밀번호입니다.

   --debug
      Mega로부터 더 많은 디버그 출력합니다.
      
      이 플래그가 설정되면 (-vv와 함께 사용되는 경우) mega 백엔드에서 추가로
      디버깅 정보를 출력합니다.

   --hard-delete
      파일을 휴지통이 아닌 영구적으로 삭제합니다.
      
      일반적으로 mega 백엔드는 모든 삭제를 휴지통에 넣고 영구적인 삭제 대신에 저장합니다.
      이 플래그를 지정하면 rclone은 객체를 영구적으로 삭제할 것입니다.

   --use-https
      전송에 HTTPS를 사용합니다.
      
      MEGA는 기본적으로 일반 텍스트 HTTP 연결을 사용합니다.
      일부 ISP는 HTTP 연결을 조절하고 이로 인해 전송 속도가 극히 느려집니다.
      이 옵션을 사용하면 MEGA가 모든 전송에 HTTPS를 사용하도록 강제합니다.
      HTTPS는 데이터가 이미 암호화되어 있기 때문에 일반적으로 필요하지 않습니다.
      사용하면 CPU 사용량이 증가하고 네트워크 오버 헤드가 추가됩니다.

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --help, -h    도움말 표시
   --pass value  비밀번호입니다. [$PASS]
   --user value  사용자 이름입니다. [$USER]

   고급

   --debug           Mega로부터 더 많은 디버그 출력합니다. (기본값: false) [$DEBUG]
   --encoding value  백엔드에 대한 인코딩입니다. (기본값: "슬래시, 부적합한 UTF-8, 점") [$ENCODING]
   --hard-delete     파일을 휴지통이 아닌 영구적으로 삭제합니다. (기본값: false) [$HARD_DELETE]
   --use-https       전송에 HTTPS를 사용합니다. (기본값: false) [$USE_HTTPS]

   일반

   --name value  스토리지의 이름입니다 (기본값: 자동 생성)
   --path value  스토리지의 경로입니다

```
{% endcode %}