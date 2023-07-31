# Mail.ru Cloud

{% code fullWidth="true" %}
```
이름:
   singularity datasource add mailru - Mail.ru Cloud

사용법:
   singularity datasource add mailru [command options] <dataset_name> <source_path>

설명:
   --mailru-check-hash
      파일 체크섬이 일치하지 않거나 유효하지 않은 경우 복사 작업이 어떻게 진행되어야 합니까?

      예시:
         | true  | 오류로 실패합니다.
         | false | 무시하고 계속 진행합니다.

   --mailru-encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --mailru-pass
      비밀번호입니다.
      
      이는 앱 비밀번호여야 합니다. rclone은 일반 비밀번호로 작업하지 않습니다. 
      앱 비밀번호의 만드는 방법은 문서에서 구성 섹션에서 확인하십시오.
      

   --mailru-quirks
      쉼표로 구분된 내부 유지 보수 플래그 목록입니다.
      
      이 옵션은 보통 사용자가 사용해서는 안됩니다. 이는 백엔드 문제의 원격 문제 해결을 위해 만들어졌습니다. 
      플래그의 엄격한 의미는 문서화되지 않았으며 릴리스 사이에 유지되지 않을 수도 있습니다. 
      백엔드가 안정화될 때까지 quirks는 제거될 것입니다.
      지원되는 quirks: atomicmkdir binlist unknowndirs

   --mailru-speedup-enable
      데이터 해시가 동일한 다른 파일이 있는 경우 전체 업로드를 건너뛰시겠습니까?
      
      이 기능을 "speedup" 또는 "put by hash"라고 합니다. 
      도서, 비디오 또는 오디오 클립과 같이 일반적으로 제공되는 파일의 경우 특히 효율적입니다. 
      파일은 모든 mailru 사용자의 모든 계정에서 해시로 검색됩니다. 
      원본 파일이 고유하거나 암호화된 경우 의미가 없으며 효과적이지 않습니다. 
      또한, rclone은 전체 업로드가 필요한지 사전에 내부적으로 내용 해시를 계산하고 결정하기 위해 
      로컬 메모리와 디스크 공간이 필요할 수 있습니다. 
      또한, rclone이 파일 크기를 실제로 알지 못하는 경우 (스트리밍 또는 부분 업로드의 경우), 
      이 최적화를 시도하지 않습니다.

      예시:
         | true  | 사용
         | false | 사용하지 않음

   --mailru-speedup-file-patterns
      속도를 향상시키기(해시로 저장하기)에 적합한 파일 이름 패턴의 쉼표로 구분된 목록입니다.
      
      패턴은 대소문자를 구분하지 않으며 '*' 또는 '?' 메타 문자를 포함할 수 있습니다.

      예시:
         | <설정하지 않음>                         | 속도를 향상시키기(해시로 저장하기)를 완전히 비활성화합니다.
         | *                     | 모든 파일을 속도를 향상시키기(해시로 저장하기)를 시도합니다.
         | *.mkv,*.avi,*.mp4,*.mp3 | 일반적인 오디오/비디오 파일만 속도를 향상시키기(해시로 저장하기)를 시도합니다.
         | *.zip,*.gz,*.rar,*.pdf  | 일반적인 아카이브 또는 PDF 도서만 속도를 향상시키기(해시로 저장하기)를 시도합니다.

   --mailru-speedup-max-disk
      이 옵션을 사용하여 대형 파일에 대한 속도를 향상시키기(해시로 저장하기)를 사용하지 않을 수 있습니다.
      
      이유는 사전 해싱이 RAM 또는 디스크 공간을 고갈시킬 수 있기 때문입니다.

      예시:
         | 0  | 속도를 향상시키기(해시로 저장하기)를 완전히 사용하지 않음.
         | 1G | 1Gb보다 큰 파일은 직접 업로드됩니다.
         | 3G | 로컬 디스크에 3Gb보다 적은 공간이 있는 경우 이 옵션을 선택하십시오.

   --mailru-speedup-max-memory
      아래에 지정된 크기보다 큰 파일은 항상 디스크에서 해싱됩니다.

      예시:
         | 0    | 사전 해싱은 항상 임시 디스크 위치에서 수행됩니다.
         | 32M  | 사전 해싱에 32Mb RAM 이상을 사용하지 마십시오.
         | 256M | 해시 계산에 대해 최대 256Mb RAM이 사용 가능합니다.

   --mailru-user
      사용자 이름(보통 이메일)입니다.

   --mailru-user-agent
      클라이언트에서 내부적으로 사용되는 HTTP 사용자 에이전트입니다.
      
      기본값은 "rclone/VERSION" 또는 명령줄에서 제공된 "--user-agent"입니다.


OPTIONS:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보낸 후 데이터셋의 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공한 스캔으로부터 이 시간이 지나면 소스 디렉터리를 자동으로 재스캔합니다. (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비됨)

   Mail.ru 옵션

   --mailru-check-hash value             파일 체크섬이 일치하지 않거나 유효하지 않은 경우 복사 작업이 어떻게 진행되어야 합니까? (기본값: "true") [$MAILRU_CHECK_HASH]
   --mailru-encoding value               백엔드의 인코딩입니다. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$MAILRU_ENCODING]
   --mailru-pass value                   비밀번호입니다. [$MAILRU_PASS]
   --mailru-speedup-enable value         데이터 해시가 동일한 다른 파일이 있는 경우 전체 업로드를 건너뛰시겠습니까? (기본값: "true") [$MAILRU_SPEEDUP_ENABLE]
   --mailru-speedup-file-patterns value  속도를 향상시키기(해시로 저장하기)에 적합한 파일 이름 패턴의 쉼표로 구분된 목록입니다. (기본값: "*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf") [$MAILRU_SPEEDUP_FILE_PATTERNS]
   --mailru-speedup-max-disk value       이 옵션을 사용하여 대형 파일에 대한 속도를 향상시키기(해시로 저장하기)를 사용하지 않을 수 있습니다. (기본값: "3Gi") [$MAILRU_SPEEDUP_MAX_DISK]
   --mailru-speedup-max-memory value     아래에 지정된 크기보다 큰 파일은 항상 디스크에서 해싱됩니다. (기본값: "32Mi") [$MAILRU_SPEEDUP_MAX_MEMORY]
   --mailru-user value                   사용자 이름(보통 이메일)입니다. [$MAILRU_USER]

```
{% endcode %}