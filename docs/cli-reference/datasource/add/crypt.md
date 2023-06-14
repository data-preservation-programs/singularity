# Encrypt/Decrypt a remote

```
NAME:
   singularity datasource add crypt - Encrypt/Decrypt a remote

USAGE:
   singularity datasource add crypt [command options] <dataset_name> <source_path>

DESCRIPTION:
   --crypt-password2
      Password or pass phrase for salt.
      
      Optional but recommended.
      Should be different to the previous password.

   --crypt-server-side-across-configs
      Allow server-side operations (e.g. copy) to work across different crypt configs.
      
      Normally this option is not what you want, but if you have two crypts
      pointing to the same backend you can use it.
      
      This can be used, for example, to change file name encryption type
      without re-uploading all the data. Just make two crypt backends
      pointing to two different directories with the single changed
      parameter and use rclone move to move the files between the crypt
      remotes.

   --crypt-no-data-encryption
      Option to either encrypt file data or leave it unencrypted.

      Examples:
         | true  | Don't encrypt file data, leave it unencrypted.
         | false | Encrypt file data.

   --crypt-filename-encoding
      How to encode the encrypted filename to text string.
      
      This option could help with shortening the encrypted filename. The 
      suitable option would depend on the way your remote count the filename
      length and if it's case sensitive.

      Examples:
         | base32    | Encode using base32. Suitable for all remote.
         | base64    | Encode using base64. Suitable for case sensitive remote.
         | base32768 | Encode using base32768. Suitable if your remote counts UTF-16 or
                     | Unicode codepoint instead of UTF-8 byte length. (Eg. Onedrive)

   --crypt-remote
      Remote to encrypt/decrypt.
      
      Normally should contain a ':' and a path, e.g. "myremote:path/to/dir",
      "myremote:bucket" or maybe "myremote:" (not recommended).

   --crypt-filename-encryption
      How to encrypt the filenames.

      Examples:
         | standard  | Encrypt the filenames.
                     | See the docs for the details.
         | obfuscate | Very simple filename obfuscation.
         | off       | Don't encrypt the file names.
                     | Adds a ".bin" extension only.

   --crypt-directory-name-encryption
      Option to either encrypt directory names or leave them intact.
      
      NB If filename_encryption is "off" then this option will do nothing.

      Examples:
         | true  | Encrypt directory names.
         | false | Don't encrypt directory names, leave them intact.

   --crypt-password
      Password or pass phrase for encryption.

   --crypt-show-mapping
      For all files listed show how the names encrypt.
      
      If this flag is set then for each file that the remote is asked to
      list, it will log (at level INFO) a line stating the decrypted file
      name and the encrypted file name.
      
      This is so you can work out which encrypted names are which decrypted
      names just in case you need to do something with the encrypted file
      names, or for debugging purposes.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for crypt

   --crypt-directory-name-encryption value   Option to either encrypt directory names or leave them intact. (default: "true") [$CRYPT_DIRECTORY_NAME_ENCRYPTION]
   --crypt-filename-encoding value           How to encode the encrypted filename to text string. (default: "base32") [$CRYPT_FILENAME_ENCODING]
   --crypt-filename-encryption value         How to encrypt the filenames. (default: "standard") [$CRYPT_FILENAME_ENCRYPTION]
   --crypt-no-data-encryption value          Option to either encrypt file data or leave it unencrypted. (default: "false") [$CRYPT_NO_DATA_ENCRYPTION]
   --crypt-password value                    Password or pass phrase for encryption. [$CRYPT_PASSWORD]
   --crypt-password2 value                   Password or pass phrase for salt. [$CRYPT_PASSWORD2]
   --crypt-remote value                      Remote to encrypt/decrypt. [$CRYPT_REMOTE]
   --crypt-server-side-across-configs value  Allow server-side operations (e.g. copy) to work across different crypt configs. (default: "false") [$CRYPT_SERVER_SIDE_ACROSS_CONFIGS]
   --crypt-show-mapping value                For all files listed show how the names encrypt. (default: "false") [$CRYPT_SHOW_MAPPING]

```
