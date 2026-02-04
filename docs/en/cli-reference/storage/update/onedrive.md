# Microsoft OneDrive

{% code fullWidth="true" %}
```
NAME:
   singularity storage update onedrive - Microsoft OneDrive

USAGE:
   singularity storage update onedrive [command options] <name|id>

DESCRIPTION:
   --client-id
      OAuth Client Id.
      
      Leave blank normally.

   --client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --token
      OAuth Access Token as a JSON blob.

   --auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --region
      Choose national cloud region for OneDrive.

      Examples:
         | global | Microsoft Cloud Global
         | us     | Microsoft Cloud for US Government
         | de     | Microsoft Cloud Germany
         | cn     | Azure and Office 365 operated by Vnet Group in China

   --chunk-size
      Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
      
      Above this size files will be chunked - must be multiple of 320k (327,680 bytes) and
      should not exceed 250M (262,144,000 bytes) else you may encounter \"Microsoft.SharePoint.Client.InvalidClientQueryException: The request message is too big.\"
      Note that the chunks will be buffered into memory.

   --drive-id
      The ID of the drive to use.

   --drive-type
      The type of the drive (personal | business | documentLibrary).

   --root-folder-id
      ID of the root folder.
      
      This isn't normally needed, but in special circumstances you might
      know the folder ID that you wish to access but not be able to get
      there through a path traversal.
      

   --access-scopes
      Set scopes to be requested by rclone.
      
      Choose or manually enter a custom space separated list with all scopes, that rclone should request.
      

      Examples:
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access | Read and write access to all resources
         | Files.Read Files.Read.All Sites.Read.All offline_access                                     | Read only access to all resources
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All offline_access                | Read and write access to all resources, without the ability to browse SharePoint sites. 
         |                                                                                             | Same as if disable_site_permission was set to true

   --disable-site-permission
      Disable the request for Sites.Read.All permission.
      
      If set to true, you will no longer be able to search for a SharePoint site when
      configuring drive ID, because rclone will not request Sites.Read.All permission.
      Set it to true if your organization didn't assign Sites.Read.All permission to the
      application, and your organization disallows users to consent app permission
      request on their own.

   --expose-onenote-files
      Set to make OneNote files show up in directory listings.
      
      By default, rclone will hide OneNote files in directory listings because
      operations like "Open" and "Update" won't work on them.  But this
      behaviour may also prevent you from deleting them.  If you want to
      delete OneNote files or otherwise want them to show up in directory
      listing, set this option.

   --server-side-across-configs
      Deprecated: use --server-side-across-configs instead.
      
      Allow server-side operations (e.g. copy) to work across different onedrive configs.
      
      This will work if you are copying between two OneDrive *Personal* drives AND the files to
      copy are already shared between them. Additionally, it should also function for a user who
      has access permissions both between Onedrive for *business* and *SharePoint* under the *same
      tenant*, and between *SharePoint* and another *SharePoint* under the *same tenant*. In other
      cases, rclone will fall back to normal copy (which will be slightly slower).

   --list-chunk
      Size of listing chunk.

   --no-versions
      Remove all versions on modifying operations.
      
      Onedrive for business creates versions when rclone uploads new files
      overwriting an existing one and when it sets the modification time.
      
      These versions take up space out of the quota.
      
      This flag checks for versions after file upload and setting
      modification time and removes all but the last version.
      
      **NB** Onedrive personal can't currently delete versions so don't use
      this flag there.
      

   --hard-delete
      Permanently delete files on removal.
      
      Normally files will get sent to the recycle bin on deletion. Setting
      this flag causes them to be permanently deleted. Use with care.
      
      OneDrive personal accounts do not support the permanentDelete API,
      it only applies to OneDrive for Business and SharePoint document libraries.
      

   --link-scope
      Set the scope of the links created by the link command.

      Examples:
         | anonymous    | Anyone with the link has access, without needing to sign in.
         |              | This may include people outside of your organization.
         |              | Anonymous link support may be disabled by an administrator.
         | organization | Anyone signed into your organization (tenant) can use the link to get access.
         |              | Only available in OneDrive for Business and SharePoint.

   --link-type
      Set the type of the links created by the link command.

      Examples:
         | view  | Creates a read-only link to the item.
         | edit  | Creates a read-write link to the item.
         | embed | Creates an embeddable link to the item.

   --link-password
      Set the password for links created by the link command.
      
      At the time of writing this only works with OneDrive personal paid accounts.
      

   --hash-type
      Specify the hash in use for the backend.
      
      This specifies the hash type in use. If set to "auto" it will use the
      default hash which is QuickXorHash.
      
      Before rclone 1.62 an SHA1 hash was used by default for Onedrive
      Personal. For 1.62 and later the default is to use a QuickXorHash for
      all onedrive types. If an SHA1 hash is desired then set this option
      accordingly.
      
      From July 2023 QuickXorHash will be the only available hash for
      both OneDrive for Business and OneDrive Personal.
      
      This can be set to "none" to not use any hashes.
      
      If the hash requested does not exist on the object, it will be
      returned as an empty string which is treated as a missing hash by
      rclone.
      

      Examples:
         | auto     | Rclone chooses the best hash
         | quickxor | QuickXor
         | sha1     | SHA1
         | sha256   | SHA256
         | crc32    | CRC32
         | none     | None - don't use any hashes

   --av-override
      Allows download of files the server thinks has a virus.
      
      The onedrive/sharepoint server may check files uploaded with an Anti
      Virus checker. If it detects any potential viruses or malware it will
      block download of the file.
      
      In this case you will see a message like this
      
          server reports this file is infected with a virus - use --onedrive-av-override to download anyway: Infected (name of virus): 403 Forbidden: 
      
      If you are 100% sure you want to download this file anyway then use
      the --onedrive-av-override flag, or av_override = true in the config
      file.
      

   --delta
      If set rclone will use delta listing to implement recursive listings.
      
      If this flag is set the onedrive backend will advertise `ListR`
      support for recursive listings.
      
      Setting this flag speeds up these things greatly:
      
          rclone lsf -R onedrive:
          rclone size onedrive:
          rclone rc vfs/refresh recursive=true
      
      **However** the delta listing API **only** works at the root of the
      drive. If you use it not at the root then it recurses from the root
      and discards all the data that is not under the directory you asked
      for. So it will be correct but may not be very efficient.
      
      This is why this flag is not set as the default.
      
      As a rule of thumb if nearly all of your data is under rclone's root
      directory (the `root/directory` in `onedrive:root/directory`) then
      using this flag will be be a big performance win. If your data is
      mostly not under the root then using this flag will be a big
      performance loss.
      
      It is recommended if you are mounting your onedrive at the root
      (or near the root when using crypt) and using rclone `rc vfs/refresh`.
      

   --metadata-permissions
      Control whether permissions should be read or written in metadata.
      
      Reading permissions metadata from files can be done quickly, but it
      isn't always desirable to set the permissions from the metadata.
      

      Examples:
         | off        | Do not read or write the value
         | read       | Read the value only
         | write      | Write the value only
         | read,write | Read and Write the value.
         | failok     | If writing fails log errors only, don't fail the transfer

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --description
      Description of the remote.


OPTIONS:
   --client-id value      OAuth Client Id. [$CLIENT_ID]
   --client-secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help
   --region value         Choose national cloud region for OneDrive. (default: "global") [$REGION]

   Advanced

   --access-scopes value         Set scopes to be requested by rclone. (default: "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access") [$ACCESS_SCOPES]
   --auth-url value              Auth server URL. [$AUTH_URL]
   --av-override                 Allows download of files the server thinks has a virus. (default: false) [$AV_OVERRIDE]
   --chunk-size value            Chunk size to upload files with - must be multiple of 320k (327,680 bytes). (default: "10Mi") [$CHUNK_SIZE]
   --delta                       If set rclone will use delta listing to implement recursive listings. (default: false) [$DELTA]
   --description value           Description of the remote. [$DESCRIPTION]
   --disable-site-permission     Disable the request for Sites.Read.All permission. (default: false) [$DISABLE_SITE_PERMISSION]
   --drive-id value              The ID of the drive to use. [$DRIVE_ID]
   --drive-type value            The type of the drive (personal | business | documentLibrary). [$DRIVE_TYPE]
   --encoding value              The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --expose-onenote-files        Set to make OneNote files show up in directory listings. (default: false) [$EXPOSE_ONENOTE_FILES]
   --hard-delete                 Permanently delete files on removal. (default: false) [$HARD_DELETE]
   --hash-type value             Specify the hash in use for the backend. (default: "auto") [$HASH_TYPE]
   --link-password value         Set the password for links created by the link command. [$LINK_PASSWORD]
   --link-scope value            Set the scope of the links created by the link command. (default: "anonymous") [$LINK_SCOPE]
   --link-type value             Set the type of the links created by the link command. (default: "view") [$LINK_TYPE]
   --list-chunk value            Size of listing chunk. (default: 1000) [$LIST_CHUNK]
   --metadata-permissions value  Control whether permissions should be read or written in metadata. (default: "off") [$METADATA_PERMISSIONS]
   --no-versions                 Remove all versions on modifying operations. (default: false) [$NO_VERSIONS]
   --root-folder-id value        ID of the root folder. [$ROOT_FOLDER_ID]
   --server-side-across-configs  Deprecated: use --server-side-across-configs instead. (default: false) [$SERVER_SIDE_ACROSS_CONFIGS]
   --token value                 OAuth Access Token as a JSON blob. [$TOKEN]
   --token-url value             Token server url. [$TOKEN_URL]

   Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers. To remove, use empty string.
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth. To remove, use empty string.
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value). This will replace the existing header values. To remove a header, use --http-header "key="". To remove all headers, use --http-header ""
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth. To remove, use empty string.
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-scan-concurrency value                  Max number of concurrent listing requests when scanning data source (default: 1)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-use-server-mod-time                     Use server modified time if possible (default: false)
   --client-user-agent value                        Set the user-agent to a specified string. To remove, use empty string. (default: rclone default)

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
