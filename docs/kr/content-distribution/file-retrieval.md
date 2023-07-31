# 파일 검색

Singularity 데이터베이스는 원본 데이터 소스에 대한 연결을 저장하기 때문에 파일 레벨의 검색도 제공할 수 있습니다. 먼저 컨텐츠 공급자를 실행하십시오.

```sh
singularitu run content-provider
```

HTTP 프로토콜을 사용하여 파일을 검색하십시오.

```
wget 127.0.0.1:8088/ipfs/bafyxxxx
```

Bitswap 프로토콜을 사용하여 파일을 검색하십시오.

```sh
ipfs daemon
ipfs swarm connect <singularity에서 표시된 multi_addr>
ipfs get bafyxxxx
```