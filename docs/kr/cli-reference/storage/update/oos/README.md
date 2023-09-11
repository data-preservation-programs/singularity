# Oracle Cloud Infrastructure Object Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage update oos - Oracle Cloud Infrastructure Object Storage

사용법:
   singularity storage update oos command [command options] [arguments...]

명령어:
   env_auth                 런타임 환경에서 자격 증명을 자동으로 가져옵니다(env). 자격 증명을 제공하는 첫 번째 사람이 인증에 성공합니다.
   instance_principal_auth  인스턴스 주체를 사용하여 API 호출을 승인합니다.
                            각 인스턴스는 고유한 ID를 갖고 있으며, 인스턴스 메타데이터에 있는 인증서를 사용하여 인증합니다.
                            https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm
   no_auth                  자격 증명이 필요하지 않습니다. 일반적으로 공개 버킷을 읽는 데 사용됩니다.
   resource_principal_auth  리소스 주체를 사용하여 API 호출을 수행합니다.
   user_principal_auth      OCI 사용자와 API 키를 사용하여 인증합니다.
                            테넌시 OCID, 사용자 OCID, 리전, 경로, API 키의 지문을 구성 파일에 입력해야 합니다.
                            https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm
   help, h                  명령어의 목록을 표시하거나 특정 명령어에 대한 도움말을 표시합니다.

옵션:
   --help, -h  도움말 표시
```
{% endcode %}