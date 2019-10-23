FROM 51785816/gindep:1.0


#将代码拷贝到工作目录空间，工作目录空间在DockerfileBase定义
COPY . .

#安装程序
RUN cd /go/src/github.com/guapo-organizations/sso; \
    go install -v

CMD ["sso"]

#执行 build -t 51785816/sso:1.0 .