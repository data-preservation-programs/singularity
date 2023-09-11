# Oracle Cloud Infrastructure 对象存储

{% code fullWidth="true" %}
```
名称:
   singularity storage update oos - Oracle Cloud Infrastructure 对象存储

用法:
   singularity storage update oos 命令 [命令选项] [参数...]

命令:
   env_auth                        自动从运行时（环境）中获取凭据，第一个提供凭据的将获胜
   instance_principal_auth         使用实例 Principal 授权实例进行 API 调用。
                                   每个实例都有自己的身份，并使用从实例元数据中读取的证书进行身份验证。
                                   https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm
   no_auth                         不需要凭据，通常用于读取公共存储桶
   resource_principal_auth         使用资源 Principal 进行 API 调用
   user_principal_auth             使用 OCI 用户和 API 密钥进行身份验证。
                                   您需要在配置文件中放置您的租户 OCID、用户 OCID、区域、路径和 API 密钥的指纹。
                                   https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm
   help, h                         显示命令列表或某个命令的帮助信息

选项:
   --help, -h  显示帮助信息
```
{% endcode %}