# 로컬 경로에서 데이터셋 준비

{% code fullWidth="true" %}
```
이름:
   singularity ez-prep - 로컬 경로에서 데이터셋을 준비합니다.

사용법:
   singularity ez-prep [command options] <경로>

카테고리:
   쉬운 명령어

설명:
   이 명령어는 최소한의 설정 가능한 매개변수로 로컬 경로에서 데이터셋을 준비하는 데 사용할 수 있습니다.
   고급 사용법은 `dataset`과 `datasource` 하위 명령어를 사용해주세요.
   또한 이 명령어를 인메모리 데이터베이스와 인라인 준비를 위한 벤치마킹에도 사용할 수 있습니다.
   예를 들어:
     mkdir dataset
     truncate -s 1024G dataset/1T.bin
     singularity ez-prep --output-dir '' --database-file '' -j $(($(nproc) / 4 + 1)) ./dataset

옵션:
   --max-size value, -M value     생성되는 CAR 파일의 최대 크기 (기본값: "31.5GiB")
   --output-dir value, -o value   CAR 파일의 출력 디렉토리입니다. 인라인 준비를 위해 빈 문자열을 사용하세요. (기본값: "./cars")
   --concurrency value, -j value  패킹을 위한 동시성 수준입니다. (기본값: 1)
   --database-file value          메타데이터를 저장하기 위한 데이터베이스 파일입니다. 인메모리 데이터베이스를 사용하려면 빈 문자열을 사용하세요. (기본값: ./ezprep-<이름>.db)
   --help, -h                     도움말을 표시합니다.
```
{% endcode %}