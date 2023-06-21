# Local Disk

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add local - Local Disk

USAGE:
   singularity datasource add local [command options] <dataset_name> <source_path>

DESCRIPTION:
   --local-case-insensitive
      Force the filesystem to report itself as case insensitive.
      
      Normally the local backend declares itself as case insensitive on
      Windows/macOS and case sensitive for everything else.  Use this flag
      to override the default choice.

   --local-no-sparse
      Disable sparse files for multi-thread downloads.
      
      On Windows platforms rclone will make sparse files when doing
      multi-thread downloads. This avoids long pauses on large files where
      the OS zeros the file. However sparse files may be undesirable as they
      cause disk fragmentation and can be slow to work with.

   --local-copy-links
      Follow symlinks and copy the pointed to item.

   --local-links
      Translate symlinks to/from regular files with a '.rclonelink' extension.

   --local-zero-size-links
      Assume the Stat size of links is zero (and read them instead) (deprecated).
      
      Rclone used to use the Stat size of links as the link size, but this fails in quite a few places:
      
      - Windows
      - On some virtual filesystems (such ash LucidLink)
      - Android
      
      So rclone now always reads the link.
      

   --local-case-sensitive
      Force the filesystem to report itself as case sensitive.
      
      Normally the local backend declares itself as case insensitive on
      Windows/macOS and case sensitive for everything else.  Use this flag
      to override the default choice.

   --local-one-file-system
      Don't cross filesystem boundaries (unix/macOS only).

   --local-no-preallocate
      Disable preallocation of disk space for transferred files.
      
      Preallocation of disk space helps prevent filesystem fragmentation.
      However, some virtual filesystem layers (such as Google Drive File
      Stream) may incorrectly set the actual file size equal to the
      preallocated space, causing checksum and file size checks to fail.
      Use this flag to disable preallocation.

   --local-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --local-skip-links
      Don't warn about skipped symlinks.
      
      This flag disables warning messages on skipped symlinks or junction
      points, as you explicitly acknowledge that they should be skipped.

   --local-unicode-normalization
      Apply unicode NFC normalization to paths and filenames.
      
      This flag can be used to normalize file names into unicode NFC form
      that are read from the local filesystem.
      
      Rclone does not normally touch the encoding of file names it reads from
      the file system.
      
      This can be useful when using macOS as it normally provides decomposed (NFD)
      unicode which in some language (eg Korean) doesn't display properly on
      some OSes.
      
      Note that rclone compares filenames with unicode normalization in the sync
      routine so this flag shouldn't normally be used.

   --local-no-set-modtime
      Disable setting modtime.
      
      Normally rclone updates modification time of files after they are done
      uploading. This can cause permissions issues on Linux platforms when 
      the user rclone is running as does not own the file uploaded, such as
      when copying to a CIFS mount owned by another user. If this option is 
      enabled, rclone will no longer update the modtime after copying a file.

   --local-nounc
      Disable UNC (long path names) conversion on Windows.

      Examples:
         | true | Disables long file names.

   --local-no-check-updated
      Don't check to see if the files change during upload.
      
      Normally rclone checks the size and modification time of files as they
      are being uploaded and aborts with a message which starts "can't copy -
      source file is being updated" if the file changes during upload.
      
      However on some file systems this modification time check may fail (e.g.
      [Glusterfs #2206](https://github.com/rclone/rclone/issues/2206)) so this
      check can be disabled with this flag.
      
      If this flag is set, rclone will use its best efforts to transfer a
      file which is being updated. If the file is only having things
      appended to it (e.g. a log) then rclone will transfer the log file with
      the size it had the first time rclone saw it.
      
      If the file is being modified throughout (not just appended to) then
      the transfer may fail with a hash check failure.
      
      In detail, once the file has had stat() called on it for the first
      time we:
      
      - Only transfer the size that stat gave
      - Only checksum the size that stat gave
      - Don't update the stat info for the file
      
      


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for local

   --local-case-insensitive value           Force the filesystem to report itself as case insensitive. (default: "false") [$LOCAL_CASE_INSENSITIVE]
   --local-case-sensitive value             Force the filesystem to report itself as case sensitive. (default: "false") [$LOCAL_CASE_SENSITIVE]
   --local-copy-links value, -L value       Follow symlinks and copy the pointed to item. (default: "false") [$LOCAL_COPY_LINKS]
   --local-encoding value                   The encoding for the backend. (default: "Slash,Dot") [$LOCAL_ENCODING]
   --local-links value, -l value            Translate symlinks to/from regular files with a '.rclonelink' extension. (default: "false") [$LOCAL_LINKS]
   --local-no-check-updated value           Don't check to see if the files change during upload. (default: "false") [$LOCAL_NO_CHECK_UPDATED]
   --local-no-preallocate value             Disable preallocation of disk space for transferred files. (default: "false") [$LOCAL_NO_PREALLOCATE]
   --local-no-set-modtime value             Disable setting modtime. (default: "false") [$LOCAL_NO_SET_MODTIME]
   --local-no-sparse value                  Disable sparse files for multi-thread downloads. (default: "false") [$LOCAL_NO_SPARSE]
   --local-nounc value                      Disable UNC (long path names) conversion on Windows. (default: "false") [$LOCAL_NOUNC]
   --local-one-file-system value, -x value  Don't cross filesystem boundaries (unix/macOS only). (default: "false") [$LOCAL_ONE_FILE_SYSTEM]
   --local-skip-links value                 Don't warn about skipped symlinks. (default: "false") [$LOCAL_SKIP_LINKS]
   --local-unicode-normalization value      Apply unicode NFC normalization to paths and filenames. (default: "false") [$LOCAL_UNICODE_NORMALIZATION]
   --local-zero-size-links value            Assume the Stat size of links is zero (and read them instead) (deprecated). (default: "false") [$LOCAL_ZERO_SIZE_LINKS]

```
{% endcode %}
