# 로컬 디스크

{% code fullWidth="true" %}
```
명령어:
   singularity storage update local - 로컬 디스크

사용법:
   singularity storage update local [command options] <name|id>

설명:
   --nounc
      Windows에서 UNC(긴 경로 이름) 변환 비활성화

      예시:
         | true | 긴 파일 이름 사용 안 함.

   --copy-links
      심볼릭 링크를 따라가고 가리키는 항목을 복사합니다.

   --links
      심볼릭 링크를 일반 파일로 변환하거나 일반 파일을 심볼릭 링크로 변환합니다(확장자 '.rclonelink' 사용).

   --skip-links
      건너뛰어진 심볼릭 링크에 대해 경고하지 않습니다.
      
      이 플래그를 사용하면 명시적으로 건너뛸 것을 확인하여, 건너뛰어진 심볼릭 링크나 점프 지점에 대한 경고 메시지가 비활성화됩니다.

   --zero-size-links
      링크의 Stat 크기를 0이라고 가정하고(그리고 링크를 읽음) (사용 중지됨).
      
      Rclone은 예전에 링크의 Stat 크기를 링크 크기로 사용했었으나, 다음과 같은 몇 가지 상황에서 실패했습니다:
      
      - Windows
      - 일부 가상 파일 시스템(예: LucidLink)
      - Android
      
      이제 rclone은 항상 링크를 읽습니다.
      

   --unicode-normalization
      경로와 파일 이름에 유니코드 NFC 정규화를 적용합니다.
      
      이 플래그는 로컬 파일 시스템에서 읽은 파일 이름을 유니코드 NFC 양식으로 정규화하는 데 사용될 수 있습니다.
      
      Rclone은 일반적으로 파일 시스템에서 읽은 파일 이름의 인코딩을 조작하지 않습니다.
      
      이 플래그는 macOS를 사용할 때 유용할 수 있습니다. macOS는 일반적으로 분해 된(NFD) 유니코드를 제공하며, 이는 어떤 OS에서(예: 한국어) 제대로 표시되지 않을 수 있습니다.
      
      또한 rclone은 동기화 루틴에서 유니코드 정규화로 파일 이름을 비교하므로 이 플래그는 일반적으로 사용하지 않아야 합니다.

   --no-check-updated
      파일이 업로드 중에 변경되었는지 확인하지 않습니다.
      
      보통 rclone은 파일을 업로드하는 동안 크기와 수정 시간을 확인하고, 파일이 업로드 중에 변경되면 "can't copy - source file is being updated"라는 메시지로 작업을 중단합니다.
      
      그러나 일부 파일 시스템에서 이 수정 시간 확인은 실패할 수 있습니다([Glusterfs #2206](https://github.com/rclone/rclone/issues/2206)). 따라서 이 플래그로 확인을 비활성화할 수 있습니다.
      
      이 플래그가 설정되면, rclone은 파일이 업데이트되고 있을 때 전송을 위한 최선의 노력을 기울일 것입니다. 파일이 단지 덧붙여지는 경우(예: 로그) rclone은 로그 파일을 전송하며, 처음 rclone이 그 파일을 확인했을 때의 크기로 전송합니다.
      
      파일이 변경되는 경우(단순히 덧붙여지기만 하는 것이 아닌) 전송은 해시 확인 실패로 실패할 수 있습니다.
      
      자세한 내용은 파일이 처음으로 stat()에 의해 호출된 이후에 다음과 같이 처리합니다:
      
      - stat에서 얻은 크기만 전송합니다.
      - stat에서 얻은 크기만 체크섬합니다.
      - 파일의 stat 정보를 업데이트하지 않습니다.
      
      

   --one-file-system
      파일 시스템 경계를 건너지 않습니다(unix/macOS 전용).

   --case-sensitive
      파일 시스템이 자신을 대소문자를 구분하는 것으로 보고하도록 강제합니다.
      
      보통 로컬 백엔드는 Windows/macOS에서 대소문자 구분을 지원하지 않으며, 다른 모든 환경에 대해 대소문자 구분을 지원합니다. 이 기본 설정을 무시하려면 이 플래그를 사용하세요.

   --case-insensitive
      파일 시스템이 자신을 대소문자를 구분하지 않는 것으로 보고하도록 강제합니다.
      
      보통 로컬 백엔드는 Windows/macOS에서 대소문자 구분을 지원하지 않으며, 다른 모든 환경에 대해 대소문자 구분을 지원합니다. 이 기본 설정을 무시하려면 이 플래그를 사용하세요.

   --no-preallocate
      파일이 전송되는 동안 디스크 공간의 사전 할당을 비활성화합니다.
      
      디스크 공간의 사전 할당은 파일 시스템 조각화를 방지하는 데 도움이 됩니다. 그러나 Google Drive File Stream과 같은 일부 가상 파일 시스템 계층은 실제 파일 크기를 미리 할당된 공간과 동일하게 설정할 수 있으며, 이로 인해 체크섬 및 파일 크기 확인이 실패할 수 있습니다. 이 플래그를 사용하여 사전 할당을 비활성화하세요.

   --no-sparse
      다중 스레드 다운로드 시 Sparse 파일을 사용하지 않습니다.
      
      Windows 플랫폼에서 rclone은 다중 스레드 다운로드 시 Sparse 파일을 생성합니다. 이렇게 함으로써 OS가 파일을 0으로 설정하는 경우 큰 파일에서 오랜 대기 시간을 방지합니다. 그러나 Sparse 파일은 디스크 조각화를 야기할 수 있으며 작업이 느려질 수 있습니다.

   --no-set-modtime
      수정 시간 설정을 비활성화합니다.
      
      보통 rclone은 파일을 업로드한 후에 수정 시간을 업데이트합니다. 이는 Linux 플랫폼에서 rclone 실행 사용자가 전송한 파일을 소유하지 않는 경우(다른 사용자가 소유한 CIFS 마운트에 복사하는 경우 등) 권한 문제를 일으킬 수 있습니다. 이 옵션을 사용하면 파일을 복사한 후에 수정 시간을 업데이트하지 않게 됩니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


OPTIONS:
   --help, -h  도움말 출력

   고급 옵션

   --case-insensitive       파일 시스템이 자신을 대소문자를 구분하지 않음으로 보고하도록 강제합니다. (기본값: false) [$CASE_INSENSITIVE]
   --case-sensitive         파일 시스템이 자신을 대소문자를 구분하는 것으로 보고하도록 강제합니다. (기본값: false) [$CASE_SENSITIVE]
   --copy-links, -L         심볼릭 링크를 따라가고 가리키는 항목을 복사합니다. (기본값: false) [$COPY_LINKS]
   --encoding value         백엔드의 인코딩입니다. (기본값: "Slash,Dot") [$ENCODING]
   --links, -l              심볼릭 링크를 일반 파일로 변환하거나 일반 파일을 심볼릭 링크로 변환합니다(확장자 '.rclonelink' 사용). (기본값: false) [$LINKS]
   --no-check-updated       파일이 업로드 중에 변경되었는지 확인하지 않습니다. (기본값: false) [$NO_CHECK_UPDATED]
   --no-preallocate         파일이 전송되는 동안 디스크 공간의 사전 할당을 비활성화합니다. (기본값: false) [$NO_PREALLOCATE]
   --no-set-modtime         수정 시간 설정을 비활성화합니다. (기본값: false) [$NO_SET_MODTIME]
   --no-sparse              다중 스레드 다운로드 시 Sparse 파일을 사용하지 않습니다. (기본값: false) [$NO_SPARSE]
   --nounc                  Windows에서 UNC(긴 경로 이름) 변환 비활성화 (기본값: false) [$NOUNC]
   --one-file-system, -x    파일 시스템 경계를 건너지 않습니다(unix/macOS 전용). (기본값: false) [$ONE_FILE_SYSTEM]
   --skip-links             건너뛰어진 심볼릭 링크에 대해 경고하지 않습니다. (기본값: false) [$SKIP_LINKS]
   --unicode-normalization  경로와 파일 이름에 유니코드 NFC 정규화를 적용합니다. (기본값: false) [$UNICODE_NORMALIZATION]
   --zero-size-links        링크의 Stat 크기를 0이라고 가정하고(그리고 링크를 읽음) (사용 중지됨). (기본값: false) [$ZERO_SIZE_LINKS]

```
{% endcode %}