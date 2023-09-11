# Mail.ru Cloud

{% code fullWidth="true" %}
```
이름:
   singularity storage create mailru - Mail.ru Cloud

사용법:
   singularity storage create mailru [command options] [arguments...]

설명:
   --user
      사용자 이름 (보통 이메일).

   --pass
      비밀번호.
      
      이는 앱 비밀번호여야 합니다. rclone은 일반 비밀번호로 작동하지 않습니다. 앱 비밀번호 생성 방법에 대해서는 설명서의 구성 섹션을 참조하세요.
      

   --speedup-enable
      동일한 데이터 해시를 가진 다른 파일이 있는 경우 전체 업로드를 건너뛸 것인지 여부.
      
      이 기능은 "speedup" 또는 "put by hash"라고도 불립니다. 보통의 책, 비디오 또는 오디오 클립과 같이 일반적으로 사용 가능한 파일의 경우 특히 효율적입니다. 파일은 모든 mailru 사용자 계정에서 해시로 검색됩니다. 소스 파일이 고유하거나 암호화되어 있는 경우에는 의미가 없으며 비효율적입니다. 또한, rclone은 미리 콘텐츠 해시를 계산하고 전체 업로드가 필요한지 여부를 결정할 수 있도록 로컬 메모리와 디스크 공간이 필요할 수 있습니다. 또한, rclone은 사이즈를 미리 알 수 없는 경우 (예: 스트리밍 또는 부분 업로드의 경우) 이 최적화를 시도하지 않습니다.

      예시:
         | true  | 활성화
         | false | 비활성화

   --speedup-file-patterns
      speedup (put by hash)을 위해 대상 파일 이름 패턴의 쉼표로 구분된 목록.
      
      패턴은 대소문자 구분이 없으며 '*' 또는 '?'와 같은 메타 문자를 포함할 수 있습니다.

      예시:
         | <unset>                 | 목록 없음. speedup (put by hash) 비활성화.
         | *                       | 모든 파일이 speedup을 위해 시도됩니다.
         | *.mkv,*.avi,*.mp4,*.mp3 | 일반적인 오디오/비디오 파일에 대해서만 speedup이 시도됩니다.
         | *.zip,*.gz,*.rar,*.pdf  | 일반적인 아카이브 또는 PDF 책에 대해서만 speedup이 시도됩니다.

   --speedup-max-disk
      대용량 파일에 대한 speedup (put by hash) 비활성화.
      
      이유는 예비 해싱이 RAM 또는 디스크 공간을 고갈시킬 수 있기 때문입니다.

      예시:
         | 0  | speedup (put by hash) 완전 비활성화.
         | 1G | 1GB보다 큰 파일은 직접 업로드됩니다.
         | 3G | 로컬 디스크에 3GB 미만 여유 공간이 있는 경우 이 옵션을 선택하세요.

   --speedup-max-memory
      아래에 제시된 크기보다 큰 파일은 항상 디스크에서 해싱됩니다.

      예시:
         | 0    | 예비 해싱은 항상 임시 디스크 위치에 수행됩니다.
         | 32M  | 예비 해싱에는 32MB 이상의 RAM을 할당하지 마십시오.
         | 256M | 해시 계산에 최대 256MB의 여유 RAM이 있습니다.

   --check-hash
      파일 체크섬이 불일치하거나 유효하지 않은 경우 복사 작업이 어떻게 처리되어야 하는지 여부.

      예시:
         | true  | 에러로 실패합니다.
         | false | 무시하고 계속합니다.

   --user-agent
      클라이언트 내부에서 사용되는 HTTP 사용자 에이전트.
      
      기본값은 "rclone/VERSION" 또는 명령줄에서 제공된 "--user-agent"입니다.

   --quirks
      내부 유지 관리 플래그의 쉼표로 구분된 목록.
      
      이 옵션은 일반 사용자가 사용해서는 안 됩니다. 이 옵션은 백엔드 문제의 원격 문제 해결을 용이하게 하는 것이 목적입니다. 플래그의 엄격한 의미는 문서화되지 않았으며 업데이트 사이에 지속되지 않음이 보장되지 않습니다. 백엔드가 안정적으로 성장하면 quirks는 제거될 것입니다.
      지원되는 quirks: atomicmkdir binlist unknowndirs

   --encoding
      백엔드의 인코딩.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --help, -h        도움말 표시
   --pass value      비밀번호. [$PASS]
   --speedup-enable  동일한 데이터 해시를 가진 다른 파일이 있는 경우 전체 업로드를 건너뛸 것인지 여부. (기본값: true) [$SPEEDUP_ENABLE]
   --user value      사용자 이름 (보통 이메일). [$USER]

   고급 옵션

   --check-hash                   파일 체크섬이 불일치하거나 유효하지 않은 경우 복사 작업이 어떻게 처리되어야 하는지 여부. (기본값: true) [$CHECK_HASH]
   --encoding value               백엔드의 인코딩. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --quirks value                 쉼표로 구분된 내부 유지 관리 플래그의 목록. [$QUIRKS]
   --speedup-file-patterns value  speedup (put by hash)을 위해 대상 파일 이름 패턴의 쉼표로 구분된 목록. (기본값: "*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf") [$SPEEDUP_FILE_PATTERNS]
   --speedup-max-disk value       대용량 파일에 대한 speedup (put by hash) 비활성화. (기본값: "3Gi") [$SPEEDUP_MAX_DISK]
   --speedup-max-memory value     아래에 제시된 크기보다 큰 파일은 항상 디스크에서 해싱됩니다. (기본값: "32Mi") [$SPEEDUP_MAX_MEMORY]
   --user-agent value             클라이언트 내부에서 사용되는 HTTP 사용자 에이전트. [$USER_AGENT]

   일반

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}