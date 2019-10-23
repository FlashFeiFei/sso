# sso
单点登录。多个子系统，在同一个认证系统中登陆



# 部署流程

## 环境

- 百度云服务器，不能翻墙！
- docker-compose
- golang框架gin
- git

## 文件说明

### Docker镜像说明

文件夹
```
|-- Dockerfile目录
 |-- DockerfileBase目录
  |-- Dockerfile文件
|-- Dockerfile文件
```

### 构建镜像

构建基础镜像
```
#登录自己的dockerhub
docker login
#构建镜像
docker build -t 51785816/goindep:1.0 .
#推送到自己的dockerhub仓库
docker push 51785816/goindep:1.0
```

构建sso单点登录的镜像
```
#登录自己的dockerhub
docker login

#构建镜像
docker build -t 51785816/sso:1.0 .

#推送到自己的dockerhub仓库
docker push 51785816/sso:1.0
```

- DockerfileBase中的Dockerfile是基础镜像
- Dockerfile中的Dockerfile是sso的镜像

Dockerfile是构建sso单点登录应用的golang镜像

DockerfileBase是Dockerfile的基础镜像，里面安装了gin框架和一些库

为什么要有DockerfileBase？ gin框架 在执行go get -u github.com/gin-gonic/gin 的时候，有些包是需要翻墙的下载的，
本人的是国内百度云的服务器，fuck you！！！！不能翻墙，所以包下载不了。构建DockerfileBase的作用是，在可以翻墙的环境下
先翻墙下载下来，然后打包成镜像，镜像推送到dockerhub，这样百度云的服务器，直接下载镜像用就行。

### fuck那些需要翻墙才能下载的包

在DockerfileBase中，golang镜像已经安装了gin框架和mysql驱动、orm、配置文件解析、jwt
所以DockerfileBase基本不需要动，除非想要装些需要翻墙才能下载的包。

总结： 

- DockerfileBase下载了框架和一些工具库


## 部署

直接执行,还没有写完
```
sudo docker-compose up -d
```

