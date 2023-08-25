# Google Drive

{% code fullWidth="true" %}
```
NAME:
   singularity storage create drive - Google Drive

USAGE:
   singularity storage create drive [command options] <name> <path>

DESCRIPTION:
   --client_id
      Google Application Client Id
      Setting your own is recommended.
      See https://rclone.org/drive/#making-your-own-client-id for how to create your own.
      If you leave this blank, it will use an internal key which is low performance.

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

   --scope
      Scope that rclone should use when requesting access from drive.

      Examples:
         | drive                   | Full access all files, excluding Application Data Folder.
         | drive.readonly          | Read-only access to file metadata and file contents.
         | drive.file              | Access to files created by rclone only.
         |                         | These are visible in the drive website.
         |                         | File authorization is revoked when the user deauthorizes the app.
         | drive.appfolder         | Allows read and write access to the Application Data folder.
         |                         | This is not visible in the drive website.
         | drive.metadata.readonly | Allows read-only access to file metadata but
         |                         | does not allow any access to read or download file content.

   --root_folder_id
      ID of the root folder.
      Leave blank normally.
      
      Fill in to access "Computers" folders (see docs), or for rclone to use
      a non root folder as its starting point.
      

   --service_account_file
      Service Account Credentials JSON file path.
      
      Leave blank normally.
      Needed only if you want use SA instead of interactive login.
      
      Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.

   --service_account_credentials
      Service Account Credentials JSON blob.
      
      Leave blank normally.
      Needed only if you want use SA instead of interactive login.

   --team_drive
      ID of the Shared Drive (Team Drive).

   --auth_owner_only
      Only consider files owned by the authenticated user.

   --use_trash
      Send files to the trash instead of deleting permanently.
      
      Defaults to true, namely sending files to the trash.
      Use `--drive-use-trash=false` to delete files permanently instead.

   --copy_shortcut_content
      Server side copy contents of shortcuts instead of the shortcut.
      
      When doing server side copies, normally rclone will copy shortcuts as
      shortcuts.
      
      If this flag is used then rclone will copy the contents of shortcuts
      rather than shortcuts themselves when doing server side copies.

   --skip_gdocs
      Skip google documents in all listings.
      
      If given, gdocs practically become invisible to rclone.

   --skip_checksum_gphotos
      Skip MD5 checksum on Google photos and videos only.
      
      Use this if you get checksum errors when transferring Google photos or
      videos.
      
      Setting this flag will cause Google photos and videos to return a
      blank MD5 checksum.
      
      Google photos are identified by being in the "photos" space.
      
      Corrupted checksums are caused by Google modifying the image/video but
      not updating the checksum.

   --shared_with_me
      Only show files that are shared with me.
      
      Instructs rclone to operate on your "Shared with me" folder (where
      Google Drive lets you access the files and folders others have shared
      with you).
      
      This works both with the "list" (lsd, lsl, etc.) and the "copy"
      commands (copy, sync, etc.), and with all other commands too.

   --trashed_only
      Only show files that are in the trash.
      
      This will show trashed files in their original directory structure.

   --starred_only
      Only show files that are starred.

   --formats
      Deprecated: See export_formats.

   --export_formats
      Comma separated list of preferred formats for downloading Google docs.

   --import_formats
      Comma separated list of preferred formats for uploading Google docs.

   --allow_import_name_change
      Allow the filetype to change when uploading Google docs.
      
      E.g. file.doc to file.docx. This will confuse sync and reupload every time.

   --use_created_date
      Use file created date instead of modified date.
      
      Useful when downloading data and you want the creation date used in
      place of the last modified date.
      
      **WARNING**: This flag may have some unexpected consequences.
      
      When uploading to your drive all files will be overwritten unless they
      haven't been modified since their creation. And the inverse will occur
      while downloading.  This side effect can be avoided by using the
      "--checksum" flag.
      
      This feature was implemented to retain photos capture date as recorded
      by google photos. You will first need to check the "Create a Google
      Photos folder" option in your google drive settings. You can then copy
      or move the photos locally and use the date the image was taken
      (created) set as the modification date.

   --use_shared_date
      Use date file was shared instead of modified date.
      
      Note that, as with "--drive-use-created-date", this flag may have
      unexpected consequences when uploading/downloading files.
      
      If both this flag and "--drive-use-created-date" are set, the created
      date is used.

   --list_chunk
      Size of listing chunk 100-1000, 0 to disable.

   --impersonate
      Impersonate this user when using a service account.

   --alternate_export
      Deprecated: No longer needed.

   --upload_cutoff
      Cutoff for switching to chunked upload.

   --chunk_size
      Upload chunk size.
      
      Must a power of 2 >= 256k.
      
      Making this larger will improve performance, but note that each chunk
      is buffered in memory one per transfer.
      
      Reducing this will reduce memory usage but decrease performance.

   --acknowledge_abuse
      Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
      
      If downloading a file returns the error "This file has been identified
      as malware or spam and cannot be downloaded" with the error code
      "cannotDownloadAbusiveFile" then supply this flag to rclone to
      indicate you acknowledge the risks of downloading the file and rclone
      will download it anyway.
      
      Note that if you are using service account it will need Manager
      permission (not Content Manager) to for this flag to work. If the SA
      does not have the right permission, Google will just ignore the flag.

   --keep_revision_forever
      Keep new head revision of each file forever.

   --size_as_quota
      Show sizes as storage quota usage, not actual size.
      
      Show the size of a file as the storage quota used. This is the
      current version plus any older versions that have been set to keep
      forever.
      
      **WARNING**: This flag may have some unexpected consequences.
      
      It is not recommended to set this flag in your config - the
      recommended usage is using the flag form --drive-size-as-quota when
      doing rclone ls/lsl/lsf/lsjson/etc only.
      
      If you do use this flag for syncing (not recommended) then you will
      need to use --ignore size also.

   --v2_download_min_size
      If Object's are greater, use drive v2 API to download.

   --pacer_min_sleep
      Minimum time to sleep between API calls.

   --pacer_burst
      Number of API calls to allow without sleeping.

   --server_side_across_configs
      Allow server-side operations (e.g. copy) to work across different drive configs.
      
      This can be useful if you wish to do a server-side copy between two
      different Google drives.  Note that this isn't enabled by default
      because it isn't easy to tell if it will work between any two
      configurations.

   --disable_http2
      Disable drive using http2.
      
      There is currently an unsolved issue with the google drive backend and
      HTTP/2.  HTTP/2 is therefore disabled by default for the drive backend
      but can be re-enabled here.  When the issue is solved this flag will
      be removed.
      
      See: https://github.com/rclone/rclone/issues/3631
      
      

   --stop_on_upload_limit
      Make upload limit errors be fatal.
      
      At the time of writing it is only possible to upload 750 GiB of data to
      Google Drive a day (this is an undocumented limit). When this limit is
      reached Google Drive produces a slightly different error message. When
      this flag is set it causes these errors to be fatal.  These will stop
      the in-progress sync.
      
      Note that this detection is relying on error message strings which
      Google don't document so it may break in the future.
      
      See: https://github.com/rclone/rclone/issues/3857
      

   --stop_on_download_limit
      Make download limit errors be fatal.
      
      At the time of writing it is only possible to download 10 TiB of data from
      Google Drive a day (this is an undocumented limit). When this limit is
      reached Google Drive produces a slightly different error message. When
      this flag is set it causes these errors to be fatal.  These will stop
      the in-progress sync.
      
      Note that this detection is relying on error message strings which
      Google don't document so it may break in the future.
      

   --skip_shortcuts
      If set skip shortcut files.
      
      Normally rclone dereferences shortcut files making them appear as if
      they are the original file (see [the shortcuts section](#shortcuts)).
      If this flag is set then rclone will ignore shortcut files completely.
      

   --skip_dangling_shortcuts
      If set skip dangling shortcut files.
      
      If this is set then rclone will not show any dangling shortcuts in listings.
      

   --resource_key
      Resource key for accessing a link-shared file.
      
      If you need to access files shared with a link like this
      
          https://drive.google.com/drive/folders/XXX?resourcekey=YYY&usp=sharing
      
      Then you will need to use the first part "XXX" as the "root_folder_id"
      and the second part "YYY" as the "resource_key" otherwise you will get
      404 not found errors when trying to access the directory.
      
      See: https://developers.google.com/drive/api/guides/resource-keys
      
      This resource key requirement only applies to a subset of old files.
      
      Note also that opening the folder once in the web interface (with the
      user you've authenticated rclone with) seems to be enough so that the
      resource key is no needed.
      

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --alternate_export            Deprecated: No longer needed. (default: false) [$ALTERNATE_EXPORT]
   --client_id value             Google Application Client Id [$CLIENT_ID]
   --client_secret value         OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h                    show help
   --scope value                 Scope that rclone should use when requesting access from drive. [$SCOPE]
   --service_account_file value  Service Account Credentials JSON file path. [$SERVICE_ACCOUNT_FILE]

   Advanced

   --acknowledge_abuse                  Set to allow files which return cannotDownloadAbusiveFile to be downloaded. (default: false) [$ACKNOWLEDGE_ABUSE]
   --allow_import_name_change           Allow the filetype to change when uploading Google docs. (default: false) [$ALLOW_IMPORT_NAME_CHANGE]
   --auth_owner_only                    Only consider files owned by the authenticated user. (default: false) [$AUTH_OWNER_ONLY]
   --auth_url value                     Auth server URL. [$AUTH_URL]
   --chunk_size value                   Upload chunk size. (default: "8Mi") [$CHUNK_SIZE]
   --copy_shortcut_content              Server side copy contents of shortcuts instead of the shortcut. (default: false) [$COPY_SHORTCUT_CONTENT]
   --disable_http2                      Disable drive using http2. (default: true) [$DISABLE_HTTP2]
   --encoding value                     The encoding for the backend. (default: "InvalidUtf8") [$ENCODING]
   --export_formats value               Comma separated list of preferred formats for downloading Google docs. (default: "docx,xlsx,pptx,svg") [$EXPORT_FORMATS]
   --formats value                      Deprecated: See export_formats. [$FORMATS]
   --impersonate value                  Impersonate this user when using a service account. [$IMPERSONATE]
   --import_formats value               Comma separated list of preferred formats for uploading Google docs. [$IMPORT_FORMATS]
   --keep_revision_forever              Keep new head revision of each file forever. (default: false) [$KEEP_REVISION_FOREVER]
   --list_chunk value                   Size of listing chunk 100-1000, 0 to disable. (default: 1000) [$LIST_CHUNK]
   --pacer_burst value                  Number of API calls to allow without sleeping. (default: 100) [$PACER_BURST]
   --pacer_min_sleep value              Minimum time to sleep between API calls. (default: "100ms") [$PACER_MIN_SLEEP]
   --resource_key value                 Resource key for accessing a link-shared file. [$RESOURCE_KEY]
   --root_folder_id value               ID of the root folder. [$ROOT_FOLDER_ID]
   --server_side_across_configs         Allow server-side operations (e.g. copy) to work across different drive configs. (default: false) [$SERVER_SIDE_ACROSS_CONFIGS]
   --service_account_credentials value  Service Account Credentials JSON blob. [$SERVICE_ACCOUNT_CREDENTIALS]
   --shared_with_me                     Only show files that are shared with me. (default: false) [$SHARED_WITH_ME]
   --size_as_quota                      Show sizes as storage quota usage, not actual size. (default: false) [$SIZE_AS_QUOTA]
   --skip_checksum_gphotos              Skip MD5 checksum on Google photos and videos only. (default: false) [$SKIP_CHECKSUM_GPHOTOS]
   --skip_dangling_shortcuts            If set skip dangling shortcut files. (default: false) [$SKIP_DANGLING_SHORTCUTS]
   --skip_gdocs                         Skip google documents in all listings. (default: false) [$SKIP_GDOCS]
   --skip_shortcuts                     If set skip shortcut files. (default: false) [$SKIP_SHORTCUTS]
   --starred_only                       Only show files that are starred. (default: false) [$STARRED_ONLY]
   --stop_on_download_limit             Make download limit errors be fatal. (default: false) [$STOP_ON_DOWNLOAD_LIMIT]
   --stop_on_upload_limit               Make upload limit errors be fatal. (default: false) [$STOP_ON_UPLOAD_LIMIT]
   --team_drive value                   ID of the Shared Drive (Team Drive). [$TEAM_DRIVE]
   --token value                        OAuth Access Token as a JSON blob. [$TOKEN]
   --token_url value                    Token server url. [$TOKEN_URL]
   --trashed_only                       Only show files that are in the trash. (default: false) [$TRASHED_ONLY]
   --upload_cutoff value                Cutoff for switching to chunked upload. (default: "8Mi") [$UPLOAD_CUTOFF]
   --use_created_date                   Use file created date instead of modified date. (default: false) [$USE_CREATED_DATE]
   --use_shared_date                    Use date file was shared instead of modified date. (default: false) [$USE_SHARED_DATE]
   --use_trash                          Send files to the trash instead of deleting permanently. (default: true) [$USE_TRASH]
   --v2_download_min_size value         If Object's are greater, use drive v2 API to download. (default: "off") [$V2_DOWNLOAD_MIN_SIZE]

```
{% endcode %}
