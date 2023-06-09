# 推送和上传

作为开发者，一旦你完成了一个项目或者更新了一个代码库，你需要将这些修改和更新推送到远程代码库上。在本文中，你将了解如何使用 Git 命令行完成推送和上传的操作。

## 推送变更

假设你已经完成了代码的修改，现在你需要把这些修改推送到远程仓库上。你可以使用 `git push` 命令来完成此操作。

```shell
git push <remote> <branch>
```

需要指定远程代码库和分支名。例如：

```shell
git push origin master
```

这条命令会将本地的 `master` 分支推送到名为 `origin` 的远程代码库中。

## 上传文件

有时候，你可能需要将本地的文件上传至远程服务器。你可以使用 `scp` 命令来完成此操作。

```shell
scp <source> <destination>
```

需要指定源文件和目标文件的路径。例如：

```shell
scp ~/Desktop/example.txt user@111.111.111.111:/home/user/
```

这条命令会将本地桌面上的 `example.txt` 文件上传至位于 `111.111.111.111` 服务器上的 `/home/user/` 目录下。

以上就是关于推送和上传的简短介绍。希望对你有所帮助！