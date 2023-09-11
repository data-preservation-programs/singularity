# 1Fichier

{% code fullWidth="true" %}
```
이름:
   singularity storage update fichier - 1Fichier

사용법:
   singularity storage update fichier [command options] <name|id>

설명:
   --api-key
      API Key를 입력하세요. https://1fichier.com/console/params.pl에서 얻을 수 있습니다.

   --shared-folder
      공유 폴더를 다운로드하려면 이 매개변수를 추가하세요.

   --file-password
      비밀번호로 보호된 공유 파일을 다운로드하려면 이 매개변수를 추가하세요.

   --folder-password
      비밀번호로 보호된 공유 폴더의 파일 목록을 조회하려면 이 매개변수를 추가하세요.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요](/overview/#encoding)의 [인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --api-key value  API Key를 입력하세요. https://1fichier.com/console/params.pl에서 얻을 수 있습니다. [$API_KEY]
   --help, -h       도움말 표시

   고급 옵션:

   --encoding value         백엔드의 인코딩입니다. (기본값: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --file-password value    비밀번호로 보호된 공유 파일을 다운로드하려면 이 매개변수를 추가하세요. [$FILE_PASSWORD]
   --folder-password value  비밀번호로 보호된 공유 폴더의 파일 목록을 조회하려면 이 매개변수를 추가하세요. [$FOLDER_PASSWORD]
   --shared-folder value    공유 폴더를 다운로드하려면 이 매개변수를 추가하세요. [$SHARED_FOLDER]
```
{% endcode %}