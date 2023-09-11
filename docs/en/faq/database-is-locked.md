# "Database is Locked" Error

When using Singularity with its default database backend (SQLite3), you might encounter the "Database is locked" error message.

## Why Does This Happen?

SQLite3 operates by using a file as its database. Every time a write operation is made, SQLite3 locks this file. If multiple write operations are attempted concurrently, the "Database is locked" error surfaces.

## What Should You Do?

- **Automatic Retry**: Singularity is designed to automatically retry operations that produce this error. Therefore, in many instances, you can safely ignore this error message.

- **Software Hang**: If you believe Singularity has become unresponsive due to this error, please report it as a bug.

## Production Recommendations

SQLite is suitable for development or light workloads, but it's not recommended for Production environments. For guidance on deploying Singularity in a Production environment with a more robust database backend, refer to the [Deploy to Production guide](../installation/deploy-to-production.md "Deploying Singularity to Production").

