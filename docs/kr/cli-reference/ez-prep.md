# 로컬 경로에서 데이터셋 준비하기

{% code fullWidth="true" %}
```
NAME:
   singularity ez-prep - 로컬 경로에서 데이터셋 준비하기

사용법:
   singularity ez-prep [command options] <경로>

카테고리:
   유틸리티

설명:
   이 명령은 최소 구성 가능한 매개변수를 사용하여 로컬 경로에서 데이터셋을 준비하는 데 사용할 수 있습니다.
   더 고급 사용법을 위해, `storage` 및 `data-prep` 하위 명령어를 사용하십시오.
   인메모리 데이터베이스와 인라인 준비를 사용하여 벤치마킹하는 데에도 이 명령을 사용할 수 있습니다.
     mkdir dataset
     truncate -s 1024G dataset/1T.bin
     singularity ez-prep --output-dir '' --database-file '' -j $(($(nproc) / 4 + 1)) ./dataset

옵션:
   --max-size value, -M value       CAR 파일로 생성될 최대 크기 (기본값: "31.5GiB")
   --output-dir value, -o value     CAR 파일의 출력 디렉터리. 인라인 준비를 위해 빈 문자열을 사용하십시오 (기본값: "./cars")
   --concurrency value, -j value    패킹을 위한 병렬 처리 수 (기본값: 1)
   --database-file value, -f value  메타데이터를 저장할 데이터베이스 파일. 인메모리 데이터베이스를 사용하려면 빈 문자열을 사용하십시오 (기본값: ./ezprep-<name>.db)
   --help, -h                       도움말 표시
```
{% endcode %}