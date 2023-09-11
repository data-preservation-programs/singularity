# Google Drive

{% code fullWidth="true" %}
```
이름:
   singularity storage create drive - Google Drive

사용법:
   singularity storage create drive [command options] [arguments...]

설명:
   --client-id
      Google 애플리케이션 클라이언트 ID
      고유의 클라이언트 ID를 설정하는 것이 좋습니다.
      자체 클라이언트 ID 생성 방법은 https://rclone.org/drive/#making-your-own-client-id에서 확인하세요.
      비워두면 성능이 낮은 내부 키를 사용합니다.

   --client-secret
      OAuth 클라이언트 비밀.

      보통 비워둡니다.

   --token
      JSON blob 형식의 OAuth 액세스 토큰.

   --auth-url
      인증 서버 URL.
      
      공급자의 기본값을 사용하려면 비워둡니다.

   --token-url
      토큰 서버 URL.
      
      공급자의 기본값을 사용하려면 비워둡니다.

   --scope
      드라이브에 액세스할 때 rclone이 사용해야 하는 범위.

      예:
         | drive                   | 모든 파일에 대한 전체 액세스, 어플리케이션 데이터 폴더를 제외합니다.
         | drive.readonly          | 파일 메타데이터 및 파일 내용에 대한 읽기 전용 액세스.
         | drive.file              | rclone에 의해 생성된 파일에 대한 액세스만 가능합니다.
         |                         | 이러한 파일은 드라이브 웹 사이트에서 볼 수 있습니다.
         |                         | 사용자가 앱의 권한을 취소하면 파일 권한이 취소됩니다.
         | drive.appfolder         | 응용 프로그램 데이터 폴더에 대한 읽기 및 쓰기 액세스가 가능합니다.
         |                         | 드라이브 웹 사이트에서 보이지 않습니다.
         | drive.metadata.readonly | 파일 메타데이터에 대한 읽기 전용 액세스입니다.
         |                         | 파일 내용에 대한 액세스 또는 다운로드가 불가능합니다.

   --root-folder-id
      루트 폴더의 ID.
      보통 비워둡니다.
      
      "컴퓨터" 폴더에 액세스하려면 채워넣으세요 (문서 참조) 또는 rclone이
      시작점으로 사용할 수 없는 루트 폴더를 사용하기 위해.

   --service-account-file
      서비스 계정 자격 증명 JSON 파일 경로.
      
      외부 로그인 대신 SA를 사용하려면 비워두세요.
      
      `~`는 파일 이름에서 확장되며 `${RCLONE_CONFIG_DIR}`와 같은 환경 변수도 확장됩니다.

   --service-account-credentials
      서비스 계정 자격 증명 JSON blob.
      
      보통 비워둡니다.
      대화형 로그인 대신 SA를 사용하려면 필요합니다.

   --team-drive
      공유 드라이브 (팀 드라이브)의 ID.

   --auth-owner-only
      인증된 사용자에게 소유권이 있는 파일만 고려합니다.

   --use-trash
      파일을 영구적으로 삭제하는 대신 휴지통에 보냅니다.
      
      기본적으로 파일을 휴지통에 보냅니다.
      대신 파일을 영구적으로 삭제하려면 `--drive-use-trash=false`를 사용하세요.

   --copy-shortcut-content
      서버 측에서 바로 가기의 내용을 복사합니다.
      
      서버 측 복사를 수행할 때 rclone은 보통 바로 가기를 바로 가기로 복사합니다.
      
      이 플래그를 사용하면 서버 측 복사를 수행할 때 바로 가기의 내용을 복사합니다.

   --skip-gdocs
      모든 목록에서 Google 문서를 건너뜁니다.
      
      지정된 경우 gdocs은 rclone에서 사실상 보이지 않습니다.

   --skip-checksum-gphotos
      Google 포토 및 비디오의 MD5 체크섬을 건너뜁니다.
      
      Google 포토 또는 비디오를 전송할 때 체크섬 오류가 발생하는 경우에 사용하세요.
      
      이 플래그를 설정하면 Google 포토 및 비디오는 빈 MD5 체크섬을 반환합니다.
      
      Google 포토는 "photos" 공간에 위치한 파일입니다.
      
      손상된 체크섬은 Google이 이미지/비디오를 수정하지만 체크섬을 업데이트하지
      않기 때문에 발생합니다.

   --shared-with-me
      공유된 파일만 표시합니다.
      
      rclone이 "공유된 내 파일" 폴더에서 작업하도록 지시합니다.
      Google 드라이브에서 다른 사람이 공유한 파일과 폴더에 액세스할 수 있습니다.
      
      이는 "list" (lsd, lsl 등) 및 "copy" (copy, sync 등) 명령뿐만 아니라 다른
      모든 명령에도 적용됩니다.

   --trashed-only
      휴지통에 있는 파일만 표시합니다.
      
      원래 디렉토리 구조에서 휴지통에 있는 파일을 표시합니다.

   --starred-only
      별표가 표시된 파일만 표시합니다.

   --formats
      삭제됨: export_formats를 참조하세요.

   --export-formats
      Google 문서 다운로드에 사용할 체계 분리된 형식의 콤마로 구분된 목록.

   --import-formats
      Google 문서 업로드에 사용할 체계 분리된 형식의 콤마로 구분된 목록.

   --allow-import-name-change
      Google 문서 업로드시 파일 유형이 변경되면 허용합니다.
      
      예: file.doc이 file.docx로 변경됩니다. 이렇게 하면 변경 내용이 매번 동기화되고
      다시 업로드됩니다.

   --use-created-date
      파일 생성일자 대신 수정일자를 사용합니다.
      
      데이터를 다운로드하고 생성일자 대신 마지막 수정일자를 사용하려는 경우 유용합니다.
      
      **경고**: 이 플래그에는 예상치 못한 결과가 발생할 수 있습니다.
      
      드라이브에 업로드할 때 파일이 수정되지 않았을 경우 파일이 모두 덮어쓰여집니다.
      그리고 다운로드하는 동안 그 반대가 일어납니다. 이 부작용은 "--checksum" 플래그를 사용하여 피할 수 있습니다.
      
      이 기능은 구글 포토에서 기록된 사진 캡처 날짜를 유지하기 위해 구현되었습니다.
      구글 드라이브 설정에서 "Google 포토 폴더 만들기" 옵션을 먼저 확인해야 합니다.
      그런 다음 로컬로 사진을 복사하거나 옮길 수 있으며 이미지 캡처 날짜 또는 생성 날짜로 설정할 수 있습니다.

   --use-shared-date
      파일이 공유된 날짜 대신 수정일자를 사용합니다.
      
      "--drive-use-created-date" 플래그와 마찬가지로 이 플래그도 파일 업로드/다운로드에
      예상치 못한 결과를 낼 수 있습니다.
      
      이 플래그와 "--drive-use-created-date" 플래그 둘 다 설정된 경우, 생성일자가 사용됩니다.

   --list-chunk
      목록 청크의 크기 100-1000, 비활성화하려면 0으로 설정합니다.

   --impersonate
      서비스 계정을 사용할 때 이 사용자를 위해 위장합니다.

   --alternate-export
      삭제됨: 더 이상 필요하지 않습니다.

   --upload-cutoff
      청크 업로드로 전환되는 파일 크기 제한.

   --chunk-size
      업로드 청크 크기.
      
      256k 이상인 2의 승수이어야 합니다.
      
      이 값을 크게 설정하면 성능이 향상되지만, 각 청크가 한 번에 하나씩 전송되므로
      메모리에 버퍼링됩니다.
      
      이 값을 줄이면 메모리 사용량이 줄어들지만 성능이 감소합니다.

   --acknowledge-abuse
      cannotDownloadAbusiveFile을 반환하는 파일을 다운로드할 수 있도록 설정합니다.
      
      "This file has been identified as malware or spam and cannot be downloaded"
      에러 메시지와 함께 파일을 다운로드하려는 경우 에러 코드 "cannotDownloadAbusiveFile"이
      반환됩니다. 이 플래그를 rclone에 제공하여 파일을 다운로드할 위험성을 인식하고 rclone이
      그래도 파일을 다운로드하도록 표시할 수 있습니다.
      
      서비스 계정을 사용하는 경우 이 플래그가 작동하려면 관리자 권한이 필요합니다
      (컨텐츠 관리자가 아님). SA에 올바른 권한이 없으면 Google은 이 플래그를 무시할 것입니다.

   --keep-revision-forever
      각 파일의 새 대표 리비전을 영구적으로 보존합니다.

   --size-as-quota
      실제 크기가 아닌 스토리지 할당량 사용으로 크기를 표시합니다.
      
      파일의 크기를 스토리지 할당량으로 사용된 크기로 표시합니다. 이는
      현재 버전과 영구히 유지된 이전 버전을 포함합니다.
      
      **경고**: 이 플래그에는 예상치 못한 결과가 발생할 수 있습니다.
      
      구성에서이 플래그를 설정하는 것은 권장되지 않습니다. 권장되는 사용 방법은 rclone
      ls/lsl/lsf/lsjson 등을 수행할 때 --drive-size-as-quota 플래그를 사용하는 것입니다.
      
      동기화에이 플래그를 사용하는 경우 (권장하지 않음) --ignore size도 사용해야 합니다.

   --v2-download-min-size
      Object가 크면 drive v2 API를 사용하여 다운로드합니다.

   --pacer-min-sleep
      API 호출 사이에 대기해야하는 최소 시간.

   --pacer-burst
      대기하지 않고 허용되는 API 호출 횟수.

   --server-side-across-configs
      서버 측 작업 (예: 복사)을 다른 드라이브 구성 간에 사용할 수 있도록 허용합니다.
      
      서로 다른 두 Google 드라이브 간에 서버 측 복사를 수행하려는 경우 이 기능이 유용할 수 있습니다.
      모든 구성 간에 작동할지 여부를 쉽게 확인할 수 없기 때문에 기본적으로 이 기능은 활성화되어 있지 않습니다.

   --disable-http2
      drive에 http2 사용 비활성화.
      
      현재 google 드라이브 백엔드와 HTTP/2에 대한 해결되지 않은 문제가 있습니다.
      따라서 드라이브 백엔드에 대한 HTTP/2는 기본적으로 비활성화되어 있지만 여기서 다시 활성화할 수 있습니다.
      해결되었을 때이 플래그가 제거될 것입니다.
      
      참조: https://github.com/rclone/rclone/issues/3631

   --stop-on-upload-limit
      업로드 제한 오류를 치명적으로 처리합니다.
      
      작성 시점에서 하루에 Google 드라이브에 750 GiB의 데이터를 업로드하는 것이
      가능한 것으로 알려져 있습니다 (이는 기록되지 않은 제한입니다). 이 제한에
      도달하면 Google 드라이브에서 약간 다른 오류 메시지가 생성됩니다.
      이 플래그가 설정되면 이러한 오류가 치명적으로 처리됩니다. 업로드 작업이 중단됩니다.
      
      이 오류 탐지는 Google이 문서화하지 않은 오류 메시지 문자열을 기반으로 하므로
      향후 변경될 수 있습니다.
      
      참조: https://github.com/rclone/rclone/issues/3857

   --stop-on-download-limit
      다운로드 제한 오류를 치명적으로 처리합니다.
      
      작성 시점에서 Google 드라이브에서 하루에 10 TiB의 데이터를 다운로드하는 것이
      가능한 것으로 알려져 있습니다 (이는 기록되지 않은 제한입니다). 이 제한에
      도달하면 Google 드라이브에서 약간 다른 오류 메시지가 생성됩니다.
      이 플래그가 설정되면 이러한 오류가 치명적으로 처리됩니다. 다운로드 작업이 중단됩니다.
      
      이 오류 탐지는 Google이 문서화하지 않은 오류 메시지 문자열을 기반으로 하므로
      향후 변경될 수 있습니다.

   --skip-shortcuts
      설정된 경우 바로 가기 파일은 건너뜁니다.
      
      일반적으로 rclone은 바로 가기 파일을 해소하여 원본 파일처럼 보이게 만듭니다
      ([바로 가기 섹션](#shortcuts) 참조). 이 플래그가 설정되면 rclone은 바로 가기 파일을
      완전히 무시합니다.

   --skip-dangling-shortcuts
      설정된 경우 매달린 바로 가기 파일을 건너뜁니다.
      
      설정된 경우 rclone은 목록에서 매달린 바로 가기를 표시하지 않습니다.

   --resource-key
      링크 공유 파일에 액세스하기 위한 리소스 키.
      
      다음과 같이 링크로 공유된 파일에 액세스해야 하는 경우
      
          https://drive.google.com/drive/folders/XXX?resourcekey=YYY&usp=sharing
      
      첫 번째 부분 "XXX"를 "root_folder_id"로 사용하고 두 번째 부분 "YYY"를
      "resource_key"로 사용해야 합니다. 그렇지 않으면 디렉토리에 액세스하려고
      할 때 404 파일을 찾을 수 없음 오류가 발생합니다.
      
      참조: https://developers.google.com/drive/api/guides/resource-keys
      
      이 리소스 키 요구 사항은 일부 구식 파일에만 적용됩니다.
      
      또한 인증된 사용자에게서 한 번 웹 인터페이스에서 폴더를 열면 리소스 키가
      필요하지 않을 것 같습니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --alternate-export            삭제됨: 더 이상 필요하지 않음. (기본값: false) [$ALTERNATE_EXPORT]
   --client-id value             Google 애플리케이션 클라이언트 ID [$CLIENT_ID]
   --client-secret value         OAuth 클라이언트 비밀. [$CLIENT_SECRET]
   --help, -h                    도움말 보기
   --scope value                 드라이브에 액세스할 때 rclone이 사용해야 하는 범위. [$SCOPE]
   --service-account-file value  서비스 계정 자격 증명 JSON 파일 경로. [$SERVICE_ACCOUNT_FILE]

고급

   --acknowledge-abuse                  cannotDownloadAbusiveFile을 반환하는 파일을 다운로드할 수 있도록 설정합니다. (기본값: false) [$ACKNOWLEDGE_ABUSE]
   --allow-import-name-change           Google 문서 업로드시 파일 유형이 변경되면 허용합니다. (기본값: false) [$ALLOW_IMPORT_NAME_CHANGE]
   --auth-owner-only                    인증된 사용자에게 소유권이 있는 파일만 고려합니다. (기본값: false) [$AUTH_OWNER_ONLY]
   --auth-url value                     인증 서버 URL. [$AUTH_URL]
   --chunk-size value                   업로드 청크 크기. (기본값: "8Mi") [$CHUNK_SIZE]
   --copy-shortcut-content              서버 측에서 바로 가기의 내용을 복사합니다. (기본값: false) [$COPY_SHORTCUT_CONTENT]
   --disable-http2                      drive에 http2 사용 비활성화. (기본값: true) [$DISABLE_HTTP2]
   --encoding value                     백엔드의 인코딩. (기본값: "InvalidUtf8") [$ENCODING]
   --export-formats value               Google 문서 다운로드에 사용할 체계 분리된 형식의 콤마로 구분된 목록. (기본값: "docx,xlsx,pptx,svg") [$EXPORT_FORMATS]
   --formats value                      삭제됨: export_formats. [$FORMATS]
   --impersonate value                  서비스 계정을 사용할 때 이 사용자를 위해 위장합니다. [$IMPERSONATE]
   --import-formats value               Google 문서 업로드에 사용할 체계 분리된 형식의 콤마로 구분된 목록. [$IMPORT_FORMATS]
   --keep-revision-forever              각 파일의 새 대표 리비전을 영구히 보존합니다. (기본값: false) [$KEEP_REVISION_FOREVER]
   --list-chunk value                   목록 청크의 크기 100-1000, 비활성화하려면 0으로 설정합니다. (기본값: 1000) [$LIST_CHUNK]
   --pacer-burst value                  대기하지 않고 허용되는 API 호출 횟수. (기본값: 100) [$PACER_BURST]
   --pacer-min-sleep value              API 호출 사이에 대기해야하는 최소 시간. (기본값: "100ms") [$PACER_MIN_SLEEP]
   --resource-key value                 링크 공유 파일에 액세스하기 위한 리소스 키. [$RESOURCE_KEY]
   --root-folder-id value               루트 폴더의 ID. [$ROOT_FOLDER_ID]
   --server-side-across-configs         서버 측 작업 (예: 복사)을 다른 drive 구성 간에 수행할 수 있도록 허용합니다. (기본값: false) [$SERVER_SIDE_ACROSS_CONFIGS]
   --service-account-credentials value  서비스 계정 자격 증명 JSON blob. [$SERVICE_ACCOUNT_CREDENTIALS]
   --shared-with-me                     공유된 파일만 표시합니다. (기본값: false) [$SHARED_WITH_ME]
   --size-as-quota                      실제 크기가 아닌 스토리지 할당량 사용으로 크기를 표시합니다. (기본값: false) [$SIZE_AS_QUOTA]
   --skip-checksum-gphotos              Google 포토 및 비디오의 MD5 체크섬을 건너뜁니다. (기본값: false) [$SKIP_CHECKSUM_GPHOTOS]
   --skip-dangling-shortcuts            설정된 경우 매달린 바로 가기 파일을 건너뜁니다. (기본값: false) [$SKIP_DANGLING_SHORTCUTS]
   --skip-gdocs                         모든 목록에서 Google 문서를 건너뜁니다. (기본값: false) [$SKIP_GDOCS]
   --skip-shortcuts                     설정된 경우 바로 가기 파일은 건너뜁니다. (기본값: false) [$SKIP_SHORTCUTS]
   --starred-only                       별표가 표시된 파일만 표시합니다. (기본값: false) [$STARRED_ONLY]
   --stop-on-download-limit             다운로드 제한 오류를 치명적으로 처리합니다. (기본값: false) [$STOP_ON_DOWNLOAD_LIMIT]
   --stop-on-upload-limit               업로드 제한 오류를 치명적으로 처리합니다. (기본값: false) [$STOP_ON_UPLOAD_LIMIT]
   --team-drive value                   공유 드라이브 (팀 드라이브)의 ID. [$TEAM_DRIVE]
   --token value                        JSON blob 형식의 OAuth 액세스 토큰. [$TOKEN]
   --token-url value                    토큰 서버 URL. [$TOKEN_URL]
   --trashed-only                       휴지통에 있는 파일만 표시합니다. (기본값: false) [$TRASHED_ONLY]
   --upload-cutoff value                청크 업로드로 전환되는 파일 크기 제한. (기본값: "8Mi") [$UPLOAD_CUTOFF]
   --use-created-date                   파일 생성일자 대신 수정일자를 사용합니다. (기본값: false) [$USE_CREATED_DATE]
   --use-shared-date                    파일이 공유된 날짜 대신 수정일자를 사용합니다. (기본값: false) [$USE_SHARED_DATE]
   --use-trash                          파일을 영구적으로 삭제하는 대신 휴지통에 보냅니다. (기본값: true) [$USE_TRASH]
   --v2-download-min-size value         Object가 크면 drive v2 API를 사용하여 다운로드합니다. (기본값: "off") [$V2_DOWNLOAD_MIN_SIZE]

   일반

   --name value  스토리지 이름 (기본값: 자동 생성)
   --path value  스토리지 경로

```
{% endcode %}