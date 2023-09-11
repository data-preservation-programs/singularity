# Oracle Cloud Infrastructure Object Storage

{% code fullWidth="true" %}
```
名称：
   singularity存储创建oos - Oracle Cloud Infrastructure Object Storage

用法：
   singularity存储创建oos命令 [命令选项] [参数...]

命令：
   env_auth                     自动从运行时(environment)获取凭据，第一个提供的认证获胜
   instance_principal_auth      使用实例主体进行身份验证来授权实例进行API调用。
                                每个实例都有自己的标识，并使用从实例元数据中读取的证书进行身份验证。
                                https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm
   no_auth                      不需要凭据，通常用于读取公共存储桶
   resource_principal_auth      使用资源主体进行API调用
   user_principal_auth          使用OCI用户和API密钥进行身份验证。
                                您需要将租户OCID、用户OCID、区域、路径和API密钥的指纹放入配置文件中。
                                https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm
   help, h                      显示命令列表或一个命令的帮助信息

选项：
   --help, -h  显示帮助信息
```
{% endcode %}