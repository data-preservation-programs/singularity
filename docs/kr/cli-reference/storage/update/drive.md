# Google Drive

## singularity storage update drive - Google 드라이브

**사용법:**
```
singularity storage update drive [command options] <name|id>
```

**설명:**
- `--client-id`
    - 사용자 지정 Google 앱 클라이언트 ID
    - 고유한 ID 사용을 권장합니다.
    - [여기](https://rclone.org/drive/#making-your-own-client-id)를 참조하여 고유한 ID를 생성하는 방법을 확인하십시오.
    - 값이 비어 있으면 성능이 저하될 수 있는 내부 키를 사용합니다.

- `--client-secret`
    - OAuth 클라이언트 시크릿입니다.
    - 일반적으로 비워 둡니다.

- `--token`
    - JSON blob으로 된 OAuth 액세스 토큰입니다.

- `--auth-url`
    - 인증 서버 URL입니다.
    - 공급자 기본값을 사용하려면 비워 둡니다.

- `--token-url`
    - 토큰 서버 URL입니다.
    - 공급자 기본값을 사용하려면 비워 둡니다.

- `--scope`
    - rclone이 drive에 액세스할 때 요청할 스코프입니다.

    예시:
    - `drive`: 모든 파일에 대한 전체 액세스(프로그램 데이터 폴더 제외)
    - `drive.readonly`: 파일 메타데이터 및 내용에 대한 읽기 전용 액세스
    - `drive.file`: rclone에 의해 생성된 파일에 대한 액세스
    - `drive.appfolder`: 응용 프로그램 데이터 폴더에 대한 읽기 및 쓰기 액세스
    - `drive.metadata.readonly`: 파일 메타데이터에 대한 읽기 전용 액세스

- `--root-folder-id`
    - 루트 폴더의 ID
    - 보통 비워 두십시오.
    - "컴퓨터" 폴더에 액세스하기 위해 입력하거나 rclone이
      시작점으로 사용할 비루트 폴더를 사용하려면 채워 넣습니다.

- `--service-account-file`
    - 서비스 계정 자격 증명 JSON 파일 경로입니다.
    - 일반적으로 비워 둡니다.
    - 대화형 로그인 대신 SA를 사용하려는 경우에만 필요합니다.
    - 선행 '~'는 파일 이름에서 확장되며 `${RCLONE_CONFIG_DIR}`과 같은 환경 변수도 확장됩니다.

- `--service-account-credentials`
    - 서비스 계정 자격 증명 JSON blob입니다.
    - 일반적으로 비워 둡니다.
    - 대화형 로그인 대신 SA를 사용하려는 경우에만 필요합니다.

- `--team-drive`
    - 공유 드라이브(팀 드라이브)의 ID입니다.

- `--auth-owner-only`
    - 인증된 사용자의 파일만 고려합니다.

- `--use-trash`
    - 파일을 영구적으로 삭제하는 대신 휴지통으로 보냅니다.
    - 기본적으로 파일을 휴지통으로 보냅니다.
    - 파일을 영구적으로 삭제하려면 `--drive-use-trash=false`를 사용하십시오.

- `--copy-shortcut-content`
    - 서버측으로 바로 가기의 내용을 복사합니다.
    - 서버측 복사를 수행하는 경우 일반적으로 rclone은 바로 가기를 복사합니다.
    - 이 플래그를 사용하면 서버측 복사를 수행할 때 바로 가기를 복사하는 대신
      바로 가기의 내용을 복사합니다.

- `--skip-gdocs`
    - 모든 목록에서 Google 문서를 건너뜁니다.
    - 지정한 경우 rclone에서 gdocs는 실제로 보이지 않게 됩니다.

- `--skip-checksum-gphotos`
    - Google 사진 및 비디오의 MD5 체크섬을 건너뜁니다.
    - Google 사진이 "photos" 공간에 있는 경우 Google 사진과 비디오를 전송할 때
      체크섬 오류가 발생하는 경우에 사용하십시오.
    - 이 플래그를 설정하면 Google 사진과 비디오가
      빈 MD5 체크섬을 반환하도록 설정됩니다.
    - 손상된 체크섬은 Google이 이미지/비디오를 수정하지만
      체크섬을 업데이트하지 않아서 발생합니다.

- `--shared-with-me`
    - 나와 공유된 파일만 표시합니다.
    - rclone이 "Shared with me" 폴더에서 작동하도록 지시합니다.
      (Google 드라이브에서 다른 사람이 공유한 파일과 폴더에 액세스하는 곳)
    - "list" (lsd, lsl 등) 및 "copy" (copy, sync 등) 명령뿐만 아니라
      다른 모든 명령에서도 작동합니다.

- `--trashed-only`
    - 휴지통에 들어 있는 파일만 표시합니다.
    - 이렇게 하면 휴지통에 있는 파일이 원래 디렉터리 구조로 표시됩니다.

- `--starred-only`
    - 즐겨찾기한 파일만 표시합니다.

- `--formats`
    - Deprecated: export_formats를 참조하십시오.

- `--export-formats`
    - Google 문서를 다운로드할 때 선호하는 형식의 쉼표로 구분된 목록입니다.

- `--import-formats`
    - Google 문서를 업로드할 때 선호하는 형식의 쉼표로 구분된 목록입니다.

- `--allow-import-name-change`
    - Google 문서를 업로드할 때 파일 유형이 변경되는 것을 허용합니다.
    - 예: `file.doc`에서 `file.docx`로 변경됩니다.
    - 이렇게 하면 동기화가 혼란스러워지고 매번 다시 업로드됩니다.

- `--use-created-date`
    - 수정된 날짜 대신 파일 생성된 날짜를 사용합니다.
    - 데이터를 다운로드하고 생성 날짜 대신 마지막 수정 날짜를 사용하려는 경우 유용합니다.
    - **주의**: 이 플래그는 예기치 않은 결과가 발생할 수 있습니다.
    - 드라이브에 업로드할 때 모든 파일이 곧 덮어쓰이므로
      파일이 생성된 후 수정되지 않았으면 대체됩니다.
      다운로드하는 경우 반대로 작동합니다.
      이 부작용은 "--checksum" 플래그를 사용하여 피할 수 있습니다.
    - 이 기능은 Google 사진에 의해 기록된 사진 캡처 날짜를 보존하기 위해 구현되었습니다.
      Google 드라이브 설정에서 "Google 사진 폴더 생성" 옵션을 확인해야 합니다.
      그런 다음 이미지를 로컬로 복사하거나 이용하여 이미지의 캡처
      날짜(생성일)를 수정 날짜로 설정할 수 있습니다.

- `--use-shared-date`
    - 수정된 날짜 대신에 파일이 공유된 날짜를 사용합니다.
    - **주의**: "--drive-use-created-date"와 같이 사용할 경우 예기치 않은 결과가 발생할 수 있습니다.
    - 이 플래그와 "--drive-use-created-date" 모두 설정된 경우 생성된 날짜가 사용됩니다.

- `--list-chunk`
    - 목록 청크의 크기(100-1000, 0은 비활성화)입니다.

- `--impersonate`
    - 서비스 계정을 사용할 때 이 사용자를 표현합니다.

- `--alternate-export`
    - Deprecated: 더 이상 필요하지 않습니다.

- `--upload-cutoff`
    - 청크 업로드로 전환하는 데 사용되는 기준입니다.

- `--chunk-size`
    - 업로드 청크 크기입니다.
    - 256k 이상인 2의 거듭 제곱이어야 합니다.
    - 이 값을 크게 설정하면 성능이 향상되지만, 각 청크는 전송당 메모리에 하나씩 버퍼링됩니다.
    - 이 값을 줄이면 메모리 사용량이 감소하지만 성능이 감소합니다.

- `--acknowledge-abuse`
    - "잠재적 위험한 파일을 다운로드할 수 없습니다."라는 오류 메시지가 표시되는 파일을 다운로드할 수 있도록 설정합니다.
    - 만약 파일 다운로드 시 "This file has been identified as malware or spam and cannot be downloaded"의 오류 코드 "cannotDownloadAbusiveFile"과 함께 에러 메시지가 표시되면, 이 플래그를 rclone에 제공하여 파일을 다운로드할 위험을 인지한다는 것을 나타낼 수 있습니다.
    - 이 플래그를 작동시키려면 서비스 계정이 관리자 권한(콘텐츠 관리자가 아님)이 필요합니다. SA에 적절한 권한이 없으면 Google은 이 플래그를 무시합니다.

- `--keep-revision-forever`
    - 각 파일의 새 헤드 버전을 영구히 유지합니다.

- `--size-as-quota`
    - 실제 크기 대신 스토리지 할당량 사용량으로 파일 크기를 표시합니다.
    - 파일의 크기를 스토리지 할당량 사용량으로 표시합니다. 이는 현재 버전과 영구히 유지된 이전 버전을 합친 값입니다.
    - **주의**: 이 플래그는 예기치 않은 결과가 발생할 수 있습니다.
    - 구성에 이 플래그를 설정하는 것이 권장되지 않으므로,
      rclone ls/lsl/lsf/lsjson 등을 사용할 때 `--drive-size-as-quota` 플래그를 사용하는 것이 권장됩니다.
    - 동기화에 이 플래그를 사용하는 경우(권장되지 않음) `--ignore size`도 사용해야 합니다.

- `--v2-download-min-size`
    - 객체의 크기가 지정한 값보다 큰 경우 drive v2 API를 사용하여 다운로드합니다.

- `--pacer-min-sleep`
    - API 호출 사이의 최소 대기 시간입니다.

- `--pacer-burst`
    - 대기 없이 허용되는 API 호출 횟수입니다.

- `--server-side-across-configs`
    - 서버측 작업(예: 복사)을 서로 다른 드라이브 구성 간에 사용할 수 있도록 허용합니다.
    - 서로 다른 두 개의 Google 드라이브 간에 서버측 복사를 수행하려는 경우에 유용할 수 있습니다.
    - 모든 구성 사이에서 작동할 수 있는지 판별하기 어렵기 때문에 기본적으로 사용되지 않습니다.

- `--disable-http2`
    - 드라이브에서 http2 사용을 비활성화합니다.
    - 현재 Google 드라이브 백엔드와 HTTP/2 사이에 해결되지 않은 문제가 있습니다.
    - 따라서 기본적으로 드라이브 백엔드에서 HTTP/2가 비활성화되어 있지만, 여기에서 다시 활성화할 수 있습니다.
    - 문제가 해결되면 이 플래그는 제거될 것입니다.
    - 참조: https://github.com/rclone/rclone/issues/3631

- `--stop-on-upload-limit`
    - 업로드 제한 오류를 치명적인 오류로 설정합니다.
    - 현재 기록 시, Google 드라이브에 하루에 750 GiB의 데이터를 업로드할 수 있습니다(기록되지 않은 제한).
      이 한계에 도달하면 Google 드라이브는 약간 다른 오류 메시지를 생성합니다.
    - 이 플래그가 설정되면 이러한 오류가 치명적인 오류로 처리됩니다. 이로 인해 진행 중인 동기화가 중지됩니다.
    - Google은 문자열 형식의 오류 메시지에 의존하기 때문에 이 감지 방법은 Google이 문서화하지 않으므로 추후에 동작이 바뀔 수 있습니다.
    - 참조: https://github.com/rclone/rclone/issues/3857

- `--stop-on-download-limit`
    - 다운로드 제한 오류를 치명적인 오류로 설정합니다.
    - 현재 기록 시, Google 드라이브에서 하루에 10 TiB의 데이터를 다운로드할 수 있습니다(기록되지 않은 제한).
      이 한계에 도달하면 Google 드라이브는 약간 다른 오류 메시지를 생성합니다.
    - 이 플래그가 설정되면 이러한 오류가 치명적인 오류로 처리됩니다. 이로 인해 진행 중인 동기화가 중지됩니다.
    - Google은 문자열 형식의 오류 메시지에 의존하기 때문에 이 감지 방법은 Google이 문서화하지 않으므로 추후에 동작이 바뀔 수 있습니다.

- `--skip-shortcuts`
    - 바로 가기 파일을 건너뜁니다.
    - 일반적으로 rclone은 바로 가기 파일을 역참조하여 원본 파일처럼 보이게 합니다.
    - 이 플래그가 설정되면 rclone은 바로 가기 파일을 완전히 무시합니다.

- `--skip-dangling-shortcuts`
    - 끊어진 바로 가기 파일을 건너뜁니다.
    - 설정된 경우 rclone은 목록에 모든 끊어진 바로 가기를 표시하지 않습니다.

- `--resource-key`
    - 링크로 공유된 파일에 액세스하기 위한 리소스 키입니다.
    - 다음과 같이 링크로 공유된 파일에 액세스해야 하는 경우
      ```
      https://drive.google.com/drive/folders/XXX?resourcekey=YYY&usp=sharing
      ```
      "XXX"를 "root_folder_id"로 사용하고 "YYY"를 "resource_key"로 사용해야 합니다.
      그렇지 않으면 디렉토리에 액세스할 때 404 오류가 발생합니다.
    - 참조: https://developers.google.com/drive/api/guides/resource-keys
    - 이 리소스 키 요구 사항은 오래된 파일의 일부에만 적용됩니다.
    - 또한 인증된 rclone 사용자로 폴더를 한 번이라도 웹 인터페이스에서 열면 리소스 키를 더 이상 필요하지 않습니다.

- `--encoding`
    - 백엔드에 대한 인코딩입니다.
    - 자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.
