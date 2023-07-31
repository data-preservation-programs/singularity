# 소스의 구성 옵션 수정

{% code fullWidth="true" %}
```
이름:
   singularity datasource update - 소스의 구성 옵션 수정

사용법:
   singularity datasource update [command options] <source_id>

옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    CAR 파일로 내보낸 후 데이터 세트의 파일 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공한 스캔으로부터 이 간격이 경과하면 소스 디렉터리를 자동으로 스캔합니다. (기본값: disabled)
   --scanning-state value   초기 스캔 상태 설정합니다. (기본값: ready)

   acd 옵션

   --acd-auth-url value            인증 서버 URL입니다. [$ACD_AUTH_URL]
   --acd-client-id value           OAuth 클라이언트 ID입니다. [$ACD_CLIENT_ID]
   --acd-client-secret value       OAuth 클라이언트 비밀입니다. [$ACD_CLIENT_SECRET]
   --acd-encoding value            백엔드에 대한 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ACD_ENCODING]
   --acd-templink-threshold value  해당 크기 이상인 파일은 tempLink를 통해 다운로드됩니다. (기본값: "9Gi") [$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               OAuth 액세스 토큰입니다. JSON blob 형식으로 입력합니다. [$ACD_TOKEN]
   --acd-token-url value           토큰 서버 URL입니다. [$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  완전한 업로드가 실패한 후 GiB 당 추가 대기시간입니다. (기본값: "3m0s") [$ACD_UPLOAD_WAIT_PER_GB]

   azureblob 옵션

   --azureblob-access-tier value                    blob의 액세스 수준: hot, cool 또는 archive입니다. [$AZUREBLOB_ACCESS_TIER]
   --azureblob-account value                        Azure Storage Account 이름입니다. [$AZUREBLOB_ACCOUNT]
   --azureblob-archive-tier-delete value            덮어쓰기 전에 아카이브 계층 blob을 삭제합니다. (기본값: "false") [$AZUREBLOB_ARCHIVE_TIER_DELETE]
   --azureblob-chunk-size value                     업로드 청크 크기입니다. (기본값: "4Mi") [$AZUREBLOB_CHUNK_SIZE]
   --azureblob-client-certificate-password value    인증서 파일의 비밀번호(선택 사항)입니다. [$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-certificate-path value        PEM 또는 PKCS12 인증서 파일의 경로를 지정합니다. [$AZUREBLOB_CLIENT_CERTIFICATE_PATH]
   --azureblob-client-id value                      사용 중인 클라이언트의 ID입니다. [$AZUREBLOB_CLIENT_ID]
   --azureblob-client-secret value                  서비스 프린시팔의 클라이언트 비밀 중 하나입니다. [$AZUREBLOB_CLIENT_SECRET]
   --azureblob-client-send-certificate-chain value  인증서 인증을 사용할 때 설명서 체인을 보냅니다. (기본값: "false") [$AZUREBLOB_CLIENT_SEND_CERTIFICATE_CHAIN]
   --azureblob-disable-checksum value               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: "false") [$AZUREBLOB_DISABLE_CHECKSUM]
   --azureblob-encoding value                       백엔드에 대한 인코딩입니다. (기본값: "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8") [$AZUREBLOB_ENCODING]
   --azureblob-endpoint value                       서비스를위한 엔드포인트입니다. [$AZUREBLOB_ENDPOINT]
   --azureblob-env-auth value                       런타임에서 자격 증명을 읽어옵니다.(환경 변수, CLI 또는 MSI) (기본값: "false") [$AZUREBLOB_ENV_AUTH]
   --azureblob-key value                            저장소 계정 공유 키입니다. [$AZUREBLOB_KEY]
   --azureblob-list-chunk value                     blob 목록의 크기입니다. (기본값: "5000") [$AZUREBLOB_LIST_CHUNK]
   --azureblob-memory-pool-flush-time value         내부 메모리 버퍼 풀이 플러시 될 수있는 시간 간격입니다. (기본값: "1m0s") [$AZUREBLOB_MEMORY_POOL_FLUSH_TIME]
   --azureblob-memory-pool-use-mmap value           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: "false") [$AZUREBLOB_MEMORY_POOL_USE_MMAP]
   --azureblob-msi-client-id value                  사용할 사용자 지정 MSI의 Object ID입니다. [$AZUREBLOB_MSI_CLIENT_ID]
   --azureblob-msi-mi-res-id value                  사용할 사용자 지정 MSI의 Azure 리소스 ID입니다. [$AZUREBLOB_MSI_MI_RES_ID]
   --azureblob-msi-object-id value                  사용할 사용자 지정 MSI의 Object ID입니다. [$AZUREBLOB_MSI_OBJECT_ID]
   --azureblob-no-check-container value             컨테이너의