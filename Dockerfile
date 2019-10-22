FROM golang:1.13.3-buster

#安装dep

WORKDIR /go/src/sso
COPY . /go/src/sso


RUN cd /go/sso; \
    go install -v

CMD ["sso"]