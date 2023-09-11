# Hadoop 분산 파일 시스템

{% code fullWidth="true" %}
```
이름:
   singularity storage update hdfs - Hadoop 분산 파일 시스템

사용법:
   singularity storage update hdfs [command options] <이름|ID>

설명:
   --namenode
      Hadoop 네임 노드와 포트입니다.
      
      예시: "namenode:8020"은 포트 8020에서 호스트 네임 노드에 연결합니다.

   --username
      Hadoop 사용자 이름입니다.

      예시:
         | root | root로 hdfs에 연결합니다.

   --service-principal-name
      네임 노드의 Kerberos 서비스 주체 이름입니다.
      
      KERBEROS 인증을 활성화합니다. 네임 노드를 위한 서비스 주체 이름(SERVICE/FQDN)을 지정합니다.
      예시: "hdfs/namenode.hadoop.docker"는 서비스 'hdfs'와 FQDN 'namenode.hadoop.docker'로 실행 중인 네임 노드입니다.

   --data-transfer-protection
      Kerberos 데이터 전송 보호: authentication|integrity|privacy입니다.
      
      데이터 노드와 통신 시 인증, 데이터 무결성 검사 및 원격 암호화가 필요한지 여부를 지정합니다.
      가능한 값은 'authentication', 'integrity' 및 'privacy'입니다. KERBEROS가 활성화된 경우에만 사용됩니다.

      예시:
         | privacy | 인증, 무결성 및 암호화가 활성화되도록 보장합니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --help, -h        도움말 표시
   --namenode value  Hadoop 네임 노드와 포트입니다. [$NAMENODE]
   --username value  Hadoop 사용자 이름입니다. [$USERNAME]

   고급

   --data-transfer-protection value  Kerberos 데이터 전송 보호: authentication|integrity|privacy입니다. [$DATA_TRANSFER_PROTECTION]
   --encoding value                  백엔드의 인코딩입니다. (기본값: "슬래시,콜론,삭제,제어문자,잘못된 UTF8,마침표") [$ENCODING]
   --service-principal-name value    네임 노드의 Kerberos 서비스 주체 이름입니다. [$SERVICE_PRINCIPAL_NAME]

```
{% endcode %}