# Microsoft OneDrive

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add onedrive - Microsoft OneDrive

USAGE:
   singularity datasource add onedrive [command options] <dataset_name> <source_path>

DESCRIPTION:
   --onedrive-access-scopes
      Set scopes to be requested by rclone.

      Choose or manually enter a custom space separated list with all scopes, that rclone should request.


      Examples:
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access | Read and write access to all resources
         | Files.Read Files.Read.All Sites.Read.All offline_access                                     | Read only access to all resources
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All offline_access                | Read and write access to all resources, without the ability to browse SharePoint sites.
                                                                                                       | Same as if disable_site_permission was set to true

   --onedrive-server-side-across-configs
      Allow server-side operations (e.g. copy) to work across different onedrive configs.

      This will only work if you are copying between two OneDrive *Personal* drives AND
      the files to copy are already shared between them.  In other cases, rclone will
      fall back to normal copy (which will be slightly slower).

   --onedrive-list-chunk
      Size of listing chunk.

   --onedrive-link-scope
      Set the scope of the links created by the link command.

      Examples:
         | anonymous    | Anyone with the link has access, without needing to sign in.
                        | This may include people outside of your organization.
                        | Anonymous link support may be disabled by an administrator.
         | organization | Anyone signed into your organization (tenant) can use the link to get access.
                        | Only available in OneDrive for Business and SharePoint.

   --onedrive-hash-type
      Specify the hash in use for the backend.

      This specifies the hash type in use. If set to "auto" it will use the
      default hash which is is QuickXorHash.

      Before rclone 1.62 an SHA1 hash was used by default for Onedrive
      Personal. For 1.62 and later the default is to use a QuickXorHash for
      all onedrive types. If an SHA1 hash is desired then set this option
      accordingly.

      From July 2023 QuickXorHash will be the only available hash for
      both OneDrive for Business and OneDriver Personal.

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

   --onedrive-region
      Choose national cloud region for OneDrive.

      Examples:
         | global | Microsoft Cloud Global
         | us     | Microsoft Cloud for US Government
         | de     | Microsoft Cloud Germany
         | cn     | Azure and Office 365 operated by Vnet Group in China

   --onedrive-chunk-size
      Chunk size to upload files with - must be multiple of 320k (327,680 bytes).

      Above this size files will be chunked - must be multiple of 320k (327,680 bytes) and
      should not exceed 250M (262,144,000 bytes) else you may encounter \"Microsoft.SharePoint.Client.InvalidClientQueryException: The request message is too big.\"
      Note that the chunks will be buffered into memory.

   --onedrive-link-type
      Set the type of the links created by the link command.

      Examples:
         | view  | Creates a read-only link to the item.
         | edit  | Creates a read-write link to the item.
         | embed | Creates an embeddable link to the item.

   --onedrive-token
      OAuth Access Token as a JSON blob.

   --onedrive-drive-id
      The ID of the drive to use.

   --onedrive-root-folder-id
      ID of the root folder.

      This isn't normally needed, but in special circumstances you might
      know the folder ID that you wish to access but not be able to get
      there through a path traversal.


   --onedrive-no-versions
      Remove all versions on modifying operations.

      Onedrive for business creates versions when rclone uploads new files
      overwriting an existing one and when it sets the modification time.

      These versions take up space out of the quota.

      This flag checks for versions after file upload and setting
      modification time and removes all but the last version.

      **NB** Onedrive personal can't currently delete versions so don't use
      this flag there.


   --onedrive-encoding
      The encoding for the backend.

      See the [encoding section in the overview](/overview/#encoding) for more info.

   --onedrive-client-id
      OAuth Client Id.

      Leave blank normally.

   --onedrive-auth-url
      Auth server URL.

      Leave blank to use the provider defaults.

   --onedrive-token-url
      Token server url.

      Leave blank to use the provider defaults.

   --onedrive-drive-type
      The type of the drive (personal | business | documentLibrary).

   --onedrive-disable-site-permission
      Disable the request for Sites.Read.All permission.

      If set to true, you will no longer be able to search for a SharePoint site when
      configuring drive ID, because rclone will not request Sites.Read.All permission.
      Set it to true if your organization didn't assign Sites.Read.All permission to the
      application, and your organization disallows users to consent app permission
      request on their own.

   --onedrive-expose-onenote-files
      Set to make OneNote files show up in directory listings.

      By default, rclone will hide OneNote files in directory listings because
      operations like "Open" and "Update" won't work on them.  But this
      behaviour may also prevent you from deleting them.  If you want to
      delete OneNote files or otherwise want them to show up in directory
      listing, set this option.

   --onedrive-link-password
      Set the password for links created by the link command.

      At the time of writing this only works with OneDrive personal paid accounts.


   --onedrive-client-secret
      OAuth Client Secret.

      Leave blank normally.


OPTIONS:
   --help, -h                      show help
   --onedrive-client-id value      OAuth Client Id. [$ONEDRIVE_CLIENT_ID]
   --onedrive-client-secret value  OAuth Client Secret. [$ONEDRIVE_CLIENT_SECRET]
   --onedrive-region value         Choose national cloud region for OneDrive. (default: "global") [$ONEDRIVE_REGION]

   Advanced Options

   --onedrive-access-scopes value               Set scopes to be requested by rclone. (default: "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access") [$ONEDRIVE_ACCESS_SCOPES]
   --onedrive-auth-url value                    Auth server URL. [$ONEDRIVE_AUTH_URL]
   --onedrive-chunk-size value                  Chunk size to upload files with - must be multiple of 320k (327,680 bytes). (default: "10Mi") [$ONEDRIVE_CHUNK_SIZE]
   --onedrive-drive-id value                    The ID of the drive to use. [$ONEDRIVE_DRIVE_ID]
   --onedrive-drive-type value                  The type of the drive (personal | business | documentLibrary). [$ONEDRIVE_DRIVE_TYPE]
   --onedrive-encoding value                    The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ONEDRIVE_ENCODING]
   --onedrive-expose-onenote-files value        Set to make OneNote files show up in directory listings. (default: "false") [$ONEDRIVE_EXPOSE_ONENOTE_FILES]
   --onedrive-hash-type value                   Specify the hash in use for the backend. (default: "auto") [$ONEDRIVE_HASH_TYPE]
   --onedrive-link-password value               Set the password for links created by the link command. [$ONEDRIVE_LINK_PASSWORD]
   --onedrive-link-scope value                  Set the scope of the links created by the link command. (default: "anonymous") [$ONEDRIVE_LINK_SCOPE]
   --onedrive-link-type value                   Set the type of the links created by the link command. (default: "view") [$ONEDRIVE_LINK_TYPE]
   --onedrive-list-chunk value                  Size of listing chunk. (default: "1000") [$ONEDRIVE_LIST_CHUNK]
   --onedrive-no-versions value                 Remove all versions on modifying operations. (default: "false") [$ONEDRIVE_NO_VERSIONS]
   --onedrive-root-folder-id value              ID of the root folder. [$ONEDRIVE_ROOT_FOLDER_ID]
   --onedrive-server-side-across-configs value  Allow server-side operations (e.g. copy) to work across different onedrive configs. (default: "false") [$ONEDRIVE_SERVER_SIDE_ACROSS_CONFIGS]
   --onedrive-token value                       OAuth Access Token as a JSON blob. [$ONEDRIVE_TOKEN]
   --onedrive-token-url value                   Token server url. [$ONEDRIVE_TOKEN_URL]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)

```
{% endcode %}
