# CAR 파일 배포하기

이제 CAR 파일을 저장 공급 업체에 배포하여 그들이 해당 파일을 자신의 서버로 가져올 수 있게 합니다. 먼저, 콘텐츠 제공자 서비스를 실행하고 준비한 데이터셋에 대한 Pieces를 다운로드합니다:

```sh
singularity run content-provider
wget 127.0.0.1:8088/piece/bagaxxxx
```

이전에 CAR를 내보내기 위해 출력 디렉토리를 지정한 경우 (인라인 준비를 비활성화하는 것) 해당 CAR 파일은 그 CAR 파일로부터 직접 제공됩니다. 그렇지 않은 경우, 인라인 준비를 사용하거나 그 CAR 파일을 실수로 삭제한 경우에는 원본 데이터 소스로부터 직접 제공됩니다.

## 다음 단계

[거래 준비 사항](../deal-making/deal-making-prerequisite.md)