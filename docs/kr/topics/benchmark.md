# 벤치마크

EZ 준비 명령어는 벤치마크를 수행하는 간단한 방법을 제공합니다.

## 테스트 데이터 준비

먼저, 벤치마크를 위해 일부 데이터를 준비하세요. 디스크 IO 시간을 고려하지 않기 위해 희소 파일을 사용합니다. 현재 CID 중복 제거를 수행하지 않기 때문에, 싱귤래리티는 해당 바이트를 무작위 바이트로 처리합니다.

```sh
mkdir dataset
truncate -s 1024G dataset/1T.bin
```

벤치마크에 디스크 IO 시간을 포함하고 싶다면, 임의의 파일을 생성하기 위해 원하는 방법을 사용할 수 있습니다. 예를 들어,

```
dd if=/dev/urandom of=dataset/8G.bin bs=1M count=8192
```

## ez-prep 실행

EZ prep 명령어는 매우 적은 커스터마이징 설정을 사용하여 로컬 폴더를 준비하기 위해 몇 가지 내부 명령을 실행하는 간단한 명령입니다.

#### 인라인 준비로 벤치마크하기&#x20;

인라인 준비는 CAR 파일을 내보내지 않고 필요한 메타데이터를 데이터베이스에 저장하는 필요성을 없애줍니다.

```sh
time singularity ez-prep --output-dir '' ./dataset
```

#### 메모리 내 데이터베이스로 벤치마크하기

디스크 IO를 추가로 줄이기 위해 메모리 내 데이터베이스를 사용할 수도 있습니다.

```sh
time singularity ez-prep --output-dir '' --database-file '' ./dataset
```

#### 여러 워커로 벤치마크하기

모든 CPU 코어를 활용하기 위해 벤치마크에 병렬성 플래그를 설정할 수 있습니다. 각 워커는 약 4개의 CPU 코어를 사용하므로 올바르게 설정해야합니다.

```sh
time singularity ez-prep --output-dir '' -j $(($(nproc) / 4 + 1)) ./dataset
```

## 결과 해석

아래와 같은 결과를 볼 수 있습니다.

```
real    0m20.379s
user    0m44.937s
sys     0m8.981s
```

`real`은 실제 시계 시간을 의미합니다. 더 많은 워커 병렬성을 사용하면 이 숫자가 감소할 가능성이 높습니다.

`user`는 사용자 공간에서 사용된 CPU 시간을 의미합니다. `user`를 `real`로 나누면 대략적으로 프로그램에서 사용한 CPU 코어 수를 나타냅니다. 작업이 변경되지 않기 때문에 더 많은 병렬성은이 숫자에 큰 영향을 미치지 않을 가능성이 높습니다.

`sys`는 디스크 IO에 소요된 커널 공간에서 사용된 CPU 시간을 의미합니다.

## 비교

아래 테스트는 임의의 8G 파일에서 수행되었습니다.

<table><thead><tr><th width="290">도구</th><th width="178.33333333333331" data-type="number">클록 시간 (sec)</th><th data-type="number">CPU 시간 (sec)</th><th data-type="number">메모리 (KB)</th></tr></thead><tbody><tr><td>Singularity with 인라인 준비</td><td>15.66</td><td>51.82</td><td>99</td></tr><tr><td>Singularity without 인라인 준비</td><td>19.13</td><td>51.51</td><td>99</td></tr><tr><td>go-fil-dataprep</td><td>16.39</td><td>43.94</td><td>83</td></tr><tr><td>generate-car</td><td>42.6</td><td>56.08</td><td>44</td></tr><tr><td>go-car + stream-commp</td><td>70.21</td><td>139.01</td><td>42</td></tr></tbody></table>