# Local Disk

{% code fullWidth="true" %}
```
NAME:
   singularity storage create local - Local Disk

USAGE:
   singularity storage create local [command options] [arguments...]

DESCRIPTION:
   --nounc
      Disable UNC (long path names) conversion on Windows.

      Examples:
         | true | Disables long file names.

   --copy-links
      Follow symlinks and copy the pointed to item.

   --links
      Translate symlinks to/from regular files with a '.rclonelink' extension.

   --skip-links
      Don't warn about skipped symlinks.
      
      This flag disables warning messages on skipped symlinks or junction
      points, as you explicitly acknowledge that they should be skipped.

   --zero-size-links
      Assume the Stat size of links is zero (and read them instead) (deprecated).
      
      Rclone used to use the Stat size of links as the link size, but this fails in quite a few places:
      
      - Windows
      - On some virtual filesystems (such ash LucidLink)
      - Android
      
      So rclone now always reads the link.
      

   --unicode-normalization
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

   --no-check-updated
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
      
      

   --one-file-system
      Don't cross filesystem boundaries (unix/macOS only).

   --case-sensitive
      Force the filesystem to report itself as case sensitive.
      
      Normally the local backend declares itself as case insensitive on
      Windows/macOS and case sensitive for everything else.  Use this flag
      to override the default choice.

   --case-insensitive
      Force the filesystem to report itself as case insensitive.
      
      Normally the local backend declares itself as case insensitive on
      Windows/macOS and case sensitive for everything else.  Use this flag
      to override the default choice.

   --no-preallocate
      Disable preallocation of disk space for transferred files.
      
      Preallocation of disk space helps prevent filesystem fragmentation.
      However, some virtual filesystem layers (such as Google Drive File
      Stream) may incorrectly set the actual file size equal to the
      preallocated space, causing checksum and file size checks to fail.
      Use this flag to disable preallocation.

   --no-sparse
      Disable sparse files for multi-thread downloads.
      
      On Windows platforms rclone will make sparse files when doing
      multi-thread downloads. This avoids long pauses on large files where
      the OS zeros the file. However sparse files may be undesirable as they
      cause disk fragmentation and can be slow to work with.

   --no-set-modtime
      Disable setting modtime.
      
      Normally rclone updates modification time of files after they are done
      uploading. This can cause permissions issues on Linux platforms when 
      the user rclone is running as does not own the file uploaded, such as
      when copying to a CIFS mount owned by another user. If this option is 
      enabled, rclone will no longer update the modtime after copying a file.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h  show help

   Advanced

   --case-insensitive       Force the filesystem to report itself as case insensitive. (default: false) [$CASE_INSENSITIVE]
   --case-sensitive         Force the filesystem to report itself as case sensitive. (default: false) [$CASE_SENSITIVE]
   --copy-links, -L         Follow symlinks and copy the pointed to item. (default: false) [$COPY_LINKS]
   --encoding value         The encoding for the backend. (default: "Slash,Dot") [$ENCODING]
   --links, -l              Translate symlinks to/from regular files with a '.rclonelink' extension. (default: false) [$LINKS]
   --no-check-updated       Don't check to see if the files change during upload. (default: false) [$NO_CHECK_UPDATED]
   --no-preallocate         Disable preallocation of disk space for transferred files. (default: false) [$NO_PREALLOCATE]
   --no-set-modtime         Disable setting modtime. (default: false) [$NO_SET_MODTIME]
   --no-sparse              Disable sparse files for multi-thread downloads. (default: false) [$NO_SPARSE]
   --nounc                  Disable UNC (long path names) conversion on Windows. (default: false) [$NOUNC]
   --one-file-system, -x    Don't cross filesystem boundaries (unix/macOS only). (default: false) [$ONE_FILE_SYSTEM]
   --skip-links             Don't warn about skipped symlinks. (default: false) [$SKIP_LINKS]
   --unicode-normalization  Apply unicode NFC normalization to paths and filenames. (default: false) [$UNICODE_NORMALIZATION]
   --zero-size-links        Assume the Stat size of links is zero (and read them instead) (deprecated). (default: false) [$ZERO_SIZE_LINKS]

   Client Config

   --client-scan-concurrency value  Max number of concurrent listing requests when scanning data source (default: 1)

   General

   --name value  Name of the storage (default: Auto generated)
   --path value  Path of the storage

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
