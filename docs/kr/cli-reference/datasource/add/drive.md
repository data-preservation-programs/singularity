# Google Drive

```
이름:
    singularity datasource add drive - Google Drive

사용법:
   singularity datasource add drive [command options] <dataset_name> <source_path>

설명:
   --drive-acknowledge-abuse
      이 플래그를 설정하면 오해 가능한 파일을 다운로드 할 수 있게됩니다. 

      다운로드 중인 파일이 "This file has been identified 
      as malware or spam and cannot be downloaded" 오류와 
      함께 오류 코드 "cannotDownloadAbusiveFile"을 반환하는 경우,
      이 플래그를 제공하여 파일을 다운로드하고 리스크에 대한 인식 여부를 
      rclone에 알리게됩니다.

      이 플래그를 사용하려면 서비스 계정이 Manager 권한
      (Content Manager가 아님)이 필요합니다. SA에 올바른 권한이 없는 경우,
      Google은 해당 플래그를 무시할 것입니다.

   --drive-allow-import-name-change
      Google 문서 업로드시 파일 유형이 변경되는 것을 허용합니다.

      예: file.doc에서 file.docx로 변경됩니다. 이렇게되면 
      동기화 및 재업로드가 계속 발생합니다.

   --drive-alternate-export
      Deprecated: 더 이상 필요하지 않음.

   --drive-auth-owner-only
      인증 된 사용자가 소유한 파일만 고려합니다.

   --drive-auth-url
      인증 서버 URL.

      기본값을 사용하려면 비워두십시오.

   --drive-chunk-size
      업로드 청크 크기.

      256k 이상의 2의 거듭제곱이어야합니다.

      이 값을 크게하면 성능이 향상되지만, 각 chunk는 메모리에 
      한 번의 전송당 하나씩 버퍼링됩니다.

      이 값을 줄이면 메모리 사용량은 줄어들지만 성능이 감소합니다.

   --drive-client-id
      Google Application Client Id
      별도의 ID를 설정하는 것이 좋습니다.
      ID 생성에 대한 자세한 내용은
      [여기](https://rclone.org/drive/#making-your-own-client-id)를 참조하십시오.
      비워두면 성능이 낮은 내부 키를 사용합니다.

   --drive-client-secret
      OAuth 클라이언트 비밀.

      일반적으로 비워둡니다.

   --drive-copy-shortcut-content
      서버측에서 바로 가기의 내용을 복사합니다.

      서버 측 복사를 수행할 때 일반적으로 rclone은 바로 가기를
      바로 가기로 복사합니다.

      이 플래그를 사용하면 rclone은 서버측 복사를 수행할 때
      바로 가기 자체가 아니라 바로 가기의 내용을 복사합니다.

   --drive-disable-http2
      http2를 사용하지 않도록 drive를 비활성화합니다.

      현재 google drive 백엔드와 HTTP / 2에 문제가 있습니다.
      이로 인해 기본적으로 드라이브 백엔드에서 HTTP / 2가 사용되지 않지만,
      여기에서 다시 활성화 할 수 있습니다.
      문제가 해결되면이 플래그가 제거됩니다.

      참조: [https://github.com/rclone/rclone/issues/3631](https://github.com/rclone/rclone/issues/3631)

   --drive-encoding
      백엔드의 인코딩.

      자세한 내용은 [링크](/overview/#encoding)를 참조하십시오.

   --drive-export-formats
      Google 문서의 다운로드에 대한 우선적으로 사용할 수 있는 형식의 
      쉼표로 구분 된 목록.

   --drive-formats
      Deprecated: export_formats 대신 사용하십시오.

   --drive-impersonate
      서비스 계정을 사용할 때 해당 사용자를 사용합니다.

   --drive-import-formats
      Google 문서를 업로드하기 전에 우선적으로 사용할 수있는 형식의 
      쉼표로 구분 된 목록.

   --drive-keep-revision-forever
      각 파일의 새로운 헤드 리비전을 영구적으로 유지합니다.

   --drive-list-chunk
      목록 청크의 크기 100-1000, 0으로 비활성화합니다.

   --drive-pacer-burst
      대기 없이 허용되는 API 호출 수입니다.

   --drive-pacer-min-sleep
      API 호출 사이의 최소 대기 시간.

   --drive-resource-key
      링크로 공유 된 파일에 액세스하기 위한 리소스 키입니다.

      다음과 같은 링크로 공유 된 파일에 액세스해야하는 경우
      
          https://drive.google.com/drive/folders/XXX?resourcekey=YYY&usp=sharing
      
      "XXX"를 "root_folder_id"로 사용하고
      "YYY"를 "resource_key"로 사용해야합니다.
      그렇지 않으면 디렉토리에 액세스 할 때 404 not found 오류가 발생합니다.

      참조: [https://developers.google.com/drive/api/guides/resource-keys](https://developers.google.com/drive/api/guides/resource-keys)

      이 리소스 키 요구 사항은 오래된 파일의 일부에만 해당됩니다.

      또한 웹 인터페이스에서 폴더를 한 번 열면 (rclone으로 인증 한 
      사용자로) 리소스 키가 필요하지 않음으로 나타납니다.

   --drive-root-folder-id
      루트 폴더의 ID.
      일반적으로 비워둡니다.

      루트 폴더에 액세스하려면 "Computers" 폴더 (문서 참조)에 
      기입하거나 rclone이 비루트 폴더를 시작점으로 사용하도록 
      채웁니다.

   --drive-scope
      드라이브에 대한 액세스 요청시 rclone이 사용해야하는 범위.

      예:
          | drive                   | Application Data 폴더를 제외한 모든 파일에 대한 완전한 액세스.
          | drive.readonly          | 파일 메타 데이터 및 파일 콘텐츠에 대한 읽기 전용 액세스.
          | drive.file              | rclone에서 생성 된 파일에 액세스합니다.
                                    | 이 파일은 드라이브 웹 사이트에서 볼 수 있습니다.
                                    | 파일 권한은 사용자가 앱의 인증을 취소하면
                                    | 취소됩니다.
          | drive.appfolder         | Application Data 폴더에 대한 읽기 및 쓰기 액세스를 허용합니다.
                                    | 드라이브 웹 사이트에서는 볼 수 없습니다.
          | drive.metadata.readonly | 파일 메타 데이터에 대한 읽기 전용 액세스를 허용하지만,
                                    | 파일 콘텐츠에 대한 액세스는 허용하지 않습니다.

   --drive-server-side-across-configs
      (예 : 복사)를 서로 다른 드라이브 구성을 통해 작동하도록 허용합니다.

      이 기능은 동일한 Google 드라이브 간에 서버 측 복사를 수행하려는 경우 유용합니다.
      모든 두 구성간에 작동 할 것인지 쉽게 판단 할 수 없기 때문에
      기본적으로 사용되지 않습니다.

   --drive-service-account-credentials
      서비스 계정 자격 증명 JSON 블롭.

      일반적으로 비워둡니다.
      대화 형 로그인 대신 SA를 사용하려는 경우에만 필요합니다.

   --drive-service-account-file
      서비스 계정 자격 증명 JSON 파일 경로.
      
      일반적으로 비워둡니다.
      대화 형 로그인 대신 SA를 사용하려는 경우에만 필요합니다.
      
      파일 이름은 '~'처리 및 `${RCLONE_CONFIG_DIR}`와 같은 
      환경 변수가 확장됩니다.

   --drive-shared-with-me
      공유 된 파일 만 표시합니다.

      rclone에 공유 된 파일 및 폴더에 액세스 할 수있게합니다.

      "list" (lsd, lsl 등) 및 "copy" (복사, 동기화 등) 명령,
      그리고 다른 모든 명령과 함께 사용할 수 있습니다.

   --drive-size-as-quota
      파일 크기를 저장소 할당량 사용으로 표시, 실제 크기가 아닌.

      파일 크기를 사용된 저장소 할당량으로 표시합니다. 이는
      현재 버전과 영구적으로 유지되도록 설정된 이전 버전을 포함합니다.

      **경고**: 이 플래그에는 예상치 못한 일부 결과가 발생할 수 있습니다.

      구성에서이 플래그를 설정하는 것이 권장되지 않으므로
      rclone ls / lsl / lsf / lsjson 등을 수행 할 때 
      --drive-size-as-quota 플래그를 사용하는 것이 권장됩니다.

      동기화에이 플래그를 사용하는 경우 (--ignore size도 사용해야 함).

   --drive-skip-checksum-gphotos
      Google 사진 및 비디오에서 MD5 체크섬 건너뜀.

      Google 사진 또는 비디오를 전송 할 때 체크섬 오류가 발생 할 경우 
      이 플래그를 사용하십시오.

      이 플래그를 설정하면 Google 사진 및 비디오가 공백 MD5 체크섬을
      반환하도록합니다.

      Google 사진은 "photos" 공간에 있는 것으로 식별됩니다.

      손상된 체크섬은 Google이 이미지 / 비디오를 수정하지만
      체크섬을 업데이트하지 않기 때문입니다.

   --drive-skip-dangling-shortcuts
      이 플래그가 설정된 경우, 리스트에서 둥둥 투기되는
      바로 가기 파일을 표시하지 않습니다.

   --drive-skip-gdocs
      모든 목록에서 Google 문서 건너 뜁니다.

      지정된 경우, gdocs는 rclone에서 실제로 볼 수 없습니다.

   --drive-skip-shortcuts
      이 플래그가 설정된 경우 바로 가기 파일을 건너 뜁니다.

      일반적으로 rclone은 바로가기 파일을 열어서 원본 파일처럼
      표시합니다 (바로가기 섹션 참조).
      이 플래그가 설정된 경우 rclone은 바로 가기 파일을 무시합니다.

   --drive-starred-only
      스타가 지정된 파일만 표시합니다.

   --drive-stop-on-download-limit
      다운로드 한도 오류를 치명적인 오류로 처리합니다.

      현재 쓰기 시간에는 하루에 10 TiB의 데이터를 Google Drive에서
      다운로드 할 수 있습니다 (이는 문서화되지 않은 제한입니다).
     이 한도에 도달하면 Google Drive는 약간 다른 오류 메시지를 생성합니다.
     이 플래그가 설정되면 이러한 오류가 치명적인 오류로 발생합니다.
      이러한 오류는 진행중인 동기화를 중지합니다.

      Google은 문서화하지 않은 오류 메시지 문자열에 의존하는이 감지
      기능을 제공하므로 향후 변경 될 수 있음을 유의하세요.

   --drive-stop-on-upload-limit
      업로드 한도 오류를 치명적인 오류로 처리합니다.

      현재라이팅 시간에는 하루에 Google Drive에 750 GiB의 데이터를 
      업로드 할 수 있습니다 (이는 문서화되지 않은 제한입니다).
     이 한도에 도달하면 Google Drive는 약간 다른 오류 메시지를 생성합니다.
     이 플래그가 설정되면 이러한 오류가 치명적인 오류로 발생합니다.
     이 오류는 진행중인 동기화를 중지합니다.

      Google은 문서화하지 않은 오류 메시지 문자열에 의존하는이 감지
      기능을 제공하므로 향후 변경 될 수 있음을 유의하세요.

   --drive-team-drive
      공유 드라이브 (팀 드라이브)의 ID.

   --drive-token
      JSON 블롭 형식으로 OAuth 액세스 토큰입니다.

   --drive-token-url
      토큰 서버 URL.

      기본값을 사용하려면 비워두십시오.

   --drive-trashed-only
      휴지통에 있는 파일 만 표시합니다.

      이렇게하면 본래 디렉토리 구조에 삭제 된 파일이 표시됩니다.

   --drive-upload-cutoff
      청크 업로드로 전환하는 기준.

   --drive-use-created-date
      수정된 날짜 대신 파일 생성된 날짜를 사용합니다.

      데이터를 다운로드하고 작성된 날짜를 수정 날짜 대신 사용하려는 경우 유용합니다.

      **경고**:이 플래그에는 예상치 못한 일부 결과가 발생할 수 있습니다.

      드라이브에 업로드 할 때 모든 파일은 수정되지 않았으면 
      덮어 쓰기 됩니다. 그리고 반대로 다운로드 할 때도 
      동일하게 발생합니다. 이 부작용은 "--checksum" 플래그를 사용하여 
      피할 수 있습니다.

      이 기능은 Google 사진이 기록 한 사진 캡처 날짜를 
      유지하기 위해 구현되었습니다. Google 드라이브 설정에서
      "Google Photos 폴더 만들기" 옵션을 먼저 확인해야합니다. 
      그런 다음 로컬로 사진을 복사 또는 이동하고 
      이미지가 캡처 된 날짜 (생성 된 날짜)가 수정 날짜로 
      설정됩니다.

   --drive-use-shared-date
      수정된 날짜 대신 파일 공유 날짜를 사용합니다.

      "--drive-use-created-date"와 마찬가지로이 플래그는
      파일 업로드 / 다운로드시 예상치 못한 결과가 발생할 수 있습니다.

     이 플래그와 "--drive-use-created-date"가 동시에 설정되면 
      작성된 날짜가 사용됩니다.

   --drive-use-trash
      파일을 영구적으로 삭제하는 대신 휴지통으로 보냅니다.

      기본적으로 파일을 휴지통으로 보냅니다.
      파일을 영구적으로 삭제하려면 `--drive-use-trash=false`를 사용하세요.
      
   --drive-v2-download-min-size
      물체가 크면 drive v2 API를 사용하여 다운로드합니다.


옵션:
   --help, -h  도움말 보기

   데이터 준비 옵션

   --delete-after-export    [위험] CAR 파일로 내보낸 데이터 세트 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공한 스캔으로부터 이 간격이 지나면 자동으로 소스 디렉터리를 다시 스캔합니다. (기본값: 비활성)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비)

   드라이브 옵션

   --drive-acknowledge-abuse value            Set to allow files which return cannotDownloadAbusiveFile to be downloaded. (기본값: "false") [$DRIVE_ACKNOWLEDGE_ABUSE]
   --drive-allow-import-name-change value     Allow the filetype to change when uploading Google docs. (기본값: "false") [$DRIVE_ALLOW_IMPORT_NAME_CHANGE]
   --drive-auth-owner-only value              Only consider files owned by the authenticated user. (기본값: "false") [$DRIVE_AUTH_OWNER_ONLY]
   --drive-auth-url value                     Auth server URL. [$DRIVE_AUTH_URL]
   --drive-chunk-size value                   Upload chunk size. (기본값: "8Mi") [$DRIVE_CHUNK_SIZE]
   --drive-client-id value                    Google Application Client Id [$DRIVE_CLIENT_ID]
   --drive-client-secret value                OAuth Client Secret. [$DRIVE_CLIENT_SECRET]
   --drive-copy-shortcut-content value        Server side copy contents of shortcuts instead of the shortcut. (기본값: "false") [$DRIVE_COPY_SHORTCUT_CONTENT]
   --drive-disable-http2 value                Disable drive using http2. (기본값: "true") [$DRIVE_DISABLE_HTTP2]
   --drive-encoding value                     The encoding for the backend. (기본값: "InvalidUtf8") [$DRIVE_ENCODING]
   --drive-export-formats value               Comma separated list of preferred formats for downloading Google docs. (기본값: "docx,xlsx,pptx,svg") [$DRIVE_EXPORT_FORMATS]
   --drive-formats value                      Deprecated: See export_formats. [$DRIVE_FORMATS]
   --drive-impersonate value                  Impersonate this user when using a service account. [$DRIVE_IMPERSONATE]
   --drive-import-formats value               Comma separated list of preferred formats for uploading Google docs. [$DRIVE_IMPORT_FORMATS]
   --drive-keep-revision-forever value        Keep new head revision of each file forever. (기본값: "false") [$DRIVE_KEEP_REVISION_FOREVER]
   --drive-list-chunk value                   Size of listing chunk 100-1000, 0 to disable. (기본값: "1000") [$DRIVE_LIST_CHUNK]
   --drive-pacer-burst value                  Number of API calls to allow without sleeping. (기본값: "100") [$DRIVE_PACER_BURST]
   --drive-pacer-min-sleep value              Minimum time to sleep between API calls. (기본값: "100ms") [$DRIVE_PACER_MIN_SLEEP]
   --drive-resource-key value                 Resource key for accessing a link-shared file. [$DRIVE_RESOURCE_KEY]
   --drive-root-folder-id value               ID of the root folder. [$DRIVE_ROOT_FOLDER_ID]
   --drive-scope value                        Scope that rclone should use when requesting access from drive. [$DRIVE_SCOPE]
   --drive-server-side-across-configs value   Allow server-side operations (e.g. copy) to work across different drive configs. (기본값: "false") [$DRIVE_SERVER_SIDE_ACROSS_CONFIGS]
   --drive-service-account-credentials value  Service Account Credentials JSON blob. [$DRIVE_SERVICE_ACCOUNT_CREDENTIALS]
   --drive-service-account-file value         Service Account Credentials JSON file path. [$DRIVE_SERVICE_ACCOUNT_FILE]
   --drive-shared-with-me value               Only show files that are shared with me. (기본값: "false") [$DRIVE_SHARED_WITH_ME]
   --drive-size-as-quota value                Show sizes as storage quota usage, not actual size. (기본값: "false") [$DRIVE_SIZE_AS_QUOTA]
   --drive-skip-checksum-gphotos value        Skip MD5 checksum on Google photos and videos only. (기본값: "false") [$DRIVE_SKIP_CHECKSUM_GPHOTOS]
   --drive-skip-dangling-shortcuts value      If set skip dangling shortcut files. (기본값: "false") [$DRIVE_SKIP_DANGLING_SHORTCUTS]
   --drive-skip-gdocs value                   Skip google documents in all listings. (기본값: "false") [$DRIVE_SKIP_GDOCS]
   --drive-skip-shortcuts value               If set skip shortcut files. (기본값: "false") [$DRIVE_SKIP_SHORTCUTS]
   --drive-starred-only value                 Only show files that are starred. (기본값: "false") [$DRIVE_STARRED_ONLY]
   --drive-stop-on-download-limit value       Make download limit errors be fatal. (기본값: "false") [$DRIVE_STOP_ON_DOWNLOAD_LIMIT]
   --drive-stop-on-upload-limit value         Make upload limit errors be fatal. (기본값: "false") [$DRIVE_STOP_ON_UPLOAD_LIMIT]
   --drive-team-drive value                   ID of the Shared Drive (Team Drive). [$DRIVE_TEAM_DRIVE]
   --drive-token value                        OAuth Access Token as a JSON blob. [$DRIVE_TOKEN]
   --drive-token-url value                    Token server url. [$DRIVE_TOKEN_URL]
   --drive-trashed-only value                 Only show files that are in the trash. (기본값: "false") [$DRIVE_TRASHED_ONLY]
   --drive-upload-cutoff value                Cutoff for switching to chunked upload. (기본값: "8Mi") [$DRIVE_UPLOAD_CUTOFF]
   --drive-use-created-date value             Use file created date instead of modified date. (기본값: "false") [$DRIVE_USE_CREATED_DATE]
   --drive-use-shared-date value              Use date file was shared instead of modified date. (기본값: "false") [$DRIVE_USE_SHARED_DATE]
   --drive-use-trash value                    Send files to the trash instead of deleting permanently. (기본값: "true") [$DRIVE_USE_TRASH]
   --drive-v2-download-min-size value         If Object's are greater, use drive v2 API to download. (기본값: "off") [$DRIVE_V2_DOWNLOAD_MIN_SIZE]

```