# 목차

## 개요 <a href="#overview" id="overview"></a>

* [Singularity란](README.md)
* [V1 또는 V2](overview/v1-or-v2.md)
* [현재 상태](overview/current-status.md)
* [관련 프로젝트](overview/related-projects.md)

## 설치 <a href="#installation" id="installation"></a>

* [소스에서 설치하기](installation/install-from-source.md)
* [도커에서 설치하기](installation/install-from-docker.md)
* [운영 환경으로 배포하기](installation/deploy-to-production.md)

## 데이터 준비 <a href="#data-preparation" id="data-preparation"></a>

* [데이터셋 생성하기](data-preparation/create-a-dataset.md)
* [데이터 소스 추가하기](data-preparation/add-a-data-source.md)
* [데이터셋 워커 시작하기](data-preparation/start-dataset-worker.md)
* [데이터 소스를 위한 DAG 생성하기](data-preparation/create-dag-for-the-data-source.md)

## 컨텐츠 배포 <a href="#content-distribution" id="content-distribution"></a>

* [CAR 파일 배포하기](content-distribution/distribute-car-files.md)
* [파일 검색 (스테이징)](content-distribution/file-retrieval-staging.md)

## 거래 생성 <a href="#deal-making" id="deal-making"></a>

* [거래 생성 전제조건](deal-making/deal-making-prerequisite.md)
* [거래 일정 생성하기](deal-making/create-a-deal-schedule.md)
* [SP 셀프 서비스](deal-making/sp-self-service.md)

## 검색 <a href="#retrieval" id="retrieval"></a>

* [개요](retrieval/overview.md)

## 주제 <a href="#topics" id="topics"></a>

* [암호화](topics/encryption.md)
* [인라인 처리](topics/inline-preparation.md)
* [데이터소스 다시 스캔](topics/datasource-rescan.md)
* [푸시 및 업로드](topics/push-and-upload.md)
* [성능 평가](topics/benchmark.md)

## 💻 CLI 참조 <a href="#cli-reference" id="cli-reference"></a>
<!-- cli begin -->

* [메뉴](cli-reference/README.md)
* [Ez Prep](cli-reference/ez-prep.md)
* [관리자](cli-reference/admin/README.md)
  * [초기화](cli-reference/admin/init.md)
  * [초기화](cli-reference/admin/reset.md)
  * [마이그레이션](cli-reference/admin/migrate.md)
* [버전](cli-reference/version.md)
* [다운로드](cli-reference/download.md)
* [거래](cli-reference/deal/README.md)
  * [일정](cli-reference/deal/schedule/README.md)
    * [생성](cli-reference/deal/schedule/create.md)
    * [목록](cli-reference/deal/schedule/list.md)
    * [일시정지](cli-reference/deal/schedule/pause.md)
    * [다시 시작](cli-reference/deal/schedule/resume.md)
  * [Spade 정책](cli-reference/deal/spade-policy/README.md)
    * [생성](cli-reference/deal/spade-policy/create.md)
    * [목록](cli-reference/deal/spade-policy/list.md)
    * [제거](cli-reference/deal/spade-policy/remove.md)
  * [수동 전송](cli-reference/deal/send-manual.md)
  * [목록](cli-reference/deal/list.md)
* [실행](cli-reference/run/README.md)
  * [API](cli-reference/run/api.md)
  * [데이터셋 워커](cli-reference/run/dataset-worker.md)
  * [컨텐츠 프로바이더](cli-reference/run/content-provider.md)
  * [거래 추적기](cli-reference/run/deal-tracker.md)
  * [Dealmaker](cli-reference/run/dealmaker.md)
  * [Spade API](cli-reference/run/spade-api.md)
* [데이터셋](cli-reference/dataset/README.md)
  * [생성](cli-reference/dataset/create.md)
  * [목록](cli-reference/dataset/list.md)
  * [업데이트](cli-reference/dataset/update.md)
  * [제거](cli-reference/dataset/remove.md)
  * [월렛 추가](cli-reference/dataset/add-wallet.md)
  * [월렛 목록](cli-reference/dataset/list-wallet.md)
  * [월렛 제거](cli-reference/dataset/remove-wallet.md)
  * [조각 추가](cli-reference/dataset/add-piece.md)
  * [조각 목록](cli-reference/dataset/list-pieces.md)
* [데이터소스](cli-reference/datasource/README.md)
  * [추가](cli-reference/datasource/add/README.md)
    * [Amazon Drive](cli-reference/datasource/add/acd.md)
    * [Microsoft Azure Blob Storage](cli-reference/datasource/add/azureblob.md)
    * [Backblaze B2](cli-reference/datasource/add/b2.md)
    * [Box](cli-reference/datasource/add/box.md)
    * [Google Drive](cli-reference/datasource/add/drive.md)
    * [Dropbox](cli-reference/datasource/add/dropbox.md)
    * [1Fichier](cli-reference/datasource/add/fichier.md)
    * [Enterprise File Fabric](cli-reference/datasource/add/filefabric.md)
    * [FTP](cli-reference/datasource/add/ftp.md)
    * [Google Cloud Storage](cli-reference/datasource/add/gcs.md)
    * [Google Photos](cli-reference/datasource/add/gphotos.md)
    * [Hadoop 분산 파일 시스템](cli-reference/datasource/add/hdfs.md)
    * [HiDrive](cli-reference/datasource/add/hidrive.md)
    * [HTTP](cli-reference/datasource/add/http.md)
    * [Internet Archive](cli-reference/datasource/add/internetarchive.md)
    * [Jottacloud](cli-reference/datasource/add/jottacloud.md)
    * [Koofr / Digi Storage](cli-reference/datasource/add/koofr.md)
    * [로컬 디스크](cli-reference/datasource/add/local.md)
    * [Mail.ru Cloud](cli-reference/datasource/add/mailru.md)
    * [Mega](cli-reference/datasource/add/mega.md)
    * [Akamai NetStorage](cli-reference/datasource/add/netstorage.md)
    * [Microsoft OneDrive](cli-reference/datasource/add/onedrive.md)
    * [OpenDrive](cli-reference/datasource/add/opendrive.md)
    * [Oracle Cloud Infrastructure Object Storage](cli-reference/datasource/add/oos.md)
    * [Pcloud](cli-reference/datasource/add/pcloud.md)
    * [premiumize.me](cli-reference/datasource/add/premiumizeme.md)
    * [Put.io](cli-reference/datasource/add/putio.md)
    * [QingCloud Object Storage](cli-reference/datasource/add/qingstor.md)
    * [AWS S3 및 호환](cli-reference/datasource/add/s3.md)
    * [Seafile](cli-reference/datasource/add/seafile.md)
    * [SSH/SFTP](cli-reference/datasource/add/sftp.md)
    * [Citrix Sharefile](cli-reference/datasource/add/sharefile.md)
    * [Sia 분산 클라우드](cli-reference/datasource/add/sia.md)
    * [SMB / CIFS](cli-reference/datasource/add/smb.md)
    * [Storj 분산 클라우드 스토리지](cli-reference/datasource/add/storj.md)
    * [Sugarsync](cli-reference/datasource/add/sugarsync.md)
    * [OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)](cli-reference/datasource/add/swift.md)
    * [Uptobox](cli-reference/datasource/add/uptobox.md)
    * [WebDAV](cli-reference/datasource/add/webdav.md)
    * [Yandex Disk](cli-reference/datasource/add/yandex.md)
    * [Zoho](cli-reference/datasource/add/zoho.md)
  * [목록](cli-reference/datasource/list.md)
  * [상태](cli-reference/datasource/status.md)
  * [제거](cli-reference/datasource/remove.md)
  * [확인](cli-reference/datasource/check.md)
  * [업데이트](cli-reference/datasource/update.md)
  * [다시 스캔](cli-reference/datasource/rescan.md)
  * [DAG 생성기](cli-reference/datasource/daggen.md)
  * [검사](cli-reference/datasource/inspect/README.md)
    * [조각들](cli-reference/datasource/inspect/chunks.md)
    * [아이템들](cli-reference/datasource/inspect/items.md)
    * [DAG들](cli-reference/datasource/inspect/dags.md)
    * [조각 세부정보](cli-reference/datasource/inspect/chunkdetail.md)
    * [아이템 세부정보](cli-reference/datasource/inspect/itemdetail.md)
    * [경로](cli-reference/datasource/inspect/path.md)
* [월렛](cli-reference/wallet/README.md)
  * [가져오기](cli-reference/wallet/import.md)
  * [목록](cli-reference/wallet/list.md)
  * [원격 추가](cli-reference/wallet/add-remote.md)
  * [제거](cli-reference/wallet/remove.md)
* [도구](cli-reference/tool/README.md)
  * [CAR 추출](cli-reference/tool/extract-car.md)

<!-- cli end -->

## 🌐 웹 API 참조 <a href="#web-api-reference" id="web-api-reference"></a>
<!-- webapi begin -->

* [관리자](web-api-reference/admin.md)
* [데이터 소스](web-api-reference/data-source.md)
* [데이터셋](web-api-reference/dataset.md)
* [거래 일정](web-api-reference/deal-schedule.md)
* [거래](web-api-reference/deal.md)
* [메타데이터](web-api-reference/metadata.md)
* [월렛 연결](web-api-reference/wallet-association.md)
* [월렛](web-api-reference/wallet.md)
* [명세](https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docs/swagger/swagger.yaml)

<!-- webapi end -->

## 개발자용 기술 아키텍처 <a href="#technical-architecture" id="technical-architecture"></a>

* [데이터 준비](technical-architecture/data-preparation.md)

## ❓ 자주 묻는 질문 <a href="#faq" id="faq"></a>

* [데이터베이스 잠김 오류](faq/database-is-locked.md)
* [파일 삭제](faq/file-deletion.md)
