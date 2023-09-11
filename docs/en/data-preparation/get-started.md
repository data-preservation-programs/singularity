# Getting Started with Singularity

Follow these steps to set up and start using Singularity.

## 1. Initialize the Database

If you're using Singularity for the first time, you'll need to initialize the database. This step is required only once.

```sh
singularity admin init
```

## 2. Connect to Storage Systems
Singularity partners with RClone to provide seamless integration with over 40 different storage systems. These storage systems can play two main roles:
* **Source Storage**: This is where the dataset is currently stored and where Singularity will source data from for preparation.
* **Output Storage**: This is the destination where Singularity will store the CAR (Content Addressable Archive) files after processing.
Choose a storage system appropriate for your needs and connect it with Singularity to start preparing your datasets.

### 2a. Add a local file system

The most command storage system is the local file system. To add a folder as a source storage to singularity:

```sh
singularity storage create local --name "my-source" --path "/mnt/dataset/folder"
```

### 2b. Add a S3 data source

Any S3 compatible storage system can be used, including AWS S3, MinIO, etc. Below is an example for public dataset

```sh
singularity storage create s3 aws --name "my-source" --path "public-dataset-test"
```

## 3. Create a preparation
```sh
singularity prep create --source "my-source" --name "my-prep"
```

## 4. Run the preparation worker
```sh
singularity prep start-scan my-prep my-source
singularity run dataset-worker
```

## 5. Check the preparation status and result
```sh
singularity prep status my-prep
singularity prep list-pieces my-prep
```
