# Database is locked

This error is normal especially when you are using the default database backend (SQlite3). This is because SQlite, which uses a file as the database has to locked the file for each write. Concurrent write will lead to this error. This error is automatically retried and you may safely ignore this error message.&#x20;

If you suspect the software has hanged due to this message, please report a bug.

SQlite is not recommended for Production usage. Checkout [deploy-to-production.md](../installation/deploy-to-production.md "mention")
