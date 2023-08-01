# 로컬 디스크

{% code fullWidth="true" %}
```
명령어:
   singularity datasource add local - 로컬 디스크

사용법:
   singularity datasource add local [command options] <dataset_name> <source_path>

설명:
   --local-case-insensitive
      파일 시스템이 대소문자를 구분하지 않는 것으로 보고하도록 강제합니다.
      
      일반적으로 로컬 백엔드는 Windows/macOS에서는 대소문자를 구분하지 않는 것으로,
      그 외의 경우에는 대소문자를 구분한다고 보고합니다.
      기본 선택사항을 무시하려면 이 플래그를 사용하세요.

   --local-case-sensitive
      파일 시스템이 대소문자를 구분한다고 보고하도록 강제합니다.
      
      일반적으로 로컬 백엔드는 Windows/macOS에서는 대소문자를 구분하지 않는 것으로,
      그 외의 경우에는 대소문자를 구분한다고 보고합니다.
      기본 선택사항을 무시하려면 이 플래그를 사용하세요.

   --local-copy-links
      심볼릭 링크를 따라가고 링크된 항목을 복사합니다.

   --local-encoding
      백엔드의 인코딩입니다.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --local-links
      심볼릭 링크를 일반 파일로 바꾸거나 반대로 일반 파일을 심볼릭 링크로 번역합니다.
      
   --local-no-check-updated
      업로드 중 파일이 변경되었는지 확인하지 않습니다.
      
      일반적으로 rclone은 파일을 업로드하는 동안 파일의 크기와 수정 시간을 확인하고,
      파일이 업로드 중에 변경되면 "can't copy - source file is being updated"로 시작하는 메시지를 출력하여 업로드를 중단합니다.
      
      그러나 몇몇 파일 시스템에서는 이러한 수정 시간 확인이 실패할 수 있습니다.
      이 플래그를 사용하면 이러한 확인을 비활성화할 수 있습니다.
      
      이 플래그가 설정되면, rclone은 업데이트 중인 파일을 전송하기 위해 최선을 다할 것입니다.
      파일이 추가될 때만(예: 로그 파일) rclone은 처음 본 시점에 해당하는 크기로 로그 파일을 전송합니다.
      
      파일이 계속 수정되면(단순히 추가되는 것이 아니라), 전송이 해시 체크 실패로 실패할 수 있습니다.
      
      자세히 말하면, 파일의 stat()이 처음 호출된 후에는:
      
      - 처음 stat에서 전달된 크기만 전송합니다.
      - 처음 stat에서 전달된 크기만 체크섬합니다.
      - 파일의 stat 정보를 업데이트하지 않습니다.

   --local-no-preallocate
      전송된 파일을 위한 디스크 공간 사전 할당 기능을 비활성화합니다.
      
      디스크 공간 사전 할당은 파일 시스템 단편화를 방지하는 데 도움이 됩니다.
      그러나 일부 가상 파일 시스템 레이어(Google Drive File Stream과 같은)는 실제 파일 크기를
      사전 할당된 공간과 동일하게 설정할 수 있으므로, 체크섬 및 파일 크기 검사가 실패할 수 있습니다.
      사전 할당을 비활성화하려면 이 플래그를 사용하세요.

   --local-no-set-modtime
      수정 시간 설정을 비활성화합니다.
      
      일반적으로 rclone은 파일 전송 후에 수정 시간을 업데이트합니다.
      이는 rclone이 파일을 업로드하는 사용자가 업로드한 파일을소유하지 않는
      Linux 플랫폼에서 권한 문제를 유발할 수 있습니다.
      예를 들어 다른 사용자가 소유한 CIFS 마운트로 복사하는 경우입니다.
      이 옵션을 사용하면 rclone은 파일을 복사한 후에 수정 시간을 더이상 업데이트하지 않습니다.

   --local-no-sparse
      멀티스레드 다운로드에 대한 희소 파일을 비활성화합니다.
      
      Windows 플랫폼에서 rclone은 멀티스레드 다운로드시 희소 파일을 생성합니다.
      이렇게 함으로써 OS가 파일을 영점으로 설정하는 것을 피합니다.
      그러나 희소 파일은 디스크 단편화를 유발하고 작업하기가 느릴 수 있으므로 원하지 않을 수 있습니다.

   --local-nounc
      Windows에서 UNC(긴 파일 경로) 변환이 비활성화됩니다.

      예:
         | true | 긴 파일 이름이 비활성화됩니다.

   --local-one-file-system
      파일 시스템 경계를 넘어가지 않습니다(unix/macOS 전용).

   --local-skip-links
      건너뛰어야 할 심볼릭 링크에 대해 경고하지 않습니다.
      
      이 플래그는 건너뛰어야 할 심볼릭 링크 또는 접점을 경고 메시지에서 비활성화하므로,
      해당 링크가 건너뛸 대상임을 명시적으로 알립니다.

   --local-unicode-normalization
      경로와 파일 이름에 유니코드 NFC 정규화를 적용합니다.
      
      이 플래그를 사용하면 로컬 파일 시스템에서 읽은 파일 이름을 유니코드 NFC 형식으로 정규화할 수 있습니다.
      
      rclone은 파일 시스템에서 읽은 파일 이름의 인코딩을 일반적으로 수정하지 않습니다.
      
      macOS를 사용할 때 유용하며, 이는 일부 다른 OS에서 (예: 한국어에서) 올바르게 표시되지 않는
      분해(분해, NFD) 유니코드를 일반적으로 제공합니다.
      
      주의: rclone은 동기화 루틴에서 유니코드 정규화로 파일 이름을 비교하므로,
      일반적으로 이 플래그를 사용하지 않아야 합니다.

   --local-zero-size-links
      링크의 Stat 크기를 0으로 가정하고 해당 링크를 읽습니다(변경됨).
      
      rclone은 이전에 링크의 Stat 크기를 링크 크기로 사용했지만,
      이는 다음 위치에서 실패합니다.
      
      - Windows
      - 일부 가상 파일 시스템(예: LucidLink)
      - Android
      
      따라서 rclone은 이제 항상 링크를 읽습니다.
      


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] 데이터셋의 파일을 CAR 파일로 내보낸 후에 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막에 성공한 스캔으로부터 이 간격이 경과하면 소스 디렉터리를 자동으로 다시 스캔합니다 (기본값: 사용 안 함)
   --scanning-state value   초기 스캔 상태를 설정합니다 (기본값: 준비 완료)

   로컬 옵션

   --local-case-insensitive value       파일 시스템이 대소문자를 구분하지 않는 것으로 보고하도록 강제합니다. (기본값: "false") [$LOCAL_CASE_INSENSITIVE]
   --local-case-sensitive value         파일 시스템이 대소문자를 구분한다고 보고하도록 강제합니다. (기본값: "false") [$LOCAL_CASE_SENSITIVE]
   --local-copy-links value             심볼릭 링크를 따라가고 링크된 항목을 복사합니다. (기본값: "false") [$LOCAL_COPY_LINKS]
   --local-encoding value               백엔드의 인코딩입니다. (기본값: "Slash,Dot") [$LOCAL_ENCODING]
   --local-links value                  심볼릭 링크를 일반 파일로 바꾸거나 반대로 일반 파일을 심볼릭 링크로 번역합니다. (기본값: "false") [$LOCAL_LINKS]
   --local-no-check-updated value       업로드 중 파일이 변경되었는지 확인하지 않습니다. (기본값: "false") [$LOCAL_NO_CHECK_UPDATED]
   --local-no-preallocate value         전송된 파일을 위한 디스크 공간 사전 할당 기능을 비활성화합니다. (기본값: "false") [$LOCAL_NO_PREALLOCATE]
   --local-no-set-modtime value         수정 시간 설정을 비활성화합니다. (기본값: "false") [$LOCAL_NO_SET_MODTIME]
   --local-no-sparse value              멀티스레드 다운로드에 대한 희소 파일을 비활성화합니다. (기본값: "false") [$LOCAL_NO_SPARSE]
   --local-nounc value                  Windows에서 UNC(긴 파일 경로) 변환이 비활성화됩니다. (기본값: "false") [$LOCAL_NOUNC]
   --local-one-file-system value        파일 시스템 경계를 넘어가지 않습니다(unix/macOS 전용). (기본값: "false") [$LOCAL_ONE_FILE_SYSTEM]
   --local-skip-links value             건너뛰어야 할 심볼릭 링크에 대해 경고하지 않습니다. (기본값: "false") [$LOCAL_SKIP_LINKS]
   --local-unicode-normalization value  경로와 파일 이름에 유니코드 NFC 정규화를 적용합니다. (기본값: "false") [$LOCAL_UNICODE_NORMALIZATION]
   --local-zero-size-links value        링크의 Stat 크기를 0으로 가정하고 해당 링크를 읽습니다(변경됨). (기본값: "false") [$LOCAL_ZERO_SIZE_LINKS]

```
{% endcode %}