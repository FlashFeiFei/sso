#!/bin/sh


#基于gin框架，构建sso镜像，并推送到自己的dockerhub仓库
sudo docker build --tag 51785816/sso:1.0 .


#登录到dockerhub
# sudo docker login

##输入账号密码



#推送到dockerhub
sudo docker push 51785816/sso:1.0