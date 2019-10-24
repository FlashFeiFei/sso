FROM golang:1.13.3-buster

#创建和指定工作区间
RUN mkdir /go/src/my-application
RUN WORKDIR /go/src/my-application

# 安装dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

#安装gin框架、mysql、orm、配置文件解析、jwt
#dep ensure -add 分开执行的话，会把没有引用的包给删除掉，日!
RUN dep ensure -add -v github.com/gin-gonic/gin \
     github.com/go-sql-driver/mysql \
     github.com/jinzhu/gorm \
     github.com/spf13/viper \
      github.com/dgrijalva/jwt-go

#安装redis驱动，支持集群的,暂时不转，官网报错了
#RUN dep ensure -add -v github.com/go-redis/redis/v7


#因为golang的某些包需要翻墙，烦死了，百度云的服务器无法下载，所以我在本地通过ss翻墙，把gin框架安装，打包到镜像里面
#执行 docker build -t 51785816/gindep:1.0 .