# 거래를 위한 사전 준비

## 스토리지 프로바이더 찾기

현재 Singularity는 거래를 허용하는 스토리지 프로바이더를 찾는 데 도움을 주지 않습니다. 우수한 품질의 스토리지 프로바이더를 찾기 위해 다음과 같은 자원을 사용할 수 있습니다. 

* TODO

## Filecoin 지갑 생성

거래를 진행하기 전에 Filecoin 지갑을 생성해야 합니다. 분산원장 지갑이나 거래소 지갑을 사용할 수 없습니다. Filecoin 지갑을 생성하려면 다음 명령을 실행하십시오.

```sh
singularity wallet create
```

이렇게 하면 지갑 주소와 해당 지갑과 연결된 비밀 키가 생성됩니다. 이 지갑은 아직 블록체인에 인식되지 않기 때문에 거래에 사용할 수 없습니다. 이제 다른 사용자들도 알게 되도록 이 지갑에 0 FIL을 이체하는 것이 좋습니다.&#x20;

이 지갑이 체인 상에 기록되면 위의 명령이 완료되며 거래를 진행할 준비가 됩니다.

또는 이미 기존 지갑이 있는 경우 다음을 사용하여 가져올 수 있습니다.

```sh
singularity wallet import xxx
```

## \[선택 사항] [데이터 용량](https://docs.filecoin.io/basics/how-storage-works/filecoin-plus/#datacap) 가져오기

현재 시장 상황에서 대부분의 스토리지 프로바이더는 일반 거래 대신 [확인된 거래](https://docs.filecoin.io/storage-provider/filecoin-deals/verified-deals/)를 선호합니다. 데이터 세트가 수십 테라바이트 이상인 경우 [Filplus 지방 정부 팀과 감사원](https://github.com/filecoin-project/notary-governance)에서 데이터 용량을 신청하는 것이 가장 좋습니다.

## 다음 단계

[거래 일정 생성](create-a-deal-schedule.md)