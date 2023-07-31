# SMB / CIFS

{% code fullWidth="true" %}
```
이름:
   singularity datasource add smb - SMB / CIFS

사용법:
   singularity datasource add smb [command options] <dataset_name> <source_path>

설명:
   --smb-case-insensitive
      서버가 대소문자를 구분하지 않도록 구성되었는지 여부입니다.
      
      Windows 공유에서는 항상 true입니다.

   --smb-domain
      NTLM 인증을 위한 도메인 이름입니다.

   --smb-encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요 섹션의 인코딩](/overview/#encoding)을 참조하세요.

   --smb-hide-special-share
      사용자가 액세스할 수 없는 특수 공유 (예: print$)을 숨깁니다.

   --smb-host
      연결할 SMB 서버 호스트 이름입니다.
      
      예: "example.com".

   --smb-idle-timeout
      대기 중인 연결을 닫기 전의 최대 시간입니다.
      
      주어진 시간 동안 연결 풀로 반환된 연결이 없으면 rclone은 연결 풀을 비웁니다.
      
      연결을 계속 유지하려면 0으로 설정합니다.
      

   --smb-pass
      SMB 암호입니다.

   --smb-port
      SMB 포트 번호입니다.

   --smb-spn
      서비스 프린시팔 이름입니다.
      
      Rclone은 이 이름을 서버에 제시합니다. 일부 서버는 이를 추가 인증으로 사용하며 클러스터에 설정해야 하는 경우가 많습니다. 예를 들어:
      
          cifs/remotehost:1020
      
      확실하지 않은 경우 공백으로 남겨둡니다.
      

   --smb-user
      SMB 사용자 이름입니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보내기 후 데이터 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공한 스캔으로부터 해당 간격이 경과하면 소스 디렉토리를 자동으로 다시 스캔합니다. (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비)

   smb 옵션

   --smb-case-insensitive value    서버가 대소문자를 구분하지 않도록 구성되었는지 여부입니다. (기본값: "true") [$SMB_CASE_INSENSITIVE]
   --smb-domain value              NTLM 인증을 위한 도메인 이름입니다. (기본값: "WORKGROUP") [$SMB_DOMAIN]
   --smb-encoding value            백엔드에 대한 인코딩입니다. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot") [$SMB_ENCODING]
   --smb-hide-special-share value  사용자가 액세스할 수 없는 특수 공유 (예: print$)을 숨깁니다. (기본값: "true") [$SMB_HIDE_SPECIAL_SHARE]
   --smb-host value                연결할 SMB 서버 호스트 이름입니다. [$SMB_HOST]
   --smb-idle-timeout value        대기 중인 연결을 닫기 전의 최대 시간입니다. (기본값: "1m0s") [$SMB_IDLE_TIMEOUT]
   --smb-pass value                SMB 암호입니다. [$SMB_PASS]
   --smb-port value                SMB 포트 번호입니다. (기본값: "445") [$SMB_PORT]
   --smb-spn value                 서비스 프린시팔 이름입니다. [$SMB_SPN]
   --smb-user value                SMB 사용자 이름입니다. (기본값: "$USER") [$SMB_USER]

```
{% endcode %}