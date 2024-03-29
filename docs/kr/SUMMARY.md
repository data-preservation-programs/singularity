# 목차

## 개요 <a href="#overview" id="overview"></a>

* [Singularity란](README.md)
* [V1 또는 V2](overview/v1-or-v2.md)

## 설치 <a href="#installation" id="installation"></a>

* [바이너리 파일 다운로드](installation/download-binaries.md)
* [도커를 통한 설치](installation/install-from-docker.md)
* [소스 코드에서 빌드](installation/install-from-source.md)
* [운영 환경에 배포](installation/deploy-to-production.md)

## 데이터 준비 <a href="#data-preparation" id="data-preparation"></a>

* [시작하기](data-preparation/get-started.md)
* [성능 조정](data-preparation/performance-tuning.md)

## 콘텐츠 배포 <a href="#content-distribution" id="content-distribution"></a>

* [CAR 파일 배포하기](content-distribution/distribute-car-files.md)

## 거래 수립 <a href="#deal-making" id="deal-making"></a>

* [거래 일정 만들기](deal-making/create-a-deal-schedule.md)

## 주제 <a href="#topics" id="topics"></a>

* [인라인 준비](topics/inline-preparation.md)
* [벤치마크](topics/benchmark.md)

## 💻 CLI 참조 <a href="#cli-reference" id="cli-reference"></a>
<!-- cli begin -->

* [메뉴](cli-reference/README.md)
* [Ez Prep](cli-reference/ez-prep.md)
* [버전](cli-reference/version.md)
* [관리자](cli-reference/admin/README.md)
  * [초기화](cli-reference/admin/init.md)
  * [재설정](cli-reference/admin/reset.md)
  * [데이터셋 이전](cli-reference/admin/migrate-dataset.md)
  * [일정 이전](cli-reference/admin/migrate-schedule.md)
* [다운로드](cli-reference/download.md)
* [CAR 추출](cli-reference/extract-car.md)
* [거래](cli-reference/deal/README.md)
  * [일정](cli-reference/deal/schedule/README.md)
    * [생성](cli-reference/deal/schedule/create.md)
    * [목록](cli-reference/deal/schedule/list.md)
    * [수정](cli-reference/deal/schedule/update.md)
    * [일시 중지](cli-reference/deal/schedule/pause.md)
    * [재개](cli-reference/deal/schedule/resume.md)
  * [수동 전송](cli-reference/deal/send-manual.md)
  * [목록](cli-reference/deal/list.md)
* [실행](cli-reference/run/README.md)
  * [API](cli-reference/run/api.md)
  * [데이터셋 워커](cli-reference/run/dataset-worker.md)
  * [콘텐츠 제공자](cli-reference/run/content-provider.md)
  * [거래 추적기](cli-reference/run/deal-tracker.md)
  * [거래 푸셔](cli-reference/run/deal-pusher.md)
  * [다운로드 서버](cli-reference/run/download-server.md)
* [월렛](cli-reference/wallet/README.md)
  * [가져오기](cli-reference/wallet/import.md)
  * [목록](cli-reference/wallet/list.md)
  * [제거](cli-reference/wallet/remove.md)
* [저장소](cli-reference/storage/README.md)
  * [생성](cli-reference/storage/create/README.md)
    * [Acd](cli-reference/storage/create/acd.md)
    * [Azureblob](cli-reference/storage/create/azureblob.md)
    * [B2](cli-reference/storage/create/b2.md)
    * [Box](cli-reference/storage/create/box.md)
    * [Drive](cli-reference/storage/create/drive.md)
    * [Dropbox](cli-reference/storage/create/dropbox.md)
    * [Fichier](cli-reference/storage/create/fichier.md)
    * [Filefabric](cli-reference/storage/create/filefabric.md)
    * [Ftp](cli-reference/storage/create/ftp.md)
    * [Google Cloud Storage](cli-reference/storage/create/gcs.md)
    * [Gphotos](cli-reference/storage/create/gphotos.md)
    * [Hdfs](cli-reference/storage/create/hdfs.md)
    * [Hidrive](cli-reference/storage/create/hidrive.md)
    * [Http](cli-reference/storage/create/http.md)
    * [Internetarchive](cli-reference/storage/create/internetarchive.md)
    * [Jottacloud](cli-reference/storage/create/jottacloud.md)
    * [Koofr / Digi Storage](cli-reference/storage/create/koofr/README.md)
      * [Digistorage](cli-reference/storage/create/koofr/digistorage.md)
      * [Koofr / Digi Storage](cli-reference/storage/create/koofr/koofr.md)
      * [기타](cli-reference/storage/create/koofr/other.md)
    * [로컬](cli-reference/storage/create/local.md)
    * [Mailru](cli-reference/storage/create/mailru.md)
    * [Mega](cli-reference/storage/create/mega.md)
    * [Netstorage](cli-reference/storage/create/netstorage.md)
    * [Onedrive](cli-reference/storage/create/onedrive.md)
    * [Opendrive](cli-reference/storage/create/opendrive.md)
    * [Oos](cli-reference/storage/create/oos/README.md)
      * [Env_auth](cli-reference/storage/create/oos/env_auth.md)
      * [Instance_principal_auth](cli-reference/storage/create/oos/instance_principal_auth.md)
      * [No_auth](cli-reference/storage/create/oos/no_auth.md)
      * [Resource_principal_auth](cli-reference/storage/create/oos/resource_principal_auth.md)
      * [User_principal_auth](cli-reference/storage/create/oos/user_principal_auth.md)
    * [Pcloud](cli-reference/storage/create/pcloud.md)
    * [Premiumizeme](cli-reference/storage/create/premiumizeme.md)
    * [Putio](cli-reference/storage/create/putio.md)
    * [Qingstor](cli-reference/storage/create/qingstor.md)
    * [AWS S3 및 호환](cli-reference/storage/create/s3/README.md)
      * [Aws](cli-reference/storage/create/s3/aws.md)
      * [Alibaba](cli-reference/storage/create/s3/alibaba.md)
      * [Arvancloud](cli-reference/storage/create/s3/arvancloud.md)
      * [Ceph](cli-reference/storage/create/s3/ceph.md)
      * [Chinamobile](cli-reference/storage/create/s3/chinamobile.md)
      * [Cloudflare](cli-reference/storage/create/s3/cloudflare.md)
      * [Digitalocean](cli-reference/storage/create/s3/digitalocean.md)
      * [Dreamhost](cli-reference/storage/create/s3/dreamhost.md)
      * [Huaweiobs](cli-reference/storage/create/s3/huaweiobs.md)
      * [Ibmcos](cli-reference/storage/create/s3/ibmcos.md)
      * [Idrive](cli-reference/storage/create/s3/idrive.md)
      * [Ionos](cli-reference/storage/create/s3/ionos.md)
      * [Liara](cli-reference/storage/create/s3/liara.md)
      * [Lyvecloud](cli-reference/storage/create/s3/lyvecloud.md)
      * [Minio](cli-reference/storage/create/s3/minio.md)
      * [Netease](cli-reference/storage/create/s3/netease.md)
      * [기타](cli-reference/storage/create/s3/other.md)
      * [Qiniu](cli-reference/storage/create/s3/qiniu.md)
      * [Rackcorp](cli-reference/storage/create/s3/rackcorp.md)
      * [Scaleway](cli-reference/storage/create/s3/scaleway.md)
      * [Seaweedfs](cli-reference/storage/create/s3/seaweedfs.md)
      * [Stackpath](cli-reference/storage/create/s3/stackpath.md)
      * [Storj](cli-reference/storage/create/s3/storj.md)
      * [Tencentcos](cli-reference/storage/create/s3/tencentcos.md)
      * [Wasabi](cli-reference/storage/create/s3/wasabi.md)
    * [Seafile](cli-reference/storage/create/seafile.md)
    * [Sftp](cli-reference/storage/create/sftp.md)
    * [Sharefile](cli-reference/storage/create/sharefile.md)
    * [Sia](cli-reference/storage/create/sia.md)
    * [Smb](cli-reference/storage/create/smb.md)
    * [Storj](cli-reference/storage/create/storj/README.md)
      * [기존](cli-reference/storage/create/storj/existing.md)
      * [새로 만들기](cli-reference/storage/create/storj/new.md)
    * [Sugarsync](cli-reference/storage/create/sugarsync.md)
    * [Swift](cli-reference/storage/create/swift.md)
    * [Uptobox](cli-reference/storage/create/uptobox.md)
    * [Webdav](cli-reference/storage/create/webdav.md)
    * [Yandex](cli-reference/storage/create/yandex.md)
    * [Zoho](cli-reference/storage/create/zoho.md)
  * [탐색](cli-reference/storage/explore.md)
  * [목록](cli-reference/storage/list.md)
  * [제거](cli-reference/storage/remove.md)
  * [갱신](cli-reference/storage/update/README.md)
    * [Acd](cli-reference/storage/update/acd.md)
    * [Azureblob](cli-reference/storage/update/azureblob.md)
    * [B2](cli-reference/storage/update/b2.md)
    * [Box](cli-reference/storage/update/box.md)
    * [Drive](cli-reference/storage/update/drive.md)
    * [Dropbox](cli-reference/storage/update/dropbox.md)
    * [Fichier](cli-reference/storage/update/fichier.md)
    * [Filefabric](cli-reference/storage/update/filefabric.md)
    * [Ftp](cli-reference/storage/update/ftp.md)
    * [Google Cloud Storage](cli-reference/storage/update/gcs.md)
    * [Gphotos](cli-reference/storage/update/gphotos.md)
    * [Hdfs](cli-reference/storage/update/hdfs.md)
    * [Hidrive](cli-reference/storage/update/hidrive.md)
    * [Http](cli-reference/storage/update/http.md)
    * [Internetarchive](cli-reference/storage/update/internetarchive.md)
    * [Jottacloud](cli-reference/storage/update/jottacloud.md)
    * [Koofr / Digi Storage](cli-reference/storage/update/koofr/README.md)
      * [Digistorage](cli-reference/storage/update/koofr/digistorage.md)
      * [Koofr / Digi Storage](cli-reference/storage/update/koofr/koofr.md)
      * [기타](cli-reference/storage/update/koofr/other.md)
    * [로컬](cli-reference/storage/update/local.md)
    * [Mailru](cli-reference/storage/update/mailru.md)
    * [Mega](cli-reference/storage/update/mega.md)
    * [Netstorage](cli-reference/storage/update/netstorage.md)
    * [Onedrive](cli-reference/storage/update/onedrive.md)
    * [Opendrive](cli-reference/storage/update/opendrive.md)
    * [Oos](cli-reference/storage/update/oos/README.md)
      * [Env_auth](cli-reference/storage/update/oos/env_auth.md)
      * [Instance_principal_auth](cli-reference/storage/update/oos/instance_principal_auth.md)
      * [No_auth](cli-reference/storage/update/oos/no_auth.md)
      * [Resource_principal_auth](cli-reference/storage/update/oos/resource_principal_auth.md)
      * [User_principal_auth](cli-reference/storage/update/oos/user_principal_auth.md)
    * [Pcloud](cli-reference/storage/update/pcloud.md)
    * [Premiumizeme](cli-reference/storage/update/premiumizeme.md)
    * [Putio](cli-reference/storage/update/putio.md)
    * [Qingstor](cli-reference/storage/update/qingstor.md)
    * [AWS S3 및 호환](cli-reference/storage/update/s3/README.md)
      * [Aws](cli-reference/storage/update/s3/aws.md)
      * [Alibaba](cli-reference/storage/update/s3/alibaba.md)
      * [Arvancloud](cli-reference/storage/update/s3/arvancloud.md)
      * [Ceph](cli-reference/storage/update/s3/ceph.md)
      * [Chinamobile](cli-reference/storage/update/s3/chinamobile.md)
      * [Cloudflare](cli-reference/storage/update/s3/cloudflare.md)
      * [Digitalocean](cli-reference/storage/update/s3/digitalocean.md)
      * [Dreamhost](cli-reference/storage/update/s3/dreamhost.md)
      * [Huaweiobs](cli-reference/storage/update/s3/huaweiobs.md)
      * [Ibmcos](cli-reference/storage/update/s3/ibmcos.md)
      * [Idrive](cli-reference/storage/update/s3/idrive.md)
      * [Ionos](cli-reference/storage/update/s3/ionos.md)
      * [Liara](cli-reference/storage/update/s3/liara.md)
      * [Lyvecloud](cli-reference/storage/update/s3/lyvecloud.md)
      * [Minio](cli-reference/storage/update/s3/minio.md)
      * [Netease](cli-reference/storage/update/s3/netease.md)
      * [기타](cli-reference/storage/update/s3/other.md)
      * [Qiniu](cli-reference/storage/update/s3/qiniu.md)
      * [Rackcorp](cli-reference/storage/update/s3/rackcorp.md)
      * [Scaleway](cli-reference/storage/update/s3/scaleway.md)
      * [Seaweedfs](cli-reference/storage/update/s3/seaweedfs.md)
      * [Stackpath](cli-reference/storage/update/s3/stackpath.md)
      * [Storj](cli-reference/storage/update/s3/storj.md)
      * [Tencentcos](cli-reference/storage/update/s3/tencentcos.md)
      * [Wasabi](cli-reference/storage/update/s3/wasabi.md)
    * [Seafile](cli-reference/storage/update/seafile.md)
    * [Sftp](cli-reference/storage/update/sftp.md)
    * [Sharefile](cli-reference/storage/update/sharefile.md)
    * [Sia](cli-reference/storage/update/sia.md)
    * [Smb](cli-reference/storage/update/smb.md)
    * [Storj](cli-reference/storage/update/storj/README.md)
      * [기존](cli-reference/storage/update/storj/existing.md)
      * [새로 만들기](cli-reference/storage/update/storj/new.md)
    * [Sugarsync](cli-reference/storage/update/sugarsync.md)
    * [Swift](cli-reference/storage/update/swift.md)
    * [Uptobox](cli-reference/storage/update/uptobox.md)
    * [Webdav](cli-reference/storage/update/webdav.md)
    * [Yandex](cli-reference/storage/update/yandex.md)
    * [Zoho](cli-reference/storage/update/zoho.md)
  * [이름 변경](cli-reference/storage/rename.md)
* [Prep](cli-reference/prep/README.md)
  * [생성](cli-reference/prep/create.md)
  * [목록](cli-reference/prep/list.md)
  * [상태](cli-reference/prep/status.md)
  * [이름 변경](cli-reference/prep/rename.md)
  * [소스 연결](cli-reference/prep/attach-source.md)
  * [출력 연결](cli-reference/prep/attach-output.md)
  * [출력 분리](cli-reference/prep/detach-output.md)
  * [스캔 시작](cli-reference/prep/start-scan.md)
  * [스캔 일시 중지](cli-reference/prep/pause-scan.md)
  * [팩 생성 시작](cli-reference/prep/start-pack.md)
  * [팩 일시 중지](cli-reference/prep/pause-pack.md)
  * [Daggen 시작](cli-reference/prep/start-daggen.md)
  * [Daggen 일시 중지](cli-reference/prep/pause-daggen.md)
  * [조각 목록](cli-reference/prep/list-pieces.md)
  * [조각 추가](cli-reference/prep/add-piece.md)
  * [탐색](cli-reference/prep/explore.md)
  * [월렛 연결](cli-reference/prep/attach-wallet.md)
  * [월렛 목록](cli-reference/prep/list-wallets.md)
  * [월렛 분리](cli-reference/prep/detach-wallet.md)

<!-- cli end -->

## 🌐 웹 API 참조 <a href="#web-api-reference" id="web-api-reference"></a>
<!-- webapi begin -->

* [거래 일정](web-api-reference/deal-schedule.md)
* [거래](web-api-reference/deal.md)
* [파일](web-api-reference/file.md)
* [작업](web-api-reference/job.md)
* [조각](web-api-reference/piece.md)
* [준비 작업](web-api-reference/preparation.md)
* [저장소](web-api-reference/storage.md)
* [월렛 연결](web-api-reference/wallet-association.md)
* [월렛](web-api-reference/wallet.md)
* [사양](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)

<!-- webapi end -->

## ❓ FAQ <a href="#faq" id="faq"></a>

* [데이터베이스가 잠김](faq/database-is-locked.md)