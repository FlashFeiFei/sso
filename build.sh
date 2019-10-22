#!/bin/sh

#基于gin框架，构建sso镜像，并推送到自己的dockerhub仓库
sudo docker build --tag 51785816/sso:1.0 .

sudo docker push 51785816/sso:1.0