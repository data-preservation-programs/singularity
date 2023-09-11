# 메타데이터 API에서 CAR 파일 다운로드

{% code fullWidth="true" %}
```
NAME:
   singularity download - 메타데이터 API에서 CAR 파일 다운로드

사용법:
   singularity download [command options] <piece_cid>

분류:
   유틸리티

옵션:
   1Fichier

   --fichier-api-key value          1Fichier 계정의 API 키. https://1fichier.com/console/params.pl에서 얻으세요. [$FICHIER_API_KEY]
   --fichier-encoding value         백엔드의 인코딩. (기본값: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$FICHIER_ENCODING]
   --fichier-file-password value    공유된 파일의 비밀번호가 있는 경우 다운로드하려면 이 매개변수를 추가하세요. [$FICHIER_FILE_PASSWORD]
   --fichier-folder-password value  공유된 폴더에서 파일 목록을 보려면 이 매개변수를 추가하세요. [$FICHIER_FOLDER_PASSWORD]
   --fichier-shared-folder value    공유 폴더를 다운로드하려면 이 매개변수를 추가하세요. [$FICHIER_SHARED_FOLDER]

   Akamai NetStorage

   --netstorage-account value   NetStorage 계정 이름을 설정하세요. [$NETSTORAGE_ACCOUNT]
   --netstorage-host value      연결할 NetStorage 호스트의 도메인+경로입니다. [$NETSTORAGE_HOST]
   --netstorage-protocol value  HTTP 또는 HTTPS 프로토콜을 선택하세요. (기본값: "https") [$NETSTORAGE_PROTOCOL]
   --netstorage-secret value    NetStorage 계정 비밀/G2O 키를 인증에 사용하세요. [$NETSTORAGE_SECRET]

   Amazon Drive

   --acd-auth-url value            인증 서버 URL입니다. [$ACD_AUTH_URL]
   --acd-checkpoint value          내부 폴링에 대한 체크포인트 (디버깅용). [$ACD_CHECKPOINT]
   --acd-client-id value           OAuth 클라이언트 ID입니다. [$ACD_CLIENT_ID]
   --acd-client-secret value       OAuth 클라이언트 시크릿입니다. [$ACD_CLIENT_SECRET]
   --acd-encoding value            백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ACD_ENCODING]
   --acd-templink-threshold value  이 크기보다 크거나 같은 파일은 tempLink를 통해 다운로드할 수 있습니다. (기본값: "9Gi") [$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               OAuth 액세스 토큰(JSON 형식)입니다. [$ACD_TOKEN]
   --acd-token-url value           토큰 서버 URL입니다. [$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  완료되지 않은 업로드 후 몇 분 기다렸다가 나타나는지 확인하기 위해 1GB당 추가로 기다리는 시간입니다. (기본값: "3m0s") [$ACD_UPLOAD_WAIT_PER_GB]

   Amazon S3 및 호환 스토리지 공급자(Alibaba, Ceph, China Mobile, Cloudflare, ArvanCloud, DigitalOcean, Dreamhost, Huawei OBS, IBM COS, IDrive e2, IONOS Cloud, Liara, Lyve Cloud, Minio, Netease, RackCorp, Scaleway, SeaweedFS, StackPath, Storj, Tencent COS, Qiniu 및 Wasabi 포함)

   --s3-access-key-id value            AWS 액세스 키 ID입니다. [$S3_ACCESS_KEY_ID]
   --s3-acl value                      버킷 및 개체를 생성할 때 사용되는 canned ACL입니다. [$S3_ACL]
   --s3-bucket-acl value               새 버킷을 생성할 때 사용되는 canned ACL입니다. [$S3_BUCKET_ACL]
   --s3-chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$S3_CHUNK_SIZE]
   --s3-copy-cutoff value              multipart 복사로 전환하기 위한 숫자. (기본값: "4.656Gi") [$S3_COPY_CUTOFF]
   --s3-decompress                     만약 설정하면 gzip으로 인코딩된 개체를 해제합니다. (기본값: false) [$S3_DECOMPRESS]
   --s3-disable-checksum               개체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$S3_DISABLE_CHECKSUM]
   --s3-disable-http2                  S3 백엔드에 대한 http2 사용을 비활성화합니다. (기본값: false) [$S3_DISABLE_HTTP2]
   --s3-download-url value