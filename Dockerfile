FROM golang:1.13.3-buster

#安装dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/guapo-organizations/sso
COPY . /go/src/github.com/guapo-organizations/sso

RUN cd /go/src/github.com/guapo-organizations/sso; \
   dep ensure -v

RUN cd /go/src/github.com/guapo-organizations/sso; \
    go install -v

CMD ["sso"]