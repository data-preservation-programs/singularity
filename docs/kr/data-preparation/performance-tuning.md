# Singularity에서의 성능 조정

Singularity는 데이터 준비 성능을 최적화하기 위해 다양한 구성을 제공합니다. 이 가이드에서는 이러한 구성을 설명하고 효과적으로 조정하는 지침을 제공합니다.

## 인라인 준비
* **설명**: 인라인 준비는 CAR 파일을 저장하는 데 필요한 추가 디스크 공간이 없어집니다. 그러나 데이터베이스 조회와 저장에 약간의 오버헤드가 발생합니다.
* **영향**: 오버헤드는 일반적으로 무시할 수 있습니다. 그러나 많은 작은 파일이 포함된 데이터셋의 경우 중요해질 수 있습니다.
* **구성**: 비활성화하려면 `singularity prep create`와 함께 `--no-inline`을 사용하십시오.
* **더 읽어보기**: [인라인 준비](../topics/inline-preparation.md)

## DAG 업데이트
* **설명**:
  준비 과정에서 Singularity는 각 디렉토리에 대해 DAG와 CID를 새로 고칩니다. 이는 변경 사항을 실시간으로 추적하는 데 유용합니다.
* **영향**:
  CAR 파일이 준비될 때마다 디렉토리가 업데이트되므로 약간의 데이터베이스 오버헤드가 발생합니다.
* **구성**:
  비활성화하려면 `singularity prep create`와 함께 `--no-dag`를 사용하십시오.

## 데이터 준비의 병렬 처리

### 스캔
* **설명**: 스캔은 소스 저장소를 탐색하여 파일 목록을 만드는 작업입니다. 로컬 저장소에서는 빠르지만 S3와 같은 원격 저장소에서는 느릴 수 있습니다.
* **구성**:
  * **병렬 처리 활성화**: `singularity storage create` 또는 `singularity storage update`와 함께 `--client-scan-concurrency <숫자>`를 사용하십시오.
  * **참고**: 활성화하면 파일이 결정적인 순서로 처리되지 않을 수 있습니다.

### 패킹
* **설명**: 패킹은 여러 파일을 단일 CAR 파일로 병합하는 것으로, CPU와 IO 모두를 많이 사용하는 작업입니다. 네트워크 제약이 있는 원격 저장소의 경우, 병렬 처리를 늘리는 것이 유리합니다.
* **구성**:
    * **병렬 처리 조정**: `singularity run dataset-worker`와 함께 `--concurrency <숫자>`를 사용하십시오.

## 서버의 수정 시간 사용
* **설명**: `AWS S3`와 같은 일부 원격 저장소는 사용자 정의 `mtime` 및 서버 측의 최종 수정 시간을 제공합니다. 기본적으로 Singularity는 사용자 정의 `mtime`을 확인하고 사용할 수 있으면 사용합니다. 그렇지 않으면 서버의 최종 수정 시간을 사용합니다.
* **영향**: 사용자 정의 `mtime` 확인을 건너뛰고 직접 서버의 최종 수정 시간을 사용하면 원격 저장소로의 요청 수를 줄일 수 있습니다.
* **구성**: `singularity storage create` 또는 `singularity storage update`와 함께 `--client-use-server-mod-time`을 사용하여 서버의 시간을 우선시하고 객체 메타데이터 검색을 우회하십시오.

## 재시도 전략
### 네트워크 요청 재시도
* **설명**: 실패한 원격 폴더 목록 또는 파일 개방에 대해 Singularity는 RClone의 재시도 메커니즘을 활용합니다.
* **구성**: 재시도 횟수를 늘리려면 `singularity storage create` 또는 `singularity storage update`와 함께 `--client-low-level-retries <숫자>`를 사용하십시오.

### 네트워크 IO 재시도
* **설명**: 성공한 네트워크 요청에도 불구하고, 불안정한 네트워크 연결로 인해 네트워크 IO가 실패할 수 있습니다. Singularity는 재시도와 마지막 성공 지점부터 재개를 지원합니다.
* **구성**: 아래 플래그를 `singularity storage create` 또는 `singularity storage update`와 함께 사용하십시오.
```shell
 --client-retry-backoff value      # IO 읽기 오류를 재시도하는 데 대한 지연 백오프 (기본값: 1초)
 --client-retry-backoff-exp value  # IO 읽기 오류를 재시도하는 데 대한 지수적인 지연 백오프 (기본값: 1.0)
 --client-retry-delay value        # IO 읽기 오류를 재시도하기 전의 초기 지연 (기본값: 1초)
 --client-retry-max value          # IO 읽기 오류의 최대 재시도 횟수 (기본값: 10)
```

## 접근 불가능한 파일 건너뛰기
* **설명**: 권한 문제로 인해 원격 저장소의 일부 파일에 액세스할 수 없는 경우가 있습니다. 이러한 문제는 파일을 열 때만 나타나며, 이로 인해 패킹 작업이 실패할 수 있습니다.
* **구성**: `singularity storage create` 또는 `singularity storage update`와 함께 `--client-skip-inaccessible-files`를 사용하여 접근할 수 없는 파일을 건너뛰십시오.