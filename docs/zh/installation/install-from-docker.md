# 通过Docker安装Singularity

利用Docker，您可以轻松地拉取和运行预配置的Singularity镜像。

## 拉取Docker镜像

执行以下命令获取预构建的Docker镜像：
```bash
docker pull ghcr.io/data-preservation-programs/singularity:main
```

## 从Docker镜像运行Singularity
### 使用默认的SQLite3后端

默认情况下，Singularity使用`sqlite3`作为其数据库后端。要运行它，您应该将本地路径挂载到容器中的主目录：
```bash
docker run -v $HOME:/root ghcr.io/datapreservationprogram/singularity -h
```

### 使用其他数据库后端（如Postgres）

如果您选择另一个数据库后端（如Postgres），请在容器执行期间设置`DATABASE_CONNECTION_STRING`环境变量：
```bash
docker run -e DATABASE_CONNECTION_STRING=your_connection_string_here ghcr.io/datapreservationprogram/singularity -h
```