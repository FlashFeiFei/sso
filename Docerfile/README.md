# sso单点登录的镜像

基于自己的创建的gin镜像

```
FROM 51785816/gindep:1.0


#将代码拷贝到工作目录空间，工作目录空间在DockerfileBase定义
COPY . .

#在工作区间安装程序
RUN go install -v

CMD ["sso"]

#执行 build -t 51785816/sso:1.0 .
```