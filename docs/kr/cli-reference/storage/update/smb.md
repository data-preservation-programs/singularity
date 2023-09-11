# SMB / CIFS

{% code fullWidth="true" %}
```
이름:
   singularity storage update smb - SMB / CIFS

사용법:
   singularity storage update smb [command options] <name|id>

설명:
   --host
      연결할 SMB 서버의 호스트 이름입니다.
      
      예: "example.com".

   --user
      SMB 사용자 이름입니다.

   --port
      SMB 포트 번호입니다.

   --pass
      SMB 비밀번호입니다.

   --domain
      NTLM 인증에 사용할 도메인 이름입니다.

   --spn
      서비스 주체 이름입니다.
      
      Rclone은 이 이름을 서버에 제공합니다. 일부 서버는 이를 추가 인증으로 사용하며 클러스터에 설정해야 할 수도 있습니다. 예를 들면:
      
          cifs/remotehost:1020
      
      확실하지 않은 경우 비워 두십시오.
      

   --idle-timeout
      유휴 연결을 닫기 전까지의 최대 시간입니다.
      
      주어진 시간 동안 연결 풀에 반환된 연결이 없으면 rclone은 연결 풀을 비웁니다.
      
      연결을 계속 유지하려면 0으로 설정하세요.
      

   --hide-special-share
      사용자가 액세스해서는 안 되는 특수 공유 (예: print$)을 숨깁니다.

   --case-insensitive
      서버가 대소문자를 구분하지 않도록 구성되어 있는지 여부입니다.
      
      Windows 공유의 경우 항상 true입니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --domain value  NTLM 인증에 사용할 도메인 이름입니다. (기본값: "WORKGROUP") [$DOMAIN]
   --help, -h      도움말 표시
   --host value    연결할 SMB 서버의 호스트 이름입니다. [$HOST]
   --pass value    SMB 비밀번호입니다. [$PASS]
   --port value    SMB 포트 번호입니다. (기본값: 445) [$PORT]
   --spn value     서비스 주체 이름입니다. [$SPN]
   --user value    SMB 사용자 이름입니다. (기본값: "$USER") [$USER]

   고급 옵션

   --case-insensitive    서버가 대소문자를 구분하지 않도록 구성되어 있는지 여부입니다. (기본값: true) [$CASE_INSENSITIVE]
   --encoding value      백엔드의 인코딩입니다. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --hide-special-share  사용자가 액세스해서는 안 되는 특수 공유 (예: print$)을 숨깁니다. (기본값: true) [$HIDE_SPECIAL_SHARE]
   --idle-timeout value  유휴 연결을 닫기 전까지의 최대 시간입니다. (기본값: "1m0s") [$IDLE_TIMEOUT]

```
{% endcode %}