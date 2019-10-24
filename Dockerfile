FROM 51785816/gindep:1.0


#将代码拷贝到工作目录空间，工作目录空间在DockerfileBase定义
COPY ./controller .
COPY ./module .
COPY ./router .
COPY ./main.go .

#在工作区间安装程序
RUN go install -v

CMD ["my-application"]

#执行 docker build -t 51785816/sso:1.0 .