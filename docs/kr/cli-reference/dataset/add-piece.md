# 거래를 위해 데이터셋에 조각 (CAR 파일)을 수동으로 등록하기

{% code fullWidth="true" %}
```

이름:
   singularity dataset add-piece - 거래를 위해 데이터셋에 조각 (CAR 파일)을 수동으로 등록하기

사용법:
   singularity dataset add-piece [command options] <dataset_name> <piece_cid> <piece_size>

설명:
   이미 CAR 파일이 있는 경우:
     singularity dataset add-piece -p <path_to_car_file> <dataset_name> <piece_cid> <piece_size>

   CAR 파일이 없지만 RootCID를 알고 있는 경우:
     singularity dataset add-piece -r <root_cid> <dataset_name> <piece_cid> <piece_size>

   둘 다 없는 경우:
     singularity dataset add-piece -r <root_cid> <dataset_name> <piece_cid> <piece_size>
   그러나 이 경우에는 거래를 수행할 때 rootCID가 올바르게 설정되지 않으므로 검색 테스트와 잘 작동하지 않을 수 있습니다.

옵션:
   --file-path value, -p value  CAR 파일의 경로. 파일 크기와 root CID를 결정하는 데 사용됩니다.
   --root-cid value, -r value   CAR 파일의 Root CID입니다. 제공되지 않으면 CAR 파일 헤더에 의해 결정됩니다. 거래의 label 필드를 채우는 데 사용됩니다.
   --help, -h                   도움말 표시
```
{% endcode %}