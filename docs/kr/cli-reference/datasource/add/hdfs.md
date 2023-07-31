# Hadoop 분산 파일 시스템

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add hdfs - Hadoop 분산 파일 시스템

USAGE:
   singularity datasource add hdfs [command options] <데이터셋_이름> <소스_경로>

DESCRIPTION:
   --hdfs-data-transfer-protection
      Kerberos 데이터 전송 보호: 인증|무결성|암호화.
      
      데이터 노드와 통신할 때 인증, 데이터 무결성 체크 및 와이어 암호화가 필요한지 여부를 지정합니다.
      가능한 값은 '인증', '무결성', '암호화'입니다. KERBEROS가 활성화된 경우에만 사용됩니다.

      예시:
         | privacy | 인증, 무결성 및 암호화가 활성화됨.

   --hdfs-encoding
      백엔드용 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --hdfs-namenode
      Hadoop 네임 노드 및 포트입니다.
      
      예: "namenode:8020"는 포트 8020에서 실행 중인 namenode에 연결합니다.

   --hdfs-service-principal-name
      네임 노드를 위한 Kerberos 서비스 주체 이름입니다.
      
      KERBEROS 인증을 활성화합니다. 네임 노드를 위한 Service Principal Name(SERVICE/FQDN)을 지정합니다.
      예: 서비스 'hdfs'를 사용하고 FQDN 'namenode.hadoop.docker'인 네임 노드의 경우 "hdfs/namenode.hadoop.docker"입니다.

   --hdfs-username
      Hadoop 사용자 이름입니다.

      예시:
         | root | root로 HDFS에 연결합니다.


OPTIONS:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] CAR 파일로 내보낸 후 데이터셋 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공한 스캔 후 이 간격이 지나면 소스 디렉토리를 자동으로 다시 스캔합니다. (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비됨)

   hdfs 옵션

   --hdfs-data-transfer-protection value  Kerberos 데이터 전송 보호: 인증|무결성|암호화. [$HDFS_DATA_TRANSFER_PROTECTION]
   --hdfs-encoding value                  백엔드용 인코딩입니다. (기본값: "Slash,Colon,Del,Ctl,InvalidUtf8,Dot") [$HDFS_ENCODING]
   --hdfs-namenode value                  Hadoop 네임 노드 및 포트입니다. [$HDFS_NAMENODE]
   --hdfs-service-principal-name value    네임 노드를 위한 Kerberos 서비스 주체 이름입니다. [$HDFS_SERVICE_PRINCIPAL_NAME]
   --hdfs-username value                  Hadoop 사용자 이름입니다. [$HDFS_USERNAME]

```
{% endcode %}