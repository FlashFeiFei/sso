FROM golang:1.13.3-buster

#安装dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/sso
COPY . /go/src/sso

RUN dep install -v

RUN go install -v ./...

CMD ["sso"]