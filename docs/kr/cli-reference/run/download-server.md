# 원격 메타데이터 API에 연결하여 CAR 파일 다운로드를 제공하는 HTTP 서버

{% code fullWidth="true" %}
```
NAME:
   singularity run download-server - 원격 메타데이터 API에 연결하여 CAR 파일 다운로드를 제공하는 HTTP 서버

사용법:
   singularity run download-server [command options] [arguments...]

설명:
   사용 예시:
      singularity run download-server --metadata-api "http://remote-metadata-api:7777" --bind "127.0.0.1:8888"

옵션:
   --help, -h  도움말 표시

   1Fichier

   --fichier-api-key value        API 키. https://1fichier.com/console/params.pl에서 가져옵니다. [$FICHIER_API_KEY]
   --fichier-file-password value  암호로 보호된 공유 파일을 다운로드하려면이 매개변수를 추가하십시오. [$FICHIER_FILE_PASSWORD]
   --fichier-folder-password value    암호로 보호된 공유 폴더의 파일 목록을 표시하려면이 매개변수를 추가하십시오. [$FICHIER_FOLDER_PASSWORD]

   Akamai NetStorage

   --netstorage-secret value  인증을 위한 NetStorage 계정 비밀 및 G2O 키 설정. [$NETSTORAGE_SECRET]

   Amazon Drive

   --acd-client-secret value  OAuth 클라이언트 비밀. [$ACD_CLIENT_SECRET]
   --acd-token value          JSON blob 형식의 OAuth 액세스 토큰. [$ACD_TOKEN]
   --acd-token-url value      토큰 서버 URL. [$ACD_TOKEN_URL]

   Amazon S3 호환 스토리지 제공자 (AWS, Alibaba, Ceph, China Mobile, Cloudflare, ArvanCloud, DigitalOcean, Dreamhost, Huawei OBS, IBM COS, IDrive e2, IONOS Cloud, Liara, Lyve Cloud, Minio, Netease, RackCorp, Scaleway, SeaweedFS, StackPath, Storj, Tencent COS, Qiniu, Wasabi 등)

   --s3-access-key-id value            AWS 엑세스 키 ID. [$S3_ACCESS_KEY_ID]
   --s3-secret-access-key value        AWS 비밀 액세스 키 (비밀번호). [$S3_SECRET_ACCESS_KEY]
   --s3-session-token value            AWS 세션 토큰. [$S3_SESSION_TOKEN]
   --s3-sse-customer-key value         데이터를 암호화/복호화하는 데 사용되는 비밀 암호화 키를 제공하려면 SSE-C를 사용하십시오. [$S3_SSE_CUSTOMER_KEY]
   --s3-sse-customer-key-base64 value  SSE-C를 사용하는 경우, 데이터를 암호화/복호화하는데 사용되는 비밀 암호화 키를 base64 형식으로 인코딩하여 제공하십시오. [$S3_SSE_CUSTOMER_KEY_BASE64]
   --s3-sse-customer-key-md5 value     SSE-C를 사용하는 경우, 비밀 암호화 키의 MD5 체크섬을 제공할 수 있습니다 (선택 사항). [$S3_SSE_CUSTOMER_KEY_MD5]
   --s3-sse-kms-key-id value           KMS ID를 사용하는 경우, 키의 ARN을 제공해야 합니다. [$S3_SSE_KMS_KEY_ID]

   Backblaze B2

   --b2-key value  응용 프로그램 키. [$B2_KEY]

   Box

   --box-access-token value   Box 앱 기본 액세스 토큰 [$BOX_ACCESS_TOKEN]
   --box-client-secret value  OAuth 클라이언트 비밀. [$BOX_CLIENT_SECRET]
   --box-token value          JSON blob 형식의 OAuth 액세스 토큰. [$BOX_TOKEN]
   --box-token-url value      토큰 서버 URL. [$BOX_TOKEN_URL]

   클라이언트 구성

   --client-ca-cert value                           서버 인증서 확인에 사용되는 CA 인증서 경로. 삭제하려면 빈 문자열을 사용하십시오.
   --client-cert value                              상호 TLS 인증을 위한 클라이언트 SSL 인증서 (PEM) 경로입니다. 삭제하려면 빈 문자열을 사용하십시오.
   --client-connect-timeout value                   HTTP 클라이언트 연결 제한시간 (기본값: 1분)
   --client-expect-continue-timeout value           HTTP에서 100-continue를 사용하는 경우의 타임아웃 (기본값: 1초)
   --client-header value [ --client-header value ]  모든 트랜잭션에 사용할 HTTP 헤더를 설정합니다 (예: key=value). 기존 헤더 값을 대체합니다. 헤더를 삭제하려면 --http-header "key=""를 사용하십시오. 모든 헤더를 삭제하려면 --http-header ""를 사용하십시오.
   --client-insecure-skip-verify                    서버 SSL 인증서를 확인하지 않습니다 (보안 위험) (기본값: false)
   --client-key value                               상호 TLS 인증을 위한 클라이언트 SSL 개인 키 (PEM) 경로입니다. 삭제하려면 빈 문자열을 사용하십시오.
   --client-no-gzip                                 Accept-Encoding: gzip을 설정하지 않습니다 (기본값: false)
   --client-scan-concurrency value                  데이터 소스 스캔 시 동시 목록 요청의 최대 수 (기본값: 1)
   --client-timeout value                           IO 유휴 시간 제한 (기본값: 5분)
   --client-use-server-mod-time                     가능한 경우 서버 수정 시간 사용 (기본값: false)
   --client-user-agent value                        사용자 에이전트를 지정된 문자열로 설정합니다. 삭제하려면 빈 문자열을 사용하십시오. (기본값: rclone/v1.62.2-DEV)

   Dropbox

   --dropbox-client-secret value  OAuth 클라이언트 비밀. [$DROPBOX_CLIENT_SECRET]
   --dropbox-token value          JSON blob 형식의 OAuth 액세스 토큰. [$DROPBOX_TOKEN]
   --dropbox-token-url value      토큰 서버 URL. [$DROPBOX_TOKEN_URL]

   기업용 File Fabric

   --filefabric-permanent-token value  영구 인증 토큰. [$FILEFABRIC_PERMANENT_TOKEN]
   --filefabric-token value            세션 토큰. [$FILEFABRIC_TOKEN]
   --filefabric-token-expiry value     토큰 만료 시간. [$FILEFABRIC_TOKEN_EXPIRY]

   FTP

   --ftp-ask-password  FTP 비밀번호 요청 허용 (기본값: false) [$FTP_ASK_PASSWORD]
   --ftp-pass value    FTP 비밀번호. [$FTP_PASS]

   일반 구성

   --bind value          HTTP 서버를 바인드할 주소 (기본값: "127.0.0.1:8888")
   --metadata-api value  메타데이터 API의 URL (기본값: "http://127.0.0.1:7777")

   Google Cloud Storage (Google Drive가 아님)

   --gcs-client-secret value  OAuth 클라이언트 비밀. [$GCS_CLIENT_SECRET]
   --gcs-token value          JSON blob 형식의 OAuth 액세스 토큰. [$GCS_TOKEN]
   --gcs-token-url value      토큰 서버 URL. [$GCS_TOKEN_URL]

   Google Drive

   --drive-client-secret value  OAuth 클라이언트 비밀. [$DRIVE_CLIENT_SECRET]
   --drive-resource-key value   링크 공유 파일에 액세스하기 위한 리소스 키. [$DRIVE_RESOURCE_KEY]
   --drive-token value          JSON blob 형식의 OAuth 액세스 토큰. [$DRIVE_TOKEN]
   --drive-token-url value      토큰 서버 URL. [$DRIVE_TOKEN_URL]

   Google 사진

   --gphotos-client-secret value  OAuth 클라이언트 비밀. [$GPHOTOS_CLIENT_SECRET]
   --gphotos-token value          JSON blob 형식의 OAuth 액세스 토큰. [$GPHOTOS_TOKEN]
   --gphotos-token-url value      토큰 서버 URL. [$GPHOTOS_TOKEN_URL]

   HiDrive

   --hidrive-client-secret value  OAuth 클라이언트 비밀. [$HIDRIVE_CLIENT_SECRET]
   --hidrive-token value          JSON blob 형식의 OAuth 액세스 토큰. [$HIDRIVE_TOKEN]
   --hidrive-token-url value      토큰 서버 URL. [$HIDRIVE_TOKEN_URL]

   Internet Archive

   --internetarchive-access-key-id value      IAS3 액세스 키. [$INTERNETARCHIVE_ACCESS_KEY_ID]
   --internetarchive-secret-access-key value  IAS3 비밀 키 (비밀번호). [$INTERNETARCHIVE_SECRET_ACCESS_KEY]

   Koofr, Digi Storage 및 기타 Koofr 호환 스토리지 제공자

   --koofr-password value  rclone의 비밀번호 (https://storage.rcs-rds.ro/app/admin/preferences/password에서 생성). [$KOOFR_PASSWORD]

   Mail.ru Cloud

   --mailru-pass value  비밀번호. [$MAILRU_PASS]

   Mega

   --mega-pass value  비밀번호. [$MEGA_PASS]

   Microsoft Azure Blob Storage

   --azureblob-client-certificate-password value  인증서 파일의 비밀번호 (선택 사항). [$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-secret value                서비스 프린시펄의 클라이언트 시크릿 중 하나 [$AZUREBLOB_CLIENT_SECRET]
   --azureblob-key value                          스토리지 계정 공유 키. [$AZUREBLOB_KEY]
   --azureblob-password value                     사용자의 비밀번호. [$AZUREBLOB_PASSWORD]

   Microsoft OneDrive

   --onedrive-client-secret value  OAuth 클라이언트 비밀. [$ONEDRIVE_CLIENT_SECRET]
   --onedrive-link-password value  링크 명령으로 생성된 링크의 비밀번호를 설정합니다. [$ONEDRIVE_LINK_PASSWORD]
   --onedrive-token value          JSON blob 형식의 OAuth 액세스 토큰. [$ONEDRIVE_TOKEN]
   --onedrive-token-url value      토큰 서버 URL. [$ONEDRIVE_TOKEN_URL]

   OpenDrive

   --opendrive-password value  비밀번호. [$OPENDRIVE_PASSWORD]

   OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

   --swift-application-credential-secret value  애플리케이션 자격 증명 비밀 (OS_APPLICATION_CREDENTIAL_SECRET). [$SWIFT_APPLICATION_CREDENTIAL_SECRET]
   --swift-auth-token value                     대체 인증을 위한 인증 토큰 - 선택 사항 (OS_AUTH_TOKEN). [$SWIFT_AUTH_TOKEN]
   --swift-key value                            API 키 또는 비밀번호 (OS_PASSWORD). [$SWIFT_KEY]

   Oracle Cloud Infrastructure Object Storage

   --oos-sse-customer-key value         SSE-C를 사용하려면 선택적 헤더인 256비트 암호화 키를 base64로 인코딩하여 제공합니다. [$OOS_SSE_CUSTOMER_KEY]
   --oos-sse-customer-key-file value    SSE-C를 사용하려면 기존 base64로 인코딩된 AES-256 암호화 키를 포함하는 파일을 제공합니다. [$OOS_SSE_CUSTOMER_KEY_FILE]
   --oos-sse-customer-key-sha256 value  SSE-C를 사용하는 경우, 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다. [$OOS_SSE_CUSTOMER_KEY_SHA256]
   --oos-sse-kms-key-id value           본인의 마스터 키를 사용하는 경우, 이 헤더로 키의 ARN을 제공할 수 있습니다. [$OOS_SSE_KMS_KEY_ID]

   Pcloud

   --pcloud-client-secret value  OAuth 클라이언트 비밀. [$PCLOUD_CLIENT_SECRET]
   --pcloud-password value       pcloud 비밀번호. [$PCLOUD_PASSWORD]
   --pcloud-token value          JSON blob 형식의 OAuth 액세스 토큰. [$PCLOUD_TOKEN]
   --pcloud-token-url value      토큰 서버 URL. [$PCLOUD_TOKEN_URL]

   QingCloud Object Storage

   --qingstor-access-key-id value      QingStor Access Key ID. [$QINGSTOR_ACCESS_KEY_ID]
   --qingstor-secret-access-key value  QingStor Secret Access Key (비밀번호). [$QINGSTOR_SECRET_ACCESS_KEY]

   재시도 전략

   --client-low-level-retries value  저수준 클라이언트 오류에 대한 최대 재시도 횟수 (기본값: 10)
   --client-retry-backoff value      IO 읽기 오류를 재시도할 때 사용하는 일정한 지연 시간 (기본값: 1초)
   --client-retry-backoff-exp value  IO 읽기 오류를 재시도할 때 사용하는 지수적인 지연 시간 (기본값: 1.0)
   --client-retry-delay value        IO 읽기 오류를 재시도하기 전의 초기 지연 시간 (기본값: 1초)
   --client-retry-max value          IO 읽기 오류에 대한 최대 재시도 횟수 (기본값: 10)
   --client-skip-inaccessible        열릴 때 접근할 수 없는 파일 건너뛰기 (기본값: false)

   SMB / CIFS

   --smb-pass value  SMB 비밀번호. [$SMB_PASS]

   SSH/SFTP

   --sftp-ask-password         SFTP 비밀번호 요청 허용 (기본값: false) [$SFTP_ASK_PASSWORD]
   --sftp-key-exchange value   우선 순위에 따라 공백으로 구분된 키 교환 알고리즘 목록. [$SFTP_KEY_EXCHANGE]
   --sftp-key-file value       PEM 인코딩된 개인 키 파일 경로. [$SFTP_KEY_FILE]
   --sftp-key-file-pass value  PEM 인코딩된 개인 키 파일을 복호화하기 위한 암호. [$SFTP_KEY_FILE_PASS]
   --sftp-key-pem value        PEM 인코딩된 개인 키. [$SFTP_KEY_PEM]
   --sftp-key-use-agent        ssh-agent 사용을 강제로 설정합니다. (기본값: false) [$SFTP_KEY_USE_AGENT]
   --sftp-pass value           SSH 비밀번호, ssh-agent를 사용하려면 비워 두십시오. [$SFTP_PASS]
   --sftp-pubkey-file value    선택 사항으로 공개 키 파일 경로입니다. [$SFTP_PUBKEY_FILE]

   Sia 분산 클라우드

   --sia-api-password value  Sia 데몬 API 비밀번호. [$SIA_API_PASSWORD]

   Storj 분산 클라우드 스토리지

   --storj-api-key value     API 키. [$STORJ_API_KEY]
   --storj-passphrase value  암호화 암호. [$STORJ_PASSPHRASE]

   Sugarsync

   --sugarsync-access-key-id value       Sugarsync 액세스 키 ID. [$SUGARSYNC_ACCESS_KEY_ID]
   --sugarsync-private-access-key value  Sugarsync 개인 액세스 키. [$SUGARSYNC_PRIVATE_ACCESS_KEY]
   --sugarsync-refresh-token value       Sugarsync 갱신 토큰. [$SUGARSYNC_REFRESH_TOKEN]

   Uptobox

   --uptobox-access-token value  액세스 토큰. [$UPTOBOX_ACCESS_TOKEN]

   WebDAV

   --webdav-bearer-token value          사용자/암호 대신 Bearer 토큰을 지정하십시오 (예: Macaroon). [$WEBDAV_BEARER_TOKEN]
   --webdav-bearer-token-command value  Bearer 토큰을 얻기 위해 실행할 명령어. [$WEBDAV_BEARER_TOKEN_COMMAND]
   --webdav-pass value                  비밀번호. [$WEBDAV_PASS]

   Yandex Disk

   --yandex-client-secret value  OAuth 클라이언트 비밀. [$YANDEX_CLIENT_SECRET]
   --yandex-token value          JSON blob 형식의 OAuth 액세스 토큰. [$YANDEX_TOKEN]
   --yandex-token-url value      토큰 서버 URL. [$YANDEX_TOKEN_URL]

   Zoho

   --zoho-client-secret value  OAuth 클라이언트 비밀. [$ZOHO_CLIENT_SECRET]
   --zoho-token value          JSON blob 형식의 OAuth 액세스 토큰. [$ZOHO_TOKEN]
   --zoho-token-url value      토큰 서버 URL. [$ZOHO_TOKEN_URL]

   premiumize.me

   --premiumizeme-api-key value  API 키. [$PREMIUMIZEME_API_KEY]

   seafile

   --seafile-auth-token value   인증 토큰. [$SEAFILE_AUTH_TOKEN]
   --seafile-library-key value  라이브러리 암호 (암호화된 라이브러리 전용). [$SEAFILE_LIBRARY_KEY]
   --seafile-pass value         비밀번호. [$SEAFILE_PASS]

```
{% endcode %}