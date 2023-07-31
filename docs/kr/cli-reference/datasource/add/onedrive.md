# Microsoft OneDrive

{% code fullWidth="true" %}
```
이름:
   singularity datasource add onedrive - Microsoft OneDrive

사용법:
   singularity datasource add onedrive [command options] <dataset_name> <source_path>

설명:
   --onedrive-access-scopes
      rclone에 요청할 스코프를 설정합니다.
      
      선택하거나 직접 사용자 정의로 스코프를 공백으로 구분하여 입력하십시오.
      

      예시:
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access | 모든 리소스에 대한 읽기 및 쓰기 권한
         | Files.Read Files.Read.All Sites.Read.All offline_access                                     | 모든 리소스에 대한 읽기 전용 권한
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All offline_access                | 모든 리소스에 대한 읽기 및 쓰기 권한, SharePoint 사이트 탐색 기능 없음. 
                                                                                                       | disable_site_permission를 true로 설정한 경우와 동일한 동작

   --onedrive-auth-url
      인증 서버 URL입니다.
      
      공백으로 둘 경우 공급자의 기본값을 사용합니다.

   --onedrive-chunk-size
      파일을 업로드하는 데 사용되는 청크 크기입니다. - 320k의 배수여야 합니다 (327,680바이트).
      
      이 크기를 초과하는 파일은 청크로 분할됩니다 - 320k의 배수여야 하며
      250M (262,144,000바이트)보다 크지 않아야 합니다. 그렇지 않으면 \"Microsoft.SharePoint.Client.InvalidClientQueryException: The request message is too big.\" 오류가 발생할 수 있습니다.
      청크는 메모리에 버퍼링됩니다.

   --onedrive-client-id
      OAuth 클라이언트 ID입니다.
      
      보통 비워 둡니다.

   --onedrive-client-secret
      OAuth 클라이언트 시크릿입니다.
      
      보통 비워 둡니다.

   --onedrive-disable-site-permission
      Sites.Read.All 권한 요청 비활성화하기.
      
      true로 설정하면, 드라이브 ID 구성 중 SharePoint 사이트를 검색할 수 없게 됩니다
      왜냐하면 rclone이 Sites.Read.All 권한을 요청하지 않기 때문입니다.
      조직에서 애플리케이션에 Sites.Read.All 권한을 할당하지 않고 사용자가
      앱 권한 요청에 동의할 수 없게 하는 경우 true로 설정하십시오.

   --onedrive-drive-id
      사용할 드라이브의 ID입니다.

   --onedrive-drive-type
      드라이브의 유형입니다 (personal | business | documentLibrary).

   --onedrive-encoding
      백엔드에 사용할 인코딩입니다.
      
      자세한 내용은 [설명서의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --onedrive-expose-onenote-files
      디렉토리 목록에 OneNote 파일이 표시되도록 설정합니다.
      
      기본적으로 rclone은 디렉토리 목록에서 OneNote 파일을 숨깁니다
      "열기" 또는 "업데이트"와 같은 작업이 작동하지 않기 때문입니다. 그러나
      이 동작은 OneNote 파일을 삭제하는 데도 방해할 수 있습니다.
      OneNote 파일을 삭제하거나 디렉토리 목록에 표시하려는 경우 이 옵션을 설정하십시오.

   --onedrive-hash-type
      백엔드에서 사용할 해시를 지정합니다.
      
      이 옵션은 사용할 해시 유형을 지정합니다. "auto"로 설정하면
      기본 해시인 QuickXorHash를 사용합니다.
      
      rclone 1.62 이전에는 Onedrive Personal의 기본값이 SHA1 해시였습니다.
      1.62 이상부터는 모든 onedrive 유형에 대해 기본값으로 QuickXorHash를 사용합니다. SHA1 해시가 필요한 경우에는 옵션을 이에 맞게 설정하십시오.
      
      2023년 7월부터는 QuickXorHash가 OneDrive for Business 및 OneDrive Personal의
      유일한 사용 가능한 해시 유형입니다.
      
      이 옵션을 "none"으로 설정하여 해시를 사용하지 않을 수 있습니다.
      
      요청한 해시가 개체에 없는 경우 빈 문자열로 반환되며 rclone에서 누락된 해시로 처리됩니다.
      

      예시:
         | auto     | Rclone이 최상의 해시를 선택합니다.
         | quickxor | QuickXor
         | sha1     | SHA1
         | sha256   | SHA256
         | crc32    | CRC32
         | none     | 사용하지 않음 - 어떤 해시도 사용하지 않습니다.

   --onedrive-link-password
      link 명령어로 생성된 링크의 암호를 설정합니다.
      
      작성 시점에는 이 기능은 개인용 OneDrive에서만 지원됩니다.
      

   --onedrive-link-scope
      link 명령으로 생성된 링크의 영역을 설정합니다.

      예시:
         | anonymous    | 링크에 대한 사인인이 필요 없이 링크에 액세스할 수 있는 사람은 누구나 가능합니다.
                        | 조직 밖 사람들도 포함될 수 있습니다.
                        | 익명 링크 지원은 관리자에 의해 비활성화될 수 있습니다.
         | organization | 조직에 로그인한 모든 사람(테넌트)이 링크를 사용하여 액세스할 수 있습니다.
                        | OneDrive for Business 및 SharePoint에서만 사용 가능합니다.

   --onedrive-link-type
      link 명령으로 생성된 링크의 유형을 설정합니다.

      예시:
         | view  | 항목에 대한 읽기 전용 링크를 생성합니다.
         | edit  | 항목에 대한 읽기/쓰기 링크를 생성합니다.
         | embed | 항목에 대한 포함 가능한 링크를 생성합니다.

   --onedrive-list-chunk
      목록 청크의 크기입니다.

   --onedrive-no-versions
      수정 작업 중 모든 버전 삭제하기.
      
      Onedrive for business는 기존 파일을 덮어쓸 때 및 수정 시간을 설정할 때
      버전을 만듭니다.
      
      이러한 버전은 할당량에서 공간을 차지합니다.
      
      이 플래그는 파일 업로드 및 수정 시간 설정 후 버전을 확인하고
      마지막 버전을 제외한 모든 버전을 삭제합니다.
      
      **참고** Onedrive personal에서는 현재 버전을 삭제할 수 없으므로 플래그를 사용하지 마십시오.
      

   --onedrive-region
      OneDrive용 국가 클라우드 지역을 선택합니다.

      예시:
         | global | Microsoft Cloud Global
         | us     | Microsoft Cloud for US Government
         | de     | Microsoft Cloud Germany
         | cn     | Azure and Office 365 operated by Vnet Group in China

   --onedrive-root-folder-id
      루트 폴더의 ID입니다.
      
      보통 필요하지 않지만 특정한 경우에는 액세스하려는 폴더의 ID를 알고,
      경로 탐색을 통해 해당 폴더에 도달할 수 없는 경우에 사용할 수 있습니다.
      

   --onedrive-server-side-across-configs
      서버 측 작업(예: 복사)을 다른 onedrive 구성 사이에서 작동하도록 허용합니다.
      
      이 기능은 두 개의 OneDrive *Personal* 드라이브 간에 복사하는 경우에만 작동합니다
      그리고 복사할 파일이 이미 공유되어 있는 경우입니다. 다른 경우에는 rclone이
      일반 복사로 되돌아갑니다 (이 경우 약간 더 느릴 수 있음).

   --onedrive-token
      OAuth 액세스 토큰(JSON blob)입니다.

   --onedrive-token-url
      토큰 서버 URL입니다.
      
      공백으로 두면 공급자의 기본값을 사용합니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보낸 후 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공한 스캔 시기로부터 경과한 시간을 기준으로 소스 디렉토리를 자동으로 다시 스캔합니다 (기본값: 해제됨)
   --scanning-state value   초기 스캔 상태 설정 (기본값: 준비 완료)

   onedrive 옵션

   --onedrive-access-scopes value               rclone에 요청할 스코프를 설정합니다. (기본값: "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access") [$ONEDRIVE_ACCESS_SCOPES]
   --onedrive-auth-url value                    인증 서버 URL입니다. [$ONEDRIVE_AUTH_URL]
   --onedrive-chunk-size value                  파일을 업로드하는 데 사용되는 청크 크기입니다. (기본값: "10Mi") [$ONEDRIVE_CHUNK_SIZE]
   --onedrive-client-id value                   OAuth 클라이언트 ID입니다. [$ONEDRIVE_CLIENT_ID]
   --onedrive-client-secret value               OAuth 클라이언트 시크릿입니다. [$ONEDRIVE_CLIENT_SECRET]
   --onedrive-drive-id value                    사용할 드라이브의 ID입니다. [$ONEDRIVE_DRIVE_ID]
   --onedrive-drive-type value                  드라이브의 유형입니다 (personal | business | documentLibrary). [$ONEDRIVE_DRIVE_TYPE]
   --onedrive-encoding value                    백엔드에 사용할 인코딩입니다. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ONEDRIVE_ENCODING]
   --onedrive-expose-onenote-files value        디렉토리 목록에 OneNote 파일이 표시되도록 설정합니다. (기본값: "false") [$ONEDRIVE_EXPOSE_ONENOTE_FILES]
   --onedrive-hash-type value                   백엔드에서 사용할 해시를 지정합니다. (기본값: "auto") [$ONEDRIVE_HASH_TYPE]
   --onedrive-link-password value               link 명령어로 생성된 링크의 암호를 설정합니다. [$ONEDRIVE_LINK_PASSWORD]
   --onedrive-link-scope value                  link 명령으로 생성된 링크의 영역을 설정합니다. (기본값: "anonymous") [$ONEDRIVE_LINK_SCOPE]
   --onedrive-link-type value                   link 명령으로 생성된 링크의 유형을 설정합니다. (기본값: "view") [$ONEDRIVE_LINK_TYPE]
   --onedrive-list-chunk value                  목록 청크의 크기입니다. (기본값: "1000") [$ONEDRIVE_LIST_CHUNK]
   --onedrive-no-versions value                 수정 작업 중 모든 버전 삭제하기. (기본값: "false") [$ONEDRIVE_NO_VERSIONS]
   --onedrive-region value                      OneDrive용 국가 클라우드 지역을 선택합니다. (기본값: "global") [$ONEDRIVE_REGION]
   --onedrive-root-folder-id value              루트 폴더의 ID입니다. [$ONEDRIVE_ROOT_FOLDER_ID]
   --onedrive-server-side-across-configs value  서버 측 작업(예: 복사)을 다른 onedrive 구성 사이에서 작동하도록 허용합니다. (기본값: "false") [$ONEDRIVE_SERVER_SIDE_ACROSS_CONFIGS]
   --onedrive-token value                       OAuth 액세스 토큰(JSON blob)입니다. [$ONEDRIVE_TOKEN]
   --onedrive-token-url value                   토큰 서버 URL입니다. [$ONEDRIVE_TOKEN_URL]

```
{% endcode %}