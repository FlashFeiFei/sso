FROM golang:1.13.3-buster

#安装dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/sso
COPY . /go/src/sso

RUN cd /go/src/sso; \
   dep ensure -v

RUN cd /go/src/sso; \
    go install -v

CMD ["sso"]