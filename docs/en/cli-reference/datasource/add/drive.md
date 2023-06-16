# Google Drive

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add drive - Google Drive

USAGE:
   singularity datasource add drive [command options] <dataset_name> <source_path>

DESCRIPTION:
   --drive-token
      OAuth Access Token as a JSON blob.

   --drive-root-folder-id
      ID of the root folder.
      Leave blank normally.
      
      Fill in to access "Computers" folders (see docs), or for rclone to use
      a non root folder as its starting point.
      

   --drive-v2-download-min-size
      If Object's are greater, use drive v2 API to download.

   --drive-stop-on-download-limit
      Make download limit errors be fatal.
      
      At the time of writing it is only possible to download 10 TiB of data from
      Google Drive a day (this is an undocumented limit). When this limit is
      reached Google Drive produces a slightly different error message. When
      this flag is set it causes these errors to be fatal.  These will stop
      the in-progress sync.
      
      Note that this detection is relying on error message strings which
      Google don't document so it may break in the future.
      

   --drive-formats
      Deprecated: See export_formats.

   --drive-import-formats
      Comma separated list of preferred formats for uploading Google docs.

   --drive-upload-cutoff
      Cutoff for switching to chunked upload.

   --drive-team-drive
      ID of the Shared Drive (Team Drive).

   --drive-auth-owner-only
      Only consider files owned by the authenticated user.

   --drive-copy-shortcut-content
      Server side copy contents of shortcuts instead of the shortcut.
      
      When doing server side copies, normally rclone will copy shortcuts as
      shortcuts.
      
      If this flag is used then rclone will copy the contents of shortcuts
      rather than shortcuts themselves when doing server side copies.

   --drive-skip-gdocs
      Skip google documents in all listings.
      
      If given, gdocs practically become invisible to rclone.

   --drive-shared-with-me
      Only show files that are shared with me.
      
      Instructs rclone to operate on your "Shared with me" folder (where
      Google Drive lets you access the files and folders others have shared
      with you).
      
      This works both with the "list" (lsd, lsl, etc.) and the "copy"
      commands (copy, sync, etc.), and with all other commands too.

   --drive-stop-on-upload-limit
      Make upload limit errors be fatal.
      
      At the time of writing it is only possible to upload 750 GiB of data to
      Google Drive a day (this is an undocumented limit). When this limit is
      reached Google Drive produces a slightly different error message. When
      this flag is set it causes these errors to be fatal.  These will stop
      the in-progress sync.
      
      Note that this detection is relying on error message strings which
      Google don't document so it may break in the future.
      
      See: https://github.com/rclone/rclone/issues/3857
      

   --drive-scope
      Scope that rclone should use when requesting access from drive.

      Examples:
         | drive                   | Full access all files, excluding Application Data Folder.
         | drive.readonly          | Read-only access to file metadata and file contents.
         | drive.file              | Access to files created by rclone only.
                                   | These are visible in the drive website.
                                   | File authorization is revoked when the user deauthorizes the app.
         | drive.appfolder         | Allows read and write access to the Application Data folder.
                                   | This is not visible in the drive website.
         | drive.metadata.readonly | Allows read-only access to file metadata but
                                   | does not allow any access to read or download file content.

   --drive-skip-dangling-shortcuts
      If set skip dangling shortcut files.
      
      If this is set then rclone will not show any dangling shortcuts in listings.
      

   --drive-alternate-export
      Deprecated: No longer needed.

   --drive-acknowledge-abuse
      Set to allow files which return cannotDownloadAbusiveFile to be downloaded.
      
      If downloading a file returns the error "This file has been identified
      as malware or spam and cannot be downloaded" with the error code
      "cannotDownloadAbusiveFile" then supply this flag to rclone to
      indicate you acknowledge the risks of downloading the file and rclone
      will download it anyway.
      
      Note that if you are using service account it will need Manager
      permission (not Content Manager) to for this flag to work. If the SA
      does not have the right permission, Google will just ignore the flag.

   --drive-keep-revision-forever
      Keep new head revision of each file forever.

   --drive-client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --drive-token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --drive-service-account-file
      Service Account Credentials JSON file path.
      
      Leave blank normally.
      Needed only if you want use SA instead of interactive login.
      
      Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.

   --drive-use-trash
      Send files to the trash instead of deleting permanently.
      
      Defaults to true, namely sending files to the trash.
      Use `--drive-use-trash=false` to delete files permanently instead.

   --drive-skip-checksum-gphotos
      Skip MD5 checksum on Google photos and videos only.
      
      Use this if you get checksum errors when transferring Google photos or
      videos.
      
      Setting this flag will cause Google photos and videos to return a
      blank MD5 checksum.
      
      Google photos are identified by being in the "photos" space.
      
      Corrupted checksums are caused by Google modifying the image/video but
      not updating the checksum.

   --drive-pacer-min-sleep
      Minimum time to sleep between API calls.

   --drive-disable-http2
      Disable drive using http2.
      
      There is currently an unsolved issue with the google drive backend and
      HTTP/2.  HTTP/2 is therefore disabled by default for the drive backend
      but can be re-enabled here.  When the issue is solved this flag will
      be removed.
      
      See: https://github.com/rclone/rclone/issues/3631
      
      

   --drive-skip-shortcuts
      If set skip shortcut files.
      
      Normally rclone dereferences shortcut files making them appear as if
      they are the original file (see [the shortcuts section](#shortcuts)).
      If this flag is set then rclone will ignore shortcut files completely.
      

   --drive-service-account-credentials
      Service Account Credentials JSON blob.
      
      Leave blank normally.
      Needed only if you want use SA instead of interactive login.

   --drive-use-created-date
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

   --drive-server-side-across-configs
      Allow server-side operations (e.g. copy) to work across different drive configs.
      
      This can be useful if you wish to do a server-side copy between two
      different Google drives.  Note that this isn't enabled by default
      because it isn't easy to tell if it will work between any two
      configurations.

   --drive-resource-key
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
      

   --drive-impersonate
      Impersonate this user when using a service account.

   --drive-chunk-size
      Upload chunk size.
      
      Must a power of 2 >= 256k.
      
      Making this larger will improve performance, but note that each chunk
      is buffered in memory one per transfer.
      
      Reducing this will reduce memory usage but decrease performance.

   --drive-size-as-quota
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

   --drive-client-id
      Google Application Client Id
      Setting your own is recommended.
      See https://rclone.org/drive/#making-your-own-client-id for how to create your own.
      If you leave this blank, it will use an internal key which is low performance.

   --drive-auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --drive-trashed-only
      Only show files that are in the trash.
      
      This will show trashed files in their original directory structure.

   --drive-allow-import-name-change
      Allow the filetype to change when uploading Google docs.
      
      E.g. file.doc to file.docx. This will confuse sync and reupload every time.

   --drive-list-chunk
      Size of listing chunk 100-1000, 0 to disable.

   --drive-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --drive-starred-only
      Only show files that are starred.

   --drive-export-formats
      Comma separated list of preferred formats for downloading Google docs.

   --drive-use-shared-date
      Use date file was shared instead of modified date.
      
      Note that, as with "--drive-use-created-date", this flag may have
      unexpected consequences when uploading/downloading files.
      
      If both this flag and "--drive-use-created-date" are set, the created
      date is used.

   --drive-pacer-burst
      Number of API calls to allow without sleeping.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for drive

   --drive-acknowledge-abuse value            Set to allow files which return cannotDownloadAbusiveFile to be downloaded. (default: "false") [$DRIVE_ACKNOWLEDGE_ABUSE]
   --drive-allow-import-name-change value     Allow the filetype to change when uploading Google docs. (default: "false") [$DRIVE_ALLOW_IMPORT_NAME_CHANGE]
   --drive-auth-owner-only value              Only consider files owned by the authenticated user. (default: "false") [$DRIVE_AUTH_OWNER_ONLY]
   --drive-auth-url value                     Auth server URL. [$DRIVE_AUTH_URL]
   --drive-chunk-size value                   Upload chunk size. (default: "8Mi") [$DRIVE_CHUNK_SIZE]
   --drive-client-id value                    Google Application Client Id [$DRIVE_CLIENT_ID]
   --drive-client-secret value                OAuth Client Secret. [$DRIVE_CLIENT_SECRET]
   --drive-copy-shortcut-content value        Server side copy contents of shortcuts instead of the shortcut. (default: "false") [$DRIVE_COPY_SHORTCUT_CONTENT]
   --drive-disable-http2 value                Disable drive using http2. (default: "true") [$DRIVE_DISABLE_HTTP2]
   --drive-encoding value                     The encoding for the backend. (default: "InvalidUtf8") [$DRIVE_ENCODING]
   --drive-export-formats value               Comma separated list of preferred formats for downloading Google docs. (default: "docx,xlsx,pptx,svg") [$DRIVE_EXPORT_FORMATS]
   --drive-formats value                      Deprecated: See export_formats. [$DRIVE_FORMATS]
   --drive-impersonate value                  Impersonate this user when using a service account. [$DRIVE_IMPERSONATE]
   --drive-import-formats value               Comma separated list of preferred formats for uploading Google docs. [$DRIVE_IMPORT_FORMATS]
   --drive-keep-revision-forever value        Keep new head revision of each file forever. (default: "false") [$DRIVE_KEEP_REVISION_FOREVER]
   --drive-list-chunk value                   Size of listing chunk 100-1000, 0 to disable. (default: "1000") [$DRIVE_LIST_CHUNK]
   --drive-pacer-burst value                  Number of API calls to allow without sleeping. (default: "100") [$DRIVE_PACER_BURST]
   --drive-pacer-min-sleep value              Minimum time to sleep between API calls. (default: "100ms") [$DRIVE_PACER_MIN_SLEEP]
   --drive-resource-key value                 Resource key for accessing a link-shared file. [$DRIVE_RESOURCE_KEY]
   --drive-root-folder-id value               ID of the root folder. [$DRIVE_ROOT_FOLDER_ID]
   --drive-scope value                        Scope that rclone should use when requesting access from drive. [$DRIVE_SCOPE]
   --drive-server-side-across-configs value   Allow server-side operations (e.g. copy) to work across different drive configs. (default: "false") [$DRIVE_SERVER_SIDE_ACROSS_CONFIGS]
   --drive-service-account-credentials value  Service Account Credentials JSON blob. [$DRIVE_SERVICE_ACCOUNT_CREDENTIALS]
   --drive-service-account-file value         Service Account Credentials JSON file path. [$DRIVE_SERVICE_ACCOUNT_FILE]
   --drive-shared-with-me value               Only show files that are shared with me. (default: "false") [$DRIVE_SHARED_WITH_ME]
   --drive-size-as-quota value                Show sizes as storage quota usage, not actual size. (default: "false") [$DRIVE_SIZE_AS_QUOTA]
   --drive-skip-checksum-gphotos value        Skip MD5 checksum on Google photos and videos only. (default: "false") [$DRIVE_SKIP_CHECKSUM_GPHOTOS]
   --drive-skip-dangling-shortcuts value      If set skip dangling shortcut files. (default: "false") [$DRIVE_SKIP_DANGLING_SHORTCUTS]
   --drive-skip-gdocs value                   Skip google documents in all listings. (default: "false") [$DRIVE_SKIP_GDOCS]
   --drive-skip-shortcuts value               If set skip shortcut files. (default: "false") [$DRIVE_SKIP_SHORTCUTS]
   --drive-starred-only value                 Only show files that are starred. (default: "false") [$DRIVE_STARRED_ONLY]
   --drive-stop-on-download-limit value       Make download limit errors be fatal. (default: "false") [$DRIVE_STOP_ON_DOWNLOAD_LIMIT]
   --drive-stop-on-upload-limit value         Make upload limit errors be fatal. (default: "false") [$DRIVE_STOP_ON_UPLOAD_LIMIT]
   --drive-team-drive value                   ID of the Shared Drive (Team Drive). [$DRIVE_TEAM_DRIVE]
   --drive-token value                        OAuth Access Token as a JSON blob. [$DRIVE_TOKEN]
   --drive-token-url value                    Token server url. [$DRIVE_TOKEN_URL]
   --drive-trashed-only value                 Only show files that are in the trash. (default: "false") [$DRIVE_TRASHED_ONLY]
   --drive-upload-cutoff value                Cutoff for switching to chunked upload. (default: "8Mi") [$DRIVE_UPLOAD_CUTOFF]
   --drive-use-created-date value             Use file created date instead of modified date. (default: "false") [$DRIVE_USE_CREATED_DATE]
   --drive-use-shared-date value              Use date file was shared instead of modified date. (default: "false") [$DRIVE_USE_SHARED_DATE]
   --drive-use-trash value                    Send files to the trash instead of deleting permanently. (default: "true") [$DRIVE_USE_TRASH]
   --drive-v2-download-min-size value         If Object's are greater, use drive v2 API to download. (default: "off") [$DRIVE_V2_DOWNLOAD_MIN_SIZE]

```
{% endcode %}
