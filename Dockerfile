FROM 51785816/goindep:1.0

#工作空间
WORKDIR /go/src/github.com/guapo-organizations/sso

#将代码拷贝下来
COPY . /go/src/github.com/guapo-organizations/sso

#dep包安装
RUN cd /go/src/github.com/guapo-organizations/sso && dep ensure -vendor-only

#安装程序
RUN cd /go/src/github.com/guapo-organizations/sso; \
    go install -v

CMD ["/go/bin/sso"]