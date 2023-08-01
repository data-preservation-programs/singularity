---
description: 데이터베이스를 초기화하고 새로운 데이터셋을 생성하는 것으로 시작하세요.
---

# 데이터셋 생성하기

## 데이터베이스 초기화하기

기본적으로, `sqlite3` 데이터베이스 백엔드를 사용하며, 데이터베이스 파일은 `$HOME/.singularity`에 초기화됩니다.

별도의 프로덕션용 데이터베이스 백엔드를 사용하려면, [deploy-to-production.md](../installation/deploy-to-production.md "mention")을 확인하세요.

```sh
singularity admin init
```

## 새로운 데이터셋 생성하기

데이터셋은 단일 데이터셋과 관련된 데이터 소스의 모음입니다. 데이터셋을 생성한 후에는 데이터 소스를 추가하고 Filecoin 지갑 주소와 연결할 수 있습니다.

```sh
singularity dataset create my_dataset
```

기본적으로, singularity는 인라인 준비 기술을 사용하여 CAR 파일로 내보내지 않습니다. 이는 대부분의 데이터 소스에 대해 데이터가 변경되지 않으며, CAR 파일이 원래 데이터 소스와 동일한 내용을 저장하기 때문입니다.&#x20;

## 다음 단계

[add-a-data-source.md](add-a-data-source.md "mention")