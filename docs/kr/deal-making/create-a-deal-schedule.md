# 거래 일정 생성하기

저장 공급 업체와 거래를 시작해 보십시오. 먼저 거래 푸시 서비스를 실행합니다.

```
singularity run deal-pusher
```

## 한 번에 모든 거래 전송하기

더 작은 데이터세트의 경우, 모든 거래를 한꺼번에 저장 공급 업체에 전송할 수 있습니다. 이를 위해서는 다음 명령을 사용할 수 있습니다.

```sh
singularity deal schedule create <preparation> <provider_id>
```

하지만 데이터세트가 큰 경우, 거래 제안이 만료되기 전에 그만큼 많은 거래를 저장 공급 업체가 처리하지 못할 수도 있으므로 일정을 만들 수 있습니다.

## 일정을 갖고 거래 전송하기

같은 명령을 사용하여 얼마나 빠르고 얼마나 자주 거래를 저장 공급 업체에 전송할지를 제어하는 자체 일정을 만들 수 있습니다.
```sh
singularity deal schedule create -h
```