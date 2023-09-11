# Singularity: An Overview

Singularity offers an end-to-end solution designed to simplify the process of onboarding datasets to Filecoin storage providers. With capabilities spanning PiB-scale data, Singularity houses everything necessary for efficient onboarding.

What makes Singularity unique is its modularity. This ensures compatibility with various data preparation and deal-making tools and services.

To date, Singularity users have successfully onboarded more than 140 PiB of data to Filecoin storage providers, establishing it as the leading data preparation tool within the Filecoin community.

## Data Preparation

Equipped with a dedicated data preparation module, Singularity processes data from either a local file system or a remote storage service. By distributing tasks across multiple data preparation workers, Singularity achieves impressive horizontal scalability.

### Compatibility with 40+ Storage Solutions

Singularity stands out as the premier tool that connects with over 40 diverse remote storage services. This includes popular consumer products like Dropbox and Google Drive as well as enterprise-grade solutions such as AWS S3, Azure Blob Storage, FTP, HDFS, and more. This compatibility ensures a seamless integration with users' existing storage setups.

### Revolutionary Inline Preparation

Singularity introduces the concept of inline preparation, eliminating the need for additional disk space for storing CAR files. Instead, a metadata database is maintained, linking CAR files to their original data source. This allows storage providers to directly import content from these original sources.

### Preserving Dataset Integrity

Users will appreciate Singularity's ability to maintain dataset hierarchies. Folder structures and file versions remain untouched, facilitating easy dataset navigation and file retrieval via paths.

## Content Distribution

Beyond preparation, Singularity doubles as a nimble storage provider, eliminating traditional storage provider operations like sealing and proving. Users can effortlessly download either CAR files or original files from Singularity using a range of protocols, including Graphsync, HTTP, and Bitswap.

### Efficient CAR Distribution

With inline preparation in play, storage providers can swiftly download CAR files from Singularity. It optimally streams the CAR file directly from the original data source. This streamlined process can also integrate with Boost market deals and is multi-threading compatible.

### Convenient Content Retrieval

For all prepared files, Singularity offers Graphsync/HTTP/Bitswap retrievals. This feature positions Singularity as an alternative to a traditional storage provider, bypassing the need for storage proofs.

## Deal Making

For the current version, Singularity requires users to provide their own list of storage providers ready to accept deals. Users have the choice of two deal-making modes:

### Push Mode

In this mode, Singularity clients dictate the deal dispatch method, be it one-time batch processes or scheduled tasks. They can preset the number of deals per time interval, specify maximum deals or deal sizes, and define the number of permissible in-flight deals. Singularity also manages deal renewals in instances of deal slashes, expirations, or any on-chain disappearances.

### Pull Mode

This mode empowers storage providers. They can independently request deal proposals, as long as they align with the policies defined by Singularity clients and the storage providers are on the approved list.

### Wallet Management

Singularity also offers a robust wallet management feature. Clients can import their Filecoin wallets using private keys and link multiple wallets to a single dataset. Periodic datacap checks on each wallet ensure optimal load balancing between wallets.
