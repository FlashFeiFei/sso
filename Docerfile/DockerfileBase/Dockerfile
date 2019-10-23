FROM golang:1.13.3-buster

#安装gin框架
RUN go get -u -v github.com/gin-gonic/gin

#安装mysql数据库驱动
RUN go get -u -v github.com/go-sql-driver/mysql

#安装orm
RUN go get -u -v github.com/jinzhu/gorm

#安装配文文件解析库
RUN go get -u -v github.com/spf13/viper

#安装jwt
RUN go get -v github.com/dgrijalva/jwt-go

#安装redis驱动，支持集群的

RUN go get -u -v github.com/go-redis/redis/v7

#因为golang的某些包需要翻墙，烦死了，百度云的服务器无法下载，所以我在本地通过ss翻墙，把gin框架安装，打包到镜像里面
#执行 docker build -t 51785816/goindep:1.0 .