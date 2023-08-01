# 암호화

## 개요

Singularity는 제공된 수신자(공개 키) 또는 YubiKey와 같은 하드웨어 PIV 토큰과 함께 파일을 암호화하는 내장 암호화 솔루션을 지원합니다. 사용자 정의 암호화 스크립트를 제공하여 외부 암호화 솔루션과 통합할 수도 있습니다 \[테스트 필요].

## 내장 암호화

비대칭 암호화를 위해 공개-비공개 키 쌍을 생성하여 시작합니다. Singularity에서 사용하는 기본 암호화 라이브러리는 [age](https://github.com/FiloSottile/age)입니다.

```sh
go install filippo.io/age/cmd/...@latest
age-keygen -o key.txt
> Public key: agexxxxxxxxxxxx
```

이제 생성된 공개 키를 사용하여 각 파일을 암호화하기 위해 데이터셋을 설정할 수 있습니다.

```sh
singularity dataset create --encryption-recipient agexxxxxxxxxxxx \
  --output-dir . test
```

인라인 준비는 비활성화되어 있습니다. 인라인 중복 암호화는 초기 암호화 프로세스 중 도입된 초기 무작위성으로 인해 다른 암호화된 내용이 생성되기 때문입니다.

이제 데이터 소스를 추가하여 데이터 준비 과정을 계속할 수 있습니다. 단, 폴더 구조는 암호화되지 않으므로 폴더 구조의 DAG를 생성하거나 `daggen` 명령을 실행하지 않을 수 있습니다. 후자의 경우 폴더 구조는 Singularity 데이터베이스 및 명령어에서만 액세스할 수 있습니다.

## 사용자 정의 암호화

Singularity는 사용자 정의 스크립트를 제공하여 파일 스트림을 암호화하는 사용자 정의 암호화도 지원합니다. 이는 키 관리 서비스 및 사용자 정의 암호화 알고리즘 또는 도구와 함께 사용될 수 있습니다.