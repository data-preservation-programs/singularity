# Microsoft OneDrive

{% code fullWidth="true" %}
```
NAME:
   singularity storage update onedrive - Microsoft OneDrive

USAGE:
   singularity storage update onedrive [command options] <name>

DESCRIPTION:
   --client_id
      OAuth Client Id.
      
      Leave blank normally.

   --client_secret
      OAuth Client Secret.
      
      Leave blank normally.

   --token
      OAuth Access Token as a JSON blob.

   --auth_url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --token_url
      Token server url.
      
      Leave blank to use the provider defaults.

   --region
      Choose national cloud region for OneDrive.

      Examples:
         | global | Microsoft Cloud Global
         | us     | Microsoft Cloud for US Government
         | de     | Microsoft Cloud Germany
         | cn     | Azure and Office 365 operated by Vnet Group in China

   --chunk_size
      Chunk size to upload files with - must be multiple of 320k (327,680 bytes).
      
      Above this size files will be chunked - must be multiple of 320k (327,680 bytes) and
      should not exceed 250M (262,144,000 bytes) else you may encounter \"Microsoft.SharePoint.Client.InvalidClientQueryException: The request message is too big.\"
      Note that the chunks will be buffered into memory.

   --drive_id
      The ID of the drive to use.

   --drive_type
      The type of the drive (personal | business | documentLibrary).

   --root_folder_id
      ID of the root folder.
      
      This isn't normally needed, but in special circumstances you might
      know the folder ID that you wish to access but not be able to get
      there through a path traversal.
      

   --access_scopes
      Set scopes to be requested by rclone.
      
      Choose or manually enter a custom space separated list with all scopes, that rclone should request.
      

      Examples:
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access | Read and write access to all resources
         | Files.Read Files.Read.All Sites.Read.All offline_access                                     | Read only access to all resources
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All offline_access                | Read and write access to all resources, without the ability to browse SharePoint sites. 
         |                                                                                             | Same as if disable_site_permission was set to true

   --disable_site_permission
      Disable the request for Sites.Read.All permission.
      
      If set to true, you will no longer be able to search for a SharePoint site when
      configuring drive ID, because rclone will not request Sites.Read.All permission.
      Set it to true if your organization didn't assign Sites.Read.All permission to the
      application, and your organization disallows users to consent app permission
      request on their own.

   --expose_onenote_files
      Set to make OneNote files show up in directory listings.
      
      By default, rclone will hide OneNote files in directory listings because
      operations like "Open" and "Update" won't work on them.  But this
      behaviour may also prevent you from deleting them.  If you want to
      delete OneNote files or otherwise want them to show up in directory
      listing, set this option.

   --server_side_across_configs
      Allow server-side operations (e.g. copy) to work across different onedrive configs.
      
      This will only work if you are copying between two OneDrive *Personal* drives AND
      the files to copy are already shared between them.  In other cases, rclone will
      fall back to normal copy (which will be slightly slower).

   --list_chunk
      Size of listing chunk.

   --no_versions
      Remove all versions on modifying operations.
      
      Onedrive for business creates versions when rclone uploads new files
      overwriting an existing one and when it sets the modification time.
      
      These versions take up space out of the quota.
      
      This flag checks for versions after file upload and setting
      modification time and removes all but the last version.
      
      **NB** Onedrive personal can't currently delete versions so don't use
      this flag there.
      

   --link_scope
      Set the scope of the links created by the link command.

      Examples:
         | anonymous    | Anyone with the link has access, without needing to sign in.
         |              | This may include people outside of your organization.
         |              | Anonymous link support may be disabled by an administrator.
         | organization | Anyone signed into your organization (tenant) can use the link to get access.
         |              | Only available in OneDrive for Business and SharePoint.

   --link_type
      Set the type of the links created by the link command.

      Examples:
         | view  | Creates a read-only link to the item.
         | edit  | Creates a read-write link to the item.
         | embed | Creates an embeddable link to the item.

   --link_password
      Set the password for links created by the link command.
      
      At the time of writing this only works with OneDrive personal paid accounts.
      

   --hash_type
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

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --client_id value      OAuth Client Id. [$CLIENT_ID]
   --client_secret value  OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h             show help
   --region value         Choose national cloud region for OneDrive. (default: "global") [$REGION]

   Advanced

   --access_scopes value         Set scopes to be requested by rclone. (default: "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access") [$ACCESS_SCOPES]
   --auth_url value              Auth server URL. [$AUTH_URL]
   --chunk_size value            Chunk size to upload files with - must be multiple of 320k (327,680 bytes). (default: "10Mi") [$CHUNK_SIZE]
   --disable_site_permission     Disable the request for Sites.Read.All permission. (default: false) [$DISABLE_SITE_PERMISSION]
   --drive_id value              The ID of the drive to use. [$DRIVE_ID]
   --drive_type value            The type of the drive (personal | business | documentLibrary). [$DRIVE_TYPE]
   --encoding value              The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --expose_onenote_files        Set to make OneNote files show up in directory listings. (default: false) [$EXPOSE_ONENOTE_FILES]
   --hash_type value             Specify the hash in use for the backend. (default: "auto") [$HASH_TYPE]
   --link_password value         Set the password for links created by the link command. [$LINK_PASSWORD]
   --link_scope value            Set the scope of the links created by the link command. (default: "anonymous") [$LINK_SCOPE]
   --link_type value             Set the type of the links created by the link command. (default: "view") [$LINK_TYPE]
   --list_chunk value            Size of listing chunk. (default: 1000) [$LIST_CHUNK]
   --no_versions                 Remove all versions on modifying operations. (default: false) [$NO_VERSIONS]
   --root_folder_id value        ID of the root folder. [$ROOT_FOLDER_ID]
   --server_side_across_configs  Allow server-side operations (e.g. copy) to work across different onedrive configs. (default: false) [$SERVER_SIDE_ACROSS_CONFIGS]
   --token value                 OAuth Access Token as a JSON blob. [$TOKEN]
   --token_url value             Token server url. [$TOKEN_URL]

```
{% endcode %}
