---
description: 개발자들이 Singularity에서 데이터 준비 과정을 이해하기를 원하는 경우, 이 문서는 기술적인 개요를 제공합니다.
---

# Singularity에서 데이터 준비의 아키텍처

![Singularity 데이터 준비 모델](data-prep-model.jpg)

# 데이터셋과 소스

데이터셋은 복제에 대한 공통 정책과 하나 이상의 Filecoin 지갑을 가진 데이터의 집합입니다.

각 데이터셋은 하나 이상의 데이터 소스를 가질 수 있습니다. 소스는 단순히 [RClone](https://github.com/rclone/rclone)에서 지원하는 로컬 파일 저장소를 포함한 스토리지에 저장된 데이터 폴더를 가리킵니다.

데이터 준비는 Singularity 사용자가 다음과 같은 작업을 수행할 때 시작됩니다:
1. 데이터셋 생성
2. 데이터셋에 소스 추가
3. 데이터 준비 워커 시작

# 스캔

데이터 준비의 첫 번째 단계는 데이터 소스를 스캔하여 해당 소스에 포함된 파일과 폴더의 디렉토리 트리 구조와 파일을 CAR 파일로 나누는 계획을 작성하는 것입니다.

데이터 소스의 디렉토리 트리 구조를 나타내는 모델은 Directory와 Item입니다. (디렉토리는 폴더를 나타내고, Item은 파일을 나타냄) 스캔 과정에서 RClone은 데이터 소스의 디렉토리 구조를 매핑하고, 이 디렉토리 구조를 기반으로 Directory와 Item 모델이 구성됩니다. 디렉토리 구조는 데이터 소스가 등록되는 시점에 생성된 루트 디렉토리부터 시작하여 구성됩니다.

데이터 소스의 각 Item에 대해 스캔 과정은 ItemPart를 생성하기도 합니다. ItemPart는 파일의 연속된 부분을 나타내는데, 이 크기는 1GB까지일 수 있습니다 (이 크기는 구성 가능합니다). ItemPart의 집합은 Chunk에 추가되는데, Chunk는 단일 CAR 파일에 저장하기에 충분한 ItemPart의 목록입니다. 이 단일 CAR 파일은 Filecoin의 32GB 조각에 저장할만큼 충분한 크기입니다.

스캔 과정의 마지막에는 Directories의 트리, 각각에는 Items의 목록이 있으며, 각각의 Item은 파일의 일부를 나타내는 ItemPart로 분할됩니다 (최대 1GB까지). 또한, ItemPart의 집합인 Chunks도 있습니다. 각 Chunk는 단일 CAR 파일에 모아질 ItemParts의 집합에 링크됩니다.

데이터 소스가 변경될 수 있으므로 파일과 폴더를 추가, 변경, 삭제할 때마다 다시 스캔할 수 있음에 유의하세요.

# 패킹

소스가 스캔되면 CAR 파일로 패킹할 준비가됩니다. 패킹은 Chunks를 실제로 작성된 개별 블록을 가진 CAR 파일로 변환하는 과정입니다.

Car 파일을 패킹하기 위해 Chunk의 각 ItemPart는 읽혀지고, 지정된 블록 크기로 IPLD Raw 블록으로 청크되어 CAR에 작성됩니다. 모든 Raw 블록이 작성된 후, ItemPart에 두 개 이상의 원본 블록이 포함되었다면 UnixFS 중간 노드 블록의 트리가 조립되어 기록되고, ItemPart에 대한 루트 CID가 생성됩니다. 이 과정이 완료되면, Chunk의 모든 ItemPart에 대한 Raw 블록과 UnixFS 중간 노드 블록이 포함된 CAR 파일이 생성됩니다.

패킹 과정이 끝나면, Singularity는 데이터베이스에 Car 파일을 나타내는 Car 모델과 CAR의 각 블록에 대한 CarBlock도 작성합니다.

각 Car를 작성하는 과정에서, Directories와 Items로 돌아갑니다. 모든 ItemPart가 작성된 각 Item에 대해, 해당 Item의 모든 ItemPart를 하나의 UnixFS 파일로 연결하는 또 다른 UnixFS 중간 노드 트리를 작성합니다. 또한, 각 Directory에 대한 UnixFS 디렉토리 노드를 조립하고 업데이트합니다. 이 데이터는 일시적으로 데이터베이스에 저장되며, Directory 객체에 연결됩니다.

패킹 과정이 완료되면, Car 파일에 데이터 소스의 모든 ItemPart를 저장할 수 있습니다. 그러나, 이 시점에서 데이터 소스의 디렉토리 (Directories) 및 파일 (Items)을 나타내는 UnixFS DAG를 직렬화하여 Filecoin에 저장하지는 않았습니다.

# Daggen

파일 및 디렉토리 구조는 시간이 지남에 따라 변경될 수 있으므로, 이 구조의 스냅샷을 Filecoin에 저장하기 위해 Daggen이라는 수동 준비 단계가 있습니다. 사용자가 이 단계를 별도로 시작하면, 패킹 중에 조립된 UnixFS DAG 트리는 CAR로 직렬화되어 Filecoin에 저장됩니다. 이 과정이 완료되면, 데이터 준비 과정에서 작성된 모든 CAR를 Filecoin에 저장한다면, 데이터 소스의 전체 스냅샷을 Filecoin에서 검색하기 위해 필요한 모든 것을 저장할 수 있습니다.