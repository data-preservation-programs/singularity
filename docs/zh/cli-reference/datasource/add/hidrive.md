# HiDrive

{% code fullWidth="true" %}
```
名称：
    singularity datasource add hidrive - HiDrive

用法：
    singularity datasource add hidrive [命令选项] <数据集名称> <源路径>

说明：
    --hidrive-token-url
        Token 服务器 URL。
        如需使用提供程序默认值，请留空。

    --hidrive-endpoint
        服务终端点。
        这是 API 调用将要执行的 URL。

    --hidrive-auth-url
        Auth 服务器 URL。
        如需使用提供程序默认值，请留空。

    --hidrive-scope-access
        rclone 请求 HiDrive 访问权限时应使用的访问权限。
        示例：
            | rw | 对资源的读取和写入权限。
            | ro | 只读权限。

    --hidrive-scope-role
        rclone 请求 HiDrive 时应使用的用户级别。
        示例：
            | user  | 管理权限的用户级别访问。
                    | 在大多数情况下，这将足够。
            | admin | 管理权限的广泛访问。
            | owner | 管理权限的完全访问。

    --hidrive-upload-concurrency
        分块上传的并发数。
        这是同时运行相同文件的传输的上限。
        将此设置为大于小于 1 的值将导致上传死锁。
        
        如果您正在通过高速链接上传较少数量的大型文件，
        并且这些上传未充分利用您的带宽，则增加此值可能有助于加快传输速度。

    --hidrive-token
        OAuth 访问令牌作为 JSON 二进制大对象。

    --hidrive-disable-fetching-member-count
        不要获取目录中对象的数量，除非绝对必要。
        如果不获取子目录中对象的数量，则请求可能更快。

    --hidrive-upload-cutoff
        分块上传的阈值/截止值。
        任何文件大于此值都将分段上传到配置的块大小。
        
        该值的上限为 2,147,483,647 字节（约 2.000Gi）。
        这是单个上传操作支持的最大字节数。
        如果将此设置为上限以上，将导致上传失败。

    --hidrive-client-id
        OAuth 客户 ID。
        正常情况下请留空。

    --hidrive-root-prefix
        所有路径的根/父文件夹。
        填写以使用指定文件夹作为所有传递给远程的路径的父文件夹。
        这样，rclone 就可以使用任何文件夹作为起点。

        示例：
            | /       | rclone 可访问的最高目录。
                      | 如果 rclone 使用常规 HiDrive 用户帐户，这将相当于 "root"。
            | root    | HiDrive 用户帐户的最顶层目录
            | <unset> | 这表示您的路径没有根前缀。
                      | 在使用此选项时，您始终需要为此远程路径指定具有有效父级的路径，例如 "remote:/path/to/dir" 或 "remote:root/path/to/dir"。

    --hidrive-chunk-size
        分块上传的块大小。
        任何文件大于配置的截止值（或未知大小的文件）都将按此大小分块上传。
        
        该值的上限为 2,147,483,647 字节（约 2.000Gi）。
        这是单个上传操作支持的最大字节数。
        如果将此设置为上限以上或负值，将导致上传失败。
        
        将此设置为较大的值可能会增加上传速度，但代价是使用更多的内存。
        它可以设置为较小的值以节省内存。

    --hidrive-encoding
        后端的编码方式。
        获取更多信息，请参阅概述中的 [编码方式部分](/overview/#encoding)。

    --hidrive-client-secret
        OAuth 客户端密码。
        正常情况下请留空。


选项：
    --help, -h  显示帮助

    数据准备选项

    --delete-after-export    [危险] 导出数据集为 CAR 文件后删除文件。
                             (默认值：false)
    --rescan-interval value  上次成功扫描后，当此时间间隔经过时自动重新扫描源目录（默认值：已禁用）

    hidrive 选项

    --hidrive-auth-url value                       Auth 服务器 URL。[$HIDRIVE_AUTH_URL]
    --hidrive-chunk-size value                     分块上传的块大小。 (默认值："48Mi") [$HIDRIVE_CHUNK_SIZE]
    --hidrive-client-id value                      OAuth 客户 ID。[$HIDRIVE_CLIENT_ID]
    --hidrive-client-secret value                  OAuth 客户端密码。[$HIDRIVE_CLIENT_SECRET]
    --hidrive-disable-fetching-member-count value  不要获取目录中对象的数量，除非绝对必要。 (默认值："false") [$HIDRIVE_DISABLE_FETCHING_MEMBER_COUNT]
    --hidrive-encoding value                       后端的编码方式。 (默认值："Slash,Dot") [$HIDRIVE_ENCODING]
    --hidrive-endpoint value                       服务终端点。 (默认值："https://api.hidrive.strato.com/2.1") [$HIDRIVE_ENDPOINT]
    --hidrive-root-prefix value                    所有路径的根/父文件夹。 (默认值："/") [$HIDRIVE_ROOT_PREFIX]
    --hidrive-scope-access value                   rclone 请求 HiDrive 访问权限时应使用的访问权限。 (默认值："rw") [$HIDRIVE_SCOPE_ACCESS]
    --hidrive-scope-role value                     rclone 请求 HiDrive 时应使用的用户级别。 (默认值："user") [$HIDRIVE_SCOPE_ROLE]
    --hidrive-token value                          OAuth 访问令牌作为 JSON 二进制大对象。[$HIDRIVE_TOKEN]
    --hidrive-token-url value                      Token 服务器 URL。[$HIDRIVE_TOKEN_URL]
    --hidrive-upload-concurrency value             分块上传的并发数。 (默认值："4") [$HIDRIVE_UPLOAD_CONCURRENCY]
    --hidrive-upload-cutoff value                  分块上传的阈值/截止值。 (默认值："96Mi") [$HIDRIVE_UPLOAD_CUTOFF]

```
{% endcode %}