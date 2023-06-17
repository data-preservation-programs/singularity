# 从元数据API下载CAR文件

{% code fullWidth="true" %}
```
命令名称：
   singularity download - 从元数据API下载CAR文件

用法：
   singularity download [命令选项] PIECE_CID

类别：
   工具

选项：
   --api value    元数据API的URL（默认值：“http://127.0.0.1:9090”）

   HTTP数据源

   --http-header value, -H value [ --http-header value, -H value ]    HTTP请求时要传输的http头（即key = value）。其值不应编码 [$ HTTP_HEADER]

   S3数据源

   --s3-access-key-id value      IAM访问密钥ID [$ AWS_ACCESS_KEY_ID]
   --s3-endpoint value           自定义S3端点[$ S3_ENDPOINT]
   --s3-region value             可用于AWS S3的S3区域[$ S3_REGION]
   --s3-secret-access-key value  IAM秘密访问密钥[$ AWS_SECRET_ACCESS_KEY]
```
{% endcode %}