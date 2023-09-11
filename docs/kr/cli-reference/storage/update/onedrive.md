# Microsoft OneDrive

{% code fullWidth="true" %}
```
이름:
   singularity storage update onedrive - Microsoft OneDrive

사용법:
   singularity storage update onedrive [옵션] <이름|ID>

설명:
   --client-id
      OAuth 클라이언트 ID.
      
      보통 비워둡니다.

   --client-secret
      OAuth 클라이언트 비밀번호.
      
      보통 비워둡니다.

   --token
      OAuth 액세스 토큰을 JSON 블롭으로 입력합니다.

   --auth-url
      인증 서버 URL.
      
      공급자 기본값을 사용하려면 비워둡니다.

   --token-url
      토큰 서버 URL.
      
      공급자 기본값을 사용하려면 비워둡니다.

   --region
      OneDrive의 국가 클라우드 지역을 선택합니다.

      예시:
         | global | Microsoft Cloud Global
         | us     | Microsoft Cloud for US Government
         | de     | Microsoft Cloud Germany
         | cn     | Azure and Office 365 operated by Vnet Group in China

   --chunk-size
      파일을 업로드할 때 사용되는 청크 크기입니다. - 320k(327,680 바이트)의 배수여야 합니다.
      
      이 크기 이상의 파일은 청크로 분할됩니다. - 320k(327,680 바이트)의 배수여야 하며
      250M(262,144,000 바이트)을 초과하지 않아야 합니다. 그렇지 않으면 "Microsoft.SharePoint.Client.InvalidClientQueryException: The request message is too big." 에러가 발생할 수 있습니다.
      청크는 메모리에 버퍼링됩니다.

   --drive-id
      사용할 드라이브의 ID입니다.

   --drive-type
      드라이브 유형(personal | business | documentLibrary)입니다.

   --root-folder-id
      루트 폴더의 ID입니다.
      
      보통은 필요하지 않지만, 특수한 경우에는 접근하려는 폴더의 ID를 알고 있지만 경로 탐색으로 갈 수 없을 수 있습니다.
      

   --access-scopes
      rclone이 요청할 스코프를 설정합니다.
      
      rclone이 요청할 모든 스코프를 지정한 스페이스로 구분된 사용자 지정 리스트로 선택하거나 입력합니다.
      

      예시:
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access | 모든 자원에 대한 읽기 및 쓰기 액세스
         | Files.Read Files.Read.All Sites.Read.All offline_access                                     | 모든 자원에 대한 읽기 전용 액세스
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All offline_access                | 모든 자원에 대한 읽기 및 쓰기 액세스, SharePoint 사이트 브라우징 기능 없음
         |                                                                                             | disable_site_permission이 true로 설정된 것과 동일

   --disable-site-permission
      Sites.Read.All 권한 요청을 비활성화합니다.
      
      이 값을 true로 설정하면 드라이브 ID를 구성하는 동안 SharePoint 사이트를 검색할 수 없습니다.
      rclone은 Sites.Read.All 권한을 요청하지 않습니다.
      조직에서 앱에 Sites.Read.All 권한을 지정하지 않았고 사용자가 자체적으로 앱 권한 승인을 거부하는 경우, 조직에서 true로 설정하세요.

   --expose-onenote-files
      디렉토리 목록에 OneNote 파일이 표시되도록 설정합니다.
      
      기본적으로 rclone은 디렉토리 목록에서 OneNote 파일을 숨깁니다.
      "열기" 및 "업데이트"와 같은 작업이 작동하지 않기 때문입니다. 하지만 이
      동작은 OneNote 파일을 삭제하는 것을 방지할 수도 있습니다. OneNote 파일을
      삭제하거나 그 밖에 디렉토리 목록에 표시하려는 경우 이 옵션을 설정하세요.

   --server-side-across-configs
      서버 측 작업(예: 복사)을 다른 onedrive 구성에서 작동할 수 있도록 허용합니다.
      
      이 작업은 두 개의 OneDrive *Personal* 드라이브 사이에서 복사하는 경우에만 작동합니다. 또한 복사할 파일은 이미 공유되어 있어야 합니다.
      그 외의 경우, rclone은 약간 느린 일반 복사로 다시 되돌아갑니다.

   --list-chunk
      목록 청크의 크기입니다.

   --no-versions
      수정 작업에서 모든 버전을 제거합니다.
      
      Onedrive for business는 기존 파일을 덮어쓰거나 수정 시 새 파일을 업로드할 때 버전을 생성합니다.
      이러한 버전은 할당량에서 공간을 차지합니다.
      
      이 플래그는 파일 업로드 및 수정 시 버전을 확인하고
      마지막 버전을 제외한 모든 버전을 제거합니다.
      
      **주의** Onedrive 개인에서는 현재 버전을 삭제할 수 없으므로 이 플래그를 사용하지 마세요.
      

   --link-scope
      링크 명령으로 생성된 링크의 범위를 설정합니다.

      예시:
         | anonymous    | 링크를 통해 로그인하지 않고 링크에 액세스할 수 있는 모든 사용자.
         |              | 조직 외부 사용자를 포함할 수 있습니다.
         |              | 익명 링크 지원이 관리자에 의해 비활성화될 수 있습니다.
         | organization | 조직에 로그인한 사용자(테넌트)만 링크를 사용하여 액세스할 수 있습니다.
         |              | OneDrive for Business와 SharePoint에서만 사용 가능합니다.

   --link-type
      링크 명령으로 생성된 링크의 유형을 설정합니다.

      예시:
         | view  | 항목에 대한 읽기 전용 링크를 생성합니다.
         | edit  | 항목에 대한 읽기-쓰기 링크를 생성합니다.
         | embed | 항목에 대한 임베드 가능한 링크를 생성합니다.

   --link-password
      링크 명령으로 생성된 링크의 암호를 설정합니다.
      
      이 기능은 현재 OneDrive 개인 유료 계정에서만 작동합니다.
      

   --hash-type
      백엔드에서 사용되는 해시를 지정합니다.
      
      이 값을 설정하면 사용되는 해시 유형이 지정됩니다. "auto"로 설정하면
      QuickXorHash가 기본 해시로 사용됩니다.
      
      rclone 1.62 이전에는 Onedrive 개인에 대해 기본적으로 SHA1 해시가 사용되었습니다.
      1.62 및 이후 버전에서는 모든 onedrive 유형에 대해 QuickXorHash를 사용하는 것이 기본값입니다.
      SHA1 해시를 원하는 경우 이 옵션을 해당 값으로 설정하세요.
      
      2023년 7월부터 OneDrive for Business와 OneDrive 개인의 유일한 사용 가능한 해시는
      QuickXorHash가 될 것입니다.
      
      이 값을 "none"으로 설정하면 해시를 사용하지 않습니다.
      
      요청한 해시가 개체에 없는 경우, rclone은 빈 문자열로 반환되며
      rclone에서 누락된 해시로 처리합니다.
      

      예시:
         | auto     | Rclone이 최상의 해시를 선택합니다.
         | quickxor | QuickXor
         | sha1     | SHA1
         | sha256   | SHA256
         | crc32    | CRC32
         | none     | 사용하지 않음 - 어떤 해시도 사용하지 않습니다.

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --client-id value      OAuth 클라이언트 ID. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 비밀번호. [$CLIENT_SECRET]
   --help, -h             도움말 표시
   --region value         OneDrive의 국가 클라우드 지역을 선택합니다. (기본값: "global") [$REGION]

   고급

   --access-scopes value         rclone이 요청할 스코프를 설정합니다. (기본값: "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access") [$ACCESS_SCOPES]
   --auth-url value              인증 서버 URL. [$AUTH_URL]
   --chunk-size value            파일을 업로드할 때 사용되는 청크 크기입니다. (기본값: "10Mi") [$CHUNK_SIZE]
   --disable-site-permission     Sites.Read.All 권한 요청을 비활성화합니다. (기본값: false) [$DISABLE_SITE_PERMISSION]
   --drive-id value              사용할 드라이브의 ID입니다. [$DRIVE_ID]
   --drive-type value            드라이브 유형(personal | business | documentLibrary)입니다. [$DRIVE_TYPE]
   --encoding value              백엔드에 대한 인코딩입니다. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --expose-onenote-files        디렉토리 목록에 OneNote 파일이 표시되도록 설정합니다. (기본값: false) [$EXPOSE_ONENOTE_FILES]
   --hash-type value             백엔드에서 사용되는 해시를 지정합니다. (기본값: "auto") [$HASH_TYPE]
   --link-password value         링크 명령으로 생성된 링크의 암호를 설정합니다. [$LINK_PASSWORD]
   --link-scope value            링크 명령으로 생성된 링크의 범위를 설정합니다. (기본값: "anonymous") [$LINK_SCOPE]
   --link-type value             링크 명령으로 생성된 링크의 유형을 설정합니다. (기본값: "view") [$LINK_TYPE]
   --list-chunk value            목록 청크의 크기입니다. (기본값: 1000) [$LIST_CHUNK]
   --no-versions                 수정 작업에서 모든 버전을 제거합니다. (기본값: false) [$NO_VERSIONS]
   --root-folder-id value        루트 폴더의 ID입니다. [$ROOT_FOLDER_ID]
   --server-side-across-configs  서버 측 작업(예: 복사)을 다른 onedrive 구성에서 작동할 수 있도록 허용합니다. (기본값: false) [$SERVER_SIDE_ACROSS_CONFIGS]
   --token value                 OAuth 액세스 토큰을 JSON 블롭으로 입력합니다. [$TOKEN]
   --token-url value             토큰 서버 URL. [$TOKEN_URL]

```
{% endcode %}