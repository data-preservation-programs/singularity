# Encryption

## Overview

Singularity supports a built-in encryption solution which encrypts the files with provided recipients (public keys), or even a hardware PIV tokens such as YubiKeys. You may also supply a custom encryption script which allows you to integrate with any external encryption solution \[Need testing].

## Built-in Encryption

Start by creating a public-private key pair for asymmetric encryption. The underlying encryption library used by Singularity is called [age](https://github.com/FiloSottile/age).

```sh
go install filippo.io/age/cmd/...@latest
age-keygen -o key.txt
> Public key: agexxxxxxxxxxxx
```

Now, we can setup a dataset to encrypt each file using the generated public key

```sh
singularity dataset create --encryption-recipient agexxxxxxxxxxxx \
  --output-dir . test
```

Inline preparation is disabled because we can't encrypting the same file a second time will yield different encrypted content due to initial randomness introduced during the encryption process.

We can then add data source as before continue our data preparation process. Do pay attention that the folder structure will not be encrypted so you can either choose to generate the DAG for the folder structure or not run the `daggen` command. In the latter case, folder structure will only be accessible from Singularity database and commands.

## Custom Encryption

Singularity also offers custom encryption by supplying a custom script to encrypt file stream. It can be potentially used with Key management service and custom encryption algorithm or tooling.

