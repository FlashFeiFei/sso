# sso单点登录的镜像

基于自己的创建的gin镜像

```
FROM 51785816/gindep:1.0


#将代码拷贝到工作目录空间，工作目录空间在DockerfileBase定义
COPY ./controller .
COPY ./module .
COPY ./router .
COPY ./main .

#在工作区间安装程序
RUN go install -v

CMD ["my-application"]

#执行 docker build -t 51785816/sso:1.0 .
```

## 注意事项

不要再非DockerfileBase基础镜像中执行

这是 dep 包管理工具的规则

```
dep ensure -add 
```
操作

不然，代码没有依赖的包dep会自动删除，fuck!!!
