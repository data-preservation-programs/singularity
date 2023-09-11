# Microsoft OneDrive

{% code fullWidth="true" %}
```
이름:
   싱귤래리티 스토리지 생성 원드라이브 - Microsoft OneDrive

사용법:
   singularity storage create onedrive [옵션] [인수...]

설명:
   --client-id
      OAuth 클라이언트 ID.
      
      보통 비워둡니다.

   --client-secret
      OAuth 클라이언트 비밀키.
      
      보통 비워둡니다.

   --token
      JSON blob 형식의 OAuth 액세스 토큰.

   --auth-url
      인증 서버 URL.
      
      공급자 기본값을 사용하려면 비워둡니다.

   --token-url
      토큰 서버 URL.
      
      공급자 기본값을 사용하려면 비워둡니다.

   --region
      원드라이브의 국가 클라우드 지역 선택.

      예:
         | global | 전 세계용 Microsoft 클라우드
         | us     | 미국 정부용 Microsoft 클라우드
         | de     | 독일용 Microsoft 클라우드
         | cn     | 중국의 Vnet Group에서 운영하는 Azure 및 Office 365

   --chunk-size
      파일 업로드에 사용할 청크 크기 - 320k(327,680 바이트)의 배수여야 함.
      
      이 크기 이상일 경우 파일이 청크로 분할됩니다. 청크 크기는 320k(327,680 바이트)의 배수이어야 하며,
      250M(262,144,000 바이트)를 초과하지 않아야 합니다. 그렇지 않으면 "Microsoft.SharePoint.Client.InvalidClientQueryException: The request message is too big." 오류가 발생할 수 있습니다.
      청크는 메모리에 버퍼링됩니다.

   --drive-id
      사용할 드라이브의 ID.

   --drive-type
      드라이브 유형(개인 | 비즈니스 | 문서 라이브러리).

   --root-folder-id
      루트 폴더의 ID.
      
      보통 이 옵션은 필요하지 않지만, 특정한 상황에서는 액세스하려는 폴더의 ID를 알고 있지만 경로 탐색을 통해 액세스할 수 없을 수도 있습니다.
      

   --access-scopes
      rclone이 요청할 스코프를 설정합니다.
      
      rclone이 요청해야 할 모든 스코프가 포함된 사용자 정의 공백으로 구분된 목록을 선택하거나 직접 입력합니다.
      

      예:
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access | 모든 리소스에 대한 읽기 및 쓰기 액세스
         | Files.Read Files.Read.All Sites.Read.All offline_access                                     | 모든 리소스에 대한 읽기 전용 액세스
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All offline_access                | 모든 리소스에 대한 읽기 및 쓰기 액세스(SharePoint 사이트를 탐색할 수 있는 기능 제외)
         |                                                                                             | disable_site_permission가 true로 설정된 것과 같습니다.

   --disable-site-permission
      Sites.Read.All 권한 요청 비활성화.
      
      true로 설정하면 SharePoint 사이트를 검색할 수 없게 되므로, 드라이브 ID를 구성하는 동안
      Sites.Read.All 권한을 rclone이 요청하지 않습니다. 조직에서 응용 프로그램에 Sites.Read.All 권한을 할당하지 않았거나,
      사용자가 스스로 앱 권한 요청에 동의할 수 없는 경우 true로 설정합니다.

   --expose-onenote-files
      디렉토리 리스트에 OneNote 파일이 표시되도록 설정합니다.
      
      기본적으로 rclone은 디렉토리 리스트에서 OneNote 파일을 숨깁니다.
      "열기" 및 "업데이트"와 같은 작업이 이러한 파일에서 작동하지 않기 때문입니다. 그러나 이
      동작은 이러한 파일을 삭제할 수 없게도 할 수 있습니다. OneNote 파일을 삭제하거나 기타 방식으로
      디렉토리 목록에 표시하려면 이 옵션을 설정합니다.

   --server-side-across-configs
      서버 측 작업(예: 복사)이 다른 onedrive 구성간에 작업할 수 있도록 합니다.
      
      이것은 두 OneDrive의 *개인용* 드라이브 사이에서 복사하는 경우에만 작동합니다.
      복사할 파일이 이미 이들 간에 공유되어 있는 경우에만 작동합니다. 다른 경우에는
      rclone은 기본 복사로 이행합니다(약간 느릴 수 있습니다).

   --list-chunk
      목록 청크 크기.

   --no-versions
      수정 작업 시 모든 버전 삭제.
      
      비즈니스용 OneDrive는 새 파일을 업로드하거나 수정 시 버전을 생성합니다.
      수정 시 생성되는 버전은 할당량 중에서 공간을 차지합니다.
      
      이 플래그는 파일 업로드 및 수정 시 버전을 확인하고
      마지막 버전을 제외한 모든 버전을 삭제합니다.
      
      **주의** 개인용 OneDrive에서는 현재 버전을 삭제할 수 없으므로 이 플래그를 사용하지 마십시오.
      

   --link-scope
      링크 명령에 의해 생성된 링크의 범위 설정.

      예:
         | anonymous    | 링크를 보유한 모든 사람이 로그인하지 않고도 액세스할 수 있습니다.
         |              | 이에는 조직 외부의 사람들도 포함될 수 있습니다.
         |              | 익명 링크 지원은 관리자에 의해 비활성화될 수 있습니다.
         | organization | 조직에 로그인한 사람(테넌트)만 링크를 사용하여 액세스할 수 있습니다.
         |              | OneDrive for Business 및 SharePoint에서만 사용할 수 있습니다.

   --link-type
      링크 명령에 의해 생성된 링크의 유형 설정.

      예:
         | view  | 항목에 대한 읽기 전용 링크를 생성합니다.
         | edit  | 항목에 대한 읽기/쓰기 링크를 생성합니다.
         | embed | 항목에 대한 임베드 가능한 링크를 생성합니다.

   --link-password
      링크 명령에 의해 생성된 링크의 비밀번호 설정.
      
      작성 시점에서 이는 OneDrive 개인용 유료 계정에서만 작동합니다.
      

   --hash-type
      백엔드에서 사용 중인 해시 지정.
      
      이는 사용 중인 해시 유형을 지정합니다. "auto"로 설정하면
      기본 해시인 QuickXorHash가 사용됩니다.
      
      rclone 1.62 이전에는 OneDrive Personal의 기본 해시로 SHA1이 사용되었습니다.
      1.62부터는 모든 onedrive 유형에 대해 기본 해시로 QuickXorHash를 사용합니다. SHA1 해시가 필요한 경우에는
      이 옵션을 해당 값으로 설정하십시오.
      
      2023년 7월부터 QuickXorHash가 OneDrive for Business와 OneDrive Personal 모두에게 사용 가능한
      유일한 해시 유형이 될 것입니다.
      
      "none"으로 설정하면 해시를 사용하지 않습니다.
      
      요청된 해시가 객체에 없을 경우, rclone은 빈 문자열로 반환됩니다. 이는 rclone에서 누락된 해시로 처리됩니다.
      

      예:
         | auto     | rclone이 최상의 해시를 선택합니다.
         | quickxor | QuickXor
         | sha1     | SHA1
         | sha256   | SHA256
         | crc32    | CRC32
         | none     | 해시를 사용하지 않음

   --encoding
      백엔드용 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --client-id value      OAuth 클라이언트 ID. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 비밀키. [$CLIENT_SECRET]
   --help, -h             도움말 표시
   --region value         원드라이브의 국가 클라우드 지역 선택. (기본값: "global") [$REGION]

   고급 옵션

   --access-scopes value         rclone이 요청할 스코프를 설정합니다. (기본값: "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access") [$ACCESS_SCOPES]
   --auth-url value              인증 서버 URL. [$AUTH_URL]
   --chunk-size value            파일 업로드에 사용할 청크 크기 - 320k(327,680 바이트)의 배수여야 함. (기본값: "10Mi") [$CHUNK_SIZE]
   --disable-site-permission     Sites.Read.All 권한 요청 비활성화. (기본값: false) [$DISABLE_SITE_PERMISSION]
   --drive-id value              사용할 드라이브의 ID. [$DRIVE_ID]
   --drive-type value            드라이브 유형(개인 | 비즈니스 | 문서 라이브러리). [$DRIVE_TYPE]
   --encoding value              백엔드용 인코딩. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --expose-onenote-files        디렉토리 리스트에 OneNote 파일이 표시되도록 설정합니다. (기본값: false) [$EXPOSE_ONENOTE_FILES]
   --hash-type value             백엔드에서 사용 중인 해시 지정. (기본값: "auto") [$HASH_TYPE]
   --link-password value         링크 명령에 의해 생성된 링크의 비밀번호 설정. [$LINK_PASSWORD]
   --link-scope value            링크 명령에 의해 생성된 링크의 범위 설정. (기본값: "anonymous") [$LINK_SCOPE]
   --link-type value             링크 명령에 의해 생성된 링크의 유형 설정. (기본값: "view") [$LINK_TYPE]
   --list-chunk value            목록 청크 크기. (기본값: 1000) [$LIST_CHUNK]
   --no-versions                 수정 작업 시 모든 버전 삭제. (기본값: false) [$NO_VERSIONS]
   --root-folder-id value        루트 폴더의 ID. [$ROOT_FOLDER_ID]
   --server-side-across-configs  서버 측 작업(예: 복사)이 다른 onedrive 구성간에 작업할 수 있도록 합니다. (기본값: false) [$SERVER_SIDE_ACROSS_CONFIGS]
   --token value                 JSON blob 형식의 OAuth 액세스 토큰. [$TOKEN]
   --token-url value             토큰 서버 URL. [$TOKEN_URL]

   일반 옵션

   --name value  스토리지의 이름(기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}