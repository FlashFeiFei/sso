# DockerfileBase下的Dockerfile说明


该Dockerfile安装的环境

- golang
- dep管理工具
- gin框架
- mysql驱动
- orm
- jwt
- 配置文件解析库


# golang项目结构代码

```
|-- bin
  |-- 应用程序
|-- src
  |--my-application
    |-- main.go
    |-- vendor (dep的包管理)
    |-- Gopkg.lock
    |-- Gopkg.tol

```

# 为什么要构建这个DockerfileBase镜像？

因为我的是百度云服务器，golang的某些包下载不能翻墙，所以，本地翻墙下载下来，推送到自己的dockerhub仓库去


## dep包管理工具使用时遇到的问题

- dep ensure 的时候需要有go的代码,我日

所以Dockerfile的时候，copy main.go到工作区间

```
COPY ./main.go .

```


- dep ensure -add 不是批量式的写法，dep会自动删除没有应用的包fuck!!

所以Dockerfile这样写
```
RUN dep ensure -add -v github.com/gin-gonic/gin \
     github.com/go-sql-driver/mysql \
     github.com/jinzhu/gorm \
     github.com/spf13/viper \
      github.com/dgrijalva/jwt-go
```


# Dockerfile打tag推送到自己的dockerhub

```
#记得登录
docker login

docker build -t 51785816/gindep:1.0 .

docker push 51785816/gindep:1.0
```