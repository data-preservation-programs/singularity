# What is Singularity

Singularity is an end to end tool to accelerate dataset onboarding to Filecoin storage providers. It contains all you need to efficiently onboard PiB scale of data.

Singularity is modular and includes a few different to work with other data preparation or deal making tools and service.

Users of Singularity has onboarded over 140 PiB of data to Filecoin storage providers and is current the top data preparation tool used by Filecoin clients.

## Data Preparation

Singularity ships with a data preparation module, which prepares the data either on local file system or remote storage service. Preparation tasks are distributed across data preparation workers achieving horizontal scalability.

### Integration with 40+ storage types

Singularity is the first tool that allows connection to 40+ different types of remote storage services, from consumer products such as Dropbox, Google Drive to enterprise cloud storage solutions such as AWS S3, Azure blob storage, FTP, HDFS, etc. This allows seamless integration with any existing storage solution the users have today.

### Inline Preparation

Singularity is the first tool that supports inline preparation, which eliminates need for extra disk spaces to store CAR files. Instead, Singularity maintains a metadata database that maps the CAR files to the original data source, so storage providers may import the content from original data source directly.

### Maintain Dataset hierarchy

Singularity maintains the folder structures and file versions of a dataset so that the users can explore the dataset folder by folder and retrieve files using paths.

### Encryption

Singularity supports a built-in encryption solution, which encrypts the files with provided recipients (public keys), or even a hardware PIV tokens such as YubiKeys. You may also supply a custom encryption script, which allows you to integrate with any external encryption solution \[Need testing].

## Content Distribution

Singularity itself can be used as a lightweight storage provider without involving any storage provider operations, i.e., sealing and proving. Users can download CAR files or original files from Singularity service using multiple protocols (Graphsync, HTTP and Bitswap)

### CAR distribution

Storage providers can download CAR files from Singularity server as Singularity streams the CAR file from original data source on the fly with minimal overhead if inline preparation is used. This download process can also be part of Boost market deals and supports multi-threading.

### Content Retrieval

Singularity can serve Graphsync/HTTP/Bitswap retrieval for all files that have already been prepared, so it can be used as a substitution of a storage provider without storage proofs.

## Deal Making

Singularity currently needs you to bring your own list of storage providers who are willing to accept your deals. There are two modes for deal making, a push mode where Singularity clients define how to send out deals, whether it is a one-off batch, or a scheduled cron. And a pull mode where Storage providers can ask for deal proposals themselves as long as it satisfies the policies set by Singularity clients.

### Push Mode

Singularity clients can set up a schedule with storage providers to send out a set number of deals for each interval time. They can also configure the maximum deals or deal size for each schedule as well as how many in-flight deals are allowed. Singularity handles deal renewal when deal is slashed, expires, or it disappears on-chain for any reason.

### Pull Mode

Singularity allows storage providers to query for list of PieceCIDs that they can propose deals to themselves as long as those storage providers are allow-listed in the deal making self-service policy. This way, storage providers have ultimate flexibility to control the deal flow within their sealing pipelines.

### Wallet Management

Singularity allows clients to import Filecoin wallet using private keys and associate multiple wallets with a dataset. It regularly checks datacap of each wallet to load balance between different wallets.

Singularity also supports remote signer, so the client can keep their wallet private key to their own.

###
