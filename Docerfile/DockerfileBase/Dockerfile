FROM golang:1.13.3-buster

#创建和指定工作区间
RUN mkdir /go/src/my-application
RUN WORKDIR /go/src/my-application

# 安装dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
#安装gin框架
RUN dep ensure -add -v github.com/gin-gonic/gin

#安装mysql数据库驱动
RUN dep ensure -add -v github.com/go-sql-driver/mysql

#安装orm
RUN dep ensure -add -v github.com/jinzhu/gorm

#安装配文文件解析库
RUN dep ensure -add -v github.com/spf13/viper

#安装jwt
RUN dep ensure -add -v github.com/dgrijalva/jwt-go
#安装redis驱动，支持集群的
#RUN dep ensure -add -v github.com/go-redis/redis/v7


#因为golang的某些包需要翻墙，烦死了，百度云的服务器无法下载，所以我在本地通过ss翻墙，把gin框架安装，打包到镜像里面
#执行 docker build -t 51785816/gindep:1.0 .