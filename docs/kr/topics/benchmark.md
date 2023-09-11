# 싱귤래리티로 벤치마킹하기

싱귤래리티의 `ez-prep` 명령을 사용하면 벤치마킹을 간편하게 진행할 수 있습니다.

## 테스트 데이터 준비

먼저 벤치마킹을 위한 데이터를 생성해야 합니다. 여기에서는 디스크 IO 시간을 벤치마킹에서 제외하기 위해 sparse 파일을 사용합니다. 현재 싱귤래리티는 CID 중복 제거를 수행하지 않으므로, 이 파일들은 무작위 바이트로 처리됩니다.

```sh
mkdir dataset
truncate -s 1024G dataset/1T.bin
```

만약 벤치마킹에 디스크 IO 시간을 포함하려면, 다음 방법을 사용하여 랜덤 파일을 생성합니다:

```sh
dd if=/dev/urandom of=dataset/8G.bin bs=1M count=8192
```

## `ez-prep` 사용하기
`ez-prep` 명령은 최소한의 설정 옵션을 가지고 로컬 폴더에서 데이터를 준비하는 작업을 간소화합니다.

### 인라인 준비 벤치마킹
인라인 준비는 CAR 파일을 내보내는 대신에, 메타데이터를 데이터베이스에 직접 저장하여 수행합니다:

```sh
time singularity ez-prep --output-dir '' ./dataset
```

### 인메모리 데이터베이스를 사용한 벤치마킹

디스크 IO를 최소화하기 위해 인메모리 데이터베이스를 사용할 수 있습니다:

```sh
time singularity ez-prep --output-dir '' --database-file '' ./dataset
```

### 다중 워커를 사용한 벤치마킹

최적의 CPU 코어 활용을 위해 벤치마킹을 위한 동시성을 설정하세요. 참고: 각 워커는 약 4개의 CPU 코어를 사용합니다:

```sh
time singularity ez-prep --output-dir '' -j $(($(nproc) / 4 + 1)) ./dataset
```

## 결과 해석하기

일반적인 출력은 다음과 유사할 것입니다:

```
real    0m20.379s
user    0m44.937s
sys     0m8.981s
```

* `real`: 실제 경과 시간입니다. 더 많은 워커를 사용하면 이 시간이 줄어들 것입니다.
* `user`: 사용자 공간에서 사용된 CPU 시간입니다. `user`를 `real`로 나누면 사용된 CPU 코어의 근사치를 얻을 수 있습니다.
* `sys`: 커널 공간에서 사용된 CPU 시간(디스크 IO를 나타냄).

## 비교

다음은 랜덤 8G 파일을 기반으로 한 벤치마킹 결과입니다:

<table><thead><tr><th width="290">도구</th><th width="178.33333333333331" data-type="number">실행 시간 (초)</th><th data-type="number">CPU 시간 (초)</th><th data-type="number">메모리 (KB)</th></tr></thead><tbody><tr><td>인라인 준비를 사용한 싱귤래리티</td><td>15.66</td><td>51.82</td><td>99</td></tr><tr><td>인라인 준비를 사용하지 않은 싱귤래리티</td><td>19.13</td><td>51.51</td><td>99</td></tr><tr><td>go-fil-dataprep</td><td>16.39</td><td>43.94</td><td>83</td></tr><tr><td>generate-car</td><td>42.6</td><td>56.08</td><td>44</td></tr><tr><td>go-car + stream-commp</td><td>70.21</td><td>139.01</td><td>42</td></tr></tbody></table>