---
description: Start by initializing the database and creating a new dataset
---

# Create a dataset

## Initialize the database

By default, it will be using the `sqlite3` database backend and initialize the database files in `$HOME/.singularity`

To use a different database backend for Production use, check [deploy-to-production.md](../installation/deploy-to-production.md "mention")

```sh
singularity init
```

## Create a new dataset

The dataset is a collection of data sources that relates to a single dataset. Once you have dataset created, you will be able to add data sources as well as associate Filecoin wallet addresses.

```sh
singularity dataset create my_dataset
```

By default, singularity uses a technical called Inline Preparation which will not export to any CAR files. That's because for most data source, it does not change and the CAR file is essentially storing the same content as the original data source.&#x20;

Next, you would want to [add-a-data-source.md](add-a-data-source.md "mention")
