# QingCloud Object Storage

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add qingstor - QingCloud Object Storage

USAGE:
   singularity datasource add qingstor [command options] <dataset_name> <source_path>

DESCRIPTION:
   --qingstor-access-key-id
     QingStor Access Key ID.
 
     익명 접근 또는 런타임 인증 정보를 사용하기 위해 공백으로 남겨둡니다.

   --qingstor-chunk-size
     업로드에 사용되는 청크 크기입니다.
     
     업로드할 파일의 크기가 업로드 임계값보다 큰 경우 이 청크 크기를 사용하여 멀티파트 업로드가 수행됩니다.
     
     알림: "--qingstor-upload-concurrency" 개수의 이러한 크기의 청크는 전송 당 메모리에 버퍼링됩니다.
     
     대역폭을 충분히 활용하고 있는 고속 링크로 대용량 파일을 전송하는 경우에는 메모리가 충분하다면 이 값을 높이면 전송 속도가 향상됩니다.

   --qingstor-connection-retries
     연결 재시도 횟수입니다.

   --qingstor-encoding
     백엔드의 인코딩입니다.
     
     자세한 정보는 [Overview](/overview/#encoding)의 "encoding" 섹션을 참조하십시오.

   --qingstor-endpoint
     QingStor API에 연결할 엔드포인트 URL을 입력합니다.
     
     비워둘 경우 기본값인 "https://qingstor.com:443"이 사용됩니다.

   --qingstor-env-auth
     런타임에서 QingStor 자격 증명을 가져옵니다.
     
     access_key_id와 secret_access_key가 공백인 경우에만 적용됩니다.
     
     예시:
        | false | 다음 단계에서 QingStor 자격 증명을 입력합니다.
        | true  | 환경 변수 또는 IAM에서 QingStor 자격 증명을 가져옵니다.

   --qingstor-secret-access-key
     QingStor Secret Access Key(암호)입니다.
     
     익명 접근 또는 런타임 인증 정보를 사용하기 위해 공백으로 남겨둡니다.

   --qingstor-upload-concurrency
     멀티파트 업로드의 동시성입니다.
     
     이 값은 동시에 업로드되는 동일한 파일의 청크 수입니다.
     
     참고: 이 값을 1보다 크게 설정하면 멀티파트 업로드의 체크섬이 손상됩니다(업로드 자체는 손상되지 않음).
     
     대역폭을 충분히 활용하지 못하는 고속 링크로 소수의 대용량 파일을 전송하는 경우에는 이 값을 높이는 것이 전송 속도 향상에 도움이 될 수 있습니다.

   --qingstor-upload-cutoff
     청크 업로드로 전환하는 임계값입니다.
     
     이 값보다 큰 파일은 청크 크기단위로 업로드됩니다.
     최소값은 0이고 최대값은 5 GiB입니다.

   --qingstor-zone
     연결할 존입니다.
     
     기본값은 "pek3a"입니다.

     예시:
        | pek3a | 베이징(중국) 3번째 존입니다.
                | 위치 제약 조건 "pek3a"가 필요합니다.
        | sh1a  | 상해(중국) 첫 번째 존입니다.
                | 위치 제약 조건 "sh1a"가 필요합니다.
        | gd2a  | 광둥(중국) 두 번째 존입니다.
                | 위치 제약 조건 "gd2a"가 필요합니다.


OPTIONS:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    CAR 파일로 내보낸 후 데이터셋 파일 삭제 [주의사항] (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔 후 이 간격만큼 경과하면 원본 디렉터리 자동으로 다시 스캔합니다 (기본값: 사용 안 함)
   --scanning-state value   초기 스캔 상태 설정 (기본값: ready)

   QingStor 옵션

   --qingstor-access-key-id value       QingStor Access Key ID. [$QINGSTOR_ACCESS_KEY_ID]
   --qingstor-chunk-size value          업로드에 사용되는 청크 크기(기본값: "4Mi") [$QINGSTOR_CHUNK_SIZE]
   --qingstor-connection-retries value  연결 재시도 횟수(기본값: "3") [$QINGSTOR_CONNECTION_RETRIES]
   --qingstor-encoding value            백엔드의 인코딩(기본값: "Slash,Ctl,InvalidUtf8") [$QINGSTOR_ENCODING]
   --qingstor-endpoint value            QingStor API에 연결할 엔드포인트 URL [$QINGSTOR_ENDPOINT]
   --qingstor-env-auth value            런타임에서 QingStor 자격 증명 가져오기(기본값: "false") [$QINGSTOR_ENV_AUTH]
   --qingstor-secret-access-key value   QingStor Secret Access Key(암호) [$QINGSTOR_SECRET_ACCESS_KEY]
   --qingstor-upload-concurrency value  멀티파트 업로드의 동시성(기본값: "1") [$QINGSTOR_UPLOAD_CONCURRENCY]
   --qingstor-upload-cutoff value       청크 업로드로 전환하는 임계값(기본값: "200Mi") [$QINGSTOR_UPLOAD_CUTOFF]
   --qingstor-zone value                연결할 존 [$QINGSTOR_ZONE]

```
{% endcode %}