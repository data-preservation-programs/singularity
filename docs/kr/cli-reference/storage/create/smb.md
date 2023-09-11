# SMB / CIFS

{% code fullWidth="true" %}
```
NAME:
   singularity storage create smb - SMB / CIFS

USAGE:
   singularity storage create smb [command options] [arguments...]

DESCRIPTION:
   --host
      연결할 SMB 서버 호스트 이름입니다.
      
      예: "example.com".

   --user
      SMB 사용자명입니다.

   --port
      SMB 포트 번호입니다.

   --pass
      SMB 암호입니다.

   --domain
      NTLM 인증용 도메인 이름입니다.

   --spn
      서비스 주체 이름입니다.
      
      Rclone은 이 이름을 서버에 제시합니다. 일부 서버는 이를 추가 인증용으로 사용하며 클러스터에 설정해야 할 수도 있습니다.
      예를 들면:
      
          cifs/remotehost:1020
      
      잘 모르면 비워 두세요.
      

   --idle-timeout
      사용되지 않는 연결을 닫는 최대 시간입니다.
      
      주어진 시간 내에 연결이 연결 풀에 반환되지 않으면 rclone은 연결 풀을 비웁니다.
      
      연결을 계속 유지하려면 0으로 설정하세요.
      

   --hide-special-share
      사용자가 액세스할 수 없는 특수 공유 (예: print$)를 숨깁니다.

   --case-insensitive
      서버가 대/소문자를 구별하지 않도록 설정되었는지 여부입니다.
      
      Windows 공유의 경우 항상 true입니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


OPTIONS:
   --domain value  NTLM 인증용 도메인 이름입니다. (기본값: "WORKGROUP") [$DOMAIN]
   --help, -h      도움말 표시
   --host value    연결할 SMB 서버 호스트 이름입니다. [$HOST]
   --pass value    SMB 암호입니다. [$PASS]
   --port value    SMB 포트 번호입니다. (기본값: 445) [$PORT]
   --spn value     서비스 주체 이름입니다. [$SPN]
   --user value    SMB 사용자명입니다. (기본값: "$USER") [$USER]

   Advanced

   --case-insensitive    서버가 대/소문자를 구별하지 않도록 설정되었는지 여부입니다. (기본값: true) [$CASE_INSENSITIVE]
   --encoding value      백엔드의 인코딩입니다. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --hide-special-share  사용자가 액세스할 수 없는 특수 공유 (예: print$)를 숨깁니다. (기본값: true) [$HIDE_SPECIAL_SHARE]
   --idle-timeout value  사용되지 않는 연결을 닫는 최대 시간입니다. (기본값: "1m0s") [$IDLE_TIMEOUT]

   General

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}