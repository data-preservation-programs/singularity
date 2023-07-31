# 数据库被锁定

这个错误在使用默认的数据库后端（SQlite3）时很常见。这是因为SQlite使用一个文件作为数据库，每次写操作都会锁定该文件。并发写操作会导致这个错误。系统会自动重试，你可以安全地忽略这个错误消息。

如果你怀疑软件因为这个消息而挂起，请报告一个错误。

不建议在生产环境中使用SQlite。请参考[deploy-to-production.md](../installation/deploy-to-production.md "mention")进行部署。