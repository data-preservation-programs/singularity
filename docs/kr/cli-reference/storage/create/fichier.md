# 1Fichier

{% code fullWidth="true" %}
```
NAME:
   singularity storage create fichier - 1Fichier

USAGE:
   singularity storage create fichier [command options] [arguments...]

DESCRIPTION:
   --api-key
      API 키, [여기](https://1fichier.com/console/params.pl)에서 얻을 수 있습니다.

   --shared-folder
      공유 폴더를 다운로드하려면 이 매개변수를 추가하십시오.

   --file-password
      암호로 보호된 공유 파일을 다운로드하려면 이 매개변수를 추가하십시오.

   --folder-password
      암호로 보호된 공유 폴더의 파일 목록을 보려면 이 매개변수를 추가하십시오.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


OPTIONS:
   --api-key value  API 키, [여기](https://1fichier.com/console/params.pl)에서 얻을 수 있습니다. [$API_KEY]
   --help, -h       도움말 표시

   고급 옵션

   --encoding value         백엔드의 인코딩입니다. (기본값: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --file-password value    암호로 보호된 공유 파일을 다운로드하려면 이 매개변수를 추가하십시오. [$FILE_PASSWORD]
   --folder-password value  암호로 보호된 공유 폴더의 파일 목록을 보려면 이 매개변수를 추가하십시오. [$FOLDER_PASSWORD]
   --shared-folder value    공유 폴더를 다운로드하려면 이 매개변수를 추가하십시오. [$SHARED_FOLDER]

   일반 옵션

   --name value  저장소의 이름 (기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}