# Mail.ru Cloud

{% code fullWidth="true" %}
```
이름:
   singularity storage update mailru - Mail.ru Cloud

사용법:
   singularity storage update mailru [command options] <name|id>

설명:
   --user
      사용자 이름 (일반적으로 이메일).

   --pass
      비밀번호.
      
      이는 앱 비밀번호여야 합니다. 일반 비밀번호로는 rclone이 작동하지 않습니다. 앱 비밀번호를 만드는 방법에 대해서는 문서의 구성 섹션을 참조하십시오.
      

   --speedup-enable
      같은 데이터 해시를 가진 다른 파일이 있을 경우 전체 업로드를 건너뛰십시오.
      
      이 기능은 "속도 향상" 또는 "해시에 의한 업로드"라고도 불립니다. 이는 일반적으로 인기 있는 도서, 비디오 또는 오디오 클립과 같은 파일에서 특히 효과적입니다.
      왜냐하면 파일들은 모든 mailru 사용자의 계정에서 해시로 검색되기 때문입니다. 파일이 고유하거나 암호화된 경우에는 의미 없고 비효율적입니다.
      또한 rclone은 전체 업로드가 필요한지 미리 결정하기 위해 내부적으로 콘텐츠 해시를 계산하고 로컬 메모리와 디스크 공간이 필요할 수 있다는 점을 유의하십시오.
      또한 rclone은 파일 크기를 미리 알지 못하는 경우 (예 : 스트리밍 또는 부분 업로드의 경우) 이 최적화를 시도하지 않을 것입니다.

      예:
         | true  | 활성화
         | false | 비활성화

   --speedup-file-patterns
      해시에 의한 업로드가 가능한 파일 이름 패턴의 쉼표로 구분된 목록입니다.
      
      패턴은 대소문자를 구분하지 않으며 '*' 또는 '?' 메타 문자를 포함할 수 있습니다.

      예:
         | <unset>                 | 빈 목록은 해시에 의한 업로드를 완전히 비활성화합니다.
         | *                       | 모든 파일이 해시에 의해 업로드를 시도할 것입니다.
         | *.mkv,*.avi,*.mp4,*.mp3 | 일반적인 오디오/비디오 파일만 해시에 의해 업로드를 시도할 것입니다.
         | *.zip,*.gz,*.rar,*.pdf  | 일반적인 아카이브 또는 PDF 도서만 해시에 의해 업로드를 시도할 것입니다.

   --speedup-max-disk
      이 옵션을 사용하여 큰 파일에 대해 해시에 의한 업로드를 비활성화할 수 있습니다.
      
      사전 해싱을 하면 RAM 또는 디스크 공간이 고갈될 수 있기 때문입니다.

      예:
         | 0  | 해시에 의한 업로드를 완전히 비활성화합니다.
         | 1G | 1GB보다 큰 파일은 직접 업로드됩니다.
         | 3G | 로컬 디스크에 3GB 미만의 여유 공간이 있는 경우 이 옵션을 선택하십시오.

   --speedup-max-memory
      아래에 지정된 크기보다 큰 파일은 항상 디스크에서 해싱됩니다.

      예:
         | 0    | 사전 해싱은 항상 임시 디스크 위치에서 수행됩니다.
         | 32M  | 사전 해싱에 32MB 이상의 RAM을 할당하지 마십시오.
         | 256M | 해시 계산을 위해 여유 있는 256MB RAM이 있습니다.

   --check-hash
      파일 체크섬이 일치하지 않거나 올바르지 않은 경우 어떻게 복사해야 하는지 여부를 결정합니다.

      예:
         | true  | 오류로 실패합니다.
         | false | 무시하고 계속 진행합니다.

   --user-agent
      클라이언트에서 내부적으로 사용되는 HTTP 사용자 에이전트입니다.
      
      기본값은 "rclone/버전"이거나 명령 줄에서 제공한 "--user-agent"입니다.

   --quirks
      내부 유지 관리 플래그의 쉼표로 구분된 목록입니다.
      
      이 옵션은 일반 사용자가 사용하지 않아야 합니다. 이는 주로 원격으로 백엔드 문제를 해결하기 위한 목적으로만 사용됩니다.
      플래그의 엄격한 의미는 문서화되지 않았으며 릴리스 간에 지속되지 않을 수 있습니다.
      백엔드가 안정되면 Quirks는 제거됩니다.
      지원되는 특징: atomicmkdir binlist unknowndirs

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --help, -h        도움말 표시
   --pass value      비밀번호. [$PASS]
   --speedup-enable  같은 데이터 해시를 가진 다른 파일이 있을 경우 전체 업로드를 건너뛰십시오. (기본값: true) [$SPEEDUP_ENABLE]
   --user value      사용자 이름 (일반적으로 이메일). [$USER]

   고급

   --check-hash                   파일 체크섬이 일치하지 않거나 올바르지 않은 경우 어떻게 복사해야 하는지 여부를 결정합니다. (기본값: true) [$CHECK_HASH]
   --encoding value               백엔드의 인코딩. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --quirks value                 내부 유지 관리 플래그의 쉼표로 구분된 목록입니다. [$QUIRKS]
   --speedup-file-patterns value  해시에 의한 업로드가 가능한 파일 이름 패턴의 쉼표로 구분된 목록입니다. (기본값: "*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf") [$SPEEDUP_FILE_PATTERNS]
   --speedup-max-disk value       이 옵션을 사용하여 큰 파일에 대해 해시에 의한 업로드를 비활성화할 수 있습니다. (기본값: "3Gi") [$SPEEDUP_MAX_DISK]
   --speedup-max-memory value     아래에 지정된 크기보다 큰 파일은 항상 디스크에서 해싱됩니다. (기본값: "32Mi") [$SPEEDUP_MAX_MEMORY]
   --user-agent value             클라이언트에서 내부적으로 사용되는 HTTP 사용자 에이전트입니다. [$USER_AGENT]
```
{% endcode %}