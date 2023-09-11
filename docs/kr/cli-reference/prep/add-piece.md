# 준비 작업에 직접적으로 조각 정보를 추가합니다. 이는 외부 도구로 준비된 조각에 유용합니다.

{% code fullWidth="true" %}
```
NAME:
   singularity prep add-piece - 준비 작업에 직접적으로 조각 정보를 추가합니다. 이는 외부 도구로 준비된 조각에 유용합니다.

사용법:
   singularity prep add-piece [command options] <preparation id|name>

카테고리:
   조각 관리

옵션:
   --piece-cid value   조각의 CID
   --piece-size value  조각의 크기 (기본값: "32GiB")
   --file-path value   CAR 파일의 경로, 파일 크기와 루트 CID를 결정하는 데 사용됩니다.
   --root-cid value    CAR 파일의 루트 CID
   --help, -h          도움말 표시
```
{% endcode %}