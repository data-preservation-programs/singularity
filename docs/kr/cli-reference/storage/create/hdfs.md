# 하둡 분산 파일 시스템

{% code fullWidth="true" %}
```
이름:
   singularity storage create hdfs - 하둡 분산 파일 시스템

사용법:
   singularity storage create hdfs [command options] [arguments...]

설명:
   --namenode
      하둡 네임노드와 포트입니다.

      예시: "namenode:8020"은 포트 8020에서 호스트 namenode에 연결합니다.

   --username
      하둡 사용자 이름입니다.

      예시:
         | root | root로 hdfs에 연결합니다.

   --service-principal-name
      네임노드를 위한 Kerberos 서비스 주체 이름입니다.

      KERBEROS 인증을 가능하게 합니다. 네임노드를 위한 서비스 주체 이름(SERVICE/FQDN)을 지정합니다. 예시로 'hdfs/namenode.hadoop.docker'은 서비스 'hdfs'를 FQDN 'namenode.hadoop.docker'로 실행하는 네임노드를 지정합니다.

   --data-transfer-protection
      Kerberos 데이터 전송 보호: 인증|무결성|개인정보보호입니다.

      데이터 노드와 통신 시 인증, 데이터 서명 무결성 검사, 와이어 암호화가 필요한지 여부를 지정합니다. '인증', '무결성', '개인정보보호' 등의 값이 가능합니다. KERBEROS가 활성화된 경우에만 사용됩니다.

      예시:
         | 개인정보보호 | 인증, 무결성 및 암호화 사용

   --encoding
      백엔드의 인코딩 방식입니다.

      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --help, -h     도움말 표시
   --namenode 값  하둡 네임노드와 포트입니다. [$NAMENODE]
   --username 값  하둡 사용자 이름입니다. [$USERNAME]

   고급

   --data-transfer-protection 값  Kerberos 데이터 전송 보호: 인증|무결성|개인정보보호입니다. [$DATA_TRANSFER_PROTECTION]
   --encoding 값                  백엔드의 인코딩 방식입니다. (기본값: "Slash,Colon,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --service-principal-name 값    네임노드를 위한 Kerberos 서비스 주체 이름입니다. [$SERVICE_PRINCIPAL_NAME]

   일반

   --name 값  저장소의 이름 (기본값: 자동 생성)
   --path 값  저장소의 경로

```
{% endcode %}