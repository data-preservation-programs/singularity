# 로컬 디스크

{% code fullWidth="true" %}
```
명령어:
   singularity storage create local - 로컬 디스크 생성

사용법:
   singularity storage create local [옵션] [인수]

설명:
   --nounc
      Windows에서 UNC (긴 경로 이름) 변환을 비활성화합니다.

      예시:
         | true | 긴 파일 이름을 비활성화합니다.

   --copy-links
      심볼릭 링크를 따라가고 링크된 항목을 복사합니다.

   --links
      심볼릭 링크를 '.rclonelink' 확장자를 가진 일반 파일로 변환하거나 그 반대로 변환합니다.

   --skip-links
      건너뜀된 심볼릭 링크에 대해 경고를 표시하지 않습니다.
      
      이 플래그는 건너뛰어져야 하는 심볼릭 링크나 정점 지점에 대한 경고 메시지를 비활성화합니다.
      명시적으로 건너뛰어야 하기 때문에 사용합니다.

   --zero-size-links
      링크의 Stat 크기가 0인 것으로 가정하고 (그리고 링크를 읽음) (사용되지 않음).
      
      Rclone은 링크의 Stat 크기를 링크 크기로 사용했었으나 이 방식은 다음과 같은 여러 상황에서 실패합니다:
      
      - Windows
      - 일부 가상 파일 시스템 (예: LucidLink)
      - 안드로이드
      
      그래서 rclone은 이제 항상 링크를 읽습니다.
      

   --unicode-normalization
      경로와 파일 이름에 유니코드 NFC 정규화를 적용합니다.
      
      이 플래그는 로컬 파일 시스템에서 읽은 파일 이름을 유니코드 NFC 형식으로 정규화하는 데 사용될 수 있습니다.
      
      Rclone은 일반적으로 파일 시스템에서 읽은 파일 이름의 인코딩을 변경하지 않습니다.
      
      이 플래그는 macOS를 사용할 때 유용합니다. macOS는 보통 분해된 (NFD) 유니코드를 제공하는데, 어떤 OS에서는 (예: 한글의 경우) 일부로 표시되지 않을 수 있습니다.
      
      이 플래그를 사용하지 않는 것이 보통입니다. rclone은 동기화 루틴에서 유니코드 정규화를 사용하여 파일 이름을 비교하기 때문입니다.

   --no-check-updated
      파일이 업로드 중에 변경되는지 확인하지 않습니다.
      
      보통 파일은 업로드하는 동안 크기와 수정 시간을 확인하고, 만약 파일이 업로드 중에 변경된다면 "can't copy - source file is being updated"로 시작하는 메시지와 함께 중단합니다.
      
      그러나 일부 파일 시스템에서 이 수정 시간 확인이 실패할 수 있습니다(예: [Glusterfs #2206](https://github.com/rclone/rclone/issues/2206)).
      따라서 이 플래그로 이 확인을 비활성화할 수 있습니다.
      
      이 플래그가 설정되어 있는 경우, rclone은 업로드 중인 파일을 전송하기 위해 최선의 노력을 다할 것입니다.
      파일이 단순히 추가만 되고 있다면(예: 로그), rclone은 로그 파일을 처음 본 때의 크기로 전송합니다.
      
      파일이 계속 수정된다면(추가만 하는 것이 아니라면) 전송이 해시 검사 실패로 실패할 수 있습니다.
      
      자세한 내용은 파일에 stat()가 처음으로 호출된 후에 다음을 수행합니다.
      
      - stat가 제공한 크기만 전송
      - stat의 크기만 체크섬 계산
      - 파일에 대한 stat 정보를 업데이트하지 않음
      
      

   --one-file-system
      파일 시스템 경계를 넘지 않습니다 (unix/macOS 전용).

   --case-sensitive
      파일 시스템을 대소문자 구분으로 표시하도록 강제합니다.
      
      일반적으로 로컬 백엔드는 Windows/macOS에서 대소문자를 구분하지 않으며, 다른 모든 경우에는 대소문자를 구분하도록 선언합니다. 이 플래그를 사용하여 기본 설정을 재정의합니다.

   --case-insensitive
      파일 시스템을 대소문자 구분 없게 표시하도록 강제합니다.
      
      일반적으로 로컬 백엔드는 Windows/macOS에서 대소문자를 구분하지 않으며, 다른 모든 경우에는 대소문자를 구분하도록 선언합니다. 이 플래그를 사용하여 기본 설정을 재정의합니다.

   --no-preallocate
      전송된 파일에 대해 디스크 공간 사전 할당을 비활성화합니다.
      
      디스크 공간 사전 할당은 파일 시스템 단편화를 방지하는 데 도움이 됩니다.
      그러나 일부 가상 파일 시스템 레이어 (예: Google Drive 파일 스트림)에서는 실제 파일 크기를 사전 할당된 공간으로 설정할 수 있어, 체크섬 및 파일 크기 확인에 실패할 수 있습니다.
      이 플래그를 사용하여 사전 할당을 비활성화합니다.

   --no-sparse
      다중 스레드 다운로드에 대해 공간 절약 파일을 비활성화합니다.
      
      Windows 플랫폼에서 rclone은 다중 스레드 다운로드 중에 공간 절약 파일을 만듭니다. 이렇게 함으로써 파일의 0 값을 설정하는 OS의 오랜 대기 시간을 피할 수 있습니다.
      그러나 공간 절약 파일은 디스크 단편화를 유발하고 처리 속도가 느릴 수 있으므로 원하지 않을 수 있습니다.

   --no-set-modtime
      수정 시간 설정을 비활성화합니다.
      
      보통 rclone은 파일을 업로드한 후 수정 시간을 업데이트합니다.
      이는 Linux 플랫폼에서 권한 문제를 발생시킬 수 있습니다. rclone이 실행되는 사용자가 업로드한 파일의 소유자가 아닌 경우, 다른 사용자에게 소유된 CIFS 마운트에 복사하는 경우 등.
      이 옵션이 활성화되면 rclone은 파일을 복사한 후 수정 시간을 더이상 업데이트하지 않습니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 확인하세요.


옵션:
   --help, -h  도움말 표시

   고급 옵션

   --case-insensitive       파일 시스템을 대소문자 구분 없게 표시하도록 강제합니다. (기본값: false) [$CASE_INSENSITIVE]
   --case-sensitive         파일 시스템을 대소문자 구분으로 표시하도록 강제합니다. (기본값: false) [$CASE_SENSITIVE]
   --copy-links, -L         심볼릭 링크를 따라가고 링크된 항목을 복사합니다. (기본값: false) [$COPY_LINKS]
   --encoding value         백엔드의 인코딩입니다. (기본값: "Slash,Dot") [$ENCODING]
   --links, -l              심볼릭 링크를 '.rclonelink' 확장자를 가진 일반 파일로 변환하거나 그 반대로 변환합니다. (기본값: false) [$LINKS]
   --no-check-updated       파일이 업로드 중에 변경되는지 확인하지 않습니다. (기본값: false) [$NO_CHECK_UPDATED]
   --no-preallocate         전송된 파일에 대해 디스크 공간 사전 할당을 비활성화합니다. (기본값: false) [$NO_PREALLOCATE]
   --no-set-modtime         수정 시간 설정을 비활성화합니다. (기본값: false) [$NO_SET_MODTIME]
   --no-sparse              다중 스레드 다운로드에 대해 공간 절약 파일을 비활성화합니다. (기본값: false) [$NO_SPARSE]
   --nounc                  Windows에서 UNC (긴 경로 이름) 변환을 비활성화합니다. (기본값: false) [$NOUNC]
   --one-file-system, -x    파일 시스템 경계를 넘지 않습니다 (unix/macOS 전용). (기본값: false) [$ONE_FILE_SYSTEM]
   --skip-links             건너뜀된 심볼릭 링크에 대해 경고를 표시하지 않습니다. (기본값: false) [$SKIP_LINKS]
   --unicode-normalization  경로와 파일 이름에 유니코드 NFC 정규화를 적용합니다. (기본값: false) [$UNICODE_NORMALIZATION]
   --zero-size-links        링크의 Stat 크기가 0인 것으로 가정하고 (그리고 링크를 읽음) (사용되지 않음). (기본값: false) [$ZERO_SIZE_LINKS]

   일반 옵션

   --name value  저장소 이름 (기본값: 자동 생성)
   --path value  저장소 경로

```
{% endcode %}